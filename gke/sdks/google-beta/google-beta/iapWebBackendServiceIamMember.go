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

type IapWebBackendServiceIamMember struct {
	pulumi.CustomResourceState

	Condition                       IapWebBackendServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                            pulumi.StringOutput                             `pulumi:"etag"`
	IapWebBackendServiceIamMemberId pulumi.StringOutput                             `pulumi:"iapWebBackendServiceIamMemberId"`
	Member                          pulumi.StringOutput                             `pulumi:"member"`
	Project                         pulumi.StringOutput                             `pulumi:"project"`
	Role                            pulumi.StringOutput                             `pulumi:"role"`
	WebBackendService               pulumi.StringOutput                             `pulumi:"webBackendService"`
}

// NewIapWebBackendServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewIapWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, args *IapWebBackendServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*IapWebBackendServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WebBackendService == nil {
		return nil, errors.New("invalid value for required argument 'WebBackendService'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IapWebBackendServiceIamMember
	err = ctx.RegisterPackageResource("google-beta:index/iapWebBackendServiceIamMember:IapWebBackendServiceIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIapWebBackendServiceIamMember gets an existing IapWebBackendServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIapWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IapWebBackendServiceIamMemberState, opts ...pulumi.ResourceOption) (*IapWebBackendServiceIamMember, error) {
	var resource IapWebBackendServiceIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/iapWebBackendServiceIamMember:IapWebBackendServiceIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IapWebBackendServiceIamMember resources.
type iapWebBackendServiceIamMemberState struct {
	Condition                       *IapWebBackendServiceIamMemberCondition `pulumi:"condition"`
	Etag                            *string                                 `pulumi:"etag"`
	IapWebBackendServiceIamMemberId *string                                 `pulumi:"iapWebBackendServiceIamMemberId"`
	Member                          *string                                 `pulumi:"member"`
	Project                         *string                                 `pulumi:"project"`
	Role                            *string                                 `pulumi:"role"`
	WebBackendService               *string                                 `pulumi:"webBackendService"`
}

type IapWebBackendServiceIamMemberState struct {
	Condition                       IapWebBackendServiceIamMemberConditionPtrInput
	Etag                            pulumi.StringPtrInput
	IapWebBackendServiceIamMemberId pulumi.StringPtrInput
	Member                          pulumi.StringPtrInput
	Project                         pulumi.StringPtrInput
	Role                            pulumi.StringPtrInput
	WebBackendService               pulumi.StringPtrInput
}

func (IapWebBackendServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iapWebBackendServiceIamMemberState)(nil)).Elem()
}

type iapWebBackendServiceIamMemberArgs struct {
	Condition                       *IapWebBackendServiceIamMemberCondition `pulumi:"condition"`
	IapWebBackendServiceIamMemberId *string                                 `pulumi:"iapWebBackendServiceIamMemberId"`
	Member                          string                                  `pulumi:"member"`
	Project                         *string                                 `pulumi:"project"`
	Role                            string                                  `pulumi:"role"`
	WebBackendService               string                                  `pulumi:"webBackendService"`
}

// The set of arguments for constructing a IapWebBackendServiceIamMember resource.
type IapWebBackendServiceIamMemberArgs struct {
	Condition                       IapWebBackendServiceIamMemberConditionPtrInput
	IapWebBackendServiceIamMemberId pulumi.StringPtrInput
	Member                          pulumi.StringInput
	Project                         pulumi.StringPtrInput
	Role                            pulumi.StringInput
	WebBackendService               pulumi.StringInput
}

func (IapWebBackendServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iapWebBackendServiceIamMemberArgs)(nil)).Elem()
}

type IapWebBackendServiceIamMemberInput interface {
	pulumi.Input

	ToIapWebBackendServiceIamMemberOutput() IapWebBackendServiceIamMemberOutput
	ToIapWebBackendServiceIamMemberOutputWithContext(ctx context.Context) IapWebBackendServiceIamMemberOutput
}

func (*IapWebBackendServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IapWebBackendServiceIamMember)(nil)).Elem()
}

func (i *IapWebBackendServiceIamMember) ToIapWebBackendServiceIamMemberOutput() IapWebBackendServiceIamMemberOutput {
	return i.ToIapWebBackendServiceIamMemberOutputWithContext(context.Background())
}

func (i *IapWebBackendServiceIamMember) ToIapWebBackendServiceIamMemberOutputWithContext(ctx context.Context) IapWebBackendServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IapWebBackendServiceIamMemberOutput)
}

type IapWebBackendServiceIamMemberOutput struct{ *pulumi.OutputState }

func (IapWebBackendServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IapWebBackendServiceIamMember)(nil)).Elem()
}

func (o IapWebBackendServiceIamMemberOutput) ToIapWebBackendServiceIamMemberOutput() IapWebBackendServiceIamMemberOutput {
	return o
}

func (o IapWebBackendServiceIamMemberOutput) ToIapWebBackendServiceIamMemberOutputWithContext(ctx context.Context) IapWebBackendServiceIamMemberOutput {
	return o
}

func (o IapWebBackendServiceIamMemberOutput) Condition() IapWebBackendServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *IapWebBackendServiceIamMember) IapWebBackendServiceIamMemberConditionPtrOutput {
		return v.Condition
	}).(IapWebBackendServiceIamMemberConditionPtrOutput)
}

func (o IapWebBackendServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebBackendServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IapWebBackendServiceIamMemberOutput) IapWebBackendServiceIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebBackendServiceIamMember) pulumi.StringOutput { return v.IapWebBackendServiceIamMemberId }).(pulumi.StringOutput)
}

func (o IapWebBackendServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebBackendServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o IapWebBackendServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebBackendServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o IapWebBackendServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebBackendServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o IapWebBackendServiceIamMemberOutput) WebBackendService() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebBackendServiceIamMember) pulumi.StringOutput { return v.WebBackendService }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IapWebBackendServiceIamMemberInput)(nil)).Elem(), &IapWebBackendServiceIamMember{})
	pulumi.RegisterOutputType(IapWebBackendServiceIamMemberOutput{})
}
