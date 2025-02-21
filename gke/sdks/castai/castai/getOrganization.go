// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package castai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOrganization(ctx *pulumi.Context, args *GetOrganizationArgs, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetOrganizationResult
	err = ctx.InvokePackage("castai:index/getOrganization:getOrganization", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationArgs struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func GetOrganizationOutput(ctx *pulumi.Context, args GetOrganizationOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOrganizationResultOutput, error) {
			args := v.(GetOrganizationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetOrganizationResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("castai:index/getOrganization:getOrganization", args, GetOrganizationResultOutput{}, options).(GetOrganizationResultOutput), nil
		}).(GetOrganizationResultOutput)
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationOutputArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringInput    `pulumi:"name"`
}

func (GetOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganization.
type GetOrganizationResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationResult)(nil)).Elem()
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutput() GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutputWithContext(ctx context.Context) GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationResultOutput{})
}
