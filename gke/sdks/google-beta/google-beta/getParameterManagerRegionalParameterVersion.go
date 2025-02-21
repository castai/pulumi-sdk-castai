// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupParameterManagerRegionalParameterVersion(ctx *pulumi.Context, args *LookupParameterManagerRegionalParameterVersionArgs, opts ...pulumi.InvokeOption) (*LookupParameterManagerRegionalParameterVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupParameterManagerRegionalParameterVersionResult
	err = ctx.InvokePackage("google-beta:index/getParameterManagerRegionalParameterVersion:getParameterManagerRegionalParameterVersion", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getParameterManagerRegionalParameterVersion.
type LookupParameterManagerRegionalParameterVersionArgs struct {
	Id                 *string `pulumi:"id"`
	Location           *string `pulumi:"location"`
	Parameter          string  `pulumi:"parameter"`
	ParameterVersionId string  `pulumi:"parameterVersionId"`
	Project            *string `pulumi:"project"`
}

// A collection of values returned by getParameterManagerRegionalParameterVersion.
type LookupParameterManagerRegionalParameterVersionResult struct {
	CreateTime         string `pulumi:"createTime"`
	Disabled           bool   `pulumi:"disabled"`
	Id                 string `pulumi:"id"`
	Location           string `pulumi:"location"`
	Name               string `pulumi:"name"`
	Parameter          string `pulumi:"parameter"`
	ParameterData      string `pulumi:"parameterData"`
	ParameterVersionId string `pulumi:"parameterVersionId"`
	Project            string `pulumi:"project"`
	UpdateTime         string `pulumi:"updateTime"`
}

func LookupParameterManagerRegionalParameterVersionOutput(ctx *pulumi.Context, args LookupParameterManagerRegionalParameterVersionOutputArgs, opts ...pulumi.InvokeOption) LookupParameterManagerRegionalParameterVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupParameterManagerRegionalParameterVersionResultOutput, error) {
			args := v.(LookupParameterManagerRegionalParameterVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupParameterManagerRegionalParameterVersionResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getParameterManagerRegionalParameterVersion:getParameterManagerRegionalParameterVersion", args, LookupParameterManagerRegionalParameterVersionResultOutput{}, options).(LookupParameterManagerRegionalParameterVersionResultOutput), nil
		}).(LookupParameterManagerRegionalParameterVersionResultOutput)
}

// A collection of arguments for invoking getParameterManagerRegionalParameterVersion.
type LookupParameterManagerRegionalParameterVersionOutputArgs struct {
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Location           pulumi.StringPtrInput `pulumi:"location"`
	Parameter          pulumi.StringInput    `pulumi:"parameter"`
	ParameterVersionId pulumi.StringInput    `pulumi:"parameterVersionId"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupParameterManagerRegionalParameterVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterManagerRegionalParameterVersionArgs)(nil)).Elem()
}

// A collection of values returned by getParameterManagerRegionalParameterVersion.
type LookupParameterManagerRegionalParameterVersionResultOutput struct{ *pulumi.OutputState }

func (LookupParameterManagerRegionalParameterVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterManagerRegionalParameterVersionResult)(nil)).Elem()
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) ToLookupParameterManagerRegionalParameterVersionResultOutput() LookupParameterManagerRegionalParameterVersionResultOutput {
	return o
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) ToLookupParameterManagerRegionalParameterVersionResultOutputWithContext(ctx context.Context) LookupParameterManagerRegionalParameterVersionResultOutput {
	return o
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) Parameter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.Parameter }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) ParameterData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.ParameterData }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) ParameterVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.ParameterVersionId }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupParameterManagerRegionalParameterVersionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerRegionalParameterVersionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterManagerRegionalParameterVersionResultOutput{})
}
