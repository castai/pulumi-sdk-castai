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

type ComputeSubnetwork struct {
	pulumi.CustomResourceState

	// Typically packets destined to IPs within the subnetwork range that do not match existing resources are dropped and
	// prevented from leaving the VPC. Setting this field to true will allow these packets to match dynamic routes injected via
	// BGP even if their destinations match existing subnet ranges.
	AllowSubnetCidrRoutesOverlap pulumi.BoolOutput   `pulumi:"allowSubnetCidrRoutesOverlap"`
	ComputeSubnetworkId          pulumi.StringOutput `pulumi:"computeSubnetworkId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only
	// at resource creation time.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIpv6Prefix pulumi.StringOutput `pulumi:"externalIpv6Prefix"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: Deprecated
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringOutput `pulumi:"gatewayAddress"`
	// The internal IPv6 address range that is assigned to this subnetwork.
	InternalIpv6Prefix pulumi.StringOutput `pulumi:"internalIpv6Prefix"`
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork.
	// For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is
	// supported. Field is optional when 'reserved_internal_range' is defined, otherwise required.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first
	// time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct
	// path. Possible values: ["EXTERNAL", "INTERNAL"]
	Ipv6AccessType pulumi.StringPtrOutput `pulumi:"ipv6AccessType"`
	// The range of internal IPv6 addresses that are owned by this subnetwork.
	Ipv6CidrRange pulumi.StringOutput `pulumi:"ipv6CidrRange"`
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud
	// Logging. Flow logging isn't supported if the subnet 'purpose' field is set to subnetwork is 'REGIONAL_MANAGED_PROXY' or
	// 'GLOBAL_MANAGED_PROXY'.
	LogConfig ComputeSubnetworkLogConfigPtrOutput `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters
	// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// 'a-z?' which means the first character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringOutput `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private
	// Google Access.
	PrivateIpGoogleAccess pulumi.BoolOutput `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess pulumi.StringOutput `pulumi:"privateIpv6GoogleAccess"`
	Project                 pulumi.StringOutput `pulumi:"project"`
	Purpose                 pulumi.StringOutput `pulumi:"purpose"`
	// The GCP region for this subnetwork.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g.
	// 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'
	ReservedInternalRange pulumi.StringPtrOutput `pulumi:"reservedInternalRange"`
	// The role of subnetwork. Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'. The value can be
	// set to 'ACTIVE' or 'BACKUP'. An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers
	// in a region. A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible
	// values: ["ACTIVE", "BACKUP"]
	Role              pulumi.StringPtrOutput                       `pulumi:"role"`
	SecondaryIpRanges ComputeSubnetworkSecondaryIpRangeArrayOutput `pulumi:"secondaryIpRanges"`
	SelfLink          pulumi.StringOutput                          `pulumi:"selfLink"`
	// Controls the removal behavior of secondary_ip_range. When false, removing secondary_ip_range from config will not
	// produce a diff as the provider will default to the API's value. When true, the provider will treat removing
	// secondary_ip_range as sending an empty list of secondary IP ranges to the API. Defaults to false.
	SendSecondaryIpRangeIfEmpty pulumi.BoolPtrOutput `pulumi:"sendSecondaryIpRangeIfEmpty"`
	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will
	// be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType pulumi.StringOutput `pulumi:"stackType"`
	// The unique identifier number for the resource. This identifier is defined by the server.
	SubnetworkId pulumi.Float64Output               `pulumi:"subnetworkId"`
	Timeouts     ComputeSubnetworkTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeSubnetwork registers a new resource with the given unique name, arguments, and options.
func NewComputeSubnetwork(ctx *pulumi.Context,
	name string, args *ComputeSubnetworkArgs, opts ...pulumi.ResourceOption) (*ComputeSubnetwork, error) {
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
	var resource ComputeSubnetwork
	err = ctx.RegisterPackageResource("google-beta:index/computeSubnetwork:ComputeSubnetwork", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeSubnetwork gets an existing ComputeSubnetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeSubnetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeSubnetworkState, opts ...pulumi.ResourceOption) (*ComputeSubnetwork, error) {
	var resource ComputeSubnetwork
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeSubnetwork:ComputeSubnetwork", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeSubnetwork resources.
type computeSubnetworkState struct {
	// Typically packets destined to IPs within the subnetwork range that do not match existing resources are dropped and
	// prevented from leaving the VPC. Setting this field to true will allow these packets to match dynamic routes injected via
	// BGP even if their destinations match existing subnet ranges.
	AllowSubnetCidrRoutesOverlap *bool   `pulumi:"allowSubnetCidrRoutesOverlap"`
	ComputeSubnetworkId          *string `pulumi:"computeSubnetworkId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only
	// at resource creation time.
	Description *string `pulumi:"description"`
	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIpv6Prefix *string `pulumi:"externalIpv6Prefix"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: Deprecated
	Fingerprint *string `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress *string `pulumi:"gatewayAddress"`
	// The internal IPv6 address range that is assigned to this subnetwork.
	InternalIpv6Prefix *string `pulumi:"internalIpv6Prefix"`
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork.
	// For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is
	// supported. Field is optional when 'reserved_internal_range' is defined, otherwise required.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first
	// time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct
	// path. Possible values: ["EXTERNAL", "INTERNAL"]
	Ipv6AccessType *string `pulumi:"ipv6AccessType"`
	// The range of internal IPv6 addresses that are owned by this subnetwork.
	Ipv6CidrRange *string `pulumi:"ipv6CidrRange"`
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud
	// Logging. Flow logging isn't supported if the subnet 'purpose' field is set to subnetwork is 'REGIONAL_MANAGED_PROXY' or
	// 'GLOBAL_MANAGED_PROXY'.
	LogConfig *ComputeSubnetworkLogConfig `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters
	// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// 'a-z?' which means the first character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
	Network *string `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private
	// Google Access.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess *string `pulumi:"privateIpv6GoogleAccess"`
	Project                 *string `pulumi:"project"`
	Purpose                 *string `pulumi:"purpose"`
	// The GCP region for this subnetwork.
	Region *string `pulumi:"region"`
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g.
	// 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'
	ReservedInternalRange *string `pulumi:"reservedInternalRange"`
	// The role of subnetwork. Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'. The value can be
	// set to 'ACTIVE' or 'BACKUP'. An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers
	// in a region. A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible
	// values: ["ACTIVE", "BACKUP"]
	Role              *string                             `pulumi:"role"`
	SecondaryIpRanges []ComputeSubnetworkSecondaryIpRange `pulumi:"secondaryIpRanges"`
	SelfLink          *string                             `pulumi:"selfLink"`
	// Controls the removal behavior of secondary_ip_range. When false, removing secondary_ip_range from config will not
	// produce a diff as the provider will default to the API's value. When true, the provider will treat removing
	// secondary_ip_range as sending an empty list of secondary IP ranges to the API. Defaults to false.
	SendSecondaryIpRangeIfEmpty *bool `pulumi:"sendSecondaryIpRangeIfEmpty"`
	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will
	// be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType *string `pulumi:"stackType"`
	// The unique identifier number for the resource. This identifier is defined by the server.
	SubnetworkId *float64                   `pulumi:"subnetworkId"`
	Timeouts     *ComputeSubnetworkTimeouts `pulumi:"timeouts"`
}

type ComputeSubnetworkState struct {
	// Typically packets destined to IPs within the subnetwork range that do not match existing resources are dropped and
	// prevented from leaving the VPC. Setting this field to true will allow these packets to match dynamic routes injected via
	// BGP even if their destinations match existing subnet ranges.
	AllowSubnetCidrRoutesOverlap pulumi.BoolPtrInput
	ComputeSubnetworkId          pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only
	// at resource creation time.
	Description pulumi.StringPtrInput
	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIpv6Prefix pulumi.StringPtrInput
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: Deprecated
	Fingerprint pulumi.StringPtrInput
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringPtrInput
	// The internal IPv6 address range that is assigned to this subnetwork.
	InternalIpv6Prefix pulumi.StringPtrInput
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork.
	// For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is
	// supported. Field is optional when 'reserved_internal_range' is defined, otherwise required.
	IpCidrRange pulumi.StringPtrInput
	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first
	// time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct
	// path. Possible values: ["EXTERNAL", "INTERNAL"]
	Ipv6AccessType pulumi.StringPtrInput
	// The range of internal IPv6 addresses that are owned by this subnetwork.
	Ipv6CidrRange pulumi.StringPtrInput
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud
	// Logging. Flow logging isn't supported if the subnet 'purpose' field is set to subnetwork is 'REGIONAL_MANAGED_PROXY' or
	// 'GLOBAL_MANAGED_PROXY'.
	LogConfig ComputeSubnetworkLogConfigPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters
	// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// 'a-z?' which means the first character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringPtrInput
	// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private
	// Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess pulumi.StringPtrInput
	Project                 pulumi.StringPtrInput
	Purpose                 pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	Region pulumi.StringPtrInput
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g.
	// 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'
	ReservedInternalRange pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'. The value can be
	// set to 'ACTIVE' or 'BACKUP'. An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers
	// in a region. A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible
	// values: ["ACTIVE", "BACKUP"]
	Role              pulumi.StringPtrInput
	SecondaryIpRanges ComputeSubnetworkSecondaryIpRangeArrayInput
	SelfLink          pulumi.StringPtrInput
	// Controls the removal behavior of secondary_ip_range. When false, removing secondary_ip_range from config will not
	// produce a diff as the provider will default to the API's value. When true, the provider will treat removing
	// secondary_ip_range as sending an empty list of secondary IP ranges to the API. Defaults to false.
	SendSecondaryIpRangeIfEmpty pulumi.BoolPtrInput
	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will
	// be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType pulumi.StringPtrInput
	// The unique identifier number for the resource. This identifier is defined by the server.
	SubnetworkId pulumi.Float64PtrInput
	Timeouts     ComputeSubnetworkTimeoutsPtrInput
}

func (ComputeSubnetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeSubnetworkState)(nil)).Elem()
}

type computeSubnetworkArgs struct {
	// Typically packets destined to IPs within the subnetwork range that do not match existing resources are dropped and
	// prevented from leaving the VPC. Setting this field to true will allow these packets to match dynamic routes injected via
	// BGP even if their destinations match existing subnet ranges.
	AllowSubnetCidrRoutesOverlap *bool   `pulumi:"allowSubnetCidrRoutesOverlap"`
	ComputeSubnetworkId          *string `pulumi:"computeSubnetworkId"`
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only
	// at resource creation time.
	Description *string `pulumi:"description"`
	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIpv6Prefix *string `pulumi:"externalIpv6Prefix"`
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork.
	// For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is
	// supported. Field is optional when 'reserved_internal_range' is defined, otherwise required.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first
	// time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct
	// path. Possible values: ["EXTERNAL", "INTERNAL"]
	Ipv6AccessType *string `pulumi:"ipv6AccessType"`
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud
	// Logging. Flow logging isn't supported if the subnet 'purpose' field is set to subnetwork is 'REGIONAL_MANAGED_PROXY' or
	// 'GLOBAL_MANAGED_PROXY'.
	LogConfig *ComputeSubnetworkLogConfig `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters
	// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// 'a-z?' which means the first character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
	Network string `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private
	// Google Access.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess *string `pulumi:"privateIpv6GoogleAccess"`
	Project                 *string `pulumi:"project"`
	Purpose                 *string `pulumi:"purpose"`
	// The GCP region for this subnetwork.
	Region *string `pulumi:"region"`
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g.
	// 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'
	ReservedInternalRange *string `pulumi:"reservedInternalRange"`
	// The role of subnetwork. Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'. The value can be
	// set to 'ACTIVE' or 'BACKUP'. An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers
	// in a region. A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible
	// values: ["ACTIVE", "BACKUP"]
	Role              *string                             `pulumi:"role"`
	SecondaryIpRanges []ComputeSubnetworkSecondaryIpRange `pulumi:"secondaryIpRanges"`
	// Controls the removal behavior of secondary_ip_range. When false, removing secondary_ip_range from config will not
	// produce a diff as the provider will default to the API's value. When true, the provider will treat removing
	// secondary_ip_range as sending an empty list of secondary IP ranges to the API. Defaults to false.
	SendSecondaryIpRangeIfEmpty *bool `pulumi:"sendSecondaryIpRangeIfEmpty"`
	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will
	// be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType *string                    `pulumi:"stackType"`
	Timeouts  *ComputeSubnetworkTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeSubnetwork resource.
type ComputeSubnetworkArgs struct {
	// Typically packets destined to IPs within the subnetwork range that do not match existing resources are dropped and
	// prevented from leaving the VPC. Setting this field to true will allow these packets to match dynamic routes injected via
	// BGP even if their destinations match existing subnet ranges.
	AllowSubnetCidrRoutesOverlap pulumi.BoolPtrInput
	ComputeSubnetworkId          pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only
	// at resource creation time.
	Description pulumi.StringPtrInput
	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIpv6Prefix pulumi.StringPtrInput
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork.
	// For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is
	// supported. Field is optional when 'reserved_internal_range' is defined, otherwise required.
	IpCidrRange pulumi.StringPtrInput
	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first
	// time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct
	// path. Possible values: ["EXTERNAL", "INTERNAL"]
	Ipv6AccessType pulumi.StringPtrInput
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud
	// Logging. Flow logging isn't supported if the subnet 'purpose' field is set to subnetwork is 'REGIONAL_MANAGED_PROXY' or
	// 'GLOBAL_MANAGED_PROXY'.
	LogConfig ComputeSubnetworkLogConfigPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters
	// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// 'a-z?' which means the first character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringInput
	// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private
	// Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess pulumi.StringPtrInput
	Project                 pulumi.StringPtrInput
	Purpose                 pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	Region pulumi.StringPtrInput
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g.
	// 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'
	ReservedInternalRange pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'. The value can be
	// set to 'ACTIVE' or 'BACKUP'. An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers
	// in a region. A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible
	// values: ["ACTIVE", "BACKUP"]
	Role              pulumi.StringPtrInput
	SecondaryIpRanges ComputeSubnetworkSecondaryIpRangeArrayInput
	// Controls the removal behavior of secondary_ip_range. When false, removing secondary_ip_range from config will not
	// produce a diff as the provider will default to the API's value. When true, the provider will treat removing
	// secondary_ip_range as sending an empty list of secondary IP ranges to the API. Defaults to false.
	SendSecondaryIpRangeIfEmpty pulumi.BoolPtrInput
	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will
	// be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	StackType pulumi.StringPtrInput
	Timeouts  ComputeSubnetworkTimeoutsPtrInput
}

func (ComputeSubnetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeSubnetworkArgs)(nil)).Elem()
}

type ComputeSubnetworkInput interface {
	pulumi.Input

	ToComputeSubnetworkOutput() ComputeSubnetworkOutput
	ToComputeSubnetworkOutputWithContext(ctx context.Context) ComputeSubnetworkOutput
}

func (*ComputeSubnetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeSubnetwork)(nil)).Elem()
}

func (i *ComputeSubnetwork) ToComputeSubnetworkOutput() ComputeSubnetworkOutput {
	return i.ToComputeSubnetworkOutputWithContext(context.Background())
}

func (i *ComputeSubnetwork) ToComputeSubnetworkOutputWithContext(ctx context.Context) ComputeSubnetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeSubnetworkOutput)
}

type ComputeSubnetworkOutput struct{ *pulumi.OutputState }

func (ComputeSubnetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeSubnetwork)(nil)).Elem()
}

func (o ComputeSubnetworkOutput) ToComputeSubnetworkOutput() ComputeSubnetworkOutput {
	return o
}

func (o ComputeSubnetworkOutput) ToComputeSubnetworkOutputWithContext(ctx context.Context) ComputeSubnetworkOutput {
	return o
}

// Typically packets destined to IPs within the subnetwork range that do not match existing resources are dropped and
// prevented from leaving the VPC. Setting this field to true will allow these packets to match dynamic routes injected via
// BGP even if their destinations match existing subnet ranges.
func (o ComputeSubnetworkOutput) AllowSubnetCidrRoutesOverlap() pulumi.BoolOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.BoolOutput { return v.AllowSubnetCidrRoutesOverlap }).(pulumi.BoolOutput)
}

func (o ComputeSubnetworkOutput) ComputeSubnetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.ComputeSubnetworkId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeSubnetworkOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource. This field can be set only
// at resource creation time.
func (o ComputeSubnetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The range of external IPv6 addresses that are owned by this subnetwork.
func (o ComputeSubnetworkOutput) ExternalIpv6Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.ExternalIpv6Prefix }).(pulumi.StringOutput)
}

// Fingerprint of this resource. This field is used internally during updates of this resource.
//
// Deprecated: Deprecated
func (o ComputeSubnetworkOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The gateway address for default routes to reach destination addresses outside this subnetwork.
func (o ComputeSubnetworkOutput) GatewayAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.GatewayAddress }).(pulumi.StringOutput)
}

// The internal IPv6 address range that is assigned to this subnetwork.
func (o ComputeSubnetworkOutput) InternalIpv6Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.InternalIpv6Prefix }).(pulumi.StringOutput)
}

// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork.
// For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is
// supported. Field is optional when 'reserved_internal_range' is defined, otherwise required.
func (o ComputeSubnetworkOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.IpCidrRange }).(pulumi.StringOutput)
}

// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first
// time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct
// path. Possible values: ["EXTERNAL", "INTERNAL"]
func (o ComputeSubnetworkOutput) Ipv6AccessType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringPtrOutput { return v.Ipv6AccessType }).(pulumi.StringPtrOutput)
}

// The range of internal IPv6 addresses that are owned by this subnetwork.
func (o ComputeSubnetworkOutput) Ipv6CidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.Ipv6CidrRange }).(pulumi.StringOutput)
}

// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud
// Logging. Flow logging isn't supported if the subnet 'purpose' field is set to subnetwork is 'REGIONAL_MANAGED_PROXY' or
// 'GLOBAL_MANAGED_PROXY'.
func (o ComputeSubnetworkOutput) LogConfig() ComputeSubnetworkLogConfigPtrOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) ComputeSubnetworkLogConfigPtrOutput { return v.LogConfig }).(ComputeSubnetworkLogConfigPtrOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters
// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// 'a-z?' which means the first character must be a lowercase letter, and all following characters must be a dash,
// lowercase letter, or digit, except the last character, which cannot be a dash.
func (o ComputeSubnetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
func (o ComputeSubnetworkOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private
// Google Access.
func (o ComputeSubnetworkOutput) PrivateIpGoogleAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.BoolOutput { return v.PrivateIpGoogleAccess }).(pulumi.BoolOutput)
}

// The private IPv6 google access type for the VMs in this subnet.
func (o ComputeSubnetworkOutput) PrivateIpv6GoogleAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.PrivateIpv6GoogleAccess }).(pulumi.StringOutput)
}

func (o ComputeSubnetworkOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeSubnetworkOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// The GCP region for this subnetwork.
func (o ComputeSubnetworkOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g.
// 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'
func (o ComputeSubnetworkOutput) ReservedInternalRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringPtrOutput { return v.ReservedInternalRange }).(pulumi.StringPtrOutput)
}

// The role of subnetwork. Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'. The value can be
// set to 'ACTIVE' or 'BACKUP'. An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers
// in a region. A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible
// values: ["ACTIVE", "BACKUP"]
func (o ComputeSubnetworkOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

func (o ComputeSubnetworkOutput) SecondaryIpRanges() ComputeSubnetworkSecondaryIpRangeArrayOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) ComputeSubnetworkSecondaryIpRangeArrayOutput { return v.SecondaryIpRanges }).(ComputeSubnetworkSecondaryIpRangeArrayOutput)
}

func (o ComputeSubnetworkOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Controls the removal behavior of secondary_ip_range. When false, removing secondary_ip_range from config will not
// produce a diff as the provider will default to the API's value. When true, the provider will treat removing
// secondary_ip_range as sending an empty list of secondary IP ranges to the API. Defaults to false.
func (o ComputeSubnetworkOutput) SendSecondaryIpRangeIfEmpty() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.BoolPtrOutput { return v.SendSecondaryIpRangeIfEmpty }).(pulumi.BoolPtrOutput)
}

// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will
// be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
func (o ComputeSubnetworkOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.StringOutput { return v.StackType }).(pulumi.StringOutput)
}

// The unique identifier number for the resource. This identifier is defined by the server.
func (o ComputeSubnetworkOutput) SubnetworkId() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeSubnetwork) pulumi.Float64Output { return v.SubnetworkId }).(pulumi.Float64Output)
}

func (o ComputeSubnetworkOutput) Timeouts() ComputeSubnetworkTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeSubnetwork) ComputeSubnetworkTimeoutsPtrOutput { return v.Timeouts }).(ComputeSubnetworkTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeSubnetworkInput)(nil)).Elem(), &ComputeSubnetwork{})
	pulumi.RegisterOutputType(ComputeSubnetworkOutput{})
}
