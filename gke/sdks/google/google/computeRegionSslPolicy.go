// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeRegionSslPolicy struct {
	pulumi.CustomResourceState

	ComputeRegionSslPolicyId pulumi.StringOutput `pulumi:"computeRegionSslPolicyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be
	// specified in this list. This field must be empty if the profile is not CUSTOM. See the [official
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
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	Project pulumi.StringOutput    `pulumi:"project"`
	// The region where the regional SSL policy resides.
	Region   pulumi.StringOutput                     `pulumi:"region"`
	SelfLink pulumi.StringOutput                     `pulumi:"selfLink"`
	Timeouts ComputeRegionSslPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeRegionSslPolicy registers a new resource with the given unique name, arguments, and options.
func NewComputeRegionSslPolicy(ctx *pulumi.Context,
	name string, args *ComputeRegionSslPolicyArgs, opts ...pulumi.ResourceOption) (*ComputeRegionSslPolicy, error) {
	if args == nil {
		args = &ComputeRegionSslPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeRegionSslPolicy
	err = ctx.RegisterPackageResource("google:index/computeRegionSslPolicy:ComputeRegionSslPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRegionSslPolicy gets an existing ComputeRegionSslPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRegionSslPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRegionSslPolicyState, opts ...pulumi.ResourceOption) (*ComputeRegionSslPolicy, error) {
	var resource ComputeRegionSslPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeRegionSslPolicy:ComputeRegionSslPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRegionSslPolicy resources.
type computeRegionSslPolicyState struct {
	ComputeRegionSslPolicyId *string `pulumi:"computeRegionSslPolicyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be
	// specified in this list. This field must be empty if the profile is not CUSTOM. See the [official
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
	Profile *string `pulumi:"profile"`
	Project *string `pulumi:"project"`
	// The region where the regional SSL policy resides.
	Region   *string                         `pulumi:"region"`
	SelfLink *string                         `pulumi:"selfLink"`
	Timeouts *ComputeRegionSslPolicyTimeouts `pulumi:"timeouts"`
}

type ComputeRegionSslPolicyState struct {
	ComputeRegionSslPolicyId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be
	// specified in this list. This field must be empty if the profile is not CUSTOM. See the [official
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
	Profile pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region where the regional SSL policy resides.
	Region   pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	Timeouts ComputeRegionSslPolicyTimeoutsPtrInput
}

func (ComputeRegionSslPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionSslPolicyState)(nil)).Elem()
}

type computeRegionSslPolicyArgs struct {
	ComputeRegionSslPolicyId *string `pulumi:"computeRegionSslPolicyId"`
	// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be
	// specified in this list. This field must be empty if the profile is not CUSTOM. See the [official
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
	Profile *string `pulumi:"profile"`
	Project *string `pulumi:"project"`
	// The region where the regional SSL policy resides.
	Region   *string                         `pulumi:"region"`
	Timeouts *ComputeRegionSslPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeRegionSslPolicy resource.
type ComputeRegionSslPolicyArgs struct {
	ComputeRegionSslPolicyId pulumi.StringPtrInput
	// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be
	// specified in this list. This field must be empty if the profile is not CUSTOM. See the [official
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
	Profile pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region where the regional SSL policy resides.
	Region   pulumi.StringPtrInput
	Timeouts ComputeRegionSslPolicyTimeoutsPtrInput
}

func (ComputeRegionSslPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionSslPolicyArgs)(nil)).Elem()
}

type ComputeRegionSslPolicyInput interface {
	pulumi.Input

	ToComputeRegionSslPolicyOutput() ComputeRegionSslPolicyOutput
	ToComputeRegionSslPolicyOutputWithContext(ctx context.Context) ComputeRegionSslPolicyOutput
}

func (*ComputeRegionSslPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionSslPolicy)(nil)).Elem()
}

func (i *ComputeRegionSslPolicy) ToComputeRegionSslPolicyOutput() ComputeRegionSslPolicyOutput {
	return i.ToComputeRegionSslPolicyOutputWithContext(context.Background())
}

func (i *ComputeRegionSslPolicy) ToComputeRegionSslPolicyOutputWithContext(ctx context.Context) ComputeRegionSslPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRegionSslPolicyOutput)
}

type ComputeRegionSslPolicyOutput struct{ *pulumi.OutputState }

func (ComputeRegionSslPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionSslPolicy)(nil)).Elem()
}

func (o ComputeRegionSslPolicyOutput) ToComputeRegionSslPolicyOutput() ComputeRegionSslPolicyOutput {
	return o
}

func (o ComputeRegionSslPolicyOutput) ToComputeRegionSslPolicyOutputWithContext(ctx context.Context) ComputeRegionSslPolicyOutput {
	return o
}

func (o ComputeRegionSslPolicyOutput) ComputeRegionSslPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringOutput { return v.ComputeRegionSslPolicyId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeRegionSslPolicyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be
// specified in this list. This field must be empty if the profile is not CUSTOM. See the [official
// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which
// ciphers are available to use. **Note**: this argument *must* be present when using the 'CUSTOM' profile. This argument
// *must not* be present when using any other profile.
func (o ComputeRegionSslPolicyOutput) CustomFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringArrayOutput { return v.CustomFeatures }).(pulumi.StringArrayOutput)
}

// An optional description of this resource.
func (o ComputeRegionSslPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The list of features enabled in the SSL policy.
func (o ComputeRegionSslPolicyOutput) EnabledFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringArrayOutput { return v.EnabledFeatures }).(pulumi.StringArrayOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
func (o ComputeRegionSslPolicyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
func (o ComputeRegionSslPolicyOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringPtrOutput { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeRegionSslPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. If
// using 'CUSTOM', the set of SSL features to enable must be specified in the 'customFeatures' field. See the [official
// documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information
// on what cipher suites each profile provides. If 'CUSTOM' is used, the 'custom_features' attribute **must be set**.
// Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
func (o ComputeRegionSslPolicyOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringPtrOutput { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o ComputeRegionSslPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region where the regional SSL policy resides.
func (o ComputeRegionSslPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeRegionSslPolicyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeRegionSslPolicyOutput) Timeouts() ComputeRegionSslPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeRegionSslPolicy) ComputeRegionSslPolicyTimeoutsPtrOutput { return v.Timeouts }).(ComputeRegionSslPolicyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRegionSslPolicyInput)(nil)).Elem(), &ComputeRegionSslPolicy{})
	pulumi.RegisterOutputType(ComputeRegionSslPolicyOutput{})
}
