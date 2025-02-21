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

type ApigeeEnvironmentIamBinding struct {
	pulumi.CustomResourceState

	ApigeeEnvironmentIamBindingId pulumi.StringOutput                           `pulumi:"apigeeEnvironmentIamBindingId"`
	Condition                     ApigeeEnvironmentIamBindingConditionPtrOutput `pulumi:"condition"`
	EnvId                         pulumi.StringOutput                           `pulumi:"envId"`
	Etag                          pulumi.StringOutput                           `pulumi:"etag"`
	Members                       pulumi.StringArrayOutput                      `pulumi:"members"`
	OrgId                         pulumi.StringOutput                           `pulumi:"orgId"`
	Role                          pulumi.StringOutput                           `pulumi:"role"`
}

// NewApigeeEnvironmentIamBinding registers a new resource with the given unique name, arguments, and options.
func NewApigeeEnvironmentIamBinding(ctx *pulumi.Context,
	name string, args *ApigeeEnvironmentIamBindingArgs, opts ...pulumi.ResourceOption) (*ApigeeEnvironmentIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvId == nil {
		return nil, errors.New("invalid value for required argument 'EnvId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApigeeEnvironmentIamBinding
	err = ctx.RegisterPackageResource("google:index/apigeeEnvironmentIamBinding:ApigeeEnvironmentIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeEnvironmentIamBinding gets an existing ApigeeEnvironmentIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeEnvironmentIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeEnvironmentIamBindingState, opts ...pulumi.ResourceOption) (*ApigeeEnvironmentIamBinding, error) {
	var resource ApigeeEnvironmentIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apigeeEnvironmentIamBinding:ApigeeEnvironmentIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeEnvironmentIamBinding resources.
type apigeeEnvironmentIamBindingState struct {
	ApigeeEnvironmentIamBindingId *string                               `pulumi:"apigeeEnvironmentIamBindingId"`
	Condition                     *ApigeeEnvironmentIamBindingCondition `pulumi:"condition"`
	EnvId                         *string                               `pulumi:"envId"`
	Etag                          *string                               `pulumi:"etag"`
	Members                       []string                              `pulumi:"members"`
	OrgId                         *string                               `pulumi:"orgId"`
	Role                          *string                               `pulumi:"role"`
}

type ApigeeEnvironmentIamBindingState struct {
	ApigeeEnvironmentIamBindingId pulumi.StringPtrInput
	Condition                     ApigeeEnvironmentIamBindingConditionPtrInput
	EnvId                         pulumi.StringPtrInput
	Etag                          pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	OrgId                         pulumi.StringPtrInput
	Role                          pulumi.StringPtrInput
}

func (ApigeeEnvironmentIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeEnvironmentIamBindingState)(nil)).Elem()
}

type apigeeEnvironmentIamBindingArgs struct {
	ApigeeEnvironmentIamBindingId *string                               `pulumi:"apigeeEnvironmentIamBindingId"`
	Condition                     *ApigeeEnvironmentIamBindingCondition `pulumi:"condition"`
	EnvId                         string                                `pulumi:"envId"`
	Members                       []string                              `pulumi:"members"`
	OrgId                         string                                `pulumi:"orgId"`
	Role                          string                                `pulumi:"role"`
}

// The set of arguments for constructing a ApigeeEnvironmentIamBinding resource.
type ApigeeEnvironmentIamBindingArgs struct {
	ApigeeEnvironmentIamBindingId pulumi.StringPtrInput
	Condition                     ApigeeEnvironmentIamBindingConditionPtrInput
	EnvId                         pulumi.StringInput
	Members                       pulumi.StringArrayInput
	OrgId                         pulumi.StringInput
	Role                          pulumi.StringInput
}

func (ApigeeEnvironmentIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeEnvironmentIamBindingArgs)(nil)).Elem()
}

type ApigeeEnvironmentIamBindingInput interface {
	pulumi.Input

	ToApigeeEnvironmentIamBindingOutput() ApigeeEnvironmentIamBindingOutput
	ToApigeeEnvironmentIamBindingOutputWithContext(ctx context.Context) ApigeeEnvironmentIamBindingOutput
}

func (*ApigeeEnvironmentIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeEnvironmentIamBinding)(nil)).Elem()
}

func (i *ApigeeEnvironmentIamBinding) ToApigeeEnvironmentIamBindingOutput() ApigeeEnvironmentIamBindingOutput {
	return i.ToApigeeEnvironmentIamBindingOutputWithContext(context.Background())
}

func (i *ApigeeEnvironmentIamBinding) ToApigeeEnvironmentIamBindingOutputWithContext(ctx context.Context) ApigeeEnvironmentIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeEnvironmentIamBindingOutput)
}

type ApigeeEnvironmentIamBindingOutput struct{ *pulumi.OutputState }

func (ApigeeEnvironmentIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeEnvironmentIamBinding)(nil)).Elem()
}

func (o ApigeeEnvironmentIamBindingOutput) ToApigeeEnvironmentIamBindingOutput() ApigeeEnvironmentIamBindingOutput {
	return o
}

func (o ApigeeEnvironmentIamBindingOutput) ToApigeeEnvironmentIamBindingOutputWithContext(ctx context.Context) ApigeeEnvironmentIamBindingOutput {
	return o
}

func (o ApigeeEnvironmentIamBindingOutput) ApigeeEnvironmentIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamBinding) pulumi.StringOutput { return v.ApigeeEnvironmentIamBindingId }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamBindingOutput) Condition() ApigeeEnvironmentIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamBinding) ApigeeEnvironmentIamBindingConditionPtrOutput { return v.Condition }).(ApigeeEnvironmentIamBindingConditionPtrOutput)
}

func (o ApigeeEnvironmentIamBindingOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamBinding) pulumi.StringOutput { return v.EnvId }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o ApigeeEnvironmentIamBindingOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamBinding) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeEnvironmentIamBindingInput)(nil)).Elem(), &ApigeeEnvironmentIamBinding{})
	pulumi.RegisterOutputType(ApigeeEnvironmentIamBindingOutput{})
}
