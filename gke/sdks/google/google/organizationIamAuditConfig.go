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

type OrganizationIamAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs OrganizationIamAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                        pulumi.StringOutput `pulumi:"orgId"`
	OrganizationIamAuditConfigId pulumi.StringOutput `pulumi:"organizationIamAuditConfigId"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewOrganizationIamAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewOrganizationIamAuditConfig(ctx *pulumi.Context,
	name string, args *OrganizationIamAuditConfigArgs, opts ...pulumi.ResourceOption) (*OrganizationIamAuditConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditLogConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AuditLogConfigs'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource OrganizationIamAuditConfig
	err = ctx.RegisterPackageResource("google:index/organizationIamAuditConfig:OrganizationIamAuditConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationIamAuditConfig gets an existing OrganizationIamAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationIamAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationIamAuditConfigState, opts ...pulumi.ResourceOption) (*OrganizationIamAuditConfig, error) {
	var resource OrganizationIamAuditConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/organizationIamAuditConfig:OrganizationIamAuditConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationIamAuditConfig resources.
type organizationIamAuditConfigState struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs []OrganizationIamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag *string `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                        *string `pulumi:"orgId"`
	OrganizationIamAuditConfigId *string `pulumi:"organizationIamAuditConfigId"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service *string `pulumi:"service"`
}

type OrganizationIamAuditConfigState struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs OrganizationIamAuditConfigAuditLogConfigArrayInput
	// The etag of iam policy
	Etag pulumi.StringPtrInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                        pulumi.StringPtrInput
	OrganizationIamAuditConfigId pulumi.StringPtrInput
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringPtrInput
}

func (OrganizationIamAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationIamAuditConfigState)(nil)).Elem()
}

type organizationIamAuditConfigArgs struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs []OrganizationIamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                        string  `pulumi:"orgId"`
	OrganizationIamAuditConfigId *string `pulumi:"organizationIamAuditConfigId"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a OrganizationIamAuditConfig resource.
type OrganizationIamAuditConfigArgs struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs OrganizationIamAuditConfigAuditLogConfigArrayInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId                        pulumi.StringInput
	OrganizationIamAuditConfigId pulumi.StringPtrInput
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringInput
}

func (OrganizationIamAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationIamAuditConfigArgs)(nil)).Elem()
}

type OrganizationIamAuditConfigInput interface {
	pulumi.Input

	ToOrganizationIamAuditConfigOutput() OrganizationIamAuditConfigOutput
	ToOrganizationIamAuditConfigOutputWithContext(ctx context.Context) OrganizationIamAuditConfigOutput
}

func (*OrganizationIamAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationIamAuditConfig)(nil)).Elem()
}

func (i *OrganizationIamAuditConfig) ToOrganizationIamAuditConfigOutput() OrganizationIamAuditConfigOutput {
	return i.ToOrganizationIamAuditConfigOutputWithContext(context.Background())
}

func (i *OrganizationIamAuditConfig) ToOrganizationIamAuditConfigOutputWithContext(ctx context.Context) OrganizationIamAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationIamAuditConfigOutput)
}

type OrganizationIamAuditConfigOutput struct{ *pulumi.OutputState }

func (OrganizationIamAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationIamAuditConfig)(nil)).Elem()
}

func (o OrganizationIamAuditConfigOutput) ToOrganizationIamAuditConfigOutput() OrganizationIamAuditConfigOutput {
	return o
}

func (o OrganizationIamAuditConfigOutput) ToOrganizationIamAuditConfigOutputWithContext(ctx context.Context) OrganizationIamAuditConfigOutput {
	return o
}

// The configuration for logging of each type of permission. This can be specified multiple times.
func (o OrganizationIamAuditConfigOutput) AuditLogConfigs() OrganizationIamAuditConfigAuditLogConfigArrayOutput {
	return o.ApplyT(func(v *OrganizationIamAuditConfig) OrganizationIamAuditConfigAuditLogConfigArrayOutput {
		return v.AuditLogConfigs
	}).(OrganizationIamAuditConfigAuditLogConfigArrayOutput)
}

// The etag of iam policy
func (o OrganizationIamAuditConfigOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamAuditConfig) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The numeric ID of the organization in which you want to manage the audit logging config.
func (o OrganizationIamAuditConfigOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamAuditConfig) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

func (o OrganizationIamAuditConfigOutput) OrganizationIamAuditConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamAuditConfig) pulumi.StringOutput { return v.OrganizationIamAuditConfigId }).(pulumi.StringOutput)
}

// Service which will be enabled for audit logging. The special value allServices covers all services.
func (o OrganizationIamAuditConfigOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationIamAuditConfig) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationIamAuditConfigInput)(nil)).Elem(), &OrganizationIamAuditConfig{})
	pulumi.RegisterOutputType(OrganizationIamAuditConfigOutput{})
}
