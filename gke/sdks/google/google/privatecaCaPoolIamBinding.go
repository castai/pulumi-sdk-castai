// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivatecaCaPoolIamBinding struct {
	pulumi.CustomResourceState

	CaPool                      pulumi.StringOutput                         `pulumi:"caPool"`
	Condition                   PrivatecaCaPoolIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                        pulumi.StringOutput                         `pulumi:"etag"`
	Location                    pulumi.StringOutput                         `pulumi:"location"`
	Members                     pulumi.StringArrayOutput                    `pulumi:"members"`
	PrivatecaCaPoolIamBindingId pulumi.StringOutput                         `pulumi:"privatecaCaPoolIamBindingId"`
	Project                     pulumi.StringOutput                         `pulumi:"project"`
	Role                        pulumi.StringOutput                         `pulumi:"role"`
}

// NewPrivatecaCaPoolIamBinding registers a new resource with the given unique name, arguments, and options.
func NewPrivatecaCaPoolIamBinding(ctx *pulumi.Context,
	name string, args *PrivatecaCaPoolIamBindingArgs, opts ...pulumi.ResourceOption) (*PrivatecaCaPoolIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
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
	var resource PrivatecaCaPoolIamBinding
	err = ctx.RegisterPackageResource("google:index/privatecaCaPoolIamBinding:PrivatecaCaPoolIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivatecaCaPoolIamBinding gets an existing PrivatecaCaPoolIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivatecaCaPoolIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivatecaCaPoolIamBindingState, opts ...pulumi.ResourceOption) (*PrivatecaCaPoolIamBinding, error) {
	var resource PrivatecaCaPoolIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/privatecaCaPoolIamBinding:PrivatecaCaPoolIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivatecaCaPoolIamBinding resources.
type privatecaCaPoolIamBindingState struct {
	CaPool                      *string                             `pulumi:"caPool"`
	Condition                   *PrivatecaCaPoolIamBindingCondition `pulumi:"condition"`
	Etag                        *string                             `pulumi:"etag"`
	Location                    *string                             `pulumi:"location"`
	Members                     []string                            `pulumi:"members"`
	PrivatecaCaPoolIamBindingId *string                             `pulumi:"privatecaCaPoolIamBindingId"`
	Project                     *string                             `pulumi:"project"`
	Role                        *string                             `pulumi:"role"`
}

type PrivatecaCaPoolIamBindingState struct {
	CaPool                      pulumi.StringPtrInput
	Condition                   PrivatecaCaPoolIamBindingConditionPtrInput
	Etag                        pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	Members                     pulumi.StringArrayInput
	PrivatecaCaPoolIamBindingId pulumi.StringPtrInput
	Project                     pulumi.StringPtrInput
	Role                        pulumi.StringPtrInput
}

func (PrivatecaCaPoolIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*privatecaCaPoolIamBindingState)(nil)).Elem()
}

type privatecaCaPoolIamBindingArgs struct {
	CaPool                      string                              `pulumi:"caPool"`
	Condition                   *PrivatecaCaPoolIamBindingCondition `pulumi:"condition"`
	Location                    *string                             `pulumi:"location"`
	Members                     []string                            `pulumi:"members"`
	PrivatecaCaPoolIamBindingId *string                             `pulumi:"privatecaCaPoolIamBindingId"`
	Project                     *string                             `pulumi:"project"`
	Role                        string                              `pulumi:"role"`
}

// The set of arguments for constructing a PrivatecaCaPoolIamBinding resource.
type PrivatecaCaPoolIamBindingArgs struct {
	CaPool                      pulumi.StringInput
	Condition                   PrivatecaCaPoolIamBindingConditionPtrInput
	Location                    pulumi.StringPtrInput
	Members                     pulumi.StringArrayInput
	PrivatecaCaPoolIamBindingId pulumi.StringPtrInput
	Project                     pulumi.StringPtrInput
	Role                        pulumi.StringInput
}

func (PrivatecaCaPoolIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privatecaCaPoolIamBindingArgs)(nil)).Elem()
}

type PrivatecaCaPoolIamBindingInput interface {
	pulumi.Input

	ToPrivatecaCaPoolIamBindingOutput() PrivatecaCaPoolIamBindingOutput
	ToPrivatecaCaPoolIamBindingOutputWithContext(ctx context.Context) PrivatecaCaPoolIamBindingOutput
}

func (*PrivatecaCaPoolIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatecaCaPoolIamBinding)(nil)).Elem()
}

func (i *PrivatecaCaPoolIamBinding) ToPrivatecaCaPoolIamBindingOutput() PrivatecaCaPoolIamBindingOutput {
	return i.ToPrivatecaCaPoolIamBindingOutputWithContext(context.Background())
}

func (i *PrivatecaCaPoolIamBinding) ToPrivatecaCaPoolIamBindingOutputWithContext(ctx context.Context) PrivatecaCaPoolIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatecaCaPoolIamBindingOutput)
}

type PrivatecaCaPoolIamBindingOutput struct{ *pulumi.OutputState }

func (PrivatecaCaPoolIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatecaCaPoolIamBinding)(nil)).Elem()
}

func (o PrivatecaCaPoolIamBindingOutput) ToPrivatecaCaPoolIamBindingOutput() PrivatecaCaPoolIamBindingOutput {
	return o
}

func (o PrivatecaCaPoolIamBindingOutput) ToPrivatecaCaPoolIamBindingOutputWithContext(ctx context.Context) PrivatecaCaPoolIamBindingOutput {
	return o
}

func (o PrivatecaCaPoolIamBindingOutput) CaPool() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) pulumi.StringOutput { return v.CaPool }).(pulumi.StringOutput)
}

func (o PrivatecaCaPoolIamBindingOutput) Condition() PrivatecaCaPoolIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) PrivatecaCaPoolIamBindingConditionPtrOutput { return v.Condition }).(PrivatecaCaPoolIamBindingConditionPtrOutput)
}

func (o PrivatecaCaPoolIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PrivatecaCaPoolIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivatecaCaPoolIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o PrivatecaCaPoolIamBindingOutput) PrivatecaCaPoolIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) pulumi.StringOutput { return v.PrivatecaCaPoolIamBindingId }).(pulumi.StringOutput)
}

func (o PrivatecaCaPoolIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o PrivatecaCaPoolIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCaPoolIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatecaCaPoolIamBindingInput)(nil)).Elem(), &PrivatecaCaPoolIamBinding{})
	pulumi.RegisterOutputType(PrivatecaCaPoolIamBindingOutput{})
}
