// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRuntimeconfigVariable(ctx *pulumi.Context, args *LookupRuntimeconfigVariableArgs, opts ...pulumi.InvokeOption) (*LookupRuntimeconfigVariableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupRuntimeconfigVariableResult
	err = ctx.InvokePackage("google-beta:index/getRuntimeconfigVariable:getRuntimeconfigVariable", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRuntimeconfigVariable.
type LookupRuntimeconfigVariableArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Parent  string  `pulumi:"parent"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getRuntimeconfigVariable.
type LookupRuntimeconfigVariableResult struct {
	Id         string  `pulumi:"id"`
	Name       string  `pulumi:"name"`
	Parent     string  `pulumi:"parent"`
	Project    *string `pulumi:"project"`
	Text       string  `pulumi:"text"`
	UpdateTime string  `pulumi:"updateTime"`
	Value      string  `pulumi:"value"`
}

func LookupRuntimeconfigVariableOutput(ctx *pulumi.Context, args LookupRuntimeconfigVariableOutputArgs, opts ...pulumi.InvokeOption) LookupRuntimeconfigVariableResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRuntimeconfigVariableResultOutput, error) {
			args := v.(LookupRuntimeconfigVariableArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupRuntimeconfigVariableResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getRuntimeconfigVariable:getRuntimeconfigVariable", args, LookupRuntimeconfigVariableResultOutput{}, options).(LookupRuntimeconfigVariableResultOutput), nil
		}).(LookupRuntimeconfigVariableResultOutput)
}

// A collection of arguments for invoking getRuntimeconfigVariable.
type LookupRuntimeconfigVariableOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Parent  pulumi.StringInput    `pulumi:"parent"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupRuntimeconfigVariableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuntimeconfigVariableArgs)(nil)).Elem()
}

// A collection of values returned by getRuntimeconfigVariable.
type LookupRuntimeconfigVariableResultOutput struct{ *pulumi.OutputState }

func (LookupRuntimeconfigVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuntimeconfigVariableResult)(nil)).Elem()
}

func (o LookupRuntimeconfigVariableResultOutput) ToLookupRuntimeconfigVariableResultOutput() LookupRuntimeconfigVariableResultOutput {
	return o
}

func (o LookupRuntimeconfigVariableResultOutput) ToLookupRuntimeconfigVariableResultOutputWithContext(ctx context.Context) LookupRuntimeconfigVariableResultOutput {
	return o
}

func (o LookupRuntimeconfigVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeconfigVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRuntimeconfigVariableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeconfigVariableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRuntimeconfigVariableResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeconfigVariableResult) string { return v.Parent }).(pulumi.StringOutput)
}

func (o LookupRuntimeconfigVariableResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuntimeconfigVariableResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupRuntimeconfigVariableResultOutput) Text() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeconfigVariableResult) string { return v.Text }).(pulumi.StringOutput)
}

func (o LookupRuntimeconfigVariableResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeconfigVariableResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func (o LookupRuntimeconfigVariableResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeconfigVariableResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuntimeconfigVariableResultOutput{})
}
