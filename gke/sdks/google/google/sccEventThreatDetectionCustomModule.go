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

type SccEventThreatDetectionCustomModule struct {
	pulumi.CustomResourceState

	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config pulumi.StringOutput `pulumi:"config"`
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringOutput `pulumi:"enablementState"`
	// The editor that last updated the custom module
	LastEditor pulumi.StringOutput `pulumi:"lastEditor"`
	// The resource name of the Event Threat Detection custom module. Its format is
	// "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}".
	Name pulumi.StringOutput `pulumi:"name"`
	// Numerical ID of the parent organization.
	Organization                          pulumi.StringOutput                                  `pulumi:"organization"`
	SccEventThreatDetectionCustomModuleId pulumi.StringOutput                                  `pulumi:"sccEventThreatDetectionCustomModuleId"`
	Timeouts                              SccEventThreatDetectionCustomModuleTimeoutsPtrOutput `pulumi:"timeouts"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSccEventThreatDetectionCustomModule registers a new resource with the given unique name, arguments, and options.
func NewSccEventThreatDetectionCustomModule(ctx *pulumi.Context,
	name string, args *SccEventThreatDetectionCustomModuleArgs, opts ...pulumi.ResourceOption) (*SccEventThreatDetectionCustomModule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.EnablementState == nil {
		return nil, errors.New("invalid value for required argument 'EnablementState'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccEventThreatDetectionCustomModule
	err = ctx.RegisterPackageResource("google:index/sccEventThreatDetectionCustomModule:SccEventThreatDetectionCustomModule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccEventThreatDetectionCustomModule gets an existing SccEventThreatDetectionCustomModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccEventThreatDetectionCustomModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccEventThreatDetectionCustomModuleState, opts ...pulumi.ResourceOption) (*SccEventThreatDetectionCustomModule, error) {
	var resource SccEventThreatDetectionCustomModule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/sccEventThreatDetectionCustomModule:SccEventThreatDetectionCustomModule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccEventThreatDetectionCustomModule resources.
type sccEventThreatDetectionCustomModuleState struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config *string `pulumi:"config"`
	// The human readable name to be displayed for the module.
	DisplayName *string `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState *string `pulumi:"enablementState"`
	// The editor that last updated the custom module
	LastEditor *string `pulumi:"lastEditor"`
	// The resource name of the Event Threat Detection custom module. Its format is
	// "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}".
	Name *string `pulumi:"name"`
	// Numerical ID of the parent organization.
	Organization                          *string                                      `pulumi:"organization"`
	SccEventThreatDetectionCustomModuleId *string                                      `pulumi:"sccEventThreatDetectionCustomModuleId"`
	Timeouts                              *SccEventThreatDetectionCustomModuleTimeouts `pulumi:"timeouts"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type *string `pulumi:"type"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type SccEventThreatDetectionCustomModuleState struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config pulumi.StringPtrInput
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringPtrInput
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrInput
	// The editor that last updated the custom module
	LastEditor pulumi.StringPtrInput
	// The resource name of the Event Threat Detection custom module. Its format is
	// "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}".
	Name pulumi.StringPtrInput
	// Numerical ID of the parent organization.
	Organization                          pulumi.StringPtrInput
	SccEventThreatDetectionCustomModuleId pulumi.StringPtrInput
	Timeouts                              SccEventThreatDetectionCustomModuleTimeoutsPtrInput
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringPtrInput
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (SccEventThreatDetectionCustomModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccEventThreatDetectionCustomModuleState)(nil)).Elem()
}

type sccEventThreatDetectionCustomModuleArgs struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config string `pulumi:"config"`
	// The human readable name to be displayed for the module.
	DisplayName *string `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState string `pulumi:"enablementState"`
	// Numerical ID of the parent organization.
	Organization                          string                                       `pulumi:"organization"`
	SccEventThreatDetectionCustomModuleId *string                                      `pulumi:"sccEventThreatDetectionCustomModuleId"`
	Timeouts                              *SccEventThreatDetectionCustomModuleTimeouts `pulumi:"timeouts"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a SccEventThreatDetectionCustomModule resource.
type SccEventThreatDetectionCustomModuleArgs struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config pulumi.StringInput
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringPtrInput
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringInput
	// Numerical ID of the parent organization.
	Organization                          pulumi.StringInput
	SccEventThreatDetectionCustomModuleId pulumi.StringPtrInput
	Timeouts                              SccEventThreatDetectionCustomModuleTimeoutsPtrInput
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringInput
}

func (SccEventThreatDetectionCustomModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccEventThreatDetectionCustomModuleArgs)(nil)).Elem()
}

type SccEventThreatDetectionCustomModuleInput interface {
	pulumi.Input

	ToSccEventThreatDetectionCustomModuleOutput() SccEventThreatDetectionCustomModuleOutput
	ToSccEventThreatDetectionCustomModuleOutputWithContext(ctx context.Context) SccEventThreatDetectionCustomModuleOutput
}

func (*SccEventThreatDetectionCustomModule) ElementType() reflect.Type {
	return reflect.TypeOf((**SccEventThreatDetectionCustomModule)(nil)).Elem()
}

func (i *SccEventThreatDetectionCustomModule) ToSccEventThreatDetectionCustomModuleOutput() SccEventThreatDetectionCustomModuleOutput {
	return i.ToSccEventThreatDetectionCustomModuleOutputWithContext(context.Background())
}

func (i *SccEventThreatDetectionCustomModule) ToSccEventThreatDetectionCustomModuleOutputWithContext(ctx context.Context) SccEventThreatDetectionCustomModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccEventThreatDetectionCustomModuleOutput)
}

type SccEventThreatDetectionCustomModuleOutput struct{ *pulumi.OutputState }

func (SccEventThreatDetectionCustomModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccEventThreatDetectionCustomModule)(nil)).Elem()
}

func (o SccEventThreatDetectionCustomModuleOutput) ToSccEventThreatDetectionCustomModuleOutput() SccEventThreatDetectionCustomModuleOutput {
	return o
}

func (o SccEventThreatDetectionCustomModuleOutput) ToSccEventThreatDetectionCustomModuleOutputWithContext(ctx context.Context) SccEventThreatDetectionCustomModuleOutput {
	return o
}

// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
// config value is inherited from the ancestor module.
func (o SccEventThreatDetectionCustomModuleOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// The human readable name to be displayed for the module.
func (o SccEventThreatDetectionCustomModuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
func (o SccEventThreatDetectionCustomModuleOutput) EnablementState() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput { return v.EnablementState }).(pulumi.StringOutput)
}

// The editor that last updated the custom module
func (o SccEventThreatDetectionCustomModuleOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput { return v.LastEditor }).(pulumi.StringOutput)
}

// The resource name of the Event Threat Detection custom module. Its format is
// "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}".
func (o SccEventThreatDetectionCustomModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Numerical ID of the parent organization.
func (o SccEventThreatDetectionCustomModuleOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

func (o SccEventThreatDetectionCustomModuleOutput) SccEventThreatDetectionCustomModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput {
		return v.SccEventThreatDetectionCustomModuleId
	}).(pulumi.StringOutput)
}

func (o SccEventThreatDetectionCustomModuleOutput) Timeouts() SccEventThreatDetectionCustomModuleTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) SccEventThreatDetectionCustomModuleTimeoutsPtrOutput {
		return v.Timeouts
	}).(SccEventThreatDetectionCustomModuleTimeoutsPtrOutput)
}

// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
func (o SccEventThreatDetectionCustomModuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccEventThreatDetectionCustomModuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccEventThreatDetectionCustomModule) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccEventThreatDetectionCustomModuleInput)(nil)).Elem(), &SccEventThreatDetectionCustomModule{})
	pulumi.RegisterOutputType(SccEventThreatDetectionCustomModuleOutput{})
}
