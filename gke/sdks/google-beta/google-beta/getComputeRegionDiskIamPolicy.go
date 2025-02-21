// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeRegionDiskIamPolicy(ctx *pulumi.Context, args *LookupComputeRegionDiskIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupComputeRegionDiskIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeRegionDiskIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getComputeRegionDiskIamPolicy:getComputeRegionDiskIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeRegionDiskIamPolicy.
type LookupComputeRegionDiskIamPolicyArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getComputeRegionDiskIamPolicy.
type LookupComputeRegionDiskIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Name       string `pulumi:"name"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Region     string `pulumi:"region"`
}

func LookupComputeRegionDiskIamPolicyOutput(ctx *pulumi.Context, args LookupComputeRegionDiskIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupComputeRegionDiskIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeRegionDiskIamPolicyResultOutput, error) {
			args := v.(LookupComputeRegionDiskIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeRegionDiskIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeRegionDiskIamPolicy:getComputeRegionDiskIamPolicy", args, LookupComputeRegionDiskIamPolicyResultOutput{}, options).(LookupComputeRegionDiskIamPolicyResultOutput), nil
		}).(LookupComputeRegionDiskIamPolicyResultOutput)
}

// A collection of arguments for invoking getComputeRegionDiskIamPolicy.
type LookupComputeRegionDiskIamPolicyOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupComputeRegionDiskIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRegionDiskIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getComputeRegionDiskIamPolicy.
type LookupComputeRegionDiskIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupComputeRegionDiskIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRegionDiskIamPolicyResult)(nil)).Elem()
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) ToLookupComputeRegionDiskIamPolicyResultOutput() LookupComputeRegionDiskIamPolicyResultOutput {
	return o
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) ToLookupComputeRegionDiskIamPolicyResultOutputWithContext(ctx context.Context) LookupComputeRegionDiskIamPolicyResultOutput {
	return o
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionDiskIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionDiskIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionDiskIamPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionDiskIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionDiskIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupComputeRegionDiskIamPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionDiskIamPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeRegionDiskIamPolicyResultOutput{})
}
