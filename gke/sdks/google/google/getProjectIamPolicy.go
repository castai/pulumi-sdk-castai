// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProjectIamPolicy(ctx *pulumi.Context, args *LookupProjectIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProjectIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupProjectIamPolicyResult
	err = ctx.InvokePackage("google:index/getProjectIamPolicy:getProjectIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectIamPolicy.
type LookupProjectIamPolicyArgs struct {
	Id      *string `pulumi:"id"`
	Project string  `pulumi:"project"`
}

// A collection of values returned by getProjectIamPolicy.
type LookupProjectIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupProjectIamPolicyOutput(ctx *pulumi.Context, args LookupProjectIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupProjectIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectIamPolicyResultOutput, error) {
			args := v.(LookupProjectIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupProjectIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getProjectIamPolicy:getProjectIamPolicy", args, LookupProjectIamPolicyResultOutput{}, options).(LookupProjectIamPolicyResultOutput), nil
		}).(LookupProjectIamPolicyResultOutput)
}

// A collection of arguments for invoking getProjectIamPolicy.
type LookupProjectIamPolicyOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringInput    `pulumi:"project"`
}

func (LookupProjectIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getProjectIamPolicy.
type LookupProjectIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupProjectIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectIamPolicyResult)(nil)).Elem()
}

func (o LookupProjectIamPolicyResultOutput) ToLookupProjectIamPolicyResultOutput() LookupProjectIamPolicyResultOutput {
	return o
}

func (o LookupProjectIamPolicyResultOutput) ToLookupProjectIamPolicyResultOutputWithContext(ctx context.Context) LookupProjectIamPolicyResultOutput {
	return o
}

func (o LookupProjectIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupProjectIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupProjectIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectIamPolicyResultOutput{})
}
