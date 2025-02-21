// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeNetworkFirewallPolicy struct {
	pulumi.CustomResourceState

	ComputeNetworkFirewallPolicyId pulumi.StringOutput `pulumi:"computeNetworkFirewallPolicyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
	// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
	// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique identifier for the resource. This identifier is defined by the server.
	NetworkFirewallPolicyId pulumi.StringOutput `pulumi:"networkFirewallPolicyId"`
	Project                 pulumi.StringOutput `pulumi:"project"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.Float64Output `pulumi:"ruleTupleCount"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput                           `pulumi:"selfLinkWithId"`
	Timeouts       ComputeNetworkFirewallPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeNetworkFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewComputeNetworkFirewallPolicy(ctx *pulumi.Context,
	name string, args *ComputeNetworkFirewallPolicyArgs, opts ...pulumi.ResourceOption) (*ComputeNetworkFirewallPolicy, error) {
	if args == nil {
		args = &ComputeNetworkFirewallPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeNetworkFirewallPolicy
	err = ctx.RegisterPackageResource("google:index/computeNetworkFirewallPolicy:ComputeNetworkFirewallPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeNetworkFirewallPolicy gets an existing ComputeNetworkFirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeNetworkFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeNetworkFirewallPolicyState, opts ...pulumi.ResourceOption) (*ComputeNetworkFirewallPolicy, error) {
	var resource ComputeNetworkFirewallPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeNetworkFirewallPolicy:ComputeNetworkFirewallPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeNetworkFirewallPolicy resources.
type computeNetworkFirewallPolicyState struct {
	ComputeNetworkFirewallPolicyId *string `pulumi:"computeNetworkFirewallPolicyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint *string `pulumi:"fingerprint"`
	// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
	// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
	// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The unique identifier for the resource. This identifier is defined by the server.
	NetworkFirewallPolicyId *string `pulumi:"networkFirewallPolicyId"`
	Project                 *string `pulumi:"project"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount *float64 `pulumi:"ruleTupleCount"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string                               `pulumi:"selfLinkWithId"`
	Timeouts       *ComputeNetworkFirewallPolicyTimeouts `pulumi:"timeouts"`
}

type ComputeNetworkFirewallPolicyState struct {
	ComputeNetworkFirewallPolicyId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringPtrInput
	// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
	// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
	// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The unique identifier for the resource. This identifier is defined by the server.
	NetworkFirewallPolicyId pulumi.StringPtrInput
	Project                 pulumi.StringPtrInput
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.Float64PtrInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	Timeouts       ComputeNetworkFirewallPolicyTimeoutsPtrInput
}

func (ComputeNetworkFirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkFirewallPolicyState)(nil)).Elem()
}

type computeNetworkFirewallPolicyArgs struct {
	ComputeNetworkFirewallPolicyId *string `pulumi:"computeNetworkFirewallPolicyId"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
	// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
	// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name     *string                               `pulumi:"name"`
	Project  *string                               `pulumi:"project"`
	Timeouts *ComputeNetworkFirewallPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeNetworkFirewallPolicy resource.
type ComputeNetworkFirewallPolicyArgs struct {
	ComputeNetworkFirewallPolicyId pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
	// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
	// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeNetworkFirewallPolicyTimeoutsPtrInput
}

func (ComputeNetworkFirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkFirewallPolicyArgs)(nil)).Elem()
}

type ComputeNetworkFirewallPolicyInput interface {
	pulumi.Input

	ToComputeNetworkFirewallPolicyOutput() ComputeNetworkFirewallPolicyOutput
	ToComputeNetworkFirewallPolicyOutputWithContext(ctx context.Context) ComputeNetworkFirewallPolicyOutput
}

func (*ComputeNetworkFirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkFirewallPolicy)(nil)).Elem()
}

func (i *ComputeNetworkFirewallPolicy) ToComputeNetworkFirewallPolicyOutput() ComputeNetworkFirewallPolicyOutput {
	return i.ToComputeNetworkFirewallPolicyOutputWithContext(context.Background())
}

func (i *ComputeNetworkFirewallPolicy) ToComputeNetworkFirewallPolicyOutputWithContext(ctx context.Context) ComputeNetworkFirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeNetworkFirewallPolicyOutput)
}

type ComputeNetworkFirewallPolicyOutput struct{ *pulumi.OutputState }

func (ComputeNetworkFirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkFirewallPolicy)(nil)).Elem()
}

func (o ComputeNetworkFirewallPolicyOutput) ToComputeNetworkFirewallPolicyOutput() ComputeNetworkFirewallPolicyOutput {
	return o
}

func (o ComputeNetworkFirewallPolicyOutput) ToComputeNetworkFirewallPolicyOutputWithContext(ctx context.Context) ComputeNetworkFirewallPolicyOutput {
	return o
}

func (o ComputeNetworkFirewallPolicyOutput) ComputeNetworkFirewallPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.ComputeNetworkFirewallPolicyId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeNetworkFirewallPolicyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o ComputeNetworkFirewallPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Fingerprint of the resource. This field is used internally during updates of this resource.
func (o ComputeNetworkFirewallPolicyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o ComputeNetworkFirewallPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique identifier for the resource. This identifier is defined by the server.
func (o ComputeNetworkFirewallPolicyOutput) NetworkFirewallPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.NetworkFirewallPolicyId }).(pulumi.StringOutput)
}

func (o ComputeNetworkFirewallPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
func (o ComputeNetworkFirewallPolicyOutput) RuleTupleCount() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.Float64Output { return v.RuleTupleCount }).(pulumi.Float64Output)
}

// Server-defined URL for the resource.
func (o ComputeNetworkFirewallPolicyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o ComputeNetworkFirewallPolicyOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

func (o ComputeNetworkFirewallPolicyOutput) Timeouts() ComputeNetworkFirewallPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicy) ComputeNetworkFirewallPolicyTimeoutsPtrOutput { return v.Timeouts }).(ComputeNetworkFirewallPolicyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeNetworkFirewallPolicyInput)(nil)).Elem(), &ComputeNetworkFirewallPolicy{})
	pulumi.RegisterOutputType(ComputeNetworkFirewallPolicyOutput{})
}
