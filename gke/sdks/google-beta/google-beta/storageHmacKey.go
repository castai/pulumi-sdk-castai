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

type StorageHmacKey struct {
	pulumi.CustomResourceState

	// The access ID of the HMAC Key.
	AccessId pulumi.StringOutput `pulumi:"accessId"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// HMAC secret key material.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// The email address of the key's associated service account.
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE",
	// "INACTIVE"]
	State            pulumi.StringPtrOutput `pulumi:"state"`
	StorageHmacKeyId pulumi.StringOutput    `pulumi:"storageHmacKeyId"`
	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated pulumi.StringOutput             `pulumi:"timeCreated"`
	Timeouts    StorageHmacKeyTimeoutsPtrOutput `pulumi:"timeouts"`
	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewStorageHmacKey registers a new resource with the given unique name, arguments, and options.
func NewStorageHmacKey(ctx *pulumi.Context,
	name string, args *StorageHmacKeyArgs, opts ...pulumi.ResourceOption) (*StorageHmacKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountEmail == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountEmail'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageHmacKey
	err = ctx.RegisterPackageResource("google-beta:index/storageHmacKey:StorageHmacKey", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageHmacKey gets an existing StorageHmacKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageHmacKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageHmacKeyState, opts ...pulumi.ResourceOption) (*StorageHmacKey, error) {
	var resource StorageHmacKey
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/storageHmacKey:StorageHmacKey", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageHmacKey resources.
type storageHmacKeyState struct {
	// The access ID of the HMAC Key.
	AccessId *string `pulumi:"accessId"`
	Project  *string `pulumi:"project"`
	// HMAC secret key material.
	Secret *string `pulumi:"secret"`
	// The email address of the key's associated service account.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE",
	// "INACTIVE"]
	State            *string `pulumi:"state"`
	StorageHmacKeyId *string `pulumi:"storageHmacKeyId"`
	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated *string                 `pulumi:"timeCreated"`
	Timeouts    *StorageHmacKeyTimeouts `pulumi:"timeouts"`
	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated *string `pulumi:"updated"`
}

type StorageHmacKeyState struct {
	// The access ID of the HMAC Key.
	AccessId pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// HMAC secret key material.
	Secret pulumi.StringPtrInput
	// The email address of the key's associated service account.
	ServiceAccountEmail pulumi.StringPtrInput
	// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE",
	// "INACTIVE"]
	State            pulumi.StringPtrInput
	StorageHmacKeyId pulumi.StringPtrInput
	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated pulumi.StringPtrInput
	Timeouts    StorageHmacKeyTimeoutsPtrInput
	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated pulumi.StringPtrInput
}

func (StorageHmacKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageHmacKeyState)(nil)).Elem()
}

type storageHmacKeyArgs struct {
	Project *string `pulumi:"project"`
	// The email address of the key's associated service account.
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE",
	// "INACTIVE"]
	State            *string                 `pulumi:"state"`
	StorageHmacKeyId *string                 `pulumi:"storageHmacKeyId"`
	Timeouts         *StorageHmacKeyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a StorageHmacKey resource.
type StorageHmacKeyArgs struct {
	Project pulumi.StringPtrInput
	// The email address of the key's associated service account.
	ServiceAccountEmail pulumi.StringInput
	// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE",
	// "INACTIVE"]
	State            pulumi.StringPtrInput
	StorageHmacKeyId pulumi.StringPtrInput
	Timeouts         StorageHmacKeyTimeoutsPtrInput
}

func (StorageHmacKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageHmacKeyArgs)(nil)).Elem()
}

type StorageHmacKeyInput interface {
	pulumi.Input

	ToStorageHmacKeyOutput() StorageHmacKeyOutput
	ToStorageHmacKeyOutputWithContext(ctx context.Context) StorageHmacKeyOutput
}

func (*StorageHmacKey) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageHmacKey)(nil)).Elem()
}

func (i *StorageHmacKey) ToStorageHmacKeyOutput() StorageHmacKeyOutput {
	return i.ToStorageHmacKeyOutputWithContext(context.Background())
}

func (i *StorageHmacKey) ToStorageHmacKeyOutputWithContext(ctx context.Context) StorageHmacKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageHmacKeyOutput)
}

type StorageHmacKeyOutput struct{ *pulumi.OutputState }

func (StorageHmacKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageHmacKey)(nil)).Elem()
}

func (o StorageHmacKeyOutput) ToStorageHmacKeyOutput() StorageHmacKeyOutput {
	return o
}

func (o StorageHmacKeyOutput) ToStorageHmacKeyOutputWithContext(ctx context.Context) StorageHmacKeyOutput {
	return o
}

// The access ID of the HMAC Key.
func (o StorageHmacKeyOutput) AccessId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringOutput { return v.AccessId }).(pulumi.StringOutput)
}

func (o StorageHmacKeyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// HMAC secret key material.
func (o StorageHmacKeyOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// The email address of the key's associated service account.
func (o StorageHmacKeyOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringOutput { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE",
// "INACTIVE"]
func (o StorageHmacKeyOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o StorageHmacKeyOutput) StorageHmacKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringOutput { return v.StorageHmacKeyId }).(pulumi.StringOutput)
}

// 'The creation time of the HMAC key in RFC 3339 format. '
func (o StorageHmacKeyOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o StorageHmacKeyOutput) Timeouts() StorageHmacKeyTimeoutsPtrOutput {
	return o.ApplyT(func(v *StorageHmacKey) StorageHmacKeyTimeoutsPtrOutput { return v.Timeouts }).(StorageHmacKeyTimeoutsPtrOutput)
}

// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
func (o StorageHmacKeyOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageHmacKey) pulumi.StringOutput { return v.Updated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageHmacKeyInput)(nil)).Elem(), &StorageHmacKey{})
	pulumi.RegisterOutputType(StorageHmacKeyOutput{})
}
