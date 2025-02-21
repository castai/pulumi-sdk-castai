// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTagsTagKeys(ctx *pulumi.Context, args *GetTagsTagKeysArgs, opts ...pulumi.InvokeOption) (*GetTagsTagKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetTagsTagKeysResult
	err = ctx.InvokePackage("google:index/getTagsTagKeys:getTagsTagKeys", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagsTagKeys.
type GetTagsTagKeysArgs struct {
	Id     *string `pulumi:"id"`
	Parent string  `pulumi:"parent"`
}

// A collection of values returned by getTagsTagKeys.
type GetTagsTagKeysResult struct {
	Id     string              `pulumi:"id"`
	Keys   []GetTagsTagKeysKey `pulumi:"keys"`
	Parent string              `pulumi:"parent"`
}

func GetTagsTagKeysOutput(ctx *pulumi.Context, args GetTagsTagKeysOutputArgs, opts ...pulumi.InvokeOption) GetTagsTagKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTagsTagKeysResultOutput, error) {
			args := v.(GetTagsTagKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetTagsTagKeysResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getTagsTagKeys:getTagsTagKeys", args, GetTagsTagKeysResultOutput{}, options).(GetTagsTagKeysResultOutput), nil
		}).(GetTagsTagKeysResultOutput)
}

// A collection of arguments for invoking getTagsTagKeys.
type GetTagsTagKeysOutputArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Parent pulumi.StringInput    `pulumi:"parent"`
}

func (GetTagsTagKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTagKeysArgs)(nil)).Elem()
}

// A collection of values returned by getTagsTagKeys.
type GetTagsTagKeysResultOutput struct{ *pulumi.OutputState }

func (GetTagsTagKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTagKeysResult)(nil)).Elem()
}

func (o GetTagsTagKeysResultOutput) ToGetTagsTagKeysResultOutput() GetTagsTagKeysResultOutput {
	return o
}

func (o GetTagsTagKeysResultOutput) ToGetTagsTagKeysResultOutputWithContext(ctx context.Context) GetTagsTagKeysResultOutput {
	return o
}

func (o GetTagsTagKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTagKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTagsTagKeysResultOutput) Keys() GetTagsTagKeysKeyArrayOutput {
	return o.ApplyT(func(v GetTagsTagKeysResult) []GetTagsTagKeysKey { return v.Keys }).(GetTagsTagKeysKeyArrayOutput)
}

func (o GetTagsTagKeysResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTagKeysResult) string { return v.Parent }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTagsTagKeysResultOutput{})
}
