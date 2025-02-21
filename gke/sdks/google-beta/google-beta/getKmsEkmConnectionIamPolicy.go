// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKmsEkmConnectionIamPolicy(ctx *pulumi.Context, args *LookupKmsEkmConnectionIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupKmsEkmConnectionIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupKmsEkmConnectionIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getKmsEkmConnectionIamPolicy:getKmsEkmConnectionIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKmsEkmConnectionIamPolicy.
type LookupKmsEkmConnectionIamPolicyArgs struct {
	Id       *string `pulumi:"id"`
	Location *string `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getKmsEkmConnectionIamPolicy.
type LookupKmsEkmConnectionIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Location   string `pulumi:"location"`
	Name       string `pulumi:"name"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupKmsEkmConnectionIamPolicyOutput(ctx *pulumi.Context, args LookupKmsEkmConnectionIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupKmsEkmConnectionIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKmsEkmConnectionIamPolicyResultOutput, error) {
			args := v.(LookupKmsEkmConnectionIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupKmsEkmConnectionIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getKmsEkmConnectionIamPolicy:getKmsEkmConnectionIamPolicy", args, LookupKmsEkmConnectionIamPolicyResultOutput{}, options).(LookupKmsEkmConnectionIamPolicyResultOutput), nil
		}).(LookupKmsEkmConnectionIamPolicyResultOutput)
}

// A collection of arguments for invoking getKmsEkmConnectionIamPolicy.
type LookupKmsEkmConnectionIamPolicyOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupKmsEkmConnectionIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKmsEkmConnectionIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getKmsEkmConnectionIamPolicy.
type LookupKmsEkmConnectionIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupKmsEkmConnectionIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKmsEkmConnectionIamPolicyResult)(nil)).Elem()
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) ToLookupKmsEkmConnectionIamPolicyResultOutput() LookupKmsEkmConnectionIamPolicyResultOutput {
	return o
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) ToLookupKmsEkmConnectionIamPolicyResultOutputWithContext(ctx context.Context) LookupKmsEkmConnectionIamPolicyResultOutput {
	return o
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsEkmConnectionIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsEkmConnectionIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsEkmConnectionIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsEkmConnectionIamPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsEkmConnectionIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupKmsEkmConnectionIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsEkmConnectionIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKmsEkmConnectionIamPolicyResultOutput{})
}
