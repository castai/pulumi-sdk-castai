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

type OrgPolicyPolicy struct {
	pulumi.CustomResourceState

	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future
	// resources if it's enforced.
	DryRunSpec OrgPolicyPolicyDryRunSpecPtrOutput `pulumi:"dryRunSpec"`
	// Optional. An opaque tag indicating the current state of the policy, used for concurrency control. This 'etag' is
	// computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the
	// constraint which this Policy configures: * 'projects/{project_number}/policies/{constraint_name}' *
	// 'folders/{folder_id}/policies/{constraint_name}' * 'organizations/{organization_id}/policies/{constraint_name}' For
	// example, "projects/123/policies/compute.disableSerialPortAccess". Note:
	// 'projects/{project_id}/policies/{constraint_name}' is also an acceptable name for API requests, but responses will
	// return the name using the equivalent project number.
	Name              pulumi.StringOutput `pulumi:"name"`
	OrgPolicyPolicyId pulumi.StringOutput `pulumi:"orgPolicyPolicyId"`
	// The parent of the resource.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Basic information about the Organization Policy.
	Spec     OrgPolicyPolicySpecPtrOutput     `pulumi:"spec"`
	Timeouts OrgPolicyPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewOrgPolicyPolicy registers a new resource with the given unique name, arguments, and options.
func NewOrgPolicyPolicy(ctx *pulumi.Context,
	name string, args *OrgPolicyPolicyArgs, opts ...pulumi.ResourceOption) (*OrgPolicyPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource OrgPolicyPolicy
	err = ctx.RegisterPackageResource("google-beta:index/orgPolicyPolicy:OrgPolicyPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgPolicyPolicy gets an existing OrgPolicyPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgPolicyPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgPolicyPolicyState, opts ...pulumi.ResourceOption) (*OrgPolicyPolicy, error) {
	var resource OrgPolicyPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/orgPolicyPolicy:OrgPolicyPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgPolicyPolicy resources.
type orgPolicyPolicyState struct {
	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future
	// resources if it's enforced.
	DryRunSpec *OrgPolicyPolicyDryRunSpec `pulumi:"dryRunSpec"`
	// Optional. An opaque tag indicating the current state of the policy, used for concurrency control. This 'etag' is
	// computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the
	// constraint which this Policy configures: * 'projects/{project_number}/policies/{constraint_name}' *
	// 'folders/{folder_id}/policies/{constraint_name}' * 'organizations/{organization_id}/policies/{constraint_name}' For
	// example, "projects/123/policies/compute.disableSerialPortAccess". Note:
	// 'projects/{project_id}/policies/{constraint_name}' is also an acceptable name for API requests, but responses will
	// return the name using the equivalent project number.
	Name              *string `pulumi:"name"`
	OrgPolicyPolicyId *string `pulumi:"orgPolicyPolicyId"`
	// The parent of the resource.
	Parent *string `pulumi:"parent"`
	// Basic information about the Organization Policy.
	Spec     *OrgPolicyPolicySpec     `pulumi:"spec"`
	Timeouts *OrgPolicyPolicyTimeouts `pulumi:"timeouts"`
}

type OrgPolicyPolicyState struct {
	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future
	// resources if it's enforced.
	DryRunSpec OrgPolicyPolicyDryRunSpecPtrInput
	// Optional. An opaque tag indicating the current state of the policy, used for concurrency control. This 'etag' is
	// computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the
	// constraint which this Policy configures: * 'projects/{project_number}/policies/{constraint_name}' *
	// 'folders/{folder_id}/policies/{constraint_name}' * 'organizations/{organization_id}/policies/{constraint_name}' For
	// example, "projects/123/policies/compute.disableSerialPortAccess". Note:
	// 'projects/{project_id}/policies/{constraint_name}' is also an acceptable name for API requests, but responses will
	// return the name using the equivalent project number.
	Name              pulumi.StringPtrInput
	OrgPolicyPolicyId pulumi.StringPtrInput
	// The parent of the resource.
	Parent pulumi.StringPtrInput
	// Basic information about the Organization Policy.
	Spec     OrgPolicyPolicySpecPtrInput
	Timeouts OrgPolicyPolicyTimeoutsPtrInput
}

func (OrgPolicyPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgPolicyPolicyState)(nil)).Elem()
}

type orgPolicyPolicyArgs struct {
	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future
	// resources if it's enforced.
	DryRunSpec *OrgPolicyPolicyDryRunSpec `pulumi:"dryRunSpec"`
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the
	// constraint which this Policy configures: * 'projects/{project_number}/policies/{constraint_name}' *
	// 'folders/{folder_id}/policies/{constraint_name}' * 'organizations/{organization_id}/policies/{constraint_name}' For
	// example, "projects/123/policies/compute.disableSerialPortAccess". Note:
	// 'projects/{project_id}/policies/{constraint_name}' is also an acceptable name for API requests, but responses will
	// return the name using the equivalent project number.
	Name              *string `pulumi:"name"`
	OrgPolicyPolicyId *string `pulumi:"orgPolicyPolicyId"`
	// The parent of the resource.
	Parent string `pulumi:"parent"`
	// Basic information about the Organization Policy.
	Spec     *OrgPolicyPolicySpec     `pulumi:"spec"`
	Timeouts *OrgPolicyPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a OrgPolicyPolicy resource.
type OrgPolicyPolicyArgs struct {
	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future
	// resources if it's enforced.
	DryRunSpec OrgPolicyPolicyDryRunSpecPtrInput
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the
	// constraint which this Policy configures: * 'projects/{project_number}/policies/{constraint_name}' *
	// 'folders/{folder_id}/policies/{constraint_name}' * 'organizations/{organization_id}/policies/{constraint_name}' For
	// example, "projects/123/policies/compute.disableSerialPortAccess". Note:
	// 'projects/{project_id}/policies/{constraint_name}' is also an acceptable name for API requests, but responses will
	// return the name using the equivalent project number.
	Name              pulumi.StringPtrInput
	OrgPolicyPolicyId pulumi.StringPtrInput
	// The parent of the resource.
	Parent pulumi.StringInput
	// Basic information about the Organization Policy.
	Spec     OrgPolicyPolicySpecPtrInput
	Timeouts OrgPolicyPolicyTimeoutsPtrInput
}

func (OrgPolicyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgPolicyPolicyArgs)(nil)).Elem()
}

type OrgPolicyPolicyInput interface {
	pulumi.Input

	ToOrgPolicyPolicyOutput() OrgPolicyPolicyOutput
	ToOrgPolicyPolicyOutputWithContext(ctx context.Context) OrgPolicyPolicyOutput
}

func (*OrgPolicyPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgPolicyPolicy)(nil)).Elem()
}

func (i *OrgPolicyPolicy) ToOrgPolicyPolicyOutput() OrgPolicyPolicyOutput {
	return i.ToOrgPolicyPolicyOutputWithContext(context.Background())
}

func (i *OrgPolicyPolicy) ToOrgPolicyPolicyOutputWithContext(ctx context.Context) OrgPolicyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgPolicyPolicyOutput)
}

type OrgPolicyPolicyOutput struct{ *pulumi.OutputState }

func (OrgPolicyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgPolicyPolicy)(nil)).Elem()
}

func (o OrgPolicyPolicyOutput) ToOrgPolicyPolicyOutput() OrgPolicyPolicyOutput {
	return o
}

func (o OrgPolicyPolicyOutput) ToOrgPolicyPolicyOutputWithContext(ctx context.Context) OrgPolicyPolicyOutput {
	return o
}

// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future
// resources if it's enforced.
func (o OrgPolicyPolicyOutput) DryRunSpec() OrgPolicyPolicyDryRunSpecPtrOutput {
	return o.ApplyT(func(v *OrgPolicyPolicy) OrgPolicyPolicyDryRunSpecPtrOutput { return v.DryRunSpec }).(OrgPolicyPolicyDryRunSpecPtrOutput)
}

// Optional. An opaque tag indicating the current state of the policy, used for concurrency control. This 'etag' is
// computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the
// client has an up-to-date value before proceeding.
func (o OrgPolicyPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgPolicyPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the
// constraint which this Policy configures: * 'projects/{project_number}/policies/{constraint_name}' *
// 'folders/{folder_id}/policies/{constraint_name}' * 'organizations/{organization_id}/policies/{constraint_name}' For
// example, "projects/123/policies/compute.disableSerialPortAccess". Note:
// 'projects/{project_id}/policies/{constraint_name}' is also an acceptable name for API requests, but responses will
// return the name using the equivalent project number.
func (o OrgPolicyPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgPolicyPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrgPolicyPolicyOutput) OrgPolicyPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgPolicyPolicy) pulumi.StringOutput { return v.OrgPolicyPolicyId }).(pulumi.StringOutput)
}

// The parent of the resource.
func (o OrgPolicyPolicyOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgPolicyPolicy) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Basic information about the Organization Policy.
func (o OrgPolicyPolicyOutput) Spec() OrgPolicyPolicySpecPtrOutput {
	return o.ApplyT(func(v *OrgPolicyPolicy) OrgPolicyPolicySpecPtrOutput { return v.Spec }).(OrgPolicyPolicySpecPtrOutput)
}

func (o OrgPolicyPolicyOutput) Timeouts() OrgPolicyPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *OrgPolicyPolicy) OrgPolicyPolicyTimeoutsPtrOutput { return v.Timeouts }).(OrgPolicyPolicyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgPolicyPolicyInput)(nil)).Elem(), &OrgPolicyPolicy{})
	pulumi.RegisterOutputType(OrgPolicyPolicyOutput{})
}
