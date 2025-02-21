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

type TpuNode struct {
	pulumi.CustomResourceState

	// The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringOutput `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The immutable name of the TPU.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of a network to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on
	// which this API has been activated. If none is provided, "default" will be used.
	Network pulumi.StringOutput `pulumi:"network"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
	// node first reach out to the first (index 0) entry.
	NetworkEndpoints TpuNodeNetworkEndpointArrayOutput `pulumi:"networkEndpoints"`
	Project          pulumi.StringOutput               `pulumi:"project"`
	// Sets the scheduling options for this TPU instance.
	SchedulingConfig TpuNodeSchedulingConfigPtrOutput `pulumi:"schedulingConfig"`
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
	// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringOutput `pulumi:"tensorflowVersion"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput   `pulumi:"terraformLabels"`
	Timeouts        TpuNodeTimeoutsPtrOutput `pulumi:"timeouts"`
	TpuNodeId       pulumi.StringOutput      `pulumi:"tpuNodeId"`
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before
	// provisioning the node. If this field is set, cidr_block field should not be specified. If the network that you want to
	// peer the TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking pulumi.BoolPtrOutput `pulumi:"useServiceNetworking"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTpuNode registers a new resource with the given unique name, arguments, and options.
func NewTpuNode(ctx *pulumi.Context,
	name string, args *TpuNodeArgs, opts ...pulumi.ResourceOption) (*TpuNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorType == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorType'")
	}
	if args.TensorflowVersion == nil {
		return nil, errors.New("invalid value for required argument 'TensorflowVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource TpuNode
	err = ctx.RegisterPackageResource("google:index/tpuNode:TpuNode", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTpuNode gets an existing TpuNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTpuNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TpuNodeState, opts ...pulumi.ResourceOption) (*TpuNode, error) {
	var resource TpuNode
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/tpuNode:TpuNode", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TpuNode resources.
type tpuNodeState struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType *string `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The immutable name of the TPU.
	Name *string `pulumi:"name"`
	// The name of a network to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on
	// which this API has been activated. If none is provided, "default" will be used.
	Network *string `pulumi:"network"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
	// node first reach out to the first (index 0) entry.
	NetworkEndpoints []TpuNodeNetworkEndpoint `pulumi:"networkEndpoints"`
	Project          *string                  `pulumi:"project"`
	// Sets the scheduling options for this TPU instance.
	SchedulingConfig *TpuNodeSchedulingConfig `pulumi:"schedulingConfig"`
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
	// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The version of Tensorflow running in the Node.
	TensorflowVersion *string `pulumi:"tensorflowVersion"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string `pulumi:"terraformLabels"`
	Timeouts        *TpuNodeTimeouts  `pulumi:"timeouts"`
	TpuNodeId       *string           `pulumi:"tpuNodeId"`
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before
	// provisioning the node. If this field is set, cidr_block field should not be specified. If the network that you want to
	// peer the TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking *bool `pulumi:"useServiceNetworking"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type TpuNodeState struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringPtrInput
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The immutable name of the TPU.
	Name pulumi.StringPtrInput
	// The name of a network to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on
	// which this API has been activated. If none is provided, "default" will be used.
	Network pulumi.StringPtrInput
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
	// node first reach out to the first (index 0) entry.
	NetworkEndpoints TpuNodeNetworkEndpointArrayInput
	Project          pulumi.StringPtrInput
	// Sets the scheduling options for this TPU instance.
	SchedulingConfig TpuNodeSchedulingConfigPtrInput
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
	// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount pulumi.StringPtrInput
	// The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        TpuNodeTimeoutsPtrInput
	TpuNodeId       pulumi.StringPtrInput
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before
	// provisioning the node. If this field is set, cidr_block field should not be specified. If the network that you want to
	// peer the TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking pulumi.BoolPtrInput
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (TpuNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*tpuNodeState)(nil)).Elem()
}

type tpuNodeArgs struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType string `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The immutable name of the TPU.
	Name *string `pulumi:"name"`
	// The name of a network to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on
	// which this API has been activated. If none is provided, "default" will be used.
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	// Sets the scheduling options for this TPU instance.
	SchedulingConfig *TpuNodeSchedulingConfig `pulumi:"schedulingConfig"`
	// The version of Tensorflow running in the Node.
	TensorflowVersion string           `pulumi:"tensorflowVersion"`
	Timeouts          *TpuNodeTimeouts `pulumi:"timeouts"`
	TpuNodeId         *string          `pulumi:"tpuNodeId"`
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before
	// provisioning the node. If this field is set, cidr_block field should not be specified. If the network that you want to
	// peer the TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking *bool `pulumi:"useServiceNetworking"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a TpuNode resource.
type TpuNodeArgs struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringInput
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The immutable name of the TPU.
	Name pulumi.StringPtrInput
	// The name of a network to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on
	// which this API has been activated. If none is provided, "default" will be used.
	Network pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Sets the scheduling options for this TPU instance.
	SchedulingConfig TpuNodeSchedulingConfigPtrInput
	// The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringInput
	Timeouts          TpuNodeTimeoutsPtrInput
	TpuNodeId         pulumi.StringPtrInput
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before
	// provisioning the node. If this field is set, cidr_block field should not be specified. If the network that you want to
	// peer the TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking pulumi.BoolPtrInput
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (TpuNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tpuNodeArgs)(nil)).Elem()
}

type TpuNodeInput interface {
	pulumi.Input

	ToTpuNodeOutput() TpuNodeOutput
	ToTpuNodeOutputWithContext(ctx context.Context) TpuNodeOutput
}

func (*TpuNode) ElementType() reflect.Type {
	return reflect.TypeOf((**TpuNode)(nil)).Elem()
}

func (i *TpuNode) ToTpuNodeOutput() TpuNodeOutput {
	return i.ToTpuNodeOutputWithContext(context.Background())
}

func (i *TpuNode) ToTpuNodeOutputWithContext(ctx context.Context) TpuNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TpuNodeOutput)
}

type TpuNodeOutput struct{ *pulumi.OutputState }

func (TpuNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TpuNode)(nil)).Elem()
}

func (o TpuNodeOutput) ToTpuNodeOutput() TpuNodeOutput {
	return o
}

func (o TpuNodeOutput) ToTpuNodeOutputWithContext(ctx context.Context) TpuNodeOutput {
	return o
}

// The type of hardware accelerators associated with this node.
func (o TpuNodeOutput) AcceleratorType() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.AcceleratorType }).(pulumi.StringOutput)
}

// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
// that is using that CIDR block.
func (o TpuNodeOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The user-supplied description of the TPU. Maximum of 512 characters.
func (o TpuNodeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TpuNodeOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o TpuNodeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The immutable name of the TPU.
func (o TpuNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of a network to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on
// which this API has been activated. If none is provided, "default" will be used.
func (o TpuNodeOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
// node first reach out to the first (index 0) entry.
func (o TpuNodeOutput) NetworkEndpoints() TpuNodeNetworkEndpointArrayOutput {
	return o.ApplyT(func(v *TpuNode) TpuNodeNetworkEndpointArrayOutput { return v.NetworkEndpoints }).(TpuNodeNetworkEndpointArrayOutput)
}

func (o TpuNodeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Sets the scheduling options for this TPU instance.
func (o TpuNodeOutput) SchedulingConfig() TpuNodeSchedulingConfigPtrOutput {
	return o.ApplyT(func(v *TpuNode) TpuNodeSchedulingConfigPtrOutput { return v.SchedulingConfig }).(TpuNodeSchedulingConfigPtrOutput)
}

// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
func (o TpuNodeOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The version of Tensorflow running in the Node.
func (o TpuNodeOutput) TensorflowVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.TensorflowVersion }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o TpuNodeOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o TpuNodeOutput) Timeouts() TpuNodeTimeoutsPtrOutput {
	return o.ApplyT(func(v *TpuNode) TpuNodeTimeoutsPtrOutput { return v.Timeouts }).(TpuNodeTimeoutsPtrOutput)
}

func (o TpuNodeOutput) TpuNodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.TpuNodeId }).(pulumi.StringOutput)
}

// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before
// provisioning the node. If this field is set, cidr_block field should not be specified. If the network that you want to
// peer the TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
func (o TpuNodeOutput) UseServiceNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.BoolPtrOutput { return v.UseServiceNetworking }).(pulumi.BoolPtrOutput)
}

// The GCP location for the TPU. If it is not provided, the provider zone is used.
func (o TpuNodeOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuNode) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TpuNodeInput)(nil)).Elem(), &TpuNode{})
	pulumi.RegisterOutputType(TpuNodeOutput{})
}
