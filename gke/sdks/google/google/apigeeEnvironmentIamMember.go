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

type ApigeeEnvironmentIamMember struct {
	pulumi.CustomResourceState

	ApigeeEnvironmentIamMemberId pulumi.StringOutput                          `pulumi:"apigeeEnvironmentIamMemberId"`
	Condition                    ApigeeEnvironmentIamMemberConditionPtrOutput `pulumi:"condition"`
	EnvId                        pulumi.StringOutput                          `pulumi:"envId"`
	Etag                         pulumi.StringOutput                          `pulumi:"etag"`
	Member                       pulumi.StringOutput                          `pulumi:"member"`
	OrgId                        pulumi.StringOutput                          `pulumi:"orgId"`
	Role                         pulumi.StringOutput                          `pulumi:"role"`
}

// NewApigeeEnvironmentIamMember registers a new resource with the given unique name, arguments, and options.
func NewApigeeEnvironmentIamMember(ctx *pulumi.Context,
	name string, args *ApigeeEnvironmentIamMemberArgs, opts ...pulumi.ResourceOption) (*ApigeeEnvironmentIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvId == nil {
		return nil, errors.New("invalid value for required argument 'EnvId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
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
	var resource ApigeeEnvironmentIamMember
	err = ctx.RegisterPackageResource("google:index/apigeeEnvironmentIamMember:ApigeeEnvironmentIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeEnvironmentIamMember gets an existing ApigeeEnvironmentIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeEnvironmentIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeEnvironmentIamMemberState, opts ...pulumi.ResourceOption) (*ApigeeEnvironmentIamMember, error) {
	var resource ApigeeEnvironmentIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apigeeEnvironmentIamMember:ApigeeEnvironmentIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeEnvironmentIamMember resources.
type apigeeEnvironmentIamMemberState struct {
	ApigeeEnvironmentIamMemberId *string                              `pulumi:"apigeeEnvironmentIamMemberId"`
	Condition                    *ApigeeEnvironmentIamMemberCondition `pulumi:"condition"`
	EnvId                        *string                              `pulumi:"envId"`
	Etag                         *string                              `pulumi:"etag"`
	Member                       *string                              `pulumi:"member"`
	OrgId                        *string                              `pulumi:"orgId"`
	Role                         *string                              `pulumi:"role"`
}

type ApigeeEnvironmentIamMemberState struct {
	ApigeeEnvironmentIamMemberId pulumi.StringPtrInput
	Condition                    ApigeeEnvironmentIamMemberConditionPtrInput
	EnvId                        pulumi.StringPtrInput
	Etag                         pulumi.StringPtrInput
	Member                       pulumi.StringPtrInput
	OrgId                        pulumi.StringPtrInput
	Role                         pulumi.StringPtrInput
}

func (ApigeeEnvironmentIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeEnvironmentIamMemberState)(nil)).Elem()
}

type apigeeEnvironmentIamMemberArgs struct {
	ApigeeEnvironmentIamMemberId *string                              `pulumi:"apigeeEnvironmentIamMemberId"`
	Condition                    *ApigeeEnvironmentIamMemberCondition `pulumi:"condition"`
	EnvId                        string                               `pulumi:"envId"`
	Member                       string                               `pulumi:"member"`
	OrgId                        string                               `pulumi:"orgId"`
	Role                         string                               `pulumi:"role"`
}

// The set of arguments for constructing a ApigeeEnvironmentIamMember resource.
type ApigeeEnvironmentIamMemberArgs struct {
	ApigeeEnvironmentIamMemberId pulumi.StringPtrInput
	Condition                    ApigeeEnvironmentIamMemberConditionPtrInput
	EnvId                        pulumi.StringInput
	Member                       pulumi.StringInput
	OrgId                        pulumi.StringInput
	Role                         pulumi.StringInput
}

func (ApigeeEnvironmentIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeEnvironmentIamMemberArgs)(nil)).Elem()
}

type ApigeeEnvironmentIamMemberInput interface {
	pulumi.Input

	ToApigeeEnvironmentIamMemberOutput() ApigeeEnvironmentIamMemberOutput
	ToApigeeEnvironmentIamMemberOutputWithContext(ctx context.Context) ApigeeEnvironmentIamMemberOutput
}

func (*ApigeeEnvironmentIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeEnvironmentIamMember)(nil)).Elem()
}

func (i *ApigeeEnvironmentIamMember) ToApigeeEnvironmentIamMemberOutput() ApigeeEnvironmentIamMemberOutput {
	return i.ToApigeeEnvironmentIamMemberOutputWithContext(context.Background())
}

func (i *ApigeeEnvironmentIamMember) ToApigeeEnvironmentIamMemberOutputWithContext(ctx context.Context) ApigeeEnvironmentIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeEnvironmentIamMemberOutput)
}

type ApigeeEnvironmentIamMemberOutput struct{ *pulumi.OutputState }

func (ApigeeEnvironmentIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeEnvironmentIamMember)(nil)).Elem()
}

func (o ApigeeEnvironmentIamMemberOutput) ToApigeeEnvironmentIamMemberOutput() ApigeeEnvironmentIamMemberOutput {
	return o
}

func (o ApigeeEnvironmentIamMemberOutput) ToApigeeEnvironmentIamMemberOutputWithContext(ctx context.Context) ApigeeEnvironmentIamMemberOutput {
	return o
}

func (o ApigeeEnvironmentIamMemberOutput) ApigeeEnvironmentIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamMember) pulumi.StringOutput { return v.ApigeeEnvironmentIamMemberId }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamMemberOutput) Condition() ApigeeEnvironmentIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamMember) ApigeeEnvironmentIamMemberConditionPtrOutput { return v.Condition }).(ApigeeEnvironmentIamMemberConditionPtrOutput)
}

func (o ApigeeEnvironmentIamMemberOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamMember) pulumi.StringOutput { return v.EnvId }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamMemberOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamMember) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

func (o ApigeeEnvironmentIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvironmentIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeEnvironmentIamMemberInput)(nil)).Elem(), &ApigeeEnvironmentIamMember{})
	pulumi.RegisterOutputType(ApigeeEnvironmentIamMemberOutput{})
}
