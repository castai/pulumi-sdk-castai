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

type OrganizationIamMember struct {
	pulumi.CustomResourceState

	Condition OrganizationIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag      pulumi.StringOutput                     `pulumi:"etag"`
	Member    pulumi.StringOutput                     `pulumi:"member"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                   pulumi.StringOutput `pulumi:"orgId"`
	OrganizationIamMemberId pulumi.StringOutput `pulumi:"organizationIamMemberId"`
	Role                    pulumi.StringOutput `pulumi:"role"`
}

// NewOrganizationIamMember registers a new resource with the given unique name, arguments, and options.
func NewOrganizationIamMember(ctx *pulumi.Context,
	name string, args *OrganizationIamMemberArgs, opts ...pulumi.ResourceOption) (*OrganizationIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource OrganizationIamMember
	err = ctx.RegisterPackageResource("google-beta:index/organizationIamMember:OrganizationIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationIamMember gets an existing OrganizationIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationIamMemberState, opts ...pulumi.ResourceOption) (*OrganizationIamMember, error) {
	var resource OrganizationIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/organizationIamMember:OrganizationIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationIamMember resources.
type organizationIamMemberState struct {
	Condition *OrganizationIamMemberCondition `pulumi:"condition"`
	Etag      *string                         `pulumi:"etag"`
	Member    *string                         `pulumi:"member"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                   *string `pulumi:"orgId"`
	OrganizationIamMemberId *string `pulumi:"organizationIamMemberId"`
	Role                    *string `pulumi:"role"`
}

type OrganizationIamMemberState struct {
	Condition OrganizationIamMemberConditionPtrInput
	Etag      pulumi.StringPtrInput
	Member    pulumi.StringPtrInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                   pulumi.StringPtrInput
	OrganizationIamMemberId pulumi.StringPtrInput
	Role                    pulumi.StringPtrInput
}

func (OrganizationIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationIamMemberState)(nil)).Elem()
}

type organizationIamMemberArgs struct {
	Condition *OrganizationIamMemberCondition `pulumi:"condition"`
	Member    string                          `pulumi:"member"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                   string  `pulumi:"orgId"`
	OrganizationIamMemberId *string `pulumi:"organizationIamMemberId"`
	Role                    string  `pulumi:"role"`
}

// The set of arguments for constructing a OrganizationIamMember resource.
type OrganizationIamMemberArgs struct {
	Condition OrganizationIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                   pulumi.StringInput
	OrganizationIamMemberId pulumi.StringPtrInput
	Role                    pulumi.StringInput
}

func (OrganizationIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationIamMemberArgs)(nil)).Elem()
}

type OrganizationIamMemberInput interface {
	pulumi.Input

	ToOrganizationIamMemberOutput() OrganizationIamMemberOutput
	ToOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationIamMemberOutput
}

func (*OrganizationIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationIamMember)(nil)).Elem()
}

func (i *OrganizationIamMember) ToOrganizationIamMemberOutput() OrganizationIamMemberOutput {
	return i.ToOrganizationIamMemberOutputWithContext(context.Background())
}

func (i *OrganizationIamMember) ToOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationIamMemberOutput)
}

type OrganizationIamMemberOutput struct{ *pulumi.OutputState }

func (OrganizationIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationIamMember)(nil)).Elem()
}

func (o OrganizationIamMemberOutput) ToOrganizationIamMemberOutput() OrganizationIamMemberOutput {
	return o
}

func (o OrganizationIamMemberOutput) ToOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationIamMemberOutput {
	return o
}

func (o OrganizationIamMemberOutput) Condition() OrganizationIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *OrganizationIamMember) OrganizationIamMemberConditionPtrOutput { return v.Condition }).(OrganizationIamMemberConditionPtrOutput)
}

func (o OrganizationIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o OrganizationIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The numeric ID of the organization in which you want to manage the audit logging config.
func (o OrganizationIamMemberOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamMember) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

func (o OrganizationIamMemberOutput) OrganizationIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamMember) pulumi.StringOutput { return v.OrganizationIamMemberId }).(pulumi.StringOutput)
}

func (o OrganizationIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationIamMemberInput)(nil)).Elem(), &OrganizationIamMember{})
	pulumi.RegisterOutputType(OrganizationIamMemberOutput{})
}
