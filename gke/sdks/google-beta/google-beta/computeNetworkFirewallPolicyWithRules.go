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

type ComputeNetworkFirewallPolicyWithRules struct {
	pulumi.CustomResourceState

	ComputeNetworkFirewallPolicyWithRulesId pulumi.StringOutput `pulumi:"computeNetworkFirewallPolicyWithRulesId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
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
	// A list of firewall policy pre-defined rules.
	PredefinedRules ComputeNetworkFirewallPolicyWithRulesPredefinedRuleArrayOutput `pulumi:"predefinedRules"`
	Project         pulumi.StringOutput                                            `pulumi:"project"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.Float64Output `pulumi:"ruleTupleCount"`
	// A list of firewall policy rules.
	Rules ComputeNetworkFirewallPolicyWithRulesRuleArrayOutput `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput                                    `pulumi:"selfLinkWithId"`
	Timeouts       ComputeNetworkFirewallPolicyWithRulesTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeNetworkFirewallPolicyWithRules registers a new resource with the given unique name, arguments, and options.
func NewComputeNetworkFirewallPolicyWithRules(ctx *pulumi.Context,
	name string, args *ComputeNetworkFirewallPolicyWithRulesArgs, opts ...pulumi.ResourceOption) (*ComputeNetworkFirewallPolicyWithRules, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeNetworkFirewallPolicyWithRules
	err = ctx.RegisterPackageResource("google-beta:index/computeNetworkFirewallPolicyWithRules:ComputeNetworkFirewallPolicyWithRules", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeNetworkFirewallPolicyWithRules gets an existing ComputeNetworkFirewallPolicyWithRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeNetworkFirewallPolicyWithRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeNetworkFirewallPolicyWithRulesState, opts ...pulumi.ResourceOption) (*ComputeNetworkFirewallPolicyWithRules, error) {
	var resource ComputeNetworkFirewallPolicyWithRules
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeNetworkFirewallPolicyWithRules:ComputeNetworkFirewallPolicyWithRules", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeNetworkFirewallPolicyWithRules resources.
type computeNetworkFirewallPolicyWithRulesState struct {
	ComputeNetworkFirewallPolicyWithRulesId *string `pulumi:"computeNetworkFirewallPolicyWithRulesId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
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
	// A list of firewall policy pre-defined rules.
	PredefinedRules []ComputeNetworkFirewallPolicyWithRulesPredefinedRule `pulumi:"predefinedRules"`
	Project         *string                                               `pulumi:"project"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount *float64 `pulumi:"ruleTupleCount"`
	// A list of firewall policy rules.
	Rules []ComputeNetworkFirewallPolicyWithRulesRule `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string                                        `pulumi:"selfLinkWithId"`
	Timeouts       *ComputeNetworkFirewallPolicyWithRulesTimeouts `pulumi:"timeouts"`
}

type ComputeNetworkFirewallPolicyWithRulesState struct {
	ComputeNetworkFirewallPolicyWithRulesId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
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
	// A list of firewall policy pre-defined rules.
	PredefinedRules ComputeNetworkFirewallPolicyWithRulesPredefinedRuleArrayInput
	Project         pulumi.StringPtrInput
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.Float64PtrInput
	// A list of firewall policy rules.
	Rules ComputeNetworkFirewallPolicyWithRulesRuleArrayInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	Timeouts       ComputeNetworkFirewallPolicyWithRulesTimeoutsPtrInput
}

func (ComputeNetworkFirewallPolicyWithRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkFirewallPolicyWithRulesState)(nil)).Elem()
}

type computeNetworkFirewallPolicyWithRulesArgs struct {
	ComputeNetworkFirewallPolicyWithRulesId *string `pulumi:"computeNetworkFirewallPolicyWithRulesId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
	// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
	// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A list of firewall policy rules.
	Rules    []ComputeNetworkFirewallPolicyWithRulesRule    `pulumi:"rules"`
	Timeouts *ComputeNetworkFirewallPolicyWithRulesTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeNetworkFirewallPolicyWithRules resource.
type ComputeNetworkFirewallPolicyWithRulesArgs struct {
	ComputeNetworkFirewallPolicyWithRulesId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
	// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
	// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A list of firewall policy rules.
	Rules    ComputeNetworkFirewallPolicyWithRulesRuleArrayInput
	Timeouts ComputeNetworkFirewallPolicyWithRulesTimeoutsPtrInput
}

func (ComputeNetworkFirewallPolicyWithRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkFirewallPolicyWithRulesArgs)(nil)).Elem()
}

type ComputeNetworkFirewallPolicyWithRulesInput interface {
	pulumi.Input

	ToComputeNetworkFirewallPolicyWithRulesOutput() ComputeNetworkFirewallPolicyWithRulesOutput
	ToComputeNetworkFirewallPolicyWithRulesOutputWithContext(ctx context.Context) ComputeNetworkFirewallPolicyWithRulesOutput
}

func (*ComputeNetworkFirewallPolicyWithRules) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkFirewallPolicyWithRules)(nil)).Elem()
}

func (i *ComputeNetworkFirewallPolicyWithRules) ToComputeNetworkFirewallPolicyWithRulesOutput() ComputeNetworkFirewallPolicyWithRulesOutput {
	return i.ToComputeNetworkFirewallPolicyWithRulesOutputWithContext(context.Background())
}

func (i *ComputeNetworkFirewallPolicyWithRules) ToComputeNetworkFirewallPolicyWithRulesOutputWithContext(ctx context.Context) ComputeNetworkFirewallPolicyWithRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeNetworkFirewallPolicyWithRulesOutput)
}

type ComputeNetworkFirewallPolicyWithRulesOutput struct{ *pulumi.OutputState }

func (ComputeNetworkFirewallPolicyWithRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkFirewallPolicyWithRules)(nil)).Elem()
}

func (o ComputeNetworkFirewallPolicyWithRulesOutput) ToComputeNetworkFirewallPolicyWithRulesOutput() ComputeNetworkFirewallPolicyWithRulesOutput {
	return o
}

func (o ComputeNetworkFirewallPolicyWithRulesOutput) ToComputeNetworkFirewallPolicyWithRulesOutputWithContext(ctx context.Context) ComputeNetworkFirewallPolicyWithRulesOutput {
	return o
}

func (o ComputeNetworkFirewallPolicyWithRulesOutput) ComputeNetworkFirewallPolicyWithRulesId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput {
		return v.ComputeNetworkFirewallPolicyWithRulesId
	}).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Fingerprint of the resource. This field is used internally during updates of this resource.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy
// is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63
// characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and
// all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique identifier for the resource. This identifier is defined by the server.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) NetworkFirewallPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput { return v.NetworkFirewallPolicyId }).(pulumi.StringOutput)
}

// A list of firewall policy pre-defined rules.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) PredefinedRules() ComputeNetworkFirewallPolicyWithRulesPredefinedRuleArrayOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) ComputeNetworkFirewallPolicyWithRulesPredefinedRuleArrayOutput {
		return v.PredefinedRules
	}).(ComputeNetworkFirewallPolicyWithRulesPredefinedRuleArrayOutput)
}

func (o ComputeNetworkFirewallPolicyWithRulesOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) RuleTupleCount() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.Float64Output { return v.RuleTupleCount }).(pulumi.Float64Output)
}

// A list of firewall policy rules.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) Rules() ComputeNetworkFirewallPolicyWithRulesRuleArrayOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) ComputeNetworkFirewallPolicyWithRulesRuleArrayOutput {
		return v.Rules
	}).(ComputeNetworkFirewallPolicyWithRulesRuleArrayOutput)
}

// Server-defined URL for the resource.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o ComputeNetworkFirewallPolicyWithRulesOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

func (o ComputeNetworkFirewallPolicyWithRulesOutput) Timeouts() ComputeNetworkFirewallPolicyWithRulesTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeNetworkFirewallPolicyWithRules) ComputeNetworkFirewallPolicyWithRulesTimeoutsPtrOutput {
		return v.Timeouts
	}).(ComputeNetworkFirewallPolicyWithRulesTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeNetworkFirewallPolicyWithRulesInput)(nil)).Elem(), &ComputeNetworkFirewallPolicyWithRules{})
	pulumi.RegisterOutputType(ComputeNetworkFirewallPolicyWithRulesOutput{})
}
