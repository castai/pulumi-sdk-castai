// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkSecurityServerTlsPolicy struct {
	pulumi.CustomResourceState

	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer
	// policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By
	// default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and
	// mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to
	// confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS
	// and non-TLS traffic reaching port :80.
	AllowOpen pulumi.BoolPtrOutput `pulumi:"allowOpen"`
	// Time the ServerTlsPolicy was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Set of label tags associated with the ServerTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location of the server tls policy. The default value is 'global'.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic
	// Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS -
	// mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If
	// allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	MtlsPolicy NetworkSecurityServerTlsPolicyMtlsPolicyPtrOutput `pulumi:"mtlsPolicy"`
	// Name of the ServerTlsPolicy resource.
	Name                             pulumi.StringOutput `pulumi:"name"`
	NetworkSecurityServerTlsPolicyId pulumi.StringOutput `pulumi:"networkSecurityServerTlsPolicyId"`
	Project                          pulumi.StringOutput `pulumi:"project"`
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ServerCertificate NetworkSecurityServerTlsPolicyServerCertificatePtrOutput `pulumi:"serverCertificate"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                          `pulumi:"terraformLabels"`
	Timeouts        NetworkSecurityServerTlsPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the ServerTlsPolicy was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecurityServerTlsPolicy registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityServerTlsPolicy(ctx *pulumi.Context,
	name string, args *NetworkSecurityServerTlsPolicyArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityServerTlsPolicy, error) {
	if args == nil {
		args = &NetworkSecurityServerTlsPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkSecurityServerTlsPolicy
	err = ctx.RegisterPackageResource("google-beta:index/networkSecurityServerTlsPolicy:NetworkSecurityServerTlsPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityServerTlsPolicy gets an existing NetworkSecurityServerTlsPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityServerTlsPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityServerTlsPolicyState, opts ...pulumi.ResourceOption) (*NetworkSecurityServerTlsPolicy, error) {
	var resource NetworkSecurityServerTlsPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkSecurityServerTlsPolicy:NetworkSecurityServerTlsPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityServerTlsPolicy resources.
type networkSecurityServerTlsPolicyState struct {
	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer
	// policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By
	// default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and
	// mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to
	// confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS
	// and non-TLS traffic reaching port :80.
	AllowOpen *bool `pulumi:"allowOpen"`
	// Time the ServerTlsPolicy was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Set of label tags associated with the ServerTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the server tls policy. The default value is 'global'.
	Location *string `pulumi:"location"`
	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic
	// Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS -
	// mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If
	// allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	MtlsPolicy *NetworkSecurityServerTlsPolicyMtlsPolicy `pulumi:"mtlsPolicy"`
	// Name of the ServerTlsPolicy resource.
	Name                             *string `pulumi:"name"`
	NetworkSecurityServerTlsPolicyId *string `pulumi:"networkSecurityServerTlsPolicyId"`
	Project                          *string `pulumi:"project"`
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ServerCertificate *NetworkSecurityServerTlsPolicyServerCertificate `pulumi:"serverCertificate"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                       `pulumi:"terraformLabels"`
	Timeouts        *NetworkSecurityServerTlsPolicyTimeouts `pulumi:"timeouts"`
	// Time the ServerTlsPolicy was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecurityServerTlsPolicyState struct {
	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer
	// policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By
	// default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and
	// mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to
	// confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS
	// and non-TLS traffic reaching port :80.
	AllowOpen pulumi.BoolPtrInput
	// Time the ServerTlsPolicy was created in UTC.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Set of label tags associated with the ServerTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location of the server tls policy. The default value is 'global'.
	Location pulumi.StringPtrInput
	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic
	// Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS -
	// mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If
	// allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	MtlsPolicy NetworkSecurityServerTlsPolicyMtlsPolicyPtrInput
	// Name of the ServerTlsPolicy resource.
	Name                             pulumi.StringPtrInput
	NetworkSecurityServerTlsPolicyId pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ServerCertificate NetworkSecurityServerTlsPolicyServerCertificatePtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkSecurityServerTlsPolicyTimeoutsPtrInput
	// Time the ServerTlsPolicy was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecurityServerTlsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityServerTlsPolicyState)(nil)).Elem()
}

type networkSecurityServerTlsPolicyArgs struct {
	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer
	// policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By
	// default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and
	// mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to
	// confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS
	// and non-TLS traffic reaching port :80.
	AllowOpen *bool `pulumi:"allowOpen"`
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the ServerTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the server tls policy. The default value is 'global'.
	Location *string `pulumi:"location"`
	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic
	// Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS -
	// mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If
	// allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	MtlsPolicy *NetworkSecurityServerTlsPolicyMtlsPolicy `pulumi:"mtlsPolicy"`
	// Name of the ServerTlsPolicy resource.
	Name                             *string `pulumi:"name"`
	NetworkSecurityServerTlsPolicyId *string `pulumi:"networkSecurityServerTlsPolicyId"`
	Project                          *string `pulumi:"project"`
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ServerCertificate *NetworkSecurityServerTlsPolicyServerCertificate `pulumi:"serverCertificate"`
	Timeouts          *NetworkSecurityServerTlsPolicyTimeouts          `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkSecurityServerTlsPolicy resource.
type NetworkSecurityServerTlsPolicyArgs struct {
	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer
	// policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By
	// default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and
	// mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to
	// confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS
	// and non-TLS traffic reaching port :80.
	AllowOpen pulumi.BoolPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the ServerTlsPolicy resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location of the server tls policy. The default value is 'global'.
	Location pulumi.StringPtrInput
	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic
	// Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS -
	// mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If
	// allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	MtlsPolicy NetworkSecurityServerTlsPolicyMtlsPolicyPtrInput
	// Name of the ServerTlsPolicy resource.
	Name                             pulumi.StringPtrInput
	NetworkSecurityServerTlsPolicyId pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
	// of this dictates mTLS.
	ServerCertificate NetworkSecurityServerTlsPolicyServerCertificatePtrInput
	Timeouts          NetworkSecurityServerTlsPolicyTimeoutsPtrInput
}

func (NetworkSecurityServerTlsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityServerTlsPolicyArgs)(nil)).Elem()
}

type NetworkSecurityServerTlsPolicyInput interface {
	pulumi.Input

	ToNetworkSecurityServerTlsPolicyOutput() NetworkSecurityServerTlsPolicyOutput
	ToNetworkSecurityServerTlsPolicyOutputWithContext(ctx context.Context) NetworkSecurityServerTlsPolicyOutput
}

func (*NetworkSecurityServerTlsPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityServerTlsPolicy)(nil)).Elem()
}

func (i *NetworkSecurityServerTlsPolicy) ToNetworkSecurityServerTlsPolicyOutput() NetworkSecurityServerTlsPolicyOutput {
	return i.ToNetworkSecurityServerTlsPolicyOutputWithContext(context.Background())
}

func (i *NetworkSecurityServerTlsPolicy) ToNetworkSecurityServerTlsPolicyOutputWithContext(ctx context.Context) NetworkSecurityServerTlsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityServerTlsPolicyOutput)
}

type NetworkSecurityServerTlsPolicyOutput struct{ *pulumi.OutputState }

func (NetworkSecurityServerTlsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityServerTlsPolicy)(nil)).Elem()
}

func (o NetworkSecurityServerTlsPolicyOutput) ToNetworkSecurityServerTlsPolicyOutput() NetworkSecurityServerTlsPolicyOutput {
	return o
}

func (o NetworkSecurityServerTlsPolicyOutput) ToNetworkSecurityServerTlsPolicyOutputWithContext(ctx context.Context) NetworkSecurityServerTlsPolicyOutput {
	return o
}

// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer
// policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By
// default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and
// mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to
// confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS
// and non-TLS traffic reaching port :80.
func (o NetworkSecurityServerTlsPolicyOutput) AllowOpen() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.BoolPtrOutput { return v.AllowOpen }).(pulumi.BoolPtrOutput)
}

// Time the ServerTlsPolicy was created in UTC.
func (o NetworkSecurityServerTlsPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o NetworkSecurityServerTlsPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityServerTlsPolicyOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Set of label tags associated with the ServerTlsPolicy resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkSecurityServerTlsPolicyOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location of the server tls policy. The default value is 'global'.
func (o NetworkSecurityServerTlsPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic
// Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS -
// mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If
// allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
func (o NetworkSecurityServerTlsPolicyOutput) MtlsPolicy() NetworkSecurityServerTlsPolicyMtlsPolicyPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) NetworkSecurityServerTlsPolicyMtlsPolicyPtrOutput {
		return v.MtlsPolicy
	}).(NetworkSecurityServerTlsPolicyMtlsPolicyPtrOutput)
}

// Name of the ServerTlsPolicy resource.
func (o NetworkSecurityServerTlsPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityServerTlsPolicyOutput) NetworkSecurityServerTlsPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringOutput { return v.NetworkSecurityServerTlsPolicyId }).(pulumi.StringOutput)
}

func (o NetworkSecurityServerTlsPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence
// of this dictates mTLS.
func (o NetworkSecurityServerTlsPolicyOutput) ServerCertificate() NetworkSecurityServerTlsPolicyServerCertificatePtrOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) NetworkSecurityServerTlsPolicyServerCertificatePtrOutput {
		return v.ServerCertificate
	}).(NetworkSecurityServerTlsPolicyServerCertificatePtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecurityServerTlsPolicyOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityServerTlsPolicyOutput) Timeouts() NetworkSecurityServerTlsPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) NetworkSecurityServerTlsPolicyTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecurityServerTlsPolicyTimeoutsPtrOutput)
}

// Time the ServerTlsPolicy was updated in UTC.
func (o NetworkSecurityServerTlsPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityServerTlsPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityServerTlsPolicyInput)(nil)).Elem(), &NetworkSecurityServerTlsPolicy{})
	pulumi.RegisterOutputType(NetworkSecurityServerTlsPolicyOutput{})
}
