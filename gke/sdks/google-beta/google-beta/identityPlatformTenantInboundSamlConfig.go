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

type IdentityPlatformTenantInboundSamlConfig struct {
	pulumi.CustomResourceState

	// Human friendly display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled                                   pulumi.BoolPtrOutput `pulumi:"enabled"`
	IdentityPlatformTenantInboundSamlConfigId pulumi.StringOutput  `pulumi:"identityPlatformTenantInboundSamlConfigId"`
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig IdentityPlatformTenantInboundSamlConfigIdpConfigOutput `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig IdentityPlatformTenantInboundSamlConfigSpConfigOutput `pulumi:"spConfig"`
	// The name of the tenant where this inbound SAML config resource exists
	Tenant   pulumi.StringOutput                                      `pulumi:"tenant"`
	Timeouts IdentityPlatformTenantInboundSamlConfigTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewIdentityPlatformTenantInboundSamlConfig registers a new resource with the given unique name, arguments, and options.
func NewIdentityPlatformTenantInboundSamlConfig(ctx *pulumi.Context,
	name string, args *IdentityPlatformTenantInboundSamlConfigArgs, opts ...pulumi.ResourceOption) (*IdentityPlatformTenantInboundSamlConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IdpConfig == nil {
		return nil, errors.New("invalid value for required argument 'IdpConfig'")
	}
	if args.SpConfig == nil {
		return nil, errors.New("invalid value for required argument 'SpConfig'")
	}
	if args.Tenant == nil {
		return nil, errors.New("invalid value for required argument 'Tenant'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IdentityPlatformTenantInboundSamlConfig
	err = ctx.RegisterPackageResource("google-beta:index/identityPlatformTenantInboundSamlConfig:IdentityPlatformTenantInboundSamlConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityPlatformTenantInboundSamlConfig gets an existing IdentityPlatformTenantInboundSamlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPlatformTenantInboundSamlConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityPlatformTenantInboundSamlConfigState, opts ...pulumi.ResourceOption) (*IdentityPlatformTenantInboundSamlConfig, error) {
	var resource IdentityPlatformTenantInboundSamlConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/identityPlatformTenantInboundSamlConfig:IdentityPlatformTenantInboundSamlConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityPlatformTenantInboundSamlConfig resources.
type identityPlatformTenantInboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled                                   *bool   `pulumi:"enabled"`
	IdentityPlatformTenantInboundSamlConfigId *string `pulumi:"identityPlatformTenantInboundSamlConfigId"`
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig *IdentityPlatformTenantInboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig *IdentityPlatformTenantInboundSamlConfigSpConfig `pulumi:"spConfig"`
	// The name of the tenant where this inbound SAML config resource exists
	Tenant   *string                                          `pulumi:"tenant"`
	Timeouts *IdentityPlatformTenantInboundSamlConfigTimeouts `pulumi:"timeouts"`
}

type IdentityPlatformTenantInboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName pulumi.StringPtrInput
	// If this config allows users to sign in with the provider.
	Enabled                                   pulumi.BoolPtrInput
	IdentityPlatformTenantInboundSamlConfigId pulumi.StringPtrInput
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig IdentityPlatformTenantInboundSamlConfigIdpConfigPtrInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig IdentityPlatformTenantInboundSamlConfigSpConfigPtrInput
	// The name of the tenant where this inbound SAML config resource exists
	Tenant   pulumi.StringPtrInput
	Timeouts IdentityPlatformTenantInboundSamlConfigTimeoutsPtrInput
}

func (IdentityPlatformTenantInboundSamlConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPlatformTenantInboundSamlConfigState)(nil)).Elem()
}

type identityPlatformTenantInboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled                                   *bool   `pulumi:"enabled"`
	IdentityPlatformTenantInboundSamlConfigId *string `pulumi:"identityPlatformTenantInboundSamlConfigId"`
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig IdentityPlatformTenantInboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig IdentityPlatformTenantInboundSamlConfigSpConfig `pulumi:"spConfig"`
	// The name of the tenant where this inbound SAML config resource exists
	Tenant   string                                           `pulumi:"tenant"`
	Timeouts *IdentityPlatformTenantInboundSamlConfigTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a IdentityPlatformTenantInboundSamlConfig resource.
type IdentityPlatformTenantInboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName pulumi.StringInput
	// If this config allows users to sign in with the provider.
	Enabled                                   pulumi.BoolPtrInput
	IdentityPlatformTenantInboundSamlConfigId pulumi.StringPtrInput
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig IdentityPlatformTenantInboundSamlConfigIdpConfigInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig IdentityPlatformTenantInboundSamlConfigSpConfigInput
	// The name of the tenant where this inbound SAML config resource exists
	Tenant   pulumi.StringInput
	Timeouts IdentityPlatformTenantInboundSamlConfigTimeoutsPtrInput
}

func (IdentityPlatformTenantInboundSamlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPlatformTenantInboundSamlConfigArgs)(nil)).Elem()
}

type IdentityPlatformTenantInboundSamlConfigInput interface {
	pulumi.Input

	ToIdentityPlatformTenantInboundSamlConfigOutput() IdentityPlatformTenantInboundSamlConfigOutput
	ToIdentityPlatformTenantInboundSamlConfigOutputWithContext(ctx context.Context) IdentityPlatformTenantInboundSamlConfigOutput
}

func (*IdentityPlatformTenantInboundSamlConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPlatformTenantInboundSamlConfig)(nil)).Elem()
}

func (i *IdentityPlatformTenantInboundSamlConfig) ToIdentityPlatformTenantInboundSamlConfigOutput() IdentityPlatformTenantInboundSamlConfigOutput {
	return i.ToIdentityPlatformTenantInboundSamlConfigOutputWithContext(context.Background())
}

func (i *IdentityPlatformTenantInboundSamlConfig) ToIdentityPlatformTenantInboundSamlConfigOutputWithContext(ctx context.Context) IdentityPlatformTenantInboundSamlConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPlatformTenantInboundSamlConfigOutput)
}

type IdentityPlatformTenantInboundSamlConfigOutput struct{ *pulumi.OutputState }

func (IdentityPlatformTenantInboundSamlConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPlatformTenantInboundSamlConfig)(nil)).Elem()
}

func (o IdentityPlatformTenantInboundSamlConfigOutput) ToIdentityPlatformTenantInboundSamlConfigOutput() IdentityPlatformTenantInboundSamlConfigOutput {
	return o
}

func (o IdentityPlatformTenantInboundSamlConfigOutput) ToIdentityPlatformTenantInboundSamlConfigOutputWithContext(ctx context.Context) IdentityPlatformTenantInboundSamlConfigOutput {
	return o
}

// Human friendly display name.
func (o IdentityPlatformTenantInboundSamlConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// If this config allows users to sign in with the provider.
func (o IdentityPlatformTenantInboundSamlConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o IdentityPlatformTenantInboundSamlConfigOutput) IdentityPlatformTenantInboundSamlConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) pulumi.StringOutput {
		return v.IdentityPlatformTenantInboundSamlConfigId
	}).(pulumi.StringOutput)
}

// SAML IdP configuration when the project acts as the relying party
func (o IdentityPlatformTenantInboundSamlConfigOutput) IdpConfig() IdentityPlatformTenantInboundSamlConfigIdpConfigOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) IdentityPlatformTenantInboundSamlConfigIdpConfigOutput {
		return v.IdpConfig
	}).(IdentityPlatformTenantInboundSamlConfigIdpConfigOutput)
}

// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
// character, and have at least 2 characters.
func (o IdentityPlatformTenantInboundSamlConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IdentityPlatformTenantInboundSamlConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
// authentication assertion issued by a SAML identity provider.
func (o IdentityPlatformTenantInboundSamlConfigOutput) SpConfig() IdentityPlatformTenantInboundSamlConfigSpConfigOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) IdentityPlatformTenantInboundSamlConfigSpConfigOutput {
		return v.SpConfig
	}).(IdentityPlatformTenantInboundSamlConfigSpConfigOutput)
}

// The name of the tenant where this inbound SAML config resource exists
func (o IdentityPlatformTenantInboundSamlConfigOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

func (o IdentityPlatformTenantInboundSamlConfigOutput) Timeouts() IdentityPlatformTenantInboundSamlConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *IdentityPlatformTenantInboundSamlConfig) IdentityPlatformTenantInboundSamlConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(IdentityPlatformTenantInboundSamlConfigTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPlatformTenantInboundSamlConfigInput)(nil)).Elem(), &IdentityPlatformTenantInboundSamlConfig{})
	pulumi.RegisterOutputType(IdentityPlatformTenantInboundSamlConfigOutput{})
}
