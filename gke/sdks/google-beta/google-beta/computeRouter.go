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

type ComputeRouter struct {
	pulumi.CustomResourceState

	// BGP information specific to this router.
	Bgp             ComputeRouterBgpPtrOutput `pulumi:"bgp"`
	ComputeRouterId pulumi.StringOutput       `pulumi:"computeRouterId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolPtrOutput `pulumi:"encryptedInterconnectRouter"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// A reference to the network to which this router belongs.
	Network pulumi.StringOutput `pulumi:"network"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the router resides.
	Region   pulumi.StringOutput            `pulumi:"region"`
	SelfLink pulumi.StringOutput            `pulumi:"selfLink"`
	Timeouts ComputeRouterTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeRouter registers a new resource with the given unique name, arguments, and options.
func NewComputeRouter(ctx *pulumi.Context,
	name string, args *ComputeRouterArgs, opts ...pulumi.ResourceOption) (*ComputeRouter, error) {
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
	var resource ComputeRouter
	err = ctx.RegisterPackageResource("google-beta:index/computeRouter:ComputeRouter", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRouter gets an existing ComputeRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRouterState, opts ...pulumi.ResourceOption) (*ComputeRouter, error) {
	var resource ComputeRouter
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeRouter:ComputeRouter", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRouter resources.
type computeRouterState struct {
	// BGP information specific to this router.
	Bgp             *ComputeRouterBgp `pulumi:"bgp"`
	ComputeRouterId *string           `pulumi:"computeRouterId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter *bool `pulumi:"encryptedInterconnectRouter"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name *string `pulumi:"name"`
	// A reference to the network to which this router belongs.
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	// Region where the router resides.
	Region   *string                `pulumi:"region"`
	SelfLink *string                `pulumi:"selfLink"`
	Timeouts *ComputeRouterTimeouts `pulumi:"timeouts"`
}

type ComputeRouterState struct {
	// BGP information specific to this router.
	Bgp             ComputeRouterBgpPtrInput
	ComputeRouterId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolPtrInput
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringPtrInput
	// A reference to the network to which this router belongs.
	Network pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Region where the router resides.
	Region   pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	Timeouts ComputeRouterTimeoutsPtrInput
}

func (ComputeRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRouterState)(nil)).Elem()
}

type computeRouterArgs struct {
	// BGP information specific to this router.
	Bgp             *ComputeRouterBgp `pulumi:"bgp"`
	ComputeRouterId *string           `pulumi:"computeRouterId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter *bool `pulumi:"encryptedInterconnectRouter"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name *string `pulumi:"name"`
	// A reference to the network to which this router belongs.
	Network string  `pulumi:"network"`
	Project *string `pulumi:"project"`
	// Region where the router resides.
	Region   *string                `pulumi:"region"`
	Timeouts *ComputeRouterTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeRouter resource.
type ComputeRouterArgs struct {
	// BGP information specific to this router.
	Bgp             ComputeRouterBgpPtrInput
	ComputeRouterId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolPtrInput
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
	// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
	// dash.
	Name pulumi.StringPtrInput
	// A reference to the network to which this router belongs.
	Network pulumi.StringInput
	Project pulumi.StringPtrInput
	// Region where the router resides.
	Region   pulumi.StringPtrInput
	Timeouts ComputeRouterTimeoutsPtrInput
}

func (ComputeRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRouterArgs)(nil)).Elem()
}

type ComputeRouterInput interface {
	pulumi.Input

	ToComputeRouterOutput() ComputeRouterOutput
	ToComputeRouterOutputWithContext(ctx context.Context) ComputeRouterOutput
}

func (*ComputeRouter) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRouter)(nil)).Elem()
}

func (i *ComputeRouter) ToComputeRouterOutput() ComputeRouterOutput {
	return i.ToComputeRouterOutputWithContext(context.Background())
}

func (i *ComputeRouter) ToComputeRouterOutputWithContext(ctx context.Context) ComputeRouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRouterOutput)
}

type ComputeRouterOutput struct{ *pulumi.OutputState }

func (ComputeRouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRouter)(nil)).Elem()
}

func (o ComputeRouterOutput) ToComputeRouterOutput() ComputeRouterOutput {
	return o
}

func (o ComputeRouterOutput) ToComputeRouterOutputWithContext(ctx context.Context) ComputeRouterOutput {
	return o
}

// BGP information specific to this router.
func (o ComputeRouterOutput) Bgp() ComputeRouterBgpPtrOutput {
	return o.ApplyT(func(v *ComputeRouter) ComputeRouterBgpPtrOutput { return v.Bgp }).(ComputeRouterBgpPtrOutput)
}

func (o ComputeRouterOutput) ComputeRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringOutput { return v.ComputeRouterId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeRouterOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeRouterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
func (o ComputeRouterOutput) EncryptedInterconnectRouter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.BoolPtrOutput { return v.EncryptedInterconnectRouter }).(pulumi.BoolPtrOutput)
}

// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be
// 1-63 characters long and match the regular expression 'a-z?' which means the first character must be a lowercase letter,
// and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a
// dash.
func (o ComputeRouterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A reference to the network to which this router belongs.
func (o ComputeRouterOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o ComputeRouterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Region where the router resides.
func (o ComputeRouterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeRouterOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouter) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeRouterOutput) Timeouts() ComputeRouterTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeRouter) ComputeRouterTimeoutsPtrOutput { return v.Timeouts }).(ComputeRouterTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRouterInput)(nil)).Elem(), &ComputeRouter{})
	pulumi.RegisterOutputType(ComputeRouterOutput{})
}
