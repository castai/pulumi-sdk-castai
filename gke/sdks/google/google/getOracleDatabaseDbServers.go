// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOracleDatabaseDbServers(ctx *pulumi.Context, args *GetOracleDatabaseDbServersArgs, opts ...pulumi.InvokeOption) (*GetOracleDatabaseDbServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetOracleDatabaseDbServersResult
	err = ctx.InvokePackage("google:index/getOracleDatabaseDbServers:getOracleDatabaseDbServers", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOracleDatabaseDbServers.
type GetOracleDatabaseDbServersArgs struct {
	CloudExadataInfrastructure string  `pulumi:"cloudExadataInfrastructure"`
	Id                         *string `pulumi:"id"`
	Location                   string  `pulumi:"location"`
	Project                    *string `pulumi:"project"`
}

// A collection of values returned by getOracleDatabaseDbServers.
type GetOracleDatabaseDbServersResult struct {
	CloudExadataInfrastructure string                               `pulumi:"cloudExadataInfrastructure"`
	DbServers                  []GetOracleDatabaseDbServersDbServer `pulumi:"dbServers"`
	Id                         string                               `pulumi:"id"`
	Location                   string                               `pulumi:"location"`
	Project                    *string                              `pulumi:"project"`
}

func GetOracleDatabaseDbServersOutput(ctx *pulumi.Context, args GetOracleDatabaseDbServersOutputArgs, opts ...pulumi.InvokeOption) GetOracleDatabaseDbServersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOracleDatabaseDbServersResultOutput, error) {
			args := v.(GetOracleDatabaseDbServersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetOracleDatabaseDbServersResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getOracleDatabaseDbServers:getOracleDatabaseDbServers", args, GetOracleDatabaseDbServersResultOutput{}, options).(GetOracleDatabaseDbServersResultOutput), nil
		}).(GetOracleDatabaseDbServersResultOutput)
}

// A collection of arguments for invoking getOracleDatabaseDbServers.
type GetOracleDatabaseDbServersOutputArgs struct {
	CloudExadataInfrastructure pulumi.StringInput    `pulumi:"cloudExadataInfrastructure"`
	Id                         pulumi.StringPtrInput `pulumi:"id"`
	Location                   pulumi.StringInput    `pulumi:"location"`
	Project                    pulumi.StringPtrInput `pulumi:"project"`
}

func (GetOracleDatabaseDbServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseDbServersArgs)(nil)).Elem()
}

// A collection of values returned by getOracleDatabaseDbServers.
type GetOracleDatabaseDbServersResultOutput struct{ *pulumi.OutputState }

func (GetOracleDatabaseDbServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseDbServersResult)(nil)).Elem()
}

func (o GetOracleDatabaseDbServersResultOutput) ToGetOracleDatabaseDbServersResultOutput() GetOracleDatabaseDbServersResultOutput {
	return o
}

func (o GetOracleDatabaseDbServersResultOutput) ToGetOracleDatabaseDbServersResultOutputWithContext(ctx context.Context) GetOracleDatabaseDbServersResultOutput {
	return o
}

func (o GetOracleDatabaseDbServersResultOutput) CloudExadataInfrastructure() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbServersResult) string { return v.CloudExadataInfrastructure }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseDbServersResultOutput) DbServers() GetOracleDatabaseDbServersDbServerArrayOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbServersResult) []GetOracleDatabaseDbServersDbServer { return v.DbServers }).(GetOracleDatabaseDbServersDbServerArrayOutput)
}

func (o GetOracleDatabaseDbServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbServersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseDbServersResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbServersResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseDbServersResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbServersResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOracleDatabaseDbServersResultOutput{})
}
