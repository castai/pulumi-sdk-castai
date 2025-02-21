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

type EssentialContactsContact struct {
	pulumi.CustomResourceState

	// The email address to send notifications to. This does not need to be a Google account.
	Email                      pulumi.StringOutput `pulumi:"email"`
	EssentialContactsContactId pulumi.StringOutput `pulumi:"essentialContactsContactId"`
	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported
	// languages.
	LanguageTag pulumi.StringOutput `pulumi:"languageTag"`
	// The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions pulumi.StringArrayOutput `pulumi:"notificationCategorySubscriptions"`
	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or
	// projects/{project_id}
	Parent   pulumi.StringOutput                       `pulumi:"parent"`
	Timeouts EssentialContactsContactTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewEssentialContactsContact registers a new resource with the given unique name, arguments, and options.
func NewEssentialContactsContact(ctx *pulumi.Context,
	name string, args *EssentialContactsContactArgs, opts ...pulumi.ResourceOption) (*EssentialContactsContact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.LanguageTag == nil {
		return nil, errors.New("invalid value for required argument 'LanguageTag'")
	}
	if args.NotificationCategorySubscriptions == nil {
		return nil, errors.New("invalid value for required argument 'NotificationCategorySubscriptions'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource EssentialContactsContact
	err = ctx.RegisterPackageResource("google-beta:index/essentialContactsContact:EssentialContactsContact", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEssentialContactsContact gets an existing EssentialContactsContact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEssentialContactsContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EssentialContactsContactState, opts ...pulumi.ResourceOption) (*EssentialContactsContact, error) {
	var resource EssentialContactsContact
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/essentialContactsContact:EssentialContactsContact", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EssentialContactsContact resources.
type essentialContactsContactState struct {
	// The email address to send notifications to. This does not need to be a Google account.
	Email                      *string `pulumi:"email"`
	EssentialContactsContactId *string `pulumi:"essentialContactsContactId"`
	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported
	// languages.
	LanguageTag *string `pulumi:"languageTag"`
	// The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
	Name *string `pulumi:"name"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []string `pulumi:"notificationCategorySubscriptions"`
	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or
	// projects/{project_id}
	Parent   *string                           `pulumi:"parent"`
	Timeouts *EssentialContactsContactTimeouts `pulumi:"timeouts"`
}

type EssentialContactsContactState struct {
	// The email address to send notifications to. This does not need to be a Google account.
	Email                      pulumi.StringPtrInput
	EssentialContactsContactId pulumi.StringPtrInput
	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported
	// languages.
	LanguageTag pulumi.StringPtrInput
	// The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
	Name pulumi.StringPtrInput
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions pulumi.StringArrayInput
	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or
	// projects/{project_id}
	Parent   pulumi.StringPtrInput
	Timeouts EssentialContactsContactTimeoutsPtrInput
}

func (EssentialContactsContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*essentialContactsContactState)(nil)).Elem()
}

type essentialContactsContactArgs struct {
	// The email address to send notifications to. This does not need to be a Google account.
	Email                      string  `pulumi:"email"`
	EssentialContactsContactId *string `pulumi:"essentialContactsContactId"`
	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported
	// languages.
	LanguageTag string `pulumi:"languageTag"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []string `pulumi:"notificationCategorySubscriptions"`
	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or
	// projects/{project_id}
	Parent   string                            `pulumi:"parent"`
	Timeouts *EssentialContactsContactTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a EssentialContactsContact resource.
type EssentialContactsContactArgs struct {
	// The email address to send notifications to. This does not need to be a Google account.
	Email                      pulumi.StringInput
	EssentialContactsContactId pulumi.StringPtrInput
	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported
	// languages.
	LanguageTag pulumi.StringInput
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions pulumi.StringArrayInput
	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or
	// projects/{project_id}
	Parent   pulumi.StringInput
	Timeouts EssentialContactsContactTimeoutsPtrInput
}

func (EssentialContactsContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*essentialContactsContactArgs)(nil)).Elem()
}

type EssentialContactsContactInput interface {
	pulumi.Input

	ToEssentialContactsContactOutput() EssentialContactsContactOutput
	ToEssentialContactsContactOutputWithContext(ctx context.Context) EssentialContactsContactOutput
}

func (*EssentialContactsContact) ElementType() reflect.Type {
	return reflect.TypeOf((**EssentialContactsContact)(nil)).Elem()
}

func (i *EssentialContactsContact) ToEssentialContactsContactOutput() EssentialContactsContactOutput {
	return i.ToEssentialContactsContactOutputWithContext(context.Background())
}

func (i *EssentialContactsContact) ToEssentialContactsContactOutputWithContext(ctx context.Context) EssentialContactsContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EssentialContactsContactOutput)
}

type EssentialContactsContactOutput struct{ *pulumi.OutputState }

func (EssentialContactsContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EssentialContactsContact)(nil)).Elem()
}

func (o EssentialContactsContactOutput) ToEssentialContactsContactOutput() EssentialContactsContactOutput {
	return o
}

func (o EssentialContactsContactOutput) ToEssentialContactsContactOutputWithContext(ctx context.Context) EssentialContactsContactOutput {
	return o
}

// The email address to send notifications to. This does not need to be a Google account.
func (o EssentialContactsContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *EssentialContactsContact) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o EssentialContactsContactOutput) EssentialContactsContactId() pulumi.StringOutput {
	return o.ApplyT(func(v *EssentialContactsContact) pulumi.StringOutput { return v.EssentialContactsContactId }).(pulumi.StringOutput)
}

// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported
// languages.
func (o EssentialContactsContactOutput) LanguageTag() pulumi.StringOutput {
	return o.ApplyT(func(v *EssentialContactsContact) pulumi.StringOutput { return v.LanguageTag }).(pulumi.StringOutput)
}

// The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
func (o EssentialContactsContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EssentialContactsContact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The categories of notifications that the contact will receive communications for.
func (o EssentialContactsContactOutput) NotificationCategorySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EssentialContactsContact) pulumi.StringArrayOutput { return v.NotificationCategorySubscriptions }).(pulumi.StringArrayOutput)
}

// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or
// projects/{project_id}
func (o EssentialContactsContactOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *EssentialContactsContact) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

func (o EssentialContactsContactOutput) Timeouts() EssentialContactsContactTimeoutsPtrOutput {
	return o.ApplyT(func(v *EssentialContactsContact) EssentialContactsContactTimeoutsPtrOutput { return v.Timeouts }).(EssentialContactsContactTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EssentialContactsContactInput)(nil)).Elem(), &EssentialContactsContact{})
	pulumi.RegisterOutputType(EssentialContactsContactOutput{})
}
