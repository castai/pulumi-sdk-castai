// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetContainerRegistryRepository(ctx *pulumi.Context, args *GetContainerRegistryRepositoryArgs, opts ...pulumi.InvokeOption) (*GetContainerRegistryRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetContainerRegistryRepositoryResult
	err = ctx.InvokePackage("google:index/getContainerRegistryRepository:getContainerRegistryRepository", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRegistryRepository.
type GetContainerRegistryRepositoryArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getContainerRegistryRepository.
type GetContainerRegistryRepositoryResult struct {
	Id            string  `pulumi:"id"`
	Project       string  `pulumi:"project"`
	Region        *string `pulumi:"region"`
	RepositoryUrl string  `pulumi:"repositoryUrl"`
}

func GetContainerRegistryRepositoryOutput(ctx *pulumi.Context, args GetContainerRegistryRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetContainerRegistryRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetContainerRegistryRepositoryResultOutput, error) {
			args := v.(GetContainerRegistryRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetContainerRegistryRepositoryResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getContainerRegistryRepository:getContainerRegistryRepository", args, GetContainerRegistryRepositoryResultOutput{}, options).(GetContainerRegistryRepositoryResultOutput), nil
		}).(GetContainerRegistryRepositoryResultOutput)
}

// A collection of arguments for invoking getContainerRegistryRepository.
type GetContainerRegistryRepositoryOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (GetContainerRegistryRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerRegistryRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRegistryRepository.
type GetContainerRegistryRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetContainerRegistryRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerRegistryRepositoryResult)(nil)).Elem()
}

func (o GetContainerRegistryRepositoryResultOutput) ToGetContainerRegistryRepositoryResultOutput() GetContainerRegistryRepositoryResultOutput {
	return o
}

func (o GetContainerRegistryRepositoryResultOutput) ToGetContainerRegistryRepositoryResultOutputWithContext(ctx context.Context) GetContainerRegistryRepositoryResultOutput {
	return o
}

func (o GetContainerRegistryRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetContainerRegistryRepositoryResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryRepositoryResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetContainerRegistryRepositoryResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerRegistryRepositoryResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetContainerRegistryRepositoryResultOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryRepositoryResult) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerRegistryRepositoryResultOutput{})
}
