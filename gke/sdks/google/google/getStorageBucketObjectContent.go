// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStorageBucketObjectContent(ctx *pulumi.Context, args *GetStorageBucketObjectContentArgs, opts ...pulumi.InvokeOption) (*GetStorageBucketObjectContentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetStorageBucketObjectContentResult
	err = ctx.InvokePackage("google:index/getStorageBucketObjectContent:getStorageBucketObjectContent", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageBucketObjectContent.
type GetStorageBucketObjectContentArgs struct {
	Bucket  string  `pulumi:"bucket"`
	Content *string `pulumi:"content"`
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
}

// A collection of values returned by getStorageBucketObjectContent.
type GetStorageBucketObjectContentResult struct {
	Bucket              string                                            `pulumi:"bucket"`
	CacheControl        string                                            `pulumi:"cacheControl"`
	Content             string                                            `pulumi:"content"`
	ContentDisposition  string                                            `pulumi:"contentDisposition"`
	ContentEncoding     string                                            `pulumi:"contentEncoding"`
	ContentLanguage     string                                            `pulumi:"contentLanguage"`
	ContentType         string                                            `pulumi:"contentType"`
	Crc32c              string                                            `pulumi:"crc32c"`
	CustomerEncryptions []GetStorageBucketObjectContentCustomerEncryption `pulumi:"customerEncryptions"`
	DetectMd5hash       string                                            `pulumi:"detectMd5hash"`
	EventBasedHold      bool                                              `pulumi:"eventBasedHold"`
	Generation          float64                                           `pulumi:"generation"`
	Id                  string                                            `pulumi:"id"`
	KmsKeyName          string                                            `pulumi:"kmsKeyName"`
	Md5hash             string                                            `pulumi:"md5hash"`
	MediaLink           string                                            `pulumi:"mediaLink"`
	Metadata            map[string]string                                 `pulumi:"metadata"`
	Name                string                                            `pulumi:"name"`
	OutputName          string                                            `pulumi:"outputName"`
	Retentions          []GetStorageBucketObjectContentRetention          `pulumi:"retentions"`
	SelfLink            string                                            `pulumi:"selfLink"`
	Source              string                                            `pulumi:"source"`
	StorageClass        string                                            `pulumi:"storageClass"`
	TemporaryHold       bool                                              `pulumi:"temporaryHold"`
}

func GetStorageBucketObjectContentOutput(ctx *pulumi.Context, args GetStorageBucketObjectContentOutputArgs, opts ...pulumi.InvokeOption) GetStorageBucketObjectContentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetStorageBucketObjectContentResultOutput, error) {
			args := v.(GetStorageBucketObjectContentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetStorageBucketObjectContentResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getStorageBucketObjectContent:getStorageBucketObjectContent", args, GetStorageBucketObjectContentResultOutput{}, options).(GetStorageBucketObjectContentResultOutput), nil
		}).(GetStorageBucketObjectContentResultOutput)
}

// A collection of arguments for invoking getStorageBucketObjectContent.
type GetStorageBucketObjectContentOutputArgs struct {
	Bucket  pulumi.StringInput    `pulumi:"bucket"`
	Content pulumi.StringPtrInput `pulumi:"content"`
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
}

func (GetStorageBucketObjectContentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStorageBucketObjectContentArgs)(nil)).Elem()
}

// A collection of values returned by getStorageBucketObjectContent.
type GetStorageBucketObjectContentResultOutput struct{ *pulumi.OutputState }

func (GetStorageBucketObjectContentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStorageBucketObjectContentResult)(nil)).Elem()
}

func (o GetStorageBucketObjectContentResultOutput) ToGetStorageBucketObjectContentResultOutput() GetStorageBucketObjectContentResultOutput {
	return o
}

func (o GetStorageBucketObjectContentResultOutput) ToGetStorageBucketObjectContentResultOutputWithContext(ctx context.Context) GetStorageBucketObjectContentResultOutput {
	return o
}

func (o GetStorageBucketObjectContentResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.Bucket }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) CacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.CacheControl }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.Content }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.ContentDisposition }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) ContentEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.ContentEncoding }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) ContentLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.ContentLanguage }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Crc32c() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.Crc32c }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) CustomerEncryptions() GetStorageBucketObjectContentCustomerEncryptionArrayOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) []GetStorageBucketObjectContentCustomerEncryption {
		return v.CustomerEncryptions
	}).(GetStorageBucketObjectContentCustomerEncryptionArrayOutput)
}

func (o GetStorageBucketObjectContentResultOutput) DetectMd5hash() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.DetectMd5hash }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) EventBasedHold() pulumi.BoolOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) bool { return v.EventBasedHold }).(pulumi.BoolOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Generation() pulumi.Float64Output {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) float64 { return v.Generation }).(pulumi.Float64Output)
}

func (o GetStorageBucketObjectContentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Md5hash() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.Md5hash }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.MediaLink }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) OutputName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.OutputName }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Retentions() GetStorageBucketObjectContentRetentionArrayOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) []GetStorageBucketObjectContentRetention {
		return v.Retentions
	}).(GetStorageBucketObjectContentRetentionArrayOutput)
}

func (o GetStorageBucketObjectContentResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.Source }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

func (o GetStorageBucketObjectContentResultOutput) TemporaryHold() pulumi.BoolOutput {
	return o.ApplyT(func(v GetStorageBucketObjectContentResult) bool { return v.TemporaryHold }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStorageBucketObjectContentResultOutput{})
}
