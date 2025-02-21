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

type LoggingFolderBucketConfig struct {
	pulumi.CustomResourceState

	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings LoggingFolderBucketConfigCmekSettingsPtrOutput `pulumi:"cmekSettings"`
	// An optional description for this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The parent resource that contains the logging bucket.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// A list of indexed fields and related configuration data.
	IndexConfigs LoggingFolderBucketConfigIndexConfigArrayOutput `pulumi:"indexConfigs"`
	// The bucket's lifecycle such as active or deleted.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location                    pulumi.StringOutput `pulumi:"location"`
	LoggingFolderBucketConfigId pulumi.StringOutput `pulumi:"loggingFolderBucketConfigId"`
	// The resource name of the bucket
	Name pulumi.StringOutput `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum
	// retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be
	// used.
	RetentionDays pulumi.Float64PtrOutput `pulumi:"retentionDays"`
}

// NewLoggingFolderBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewLoggingFolderBucketConfig(ctx *pulumi.Context,
	name string, args *LoggingFolderBucketConfigArgs, opts ...pulumi.ResourceOption) (*LoggingFolderBucketConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource LoggingFolderBucketConfig
	err = ctx.RegisterPackageResource("google-beta:index/loggingFolderBucketConfig:LoggingFolderBucketConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingFolderBucketConfig gets an existing LoggingFolderBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingFolderBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingFolderBucketConfigState, opts ...pulumi.ResourceOption) (*LoggingFolderBucketConfig, error) {
	var resource LoggingFolderBucketConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/loggingFolderBucketConfig:LoggingFolderBucketConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingFolderBucketConfig resources.
type loggingFolderBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	BucketId *string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings *LoggingFolderBucketConfigCmekSettings `pulumi:"cmekSettings"`
	// An optional description for this bucket.
	Description *string `pulumi:"description"`
	// The parent resource that contains the logging bucket.
	Folder *string `pulumi:"folder"`
	// A list of indexed fields and related configuration data.
	IndexConfigs []LoggingFolderBucketConfigIndexConfig `pulumi:"indexConfigs"`
	// The bucket's lifecycle such as active or deleted.
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location                    *string `pulumi:"location"`
	LoggingFolderBucketConfigId *string `pulumi:"loggingFolderBucketConfigId"`
	// The resource name of the bucket
	Name *string `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum
	// retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be
	// used.
	RetentionDays *float64 `pulumi:"retentionDays"`
}

type LoggingFolderBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	BucketId pulumi.StringPtrInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings LoggingFolderBucketConfigCmekSettingsPtrInput
	// An optional description for this bucket.
	Description pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Folder pulumi.StringPtrInput
	// A list of indexed fields and related configuration data.
	IndexConfigs LoggingFolderBucketConfigIndexConfigArrayInput
	// The bucket's lifecycle such as active or deleted.
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket.
	Location                    pulumi.StringPtrInput
	LoggingFolderBucketConfigId pulumi.StringPtrInput
	// The resource name of the bucket
	Name pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum
	// retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be
	// used.
	RetentionDays pulumi.Float64PtrInput
}

func (LoggingFolderBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingFolderBucketConfigState)(nil)).Elem()
}

type loggingFolderBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	BucketId string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings *LoggingFolderBucketConfigCmekSettings `pulumi:"cmekSettings"`
	// An optional description for this bucket.
	Description *string `pulumi:"description"`
	// The parent resource that contains the logging bucket.
	Folder string `pulumi:"folder"`
	// A list of indexed fields and related configuration data.
	IndexConfigs []LoggingFolderBucketConfigIndexConfig `pulumi:"indexConfigs"`
	// The location of the bucket.
	Location                    string  `pulumi:"location"`
	LoggingFolderBucketConfigId *string `pulumi:"loggingFolderBucketConfigId"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum
	// retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be
	// used.
	RetentionDays *float64 `pulumi:"retentionDays"`
}

// The set of arguments for constructing a LoggingFolderBucketConfig resource.
type LoggingFolderBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	BucketId pulumi.StringInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings LoggingFolderBucketConfigCmekSettingsPtrInput
	// An optional description for this bucket.
	Description pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Folder pulumi.StringInput
	// A list of indexed fields and related configuration data.
	IndexConfigs LoggingFolderBucketConfigIndexConfigArrayInput
	// The location of the bucket.
	Location                    pulumi.StringInput
	LoggingFolderBucketConfigId pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum
	// retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be
	// used.
	RetentionDays pulumi.Float64PtrInput
}

func (LoggingFolderBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingFolderBucketConfigArgs)(nil)).Elem()
}

type LoggingFolderBucketConfigInput interface {
	pulumi.Input

	ToLoggingFolderBucketConfigOutput() LoggingFolderBucketConfigOutput
	ToLoggingFolderBucketConfigOutputWithContext(ctx context.Context) LoggingFolderBucketConfigOutput
}

func (*LoggingFolderBucketConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingFolderBucketConfig)(nil)).Elem()
}

func (i *LoggingFolderBucketConfig) ToLoggingFolderBucketConfigOutput() LoggingFolderBucketConfigOutput {
	return i.ToLoggingFolderBucketConfigOutputWithContext(context.Background())
}

func (i *LoggingFolderBucketConfig) ToLoggingFolderBucketConfigOutputWithContext(ctx context.Context) LoggingFolderBucketConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingFolderBucketConfigOutput)
}

type LoggingFolderBucketConfigOutput struct{ *pulumi.OutputState }

func (LoggingFolderBucketConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingFolderBucketConfig)(nil)).Elem()
}

func (o LoggingFolderBucketConfigOutput) ToLoggingFolderBucketConfigOutput() LoggingFolderBucketConfigOutput {
	return o
}

func (o LoggingFolderBucketConfigOutput) ToLoggingFolderBucketConfigOutputWithContext(ctx context.Context) LoggingFolderBucketConfigOutput {
	return o
}

// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
func (o LoggingFolderBucketConfigOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
// updating the log bucket. Changing the KMS key is allowed.
func (o LoggingFolderBucketConfigOutput) CmekSettings() LoggingFolderBucketConfigCmekSettingsPtrOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) LoggingFolderBucketConfigCmekSettingsPtrOutput {
		return v.CmekSettings
	}).(LoggingFolderBucketConfigCmekSettingsPtrOutput)
}

// An optional description for this bucket.
func (o LoggingFolderBucketConfigOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The parent resource that contains the logging bucket.
func (o LoggingFolderBucketConfigOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

// A list of indexed fields and related configuration data.
func (o LoggingFolderBucketConfigOutput) IndexConfigs() LoggingFolderBucketConfigIndexConfigArrayOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) LoggingFolderBucketConfigIndexConfigArrayOutput {
		return v.IndexConfigs
	}).(LoggingFolderBucketConfigIndexConfigArrayOutput)
}

// The bucket's lifecycle such as active or deleted.
func (o LoggingFolderBucketConfigOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// The location of the bucket.
func (o LoggingFolderBucketConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LoggingFolderBucketConfigOutput) LoggingFolderBucketConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.StringOutput { return v.LoggingFolderBucketConfigId }).(pulumi.StringOutput)
}

// The resource name of the bucket
func (o LoggingFolderBucketConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum
// retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be
// used.
func (o LoggingFolderBucketConfigOutput) RetentionDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggingFolderBucketConfig) pulumi.Float64PtrOutput { return v.RetentionDays }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingFolderBucketConfigInput)(nil)).Elem(), &LoggingFolderBucketConfig{})
	pulumi.RegisterOutputType(LoggingFolderBucketConfigOutput{})
}
