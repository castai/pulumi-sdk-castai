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

type Folder struct {
	pulumi.CustomResourceState

	// Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to
	// nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime         pulumi.StringOutput  `pulumi:"createTime"`
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the
	// same parent can share the same display name. The display name must start and end with a letter or digit, may contain
	// letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The folder id from the name "folders/{folder_id}"
	FolderId       pulumi.StringOutput `pulumi:"folderId"`
	GoogleFolderId pulumi.StringOutput `pulumi:"googleFolderId"`
	// The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource name of the parent Folder or Organization. Must be of the form folders/{folder_id} or
	// organizations/{org_id}.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. This field is only set at create time and modifying this field after creation will trigger recreation. To apply
	// tags to an existing resource, see the google.TagsTagValue resource.
	Tags     pulumi.StringMapOutput  `pulumi:"tags"`
	Timeouts FolderTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOption) (*Folder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource Folder
	err = ctx.RegisterPackageResource("google:index/folder:Folder", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderState, opts ...pulumi.ResourceOption) (*Folder, error) {
	var resource Folder
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/folder:Folder", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Folder resources.
type folderState struct {
	// Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to
	// nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime         *string `pulumi:"createTime"`
	DeletionProtection *bool   `pulumi:"deletionProtection"`
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the
	// same parent can share the same display name. The display name must start and end with a letter or digit, may contain
	// letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName *string `pulumi:"displayName"`
	// The folder id from the name "folders/{folder_id}"
	FolderId       *string `pulumi:"folderId"`
	GoogleFolderId *string `pulumi:"googleFolderId"`
	// The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.
	LifecycleState *string `pulumi:"lifecycleState"`
	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name *string `pulumi:"name"`
	// The resource name of the parent Folder or Organization. Must be of the form folders/{folder_id} or
	// organizations/{org_id}.
	Parent *string `pulumi:"parent"`
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. This field is only set at create time and modifying this field after creation will trigger recreation. To apply
	// tags to an existing resource, see the google.TagsTagValue resource.
	Tags     map[string]string `pulumi:"tags"`
	Timeouts *FolderTimeouts   `pulumi:"timeouts"`
}

type FolderState struct {
	// Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to
	// nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime         pulumi.StringPtrInput
	DeletionProtection pulumi.BoolPtrInput
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the
	// same parent can share the same display name. The display name must start and end with a letter or digit, may contain
	// letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName pulumi.StringPtrInput
	// The folder id from the name "folders/{folder_id}"
	FolderId       pulumi.StringPtrInput
	GoogleFolderId pulumi.StringPtrInput
	// The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.
	LifecycleState pulumi.StringPtrInput
	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name pulumi.StringPtrInput
	// The resource name of the parent Folder or Organization. Must be of the form folders/{folder_id} or
	// organizations/{org_id}.
	Parent pulumi.StringPtrInput
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. This field is only set at create time and modifying this field after creation will trigger recreation. To apply
	// tags to an existing resource, see the google.TagsTagValue resource.
	Tags     pulumi.StringMapInput
	Timeouts FolderTimeoutsPtrInput
}

func (FolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderState)(nil)).Elem()
}

type folderArgs struct {
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the
	// same parent can share the same display name. The display name must start and end with a letter or digit, may contain
	// letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName    string  `pulumi:"displayName"`
	GoogleFolderId *string `pulumi:"googleFolderId"`
	// The resource name of the parent Folder or Organization. Must be of the form folders/{folder_id} or
	// organizations/{org_id}.
	Parent string `pulumi:"parent"`
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. This field is only set at create time and modifying this field after creation will trigger recreation. To apply
	// tags to an existing resource, see the google.TagsTagValue resource.
	Tags     map[string]string `pulumi:"tags"`
	Timeouts *FolderTimeouts   `pulumi:"timeouts"`
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	DeletionProtection pulumi.BoolPtrInput
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the
	// same parent can share the same display name. The display name must start and end with a letter or digit, may contain
	// letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName    pulumi.StringInput
	GoogleFolderId pulumi.StringPtrInput
	// The resource name of the parent Folder or Organization. Must be of the form folders/{folder_id} or
	// organizations/{org_id}.
	Parent pulumi.StringInput
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. This field is only set at create time and modifying this field after creation will trigger recreation. To apply
	// tags to an existing resource, see the google.TagsTagValue resource.
	Tags     pulumi.StringMapInput
	Timeouts FolderTimeoutsPtrInput
}

func (FolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}

type FolderInput interface {
	pulumi.Input

	ToFolderOutput() FolderOutput
	ToFolderOutputWithContext(ctx context.Context) FolderOutput
}

func (*Folder) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (i *Folder) ToFolderOutput() FolderOutput {
	return i.ToFolderOutputWithContext(context.Background())
}

func (i *Folder) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderOutput)
}

type FolderOutput struct{ *pulumi.OutputState }

func (FolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (o FolderOutput) ToFolderOutput() FolderOutput {
	return o
}

func (o FolderOutput) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return o
}

// Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to
// nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o FolderOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o FolderOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Folder) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the
// same parent can share the same display name. The display name must start and end with a letter or digit, may contain
// letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
func (o FolderOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The folder id from the name "folders/{folder_id}"
func (o FolderOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

func (o FolderOutput) GoogleFolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.GoogleFolderId }).(pulumi.StringOutput)
}

// The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.
func (o FolderOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// The resource name of the Folder. Its format is folders/{folder_id}.
func (o FolderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource name of the parent Folder or Organization. Must be of the form folders/{folder_id} or
// organizations/{org_id}.
func (o FolderOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
// empty. This field is only set at create time and modifying this field after creation will trigger recreation. To apply
// tags to an existing resource, see the google.TagsTagValue resource.
func (o FolderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FolderOutput) Timeouts() FolderTimeoutsPtrOutput {
	return o.ApplyT(func(v *Folder) FolderTimeoutsPtrOutput { return v.Timeouts }).(FolderTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderInput)(nil)).Elem(), &Folder{})
	pulumi.RegisterOutputType(FolderOutput{})
}
