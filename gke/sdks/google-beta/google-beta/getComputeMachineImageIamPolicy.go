// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeMachineImageIamPolicy(ctx *pulumi.Context, args *LookupComputeMachineImageIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupComputeMachineImageIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeMachineImageIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getComputeMachineImageIamPolicy:getComputeMachineImageIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeMachineImageIamPolicy.
type LookupComputeMachineImageIamPolicyArgs struct {
	Id           *string `pulumi:"id"`
	MachineImage string  `pulumi:"machineImage"`
	Project      *string `pulumi:"project"`
}

// A collection of values returned by getComputeMachineImageIamPolicy.
type LookupComputeMachineImageIamPolicyResult struct {
	Etag         string `pulumi:"etag"`
	Id           string `pulumi:"id"`
	MachineImage string `pulumi:"machineImage"`
	PolicyData   string `pulumi:"policyData"`
	Project      string `pulumi:"project"`
}

func LookupComputeMachineImageIamPolicyOutput(ctx *pulumi.Context, args LookupComputeMachineImageIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupComputeMachineImageIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeMachineImageIamPolicyResultOutput, error) {
			args := v.(LookupComputeMachineImageIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeMachineImageIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeMachineImageIamPolicy:getComputeMachineImageIamPolicy", args, LookupComputeMachineImageIamPolicyResultOutput{}, options).(LookupComputeMachineImageIamPolicyResultOutput), nil
		}).(LookupComputeMachineImageIamPolicyResultOutput)
}

// A collection of arguments for invoking getComputeMachineImageIamPolicy.
type LookupComputeMachineImageIamPolicyOutputArgs struct {
	Id           pulumi.StringPtrInput `pulumi:"id"`
	MachineImage pulumi.StringInput    `pulumi:"machineImage"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupComputeMachineImageIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeMachineImageIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getComputeMachineImageIamPolicy.
type LookupComputeMachineImageIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupComputeMachineImageIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeMachineImageIamPolicyResult)(nil)).Elem()
}

func (o LookupComputeMachineImageIamPolicyResultOutput) ToLookupComputeMachineImageIamPolicyResultOutput() LookupComputeMachineImageIamPolicyResultOutput {
	return o
}

func (o LookupComputeMachineImageIamPolicyResultOutput) ToLookupComputeMachineImageIamPolicyResultOutputWithContext(ctx context.Context) LookupComputeMachineImageIamPolicyResultOutput {
	return o
}

func (o LookupComputeMachineImageIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeMachineImageIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupComputeMachineImageIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeMachineImageIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeMachineImageIamPolicyResultOutput) MachineImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeMachineImageIamPolicyResult) string { return v.MachineImage }).(pulumi.StringOutput)
}

func (o LookupComputeMachineImageIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeMachineImageIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupComputeMachineImageIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeMachineImageIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeMachineImageIamPolicyResultOutput{})
}
