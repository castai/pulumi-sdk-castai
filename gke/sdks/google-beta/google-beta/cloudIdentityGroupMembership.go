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

type CloudIdentityGroupMembership struct {
	pulumi.CustomResourceState

	CloudIdentityGroupMembershipId pulumi.StringOutput `pulumi:"cloudIdentityGroupMembershipId"`
	// The time when the Membership was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The name of the Group to create this membership in.
	Group pulumi.StringOutput `pulumi:"group"`
	// EntityKey of the member.
	MemberKey CloudIdentityGroupMembershipMemberKeyPtrOutput `pulumi:"memberKey"`
	// The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// EntityKey of the member.
	PreferredMemberKey CloudIdentityGroupMembershipPreferredMemberKeyPtrOutput `pulumi:"preferredMemberKey"`
	// The MembershipRoles that apply to the Membership. Must not contain duplicate MembershipRoles with the same name.
	Roles    CloudIdentityGroupMembershipRoleArrayOutput   `pulumi:"roles"`
	Timeouts CloudIdentityGroupMembershipTimeoutsPtrOutput `pulumi:"timeouts"`
	// The type of the membership.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time when the Membership was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCloudIdentityGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewCloudIdentityGroupMembership(ctx *pulumi.Context,
	name string, args *CloudIdentityGroupMembershipArgs, opts ...pulumi.ResourceOption) (*CloudIdentityGroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudIdentityGroupMembership
	err = ctx.RegisterPackageResource("google-beta:index/cloudIdentityGroupMembership:CloudIdentityGroupMembership", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudIdentityGroupMembership gets an existing CloudIdentityGroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudIdentityGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudIdentityGroupMembershipState, opts ...pulumi.ResourceOption) (*CloudIdentityGroupMembership, error) {
	var resource CloudIdentityGroupMembership
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/cloudIdentityGroupMembership:CloudIdentityGroupMembership", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudIdentityGroupMembership resources.
type cloudIdentityGroupMembershipState struct {
	CloudIdentityGroupMembershipId *string `pulumi:"cloudIdentityGroupMembershipId"`
	// The time when the Membership was created.
	CreateTime *string `pulumi:"createTime"`
	// The name of the Group to create this membership in.
	Group *string `pulumi:"group"`
	// EntityKey of the member.
	MemberKey *CloudIdentityGroupMembershipMemberKey `pulumi:"memberKey"`
	// The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.
	Name *string `pulumi:"name"`
	// EntityKey of the member.
	PreferredMemberKey *CloudIdentityGroupMembershipPreferredMemberKey `pulumi:"preferredMemberKey"`
	// The MembershipRoles that apply to the Membership. Must not contain duplicate MembershipRoles with the same name.
	Roles    []CloudIdentityGroupMembershipRole    `pulumi:"roles"`
	Timeouts *CloudIdentityGroupMembershipTimeouts `pulumi:"timeouts"`
	// The type of the membership.
	Type *string `pulumi:"type"`
	// The time when the Membership was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type CloudIdentityGroupMembershipState struct {
	CloudIdentityGroupMembershipId pulumi.StringPtrInput
	// The time when the Membership was created.
	CreateTime pulumi.StringPtrInput
	// The name of the Group to create this membership in.
	Group pulumi.StringPtrInput
	// EntityKey of the member.
	MemberKey CloudIdentityGroupMembershipMemberKeyPtrInput
	// The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.
	Name pulumi.StringPtrInput
	// EntityKey of the member.
	PreferredMemberKey CloudIdentityGroupMembershipPreferredMemberKeyPtrInput
	// The MembershipRoles that apply to the Membership. Must not contain duplicate MembershipRoles with the same name.
	Roles    CloudIdentityGroupMembershipRoleArrayInput
	Timeouts CloudIdentityGroupMembershipTimeoutsPtrInput
	// The type of the membership.
	Type pulumi.StringPtrInput
	// The time when the Membership was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (CloudIdentityGroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIdentityGroupMembershipState)(nil)).Elem()
}

type cloudIdentityGroupMembershipArgs struct {
	CloudIdentityGroupMembershipId *string `pulumi:"cloudIdentityGroupMembershipId"`
	// The name of the Group to create this membership in.
	Group string `pulumi:"group"`
	// EntityKey of the member.
	MemberKey *CloudIdentityGroupMembershipMemberKey `pulumi:"memberKey"`
	// EntityKey of the member.
	PreferredMemberKey *CloudIdentityGroupMembershipPreferredMemberKey `pulumi:"preferredMemberKey"`
	// The MembershipRoles that apply to the Membership. Must not contain duplicate MembershipRoles with the same name.
	Roles    []CloudIdentityGroupMembershipRole    `pulumi:"roles"`
	Timeouts *CloudIdentityGroupMembershipTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a CloudIdentityGroupMembership resource.
type CloudIdentityGroupMembershipArgs struct {
	CloudIdentityGroupMembershipId pulumi.StringPtrInput
	// The name of the Group to create this membership in.
	Group pulumi.StringInput
	// EntityKey of the member.
	MemberKey CloudIdentityGroupMembershipMemberKeyPtrInput
	// EntityKey of the member.
	PreferredMemberKey CloudIdentityGroupMembershipPreferredMemberKeyPtrInput
	// The MembershipRoles that apply to the Membership. Must not contain duplicate MembershipRoles with the same name.
	Roles    CloudIdentityGroupMembershipRoleArrayInput
	Timeouts CloudIdentityGroupMembershipTimeoutsPtrInput
}

func (CloudIdentityGroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIdentityGroupMembershipArgs)(nil)).Elem()
}

type CloudIdentityGroupMembershipInput interface {
	pulumi.Input

	ToCloudIdentityGroupMembershipOutput() CloudIdentityGroupMembershipOutput
	ToCloudIdentityGroupMembershipOutputWithContext(ctx context.Context) CloudIdentityGroupMembershipOutput
}

func (*CloudIdentityGroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIdentityGroupMembership)(nil)).Elem()
}

func (i *CloudIdentityGroupMembership) ToCloudIdentityGroupMembershipOutput() CloudIdentityGroupMembershipOutput {
	return i.ToCloudIdentityGroupMembershipOutputWithContext(context.Background())
}

func (i *CloudIdentityGroupMembership) ToCloudIdentityGroupMembershipOutputWithContext(ctx context.Context) CloudIdentityGroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIdentityGroupMembershipOutput)
}

type CloudIdentityGroupMembershipOutput struct{ *pulumi.OutputState }

func (CloudIdentityGroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIdentityGroupMembership)(nil)).Elem()
}

func (o CloudIdentityGroupMembershipOutput) ToCloudIdentityGroupMembershipOutput() CloudIdentityGroupMembershipOutput {
	return o
}

func (o CloudIdentityGroupMembershipOutput) ToCloudIdentityGroupMembershipOutputWithContext(ctx context.Context) CloudIdentityGroupMembershipOutput {
	return o
}

func (o CloudIdentityGroupMembershipOutput) CloudIdentityGroupMembershipId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) pulumi.StringOutput { return v.CloudIdentityGroupMembershipId }).(pulumi.StringOutput)
}

// The time when the Membership was created.
func (o CloudIdentityGroupMembershipOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The name of the Group to create this membership in.
func (o CloudIdentityGroupMembershipOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// EntityKey of the member.
func (o CloudIdentityGroupMembershipOutput) MemberKey() CloudIdentityGroupMembershipMemberKeyPtrOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) CloudIdentityGroupMembershipMemberKeyPtrOutput {
		return v.MemberKey
	}).(CloudIdentityGroupMembershipMemberKeyPtrOutput)
}

// The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.
func (o CloudIdentityGroupMembershipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// EntityKey of the member.
func (o CloudIdentityGroupMembershipOutput) PreferredMemberKey() CloudIdentityGroupMembershipPreferredMemberKeyPtrOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) CloudIdentityGroupMembershipPreferredMemberKeyPtrOutput {
		return v.PreferredMemberKey
	}).(CloudIdentityGroupMembershipPreferredMemberKeyPtrOutput)
}

// The MembershipRoles that apply to the Membership. Must not contain duplicate MembershipRoles with the same name.
func (o CloudIdentityGroupMembershipOutput) Roles() CloudIdentityGroupMembershipRoleArrayOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) CloudIdentityGroupMembershipRoleArrayOutput { return v.Roles }).(CloudIdentityGroupMembershipRoleArrayOutput)
}

func (o CloudIdentityGroupMembershipOutput) Timeouts() CloudIdentityGroupMembershipTimeoutsPtrOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) CloudIdentityGroupMembershipTimeoutsPtrOutput { return v.Timeouts }).(CloudIdentityGroupMembershipTimeoutsPtrOutput)
}

// The type of the membership.
func (o CloudIdentityGroupMembershipOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time when the Membership was last updated.
func (o CloudIdentityGroupMembershipOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIdentityGroupMembership) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIdentityGroupMembershipInput)(nil)).Elem(), &CloudIdentityGroupMembership{})
	pulumi.RegisterOutputType(CloudIdentityGroupMembershipOutput{})
}
