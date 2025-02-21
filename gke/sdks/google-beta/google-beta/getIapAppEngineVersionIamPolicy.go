// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIapAppEngineVersionIamPolicy(ctx *pulumi.Context, args *LookupIapAppEngineVersionIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupIapAppEngineVersionIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupIapAppEngineVersionIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getIapAppEngineVersionIamPolicy:getIapAppEngineVersionIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIapAppEngineVersionIamPolicy.
type LookupIapAppEngineVersionIamPolicyArgs struct {
	AppId     string  `pulumi:"appId"`
	Id        *string `pulumi:"id"`
	Project   *string `pulumi:"project"`
	Service   string  `pulumi:"service"`
	VersionId string  `pulumi:"versionId"`
}

// A collection of values returned by getIapAppEngineVersionIamPolicy.
type LookupIapAppEngineVersionIamPolicyResult struct {
	AppId      string `pulumi:"appId"`
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Service    string `pulumi:"service"`
	VersionId  string `pulumi:"versionId"`
}

func LookupIapAppEngineVersionIamPolicyOutput(ctx *pulumi.Context, args LookupIapAppEngineVersionIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupIapAppEngineVersionIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIapAppEngineVersionIamPolicyResultOutput, error) {
			args := v.(LookupIapAppEngineVersionIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupIapAppEngineVersionIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getIapAppEngineVersionIamPolicy:getIapAppEngineVersionIamPolicy", args, LookupIapAppEngineVersionIamPolicyResultOutput{}, options).(LookupIapAppEngineVersionIamPolicyResultOutput), nil
		}).(LookupIapAppEngineVersionIamPolicyResultOutput)
}

// A collection of arguments for invoking getIapAppEngineVersionIamPolicy.
type LookupIapAppEngineVersionIamPolicyOutputArgs struct {
	AppId     pulumi.StringInput    `pulumi:"appId"`
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	Service   pulumi.StringInput    `pulumi:"service"`
	VersionId pulumi.StringInput    `pulumi:"versionId"`
}

func (LookupIapAppEngineVersionIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapAppEngineVersionIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getIapAppEngineVersionIamPolicy.
type LookupIapAppEngineVersionIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupIapAppEngineVersionIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapAppEngineVersionIamPolicyResult)(nil)).Elem()
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) ToLookupIapAppEngineVersionIamPolicyResultOutput() LookupIapAppEngineVersionIamPolicyResultOutput {
	return o
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) ToLookupIapAppEngineVersionIamPolicyResultOutputWithContext(ctx context.Context) LookupIapAppEngineVersionIamPolicyResultOutput {
	return o
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapAppEngineVersionIamPolicyResult) string { return v.AppId }).(pulumi.StringOutput)
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapAppEngineVersionIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapAppEngineVersionIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapAppEngineVersionIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapAppEngineVersionIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapAppEngineVersionIamPolicyResult) string { return v.Service }).(pulumi.StringOutput)
}

func (o LookupIapAppEngineVersionIamPolicyResultOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapAppEngineVersionIamPolicyResult) string { return v.VersionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIapAppEngineVersionIamPolicyResultOutput{})
}
