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

type SccManagementOrganizationEventThreatDetectionCustomModule struct {
	pulumi.CustomResourceState

	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrOutput `pulumi:"enablementState"`
	// The editor that last updated the custom module
	LastEditor pulumi.StringOutput `pulumi:"lastEditor"`
	// Location ID of the parent organization. Only global is supported at the moment.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name of the Event Threat Detection custom module. Its format is
	// "organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventThreatDetectionCustomModule}".
	Name pulumi.StringOutput `pulumi:"name"`
	// Numerical ID of the parent organization.
	Organization                                                pulumi.StringOutput                                                        `pulumi:"organization"`
	SccManagementOrganizationEventThreatDetectionCustomModuleId pulumi.StringOutput                                                        `pulumi:"sccManagementOrganizationEventThreatDetectionCustomModuleId"`
	Timeouts                                                    SccManagementOrganizationEventThreatDetectionCustomModuleTimeoutsPtrOutput `pulumi:"timeouts"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSccManagementOrganizationEventThreatDetectionCustomModule registers a new resource with the given unique name, arguments, and options.
func NewSccManagementOrganizationEventThreatDetectionCustomModule(ctx *pulumi.Context,
	name string, args *SccManagementOrganizationEventThreatDetectionCustomModuleArgs, opts ...pulumi.ResourceOption) (*SccManagementOrganizationEventThreatDetectionCustomModule, error) {
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
	var resource SccManagementOrganizationEventThreatDetectionCustomModule
	err = ctx.RegisterPackageResource("google-beta:index/sccManagementOrganizationEventThreatDetectionCustomModule:SccManagementOrganizationEventThreatDetectionCustomModule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccManagementOrganizationEventThreatDetectionCustomModule gets an existing SccManagementOrganizationEventThreatDetectionCustomModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccManagementOrganizationEventThreatDetectionCustomModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccManagementOrganizationEventThreatDetectionCustomModuleState, opts ...pulumi.ResourceOption) (*SccManagementOrganizationEventThreatDetectionCustomModule, error) {
	var resource SccManagementOrganizationEventThreatDetectionCustomModule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/sccManagementOrganizationEventThreatDetectionCustomModule:SccManagementOrganizationEventThreatDetectionCustomModule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccManagementOrganizationEventThreatDetectionCustomModule resources.
type sccManagementOrganizationEventThreatDetectionCustomModuleState struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config *string `pulumi:"config"`
	// The human readable name to be displayed for the module.
	DisplayName *string `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState *string `pulumi:"enablementState"`
	// The editor that last updated the custom module
	LastEditor *string `pulumi:"lastEditor"`
	// Location ID of the parent organization. Only global is supported at the moment.
	Location *string `pulumi:"location"`
	// The resource name of the Event Threat Detection custom module. Its format is
	// "organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventThreatDetectionCustomModule}".
	Name *string `pulumi:"name"`
	// Numerical ID of the parent organization.
	Organization                                                *string                                                            `pulumi:"organization"`
	SccManagementOrganizationEventThreatDetectionCustomModuleId *string                                                            `pulumi:"sccManagementOrganizationEventThreatDetectionCustomModuleId"`
	Timeouts                                                    *SccManagementOrganizationEventThreatDetectionCustomModuleTimeouts `pulumi:"timeouts"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type *string `pulumi:"type"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type SccManagementOrganizationEventThreatDetectionCustomModuleState struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config pulumi.StringPtrInput
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringPtrInput
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrInput
	// The editor that last updated the custom module
	LastEditor pulumi.StringPtrInput
	// Location ID of the parent organization. Only global is supported at the moment.
	Location pulumi.StringPtrInput
	// The resource name of the Event Threat Detection custom module. Its format is
	// "organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventThreatDetectionCustomModule}".
	Name pulumi.StringPtrInput
	// Numerical ID of the parent organization.
	Organization                                                pulumi.StringPtrInput
	SccManagementOrganizationEventThreatDetectionCustomModuleId pulumi.StringPtrInput
	Timeouts                                                    SccManagementOrganizationEventThreatDetectionCustomModuleTimeoutsPtrInput
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringPtrInput
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (SccManagementOrganizationEventThreatDetectionCustomModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccManagementOrganizationEventThreatDetectionCustomModuleState)(nil)).Elem()
}

type sccManagementOrganizationEventThreatDetectionCustomModuleArgs struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config *string `pulumi:"config"`
	// The human readable name to be displayed for the module.
	DisplayName *string `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState *string `pulumi:"enablementState"`
	// Location ID of the parent organization. Only global is supported at the moment.
	Location *string `pulumi:"location"`
	// Numerical ID of the parent organization.
	Organization                                                string                                                             `pulumi:"organization"`
	SccManagementOrganizationEventThreatDetectionCustomModuleId *string                                                            `pulumi:"sccManagementOrganizationEventThreatDetectionCustomModuleId"`
	Timeouts                                                    *SccManagementOrganizationEventThreatDetectionCustomModuleTimeouts `pulumi:"timeouts"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SccManagementOrganizationEventThreatDetectionCustomModule resource.
type SccManagementOrganizationEventThreatDetectionCustomModuleArgs struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
	// config value is inherited from the ancestor module.
	Config pulumi.StringPtrInput
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringPtrInput
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrInput
	// Location ID of the parent organization. Only global is supported at the moment.
	Location pulumi.StringPtrInput
	// Numerical ID of the parent organization.
	Organization                                                pulumi.StringInput
	SccManagementOrganizationEventThreatDetectionCustomModuleId pulumi.StringPtrInput
	Timeouts                                                    SccManagementOrganizationEventThreatDetectionCustomModuleTimeoutsPtrInput
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringPtrInput
}

func (SccManagementOrganizationEventThreatDetectionCustomModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccManagementOrganizationEventThreatDetectionCustomModuleArgs)(nil)).Elem()
}

type SccManagementOrganizationEventThreatDetectionCustomModuleInput interface {
	pulumi.Input

	ToSccManagementOrganizationEventThreatDetectionCustomModuleOutput() SccManagementOrganizationEventThreatDetectionCustomModuleOutput
	ToSccManagementOrganizationEventThreatDetectionCustomModuleOutputWithContext(ctx context.Context) SccManagementOrganizationEventThreatDetectionCustomModuleOutput
}

func (*SccManagementOrganizationEventThreatDetectionCustomModule) ElementType() reflect.Type {
	return reflect.TypeOf((**SccManagementOrganizationEventThreatDetectionCustomModule)(nil)).Elem()
}

func (i *SccManagementOrganizationEventThreatDetectionCustomModule) ToSccManagementOrganizationEventThreatDetectionCustomModuleOutput() SccManagementOrganizationEventThreatDetectionCustomModuleOutput {
	return i.ToSccManagementOrganizationEventThreatDetectionCustomModuleOutputWithContext(context.Background())
}

func (i *SccManagementOrganizationEventThreatDetectionCustomModule) ToSccManagementOrganizationEventThreatDetectionCustomModuleOutputWithContext(ctx context.Context) SccManagementOrganizationEventThreatDetectionCustomModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccManagementOrganizationEventThreatDetectionCustomModuleOutput)
}

type SccManagementOrganizationEventThreatDetectionCustomModuleOutput struct{ *pulumi.OutputState }

func (SccManagementOrganizationEventThreatDetectionCustomModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccManagementOrganizationEventThreatDetectionCustomModule)(nil)).Elem()
}

func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) ToSccManagementOrganizationEventThreatDetectionCustomModuleOutput() SccManagementOrganizationEventThreatDetectionCustomModuleOutput {
	return o
}

func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) ToSccManagementOrganizationEventThreatDetectionCustomModuleOutputWithContext(ctx context.Context) SccManagementOrganizationEventThreatDetectionCustomModuleOutput {
	return o
}

// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its
// config value is inherited from the ancestor module.
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringPtrOutput {
		return v.Config
	}).(pulumi.StringPtrOutput)
}

// The human readable name to be displayed for the module.
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringPtrOutput {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) EnablementState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringPtrOutput {
		return v.EnablementState
	}).(pulumi.StringPtrOutput)
}

// The editor that last updated the custom module
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringOutput {
		return v.LastEditor
	}).(pulumi.StringOutput)
}

// Location ID of the parent organization. Only global is supported at the moment.
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringPtrOutput {
		return v.Location
	}).(pulumi.StringPtrOutput)
}

// The resource name of the Event Threat Detection custom module. Its format is
// "organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventThreatDetectionCustomModule}".
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Numerical ID of the parent organization.
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringOutput {
		return v.Organization
	}).(pulumi.StringOutput)
}

func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) SccManagementOrganizationEventThreatDetectionCustomModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringOutput {
		return v.SccManagementOrganizationEventThreatDetectionCustomModuleId
	}).(pulumi.StringOutput)
}

func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) Timeouts() SccManagementOrganizationEventThreatDetectionCustomModuleTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) SccManagementOrganizationEventThreatDetectionCustomModuleTimeoutsPtrOutput {
		return v.Timeouts
	}).(SccManagementOrganizationEventThreatDetectionCustomModuleTimeoutsPtrOutput)
}

// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringPtrOutput {
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccManagementOrganizationEventThreatDetectionCustomModuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementOrganizationEventThreatDetectionCustomModule) pulumi.StringOutput {
		return v.UpdateTime
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccManagementOrganizationEventThreatDetectionCustomModuleInput)(nil)).Elem(), &SccManagementOrganizationEventThreatDetectionCustomModule{})
	pulumi.RegisterOutputType(SccManagementOrganizationEventThreatDetectionCustomModuleOutput{})
}
