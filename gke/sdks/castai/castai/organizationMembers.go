// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package castai

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrganizationMembers struct {
	pulumi.CustomResourceState

	// A list of email addresses corresponding to users who should be given member access to the organization.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// CAST AI organization ID.
	OrganizationId        pulumi.StringOutput `pulumi:"organizationId"`
	OrganizationMembersId pulumi.StringOutput `pulumi:"organizationMembersId"`
	// A list of email addresses corresponding to users who should be given owner access to the organization.
	Owners   pulumi.StringArrayOutput             `pulumi:"owners"`
	Timeouts OrganizationMembersTimeoutsPtrOutput `pulumi:"timeouts"`
	// A list of email addresses corresponding to users who should be given viewer access to the organization.
	Viewers pulumi.StringArrayOutput `pulumi:"viewers"`
}

// NewOrganizationMembers registers a new resource with the given unique name, arguments, and options.
func NewOrganizationMembers(ctx *pulumi.Context,
	name string, args *OrganizationMembersArgs, opts ...pulumi.ResourceOption) (*OrganizationMembers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource OrganizationMembers
	err = ctx.RegisterPackageResource("castai:index/organizationMembers:OrganizationMembers", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationMembers gets an existing OrganizationMembers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationMembers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationMembersState, opts ...pulumi.ResourceOption) (*OrganizationMembers, error) {
	var resource OrganizationMembers
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("castai:index/organizationMembers:OrganizationMembers", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationMembers resources.
type organizationMembersState struct {
	// A list of email addresses corresponding to users who should be given member access to the organization.
	Members []string `pulumi:"members"`
	// CAST AI organization ID.
	OrganizationId        *string `pulumi:"organizationId"`
	OrganizationMembersId *string `pulumi:"organizationMembersId"`
	// A list of email addresses corresponding to users who should be given owner access to the organization.
	Owners   []string                     `pulumi:"owners"`
	Timeouts *OrganizationMembersTimeouts `pulumi:"timeouts"`
	// A list of email addresses corresponding to users who should be given viewer access to the organization.
	Viewers []string `pulumi:"viewers"`
}

type OrganizationMembersState struct {
	// A list of email addresses corresponding to users who should be given member access to the organization.
	Members pulumi.StringArrayInput
	// CAST AI organization ID.
	OrganizationId        pulumi.StringPtrInput
	OrganizationMembersId pulumi.StringPtrInput
	// A list of email addresses corresponding to users who should be given owner access to the organization.
	Owners   pulumi.StringArrayInput
	Timeouts OrganizationMembersTimeoutsPtrInput
	// A list of email addresses corresponding to users who should be given viewer access to the organization.
	Viewers pulumi.StringArrayInput
}

func (OrganizationMembersState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationMembersState)(nil)).Elem()
}

type organizationMembersArgs struct {
	// A list of email addresses corresponding to users who should be given member access to the organization.
	Members []string `pulumi:"members"`
	// CAST AI organization ID.
	OrganizationId        string  `pulumi:"organizationId"`
	OrganizationMembersId *string `pulumi:"organizationMembersId"`
	// A list of email addresses corresponding to users who should be given owner access to the organization.
	Owners   []string                     `pulumi:"owners"`
	Timeouts *OrganizationMembersTimeouts `pulumi:"timeouts"`
	// A list of email addresses corresponding to users who should be given viewer access to the organization.
	Viewers []string `pulumi:"viewers"`
}

// The set of arguments for constructing a OrganizationMembers resource.
type OrganizationMembersArgs struct {
	// A list of email addresses corresponding to users who should be given member access to the organization.
	Members pulumi.StringArrayInput
	// CAST AI organization ID.
	OrganizationId        pulumi.StringInput
	OrganizationMembersId pulumi.StringPtrInput
	// A list of email addresses corresponding to users who should be given owner access to the organization.
	Owners   pulumi.StringArrayInput
	Timeouts OrganizationMembersTimeoutsPtrInput
	// A list of email addresses corresponding to users who should be given viewer access to the organization.
	Viewers pulumi.StringArrayInput
}

func (OrganizationMembersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationMembersArgs)(nil)).Elem()
}

type OrganizationMembersInput interface {
	pulumi.Input

	ToOrganizationMembersOutput() OrganizationMembersOutput
	ToOrganizationMembersOutputWithContext(ctx context.Context) OrganizationMembersOutput
}

func (*OrganizationMembers) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationMembers)(nil)).Elem()
}

func (i *OrganizationMembers) ToOrganizationMembersOutput() OrganizationMembersOutput {
	return i.ToOrganizationMembersOutputWithContext(context.Background())
}

func (i *OrganizationMembers) ToOrganizationMembersOutputWithContext(ctx context.Context) OrganizationMembersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationMembersOutput)
}

type OrganizationMembersOutput struct{ *pulumi.OutputState }

func (OrganizationMembersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationMembers)(nil)).Elem()
}

func (o OrganizationMembersOutput) ToOrganizationMembersOutput() OrganizationMembersOutput {
	return o
}

func (o OrganizationMembersOutput) ToOrganizationMembersOutputWithContext(ctx context.Context) OrganizationMembersOutput {
	return o
}

// A list of email addresses corresponding to users who should be given member access to the organization.
func (o OrganizationMembersOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationMembers) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// CAST AI organization ID.
func (o OrganizationMembersOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMembers) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o OrganizationMembersOutput) OrganizationMembersId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMembers) pulumi.StringOutput { return v.OrganizationMembersId }).(pulumi.StringOutput)
}

// A list of email addresses corresponding to users who should be given owner access to the organization.
func (o OrganizationMembersOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationMembers) pulumi.StringArrayOutput { return v.Owners }).(pulumi.StringArrayOutput)
}

func (o OrganizationMembersOutput) Timeouts() OrganizationMembersTimeoutsPtrOutput {
	return o.ApplyT(func(v *OrganizationMembers) OrganizationMembersTimeoutsPtrOutput { return v.Timeouts }).(OrganizationMembersTimeoutsPtrOutput)
}

// A list of email addresses corresponding to users who should be given viewer access to the organization.
func (o OrganizationMembersOutput) Viewers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationMembers) pulumi.StringArrayOutput { return v.Viewers }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMembersInput)(nil)).Elem(), &OrganizationMembers{})
	pulumi.RegisterOutputType(OrganizationMembersOutput{})
}
