// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSqlDatabases(ctx *pulumi.Context, args *GetSqlDatabasesArgs, opts ...pulumi.InvokeOption) (*GetSqlDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetSqlDatabasesResult
	err = ctx.InvokePackage("google:index/getSqlDatabases:getSqlDatabases", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSqlDatabases.
type GetSqlDatabasesArgs struct {
	Id       *string `pulumi:"id"`
	Instance string  `pulumi:"instance"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getSqlDatabases.
type GetSqlDatabasesResult struct {
	Databases []GetSqlDatabasesDatabase `pulumi:"databases"`
	Id        string                    `pulumi:"id"`
	Instance  string                    `pulumi:"instance"`
	Project   *string                   `pulumi:"project"`
}

func GetSqlDatabasesOutput(ctx *pulumi.Context, args GetSqlDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetSqlDatabasesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetSqlDatabasesResultOutput, error) {
			args := v.(GetSqlDatabasesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetSqlDatabasesResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getSqlDatabases:getSqlDatabases", args, GetSqlDatabasesResultOutput{}, options).(GetSqlDatabasesResultOutput), nil
		}).(GetSqlDatabasesResultOutput)
}

// A collection of arguments for invoking getSqlDatabases.
type GetSqlDatabasesOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Instance pulumi.StringInput    `pulumi:"instance"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (GetSqlDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSqlDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getSqlDatabases.
type GetSqlDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetSqlDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSqlDatabasesResult)(nil)).Elem()
}

func (o GetSqlDatabasesResultOutput) ToGetSqlDatabasesResultOutput() GetSqlDatabasesResultOutput {
	return o
}

func (o GetSqlDatabasesResultOutput) ToGetSqlDatabasesResultOutputWithContext(ctx context.Context) GetSqlDatabasesResultOutput {
	return o
}

func (o GetSqlDatabasesResultOutput) Databases() GetSqlDatabasesDatabaseArrayOutput {
	return o.ApplyT(func(v GetSqlDatabasesResult) []GetSqlDatabasesDatabase { return v.Databases }).(GetSqlDatabasesDatabaseArrayOutput)
}

func (o GetSqlDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSqlDatabasesResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlDatabasesResult) string { return v.Instance }).(pulumi.StringOutput)
}

func (o GetSqlDatabasesResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSqlDatabasesResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSqlDatabasesResultOutput{})
}
