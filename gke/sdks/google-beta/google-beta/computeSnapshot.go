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

type ComputeSnapshot struct {
	pulumi.CustomResourceState

	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters
	// long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate
	// snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible
	// only if it has a non-empty value.
	ChainName         pulumi.StringPtrOutput `pulumi:"chainName"`
	ComputeSnapshotId pulumi.StringOutput    `pulumi:"computeSnapshotId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Size of the snapshot, specified in GB.
	DiskSizeGb      pulumi.Float64Output   `pulumi:"diskSizeGb"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this Snapshot. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
	// attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
	// encryption key.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name     pulumi.StringOutput `pulumi:"name"`
	Project  pulumi.StringOutput `pulumi:"project"`
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied
	// key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when
	// you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect
	// access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the
	// snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot
	// later.
	SnapshotEncryptionKey ComputeSnapshotSnapshotEncryptionKeyPtrOutput `pulumi:"snapshotEncryptionKey"`
	// The unique identifier for the resource.
	SnapshotId pulumi.Float64Output `pulumi:"snapshotId"`
	// A reference to the disk used to create this snapshot.
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceDiskEncryptionKey ComputeSnapshotSourceDiskEncryptionKeyPtrOutput `pulumi:"sourceDiskEncryptionKey"`
	// A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot
	// creation/deletion.
	StorageBytes pulumi.Float64Output `pulumi:"storageBytes"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput           `pulumi:"terraformLabels"`
	Timeouts        ComputeSnapshotTimeoutsPtrOutput `pulumi:"timeouts"`
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputeSnapshot registers a new resource with the given unique name, arguments, and options.
func NewComputeSnapshot(ctx *pulumi.Context,
	name string, args *ComputeSnapshotArgs, opts ...pulumi.ResourceOption) (*ComputeSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceDisk == nil {
		return nil, errors.New("invalid value for required argument 'SourceDisk'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeSnapshot
	err = ctx.RegisterPackageResource("google-beta:index/computeSnapshot:ComputeSnapshot", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeSnapshot gets an existing ComputeSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeSnapshotState, opts ...pulumi.ResourceOption) (*ComputeSnapshot, error) {
	var resource ComputeSnapshot
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeSnapshot:ComputeSnapshot", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeSnapshot resources.
type computeSnapshotState struct {
	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters
	// long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate
	// snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible
	// only if it has a non-empty value.
	ChainName         *string `pulumi:"chainName"`
	ComputeSnapshotId *string `pulumi:"computeSnapshotId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Size of the snapshot, specified in GB.
	DiskSizeGb      *float64          `pulumi:"diskSizeGb"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this Snapshot. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
	// attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
	// encryption key.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name     *string `pulumi:"name"`
	Project  *string `pulumi:"project"`
	SelfLink *string `pulumi:"selfLink"`
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied
	// key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when
	// you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect
	// access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the
	// snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot
	// later.
	SnapshotEncryptionKey *ComputeSnapshotSnapshotEncryptionKey `pulumi:"snapshotEncryptionKey"`
	// The unique identifier for the resource.
	SnapshotId *float64 `pulumi:"snapshotId"`
	// A reference to the disk used to create this snapshot.
	SourceDisk *string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceDiskEncryptionKey *ComputeSnapshotSourceDiskEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	// A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot
	// creation/deletion.
	StorageBytes *float64 `pulumi:"storageBytes"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations []string `pulumi:"storageLocations"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string        `pulumi:"terraformLabels"`
	Timeouts        *ComputeSnapshotTimeouts `pulumi:"timeouts"`
	// A reference to the zone where the disk is hosted.
	Zone *string `pulumi:"zone"`
}

type ComputeSnapshotState struct {
	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters
	// long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate
	// snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible
	// only if it has a non-empty value.
	ChainName         pulumi.StringPtrInput
	ComputeSnapshotId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Size of the snapshot, specified in GB.
	DiskSizeGb      pulumi.Float64PtrInput
	EffectiveLabels pulumi.StringMapInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this Snapshot. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
	// attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
	// encryption key.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied
	// key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when
	// you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect
	// access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the
	// snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot
	// later.
	SnapshotEncryptionKey ComputeSnapshotSnapshotEncryptionKeyPtrInput
	// The unique identifier for the resource.
	SnapshotId pulumi.Float64PtrInput
	// A reference to the disk used to create this snapshot.
	SourceDisk pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceDiskEncryptionKey ComputeSnapshotSourceDiskEncryptionKeyPtrInput
	// A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot
	// creation/deletion.
	StorageBytes pulumi.Float64PtrInput
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ComputeSnapshotTimeoutsPtrInput
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringPtrInput
}

func (ComputeSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeSnapshotState)(nil)).Elem()
}

type computeSnapshotArgs struct {
	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters
	// long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate
	// snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible
	// only if it has a non-empty value.
	ChainName         *string `pulumi:"chainName"`
	ComputeSnapshotId *string `pulumi:"computeSnapshotId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Labels to apply to this Snapshot. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied
	// key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when
	// you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect
	// access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the
	// snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot
	// later.
	SnapshotEncryptionKey *ComputeSnapshotSnapshotEncryptionKey `pulumi:"snapshotEncryptionKey"`
	// A reference to the disk used to create this snapshot.
	SourceDisk string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceDiskEncryptionKey *ComputeSnapshotSourceDiskEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations []string                 `pulumi:"storageLocations"`
	Timeouts         *ComputeSnapshotTimeouts `pulumi:"timeouts"`
	// A reference to the zone where the disk is hosted.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeSnapshot resource.
type ComputeSnapshotArgs struct {
	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters
	// long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate
	// snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible
	// only if it has a non-empty value.
	ChainName         pulumi.StringPtrInput
	ComputeSnapshotId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Labels to apply to this Snapshot. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied
	// key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when
	// you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect
	// access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the
	// snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot
	// later.
	SnapshotEncryptionKey ComputeSnapshotSnapshotEncryptionKeyPtrInput
	// A reference to the disk used to create this snapshot.
	SourceDisk pulumi.StringInput
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceDiskEncryptionKey ComputeSnapshotSourceDiskEncryptionKeyPtrInput
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayInput
	Timeouts         ComputeSnapshotTimeoutsPtrInput
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringPtrInput
}

func (ComputeSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeSnapshotArgs)(nil)).Elem()
}

type ComputeSnapshotInput interface {
	pulumi.Input

	ToComputeSnapshotOutput() ComputeSnapshotOutput
	ToComputeSnapshotOutputWithContext(ctx context.Context) ComputeSnapshotOutput
}

func (*ComputeSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeSnapshot)(nil)).Elem()
}

func (i *ComputeSnapshot) ToComputeSnapshotOutput() ComputeSnapshotOutput {
	return i.ToComputeSnapshotOutputWithContext(context.Background())
}

func (i *ComputeSnapshot) ToComputeSnapshotOutputWithContext(ctx context.Context) ComputeSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeSnapshotOutput)
}

type ComputeSnapshotOutput struct{ *pulumi.OutputState }

func (ComputeSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeSnapshot)(nil)).Elem()
}

func (o ComputeSnapshotOutput) ToComputeSnapshotOutput() ComputeSnapshotOutput {
	return o
}

func (o ComputeSnapshotOutput) ToComputeSnapshotOutputWithContext(ctx context.Context) ComputeSnapshotOutput {
	return o
}

// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters
// long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate
// snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible
// only if it has a non-empty value.
func (o ComputeSnapshotOutput) ChainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringPtrOutput { return v.ChainName }).(pulumi.StringPtrOutput)
}

func (o ComputeSnapshotOutput) ComputeSnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.ComputeSnapshotId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeSnapshotOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeSnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Size of the snapshot, specified in GB.
func (o ComputeSnapshotOutput) DiskSizeGb() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.Float64Output { return v.DiskSizeGb }).(pulumi.Float64Output)
}

func (o ComputeSnapshotOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (o ComputeSnapshotOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this Snapshot. **Note**: This field is non-authoritative, and will only manage the labels present in
// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o ComputeSnapshotOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
// attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
// encryption key.
func (o ComputeSnapshotOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringArrayOutput { return v.Licenses }).(pulumi.StringArrayOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeSnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeSnapshotOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeSnapshotOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied
// key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when
// you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect
// access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the
// snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot
// later.
func (o ComputeSnapshotOutput) SnapshotEncryptionKey() ComputeSnapshotSnapshotEncryptionKeyPtrOutput {
	return o.ApplyT(func(v *ComputeSnapshot) ComputeSnapshotSnapshotEncryptionKeyPtrOutput { return v.SnapshotEncryptionKey }).(ComputeSnapshotSnapshotEncryptionKeyPtrOutput)
}

// The unique identifier for the resource.
func (o ComputeSnapshotOutput) SnapshotId() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.Float64Output { return v.SnapshotId }).(pulumi.Float64Output)
}

// A reference to the disk used to create this snapshot.
func (o ComputeSnapshotOutput) SourceDisk() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.SourceDisk }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
// customer-supplied encryption key.
func (o ComputeSnapshotOutput) SourceDiskEncryptionKey() ComputeSnapshotSourceDiskEncryptionKeyPtrOutput {
	return o.ApplyT(func(v *ComputeSnapshot) ComputeSnapshotSourceDiskEncryptionKeyPtrOutput {
		return v.SourceDiskEncryptionKey
	}).(ComputeSnapshotSourceDiskEncryptionKeyPtrOutput)
}

// A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot
// creation/deletion.
func (o ComputeSnapshotOutput) StorageBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.Float64Output { return v.StorageBytes }).(pulumi.Float64Output)
}

// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
func (o ComputeSnapshotOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringArrayOutput { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ComputeSnapshotOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ComputeSnapshotOutput) Timeouts() ComputeSnapshotTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeSnapshot) ComputeSnapshotTimeoutsPtrOutput { return v.Timeouts }).(ComputeSnapshotTimeoutsPtrOutput)
}

// A reference to the zone where the disk is hosted.
func (o ComputeSnapshotOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeSnapshot) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeSnapshotInput)(nil)).Elem(), &ComputeSnapshot{})
	pulumi.RegisterOutputType(ComputeSnapshotOutput{})
}
