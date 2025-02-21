// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiGatewayGatewayIamPolicy(ctx *pulumi.Context, args *LookupApiGatewayGatewayIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiGatewayGatewayIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupApiGatewayGatewayIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getApiGatewayGatewayIamPolicy:getApiGatewayGatewayIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiGatewayGatewayIamPolicy.
type LookupApiGatewayGatewayIamPolicyArgs struct {
	Gateway string  `pulumi:"gateway"`
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getApiGatewayGatewayIamPolicy.
type LookupApiGatewayGatewayIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Gateway    string `pulumi:"gateway"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Region     string `pulumi:"region"`
}

func LookupApiGatewayGatewayIamPolicyOutput(ctx *pulumi.Context, args LookupApiGatewayGatewayIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupApiGatewayGatewayIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApiGatewayGatewayIamPolicyResultOutput, error) {
			args := v.(LookupApiGatewayGatewayIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupApiGatewayGatewayIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getApiGatewayGatewayIamPolicy:getApiGatewayGatewayIamPolicy", args, LookupApiGatewayGatewayIamPolicyResultOutput{}, options).(LookupApiGatewayGatewayIamPolicyResultOutput), nil
		}).(LookupApiGatewayGatewayIamPolicyResultOutput)
}

// A collection of arguments for invoking getApiGatewayGatewayIamPolicy.
type LookupApiGatewayGatewayIamPolicyOutputArgs struct {
	Gateway pulumi.StringInput    `pulumi:"gateway"`
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupApiGatewayGatewayIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayGatewayIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getApiGatewayGatewayIamPolicy.
type LookupApiGatewayGatewayIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupApiGatewayGatewayIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayGatewayIamPolicyResult)(nil)).Elem()
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) ToLookupApiGatewayGatewayIamPolicyResultOutput() LookupApiGatewayGatewayIamPolicyResultOutput {
	return o
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) ToLookupApiGatewayGatewayIamPolicyResultOutputWithContext(ctx context.Context) LookupApiGatewayGatewayIamPolicyResultOutput {
	return o
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayGatewayIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayGatewayIamPolicyResult) string { return v.Gateway }).(pulumi.StringOutput)
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayGatewayIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayGatewayIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayGatewayIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupApiGatewayGatewayIamPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayGatewayIamPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiGatewayGatewayIamPolicyResultOutput{})
}
