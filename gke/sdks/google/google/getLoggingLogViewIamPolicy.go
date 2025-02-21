// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoggingLogViewIamPolicy(ctx *pulumi.Context, args *LookupLoggingLogViewIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupLoggingLogViewIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupLoggingLogViewIamPolicyResult
	err = ctx.InvokePackage("google:index/getLoggingLogViewIamPolicy:getLoggingLogViewIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoggingLogViewIamPolicy.
type LookupLoggingLogViewIamPolicyArgs struct {
	Bucket   string  `pulumi:"bucket"`
	Id       *string `pulumi:"id"`
	Location *string `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Parent   string  `pulumi:"parent"`
}

// A collection of values returned by getLoggingLogViewIamPolicy.
type LookupLoggingLogViewIamPolicyResult struct {
	Bucket     string `pulumi:"bucket"`
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Location   string `pulumi:"location"`
	Name       string `pulumi:"name"`
	Parent     string `pulumi:"parent"`
	PolicyData string `pulumi:"policyData"`
}

func LookupLoggingLogViewIamPolicyOutput(ctx *pulumi.Context, args LookupLoggingLogViewIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupLoggingLogViewIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLoggingLogViewIamPolicyResultOutput, error) {
			args := v.(LookupLoggingLogViewIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupLoggingLogViewIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getLoggingLogViewIamPolicy:getLoggingLogViewIamPolicy", args, LookupLoggingLogViewIamPolicyResultOutput{}, options).(LookupLoggingLogViewIamPolicyResultOutput), nil
		}).(LookupLoggingLogViewIamPolicyResultOutput)
}

// A collection of arguments for invoking getLoggingLogViewIamPolicy.
type LookupLoggingLogViewIamPolicyOutputArgs struct {
	Bucket   pulumi.StringInput    `pulumi:"bucket"`
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Parent   pulumi.StringInput    `pulumi:"parent"`
}

func (LookupLoggingLogViewIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggingLogViewIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getLoggingLogViewIamPolicy.
type LookupLoggingLogViewIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupLoggingLogViewIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggingLogViewIamPolicyResult)(nil)).Elem()
}

func (o LookupLoggingLogViewIamPolicyResultOutput) ToLookupLoggingLogViewIamPolicyResultOutput() LookupLoggingLogViewIamPolicyResultOutput {
	return o
}

func (o LookupLoggingLogViewIamPolicyResultOutput) ToLookupLoggingLogViewIamPolicyResultOutputWithContext(ctx context.Context) LookupLoggingLogViewIamPolicyResultOutput {
	return o
}

func (o LookupLoggingLogViewIamPolicyResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingLogViewIamPolicyResult) string { return v.Bucket }).(pulumi.StringOutput)
}

func (o LookupLoggingLogViewIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingLogViewIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupLoggingLogViewIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingLogViewIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLoggingLogViewIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingLogViewIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLoggingLogViewIamPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingLogViewIamPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLoggingLogViewIamPolicyResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingLogViewIamPolicyResult) string { return v.Parent }).(pulumi.StringOutput)
}

func (o LookupLoggingLogViewIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingLogViewIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoggingLogViewIamPolicyResultOutput{})
}
