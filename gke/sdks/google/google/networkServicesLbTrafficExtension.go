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

type NetworkServicesLbTrafficExtension struct {
	pulumi.CustomResourceState

	// A human-readable description of the resource.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each
	// extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that
	// matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per
	// resource.
	ExtensionChains NetworkServicesLbTrafficExtensionExtensionChainArrayOutput `pulumi:"extensionChains"`
	// A list of references to the forwarding rules to which this service extension is attached to. At least one forwarding
	// rule is required. There can be only one LBTrafficExtension resource per forwarding rule.
	ForwardingRules pulumi.StringArrayOutput `pulumi:"forwardingRules"`
	// Set of labels associated with the LbTrafficExtension resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. For
	// more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and
	// [Supported application load
	// balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values:
	// ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	LoadBalancingScheme pulumi.StringPtrOutput `pulumi:"loadBalancingScheme"`
	// The location of the traffic extension
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the LbTrafficExtension resource in the following format:
	// projects/{project}/locations/{location}/lbTrafficExtensions/{lbTrafficExtension}.
	Name                                pulumi.StringOutput `pulumi:"name"`
	NetworkServicesLbTrafficExtensionId pulumi.StringOutput `pulumi:"networkServicesLbTrafficExtensionId"`
	Project                             pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                             `pulumi:"terraformLabels"`
	Timeouts        NetworkServicesLbTrafficExtensionTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewNetworkServicesLbTrafficExtension registers a new resource with the given unique name, arguments, and options.
func NewNetworkServicesLbTrafficExtension(ctx *pulumi.Context,
	name string, args *NetworkServicesLbTrafficExtensionArgs, opts ...pulumi.ResourceOption) (*NetworkServicesLbTrafficExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtensionChains == nil {
		return nil, errors.New("invalid value for required argument 'ExtensionChains'")
	}
	if args.ForwardingRules == nil {
		return nil, errors.New("invalid value for required argument 'ForwardingRules'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkServicesLbTrafficExtension
	err = ctx.RegisterPackageResource("google:index/networkServicesLbTrafficExtension:NetworkServicesLbTrafficExtension", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkServicesLbTrafficExtension gets an existing NetworkServicesLbTrafficExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkServicesLbTrafficExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkServicesLbTrafficExtensionState, opts ...pulumi.ResourceOption) (*NetworkServicesLbTrafficExtension, error) {
	var resource NetworkServicesLbTrafficExtension
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/networkServicesLbTrafficExtension:NetworkServicesLbTrafficExtension", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkServicesLbTrafficExtension resources.
type networkServicesLbTrafficExtensionState struct {
	// A human-readable description of the resource.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each
	// extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that
	// matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per
	// resource.
	ExtensionChains []NetworkServicesLbTrafficExtensionExtensionChain `pulumi:"extensionChains"`
	// A list of references to the forwarding rules to which this service extension is attached to. At least one forwarding
	// rule is required. There can be only one LBTrafficExtension resource per forwarding rule.
	ForwardingRules []string `pulumi:"forwardingRules"`
	// Set of labels associated with the LbTrafficExtension resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. For
	// more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and
	// [Supported application load
	// balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values:
	// ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// The location of the traffic extension
	Location *string `pulumi:"location"`
	// Name of the LbTrafficExtension resource in the following format:
	// projects/{project}/locations/{location}/lbTrafficExtensions/{lbTrafficExtension}.
	Name                                *string `pulumi:"name"`
	NetworkServicesLbTrafficExtensionId *string `pulumi:"networkServicesLbTrafficExtensionId"`
	Project                             *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                          `pulumi:"terraformLabels"`
	Timeouts        *NetworkServicesLbTrafficExtensionTimeouts `pulumi:"timeouts"`
}

type NetworkServicesLbTrafficExtensionState struct {
	// A human-readable description of the resource.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each
	// extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that
	// matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per
	// resource.
	ExtensionChains NetworkServicesLbTrafficExtensionExtensionChainArrayInput
	// A list of references to the forwarding rules to which this service extension is attached to. At least one forwarding
	// rule is required. There can be only one LBTrafficExtension resource per forwarding rule.
	ForwardingRules pulumi.StringArrayInput
	// Set of labels associated with the LbTrafficExtension resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. For
	// more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and
	// [Supported application load
	// balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values:
	// ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	LoadBalancingScheme pulumi.StringPtrInput
	// The location of the traffic extension
	Location pulumi.StringPtrInput
	// Name of the LbTrafficExtension resource in the following format:
	// projects/{project}/locations/{location}/lbTrafficExtensions/{lbTrafficExtension}.
	Name                                pulumi.StringPtrInput
	NetworkServicesLbTrafficExtensionId pulumi.StringPtrInput
	Project                             pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkServicesLbTrafficExtensionTimeoutsPtrInput
}

func (NetworkServicesLbTrafficExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesLbTrafficExtensionState)(nil)).Elem()
}

type networkServicesLbTrafficExtensionArgs struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each
	// extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that
	// matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per
	// resource.
	ExtensionChains []NetworkServicesLbTrafficExtensionExtensionChain `pulumi:"extensionChains"`
	// A list of references to the forwarding rules to which this service extension is attached to. At least one forwarding
	// rule is required. There can be only one LBTrafficExtension resource per forwarding rule.
	ForwardingRules []string `pulumi:"forwardingRules"`
	// Set of labels associated with the LbTrafficExtension resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. For
	// more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and
	// [Supported application load
	// balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values:
	// ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// The location of the traffic extension
	Location string `pulumi:"location"`
	// Name of the LbTrafficExtension resource in the following format:
	// projects/{project}/locations/{location}/lbTrafficExtensions/{lbTrafficExtension}.
	Name                                *string                                    `pulumi:"name"`
	NetworkServicesLbTrafficExtensionId *string                                    `pulumi:"networkServicesLbTrafficExtensionId"`
	Project                             *string                                    `pulumi:"project"`
	Timeouts                            *NetworkServicesLbTrafficExtensionTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkServicesLbTrafficExtension resource.
type NetworkServicesLbTrafficExtensionArgs struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each
	// extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that
	// matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per
	// resource.
	ExtensionChains NetworkServicesLbTrafficExtensionExtensionChainArrayInput
	// A list of references to the forwarding rules to which this service extension is attached to. At least one forwarding
	// rule is required. There can be only one LBTrafficExtension resource per forwarding rule.
	ForwardingRules pulumi.StringArrayInput
	// Set of labels associated with the LbTrafficExtension resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. For
	// more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and
	// [Supported application load
	// balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values:
	// ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	LoadBalancingScheme pulumi.StringPtrInput
	// The location of the traffic extension
	Location pulumi.StringInput
	// Name of the LbTrafficExtension resource in the following format:
	// projects/{project}/locations/{location}/lbTrafficExtensions/{lbTrafficExtension}.
	Name                                pulumi.StringPtrInput
	NetworkServicesLbTrafficExtensionId pulumi.StringPtrInput
	Project                             pulumi.StringPtrInput
	Timeouts                            NetworkServicesLbTrafficExtensionTimeoutsPtrInput
}

func (NetworkServicesLbTrafficExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesLbTrafficExtensionArgs)(nil)).Elem()
}

type NetworkServicesLbTrafficExtensionInput interface {
	pulumi.Input

	ToNetworkServicesLbTrafficExtensionOutput() NetworkServicesLbTrafficExtensionOutput
	ToNetworkServicesLbTrafficExtensionOutputWithContext(ctx context.Context) NetworkServicesLbTrafficExtensionOutput
}

func (*NetworkServicesLbTrafficExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesLbTrafficExtension)(nil)).Elem()
}

func (i *NetworkServicesLbTrafficExtension) ToNetworkServicesLbTrafficExtensionOutput() NetworkServicesLbTrafficExtensionOutput {
	return i.ToNetworkServicesLbTrafficExtensionOutputWithContext(context.Background())
}

func (i *NetworkServicesLbTrafficExtension) ToNetworkServicesLbTrafficExtensionOutputWithContext(ctx context.Context) NetworkServicesLbTrafficExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkServicesLbTrafficExtensionOutput)
}

type NetworkServicesLbTrafficExtensionOutput struct{ *pulumi.OutputState }

func (NetworkServicesLbTrafficExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesLbTrafficExtension)(nil)).Elem()
}

func (o NetworkServicesLbTrafficExtensionOutput) ToNetworkServicesLbTrafficExtensionOutput() NetworkServicesLbTrafficExtensionOutput {
	return o
}

func (o NetworkServicesLbTrafficExtensionOutput) ToNetworkServicesLbTrafficExtensionOutputWithContext(ctx context.Context) NetworkServicesLbTrafficExtensionOutput {
	return o
}

// A human-readable description of the resource.
func (o NetworkServicesLbTrafficExtensionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkServicesLbTrafficExtensionOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each
// extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that
// matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per
// resource.
func (o NetworkServicesLbTrafficExtensionOutput) ExtensionChains() NetworkServicesLbTrafficExtensionExtensionChainArrayOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) NetworkServicesLbTrafficExtensionExtensionChainArrayOutput {
		return v.ExtensionChains
	}).(NetworkServicesLbTrafficExtensionExtensionChainArrayOutput)
}

// A list of references to the forwarding rules to which this service extension is attached to. At least one forwarding
// rule is required. There can be only one LBTrafficExtension resource per forwarding rule.
func (o NetworkServicesLbTrafficExtensionOutput) ForwardingRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringArrayOutput { return v.ForwardingRules }).(pulumi.StringArrayOutput)
}

// Set of labels associated with the LbTrafficExtension resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkServicesLbTrafficExtensionOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. For
// more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and
// [Supported application load
// balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values:
// ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
func (o NetworkServicesLbTrafficExtensionOutput) LoadBalancingScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringPtrOutput { return v.LoadBalancingScheme }).(pulumi.StringPtrOutput)
}

// The location of the traffic extension
func (o NetworkServicesLbTrafficExtensionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the LbTrafficExtension resource in the following format:
// projects/{project}/locations/{location}/lbTrafficExtensions/{lbTrafficExtension}.
func (o NetworkServicesLbTrafficExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkServicesLbTrafficExtensionOutput) NetworkServicesLbTrafficExtensionId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringOutput {
		return v.NetworkServicesLbTrafficExtensionId
	}).(pulumi.StringOutput)
}

func (o NetworkServicesLbTrafficExtensionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkServicesLbTrafficExtensionOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkServicesLbTrafficExtensionOutput) Timeouts() NetworkServicesLbTrafficExtensionTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkServicesLbTrafficExtension) NetworkServicesLbTrafficExtensionTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkServicesLbTrafficExtensionTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkServicesLbTrafficExtensionInput)(nil)).Elem(), &NetworkServicesLbTrafficExtension{})
	pulumi.RegisterOutputType(NetworkServicesLbTrafficExtensionOutput{})
}
