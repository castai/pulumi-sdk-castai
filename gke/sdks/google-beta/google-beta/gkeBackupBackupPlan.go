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

type GkeBackupBackupPlan struct {
	pulumi.CustomResourceState

	// Defines the configuration of Backups created via this BackupPlan.
	BackupConfig GkeBackupBackupPlanBackupConfigPtrOutput `pulumi:"backupConfig"`
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	BackupSchedule GkeBackupBackupPlanBackupSchedulePtrOutput `pulumi:"backupSchedule"`
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such
	// that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any
	// new Backups from being created via this BackupPlan (including scheduled Backups).
	Deactivated pulumi.BoolOutput `pulumi:"deactivated"`
	// User specified descriptive string for this BackupPlan.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup plan from
	// overwriting each other. It is strongly suggested that systems make use of the 'etag' in the read-modify-write cycle to
	// perform BackupPlan updates in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
	// and systems are expected to put that etag in the request to backupPlans.patch or backupPlans.delete to ensure that their
	// change will be applied to the same version of the resource.
	Etag                  pulumi.StringOutput `pulumi:"etag"`
	GkeBackupBackupPlanId pulumi.StringOutput `pulumi:"gkeBackupBackupPlanId"`
	// Description: A set of custom labels supplied by the user. A list of key->value pairs. Example: { "name": "wrench",
	// "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The region of the Backup Plan.
	Location pulumi.StringOutput `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
	ProtectedPodCount pulumi.Float64Output `pulumi:"protectedPodCount"`
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	RetentionPolicy GkeBackupBackupPlanRetentionPolicyPtrOutput `pulumi:"retentionPolicy"`
	// The State of the BackupPlan.
	State pulumi.StringOutput `pulumi:"state"`
	// Detailed description of why BackupPlan is in its current state.
	StateReason pulumi.StringOutput `pulumi:"stateReason"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput               `pulumi:"terraformLabels"`
	Timeouts        GkeBackupBackupPlanTimeoutsPtrOutput `pulumi:"timeouts"`
	// Server generated, unique identifier of UUID format.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewGkeBackupBackupPlan registers a new resource with the given unique name, arguments, and options.
func NewGkeBackupBackupPlan(ctx *pulumi.Context,
	name string, args *GkeBackupBackupPlanArgs, opts ...pulumi.ResourceOption) (*GkeBackupBackupPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GkeBackupBackupPlan
	err = ctx.RegisterPackageResource("google-beta:index/gkeBackupBackupPlan:GkeBackupBackupPlan", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeBackupBackupPlan gets an existing GkeBackupBackupPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeBackupBackupPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeBackupBackupPlanState, opts ...pulumi.ResourceOption) (*GkeBackupBackupPlan, error) {
	var resource GkeBackupBackupPlan
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/gkeBackupBackupPlan:GkeBackupBackupPlan", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeBackupBackupPlan resources.
type gkeBackupBackupPlanState struct {
	// Defines the configuration of Backups created via this BackupPlan.
	BackupConfig *GkeBackupBackupPlanBackupConfig `pulumi:"backupConfig"`
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	BackupSchedule *GkeBackupBackupPlanBackupSchedule `pulumi:"backupSchedule"`
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster *string `pulumi:"cluster"`
	// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such
	// that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any
	// new Backups from being created via this BackupPlan (including scheduled Backups).
	Deactivated *bool `pulumi:"deactivated"`
	// User specified descriptive string for this BackupPlan.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup plan from
	// overwriting each other. It is strongly suggested that systems make use of the 'etag' in the read-modify-write cycle to
	// perform BackupPlan updates in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
	// and systems are expected to put that etag in the request to backupPlans.patch or backupPlans.delete to ensure that their
	// change will be applied to the same version of the resource.
	Etag                  *string `pulumi:"etag"`
	GkeBackupBackupPlanId *string `pulumi:"gkeBackupBackupPlanId"`
	// Description: A set of custom labels supplied by the user. A list of key->value pairs. Example: { "name": "wrench",
	// "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The region of the Backup Plan.
	Location *string `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
	ProtectedPodCount *float64 `pulumi:"protectedPodCount"`
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	RetentionPolicy *GkeBackupBackupPlanRetentionPolicy `pulumi:"retentionPolicy"`
	// The State of the BackupPlan.
	State *string `pulumi:"state"`
	// Detailed description of why BackupPlan is in its current state.
	StateReason *string `pulumi:"stateReason"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string            `pulumi:"terraformLabels"`
	Timeouts        *GkeBackupBackupPlanTimeouts `pulumi:"timeouts"`
	// Server generated, unique identifier of UUID format.
	Uid *string `pulumi:"uid"`
}

type GkeBackupBackupPlanState struct {
	// Defines the configuration of Backups created via this BackupPlan.
	BackupConfig GkeBackupBackupPlanBackupConfigPtrInput
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	BackupSchedule GkeBackupBackupPlanBackupSchedulePtrInput
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster pulumi.StringPtrInput
	// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such
	// that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any
	// new Backups from being created via this BackupPlan (including scheduled Backups).
	Deactivated pulumi.BoolPtrInput
	// User specified descriptive string for this BackupPlan.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup plan from
	// overwriting each other. It is strongly suggested that systems make use of the 'etag' in the read-modify-write cycle to
	// perform BackupPlan updates in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
	// and systems are expected to put that etag in the request to backupPlans.patch or backupPlans.delete to ensure that their
	// change will be applied to the same version of the resource.
	Etag                  pulumi.StringPtrInput
	GkeBackupBackupPlanId pulumi.StringPtrInput
	// Description: A set of custom labels supplied by the user. A list of key->value pairs. Example: { "name": "wrench",
	// "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The region of the Backup Plan.
	Location pulumi.StringPtrInput
	// The full name of the BackupPlan Resource.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
	ProtectedPodCount pulumi.Float64PtrInput
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	RetentionPolicy GkeBackupBackupPlanRetentionPolicyPtrInput
	// The State of the BackupPlan.
	State pulumi.StringPtrInput
	// Detailed description of why BackupPlan is in its current state.
	StateReason pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        GkeBackupBackupPlanTimeoutsPtrInput
	// Server generated, unique identifier of UUID format.
	Uid pulumi.StringPtrInput
}

func (GkeBackupBackupPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeBackupBackupPlanState)(nil)).Elem()
}

type gkeBackupBackupPlanArgs struct {
	// Defines the configuration of Backups created via this BackupPlan.
	BackupConfig *GkeBackupBackupPlanBackupConfig `pulumi:"backupConfig"`
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	BackupSchedule *GkeBackupBackupPlanBackupSchedule `pulumi:"backupSchedule"`
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster string `pulumi:"cluster"`
	// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such
	// that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any
	// new Backups from being created via this BackupPlan (including scheduled Backups).
	Deactivated *bool `pulumi:"deactivated"`
	// User specified descriptive string for this BackupPlan.
	Description           *string `pulumi:"description"`
	GkeBackupBackupPlanId *string `pulumi:"gkeBackupBackupPlanId"`
	// Description: A set of custom labels supplied by the user. A list of key->value pairs. Example: { "name": "wrench",
	// "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The region of the Backup Plan.
	Location string `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	RetentionPolicy *GkeBackupBackupPlanRetentionPolicy `pulumi:"retentionPolicy"`
	Timeouts        *GkeBackupBackupPlanTimeouts        `pulumi:"timeouts"`
}

// The set of arguments for constructing a GkeBackupBackupPlan resource.
type GkeBackupBackupPlanArgs struct {
	// Defines the configuration of Backups created via this BackupPlan.
	BackupConfig GkeBackupBackupPlanBackupConfigPtrInput
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	BackupSchedule GkeBackupBackupPlanBackupSchedulePtrInput
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster pulumi.StringInput
	// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such
	// that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any
	// new Backups from being created via this BackupPlan (including scheduled Backups).
	Deactivated pulumi.BoolPtrInput
	// User specified descriptive string for this BackupPlan.
	Description           pulumi.StringPtrInput
	GkeBackupBackupPlanId pulumi.StringPtrInput
	// Description: A set of custom labels supplied by the user. A list of key->value pairs. Example: { "name": "wrench",
	// "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The region of the Backup Plan.
	Location pulumi.StringInput
	// The full name of the BackupPlan Resource.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	RetentionPolicy GkeBackupBackupPlanRetentionPolicyPtrInput
	Timeouts        GkeBackupBackupPlanTimeoutsPtrInput
}

func (GkeBackupBackupPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeBackupBackupPlanArgs)(nil)).Elem()
}

type GkeBackupBackupPlanInput interface {
	pulumi.Input

	ToGkeBackupBackupPlanOutput() GkeBackupBackupPlanOutput
	ToGkeBackupBackupPlanOutputWithContext(ctx context.Context) GkeBackupBackupPlanOutput
}

func (*GkeBackupBackupPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeBackupBackupPlan)(nil)).Elem()
}

func (i *GkeBackupBackupPlan) ToGkeBackupBackupPlanOutput() GkeBackupBackupPlanOutput {
	return i.ToGkeBackupBackupPlanOutputWithContext(context.Background())
}

func (i *GkeBackupBackupPlan) ToGkeBackupBackupPlanOutputWithContext(ctx context.Context) GkeBackupBackupPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeBackupBackupPlanOutput)
}

type GkeBackupBackupPlanOutput struct{ *pulumi.OutputState }

func (GkeBackupBackupPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeBackupBackupPlan)(nil)).Elem()
}

func (o GkeBackupBackupPlanOutput) ToGkeBackupBackupPlanOutput() GkeBackupBackupPlanOutput {
	return o
}

func (o GkeBackupBackupPlanOutput) ToGkeBackupBackupPlanOutputWithContext(ctx context.Context) GkeBackupBackupPlanOutput {
	return o
}

// Defines the configuration of Backups created via this BackupPlan.
func (o GkeBackupBackupPlanOutput) BackupConfig() GkeBackupBackupPlanBackupConfigPtrOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) GkeBackupBackupPlanBackupConfigPtrOutput { return v.BackupConfig }).(GkeBackupBackupPlanBackupConfigPtrOutput)
}

// Defines a schedule for automatic Backup creation via this BackupPlan.
func (o GkeBackupBackupPlanOutput) BackupSchedule() GkeBackupBackupPlanBackupSchedulePtrOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) GkeBackupBackupPlanBackupSchedulePtrOutput { return v.BackupSchedule }).(GkeBackupBackupPlanBackupSchedulePtrOutput)
}

// The source cluster from which Backups will be created via this BackupPlan.
func (o GkeBackupBackupPlanOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such
// that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any
// new Backups from being created via this BackupPlan (including scheduled Backups).
func (o GkeBackupBackupPlanOutput) Deactivated() pulumi.BoolOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.BoolOutput { return v.Deactivated }).(pulumi.BoolOutput)
}

// User specified descriptive string for this BackupPlan.
func (o GkeBackupBackupPlanOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GkeBackupBackupPlanOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup plan from
// overwriting each other. It is strongly suggested that systems make use of the 'etag' in the read-modify-write cycle to
// perform BackupPlan updates in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
// and systems are expected to put that etag in the request to backupPlans.patch or backupPlans.delete to ensure that their
// change will be applied to the same version of the resource.
func (o GkeBackupBackupPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GkeBackupBackupPlanOutput) GkeBackupBackupPlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.GkeBackupBackupPlanId }).(pulumi.StringOutput)
}

// Description: A set of custom labels supplied by the user. A list of key->value pairs. Example: { "name": "wrench",
// "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels present in
// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o GkeBackupBackupPlanOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The region of the Backup Plan.
func (o GkeBackupBackupPlanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The full name of the BackupPlan Resource.
func (o GkeBackupBackupPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GkeBackupBackupPlanOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
func (o GkeBackupBackupPlanOutput) ProtectedPodCount() pulumi.Float64Output {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.Float64Output { return v.ProtectedPodCount }).(pulumi.Float64Output)
}

// RetentionPolicy governs lifecycle of Backups created under this plan.
func (o GkeBackupBackupPlanOutput) RetentionPolicy() GkeBackupBackupPlanRetentionPolicyPtrOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) GkeBackupBackupPlanRetentionPolicyPtrOutput { return v.RetentionPolicy }).(GkeBackupBackupPlanRetentionPolicyPtrOutput)
}

// The State of the BackupPlan.
func (o GkeBackupBackupPlanOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Detailed description of why BackupPlan is in its current state.
func (o GkeBackupBackupPlanOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o GkeBackupBackupPlanOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o GkeBackupBackupPlanOutput) Timeouts() GkeBackupBackupPlanTimeoutsPtrOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) GkeBackupBackupPlanTimeoutsPtrOutput { return v.Timeouts }).(GkeBackupBackupPlanTimeoutsPtrOutput)
}

// Server generated, unique identifier of UUID format.
func (o GkeBackupBackupPlanOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupBackupPlan) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeBackupBackupPlanInput)(nil)).Elem(), &GkeBackupBackupPlan{})
	pulumi.RegisterOutputType(GkeBackupBackupPlanOutput{})
}
