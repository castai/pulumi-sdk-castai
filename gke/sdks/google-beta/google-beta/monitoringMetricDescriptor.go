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

type MonitoringMetricDescriptor struct {
	pulumi.CustomResourceState

	// A detailed description of the metric, which can be used in documentation.
	Description pulumi.StringOutput `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period,
	// for example "Request count".
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the
	// entire resource must be deleted, then created with the desired labels.
	Labels MonitoringMetricDescriptorLabelArrayOutput `pulumi:"labels"`
	// The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH",
	// "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	LaunchStage pulumi.StringPtrOutput `pulumi:"launchStage"`
	// Metadata which can be used to guide usage of the metric.
	Metadata MonitoringMetricDescriptorMetadataPtrOutput `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType
	// might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]
	MetricKind pulumi.StringOutput `pulumi:"metricKind"`
	// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
	// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
	// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
	// this list.
	MonitoredResourceTypes       pulumi.StringArrayOutput `pulumi:"monitoredResourceTypes"`
	MonitoringMetricDescriptorId pulumi.StringOutput      `pulumi:"monitoringMetricDescriptorId"`
	// The resource name of the metric descriptor.
	Name     pulumi.StringOutput                         `pulumi:"name"`
	Project  pulumi.StringOutput                         `pulumi:"project"`
	Timeouts MonitoringMetricDescriptorTimeoutsPtrOutput `pulumi:"timeouts"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be
	// prefixed with the service name, in the format of {service name}/{relative metric name}, such as
	// cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters,
	// digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the
	// relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com,
	// external.googleapis.com, or logging.googleapis.com/user/.
	Type pulumi.StringOutput `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the valueType is INT64, DOUBLE, or
	// DISTRIBUTION. The unit defines the representation of the stored metric values. Different systems may scale the values to
	// be more easily displayed (so a value of 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in thousands of bytes, no matter how it may
	// be displayed. If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an
	// INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds,
	// then the value is written as 12005. Alternatively, if you want a custom metric to record data in a more granular way,
	// you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000),
	// or use Kis{CPU} and write 11.723 (which is 12005/1024). The supported units are a subset of The Unified Code for Units
	// of Measure standard. More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit pulumi.StringPtrOutput `pulumi:"unit"`
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might
	// not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]
	ValueType pulumi.StringOutput `pulumi:"valueType"`
}

// NewMonitoringMetricDescriptor registers a new resource with the given unique name, arguments, and options.
func NewMonitoringMetricDescriptor(ctx *pulumi.Context,
	name string, args *MonitoringMetricDescriptorArgs, opts ...pulumi.ResourceOption) (*MonitoringMetricDescriptor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MetricKind == nil {
		return nil, errors.New("invalid value for required argument 'MetricKind'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.ValueType == nil {
		return nil, errors.New("invalid value for required argument 'ValueType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource MonitoringMetricDescriptor
	err = ctx.RegisterPackageResource("google-beta:index/monitoringMetricDescriptor:MonitoringMetricDescriptor", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringMetricDescriptor gets an existing MonitoringMetricDescriptor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringMetricDescriptor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringMetricDescriptorState, opts ...pulumi.ResourceOption) (*MonitoringMetricDescriptor, error) {
	var resource MonitoringMetricDescriptor
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/monitoringMetricDescriptor:MonitoringMetricDescriptor", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringMetricDescriptor resources.
type monitoringMetricDescriptorState struct {
	// A detailed description of the metric, which can be used in documentation.
	Description *string `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period,
	// for example "Request count".
	DisplayName *string `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the
	// entire resource must be deleted, then created with the desired labels.
	Labels []MonitoringMetricDescriptorLabel `pulumi:"labels"`
	// The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH",
	// "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	LaunchStage *string `pulumi:"launchStage"`
	// Metadata which can be used to guide usage of the metric.
	Metadata *MonitoringMetricDescriptorMetadata `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType
	// might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]
	MetricKind *string `pulumi:"metricKind"`
	// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
	// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
	// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
	// this list.
	MonitoredResourceTypes       []string `pulumi:"monitoredResourceTypes"`
	MonitoringMetricDescriptorId *string  `pulumi:"monitoringMetricDescriptorId"`
	// The resource name of the metric descriptor.
	Name     *string                             `pulumi:"name"`
	Project  *string                             `pulumi:"project"`
	Timeouts *MonitoringMetricDescriptorTimeouts `pulumi:"timeouts"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be
	// prefixed with the service name, in the format of {service name}/{relative metric name}, such as
	// cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters,
	// digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the
	// relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com,
	// external.googleapis.com, or logging.googleapis.com/user/.
	Type *string `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the valueType is INT64, DOUBLE, or
	// DISTRIBUTION. The unit defines the representation of the stored metric values. Different systems may scale the values to
	// be more easily displayed (so a value of 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in thousands of bytes, no matter how it may
	// be displayed. If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an
	// INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds,
	// then the value is written as 12005. Alternatively, if you want a custom metric to record data in a more granular way,
	// you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000),
	// or use Kis{CPU} and write 11.723 (which is 12005/1024). The supported units are a subset of The Unified Code for Units
	// of Measure standard. More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit *string `pulumi:"unit"`
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might
	// not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]
	ValueType *string `pulumi:"valueType"`
}

type MonitoringMetricDescriptorState struct {
	// A detailed description of the metric, which can be used in documentation.
	Description pulumi.StringPtrInput
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period,
	// for example "Request count".
	DisplayName pulumi.StringPtrInput
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the
	// entire resource must be deleted, then created with the desired labels.
	Labels MonitoringMetricDescriptorLabelArrayInput
	// The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH",
	// "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	LaunchStage pulumi.StringPtrInput
	// Metadata which can be used to guide usage of the metric.
	Metadata MonitoringMetricDescriptorMetadataPtrInput
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType
	// might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]
	MetricKind pulumi.StringPtrInput
	// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
	// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
	// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
	// this list.
	MonitoredResourceTypes       pulumi.StringArrayInput
	MonitoringMetricDescriptorId pulumi.StringPtrInput
	// The resource name of the metric descriptor.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts MonitoringMetricDescriptorTimeoutsPtrInput
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be
	// prefixed with the service name, in the format of {service name}/{relative metric name}, such as
	// cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters,
	// digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the
	// relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com,
	// external.googleapis.com, or logging.googleapis.com/user/.
	Type pulumi.StringPtrInput
	// The units in which the metric value is reported. It is only applicable if the valueType is INT64, DOUBLE, or
	// DISTRIBUTION. The unit defines the representation of the stored metric values. Different systems may scale the values to
	// be more easily displayed (so a value of 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in thousands of bytes, no matter how it may
	// be displayed. If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an
	// INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds,
	// then the value is written as 12005. Alternatively, if you want a custom metric to record data in a more granular way,
	// you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000),
	// or use Kis{CPU} and write 11.723 (which is 12005/1024). The supported units are a subset of The Unified Code for Units
	// of Measure standard. More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit pulumi.StringPtrInput
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might
	// not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]
	ValueType pulumi.StringPtrInput
}

func (MonitoringMetricDescriptorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringMetricDescriptorState)(nil)).Elem()
}

type monitoringMetricDescriptorArgs struct {
	// A detailed description of the metric, which can be used in documentation.
	Description string `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period,
	// for example "Request count".
	DisplayName string `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the
	// entire resource must be deleted, then created with the desired labels.
	Labels []MonitoringMetricDescriptorLabel `pulumi:"labels"`
	// The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH",
	// "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	LaunchStage *string `pulumi:"launchStage"`
	// Metadata which can be used to guide usage of the metric.
	Metadata *MonitoringMetricDescriptorMetadata `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType
	// might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]
	MetricKind                   string                              `pulumi:"metricKind"`
	MonitoringMetricDescriptorId *string                             `pulumi:"monitoringMetricDescriptorId"`
	Project                      *string                             `pulumi:"project"`
	Timeouts                     *MonitoringMetricDescriptorTimeouts `pulumi:"timeouts"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be
	// prefixed with the service name, in the format of {service name}/{relative metric name}, such as
	// cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters,
	// digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the
	// relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com,
	// external.googleapis.com, or logging.googleapis.com/user/.
	Type string `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the valueType is INT64, DOUBLE, or
	// DISTRIBUTION. The unit defines the representation of the stored metric values. Different systems may scale the values to
	// be more easily displayed (so a value of 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in thousands of bytes, no matter how it may
	// be displayed. If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an
	// INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds,
	// then the value is written as 12005. Alternatively, if you want a custom metric to record data in a more granular way,
	// you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000),
	// or use Kis{CPU} and write 11.723 (which is 12005/1024). The supported units are a subset of The Unified Code for Units
	// of Measure standard. More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit *string `pulumi:"unit"`
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might
	// not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]
	ValueType string `pulumi:"valueType"`
}

// The set of arguments for constructing a MonitoringMetricDescriptor resource.
type MonitoringMetricDescriptorArgs struct {
	// A detailed description of the metric, which can be used in documentation.
	Description pulumi.StringInput
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period,
	// for example "Request count".
	DisplayName pulumi.StringInput
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the
	// entire resource must be deleted, then created with the desired labels.
	Labels MonitoringMetricDescriptorLabelArrayInput
	// The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH",
	// "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	LaunchStage pulumi.StringPtrInput
	// Metadata which can be used to guide usage of the metric.
	Metadata MonitoringMetricDescriptorMetadataPtrInput
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType
	// might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]
	MetricKind                   pulumi.StringInput
	MonitoringMetricDescriptorId pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Timeouts                     MonitoringMetricDescriptorTimeoutsPtrInput
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be
	// prefixed with the service name, in the format of {service name}/{relative metric name}, such as
	// cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters,
	// digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the
	// relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com,
	// external.googleapis.com, or logging.googleapis.com/user/.
	Type pulumi.StringInput
	// The units in which the metric value is reported. It is only applicable if the valueType is INT64, DOUBLE, or
	// DISTRIBUTION. The unit defines the representation of the stored metric values. Different systems may scale the values to
	// be more easily displayed (so a value of 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in thousands of bytes, no matter how it may
	// be displayed. If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an
	// INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds,
	// then the value is written as 12005. Alternatively, if you want a custom metric to record data in a more granular way,
	// you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000),
	// or use Kis{CPU} and write 11.723 (which is 12005/1024). The supported units are a subset of The Unified Code for Units
	// of Measure standard. More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit pulumi.StringPtrInput
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might
	// not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]
	ValueType pulumi.StringInput
}

func (MonitoringMetricDescriptorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringMetricDescriptorArgs)(nil)).Elem()
}

type MonitoringMetricDescriptorInput interface {
	pulumi.Input

	ToMonitoringMetricDescriptorOutput() MonitoringMetricDescriptorOutput
	ToMonitoringMetricDescriptorOutputWithContext(ctx context.Context) MonitoringMetricDescriptorOutput
}

func (*MonitoringMetricDescriptor) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringMetricDescriptor)(nil)).Elem()
}

func (i *MonitoringMetricDescriptor) ToMonitoringMetricDescriptorOutput() MonitoringMetricDescriptorOutput {
	return i.ToMonitoringMetricDescriptorOutputWithContext(context.Background())
}

func (i *MonitoringMetricDescriptor) ToMonitoringMetricDescriptorOutputWithContext(ctx context.Context) MonitoringMetricDescriptorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringMetricDescriptorOutput)
}

type MonitoringMetricDescriptorOutput struct{ *pulumi.OutputState }

func (MonitoringMetricDescriptorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringMetricDescriptor)(nil)).Elem()
}

func (o MonitoringMetricDescriptorOutput) ToMonitoringMetricDescriptorOutput() MonitoringMetricDescriptorOutput {
	return o
}

func (o MonitoringMetricDescriptorOutput) ToMonitoringMetricDescriptorOutputWithContext(ctx context.Context) MonitoringMetricDescriptorOutput {
	return o
}

// A detailed description of the metric, which can be used in documentation.
func (o MonitoringMetricDescriptorOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period,
// for example "Request count".
func (o MonitoringMetricDescriptorOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the
// entire resource must be deleted, then created with the desired labels.
func (o MonitoringMetricDescriptorOutput) Labels() MonitoringMetricDescriptorLabelArrayOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) MonitoringMetricDescriptorLabelArrayOutput { return v.Labels }).(MonitoringMetricDescriptorLabelArrayOutput)
}

// The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH",
// "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
func (o MonitoringMetricDescriptorOutput) LaunchStage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringPtrOutput { return v.LaunchStage }).(pulumi.StringPtrOutput)
}

// Metadata which can be used to guide usage of the metric.
func (o MonitoringMetricDescriptorOutput) Metadata() MonitoringMetricDescriptorMetadataPtrOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) MonitoringMetricDescriptorMetadataPtrOutput { return v.Metadata }).(MonitoringMetricDescriptorMetadataPtrOutput)
}

// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType
// might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]
func (o MonitoringMetricDescriptorOutput) MetricKind() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.MetricKind }).(pulumi.StringOutput)
}

// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
// this list.
func (o MonitoringMetricDescriptorOutput) MonitoredResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringArrayOutput { return v.MonitoredResourceTypes }).(pulumi.StringArrayOutput)
}

func (o MonitoringMetricDescriptorOutput) MonitoringMetricDescriptorId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.MonitoringMetricDescriptorId }).(pulumi.StringOutput)
}

// The resource name of the metric descriptor.
func (o MonitoringMetricDescriptorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MonitoringMetricDescriptorOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o MonitoringMetricDescriptorOutput) Timeouts() MonitoringMetricDescriptorTimeoutsPtrOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) MonitoringMetricDescriptorTimeoutsPtrOutput { return v.Timeouts }).(MonitoringMetricDescriptorTimeoutsPtrOutput)
}

// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be
// prefixed with the service name, in the format of {service name}/{relative metric name}, such as
// cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters,
// digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the
// relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com,
// external.googleapis.com, or logging.googleapis.com/user/.
func (o MonitoringMetricDescriptorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The units in which the metric value is reported. It is only applicable if the valueType is INT64, DOUBLE, or
// DISTRIBUTION. The unit defines the representation of the stored metric values. Different systems may scale the values to
// be more easily displayed (so a value of 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in thousands of bytes, no matter how it may
// be displayed. If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an
// INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds,
// then the value is written as 12005. Alternatively, if you want a custom metric to record data in a more granular way,
// you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000),
// or use Kis{CPU} and write 11.723 (which is 12005/1024). The supported units are a subset of The Unified Code for Units
// of Measure standard. More info can be found in the API documentation
// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
func (o MonitoringMetricDescriptorOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringPtrOutput { return v.Unit }).(pulumi.StringPtrOutput)
}

// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might
// not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]
func (o MonitoringMetricDescriptorOutput) ValueType() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringMetricDescriptor) pulumi.StringOutput { return v.ValueType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringMetricDescriptorInput)(nil)).Elem(), &MonitoringMetricDescriptor{})
	pulumi.RegisterOutputType(MonitoringMetricDescriptorOutput{})
}
