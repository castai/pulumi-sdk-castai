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

type LoggingOrganizationSettings struct {
	pulumi.CustomResourceState

	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink pulumi.BoolOutput `pulumi:"disableDefaultSink"`
	// The resource name for the configured Cloud KMS key.
	KmsKeyName pulumi.StringOutput `pulumi:"kmsKeyName"`
	// The service account that will be used by the Log Router to access your Cloud KMS key.
	KmsServiceAccountId           pulumi.StringOutput `pulumi:"kmsServiceAccountId"`
	LoggingOrganizationSettingsId pulumi.StringOutput `pulumi:"loggingOrganizationSettingsId"`
	// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
	// account is provided.
	LoggingServiceAccountId pulumi.StringOutput `pulumi:"loggingServiceAccountId"`
	// The resource name of the settings.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization for which to retrieve or configure settings.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation pulumi.StringOutput                          `pulumi:"storageLocation"`
	Timeouts        LoggingOrganizationSettingsTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewLoggingOrganizationSettings registers a new resource with the given unique name, arguments, and options.
func NewLoggingOrganizationSettings(ctx *pulumi.Context,
	name string, args *LoggingOrganizationSettingsArgs, opts ...pulumi.ResourceOption) (*LoggingOrganizationSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource LoggingOrganizationSettings
	err = ctx.RegisterPackageResource("google:index/loggingOrganizationSettings:LoggingOrganizationSettings", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingOrganizationSettings gets an existing LoggingOrganizationSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingOrganizationSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingOrganizationSettingsState, opts ...pulumi.ResourceOption) (*LoggingOrganizationSettings, error) {
	var resource LoggingOrganizationSettings
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/loggingOrganizationSettings:LoggingOrganizationSettings", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingOrganizationSettings resources.
type loggingOrganizationSettingsState struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink *bool `pulumi:"disableDefaultSink"`
	// The resource name for the configured Cloud KMS key.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// The service account that will be used by the Log Router to access your Cloud KMS key.
	KmsServiceAccountId           *string `pulumi:"kmsServiceAccountId"`
	LoggingOrganizationSettingsId *string `pulumi:"loggingOrganizationSettingsId"`
	// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
	// account is provided.
	LoggingServiceAccountId *string `pulumi:"loggingServiceAccountId"`
	// The resource name of the settings.
	Name *string `pulumi:"name"`
	// The organization for which to retrieve or configure settings.
	Organization *string `pulumi:"organization"`
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation *string                              `pulumi:"storageLocation"`
	Timeouts        *LoggingOrganizationSettingsTimeouts `pulumi:"timeouts"`
}

type LoggingOrganizationSettingsState struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink pulumi.BoolPtrInput
	// The resource name for the configured Cloud KMS key.
	KmsKeyName pulumi.StringPtrInput
	// The service account that will be used by the Log Router to access your Cloud KMS key.
	KmsServiceAccountId           pulumi.StringPtrInput
	LoggingOrganizationSettingsId pulumi.StringPtrInput
	// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
	// account is provided.
	LoggingServiceAccountId pulumi.StringPtrInput
	// The resource name of the settings.
	Name pulumi.StringPtrInput
	// The organization for which to retrieve or configure settings.
	Organization pulumi.StringPtrInput
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation pulumi.StringPtrInput
	Timeouts        LoggingOrganizationSettingsTimeoutsPtrInput
}

func (LoggingOrganizationSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingOrganizationSettingsState)(nil)).Elem()
}

type loggingOrganizationSettingsArgs struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink *bool `pulumi:"disableDefaultSink"`
	// The resource name for the configured Cloud KMS key.
	KmsKeyName                    *string `pulumi:"kmsKeyName"`
	LoggingOrganizationSettingsId *string `pulumi:"loggingOrganizationSettingsId"`
	// The organization for which to retrieve or configure settings.
	Organization string `pulumi:"organization"`
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation *string                              `pulumi:"storageLocation"`
	Timeouts        *LoggingOrganizationSettingsTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a LoggingOrganizationSettings resource.
type LoggingOrganizationSettingsArgs struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
	// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
	// _Default sink can be re-enabled manually if needed.
	DisableDefaultSink pulumi.BoolPtrInput
	// The resource name for the configured Cloud KMS key.
	KmsKeyName                    pulumi.StringPtrInput
	LoggingOrganizationSettingsId pulumi.StringPtrInput
	// The organization for which to retrieve or configure settings.
	Organization pulumi.StringInput
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
	// provided.
	StorageLocation pulumi.StringPtrInput
	Timeouts        LoggingOrganizationSettingsTimeoutsPtrInput
}

func (LoggingOrganizationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingOrganizationSettingsArgs)(nil)).Elem()
}

type LoggingOrganizationSettingsInput interface {
	pulumi.Input

	ToLoggingOrganizationSettingsOutput() LoggingOrganizationSettingsOutput
	ToLoggingOrganizationSettingsOutputWithContext(ctx context.Context) LoggingOrganizationSettingsOutput
}

func (*LoggingOrganizationSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingOrganizationSettings)(nil)).Elem()
}

func (i *LoggingOrganizationSettings) ToLoggingOrganizationSettingsOutput() LoggingOrganizationSettingsOutput {
	return i.ToLoggingOrganizationSettingsOutputWithContext(context.Background())
}

func (i *LoggingOrganizationSettings) ToLoggingOrganizationSettingsOutputWithContext(ctx context.Context) LoggingOrganizationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingOrganizationSettingsOutput)
}

type LoggingOrganizationSettingsOutput struct{ *pulumi.OutputState }

func (LoggingOrganizationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingOrganizationSettings)(nil)).Elem()
}

func (o LoggingOrganizationSettingsOutput) ToLoggingOrganizationSettingsOutput() LoggingOrganizationSettingsOutput {
	return o
}

func (o LoggingOrganizationSettingsOutput) ToLoggingOrganizationSettingsOutputWithContext(ctx context.Context) LoggingOrganizationSettingsOutput {
	return o
}

// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be
// used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The
// _Default sink can be re-enabled manually if needed.
func (o LoggingOrganizationSettingsOutput) DisableDefaultSink() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.BoolOutput { return v.DisableDefaultSink }).(pulumi.BoolOutput)
}

// The resource name for the configured Cloud KMS key.
func (o LoggingOrganizationSettingsOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.StringOutput { return v.KmsKeyName }).(pulumi.StringOutput)
}

// The service account that will be used by the Log Router to access your Cloud KMS key.
func (o LoggingOrganizationSettingsOutput) KmsServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.StringOutput { return v.KmsServiceAccountId }).(pulumi.StringOutput)
}

func (o LoggingOrganizationSettingsOutput) LoggingOrganizationSettingsId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.StringOutput { return v.LoggingOrganizationSettingsId }).(pulumi.StringOutput)
}

// The service account for the given container. Sinks use this service account as their writerIdentity if no custom service
// account is provided.
func (o LoggingOrganizationSettingsOutput) LoggingServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.StringOutput { return v.LoggingServiceAccountId }).(pulumi.StringOutput)
}

// The resource name of the settings.
func (o LoggingOrganizationSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization for which to retrieve or configure settings.
func (o LoggingOrganizationSettingsOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly
// provided.
func (o LoggingOrganizationSettingsOutput) StorageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) pulumi.StringOutput { return v.StorageLocation }).(pulumi.StringOutput)
}

func (o LoggingOrganizationSettingsOutput) Timeouts() LoggingOrganizationSettingsTimeoutsPtrOutput {
	return o.ApplyT(func(v *LoggingOrganizationSettings) LoggingOrganizationSettingsTimeoutsPtrOutput { return v.Timeouts }).(LoggingOrganizationSettingsTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingOrganizationSettingsInput)(nil)).Elem(), &LoggingOrganizationSettings{})
	pulumi.RegisterOutputType(LoggingOrganizationSettingsOutput{})
}
