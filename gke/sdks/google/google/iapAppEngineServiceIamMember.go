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

type IapAppEngineServiceIamMember struct {
	pulumi.CustomResourceState

	AppId                          pulumi.StringOutput                            `pulumi:"appId"`
	Condition                      IapAppEngineServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                           pulumi.StringOutput                            `pulumi:"etag"`
	IapAppEngineServiceIamMemberId pulumi.StringOutput                            `pulumi:"iapAppEngineServiceIamMemberId"`
	Member                         pulumi.StringOutput                            `pulumi:"member"`
	Project                        pulumi.StringOutput                            `pulumi:"project"`
	Role                           pulumi.StringOutput                            `pulumi:"role"`
	Service                        pulumi.StringOutput                            `pulumi:"service"`
}

// NewIapAppEngineServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewIapAppEngineServiceIamMember(ctx *pulumi.Context,
	name string, args *IapAppEngineServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*IapAppEngineServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IapAppEngineServiceIamMember
	err = ctx.RegisterPackageResource("google:index/iapAppEngineServiceIamMember:IapAppEngineServiceIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIapAppEngineServiceIamMember gets an existing IapAppEngineServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIapAppEngineServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IapAppEngineServiceIamMemberState, opts ...pulumi.ResourceOption) (*IapAppEngineServiceIamMember, error) {
	var resource IapAppEngineServiceIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/iapAppEngineServiceIamMember:IapAppEngineServiceIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IapAppEngineServiceIamMember resources.
type iapAppEngineServiceIamMemberState struct {
	AppId                          *string                                `pulumi:"appId"`
	Condition                      *IapAppEngineServiceIamMemberCondition `pulumi:"condition"`
	Etag                           *string                                `pulumi:"etag"`
	IapAppEngineServiceIamMemberId *string                                `pulumi:"iapAppEngineServiceIamMemberId"`
	Member                         *string                                `pulumi:"member"`
	Project                        *string                                `pulumi:"project"`
	Role                           *string                                `pulumi:"role"`
	Service                        *string                                `pulumi:"service"`
}

type IapAppEngineServiceIamMemberState struct {
	AppId                          pulumi.StringPtrInput
	Condition                      IapAppEngineServiceIamMemberConditionPtrInput
	Etag                           pulumi.StringPtrInput
	IapAppEngineServiceIamMemberId pulumi.StringPtrInput
	Member                         pulumi.StringPtrInput
	Project                        pulumi.StringPtrInput
	Role                           pulumi.StringPtrInput
	Service                        pulumi.StringPtrInput
}

func (IapAppEngineServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iapAppEngineServiceIamMemberState)(nil)).Elem()
}

type iapAppEngineServiceIamMemberArgs struct {
	AppId                          string                                 `pulumi:"appId"`
	Condition                      *IapAppEngineServiceIamMemberCondition `pulumi:"condition"`
	IapAppEngineServiceIamMemberId *string                                `pulumi:"iapAppEngineServiceIamMemberId"`
	Member                         string                                 `pulumi:"member"`
	Project                        *string                                `pulumi:"project"`
	Role                           string                                 `pulumi:"role"`
	Service                        string                                 `pulumi:"service"`
}

// The set of arguments for constructing a IapAppEngineServiceIamMember resource.
type IapAppEngineServiceIamMemberArgs struct {
	AppId                          pulumi.StringInput
	Condition                      IapAppEngineServiceIamMemberConditionPtrInput
	IapAppEngineServiceIamMemberId pulumi.StringPtrInput
	Member                         pulumi.StringInput
	Project                        pulumi.StringPtrInput
	Role                           pulumi.StringInput
	Service                        pulumi.StringInput
}

func (IapAppEngineServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iapAppEngineServiceIamMemberArgs)(nil)).Elem()
}

type IapAppEngineServiceIamMemberInput interface {
	pulumi.Input

	ToIapAppEngineServiceIamMemberOutput() IapAppEngineServiceIamMemberOutput
	ToIapAppEngineServiceIamMemberOutputWithContext(ctx context.Context) IapAppEngineServiceIamMemberOutput
}

func (*IapAppEngineServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IapAppEngineServiceIamMember)(nil)).Elem()
}

func (i *IapAppEngineServiceIamMember) ToIapAppEngineServiceIamMemberOutput() IapAppEngineServiceIamMemberOutput {
	return i.ToIapAppEngineServiceIamMemberOutputWithContext(context.Background())
}

func (i *IapAppEngineServiceIamMember) ToIapAppEngineServiceIamMemberOutputWithContext(ctx context.Context) IapAppEngineServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IapAppEngineServiceIamMemberOutput)
}

type IapAppEngineServiceIamMemberOutput struct{ *pulumi.OutputState }

func (IapAppEngineServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IapAppEngineServiceIamMember)(nil)).Elem()
}

func (o IapAppEngineServiceIamMemberOutput) ToIapAppEngineServiceIamMemberOutput() IapAppEngineServiceIamMemberOutput {
	return o
}

func (o IapAppEngineServiceIamMemberOutput) ToIapAppEngineServiceIamMemberOutputWithContext(ctx context.Context) IapAppEngineServiceIamMemberOutput {
	return o
}

func (o IapAppEngineServiceIamMemberOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o IapAppEngineServiceIamMemberOutput) Condition() IapAppEngineServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) IapAppEngineServiceIamMemberConditionPtrOutput {
		return v.Condition
	}).(IapAppEngineServiceIamMemberConditionPtrOutput)
}

func (o IapAppEngineServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IapAppEngineServiceIamMemberOutput) IapAppEngineServiceIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) pulumi.StringOutput { return v.IapAppEngineServiceIamMemberId }).(pulumi.StringOutput)
}

func (o IapAppEngineServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o IapAppEngineServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o IapAppEngineServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o IapAppEngineServiceIamMemberOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *IapAppEngineServiceIamMember) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IapAppEngineServiceIamMemberInput)(nil)).Elem(), &IapAppEngineServiceIamMember{})
	pulumi.RegisterOutputType(IapAppEngineServiceIamMemberOutput{})
}
