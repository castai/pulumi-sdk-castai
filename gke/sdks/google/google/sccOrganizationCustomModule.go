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

type SccOrganizationCustomModule struct {
	pulumi.CustomResourceState

	// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing
	// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
	// inherited.
	AncestorModule pulumi.StringOutput `pulumi:"ancestorModule"`
	// The user specified custom configuration for the module.
	CustomConfig SccOrganizationCustomModuleCustomConfigOutput `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringOutput `pulumi:"enablementState"`
	// The editor that last updated the custom module.
	LastEditor pulumi.StringOutput `pulumi:"lastEditor"`
	// The resource name of the custom module. Its format is
	// "organizations/{org_id}/securityHealthAnalyticsSettings/customModules/{customModule}". The id {customModule} is
	// server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name pulumi.StringOutput `pulumi:"name"`
	// Numerical ID of the parent organization.
	Organization                  pulumi.StringOutput                          `pulumi:"organization"`
	SccOrganizationCustomModuleId pulumi.StringOutput                          `pulumi:"sccOrganizationCustomModuleId"`
	Timeouts                      SccOrganizationCustomModuleTimeoutsPtrOutput `pulumi:"timeouts"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSccOrganizationCustomModule registers a new resource with the given unique name, arguments, and options.
func NewSccOrganizationCustomModule(ctx *pulumi.Context,
	name string, args *SccOrganizationCustomModuleArgs, opts ...pulumi.ResourceOption) (*SccOrganizationCustomModule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomConfig == nil {
		return nil, errors.New("invalid value for required argument 'CustomConfig'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.EnablementState == nil {
		return nil, errors.New("invalid value for required argument 'EnablementState'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccOrganizationCustomModule
	err = ctx.RegisterPackageResource("google:index/sccOrganizationCustomModule:SccOrganizationCustomModule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccOrganizationCustomModule gets an existing SccOrganizationCustomModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccOrganizationCustomModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccOrganizationCustomModuleState, opts ...pulumi.ResourceOption) (*SccOrganizationCustomModule, error) {
	var resource SccOrganizationCustomModule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/sccOrganizationCustomModule:SccOrganizationCustomModule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccOrganizationCustomModule resources.
type sccOrganizationCustomModuleState struct {
	// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing
	// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
	// inherited.
	AncestorModule *string `pulumi:"ancestorModule"`
	// The user specified custom configuration for the module.
	CustomConfig *SccOrganizationCustomModuleCustomConfig `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName *string `pulumi:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState *string `pulumi:"enablementState"`
	// The editor that last updated the custom module.
	LastEditor *string `pulumi:"lastEditor"`
	// The resource name of the custom module. Its format is
	// "organizations/{org_id}/securityHealthAnalyticsSettings/customModules/{customModule}". The id {customModule} is
	// server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name *string `pulumi:"name"`
	// Numerical ID of the parent organization.
	Organization                  *string                              `pulumi:"organization"`
	SccOrganizationCustomModuleId *string                              `pulumi:"sccOrganizationCustomModuleId"`
	Timeouts                      *SccOrganizationCustomModuleTimeouts `pulumi:"timeouts"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type SccOrganizationCustomModuleState struct {
	// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing
	// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
	// inherited.
	AncestorModule pulumi.StringPtrInput
	// The user specified custom configuration for the module.
	CustomConfig SccOrganizationCustomModuleCustomConfigPtrInput
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringPtrInput
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrInput
	// The editor that last updated the custom module.
	LastEditor pulumi.StringPtrInput
	// The resource name of the custom module. Its format is
	// "organizations/{org_id}/securityHealthAnalyticsSettings/customModules/{customModule}". The id {customModule} is
	// server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name pulumi.StringPtrInput
	// Numerical ID of the parent organization.
	Organization                  pulumi.StringPtrInput
	SccOrganizationCustomModuleId pulumi.StringPtrInput
	Timeouts                      SccOrganizationCustomModuleTimeoutsPtrInput
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (SccOrganizationCustomModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccOrganizationCustomModuleState)(nil)).Elem()
}

type sccOrganizationCustomModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig SccOrganizationCustomModuleCustomConfig `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName string `pulumi:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState string `pulumi:"enablementState"`
	// Numerical ID of the parent organization.
	Organization                  string                               `pulumi:"organization"`
	SccOrganizationCustomModuleId *string                              `pulumi:"sccOrganizationCustomModuleId"`
	Timeouts                      *SccOrganizationCustomModuleTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a SccOrganizationCustomModule resource.
type SccOrganizationCustomModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig SccOrganizationCustomModuleCustomConfigInput
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringInput
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringInput
	// Numerical ID of the parent organization.
	Organization                  pulumi.StringInput
	SccOrganizationCustomModuleId pulumi.StringPtrInput
	Timeouts                      SccOrganizationCustomModuleTimeoutsPtrInput
}

func (SccOrganizationCustomModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccOrganizationCustomModuleArgs)(nil)).Elem()
}

type SccOrganizationCustomModuleInput interface {
	pulumi.Input

	ToSccOrganizationCustomModuleOutput() SccOrganizationCustomModuleOutput
	ToSccOrganizationCustomModuleOutputWithContext(ctx context.Context) SccOrganizationCustomModuleOutput
}

func (*SccOrganizationCustomModule) ElementType() reflect.Type {
	return reflect.TypeOf((**SccOrganizationCustomModule)(nil)).Elem()
}

func (i *SccOrganizationCustomModule) ToSccOrganizationCustomModuleOutput() SccOrganizationCustomModuleOutput {
	return i.ToSccOrganizationCustomModuleOutputWithContext(context.Background())
}

func (i *SccOrganizationCustomModule) ToSccOrganizationCustomModuleOutputWithContext(ctx context.Context) SccOrganizationCustomModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccOrganizationCustomModuleOutput)
}

type SccOrganizationCustomModuleOutput struct{ *pulumi.OutputState }

func (SccOrganizationCustomModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccOrganizationCustomModule)(nil)).Elem()
}

func (o SccOrganizationCustomModuleOutput) ToSccOrganizationCustomModuleOutput() SccOrganizationCustomModuleOutput {
	return o
}

func (o SccOrganizationCustomModuleOutput) ToSccOrganizationCustomModuleOutputWithContext(ctx context.Context) SccOrganizationCustomModuleOutput {
	return o
}

// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing
// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
// inherited.
func (o SccOrganizationCustomModuleOutput) AncestorModule() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.AncestorModule }).(pulumi.StringOutput)
}

// The user specified custom configuration for the module.
func (o SccOrganizationCustomModuleOutput) CustomConfig() SccOrganizationCustomModuleCustomConfigOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) SccOrganizationCustomModuleCustomConfigOutput {
		return v.CustomConfig
	}).(SccOrganizationCustomModuleCustomConfigOutput)
}

// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
// lowercase letter, and contain alphanumeric characters or underscores only.
func (o SccOrganizationCustomModuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
func (o SccOrganizationCustomModuleOutput) EnablementState() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.EnablementState }).(pulumi.StringOutput)
}

// The editor that last updated the custom module.
func (o SccOrganizationCustomModuleOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.LastEditor }).(pulumi.StringOutput)
}

// The resource name of the custom module. Its format is
// "organizations/{org_id}/securityHealthAnalyticsSettings/customModules/{customModule}". The id {customModule} is
// server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
func (o SccOrganizationCustomModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Numerical ID of the parent organization.
func (o SccOrganizationCustomModuleOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

func (o SccOrganizationCustomModuleOutput) SccOrganizationCustomModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.SccOrganizationCustomModuleId }).(pulumi.StringOutput)
}

func (o SccOrganizationCustomModuleOutput) Timeouts() SccOrganizationCustomModuleTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) SccOrganizationCustomModuleTimeoutsPtrOutput { return v.Timeouts }).(SccOrganizationCustomModuleTimeoutsPtrOutput)
}

// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccOrganizationCustomModuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationCustomModule) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccOrganizationCustomModuleInput)(nil)).Elem(), &SccOrganizationCustomModule{})
	pulumi.RegisterOutputType(SccOrganizationCustomModuleOutput{})
}
