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

type FolderIamAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs FolderIamAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag                   pulumi.StringOutput `pulumi:"etag"`
	Folder                 pulumi.StringOutput `pulumi:"folder"`
	FolderIamAuditConfigId pulumi.StringOutput `pulumi:"folderIamAuditConfigId"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewFolderIamAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewFolderIamAuditConfig(ctx *pulumi.Context,
	name string, args *FolderIamAuditConfigArgs, opts ...pulumi.ResourceOption) (*FolderIamAuditConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditLogConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AuditLogConfigs'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FolderIamAuditConfig
	err = ctx.RegisterPackageResource("google-beta:index/folderIamAuditConfig:FolderIamAuditConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderIamAuditConfig gets an existing FolderIamAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderIamAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderIamAuditConfigState, opts ...pulumi.ResourceOption) (*FolderIamAuditConfig, error) {
	var resource FolderIamAuditConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/folderIamAuditConfig:FolderIamAuditConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderIamAuditConfig resources.
type folderIamAuditConfigState struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs []FolderIamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag                   *string `pulumi:"etag"`
	Folder                 *string `pulumi:"folder"`
	FolderIamAuditConfigId *string `pulumi:"folderIamAuditConfigId"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service *string `pulumi:"service"`
}

type FolderIamAuditConfigState struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs FolderIamAuditConfigAuditLogConfigArrayInput
	// The etag of iam policy
	Etag                   pulumi.StringPtrInput
	Folder                 pulumi.StringPtrInput
	FolderIamAuditConfigId pulumi.StringPtrInput
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringPtrInput
}

func (FolderIamAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderIamAuditConfigState)(nil)).Elem()
}

type folderIamAuditConfigArgs struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs        []FolderIamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	Folder                 string                               `pulumi:"folder"`
	FolderIamAuditConfigId *string                              `pulumi:"folderIamAuditConfigId"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a FolderIamAuditConfig resource.
type FolderIamAuditConfigArgs struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs        FolderIamAuditConfigAuditLogConfigArrayInput
	Folder                 pulumi.StringInput
	FolderIamAuditConfigId pulumi.StringPtrInput
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringInput
}

func (FolderIamAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderIamAuditConfigArgs)(nil)).Elem()
}

type FolderIamAuditConfigInput interface {
	pulumi.Input

	ToFolderIamAuditConfigOutput() FolderIamAuditConfigOutput
	ToFolderIamAuditConfigOutputWithContext(ctx context.Context) FolderIamAuditConfigOutput
}

func (*FolderIamAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderIamAuditConfig)(nil)).Elem()
}

func (i *FolderIamAuditConfig) ToFolderIamAuditConfigOutput() FolderIamAuditConfigOutput {
	return i.ToFolderIamAuditConfigOutputWithContext(context.Background())
}

func (i *FolderIamAuditConfig) ToFolderIamAuditConfigOutputWithContext(ctx context.Context) FolderIamAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderIamAuditConfigOutput)
}

type FolderIamAuditConfigOutput struct{ *pulumi.OutputState }

func (FolderIamAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderIamAuditConfig)(nil)).Elem()
}

func (o FolderIamAuditConfigOutput) ToFolderIamAuditConfigOutput() FolderIamAuditConfigOutput {
	return o
}

func (o FolderIamAuditConfigOutput) ToFolderIamAuditConfigOutputWithContext(ctx context.Context) FolderIamAuditConfigOutput {
	return o
}

// The configuration for logging of each type of permission. This can be specified multiple times.
func (o FolderIamAuditConfigOutput) AuditLogConfigs() FolderIamAuditConfigAuditLogConfigArrayOutput {
	return o.ApplyT(func(v *FolderIamAuditConfig) FolderIamAuditConfigAuditLogConfigArrayOutput { return v.AuditLogConfigs }).(FolderIamAuditConfigAuditLogConfigArrayOutput)
}

// The etag of iam policy
func (o FolderIamAuditConfigOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamAuditConfig) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FolderIamAuditConfigOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamAuditConfig) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

func (o FolderIamAuditConfigOutput) FolderIamAuditConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamAuditConfig) pulumi.StringOutput { return v.FolderIamAuditConfigId }).(pulumi.StringOutput)
}

// Service which will be enabled for audit logging. The special value allServices covers all services.
func (o FolderIamAuditConfigOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamAuditConfig) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderIamAuditConfigInput)(nil)).Elem(), &FolderIamAuditConfig{})
	pulumi.RegisterOutputType(FolderIamAuditConfigOutput{})
}
