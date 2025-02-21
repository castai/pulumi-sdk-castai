// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkSecuritySecurityProfileGroup struct {
	pulumi.CustomResourceState

	// Time the security profile group was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Reference to a SecurityProfile with the CustomIntercept configuration.
	CustomInterceptProfile pulumi.StringPtrOutput `pulumi:"customInterceptProfile"`
	// Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.
	CustomMirroringProfile pulumi.StringPtrOutput `pulumi:"customMirroringProfile"`
	// An optional description of the profile. The Max length is 512 characters.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
	// requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location of the security profile group. The default value is 'global'.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the security profile group resource.
	Name                                  pulumi.StringOutput `pulumi:"name"`
	NetworkSecuritySecurityProfileGroupId pulumi.StringOutput `pulumi:"networkSecuritySecurityProfileGroupId"`
	// The name of the parent this security profile group belongs to. Format: organizations/{organization_id}.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput `pulumi:"terraformLabels"`
	// Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.
	ThreatPreventionProfile pulumi.StringPtrOutput                               `pulumi:"threatPreventionProfile"`
	Timeouts                NetworkSecuritySecurityProfileGroupTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the security profile group was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecuritySecurityProfileGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecuritySecurityProfileGroup(ctx *pulumi.Context,
	name string, args *NetworkSecuritySecurityProfileGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecuritySecurityProfileGroup, error) {
	if args == nil {
		args = &NetworkSecuritySecurityProfileGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkSecuritySecurityProfileGroup
	err = ctx.RegisterPackageResource("google:index/networkSecuritySecurityProfileGroup:NetworkSecuritySecurityProfileGroup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecuritySecurityProfileGroup gets an existing NetworkSecuritySecurityProfileGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecuritySecurityProfileGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecuritySecurityProfileGroupState, opts ...pulumi.ResourceOption) (*NetworkSecuritySecurityProfileGroup, error) {
	var resource NetworkSecuritySecurityProfileGroup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/networkSecuritySecurityProfileGroup:NetworkSecuritySecurityProfileGroup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecuritySecurityProfileGroup resources.
type networkSecuritySecurityProfileGroupState struct {
	// Time the security profile group was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Reference to a SecurityProfile with the CustomIntercept configuration.
	CustomInterceptProfile *string `pulumi:"customInterceptProfile"`
	// Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.
	CustomMirroringProfile *string `pulumi:"customMirroringProfile"`
	// An optional description of the profile. The Max length is 512 characters.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
	// requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the security profile group. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The name of the security profile group resource.
	Name                                  *string `pulumi:"name"`
	NetworkSecuritySecurityProfileGroupId *string `pulumi:"networkSecuritySecurityProfileGroupId"`
	// The name of the parent this security profile group belongs to. Format: organizations/{organization_id}.
	Parent *string `pulumi:"parent"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string `pulumi:"terraformLabels"`
	// Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.
	ThreatPreventionProfile *string                                      `pulumi:"threatPreventionProfile"`
	Timeouts                *NetworkSecuritySecurityProfileGroupTimeouts `pulumi:"timeouts"`
	// Time the security profile group was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecuritySecurityProfileGroupState struct {
	// Time the security profile group was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Reference to a SecurityProfile with the CustomIntercept configuration.
	CustomInterceptProfile pulumi.StringPtrInput
	// Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.
	CustomMirroringProfile pulumi.StringPtrInput
	// An optional description of the profile. The Max length is 512 characters.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
	// requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location of the security profile group. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The name of the security profile group resource.
	Name                                  pulumi.StringPtrInput
	NetworkSecuritySecurityProfileGroupId pulumi.StringPtrInput
	// The name of the parent this security profile group belongs to. Format: organizations/{organization_id}.
	Parent pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	// Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.
	ThreatPreventionProfile pulumi.StringPtrInput
	Timeouts                NetworkSecuritySecurityProfileGroupTimeoutsPtrInput
	// Time the security profile group was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecuritySecurityProfileGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecuritySecurityProfileGroupState)(nil)).Elem()
}

type networkSecuritySecurityProfileGroupArgs struct {
	// Reference to a SecurityProfile with the CustomIntercept configuration.
	CustomInterceptProfile *string `pulumi:"customInterceptProfile"`
	// Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.
	CustomMirroringProfile *string `pulumi:"customMirroringProfile"`
	// An optional description of the profile. The Max length is 512 characters.
	Description *string `pulumi:"description"`
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the security profile group. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The name of the security profile group resource.
	Name                                  *string `pulumi:"name"`
	NetworkSecuritySecurityProfileGroupId *string `pulumi:"networkSecuritySecurityProfileGroupId"`
	// The name of the parent this security profile group belongs to. Format: organizations/{organization_id}.
	Parent *string `pulumi:"parent"`
	// Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.
	ThreatPreventionProfile *string                                      `pulumi:"threatPreventionProfile"`
	Timeouts                *NetworkSecuritySecurityProfileGroupTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkSecuritySecurityProfileGroup resource.
type NetworkSecuritySecurityProfileGroupArgs struct {
	// Reference to a SecurityProfile with the CustomIntercept configuration.
	CustomInterceptProfile pulumi.StringPtrInput
	// Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.
	CustomMirroringProfile pulumi.StringPtrInput
	// An optional description of the profile. The Max length is 512 characters.
	Description pulumi.StringPtrInput
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location of the security profile group. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The name of the security profile group resource.
	Name                                  pulumi.StringPtrInput
	NetworkSecuritySecurityProfileGroupId pulumi.StringPtrInput
	// The name of the parent this security profile group belongs to. Format: organizations/{organization_id}.
	Parent pulumi.StringPtrInput
	// Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.
	ThreatPreventionProfile pulumi.StringPtrInput
	Timeouts                NetworkSecuritySecurityProfileGroupTimeoutsPtrInput
}

func (NetworkSecuritySecurityProfileGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecuritySecurityProfileGroupArgs)(nil)).Elem()
}

type NetworkSecuritySecurityProfileGroupInput interface {
	pulumi.Input

	ToNetworkSecuritySecurityProfileGroupOutput() NetworkSecuritySecurityProfileGroupOutput
	ToNetworkSecuritySecurityProfileGroupOutputWithContext(ctx context.Context) NetworkSecuritySecurityProfileGroupOutput
}

func (*NetworkSecuritySecurityProfileGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecuritySecurityProfileGroup)(nil)).Elem()
}

func (i *NetworkSecuritySecurityProfileGroup) ToNetworkSecuritySecurityProfileGroupOutput() NetworkSecuritySecurityProfileGroupOutput {
	return i.ToNetworkSecuritySecurityProfileGroupOutputWithContext(context.Background())
}

func (i *NetworkSecuritySecurityProfileGroup) ToNetworkSecuritySecurityProfileGroupOutputWithContext(ctx context.Context) NetworkSecuritySecurityProfileGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecuritySecurityProfileGroupOutput)
}

type NetworkSecuritySecurityProfileGroupOutput struct{ *pulumi.OutputState }

func (NetworkSecuritySecurityProfileGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecuritySecurityProfileGroup)(nil)).Elem()
}

func (o NetworkSecuritySecurityProfileGroupOutput) ToNetworkSecuritySecurityProfileGroupOutput() NetworkSecuritySecurityProfileGroupOutput {
	return o
}

func (o NetworkSecuritySecurityProfileGroupOutput) ToNetworkSecuritySecurityProfileGroupOutputWithContext(ctx context.Context) NetworkSecuritySecurityProfileGroupOutput {
	return o
}

// Time the security profile group was created in UTC.
func (o NetworkSecuritySecurityProfileGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Reference to a SecurityProfile with the CustomIntercept configuration.
func (o NetworkSecuritySecurityProfileGroupOutput) CustomInterceptProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringPtrOutput { return v.CustomInterceptProfile }).(pulumi.StringPtrOutput)
}

// Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.
func (o NetworkSecuritySecurityProfileGroupOutput) CustomMirroringProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringPtrOutput { return v.CustomMirroringProfile }).(pulumi.StringPtrOutput)
}

// An optional description of the profile. The Max length is 512 characters.
func (o NetworkSecuritySecurityProfileGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkSecuritySecurityProfileGroupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
// requests to ensure the client has an up-to-date value before proceeding.
func (o NetworkSecuritySecurityProfileGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkSecuritySecurityProfileGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location of the security profile group. The default value is 'global'.
func (o NetworkSecuritySecurityProfileGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the security profile group resource.
func (o NetworkSecuritySecurityProfileGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecuritySecurityProfileGroupOutput) NetworkSecuritySecurityProfileGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringOutput {
		return v.NetworkSecuritySecurityProfileGroupId
	}).(pulumi.StringOutput)
}

// The name of the parent this security profile group belongs to. Format: organizations/{organization_id}.
func (o NetworkSecuritySecurityProfileGroupOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecuritySecurityProfileGroupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

// Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.
func (o NetworkSecuritySecurityProfileGroupOutput) ThreatPreventionProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringPtrOutput { return v.ThreatPreventionProfile }).(pulumi.StringPtrOutput)
}

func (o NetworkSecuritySecurityProfileGroupOutput) Timeouts() NetworkSecuritySecurityProfileGroupTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) NetworkSecuritySecurityProfileGroupTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecuritySecurityProfileGroupTimeoutsPtrOutput)
}

// Time the security profile group was updated in UTC.
func (o NetworkSecuritySecurityProfileGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecuritySecurityProfileGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecuritySecurityProfileGroupInput)(nil)).Elem(), &NetworkSecuritySecurityProfileGroup{})
	pulumi.RegisterOutputType(NetworkSecuritySecurityProfileGroupOutput{})
}
