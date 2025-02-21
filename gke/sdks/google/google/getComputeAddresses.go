// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetComputeAddresses(ctx *pulumi.Context, args *GetComputeAddressesArgs, opts ...pulumi.InvokeOption) (*GetComputeAddressesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetComputeAddressesResult
	err = ctx.InvokePackage("google:index/getComputeAddresses:getComputeAddresses", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeAddresses.
type GetComputeAddressesArgs struct {
	Filter  *string `pulumi:"filter"`
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getComputeAddresses.
type GetComputeAddressesResult struct {
	Addresses []GetComputeAddressesAddress `pulumi:"addresses"`
	Filter    *string                      `pulumi:"filter"`
	Id        string                       `pulumi:"id"`
	Project   string                       `pulumi:"project"`
	Region    *string                      `pulumi:"region"`
}

func GetComputeAddressesOutput(ctx *pulumi.Context, args GetComputeAddressesOutputArgs, opts ...pulumi.InvokeOption) GetComputeAddressesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetComputeAddressesResultOutput, error) {
			args := v.(GetComputeAddressesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetComputeAddressesResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getComputeAddresses:getComputeAddresses", args, GetComputeAddressesResultOutput{}, options).(GetComputeAddressesResultOutput), nil
		}).(GetComputeAddressesResultOutput)
}

// A collection of arguments for invoking getComputeAddresses.
type GetComputeAddressesOutputArgs struct {
	Filter  pulumi.StringPtrInput `pulumi:"filter"`
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (GetComputeAddressesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeAddressesArgs)(nil)).Elem()
}

// A collection of values returned by getComputeAddresses.
type GetComputeAddressesResultOutput struct{ *pulumi.OutputState }

func (GetComputeAddressesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeAddressesResult)(nil)).Elem()
}

func (o GetComputeAddressesResultOutput) ToGetComputeAddressesResultOutput() GetComputeAddressesResultOutput {
	return o
}

func (o GetComputeAddressesResultOutput) ToGetComputeAddressesResultOutputWithContext(ctx context.Context) GetComputeAddressesResultOutput {
	return o
}

func (o GetComputeAddressesResultOutput) Addresses() GetComputeAddressesAddressArrayOutput {
	return o.ApplyT(func(v GetComputeAddressesResult) []GetComputeAddressesAddress { return v.Addresses }).(GetComputeAddressesAddressArrayOutput)
}

func (o GetComputeAddressesResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeAddressesResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetComputeAddressesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeAddressesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetComputeAddressesResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeAddressesResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetComputeAddressesResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeAddressesResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetComputeAddressesResultOutput{})
}
