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

type KmsAutokeyConfig struct {
	pulumi.CustomResourceState

	// The folder for which to retrieve config.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// The target key project for a given folder where KMS Autokey will provision a CryptoKey for any new KeyHandle the
	// Developer creates. Should have the form 'projects/<project_id_or_number>'.
	KeyProject         pulumi.StringPtrOutput            `pulumi:"keyProject"`
	KmsAutokeyConfigId pulumi.StringOutput               `pulumi:"kmsAutokeyConfigId"`
	Timeouts           KmsAutokeyConfigTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewKmsAutokeyConfig registers a new resource with the given unique name, arguments, and options.
func NewKmsAutokeyConfig(ctx *pulumi.Context,
	name string, args *KmsAutokeyConfigArgs, opts ...pulumi.ResourceOption) (*KmsAutokeyConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource KmsAutokeyConfig
	err = ctx.RegisterPackageResource("google-beta:index/kmsAutokeyConfig:KmsAutokeyConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsAutokeyConfig gets an existing KmsAutokeyConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsAutokeyConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsAutokeyConfigState, opts ...pulumi.ResourceOption) (*KmsAutokeyConfig, error) {
	var resource KmsAutokeyConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/kmsAutokeyConfig:KmsAutokeyConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsAutokeyConfig resources.
type kmsAutokeyConfigState struct {
	// The folder for which to retrieve config.
	Folder *string `pulumi:"folder"`
	// The target key project for a given folder where KMS Autokey will provision a CryptoKey for any new KeyHandle the
	// Developer creates. Should have the form 'projects/<project_id_or_number>'.
	KeyProject         *string                   `pulumi:"keyProject"`
	KmsAutokeyConfigId *string                   `pulumi:"kmsAutokeyConfigId"`
	Timeouts           *KmsAutokeyConfigTimeouts `pulumi:"timeouts"`
}

type KmsAutokeyConfigState struct {
	// The folder for which to retrieve config.
	Folder pulumi.StringPtrInput
	// The target key project for a given folder where KMS Autokey will provision a CryptoKey for any new KeyHandle the
	// Developer creates. Should have the form 'projects/<project_id_or_number>'.
	KeyProject         pulumi.StringPtrInput
	KmsAutokeyConfigId pulumi.StringPtrInput
	Timeouts           KmsAutokeyConfigTimeoutsPtrInput
}

func (KmsAutokeyConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAutokeyConfigState)(nil)).Elem()
}

type kmsAutokeyConfigArgs struct {
	// The folder for which to retrieve config.
	Folder string `pulumi:"folder"`
	// The target key project for a given folder where KMS Autokey will provision a CryptoKey for any new KeyHandle the
	// Developer creates. Should have the form 'projects/<project_id_or_number>'.
	KeyProject         *string                   `pulumi:"keyProject"`
	KmsAutokeyConfigId *string                   `pulumi:"kmsAutokeyConfigId"`
	Timeouts           *KmsAutokeyConfigTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a KmsAutokeyConfig resource.
type KmsAutokeyConfigArgs struct {
	// The folder for which to retrieve config.
	Folder pulumi.StringInput
	// The target key project for a given folder where KMS Autokey will provision a CryptoKey for any new KeyHandle the
	// Developer creates. Should have the form 'projects/<project_id_or_number>'.
	KeyProject         pulumi.StringPtrInput
	KmsAutokeyConfigId pulumi.StringPtrInput
	Timeouts           KmsAutokeyConfigTimeoutsPtrInput
}

func (KmsAutokeyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAutokeyConfigArgs)(nil)).Elem()
}

type KmsAutokeyConfigInput interface {
	pulumi.Input

	ToKmsAutokeyConfigOutput() KmsAutokeyConfigOutput
	ToKmsAutokeyConfigOutputWithContext(ctx context.Context) KmsAutokeyConfigOutput
}

func (*KmsAutokeyConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAutokeyConfig)(nil)).Elem()
}

func (i *KmsAutokeyConfig) ToKmsAutokeyConfigOutput() KmsAutokeyConfigOutput {
	return i.ToKmsAutokeyConfigOutputWithContext(context.Background())
}

func (i *KmsAutokeyConfig) ToKmsAutokeyConfigOutputWithContext(ctx context.Context) KmsAutokeyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAutokeyConfigOutput)
}

type KmsAutokeyConfigOutput struct{ *pulumi.OutputState }

func (KmsAutokeyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAutokeyConfig)(nil)).Elem()
}

func (o KmsAutokeyConfigOutput) ToKmsAutokeyConfigOutput() KmsAutokeyConfigOutput {
	return o
}

func (o KmsAutokeyConfigOutput) ToKmsAutokeyConfigOutputWithContext(ctx context.Context) KmsAutokeyConfigOutput {
	return o
}

// The folder for which to retrieve config.
func (o KmsAutokeyConfigOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAutokeyConfig) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

// The target key project for a given folder where KMS Autokey will provision a CryptoKey for any new KeyHandle the
// Developer creates. Should have the form 'projects/<project_id_or_number>'.
func (o KmsAutokeyConfigOutput) KeyProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KmsAutokeyConfig) pulumi.StringPtrOutput { return v.KeyProject }).(pulumi.StringPtrOutput)
}

func (o KmsAutokeyConfigOutput) KmsAutokeyConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAutokeyConfig) pulumi.StringOutput { return v.KmsAutokeyConfigId }).(pulumi.StringOutput)
}

func (o KmsAutokeyConfigOutput) Timeouts() KmsAutokeyConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *KmsAutokeyConfig) KmsAutokeyConfigTimeoutsPtrOutput { return v.Timeouts }).(KmsAutokeyConfigTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAutokeyConfigInput)(nil)).Elem(), &KmsAutokeyConfig{})
	pulumi.RegisterOutputType(KmsAutokeyConfigOutput{})
}
