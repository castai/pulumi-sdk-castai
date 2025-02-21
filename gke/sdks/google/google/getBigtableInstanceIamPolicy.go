// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBigtableInstanceIamPolicy(ctx *pulumi.Context, args *LookupBigtableInstanceIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBigtableInstanceIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupBigtableInstanceIamPolicyResult
	err = ctx.InvokePackage("google:index/getBigtableInstanceIamPolicy:getBigtableInstanceIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBigtableInstanceIamPolicy.
type LookupBigtableInstanceIamPolicyArgs struct {
	Id       *string `pulumi:"id"`
	Instance string  `pulumi:"instance"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getBigtableInstanceIamPolicy.
type LookupBigtableInstanceIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Instance   string `pulumi:"instance"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupBigtableInstanceIamPolicyOutput(ctx *pulumi.Context, args LookupBigtableInstanceIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBigtableInstanceIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBigtableInstanceIamPolicyResultOutput, error) {
			args := v.(LookupBigtableInstanceIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupBigtableInstanceIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getBigtableInstanceIamPolicy:getBigtableInstanceIamPolicy", args, LookupBigtableInstanceIamPolicyResultOutput{}, options).(LookupBigtableInstanceIamPolicyResultOutput), nil
		}).(LookupBigtableInstanceIamPolicyResultOutput)
}

// A collection of arguments for invoking getBigtableInstanceIamPolicy.
type LookupBigtableInstanceIamPolicyOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Instance pulumi.StringInput    `pulumi:"instance"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBigtableInstanceIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigtableInstanceIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getBigtableInstanceIamPolicy.
type LookupBigtableInstanceIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBigtableInstanceIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigtableInstanceIamPolicyResult)(nil)).Elem()
}

func (o LookupBigtableInstanceIamPolicyResultOutput) ToLookupBigtableInstanceIamPolicyResultOutput() LookupBigtableInstanceIamPolicyResultOutput {
	return o
}

func (o LookupBigtableInstanceIamPolicyResultOutput) ToLookupBigtableInstanceIamPolicyResultOutputWithContext(ctx context.Context) LookupBigtableInstanceIamPolicyResultOutput {
	return o
}

func (o LookupBigtableInstanceIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigtableInstanceIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupBigtableInstanceIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigtableInstanceIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBigtableInstanceIamPolicyResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigtableInstanceIamPolicyResult) string { return v.Instance }).(pulumi.StringOutput)
}

func (o LookupBigtableInstanceIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigtableInstanceIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupBigtableInstanceIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigtableInstanceIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBigtableInstanceIamPolicyResultOutput{})
}
