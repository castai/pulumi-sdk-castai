// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificateManagerCertificateMap(ctx *pulumi.Context, args *LookupCertificateManagerCertificateMapArgs, opts ...pulumi.InvokeOption) (*LookupCertificateManagerCertificateMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupCertificateManagerCertificateMapResult
	err = ctx.InvokePackage("google-beta:index/getCertificateManagerCertificateMap:getCertificateManagerCertificateMap", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificateManagerCertificateMap.
type LookupCertificateManagerCertificateMapArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getCertificateManagerCertificateMap.
type LookupCertificateManagerCertificateMapResult struct {
	CreateTime      string                                          `pulumi:"createTime"`
	Description     string                                          `pulumi:"description"`
	EffectiveLabels map[string]string                               `pulumi:"effectiveLabels"`
	GclbTargets     []GetCertificateManagerCertificateMapGclbTarget `pulumi:"gclbTargets"`
	Id              string                                          `pulumi:"id"`
	Labels          map[string]string                               `pulumi:"labels"`
	Name            string                                          `pulumi:"name"`
	Project         *string                                         `pulumi:"project"`
	TerraformLabels map[string]string                               `pulumi:"terraformLabels"`
	UpdateTime      string                                          `pulumi:"updateTime"`
}

func LookupCertificateManagerCertificateMapOutput(ctx *pulumi.Context, args LookupCertificateManagerCertificateMapOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateManagerCertificateMapResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCertificateManagerCertificateMapResultOutput, error) {
			args := v.(LookupCertificateManagerCertificateMapArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupCertificateManagerCertificateMapResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getCertificateManagerCertificateMap:getCertificateManagerCertificateMap", args, LookupCertificateManagerCertificateMapResultOutput{}, options).(LookupCertificateManagerCertificateMapResultOutput), nil
		}).(LookupCertificateManagerCertificateMapResultOutput)
}

// A collection of arguments for invoking getCertificateManagerCertificateMap.
type LookupCertificateManagerCertificateMapOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupCertificateManagerCertificateMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateManagerCertificateMapArgs)(nil)).Elem()
}

// A collection of values returned by getCertificateManagerCertificateMap.
type LookupCertificateManagerCertificateMapResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateManagerCertificateMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateManagerCertificateMapResult)(nil)).Elem()
}

func (o LookupCertificateManagerCertificateMapResultOutput) ToLookupCertificateManagerCertificateMapResultOutput() LookupCertificateManagerCertificateMapResultOutput {
	return o
}

func (o LookupCertificateManagerCertificateMapResultOutput) ToLookupCertificateManagerCertificateMapResultOutputWithContext(ctx context.Context) LookupCertificateManagerCertificateMapResultOutput {
	return o
}

func (o LookupCertificateManagerCertificateMapResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) GclbTargets() GetCertificateManagerCertificateMapGclbTargetArrayOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) []GetCertificateManagerCertificateMapGclbTarget {
		return v.GclbTargets
	}).(GetCertificateManagerCertificateMapGclbTargetArrayOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) map[string]string { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o LookupCertificateManagerCertificateMapResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateManagerCertificateMapResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateManagerCertificateMapResultOutput{})
}
