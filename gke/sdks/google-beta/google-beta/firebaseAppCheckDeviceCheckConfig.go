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

type FirebaseAppCheckDeviceCheckConfig struct {
	pulumi.CustomResourceState

	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                               pulumi.StringOutput `pulumi:"appId"`
	FirebaseAppCheckDeviceCheckConfigId pulumi.StringOutput `pulumi:"firebaseAppCheckDeviceCheckConfigId"`
	// The key identifier of a private key enabled with DeviceCheck, created in your Apple Developer account.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The relative resource name of the DeviceCheck configuration object
	Name pulumi.StringOutput `pulumi:"name"`
	// The contents of the private key (.p8) file associated with the key specified by keyId.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// Whether the privateKey field was previously set. Since App Check will never return the privateKey field, this field is
	// the only way to find out whether it was previously set.
	PrivateKeySet pulumi.BoolOutput                                  `pulumi:"privateKeySet"`
	Project       pulumi.StringOutput                                `pulumi:"project"`
	Timeouts      FirebaseAppCheckDeviceCheckConfigTimeoutsPtrOutput `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from DeviceCheck artifacts will be valid. If unset, a
	// default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringOutput `pulumi:"tokenTtl"`
}

// NewFirebaseAppCheckDeviceCheckConfig registers a new resource with the given unique name, arguments, and options.
func NewFirebaseAppCheckDeviceCheckConfig(ctx *pulumi.Context,
	name string, args *FirebaseAppCheckDeviceCheckConfigArgs, opts ...pulumi.ResourceOption) (*FirebaseAppCheckDeviceCheckConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseAppCheckDeviceCheckConfig
	err = ctx.RegisterPackageResource("google-beta:index/firebaseAppCheckDeviceCheckConfig:FirebaseAppCheckDeviceCheckConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseAppCheckDeviceCheckConfig gets an existing FirebaseAppCheckDeviceCheckConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseAppCheckDeviceCheckConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseAppCheckDeviceCheckConfigState, opts ...pulumi.ResourceOption) (*FirebaseAppCheckDeviceCheckConfig, error) {
	var resource FirebaseAppCheckDeviceCheckConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/firebaseAppCheckDeviceCheckConfig:FirebaseAppCheckDeviceCheckConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseAppCheckDeviceCheckConfig resources.
type firebaseAppCheckDeviceCheckConfigState struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                               *string `pulumi:"appId"`
	FirebaseAppCheckDeviceCheckConfigId *string `pulumi:"firebaseAppCheckDeviceCheckConfigId"`
	// The key identifier of a private key enabled with DeviceCheck, created in your Apple Developer account.
	KeyId *string `pulumi:"keyId"`
	// The relative resource name of the DeviceCheck configuration object
	Name *string `pulumi:"name"`
	// The contents of the private key (.p8) file associated with the key specified by keyId.
	PrivateKey *string `pulumi:"privateKey"`
	// Whether the privateKey field was previously set. Since App Check will never return the privateKey field, this field is
	// the only way to find out whether it was previously set.
	PrivateKeySet *bool                                      `pulumi:"privateKeySet"`
	Project       *string                                    `pulumi:"project"`
	Timeouts      *FirebaseAppCheckDeviceCheckConfigTimeouts `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from DeviceCheck artifacts will be valid. If unset, a
	// default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl *string `pulumi:"tokenTtl"`
}

type FirebaseAppCheckDeviceCheckConfigState struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                               pulumi.StringPtrInput
	FirebaseAppCheckDeviceCheckConfigId pulumi.StringPtrInput
	// The key identifier of a private key enabled with DeviceCheck, created in your Apple Developer account.
	KeyId pulumi.StringPtrInput
	// The relative resource name of the DeviceCheck configuration object
	Name pulumi.StringPtrInput
	// The contents of the private key (.p8) file associated with the key specified by keyId.
	PrivateKey pulumi.StringPtrInput
	// Whether the privateKey field was previously set. Since App Check will never return the privateKey field, this field is
	// the only way to find out whether it was previously set.
	PrivateKeySet pulumi.BoolPtrInput
	Project       pulumi.StringPtrInput
	Timeouts      FirebaseAppCheckDeviceCheckConfigTimeoutsPtrInput
	// Specifies the duration for which App Check tokens exchanged from DeviceCheck artifacts will be valid. If unset, a
	// default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringPtrInput
}

func (FirebaseAppCheckDeviceCheckConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckDeviceCheckConfigState)(nil)).Elem()
}

type firebaseAppCheckDeviceCheckConfigArgs struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                               string  `pulumi:"appId"`
	FirebaseAppCheckDeviceCheckConfigId *string `pulumi:"firebaseAppCheckDeviceCheckConfigId"`
	// The key identifier of a private key enabled with DeviceCheck, created in your Apple Developer account.
	KeyId string `pulumi:"keyId"`
	// The contents of the private key (.p8) file associated with the key specified by keyId.
	PrivateKey string                                     `pulumi:"privateKey"`
	Project    *string                                    `pulumi:"project"`
	Timeouts   *FirebaseAppCheckDeviceCheckConfigTimeouts `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from DeviceCheck artifacts will be valid. If unset, a
	// default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl *string `pulumi:"tokenTtl"`
}

// The set of arguments for constructing a FirebaseAppCheckDeviceCheckConfig resource.
type FirebaseAppCheckDeviceCheckConfigArgs struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                               pulumi.StringInput
	FirebaseAppCheckDeviceCheckConfigId pulumi.StringPtrInput
	// The key identifier of a private key enabled with DeviceCheck, created in your Apple Developer account.
	KeyId pulumi.StringInput
	// The contents of the private key (.p8) file associated with the key specified by keyId.
	PrivateKey pulumi.StringInput
	Project    pulumi.StringPtrInput
	Timeouts   FirebaseAppCheckDeviceCheckConfigTimeoutsPtrInput
	// Specifies the duration for which App Check tokens exchanged from DeviceCheck artifacts will be valid. If unset, a
	// default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringPtrInput
}

func (FirebaseAppCheckDeviceCheckConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckDeviceCheckConfigArgs)(nil)).Elem()
}

type FirebaseAppCheckDeviceCheckConfigInput interface {
	pulumi.Input

	ToFirebaseAppCheckDeviceCheckConfigOutput() FirebaseAppCheckDeviceCheckConfigOutput
	ToFirebaseAppCheckDeviceCheckConfigOutputWithContext(ctx context.Context) FirebaseAppCheckDeviceCheckConfigOutput
}

func (*FirebaseAppCheckDeviceCheckConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckDeviceCheckConfig)(nil)).Elem()
}

func (i *FirebaseAppCheckDeviceCheckConfig) ToFirebaseAppCheckDeviceCheckConfigOutput() FirebaseAppCheckDeviceCheckConfigOutput {
	return i.ToFirebaseAppCheckDeviceCheckConfigOutputWithContext(context.Background())
}

func (i *FirebaseAppCheckDeviceCheckConfig) ToFirebaseAppCheckDeviceCheckConfigOutputWithContext(ctx context.Context) FirebaseAppCheckDeviceCheckConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseAppCheckDeviceCheckConfigOutput)
}

type FirebaseAppCheckDeviceCheckConfigOutput struct{ *pulumi.OutputState }

func (FirebaseAppCheckDeviceCheckConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckDeviceCheckConfig)(nil)).Elem()
}

func (o FirebaseAppCheckDeviceCheckConfigOutput) ToFirebaseAppCheckDeviceCheckConfigOutput() FirebaseAppCheckDeviceCheckConfigOutput {
	return o
}

func (o FirebaseAppCheckDeviceCheckConfigOutput) ToFirebaseAppCheckDeviceCheckConfigOutputWithContext(ctx context.Context) FirebaseAppCheckDeviceCheckConfigOutput {
	return o
}

// The ID of an [Apple
// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
func (o FirebaseAppCheckDeviceCheckConfigOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckDeviceCheckConfigOutput) FirebaseAppCheckDeviceCheckConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.StringOutput {
		return v.FirebaseAppCheckDeviceCheckConfigId
	}).(pulumi.StringOutput)
}

// The key identifier of a private key enabled with DeviceCheck, created in your Apple Developer account.
func (o FirebaseAppCheckDeviceCheckConfigOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The relative resource name of the DeviceCheck configuration object
func (o FirebaseAppCheckDeviceCheckConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The contents of the private key (.p8) file associated with the key specified by keyId.
func (o FirebaseAppCheckDeviceCheckConfigOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// Whether the privateKey field was previously set. Since App Check will never return the privateKey field, this field is
// the only way to find out whether it was previously set.
func (o FirebaseAppCheckDeviceCheckConfigOutput) PrivateKeySet() pulumi.BoolOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.BoolOutput { return v.PrivateKeySet }).(pulumi.BoolOutput)
}

func (o FirebaseAppCheckDeviceCheckConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckDeviceCheckConfigOutput) Timeouts() FirebaseAppCheckDeviceCheckConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) FirebaseAppCheckDeviceCheckConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(FirebaseAppCheckDeviceCheckConfigTimeoutsPtrOutput)
}

// Specifies the duration for which App Check tokens exchanged from DeviceCheck artifacts will be valid. If unset, a
// default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
// nine fractional digits, ending with 's'. Example: "3.5s".
func (o FirebaseAppCheckDeviceCheckConfigOutput) TokenTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDeviceCheckConfig) pulumi.StringOutput { return v.TokenTtl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseAppCheckDeviceCheckConfigInput)(nil)).Elem(), &FirebaseAppCheckDeviceCheckConfig{})
	pulumi.RegisterOutputType(FirebaseAppCheckDeviceCheckConfigOutput{})
}
