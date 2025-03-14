// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package castai

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GkeCluster struct {
	pulumi.CustomResourceState

	// CAST.AI agent cluster token
	ClusterToken pulumi.StringOutput `pulumi:"clusterToken"`
	// CAST AI credentials id for cluster
	CredentialsId pulumi.StringOutput `pulumi:"credentialsId"`
	// GCP credentials.json from ServiceAccount with credentials for CAST AI
	CredentialsJson pulumi.StringPtrOutput `pulumi:"credentialsJson"`
	// Should CAST AI remove nodes managed by CAST.AI on disconnect
	DeleteNodesOnDisconnect pulumi.BoolPtrOutput `pulumi:"deleteNodesOnDisconnect"`
	GkeClusterId            pulumi.StringOutput  `pulumi:"gkeClusterId"`
	// GCP cluster zone in case of zonal or region in case of regional cluster
	Location pulumi.StringOutput `pulumi:"location"`
	// GKE cluster name
	Name pulumi.StringOutput `pulumi:"name"`
	// GCP project id
	ProjectId pulumi.StringOutput         `pulumi:"projectId"`
	Timeouts  GkeClusterTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewGkeCluster registers a new resource with the given unique name, arguments, and options.
func NewGkeCluster(ctx *pulumi.Context,
	name string, args *GkeClusterArgs, opts ...pulumi.ResourceOption) (*GkeCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.CredentialsJson != nil {
		args.CredentialsJson = pulumi.ToSecret(args.CredentialsJson).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clusterToken",
		"credentialsJson",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GkeCluster
	err = ctx.RegisterPackageResource("castai:index/gkeCluster:GkeCluster", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeCluster gets an existing GkeCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeClusterState, opts ...pulumi.ResourceOption) (*GkeCluster, error) {
	var resource GkeCluster
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("castai:index/gkeCluster:GkeCluster", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeCluster resources.
type gkeClusterState struct {
	// CAST.AI agent cluster token
	ClusterToken *string `pulumi:"clusterToken"`
	// CAST AI credentials id for cluster
	CredentialsId *string `pulumi:"credentialsId"`
	// GCP credentials.json from ServiceAccount with credentials for CAST AI
	CredentialsJson *string `pulumi:"credentialsJson"`
	// Should CAST AI remove nodes managed by CAST.AI on disconnect
	DeleteNodesOnDisconnect *bool   `pulumi:"deleteNodesOnDisconnect"`
	GkeClusterId            *string `pulumi:"gkeClusterId"`
	// GCP cluster zone in case of zonal or region in case of regional cluster
	Location *string `pulumi:"location"`
	// GKE cluster name
	Name *string `pulumi:"name"`
	// GCP project id
	ProjectId *string             `pulumi:"projectId"`
	Timeouts  *GkeClusterTimeouts `pulumi:"timeouts"`
}

type GkeClusterState struct {
	// CAST.AI agent cluster token
	ClusterToken pulumi.StringPtrInput
	// CAST AI credentials id for cluster
	CredentialsId pulumi.StringPtrInput
	// GCP credentials.json from ServiceAccount with credentials for CAST AI
	CredentialsJson pulumi.StringPtrInput
	// Should CAST AI remove nodes managed by CAST.AI on disconnect
	DeleteNodesOnDisconnect pulumi.BoolPtrInput
	GkeClusterId            pulumi.StringPtrInput
	// GCP cluster zone in case of zonal or region in case of regional cluster
	Location pulumi.StringPtrInput
	// GKE cluster name
	Name pulumi.StringPtrInput
	// GCP project id
	ProjectId pulumi.StringPtrInput
	Timeouts  GkeClusterTimeoutsPtrInput
}

func (GkeClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeClusterState)(nil)).Elem()
}

type gkeClusterArgs struct {
	// GCP credentials.json from ServiceAccount with credentials for CAST AI
	CredentialsJson *string `pulumi:"credentialsJson"`
	// Should CAST AI remove nodes managed by CAST.AI on disconnect
	DeleteNodesOnDisconnect *bool   `pulumi:"deleteNodesOnDisconnect"`
	GkeClusterId            *string `pulumi:"gkeClusterId"`
	// GCP cluster zone in case of zonal or region in case of regional cluster
	Location string `pulumi:"location"`
	// GKE cluster name
	Name *string `pulumi:"name"`
	// GCP project id
	ProjectId string              `pulumi:"projectId"`
	Timeouts  *GkeClusterTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a GkeCluster resource.
type GkeClusterArgs struct {
	// GCP credentials.json from ServiceAccount with credentials for CAST AI
	CredentialsJson pulumi.StringPtrInput
	// Should CAST AI remove nodes managed by CAST.AI on disconnect
	DeleteNodesOnDisconnect pulumi.BoolPtrInput
	GkeClusterId            pulumi.StringPtrInput
	// GCP cluster zone in case of zonal or region in case of regional cluster
	Location pulumi.StringInput
	// GKE cluster name
	Name pulumi.StringPtrInput
	// GCP project id
	ProjectId pulumi.StringInput
	Timeouts  GkeClusterTimeoutsPtrInput
}

func (GkeClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeClusterArgs)(nil)).Elem()
}

type GkeClusterInput interface {
	pulumi.Input

	ToGkeClusterOutput() GkeClusterOutput
	ToGkeClusterOutputWithContext(ctx context.Context) GkeClusterOutput
}

func (*GkeCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeCluster)(nil)).Elem()
}

func (i *GkeCluster) ToGkeClusterOutput() GkeClusterOutput {
	return i.ToGkeClusterOutputWithContext(context.Background())
}

func (i *GkeCluster) ToGkeClusterOutputWithContext(ctx context.Context) GkeClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeClusterOutput)
}

type GkeClusterOutput struct{ *pulumi.OutputState }

func (GkeClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeCluster)(nil)).Elem()
}

func (o GkeClusterOutput) ToGkeClusterOutput() GkeClusterOutput {
	return o
}

func (o GkeClusterOutput) ToGkeClusterOutputWithContext(ctx context.Context) GkeClusterOutput {
	return o
}

// CAST.AI agent cluster token
func (o GkeClusterOutput) ClusterToken() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.StringOutput { return v.ClusterToken }).(pulumi.StringOutput)
}

// CAST AI credentials id for cluster
func (o GkeClusterOutput) CredentialsId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.StringOutput { return v.CredentialsId }).(pulumi.StringOutput)
}

// GCP credentials.json from ServiceAccount with credentials for CAST AI
func (o GkeClusterOutput) CredentialsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.StringPtrOutput { return v.CredentialsJson }).(pulumi.StringPtrOutput)
}

// Should CAST AI remove nodes managed by CAST.AI on disconnect
func (o GkeClusterOutput) DeleteNodesOnDisconnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.BoolPtrOutput { return v.DeleteNodesOnDisconnect }).(pulumi.BoolPtrOutput)
}

func (o GkeClusterOutput) GkeClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.StringOutput { return v.GkeClusterId }).(pulumi.StringOutput)
}

// GCP cluster zone in case of zonal or region in case of regional cluster
func (o GkeClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// GKE cluster name
func (o GkeClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// GCP project id
func (o GkeClusterOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeCluster) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GkeClusterOutput) Timeouts() GkeClusterTimeoutsPtrOutput {
	return o.ApplyT(func(v *GkeCluster) GkeClusterTimeoutsPtrOutput { return v.Timeouts }).(GkeClusterTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeClusterInput)(nil)).Elem(), &GkeCluster{})
	pulumi.RegisterOutputType(GkeClusterOutput{})
}
