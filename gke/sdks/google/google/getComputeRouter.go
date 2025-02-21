// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeRouter(ctx *pulumi.Context, args *LookupComputeRouterArgs, opts ...pulumi.InvokeOption) (*LookupComputeRouterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeRouterResult
	err = ctx.InvokePackage("google:index/getComputeRouter:getComputeRouter", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeRouter.
type LookupComputeRouterArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Network string  `pulumi:"network"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getComputeRouter.
type LookupComputeRouterResult struct {
	Bgps                        []GetComputeRouterBgp `pulumi:"bgps"`
	CreationTimestamp           string                `pulumi:"creationTimestamp"`
	Description                 string                `pulumi:"description"`
	EncryptedInterconnectRouter bool                  `pulumi:"encryptedInterconnectRouter"`
	Id                          string                `pulumi:"id"`
	Name                        string                `pulumi:"name"`
	Network                     string                `pulumi:"network"`
	Project                     *string               `pulumi:"project"`
	Region                      *string               `pulumi:"region"`
	SelfLink                    string                `pulumi:"selfLink"`
}

func LookupComputeRouterOutput(ctx *pulumi.Context, args LookupComputeRouterOutputArgs, opts ...pulumi.InvokeOption) LookupComputeRouterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeRouterResultOutput, error) {
			args := v.(LookupComputeRouterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeRouterResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getComputeRouter:getComputeRouter", args, LookupComputeRouterResultOutput{}, options).(LookupComputeRouterResultOutput), nil
		}).(LookupComputeRouterResultOutput)
}

// A collection of arguments for invoking getComputeRouter.
type LookupComputeRouterOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Network pulumi.StringInput    `pulumi:"network"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupComputeRouterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRouterArgs)(nil)).Elem()
}

// A collection of values returned by getComputeRouter.
type LookupComputeRouterResultOutput struct{ *pulumi.OutputState }

func (LookupComputeRouterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRouterResult)(nil)).Elem()
}

func (o LookupComputeRouterResultOutput) ToLookupComputeRouterResultOutput() LookupComputeRouterResultOutput {
	return o
}

func (o LookupComputeRouterResultOutput) ToLookupComputeRouterResultOutputWithContext(ctx context.Context) LookupComputeRouterResultOutput {
	return o
}

func (o LookupComputeRouterResultOutput) Bgps() GetComputeRouterBgpArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) []GetComputeRouterBgp { return v.Bgps }).(GetComputeRouterBgpArrayOutput)
}

func (o LookupComputeRouterResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o LookupComputeRouterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeRouterResultOutput) EncryptedInterconnectRouter() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) bool { return v.EncryptedInterconnectRouter }).(pulumi.BoolOutput)
}

func (o LookupComputeRouterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeRouterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeRouterResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupComputeRouterResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeRouterResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupComputeRouterResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeRouterResultOutput{})
}
