package gke

import (
	"fmt"

	"crypto/sha1"
	"encoding/hex"

	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-std/sdk/go/std"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GkeIamArgs struct {
	ProjectId                         pulumi.String
	GkeClusterName                    pulumi.String
	ServiceAccountsUniqueIds          pulumi.String
	Location                          pulumi.String
	ComputeManagerProjectIds          []pulumi.StringInput
	CreateServiceAccount              pulumi.Bool
	SetupCloudProxyWorkloadIdentity   pulumi.BoolInput
	WorkloadIdentityNamespace         pulumi.StringInput
	CloudProxyServiceAccountNamespace pulumi.StringInput
	CloudProxyServiceAccountName      pulumi.StringInput
	CastaiRolePermissions             []pulumi.StringInput
	ComputeManagerPermissions         []pulumi.StringInput
}

type GkeIamResult struct {
	pulumi.ResourceState
	PrivateKey                       pulumi.StringOutput
	ServiceAccountId                 pulumi.AnyOutput
	ServiceAccountEmail              pulumi.AnyOutput
	DefaultComputeManagerPermissions pulumi.AnyOutput
	DefaultCastaiRolePermissions     pulumi.AnyOutput
}

var DefaultCastaiRolePermissions = []string{
	"container.clusters.get",
	"container.clusters.update",
	"container.certificateSigningRequests.approve",
	"compute.instances.get",
	"compute.instances.list",
	"compute.instances.create",
	"compute.instances.start",
	"compute.instances.stop",
	"compute.instances.delete",
	"compute.instances.setLabels",
	"compute.instances.setServiceAccount",
	"compute.instances.setMetadata",
	"compute.instances.setTags",
	"compute.instanceGroupManagers.get",
	"compute.instanceGroupManagers.update",
	"compute.instanceGroups.get",
	"compute.networks.use",
	"compute.networks.useExternalIp",
	"compute.subnetworks.get",
	"compute.subnetworks.use",
	"compute.subnetworks.useExternalIp",
	"compute.addresses.use",
	"compute.disks.use",
	"compute.disks.create",
	"compute.disks.setLabels",
	"compute.images.get",
	"compute.images.useReadOnly",
	"compute.instanceTemplates.get",
	"compute.instanceTemplates.list",
	"compute.instanceTemplates.create",
	"compute.instanceTemplates.delete",
	"compute.regionOperations.get",
	"compute.zoneOperations.get",
	"compute.zones.list",
	"compute.zones.get",
	"serviceusage.services.list",
	"resourcemanager.projects.getIamPolicy",
	"compute.targetPools.get",
	"compute.targetPools.addInstance",
	"compute.targetPools.removeInstance",
	"compute.instances.use",
}

// NewGkeIam component creates the necessary permission
// to your GCP project to manage nodes in GKE
func NewGkeIam(ctx *pulumi.Context, name string, args *GkeIamArgs, opts ...pulumi.ResourceOption) (*GkeIamResult, error) {
	var componentResource GkeIamResult
	err := ctx.RegisterComponentResource("components:index:gkeiam", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}

	// Custom IAM roles prep
	// Create a has of the cluster name
	clusterNameHash := sha1.New()
	clusterNameHash.Write([]byte(args.GkeClusterName))
	clusterHashString := hex.EncodeToString(clusterNameHash.Sum(nil))

	// Set Custom Role ID
	castaiCustomClusterRoleId := pulumi.String(fmt.Sprintf("cluster.castai.gkeAccess.%v.tf", clusterHashString[:8]))

	var castaiRolePermissions pulumi.StringArray
	if float64(len(args.CastaiRolePermissions)) > 0 {
		castaiRolePermissions = args.CastaiRolePermissions
	} else {
		castaiRolePermissions = pulumi.ToStringArray(DefaultCastaiRolePermissions)
	}

	// CREATING THE CLUSTER IAM ROLES
	castaiClusterCustomRoleResult, err := projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s:gkeiam:custom_role", name), &projects.IAMCustomRoleArgs{
		RoleId:      pulumi.String(castaiCustomClusterRoleId),
		Title:       pulumi.String("Role to manage GKE cluster via CAST AI"),
		Description: pulumi.String("Role to manage GKE cluster via CAST AI"),
		Permissions: castaiRolePermissions,
		Project:     pulumi.String(args.ProjectId),
		Stage:       pulumi.String("GA"),
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// ÇREATE SERVICE ACCOUNT
	serviceAccountHash := sha1.New()
	serviceAccountHash.Write([]byte(args.GkeClusterName))
	serviceAccountHashString := hex.EncodeToString(serviceAccountHash.Sum(nil))

	serviceAccountId := pulumi.String(fmt.Sprintf("castai-gke-%v", serviceAccountHashString[:8]))
	serviceAccountEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", serviceAccountId, args.ProjectId)

	serviceAccountResult, err := serviceaccount.NewAccount(ctx, fmt.Sprintf("%s:gkeiam:service_account", name), &serviceaccount.AccountArgs{
		AccountId:   pulumi.String(serviceAccountId),
		DisplayName: pulumi.Sprintf("Service account to manage %v cluster via CAST", args.GkeClusterName),
		Project:     pulumi.String(args.ProjectId),
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	serviceAccountKeyResult, err := serviceaccount.NewKey(ctx, fmt.Sprintf("%s:gkeiam:service_account_key", name), &serviceaccount.KeyArgs{
		ServiceAccountId: serviceAccountResult.Name,
		PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		serviceAccountResult,
	}))
	if err != nil {
		return nil, err
	}

	_, err = projects.NewIAMMember(ctx, fmt.Sprintf("%s:gkeiam:member_project", name), &projects.IAMMemberArgs{
		Project: pulumi.String(args.ProjectId),
		Role:    pulumi.Sprintf("projects/%s/roles/%s", args.ProjectId, castaiCustomClusterRoleId),
		Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiClusterCustomRoleResult, serviceAccountResult,
	}))
	if err != nil {
		return nil, err
	}

	_, err = projects.NewIAMMember(ctx, fmt.Sprintf("%s:gkeiam:member_service_account_user", name), &projects.IAMMemberArgs{
		Project: pulumi.String(args.ProjectId),
		Role:    pulumi.String("roles/iam.serviceAccountUser"),
		Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		serviceAccountResult,
	}))
	if err != nil {
		return nil, err
	}

	// Base64decode the privatekey of Serviceaccount created
	decoded := std.Base64decodeOutput(ctx, std.Base64decodeOutputArgs{
		Input: serviceAccountKeyResult.PrivateKey,
	}, nil).ApplyT(func(invoke std.Base64decodeResult) (string, error) {
		return invoke.Result, nil
	}).(pulumi.StringOutput)

	componentResource.PrivateKey = decoded

	err = ctx.RegisterResourceOutputs(&componentResource, pulumi.Map{
		"privateKey":          decoded,
		"serviceAccountId":    serviceAccountResult.AccountId,
		"serviceAccountEmail": serviceAccountResult.Email,
	})
	if err != nil {
		return nil, err
	}

	return &componentResource, nil
}
