// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerRegistry struct {
	pulumi.CustomResourceState

	// The URI of the created resource.
	BucketSelfLink      pulumi.StringOutput `pulumi:"bucketSelfLink"`
	ContainerRegistryId pulumi.StringOutput `pulumi:"containerRegistryId"`
	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information
	// on registry locations.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistry(ctx *pulumi.Context,
	name string, args *ContainerRegistryArgs, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	if args == nil {
		args = &ContainerRegistryArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ContainerRegistry
	err = ctx.RegisterPackageResource("google-beta:index/containerRegistry:ContainerRegistry", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistry gets an existing ContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryState, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	var resource ContainerRegistry
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/containerRegistry:ContainerRegistry", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistry resources.
type containerRegistryState struct {
	// The URI of the created resource.
	BucketSelfLink      *string `pulumi:"bucketSelfLink"`
	ContainerRegistryId *string `pulumi:"containerRegistryId"`
	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information
	// on registry locations.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ContainerRegistryState struct {
	// The URI of the created resource.
	BucketSelfLink      pulumi.StringPtrInput
	ContainerRegistryId pulumi.StringPtrInput
	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information
	// on registry locations.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryState)(nil)).Elem()
}

type containerRegistryArgs struct {
	ContainerRegistryId *string `pulumi:"containerRegistryId"`
	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information
	// on registry locations.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ContainerRegistry resource.
type ContainerRegistryArgs struct {
	ContainerRegistryId pulumi.StringPtrInput
	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information
	// on registry locations.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryArgs)(nil)).Elem()
}

type ContainerRegistryInput interface {
	pulumi.Input

	ToContainerRegistryOutput() ContainerRegistryOutput
	ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput
}

func (*ContainerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (i *ContainerRegistry) ToContainerRegistryOutput() ContainerRegistryOutput {
	return i.ToContainerRegistryOutputWithContext(context.Background())
}

func (i *ContainerRegistry) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOutput)
}

type ContainerRegistryOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryOutput) ToContainerRegistryOutput() ContainerRegistryOutput {
	return o
}

func (o ContainerRegistryOutput) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return o
}

// The URI of the created resource.
func (o ContainerRegistryOutput) BucketSelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.BucketSelfLink }).(pulumi.StringOutput)
}

func (o ContainerRegistryOutput) ContainerRegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.ContainerRegistryId }).(pulumi.StringOutput)
}

// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information
// on registry locations.
func (o ContainerRegistryOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
func (o ContainerRegistryOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryInput)(nil)).Elem(), &ContainerRegistry{})
	pulumi.RegisterOutputType(ContainerRegistryOutput{})
}
