// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOracleDatabaseDbNodes(ctx *pulumi.Context, args *GetOracleDatabaseDbNodesArgs, opts ...pulumi.InvokeOption) (*GetOracleDatabaseDbNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetOracleDatabaseDbNodesResult
	err = ctx.InvokePackage("google-beta:index/getOracleDatabaseDbNodes:getOracleDatabaseDbNodes", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOracleDatabaseDbNodes.
type GetOracleDatabaseDbNodesArgs struct {
	CloudVmCluster string  `pulumi:"cloudVmCluster"`
	Id             *string `pulumi:"id"`
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
}

// A collection of values returned by getOracleDatabaseDbNodes.
type GetOracleDatabaseDbNodesResult struct {
	CloudVmCluster string                           `pulumi:"cloudVmCluster"`
	DbNodes        []GetOracleDatabaseDbNodesDbNode `pulumi:"dbNodes"`
	Id             string                           `pulumi:"id"`
	Location       string                           `pulumi:"location"`
	Project        *string                          `pulumi:"project"`
}

func GetOracleDatabaseDbNodesOutput(ctx *pulumi.Context, args GetOracleDatabaseDbNodesOutputArgs, opts ...pulumi.InvokeOption) GetOracleDatabaseDbNodesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOracleDatabaseDbNodesResultOutput, error) {
			args := v.(GetOracleDatabaseDbNodesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetOracleDatabaseDbNodesResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getOracleDatabaseDbNodes:getOracleDatabaseDbNodes", args, GetOracleDatabaseDbNodesResultOutput{}, options).(GetOracleDatabaseDbNodesResultOutput), nil
		}).(GetOracleDatabaseDbNodesResultOutput)
}

// A collection of arguments for invoking getOracleDatabaseDbNodes.
type GetOracleDatabaseDbNodesOutputArgs struct {
	CloudVmCluster pulumi.StringInput    `pulumi:"cloudVmCluster"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (GetOracleDatabaseDbNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseDbNodesArgs)(nil)).Elem()
}

// A collection of values returned by getOracleDatabaseDbNodes.
type GetOracleDatabaseDbNodesResultOutput struct{ *pulumi.OutputState }

func (GetOracleDatabaseDbNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseDbNodesResult)(nil)).Elem()
}

func (o GetOracleDatabaseDbNodesResultOutput) ToGetOracleDatabaseDbNodesResultOutput() GetOracleDatabaseDbNodesResultOutput {
	return o
}

func (o GetOracleDatabaseDbNodesResultOutput) ToGetOracleDatabaseDbNodesResultOutputWithContext(ctx context.Context) GetOracleDatabaseDbNodesResultOutput {
	return o
}

func (o GetOracleDatabaseDbNodesResultOutput) CloudVmCluster() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbNodesResult) string { return v.CloudVmCluster }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseDbNodesResultOutput) DbNodes() GetOracleDatabaseDbNodesDbNodeArrayOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbNodesResult) []GetOracleDatabaseDbNodesDbNode { return v.DbNodes }).(GetOracleDatabaseDbNodesDbNodeArrayOutput)
}

func (o GetOracleDatabaseDbNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseDbNodesResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbNodesResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseDbNodesResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOracleDatabaseDbNodesResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOracleDatabaseDbNodesResultOutput{})
}
