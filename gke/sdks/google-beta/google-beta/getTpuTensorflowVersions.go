// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTpuTensorflowVersions(ctx *pulumi.Context, args *GetTpuTensorflowVersionsArgs, opts ...pulumi.InvokeOption) (*GetTpuTensorflowVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetTpuTensorflowVersionsResult
	err = ctx.InvokePackage("google-beta:index/getTpuTensorflowVersions:getTpuTensorflowVersions", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTpuTensorflowVersions.
type GetTpuTensorflowVersionsArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
	Zone    *string `pulumi:"zone"`
}

// A collection of values returned by getTpuTensorflowVersions.
type GetTpuTensorflowVersionsResult struct {
	Id       string   `pulumi:"id"`
	Project  string   `pulumi:"project"`
	Versions []string `pulumi:"versions"`
	Zone     string   `pulumi:"zone"`
}

func GetTpuTensorflowVersionsOutput(ctx *pulumi.Context, args GetTpuTensorflowVersionsOutputArgs, opts ...pulumi.InvokeOption) GetTpuTensorflowVersionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTpuTensorflowVersionsResultOutput, error) {
			args := v.(GetTpuTensorflowVersionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetTpuTensorflowVersionsResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getTpuTensorflowVersions:getTpuTensorflowVersions", args, GetTpuTensorflowVersionsResultOutput{}, options).(GetTpuTensorflowVersionsResultOutput), nil
		}).(GetTpuTensorflowVersionsResultOutput)
}

// A collection of arguments for invoking getTpuTensorflowVersions.
type GetTpuTensorflowVersionsOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Zone    pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetTpuTensorflowVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTpuTensorflowVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getTpuTensorflowVersions.
type GetTpuTensorflowVersionsResultOutput struct{ *pulumi.OutputState }

func (GetTpuTensorflowVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTpuTensorflowVersionsResult)(nil)).Elem()
}

func (o GetTpuTensorflowVersionsResultOutput) ToGetTpuTensorflowVersionsResultOutput() GetTpuTensorflowVersionsResultOutput {
	return o
}

func (o GetTpuTensorflowVersionsResultOutput) ToGetTpuTensorflowVersionsResultOutputWithContext(ctx context.Context) GetTpuTensorflowVersionsResultOutput {
	return o
}

func (o GetTpuTensorflowVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTpuTensorflowVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTpuTensorflowVersionsResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetTpuTensorflowVersionsResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetTpuTensorflowVersionsResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTpuTensorflowVersionsResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func (o GetTpuTensorflowVersionsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetTpuTensorflowVersionsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTpuTensorflowVersionsResultOutput{})
}
