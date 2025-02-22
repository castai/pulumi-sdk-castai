package main

import (
	"github.com/castai/pulumi-convert/castai"
	"github.com/castai/pulumi-convert/gke"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Set configs
		apiToken := config.Get(ctx, "castai:apiToken")

		projectId := "<YOUR_PROJECT>"
		clusterName := "<EXISTING_GKE_CLUSTER_NAME"
		gcpLocation := "<GKE_LOCATION>"
		subnets := []string{
			"<SUBNET>",
		}
		enableAutoScaler := true
		enableUnschedulablePods := true
		aggressiveModeEnabled := false
		enableEvictor := true
		enableDownScaler := true

		gkeIamResult, err := gke.NewGkeIam(ctx, "castai:gke:iam", &gke.GkeIamArgs{
			ProjectId:      pulumi.String(projectId),
			GkeClusterName: pulumi.String(clusterName),
			Location:       pulumi.String(gcpLocation),
		})
		if err != nil {
			return err
		}

		_, err = castai.NewGkeCluster(ctx, "castai:gke:cluster", &castai.GkeClusterArgs{
			ApiUrl:                    pulumi.String("https://api.cast.ai"),
			CastaiApiToken:            pulumi.String(apiToken),
			GrpcUrl:                   pulumi.String("grpc.cast.ai:443"),
			WaitForClusterReady:       pulumi.Bool(false),
			ProjectId:                 pulumi.String(projectId),
			GkeClusterName:            pulumi.String(clusterName),
			GkeClusterLocation:        pulumi.String(gcpLocation),
			GkeCredentials:            gkeIamResult.PrivateKey,
			DeleteNodesOnDisconnect:   pulumi.Bool(true),
			InstallWorkloadAutoscaler: pulumi.Bool(true),
			AgentVersion:              pulumi.String("0.97.0"),
			ClusterControllerVersion:  pulumi.String("0.80.1"),
			EvictorVersion:            pulumi.String("0.31.42"),
			SpotHandlerVersion:        pulumi.String("0.26.0"),
			KvisorVersion:             pulumi.String("1.0.71"),
			Subnets:                   subnets,
			EnableAutoScaler:          pulumi.Bool(enableAutoScaler),
			EnableUnschedulablePods:   pulumi.Bool(enableUnschedulablePods),
			AggressiveModeEnabled:     pulumi.Bool(aggressiveModeEnabled),
			EnableEvictor:             pulumi.Bool(enableEvictor),
			EnableDownScaler:          pulumi.Bool(enableDownScaler),
			NodeConfigurations:        pulumi.Map{},
		}, pulumi.DependsOn([]pulumi.Resource{
			gkeIamResult,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
