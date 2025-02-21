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

type SccSourceIamMember struct {
	pulumi.CustomResourceState

	Condition            SccSourceIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                 pulumi.StringOutput                  `pulumi:"etag"`
	Member               pulumi.StringOutput                  `pulumi:"member"`
	Organization         pulumi.StringOutput                  `pulumi:"organization"`
	Role                 pulumi.StringOutput                  `pulumi:"role"`
	SccSourceIamMemberId pulumi.StringOutput                  `pulumi:"sccSourceIamMemberId"`
	Source               pulumi.StringOutput                  `pulumi:"source"`
}

// NewSccSourceIamMember registers a new resource with the given unique name, arguments, and options.
func NewSccSourceIamMember(ctx *pulumi.Context,
	name string, args *SccSourceIamMemberArgs, opts ...pulumi.ResourceOption) (*SccSourceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccSourceIamMember
	err = ctx.RegisterPackageResource("google-beta:index/sccSourceIamMember:SccSourceIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccSourceIamMember gets an existing SccSourceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccSourceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccSourceIamMemberState, opts ...pulumi.ResourceOption) (*SccSourceIamMember, error) {
	var resource SccSourceIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/sccSourceIamMember:SccSourceIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccSourceIamMember resources.
type sccSourceIamMemberState struct {
	Condition            *SccSourceIamMemberCondition `pulumi:"condition"`
	Etag                 *string                      `pulumi:"etag"`
	Member               *string                      `pulumi:"member"`
	Organization         *string                      `pulumi:"organization"`
	Role                 *string                      `pulumi:"role"`
	SccSourceIamMemberId *string                      `pulumi:"sccSourceIamMemberId"`
	Source               *string                      `pulumi:"source"`
}

type SccSourceIamMemberState struct {
	Condition            SccSourceIamMemberConditionPtrInput
	Etag                 pulumi.StringPtrInput
	Member               pulumi.StringPtrInput
	Organization         pulumi.StringPtrInput
	Role                 pulumi.StringPtrInput
	SccSourceIamMemberId pulumi.StringPtrInput
	Source               pulumi.StringPtrInput
}

func (SccSourceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccSourceIamMemberState)(nil)).Elem()
}

type sccSourceIamMemberArgs struct {
	Condition            *SccSourceIamMemberCondition `pulumi:"condition"`
	Member               string                       `pulumi:"member"`
	Organization         string                       `pulumi:"organization"`
	Role                 string                       `pulumi:"role"`
	SccSourceIamMemberId *string                      `pulumi:"sccSourceIamMemberId"`
	Source               string                       `pulumi:"source"`
}

// The set of arguments for constructing a SccSourceIamMember resource.
type SccSourceIamMemberArgs struct {
	Condition            SccSourceIamMemberConditionPtrInput
	Member               pulumi.StringInput
	Organization         pulumi.StringInput
	Role                 pulumi.StringInput
	SccSourceIamMemberId pulumi.StringPtrInput
	Source               pulumi.StringInput
}

func (SccSourceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccSourceIamMemberArgs)(nil)).Elem()
}

type SccSourceIamMemberInput interface {
	pulumi.Input

	ToSccSourceIamMemberOutput() SccSourceIamMemberOutput
	ToSccSourceIamMemberOutputWithContext(ctx context.Context) SccSourceIamMemberOutput
}

func (*SccSourceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**SccSourceIamMember)(nil)).Elem()
}

func (i *SccSourceIamMember) ToSccSourceIamMemberOutput() SccSourceIamMemberOutput {
	return i.ToSccSourceIamMemberOutputWithContext(context.Background())
}

func (i *SccSourceIamMember) ToSccSourceIamMemberOutputWithContext(ctx context.Context) SccSourceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccSourceIamMemberOutput)
}

type SccSourceIamMemberOutput struct{ *pulumi.OutputState }

func (SccSourceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccSourceIamMember)(nil)).Elem()
}

func (o SccSourceIamMemberOutput) ToSccSourceIamMemberOutput() SccSourceIamMemberOutput {
	return o
}

func (o SccSourceIamMemberOutput) ToSccSourceIamMemberOutputWithContext(ctx context.Context) SccSourceIamMemberOutput {
	return o
}

func (o SccSourceIamMemberOutput) Condition() SccSourceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *SccSourceIamMember) SccSourceIamMemberConditionPtrOutput { return v.Condition }).(SccSourceIamMemberConditionPtrOutput)
}

func (o SccSourceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SccSourceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SccSourceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *SccSourceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o SccSourceIamMemberOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccSourceIamMember) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

func (o SccSourceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SccSourceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o SccSourceIamMemberOutput) SccSourceIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccSourceIamMember) pulumi.StringOutput { return v.SccSourceIamMemberId }).(pulumi.StringOutput)
}

func (o SccSourceIamMemberOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *SccSourceIamMember) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccSourceIamMemberInput)(nil)).Elem(), &SccSourceIamMember{})
	pulumi.RegisterOutputType(SccSourceIamMemberOutput{})
}
