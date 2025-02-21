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

type NetappBackupVault struct {
	pulumi.CustomResourceState

	// Create time of the backup vault. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// An optional description of this resource.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location (region) of the backup vault.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the backup vault. Needs to be unique per location.
	Name                pulumi.StringOutput `pulumi:"name"`
	NetappBackupVaultId pulumi.StringOutput `pulumi:"netappBackupVaultId"`
	Project             pulumi.StringOutput `pulumi:"project"`
	// The state of the Backup Vault.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput             `pulumi:"terraformLabels"`
	Timeouts        NetappBackupVaultTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewNetappBackupVault registers a new resource with the given unique name, arguments, and options.
func NewNetappBackupVault(ctx *pulumi.Context,
	name string, args *NetappBackupVaultArgs, opts ...pulumi.ResourceOption) (*NetappBackupVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetappBackupVault
	err = ctx.RegisterPackageResource("google-beta:index/netappBackupVault:NetappBackupVault", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetappBackupVault gets an existing NetappBackupVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetappBackupVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetappBackupVaultState, opts ...pulumi.ResourceOption) (*NetappBackupVault, error) {
	var resource NetappBackupVault
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/netappBackupVault:NetappBackupVault", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetappBackupVault resources.
type netappBackupVaultState struct {
	// Create time of the backup vault. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	CreateTime *string `pulumi:"createTime"`
	// An optional description of this resource.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location (region) of the backup vault.
	Location *string `pulumi:"location"`
	// The resource name of the backup vault. Needs to be unique per location.
	Name                *string `pulumi:"name"`
	NetappBackupVaultId *string `pulumi:"netappBackupVaultId"`
	Project             *string `pulumi:"project"`
	// The state of the Backup Vault.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string          `pulumi:"terraformLabels"`
	Timeouts        *NetappBackupVaultTimeouts `pulumi:"timeouts"`
}

type NetappBackupVaultState struct {
	// Create time of the backup vault. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	CreateTime pulumi.StringPtrInput
	// An optional description of this resource.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location (region) of the backup vault.
	Location pulumi.StringPtrInput
	// The resource name of the backup vault. Needs to be unique per location.
	Name                pulumi.StringPtrInput
	NetappBackupVaultId pulumi.StringPtrInput
	Project             pulumi.StringPtrInput
	// The state of the Backup Vault.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetappBackupVaultTimeoutsPtrInput
}

func (NetappBackupVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*netappBackupVaultState)(nil)).Elem()
}

type netappBackupVaultArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location (region) of the backup vault.
	Location string `pulumi:"location"`
	// The resource name of the backup vault. Needs to be unique per location.
	Name                *string                    `pulumi:"name"`
	NetappBackupVaultId *string                    `pulumi:"netappBackupVaultId"`
	Project             *string                    `pulumi:"project"`
	Timeouts            *NetappBackupVaultTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetappBackupVault resource.
type NetappBackupVaultArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location (region) of the backup vault.
	Location pulumi.StringInput
	// The resource name of the backup vault. Needs to be unique per location.
	Name                pulumi.StringPtrInput
	NetappBackupVaultId pulumi.StringPtrInput
	Project             pulumi.StringPtrInput
	Timeouts            NetappBackupVaultTimeoutsPtrInput
}

func (NetappBackupVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*netappBackupVaultArgs)(nil)).Elem()
}

type NetappBackupVaultInput interface {
	pulumi.Input

	ToNetappBackupVaultOutput() NetappBackupVaultOutput
	ToNetappBackupVaultOutputWithContext(ctx context.Context) NetappBackupVaultOutput
}

func (*NetappBackupVault) ElementType() reflect.Type {
	return reflect.TypeOf((**NetappBackupVault)(nil)).Elem()
}

func (i *NetappBackupVault) ToNetappBackupVaultOutput() NetappBackupVaultOutput {
	return i.ToNetappBackupVaultOutputWithContext(context.Background())
}

func (i *NetappBackupVault) ToNetappBackupVaultOutputWithContext(ctx context.Context) NetappBackupVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetappBackupVaultOutput)
}

type NetappBackupVaultOutput struct{ *pulumi.OutputState }

func (NetappBackupVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetappBackupVault)(nil)).Elem()
}

func (o NetappBackupVaultOutput) ToNetappBackupVaultOutput() NetappBackupVaultOutput {
	return o
}

func (o NetappBackupVaultOutput) ToNetappBackupVaultOutputWithContext(ctx context.Context) NetappBackupVaultOutput {
	return o
}

// Create time of the backup vault. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
func (o NetappBackupVaultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o NetappBackupVaultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetappBackupVaultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
// 'effective_labels' for all of the labels present on the resource.
func (o NetappBackupVaultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Location (region) of the backup vault.
func (o NetappBackupVaultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the backup vault. Needs to be unique per location.
func (o NetappBackupVaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetappBackupVaultOutput) NetappBackupVaultId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringOutput { return v.NetappBackupVaultId }).(pulumi.StringOutput)
}

func (o NetappBackupVaultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The state of the Backup Vault.
func (o NetappBackupVaultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetappBackupVaultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetappBackupVault) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetappBackupVaultOutput) Timeouts() NetappBackupVaultTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetappBackupVault) NetappBackupVaultTimeoutsPtrOutput { return v.Timeouts }).(NetappBackupVaultTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetappBackupVaultInput)(nil)).Elem(), &NetappBackupVault{})
	pulumi.RegisterOutputType(NetappBackupVaultOutput{})
}
