// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOrganization(ctx *pulumi.Context, args *GetOrganizationArgs, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetOrganizationResult
	err = ctx.InvokePackage("google-beta:index/getOrganization:getOrganization", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationArgs struct {
	Domain       *string `pulumi:"domain"`
	Id           *string `pulumi:"id"`
	Organization *string `pulumi:"organization"`
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	CreateTime          string  `pulumi:"createTime"`
	DirectoryCustomerId string  `pulumi:"directoryCustomerId"`
	Domain              string  `pulumi:"domain"`
	Id                  string  `pulumi:"id"`
	LifecycleState      string  `pulumi:"lifecycleState"`
	Name                string  `pulumi:"name"`
	OrgId               string  `pulumi:"orgId"`
	Organization        *string `pulumi:"organization"`
}

func GetOrganizationOutput(ctx *pulumi.Context, args GetOrganizationOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOrganizationResultOutput, error) {
			args := v.(GetOrganizationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetOrganizationResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getOrganization:getOrganization", args, GetOrganizationResultOutput{}, options).(GetOrganizationResultOutput), nil
		}).(GetOrganizationResultOutput)
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationOutputArgs struct {
	Domain       pulumi.StringPtrInput `pulumi:"domain"`
	Id           pulumi.StringPtrInput `pulumi:"id"`
	Organization pulumi.StringPtrInput `pulumi:"organization"`
}

func (GetOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganization.
type GetOrganizationResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationResult)(nil)).Elem()
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutput() GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutputWithContext(ctx context.Context) GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) DirectoryCustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.DirectoryCustomerId }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.LifecycleState }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrganizationResult) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationResultOutput{})
}
