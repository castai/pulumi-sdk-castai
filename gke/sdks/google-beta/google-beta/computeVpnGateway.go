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

type ComputeVpnGateway struct {
	pulumi.CustomResourceState

	ComputeVpnGatewayId pulumi.StringOutput `pulumi:"computeVpnGatewayId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The unique identifier for the resource.
	GatewayId pulumi.Float64Output `pulumi:"gatewayId"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringOutput `pulumi:"network"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The region this gateway should sit in.
	Region   pulumi.StringOutput                `pulumi:"region"`
	SelfLink pulumi.StringOutput                `pulumi:"selfLink"`
	Timeouts ComputeVpnGatewayTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewComputeVpnGateway(ctx *pulumi.Context,
	name string, args *ComputeVpnGatewayArgs, opts ...pulumi.ResourceOption) (*ComputeVpnGateway, error) {
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
	var resource ComputeVpnGateway
	err = ctx.RegisterPackageResource("google-beta:index/computeVpnGateway:ComputeVpnGateway", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeVpnGateway gets an existing ComputeVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeVpnGatewayState, opts ...pulumi.ResourceOption) (*ComputeVpnGateway, error) {
	var resource ComputeVpnGateway
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeVpnGateway:ComputeVpnGateway", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeVpnGateway resources.
type computeVpnGatewayState struct {
	ComputeVpnGatewayId *string `pulumi:"computeVpnGatewayId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The unique identifier for the resource.
	GatewayId *float64 `pulumi:"gatewayId"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region   *string                    `pulumi:"region"`
	SelfLink *string                    `pulumi:"selfLink"`
	Timeouts *ComputeVpnGatewayTimeouts `pulumi:"timeouts"`
}

type ComputeVpnGatewayState struct {
	ComputeVpnGatewayId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The unique identifier for the resource.
	GatewayId pulumi.Float64PtrInput
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
	Timeouts ComputeVpnGatewayTimeoutsPtrInput
}

func (ComputeVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeVpnGatewayState)(nil)).Elem()
}

type computeVpnGatewayArgs struct {
	ComputeVpnGatewayId *string `pulumi:"computeVpnGatewayId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network string  `pulumi:"network"`
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region   *string                    `pulumi:"region"`
	Timeouts *ComputeVpnGatewayTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeVpnGateway resource.
type ComputeVpnGatewayArgs struct {
	ComputeVpnGatewayId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringInput
	Project pulumi.StringPtrInput
	// The region this gateway should sit in.
	Region   pulumi.StringPtrInput
	Timeouts ComputeVpnGatewayTimeoutsPtrInput
}

func (ComputeVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeVpnGatewayArgs)(nil)).Elem()
}

type ComputeVpnGatewayInput interface {
	pulumi.Input

	ToComputeVpnGatewayOutput() ComputeVpnGatewayOutput
	ToComputeVpnGatewayOutputWithContext(ctx context.Context) ComputeVpnGatewayOutput
}

func (*ComputeVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeVpnGateway)(nil)).Elem()
}

func (i *ComputeVpnGateway) ToComputeVpnGatewayOutput() ComputeVpnGatewayOutput {
	return i.ToComputeVpnGatewayOutputWithContext(context.Background())
}

func (i *ComputeVpnGateway) ToComputeVpnGatewayOutputWithContext(ctx context.Context) ComputeVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeVpnGatewayOutput)
}

type ComputeVpnGatewayOutput struct{ *pulumi.OutputState }

func (ComputeVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeVpnGateway)(nil)).Elem()
}

func (o ComputeVpnGatewayOutput) ToComputeVpnGatewayOutput() ComputeVpnGatewayOutput {
	return o
}

func (o ComputeVpnGatewayOutput) ToComputeVpnGatewayOutputWithContext(ctx context.Context) ComputeVpnGatewayOutput {
	return o
}

func (o ComputeVpnGatewayOutput) ComputeVpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringOutput { return v.ComputeVpnGatewayId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeVpnGatewayOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeVpnGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The unique identifier for the resource.
func (o ComputeVpnGatewayOutput) GatewayId() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.Float64Output { return v.GatewayId }).(pulumi.Float64Output)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeVpnGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network this VPN gateway is accepting traffic for.
func (o ComputeVpnGatewayOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o ComputeVpnGatewayOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region this gateway should sit in.
func (o ComputeVpnGatewayOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeVpnGatewayOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeVpnGatewayOutput) Timeouts() ComputeVpnGatewayTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeVpnGateway) ComputeVpnGatewayTimeoutsPtrOutput { return v.Timeouts }).(ComputeVpnGatewayTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeVpnGatewayInput)(nil)).Elem(), &ComputeVpnGateway{})
	pulumi.RegisterOutputType(ComputeVpnGatewayOutput{})
}
