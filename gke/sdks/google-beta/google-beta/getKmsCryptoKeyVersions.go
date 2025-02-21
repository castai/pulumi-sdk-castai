// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKmsCryptoKeyVersions(ctx *pulumi.Context, args *GetKmsCryptoKeyVersionsArgs, opts ...pulumi.InvokeOption) (*GetKmsCryptoKeyVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetKmsCryptoKeyVersionsResult
	err = ctx.InvokePackage("google-beta:index/getKmsCryptoKeyVersions:getKmsCryptoKeyVersions", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKmsCryptoKeyVersions.
type GetKmsCryptoKeyVersionsArgs struct {
	CryptoKey string  `pulumi:"cryptoKey"`
	Filter    *string `pulumi:"filter"`
	Id        *string `pulumi:"id"`
}

// A collection of values returned by getKmsCryptoKeyVersions.
type GetKmsCryptoKeyVersionsResult struct {
	CryptoKey  string                             `pulumi:"cryptoKey"`
	Filter     *string                            `pulumi:"filter"`
	Id         string                             `pulumi:"id"`
	PublicKeys []GetKmsCryptoKeyVersionsPublicKey `pulumi:"publicKeys"`
	Versions   []GetKmsCryptoKeyVersionsVersion   `pulumi:"versions"`
}

func GetKmsCryptoKeyVersionsOutput(ctx *pulumi.Context, args GetKmsCryptoKeyVersionsOutputArgs, opts ...pulumi.InvokeOption) GetKmsCryptoKeyVersionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKmsCryptoKeyVersionsResultOutput, error) {
			args := v.(GetKmsCryptoKeyVersionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetKmsCryptoKeyVersionsResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getKmsCryptoKeyVersions:getKmsCryptoKeyVersions", args, GetKmsCryptoKeyVersionsResultOutput{}, options).(GetKmsCryptoKeyVersionsResultOutput), nil
		}).(GetKmsCryptoKeyVersionsResultOutput)
}

// A collection of arguments for invoking getKmsCryptoKeyVersions.
type GetKmsCryptoKeyVersionsOutputArgs struct {
	CryptoKey pulumi.StringInput    `pulumi:"cryptoKey"`
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Id        pulumi.StringPtrInput `pulumi:"id"`
}

func (GetKmsCryptoKeyVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKmsCryptoKeyVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getKmsCryptoKeyVersions.
type GetKmsCryptoKeyVersionsResultOutput struct{ *pulumi.OutputState }

func (GetKmsCryptoKeyVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKmsCryptoKeyVersionsResult)(nil)).Elem()
}

func (o GetKmsCryptoKeyVersionsResultOutput) ToGetKmsCryptoKeyVersionsResultOutput() GetKmsCryptoKeyVersionsResultOutput {
	return o
}

func (o GetKmsCryptoKeyVersionsResultOutput) ToGetKmsCryptoKeyVersionsResultOutputWithContext(ctx context.Context) GetKmsCryptoKeyVersionsResultOutput {
	return o
}

func (o GetKmsCryptoKeyVersionsResultOutput) CryptoKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyVersionsResult) string { return v.CryptoKey }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyVersionsResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyVersionsResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetKmsCryptoKeyVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyVersionsResultOutput) PublicKeys() GetKmsCryptoKeyVersionsPublicKeyArrayOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyVersionsResult) []GetKmsCryptoKeyVersionsPublicKey { return v.PublicKeys }).(GetKmsCryptoKeyVersionsPublicKeyArrayOutput)
}

func (o GetKmsCryptoKeyVersionsResultOutput) Versions() GetKmsCryptoKeyVersionsVersionArrayOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyVersionsResult) []GetKmsCryptoKeyVersionsVersion { return v.Versions }).(GetKmsCryptoKeyVersionsVersionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKmsCryptoKeyVersionsResultOutput{})
}
