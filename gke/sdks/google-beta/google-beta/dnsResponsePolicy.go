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

type DnsResponsePolicy struct {
	pulumi.CustomResourceState

	// The description of the response policy, such as 'My new response policy'.
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	DnsResponsePolicyId pulumi.StringOutput    `pulumi:"dnsResponsePolicyId"`
	// The list of Google Kubernetes Engine clusters that can see this zone.
	GkeClusters DnsResponsePolicyGkeClusterArrayOutput `pulumi:"gkeClusters"`
	// The list of network names specifying networks to which this policy is applied.
	Networks DnsResponsePolicyNetworkArrayOutput `pulumi:"networks"`
	Project  pulumi.StringOutput                 `pulumi:"project"`
	// The user assigned name for this Response Policy, such as 'myresponsepolicy'.
	ResponsePolicyName pulumi.StringOutput                `pulumi:"responsePolicyName"`
	Timeouts           DnsResponsePolicyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDnsResponsePolicy registers a new resource with the given unique name, arguments, and options.
func NewDnsResponsePolicy(ctx *pulumi.Context,
	name string, args *DnsResponsePolicyArgs, opts ...pulumi.ResourceOption) (*DnsResponsePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResponsePolicyName == nil {
		return nil, errors.New("invalid value for required argument 'ResponsePolicyName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DnsResponsePolicy
	err = ctx.RegisterPackageResource("google-beta:index/dnsResponsePolicy:DnsResponsePolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsResponsePolicy gets an existing DnsResponsePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsResponsePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsResponsePolicyState, opts ...pulumi.ResourceOption) (*DnsResponsePolicy, error) {
	var resource DnsResponsePolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dnsResponsePolicy:DnsResponsePolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsResponsePolicy resources.
type dnsResponsePolicyState struct {
	// The description of the response policy, such as 'My new response policy'.
	Description         *string `pulumi:"description"`
	DnsResponsePolicyId *string `pulumi:"dnsResponsePolicyId"`
	// The list of Google Kubernetes Engine clusters that can see this zone.
	GkeClusters []DnsResponsePolicyGkeCluster `pulumi:"gkeClusters"`
	// The list of network names specifying networks to which this policy is applied.
	Networks []DnsResponsePolicyNetwork `pulumi:"networks"`
	Project  *string                    `pulumi:"project"`
	// The user assigned name for this Response Policy, such as 'myresponsepolicy'.
	ResponsePolicyName *string                    `pulumi:"responsePolicyName"`
	Timeouts           *DnsResponsePolicyTimeouts `pulumi:"timeouts"`
}

type DnsResponsePolicyState struct {
	// The description of the response policy, such as 'My new response policy'.
	Description         pulumi.StringPtrInput
	DnsResponsePolicyId pulumi.StringPtrInput
	// The list of Google Kubernetes Engine clusters that can see this zone.
	GkeClusters DnsResponsePolicyGkeClusterArrayInput
	// The list of network names specifying networks to which this policy is applied.
	Networks DnsResponsePolicyNetworkArrayInput
	Project  pulumi.StringPtrInput
	// The user assigned name for this Response Policy, such as 'myresponsepolicy'.
	ResponsePolicyName pulumi.StringPtrInput
	Timeouts           DnsResponsePolicyTimeoutsPtrInput
}

func (DnsResponsePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsResponsePolicyState)(nil)).Elem()
}

type dnsResponsePolicyArgs struct {
	// The description of the response policy, such as 'My new response policy'.
	Description         *string `pulumi:"description"`
	DnsResponsePolicyId *string `pulumi:"dnsResponsePolicyId"`
	// The list of Google Kubernetes Engine clusters that can see this zone.
	GkeClusters []DnsResponsePolicyGkeCluster `pulumi:"gkeClusters"`
	// The list of network names specifying networks to which this policy is applied.
	Networks []DnsResponsePolicyNetwork `pulumi:"networks"`
	Project  *string                    `pulumi:"project"`
	// The user assigned name for this Response Policy, such as 'myresponsepolicy'.
	ResponsePolicyName string                     `pulumi:"responsePolicyName"`
	Timeouts           *DnsResponsePolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DnsResponsePolicy resource.
type DnsResponsePolicyArgs struct {
	// The description of the response policy, such as 'My new response policy'.
	Description         pulumi.StringPtrInput
	DnsResponsePolicyId pulumi.StringPtrInput
	// The list of Google Kubernetes Engine clusters that can see this zone.
	GkeClusters DnsResponsePolicyGkeClusterArrayInput
	// The list of network names specifying networks to which this policy is applied.
	Networks DnsResponsePolicyNetworkArrayInput
	Project  pulumi.StringPtrInput
	// The user assigned name for this Response Policy, such as 'myresponsepolicy'.
	ResponsePolicyName pulumi.StringInput
	Timeouts           DnsResponsePolicyTimeoutsPtrInput
}

func (DnsResponsePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsResponsePolicyArgs)(nil)).Elem()
}

type DnsResponsePolicyInput interface {
	pulumi.Input

	ToDnsResponsePolicyOutput() DnsResponsePolicyOutput
	ToDnsResponsePolicyOutputWithContext(ctx context.Context) DnsResponsePolicyOutput
}

func (*DnsResponsePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsResponsePolicy)(nil)).Elem()
}

func (i *DnsResponsePolicy) ToDnsResponsePolicyOutput() DnsResponsePolicyOutput {
	return i.ToDnsResponsePolicyOutputWithContext(context.Background())
}

func (i *DnsResponsePolicy) ToDnsResponsePolicyOutputWithContext(ctx context.Context) DnsResponsePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsResponsePolicyOutput)
}

type DnsResponsePolicyOutput struct{ *pulumi.OutputState }

func (DnsResponsePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsResponsePolicy)(nil)).Elem()
}

func (o DnsResponsePolicyOutput) ToDnsResponsePolicyOutput() DnsResponsePolicyOutput {
	return o
}

func (o DnsResponsePolicyOutput) ToDnsResponsePolicyOutputWithContext(ctx context.Context) DnsResponsePolicyOutput {
	return o
}

// The description of the response policy, such as 'My new response policy'.
func (o DnsResponsePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsResponsePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DnsResponsePolicyOutput) DnsResponsePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResponsePolicy) pulumi.StringOutput { return v.DnsResponsePolicyId }).(pulumi.StringOutput)
}

// The list of Google Kubernetes Engine clusters that can see this zone.
func (o DnsResponsePolicyOutput) GkeClusters() DnsResponsePolicyGkeClusterArrayOutput {
	return o.ApplyT(func(v *DnsResponsePolicy) DnsResponsePolicyGkeClusterArrayOutput { return v.GkeClusters }).(DnsResponsePolicyGkeClusterArrayOutput)
}

// The list of network names specifying networks to which this policy is applied.
func (o DnsResponsePolicyOutput) Networks() DnsResponsePolicyNetworkArrayOutput {
	return o.ApplyT(func(v *DnsResponsePolicy) DnsResponsePolicyNetworkArrayOutput { return v.Networks }).(DnsResponsePolicyNetworkArrayOutput)
}

func (o DnsResponsePolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResponsePolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The user assigned name for this Response Policy, such as 'myresponsepolicy'.
func (o DnsResponsePolicyOutput) ResponsePolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResponsePolicy) pulumi.StringOutput { return v.ResponsePolicyName }).(pulumi.StringOutput)
}

func (o DnsResponsePolicyOutput) Timeouts() DnsResponsePolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *DnsResponsePolicy) DnsResponsePolicyTimeoutsPtrOutput { return v.Timeouts }).(DnsResponsePolicyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsResponsePolicyInput)(nil)).Elem(), &DnsResponsePolicy{})
	pulumi.RegisterOutputType(DnsResponsePolicyOutput{})
}
