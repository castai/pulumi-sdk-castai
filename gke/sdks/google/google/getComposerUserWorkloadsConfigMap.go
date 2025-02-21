// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComposerUserWorkloadsConfigMap(ctx *pulumi.Context, args *LookupComposerUserWorkloadsConfigMapArgs, opts ...pulumi.InvokeOption) (*LookupComposerUserWorkloadsConfigMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComposerUserWorkloadsConfigMapResult
	err = ctx.InvokePackage("google:index/getComposerUserWorkloadsConfigMap:getComposerUserWorkloadsConfigMap", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComposerUserWorkloadsConfigMap.
type LookupComposerUserWorkloadsConfigMapArgs struct {
	Environment string  `pulumi:"environment"`
	Id          *string `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Project     *string `pulumi:"project"`
	Region      *string `pulumi:"region"`
}

// A collection of values returned by getComposerUserWorkloadsConfigMap.
type LookupComposerUserWorkloadsConfigMapResult struct {
	Data        map[string]string `pulumi:"data"`
	Environment string            `pulumi:"environment"`
	Id          string            `pulumi:"id"`
	Name        string            `pulumi:"name"`
	Project     *string           `pulumi:"project"`
	Region      *string           `pulumi:"region"`
}

func LookupComposerUserWorkloadsConfigMapOutput(ctx *pulumi.Context, args LookupComposerUserWorkloadsConfigMapOutputArgs, opts ...pulumi.InvokeOption) LookupComposerUserWorkloadsConfigMapResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComposerUserWorkloadsConfigMapResultOutput, error) {
			args := v.(LookupComposerUserWorkloadsConfigMapArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComposerUserWorkloadsConfigMapResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getComposerUserWorkloadsConfigMap:getComposerUserWorkloadsConfigMap", args, LookupComposerUserWorkloadsConfigMapResultOutput{}, options).(LookupComposerUserWorkloadsConfigMapResultOutput), nil
		}).(LookupComposerUserWorkloadsConfigMapResultOutput)
}

// A collection of arguments for invoking getComposerUserWorkloadsConfigMap.
type LookupComposerUserWorkloadsConfigMapOutputArgs struct {
	Environment pulumi.StringInput    `pulumi:"environment"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	Region      pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupComposerUserWorkloadsConfigMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComposerUserWorkloadsConfigMapArgs)(nil)).Elem()
}

// A collection of values returned by getComposerUserWorkloadsConfigMap.
type LookupComposerUserWorkloadsConfigMapResultOutput struct{ *pulumi.OutputState }

func (LookupComposerUserWorkloadsConfigMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComposerUserWorkloadsConfigMapResult)(nil)).Elem()
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) ToLookupComposerUserWorkloadsConfigMapResultOutput() LookupComposerUserWorkloadsConfigMapResultOutput {
	return o
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) ToLookupComposerUserWorkloadsConfigMapResultOutputWithContext(ctx context.Context) LookupComposerUserWorkloadsConfigMapResultOutput {
	return o
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComposerUserWorkloadsConfigMapResult) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComposerUserWorkloadsConfigMapResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComposerUserWorkloadsConfigMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComposerUserWorkloadsConfigMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComposerUserWorkloadsConfigMapResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComposerUserWorkloadsConfigMapResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComposerUserWorkloadsConfigMapResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComposerUserWorkloadsConfigMapResultOutput{})
}
