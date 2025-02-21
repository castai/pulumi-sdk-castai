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

type EdgenetworkSubnet struct {
	pulumi.CustomResourceState

	// The time when the subnet was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	EdgenetworkSubnetId pulumi.StringOutput    `pulumi:"edgenetworkSubnetId"`
	EffectiveLabels     pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	Ipv4Cidrs pulumi.StringArrayOutput `pulumi:"ipv4Cidrs"`
	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	Ipv6Cidrs pulumi.StringArrayOutput `pulumi:"ipv6Cidrs"`
	// Labels associated with this resource. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location pulumi.StringOutput `pulumi:"location"`
	// The canonical name of this resource, with format
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/subnets/{{subnet_id}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the network to which this router belongs. Must be of the form:
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'
	Network pulumi.StringOutput `pulumi:"network"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Current stage of the resource to the device by config push.
	State pulumi.StringOutput `pulumi:"state"`
	// A unique ID that identifies this subnet.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput             `pulumi:"terraformLabels"`
	Timeouts        EdgenetworkSubnetTimeoutsPtrOutput `pulumi:"timeouts"`
	// The time when the subnet was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	VlanId pulumi.Float64Output `pulumi:"vlanId"`
	// The name of the target Distributed Cloud Edge zone.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewEdgenetworkSubnet registers a new resource with the given unique name, arguments, and options.
func NewEdgenetworkSubnet(ctx *pulumi.Context,
	name string, args *EdgenetworkSubnetArgs, opts ...pulumi.ResourceOption) (*EdgenetworkSubnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource EdgenetworkSubnet
	err = ctx.RegisterPackageResource("google-beta:index/edgenetworkSubnet:EdgenetworkSubnet", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgenetworkSubnet gets an existing EdgenetworkSubnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgenetworkSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgenetworkSubnetState, opts ...pulumi.ResourceOption) (*EdgenetworkSubnet, error) {
	var resource EdgenetworkSubnet
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/edgenetworkSubnet:EdgenetworkSubnet", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgenetworkSubnet resources.
type edgenetworkSubnetState struct {
	// The time when the subnet was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description         *string           `pulumi:"description"`
	EdgenetworkSubnetId *string           `pulumi:"edgenetworkSubnetId"`
	EffectiveLabels     map[string]string `pulumi:"effectiveLabels"`
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	Ipv4Cidrs []string `pulumi:"ipv4Cidrs"`
	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	Ipv6Cidrs []string `pulumi:"ipv6Cidrs"`
	// Labels associated with this resource. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location *string `pulumi:"location"`
	// The canonical name of this resource, with format
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/subnets/{{subnet_id}}'
	Name *string `pulumi:"name"`
	// The ID of the network to which this router belongs. Must be of the form:
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	// Current stage of the resource to the device by config push.
	State *string `pulumi:"state"`
	// A unique ID that identifies this subnet.
	SubnetId *string `pulumi:"subnetId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string          `pulumi:"terraformLabels"`
	Timeouts        *EdgenetworkSubnetTimeouts `pulumi:"timeouts"`
	// The time when the subnet was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
	UpdateTime *string `pulumi:"updateTime"`
	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	VlanId *float64 `pulumi:"vlanId"`
	// The name of the target Distributed Cloud Edge zone.
	Zone *string `pulumi:"zone"`
}

type EdgenetworkSubnetState struct {
	// The time when the subnet was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description         pulumi.StringPtrInput
	EdgenetworkSubnetId pulumi.StringPtrInput
	EffectiveLabels     pulumi.StringMapInput
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	Ipv4Cidrs pulumi.StringArrayInput
	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	Ipv6Cidrs pulumi.StringArrayInput
	// Labels associated with this resource. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location pulumi.StringPtrInput
	// The canonical name of this resource, with format
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/subnets/{{subnet_id}}'
	Name pulumi.StringPtrInput
	// The ID of the network to which this router belongs. Must be of the form:
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'
	Network pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Current stage of the resource to the device by config push.
	State pulumi.StringPtrInput
	// A unique ID that identifies this subnet.
	SubnetId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        EdgenetworkSubnetTimeoutsPtrInput
	// The time when the subnet was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
	UpdateTime pulumi.StringPtrInput
	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	VlanId pulumi.Float64PtrInput
	// The name of the target Distributed Cloud Edge zone.
	Zone pulumi.StringPtrInput
}

func (EdgenetworkSubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgenetworkSubnetState)(nil)).Elem()
}

type edgenetworkSubnetArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description         *string `pulumi:"description"`
	EdgenetworkSubnetId *string `pulumi:"edgenetworkSubnetId"`
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	Ipv4Cidrs []string `pulumi:"ipv4Cidrs"`
	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	Ipv6Cidrs []string `pulumi:"ipv6Cidrs"`
	// Labels associated with this resource. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location string `pulumi:"location"`
	// The ID of the network to which this router belongs. Must be of the form:
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'
	Network string  `pulumi:"network"`
	Project *string `pulumi:"project"`
	// A unique ID that identifies this subnet.
	SubnetId string                     `pulumi:"subnetId"`
	Timeouts *EdgenetworkSubnetTimeouts `pulumi:"timeouts"`
	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	VlanId *float64 `pulumi:"vlanId"`
	// The name of the target Distributed Cloud Edge zone.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a EdgenetworkSubnet resource.
type EdgenetworkSubnetArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description         pulumi.StringPtrInput
	EdgenetworkSubnetId pulumi.StringPtrInput
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	Ipv4Cidrs pulumi.StringArrayInput
	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	Ipv6Cidrs pulumi.StringArrayInput
	// Labels associated with this resource. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location pulumi.StringInput
	// The ID of the network to which this router belongs. Must be of the form:
	// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'
	Network pulumi.StringInput
	Project pulumi.StringPtrInput
	// A unique ID that identifies this subnet.
	SubnetId pulumi.StringInput
	Timeouts EdgenetworkSubnetTimeoutsPtrInput
	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	VlanId pulumi.Float64PtrInput
	// The name of the target Distributed Cloud Edge zone.
	Zone pulumi.StringInput
}

func (EdgenetworkSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgenetworkSubnetArgs)(nil)).Elem()
}

type EdgenetworkSubnetInput interface {
	pulumi.Input

	ToEdgenetworkSubnetOutput() EdgenetworkSubnetOutput
	ToEdgenetworkSubnetOutputWithContext(ctx context.Context) EdgenetworkSubnetOutput
}

func (*EdgenetworkSubnet) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgenetworkSubnet)(nil)).Elem()
}

func (i *EdgenetworkSubnet) ToEdgenetworkSubnetOutput() EdgenetworkSubnetOutput {
	return i.ToEdgenetworkSubnetOutputWithContext(context.Background())
}

func (i *EdgenetworkSubnet) ToEdgenetworkSubnetOutputWithContext(ctx context.Context) EdgenetworkSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgenetworkSubnetOutput)
}

type EdgenetworkSubnetOutput struct{ *pulumi.OutputState }

func (EdgenetworkSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgenetworkSubnet)(nil)).Elem()
}

func (o EdgenetworkSubnetOutput) ToEdgenetworkSubnetOutput() EdgenetworkSubnetOutput {
	return o
}

func (o EdgenetworkSubnetOutput) ToEdgenetworkSubnetOutputWithContext(ctx context.Context) EdgenetworkSubnetOutput {
	return o
}

// The time when the subnet was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
// nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
func (o EdgenetworkSubnetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o EdgenetworkSubnetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EdgenetworkSubnetOutput) EdgenetworkSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.EdgenetworkSubnetId }).(pulumi.StringOutput)
}

func (o EdgenetworkSubnetOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
func (o EdgenetworkSubnetOutput) Ipv4Cidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringArrayOutput { return v.Ipv4Cidrs }).(pulumi.StringArrayOutput)
}

// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
func (o EdgenetworkSubnetOutput) Ipv6Cidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringArrayOutput { return v.Ipv6Cidrs }).(pulumi.StringArrayOutput)
}

// Labels associated with this resource. **Note**: This field is non-authoritative, and will only manage the labels present
// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o EdgenetworkSubnetOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
func (o EdgenetworkSubnetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The canonical name of this resource, with format
// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/subnets/{{subnet_id}}'
func (o EdgenetworkSubnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the network to which this router belongs. Must be of the form:
// 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'
func (o EdgenetworkSubnetOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o EdgenetworkSubnetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Current stage of the resource to the device by config push.
func (o EdgenetworkSubnetOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A unique ID that identifies this subnet.
func (o EdgenetworkSubnetOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o EdgenetworkSubnetOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o EdgenetworkSubnetOutput) Timeouts() EdgenetworkSubnetTimeoutsPtrOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) EdgenetworkSubnetTimeoutsPtrOutput { return v.Timeouts }).(EdgenetworkSubnetTimeoutsPtrOutput)
}

// The time when the subnet was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
// to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
func (o EdgenetworkSubnetOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
func (o EdgenetworkSubnetOutput) VlanId() pulumi.Float64Output {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.Float64Output { return v.VlanId }).(pulumi.Float64Output)
}

// The name of the target Distributed Cloud Edge zone.
func (o EdgenetworkSubnetOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgenetworkSubnet) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgenetworkSubnetInput)(nil)).Elem(), &EdgenetworkSubnet{})
	pulumi.RegisterOutputType(EdgenetworkSubnetOutput{})
}
