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

type VmwareengineNetworkPeering struct {
	pulumi.CustomResourceState

	// Creation time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User-provided description for this network peering.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// True if custom routes are exported to the peered network; false otherwise.
	ExportCustomRoutes pulumi.BoolPtrOutput `pulumi:"exportCustomRoutes"`
	// True if all subnet routes with a public IP address range are exported; false otherwise.
	ExportCustomRoutesWithPublicIp pulumi.BoolPtrOutput `pulumi:"exportCustomRoutesWithPublicIp"`
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutes pulumi.BoolPtrOutput `pulumi:"importCustomRoutes"`
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutesWithPublicIp pulumi.BoolPtrOutput `pulumi:"importCustomRoutesWithPublicIp"`
	// The ID of the Network Peering.
	Name pulumi.StringOutput `pulumi:"name"`
	// The relative resource name of the network to peer with a standard VMware Engine network. The provided network can be a
	// consumer VPC network or another standard VMware Engine network.
	PeerNetwork pulumi.StringOutput `pulumi:"peerNetwork"`
	// The type of the network to peer with the VMware Engine network. Possible values: ["STANDARD", "VMWARE_ENGINE_NETWORK",
	// "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE"]
	PeerNetworkType pulumi.StringOutput `pulumi:"peerNetworkType"`
	Project         pulumi.StringOutput `pulumi:"project"`
	// State of the network peering. This field has a value of 'ACTIVE' when there's a matching configuration in the peer
	// network. New values may be added to this enum when appropriate.
	State pulumi.StringOutput `pulumi:"state"`
	// Details about the current state of the network peering.
	StateDetails pulumi.StringOutput                         `pulumi:"stateDetails"`
	Timeouts     VmwareengineNetworkPeeringTimeoutsPtrOutput `pulumi:"timeouts"`
	// System-generated unique identifier for the resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Last updated time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The relative resource name of the VMware Engine network. Specify the name in the following form:
	// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project} can either be a
	// project number or a project ID.
	VmwareEngineNetwork pulumi.StringOutput `pulumi:"vmwareEngineNetwork"`
	// The canonical name of the VMware Engine network in the form:
	// projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
	VmwareEngineNetworkCanonical pulumi.StringOutput `pulumi:"vmwareEngineNetworkCanonical"`
	VmwareengineNetworkPeeringId pulumi.StringOutput `pulumi:"vmwareengineNetworkPeeringId"`
}

// NewVmwareengineNetworkPeering registers a new resource with the given unique name, arguments, and options.
func NewVmwareengineNetworkPeering(ctx *pulumi.Context,
	name string, args *VmwareengineNetworkPeeringArgs, opts ...pulumi.ResourceOption) (*VmwareengineNetworkPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerNetwork == nil {
		return nil, errors.New("invalid value for required argument 'PeerNetwork'")
	}
	if args.PeerNetworkType == nil {
		return nil, errors.New("invalid value for required argument 'PeerNetworkType'")
	}
	if args.VmwareEngineNetwork == nil {
		return nil, errors.New("invalid value for required argument 'VmwareEngineNetwork'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource VmwareengineNetworkPeering
	err = ctx.RegisterPackageResource("google:index/vmwareengineNetworkPeering:VmwareengineNetworkPeering", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmwareengineNetworkPeering gets an existing VmwareengineNetworkPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmwareengineNetworkPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmwareengineNetworkPeeringState, opts ...pulumi.ResourceOption) (*VmwareengineNetworkPeering, error) {
	var resource VmwareengineNetworkPeering
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/vmwareengineNetworkPeering:VmwareengineNetworkPeering", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmwareengineNetworkPeering resources.
type vmwareengineNetworkPeeringState struct {
	// Creation time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// User-provided description for this network peering.
	Description *string `pulumi:"description"`
	// True if custom routes are exported to the peered network; false otherwise.
	ExportCustomRoutes *bool `pulumi:"exportCustomRoutes"`
	// True if all subnet routes with a public IP address range are exported; false otherwise.
	ExportCustomRoutesWithPublicIp *bool `pulumi:"exportCustomRoutesWithPublicIp"`
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutes *bool `pulumi:"importCustomRoutes"`
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutesWithPublicIp *bool `pulumi:"importCustomRoutesWithPublicIp"`
	// The ID of the Network Peering.
	Name *string `pulumi:"name"`
	// The relative resource name of the network to peer with a standard VMware Engine network. The provided network can be a
	// consumer VPC network or another standard VMware Engine network.
	PeerNetwork *string `pulumi:"peerNetwork"`
	// The type of the network to peer with the VMware Engine network. Possible values: ["STANDARD", "VMWARE_ENGINE_NETWORK",
	// "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE"]
	PeerNetworkType *string `pulumi:"peerNetworkType"`
	Project         *string `pulumi:"project"`
	// State of the network peering. This field has a value of 'ACTIVE' when there's a matching configuration in the peer
	// network. New values may be added to this enum when appropriate.
	State *string `pulumi:"state"`
	// Details about the current state of the network peering.
	StateDetails *string                             `pulumi:"stateDetails"`
	Timeouts     *VmwareengineNetworkPeeringTimeouts `pulumi:"timeouts"`
	// System-generated unique identifier for the resource.
	Uid *string `pulumi:"uid"`
	// Last updated time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
	// The relative resource name of the VMware Engine network. Specify the name in the following form:
	// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project} can either be a
	// project number or a project ID.
	VmwareEngineNetwork *string `pulumi:"vmwareEngineNetwork"`
	// The canonical name of the VMware Engine network in the form:
	// projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
	VmwareEngineNetworkCanonical *string `pulumi:"vmwareEngineNetworkCanonical"`
	VmwareengineNetworkPeeringId *string `pulumi:"vmwareengineNetworkPeeringId"`
}

type VmwareengineNetworkPeeringState struct {
	// Creation time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// User-provided description for this network peering.
	Description pulumi.StringPtrInput
	// True if custom routes are exported to the peered network; false otherwise.
	ExportCustomRoutes pulumi.BoolPtrInput
	// True if all subnet routes with a public IP address range are exported; false otherwise.
	ExportCustomRoutesWithPublicIp pulumi.BoolPtrInput
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutes pulumi.BoolPtrInput
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutesWithPublicIp pulumi.BoolPtrInput
	// The ID of the Network Peering.
	Name pulumi.StringPtrInput
	// The relative resource name of the network to peer with a standard VMware Engine network. The provided network can be a
	// consumer VPC network or another standard VMware Engine network.
	PeerNetwork pulumi.StringPtrInput
	// The type of the network to peer with the VMware Engine network. Possible values: ["STANDARD", "VMWARE_ENGINE_NETWORK",
	// "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE"]
	PeerNetworkType pulumi.StringPtrInput
	Project         pulumi.StringPtrInput
	// State of the network peering. This field has a value of 'ACTIVE' when there's a matching configuration in the peer
	// network. New values may be added to this enum when appropriate.
	State pulumi.StringPtrInput
	// Details about the current state of the network peering.
	StateDetails pulumi.StringPtrInput
	Timeouts     VmwareengineNetworkPeeringTimeoutsPtrInput
	// System-generated unique identifier for the resource.
	Uid pulumi.StringPtrInput
	// Last updated time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
	// The relative resource name of the VMware Engine network. Specify the name in the following form:
	// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project} can either be a
	// project number or a project ID.
	VmwareEngineNetwork pulumi.StringPtrInput
	// The canonical name of the VMware Engine network in the form:
	// projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
	VmwareEngineNetworkCanonical pulumi.StringPtrInput
	VmwareengineNetworkPeeringId pulumi.StringPtrInput
}

func (VmwareengineNetworkPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareengineNetworkPeeringState)(nil)).Elem()
}

type vmwareengineNetworkPeeringArgs struct {
	// User-provided description for this network peering.
	Description *string `pulumi:"description"`
	// True if custom routes are exported to the peered network; false otherwise.
	ExportCustomRoutes *bool `pulumi:"exportCustomRoutes"`
	// True if all subnet routes with a public IP address range are exported; false otherwise.
	ExportCustomRoutesWithPublicIp *bool `pulumi:"exportCustomRoutesWithPublicIp"`
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutes *bool `pulumi:"importCustomRoutes"`
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutesWithPublicIp *bool `pulumi:"importCustomRoutesWithPublicIp"`
	// The ID of the Network Peering.
	Name *string `pulumi:"name"`
	// The relative resource name of the network to peer with a standard VMware Engine network. The provided network can be a
	// consumer VPC network or another standard VMware Engine network.
	PeerNetwork string `pulumi:"peerNetwork"`
	// The type of the network to peer with the VMware Engine network. Possible values: ["STANDARD", "VMWARE_ENGINE_NETWORK",
	// "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE"]
	PeerNetworkType string                              `pulumi:"peerNetworkType"`
	Project         *string                             `pulumi:"project"`
	Timeouts        *VmwareengineNetworkPeeringTimeouts `pulumi:"timeouts"`
	// The relative resource name of the VMware Engine network. Specify the name in the following form:
	// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project} can either be a
	// project number or a project ID.
	VmwareEngineNetwork          string  `pulumi:"vmwareEngineNetwork"`
	VmwareengineNetworkPeeringId *string `pulumi:"vmwareengineNetworkPeeringId"`
}

// The set of arguments for constructing a VmwareengineNetworkPeering resource.
type VmwareengineNetworkPeeringArgs struct {
	// User-provided description for this network peering.
	Description pulumi.StringPtrInput
	// True if custom routes are exported to the peered network; false otherwise.
	ExportCustomRoutes pulumi.BoolPtrInput
	// True if all subnet routes with a public IP address range are exported; false otherwise.
	ExportCustomRoutesWithPublicIp pulumi.BoolPtrInput
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutes pulumi.BoolPtrInput
	// True if custom routes are imported from the peered network; false otherwise.
	ImportCustomRoutesWithPublicIp pulumi.BoolPtrInput
	// The ID of the Network Peering.
	Name pulumi.StringPtrInput
	// The relative resource name of the network to peer with a standard VMware Engine network. The provided network can be a
	// consumer VPC network or another standard VMware Engine network.
	PeerNetwork pulumi.StringInput
	// The type of the network to peer with the VMware Engine network. Possible values: ["STANDARD", "VMWARE_ENGINE_NETWORK",
	// "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE"]
	PeerNetworkType pulumi.StringInput
	Project         pulumi.StringPtrInput
	Timeouts        VmwareengineNetworkPeeringTimeoutsPtrInput
	// The relative resource name of the VMware Engine network. Specify the name in the following form:
	// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project} can either be a
	// project number or a project ID.
	VmwareEngineNetwork          pulumi.StringInput
	VmwareengineNetworkPeeringId pulumi.StringPtrInput
}

func (VmwareengineNetworkPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareengineNetworkPeeringArgs)(nil)).Elem()
}

type VmwareengineNetworkPeeringInput interface {
	pulumi.Input

	ToVmwareengineNetworkPeeringOutput() VmwareengineNetworkPeeringOutput
	ToVmwareengineNetworkPeeringOutputWithContext(ctx context.Context) VmwareengineNetworkPeeringOutput
}

func (*VmwareengineNetworkPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareengineNetworkPeering)(nil)).Elem()
}

func (i *VmwareengineNetworkPeering) ToVmwareengineNetworkPeeringOutput() VmwareengineNetworkPeeringOutput {
	return i.ToVmwareengineNetworkPeeringOutputWithContext(context.Background())
}

func (i *VmwareengineNetworkPeering) ToVmwareengineNetworkPeeringOutputWithContext(ctx context.Context) VmwareengineNetworkPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmwareengineNetworkPeeringOutput)
}

type VmwareengineNetworkPeeringOutput struct{ *pulumi.OutputState }

func (VmwareengineNetworkPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareengineNetworkPeering)(nil)).Elem()
}

func (o VmwareengineNetworkPeeringOutput) ToVmwareengineNetworkPeeringOutput() VmwareengineNetworkPeeringOutput {
	return o
}

func (o VmwareengineNetworkPeeringOutput) ToVmwareengineNetworkPeeringOutputWithContext(ctx context.Context) VmwareengineNetworkPeeringOutput {
	return o
}

// Creation time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o VmwareengineNetworkPeeringOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// User-provided description for this network peering.
func (o VmwareengineNetworkPeeringOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// True if custom routes are exported to the peered network; false otherwise.
func (o VmwareengineNetworkPeeringOutput) ExportCustomRoutes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.BoolPtrOutput { return v.ExportCustomRoutes }).(pulumi.BoolPtrOutput)
}

// True if all subnet routes with a public IP address range are exported; false otherwise.
func (o VmwareengineNetworkPeeringOutput) ExportCustomRoutesWithPublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.BoolPtrOutput { return v.ExportCustomRoutesWithPublicIp }).(pulumi.BoolPtrOutput)
}

// True if custom routes are imported from the peered network; false otherwise.
func (o VmwareengineNetworkPeeringOutput) ImportCustomRoutes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.BoolPtrOutput { return v.ImportCustomRoutes }).(pulumi.BoolPtrOutput)
}

// True if custom routes are imported from the peered network; false otherwise.
func (o VmwareengineNetworkPeeringOutput) ImportCustomRoutesWithPublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.BoolPtrOutput { return v.ImportCustomRoutesWithPublicIp }).(pulumi.BoolPtrOutput)
}

// The ID of the Network Peering.
func (o VmwareengineNetworkPeeringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The relative resource name of the network to peer with a standard VMware Engine network. The provided network can be a
// consumer VPC network or another standard VMware Engine network.
func (o VmwareengineNetworkPeeringOutput) PeerNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.PeerNetwork }).(pulumi.StringOutput)
}

// The type of the network to peer with the VMware Engine network. Possible values: ["STANDARD", "VMWARE_ENGINE_NETWORK",
// "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE"]
func (o VmwareengineNetworkPeeringOutput) PeerNetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.PeerNetworkType }).(pulumi.StringOutput)
}

func (o VmwareengineNetworkPeeringOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// State of the network peering. This field has a value of 'ACTIVE' when there's a matching configuration in the peer
// network. New values may be added to this enum when appropriate.
func (o VmwareengineNetworkPeeringOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Details about the current state of the network peering.
func (o VmwareengineNetworkPeeringOutput) StateDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.StateDetails }).(pulumi.StringOutput)
}

func (o VmwareengineNetworkPeeringOutput) Timeouts() VmwareengineNetworkPeeringTimeoutsPtrOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) VmwareengineNetworkPeeringTimeoutsPtrOutput { return v.Timeouts }).(VmwareengineNetworkPeeringTimeoutsPtrOutput)
}

// System-generated unique identifier for the resource.
func (o VmwareengineNetworkPeeringOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Last updated time of this resource. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o VmwareengineNetworkPeeringOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The relative resource name of the VMware Engine network. Specify the name in the following form:
// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project} can either be a
// project number or a project ID.
func (o VmwareengineNetworkPeeringOutput) VmwareEngineNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.VmwareEngineNetwork }).(pulumi.StringOutput)
}

// The canonical name of the VMware Engine network in the form:
// projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
func (o VmwareengineNetworkPeeringOutput) VmwareEngineNetworkCanonical() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.VmwareEngineNetworkCanonical }).(pulumi.StringOutput)
}

func (o VmwareengineNetworkPeeringOutput) VmwareengineNetworkPeeringId() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareengineNetworkPeering) pulumi.StringOutput { return v.VmwareengineNetworkPeeringId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmwareengineNetworkPeeringInput)(nil)).Elem(), &VmwareengineNetworkPeering{})
	pulumi.RegisterOutputType(VmwareengineNetworkPeeringOutput{})
}
