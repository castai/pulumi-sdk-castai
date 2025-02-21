// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceAccountKey(ctx *pulumi.Context, args *LookupServiceAccountKeyArgs, opts ...pulumi.InvokeOption) (*LookupServiceAccountKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupServiceAccountKeyResult
	err = ctx.InvokePackage("google:index/getServiceAccountKey:getServiceAccountKey", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccountKey.
type LookupServiceAccountKeyArgs struct {
	Id            *string `pulumi:"id"`
	Name          string  `pulumi:"name"`
	Project       *string `pulumi:"project"`
	PublicKeyType *string `pulumi:"publicKeyType"`
}

// A collection of values returned by getServiceAccountKey.
type LookupServiceAccountKeyResult struct {
	Id            string  `pulumi:"id"`
	KeyAlgorithm  string  `pulumi:"keyAlgorithm"`
	Name          string  `pulumi:"name"`
	Project       *string `pulumi:"project"`
	PublicKey     string  `pulumi:"publicKey"`
	PublicKeyType *string `pulumi:"publicKeyType"`
}

func LookupServiceAccountKeyOutput(ctx *pulumi.Context, args LookupServiceAccountKeyOutputArgs, opts ...pulumi.InvokeOption) LookupServiceAccountKeyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServiceAccountKeyResultOutput, error) {
			args := v.(LookupServiceAccountKeyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupServiceAccountKeyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getServiceAccountKey:getServiceAccountKey", args, LookupServiceAccountKeyResultOutput{}, options).(LookupServiceAccountKeyResultOutput), nil
		}).(LookupServiceAccountKeyResultOutput)
}

// A collection of arguments for invoking getServiceAccountKey.
type LookupServiceAccountKeyOutputArgs struct {
	Id            pulumi.StringPtrInput `pulumi:"id"`
	Name          pulumi.StringInput    `pulumi:"name"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	PublicKeyType pulumi.StringPtrInput `pulumi:"publicKeyType"`
}

func (LookupServiceAccountKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountKeyArgs)(nil)).Elem()
}

// A collection of values returned by getServiceAccountKey.
type LookupServiceAccountKeyResultOutput struct{ *pulumi.OutputState }

func (LookupServiceAccountKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountKeyResult)(nil)).Elem()
}

func (o LookupServiceAccountKeyResultOutput) ToLookupServiceAccountKeyResultOutput() LookupServiceAccountKeyResultOutput {
	return o
}

func (o LookupServiceAccountKeyResultOutput) ToLookupServiceAccountKeyResultOutputWithContext(ctx context.Context) LookupServiceAccountKeyResultOutput {
	return o
}

func (o LookupServiceAccountKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceAccountKeyResultOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountKeyResult) string { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

func (o LookupServiceAccountKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceAccountKeyResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceAccountKeyResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupServiceAccountKeyResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountKeyResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o LookupServiceAccountKeyResultOutput) PublicKeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceAccountKeyResult) *string { return v.PublicKeyType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceAccountKeyResultOutput{})
}
