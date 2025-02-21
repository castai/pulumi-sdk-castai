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

type FolderAccessApprovalSettings struct {
	pulumi.CustomResourceState

	// The asymmetric crypto key version to use for signing approval requests. Empty active_key_version indicates that a
	// Google-managed key should be used for signing. This property will be ignored if set by an ancestor of the resource, and
	// new non-empty values may not be set.
	ActiveKeyVersion pulumi.StringPtrOutput `pulumi:"activeKeyVersion"`
	// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
	AncestorHasActiveKeyVersion pulumi.BoolOutput `pulumi:"ancestorHasActiveKeyVersion"`
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
	// of the Folder.
	EnrolledAncestor pulumi.BoolOutput `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices               FolderAccessApprovalSettingsEnrolledServiceArrayOutput `pulumi:"enrolledServices"`
	FolderAccessApprovalSettingsId pulumi.StringOutput                                    `pulumi:"folderAccessApprovalSettingsId"`
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// If the field is true, that indicates that there is some configuration issue with the active_key_version configured on
	// this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the correct permissions on it,
	// etc.) This key version is not necessarily the effective key version at this level, as key versions are inherited
	// top-down.
	InvalidKeyVersion pulumi.BoolOutput `pulumi:"invalidKeyVersion"`
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayOutput                      `pulumi:"notificationEmails"`
	Timeouts           FolderAccessApprovalSettingsTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewFolderAccessApprovalSettings registers a new resource with the given unique name, arguments, and options.
func NewFolderAccessApprovalSettings(ctx *pulumi.Context,
	name string, args *FolderAccessApprovalSettingsArgs, opts ...pulumi.ResourceOption) (*FolderAccessApprovalSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnrolledServices == nil {
		return nil, errors.New("invalid value for required argument 'EnrolledServices'")
	}
	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FolderAccessApprovalSettings
	err = ctx.RegisterPackageResource("google-beta:index/folderAccessApprovalSettings:FolderAccessApprovalSettings", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderAccessApprovalSettings gets an existing FolderAccessApprovalSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderAccessApprovalSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderAccessApprovalSettingsState, opts ...pulumi.ResourceOption) (*FolderAccessApprovalSettings, error) {
	var resource FolderAccessApprovalSettings
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/folderAccessApprovalSettings:FolderAccessApprovalSettings", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderAccessApprovalSettings resources.
type folderAccessApprovalSettingsState struct {
	// The asymmetric crypto key version to use for signing approval requests. Empty active_key_version indicates that a
	// Google-managed key should be used for signing. This property will be ignored if set by an ancestor of the resource, and
	// new non-empty values may not be set.
	ActiveKeyVersion *string `pulumi:"activeKeyVersion"`
	// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
	AncestorHasActiveKeyVersion *bool `pulumi:"ancestorHasActiveKeyVersion"`
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
	// of the Folder.
	EnrolledAncestor *bool `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices               []FolderAccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	FolderAccessApprovalSettingsId *string                                       `pulumi:"folderAccessApprovalSettingsId"`
	// ID of the folder of the access approval settings.
	FolderId *string `pulumi:"folderId"`
	// If the field is true, that indicates that there is some configuration issue with the active_key_version configured on
	// this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the correct permissions on it,
	// etc.) This key version is not necessarily the effective key version at this level, as key versions are inherited
	// top-down.
	InvalidKeyVersion *bool `pulumi:"invalidKeyVersion"`
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name *string `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails []string                              `pulumi:"notificationEmails"`
	Timeouts           *FolderAccessApprovalSettingsTimeouts `pulumi:"timeouts"`
}

type FolderAccessApprovalSettingsState struct {
	// The asymmetric crypto key version to use for signing approval requests. Empty active_key_version indicates that a
	// Google-managed key should be used for signing. This property will be ignored if set by an ancestor of the resource, and
	// new non-empty values may not be set.
	ActiveKeyVersion pulumi.StringPtrInput
	// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
	AncestorHasActiveKeyVersion pulumi.BoolPtrInput
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
	// of the Folder.
	EnrolledAncestor pulumi.BoolPtrInput
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices               FolderAccessApprovalSettingsEnrolledServiceArrayInput
	FolderAccessApprovalSettingsId pulumi.StringPtrInput
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringPtrInput
	// If the field is true, that indicates that there is some configuration issue with the active_key_version configured on
	// this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the correct permissions on it,
	// etc.) This key version is not necessarily the effective key version at this level, as key versions are inherited
	// top-down.
	InvalidKeyVersion pulumi.BoolPtrInput
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name pulumi.StringPtrInput
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
	Timeouts           FolderAccessApprovalSettingsTimeoutsPtrInput
}

func (FolderAccessApprovalSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderAccessApprovalSettingsState)(nil)).Elem()
}

type folderAccessApprovalSettingsArgs struct {
	// The asymmetric crypto key version to use for signing approval requests. Empty active_key_version indicates that a
	// Google-managed key should be used for signing. This property will be ignored if set by an ancestor of the resource, and
	// new non-empty values may not be set.
	ActiveKeyVersion *string `pulumi:"activeKeyVersion"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices               []FolderAccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	FolderAccessApprovalSettingsId *string                                       `pulumi:"folderAccessApprovalSettingsId"`
	// ID of the folder of the access approval settings.
	FolderId string `pulumi:"folderId"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails []string                              `pulumi:"notificationEmails"`
	Timeouts           *FolderAccessApprovalSettingsTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a FolderAccessApprovalSettings resource.
type FolderAccessApprovalSettingsArgs struct {
	// The asymmetric crypto key version to use for signing approval requests. Empty active_key_version indicates that a
	// Google-managed key should be used for signing. This property will be ignored if set by an ancestor of the resource, and
	// new non-empty values may not be set.
	ActiveKeyVersion pulumi.StringPtrInput
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices               FolderAccessApprovalSettingsEnrolledServiceArrayInput
	FolderAccessApprovalSettingsId pulumi.StringPtrInput
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringInput
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
	Timeouts           FolderAccessApprovalSettingsTimeoutsPtrInput
}

func (FolderAccessApprovalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderAccessApprovalSettingsArgs)(nil)).Elem()
}

type FolderAccessApprovalSettingsInput interface {
	pulumi.Input

	ToFolderAccessApprovalSettingsOutput() FolderAccessApprovalSettingsOutput
	ToFolderAccessApprovalSettingsOutputWithContext(ctx context.Context) FolderAccessApprovalSettingsOutput
}

func (*FolderAccessApprovalSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderAccessApprovalSettings)(nil)).Elem()
}

func (i *FolderAccessApprovalSettings) ToFolderAccessApprovalSettingsOutput() FolderAccessApprovalSettingsOutput {
	return i.ToFolderAccessApprovalSettingsOutputWithContext(context.Background())
}

func (i *FolderAccessApprovalSettings) ToFolderAccessApprovalSettingsOutputWithContext(ctx context.Context) FolderAccessApprovalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderAccessApprovalSettingsOutput)
}

type FolderAccessApprovalSettingsOutput struct{ *pulumi.OutputState }

func (FolderAccessApprovalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderAccessApprovalSettings)(nil)).Elem()
}

func (o FolderAccessApprovalSettingsOutput) ToFolderAccessApprovalSettingsOutput() FolderAccessApprovalSettingsOutput {
	return o
}

func (o FolderAccessApprovalSettingsOutput) ToFolderAccessApprovalSettingsOutputWithContext(ctx context.Context) FolderAccessApprovalSettingsOutput {
	return o
}

// The asymmetric crypto key version to use for signing approval requests. Empty active_key_version indicates that a
// Google-managed key should be used for signing. This property will be ignored if set by an ancestor of the resource, and
// new non-empty values may not be set.
func (o FolderAccessApprovalSettingsOutput) ActiveKeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.StringPtrOutput { return v.ActiveKeyVersion }).(pulumi.StringPtrOutput)
}

// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
func (o FolderAccessApprovalSettingsOutput) AncestorHasActiveKeyVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.BoolOutput { return v.AncestorHasActiveKeyVersion }).(pulumi.BoolOutput)
}

// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
// of the Folder.
func (o FolderAccessApprovalSettingsOutput) EnrolledAncestor() pulumi.BoolOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.BoolOutput { return v.EnrolledAncestor }).(pulumi.BoolOutput)
}

// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
// resource given by name against any of these services contained here will be required to have explicit approval.
// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
// expanded as the set of supported services is expanded.
func (o FolderAccessApprovalSettingsOutput) EnrolledServices() FolderAccessApprovalSettingsEnrolledServiceArrayOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) FolderAccessApprovalSettingsEnrolledServiceArrayOutput {
		return v.EnrolledServices
	}).(FolderAccessApprovalSettingsEnrolledServiceArrayOutput)
}

func (o FolderAccessApprovalSettingsOutput) FolderAccessApprovalSettingsId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.StringOutput { return v.FolderAccessApprovalSettingsId }).(pulumi.StringOutput)
}

// ID of the folder of the access approval settings.
func (o FolderAccessApprovalSettingsOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// If the field is true, that indicates that there is some configuration issue with the active_key_version configured on
// this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the correct permissions on it,
// etc.) This key version is not necessarily the effective key version at this level, as key versions are inherited
// top-down.
func (o FolderAccessApprovalSettingsOutput) InvalidKeyVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.BoolOutput { return v.InvalidKeyVersion }).(pulumi.BoolOutput)
}

// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
func (o FolderAccessApprovalSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
// addresses are allowed.
func (o FolderAccessApprovalSettingsOutput) NotificationEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) pulumi.StringArrayOutput { return v.NotificationEmails }).(pulumi.StringArrayOutput)
}

func (o FolderAccessApprovalSettingsOutput) Timeouts() FolderAccessApprovalSettingsTimeoutsPtrOutput {
	return o.ApplyT(func(v *FolderAccessApprovalSettings) FolderAccessApprovalSettingsTimeoutsPtrOutput { return v.Timeouts }).(FolderAccessApprovalSettingsTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderAccessApprovalSettingsInput)(nil)).Elem(), &FolderAccessApprovalSettings{})
	pulumi.RegisterOutputType(FolderAccessApprovalSettingsOutput{})
}
