// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RuntimeconfigConfig struct {
	pulumi.CustomResourceState

	// The description to associate with the runtime config.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the runtime config.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project               pulumi.StringOutput `pulumi:"project"`
	RuntimeconfigConfigId pulumi.StringOutput `pulumi:"runtimeconfigConfigId"`
}

// NewRuntimeconfigConfig registers a new resource with the given unique name, arguments, and options.
func NewRuntimeconfigConfig(ctx *pulumi.Context,
	name string, args *RuntimeconfigConfigArgs, opts ...pulumi.ResourceOption) (*RuntimeconfigConfig, error) {
	if args == nil {
		args = &RuntimeconfigConfigArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource RuntimeconfigConfig
	err = ctx.RegisterPackageResource("google-beta:index/runtimeconfigConfig:RuntimeconfigConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeconfigConfig gets an existing RuntimeconfigConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeconfigConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeconfigConfigState, opts ...pulumi.ResourceOption) (*RuntimeconfigConfig, error) {
	var resource RuntimeconfigConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/runtimeconfigConfig:RuntimeconfigConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeconfigConfig resources.
type runtimeconfigConfigState struct {
	// The description to associate with the runtime config.
	Description *string `pulumi:"description"`
	// The name of the runtime config.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project               *string `pulumi:"project"`
	RuntimeconfigConfigId *string `pulumi:"runtimeconfigConfigId"`
}

type RuntimeconfigConfigState struct {
	// The description to associate with the runtime config.
	Description pulumi.StringPtrInput
	// The name of the runtime config.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project               pulumi.StringPtrInput
	RuntimeconfigConfigId pulumi.StringPtrInput
}

func (RuntimeconfigConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeconfigConfigState)(nil)).Elem()
}

type runtimeconfigConfigArgs struct {
	// The description to associate with the runtime config.
	Description *string `pulumi:"description"`
	// The name of the runtime config.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project               *string `pulumi:"project"`
	RuntimeconfigConfigId *string `pulumi:"runtimeconfigConfigId"`
}

// The set of arguments for constructing a RuntimeconfigConfig resource.
type RuntimeconfigConfigArgs struct {
	// The description to associate with the runtime config.
	Description pulumi.StringPtrInput
	// The name of the runtime config.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project               pulumi.StringPtrInput
	RuntimeconfigConfigId pulumi.StringPtrInput
}

func (RuntimeconfigConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeconfigConfigArgs)(nil)).Elem()
}

type RuntimeconfigConfigInput interface {
	pulumi.Input

	ToRuntimeconfigConfigOutput() RuntimeconfigConfigOutput
	ToRuntimeconfigConfigOutputWithContext(ctx context.Context) RuntimeconfigConfigOutput
}

func (*RuntimeconfigConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeconfigConfig)(nil)).Elem()
}

func (i *RuntimeconfigConfig) ToRuntimeconfigConfigOutput() RuntimeconfigConfigOutput {
	return i.ToRuntimeconfigConfigOutputWithContext(context.Background())
}

func (i *RuntimeconfigConfig) ToRuntimeconfigConfigOutputWithContext(ctx context.Context) RuntimeconfigConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeconfigConfigOutput)
}

type RuntimeconfigConfigOutput struct{ *pulumi.OutputState }

func (RuntimeconfigConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeconfigConfig)(nil)).Elem()
}

func (o RuntimeconfigConfigOutput) ToRuntimeconfigConfigOutput() RuntimeconfigConfigOutput {
	return o
}

func (o RuntimeconfigConfigOutput) ToRuntimeconfigConfigOutputWithContext(ctx context.Context) RuntimeconfigConfigOutput {
	return o
}

// The description to associate with the runtime config.
func (o RuntimeconfigConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeconfigConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the runtime config.
func (o RuntimeconfigConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
func (o RuntimeconfigConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RuntimeconfigConfigOutput) RuntimeconfigConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfig) pulumi.StringOutput { return v.RuntimeconfigConfigId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeconfigConfigInput)(nil)).Elem(), &RuntimeconfigConfig{})
	pulumi.RegisterOutputType(RuntimeconfigConfigOutput{})
}
