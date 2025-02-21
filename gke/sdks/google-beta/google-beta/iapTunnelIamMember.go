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

type IapTunnelIamMember struct {
	pulumi.CustomResourceState

	Condition            IapTunnelIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                 pulumi.StringOutput                  `pulumi:"etag"`
	IapTunnelIamMemberId pulumi.StringOutput                  `pulumi:"iapTunnelIamMemberId"`
	Member               pulumi.StringOutput                  `pulumi:"member"`
	Project              pulumi.StringOutput                  `pulumi:"project"`
	Role                 pulumi.StringOutput                  `pulumi:"role"`
}

// NewIapTunnelIamMember registers a new resource with the given unique name, arguments, and options.
func NewIapTunnelIamMember(ctx *pulumi.Context,
	name string, args *IapTunnelIamMemberArgs, opts ...pulumi.ResourceOption) (*IapTunnelIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IapTunnelIamMember
	err = ctx.RegisterPackageResource("google-beta:index/iapTunnelIamMember:IapTunnelIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIapTunnelIamMember gets an existing IapTunnelIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIapTunnelIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IapTunnelIamMemberState, opts ...pulumi.ResourceOption) (*IapTunnelIamMember, error) {
	var resource IapTunnelIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/iapTunnelIamMember:IapTunnelIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IapTunnelIamMember resources.
type iapTunnelIamMemberState struct {
	Condition            *IapTunnelIamMemberCondition `pulumi:"condition"`
	Etag                 *string                      `pulumi:"etag"`
	IapTunnelIamMemberId *string                      `pulumi:"iapTunnelIamMemberId"`
	Member               *string                      `pulumi:"member"`
	Project              *string                      `pulumi:"project"`
	Role                 *string                      `pulumi:"role"`
}

type IapTunnelIamMemberState struct {
	Condition            IapTunnelIamMemberConditionPtrInput
	Etag                 pulumi.StringPtrInput
	IapTunnelIamMemberId pulumi.StringPtrInput
	Member               pulumi.StringPtrInput
	Project              pulumi.StringPtrInput
	Role                 pulumi.StringPtrInput
}

func (IapTunnelIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iapTunnelIamMemberState)(nil)).Elem()
}

type iapTunnelIamMemberArgs struct {
	Condition            *IapTunnelIamMemberCondition `pulumi:"condition"`
	IapTunnelIamMemberId *string                      `pulumi:"iapTunnelIamMemberId"`
	Member               string                       `pulumi:"member"`
	Project              *string                      `pulumi:"project"`
	Role                 string                       `pulumi:"role"`
}

// The set of arguments for constructing a IapTunnelIamMember resource.
type IapTunnelIamMemberArgs struct {
	Condition            IapTunnelIamMemberConditionPtrInput
	IapTunnelIamMemberId pulumi.StringPtrInput
	Member               pulumi.StringInput
	Project              pulumi.StringPtrInput
	Role                 pulumi.StringInput
}

func (IapTunnelIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iapTunnelIamMemberArgs)(nil)).Elem()
}

type IapTunnelIamMemberInput interface {
	pulumi.Input

	ToIapTunnelIamMemberOutput() IapTunnelIamMemberOutput
	ToIapTunnelIamMemberOutputWithContext(ctx context.Context) IapTunnelIamMemberOutput
}

func (*IapTunnelIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IapTunnelIamMember)(nil)).Elem()
}

func (i *IapTunnelIamMember) ToIapTunnelIamMemberOutput() IapTunnelIamMemberOutput {
	return i.ToIapTunnelIamMemberOutputWithContext(context.Background())
}

func (i *IapTunnelIamMember) ToIapTunnelIamMemberOutputWithContext(ctx context.Context) IapTunnelIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IapTunnelIamMemberOutput)
}

type IapTunnelIamMemberOutput struct{ *pulumi.OutputState }

func (IapTunnelIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IapTunnelIamMember)(nil)).Elem()
}

func (o IapTunnelIamMemberOutput) ToIapTunnelIamMemberOutput() IapTunnelIamMemberOutput {
	return o
}

func (o IapTunnelIamMemberOutput) ToIapTunnelIamMemberOutputWithContext(ctx context.Context) IapTunnelIamMemberOutput {
	return o
}

func (o IapTunnelIamMemberOutput) Condition() IapTunnelIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *IapTunnelIamMember) IapTunnelIamMemberConditionPtrOutput { return v.Condition }).(IapTunnelIamMemberConditionPtrOutput)
}

func (o IapTunnelIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IapTunnelIamMemberOutput) IapTunnelIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamMember) pulumi.StringOutput { return v.IapTunnelIamMemberId }).(pulumi.StringOutput)
}

func (o IapTunnelIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o IapTunnelIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o IapTunnelIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IapTunnelIamMemberInput)(nil)).Elem(), &IapTunnelIamMember{})
	pulumi.RegisterOutputType(IapTunnelIamMemberOutput{})
}
