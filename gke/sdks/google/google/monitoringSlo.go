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

type MonitoringSlo struct {
	pulumi.CustomResourceState

	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli',
	// 'windows_based_sli'
	BasicSli MonitoringSloBasicSliPtrOutput `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>". Possible values: ["DAY", "WEEK",
	// "FORTNIGHT", "MONTH"]
	CalendarPeriod pulumi.StringPtrOutput `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal            pulumi.Float64Output `pulumi:"goal"`
	MonitoringSloId pulumi.StringOutput  `pulumi:"monitoringSloId"`
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A request-based SLI defines a SLI for which atomic units of service are counted directly. A SLI describes a good
	// service. It is used to measure and calculate the quality of the Service's performance with respect to a single aspect of
	// service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	RequestBasedSli MonitoringSloRequestBasedSliPtrOutput `pulumi:"requestBasedSli"`
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.Float64PtrOutput `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	Service pulumi.StringOutput `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId    pulumi.StringOutput            `pulumi:"sloId"`
	Timeouts MonitoringSloTimeoutsPtrOutput `pulumi:"timeouts"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapOutput `pulumi:"userLabels"`
	// A windows-based SLI defines the criteria for time windows. good_service is defined based off the count of these time
	// windows for which the provided service was of good quality. A SLI describes a good service. It is used to measure and
	// calculate the quality of the Service's performance with respect to a single aspect of service quality. Exactly one of
	// the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	WindowsBasedSli MonitoringSloWindowsBasedSliPtrOutput `pulumi:"windowsBasedSli"`
}

// NewMonitoringSlo registers a new resource with the given unique name, arguments, and options.
func NewMonitoringSlo(ctx *pulumi.Context,
	name string, args *MonitoringSloArgs, opts ...pulumi.ResourceOption) (*MonitoringSlo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Goal == nil {
		return nil, errors.New("invalid value for required argument 'Goal'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource MonitoringSlo
	err = ctx.RegisterPackageResource("google:index/monitoringSlo:MonitoringSlo", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringSlo gets an existing MonitoringSlo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringSlo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringSloState, opts ...pulumi.ResourceOption) (*MonitoringSlo, error) {
	var resource MonitoringSlo
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/monitoringSlo:MonitoringSlo", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringSlo resources.
type monitoringSloState struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli',
	// 'windows_based_sli'
	BasicSli *MonitoringSloBasicSli `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>". Possible values: ["DAY", "WEEK",
	// "FORTNIGHT", "MONTH"]
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal            *float64 `pulumi:"goal"`
	MonitoringSloId *string  `pulumi:"monitoringSloId"`
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A request-based SLI defines a SLI for which atomic units of service are counted directly. A SLI describes a good
	// service. It is used to measure and calculate the quality of the Service's performance with respect to a single aspect of
	// service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	RequestBasedSli *MonitoringSloRequestBasedSli `pulumi:"requestBasedSli"`
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays *float64 `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	Service *string `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId    *string                `pulumi:"sloId"`
	Timeouts *MonitoringSloTimeouts `pulumi:"timeouts"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
	// A windows-based SLI defines the criteria for time windows. good_service is defined based off the count of these time
	// windows for which the provided service was of good quality. A SLI describes a good service. It is used to measure and
	// calculate the quality of the Service's performance with respect to a single aspect of service quality. Exactly one of
	// the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	WindowsBasedSli *MonitoringSloWindowsBasedSli `pulumi:"windowsBasedSli"`
}

type MonitoringSloState struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli',
	// 'windows_based_sli'
	BasicSli MonitoringSloBasicSliPtrInput
	// A calendar period, semantically "since the start of the current <calendarPeriod>". Possible values: ["DAY", "WEEK",
	// "FORTNIGHT", "MONTH"]
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal            pulumi.Float64PtrInput
	MonitoringSloId pulumi.StringPtrInput
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A request-based SLI defines a SLI for which atomic units of service are counted directly. A SLI describes a good
	// service. It is used to measure and calculate the quality of the Service's performance with respect to a single aspect of
	// service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	RequestBasedSli MonitoringSloRequestBasedSliPtrInput
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.Float64PtrInput
	// ID of the service to which this SLO belongs.
	Service pulumi.StringPtrInput
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId    pulumi.StringPtrInput
	Timeouts MonitoringSloTimeoutsPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapInput
	// A windows-based SLI defines the criteria for time windows. good_service is defined based off the count of these time
	// windows for which the provided service was of good quality. A SLI describes a good service. It is used to measure and
	// calculate the quality of the Service's performance with respect to a single aspect of service quality. Exactly one of
	// the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	WindowsBasedSli MonitoringSloWindowsBasedSliPtrInput
}

func (MonitoringSloState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSloState)(nil)).Elem()
}

type monitoringSloArgs struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli',
	// 'windows_based_sli'
	BasicSli *MonitoringSloBasicSli `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>". Possible values: ["DAY", "WEEK",
	// "FORTNIGHT", "MONTH"]
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal            float64 `pulumi:"goal"`
	MonitoringSloId *string `pulumi:"monitoringSloId"`
	Project         *string `pulumi:"project"`
	// A request-based SLI defines a SLI for which atomic units of service are counted directly. A SLI describes a good
	// service. It is used to measure and calculate the quality of the Service's performance with respect to a single aspect of
	// service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	RequestBasedSli *MonitoringSloRequestBasedSli `pulumi:"requestBasedSli"`
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays *float64 `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	Service string `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId    *string                `pulumi:"sloId"`
	Timeouts *MonitoringSloTimeouts `pulumi:"timeouts"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
	// A windows-based SLI defines the criteria for time windows. good_service is defined based off the count of these time
	// windows for which the provided service was of good quality. A SLI describes a good service. It is used to measure and
	// calculate the quality of the Service's performance with respect to a single aspect of service quality. Exactly one of
	// the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	WindowsBasedSli *MonitoringSloWindowsBasedSli `pulumi:"windowsBasedSli"`
}

// The set of arguments for constructing a MonitoringSlo resource.
type MonitoringSloArgs struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli',
	// 'windows_based_sli'
	BasicSli MonitoringSloBasicSliPtrInput
	// A calendar period, semantically "since the start of the current <calendarPeriod>". Possible values: ["DAY", "WEEK",
	// "FORTNIGHT", "MONTH"]
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal            pulumi.Float64Input
	MonitoringSloId pulumi.StringPtrInput
	Project         pulumi.StringPtrInput
	// A request-based SLI defines a SLI for which atomic units of service are counted directly. A SLI describes a good
	// service. It is used to measure and calculate the quality of the Service's performance with respect to a single aspect of
	// service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	RequestBasedSli MonitoringSloRequestBasedSliPtrInput
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.Float64PtrInput
	// ID of the service to which this SLO belongs.
	Service pulumi.StringInput
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId    pulumi.StringPtrInput
	Timeouts MonitoringSloTimeoutsPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapInput
	// A windows-based SLI defines the criteria for time windows. good_service is defined based off the count of these time
	// windows for which the provided service was of good quality. A SLI describes a good service. It is used to measure and
	// calculate the quality of the Service's performance with respect to a single aspect of service quality. Exactly one of
	// the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
	WindowsBasedSli MonitoringSloWindowsBasedSliPtrInput
}

func (MonitoringSloArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSloArgs)(nil)).Elem()
}

type MonitoringSloInput interface {
	pulumi.Input

	ToMonitoringSloOutput() MonitoringSloOutput
	ToMonitoringSloOutputWithContext(ctx context.Context) MonitoringSloOutput
}

func (*MonitoringSlo) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSlo)(nil)).Elem()
}

func (i *MonitoringSlo) ToMonitoringSloOutput() MonitoringSloOutput {
	return i.ToMonitoringSloOutputWithContext(context.Background())
}

func (i *MonitoringSlo) ToMonitoringSloOutputWithContext(ctx context.Context) MonitoringSloOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSloOutput)
}

type MonitoringSloOutput struct{ *pulumi.OutputState }

func (MonitoringSloOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSlo)(nil)).Elem()
}

func (o MonitoringSloOutput) ToMonitoringSloOutput() MonitoringSloOutput {
	return o
}

func (o MonitoringSloOutput) ToMonitoringSloOutputWithContext(ctx context.Context) MonitoringSloOutput {
	return o
}

// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
// single aspect of service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli',
// 'windows_based_sli'
func (o MonitoringSloOutput) BasicSli() MonitoringSloBasicSliPtrOutput {
	return o.ApplyT(func(v *MonitoringSlo) MonitoringSloBasicSliPtrOutput { return v.BasicSli }).(MonitoringSloBasicSliPtrOutput)
}

// A calendar period, semantically "since the start of the current <calendarPeriod>". Possible values: ["DAY", "WEEK",
// "FORTNIGHT", "MONTH"]
func (o MonitoringSloOutput) CalendarPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringPtrOutput { return v.CalendarPeriod }).(pulumi.StringPtrOutput)
}

// Name used for UI elements listing this SLO.
func (o MonitoringSloOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
func (o MonitoringSloOutput) Goal() pulumi.Float64Output {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.Float64Output { return v.Goal }).(pulumi.Float64Output)
}

func (o MonitoringSloOutput) MonitoringSloId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringOutput { return v.MonitoringSloId }).(pulumi.StringOutput)
}

// The full resource name for this service. The syntax is:
// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
func (o MonitoringSloOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MonitoringSloOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A request-based SLI defines a SLI for which atomic units of service are counted directly. A SLI describes a good
// service. It is used to measure and calculate the quality of the Service's performance with respect to a single aspect of
// service quality. Exactly one of the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
func (o MonitoringSloOutput) RequestBasedSli() MonitoringSloRequestBasedSliPtrOutput {
	return o.ApplyT(func(v *MonitoringSlo) MonitoringSloRequestBasedSliPtrOutput { return v.RequestBasedSli }).(MonitoringSloRequestBasedSliPtrOutput)
}

// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
func (o MonitoringSloOutput) RollingPeriodDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.Float64PtrOutput { return v.RollingPeriodDays }).(pulumi.Float64PtrOutput)
}

// ID of the service to which this SLO belongs.
func (o MonitoringSloOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
func (o MonitoringSloOutput) SloId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringOutput { return v.SloId }).(pulumi.StringOutput)
}

func (o MonitoringSloOutput) Timeouts() MonitoringSloTimeoutsPtrOutput {
	return o.ApplyT(func(v *MonitoringSlo) MonitoringSloTimeoutsPtrOutput { return v.Timeouts }).(MonitoringSloTimeoutsPtrOutput)
}

// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
func (o MonitoringSloOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MonitoringSlo) pulumi.StringMapOutput { return v.UserLabels }).(pulumi.StringMapOutput)
}

// A windows-based SLI defines the criteria for time windows. good_service is defined based off the count of these time
// windows for which the provided service was of good quality. A SLI describes a good service. It is used to measure and
// calculate the quality of the Service's performance with respect to a single aspect of service quality. Exactly one of
// the following must be set: 'basic_sli', 'request_based_sli', 'windows_based_sli'
func (o MonitoringSloOutput) WindowsBasedSli() MonitoringSloWindowsBasedSliPtrOutput {
	return o.ApplyT(func(v *MonitoringSlo) MonitoringSloWindowsBasedSliPtrOutput { return v.WindowsBasedSli }).(MonitoringSloWindowsBasedSliPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringSloInput)(nil)).Elem(), &MonitoringSlo{})
	pulumi.RegisterOutputType(MonitoringSloOutput{})
}
