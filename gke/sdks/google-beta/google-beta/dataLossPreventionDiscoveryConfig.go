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

type DataLossPreventionDiscoveryConfig struct {
	pulumi.CustomResourceState

	// Actions to execute at the completion of scanning
	Actions DataLossPreventionDiscoveryConfigActionArrayOutput `pulumi:"actions"`
	// Output only. The creation timestamp of a DiscoveryConfig.
	CreateTime                          pulumi.StringOutput `pulumi:"createTime"`
	DataLossPreventionDiscoveryConfigId pulumi.StringOutput `pulumi:"dataLossPreventionDiscoveryConfigId"`
	// Display Name (max 1000 Chars)
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Output only. A stream of errors encountered when the config was activated. Repeated errors may result in the config
	// automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this
	// list will be cleared.
	Errors DataLossPreventionDiscoveryConfigErrorArrayOutput `pulumi:"errors"`
	// Detection logic for profile generation
	InspectTemplates pulumi.StringArrayOutput `pulumi:"inspectTemplates"`
	// Output only. The timestamp of the last time this config was executed
	LastRunTime pulumi.StringOutput `pulumi:"lastRunTime"`
	// Location to create the discovery config in.
	Location pulumi.StringOutput `pulumi:"location"`
	// Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A nested object resource.
	OrgConfig DataLossPreventionDiscoveryConfigOrgConfigPtrOutput `pulumi:"orgConfig"`
	// The parent of the discovery config in any of the following formats: * 'projects/{{project}}/locations/{{location}}' *
	// 'organizations/{{organization_id}}/locations/{{location}}'
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Required. A status for this configuration Possible values: ["RUNNING", "PAUSED"]
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Target to match against for determining what to scan and how frequently
	Targets  DataLossPreventionDiscoveryConfigTargetArrayOutput `pulumi:"targets"`
	Timeouts DataLossPreventionDiscoveryConfigTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. The last update timestamp of a DiscoveryConfig.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDataLossPreventionDiscoveryConfig registers a new resource with the given unique name, arguments, and options.
func NewDataLossPreventionDiscoveryConfig(ctx *pulumi.Context,
	name string, args *DataLossPreventionDiscoveryConfigArgs, opts ...pulumi.ResourceOption) (*DataLossPreventionDiscoveryConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataLossPreventionDiscoveryConfig
	err = ctx.RegisterPackageResource("google-beta:index/dataLossPreventionDiscoveryConfig:DataLossPreventionDiscoveryConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataLossPreventionDiscoveryConfig gets an existing DataLossPreventionDiscoveryConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataLossPreventionDiscoveryConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataLossPreventionDiscoveryConfigState, opts ...pulumi.ResourceOption) (*DataLossPreventionDiscoveryConfig, error) {
	var resource DataLossPreventionDiscoveryConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataLossPreventionDiscoveryConfig:DataLossPreventionDiscoveryConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataLossPreventionDiscoveryConfig resources.
type dataLossPreventionDiscoveryConfigState struct {
	// Actions to execute at the completion of scanning
	Actions []DataLossPreventionDiscoveryConfigAction `pulumi:"actions"`
	// Output only. The creation timestamp of a DiscoveryConfig.
	CreateTime                          *string `pulumi:"createTime"`
	DataLossPreventionDiscoveryConfigId *string `pulumi:"dataLossPreventionDiscoveryConfigId"`
	// Display Name (max 1000 Chars)
	DisplayName *string `pulumi:"displayName"`
	// Output only. A stream of errors encountered when the config was activated. Repeated errors may result in the config
	// automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this
	// list will be cleared.
	Errors []DataLossPreventionDiscoveryConfigError `pulumi:"errors"`
	// Detection logic for profile generation
	InspectTemplates []string `pulumi:"inspectTemplates"`
	// Output only. The timestamp of the last time this config was executed
	LastRunTime *string `pulumi:"lastRunTime"`
	// Location to create the discovery config in.
	Location *string `pulumi:"location"`
	// Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created.
	Name *string `pulumi:"name"`
	// A nested object resource.
	OrgConfig *DataLossPreventionDiscoveryConfigOrgConfig `pulumi:"orgConfig"`
	// The parent of the discovery config in any of the following formats: * 'projects/{{project}}/locations/{{location}}' *
	// 'organizations/{{organization_id}}/locations/{{location}}'
	Parent *string `pulumi:"parent"`
	// Required. A status for this configuration Possible values: ["RUNNING", "PAUSED"]
	Status *string `pulumi:"status"`
	// Target to match against for determining what to scan and how frequently
	Targets  []DataLossPreventionDiscoveryConfigTarget  `pulumi:"targets"`
	Timeouts *DataLossPreventionDiscoveryConfigTimeouts `pulumi:"timeouts"`
	// Output only. The last update timestamp of a DiscoveryConfig.
	UpdateTime *string `pulumi:"updateTime"`
}

type DataLossPreventionDiscoveryConfigState struct {
	// Actions to execute at the completion of scanning
	Actions DataLossPreventionDiscoveryConfigActionArrayInput
	// Output only. The creation timestamp of a DiscoveryConfig.
	CreateTime                          pulumi.StringPtrInput
	DataLossPreventionDiscoveryConfigId pulumi.StringPtrInput
	// Display Name (max 1000 Chars)
	DisplayName pulumi.StringPtrInput
	// Output only. A stream of errors encountered when the config was activated. Repeated errors may result in the config
	// automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this
	// list will be cleared.
	Errors DataLossPreventionDiscoveryConfigErrorArrayInput
	// Detection logic for profile generation
	InspectTemplates pulumi.StringArrayInput
	// Output only. The timestamp of the last time this config was executed
	LastRunTime pulumi.StringPtrInput
	// Location to create the discovery config in.
	Location pulumi.StringPtrInput
	// Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created.
	Name pulumi.StringPtrInput
	// A nested object resource.
	OrgConfig DataLossPreventionDiscoveryConfigOrgConfigPtrInput
	// The parent of the discovery config in any of the following formats: * 'projects/{{project}}/locations/{{location}}' *
	// 'organizations/{{organization_id}}/locations/{{location}}'
	Parent pulumi.StringPtrInput
	// Required. A status for this configuration Possible values: ["RUNNING", "PAUSED"]
	Status pulumi.StringPtrInput
	// Target to match against for determining what to scan and how frequently
	Targets  DataLossPreventionDiscoveryConfigTargetArrayInput
	Timeouts DataLossPreventionDiscoveryConfigTimeoutsPtrInput
	// Output only. The last update timestamp of a DiscoveryConfig.
	UpdateTime pulumi.StringPtrInput
}

func (DataLossPreventionDiscoveryConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLossPreventionDiscoveryConfigState)(nil)).Elem()
}

type dataLossPreventionDiscoveryConfigArgs struct {
	// Actions to execute at the completion of scanning
	Actions                             []DataLossPreventionDiscoveryConfigAction `pulumi:"actions"`
	DataLossPreventionDiscoveryConfigId *string                                   `pulumi:"dataLossPreventionDiscoveryConfigId"`
	// Display Name (max 1000 Chars)
	DisplayName *string `pulumi:"displayName"`
	// Detection logic for profile generation
	InspectTemplates []string `pulumi:"inspectTemplates"`
	// Location to create the discovery config in.
	Location string `pulumi:"location"`
	// A nested object resource.
	OrgConfig *DataLossPreventionDiscoveryConfigOrgConfig `pulumi:"orgConfig"`
	// The parent of the discovery config in any of the following formats: * 'projects/{{project}}/locations/{{location}}' *
	// 'organizations/{{organization_id}}/locations/{{location}}'
	Parent string `pulumi:"parent"`
	// Required. A status for this configuration Possible values: ["RUNNING", "PAUSED"]
	Status *string `pulumi:"status"`
	// Target to match against for determining what to scan and how frequently
	Targets  []DataLossPreventionDiscoveryConfigTarget  `pulumi:"targets"`
	Timeouts *DataLossPreventionDiscoveryConfigTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DataLossPreventionDiscoveryConfig resource.
type DataLossPreventionDiscoveryConfigArgs struct {
	// Actions to execute at the completion of scanning
	Actions                             DataLossPreventionDiscoveryConfigActionArrayInput
	DataLossPreventionDiscoveryConfigId pulumi.StringPtrInput
	// Display Name (max 1000 Chars)
	DisplayName pulumi.StringPtrInput
	// Detection logic for profile generation
	InspectTemplates pulumi.StringArrayInput
	// Location to create the discovery config in.
	Location pulumi.StringInput
	// A nested object resource.
	OrgConfig DataLossPreventionDiscoveryConfigOrgConfigPtrInput
	// The parent of the discovery config in any of the following formats: * 'projects/{{project}}/locations/{{location}}' *
	// 'organizations/{{organization_id}}/locations/{{location}}'
	Parent pulumi.StringInput
	// Required. A status for this configuration Possible values: ["RUNNING", "PAUSED"]
	Status pulumi.StringPtrInput
	// Target to match against for determining what to scan and how frequently
	Targets  DataLossPreventionDiscoveryConfigTargetArrayInput
	Timeouts DataLossPreventionDiscoveryConfigTimeoutsPtrInput
}

func (DataLossPreventionDiscoveryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLossPreventionDiscoveryConfigArgs)(nil)).Elem()
}

type DataLossPreventionDiscoveryConfigInput interface {
	pulumi.Input

	ToDataLossPreventionDiscoveryConfigOutput() DataLossPreventionDiscoveryConfigOutput
	ToDataLossPreventionDiscoveryConfigOutputWithContext(ctx context.Context) DataLossPreventionDiscoveryConfigOutput
}

func (*DataLossPreventionDiscoveryConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLossPreventionDiscoveryConfig)(nil)).Elem()
}

func (i *DataLossPreventionDiscoveryConfig) ToDataLossPreventionDiscoveryConfigOutput() DataLossPreventionDiscoveryConfigOutput {
	return i.ToDataLossPreventionDiscoveryConfigOutputWithContext(context.Background())
}

func (i *DataLossPreventionDiscoveryConfig) ToDataLossPreventionDiscoveryConfigOutputWithContext(ctx context.Context) DataLossPreventionDiscoveryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLossPreventionDiscoveryConfigOutput)
}

type DataLossPreventionDiscoveryConfigOutput struct{ *pulumi.OutputState }

func (DataLossPreventionDiscoveryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLossPreventionDiscoveryConfig)(nil)).Elem()
}

func (o DataLossPreventionDiscoveryConfigOutput) ToDataLossPreventionDiscoveryConfigOutput() DataLossPreventionDiscoveryConfigOutput {
	return o
}

func (o DataLossPreventionDiscoveryConfigOutput) ToDataLossPreventionDiscoveryConfigOutputWithContext(ctx context.Context) DataLossPreventionDiscoveryConfigOutput {
	return o
}

// Actions to execute at the completion of scanning
func (o DataLossPreventionDiscoveryConfigOutput) Actions() DataLossPreventionDiscoveryConfigActionArrayOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) DataLossPreventionDiscoveryConfigActionArrayOutput {
		return v.Actions
	}).(DataLossPreventionDiscoveryConfigActionArrayOutput)
}

// Output only. The creation timestamp of a DiscoveryConfig.
func (o DataLossPreventionDiscoveryConfigOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o DataLossPreventionDiscoveryConfigOutput) DataLossPreventionDiscoveryConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringOutput {
		return v.DataLossPreventionDiscoveryConfigId
	}).(pulumi.StringOutput)
}

// Display Name (max 1000 Chars)
func (o DataLossPreventionDiscoveryConfigOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Output only. A stream of errors encountered when the config was activated. Repeated errors may result in the config
// automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this
// list will be cleared.
func (o DataLossPreventionDiscoveryConfigOutput) Errors() DataLossPreventionDiscoveryConfigErrorArrayOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) DataLossPreventionDiscoveryConfigErrorArrayOutput {
		return v.Errors
	}).(DataLossPreventionDiscoveryConfigErrorArrayOutput)
}

// Detection logic for profile generation
func (o DataLossPreventionDiscoveryConfigOutput) InspectTemplates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringArrayOutput { return v.InspectTemplates }).(pulumi.StringArrayOutput)
}

// Output only. The timestamp of the last time this config was executed
func (o DataLossPreventionDiscoveryConfigOutput) LastRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringOutput { return v.LastRunTime }).(pulumi.StringOutput)
}

// Location to create the discovery config in.
func (o DataLossPreventionDiscoveryConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created.
func (o DataLossPreventionDiscoveryConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A nested object resource.
func (o DataLossPreventionDiscoveryConfigOutput) OrgConfig() DataLossPreventionDiscoveryConfigOrgConfigPtrOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) DataLossPreventionDiscoveryConfigOrgConfigPtrOutput {
		return v.OrgConfig
	}).(DataLossPreventionDiscoveryConfigOrgConfigPtrOutput)
}

// The parent of the discovery config in any of the following formats: * 'projects/{{project}}/locations/{{location}}' *
// 'organizations/{{organization_id}}/locations/{{location}}'
func (o DataLossPreventionDiscoveryConfigOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Required. A status for this configuration Possible values: ["RUNNING", "PAUSED"]
func (o DataLossPreventionDiscoveryConfigOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Target to match against for determining what to scan and how frequently
func (o DataLossPreventionDiscoveryConfigOutput) Targets() DataLossPreventionDiscoveryConfigTargetArrayOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) DataLossPreventionDiscoveryConfigTargetArrayOutput {
		return v.Targets
	}).(DataLossPreventionDiscoveryConfigTargetArrayOutput)
}

func (o DataLossPreventionDiscoveryConfigOutput) Timeouts() DataLossPreventionDiscoveryConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) DataLossPreventionDiscoveryConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(DataLossPreventionDiscoveryConfigTimeoutsPtrOutput)
}

// Output only. The last update timestamp of a DiscoveryConfig.
func (o DataLossPreventionDiscoveryConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLossPreventionDiscoveryConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataLossPreventionDiscoveryConfigInput)(nil)).Elem(), &DataLossPreventionDiscoveryConfig{})
	pulumi.RegisterOutputType(DataLossPreventionDiscoveryConfigOutput{})
}
