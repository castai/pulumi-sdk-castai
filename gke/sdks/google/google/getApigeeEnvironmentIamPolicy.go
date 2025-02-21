// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApigeeEnvironmentIamPolicy(ctx *pulumi.Context, args *LookupApigeeEnvironmentIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApigeeEnvironmentIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupApigeeEnvironmentIamPolicyResult
	err = ctx.InvokePackage("google:index/getApigeeEnvironmentIamPolicy:getApigeeEnvironmentIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApigeeEnvironmentIamPolicy.
type LookupApigeeEnvironmentIamPolicyArgs struct {
	EnvId string  `pulumi:"envId"`
	Id    *string `pulumi:"id"`
	OrgId string  `pulumi:"orgId"`
}

// A collection of values returned by getApigeeEnvironmentIamPolicy.
type LookupApigeeEnvironmentIamPolicyResult struct {
	EnvId      string `pulumi:"envId"`
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	OrgId      string `pulumi:"orgId"`
	PolicyData string `pulumi:"policyData"`
}

func LookupApigeeEnvironmentIamPolicyOutput(ctx *pulumi.Context, args LookupApigeeEnvironmentIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupApigeeEnvironmentIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApigeeEnvironmentIamPolicyResultOutput, error) {
			args := v.(LookupApigeeEnvironmentIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupApigeeEnvironmentIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getApigeeEnvironmentIamPolicy:getApigeeEnvironmentIamPolicy", args, LookupApigeeEnvironmentIamPolicyResultOutput{}, options).(LookupApigeeEnvironmentIamPolicyResultOutput), nil
		}).(LookupApigeeEnvironmentIamPolicyResultOutput)
}

// A collection of arguments for invoking getApigeeEnvironmentIamPolicy.
type LookupApigeeEnvironmentIamPolicyOutputArgs struct {
	EnvId pulumi.StringInput    `pulumi:"envId"`
	Id    pulumi.StringPtrInput `pulumi:"id"`
	OrgId pulumi.StringInput    `pulumi:"orgId"`
}

func (LookupApigeeEnvironmentIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApigeeEnvironmentIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getApigeeEnvironmentIamPolicy.
type LookupApigeeEnvironmentIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupApigeeEnvironmentIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApigeeEnvironmentIamPolicyResult)(nil)).Elem()
}

func (o LookupApigeeEnvironmentIamPolicyResultOutput) ToLookupApigeeEnvironmentIamPolicyResultOutput() LookupApigeeEnvironmentIamPolicyResultOutput {
	return o
}

func (o LookupApigeeEnvironmentIamPolicyResultOutput) ToLookupApigeeEnvironmentIamPolicyResultOutputWithContext(ctx context.Context) LookupApigeeEnvironmentIamPolicyResultOutput {
	return o
}

func (o LookupApigeeEnvironmentIamPolicyResultOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApigeeEnvironmentIamPolicyResult) string { return v.EnvId }).(pulumi.StringOutput)
}

func (o LookupApigeeEnvironmentIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApigeeEnvironmentIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApigeeEnvironmentIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApigeeEnvironmentIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApigeeEnvironmentIamPolicyResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApigeeEnvironmentIamPolicyResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o LookupApigeeEnvironmentIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApigeeEnvironmentIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApigeeEnvironmentIamPolicyResultOutput{})
}
