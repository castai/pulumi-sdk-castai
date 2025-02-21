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

type SecureSourceManagerInstance struct {
	pulumi.CustomResourceState

	// Time the Instance was created in UTC.
	CreateTime      pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// A list of hostnames for this instance.
	HostConfigs SecureSourceManagerInstanceHostConfigArrayOutput `pulumi:"hostConfigs"`
	// The name for the Instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.
	KmsKey pulumi.StringPtrOutput `pulumi:"kmsKey"`
	// Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the Instance.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the Instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private settings for private instance.
	PrivateConfig                 SecureSourceManagerInstancePrivateConfigPtrOutput `pulumi:"privateConfig"`
	Project                       pulumi.StringOutput                               `pulumi:"project"`
	SecureSourceManagerInstanceId pulumi.StringOutput                               `pulumi:"secureSourceManagerInstanceId"`
	// The current state of the Instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Provides information about the current instance state.
	StateNote pulumi.StringOutput `pulumi:"stateNote"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                       `pulumi:"terraformLabels"`
	Timeouts        SecureSourceManagerInstanceTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the Instance was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Configuration for Workforce Identity Federation to support third party identity provider. If unset, defaults to the
	// Google OIDC IdP.
	WorkforceIdentityFederationConfig SecureSourceManagerInstanceWorkforceIdentityFederationConfigPtrOutput `pulumi:"workforceIdentityFederationConfig"`
}

// NewSecureSourceManagerInstance registers a new resource with the given unique name, arguments, and options.
func NewSecureSourceManagerInstance(ctx *pulumi.Context,
	name string, args *SecureSourceManagerInstanceArgs, opts ...pulumi.ResourceOption) (*SecureSourceManagerInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SecureSourceManagerInstance
	err = ctx.RegisterPackageResource("google-beta:index/secureSourceManagerInstance:SecureSourceManagerInstance", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecureSourceManagerInstance gets an existing SecureSourceManagerInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecureSourceManagerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecureSourceManagerInstanceState, opts ...pulumi.ResourceOption) (*SecureSourceManagerInstance, error) {
	var resource SecureSourceManagerInstance
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/secureSourceManagerInstance:SecureSourceManagerInstance", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecureSourceManagerInstance resources.
type secureSourceManagerInstanceState struct {
	// Time the Instance was created in UTC.
	CreateTime      *string           `pulumi:"createTime"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// A list of hostnames for this instance.
	HostConfigs []SecureSourceManagerInstanceHostConfig `pulumi:"hostConfigs"`
	// The name for the Instance.
	InstanceId *string `pulumi:"instanceId"`
	// Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.
	KmsKey *string `pulumi:"kmsKey"`
	// Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the Instance.
	Location *string `pulumi:"location"`
	// The resource name for the Instance.
	Name *string `pulumi:"name"`
	// Private settings for private instance.
	PrivateConfig                 *SecureSourceManagerInstancePrivateConfig `pulumi:"privateConfig"`
	Project                       *string                                   `pulumi:"project"`
	SecureSourceManagerInstanceId *string                                   `pulumi:"secureSourceManagerInstanceId"`
	// The current state of the Instance.
	State *string `pulumi:"state"`
	// Provides information about the current instance state.
	StateNote *string `pulumi:"stateNote"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                    `pulumi:"terraformLabels"`
	Timeouts        *SecureSourceManagerInstanceTimeouts `pulumi:"timeouts"`
	// Time the Instance was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
	// Configuration for Workforce Identity Federation to support third party identity provider. If unset, defaults to the
	// Google OIDC IdP.
	WorkforceIdentityFederationConfig *SecureSourceManagerInstanceWorkforceIdentityFederationConfig `pulumi:"workforceIdentityFederationConfig"`
}

type SecureSourceManagerInstanceState struct {
	// Time the Instance was created in UTC.
	CreateTime      pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// A list of hostnames for this instance.
	HostConfigs SecureSourceManagerInstanceHostConfigArrayInput
	// The name for the Instance.
	InstanceId pulumi.StringPtrInput
	// Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.
	KmsKey pulumi.StringPtrInput
	// Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the Instance.
	Location pulumi.StringPtrInput
	// The resource name for the Instance.
	Name pulumi.StringPtrInput
	// Private settings for private instance.
	PrivateConfig                 SecureSourceManagerInstancePrivateConfigPtrInput
	Project                       pulumi.StringPtrInput
	SecureSourceManagerInstanceId pulumi.StringPtrInput
	// The current state of the Instance.
	State pulumi.StringPtrInput
	// Provides information about the current instance state.
	StateNote pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        SecureSourceManagerInstanceTimeoutsPtrInput
	// Time the Instance was updated in UTC.
	UpdateTime pulumi.StringPtrInput
	// Configuration for Workforce Identity Federation to support third party identity provider. If unset, defaults to the
	// Google OIDC IdP.
	WorkforceIdentityFederationConfig SecureSourceManagerInstanceWorkforceIdentityFederationConfigPtrInput
}

func (SecureSourceManagerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerInstanceState)(nil)).Elem()
}

type secureSourceManagerInstanceArgs struct {
	// The name for the Instance.
	InstanceId string `pulumi:"instanceId"`
	// Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.
	KmsKey *string `pulumi:"kmsKey"`
	// Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the Instance.
	Location string `pulumi:"location"`
	// Private settings for private instance.
	PrivateConfig                 *SecureSourceManagerInstancePrivateConfig `pulumi:"privateConfig"`
	Project                       *string                                   `pulumi:"project"`
	SecureSourceManagerInstanceId *string                                   `pulumi:"secureSourceManagerInstanceId"`
	Timeouts                      *SecureSourceManagerInstanceTimeouts      `pulumi:"timeouts"`
	// Configuration for Workforce Identity Federation to support third party identity provider. If unset, defaults to the
	// Google OIDC IdP.
	WorkforceIdentityFederationConfig *SecureSourceManagerInstanceWorkforceIdentityFederationConfig `pulumi:"workforceIdentityFederationConfig"`
}

// The set of arguments for constructing a SecureSourceManagerInstance resource.
type SecureSourceManagerInstanceArgs struct {
	// The name for the Instance.
	InstanceId pulumi.StringInput
	// Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.
	KmsKey pulumi.StringPtrInput
	// Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the Instance.
	Location pulumi.StringInput
	// Private settings for private instance.
	PrivateConfig                 SecureSourceManagerInstancePrivateConfigPtrInput
	Project                       pulumi.StringPtrInput
	SecureSourceManagerInstanceId pulumi.StringPtrInput
	Timeouts                      SecureSourceManagerInstanceTimeoutsPtrInput
	// Configuration for Workforce Identity Federation to support third party identity provider. If unset, defaults to the
	// Google OIDC IdP.
	WorkforceIdentityFederationConfig SecureSourceManagerInstanceWorkforceIdentityFederationConfigPtrInput
}

func (SecureSourceManagerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerInstanceArgs)(nil)).Elem()
}

type SecureSourceManagerInstanceInput interface {
	pulumi.Input

	ToSecureSourceManagerInstanceOutput() SecureSourceManagerInstanceOutput
	ToSecureSourceManagerInstanceOutputWithContext(ctx context.Context) SecureSourceManagerInstanceOutput
}

func (*SecureSourceManagerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerInstance)(nil)).Elem()
}

func (i *SecureSourceManagerInstance) ToSecureSourceManagerInstanceOutput() SecureSourceManagerInstanceOutput {
	return i.ToSecureSourceManagerInstanceOutputWithContext(context.Background())
}

func (i *SecureSourceManagerInstance) ToSecureSourceManagerInstanceOutputWithContext(ctx context.Context) SecureSourceManagerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureSourceManagerInstanceOutput)
}

type SecureSourceManagerInstanceOutput struct{ *pulumi.OutputState }

func (SecureSourceManagerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerInstance)(nil)).Elem()
}

func (o SecureSourceManagerInstanceOutput) ToSecureSourceManagerInstanceOutput() SecureSourceManagerInstanceOutput {
	return o
}

func (o SecureSourceManagerInstanceOutput) ToSecureSourceManagerInstanceOutputWithContext(ctx context.Context) SecureSourceManagerInstanceOutput {
	return o
}

// Time the Instance was created in UTC.
func (o SecureSourceManagerInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o SecureSourceManagerInstanceOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// A list of hostnames for this instance.
func (o SecureSourceManagerInstanceOutput) HostConfigs() SecureSourceManagerInstanceHostConfigArrayOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) SecureSourceManagerInstanceHostConfigArrayOutput {
		return v.HostConfigs
	}).(SecureSourceManagerInstanceHostConfigArrayOutput)
}

// The name for the Instance.
func (o SecureSourceManagerInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.
func (o SecureSourceManagerInstanceOutput) KmsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringPtrOutput { return v.KmsKey }).(pulumi.StringPtrOutput)
}

// Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present in your
// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o SecureSourceManagerInstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the Instance.
func (o SecureSourceManagerInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for the Instance.
func (o SecureSourceManagerInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private settings for private instance.
func (o SecureSourceManagerInstanceOutput) PrivateConfig() SecureSourceManagerInstancePrivateConfigPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) SecureSourceManagerInstancePrivateConfigPtrOutput {
		return v.PrivateConfig
	}).(SecureSourceManagerInstancePrivateConfigPtrOutput)
}

func (o SecureSourceManagerInstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SecureSourceManagerInstanceOutput) SecureSourceManagerInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.SecureSourceManagerInstanceId }).(pulumi.StringOutput)
}

// The current state of the Instance.
func (o SecureSourceManagerInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Provides information about the current instance state.
func (o SecureSourceManagerInstanceOutput) StateNote() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.StateNote }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o SecureSourceManagerInstanceOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o SecureSourceManagerInstanceOutput) Timeouts() SecureSourceManagerInstanceTimeoutsPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) SecureSourceManagerInstanceTimeoutsPtrOutput { return v.Timeouts }).(SecureSourceManagerInstanceTimeoutsPtrOutput)
}

// Time the Instance was updated in UTC.
func (o SecureSourceManagerInstanceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Configuration for Workforce Identity Federation to support third party identity provider. If unset, defaults to the
// Google OIDC IdP.
func (o SecureSourceManagerInstanceOutput) WorkforceIdentityFederationConfig() SecureSourceManagerInstanceWorkforceIdentityFederationConfigPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstance) SecureSourceManagerInstanceWorkforceIdentityFederationConfigPtrOutput {
		return v.WorkforceIdentityFederationConfig
	}).(SecureSourceManagerInstanceWorkforceIdentityFederationConfigPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecureSourceManagerInstanceInput)(nil)).Elem(), &SecureSourceManagerInstance{})
	pulumi.RegisterOutputType(SecureSourceManagerInstanceOutput{})
}
