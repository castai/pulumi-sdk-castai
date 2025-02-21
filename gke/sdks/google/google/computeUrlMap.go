// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeUrlMap struct {
	pulumi.CustomResourceState

	ComputeUrlMapId pulumi.StringOutput `pulumi:"computeUrlMapId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions
	// like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If
	// defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService
	// is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or
	// defaultUrlRedirect must be set.
	DefaultRouteAction ComputeUrlMapDefaultRouteActionPtrOutput `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given rules match.
	DefaultService pulumi.StringPtrOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If
	// defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	DefaultUrlRedirect ComputeUrlMapDefaultUrlRedirectPtrOutput `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The
	// headerAction specified here take effect after headerAction specified under pathMatcher.
	HeaderAction ComputeUrlMapHeaderActionPtrOutput `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	HostRules ComputeUrlMapHostRuleArrayOutput `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId pulumi.Float64Output `pulumi:"mapId"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers ComputeUrlMapPathMatcherArrayOutput `pulumi:"pathMatchers"`
	Project      pulumi.StringOutput                 `pulumi:"project"`
	SelfLink     pulumi.StringOutput                 `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass.
	// You can specify a maximum of 100 tests per UrlMap.
	Tests    ComputeUrlMapTestArrayOutput   `pulumi:"tests"`
	Timeouts ComputeUrlMapTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeUrlMap registers a new resource with the given unique name, arguments, and options.
func NewComputeUrlMap(ctx *pulumi.Context,
	name string, args *ComputeUrlMapArgs, opts ...pulumi.ResourceOption) (*ComputeUrlMap, error) {
	if args == nil {
		args = &ComputeUrlMapArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeUrlMap
	err = ctx.RegisterPackageResource("google:index/computeUrlMap:ComputeUrlMap", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeUrlMap gets an existing ComputeUrlMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeUrlMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeUrlMapState, opts ...pulumi.ResourceOption) (*ComputeUrlMap, error) {
	var resource ComputeUrlMap
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeUrlMap:ComputeUrlMap", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeUrlMap resources.
type computeUrlMapState struct {
	ComputeUrlMapId *string `pulumi:"computeUrlMapId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions
	// like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If
	// defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService
	// is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or
	// defaultUrlRedirect must be set.
	DefaultRouteAction *ComputeUrlMapDefaultRouteAction `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given rules match.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If
	// defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	DefaultUrlRedirect *ComputeUrlMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The
	// headerAction specified here take effect after headerAction specified under pathMatcher.
	HeaderAction *ComputeUrlMapHeaderAction `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	HostRules []ComputeUrlMapHostRule `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId *float64 `pulumi:"mapId"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers []ComputeUrlMapPathMatcher `pulumi:"pathMatchers"`
	Project      *string                    `pulumi:"project"`
	SelfLink     *string                    `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass.
	// You can specify a maximum of 100 tests per UrlMap.
	Tests    []ComputeUrlMapTest    `pulumi:"tests"`
	Timeouts *ComputeUrlMapTimeouts `pulumi:"timeouts"`
}

type ComputeUrlMapState struct {
	ComputeUrlMapId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions
	// like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If
	// defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService
	// is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or
	// defaultUrlRedirect must be set.
	DefaultRouteAction ComputeUrlMapDefaultRouteActionPtrInput
	// The backend service or backend bucket to use when none of the given rules match.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If
	// defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	DefaultUrlRedirect ComputeUrlMapDefaultUrlRedirectPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The
	// headerAction specified here take effect after headerAction specified under pathMatcher.
	HeaderAction ComputeUrlMapHeaderActionPtrInput
	// The list of HostRules to use against the URL.
	HostRules ComputeUrlMapHostRuleArrayInput
	// The unique identifier for the resource.
	MapId pulumi.Float64PtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The list of named PathMatchers to use against the URL.
	PathMatchers ComputeUrlMapPathMatcherArrayInput
	Project      pulumi.StringPtrInput
	SelfLink     pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass.
	// You can specify a maximum of 100 tests per UrlMap.
	Tests    ComputeUrlMapTestArrayInput
	Timeouts ComputeUrlMapTimeoutsPtrInput
}

func (ComputeUrlMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeUrlMapState)(nil)).Elem()
}

type computeUrlMapArgs struct {
	ComputeUrlMapId *string `pulumi:"computeUrlMapId"`
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions
	// like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If
	// defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService
	// is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or
	// defaultUrlRedirect must be set.
	DefaultRouteAction *ComputeUrlMapDefaultRouteAction `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given rules match.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If
	// defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	DefaultUrlRedirect *ComputeUrlMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The
	// headerAction specified here take effect after headerAction specified under pathMatcher.
	HeaderAction *ComputeUrlMapHeaderAction `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	HostRules []ComputeUrlMapHostRule `pulumi:"hostRules"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers []ComputeUrlMapPathMatcher `pulumi:"pathMatchers"`
	Project      *string                    `pulumi:"project"`
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass.
	// You can specify a maximum of 100 tests per UrlMap.
	Tests    []ComputeUrlMapTest    `pulumi:"tests"`
	Timeouts *ComputeUrlMapTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeUrlMap resource.
type ComputeUrlMapArgs struct {
	ComputeUrlMapId pulumi.StringPtrInput
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions
	// like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If
	// defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService
	// is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or
	// defaultUrlRedirect must be set.
	DefaultRouteAction ComputeUrlMapDefaultRouteActionPtrInput
	// The backend service or backend bucket to use when none of the given rules match.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If
	// defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	DefaultUrlRedirect ComputeUrlMapDefaultUrlRedirectPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The
	// headerAction specified here take effect after headerAction specified under pathMatcher.
	HeaderAction ComputeUrlMapHeaderActionPtrInput
	// The list of HostRules to use against the URL.
	HostRules ComputeUrlMapHostRuleArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The list of named PathMatchers to use against the URL.
	PathMatchers ComputeUrlMapPathMatcherArrayInput
	Project      pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass.
	// You can specify a maximum of 100 tests per UrlMap.
	Tests    ComputeUrlMapTestArrayInput
	Timeouts ComputeUrlMapTimeoutsPtrInput
}

func (ComputeUrlMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeUrlMapArgs)(nil)).Elem()
}

type ComputeUrlMapInput interface {
	pulumi.Input

	ToComputeUrlMapOutput() ComputeUrlMapOutput
	ToComputeUrlMapOutputWithContext(ctx context.Context) ComputeUrlMapOutput
}

func (*ComputeUrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeUrlMap)(nil)).Elem()
}

func (i *ComputeUrlMap) ToComputeUrlMapOutput() ComputeUrlMapOutput {
	return i.ToComputeUrlMapOutputWithContext(context.Background())
}

func (i *ComputeUrlMap) ToComputeUrlMapOutputWithContext(ctx context.Context) ComputeUrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeUrlMapOutput)
}

type ComputeUrlMapOutput struct{ *pulumi.OutputState }

func (ComputeUrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeUrlMap)(nil)).Elem()
}

func (o ComputeUrlMapOutput) ToComputeUrlMapOutput() ComputeUrlMapOutput {
	return o
}

func (o ComputeUrlMapOutput) ToComputeUrlMapOutputWithContext(ctx context.Context) ComputeUrlMapOutput {
	return o
}

func (o ComputeUrlMapOutput) ComputeUrlMapId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringOutput { return v.ComputeUrlMapId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeUrlMapOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions
// like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If
// defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService
// is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or
// defaultUrlRedirect must be set.
func (o ComputeUrlMapOutput) DefaultRouteAction() ComputeUrlMapDefaultRouteActionPtrOutput {
	return o.ApplyT(func(v *ComputeUrlMap) ComputeUrlMapDefaultRouteActionPtrOutput { return v.DefaultRouteAction }).(ComputeUrlMapDefaultRouteActionPtrOutput)
}

// The backend service or backend bucket to use when none of the given rules match.
func (o ComputeUrlMapOutput) DefaultService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringPtrOutput { return v.DefaultService }).(pulumi.StringPtrOutput)
}

// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If
// defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
func (o ComputeUrlMapOutput) DefaultUrlRedirect() ComputeUrlMapDefaultUrlRedirectPtrOutput {
	return o.ApplyT(func(v *ComputeUrlMap) ComputeUrlMapDefaultUrlRedirectPtrOutput { return v.DefaultUrlRedirect }).(ComputeUrlMapDefaultUrlRedirectPtrOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o ComputeUrlMapOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
func (o ComputeUrlMapOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Specifies changes to request and response headers that need to take effect for the selected backendService. The
// headerAction specified here take effect after headerAction specified under pathMatcher.
func (o ComputeUrlMapOutput) HeaderAction() ComputeUrlMapHeaderActionPtrOutput {
	return o.ApplyT(func(v *ComputeUrlMap) ComputeUrlMapHeaderActionPtrOutput { return v.HeaderAction }).(ComputeUrlMapHeaderActionPtrOutput)
}

// The list of HostRules to use against the URL.
func (o ComputeUrlMapOutput) HostRules() ComputeUrlMapHostRuleArrayOutput {
	return o.ApplyT(func(v *ComputeUrlMap) ComputeUrlMapHostRuleArrayOutput { return v.HostRules }).(ComputeUrlMapHostRuleArrayOutput)
}

// The unique identifier for the resource.
func (o ComputeUrlMapOutput) MapId() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.Float64Output { return v.MapId }).(pulumi.Float64Output)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeUrlMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of named PathMatchers to use against the URL.
func (o ComputeUrlMapOutput) PathMatchers() ComputeUrlMapPathMatcherArrayOutput {
	return o.ApplyT(func(v *ComputeUrlMap) ComputeUrlMapPathMatcherArrayOutput { return v.PathMatchers }).(ComputeUrlMapPathMatcherArrayOutput)
}

func (o ComputeUrlMapOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeUrlMapOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeUrlMap) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass.
// You can specify a maximum of 100 tests per UrlMap.
func (o ComputeUrlMapOutput) Tests() ComputeUrlMapTestArrayOutput {
	return o.ApplyT(func(v *ComputeUrlMap) ComputeUrlMapTestArrayOutput { return v.Tests }).(ComputeUrlMapTestArrayOutput)
}

func (o ComputeUrlMapOutput) Timeouts() ComputeUrlMapTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeUrlMap) ComputeUrlMapTimeoutsPtrOutput { return v.Timeouts }).(ComputeUrlMapTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeUrlMapInput)(nil)).Elem(), &ComputeUrlMap{})
	pulumi.RegisterOutputType(ComputeUrlMapOutput{})
}
