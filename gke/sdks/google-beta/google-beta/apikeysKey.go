// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApikeysKey struct {
	pulumi.CustomResourceState

	ApikeysKeyId pulumi.StringOutput `pulumi:"apikeysKeyId"`
	// Human-readable display name of this API key. Modifiable by user.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Output only. An encrypted and signed value held by this key. This field can be accessed only through the `GetKeyString`
	// method.
	KeyString pulumi.StringOutput `pulumi:"keyString"`
	// The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to
	// lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular
	// expression: `a-z?`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Key restrictions.
	Restrictions ApikeysKeyRestrictionsPtrOutput `pulumi:"restrictions"`
	Timeouts     ApikeysKeyTimeoutsPtrOutput     `pulumi:"timeouts"`
	// Output only. Unique id in UUID4 format.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewApikeysKey registers a new resource with the given unique name, arguments, and options.
func NewApikeysKey(ctx *pulumi.Context,
	name string, args *ApikeysKeyArgs, opts ...pulumi.ResourceOption) (*ApikeysKey, error) {
	if args == nil {
		args = &ApikeysKeyArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keyString",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApikeysKey
	err = ctx.RegisterPackageResource("google-beta:index/apikeysKey:ApikeysKey", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApikeysKey gets an existing ApikeysKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApikeysKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApikeysKeyState, opts ...pulumi.ResourceOption) (*ApikeysKey, error) {
	var resource ApikeysKey
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/apikeysKey:ApikeysKey", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApikeysKey resources.
type apikeysKeyState struct {
	ApikeysKeyId *string `pulumi:"apikeysKeyId"`
	// Human-readable display name of this API key. Modifiable by user.
	DisplayName *string `pulumi:"displayName"`
	// Output only. An encrypted and signed value held by this key. This field can be accessed only through the `GetKeyString`
	// method.
	KeyString *string `pulumi:"keyString"`
	// The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to
	// lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular
	// expression: `a-z?`.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Key restrictions.
	Restrictions *ApikeysKeyRestrictions `pulumi:"restrictions"`
	Timeouts     *ApikeysKeyTimeouts     `pulumi:"timeouts"`
	// Output only. Unique id in UUID4 format.
	Uid *string `pulumi:"uid"`
}

type ApikeysKeyState struct {
	ApikeysKeyId pulumi.StringPtrInput
	// Human-readable display name of this API key. Modifiable by user.
	DisplayName pulumi.StringPtrInput
	// Output only. An encrypted and signed value held by this key. This field can be accessed only through the `GetKeyString`
	// method.
	KeyString pulumi.StringPtrInput
	// The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to
	// lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular
	// expression: `a-z?`.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Key restrictions.
	Restrictions ApikeysKeyRestrictionsPtrInput
	Timeouts     ApikeysKeyTimeoutsPtrInput
	// Output only. Unique id in UUID4 format.
	Uid pulumi.StringPtrInput
}

func (ApikeysKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeysKeyState)(nil)).Elem()
}

type apikeysKeyArgs struct {
	ApikeysKeyId *string `pulumi:"apikeysKeyId"`
	// Human-readable display name of this API key. Modifiable by user.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to
	// lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular
	// expression: `a-z?`.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Key restrictions.
	Restrictions *ApikeysKeyRestrictions `pulumi:"restrictions"`
	Timeouts     *ApikeysKeyTimeouts     `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApikeysKey resource.
type ApikeysKeyArgs struct {
	ApikeysKeyId pulumi.StringPtrInput
	// Human-readable display name of this API key. Modifiable by user.
	DisplayName pulumi.StringPtrInput
	// The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to
	// lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular
	// expression: `a-z?`.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Key restrictions.
	Restrictions ApikeysKeyRestrictionsPtrInput
	Timeouts     ApikeysKeyTimeoutsPtrInput
}

func (ApikeysKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeysKeyArgs)(nil)).Elem()
}

type ApikeysKeyInput interface {
	pulumi.Input

	ToApikeysKeyOutput() ApikeysKeyOutput
	ToApikeysKeyOutputWithContext(ctx context.Context) ApikeysKeyOutput
}

func (*ApikeysKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ApikeysKey)(nil)).Elem()
}

func (i *ApikeysKey) ToApikeysKeyOutput() ApikeysKeyOutput {
	return i.ToApikeysKeyOutputWithContext(context.Background())
}

func (i *ApikeysKey) ToApikeysKeyOutputWithContext(ctx context.Context) ApikeysKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApikeysKeyOutput)
}

type ApikeysKeyOutput struct{ *pulumi.OutputState }

func (ApikeysKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApikeysKey)(nil)).Elem()
}

func (o ApikeysKeyOutput) ToApikeysKeyOutput() ApikeysKeyOutput {
	return o
}

func (o ApikeysKeyOutput) ToApikeysKeyOutputWithContext(ctx context.Context) ApikeysKeyOutput {
	return o
}

func (o ApikeysKeyOutput) ApikeysKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApikeysKey) pulumi.StringOutput { return v.ApikeysKeyId }).(pulumi.StringOutput)
}

// Human-readable display name of this API key. Modifiable by user.
func (o ApikeysKeyOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApikeysKey) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Output only. An encrypted and signed value held by this key. This field can be accessed only through the `GetKeyString`
// method.
func (o ApikeysKeyOutput) KeyString() pulumi.StringOutput {
	return o.ApplyT(func(v *ApikeysKey) pulumi.StringOutput { return v.KeyString }).(pulumi.StringOutput)
}

// The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to
// lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular
// expression: `a-z?`.
func (o ApikeysKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApikeysKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o ApikeysKeyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApikeysKey) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Key restrictions.
func (o ApikeysKeyOutput) Restrictions() ApikeysKeyRestrictionsPtrOutput {
	return o.ApplyT(func(v *ApikeysKey) ApikeysKeyRestrictionsPtrOutput { return v.Restrictions }).(ApikeysKeyRestrictionsPtrOutput)
}

func (o ApikeysKeyOutput) Timeouts() ApikeysKeyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApikeysKey) ApikeysKeyTimeoutsPtrOutput { return v.Timeouts }).(ApikeysKeyTimeoutsPtrOutput)
}

// Output only. Unique id in UUID4 format.
func (o ApikeysKeyOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ApikeysKey) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApikeysKeyInput)(nil)).Elem(), &ApikeysKey{})
	pulumi.RegisterOutputType(ApikeysKeyOutput{})
}
