// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComposerEnvironment(ctx *pulumi.Context, args *LookupComposerEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupComposerEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComposerEnvironmentResult
	err = ctx.InvokePackage("google:index/getComposerEnvironment:getComposerEnvironment", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComposerEnvironment.
type LookupComposerEnvironmentArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getComposerEnvironment.
type LookupComposerEnvironmentResult struct {
	Configs         []GetComposerEnvironmentConfig        `pulumi:"configs"`
	EffectiveLabels map[string]string                     `pulumi:"effectiveLabels"`
	Id              string                                `pulumi:"id"`
	Labels          map[string]string                     `pulumi:"labels"`
	Name            string                                `pulumi:"name"`
	Project         *string                               `pulumi:"project"`
	Region          *string                               `pulumi:"region"`
	StorageConfigs  []GetComposerEnvironmentStorageConfig `pulumi:"storageConfigs"`
	TerraformLabels map[string]string                     `pulumi:"terraformLabels"`
}

func LookupComposerEnvironmentOutput(ctx *pulumi.Context, args LookupComposerEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupComposerEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComposerEnvironmentResultOutput, error) {
			args := v.(LookupComposerEnvironmentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComposerEnvironmentResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getComposerEnvironment:getComposerEnvironment", args, LookupComposerEnvironmentResultOutput{}, options).(LookupComposerEnvironmentResultOutput), nil
		}).(LookupComposerEnvironmentResultOutput)
}

// A collection of arguments for invoking getComposerEnvironment.
type LookupComposerEnvironmentOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupComposerEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComposerEnvironmentArgs)(nil)).Elem()
}

// A collection of values returned by getComposerEnvironment.
type LookupComposerEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupComposerEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComposerEnvironmentResult)(nil)).Elem()
}

func (o LookupComposerEnvironmentResultOutput) ToLookupComposerEnvironmentResultOutput() LookupComposerEnvironmentResultOutput {
	return o
}

func (o LookupComposerEnvironmentResultOutput) ToLookupComposerEnvironmentResultOutputWithContext(ctx context.Context) LookupComposerEnvironmentResultOutput {
	return o
}

func (o LookupComposerEnvironmentResultOutput) Configs() GetComposerEnvironmentConfigArrayOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) []GetComposerEnvironmentConfig { return v.Configs }).(GetComposerEnvironmentConfigArrayOutput)
}

func (o LookupComposerEnvironmentResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o LookupComposerEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComposerEnvironmentResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupComposerEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComposerEnvironmentResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComposerEnvironmentResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupComposerEnvironmentResultOutput) StorageConfigs() GetComposerEnvironmentStorageConfigArrayOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) []GetComposerEnvironmentStorageConfig { return v.StorageConfigs }).(GetComposerEnvironmentStorageConfigArrayOutput)
}

func (o LookupComposerEnvironmentResultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComposerEnvironmentResult) map[string]string { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComposerEnvironmentResultOutput{})
}
