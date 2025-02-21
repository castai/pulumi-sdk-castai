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

type OsConfigOsPolicyAssignment struct {
	pulumi.CustomResourceState

	// Output only. Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS
	// policies from this revision. For a given OS policy assignment, there is only one revision with a value of 'true' for
	// this field.
	Baseline pulumi.BoolOutput `pulumi:"baseline"`
	// Output only. Indicates that this revision deletes the OS policy assignment.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Filter to select VMs.
	InstanceFilter OsConfigOsPolicyAssignmentInstanceFilterOutput `pulumi:"instanceFilter"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name                         pulumi.StringOutput `pulumi:"name"`
	OsConfigOsPolicyAssignmentId pulumi.StringOutput `pulumi:"osConfigOsPolicyAssignmentId"`
	// List of OS policies to be applied to the VMs.
	OsPolicies OsConfigOsPolicyAssignmentOsPolicyArrayOutput `pulumi:"osPolicies"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. Indicates that reconciliation is in progress for the revision. This value is 'true' when the
	// 'rollout_state' is one of: * IN_PROGRESS * CANCELLING
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Output only. The timestamp that the revision was created.
	RevisionCreateTime pulumi.StringOutput `pulumi:"revisionCreateTime"`
	// Output only. The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy
	// assignment
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is
	// created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: -
	// instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OsConfigOsPolicyAssignmentRolloutOutput `pulumi:"rollout"`
	// Output only. OS policy assignment rollout state
	RolloutState pulumi.StringOutput `pulumi:"rolloutState"`
	// Set to true to skip awaiting rollout during resource creation and update.
	SkipAwaitRollout pulumi.BoolPtrOutput                        `pulumi:"skipAwaitRollout"`
	Timeouts         OsConfigOsPolicyAssignmentTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. Server generated unique id for the OS policy assignment resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewOsConfigOsPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewOsConfigOsPolicyAssignment(ctx *pulumi.Context,
	name string, args *OsConfigOsPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*OsConfigOsPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceFilter == nil {
		return nil, errors.New("invalid value for required argument 'InstanceFilter'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.OsPolicies == nil {
		return nil, errors.New("invalid value for required argument 'OsPolicies'")
	}
	if args.Rollout == nil {
		return nil, errors.New("invalid value for required argument 'Rollout'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource OsConfigOsPolicyAssignment
	err = ctx.RegisterPackageResource("google:index/osConfigOsPolicyAssignment:OsConfigOsPolicyAssignment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOsConfigOsPolicyAssignment gets an existing OsConfigOsPolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOsConfigOsPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OsConfigOsPolicyAssignmentState, opts ...pulumi.ResourceOption) (*OsConfigOsPolicyAssignment, error) {
	var resource OsConfigOsPolicyAssignment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/osConfigOsPolicyAssignment:OsConfigOsPolicyAssignment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OsConfigOsPolicyAssignment resources.
type osConfigOsPolicyAssignmentState struct {
	// Output only. Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS
	// policies from this revision. For a given OS policy assignment, there is only one revision with a value of 'true' for
	// this field.
	Baseline *bool `pulumi:"baseline"`
	// Output only. Indicates that this revision deletes the OS policy assignment.
	Deleted *bool `pulumi:"deleted"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	Etag *string `pulumi:"etag"`
	// Filter to select VMs.
	InstanceFilter *OsConfigOsPolicyAssignmentInstanceFilter `pulumi:"instanceFilter"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Resource name.
	Name                         *string `pulumi:"name"`
	OsConfigOsPolicyAssignmentId *string `pulumi:"osConfigOsPolicyAssignmentId"`
	// List of OS policies to be applied to the VMs.
	OsPolicies []OsConfigOsPolicyAssignmentOsPolicy `pulumi:"osPolicies"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. Indicates that reconciliation is in progress for the revision. This value is 'true' when the
	// 'rollout_state' is one of: * IN_PROGRESS * CANCELLING
	Reconciling *bool `pulumi:"reconciling"`
	// Output only. The timestamp that the revision was created.
	RevisionCreateTime *string `pulumi:"revisionCreateTime"`
	// Output only. The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy
	// assignment
	RevisionId *string `pulumi:"revisionId"`
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is
	// created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: -
	// instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout *OsConfigOsPolicyAssignmentRollout `pulumi:"rollout"`
	// Output only. OS policy assignment rollout state
	RolloutState *string `pulumi:"rolloutState"`
	// Set to true to skip awaiting rollout during resource creation and update.
	SkipAwaitRollout *bool                               `pulumi:"skipAwaitRollout"`
	Timeouts         *OsConfigOsPolicyAssignmentTimeouts `pulumi:"timeouts"`
	// Output only. Server generated unique id for the OS policy assignment resource.
	Uid *string `pulumi:"uid"`
}

type OsConfigOsPolicyAssignmentState struct {
	// Output only. Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS
	// policies from this revision. For a given OS policy assignment, there is only one revision with a value of 'true' for
	// this field.
	Baseline pulumi.BoolPtrInput
	// Output only. Indicates that this revision deletes the OS policy assignment.
	Deleted pulumi.BoolPtrInput
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringPtrInput
	// Filter to select VMs.
	InstanceFilter OsConfigOsPolicyAssignmentInstanceFilterPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Resource name.
	Name                         pulumi.StringPtrInput
	OsConfigOsPolicyAssignmentId pulumi.StringPtrInput
	// List of OS policies to be applied to the VMs.
	OsPolicies OsConfigOsPolicyAssignmentOsPolicyArrayInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. Indicates that reconciliation is in progress for the revision. This value is 'true' when the
	// 'rollout_state' is one of: * IN_PROGRESS * CANCELLING
	Reconciling pulumi.BoolPtrInput
	// Output only. The timestamp that the revision was created.
	RevisionCreateTime pulumi.StringPtrInput
	// Output only. The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy
	// assignment
	RevisionId pulumi.StringPtrInput
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is
	// created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: -
	// instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OsConfigOsPolicyAssignmentRolloutPtrInput
	// Output only. OS policy assignment rollout state
	RolloutState pulumi.StringPtrInput
	// Set to true to skip awaiting rollout during resource creation and update.
	SkipAwaitRollout pulumi.BoolPtrInput
	Timeouts         OsConfigOsPolicyAssignmentTimeoutsPtrInput
	// Output only. Server generated unique id for the OS policy assignment resource.
	Uid pulumi.StringPtrInput
}

func (OsConfigOsPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*osConfigOsPolicyAssignmentState)(nil)).Elem()
}

type osConfigOsPolicyAssignmentArgs struct {
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// Filter to select VMs.
	InstanceFilter OsConfigOsPolicyAssignmentInstanceFilter `pulumi:"instanceFilter"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Resource name.
	Name                         *string `pulumi:"name"`
	OsConfigOsPolicyAssignmentId *string `pulumi:"osConfigOsPolicyAssignmentId"`
	// List of OS policies to be applied to the VMs.
	OsPolicies []OsConfigOsPolicyAssignmentOsPolicy `pulumi:"osPolicies"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is
	// created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: -
	// instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OsConfigOsPolicyAssignmentRollout `pulumi:"rollout"`
	// Set to true to skip awaiting rollout during resource creation and update.
	SkipAwaitRollout *bool                               `pulumi:"skipAwaitRollout"`
	Timeouts         *OsConfigOsPolicyAssignmentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a OsConfigOsPolicyAssignment resource.
type OsConfigOsPolicyAssignmentArgs struct {
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// Filter to select VMs.
	InstanceFilter OsConfigOsPolicyAssignmentInstanceFilterInput
	// The location for the resource
	Location pulumi.StringInput
	// Resource name.
	Name                         pulumi.StringPtrInput
	OsConfigOsPolicyAssignmentId pulumi.StringPtrInput
	// List of OS policies to be applied to the VMs.
	OsPolicies OsConfigOsPolicyAssignmentOsPolicyArrayInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is
	// created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: -
	// instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OsConfigOsPolicyAssignmentRolloutInput
	// Set to true to skip awaiting rollout during resource creation and update.
	SkipAwaitRollout pulumi.BoolPtrInput
	Timeouts         OsConfigOsPolicyAssignmentTimeoutsPtrInput
}

func (OsConfigOsPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*osConfigOsPolicyAssignmentArgs)(nil)).Elem()
}

type OsConfigOsPolicyAssignmentInput interface {
	pulumi.Input

	ToOsConfigOsPolicyAssignmentOutput() OsConfigOsPolicyAssignmentOutput
	ToOsConfigOsPolicyAssignmentOutputWithContext(ctx context.Context) OsConfigOsPolicyAssignmentOutput
}

func (*OsConfigOsPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**OsConfigOsPolicyAssignment)(nil)).Elem()
}

func (i *OsConfigOsPolicyAssignment) ToOsConfigOsPolicyAssignmentOutput() OsConfigOsPolicyAssignmentOutput {
	return i.ToOsConfigOsPolicyAssignmentOutputWithContext(context.Background())
}

func (i *OsConfigOsPolicyAssignment) ToOsConfigOsPolicyAssignmentOutputWithContext(ctx context.Context) OsConfigOsPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsConfigOsPolicyAssignmentOutput)
}

type OsConfigOsPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (OsConfigOsPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsConfigOsPolicyAssignment)(nil)).Elem()
}

func (o OsConfigOsPolicyAssignmentOutput) ToOsConfigOsPolicyAssignmentOutput() OsConfigOsPolicyAssignmentOutput {
	return o
}

func (o OsConfigOsPolicyAssignmentOutput) ToOsConfigOsPolicyAssignmentOutputWithContext(ctx context.Context) OsConfigOsPolicyAssignmentOutput {
	return o
}

// Output only. Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS
// policies from this revision. For a given OS policy assignment, there is only one revision with a value of 'true' for
// this field.
func (o OsConfigOsPolicyAssignmentOutput) Baseline() pulumi.BoolOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.BoolOutput { return v.Baseline }).(pulumi.BoolOutput)
}

// Output only. Indicates that this revision deletes the OS policy assignment.
func (o OsConfigOsPolicyAssignmentOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.BoolOutput { return v.Deleted }).(pulumi.BoolOutput)
}

// OS policy assignment description. Length of the description is limited to 1024 characters.
func (o OsConfigOsPolicyAssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
func (o OsConfigOsPolicyAssignmentOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Filter to select VMs.
func (o OsConfigOsPolicyAssignmentOutput) InstanceFilter() OsConfigOsPolicyAssignmentInstanceFilterOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) OsConfigOsPolicyAssignmentInstanceFilterOutput {
		return v.InstanceFilter
	}).(OsConfigOsPolicyAssignmentInstanceFilterOutput)
}

// The location for the resource
func (o OsConfigOsPolicyAssignmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o OsConfigOsPolicyAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OsConfigOsPolicyAssignmentOutput) OsConfigOsPolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.OsConfigOsPolicyAssignmentId }).(pulumi.StringOutput)
}

// List of OS policies to be applied to the VMs.
func (o OsConfigOsPolicyAssignmentOutput) OsPolicies() OsConfigOsPolicyAssignmentOsPolicyArrayOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) OsConfigOsPolicyAssignmentOsPolicyArrayOutput { return v.OsPolicies }).(OsConfigOsPolicyAssignmentOsPolicyArrayOutput)
}

// The project for the resource
func (o OsConfigOsPolicyAssignmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Indicates that reconciliation is in progress for the revision. This value is 'true' when the
// 'rollout_state' is one of: * IN_PROGRESS * CANCELLING
func (o OsConfigOsPolicyAssignmentOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Output only. The timestamp that the revision was created.
func (o OsConfigOsPolicyAssignmentOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// Output only. The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy
// assignment
func (o OsConfigOsPolicyAssignmentOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is
// created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: -
// instance_filter - os_policies 3) OSPolicyAssignment is deleted.
func (o OsConfigOsPolicyAssignmentOutput) Rollout() OsConfigOsPolicyAssignmentRolloutOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) OsConfigOsPolicyAssignmentRolloutOutput { return v.Rollout }).(OsConfigOsPolicyAssignmentRolloutOutput)
}

// Output only. OS policy assignment rollout state
func (o OsConfigOsPolicyAssignmentOutput) RolloutState() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.RolloutState }).(pulumi.StringOutput)
}

// Set to true to skip awaiting rollout during resource creation and update.
func (o OsConfigOsPolicyAssignmentOutput) SkipAwaitRollout() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.BoolPtrOutput { return v.SkipAwaitRollout }).(pulumi.BoolPtrOutput)
}

func (o OsConfigOsPolicyAssignmentOutput) Timeouts() OsConfigOsPolicyAssignmentTimeoutsPtrOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) OsConfigOsPolicyAssignmentTimeoutsPtrOutput { return v.Timeouts }).(OsConfigOsPolicyAssignmentTimeoutsPtrOutput)
}

// Output only. Server generated unique id for the OS policy assignment resource.
func (o OsConfigOsPolicyAssignmentOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *OsConfigOsPolicyAssignment) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OsConfigOsPolicyAssignmentInput)(nil)).Elem(), &OsConfigOsPolicyAssignment{})
	pulumi.RegisterOutputType(OsConfigOsPolicyAssignmentOutput{})
}
