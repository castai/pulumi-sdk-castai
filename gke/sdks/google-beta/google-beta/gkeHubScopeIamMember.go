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

type GkeHubScopeIamMember struct {
	pulumi.CustomResourceState

	Condition              GkeHubScopeIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                   pulumi.StringOutput                    `pulumi:"etag"`
	GkeHubScopeIamMemberId pulumi.StringOutput                    `pulumi:"gkeHubScopeIamMemberId"`
	Member                 pulumi.StringOutput                    `pulumi:"member"`
	Project                pulumi.StringOutput                    `pulumi:"project"`
	Role                   pulumi.StringOutput                    `pulumi:"role"`
	ScopeId                pulumi.StringOutput                    `pulumi:"scopeId"`
}

// NewGkeHubScopeIamMember registers a new resource with the given unique name, arguments, and options.
func NewGkeHubScopeIamMember(ctx *pulumi.Context,
	name string, args *GkeHubScopeIamMemberArgs, opts ...pulumi.ResourceOption) (*GkeHubScopeIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GkeHubScopeIamMember
	err = ctx.RegisterPackageResource("google-beta:index/gkeHubScopeIamMember:GkeHubScopeIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeHubScopeIamMember gets an existing GkeHubScopeIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeHubScopeIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeHubScopeIamMemberState, opts ...pulumi.ResourceOption) (*GkeHubScopeIamMember, error) {
	var resource GkeHubScopeIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/gkeHubScopeIamMember:GkeHubScopeIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeHubScopeIamMember resources.
type gkeHubScopeIamMemberState struct {
	Condition              *GkeHubScopeIamMemberCondition `pulumi:"condition"`
	Etag                   *string                        `pulumi:"etag"`
	GkeHubScopeIamMemberId *string                        `pulumi:"gkeHubScopeIamMemberId"`
	Member                 *string                        `pulumi:"member"`
	Project                *string                        `pulumi:"project"`
	Role                   *string                        `pulumi:"role"`
	ScopeId                *string                        `pulumi:"scopeId"`
}

type GkeHubScopeIamMemberState struct {
	Condition              GkeHubScopeIamMemberConditionPtrInput
	Etag                   pulumi.StringPtrInput
	GkeHubScopeIamMemberId pulumi.StringPtrInput
	Member                 pulumi.StringPtrInput
	Project                pulumi.StringPtrInput
	Role                   pulumi.StringPtrInput
	ScopeId                pulumi.StringPtrInput
}

func (GkeHubScopeIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubScopeIamMemberState)(nil)).Elem()
}

type gkeHubScopeIamMemberArgs struct {
	Condition              *GkeHubScopeIamMemberCondition `pulumi:"condition"`
	GkeHubScopeIamMemberId *string                        `pulumi:"gkeHubScopeIamMemberId"`
	Member                 string                         `pulumi:"member"`
	Project                *string                        `pulumi:"project"`
	Role                   string                         `pulumi:"role"`
	ScopeId                string                         `pulumi:"scopeId"`
}

// The set of arguments for constructing a GkeHubScopeIamMember resource.
type GkeHubScopeIamMemberArgs struct {
	Condition              GkeHubScopeIamMemberConditionPtrInput
	GkeHubScopeIamMemberId pulumi.StringPtrInput
	Member                 pulumi.StringInput
	Project                pulumi.StringPtrInput
	Role                   pulumi.StringInput
	ScopeId                pulumi.StringInput
}

func (GkeHubScopeIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubScopeIamMemberArgs)(nil)).Elem()
}

type GkeHubScopeIamMemberInput interface {
	pulumi.Input

	ToGkeHubScopeIamMemberOutput() GkeHubScopeIamMemberOutput
	ToGkeHubScopeIamMemberOutputWithContext(ctx context.Context) GkeHubScopeIamMemberOutput
}

func (*GkeHubScopeIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubScopeIamMember)(nil)).Elem()
}

func (i *GkeHubScopeIamMember) ToGkeHubScopeIamMemberOutput() GkeHubScopeIamMemberOutput {
	return i.ToGkeHubScopeIamMemberOutputWithContext(context.Background())
}

func (i *GkeHubScopeIamMember) ToGkeHubScopeIamMemberOutputWithContext(ctx context.Context) GkeHubScopeIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeHubScopeIamMemberOutput)
}

type GkeHubScopeIamMemberOutput struct{ *pulumi.OutputState }

func (GkeHubScopeIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubScopeIamMember)(nil)).Elem()
}

func (o GkeHubScopeIamMemberOutput) ToGkeHubScopeIamMemberOutput() GkeHubScopeIamMemberOutput {
	return o
}

func (o GkeHubScopeIamMemberOutput) ToGkeHubScopeIamMemberOutputWithContext(ctx context.Context) GkeHubScopeIamMemberOutput {
	return o
}

func (o GkeHubScopeIamMemberOutput) Condition() GkeHubScopeIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *GkeHubScopeIamMember) GkeHubScopeIamMemberConditionPtrOutput { return v.Condition }).(GkeHubScopeIamMemberConditionPtrOutput)
}

func (o GkeHubScopeIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamMemberOutput) GkeHubScopeIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamMember) pulumi.StringOutput { return v.GkeHubScopeIamMemberId }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamMemberOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamMember) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeHubScopeIamMemberInput)(nil)).Elem(), &GkeHubScopeIamMember{})
	pulumi.RegisterOutputType(GkeHubScopeIamMemberOutput{})
}
