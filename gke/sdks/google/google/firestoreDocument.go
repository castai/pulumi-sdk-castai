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

type FirestoreDocument struct {
	pulumi.CustomResourceState

	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection pulumi.StringOutput `pulumi:"collection"`
	// Creation timestamp in RFC3339 format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// The client-assigned document ID to use for this document during creation.
	DocumentId pulumi.StringOutput `pulumi:"documentId"`
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated
	// as a json string.
	Fields              pulumi.StringOutput `pulumi:"fields"`
	FirestoreDocumentId pulumi.StringOutput `pulumi:"firestoreDocumentId"`
	// A server defined name for this document. Format:
	// 'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// A relative path to the collection this document exists within
	Path     pulumi.StringOutput                `pulumi:"path"`
	Project  pulumi.StringOutput                `pulumi:"project"`
	Timeouts FirestoreDocumentTimeoutsPtrOutput `pulumi:"timeouts"`
	// Last update timestamp in RFC3339 format.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewFirestoreDocument registers a new resource with the given unique name, arguments, and options.
func NewFirestoreDocument(ctx *pulumi.Context,
	name string, args *FirestoreDocumentArgs, opts ...pulumi.ResourceOption) (*FirestoreDocument, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Collection == nil {
		return nil, errors.New("invalid value for required argument 'Collection'")
	}
	if args.DocumentId == nil {
		return nil, errors.New("invalid value for required argument 'DocumentId'")
	}
	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirestoreDocument
	err = ctx.RegisterPackageResource("google:index/firestoreDocument:FirestoreDocument", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirestoreDocument gets an existing FirestoreDocument resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirestoreDocument(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirestoreDocumentState, opts ...pulumi.ResourceOption) (*FirestoreDocument, error) {
	var resource FirestoreDocument
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/firestoreDocument:FirestoreDocument", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirestoreDocument resources.
type firestoreDocumentState struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection *string `pulumi:"collection"`
	// Creation timestamp in RFC3339 format.
	CreateTime *string `pulumi:"createTime"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database *string `pulumi:"database"`
	// The client-assigned document ID to use for this document during creation.
	DocumentId *string `pulumi:"documentId"`
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated
	// as a json string.
	Fields              *string `pulumi:"fields"`
	FirestoreDocumentId *string `pulumi:"firestoreDocumentId"`
	// A server defined name for this document. Format:
	// 'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
	Name *string `pulumi:"name"`
	// A relative path to the collection this document exists within
	Path     *string                    `pulumi:"path"`
	Project  *string                    `pulumi:"project"`
	Timeouts *FirestoreDocumentTimeouts `pulumi:"timeouts"`
	// Last update timestamp in RFC3339 format.
	UpdateTime *string `pulumi:"updateTime"`
}

type FirestoreDocumentState struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection pulumi.StringPtrInput
	// Creation timestamp in RFC3339 format.
	CreateTime pulumi.StringPtrInput
	// The Firestore database id. Defaults to '"(default)"'.
	Database pulumi.StringPtrInput
	// The client-assigned document ID to use for this document during creation.
	DocumentId pulumi.StringPtrInput
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated
	// as a json string.
	Fields              pulumi.StringPtrInput
	FirestoreDocumentId pulumi.StringPtrInput
	// A server defined name for this document. Format:
	// 'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
	Name pulumi.StringPtrInput
	// A relative path to the collection this document exists within
	Path     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts FirestoreDocumentTimeoutsPtrInput
	// Last update timestamp in RFC3339 format.
	UpdateTime pulumi.StringPtrInput
}

func (FirestoreDocumentState) ElementType() reflect.Type {
	return reflect.TypeOf((*firestoreDocumentState)(nil)).Elem()
}

type firestoreDocumentArgs struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection string `pulumi:"collection"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database *string `pulumi:"database"`
	// The client-assigned document ID to use for this document during creation.
	DocumentId string `pulumi:"documentId"`
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated
	// as a json string.
	Fields              string                     `pulumi:"fields"`
	FirestoreDocumentId *string                    `pulumi:"firestoreDocumentId"`
	Project             *string                    `pulumi:"project"`
	Timeouts            *FirestoreDocumentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a FirestoreDocument resource.
type FirestoreDocumentArgs struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection pulumi.StringInput
	// The Firestore database id. Defaults to '"(default)"'.
	Database pulumi.StringPtrInput
	// The client-assigned document ID to use for this document during creation.
	DocumentId pulumi.StringInput
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated
	// as a json string.
	Fields              pulumi.StringInput
	FirestoreDocumentId pulumi.StringPtrInput
	Project             pulumi.StringPtrInput
	Timeouts            FirestoreDocumentTimeoutsPtrInput
}

func (FirestoreDocumentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firestoreDocumentArgs)(nil)).Elem()
}

type FirestoreDocumentInput interface {
	pulumi.Input

	ToFirestoreDocumentOutput() FirestoreDocumentOutput
	ToFirestoreDocumentOutputWithContext(ctx context.Context) FirestoreDocumentOutput
}

func (*FirestoreDocument) ElementType() reflect.Type {
	return reflect.TypeOf((**FirestoreDocument)(nil)).Elem()
}

func (i *FirestoreDocument) ToFirestoreDocumentOutput() FirestoreDocumentOutput {
	return i.ToFirestoreDocumentOutputWithContext(context.Background())
}

func (i *FirestoreDocument) ToFirestoreDocumentOutputWithContext(ctx context.Context) FirestoreDocumentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirestoreDocumentOutput)
}

type FirestoreDocumentOutput struct{ *pulumi.OutputState }

func (FirestoreDocumentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirestoreDocument)(nil)).Elem()
}

func (o FirestoreDocumentOutput) ToFirestoreDocumentOutput() FirestoreDocumentOutput {
	return o
}

func (o FirestoreDocumentOutput) ToFirestoreDocumentOutputWithContext(ctx context.Context) FirestoreDocumentOutput {
	return o
}

// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
func (o FirestoreDocumentOutput) Collection() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.Collection }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 format.
func (o FirestoreDocumentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The Firestore database id. Defaults to '"(default)"'.
func (o FirestoreDocumentOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringPtrOutput { return v.Database }).(pulumi.StringPtrOutput)
}

// The client-assigned document ID to use for this document during creation.
func (o FirestoreDocumentOutput) DocumentId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.DocumentId }).(pulumi.StringOutput)
}

// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated
// as a json string.
func (o FirestoreDocumentOutput) Fields() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.Fields }).(pulumi.StringOutput)
}

func (o FirestoreDocumentOutput) FirestoreDocumentId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.FirestoreDocumentId }).(pulumi.StringOutput)
}

// A server defined name for this document. Format:
// 'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
func (o FirestoreDocumentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A relative path to the collection this document exists within
func (o FirestoreDocumentOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o FirestoreDocumentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o FirestoreDocumentOutput) Timeouts() FirestoreDocumentTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirestoreDocument) FirestoreDocumentTimeoutsPtrOutput { return v.Timeouts }).(FirestoreDocumentTimeoutsPtrOutput)
}

// Last update timestamp in RFC3339 format.
func (o FirestoreDocumentOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreDocument) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirestoreDocumentInput)(nil)).Elem(), &FirestoreDocument{})
	pulumi.RegisterOutputType(FirestoreDocumentOutput{})
}
