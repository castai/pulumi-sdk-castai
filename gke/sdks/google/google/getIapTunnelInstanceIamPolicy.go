// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIapTunnelInstanceIamPolicy(ctx *pulumi.Context, args *LookupIapTunnelInstanceIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupIapTunnelInstanceIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupIapTunnelInstanceIamPolicyResult
	err = ctx.InvokePackage("google:index/getIapTunnelInstanceIamPolicy:getIapTunnelInstanceIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIapTunnelInstanceIamPolicy.
type LookupIapTunnelInstanceIamPolicyArgs struct {
	Id       *string `pulumi:"id"`
	Instance string  `pulumi:"instance"`
	Project  *string `pulumi:"project"`
	Zone     *string `pulumi:"zone"`
}

// A collection of values returned by getIapTunnelInstanceIamPolicy.
type LookupIapTunnelInstanceIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Instance   string `pulumi:"instance"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Zone       string `pulumi:"zone"`
}

func LookupIapTunnelInstanceIamPolicyOutput(ctx *pulumi.Context, args LookupIapTunnelInstanceIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupIapTunnelInstanceIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIapTunnelInstanceIamPolicyResultOutput, error) {
			args := v.(LookupIapTunnelInstanceIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupIapTunnelInstanceIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getIapTunnelInstanceIamPolicy:getIapTunnelInstanceIamPolicy", args, LookupIapTunnelInstanceIamPolicyResultOutput{}, options).(LookupIapTunnelInstanceIamPolicyResultOutput), nil
		}).(LookupIapTunnelInstanceIamPolicyResultOutput)
}

// A collection of arguments for invoking getIapTunnelInstanceIamPolicy.
type LookupIapTunnelInstanceIamPolicyOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Instance pulumi.StringInput    `pulumi:"instance"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	Zone     pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupIapTunnelInstanceIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapTunnelInstanceIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getIapTunnelInstanceIamPolicy.
type LookupIapTunnelInstanceIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupIapTunnelInstanceIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapTunnelInstanceIamPolicyResult)(nil)).Elem()
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) ToLookupIapTunnelInstanceIamPolicyResultOutput() LookupIapTunnelInstanceIamPolicyResultOutput {
	return o
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) ToLookupIapTunnelInstanceIamPolicyResultOutputWithContext(ctx context.Context) LookupIapTunnelInstanceIamPolicyResultOutput {
	return o
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelInstanceIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelInstanceIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelInstanceIamPolicyResult) string { return v.Instance }).(pulumi.StringOutput)
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelInstanceIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelInstanceIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupIapTunnelInstanceIamPolicyResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapTunnelInstanceIamPolicyResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIapTunnelInstanceIamPolicyResultOutput{})
}
