package castai

import (
	"fmt"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GkeClusterArgs struct {
	ApiUrl                       pulumi.StringInput
	CastaiApiToken               pulumi.StringInput
	GrpcUrl                      pulumi.StringInput
	ApiGrpcAddr                  pulumi.StringInput
	KvisorControllerExtraArgs    map[string]pulumi.StringInput
	ProjectId                    pulumi.String
	GkeClusterName               pulumi.String
	AutoscalerPoliciesJson       pulumi.StringInput
	DeleteNodesOnDisconnect      pulumi.BoolInput
	GkeClusterLocation           pulumi.String
	Subnets                      []string
	GkeCredentials               pulumi.StringInput
	CastaiComponentsLabels       map[string]interface{}
	EnableAutoScaler             pulumi.Bool
	EnableUnschedulablePods      pulumi.Bool
	AggressiveModeEnabled        pulumi.Bool
	EnableEvictor                pulumi.Bool
	EnableDownScaler             pulumi.Bool
	NodeConfigurations           pulumi.Map
	DefaultNodeConfiguration     pulumi.StringInput
	DefaultNodeConfigurationName pulumi.StringInput
	NodeTemplates                pulumi.Map
	WorkloadScalingPolicies      pulumi.Map
	InstallSecurityAgent         pulumi.BoolInput
	AgentVersion                 pulumi.StringInput
	ClusterControllerVersion     pulumi.StringInput
	EvictorVersion               pulumi.StringInput
	EvictorExtVersion            pulumi.StringInput
	PodPinnerVersion             pulumi.StringInput
	SpotHandlerVersion           pulumi.StringInput
	KvisorVersion                pulumi.StringInput
	AgentValues                  []pulumi.StringInput
	SpotHandlerValues            []pulumi.StringInput
	ClusterControllerValues      []pulumi.StringInput
	EvictorValues                []pulumi.StringInput
	EvictorExtValues             []pulumi.StringInput
	PodPinnerValues              []pulumi.StringInput
	KvisorValues                 []pulumi.StringInput
	SelfManaged                  pulumi.BoolInput
	WaitForClusterReady          pulumi.BoolInput
	InstallWorkloadAutoscaler    pulumi.BoolInput
	WorkloadAutoscalerVersion    pulumi.StringInput
	WorkloadAutoscalerValues     []pulumi.StringInput
	InstallCloudProxy            pulumi.BoolInput
	CloudProxyVersion            pulumi.StringInput
	CloudProxyValues             []pulumi.StringInput
	CloudProxyGrpcUrlOverride    pulumi.StringInput
}

type GkeClusterResult struct {
	pulumi.ResourceState
	ClusterId                pulumi.AnyOutput
	CastaiNodeConfigurations pulumi.AnyOutput
	CastaiNodeTemplates      pulumi.AnyOutput
}

// NewGkeCluster component connects a GKE cluster to CASTAI
// and deploys the in-cluster components
func NewGkeCluster(ctx *pulumi.Context, name string, args *GkeClusterArgs, opts ...pulumi.ResourceOption) (*GkeClusterResult, error) {
	var componentResource GkeClusterResult
	err := ctx.RegisterComponentResource("components:index:gkecluster", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}

	// Joins a cluster into CASTAI
	castaiCluster, err := castai.NewGkeCluster(ctx, fmt.Sprintf("%s:castai:cluster", name), &castai.GkeClusterArgs{
		ProjectId:               pulumi.String(args.ProjectId),
		Location:                pulumi.String(args.GkeClusterLocation),
		Name:                    pulumi.String(args.GkeClusterName),
		DeleteNodesOnDisconnect: args.DeleteNodesOnDisconnect,
		CredentialsJson:         args.GkeCredentials,
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// Create the castai-agent namespace
	namespace, err := corev1.NewNamespace(ctx, fmt.Sprintf("%s:castai:namespace", name), &corev1.NamespaceArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name: pulumi.String("castai-agent"),
		},
	})
	if err != nil {
		return nil, err
	}

	// Helm releases of in-cluster components
	agentValues := pulumi.Map{
		"provider": pulumi.String("gke"),
		"additionalEnv": pulumi.Map{
			"STATIC_CLUSTER_ID": castaiCluster.GkeClusterId,
		},
		"apiKey":          pulumi.StringInput(args.CastaiApiToken),
		"createNamespace": pulumi.Bool(false),
	}

	castAgent, err := helm.NewRelease(ctx, fmt.Sprintf("%s:castai:helm:agent", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-agent"),
		Chart:           pulumi.String("castai-agent"),
		Namespace:       namespace.Metadata.Name(),
		CreateNamespace: pulumi.Bool(true),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.AgentVersion,
		Values:          agentValues,
		Timeout:         pulumi.Int(120),
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	controllerValues := pulumi.Map{
		"castai": pulumi.Map{
			"apiKey":    pulumi.StringInput(castaiCluster.ClusterToken),
			"clusterID": pulumi.StringInput(castaiCluster.GkeClusterId),
		},
	}

	clusterController, err := helm.NewRelease(ctx, fmt.Sprintf("%s:castai:helm:cluster_controller", name), &helm.ReleaseArgs{
		Name:            pulumi.String("cluster-controller"),
		Chart:           pulumi.String("castai-cluster-controller"),
		Namespace:       namespace.Metadata.Name(),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.ClusterControllerVersion,
		Values:          controllerValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castAgent,
	}))
	if err != nil {
		return nil, err
	}

	evictorValues := pulumi.Map{}
	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s:castai:helm:castai_evictor", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-evictor"),
		Chart:           pulumi.String("castai-evictor"),
		Namespace:       namespace.Metadata.Name(),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.EvictorVersion,
		Values:          evictorValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castAgent, clusterController,
	}))
	if err != nil {
		return nil, err
	}

	spotHandlerValues := pulumi.Map{
		"castai": pulumi.Map{
			"provider":  pulumi.String("gcp"),
			"clusterID": pulumi.StringInput(castaiCluster.GkeClusterId),
		},
	}

	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s:castai:helm:castai_spot_handler", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-spot-handler"),
		Chart:           pulumi.String("castai-spot-handler"),
		Namespace:       namespace.Metadata.Name(),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.SpotHandlerVersion,
		Values:          spotHandlerValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castAgent, clusterController,
	}))
	if err != nil {
		return nil, err
	}

	kvisorValues := pulumi.Map{
		"castai": pulumi.Map{
			"clusterID": castaiCluster.GkeClusterId,
			"apiKey":    pulumi.StringInput(args.CastaiApiToken),
		},
		"controller": pulumi.Map{
			"extraArgs": pulumi.Map{
				"kube-bench-cloud-provider": pulumi.String("gke"),
			},
		},
	}
	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s:castai:helm:castai_kvisor", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-kvisor"),
		Chart:           pulumi.String("castai-kvisor"),
		Namespace:       namespace.Metadata.Name(),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.KvisorVersion,
		Values:          kvisorValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castAgent, clusterController,
	}))
	if err != nil {
		return nil, err
	}

	// Node configuration
	var networkTags pulumi.StringArray

	nodeConfig := &castai.NodeConfigurationGkeArgs{
		MaxPodsPerNode:              pulumi.Float64(100),
		NetworkTags:                 networkTags,
		DiskType:                    pulumi.String("pd-standard"),
		UseEphemeralStorageLocalSsd: pulumi.Bool(false),
	}

	newNodeConfigurationRes, err := castai.NewNodeConfiguration(ctx, fmt.Sprintf("%s:castai:cluster:nodeconfiguration", name), &castai.NodeConfigurationArgs{
		Gke:             nodeConfig,
		ClusterId:       castaiCluster.GkeClusterId,
		Name:            pulumi.String("default"),
		DiskCpuRatio:    pulumi.Float64(0),
		DrainTimeoutSec: pulumi.Float64(100),
		MinDiskSize:     pulumi.Float64(100),
		Subnets:         pulumi.ToStringArray(args.Subnets),
		Tags: pulumi.StringMap{
			"name": pulumi.String("test"),
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// Node template
	nodeTemplate, err := castai.NewNodeTemplate(ctx, fmt.Sprintf("%s:castai:cluster:nodetemplate", name), &castai.NodeTemplateArgs{
		CustomTaints: &castai.NodeTemplateCustomTaintArray{},
		CustomLabels: pulumi.StringMap{
			"template": pulumi.String("example_nodetemplate"),
		},
		Constraints:                              &castai.NodeTemplateConstraintsArgs{},
		ClusterId:                                castaiCluster.GkeClusterId,
		Name:                                     pulumi.String("example_nodetemplate"),
		ConfigurationId:                          newNodeConfigurationRes.ID(),
		IsDefault:                                pulumi.Bool(false),
		IsEnabled:                                pulumi.Bool(true),
		ShouldTaint:                              pulumi.Bool(false),
		CustomInstancesEnabled:                   pulumi.Bool(false),
		CustomInstancesWithExtendedMemoryEnabled: pulumi.Bool(false),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		newNodeConfigurationRes,
	}))
	if err != nil {
		return nil, err
	}

	// Autoscaler Settings
	_, err = castai.NewAutoscaler(ctx, fmt.Sprintf("%s:castai:cluster:autoscaler_policies", name), &castai.AutoscalerArgs{
		AutoscalerSettings: &castai.AutoscalerAutoscalerSettingsArgs{
			Enabled: args.EnableAutoScaler,
			UnschedulablePods: &castai.AutoscalerAutoscalerSettingsUnschedulablePodsArgs{
				Enabled: args.EnableUnschedulablePods,
			},
			NodeDownscaler: &castai.AutoscalerAutoscalerSettingsNodeDownscalerArgs{
				Enabled: args.EnableDownScaler,
				Evictor: &castai.AutoscalerAutoscalerSettingsNodeDownscalerEvictorArgs{
					Enabled:        args.EnableEvictor,
					AggressiveMode: args.AggressiveModeEnabled,
				},
			},
		},
		ClusterId:              castaiCluster.GkeClusterId,
		AutoscalerPoliciesJson: args.AutoscalerPoliciesJson,
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		newNodeConfigurationRes,
	}))
	if err != nil {
		return nil, err
	}

	err = ctx.RegisterResourceOutputs(&componentResource, pulumi.Map{
		"clusterId":                castaiCluster.GkeClusterId,
		"castaiNodeConfigurations": newNodeConfigurationRes,
		"castaiNodeTemplates":      nodeTemplate,
	})
	if err != nil {
		return nil, err
	}
	componentResource.ClusterId = pulumi.AnyOutput(castaiCluster.GkeClusterId)
	componentResource.CastaiNodeConfigurations = pulumi.AnyOutput(newNodeConfigurationRes.NodeConfigurationId)
	componentResource.CastaiNodeTemplates = pulumi.AnyOutput(nodeTemplate.NodeTemplateId)
	return &componentResource, nil
}
