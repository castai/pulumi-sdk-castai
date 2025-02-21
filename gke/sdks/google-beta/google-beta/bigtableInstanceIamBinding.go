// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BigtableInstanceIamBinding struct {
	pulumi.CustomResourceState

	BigtableInstanceIamBindingId pulumi.StringOutput                          `pulumi:"bigtableInstanceIamBindingId"`
	Condition                    BigtableInstanceIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                         pulumi.StringOutput                          `pulumi:"etag"`
	Instance                     pulumi.StringOutput                          `pulumi:"instance"`
	Members                      pulumi.StringArrayOutput                     `pulumi:"members"`
	Project                      pulumi.StringOutput                          `pulumi:"project"`
	Role                         pulumi.StringOutput                          `pulumi:"role"`
}

// NewBigtableInstanceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewBigtableInstanceIamBinding(ctx *pulumi.Context,
	name string, args *BigtableInstanceIamBindingArgs, opts ...pulumi.ResourceOption) (*BigtableInstanceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigtableInstanceIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/bigtableInstanceIamBinding:BigtableInstanceIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigtableInstanceIamBinding gets an existing BigtableInstanceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigtableInstanceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigtableInstanceIamBindingState, opts ...pulumi.ResourceOption) (*BigtableInstanceIamBinding, error) {
	var resource BigtableInstanceIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/bigtableInstanceIamBinding:BigtableInstanceIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigtableInstanceIamBinding resources.
type bigtableInstanceIamBindingState struct {
	BigtableInstanceIamBindingId *string                              `pulumi:"bigtableInstanceIamBindingId"`
	Condition                    *BigtableInstanceIamBindingCondition `pulumi:"condition"`
	Etag                         *string                              `pulumi:"etag"`
	Instance                     *string                              `pulumi:"instance"`
	Members                      []string                             `pulumi:"members"`
	Project                      *string                              `pulumi:"project"`
	Role                         *string                              `pulumi:"role"`
}

type BigtableInstanceIamBindingState struct {
	BigtableInstanceIamBindingId pulumi.StringPtrInput
	Condition                    BigtableInstanceIamBindingConditionPtrInput
	Etag                         pulumi.StringPtrInput
	Instance                     pulumi.StringPtrInput
	Members                      pulumi.StringArrayInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringPtrInput
}

func (BigtableInstanceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigtableInstanceIamBindingState)(nil)).Elem()
}

type bigtableInstanceIamBindingArgs struct {
	BigtableInstanceIamBindingId *string                              `pulumi:"bigtableInstanceIamBindingId"`
	Condition                    *BigtableInstanceIamBindingCondition `pulumi:"condition"`
	Instance                     string                               `pulumi:"instance"`
	Members                      []string                             `pulumi:"members"`
	Project                      *string                              `pulumi:"project"`
	Role                         string                               `pulumi:"role"`
}

// The set of arguments for constructing a BigtableInstanceIamBinding resource.
type BigtableInstanceIamBindingArgs struct {
	BigtableInstanceIamBindingId pulumi.StringPtrInput
	Condition                    BigtableInstanceIamBindingConditionPtrInput
	Instance                     pulumi.StringInput
	Members                      pulumi.StringArrayInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringInput
}

func (BigtableInstanceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigtableInstanceIamBindingArgs)(nil)).Elem()
}

type BigtableInstanceIamBindingInput interface {
	pulumi.Input

	ToBigtableInstanceIamBindingOutput() BigtableInstanceIamBindingOutput
	ToBigtableInstanceIamBindingOutputWithContext(ctx context.Context) BigtableInstanceIamBindingOutput
}

func (*BigtableInstanceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**BigtableInstanceIamBinding)(nil)).Elem()
}

func (i *BigtableInstanceIamBinding) ToBigtableInstanceIamBindingOutput() BigtableInstanceIamBindingOutput {
	return i.ToBigtableInstanceIamBindingOutputWithContext(context.Background())
}

func (i *BigtableInstanceIamBinding) ToBigtableInstanceIamBindingOutputWithContext(ctx context.Context) BigtableInstanceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigtableInstanceIamBindingOutput)
}

type BigtableInstanceIamBindingOutput struct{ *pulumi.OutputState }

func (BigtableInstanceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigtableInstanceIamBinding)(nil)).Elem()
}

func (o BigtableInstanceIamBindingOutput) ToBigtableInstanceIamBindingOutput() BigtableInstanceIamBindingOutput {
	return o
}

func (o BigtableInstanceIamBindingOutput) ToBigtableInstanceIamBindingOutputWithContext(ctx context.Context) BigtableInstanceIamBindingOutput {
	return o
}

func (o BigtableInstanceIamBindingOutput) BigtableInstanceIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableInstanceIamBinding) pulumi.StringOutput { return v.BigtableInstanceIamBindingId }).(pulumi.StringOutput)
}

func (o BigtableInstanceIamBindingOutput) Condition() BigtableInstanceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *BigtableInstanceIamBinding) BigtableInstanceIamBindingConditionPtrOutput { return v.Condition }).(BigtableInstanceIamBindingConditionPtrOutput)
}

func (o BigtableInstanceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableInstanceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BigtableInstanceIamBindingOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableInstanceIamBinding) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

func (o BigtableInstanceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BigtableInstanceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o BigtableInstanceIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableInstanceIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o BigtableInstanceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableInstanceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigtableInstanceIamBindingInput)(nil)).Elem(), &BigtableInstanceIamBinding{})
	pulumi.RegisterOutputType(BigtableInstanceIamBindingOutput{})
}
