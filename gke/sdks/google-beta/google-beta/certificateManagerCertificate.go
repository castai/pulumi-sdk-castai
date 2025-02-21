// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CertificateManagerCertificate struct {
	pulumi.CustomResourceState

	CertificateManagerCertificateId pulumi.StringOutput `pulumi:"certificateManagerCertificateId"`
	// A human-readable description of the resource.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Set of label tags associated with the Certificate resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The Certificate Manager location. If not specified, "global" is used.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Configuration and state of a Managed Certificate. Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	Managed CertificateManagerCertificateManagedPtrOutput `pulumi:"managed"`
	// A user-defined name of the certificate. Certificate names must be unique The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The list of Subject Alternative Names of dnsName type defined in the certificate (see RFC 5280 4.2.1.6)
	SanDnsnames pulumi.StringArrayOutput `pulumi:"sanDnsnames"`
	// The scope of the certificate. DEFAULT: Certificates with default scope are served from core Google data centers. If
	// unsure, choose this option. EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served
	// from Edge Points of Presence. See https://cloud.google.com/vpc/docs/edge-locations. ALL_REGIONS: Certificates with
	// ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs). See
	// https://cloud.google.com/compute/docs/regions-zones
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Certificate data for a SelfManaged Certificate. SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	SelfManaged CertificateManagerCertificateSelfManagedPtrOutput `pulumi:"selfManaged"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                         `pulumi:"terraformLabels"`
	Timeouts        CertificateManagerCertificateTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewCertificateManagerCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificateManagerCertificate(ctx *pulumi.Context,
	name string, args *CertificateManagerCertificateArgs, opts ...pulumi.ResourceOption) (*CertificateManagerCertificate, error) {
	if args == nil {
		args = &CertificateManagerCertificateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CertificateManagerCertificate
	err = ctx.RegisterPackageResource("google-beta:index/certificateManagerCertificate:CertificateManagerCertificate", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateManagerCertificate gets an existing CertificateManagerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateManagerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateManagerCertificateState, opts ...pulumi.ResourceOption) (*CertificateManagerCertificate, error) {
	var resource CertificateManagerCertificate
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/certificateManagerCertificate:CertificateManagerCertificate", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateManagerCertificate resources.
type certificateManagerCertificateState struct {
	CertificateManagerCertificateId *string `pulumi:"certificateManagerCertificateId"`
	// A human-readable description of the resource.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Set of label tags associated with the Certificate resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The Certificate Manager location. If not specified, "global" is used.
	Location *string `pulumi:"location"`
	// Configuration and state of a Managed Certificate. Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	Managed *CertificateManagerCertificateManaged `pulumi:"managed"`
	// A user-defined name of the certificate. Certificate names must be unique The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The list of Subject Alternative Names of dnsName type defined in the certificate (see RFC 5280 4.2.1.6)
	SanDnsnames []string `pulumi:"sanDnsnames"`
	// The scope of the certificate. DEFAULT: Certificates with default scope are served from core Google data centers. If
	// unsure, choose this option. EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served
	// from Edge Points of Presence. See https://cloud.google.com/vpc/docs/edge-locations. ALL_REGIONS: Certificates with
	// ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs). See
	// https://cloud.google.com/compute/docs/regions-zones
	Scope *string `pulumi:"scope"`
	// Certificate data for a SelfManaged Certificate. SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	SelfManaged *CertificateManagerCertificateSelfManaged `pulumi:"selfManaged"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                      `pulumi:"terraformLabels"`
	Timeouts        *CertificateManagerCertificateTimeouts `pulumi:"timeouts"`
}

type CertificateManagerCertificateState struct {
	CertificateManagerCertificateId pulumi.StringPtrInput
	// A human-readable description of the resource.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Set of label tags associated with the Certificate resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The Certificate Manager location. If not specified, "global" is used.
	Location pulumi.StringPtrInput
	// Configuration and state of a Managed Certificate. Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	Managed CertificateManagerCertificateManagedPtrInput
	// A user-defined name of the certificate. Certificate names must be unique The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The list of Subject Alternative Names of dnsName type defined in the certificate (see RFC 5280 4.2.1.6)
	SanDnsnames pulumi.StringArrayInput
	// The scope of the certificate. DEFAULT: Certificates with default scope are served from core Google data centers. If
	// unsure, choose this option. EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served
	// from Edge Points of Presence. See https://cloud.google.com/vpc/docs/edge-locations. ALL_REGIONS: Certificates with
	// ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs). See
	// https://cloud.google.com/compute/docs/regions-zones
	Scope pulumi.StringPtrInput
	// Certificate data for a SelfManaged Certificate. SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	SelfManaged CertificateManagerCertificateSelfManagedPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        CertificateManagerCertificateTimeoutsPtrInput
}

func (CertificateManagerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateManagerCertificateState)(nil)).Elem()
}

type certificateManagerCertificateArgs struct {
	CertificateManagerCertificateId *string `pulumi:"certificateManagerCertificateId"`
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the Certificate resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The Certificate Manager location. If not specified, "global" is used.
	Location *string `pulumi:"location"`
	// Configuration and state of a Managed Certificate. Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	Managed *CertificateManagerCertificateManaged `pulumi:"managed"`
	// A user-defined name of the certificate. Certificate names must be unique The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The scope of the certificate. DEFAULT: Certificates with default scope are served from core Google data centers. If
	// unsure, choose this option. EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served
	// from Edge Points of Presence. See https://cloud.google.com/vpc/docs/edge-locations. ALL_REGIONS: Certificates with
	// ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs). See
	// https://cloud.google.com/compute/docs/regions-zones
	Scope *string `pulumi:"scope"`
	// Certificate data for a SelfManaged Certificate. SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	SelfManaged *CertificateManagerCertificateSelfManaged `pulumi:"selfManaged"`
	Timeouts    *CertificateManagerCertificateTimeouts    `pulumi:"timeouts"`
}

// The set of arguments for constructing a CertificateManagerCertificate resource.
type CertificateManagerCertificateArgs struct {
	CertificateManagerCertificateId pulumi.StringPtrInput
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the Certificate resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The Certificate Manager location. If not specified, "global" is used.
	Location pulumi.StringPtrInput
	// Configuration and state of a Managed Certificate. Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	Managed CertificateManagerCertificateManagedPtrInput
	// A user-defined name of the certificate. Certificate names must be unique The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The scope of the certificate. DEFAULT: Certificates with default scope are served from core Google data centers. If
	// unsure, choose this option. EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served
	// from Edge Points of Presence. See https://cloud.google.com/vpc/docs/edge-locations. ALL_REGIONS: Certificates with
	// ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs). See
	// https://cloud.google.com/compute/docs/regions-zones
	Scope pulumi.StringPtrInput
	// Certificate data for a SelfManaged Certificate. SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	SelfManaged CertificateManagerCertificateSelfManagedPtrInput
	Timeouts    CertificateManagerCertificateTimeoutsPtrInput
}

func (CertificateManagerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateManagerCertificateArgs)(nil)).Elem()
}

type CertificateManagerCertificateInput interface {
	pulumi.Input

	ToCertificateManagerCertificateOutput() CertificateManagerCertificateOutput
	ToCertificateManagerCertificateOutputWithContext(ctx context.Context) CertificateManagerCertificateOutput
}

func (*CertificateManagerCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateManagerCertificate)(nil)).Elem()
}

func (i *CertificateManagerCertificate) ToCertificateManagerCertificateOutput() CertificateManagerCertificateOutput {
	return i.ToCertificateManagerCertificateOutputWithContext(context.Background())
}

func (i *CertificateManagerCertificate) ToCertificateManagerCertificateOutputWithContext(ctx context.Context) CertificateManagerCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateManagerCertificateOutput)
}

type CertificateManagerCertificateOutput struct{ *pulumi.OutputState }

func (CertificateManagerCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateManagerCertificate)(nil)).Elem()
}

func (o CertificateManagerCertificateOutput) ToCertificateManagerCertificateOutput() CertificateManagerCertificateOutput {
	return o
}

func (o CertificateManagerCertificateOutput) ToCertificateManagerCertificateOutputWithContext(ctx context.Context) CertificateManagerCertificateOutput {
	return o
}

func (o CertificateManagerCertificateOutput) CertificateManagerCertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringOutput { return v.CertificateManagerCertificateId }).(pulumi.StringOutput)
}

// A human-readable description of the resource.
func (o CertificateManagerCertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CertificateManagerCertificateOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Set of label tags associated with the Certificate resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o CertificateManagerCertificateOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The Certificate Manager location. If not specified, "global" is used.
func (o CertificateManagerCertificateOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Configuration and state of a Managed Certificate. Certificate Manager provisions and renews Managed Certificates
// automatically, for as long as it's authorized to do so.
func (o CertificateManagerCertificateOutput) Managed() CertificateManagerCertificateManagedPtrOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) CertificateManagerCertificateManagedPtrOutput { return v.Managed }).(CertificateManagerCertificateManagedPtrOutput)
}

// A user-defined name of the certificate. Certificate names must be unique The name must be 1-64 characters long, and
// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
// characters must be a dash, underscore, letter or digit.
func (o CertificateManagerCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateManagerCertificateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The list of Subject Alternative Names of dnsName type defined in the certificate (see RFC 5280 4.2.1.6)
func (o CertificateManagerCertificateOutput) SanDnsnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringArrayOutput { return v.SanDnsnames }).(pulumi.StringArrayOutput)
}

// The scope of the certificate. DEFAULT: Certificates with default scope are served from core Google data centers. If
// unsure, choose this option. EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served
// from Edge Points of Presence. See https://cloud.google.com/vpc/docs/edge-locations. ALL_REGIONS: Certificates with
// ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs). See
// https://cloud.google.com/compute/docs/regions-zones
func (o CertificateManagerCertificateOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Certificate data for a SelfManaged Certificate. SelfManaged Certificates are uploaded by the user. Updating such
// certificates before they expire remains the user's responsibility.
func (o CertificateManagerCertificateOutput) SelfManaged() CertificateManagerCertificateSelfManagedPtrOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) CertificateManagerCertificateSelfManagedPtrOutput {
		return v.SelfManaged
	}).(CertificateManagerCertificateSelfManagedPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o CertificateManagerCertificateOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o CertificateManagerCertificateOutput) Timeouts() CertificateManagerCertificateTimeoutsPtrOutput {
	return o.ApplyT(func(v *CertificateManagerCertificate) CertificateManagerCertificateTimeoutsPtrOutput {
		return v.Timeouts
	}).(CertificateManagerCertificateTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateManagerCertificateInput)(nil)).Elem(), &CertificateManagerCertificate{})
	pulumi.RegisterOutputType(CertificateManagerCertificateOutput{})
}
