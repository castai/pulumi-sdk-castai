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

type ComputePacketMirroring struct {
	pulumi.CustomResourceState

	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
	// traffic. The specified forwarding rule must have is_mirroring_collector set to true.
	CollectorIlb             ComputePacketMirroringCollectorIlbOutput `pulumi:"collectorIlb"`
	ComputePacketMirroringId pulumi.StringOutput                      `pulumi:"computePacketMirroringId"`
	// A human-readable description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A filter for mirrored traffic. If unset, all traffic is mirrored.
	Filter ComputePacketMirroringFilterPtrOutput `pulumi:"filter"`
	// A means of specifying which resources to mirror.
	MirroredResources ComputePacketMirroringMirroredResourcesOutput `pulumi:"mirroredResources"`
	// The name of the packet mirroring rule
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in
	// the given network. All mirrored subnetworks should belong to the given network.
	Network ComputePacketMirroringNetworkOutput `pulumi:"network"`
	// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to the
	// same instances.
	Priority pulumi.Float64Output `pulumi:"priority"`
	Project  pulumi.StringOutput  `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringOutput                     `pulumi:"region"`
	Timeouts ComputePacketMirroringTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputePacketMirroring registers a new resource with the given unique name, arguments, and options.
func NewComputePacketMirroring(ctx *pulumi.Context,
	name string, args *ComputePacketMirroringArgs, opts ...pulumi.ResourceOption) (*ComputePacketMirroring, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectorIlb == nil {
		return nil, errors.New("invalid value for required argument 'CollectorIlb'")
	}
	if args.MirroredResources == nil {
		return nil, errors.New("invalid value for required argument 'MirroredResources'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputePacketMirroring
	err = ctx.RegisterPackageResource("google-beta:index/computePacketMirroring:ComputePacketMirroring", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputePacketMirroring gets an existing ComputePacketMirroring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputePacketMirroring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputePacketMirroringState, opts ...pulumi.ResourceOption) (*ComputePacketMirroring, error) {
	var resource ComputePacketMirroring
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computePacketMirroring:ComputePacketMirroring", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputePacketMirroring resources.
type computePacketMirroringState struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
	// traffic. The specified forwarding rule must have is_mirroring_collector set to true.
	CollectorIlb             *ComputePacketMirroringCollectorIlb `pulumi:"collectorIlb"`
	ComputePacketMirroringId *string                             `pulumi:"computePacketMirroringId"`
	// A human-readable description of the rule.
	Description *string `pulumi:"description"`
	// A filter for mirrored traffic. If unset, all traffic is mirrored.
	Filter *ComputePacketMirroringFilter `pulumi:"filter"`
	// A means of specifying which resources to mirror.
	MirroredResources *ComputePacketMirroringMirroredResources `pulumi:"mirroredResources"`
	// The name of the packet mirroring rule
	Name *string `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in
	// the given network. All mirrored subnetworks should belong to the given network.
	Network *ComputePacketMirroringNetwork `pulumi:"network"`
	// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to the
	// same instances.
	Priority *float64 `pulumi:"priority"`
	Project  *string  `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   *string                         `pulumi:"region"`
	Timeouts *ComputePacketMirroringTimeouts `pulumi:"timeouts"`
}

type ComputePacketMirroringState struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
	// traffic. The specified forwarding rule must have is_mirroring_collector set to true.
	CollectorIlb             ComputePacketMirroringCollectorIlbPtrInput
	ComputePacketMirroringId pulumi.StringPtrInput
	// A human-readable description of the rule.
	Description pulumi.StringPtrInput
	// A filter for mirrored traffic. If unset, all traffic is mirrored.
	Filter ComputePacketMirroringFilterPtrInput
	// A means of specifying which resources to mirror.
	MirroredResources ComputePacketMirroringMirroredResourcesPtrInput
	// The name of the packet mirroring rule
	Name pulumi.StringPtrInput
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in
	// the given network. All mirrored subnetworks should belong to the given network.
	Network ComputePacketMirroringNetworkPtrInput
	// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to the
	// same instances.
	Priority pulumi.Float64PtrInput
	Project  pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringPtrInput
	Timeouts ComputePacketMirroringTimeoutsPtrInput
}

func (ComputePacketMirroringState) ElementType() reflect.Type {
	return reflect.TypeOf((*computePacketMirroringState)(nil)).Elem()
}

type computePacketMirroringArgs struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
	// traffic. The specified forwarding rule must have is_mirroring_collector set to true.
	CollectorIlb             ComputePacketMirroringCollectorIlb `pulumi:"collectorIlb"`
	ComputePacketMirroringId *string                            `pulumi:"computePacketMirroringId"`
	// A human-readable description of the rule.
	Description *string `pulumi:"description"`
	// A filter for mirrored traffic. If unset, all traffic is mirrored.
	Filter *ComputePacketMirroringFilter `pulumi:"filter"`
	// A means of specifying which resources to mirror.
	MirroredResources ComputePacketMirroringMirroredResources `pulumi:"mirroredResources"`
	// The name of the packet mirroring rule
	Name *string `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in
	// the given network. All mirrored subnetworks should belong to the given network.
	Network ComputePacketMirroringNetwork `pulumi:"network"`
	// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to the
	// same instances.
	Priority *float64 `pulumi:"priority"`
	Project  *string  `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   *string                         `pulumi:"region"`
	Timeouts *ComputePacketMirroringTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputePacketMirroring resource.
type ComputePacketMirroringArgs struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
	// traffic. The specified forwarding rule must have is_mirroring_collector set to true.
	CollectorIlb             ComputePacketMirroringCollectorIlbInput
	ComputePacketMirroringId pulumi.StringPtrInput
	// A human-readable description of the rule.
	Description pulumi.StringPtrInput
	// A filter for mirrored traffic. If unset, all traffic is mirrored.
	Filter ComputePacketMirroringFilterPtrInput
	// A means of specifying which resources to mirror.
	MirroredResources ComputePacketMirroringMirroredResourcesInput
	// The name of the packet mirroring rule
	Name pulumi.StringPtrInput
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in
	// the given network. All mirrored subnetworks should belong to the given network.
	Network ComputePacketMirroringNetworkInput
	// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to the
	// same instances.
	Priority pulumi.Float64PtrInput
	Project  pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringPtrInput
	Timeouts ComputePacketMirroringTimeoutsPtrInput
}

func (ComputePacketMirroringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computePacketMirroringArgs)(nil)).Elem()
}

type ComputePacketMirroringInput interface {
	pulumi.Input

	ToComputePacketMirroringOutput() ComputePacketMirroringOutput
	ToComputePacketMirroringOutputWithContext(ctx context.Context) ComputePacketMirroringOutput
}

func (*ComputePacketMirroring) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputePacketMirroring)(nil)).Elem()
}

func (i *ComputePacketMirroring) ToComputePacketMirroringOutput() ComputePacketMirroringOutput {
	return i.ToComputePacketMirroringOutputWithContext(context.Background())
}

func (i *ComputePacketMirroring) ToComputePacketMirroringOutputWithContext(ctx context.Context) ComputePacketMirroringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputePacketMirroringOutput)
}

type ComputePacketMirroringOutput struct{ *pulumi.OutputState }

func (ComputePacketMirroringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputePacketMirroring)(nil)).Elem()
}

func (o ComputePacketMirroringOutput) ToComputePacketMirroringOutput() ComputePacketMirroringOutput {
	return o
}

func (o ComputePacketMirroringOutput) ToComputePacketMirroringOutputWithContext(ctx context.Context) ComputePacketMirroringOutput {
	return o
}

// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
// traffic. The specified forwarding rule must have is_mirroring_collector set to true.
func (o ComputePacketMirroringOutput) CollectorIlb() ComputePacketMirroringCollectorIlbOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) ComputePacketMirroringCollectorIlbOutput { return v.CollectorIlb }).(ComputePacketMirroringCollectorIlbOutput)
}

func (o ComputePacketMirroringOutput) ComputePacketMirroringId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) pulumi.StringOutput { return v.ComputePacketMirroringId }).(pulumi.StringOutput)
}

// A human-readable description of the rule.
func (o ComputePacketMirroringOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A filter for mirrored traffic. If unset, all traffic is mirrored.
func (o ComputePacketMirroringOutput) Filter() ComputePacketMirroringFilterPtrOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) ComputePacketMirroringFilterPtrOutput { return v.Filter }).(ComputePacketMirroringFilterPtrOutput)
}

// A means of specifying which resources to mirror.
func (o ComputePacketMirroringOutput) MirroredResources() ComputePacketMirroringMirroredResourcesOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) ComputePacketMirroringMirroredResourcesOutput {
		return v.MirroredResources
	}).(ComputePacketMirroringMirroredResourcesOutput)
}

// The name of the packet mirroring rule
func (o ComputePacketMirroringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in
// the given network. All mirrored subnetworks should belong to the given network.
func (o ComputePacketMirroringOutput) Network() ComputePacketMirroringNetworkOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) ComputePacketMirroringNetworkOutput { return v.Network }).(ComputePacketMirroringNetworkOutput)
}

// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to the
// same instances.
func (o ComputePacketMirroringOutput) Priority() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputePacketMirroring) pulumi.Float64Output { return v.Priority }).(pulumi.Float64Output)
}

func (o ComputePacketMirroringOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The Region in which the created address should reside. If it is not provided, the provider region is used.
func (o ComputePacketMirroringOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputePacketMirroringOutput) Timeouts() ComputePacketMirroringTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputePacketMirroring) ComputePacketMirroringTimeoutsPtrOutput { return v.Timeouts }).(ComputePacketMirroringTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputePacketMirroringInput)(nil)).Elem(), &ComputePacketMirroring{})
	pulumi.RegisterOutputType(ComputePacketMirroringOutput{})
}
