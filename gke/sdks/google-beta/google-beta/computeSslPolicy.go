// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeSslPolicy struct {
	pulumi.CustomResourceState

	ComputeSslPolicyId pulumi.StringOutput `pulumi:"computeSslPolicyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This
	// can be one of 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM', the set of SSL features to enable
	// must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which
	// ciphers are available to use. **Note**: this argument *must* be present when using the 'CUSTOM' profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayOutput `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayOutput `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
	// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
	MinTlsVersion pulumi.StringPtrOutput `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. If
	// using 'CUSTOM', the set of SSL features to enable must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information
	// on what cipher suites each profile provides. If 'CUSTOM' is used, the 'custom_features' attribute **must be set**.
	// Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
	Profile  pulumi.StringPtrOutput            `pulumi:"profile"`
	Project  pulumi.StringOutput               `pulumi:"project"`
	SelfLink pulumi.StringOutput               `pulumi:"selfLink"`
	Timeouts ComputeSslPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeSslPolicy registers a new resource with the given unique name, arguments, and options.
func NewComputeSslPolicy(ctx *pulumi.Context,
	name string, args *ComputeSslPolicyArgs, opts ...pulumi.ResourceOption) (*ComputeSslPolicy, error) {
	if args == nil {
		args = &ComputeSslPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeSslPolicy
	err = ctx.RegisterPackageResource("google-beta:index/computeSslPolicy:ComputeSslPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeSslPolicy gets an existing ComputeSslPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeSslPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeSslPolicyState, opts ...pulumi.ResourceOption) (*ComputeSslPolicy, error) {
	var resource ComputeSslPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeSslPolicy:ComputeSslPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeSslPolicy resources.
type computeSslPolicyState struct {
	ComputeSslPolicyId *string `pulumi:"computeSslPolicyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This
	// can be one of 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM', the set of SSL features to enable
	// must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which
	// ciphers are available to use. **Note**: this argument *must* be present when using the 'CUSTOM' profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures []string `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
	// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. If
	// using 'CUSTOM', the set of SSL features to enable must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information
	// on what cipher suites each profile provides. If 'CUSTOM' is used, the 'custom_features' attribute **must be set**.
	// Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
	Profile  *string                   `pulumi:"profile"`
	Project  *string                   `pulumi:"project"`
	SelfLink *string                   `pulumi:"selfLink"`
	Timeouts *ComputeSslPolicyTimeouts `pulumi:"timeouts"`
}

type ComputeSslPolicyState struct {
	ComputeSslPolicyId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This
	// can be one of 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM', the set of SSL features to enable
	// must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which
	// ciphers are available to use. **Note**: this argument *must* be present when using the 'CUSTOM' profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
	// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
	MinTlsVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. If
	// using 'CUSTOM', the set of SSL features to enable must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information
	// on what cipher suites each profile provides. If 'CUSTOM' is used, the 'custom_features' attribute **must be set**.
	// Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
	Profile  pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	Timeouts ComputeSslPolicyTimeoutsPtrInput
}

func (ComputeSslPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeSslPolicyState)(nil)).Elem()
}

type computeSslPolicyArgs struct {
	ComputeSslPolicyId *string `pulumi:"computeSslPolicyId"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This
	// can be one of 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM', the set of SSL features to enable
	// must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which
	// ciphers are available to use. **Note**: this argument *must* be present when using the 'CUSTOM' profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
	// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. If
	// using 'CUSTOM', the set of SSL features to enable must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information
	// on what cipher suites each profile provides. If 'CUSTOM' is used, the 'custom_features' attribute **must be set**.
	// Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
	Profile  *string                   `pulumi:"profile"`
	Project  *string                   `pulumi:"project"`
	Timeouts *ComputeSslPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeSslPolicy resource.
type ComputeSslPolicyArgs struct {
	ComputeSslPolicyId pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This
	// can be one of 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM', the set of SSL features to enable
	// must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which
	// ciphers are available to use. **Note**: this argument *must* be present when using the 'CUSTOM' profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
	// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
	MinTlsVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. If
	// using 'CUSTOM', the set of SSL features to enable must be specified in the 'customFeatures' field. See the [official
	// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information
	// on what cipher suites each profile provides. If 'CUSTOM' is used, the 'custom_features' attribute **must be set**.
	// Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
	Profile  pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeSslPolicyTimeoutsPtrInput
}

func (ComputeSslPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeSslPolicyArgs)(nil)).Elem()
}

type ComputeSslPolicyInput interface {
	pulumi.Input

	ToComputeSslPolicyOutput() ComputeSslPolicyOutput
	ToComputeSslPolicyOutputWithContext(ctx context.Context) ComputeSslPolicyOutput
}

func (*ComputeSslPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeSslPolicy)(nil)).Elem()
}

func (i *ComputeSslPolicy) ToComputeSslPolicyOutput() ComputeSslPolicyOutput {
	return i.ToComputeSslPolicyOutputWithContext(context.Background())
}

func (i *ComputeSslPolicy) ToComputeSslPolicyOutputWithContext(ctx context.Context) ComputeSslPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeSslPolicyOutput)
}

type ComputeSslPolicyOutput struct{ *pulumi.OutputState }

func (ComputeSslPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeSslPolicy)(nil)).Elem()
}

func (o ComputeSslPolicyOutput) ToComputeSslPolicyOutput() ComputeSslPolicyOutput {
	return o
}

func (o ComputeSslPolicyOutput) ToComputeSslPolicyOutputWithContext(ctx context.Context) ComputeSslPolicyOutput {
	return o
}

func (o ComputeSslPolicyOutput) ComputeSslPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringOutput { return v.ComputeSslPolicyId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeSslPolicyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This
// can be one of 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM', the set of SSL features to enable
// must be specified in the 'customFeatures' field. See the [official
// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which
// ciphers are available to use. **Note**: this argument *must* be present when using the 'CUSTOM' profile. This argument
// *must not* be present when using any other profile.
func (o ComputeSslPolicyOutput) CustomFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringArrayOutput { return v.CustomFeatures }).(pulumi.StringArrayOutput)
}

// An optional description of this resource.
func (o ComputeSslPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The list of features enabled in the SSL policy.
func (o ComputeSslPolicyOutput) EnabledFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringArrayOutput { return v.EnabledFeatures }).(pulumi.StringArrayOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
func (o ComputeSslPolicyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
func (o ComputeSslPolicyOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringPtrOutput { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeSslPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. If
// using 'CUSTOM', the set of SSL features to enable must be specified in the 'customFeatures' field. See the [official
// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information
// on what cipher suites each profile provides. If 'CUSTOM' is used, the 'custom_features' attribute **must be set**.
// Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
func (o ComputeSslPolicyOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringPtrOutput { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o ComputeSslPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeSslPolicyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeSslPolicyOutput) Timeouts() ComputeSslPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeSslPolicy) ComputeSslPolicyTimeoutsPtrOutput { return v.Timeouts }).(ComputeSslPolicyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeSslPolicyInput)(nil)).Elem(), &ComputeSslPolicy{})
	pulumi.RegisterOutputType(ComputeSslPolicyOutput{})
}
