// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudAssetResourcesSearchAll(ctx *pulumi.Context, args *LookupCloudAssetResourcesSearchAllArgs, opts ...pulumi.InvokeOption) (*LookupCloudAssetResourcesSearchAllResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupCloudAssetResourcesSearchAllResult
	err = ctx.InvokePackage("google-beta:index/getCloudAssetResourcesSearchAll:getCloudAssetResourcesSearchAll", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudAssetResourcesSearchAll.
type LookupCloudAssetResourcesSearchAllArgs struct {
	AssetTypes []string `pulumi:"assetTypes"`
	Id         *string  `pulumi:"id"`
	Query      *string  `pulumi:"query"`
	Scope      string   `pulumi:"scope"`
}

// A collection of values returned by getCloudAssetResourcesSearchAll.
type LookupCloudAssetResourcesSearchAllResult struct {
	AssetTypes []string                                `pulumi:"assetTypes"`
	Id         string                                  `pulumi:"id"`
	Query      *string                                 `pulumi:"query"`
	Results    []GetCloudAssetResourcesSearchAllResult `pulumi:"results"`
	Scope      string                                  `pulumi:"scope"`
}

func LookupCloudAssetResourcesSearchAllOutput(ctx *pulumi.Context, args LookupCloudAssetResourcesSearchAllOutputArgs, opts ...pulumi.InvokeOption) LookupCloudAssetResourcesSearchAllResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCloudAssetResourcesSearchAllResultOutput, error) {
			args := v.(LookupCloudAssetResourcesSearchAllArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupCloudAssetResourcesSearchAllResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getCloudAssetResourcesSearchAll:getCloudAssetResourcesSearchAll", args, LookupCloudAssetResourcesSearchAllResultOutput{}, options).(LookupCloudAssetResourcesSearchAllResultOutput), nil
		}).(LookupCloudAssetResourcesSearchAllResultOutput)
}

// A collection of arguments for invoking getCloudAssetResourcesSearchAll.
type LookupCloudAssetResourcesSearchAllOutputArgs struct {
	AssetTypes pulumi.StringArrayInput `pulumi:"assetTypes"`
	Id         pulumi.StringPtrInput   `pulumi:"id"`
	Query      pulumi.StringPtrInput   `pulumi:"query"`
	Scope      pulumi.StringInput      `pulumi:"scope"`
}

func (LookupCloudAssetResourcesSearchAllOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAssetResourcesSearchAllArgs)(nil)).Elem()
}

// A collection of values returned by getCloudAssetResourcesSearchAll.
type LookupCloudAssetResourcesSearchAllResultOutput struct{ *pulumi.OutputState }

func (LookupCloudAssetResourcesSearchAllResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAssetResourcesSearchAllResult)(nil)).Elem()
}

func (o LookupCloudAssetResourcesSearchAllResultOutput) ToLookupCloudAssetResourcesSearchAllResultOutput() LookupCloudAssetResourcesSearchAllResultOutput {
	return o
}

func (o LookupCloudAssetResourcesSearchAllResultOutput) ToLookupCloudAssetResourcesSearchAllResultOutputWithContext(ctx context.Context) LookupCloudAssetResourcesSearchAllResultOutput {
	return o
}

func (o LookupCloudAssetResourcesSearchAllResultOutput) AssetTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudAssetResourcesSearchAllResult) []string { return v.AssetTypes }).(pulumi.StringArrayOutput)
}

func (o LookupCloudAssetResourcesSearchAllResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAssetResourcesSearchAllResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudAssetResourcesSearchAllResultOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudAssetResourcesSearchAllResult) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o LookupCloudAssetResourcesSearchAllResultOutput) Results() GetCloudAssetResourcesSearchAllResultArrayOutput {
	return o.ApplyT(func(v LookupCloudAssetResourcesSearchAllResult) []GetCloudAssetResourcesSearchAllResult {
		return v.Results
	}).(GetCloudAssetResourcesSearchAllResultArrayOutput)
}

func (o LookupCloudAssetResourcesSearchAllResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAssetResourcesSearchAllResult) string { return v.Scope }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudAssetResourcesSearchAllResultOutput{})
}
