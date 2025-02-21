// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourcerepoRepositoryIamPolicy(ctx *pulumi.Context, args *LookupSourcerepoRepositoryIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSourcerepoRepositoryIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupSourcerepoRepositoryIamPolicyResult
	err = ctx.InvokePackage("google:index/getSourcerepoRepositoryIamPolicy:getSourcerepoRepositoryIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcerepoRepositoryIamPolicy.
type LookupSourcerepoRepositoryIamPolicyArgs struct {
	Id         *string `pulumi:"id"`
	Project    *string `pulumi:"project"`
	Repository string  `pulumi:"repository"`
}

// A collection of values returned by getSourcerepoRepositoryIamPolicy.
type LookupSourcerepoRepositoryIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Repository string `pulumi:"repository"`
}

func LookupSourcerepoRepositoryIamPolicyOutput(ctx *pulumi.Context, args LookupSourcerepoRepositoryIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSourcerepoRepositoryIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSourcerepoRepositoryIamPolicyResultOutput, error) {
			args := v.(LookupSourcerepoRepositoryIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupSourcerepoRepositoryIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getSourcerepoRepositoryIamPolicy:getSourcerepoRepositoryIamPolicy", args, LookupSourcerepoRepositoryIamPolicyResultOutput{}, options).(LookupSourcerepoRepositoryIamPolicyResultOutput), nil
		}).(LookupSourcerepoRepositoryIamPolicyResultOutput)
}

// A collection of arguments for invoking getSourcerepoRepositoryIamPolicy.
type LookupSourcerepoRepositoryIamPolicyOutputArgs struct {
	Id         pulumi.StringPtrInput `pulumi:"id"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
	Repository pulumi.StringInput    `pulumi:"repository"`
}

func (LookupSourcerepoRepositoryIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcerepoRepositoryIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getSourcerepoRepositoryIamPolicy.
type LookupSourcerepoRepositoryIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSourcerepoRepositoryIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcerepoRepositoryIamPolicyResult)(nil)).Elem()
}

func (o LookupSourcerepoRepositoryIamPolicyResultOutput) ToLookupSourcerepoRepositoryIamPolicyResultOutput() LookupSourcerepoRepositoryIamPolicyResultOutput {
	return o
}

func (o LookupSourcerepoRepositoryIamPolicyResultOutput) ToLookupSourcerepoRepositoryIamPolicyResultOutputWithContext(ctx context.Context) LookupSourcerepoRepositoryIamPolicyResultOutput {
	return o
}

func (o LookupSourcerepoRepositoryIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcerepoRepositoryIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSourcerepoRepositoryIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcerepoRepositoryIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcerepoRepositoryIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcerepoRepositoryIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupSourcerepoRepositoryIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcerepoRepositoryIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupSourcerepoRepositoryIamPolicyResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcerepoRepositoryIamPolicyResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcerepoRepositoryIamPolicyResultOutput{})
}
