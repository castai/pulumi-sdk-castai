// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeExternalVpnGateway struct {
	pulumi.CustomResourceState

	ComputeExternalVpnGatewayId pulumi.StringOutput `pulumi:"computeExternalVpnGatewayId"`
	// An optional description of this resource.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// A list of interfaces on this external VPN gateway.
	Interfaces ComputeExternalVpnGatewayInterfaceArrayOutput `pulumi:"interfaces"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for the external VPN gateway resource. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY",
	// "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"]
	RedundancyType pulumi.StringPtrOutput `pulumi:"redundancyType"`
	SelfLink       pulumi.StringOutput    `pulumi:"selfLink"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                     `pulumi:"terraformLabels"`
	Timeouts        ComputeExternalVpnGatewayTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeExternalVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewComputeExternalVpnGateway(ctx *pulumi.Context,
	name string, args *ComputeExternalVpnGatewayArgs, opts ...pulumi.ResourceOption) (*ComputeExternalVpnGateway, error) {
	if args == nil {
		args = &ComputeExternalVpnGatewayArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeExternalVpnGateway
	err = ctx.RegisterPackageResource("google:index/computeExternalVpnGateway:ComputeExternalVpnGateway", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeExternalVpnGateway gets an existing ComputeExternalVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeExternalVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeExternalVpnGatewayState, opts ...pulumi.ResourceOption) (*ComputeExternalVpnGateway, error) {
	var resource ComputeExternalVpnGateway
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeExternalVpnGateway:ComputeExternalVpnGateway", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeExternalVpnGateway resources.
type computeExternalVpnGatewayState struct {
	ComputeExternalVpnGatewayId *string `pulumi:"computeExternalVpnGatewayId"`
	// An optional description of this resource.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// A list of interfaces on this external VPN gateway.
	Interfaces []ComputeExternalVpnGatewayInterface `pulumi:"interfaces"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels for the external VPN gateway resource. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY",
	// "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"]
	RedundancyType *string `pulumi:"redundancyType"`
	SelfLink       *string `pulumi:"selfLink"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                  `pulumi:"terraformLabels"`
	Timeouts        *ComputeExternalVpnGatewayTimeouts `pulumi:"timeouts"`
}

type ComputeExternalVpnGatewayState struct {
	ComputeExternalVpnGatewayId pulumi.StringPtrInput
	// An optional description of this resource.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// A list of interfaces on this external VPN gateway.
	Interfaces ComputeExternalVpnGatewayInterfaceArrayInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels for the external VPN gateway resource. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY",
	// "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"]
	RedundancyType pulumi.StringPtrInput
	SelfLink       pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ComputeExternalVpnGatewayTimeoutsPtrInput
}

func (ComputeExternalVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeExternalVpnGatewayState)(nil)).Elem()
}

type computeExternalVpnGatewayArgs struct {
	ComputeExternalVpnGatewayId *string `pulumi:"computeExternalVpnGatewayId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	Interfaces []ComputeExternalVpnGatewayInterface `pulumi:"interfaces"`
	// Labels for the external VPN gateway resource. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY",
	// "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"]
	RedundancyType *string                            `pulumi:"redundancyType"`
	Timeouts       *ComputeExternalVpnGatewayTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeExternalVpnGateway resource.
type ComputeExternalVpnGatewayArgs struct {
	ComputeExternalVpnGatewayId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// A list of interfaces on this external VPN gateway.
	Interfaces ComputeExternalVpnGatewayInterfaceArrayInput
	// Labels for the external VPN gateway resource. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY",
	// "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"]
	RedundancyType pulumi.StringPtrInput
	Timeouts       ComputeExternalVpnGatewayTimeoutsPtrInput
}

func (ComputeExternalVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeExternalVpnGatewayArgs)(nil)).Elem()
}

type ComputeExternalVpnGatewayInput interface {
	pulumi.Input

	ToComputeExternalVpnGatewayOutput() ComputeExternalVpnGatewayOutput
	ToComputeExternalVpnGatewayOutputWithContext(ctx context.Context) ComputeExternalVpnGatewayOutput
}

func (*ComputeExternalVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeExternalVpnGateway)(nil)).Elem()
}

func (i *ComputeExternalVpnGateway) ToComputeExternalVpnGatewayOutput() ComputeExternalVpnGatewayOutput {
	return i.ToComputeExternalVpnGatewayOutputWithContext(context.Background())
}

func (i *ComputeExternalVpnGateway) ToComputeExternalVpnGatewayOutputWithContext(ctx context.Context) ComputeExternalVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeExternalVpnGatewayOutput)
}

type ComputeExternalVpnGatewayOutput struct{ *pulumi.OutputState }

func (ComputeExternalVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeExternalVpnGateway)(nil)).Elem()
}

func (o ComputeExternalVpnGatewayOutput) ToComputeExternalVpnGatewayOutput() ComputeExternalVpnGatewayOutput {
	return o
}

func (o ComputeExternalVpnGatewayOutput) ToComputeExternalVpnGatewayOutputWithContext(ctx context.Context) ComputeExternalVpnGatewayOutput {
	return o
}

func (o ComputeExternalVpnGatewayOutput) ComputeExternalVpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringOutput { return v.ComputeExternalVpnGatewayId }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeExternalVpnGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComputeExternalVpnGatewayOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// A list of interfaces on this external VPN gateway.
func (o ComputeExternalVpnGatewayOutput) Interfaces() ComputeExternalVpnGatewayInterfaceArrayOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) ComputeExternalVpnGatewayInterfaceArrayOutput { return v.Interfaces }).(ComputeExternalVpnGatewayInterfaceArrayOutput)
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (o ComputeExternalVpnGatewayOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for the external VPN gateway resource. **Note**: This field is non-authoritative, and will only manage the labels
// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o ComputeExternalVpnGatewayOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeExternalVpnGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeExternalVpnGatewayOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY",
// "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"]
func (o ComputeExternalVpnGatewayOutput) RedundancyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringPtrOutput { return v.RedundancyType }).(pulumi.StringPtrOutput)
}

func (o ComputeExternalVpnGatewayOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ComputeExternalVpnGatewayOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ComputeExternalVpnGatewayOutput) Timeouts() ComputeExternalVpnGatewayTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeExternalVpnGateway) ComputeExternalVpnGatewayTimeoutsPtrOutput { return v.Timeouts }).(ComputeExternalVpnGatewayTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeExternalVpnGatewayInput)(nil)).Elem(), &ComputeExternalVpnGateway{})
	pulumi.RegisterOutputType(ComputeExternalVpnGatewayOutput{})
}
