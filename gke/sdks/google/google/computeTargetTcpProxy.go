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

type ComputeTargetTcpProxy struct {
	pulumi.CustomResourceState

	// A reference to the BackendService resource.
	BackendService          pulumi.StringOutput `pulumi:"backendService"`
	ComputeTargetTcpProxyId pulumi.StringOutput `pulumi:"computeTargetTcpProxyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader pulumi.StringPtrOutput `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId  pulumi.Float64Output                   `pulumi:"proxyId"`
	SelfLink pulumi.StringOutput                    `pulumi:"selfLink"`
	Timeouts ComputeTargetTcpProxyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeTargetTcpProxy registers a new resource with the given unique name, arguments, and options.
func NewComputeTargetTcpProxy(ctx *pulumi.Context,
	name string, args *ComputeTargetTcpProxyArgs, opts ...pulumi.ResourceOption) (*ComputeTargetTcpProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendService == nil {
		return nil, errors.New("invalid value for required argument 'BackendService'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeTargetTcpProxy
	err = ctx.RegisterPackageResource("google:index/computeTargetTcpProxy:ComputeTargetTcpProxy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeTargetTcpProxy gets an existing ComputeTargetTcpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeTargetTcpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeTargetTcpProxyState, opts ...pulumi.ResourceOption) (*ComputeTargetTcpProxy, error) {
	var resource ComputeTargetTcpProxy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeTargetTcpProxy:ComputeTargetTcpProxy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeTargetTcpProxy resources.
type computeTargetTcpProxyState struct {
	// A reference to the BackendService resource.
	BackendService          *string `pulumi:"backendService"`
	ComputeTargetTcpProxyId *string `pulumi:"computeTargetTcpProxyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader *string `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId  *float64                       `pulumi:"proxyId"`
	SelfLink *string                        `pulumi:"selfLink"`
	Timeouts *ComputeTargetTcpProxyTimeouts `pulumi:"timeouts"`
}

type ComputeTargetTcpProxyState struct {
	// A reference to the BackendService resource.
	BackendService          pulumi.StringPtrInput
	ComputeTargetTcpProxyId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader pulumi.StringPtrInput
	// The unique identifier for the resource.
	ProxyId  pulumi.Float64PtrInput
	SelfLink pulumi.StringPtrInput
	Timeouts ComputeTargetTcpProxyTimeoutsPtrInput
}

func (ComputeTargetTcpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeTargetTcpProxyState)(nil)).Elem()
}

type computeTargetTcpProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService          string  `pulumi:"backendService"`
	ComputeTargetTcpProxyId *string `pulumi:"computeTargetTcpProxyId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader *string                        `pulumi:"proxyHeader"`
	Timeouts    *ComputeTargetTcpProxyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeTargetTcpProxy resource.
type ComputeTargetTcpProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService          pulumi.StringInput
	ComputeTargetTcpProxyId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader pulumi.StringPtrInput
	Timeouts    ComputeTargetTcpProxyTimeoutsPtrInput
}

func (ComputeTargetTcpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeTargetTcpProxyArgs)(nil)).Elem()
}

type ComputeTargetTcpProxyInput interface {
	pulumi.Input

	ToComputeTargetTcpProxyOutput() ComputeTargetTcpProxyOutput
	ToComputeTargetTcpProxyOutputWithContext(ctx context.Context) ComputeTargetTcpProxyOutput
}

func (*ComputeTargetTcpProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeTargetTcpProxy)(nil)).Elem()
}

func (i *ComputeTargetTcpProxy) ToComputeTargetTcpProxyOutput() ComputeTargetTcpProxyOutput {
	return i.ToComputeTargetTcpProxyOutputWithContext(context.Background())
}

func (i *ComputeTargetTcpProxy) ToComputeTargetTcpProxyOutputWithContext(ctx context.Context) ComputeTargetTcpProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeTargetTcpProxyOutput)
}

type ComputeTargetTcpProxyOutput struct{ *pulumi.OutputState }

func (ComputeTargetTcpProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeTargetTcpProxy)(nil)).Elem()
}

func (o ComputeTargetTcpProxyOutput) ToComputeTargetTcpProxyOutput() ComputeTargetTcpProxyOutput {
	return o
}

func (o ComputeTargetTcpProxyOutput) ToComputeTargetTcpProxyOutputWithContext(ctx context.Context) ComputeTargetTcpProxyOutput {
	return o
}

// A reference to the BackendService resource.
func (o ComputeTargetTcpProxyOutput) BackendService() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringOutput { return v.BackendService }).(pulumi.StringOutput)
}

func (o ComputeTargetTcpProxyOutput) ComputeTargetTcpProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringOutput { return v.ComputeTargetTcpProxyId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeTargetTcpProxyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeTargetTcpProxyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeTargetTcpProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeTargetTcpProxyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
// INTERNAL_SELF_MANAGED.
func (o ComputeTargetTcpProxyOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.BoolOutput { return v.ProxyBind }).(pulumi.BoolOutput)
}

// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
// ["NONE", "PROXY_V1"]
func (o ComputeTargetTcpProxyOutput) ProxyHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringPtrOutput { return v.ProxyHeader }).(pulumi.StringPtrOutput)
}

// The unique identifier for the resource.
func (o ComputeTargetTcpProxyOutput) ProxyId() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.Float64Output { return v.ProxyId }).(pulumi.Float64Output)
}

func (o ComputeTargetTcpProxyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeTargetTcpProxyOutput) Timeouts() ComputeTargetTcpProxyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeTargetTcpProxy) ComputeTargetTcpProxyTimeoutsPtrOutput { return v.Timeouts }).(ComputeTargetTcpProxyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeTargetTcpProxyInput)(nil)).Elem(), &ComputeTargetTcpProxy{})
	pulumi.RegisterOutputType(ComputeTargetTcpProxyOutput{})
}
