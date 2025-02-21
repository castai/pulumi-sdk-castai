// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KmsSecretCiphertext struct {
	pulumi.CustomResourceState

	// The additional authenticated data used for integrity checks during encryption and decryption.
	AdditionalAuthenticatedData pulumi.StringPtrOutput `pulumi:"additionalAuthenticatedData"`
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext pulumi.StringOutput `pulumi:"ciphertext"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
	CryptoKey             pulumi.StringOutput `pulumi:"cryptoKey"`
	KmsSecretCiphertextId pulumi.StringOutput `pulumi:"kmsSecretCiphertextId"`
	// The plaintext to be encrypted.
	Plaintext pulumi.StringOutput                  `pulumi:"plaintext"`
	Timeouts  KmsSecretCiphertextTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewKmsSecretCiphertext registers a new resource with the given unique name, arguments, and options.
func NewKmsSecretCiphertext(ctx *pulumi.Context,
	name string, args *KmsSecretCiphertextArgs, opts ...pulumi.ResourceOption) (*KmsSecretCiphertext, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptoKey == nil {
		return nil, errors.New("invalid value for required argument 'CryptoKey'")
	}
	if args.Plaintext == nil {
		return nil, errors.New("invalid value for required argument 'Plaintext'")
	}
	if args.AdditionalAuthenticatedData != nil {
		args.AdditionalAuthenticatedData = pulumi.ToSecret(args.AdditionalAuthenticatedData).(pulumi.StringPtrInput)
	}
	if args.Plaintext != nil {
		args.Plaintext = pulumi.ToSecret(args.Plaintext).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"additionalAuthenticatedData",
		"plaintext",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource KmsSecretCiphertext
	err = ctx.RegisterPackageResource("google-beta:index/kmsSecretCiphertext:KmsSecretCiphertext", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsSecretCiphertext gets an existing KmsSecretCiphertext resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsSecretCiphertext(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsSecretCiphertextState, opts ...pulumi.ResourceOption) (*KmsSecretCiphertext, error) {
	var resource KmsSecretCiphertext
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/kmsSecretCiphertext:KmsSecretCiphertext", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsSecretCiphertext resources.
type kmsSecretCiphertextState struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	AdditionalAuthenticatedData *string `pulumi:"additionalAuthenticatedData"`
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext *string `pulumi:"ciphertext"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
	CryptoKey             *string `pulumi:"cryptoKey"`
	KmsSecretCiphertextId *string `pulumi:"kmsSecretCiphertextId"`
	// The plaintext to be encrypted.
	Plaintext *string                      `pulumi:"plaintext"`
	Timeouts  *KmsSecretCiphertextTimeouts `pulumi:"timeouts"`
}

type KmsSecretCiphertextState struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	AdditionalAuthenticatedData pulumi.StringPtrInput
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext pulumi.StringPtrInput
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
	CryptoKey             pulumi.StringPtrInput
	KmsSecretCiphertextId pulumi.StringPtrInput
	// The plaintext to be encrypted.
	Plaintext pulumi.StringPtrInput
	Timeouts  KmsSecretCiphertextTimeoutsPtrInput
}

func (KmsSecretCiphertextState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSecretCiphertextState)(nil)).Elem()
}

type kmsSecretCiphertextArgs struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	AdditionalAuthenticatedData *string `pulumi:"additionalAuthenticatedData"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
	CryptoKey             string  `pulumi:"cryptoKey"`
	KmsSecretCiphertextId *string `pulumi:"kmsSecretCiphertextId"`
	// The plaintext to be encrypted.
	Plaintext string                       `pulumi:"plaintext"`
	Timeouts  *KmsSecretCiphertextTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a KmsSecretCiphertext resource.
type KmsSecretCiphertextArgs struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	AdditionalAuthenticatedData pulumi.StringPtrInput
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
	CryptoKey             pulumi.StringInput
	KmsSecretCiphertextId pulumi.StringPtrInput
	// The plaintext to be encrypted.
	Plaintext pulumi.StringInput
	Timeouts  KmsSecretCiphertextTimeoutsPtrInput
}

func (KmsSecretCiphertextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSecretCiphertextArgs)(nil)).Elem()
}

type KmsSecretCiphertextInput interface {
	pulumi.Input

	ToKmsSecretCiphertextOutput() KmsSecretCiphertextOutput
	ToKmsSecretCiphertextOutputWithContext(ctx context.Context) KmsSecretCiphertextOutput
}

func (*KmsSecretCiphertext) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSecretCiphertext)(nil)).Elem()
}

func (i *KmsSecretCiphertext) ToKmsSecretCiphertextOutput() KmsSecretCiphertextOutput {
	return i.ToKmsSecretCiphertextOutputWithContext(context.Background())
}

func (i *KmsSecretCiphertext) ToKmsSecretCiphertextOutputWithContext(ctx context.Context) KmsSecretCiphertextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSecretCiphertextOutput)
}

type KmsSecretCiphertextOutput struct{ *pulumi.OutputState }

func (KmsSecretCiphertextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSecretCiphertext)(nil)).Elem()
}

func (o KmsSecretCiphertextOutput) ToKmsSecretCiphertextOutput() KmsSecretCiphertextOutput {
	return o
}

func (o KmsSecretCiphertextOutput) ToKmsSecretCiphertextOutputWithContext(ctx context.Context) KmsSecretCiphertextOutput {
	return o
}

// The additional authenticated data used for integrity checks during encryption and decryption.
func (o KmsSecretCiphertextOutput) AdditionalAuthenticatedData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KmsSecretCiphertext) pulumi.StringPtrOutput { return v.AdditionalAuthenticatedData }).(pulumi.StringPtrOutput)
}

// Contains the result of encrypting the provided plaintext, encoded in base64.
func (o KmsSecretCiphertextOutput) Ciphertext() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSecretCiphertext) pulumi.StringOutput { return v.Ciphertext }).(pulumi.StringOutput)
}

// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
// ”projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}”
func (o KmsSecretCiphertextOutput) CryptoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSecretCiphertext) pulumi.StringOutput { return v.CryptoKey }).(pulumi.StringOutput)
}

func (o KmsSecretCiphertextOutput) KmsSecretCiphertextId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSecretCiphertext) pulumi.StringOutput { return v.KmsSecretCiphertextId }).(pulumi.StringOutput)
}

// The plaintext to be encrypted.
func (o KmsSecretCiphertextOutput) Plaintext() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSecretCiphertext) pulumi.StringOutput { return v.Plaintext }).(pulumi.StringOutput)
}

func (o KmsSecretCiphertextOutput) Timeouts() KmsSecretCiphertextTimeoutsPtrOutput {
	return o.ApplyT(func(v *KmsSecretCiphertext) KmsSecretCiphertextTimeoutsPtrOutput { return v.Timeouts }).(KmsSecretCiphertextTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSecretCiphertextInput)(nil)).Elem(), &KmsSecretCiphertext{})
	pulumi.RegisterOutputType(KmsSecretCiphertextOutput{})
}
