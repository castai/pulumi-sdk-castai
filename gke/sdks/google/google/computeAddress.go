// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeAddress struct {
	pulumi.CustomResourceState

	// The static external IP address represented by this resource. The IP address must be inside the specified subnetwork, if
	// any. Set by the API if undefined.
	Address pulumi.StringOutput `pulumi:"address"`
	// The type of address to reserve. Note: if you set this argument's value as 'INTERNAL' you need to leave the
	// 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL",
	// "EXTERNAL"]
	AddressType      pulumi.StringPtrOutput `pulumi:"addressType"`
	ComputeAddressId pulumi.StringOutput    `pulumi:"computeAddressId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this
	// address can be used after the external IPv6 address reservation. Possible values: ["VM", "NETLB"]
	Ipv6EndpointType pulumi.StringPtrOutput `pulumi:"ipv6EndpointType"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this address. A list of key->value pairs. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the
	// VPC_PEERING and IPSEC_INTERCONNECT purposes.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// The networking tier used for configuring this address. If this field is not specified, it is assumed to be PREMIUM. This
	// argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal
	// traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM",
	// "STANDARD"]
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.Float64Output `pulumi:"prefixLength"`
	Project      pulumi.StringOutput  `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values. * GCE_ENDPOINT for addresses that are used by VM
	// instances, alias IP ranges, load balancers, and similar resources. * SHARED_LOADBALANCER_VIP for an address that can be
	// used by multiple internal load balancers. * VPC_PEERING for addresses that are reserved for VPC peer networks. *
	// IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an HA VPN
	// over Cloud Interconnect configuration. These addresses are regional resources. * PRIVATE_SERVICE_CONNECT for a private
	// network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	// This should only be set when using an Internal address.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringOutput `pulumi:"region"`
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the
	// subnetwork's IP range. This field can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput          `pulumi:"terraformLabels"`
	Timeouts        ComputeAddressTimeoutsPtrOutput `pulumi:"timeouts"`
	// The URLs of the resources that are using this address.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewComputeAddress registers a new resource with the given unique name, arguments, and options.
func NewComputeAddress(ctx *pulumi.Context,
	name string, args *ComputeAddressArgs, opts ...pulumi.ResourceOption) (*ComputeAddress, error) {
	if args == nil {
		args = &ComputeAddressArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeAddress
	err = ctx.RegisterPackageResource("google:index/computeAddress:ComputeAddress", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeAddress gets an existing ComputeAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeAddressState, opts ...pulumi.ResourceOption) (*ComputeAddress, error) {
	var resource ComputeAddress
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeAddress:ComputeAddress", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeAddress resources.
type computeAddressState struct {
	// The static external IP address represented by this resource. The IP address must be inside the specified subnetwork, if
	// any. Set by the API if undefined.
	Address *string `pulumi:"address"`
	// The type of address to reserve. Note: if you set this argument's value as 'INTERNAL' you need to leave the
	// 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL",
	// "EXTERNAL"]
	AddressType      *string `pulumi:"addressType"`
	ComputeAddressId *string `pulumi:"computeAddressId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]
	IpVersion *string `pulumi:"ipVersion"`
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this
	// address can be used after the external IPv6 address reservation. Possible values: ["VM", "NETLB"]
	Ipv6EndpointType *string `pulumi:"ipv6EndpointType"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this address. A list of key->value pairs. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name *string `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the
	// VPC_PEERING and IPSEC_INTERCONNECT purposes.
	Network *string `pulumi:"network"`
	// The networking tier used for configuring this address. If this field is not specified, it is assumed to be PREMIUM. This
	// argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal
	// traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM",
	// "STANDARD"]
	NetworkTier *string `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength *float64 `pulumi:"prefixLength"`
	Project      *string  `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values. * GCE_ENDPOINT for addresses that are used by VM
	// instances, alias IP ranges, load balancers, and similar resources. * SHARED_LOADBALANCER_VIP for an address that can be
	// used by multiple internal load balancers. * VPC_PEERING for addresses that are reserved for VPC peer networks. *
	// IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an HA VPN
	// over Cloud Interconnect configuration. These addresses are regional resources. * PRIVATE_SERVICE_CONNECT for a private
	// network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	// This should only be set when using an Internal address.
	Purpose *string `pulumi:"purpose"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   *string `pulumi:"region"`
	SelfLink *string `pulumi:"selfLink"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the
	// subnetwork's IP range. This field can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork *string `pulumi:"subnetwork"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string       `pulumi:"terraformLabels"`
	Timeouts        *ComputeAddressTimeouts `pulumi:"timeouts"`
	// The URLs of the resources that are using this address.
	Users []string `pulumi:"users"`
}

type ComputeAddressState struct {
	// The static external IP address represented by this resource. The IP address must be inside the specified subnetwork, if
	// any. Set by the API if undefined.
	Address pulumi.StringPtrInput
	// The type of address to reserve. Note: if you set this argument's value as 'INTERNAL' you need to leave the
	// 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL",
	// "EXTERNAL"]
	AddressType      pulumi.StringPtrInput
	ComputeAddressId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]
	IpVersion pulumi.StringPtrInput
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this
	// address can be used after the external IPv6 address reservation. Possible values: ["VM", "NETLB"]
	Ipv6EndpointType pulumi.StringPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this address. A list of key->value pairs. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringPtrInput
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the
	// VPC_PEERING and IPSEC_INTERCONNECT purposes.
	Network pulumi.StringPtrInput
	// The networking tier used for configuring this address. If this field is not specified, it is assumed to be PREMIUM. This
	// argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal
	// traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM",
	// "STANDARD"]
	NetworkTier pulumi.StringPtrInput
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.Float64PtrInput
	Project      pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values. * GCE_ENDPOINT for addresses that are used by VM
	// instances, alias IP ranges, load balancers, and similar resources. * SHARED_LOADBALANCER_VIP for an address that can be
	// used by multiple internal load balancers. * VPC_PEERING for addresses that are reserved for VPC peer networks. *
	// IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an HA VPN
	// over Cloud Interconnect configuration. These addresses are regional resources. * PRIVATE_SERVICE_CONNECT for a private
	// network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	// This should only be set when using an Internal address.
	Purpose pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the
	// subnetwork's IP range. This field can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ComputeAddressTimeoutsPtrInput
	// The URLs of the resources that are using this address.
	Users pulumi.StringArrayInput
}

func (ComputeAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeAddressState)(nil)).Elem()
}

type computeAddressArgs struct {
	// The static external IP address represented by this resource. The IP address must be inside the specified subnetwork, if
	// any. Set by the API if undefined.
	Address *string `pulumi:"address"`
	// The type of address to reserve. Note: if you set this argument's value as 'INTERNAL' you need to leave the
	// 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL",
	// "EXTERNAL"]
	AddressType      *string `pulumi:"addressType"`
	ComputeAddressId *string `pulumi:"computeAddressId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]
	IpVersion *string `pulumi:"ipVersion"`
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this
	// address can be used after the external IPv6 address reservation. Possible values: ["VM", "NETLB"]
	Ipv6EndpointType *string `pulumi:"ipv6EndpointType"`
	// Labels to apply to this address. A list of key->value pairs. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name *string `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the
	// VPC_PEERING and IPSEC_INTERCONNECT purposes.
	Network *string `pulumi:"network"`
	// The networking tier used for configuring this address. If this field is not specified, it is assumed to be PREMIUM. This
	// argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal
	// traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM",
	// "STANDARD"]
	NetworkTier *string `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength *float64 `pulumi:"prefixLength"`
	Project      *string  `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values. * GCE_ENDPOINT for addresses that are used by VM
	// instances, alias IP ranges, load balancers, and similar resources. * SHARED_LOADBALANCER_VIP for an address that can be
	// used by multiple internal load balancers. * VPC_PEERING for addresses that are reserved for VPC peer networks. *
	// IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an HA VPN
	// over Cloud Interconnect configuration. These addresses are regional resources. * PRIVATE_SERVICE_CONNECT for a private
	// network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	// This should only be set when using an Internal address.
	Purpose *string `pulumi:"purpose"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the
	// subnetwork's IP range. This field can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork *string                 `pulumi:"subnetwork"`
	Timeouts   *ComputeAddressTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeAddress resource.
type ComputeAddressArgs struct {
	// The static external IP address represented by this resource. The IP address must be inside the specified subnetwork, if
	// any. Set by the API if undefined.
	Address pulumi.StringPtrInput
	// The type of address to reserve. Note: if you set this argument's value as 'INTERNAL' you need to leave the
	// 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL",
	// "EXTERNAL"]
	AddressType      pulumi.StringPtrInput
	ComputeAddressId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]
	IpVersion pulumi.StringPtrInput
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this
	// address can be used after the external IPv6 address reservation. Possible values: ["VM", "NETLB"]
	Ipv6EndpointType pulumi.StringPtrInput
	// Labels to apply to this address. A list of key->value pairs. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringPtrInput
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the
	// VPC_PEERING and IPSEC_INTERCONNECT purposes.
	Network pulumi.StringPtrInput
	// The networking tier used for configuring this address. If this field is not specified, it is assumed to be PREMIUM. This
	// argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal
	// traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM",
	// "STANDARD"]
	NetworkTier pulumi.StringPtrInput
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.Float64PtrInput
	Project      pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values. * GCE_ENDPOINT for addresses that are used by VM
	// instances, alias IP ranges, load balancers, and similar resources. * SHARED_LOADBALANCER_VIP for an address that can be
	// used by multiple internal load balancers. * VPC_PEERING for addresses that are reserved for VPC peer networks. *
	// IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an HA VPN
	// over Cloud Interconnect configuration. These addresses are regional resources. * PRIVATE_SERVICE_CONNECT for a private
	// network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	// This should only be set when using an Internal address.
	Purpose pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the
	// subnetwork's IP range. This field can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork pulumi.StringPtrInput
	Timeouts   ComputeAddressTimeoutsPtrInput
}

func (ComputeAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeAddressArgs)(nil)).Elem()
}

type ComputeAddressInput interface {
	pulumi.Input

	ToComputeAddressOutput() ComputeAddressOutput
	ToComputeAddressOutputWithContext(ctx context.Context) ComputeAddressOutput
}

func (*ComputeAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeAddress)(nil)).Elem()
}

func (i *ComputeAddress) ToComputeAddressOutput() ComputeAddressOutput {
	return i.ToComputeAddressOutputWithContext(context.Background())
}

func (i *ComputeAddress) ToComputeAddressOutputWithContext(ctx context.Context) ComputeAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeAddressOutput)
}

type ComputeAddressOutput struct{ *pulumi.OutputState }

func (ComputeAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeAddress)(nil)).Elem()
}

func (o ComputeAddressOutput) ToComputeAddressOutput() ComputeAddressOutput {
	return o
}

func (o ComputeAddressOutput) ToComputeAddressOutputWithContext(ctx context.Context) ComputeAddressOutput {
	return o
}

// The static external IP address represented by this resource. The IP address must be inside the specified subnetwork, if
// any. Set by the API if undefined.
func (o ComputeAddressOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The type of address to reserve. Note: if you set this argument's value as 'INTERNAL' you need to leave the
// 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL",
// "EXTERNAL"]
func (o ComputeAddressOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringPtrOutput { return v.AddressType }).(pulumi.StringPtrOutput)
}

func (o ComputeAddressOutput) ComputeAddressId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.ComputeAddressId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeAddressOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeAddressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComputeAddressOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]
func (o ComputeAddressOutput) IpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringPtrOutput { return v.IpVersion }).(pulumi.StringPtrOutput)
}

// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this
// address can be used after the external IPv6 address reservation. Possible values: ["VM", "NETLB"]
func (o ComputeAddressOutput) Ipv6EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringPtrOutput { return v.Ipv6EndpointType }).(pulumi.StringPtrOutput)
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (o ComputeAddressOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this address. A list of key->value pairs. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o ComputeAddressOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
// dash.
func (o ComputeAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the
// VPC_PEERING and IPSEC_INTERCONNECT purposes.
func (o ComputeAddressOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// The networking tier used for configuring this address. If this field is not specified, it is assumed to be PREMIUM. This
// argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal
// traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM",
// "STANDARD"]
func (o ComputeAddressOutput) NetworkTier() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.NetworkTier }).(pulumi.StringOutput)
}

// The prefix length if the resource represents an IP range.
func (o ComputeAddressOutput) PrefixLength() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeAddress) pulumi.Float64Output { return v.PrefixLength }).(pulumi.Float64Output)
}

func (o ComputeAddressOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The purpose of this resource, which can be one of the following values. * GCE_ENDPOINT for addresses that are used by VM
// instances, alias IP ranges, load balancers, and similar resources. * SHARED_LOADBALANCER_VIP for an address that can be
// used by multiple internal load balancers. * VPC_PEERING for addresses that are reserved for VPC peer networks. *
// IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an HA VPN
// over Cloud Interconnect configuration. These addresses are regional resources. * PRIVATE_SERVICE_CONNECT for a private
// network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
// This should only be set when using an Internal address.
func (o ComputeAddressOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// The Region in which the created address should reside. If it is not provided, the provider region is used.
func (o ComputeAddressOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeAddressOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the
// subnetwork's IP range. This field can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.
func (o ComputeAddressOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringOutput { return v.Subnetwork }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ComputeAddressOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ComputeAddressOutput) Timeouts() ComputeAddressTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeAddress) ComputeAddressTimeoutsPtrOutput { return v.Timeouts }).(ComputeAddressTimeoutsPtrOutput)
}

// The URLs of the resources that are using this address.
func (o ComputeAddressOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeAddress) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeAddressInput)(nil)).Elem(), &ComputeAddress{})
	pulumi.RegisterOutputType(ComputeAddressOutput{})
}
