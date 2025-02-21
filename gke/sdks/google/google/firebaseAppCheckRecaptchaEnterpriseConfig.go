// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirebaseAppCheckRecaptchaEnterpriseConfig struct {
	pulumi.CustomResourceState

	// The ID of an [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
	AppId                                       pulumi.StringOutput `pulumi:"appId"`
	FirebaseAppCheckRecaptchaEnterpriseConfigId pulumi.StringOutput `pulumi:"firebaseAppCheckRecaptchaEnterpriseConfigId"`
	// The relative resource name of the reCAPTCHA Enterprise configuration object
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for
	// your application. **Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based
	// reCAPTCHA Enterprise site key.
	SiteKey  pulumi.StringOutput                                        `pulumi:"siteKey"`
	Timeouts FirebaseAppCheckRecaptchaEnterpriseConfigTimeoutsPtrOutput `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid. If unset,
	// a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringOutput `pulumi:"tokenTtl"`
}

// NewFirebaseAppCheckRecaptchaEnterpriseConfig registers a new resource with the given unique name, arguments, and options.
func NewFirebaseAppCheckRecaptchaEnterpriseConfig(ctx *pulumi.Context,
	name string, args *FirebaseAppCheckRecaptchaEnterpriseConfigArgs, opts ...pulumi.ResourceOption) (*FirebaseAppCheckRecaptchaEnterpriseConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.SiteKey == nil {
		return nil, errors.New("invalid value for required argument 'SiteKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseAppCheckRecaptchaEnterpriseConfig
	err = ctx.RegisterPackageResource("google:index/firebaseAppCheckRecaptchaEnterpriseConfig:FirebaseAppCheckRecaptchaEnterpriseConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseAppCheckRecaptchaEnterpriseConfig gets an existing FirebaseAppCheckRecaptchaEnterpriseConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseAppCheckRecaptchaEnterpriseConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseAppCheckRecaptchaEnterpriseConfigState, opts ...pulumi.ResourceOption) (*FirebaseAppCheckRecaptchaEnterpriseConfig, error) {
	var resource FirebaseAppCheckRecaptchaEnterpriseConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/firebaseAppCheckRecaptchaEnterpriseConfig:FirebaseAppCheckRecaptchaEnterpriseConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseAppCheckRecaptchaEnterpriseConfig resources.
type firebaseAppCheckRecaptchaEnterpriseConfigState struct {
	// The ID of an [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
	AppId                                       *string `pulumi:"appId"`
	FirebaseAppCheckRecaptchaEnterpriseConfigId *string `pulumi:"firebaseAppCheckRecaptchaEnterpriseConfigId"`
	// The relative resource name of the reCAPTCHA Enterprise configuration object
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for
	// your application. **Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based
	// reCAPTCHA Enterprise site key.
	SiteKey  *string                                            `pulumi:"siteKey"`
	Timeouts *FirebaseAppCheckRecaptchaEnterpriseConfigTimeouts `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid. If unset,
	// a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl *string `pulumi:"tokenTtl"`
}

type FirebaseAppCheckRecaptchaEnterpriseConfigState struct {
	// The ID of an [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
	AppId                                       pulumi.StringPtrInput
	FirebaseAppCheckRecaptchaEnterpriseConfigId pulumi.StringPtrInput
	// The relative resource name of the reCAPTCHA Enterprise configuration object
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for
	// your application. **Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based
	// reCAPTCHA Enterprise site key.
	SiteKey  pulumi.StringPtrInput
	Timeouts FirebaseAppCheckRecaptchaEnterpriseConfigTimeoutsPtrInput
	// Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid. If unset,
	// a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringPtrInput
}

func (FirebaseAppCheckRecaptchaEnterpriseConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckRecaptchaEnterpriseConfigState)(nil)).Elem()
}

type firebaseAppCheckRecaptchaEnterpriseConfigArgs struct {
	// The ID of an [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
	AppId                                       string  `pulumi:"appId"`
	FirebaseAppCheckRecaptchaEnterpriseConfigId *string `pulumi:"firebaseAppCheckRecaptchaEnterpriseConfigId"`
	Project                                     *string `pulumi:"project"`
	// The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for
	// your application. **Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based
	// reCAPTCHA Enterprise site key.
	SiteKey  string                                             `pulumi:"siteKey"`
	Timeouts *FirebaseAppCheckRecaptchaEnterpriseConfigTimeouts `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid. If unset,
	// a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl *string `pulumi:"tokenTtl"`
}

// The set of arguments for constructing a FirebaseAppCheckRecaptchaEnterpriseConfig resource.
type FirebaseAppCheckRecaptchaEnterpriseConfigArgs struct {
	// The ID of an [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
	AppId                                       pulumi.StringInput
	FirebaseAppCheckRecaptchaEnterpriseConfigId pulumi.StringPtrInput
	Project                                     pulumi.StringPtrInput
	// The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for
	// your application. **Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based
	// reCAPTCHA Enterprise site key.
	SiteKey  pulumi.StringInput
	Timeouts FirebaseAppCheckRecaptchaEnterpriseConfigTimeoutsPtrInput
	// Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid. If unset,
	// a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
	// nine fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringPtrInput
}

func (FirebaseAppCheckRecaptchaEnterpriseConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckRecaptchaEnterpriseConfigArgs)(nil)).Elem()
}

type FirebaseAppCheckRecaptchaEnterpriseConfigInput interface {
	pulumi.Input

	ToFirebaseAppCheckRecaptchaEnterpriseConfigOutput() FirebaseAppCheckRecaptchaEnterpriseConfigOutput
	ToFirebaseAppCheckRecaptchaEnterpriseConfigOutputWithContext(ctx context.Context) FirebaseAppCheckRecaptchaEnterpriseConfigOutput
}

func (*FirebaseAppCheckRecaptchaEnterpriseConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckRecaptchaEnterpriseConfig)(nil)).Elem()
}

func (i *FirebaseAppCheckRecaptchaEnterpriseConfig) ToFirebaseAppCheckRecaptchaEnterpriseConfigOutput() FirebaseAppCheckRecaptchaEnterpriseConfigOutput {
	return i.ToFirebaseAppCheckRecaptchaEnterpriseConfigOutputWithContext(context.Background())
}

func (i *FirebaseAppCheckRecaptchaEnterpriseConfig) ToFirebaseAppCheckRecaptchaEnterpriseConfigOutputWithContext(ctx context.Context) FirebaseAppCheckRecaptchaEnterpriseConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseAppCheckRecaptchaEnterpriseConfigOutput)
}

type FirebaseAppCheckRecaptchaEnterpriseConfigOutput struct{ *pulumi.OutputState }

func (FirebaseAppCheckRecaptchaEnterpriseConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckRecaptchaEnterpriseConfig)(nil)).Elem()
}

func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) ToFirebaseAppCheckRecaptchaEnterpriseConfigOutput() FirebaseAppCheckRecaptchaEnterpriseConfigOutput {
	return o
}

func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) ToFirebaseAppCheckRecaptchaEnterpriseConfigOutputWithContext(ctx context.Context) FirebaseAppCheckRecaptchaEnterpriseConfigOutput {
	return o
}

// The ID of an [Web
// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckRecaptchaEnterpriseConfig) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) FirebaseAppCheckRecaptchaEnterpriseConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckRecaptchaEnterpriseConfig) pulumi.StringOutput {
		return v.FirebaseAppCheckRecaptchaEnterpriseConfigId
	}).(pulumi.StringOutput)
}

// The relative resource name of the reCAPTCHA Enterprise configuration object
func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckRecaptchaEnterpriseConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckRecaptchaEnterpriseConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for
// your application. **Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based
// reCAPTCHA Enterprise site key.
func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) SiteKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckRecaptchaEnterpriseConfig) pulumi.StringOutput { return v.SiteKey }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) Timeouts() FirebaseAppCheckRecaptchaEnterpriseConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseAppCheckRecaptchaEnterpriseConfig) FirebaseAppCheckRecaptchaEnterpriseConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(FirebaseAppCheckRecaptchaEnterpriseConfigTimeoutsPtrOutput)
}

// Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid. If unset,
// a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to
// nine fractional digits, ending with 's'. Example: "3.5s".
func (o FirebaseAppCheckRecaptchaEnterpriseConfigOutput) TokenTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckRecaptchaEnterpriseConfig) pulumi.StringOutput { return v.TokenTtl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseAppCheckRecaptchaEnterpriseConfigInput)(nil)).Elem(), &FirebaseAppCheckRecaptchaEnterpriseConfig{})
	pulumi.RegisterOutputType(FirebaseAppCheckRecaptchaEnterpriseConfigOutput{})
}
