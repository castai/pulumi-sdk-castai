// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKmsCryptoKeyLatestVersion(ctx *pulumi.Context, args *GetKmsCryptoKeyLatestVersionArgs, opts ...pulumi.InvokeOption) (*GetKmsCryptoKeyLatestVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetKmsCryptoKeyLatestVersionResult
	err = ctx.InvokePackage("google:index/getKmsCryptoKeyLatestVersion:getKmsCryptoKeyLatestVersion", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKmsCryptoKeyLatestVersion.
type GetKmsCryptoKeyLatestVersionArgs struct {
	CryptoKey string  `pulumi:"cryptoKey"`
	Filter    *string `pulumi:"filter"`
	Id        *string `pulumi:"id"`
}

// A collection of values returned by getKmsCryptoKeyLatestVersion.
type GetKmsCryptoKeyLatestVersionResult struct {
	Algorithm       string                                  `pulumi:"algorithm"`
	CryptoKey       string                                  `pulumi:"cryptoKey"`
	Filter          *string                                 `pulumi:"filter"`
	Id              string                                  `pulumi:"id"`
	Name            string                                  `pulumi:"name"`
	ProtectionLevel string                                  `pulumi:"protectionLevel"`
	PublicKeys      []GetKmsCryptoKeyLatestVersionPublicKey `pulumi:"publicKeys"`
	State           string                                  `pulumi:"state"`
	Version         float64                                 `pulumi:"version"`
}

func GetKmsCryptoKeyLatestVersionOutput(ctx *pulumi.Context, args GetKmsCryptoKeyLatestVersionOutputArgs, opts ...pulumi.InvokeOption) GetKmsCryptoKeyLatestVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKmsCryptoKeyLatestVersionResultOutput, error) {
			args := v.(GetKmsCryptoKeyLatestVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetKmsCryptoKeyLatestVersionResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getKmsCryptoKeyLatestVersion:getKmsCryptoKeyLatestVersion", args, GetKmsCryptoKeyLatestVersionResultOutput{}, options).(GetKmsCryptoKeyLatestVersionResultOutput), nil
		}).(GetKmsCryptoKeyLatestVersionResultOutput)
}

// A collection of arguments for invoking getKmsCryptoKeyLatestVersion.
type GetKmsCryptoKeyLatestVersionOutputArgs struct {
	CryptoKey pulumi.StringInput    `pulumi:"cryptoKey"`
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Id        pulumi.StringPtrInput `pulumi:"id"`
}

func (GetKmsCryptoKeyLatestVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKmsCryptoKeyLatestVersionArgs)(nil)).Elem()
}

// A collection of values returned by getKmsCryptoKeyLatestVersion.
type GetKmsCryptoKeyLatestVersionResultOutput struct{ *pulumi.OutputState }

func (GetKmsCryptoKeyLatestVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKmsCryptoKeyLatestVersionResult)(nil)).Elem()
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) ToGetKmsCryptoKeyLatestVersionResultOutput() GetKmsCryptoKeyLatestVersionResultOutput {
	return o
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) ToGetKmsCryptoKeyLatestVersionResultOutputWithContext(ctx context.Context) GetKmsCryptoKeyLatestVersionResultOutput {
	return o
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) CryptoKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) string { return v.CryptoKey }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) ProtectionLevel() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) string { return v.ProtectionLevel }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) PublicKeys() GetKmsCryptoKeyLatestVersionPublicKeyArrayOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) []GetKmsCryptoKeyLatestVersionPublicKey {
		return v.PublicKeys
	}).(GetKmsCryptoKeyLatestVersionPublicKeyArrayOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetKmsCryptoKeyLatestVersionResultOutput) Version() pulumi.Float64Output {
	return o.ApplyT(func(v GetKmsCryptoKeyLatestVersionResult) float64 { return v.Version }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(GetKmsCryptoKeyLatestVersionResultOutput{})
}
