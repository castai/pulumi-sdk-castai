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

type ComputeVpnTunnel struct {
	pulumi.CustomResourceState

	ComputeVpnTunnelId pulumi.StringOutput `pulumi:"computeVpnTunnelId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Detailed status message for the VPN tunnel.
	DetailedStatus  pulumi.StringOutput    `pulumi:"detailedStatus"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2.
	// Default version is 2.
	IkeVersion pulumi.Float64PtrOutput `pulumi:"ikeVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this VpnTunnel. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	LocalTrafficSelectors pulumi.StringArrayOutput `pulumi:"localTrafficSelectors"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway pulumi.StringPtrOutput `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface pulumi.Float64PtrOutput `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will
	// automatically use the same vpn_gateway_interface ID in the peer GCP VPN gateway. This field must reference a
	// 'google_compute_ha_vpn_gateway' resource.
	PeerGcpGateway pulumi.StringPtrOutput `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  pulumi.StringOutput `pulumi:"peerIp"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	Region pulumi.StringOutput `pulumi:"region"`
	// Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelectors pulumi.StringArrayOutput `pulumi:"remoteTrafficSelectors"`
	// URL of router resource to be used for dynamic routing.
	Router   pulumi.StringPtrOutput `pulumi:"router"`
	SelfLink pulumi.StringOutput    `pulumi:"selfLink"`
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret pulumi.StringOutput `pulumi:"sharedSecret"`
	// Hash of the shared secret.
	SharedSecretHash pulumi.StringOutput `pulumi:"sharedSecretHash"`
	// URL of the Target VPN gateway with which this VPN tunnel is associated.
	TargetVpnGateway pulumi.StringPtrOutput `pulumi:"targetVpnGateway"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput            `pulumi:"terraformLabels"`
	Timeouts        ComputeVpnTunnelTimeoutsPtrOutput `pulumi:"timeouts"`
	// The unique identifier for the resource. This identifier is defined by the server.
	TunnelId pulumi.StringOutput `pulumi:"tunnelId"`
	// URL of the VPN gateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway
	// resource is created. This field must reference a 'google_compute_ha_vpn_gateway' resource.
	VpnGateway pulumi.StringPtrOutput `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface pulumi.Float64PtrOutput `pulumi:"vpnGatewayInterface"`
}

// NewComputeVpnTunnel registers a new resource with the given unique name, arguments, and options.
func NewComputeVpnTunnel(ctx *pulumi.Context,
	name string, args *ComputeVpnTunnelArgs, opts ...pulumi.ResourceOption) (*ComputeVpnTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SharedSecret == nil {
		return nil, errors.New("invalid value for required argument 'SharedSecret'")
	}
	if args.SharedSecret != nil {
		args.SharedSecret = pulumi.ToSecret(args.SharedSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"sharedSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeVpnTunnel
	err = ctx.RegisterPackageResource("google-beta:index/computeVpnTunnel:ComputeVpnTunnel", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeVpnTunnel gets an existing ComputeVpnTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeVpnTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeVpnTunnelState, opts ...pulumi.ResourceOption) (*ComputeVpnTunnel, error) {
	var resource ComputeVpnTunnel
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeVpnTunnel:ComputeVpnTunnel", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeVpnTunnel resources.
type computeVpnTunnelState struct {
	ComputeVpnTunnelId *string `pulumi:"computeVpnTunnelId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Detailed status message for the VPN tunnel.
	DetailedStatus  *string           `pulumi:"detailedStatus"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2.
	// Default version is 2.
	IkeVersion *float64 `pulumi:"ikeVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this VpnTunnel. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	LocalTrafficSelectors []string `pulumi:"localTrafficSelectors"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name *string `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway *string `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface *float64 `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will
	// automatically use the same vpn_gateway_interface ID in the peer GCP VPN gateway. This field must reference a
	// 'google_compute_ha_vpn_gateway' resource.
	PeerGcpGateway *string `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  *string `pulumi:"peerIp"`
	Project *string `pulumi:"project"`
	// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	Region *string `pulumi:"region"`
	// Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelectors []string `pulumi:"remoteTrafficSelectors"`
	// URL of router resource to be used for dynamic routing.
	Router   *string `pulumi:"router"`
	SelfLink *string `pulumi:"selfLink"`
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret *string `pulumi:"sharedSecret"`
	// Hash of the shared secret.
	SharedSecretHash *string `pulumi:"sharedSecretHash"`
	// URL of the Target VPN gateway with which this VPN tunnel is associated.
	TargetVpnGateway *string `pulumi:"targetVpnGateway"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string         `pulumi:"terraformLabels"`
	Timeouts        *ComputeVpnTunnelTimeouts `pulumi:"timeouts"`
	// The unique identifier for the resource. This identifier is defined by the server.
	TunnelId *string `pulumi:"tunnelId"`
	// URL of the VPN gateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway
	// resource is created. This field must reference a 'google_compute_ha_vpn_gateway' resource.
	VpnGateway *string `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface *float64 `pulumi:"vpnGatewayInterface"`
}

type ComputeVpnTunnelState struct {
	ComputeVpnTunnelId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Detailed status message for the VPN tunnel.
	DetailedStatus  pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2.
	// Default version is 2.
	IkeVersion pulumi.Float64PtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this VpnTunnel. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	LocalTrafficSelectors pulumi.StringArrayInput
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringPtrInput
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway pulumi.StringPtrInput
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface pulumi.Float64PtrInput
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will
	// automatically use the same vpn_gateway_interface ID in the peer GCP VPN gateway. This field must reference a
	// 'google_compute_ha_vpn_gateway' resource.
	PeerGcpGateway pulumi.StringPtrInput
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	Region pulumi.StringPtrInput
	// Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelectors pulumi.StringArrayInput
	// URL of router resource to be used for dynamic routing.
	Router   pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret pulumi.StringPtrInput
	// Hash of the shared secret.
	SharedSecretHash pulumi.StringPtrInput
	// URL of the Target VPN gateway with which this VPN tunnel is associated.
	TargetVpnGateway pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ComputeVpnTunnelTimeoutsPtrInput
	// The unique identifier for the resource. This identifier is defined by the server.
	TunnelId pulumi.StringPtrInput
	// URL of the VPN gateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway
	// resource is created. This field must reference a 'google_compute_ha_vpn_gateway' resource.
	VpnGateway pulumi.StringPtrInput
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface pulumi.Float64PtrInput
}

func (ComputeVpnTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeVpnTunnelState)(nil)).Elem()
}

type computeVpnTunnelArgs struct {
	ComputeVpnTunnelId *string `pulumi:"computeVpnTunnelId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2.
	// Default version is 2.
	IkeVersion *float64 `pulumi:"ikeVersion"`
	// Labels to apply to this VpnTunnel. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	LocalTrafficSelectors []string `pulumi:"localTrafficSelectors"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name *string `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway *string `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface *float64 `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will
	// automatically use the same vpn_gateway_interface ID in the peer GCP VPN gateway. This field must reference a
	// 'google_compute_ha_vpn_gateway' resource.
	PeerGcpGateway *string `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  *string `pulumi:"peerIp"`
	Project *string `pulumi:"project"`
	// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	Region *string `pulumi:"region"`
	// Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelectors []string `pulumi:"remoteTrafficSelectors"`
	// URL of router resource to be used for dynamic routing.
	Router *string `pulumi:"router"`
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret string `pulumi:"sharedSecret"`
	// URL of the Target VPN gateway with which this VPN tunnel is associated.
	TargetVpnGateway *string                   `pulumi:"targetVpnGateway"`
	Timeouts         *ComputeVpnTunnelTimeouts `pulumi:"timeouts"`
	// URL of the VPN gateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway
	// resource is created. This field must reference a 'google_compute_ha_vpn_gateway' resource.
	VpnGateway *string `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface *float64 `pulumi:"vpnGatewayInterface"`
}

// The set of arguments for constructing a ComputeVpnTunnel resource.
type ComputeVpnTunnelArgs struct {
	ComputeVpnTunnelId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2.
	// Default version is 2.
	IkeVersion pulumi.Float64PtrInput
	// Labels to apply to this VpnTunnel. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	LocalTrafficSelectors pulumi.StringArrayInput
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringPtrInput
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway pulumi.StringPtrInput
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface pulumi.Float64PtrInput
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will
	// automatically use the same vpn_gateway_interface ID in the peer GCP VPN gateway. This field must reference a
	// 'google_compute_ha_vpn_gateway' resource.
	PeerGcpGateway pulumi.StringPtrInput
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	Region pulumi.StringPtrInput
	// Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
	// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelectors pulumi.StringArrayInput
	// URL of router resource to be used for dynamic routing.
	Router pulumi.StringPtrInput
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret pulumi.StringInput
	// URL of the Target VPN gateway with which this VPN tunnel is associated.
	TargetVpnGateway pulumi.StringPtrInput
	Timeouts         ComputeVpnTunnelTimeoutsPtrInput
	// URL of the VPN gateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway
	// resource is created. This field must reference a 'google_compute_ha_vpn_gateway' resource.
	VpnGateway pulumi.StringPtrInput
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface pulumi.Float64PtrInput
}

func (ComputeVpnTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeVpnTunnelArgs)(nil)).Elem()
}

type ComputeVpnTunnelInput interface {
	pulumi.Input

	ToComputeVpnTunnelOutput() ComputeVpnTunnelOutput
	ToComputeVpnTunnelOutputWithContext(ctx context.Context) ComputeVpnTunnelOutput
}

func (*ComputeVpnTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeVpnTunnel)(nil)).Elem()
}

func (i *ComputeVpnTunnel) ToComputeVpnTunnelOutput() ComputeVpnTunnelOutput {
	return i.ToComputeVpnTunnelOutputWithContext(context.Background())
}

func (i *ComputeVpnTunnel) ToComputeVpnTunnelOutputWithContext(ctx context.Context) ComputeVpnTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeVpnTunnelOutput)
}

type ComputeVpnTunnelOutput struct{ *pulumi.OutputState }

func (ComputeVpnTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeVpnTunnel)(nil)).Elem()
}

func (o ComputeVpnTunnelOutput) ToComputeVpnTunnelOutput() ComputeVpnTunnelOutput {
	return o
}

func (o ComputeVpnTunnelOutput) ToComputeVpnTunnelOutputWithContext(ctx context.Context) ComputeVpnTunnelOutput {
	return o
}

func (o ComputeVpnTunnelOutput) ComputeVpnTunnelId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.ComputeVpnTunnelId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeVpnTunnelOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeVpnTunnelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Detailed status message for the VPN tunnel.
func (o ComputeVpnTunnelOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

func (o ComputeVpnTunnelOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2.
// Default version is 2.
func (o ComputeVpnTunnelOutput) IkeVersion() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.Float64PtrOutput { return v.IkeVersion }).(pulumi.Float64PtrOutput)
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (o ComputeVpnTunnelOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this VpnTunnel. **Note**: This field is non-authoritative, and will only manage the labels present in
// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o ComputeVpnTunnelOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
func (o ComputeVpnTunnelOutput) LocalTrafficSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringArrayOutput { return v.LocalTrafficSelectors }).(pulumi.StringArrayOutput)
}

// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
// dash.
func (o ComputeVpnTunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
func (o ComputeVpnTunnelOutput) PeerExternalGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringPtrOutput { return v.PeerExternalGateway }).(pulumi.StringPtrOutput)
}

// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
func (o ComputeVpnTunnelOutput) PeerExternalGatewayInterface() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.Float64PtrOutput { return v.PeerExternalGatewayInterface }).(pulumi.Float64PtrOutput)
}

// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will
// automatically use the same vpn_gateway_interface ID in the peer GCP VPN gateway. This field must reference a
// 'google_compute_ha_vpn_gateway' resource.
func (o ComputeVpnTunnelOutput) PeerGcpGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringPtrOutput { return v.PeerGcpGateway }).(pulumi.StringPtrOutput)
}

// IP address of the peer VPN gateway. Only IPv4 is supported.
func (o ComputeVpnTunnelOutput) PeerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.PeerIp }).(pulumi.StringOutput)
}

func (o ComputeVpnTunnelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
func (o ComputeVpnTunnelOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR
// formatted string, for example '192.168.0.0/16'. The ranges should be disjoint. Only IPv4 is supported.
func (o ComputeVpnTunnelOutput) RemoteTrafficSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringArrayOutput { return v.RemoteTrafficSelectors }).(pulumi.StringArrayOutput)
}

// URL of router resource to be used for dynamic routing.
func (o ComputeVpnTunnelOutput) Router() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringPtrOutput { return v.Router }).(pulumi.StringPtrOutput)
}

func (o ComputeVpnTunnelOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
func (o ComputeVpnTunnelOutput) SharedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.SharedSecret }).(pulumi.StringOutput)
}

// Hash of the shared secret.
func (o ComputeVpnTunnelOutput) SharedSecretHash() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.SharedSecretHash }).(pulumi.StringOutput)
}

// URL of the Target VPN gateway with which this VPN tunnel is associated.
func (o ComputeVpnTunnelOutput) TargetVpnGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringPtrOutput { return v.TargetVpnGateway }).(pulumi.StringPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ComputeVpnTunnelOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ComputeVpnTunnelOutput) Timeouts() ComputeVpnTunnelTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) ComputeVpnTunnelTimeoutsPtrOutput { return v.Timeouts }).(ComputeVpnTunnelTimeoutsPtrOutput)
}

// The unique identifier for the resource. This identifier is defined by the server.
func (o ComputeVpnTunnelOutput) TunnelId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringOutput { return v.TunnelId }).(pulumi.StringOutput)
}

// URL of the VPN gateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway
// resource is created. This field must reference a 'google_compute_ha_vpn_gateway' resource.
func (o ComputeVpnTunnelOutput) VpnGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.StringPtrOutput { return v.VpnGateway }).(pulumi.StringPtrOutput)
}

// The interface ID of the VPN gateway with which this VPN tunnel is associated.
func (o ComputeVpnTunnelOutput) VpnGatewayInterface() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ComputeVpnTunnel) pulumi.Float64PtrOutput { return v.VpnGatewayInterface }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeVpnTunnelInput)(nil)).Elem(), &ComputeVpnTunnel{})
	pulumi.RegisterOutputType(ComputeVpnTunnelOutput{})
}
