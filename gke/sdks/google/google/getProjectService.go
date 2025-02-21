// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProjectService(ctx *pulumi.Context, args *LookupProjectServiceArgs, opts ...pulumi.InvokeOption) (*LookupProjectServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupProjectServiceResult
	err = ctx.InvokePackage("google:index/getProjectService:getProjectService", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectService.
type LookupProjectServiceArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
	Service string  `pulumi:"service"`
}

// A collection of values returned by getProjectService.
type LookupProjectServiceResult struct {
	DisableDependentServices bool    `pulumi:"disableDependentServices"`
	DisableOnDestroy         bool    `pulumi:"disableOnDestroy"`
	Id                       string  `pulumi:"id"`
	Project                  *string `pulumi:"project"`
	Service                  string  `pulumi:"service"`
}

func LookupProjectServiceOutput(ctx *pulumi.Context, args LookupProjectServiceOutputArgs, opts ...pulumi.InvokeOption) LookupProjectServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectServiceResultOutput, error) {
			args := v.(LookupProjectServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupProjectServiceResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getProjectService:getProjectService", args, LookupProjectServiceResultOutput{}, options).(LookupProjectServiceResultOutput), nil
		}).(LookupProjectServiceResultOutput)
}

// A collection of arguments for invoking getProjectService.
type LookupProjectServiceOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Service pulumi.StringInput    `pulumi:"service"`
}

func (LookupProjectServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectServiceArgs)(nil)).Elem()
}

// A collection of values returned by getProjectService.
type LookupProjectServiceResultOutput struct{ *pulumi.OutputState }

func (LookupProjectServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectServiceResult)(nil)).Elem()
}

func (o LookupProjectServiceResultOutput) ToLookupProjectServiceResultOutput() LookupProjectServiceResultOutput {
	return o
}

func (o LookupProjectServiceResultOutput) ToLookupProjectServiceResultOutputWithContext(ctx context.Context) LookupProjectServiceResultOutput {
	return o
}

func (o LookupProjectServiceResultOutput) DisableDependentServices() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectServiceResult) bool { return v.DisableDependentServices }).(pulumi.BoolOutput)
}

func (o LookupProjectServiceResultOutput) DisableOnDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectServiceResult) bool { return v.DisableOnDestroy }).(pulumi.BoolOutput)
}

func (o LookupProjectServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectServiceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectServiceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupProjectServiceResultOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectServiceResult) string { return v.Service }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectServiceResultOutput{})
}
