// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeNetwork struct {
	pulumi.CustomResourceState

	// When set to 'true', the network is created in "auto subnet mode" and it will create a subnet for each region
	// automatically across the '10.128.0.0/9' address range. When set to 'false', the network is created in "custom subnet
	// mode" so the user can explicitly connect subnetwork resources.
	AutoCreateSubnetworks pulumi.BoolPtrOutput `pulumi:"autoCreateSubnetworks"`
	// Enables/disables the comparison of MED across routes with different Neighbor ASNs. This value can only be set if the
	// --bgp-best-path-selection-mode is STANDARD
	BgpAlwaysCompareMed pulumi.BoolOutput `pulumi:"bgpAlwaysCompareMed"`
	// The BGP best selection algorithm to be employed. MODE can be LEGACY or STANDARD. Possible values: ["LEGACY", "STANDARD"]
	BgpBestPathSelectionMode pulumi.StringOutput `pulumi:"bgpBestPathSelectionMode"`
	// Choice of the behavior of inter-regional cost and MED in the BPS algorithm. Possible values: ["DEFAULT",
	// "ADD_COST_TO_MED"]
	BgpInterRegionCost pulumi.StringOutput `pulumi:"bgpInterRegionCost"`
	ComputeNetworkId   pulumi.StringOutput `pulumi:"computeNetworkId"`
	// If set to 'true', default routes ('0.0.0.0/0') will be deleted immediately after network creation. Defaults to 'false'.
	DeleteDefaultRoutesOnCreate pulumi.BoolPtrOutput `pulumi:"deleteDefaultRoutesOnCreate"`
	// An optional description of this resource. The resource must be recreated to modify this field.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix
	// fd20::/20.
	EnableUlaInternalIpv6 pulumi.BoolPtrOutput `pulumi:"enableUlaInternalIpv6"`
	// The gateway address for default routing out of the network. This value is selected by GCP.
	GatewayIpv4 pulumi.StringOutput `pulumi:"gatewayIpv4"`
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA
	// prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail
	// if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be
	// randomly allocated from fd20::/20 and returned via this field.
	InternalIpv6Range pulumi.StringOutput `pulumi:"internalIpv6Range"`
	// Maximum Transmission Unit in bytes. The default value is 1460 bytes. The minimum value for this field is 1300 and the
	// maximum value is 8896 bytes (jumbo frames). Note that packets larger than 1500 bytes (standard Ethernet) can be subject
	// to TCP-MSS clamping or dropped with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or
	// other VPCs with varying MTUs.
	Mtu pulumi.Float64Output `pulumi:"mtu"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible
	// values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]
	NetworkFirewallPolicyEnforcementOrder pulumi.StringPtrOutput `pulumi:"networkFirewallPolicyEnforcementOrder"`
	// The unique identifier for the resource. This identifier is defined by the server.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// A full or partial URL of the network profile to apply to this network. This field can be set only at resource creation
	// time. For example, the following are valid URLs: *
	// https://www.googleapis.com/compute/v1/projects/{projectId}/global/networkProfiles/{network_profile_name} *
	// projects/{projectId}/global/networkProfiles/{network_profile_name}
	NetworkProfile pulumi.StringPtrOutput `pulumi:"networkProfile"`
	// The unique identifier for the resource. This identifier is defined by the server.
	//
	// Deprecated: Deprecated
	NumericId pulumi.StringOutput `pulumi:"numericId"`
	Project   pulumi.StringOutput `pulumi:"project"`
	// The network-wide routing mode to use. If set to 'REGIONAL', this network's cloud routers will only advertise routes with
	// subnetworks of this network in the same region as the router. If set to 'GLOBAL', this network's cloud routers will
	// advertise routes with all subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
	RoutingMode pulumi.StringOutput             `pulumi:"routingMode"`
	SelfLink    pulumi.StringOutput             `pulumi:"selfLink"`
	Timeouts    ComputeNetworkTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeNetwork registers a new resource with the given unique name, arguments, and options.
func NewComputeNetwork(ctx *pulumi.Context,
	name string, args *ComputeNetworkArgs, opts ...pulumi.ResourceOption) (*ComputeNetwork, error) {
	if args == nil {
		args = &ComputeNetworkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeNetwork
	err = ctx.RegisterPackageResource("google-beta:index/computeNetwork:ComputeNetwork", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeNetwork gets an existing ComputeNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeNetworkState, opts ...pulumi.ResourceOption) (*ComputeNetwork, error) {
	var resource ComputeNetwork
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeNetwork:ComputeNetwork", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeNetwork resources.
type computeNetworkState struct {
	// When set to 'true', the network is created in "auto subnet mode" and it will create a subnet for each region
	// automatically across the '10.128.0.0/9' address range. When set to 'false', the network is created in "custom subnet
	// mode" so the user can explicitly connect subnetwork resources.
	AutoCreateSubnetworks *bool `pulumi:"autoCreateSubnetworks"`
	// Enables/disables the comparison of MED across routes with different Neighbor ASNs. This value can only be set if the
	// --bgp-best-path-selection-mode is STANDARD
	BgpAlwaysCompareMed *bool `pulumi:"bgpAlwaysCompareMed"`
	// The BGP best selection algorithm to be employed. MODE can be LEGACY or STANDARD. Possible values: ["LEGACY", "STANDARD"]
	BgpBestPathSelectionMode *string `pulumi:"bgpBestPathSelectionMode"`
	// Choice of the behavior of inter-regional cost and MED in the BPS algorithm. Possible values: ["DEFAULT",
	// "ADD_COST_TO_MED"]
	BgpInterRegionCost *string `pulumi:"bgpInterRegionCost"`
	ComputeNetworkId   *string `pulumi:"computeNetworkId"`
	// If set to 'true', default routes ('0.0.0.0/0') will be deleted immediately after network creation. Defaults to 'false'.
	DeleteDefaultRoutesOnCreate *bool `pulumi:"deleteDefaultRoutesOnCreate"`
	// An optional description of this resource. The resource must be recreated to modify this field.
	Description *string `pulumi:"description"`
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix
	// fd20::/20.
	EnableUlaInternalIpv6 *bool `pulumi:"enableUlaInternalIpv6"`
	// The gateway address for default routing out of the network. This value is selected by GCP.
	GatewayIpv4 *string `pulumi:"gatewayIpv4"`
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA
	// prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail
	// if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be
	// randomly allocated from fd20::/20 and returned via this field.
	InternalIpv6Range *string `pulumi:"internalIpv6Range"`
	// Maximum Transmission Unit in bytes. The default value is 1460 bytes. The minimum value for this field is 1300 and the
	// maximum value is 8896 bytes (jumbo frames). Note that packets larger than 1500 bytes (standard Ethernet) can be subject
	// to TCP-MSS clamping or dropped with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or
	// other VPCs with varying MTUs.
	Mtu *float64 `pulumi:"mtu"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible
	// values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]
	NetworkFirewallPolicyEnforcementOrder *string `pulumi:"networkFirewallPolicyEnforcementOrder"`
	// The unique identifier for the resource. This identifier is defined by the server.
	NetworkId *string `pulumi:"networkId"`
	// A full or partial URL of the network profile to apply to this network. This field can be set only at resource creation
	// time. For example, the following are valid URLs: *
	// https://www.googleapis.com/compute/v1/projects/{projectId}/global/networkProfiles/{network_profile_name} *
	// projects/{projectId}/global/networkProfiles/{network_profile_name}
	NetworkProfile *string `pulumi:"networkProfile"`
	// The unique identifier for the resource. This identifier is defined by the server.
	//
	// Deprecated: Deprecated
	NumericId *string `pulumi:"numericId"`
	Project   *string `pulumi:"project"`
	// The network-wide routing mode to use. If set to 'REGIONAL', this network's cloud routers will only advertise routes with
	// subnetworks of this network in the same region as the router. If set to 'GLOBAL', this network's cloud routers will
	// advertise routes with all subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
	RoutingMode *string                 `pulumi:"routingMode"`
	SelfLink    *string                 `pulumi:"selfLink"`
	Timeouts    *ComputeNetworkTimeouts `pulumi:"timeouts"`
}

type ComputeNetworkState struct {
	// When set to 'true', the network is created in "auto subnet mode" and it will create a subnet for each region
	// automatically across the '10.128.0.0/9' address range. When set to 'false', the network is created in "custom subnet
	// mode" so the user can explicitly connect subnetwork resources.
	AutoCreateSubnetworks pulumi.BoolPtrInput
	// Enables/disables the comparison of MED across routes with different Neighbor ASNs. This value can only be set if the
	// --bgp-best-path-selection-mode is STANDARD
	BgpAlwaysCompareMed pulumi.BoolPtrInput
	// The BGP best selection algorithm to be employed. MODE can be LEGACY or STANDARD. Possible values: ["LEGACY", "STANDARD"]
	BgpBestPathSelectionMode pulumi.StringPtrInput
	// Choice of the behavior of inter-regional cost and MED in the BPS algorithm. Possible values: ["DEFAULT",
	// "ADD_COST_TO_MED"]
	BgpInterRegionCost pulumi.StringPtrInput
	ComputeNetworkId   pulumi.StringPtrInput
	// If set to 'true', default routes ('0.0.0.0/0') will be deleted immediately after network creation. Defaults to 'false'.
	DeleteDefaultRoutesOnCreate pulumi.BoolPtrInput
	// An optional description of this resource. The resource must be recreated to modify this field.
	Description pulumi.StringPtrInput
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix
	// fd20::/20.
	EnableUlaInternalIpv6 pulumi.BoolPtrInput
	// The gateway address for default routing out of the network. This value is selected by GCP.
	GatewayIpv4 pulumi.StringPtrInput
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA
	// prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail
	// if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be
	// randomly allocated from fd20::/20 and returned via this field.
	InternalIpv6Range pulumi.StringPtrInput
	// Maximum Transmission Unit in bytes. The default value is 1460 bytes. The minimum value for this field is 1300 and the
	// maximum value is 8896 bytes (jumbo frames). Note that packets larger than 1500 bytes (standard Ethernet) can be subject
	// to TCP-MSS clamping or dropped with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or
	// other VPCs with varying MTUs.
	Mtu pulumi.Float64PtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible
	// values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]
	NetworkFirewallPolicyEnforcementOrder pulumi.StringPtrInput
	// The unique identifier for the resource. This identifier is defined by the server.
	NetworkId pulumi.StringPtrInput
	// A full or partial URL of the network profile to apply to this network. This field can be set only at resource creation
	// time. For example, the following are valid URLs: *
	// https://www.googleapis.com/compute/v1/projects/{projectId}/global/networkProfiles/{network_profile_name} *
	// projects/{projectId}/global/networkProfiles/{network_profile_name}
	NetworkProfile pulumi.StringPtrInput
	// The unique identifier for the resource. This identifier is defined by the server.
	//
	// Deprecated: Deprecated
	NumericId pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	// The network-wide routing mode to use. If set to 'REGIONAL', this network's cloud routers will only advertise routes with
	// subnetworks of this network in the same region as the router. If set to 'GLOBAL', this network's cloud routers will
	// advertise routes with all subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
	RoutingMode pulumi.StringPtrInput
	SelfLink    pulumi.StringPtrInput
	Timeouts    ComputeNetworkTimeoutsPtrInput
}

func (ComputeNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkState)(nil)).Elem()
}

type computeNetworkArgs struct {
	// When set to 'true', the network is created in "auto subnet mode" and it will create a subnet for each region
	// automatically across the '10.128.0.0/9' address range. When set to 'false', the network is created in "custom subnet
	// mode" so the user can explicitly connect subnetwork resources.
	AutoCreateSubnetworks *bool `pulumi:"autoCreateSubnetworks"`
	// Enables/disables the comparison of MED across routes with different Neighbor ASNs. This value can only be set if the
	// --bgp-best-path-selection-mode is STANDARD
	BgpAlwaysCompareMed *bool `pulumi:"bgpAlwaysCompareMed"`
	// The BGP best selection algorithm to be employed. MODE can be LEGACY or STANDARD. Possible values: ["LEGACY", "STANDARD"]
	BgpBestPathSelectionMode *string `pulumi:"bgpBestPathSelectionMode"`
	// Choice of the behavior of inter-regional cost and MED in the BPS algorithm. Possible values: ["DEFAULT",
	// "ADD_COST_TO_MED"]
	BgpInterRegionCost *string `pulumi:"bgpInterRegionCost"`
	ComputeNetworkId   *string `pulumi:"computeNetworkId"`
	// If set to 'true', default routes ('0.0.0.0/0') will be deleted immediately after network creation. Defaults to 'false'.
	DeleteDefaultRoutesOnCreate *bool `pulumi:"deleteDefaultRoutesOnCreate"`
	// An optional description of this resource. The resource must be recreated to modify this field.
	Description *string `pulumi:"description"`
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix
	// fd20::/20.
	EnableUlaInternalIpv6 *bool `pulumi:"enableUlaInternalIpv6"`
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA
	// prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail
	// if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be
	// randomly allocated from fd20::/20 and returned via this field.
	InternalIpv6Range *string `pulumi:"internalIpv6Range"`
	// Maximum Transmission Unit in bytes. The default value is 1460 bytes. The minimum value for this field is 1300 and the
	// maximum value is 8896 bytes (jumbo frames). Note that packets larger than 1500 bytes (standard Ethernet) can be subject
	// to TCP-MSS clamping or dropped with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or
	// other VPCs with varying MTUs.
	Mtu *float64 `pulumi:"mtu"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible
	// values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]
	NetworkFirewallPolicyEnforcementOrder *string `pulumi:"networkFirewallPolicyEnforcementOrder"`
	// A full or partial URL of the network profile to apply to this network. This field can be set only at resource creation
	// time. For example, the following are valid URLs: *
	// https://www.googleapis.com/compute/v1/projects/{projectId}/global/networkProfiles/{network_profile_name} *
	// projects/{projectId}/global/networkProfiles/{network_profile_name}
	NetworkProfile *string `pulumi:"networkProfile"`
	Project        *string `pulumi:"project"`
	// The network-wide routing mode to use. If set to 'REGIONAL', this network's cloud routers will only advertise routes with
	// subnetworks of this network in the same region as the router. If set to 'GLOBAL', this network's cloud routers will
	// advertise routes with all subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
	RoutingMode *string                 `pulumi:"routingMode"`
	Timeouts    *ComputeNetworkTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeNetwork resource.
type ComputeNetworkArgs struct {
	// When set to 'true', the network is created in "auto subnet mode" and it will create a subnet for each region
	// automatically across the '10.128.0.0/9' address range. When set to 'false', the network is created in "custom subnet
	// mode" so the user can explicitly connect subnetwork resources.
	AutoCreateSubnetworks pulumi.BoolPtrInput
	// Enables/disables the comparison of MED across routes with different Neighbor ASNs. This value can only be set if the
	// --bgp-best-path-selection-mode is STANDARD
	BgpAlwaysCompareMed pulumi.BoolPtrInput
	// The BGP best selection algorithm to be employed. MODE can be LEGACY or STANDARD. Possible values: ["LEGACY", "STANDARD"]
	BgpBestPathSelectionMode pulumi.StringPtrInput
	// Choice of the behavior of inter-regional cost and MED in the BPS algorithm. Possible values: ["DEFAULT",
	// "ADD_COST_TO_MED"]
	BgpInterRegionCost pulumi.StringPtrInput
	ComputeNetworkId   pulumi.StringPtrInput
	// If set to 'true', default routes ('0.0.0.0/0') will be deleted immediately after network creation. Defaults to 'false'.
	DeleteDefaultRoutesOnCreate pulumi.BoolPtrInput
	// An optional description of this resource. The resource must be recreated to modify this field.
	Description pulumi.StringPtrInput
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix
	// fd20::/20.
	EnableUlaInternalIpv6 pulumi.BoolPtrInput
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA
	// prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail
	// if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be
	// randomly allocated from fd20::/20 and returned via this field.
	InternalIpv6Range pulumi.StringPtrInput
	// Maximum Transmission Unit in bytes. The default value is 1460 bytes. The minimum value for this field is 1300 and the
	// maximum value is 8896 bytes (jumbo frames). Note that packets larger than 1500 bytes (standard Ethernet) can be subject
	// to TCP-MSS clamping or dropped with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or
	// other VPCs with varying MTUs.
	Mtu pulumi.Float64PtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible
	// values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]
	NetworkFirewallPolicyEnforcementOrder pulumi.StringPtrInput
	// A full or partial URL of the network profile to apply to this network. This field can be set only at resource creation
	// time. For example, the following are valid URLs: *
	// https://www.googleapis.com/compute/v1/projects/{projectId}/global/networkProfiles/{network_profile_name} *
	// projects/{projectId}/global/networkProfiles/{network_profile_name}
	NetworkProfile pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
	// The network-wide routing mode to use. If set to 'REGIONAL', this network's cloud routers will only advertise routes with
	// subnetworks of this network in the same region as the router. If set to 'GLOBAL', this network's cloud routers will
	// advertise routes with all subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
	RoutingMode pulumi.StringPtrInput
	Timeouts    ComputeNetworkTimeoutsPtrInput
}

func (ComputeNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkArgs)(nil)).Elem()
}

type ComputeNetworkInput interface {
	pulumi.Input

	ToComputeNetworkOutput() ComputeNetworkOutput
	ToComputeNetworkOutputWithContext(ctx context.Context) ComputeNetworkOutput
}

func (*ComputeNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetwork)(nil)).Elem()
}

func (i *ComputeNetwork) ToComputeNetworkOutput() ComputeNetworkOutput {
	return i.ToComputeNetworkOutputWithContext(context.Background())
}

func (i *ComputeNetwork) ToComputeNetworkOutputWithContext(ctx context.Context) ComputeNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeNetworkOutput)
}

type ComputeNetworkOutput struct{ *pulumi.OutputState }

func (ComputeNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetwork)(nil)).Elem()
}

func (o ComputeNetworkOutput) ToComputeNetworkOutput() ComputeNetworkOutput {
	return o
}

func (o ComputeNetworkOutput) ToComputeNetworkOutputWithContext(ctx context.Context) ComputeNetworkOutput {
	return o
}

// When set to 'true', the network is created in "auto subnet mode" and it will create a subnet for each region
// automatically across the '10.128.0.0/9' address range. When set to 'false', the network is created in "custom subnet
// mode" so the user can explicitly connect subnetwork resources.
func (o ComputeNetworkOutput) AutoCreateSubnetworks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.BoolPtrOutput { return v.AutoCreateSubnetworks }).(pulumi.BoolPtrOutput)
}

// Enables/disables the comparison of MED across routes with different Neighbor ASNs. This value can only be set if the
// --bgp-best-path-selection-mode is STANDARD
func (o ComputeNetworkOutput) BgpAlwaysCompareMed() pulumi.BoolOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.BoolOutput { return v.BgpAlwaysCompareMed }).(pulumi.BoolOutput)
}

// The BGP best selection algorithm to be employed. MODE can be LEGACY or STANDARD. Possible values: ["LEGACY", "STANDARD"]
func (o ComputeNetworkOutput) BgpBestPathSelectionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.BgpBestPathSelectionMode }).(pulumi.StringOutput)
}

// Choice of the behavior of inter-regional cost and MED in the BPS algorithm. Possible values: ["DEFAULT",
// "ADD_COST_TO_MED"]
func (o ComputeNetworkOutput) BgpInterRegionCost() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.BgpInterRegionCost }).(pulumi.StringOutput)
}

func (o ComputeNetworkOutput) ComputeNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.ComputeNetworkId }).(pulumi.StringOutput)
}

// If set to 'true', default routes ('0.0.0.0/0') will be deleted immediately after network creation. Defaults to 'false'.
func (o ComputeNetworkOutput) DeleteDefaultRoutesOnCreate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.BoolPtrOutput { return v.DeleteDefaultRoutesOnCreate }).(pulumi.BoolPtrOutput)
}

// An optional description of this resource. The resource must be recreated to modify this field.
func (o ComputeNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix
// fd20::/20.
func (o ComputeNetworkOutput) EnableUlaInternalIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.BoolPtrOutput { return v.EnableUlaInternalIpv6 }).(pulumi.BoolPtrOutput)
}

// The gateway address for default routing out of the network. This value is selected by GCP.
func (o ComputeNetworkOutput) GatewayIpv4() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.GatewayIpv4 }).(pulumi.StringOutput)
}

// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA
// prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail
// if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be
// randomly allocated from fd20::/20 and returned via this field.
func (o ComputeNetworkOutput) InternalIpv6Range() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.InternalIpv6Range }).(pulumi.StringOutput)
}

// Maximum Transmission Unit in bytes. The default value is 1460 bytes. The minimum value for this field is 1300 and the
// maximum value is 8896 bytes (jumbo frames). Note that packets larger than 1500 bytes (standard Ethernet) can be subject
// to TCP-MSS clamping or dropped with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or
// other VPCs with varying MTUs.
func (o ComputeNetworkOutput) Mtu() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.Float64Output { return v.Mtu }).(pulumi.Float64Output)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible
// values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]
func (o ComputeNetworkOutput) NetworkFirewallPolicyEnforcementOrder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringPtrOutput { return v.NetworkFirewallPolicyEnforcementOrder }).(pulumi.StringPtrOutput)
}

// The unique identifier for the resource. This identifier is defined by the server.
func (o ComputeNetworkOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// A full or partial URL of the network profile to apply to this network. This field can be set only at resource creation
// time. For example, the following are valid URLs: *
// https://www.googleapis.com/compute/v1/projects/{projectId}/global/networkProfiles/{network_profile_name} *
// projects/{projectId}/global/networkProfiles/{network_profile_name}
func (o ComputeNetworkOutput) NetworkProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringPtrOutput { return v.NetworkProfile }).(pulumi.StringPtrOutput)
}

// The unique identifier for the resource. This identifier is defined by the server.
//
// Deprecated: Deprecated
func (o ComputeNetworkOutput) NumericId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.NumericId }).(pulumi.StringOutput)
}

func (o ComputeNetworkOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The network-wide routing mode to use. If set to 'REGIONAL', this network's cloud routers will only advertise routes with
// subnetworks of this network in the same region as the router. If set to 'GLOBAL', this network's cloud routers will
// advertise routes with all subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
func (o ComputeNetworkOutput) RoutingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.RoutingMode }).(pulumi.StringOutput)
}

func (o ComputeNetworkOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetwork) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeNetworkOutput) Timeouts() ComputeNetworkTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeNetwork) ComputeNetworkTimeoutsPtrOutput { return v.Timeouts }).(ComputeNetworkTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeNetworkInput)(nil)).Elem(), &ComputeNetwork{})
	pulumi.RegisterOutputType(ComputeNetworkOutput{})
}
