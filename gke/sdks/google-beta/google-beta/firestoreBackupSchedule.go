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

type FirestoreBackupSchedule struct {
	pulumi.CustomResourceState

	// For a schedule that runs daily.
	DailyRecurrence FirestoreBackupScheduleDailyRecurrencePtrOutput `pulumi:"dailyRecurrence"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database                  pulumi.StringPtrOutput `pulumi:"database"`
	FirestoreBackupScheduleId pulumi.StringOutput    `pulumi:"firestoreBackupScheduleId"`
	// The unique backup schedule identifier across all locations and databases for the given project. Format:
	// 'projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}'
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for
	// 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". You can set this to a
	// value up to 14 weeks.
	Retention pulumi.StringOutput                      `pulumi:"retention"`
	Timeouts  FirestoreBackupScheduleTimeoutsPtrOutput `pulumi:"timeouts"`
	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence FirestoreBackupScheduleWeeklyRecurrencePtrOutput `pulumi:"weeklyRecurrence"`
}

// NewFirestoreBackupSchedule registers a new resource with the given unique name, arguments, and options.
func NewFirestoreBackupSchedule(ctx *pulumi.Context,
	name string, args *FirestoreBackupScheduleArgs, opts ...pulumi.ResourceOption) (*FirestoreBackupSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Retention == nil {
		return nil, errors.New("invalid value for required argument 'Retention'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirestoreBackupSchedule
	err = ctx.RegisterPackageResource("google-beta:index/firestoreBackupSchedule:FirestoreBackupSchedule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirestoreBackupSchedule gets an existing FirestoreBackupSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirestoreBackupSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirestoreBackupScheduleState, opts ...pulumi.ResourceOption) (*FirestoreBackupSchedule, error) {
	var resource FirestoreBackupSchedule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/firestoreBackupSchedule:FirestoreBackupSchedule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirestoreBackupSchedule resources.
type firestoreBackupScheduleState struct {
	// For a schedule that runs daily.
	DailyRecurrence *FirestoreBackupScheduleDailyRecurrence `pulumi:"dailyRecurrence"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database                  *string `pulumi:"database"`
	FirestoreBackupScheduleId *string `pulumi:"firestoreBackupScheduleId"`
	// The unique backup schedule identifier across all locations and databases for the given project. Format:
	// 'projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}'
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for
	// 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". You can set this to a
	// value up to 14 weeks.
	Retention *string                          `pulumi:"retention"`
	Timeouts  *FirestoreBackupScheduleTimeouts `pulumi:"timeouts"`
	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence *FirestoreBackupScheduleWeeklyRecurrence `pulumi:"weeklyRecurrence"`
}

type FirestoreBackupScheduleState struct {
	// For a schedule that runs daily.
	DailyRecurrence FirestoreBackupScheduleDailyRecurrencePtrInput
	// The Firestore database id. Defaults to '"(default)"'.
	Database                  pulumi.StringPtrInput
	FirestoreBackupScheduleId pulumi.StringPtrInput
	// The unique backup schedule identifier across all locations and databases for the given project. Format:
	// 'projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}'
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for
	// 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". You can set this to a
	// value up to 14 weeks.
	Retention pulumi.StringPtrInput
	Timeouts  FirestoreBackupScheduleTimeoutsPtrInput
	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence FirestoreBackupScheduleWeeklyRecurrencePtrInput
}

func (FirestoreBackupScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firestoreBackupScheduleState)(nil)).Elem()
}

type firestoreBackupScheduleArgs struct {
	// For a schedule that runs daily.
	DailyRecurrence *FirestoreBackupScheduleDailyRecurrence `pulumi:"dailyRecurrence"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database                  *string `pulumi:"database"`
	FirestoreBackupScheduleId *string `pulumi:"firestoreBackupScheduleId"`
	Project                   *string `pulumi:"project"`
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for
	// 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". You can set this to a
	// value up to 14 weeks.
	Retention string                           `pulumi:"retention"`
	Timeouts  *FirestoreBackupScheduleTimeouts `pulumi:"timeouts"`
	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence *FirestoreBackupScheduleWeeklyRecurrence `pulumi:"weeklyRecurrence"`
}

// The set of arguments for constructing a FirestoreBackupSchedule resource.
type FirestoreBackupScheduleArgs struct {
	// For a schedule that runs daily.
	DailyRecurrence FirestoreBackupScheduleDailyRecurrencePtrInput
	// The Firestore database id. Defaults to '"(default)"'.
	Database                  pulumi.StringPtrInput
	FirestoreBackupScheduleId pulumi.StringPtrInput
	Project                   pulumi.StringPtrInput
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for
	// 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". You can set this to a
	// value up to 14 weeks.
	Retention pulumi.StringInput
	Timeouts  FirestoreBackupScheduleTimeoutsPtrInput
	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence FirestoreBackupScheduleWeeklyRecurrencePtrInput
}

func (FirestoreBackupScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firestoreBackupScheduleArgs)(nil)).Elem()
}

type FirestoreBackupScheduleInput interface {
	pulumi.Input

	ToFirestoreBackupScheduleOutput() FirestoreBackupScheduleOutput
	ToFirestoreBackupScheduleOutputWithContext(ctx context.Context) FirestoreBackupScheduleOutput
}

func (*FirestoreBackupSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirestoreBackupSchedule)(nil)).Elem()
}

func (i *FirestoreBackupSchedule) ToFirestoreBackupScheduleOutput() FirestoreBackupScheduleOutput {
	return i.ToFirestoreBackupScheduleOutputWithContext(context.Background())
}

func (i *FirestoreBackupSchedule) ToFirestoreBackupScheduleOutputWithContext(ctx context.Context) FirestoreBackupScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirestoreBackupScheduleOutput)
}

type FirestoreBackupScheduleOutput struct{ *pulumi.OutputState }

func (FirestoreBackupScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirestoreBackupSchedule)(nil)).Elem()
}

func (o FirestoreBackupScheduleOutput) ToFirestoreBackupScheduleOutput() FirestoreBackupScheduleOutput {
	return o
}

func (o FirestoreBackupScheduleOutput) ToFirestoreBackupScheduleOutputWithContext(ctx context.Context) FirestoreBackupScheduleOutput {
	return o
}

// For a schedule that runs daily.
func (o FirestoreBackupScheduleOutput) DailyRecurrence() FirestoreBackupScheduleDailyRecurrencePtrOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) FirestoreBackupScheduleDailyRecurrencePtrOutput {
		return v.DailyRecurrence
	}).(FirestoreBackupScheduleDailyRecurrencePtrOutput)
}

// The Firestore database id. Defaults to '"(default)"'.
func (o FirestoreBackupScheduleOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) pulumi.StringPtrOutput { return v.Database }).(pulumi.StringPtrOutput)
}

func (o FirestoreBackupScheduleOutput) FirestoreBackupScheduleId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) pulumi.StringOutput { return v.FirestoreBackupScheduleId }).(pulumi.StringOutput)
}

// The unique backup schedule identifier across all locations and databases for the given project. Format:
// 'projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}'
func (o FirestoreBackupScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirestoreBackupScheduleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for
// 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". You can set this to a
// value up to 14 weeks.
func (o FirestoreBackupScheduleOutput) Retention() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) pulumi.StringOutput { return v.Retention }).(pulumi.StringOutput)
}

func (o FirestoreBackupScheduleOutput) Timeouts() FirestoreBackupScheduleTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) FirestoreBackupScheduleTimeoutsPtrOutput { return v.Timeouts }).(FirestoreBackupScheduleTimeoutsPtrOutput)
}

// For a schedule that runs weekly on a specific day.
func (o FirestoreBackupScheduleOutput) WeeklyRecurrence() FirestoreBackupScheduleWeeklyRecurrencePtrOutput {
	return o.ApplyT(func(v *FirestoreBackupSchedule) FirestoreBackupScheduleWeeklyRecurrencePtrOutput {
		return v.WeeklyRecurrence
	}).(FirestoreBackupScheduleWeeklyRecurrencePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirestoreBackupScheduleInput)(nil)).Elem(), &FirestoreBackupSchedule{})
	pulumi.RegisterOutputType(FirestoreBackupScheduleOutput{})
}
