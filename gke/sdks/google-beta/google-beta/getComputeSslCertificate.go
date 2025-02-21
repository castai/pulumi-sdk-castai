// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeSslCertificate(ctx *pulumi.Context, args *LookupComputeSslCertificateArgs, opts ...pulumi.InvokeOption) (*LookupComputeSslCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeSslCertificateResult
	err = ctx.InvokePackage("google-beta:index/getComputeSslCertificate:getComputeSslCertificate", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeSslCertificate.
type LookupComputeSslCertificateArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getComputeSslCertificate.
type LookupComputeSslCertificateResult struct {
	Certificate       string  `pulumi:"certificate"`
	CertificateId     float64 `pulumi:"certificateId"`
	CreationTimestamp string  `pulumi:"creationTimestamp"`
	Description       string  `pulumi:"description"`
	ExpireTime        string  `pulumi:"expireTime"`
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	NamePrefix        string  `pulumi:"namePrefix"`
	PrivateKey        string  `pulumi:"privateKey"`
	Project           *string `pulumi:"project"`
	SelfLink          string  `pulumi:"selfLink"`
}

func LookupComputeSslCertificateOutput(ctx *pulumi.Context, args LookupComputeSslCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupComputeSslCertificateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeSslCertificateResultOutput, error) {
			args := v.(LookupComputeSslCertificateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeSslCertificateResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeSslCertificate:getComputeSslCertificate", args, LookupComputeSslCertificateResultOutput{}, options).(LookupComputeSslCertificateResultOutput), nil
		}).(LookupComputeSslCertificateResultOutput)
}

// A collection of arguments for invoking getComputeSslCertificate.
type LookupComputeSslCertificateOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupComputeSslCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeSslCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getComputeSslCertificate.
type LookupComputeSslCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupComputeSslCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeSslCertificateResult)(nil)).Elem()
}

func (o LookupComputeSslCertificateResultOutput) ToLookupComputeSslCertificateResultOutput() LookupComputeSslCertificateResultOutput {
	return o
}

func (o LookupComputeSslCertificateResultOutput) ToLookupComputeSslCertificateResultOutputWithContext(ctx context.Context) LookupComputeSslCertificateResultOutput {
	return o
}

func (o LookupComputeSslCertificateResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) CertificateId() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) float64 { return v.CertificateId }).(pulumi.Float64Output)
}

func (o LookupComputeSslCertificateResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.NamePrefix }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o LookupComputeSslCertificateResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeSslCertificateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSslCertificateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeSslCertificateResultOutput{})
}
