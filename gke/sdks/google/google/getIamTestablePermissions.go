// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIamTestablePermissions(ctx *pulumi.Context, args *GetIamTestablePermissionsArgs, opts ...pulumi.InvokeOption) (*GetIamTestablePermissionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetIamTestablePermissionsResult
	err = ctx.InvokePackage("google:index/getIamTestablePermissions:getIamTestablePermissions", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamTestablePermissions.
type GetIamTestablePermissionsArgs struct {
	CustomSupportLevel *string  `pulumi:"customSupportLevel"`
	FullResourceName   string   `pulumi:"fullResourceName"`
	Id                 *string  `pulumi:"id"`
	Stages             []string `pulumi:"stages"`
}

// A collection of values returned by getIamTestablePermissions.
type GetIamTestablePermissionsResult struct {
	CustomSupportLevel *string                               `pulumi:"customSupportLevel"`
	FullResourceName   string                                `pulumi:"fullResourceName"`
	Id                 string                                `pulumi:"id"`
	Permissions        []GetIamTestablePermissionsPermission `pulumi:"permissions"`
	Stages             []string                              `pulumi:"stages"`
}

func GetIamTestablePermissionsOutput(ctx *pulumi.Context, args GetIamTestablePermissionsOutputArgs, opts ...pulumi.InvokeOption) GetIamTestablePermissionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIamTestablePermissionsResultOutput, error) {
			args := v.(GetIamTestablePermissionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetIamTestablePermissionsResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getIamTestablePermissions:getIamTestablePermissions", args, GetIamTestablePermissionsResultOutput{}, options).(GetIamTestablePermissionsResultOutput), nil
		}).(GetIamTestablePermissionsResultOutput)
}

// A collection of arguments for invoking getIamTestablePermissions.
type GetIamTestablePermissionsOutputArgs struct {
	CustomSupportLevel pulumi.StringPtrInput   `pulumi:"customSupportLevel"`
	FullResourceName   pulumi.StringInput      `pulumi:"fullResourceName"`
	Id                 pulumi.StringPtrInput   `pulumi:"id"`
	Stages             pulumi.StringArrayInput `pulumi:"stages"`
}

func (GetIamTestablePermissionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamTestablePermissionsArgs)(nil)).Elem()
}

// A collection of values returned by getIamTestablePermissions.
type GetIamTestablePermissionsResultOutput struct{ *pulumi.OutputState }

func (GetIamTestablePermissionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamTestablePermissionsResult)(nil)).Elem()
}

func (o GetIamTestablePermissionsResultOutput) ToGetIamTestablePermissionsResultOutput() GetIamTestablePermissionsResultOutput {
	return o
}

func (o GetIamTestablePermissionsResultOutput) ToGetIamTestablePermissionsResultOutputWithContext(ctx context.Context) GetIamTestablePermissionsResultOutput {
	return o
}

func (o GetIamTestablePermissionsResultOutput) CustomSupportLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamTestablePermissionsResult) *string { return v.CustomSupportLevel }).(pulumi.StringPtrOutput)
}

func (o GetIamTestablePermissionsResultOutput) FullResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamTestablePermissionsResult) string { return v.FullResourceName }).(pulumi.StringOutput)
}

func (o GetIamTestablePermissionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamTestablePermissionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIamTestablePermissionsResultOutput) Permissions() GetIamTestablePermissionsPermissionArrayOutput {
	return o.ApplyT(func(v GetIamTestablePermissionsResult) []GetIamTestablePermissionsPermission { return v.Permissions }).(GetIamTestablePermissionsPermissionArrayOutput)
}

func (o GetIamTestablePermissionsResultOutput) Stages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIamTestablePermissionsResult) []string { return v.Stages }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIamTestablePermissionsResultOutput{})
}
