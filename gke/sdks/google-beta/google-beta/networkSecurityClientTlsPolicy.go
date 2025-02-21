// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkSecurityClientTlsPolicy struct {
	pulumi.CustomResourceState

	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ClientCertificate NetworkSecurityClientTlsPolicyClientCertificatePtrOutput `pulumi:"clientCertificate"`
	// Time the ClientTlsPolicy was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Set of label tags associated with the ClientTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location of the client tls policy. The default value is 'global'.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the ClientTlsPolicy resource.
	Name                             pulumi.StringOutput `pulumi:"name"`
	NetworkSecurityClientTlsPolicyId pulumi.StringOutput `pulumi:"networkSecurityClientTlsPolicyId"`
	Project                          pulumi.StringOutput `pulumi:"project"`
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty,
	// client does not validate the server certificate.
	ServerValidationCas NetworkSecurityClientTlsPolicyServerValidationCaArrayOutput `pulumi:"serverValidationCas"`
	// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	Sni pulumi.StringPtrOutput `pulumi:"sni"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                          `pulumi:"terraformLabels"`
	Timeouts        NetworkSecurityClientTlsPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the ClientTlsPolicy was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecurityClientTlsPolicy registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityClientTlsPolicy(ctx *pulumi.Context,
	name string, args *NetworkSecurityClientTlsPolicyArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityClientTlsPolicy, error) {
	if args == nil {
		args = &NetworkSecurityClientTlsPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkSecurityClientTlsPolicy
	err = ctx.RegisterPackageResource("google-beta:index/networkSecurityClientTlsPolicy:NetworkSecurityClientTlsPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityClientTlsPolicy gets an existing NetworkSecurityClientTlsPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityClientTlsPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityClientTlsPolicyState, opts ...pulumi.ResourceOption) (*NetworkSecurityClientTlsPolicy, error) {
	var resource NetworkSecurityClientTlsPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkSecurityClientTlsPolicy:NetworkSecurityClientTlsPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityClientTlsPolicy resources.
type networkSecurityClientTlsPolicyState struct {
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ClientCertificate *NetworkSecurityClientTlsPolicyClientCertificate `pulumi:"clientCertificate"`
	// Time the ClientTlsPolicy was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Set of label tags associated with the ClientTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the client tls policy. The default value is 'global'.
	Location *string `pulumi:"location"`
	// Name of the ClientTlsPolicy resource.
	Name                             *string `pulumi:"name"`
	NetworkSecurityClientTlsPolicyId *string `pulumi:"networkSecurityClientTlsPolicyId"`
	Project                          *string `pulumi:"project"`
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty,
	// client does not validate the server certificate.
	ServerValidationCas []NetworkSecurityClientTlsPolicyServerValidationCa `pulumi:"serverValidationCas"`
	// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	Sni *string `pulumi:"sni"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                       `pulumi:"terraformLabels"`
	Timeouts        *NetworkSecurityClientTlsPolicyTimeouts `pulumi:"timeouts"`
	// Time the ClientTlsPolicy was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecurityClientTlsPolicyState struct {
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ClientCertificate NetworkSecurityClientTlsPolicyClientCertificatePtrInput
	// Time the ClientTlsPolicy was created in UTC.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Set of label tags associated with the ClientTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location of the client tls policy. The default value is 'global'.
	Location pulumi.StringPtrInput
	// Name of the ClientTlsPolicy resource.
	Name                             pulumi.StringPtrInput
	NetworkSecurityClientTlsPolicyId pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty,
	// client does not validate the server certificate.
	ServerValidationCas NetworkSecurityClientTlsPolicyServerValidationCaArrayInput
	// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	Sni pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkSecurityClientTlsPolicyTimeoutsPtrInput
	// Time the ClientTlsPolicy was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecurityClientTlsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityClientTlsPolicyState)(nil)).Elem()
}

type networkSecurityClientTlsPolicyArgs struct {
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ClientCertificate *NetworkSecurityClientTlsPolicyClientCertificate `pulumi:"clientCertificate"`
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the ClientTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the client tls policy. The default value is 'global'.
	Location *string `pulumi:"location"`
	// Name of the ClientTlsPolicy resource.
	Name                             *string `pulumi:"name"`
	NetworkSecurityClientTlsPolicyId *string `pulumi:"networkSecurityClientTlsPolicyId"`
	Project                          *string `pulumi:"project"`
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty,
	// client does not validate the server certificate.
	ServerValidationCas []NetworkSecurityClientTlsPolicyServerValidationCa `pulumi:"serverValidationCas"`
	// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	Sni      *string                                 `pulumi:"sni"`
	Timeouts *NetworkSecurityClientTlsPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkSecurityClientTlsPolicy resource.
type NetworkSecurityClientTlsPolicyArgs struct {
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ClientCertificate NetworkSecurityClientTlsPolicyClientCertificatePtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the ClientTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location of the client tls policy. The default value is 'global'.
	Location pulumi.StringPtrInput
	// Name of the ClientTlsPolicy resource.
	Name                             pulumi.StringPtrInput
	NetworkSecurityClientTlsPolicyId pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty,
	// client does not validate the server certificate.
	ServerValidationCas NetworkSecurityClientTlsPolicyServerValidationCaArrayInput
	// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	Sni      pulumi.StringPtrInput
	Timeouts NetworkSecurityClientTlsPolicyTimeoutsPtrInput
}

func (NetworkSecurityClientTlsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityClientTlsPolicyArgs)(nil)).Elem()
}

type NetworkSecurityClientTlsPolicyInput interface {
	pulumi.Input

	ToNetworkSecurityClientTlsPolicyOutput() NetworkSecurityClientTlsPolicyOutput
	ToNetworkSecurityClientTlsPolicyOutputWithContext(ctx context.Context) NetworkSecurityClientTlsPolicyOutput
}

func (*NetworkSecurityClientTlsPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityClientTlsPolicy)(nil)).Elem()
}

func (i *NetworkSecurityClientTlsPolicy) ToNetworkSecurityClientTlsPolicyOutput() NetworkSecurityClientTlsPolicyOutput {
	return i.ToNetworkSecurityClientTlsPolicyOutputWithContext(context.Background())
}

func (i *NetworkSecurityClientTlsPolicy) ToNetworkSecurityClientTlsPolicyOutputWithContext(ctx context.Context) NetworkSecurityClientTlsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityClientTlsPolicyOutput)
}

type NetworkSecurityClientTlsPolicyOutput struct{ *pulumi.OutputState }

func (NetworkSecurityClientTlsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityClientTlsPolicy)(nil)).Elem()
}

func (o NetworkSecurityClientTlsPolicyOutput) ToNetworkSecurityClientTlsPolicyOutput() NetworkSecurityClientTlsPolicyOutput {
	return o
}

func (o NetworkSecurityClientTlsPolicyOutput) ToNetworkSecurityClientTlsPolicyOutputWithContext(ctx context.Context) NetworkSecurityClientTlsPolicyOutput {
	return o
}

// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
// of this dictates mTLS.
func (o NetworkSecurityClientTlsPolicyOutput) ClientCertificate() NetworkSecurityClientTlsPolicyClientCertificatePtrOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) NetworkSecurityClientTlsPolicyClientCertificatePtrOutput {
		return v.ClientCertificate
	}).(NetworkSecurityClientTlsPolicyClientCertificatePtrOutput)
}

// Time the ClientTlsPolicy was created in UTC.
func (o NetworkSecurityClientTlsPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o NetworkSecurityClientTlsPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityClientTlsPolicyOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Set of label tags associated with the ClientTlsPolicy resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkSecurityClientTlsPolicyOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location of the client tls policy. The default value is 'global'.
func (o NetworkSecurityClientTlsPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the ClientTlsPolicy resource.
func (o NetworkSecurityClientTlsPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityClientTlsPolicyOutput) NetworkSecurityClientTlsPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringOutput { return v.NetworkSecurityClientTlsPolicyId }).(pulumi.StringOutput)
}

func (o NetworkSecurityClientTlsPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty,
// client does not validate the server certificate.
func (o NetworkSecurityClientTlsPolicyOutput) ServerValidationCas() NetworkSecurityClientTlsPolicyServerValidationCaArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) NetworkSecurityClientTlsPolicyServerValidationCaArrayOutput {
		return v.ServerValidationCas
	}).(NetworkSecurityClientTlsPolicyServerValidationCaArrayOutput)
}

// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
func (o NetworkSecurityClientTlsPolicyOutput) Sni() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringPtrOutput { return v.Sni }).(pulumi.StringPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecurityClientTlsPolicyOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityClientTlsPolicyOutput) Timeouts() NetworkSecurityClientTlsPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) NetworkSecurityClientTlsPolicyTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecurityClientTlsPolicyTimeoutsPtrOutput)
}

// Time the ClientTlsPolicy was updated in UTC.
func (o NetworkSecurityClientTlsPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityClientTlsPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityClientTlsPolicyInput)(nil)).Elem(), &NetworkSecurityClientTlsPolicy{})
	pulumi.RegisterOutputType(NetworkSecurityClientTlsPolicyOutput{})
}
