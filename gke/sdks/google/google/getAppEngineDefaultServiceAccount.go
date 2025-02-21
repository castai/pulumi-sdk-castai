// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppEngineDefaultServiceAccount(ctx *pulumi.Context, args *GetAppEngineDefaultServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetAppEngineDefaultServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetAppEngineDefaultServiceAccountResult
	err = ctx.InvokePackage("google:index/getAppEngineDefaultServiceAccount:getAppEngineDefaultServiceAccount", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppEngineDefaultServiceAccount.
type GetAppEngineDefaultServiceAccountArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getAppEngineDefaultServiceAccount.
type GetAppEngineDefaultServiceAccountResult struct {
	DisplayName string `pulumi:"displayName"`
	Email       string `pulumi:"email"`
	Id          string `pulumi:"id"`
	Member      string `pulumi:"member"`
	Name        string `pulumi:"name"`
	Project     string `pulumi:"project"`
	UniqueId    string `pulumi:"uniqueId"`
}

func GetAppEngineDefaultServiceAccountOutput(ctx *pulumi.Context, args GetAppEngineDefaultServiceAccountOutputArgs, opts ...pulumi.InvokeOption) GetAppEngineDefaultServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetAppEngineDefaultServiceAccountResultOutput, error) {
			args := v.(GetAppEngineDefaultServiceAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetAppEngineDefaultServiceAccountResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getAppEngineDefaultServiceAccount:getAppEngineDefaultServiceAccount", args, GetAppEngineDefaultServiceAccountResultOutput{}, options).(GetAppEngineDefaultServiceAccountResultOutput), nil
		}).(GetAppEngineDefaultServiceAccountResultOutput)
}

// A collection of arguments for invoking getAppEngineDefaultServiceAccount.
type GetAppEngineDefaultServiceAccountOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetAppEngineDefaultServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppEngineDefaultServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getAppEngineDefaultServiceAccount.
type GetAppEngineDefaultServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetAppEngineDefaultServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppEngineDefaultServiceAccountResult)(nil)).Elem()
}

func (o GetAppEngineDefaultServiceAccountResultOutput) ToGetAppEngineDefaultServiceAccountResultOutput() GetAppEngineDefaultServiceAccountResultOutput {
	return o
}

func (o GetAppEngineDefaultServiceAccountResultOutput) ToGetAppEngineDefaultServiceAccountResultOutputWithContext(ctx context.Context) GetAppEngineDefaultServiceAccountResultOutput {
	return o
}

func (o GetAppEngineDefaultServiceAccountResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineDefaultServiceAccountResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GetAppEngineDefaultServiceAccountResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineDefaultServiceAccountResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetAppEngineDefaultServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineDefaultServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppEngineDefaultServiceAccountResultOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineDefaultServiceAccountResult) string { return v.Member }).(pulumi.StringOutput)
}

func (o GetAppEngineDefaultServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineDefaultServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetAppEngineDefaultServiceAccountResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineDefaultServiceAccountResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetAppEngineDefaultServiceAccountResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppEngineDefaultServiceAccountResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppEngineDefaultServiceAccountResultOutput{})
}
