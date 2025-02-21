// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataprocJobIamPolicy(ctx *pulumi.Context, args *LookupDataprocJobIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDataprocJobIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupDataprocJobIamPolicyResult
	err = ctx.InvokePackage("google:index/getDataprocJobIamPolicy:getDataprocJobIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataprocJobIamPolicy.
type LookupDataprocJobIamPolicyArgs struct {
	Id      *string `pulumi:"id"`
	JobId   string  `pulumi:"jobId"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getDataprocJobIamPolicy.
type LookupDataprocJobIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	JobId      string `pulumi:"jobId"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Region     string `pulumi:"region"`
}

func LookupDataprocJobIamPolicyOutput(ctx *pulumi.Context, args LookupDataprocJobIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDataprocJobIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataprocJobIamPolicyResultOutput, error) {
			args := v.(LookupDataprocJobIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupDataprocJobIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getDataprocJobIamPolicy:getDataprocJobIamPolicy", args, LookupDataprocJobIamPolicyResultOutput{}, options).(LookupDataprocJobIamPolicyResultOutput), nil
		}).(LookupDataprocJobIamPolicyResultOutput)
}

// A collection of arguments for invoking getDataprocJobIamPolicy.
type LookupDataprocJobIamPolicyOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	JobId   pulumi.StringInput    `pulumi:"jobId"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupDataprocJobIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataprocJobIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getDataprocJobIamPolicy.
type LookupDataprocJobIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDataprocJobIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataprocJobIamPolicyResult)(nil)).Elem()
}

func (o LookupDataprocJobIamPolicyResultOutput) ToLookupDataprocJobIamPolicyResultOutput() LookupDataprocJobIamPolicyResultOutput {
	return o
}

func (o LookupDataprocJobIamPolicyResultOutput) ToLookupDataprocJobIamPolicyResultOutputWithContext(ctx context.Context) LookupDataprocJobIamPolicyResultOutput {
	return o
}

func (o LookupDataprocJobIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocJobIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDataprocJobIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocJobIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataprocJobIamPolicyResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocJobIamPolicyResult) string { return v.JobId }).(pulumi.StringOutput)
}

func (o LookupDataprocJobIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocJobIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupDataprocJobIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocJobIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupDataprocJobIamPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocJobIamPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataprocJobIamPolicyResultOutput{})
}
