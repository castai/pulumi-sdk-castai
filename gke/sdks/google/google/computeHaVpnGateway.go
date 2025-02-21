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

type ComputeHaVpnGateway struct {
	pulumi.CustomResourceState

	ComputeHaVpnGatewayId pulumi.StringOutput `pulumi:"computeHaVpnGatewayId"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value:
	// "IPV4" Possible values: ["IPV4", "IPV6"]
	GatewayIpVersion pulumi.StringPtrOutput `pulumi:"gatewayIpVersion"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringOutput `pulumi:"network"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The region this gateway should sit in.
	Region   pulumi.StringOutput `pulumi:"region"`
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The stack type for this VPN gateway to identify the IP protocols that are enabled. If not specified, IPV4_ONLY will be
	// used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType pulumi.StringPtrOutput               `pulumi:"stackType"`
	Timeouts  ComputeHaVpnGatewayTimeoutsPtrOutput `pulumi:"timeouts"`
	// A list of interfaces on this VPN gateway.
	VpnInterfaces ComputeHaVpnGatewayVpnInterfaceArrayOutput `pulumi:"vpnInterfaces"`
}

// NewComputeHaVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewComputeHaVpnGateway(ctx *pulumi.Context,
	name string, args *ComputeHaVpnGatewayArgs, opts ...pulumi.ResourceOption) (*ComputeHaVpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeHaVpnGateway
	err = ctx.RegisterPackageResource("google:index/computeHaVpnGateway:ComputeHaVpnGateway", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeHaVpnGateway gets an existing ComputeHaVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeHaVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeHaVpnGatewayState, opts ...pulumi.ResourceOption) (*ComputeHaVpnGateway, error) {
	var resource ComputeHaVpnGateway
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeHaVpnGateway:ComputeHaVpnGateway", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeHaVpnGateway resources.
type computeHaVpnGatewayState struct {
	ComputeHaVpnGatewayId *string `pulumi:"computeHaVpnGatewayId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value:
	// "IPV4" Possible values: ["IPV4", "IPV6"]
	GatewayIpVersion *string `pulumi:"gatewayIpVersion"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region   *string `pulumi:"region"`
	SelfLink *string `pulumi:"selfLink"`
	// The stack type for this VPN gateway to identify the IP protocols that are enabled. If not specified, IPV4_ONLY will be
	// used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType *string                      `pulumi:"stackType"`
	Timeouts  *ComputeHaVpnGatewayTimeouts `pulumi:"timeouts"`
	// A list of interfaces on this VPN gateway.
	VpnInterfaces []ComputeHaVpnGatewayVpnInterface `pulumi:"vpnInterfaces"`
}

type ComputeHaVpnGatewayState struct {
	ComputeHaVpnGatewayId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value:
	// "IPV4" Possible values: ["IPV4", "IPV6"]
	GatewayIpVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region this gateway should sit in.
	Region   pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	// The stack type for this VPN gateway to identify the IP protocols that are enabled. If not specified, IPV4_ONLY will be
	// used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType pulumi.StringPtrInput
	Timeouts  ComputeHaVpnGatewayTimeoutsPtrInput
	// A list of interfaces on this VPN gateway.
	VpnInterfaces ComputeHaVpnGatewayVpnInterfaceArrayInput
}

func (ComputeHaVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeHaVpnGatewayState)(nil)).Elem()
}

type computeHaVpnGatewayArgs struct {
	ComputeHaVpnGatewayId *string `pulumi:"computeHaVpnGatewayId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value:
	// "IPV4" Possible values: ["IPV4", "IPV6"]
	GatewayIpVersion *string `pulumi:"gatewayIpVersion"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network string  `pulumi:"network"`
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region *string `pulumi:"region"`
	// The stack type for this VPN gateway to identify the IP protocols that are enabled. If not specified, IPV4_ONLY will be
	// used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType *string                      `pulumi:"stackType"`
	Timeouts  *ComputeHaVpnGatewayTimeouts `pulumi:"timeouts"`
	// A list of interfaces on this VPN gateway.
	VpnInterfaces []ComputeHaVpnGatewayVpnInterface `pulumi:"vpnInterfaces"`
}

// The set of arguments for constructing a ComputeHaVpnGateway resource.
type ComputeHaVpnGatewayArgs struct {
	ComputeHaVpnGatewayId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value:
	// "IPV4" Possible values: ["IPV4", "IPV6"]
	GatewayIpVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringInput
	Project pulumi.StringPtrInput
	// The region this gateway should sit in.
	Region pulumi.StringPtrInput
	// The stack type for this VPN gateway to identify the IP protocols that are enabled. If not specified, IPV4_ONLY will be
	// used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType pulumi.StringPtrInput
	Timeouts  ComputeHaVpnGatewayTimeoutsPtrInput
	// A list of interfaces on this VPN gateway.
	VpnInterfaces ComputeHaVpnGatewayVpnInterfaceArrayInput
}

func (ComputeHaVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeHaVpnGatewayArgs)(nil)).Elem()
}

type ComputeHaVpnGatewayInput interface {
	pulumi.Input

	ToComputeHaVpnGatewayOutput() ComputeHaVpnGatewayOutput
	ToComputeHaVpnGatewayOutputWithContext(ctx context.Context) ComputeHaVpnGatewayOutput
}

func (*ComputeHaVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeHaVpnGateway)(nil)).Elem()
}

func (i *ComputeHaVpnGateway) ToComputeHaVpnGatewayOutput() ComputeHaVpnGatewayOutput {
	return i.ToComputeHaVpnGatewayOutputWithContext(context.Background())
}

func (i *ComputeHaVpnGateway) ToComputeHaVpnGatewayOutputWithContext(ctx context.Context) ComputeHaVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeHaVpnGatewayOutput)
}

type ComputeHaVpnGatewayOutput struct{ *pulumi.OutputState }

func (ComputeHaVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeHaVpnGateway)(nil)).Elem()
}

func (o ComputeHaVpnGatewayOutput) ToComputeHaVpnGatewayOutput() ComputeHaVpnGatewayOutput {
	return o
}

func (o ComputeHaVpnGatewayOutput) ToComputeHaVpnGatewayOutputWithContext(ctx context.Context) ComputeHaVpnGatewayOutput {
	return o
}

func (o ComputeHaVpnGatewayOutput) ComputeHaVpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringOutput { return v.ComputeHaVpnGatewayId }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeHaVpnGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value:
// "IPV4" Possible values: ["IPV4", "IPV6"]
func (o ComputeHaVpnGatewayOutput) GatewayIpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringPtrOutput { return v.GatewayIpVersion }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeHaVpnGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network this VPN gateway is accepting traffic for.
func (o ComputeHaVpnGatewayOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o ComputeHaVpnGatewayOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region this gateway should sit in.
func (o ComputeHaVpnGatewayOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeHaVpnGatewayOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The stack type for this VPN gateway to identify the IP protocols that are enabled. If not specified, IPV4_ONLY will be
// used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
func (o ComputeHaVpnGatewayOutput) StackType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) pulumi.StringPtrOutput { return v.StackType }).(pulumi.StringPtrOutput)
}

func (o ComputeHaVpnGatewayOutput) Timeouts() ComputeHaVpnGatewayTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) ComputeHaVpnGatewayTimeoutsPtrOutput { return v.Timeouts }).(ComputeHaVpnGatewayTimeoutsPtrOutput)
}

// A list of interfaces on this VPN gateway.
func (o ComputeHaVpnGatewayOutput) VpnInterfaces() ComputeHaVpnGatewayVpnInterfaceArrayOutput {
	return o.ApplyT(func(v *ComputeHaVpnGateway) ComputeHaVpnGatewayVpnInterfaceArrayOutput { return v.VpnInterfaces }).(ComputeHaVpnGatewayVpnInterfaceArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeHaVpnGatewayInput)(nil)).Elem(), &ComputeHaVpnGateway{})
	pulumi.RegisterOutputType(ComputeHaVpnGatewayOutput{})
}
