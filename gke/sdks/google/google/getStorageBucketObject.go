// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageBucketObject(ctx *pulumi.Context, args *LookupStorageBucketObjectArgs, opts ...pulumi.InvokeOption) (*LookupStorageBucketObjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupStorageBucketObjectResult
	err = ctx.InvokePackage("google:index/getStorageBucketObject:getStorageBucketObject", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageBucketObject.
type LookupStorageBucketObjectArgs struct {
	Bucket *string `pulumi:"bucket"`
	Id     *string `pulumi:"id"`
	Name   *string `pulumi:"name"`
}

// A collection of values returned by getStorageBucketObject.
type LookupStorageBucketObjectResult struct {
	Bucket              *string                                    `pulumi:"bucket"`
	CacheControl        string                                     `pulumi:"cacheControl"`
	Content             string                                     `pulumi:"content"`
	ContentDisposition  string                                     `pulumi:"contentDisposition"`
	ContentEncoding     string                                     `pulumi:"contentEncoding"`
	ContentLanguage     string                                     `pulumi:"contentLanguage"`
	ContentType         string                                     `pulumi:"contentType"`
	Crc32c              string                                     `pulumi:"crc32c"`
	CustomerEncryptions []GetStorageBucketObjectCustomerEncryption `pulumi:"customerEncryptions"`
	DetectMd5hash       string                                     `pulumi:"detectMd5hash"`
	EventBasedHold      bool                                       `pulumi:"eventBasedHold"`
	Generation          float64                                    `pulumi:"generation"`
	Id                  string                                     `pulumi:"id"`
	KmsKeyName          string                                     `pulumi:"kmsKeyName"`
	Md5hash             string                                     `pulumi:"md5hash"`
	MediaLink           string                                     `pulumi:"mediaLink"`
	Metadata            map[string]string                          `pulumi:"metadata"`
	Name                *string                                    `pulumi:"name"`
	OutputName          string                                     `pulumi:"outputName"`
	Retentions          []GetStorageBucketObjectRetention          `pulumi:"retentions"`
	SelfLink            string                                     `pulumi:"selfLink"`
	Source              string                                     `pulumi:"source"`
	StorageClass        string                                     `pulumi:"storageClass"`
	TemporaryHold       bool                                       `pulumi:"temporaryHold"`
}

func LookupStorageBucketObjectOutput(ctx *pulumi.Context, args LookupStorageBucketObjectOutputArgs, opts ...pulumi.InvokeOption) LookupStorageBucketObjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStorageBucketObjectResultOutput, error) {
			args := v.(LookupStorageBucketObjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupStorageBucketObjectResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getStorageBucketObject:getStorageBucketObject", args, LookupStorageBucketObjectResultOutput{}, options).(LookupStorageBucketObjectResultOutput), nil
		}).(LookupStorageBucketObjectResultOutput)
}

// A collection of arguments for invoking getStorageBucketObject.
type LookupStorageBucketObjectOutputArgs struct {
	Bucket pulumi.StringPtrInput `pulumi:"bucket"`
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupStorageBucketObjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageBucketObjectArgs)(nil)).Elem()
}

// A collection of values returned by getStorageBucketObject.
type LookupStorageBucketObjectResultOutput struct{ *pulumi.OutputState }

func (LookupStorageBucketObjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageBucketObjectResult)(nil)).Elem()
}

func (o LookupStorageBucketObjectResultOutput) ToLookupStorageBucketObjectResultOutput() LookupStorageBucketObjectResultOutput {
	return o
}

func (o LookupStorageBucketObjectResultOutput) ToLookupStorageBucketObjectResultOutputWithContext(ctx context.Context) LookupStorageBucketObjectResultOutput {
	return o
}

func (o LookupStorageBucketObjectResultOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) *string { return v.Bucket }).(pulumi.StringPtrOutput)
}

func (o LookupStorageBucketObjectResultOutput) CacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.CacheControl }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.Content }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.ContentDisposition }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) ContentEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.ContentEncoding }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) ContentLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.ContentLanguage }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) Crc32c() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.Crc32c }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) CustomerEncryptions() GetStorageBucketObjectCustomerEncryptionArrayOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) []GetStorageBucketObjectCustomerEncryption {
		return v.CustomerEncryptions
	}).(GetStorageBucketObjectCustomerEncryptionArrayOutput)
}

func (o LookupStorageBucketObjectResultOutput) DetectMd5hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.DetectMd5hash }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) EventBasedHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) bool { return v.EventBasedHold }).(pulumi.BoolOutput)
}

func (o LookupStorageBucketObjectResultOutput) Generation() pulumi.Float64Output {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) float64 { return v.Generation }).(pulumi.Float64Output)
}

func (o LookupStorageBucketObjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) Md5hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.Md5hash }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.MediaLink }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupStorageBucketObjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupStorageBucketObjectResultOutput) OutputName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.OutputName }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) Retentions() GetStorageBucketObjectRetentionArrayOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) []GetStorageBucketObjectRetention { return v.Retentions }).(GetStorageBucketObjectRetentionArrayOutput)
}

func (o LookupStorageBucketObjectResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.Source }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

func (o LookupStorageBucketObjectResultOutput) TemporaryHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageBucketObjectResult) bool { return v.TemporaryHold }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageBucketObjectResultOutput{})
}
