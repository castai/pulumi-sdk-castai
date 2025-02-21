// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIapTunnelDestGroupIamPolicy(ctx *pulumi.Context, args *LookupIapTunnelDestGroupIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupIapTunnelDestGroupIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupIapTunnelDestGroupIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getIapTunnelDestGroupIamPolicy:getIapTunnelDestGroupIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIapTunnelDestGroupIamPolicy.
type LookupIapTunnelDestGroupIamPolicyArgs struct {
	DestGroup string  `pulumi:"destGroup"`
	Id        *string `pulumi:"id"`
	Project   *string `pulumi:"project"`
	Region    *string `pulumi:"region"`
}

// A collection of values returned by getIapTunnelDestGroupIamPolicy.
type LookupIapTunnelDestGroupIamPolicyResult struct {
	DestGroup  string `pulumi:"destGroup"`
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Region     string `pulumi:"region"`
}

func LookupIapTunnelDestGroupIamPolicyOutput(ctx *pulumi.Context, args LookupIapTunnelDestGroupIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupIapTunnelDestGroupIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIapTunnelDestGroupIamPolicyResultOutput, error) {
			args := v.(LookupIapTunnelDestGroupIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupIapTunnelDestGroupIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getIapTunnelDestGroupIamPolicy:getIapTunnelDestGroupIamPolicy", args, LookupIapTunnelDestGroupIamPolicyResultOutput{}, options).(LookupIapTunnelDestGroupIamPolicyResultOutput), nil
		}).(LookupIapTunnelDestGroupIamPolicyResultOutput)
}

// A collection of arguments for invoking getIapTunnelDestGroupIamPolicy.
type LookupIapTunnelDestGroupIamPolicyOutputArgs struct {
	DestGroup pulumi.StringInput    `pulumi:"destGroup"`
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	Region    pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupIapTunnelDestGroupIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapTunnelDestGroupIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getIapTunnelDestGroupIamPolicy.
type LookupIapTunnelDestGroupIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupIapTunnelDestGroupIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapTunnelDestGroupIamPolicyResult)(nil)).Elem()
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) ToLookupIapTunnelDestGroupIamPolicyResultOutput() LookupIapTunnelDestGroupIamPolicyResultOutput {
	return o
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) ToLookupIapTunnelDestGroupIamPolicyResultOutputWithContext(ctx context.Context) LookupIapTunnelDestGroupIamPolicyResultOutput {
	return o
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) DestGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelDestGroupIamPolicyResult) string { return v.DestGroup }).(pulumi.StringOutput)
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelDestGroupIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelDestGroupIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelDestGroupIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelDestGroupIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupIapTunnelDestGroupIamPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelDestGroupIamPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIapTunnelDestGroupIamPolicyResultOutput{})
}
