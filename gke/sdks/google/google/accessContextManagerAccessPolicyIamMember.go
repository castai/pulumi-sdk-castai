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

type AccessContextManagerAccessPolicyIamMember struct {
	pulumi.CustomResourceState

	AccessContextManagerAccessPolicyIamMemberId pulumi.StringOutput                                         `pulumi:"accessContextManagerAccessPolicyIamMemberId"`
	Condition                                   AccessContextManagerAccessPolicyIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                                        pulumi.StringOutput                                         `pulumi:"etag"`
	Member                                      pulumi.StringOutput                                         `pulumi:"member"`
	Name                                        pulumi.StringOutput                                         `pulumi:"name"`
	Role                                        pulumi.StringOutput                                         `pulumi:"role"`
}

// NewAccessContextManagerAccessPolicyIamMember registers a new resource with the given unique name, arguments, and options.
func NewAccessContextManagerAccessPolicyIamMember(ctx *pulumi.Context,
	name string, args *AccessContextManagerAccessPolicyIamMemberArgs, opts ...pulumi.ResourceOption) (*AccessContextManagerAccessPolicyIamMember, error) {
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
	var resource AccessContextManagerAccessPolicyIamMember
	err = ctx.RegisterPackageResource("google:index/accessContextManagerAccessPolicyIamMember:AccessContextManagerAccessPolicyIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessContextManagerAccessPolicyIamMember gets an existing AccessContextManagerAccessPolicyIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessContextManagerAccessPolicyIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessContextManagerAccessPolicyIamMemberState, opts ...pulumi.ResourceOption) (*AccessContextManagerAccessPolicyIamMember, error) {
	var resource AccessContextManagerAccessPolicyIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/accessContextManagerAccessPolicyIamMember:AccessContextManagerAccessPolicyIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessContextManagerAccessPolicyIamMember resources.
type accessContextManagerAccessPolicyIamMemberState struct {
	AccessContextManagerAccessPolicyIamMemberId *string                                             `pulumi:"accessContextManagerAccessPolicyIamMemberId"`
	Condition                                   *AccessContextManagerAccessPolicyIamMemberCondition `pulumi:"condition"`
	Etag                                        *string                                             `pulumi:"etag"`
	Member                                      *string                                             `pulumi:"member"`
	Name                                        *string                                             `pulumi:"name"`
	Role                                        *string                                             `pulumi:"role"`
}

type AccessContextManagerAccessPolicyIamMemberState struct {
	AccessContextManagerAccessPolicyIamMemberId pulumi.StringPtrInput
	Condition                                   AccessContextManagerAccessPolicyIamMemberConditionPtrInput
	Etag                                        pulumi.StringPtrInput
	Member                                      pulumi.StringPtrInput
	Name                                        pulumi.StringPtrInput
	Role                                        pulumi.StringPtrInput
}

func (AccessContextManagerAccessPolicyIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerAccessPolicyIamMemberState)(nil)).Elem()
}

type accessContextManagerAccessPolicyIamMemberArgs struct {
	AccessContextManagerAccessPolicyIamMemberId *string                                             `pulumi:"accessContextManagerAccessPolicyIamMemberId"`
	Condition                                   *AccessContextManagerAccessPolicyIamMemberCondition `pulumi:"condition"`
	Member                                      string                                              `pulumi:"member"`
	Name                                        *string                                             `pulumi:"name"`
	Role                                        string                                              `pulumi:"role"`
}

// The set of arguments for constructing a AccessContextManagerAccessPolicyIamMember resource.
type AccessContextManagerAccessPolicyIamMemberArgs struct {
	AccessContextManagerAccessPolicyIamMemberId pulumi.StringPtrInput
	Condition                                   AccessContextManagerAccessPolicyIamMemberConditionPtrInput
	Member                                      pulumi.StringInput
	Name                                        pulumi.StringPtrInput
	Role                                        pulumi.StringInput
}

func (AccessContextManagerAccessPolicyIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerAccessPolicyIamMemberArgs)(nil)).Elem()
}

type AccessContextManagerAccessPolicyIamMemberInput interface {
	pulumi.Input

	ToAccessContextManagerAccessPolicyIamMemberOutput() AccessContextManagerAccessPolicyIamMemberOutput
	ToAccessContextManagerAccessPolicyIamMemberOutputWithContext(ctx context.Context) AccessContextManagerAccessPolicyIamMemberOutput
}

func (*AccessContextManagerAccessPolicyIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerAccessPolicyIamMember)(nil)).Elem()
}

func (i *AccessContextManagerAccessPolicyIamMember) ToAccessContextManagerAccessPolicyIamMemberOutput() AccessContextManagerAccessPolicyIamMemberOutput {
	return i.ToAccessContextManagerAccessPolicyIamMemberOutputWithContext(context.Background())
}

func (i *AccessContextManagerAccessPolicyIamMember) ToAccessContextManagerAccessPolicyIamMemberOutputWithContext(ctx context.Context) AccessContextManagerAccessPolicyIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessContextManagerAccessPolicyIamMemberOutput)
}

type AccessContextManagerAccessPolicyIamMemberOutput struct{ *pulumi.OutputState }

func (AccessContextManagerAccessPolicyIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerAccessPolicyIamMember)(nil)).Elem()
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) ToAccessContextManagerAccessPolicyIamMemberOutput() AccessContextManagerAccessPolicyIamMemberOutput {
	return o
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) ToAccessContextManagerAccessPolicyIamMemberOutputWithContext(ctx context.Context) AccessContextManagerAccessPolicyIamMemberOutput {
	return o
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) AccessContextManagerAccessPolicyIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessPolicyIamMember) pulumi.StringOutput {
		return v.AccessContextManagerAccessPolicyIamMemberId
	}).(pulumi.StringOutput)
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) Condition() AccessContextManagerAccessPolicyIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessPolicyIamMember) AccessContextManagerAccessPolicyIamMemberConditionPtrOutput {
		return v.Condition
	}).(AccessContextManagerAccessPolicyIamMemberConditionPtrOutput)
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessPolicyIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessPolicyIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessPolicyIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccessContextManagerAccessPolicyIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessPolicyIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessContextManagerAccessPolicyIamMemberInput)(nil)).Elem(), &AccessContextManagerAccessPolicyIamMember{})
	pulumi.RegisterOutputType(AccessContextManagerAccessPolicyIamMemberOutput{})
}
