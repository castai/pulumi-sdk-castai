// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApphubDiscoveredService(ctx *pulumi.Context, args *GetApphubDiscoveredServiceArgs, opts ...pulumi.InvokeOption) (*GetApphubDiscoveredServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetApphubDiscoveredServiceResult
	err = ctx.InvokePackage("google:index/getApphubDiscoveredService:getApphubDiscoveredService", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApphubDiscoveredService.
type GetApphubDiscoveredServiceArgs struct {
	Id         *string `pulumi:"id"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
	ServiceUri string  `pulumi:"serviceUri"`
}

// A collection of values returned by getApphubDiscoveredService.
type GetApphubDiscoveredServiceResult struct {
	Id                string                                       `pulumi:"id"`
	Location          string                                       `pulumi:"location"`
	Name              string                                       `pulumi:"name"`
	Project           *string                                      `pulumi:"project"`
	ServiceProperties []GetApphubDiscoveredServiceServiceProperty  `pulumi:"serviceProperties"`
	ServiceReferences []GetApphubDiscoveredServiceServiceReference `pulumi:"serviceReferences"`
	ServiceUri        string                                       `pulumi:"serviceUri"`
}

func GetApphubDiscoveredServiceOutput(ctx *pulumi.Context, args GetApphubDiscoveredServiceOutputArgs, opts ...pulumi.InvokeOption) GetApphubDiscoveredServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApphubDiscoveredServiceResultOutput, error) {
			args := v.(GetApphubDiscoveredServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetApphubDiscoveredServiceResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getApphubDiscoveredService:getApphubDiscoveredService", args, GetApphubDiscoveredServiceResultOutput{}, options).(GetApphubDiscoveredServiceResultOutput), nil
		}).(GetApphubDiscoveredServiceResultOutput)
}

// A collection of arguments for invoking getApphubDiscoveredService.
type GetApphubDiscoveredServiceOutputArgs struct {
	Id         pulumi.StringPtrInput `pulumi:"id"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
	ServiceUri pulumi.StringInput    `pulumi:"serviceUri"`
}

func (GetApphubDiscoveredServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApphubDiscoveredServiceArgs)(nil)).Elem()
}

// A collection of values returned by getApphubDiscoveredService.
type GetApphubDiscoveredServiceResultOutput struct{ *pulumi.OutputState }

func (GetApphubDiscoveredServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApphubDiscoveredServiceResult)(nil)).Elem()
}

func (o GetApphubDiscoveredServiceResultOutput) ToGetApphubDiscoveredServiceResultOutput() GetApphubDiscoveredServiceResultOutput {
	return o
}

func (o GetApphubDiscoveredServiceResultOutput) ToGetApphubDiscoveredServiceResultOutputWithContext(ctx context.Context) GetApphubDiscoveredServiceResultOutput {
	return o
}

func (o GetApphubDiscoveredServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApphubDiscoveredServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetApphubDiscoveredServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetApphubDiscoveredServiceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApphubDiscoveredServiceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o GetApphubDiscoveredServiceResultOutput) ServiceProperties() GetApphubDiscoveredServiceServicePropertyArrayOutput {
	return o.ApplyT(func(v GetApphubDiscoveredServiceResult) []GetApphubDiscoveredServiceServiceProperty {
		return v.ServiceProperties
	}).(GetApphubDiscoveredServiceServicePropertyArrayOutput)
}

func (o GetApphubDiscoveredServiceResultOutput) ServiceReferences() GetApphubDiscoveredServiceServiceReferenceArrayOutput {
	return o.ApplyT(func(v GetApphubDiscoveredServiceResult) []GetApphubDiscoveredServiceServiceReference {
		return v.ServiceReferences
	}).(GetApphubDiscoveredServiceServiceReferenceArrayOutput)
}

func (o GetApphubDiscoveredServiceResultOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredServiceResult) string { return v.ServiceUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApphubDiscoveredServiceResultOutput{})
}
