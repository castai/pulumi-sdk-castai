// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecureSourceManagerInstanceIamPolicy(ctx *pulumi.Context, args *LookupSecureSourceManagerInstanceIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSecureSourceManagerInstanceIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupSecureSourceManagerInstanceIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getSecureSourceManagerInstanceIamPolicy:getSecureSourceManagerInstanceIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecureSourceManagerInstanceIamPolicy.
type LookupSecureSourceManagerInstanceIamPolicyArgs struct {
	Id         *string `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	Location   *string `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

// A collection of values returned by getSecureSourceManagerInstanceIamPolicy.
type LookupSecureSourceManagerInstanceIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	Location   string `pulumi:"location"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupSecureSourceManagerInstanceIamPolicyOutput(ctx *pulumi.Context, args LookupSecureSourceManagerInstanceIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSecureSourceManagerInstanceIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSecureSourceManagerInstanceIamPolicyResultOutput, error) {
			args := v.(LookupSecureSourceManagerInstanceIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupSecureSourceManagerInstanceIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getSecureSourceManagerInstanceIamPolicy:getSecureSourceManagerInstanceIamPolicy", args, LookupSecureSourceManagerInstanceIamPolicyResultOutput{}, options).(LookupSecureSourceManagerInstanceIamPolicyResultOutput), nil
		}).(LookupSecureSourceManagerInstanceIamPolicyResultOutput)
}

// A collection of arguments for invoking getSecureSourceManagerInstanceIamPolicy.
type LookupSecureSourceManagerInstanceIamPolicyOutputArgs struct {
	Id         pulumi.StringPtrInput `pulumi:"id"`
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Location   pulumi.StringPtrInput `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupSecureSourceManagerInstanceIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecureSourceManagerInstanceIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getSecureSourceManagerInstanceIamPolicy.
type LookupSecureSourceManagerInstanceIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSecureSourceManagerInstanceIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecureSourceManagerInstanceIamPolicyResult)(nil)).Elem()
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) ToLookupSecureSourceManagerInstanceIamPolicyResultOutput() LookupSecureSourceManagerInstanceIamPolicyResultOutput {
	return o
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) ToLookupSecureSourceManagerInstanceIamPolicyResultOutputWithContext(ctx context.Context) LookupSecureSourceManagerInstanceIamPolicyResultOutput {
	return o
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecureSourceManagerInstanceIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecureSourceManagerInstanceIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecureSourceManagerInstanceIamPolicyResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecureSourceManagerInstanceIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecureSourceManagerInstanceIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupSecureSourceManagerInstanceIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecureSourceManagerInstanceIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecureSourceManagerInstanceIamPolicyResultOutput{})
}
