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

type LoggingFolderSettings struct {
	pulumi.CustomResourceState

	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink pulumi.BoolOutput `pulumi:"disableDefaultSink"`
	// The folder for which to retrieve settings.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// The resource name for the configured Cloud KMS key.
	KmsKeyName pulumi.StringOutput `pulumi:"kmsKeyName"`
	// The service account that will be used by the Log Router to access your Cloud KMS key.
	KmsServiceAccountId     pulumi.StringOutput `pulumi:"kmsServiceAccountId"`
	LoggingFolderSettingsId pulumi.StringOutput `pulumi:"loggingFolderSettingsId"`
	// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
	// account is provided.
	LoggingServiceAccountId pulumi.StringOutput `pulumi:"loggingServiceAccountId"`
	// The resource name of the settings.
	Name pulumi.StringOutput `pulumi:"name"`
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation pulumi.StringOutput                    `pulumi:"storageLocation"`
	Timeouts        LoggingFolderSettingsTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewLoggingFolderSettings registers a new resource with the given unique name, arguments, and options.
func NewLoggingFolderSettings(ctx *pulumi.Context,
	name string, args *LoggingFolderSettingsArgs, opts ...pulumi.ResourceOption) (*LoggingFolderSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource LoggingFolderSettings
	err = ctx.RegisterPackageResource("google-beta:index/loggingFolderSettings:LoggingFolderSettings", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingFolderSettings gets an existing LoggingFolderSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingFolderSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingFolderSettingsState, opts ...pulumi.ResourceOption) (*LoggingFolderSettings, error) {
	var resource LoggingFolderSettings
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/loggingFolderSettings:LoggingFolderSettings", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingFolderSettings resources.
type loggingFolderSettingsState struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink *bool `pulumi:"disableDefaultSink"`
	// The folder for which to retrieve settings.
	Folder *string `pulumi:"folder"`
	// The resource name for the configured Cloud KMS key.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// The service account that will be used by the Log Router to access your Cloud KMS key.
	KmsServiceAccountId     *string `pulumi:"kmsServiceAccountId"`
	LoggingFolderSettingsId *string `pulumi:"loggingFolderSettingsId"`
	// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
	// account is provided.
	LoggingServiceAccountId *string `pulumi:"loggingServiceAccountId"`
	// The resource name of the settings.
	Name *string `pulumi:"name"`
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation *string                        `pulumi:"storageLocation"`
	Timeouts        *LoggingFolderSettingsTimeouts `pulumi:"timeouts"`
}

type LoggingFolderSettingsState struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink pulumi.BoolPtrInput
	// The folder for which to retrieve settings.
	Folder pulumi.StringPtrInput
	// The resource name for the configured Cloud KMS key.
	KmsKeyName pulumi.StringPtrInput
	// The service account that will be used by the Log Router to access your Cloud KMS key.
	KmsServiceAccountId     pulumi.StringPtrInput
	LoggingFolderSettingsId pulumi.StringPtrInput
	// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
	// account is provided.
	LoggingServiceAccountId pulumi.StringPtrInput
	// The resource name of the settings.
	Name pulumi.StringPtrInput
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation pulumi.StringPtrInput
	Timeouts        LoggingFolderSettingsTimeoutsPtrInput
}

func (LoggingFolderSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingFolderSettingsState)(nil)).Elem()
}

type loggingFolderSettingsArgs struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink *bool `pulumi:"disableDefaultSink"`
	// The folder for which to retrieve settings.
	Folder string `pulumi:"folder"`
	// The resource name for the configured Cloud KMS key.
	KmsKeyName              *string `pulumi:"kmsKeyName"`
	LoggingFolderSettingsId *string `pulumi:"loggingFolderSettingsId"`
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation *string                        `pulumi:"storageLocation"`
	Timeouts        *LoggingFolderSettingsTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a LoggingFolderSettings resource.
type LoggingFolderSettingsArgs struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink pulumi.BoolPtrInput
	// The folder for which to retrieve settings.
	Folder pulumi.StringInput
	// The resource name for the configured Cloud KMS key.
	KmsKeyName              pulumi.StringPtrInput
	LoggingFolderSettingsId pulumi.StringPtrInput
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation pulumi.StringPtrInput
	Timeouts        LoggingFolderSettingsTimeoutsPtrInput
}

func (LoggingFolderSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingFolderSettingsArgs)(nil)).Elem()
}

type LoggingFolderSettingsInput interface {
	pulumi.Input

	ToLoggingFolderSettingsOutput() LoggingFolderSettingsOutput
	ToLoggingFolderSettingsOutputWithContext(ctx context.Context) LoggingFolderSettingsOutput
}

func (*LoggingFolderSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingFolderSettings)(nil)).Elem()
}

func (i *LoggingFolderSettings) ToLoggingFolderSettingsOutput() LoggingFolderSettingsOutput {
	return i.ToLoggingFolderSettingsOutputWithContext(context.Background())
}

func (i *LoggingFolderSettings) ToLoggingFolderSettingsOutputWithContext(ctx context.Context) LoggingFolderSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingFolderSettingsOutput)
}

type LoggingFolderSettingsOutput struct{ *pulumi.OutputState }

func (LoggingFolderSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingFolderSettings)(nil)).Elem()
}

func (o LoggingFolderSettingsOutput) ToLoggingFolderSettingsOutput() LoggingFolderSettingsOutput {
	return o
}

func (o LoggingFolderSettingsOutput) ToLoggingFolderSettingsOutputWithContext(ctx context.Context) LoggingFolderSettingsOutput {
	return o
}

// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
// _Default sink can be re-enabled manually if needed.
func (o LoggingFolderSettingsOutput) DisableDefaultSink() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.BoolOutput { return v.DisableDefaultSink }).(pulumi.BoolOutput)
}

// The folder for which to retrieve settings.
func (o LoggingFolderSettingsOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

// The resource name for the configured Cloud KMS key.
func (o LoggingFolderSettingsOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.StringOutput { return v.KmsKeyName }).(pulumi.StringOutput)
}

// The service account that will be used by the Log Router to access your Cloud KMS key.
func (o LoggingFolderSettingsOutput) KmsServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.StringOutput { return v.KmsServiceAccountId }).(pulumi.StringOutput)
}

func (o LoggingFolderSettingsOutput) LoggingFolderSettingsId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.StringOutput { return v.LoggingFolderSettingsId }).(pulumi.StringOutput)
}

// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
// account is provided.
func (o LoggingFolderSettingsOutput) LoggingServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.StringOutput { return v.LoggingServiceAccountId }).(pulumi.StringOutput)
}

// The resource name of the settings.
func (o LoggingFolderSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
// provided.
func (o LoggingFolderSettingsOutput) StorageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) pulumi.StringOutput { return v.StorageLocation }).(pulumi.StringOutput)
}

func (o LoggingFolderSettingsOutput) Timeouts() LoggingFolderSettingsTimeoutsPtrOutput {
	return o.ApplyT(func(v *LoggingFolderSettings) LoggingFolderSettingsTimeoutsPtrOutput { return v.Timeouts }).(LoggingFolderSettingsTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingFolderSettingsInput)(nil)).Elem(), &LoggingFolderSettings{})
	pulumi.RegisterOutputType(LoggingFolderSettingsOutput{})
}
