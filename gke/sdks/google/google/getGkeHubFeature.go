// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGkeHubFeature(ctx *pulumi.Context, args *LookupGkeHubFeatureArgs, opts ...pulumi.InvokeOption) (*LookupGkeHubFeatureResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupGkeHubFeatureResult
	err = ctx.InvokePackage("google:index/getGkeHubFeature:getGkeHubFeature", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGkeHubFeature.
type LookupGkeHubFeatureArgs struct {
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getGkeHubFeature.
type LookupGkeHubFeatureResult struct {
	CreateTime                string                                     `pulumi:"createTime"`
	DeleteTime                string                                     `pulumi:"deleteTime"`
	EffectiveLabels           map[string]string                          `pulumi:"effectiveLabels"`
	FleetDefaultMemberConfigs []GetGkeHubFeatureFleetDefaultMemberConfig `pulumi:"fleetDefaultMemberConfigs"`
	Id                        string                                     `pulumi:"id"`
	Labels                    map[string]string                          `pulumi:"labels"`
	Location                  string                                     `pulumi:"location"`
	Name                      string                                     `pulumi:"name"`
	Project                   *string                                    `pulumi:"project"`
	ResourceStates            []GetGkeHubFeatureResourceState            `pulumi:"resourceStates"`
	Specs                     []GetGkeHubFeatureSpec                     `pulumi:"specs"`
	States                    []GetGkeHubFeatureState                    `pulumi:"states"`
	TerraformLabels           map[string]string                          `pulumi:"terraformLabels"`
	UpdateTime                string                                     `pulumi:"updateTime"`
}

func LookupGkeHubFeatureOutput(ctx *pulumi.Context, args LookupGkeHubFeatureOutputArgs, opts ...pulumi.InvokeOption) LookupGkeHubFeatureResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGkeHubFeatureResultOutput, error) {
			args := v.(LookupGkeHubFeatureArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupGkeHubFeatureResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getGkeHubFeature:getGkeHubFeature", args, LookupGkeHubFeatureResultOutput{}, options).(LookupGkeHubFeatureResultOutput), nil
		}).(LookupGkeHubFeatureResultOutput)
}

// A collection of arguments for invoking getGkeHubFeature.
type LookupGkeHubFeatureOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGkeHubFeatureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubFeatureArgs)(nil)).Elem()
}

// A collection of values returned by getGkeHubFeature.
type LookupGkeHubFeatureResultOutput struct{ *pulumi.OutputState }

func (LookupGkeHubFeatureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubFeatureResult)(nil)).Elem()
}

func (o LookupGkeHubFeatureResultOutput) ToLookupGkeHubFeatureResultOutput() LookupGkeHubFeatureResultOutput {
	return o
}

func (o LookupGkeHubFeatureResultOutput) ToLookupGkeHubFeatureResultOutputWithContext(ctx context.Context) LookupGkeHubFeatureResultOutput {
	return o
}

func (o LookupGkeHubFeatureResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupGkeHubFeatureResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

func (o LookupGkeHubFeatureResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o LookupGkeHubFeatureResultOutput) FleetDefaultMemberConfigs() GetGkeHubFeatureFleetDefaultMemberConfigArrayOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) []GetGkeHubFeatureFleetDefaultMemberConfig {
		return v.FleetDefaultMemberConfigs
	}).(GetGkeHubFeatureFleetDefaultMemberConfigArrayOutput)
}

func (o LookupGkeHubFeatureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGkeHubFeatureResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupGkeHubFeatureResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGkeHubFeatureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGkeHubFeatureResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupGkeHubFeatureResultOutput) ResourceStates() GetGkeHubFeatureResourceStateArrayOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) []GetGkeHubFeatureResourceState { return v.ResourceStates }).(GetGkeHubFeatureResourceStateArrayOutput)
}

func (o LookupGkeHubFeatureResultOutput) Specs() GetGkeHubFeatureSpecArrayOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) []GetGkeHubFeatureSpec { return v.Specs }).(GetGkeHubFeatureSpecArrayOutput)
}

func (o LookupGkeHubFeatureResultOutput) States() GetGkeHubFeatureStateArrayOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) []GetGkeHubFeatureState { return v.States }).(GetGkeHubFeatureStateArrayOutput)
}

func (o LookupGkeHubFeatureResultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) map[string]string { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o LookupGkeHubFeatureResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubFeatureResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGkeHubFeatureResultOutput{})
}
