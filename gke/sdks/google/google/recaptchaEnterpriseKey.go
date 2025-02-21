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

type RecaptchaEnterpriseKey struct {
	pulumi.CustomResourceState

	// Settings for keys that can be used by Android apps.
	AndroidSettings RecaptchaEnterpriseKeyAndroidSettingsPtrOutput `pulumi:"androidSettings"`
	// The timestamp corresponding to the creation of this Key.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human-readable display name of this key. Modifiable by user.
	DisplayName     pulumi.StringOutput    `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Settings for keys that can be used by iOS apps.
	IosSettings RecaptchaEnterpriseKeyIosSettingsPtrOutput `pulumi:"iosSettings"`
	// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// `effective_labels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource id for the Key, which is the same as the Site Key itself.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project                  pulumi.StringOutput `pulumi:"project"`
	RecaptchaEnterpriseKeyId pulumi.StringOutput `pulumi:"recaptchaEnterpriseKeyId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput `pulumi:"terraformLabels"`
	// Options for user acceptance testing.
	TestingOptions RecaptchaEnterpriseKeyTestingOptionsPtrOutput `pulumi:"testingOptions"`
	Timeouts       RecaptchaEnterpriseKeyTimeoutsPtrOutput       `pulumi:"timeouts"`
	// Settings specific to keys that can be used for WAF (Web Application Firewall).
	WafSettings RecaptchaEnterpriseKeyWafSettingsPtrOutput `pulumi:"wafSettings"`
	// Settings for keys that can be used by websites.
	WebSettings RecaptchaEnterpriseKeyWebSettingsPtrOutput `pulumi:"webSettings"`
}

// NewRecaptchaEnterpriseKey registers a new resource with the given unique name, arguments, and options.
func NewRecaptchaEnterpriseKey(ctx *pulumi.Context,
	name string, args *RecaptchaEnterpriseKeyArgs, opts ...pulumi.ResourceOption) (*RecaptchaEnterpriseKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource RecaptchaEnterpriseKey
	err = ctx.RegisterPackageResource("google:index/recaptchaEnterpriseKey:RecaptchaEnterpriseKey", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecaptchaEnterpriseKey gets an existing RecaptchaEnterpriseKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecaptchaEnterpriseKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecaptchaEnterpriseKeyState, opts ...pulumi.ResourceOption) (*RecaptchaEnterpriseKey, error) {
	var resource RecaptchaEnterpriseKey
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/recaptchaEnterpriseKey:RecaptchaEnterpriseKey", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecaptchaEnterpriseKey resources.
type recaptchaEnterpriseKeyState struct {
	// Settings for keys that can be used by Android apps.
	AndroidSettings *RecaptchaEnterpriseKeyAndroidSettings `pulumi:"androidSettings"`
	// The timestamp corresponding to the creation of this Key.
	CreateTime *string `pulumi:"createTime"`
	// Human-readable display name of this key. Modifiable by user.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Settings for keys that can be used by iOS apps.
	IosSettings *RecaptchaEnterpriseKeyIosSettings `pulumi:"iosSettings"`
	// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// `effective_labels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource id for the Key, which is the same as the Site Key itself.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project                  *string `pulumi:"project"`
	RecaptchaEnterpriseKeyId *string `pulumi:"recaptchaEnterpriseKeyId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string `pulumi:"terraformLabels"`
	// Options for user acceptance testing.
	TestingOptions *RecaptchaEnterpriseKeyTestingOptions `pulumi:"testingOptions"`
	Timeouts       *RecaptchaEnterpriseKeyTimeouts       `pulumi:"timeouts"`
	// Settings specific to keys that can be used for WAF (Web Application Firewall).
	WafSettings *RecaptchaEnterpriseKeyWafSettings `pulumi:"wafSettings"`
	// Settings for keys that can be used by websites.
	WebSettings *RecaptchaEnterpriseKeyWebSettings `pulumi:"webSettings"`
}

type RecaptchaEnterpriseKeyState struct {
	// Settings for keys that can be used by Android apps.
	AndroidSettings RecaptchaEnterpriseKeyAndroidSettingsPtrInput
	// The timestamp corresponding to the creation of this Key.
	CreateTime pulumi.StringPtrInput
	// Human-readable display name of this key. Modifiable by user.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Settings for keys that can be used by iOS apps.
	IosSettings RecaptchaEnterpriseKeyIosSettingsPtrInput
	// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// `effective_labels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The resource id for the Key, which is the same as the Site Key itself.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project                  pulumi.StringPtrInput
	RecaptchaEnterpriseKeyId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	// Options for user acceptance testing.
	TestingOptions RecaptchaEnterpriseKeyTestingOptionsPtrInput
	Timeouts       RecaptchaEnterpriseKeyTimeoutsPtrInput
	// Settings specific to keys that can be used for WAF (Web Application Firewall).
	WafSettings RecaptchaEnterpriseKeyWafSettingsPtrInput
	// Settings for keys that can be used by websites.
	WebSettings RecaptchaEnterpriseKeyWebSettingsPtrInput
}

func (RecaptchaEnterpriseKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*recaptchaEnterpriseKeyState)(nil)).Elem()
}

type recaptchaEnterpriseKeyArgs struct {
	// Settings for keys that can be used by Android apps.
	AndroidSettings *RecaptchaEnterpriseKeyAndroidSettings `pulumi:"androidSettings"`
	// Human-readable display name of this key. Modifiable by user.
	DisplayName string `pulumi:"displayName"`
	// Settings for keys that can be used by iOS apps.
	IosSettings *RecaptchaEnterpriseKeyIosSettings `pulumi:"iosSettings"`
	// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// `effective_labels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The project for the resource
	Project                  *string `pulumi:"project"`
	RecaptchaEnterpriseKeyId *string `pulumi:"recaptchaEnterpriseKeyId"`
	// Options for user acceptance testing.
	TestingOptions *RecaptchaEnterpriseKeyTestingOptions `pulumi:"testingOptions"`
	Timeouts       *RecaptchaEnterpriseKeyTimeouts       `pulumi:"timeouts"`
	// Settings specific to keys that can be used for WAF (Web Application Firewall).
	WafSettings *RecaptchaEnterpriseKeyWafSettings `pulumi:"wafSettings"`
	// Settings for keys that can be used by websites.
	WebSettings *RecaptchaEnterpriseKeyWebSettings `pulumi:"webSettings"`
}

// The set of arguments for constructing a RecaptchaEnterpriseKey resource.
type RecaptchaEnterpriseKeyArgs struct {
	// Settings for keys that can be used by Android apps.
	AndroidSettings RecaptchaEnterpriseKeyAndroidSettingsPtrInput
	// Human-readable display name of this key. Modifiable by user.
	DisplayName pulumi.StringInput
	// Settings for keys that can be used by iOS apps.
	IosSettings RecaptchaEnterpriseKeyIosSettingsPtrInput
	// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// `effective_labels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The project for the resource
	Project                  pulumi.StringPtrInput
	RecaptchaEnterpriseKeyId pulumi.StringPtrInput
	// Options for user acceptance testing.
	TestingOptions RecaptchaEnterpriseKeyTestingOptionsPtrInput
	Timeouts       RecaptchaEnterpriseKeyTimeoutsPtrInput
	// Settings specific to keys that can be used for WAF (Web Application Firewall).
	WafSettings RecaptchaEnterpriseKeyWafSettingsPtrInput
	// Settings for keys that can be used by websites.
	WebSettings RecaptchaEnterpriseKeyWebSettingsPtrInput
}

func (RecaptchaEnterpriseKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recaptchaEnterpriseKeyArgs)(nil)).Elem()
}

type RecaptchaEnterpriseKeyInput interface {
	pulumi.Input

	ToRecaptchaEnterpriseKeyOutput() RecaptchaEnterpriseKeyOutput
	ToRecaptchaEnterpriseKeyOutputWithContext(ctx context.Context) RecaptchaEnterpriseKeyOutput
}

func (*RecaptchaEnterpriseKey) ElementType() reflect.Type {
	return reflect.TypeOf((**RecaptchaEnterpriseKey)(nil)).Elem()
}

func (i *RecaptchaEnterpriseKey) ToRecaptchaEnterpriseKeyOutput() RecaptchaEnterpriseKeyOutput {
	return i.ToRecaptchaEnterpriseKeyOutputWithContext(context.Background())
}

func (i *RecaptchaEnterpriseKey) ToRecaptchaEnterpriseKeyOutputWithContext(ctx context.Context) RecaptchaEnterpriseKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecaptchaEnterpriseKeyOutput)
}

type RecaptchaEnterpriseKeyOutput struct{ *pulumi.OutputState }

func (RecaptchaEnterpriseKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecaptchaEnterpriseKey)(nil)).Elem()
}

func (o RecaptchaEnterpriseKeyOutput) ToRecaptchaEnterpriseKeyOutput() RecaptchaEnterpriseKeyOutput {
	return o
}

func (o RecaptchaEnterpriseKeyOutput) ToRecaptchaEnterpriseKeyOutputWithContext(ctx context.Context) RecaptchaEnterpriseKeyOutput {
	return o
}

// Settings for keys that can be used by Android apps.
func (o RecaptchaEnterpriseKeyOutput) AndroidSettings() RecaptchaEnterpriseKeyAndroidSettingsPtrOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) RecaptchaEnterpriseKeyAndroidSettingsPtrOutput {
		return v.AndroidSettings
	}).(RecaptchaEnterpriseKeyAndroidSettingsPtrOutput)
}

// The timestamp corresponding to the creation of this Key.
func (o RecaptchaEnterpriseKeyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Human-readable display name of this key. Modifiable by user.
func (o RecaptchaEnterpriseKeyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o RecaptchaEnterpriseKeyOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Settings for keys that can be used by iOS apps.
func (o RecaptchaEnterpriseKeyOutput) IosSettings() RecaptchaEnterpriseKeyIosSettingsPtrOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) RecaptchaEnterpriseKeyIosSettingsPtrOutput { return v.IosSettings }).(RecaptchaEnterpriseKeyIosSettingsPtrOutput)
}

// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels). **Note**: This field is
// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
// `effective_labels` for all of the labels present on the resource.
func (o RecaptchaEnterpriseKeyOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource id for the Key, which is the same as the Site Key itself.
func (o RecaptchaEnterpriseKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o RecaptchaEnterpriseKeyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RecaptchaEnterpriseKeyOutput) RecaptchaEnterpriseKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringOutput { return v.RecaptchaEnterpriseKeyId }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o RecaptchaEnterpriseKeyOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

// Options for user acceptance testing.
func (o RecaptchaEnterpriseKeyOutput) TestingOptions() RecaptchaEnterpriseKeyTestingOptionsPtrOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) RecaptchaEnterpriseKeyTestingOptionsPtrOutput { return v.TestingOptions }).(RecaptchaEnterpriseKeyTestingOptionsPtrOutput)
}

func (o RecaptchaEnterpriseKeyOutput) Timeouts() RecaptchaEnterpriseKeyTimeoutsPtrOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) RecaptchaEnterpriseKeyTimeoutsPtrOutput { return v.Timeouts }).(RecaptchaEnterpriseKeyTimeoutsPtrOutput)
}

// Settings specific to keys that can be used for WAF (Web Application Firewall).
func (o RecaptchaEnterpriseKeyOutput) WafSettings() RecaptchaEnterpriseKeyWafSettingsPtrOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) RecaptchaEnterpriseKeyWafSettingsPtrOutput { return v.WafSettings }).(RecaptchaEnterpriseKeyWafSettingsPtrOutput)
}

// Settings for keys that can be used by websites.
func (o RecaptchaEnterpriseKeyOutput) WebSettings() RecaptchaEnterpriseKeyWebSettingsPtrOutput {
	return o.ApplyT(func(v *RecaptchaEnterpriseKey) RecaptchaEnterpriseKeyWebSettingsPtrOutput { return v.WebSettings }).(RecaptchaEnterpriseKeyWebSettingsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecaptchaEnterpriseKeyInput)(nil)).Elem(), &RecaptchaEnterpriseKey{})
	pulumi.RegisterOutputType(RecaptchaEnterpriseKeyOutput{})
}
