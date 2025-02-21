// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeBackendBucket(ctx *pulumi.Context, args *LookupComputeBackendBucketArgs, opts ...pulumi.InvokeOption) (*LookupComputeBackendBucketResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeBackendBucketResult
	err = ctx.InvokePackage("google:index/getComputeBackendBucket:getComputeBackendBucket", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeBackendBucket.
type LookupComputeBackendBucketArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getComputeBackendBucket.
type LookupComputeBackendBucketResult struct {
	BucketName            string                             `pulumi:"bucketName"`
	CdnPolicies           []GetComputeBackendBucketCdnPolicy `pulumi:"cdnPolicies"`
	CompressionMode       string                             `pulumi:"compressionMode"`
	CreationTimestamp     string                             `pulumi:"creationTimestamp"`
	CustomResponseHeaders []string                           `pulumi:"customResponseHeaders"`
	Description           string                             `pulumi:"description"`
	EdgeSecurityPolicy    string                             `pulumi:"edgeSecurityPolicy"`
	EnableCdn             bool                               `pulumi:"enableCdn"`
	Id                    string                             `pulumi:"id"`
	Name                  string                             `pulumi:"name"`
	Project               *string                            `pulumi:"project"`
	SelfLink              string                             `pulumi:"selfLink"`
}

func LookupComputeBackendBucketOutput(ctx *pulumi.Context, args LookupComputeBackendBucketOutputArgs, opts ...pulumi.InvokeOption) LookupComputeBackendBucketResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeBackendBucketResultOutput, error) {
			args := v.(LookupComputeBackendBucketArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeBackendBucketResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getComputeBackendBucket:getComputeBackendBucket", args, LookupComputeBackendBucketResultOutput{}, options).(LookupComputeBackendBucketResultOutput), nil
		}).(LookupComputeBackendBucketResultOutput)
}

// A collection of arguments for invoking getComputeBackendBucket.
type LookupComputeBackendBucketOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupComputeBackendBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeBackendBucketArgs)(nil)).Elem()
}

// A collection of values returned by getComputeBackendBucket.
type LookupComputeBackendBucketResultOutput struct{ *pulumi.OutputState }

func (LookupComputeBackendBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeBackendBucketResult)(nil)).Elem()
}

func (o LookupComputeBackendBucketResultOutput) ToLookupComputeBackendBucketResultOutput() LookupComputeBackendBucketResultOutput {
	return o
}

func (o LookupComputeBackendBucketResultOutput) ToLookupComputeBackendBucketResultOutputWithContext(ctx context.Context) LookupComputeBackendBucketResultOutput {
	return o
}

func (o LookupComputeBackendBucketResultOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.BucketName }).(pulumi.StringOutput)
}

func (o LookupComputeBackendBucketResultOutput) CdnPolicies() GetComputeBackendBucketCdnPolicyArrayOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) []GetComputeBackendBucketCdnPolicy { return v.CdnPolicies }).(GetComputeBackendBucketCdnPolicyArrayOutput)
}

func (o LookupComputeBackendBucketResultOutput) CompressionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.CompressionMode }).(pulumi.StringOutput)
}

func (o LookupComputeBackendBucketResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o LookupComputeBackendBucketResultOutput) CustomResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) []string { return v.CustomResponseHeaders }).(pulumi.StringArrayOutput)
}

func (o LookupComputeBackendBucketResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeBackendBucketResultOutput) EdgeSecurityPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.EdgeSecurityPolicy }).(pulumi.StringOutput)
}

func (o LookupComputeBackendBucketResultOutput) EnableCdn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) bool { return v.EnableCdn }).(pulumi.BoolOutput)
}

func (o LookupComputeBackendBucketResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeBackendBucketResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeBackendBucketResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeBackendBucketResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeBackendBucketResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeBackendBucketResultOutput{})
}
