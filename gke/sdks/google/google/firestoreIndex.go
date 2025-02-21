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

type FirestoreIndex struct {
	pulumi.CustomResourceState

	// The API scope at which a query is run. Default value: "ANY_API" Possible values: ["ANY_API", "DATASTORE_MODE_API"]
	ApiScope pulumi.StringPtrOutput `pulumi:"apiScope"`
	// The collection being indexed.
	Collection pulumi.StringOutput `pulumi:"collection"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// The fields supported by this index. The last non-stored field entry is always for the field path '__name__'. If, on
	// creation, '__name__' was not specified as the last field, it will be added automatically with the same direction as that
	// of the last field defined. If the final field in a composite index is not directional, the '__name__' will be ordered
	// '"ASCENDING"' (unless explicitly specified otherwise).
	Fields           FirestoreIndexFieldArrayOutput `pulumi:"fields"`
	FirestoreIndexId pulumi.StringOutput            `pulumi:"firestoreIndexId"`
	// A server defined name for this index. Format:
	// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP",
	// "COLLECTION_RECURSIVE"]
	QueryScope pulumi.StringPtrOutput          `pulumi:"queryScope"`
	Timeouts   FirestoreIndexTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewFirestoreIndex registers a new resource with the given unique name, arguments, and options.
func NewFirestoreIndex(ctx *pulumi.Context,
	name string, args *FirestoreIndexArgs, opts ...pulumi.ResourceOption) (*FirestoreIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Collection == nil {
		return nil, errors.New("invalid value for required argument 'Collection'")
	}
	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirestoreIndex
	err = ctx.RegisterPackageResource("google:index/firestoreIndex:FirestoreIndex", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirestoreIndex gets an existing FirestoreIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirestoreIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirestoreIndexState, opts ...pulumi.ResourceOption) (*FirestoreIndex, error) {
	var resource FirestoreIndex
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/firestoreIndex:FirestoreIndex", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirestoreIndex resources.
type firestoreIndexState struct {
	// The API scope at which a query is run. Default value: "ANY_API" Possible values: ["ANY_API", "DATASTORE_MODE_API"]
	ApiScope *string `pulumi:"apiScope"`
	// The collection being indexed.
	Collection *string `pulumi:"collection"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database *string `pulumi:"database"`
	// The fields supported by this index. The last non-stored field entry is always for the field path '__name__'. If, on
	// creation, '__name__' was not specified as the last field, it will be added automatically with the same direction as that
	// of the last field defined. If the final field in a composite index is not directional, the '__name__' will be ordered
	// '"ASCENDING"' (unless explicitly specified otherwise).
	Fields           []FirestoreIndexField `pulumi:"fields"`
	FirestoreIndexId *string               `pulumi:"firestoreIndexId"`
	// A server defined name for this index. Format:
	// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP",
	// "COLLECTION_RECURSIVE"]
	QueryScope *string                 `pulumi:"queryScope"`
	Timeouts   *FirestoreIndexTimeouts `pulumi:"timeouts"`
}

type FirestoreIndexState struct {
	// The API scope at which a query is run. Default value: "ANY_API" Possible values: ["ANY_API", "DATASTORE_MODE_API"]
	ApiScope pulumi.StringPtrInput
	// The collection being indexed.
	Collection pulumi.StringPtrInput
	// The Firestore database id. Defaults to '"(default)"'.
	Database pulumi.StringPtrInput
	// The fields supported by this index. The last non-stored field entry is always for the field path '__name__'. If, on
	// creation, '__name__' was not specified as the last field, it will be added automatically with the same direction as that
	// of the last field defined. If the final field in a composite index is not directional, the '__name__' will be ordered
	// '"ASCENDING"' (unless explicitly specified otherwise).
	Fields           FirestoreIndexFieldArrayInput
	FirestoreIndexId pulumi.StringPtrInput
	// A server defined name for this index. Format:
	// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP",
	// "COLLECTION_RECURSIVE"]
	QueryScope pulumi.StringPtrInput
	Timeouts   FirestoreIndexTimeoutsPtrInput
}

func (FirestoreIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*firestoreIndexState)(nil)).Elem()
}

type firestoreIndexArgs struct {
	// The API scope at which a query is run. Default value: "ANY_API" Possible values: ["ANY_API", "DATASTORE_MODE_API"]
	ApiScope *string `pulumi:"apiScope"`
	// The collection being indexed.
	Collection string `pulumi:"collection"`
	// The Firestore database id. Defaults to '"(default)"'.
	Database *string `pulumi:"database"`
	// The fields supported by this index. The last non-stored field entry is always for the field path '__name__'. If, on
	// creation, '__name__' was not specified as the last field, it will be added automatically with the same direction as that
	// of the last field defined. If the final field in a composite index is not directional, the '__name__' will be ordered
	// '"ASCENDING"' (unless explicitly specified otherwise).
	Fields           []FirestoreIndexField `pulumi:"fields"`
	FirestoreIndexId *string               `pulumi:"firestoreIndexId"`
	Project          *string               `pulumi:"project"`
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP",
	// "COLLECTION_RECURSIVE"]
	QueryScope *string                 `pulumi:"queryScope"`
	Timeouts   *FirestoreIndexTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a FirestoreIndex resource.
type FirestoreIndexArgs struct {
	// The API scope at which a query is run. Default value: "ANY_API" Possible values: ["ANY_API", "DATASTORE_MODE_API"]
	ApiScope pulumi.StringPtrInput
	// The collection being indexed.
	Collection pulumi.StringInput
	// The Firestore database id. Defaults to '"(default)"'.
	Database pulumi.StringPtrInput
	// The fields supported by this index. The last non-stored field entry is always for the field path '__name__'. If, on
	// creation, '__name__' was not specified as the last field, it will be added automatically with the same direction as that
	// of the last field defined. If the final field in a composite index is not directional, the '__name__' will be ordered
	// '"ASCENDING"' (unless explicitly specified otherwise).
	Fields           FirestoreIndexFieldArrayInput
	FirestoreIndexId pulumi.StringPtrInput
	Project          pulumi.StringPtrInput
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP",
	// "COLLECTION_RECURSIVE"]
	QueryScope pulumi.StringPtrInput
	Timeouts   FirestoreIndexTimeoutsPtrInput
}

func (FirestoreIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firestoreIndexArgs)(nil)).Elem()
}

type FirestoreIndexInput interface {
	pulumi.Input

	ToFirestoreIndexOutput() FirestoreIndexOutput
	ToFirestoreIndexOutputWithContext(ctx context.Context) FirestoreIndexOutput
}

func (*FirestoreIndex) ElementType() reflect.Type {
	return reflect.TypeOf((**FirestoreIndex)(nil)).Elem()
}

func (i *FirestoreIndex) ToFirestoreIndexOutput() FirestoreIndexOutput {
	return i.ToFirestoreIndexOutputWithContext(context.Background())
}

func (i *FirestoreIndex) ToFirestoreIndexOutputWithContext(ctx context.Context) FirestoreIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirestoreIndexOutput)
}

type FirestoreIndexOutput struct{ *pulumi.OutputState }

func (FirestoreIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirestoreIndex)(nil)).Elem()
}

func (o FirestoreIndexOutput) ToFirestoreIndexOutput() FirestoreIndexOutput {
	return o
}

func (o FirestoreIndexOutput) ToFirestoreIndexOutputWithContext(ctx context.Context) FirestoreIndexOutput {
	return o
}

// The API scope at which a query is run. Default value: "ANY_API" Possible values: ["ANY_API", "DATASTORE_MODE_API"]
func (o FirestoreIndexOutput) ApiScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirestoreIndex) pulumi.StringPtrOutput { return v.ApiScope }).(pulumi.StringPtrOutput)
}

// The collection being indexed.
func (o FirestoreIndexOutput) Collection() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreIndex) pulumi.StringOutput { return v.Collection }).(pulumi.StringOutput)
}

// The Firestore database id. Defaults to '"(default)"'.
func (o FirestoreIndexOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirestoreIndex) pulumi.StringPtrOutput { return v.Database }).(pulumi.StringPtrOutput)
}

// The fields supported by this index. The last non-stored field entry is always for the field path '__name__'. If, on
// creation, '__name__' was not specified as the last field, it will be added automatically with the same direction as that
// of the last field defined. If the final field in a composite index is not directional, the '__name__' will be ordered
// '"ASCENDING"' (unless explicitly specified otherwise).
func (o FirestoreIndexOutput) Fields() FirestoreIndexFieldArrayOutput {
	return o.ApplyT(func(v *FirestoreIndex) FirestoreIndexFieldArrayOutput { return v.Fields }).(FirestoreIndexFieldArrayOutput)
}

func (o FirestoreIndexOutput) FirestoreIndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreIndex) pulumi.StringOutput { return v.FirestoreIndexId }).(pulumi.StringOutput)
}

// A server defined name for this index. Format:
// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
func (o FirestoreIndexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreIndex) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirestoreIndexOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirestoreIndex) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP",
// "COLLECTION_RECURSIVE"]
func (o FirestoreIndexOutput) QueryScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirestoreIndex) pulumi.StringPtrOutput { return v.QueryScope }).(pulumi.StringPtrOutput)
}

func (o FirestoreIndexOutput) Timeouts() FirestoreIndexTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirestoreIndex) FirestoreIndexTimeoutsPtrOutput { return v.Timeouts }).(FirestoreIndexTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirestoreIndexInput)(nil)).Elem(), &FirestoreIndex{})
	pulumi.RegisterOutputType(FirestoreIndexOutput{})
}
