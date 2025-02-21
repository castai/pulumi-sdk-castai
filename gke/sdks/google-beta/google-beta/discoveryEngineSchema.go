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

type DiscoveryEngineSchema struct {
	pulumi.CustomResourceState

	// The unique id of the data store.
	DataStoreId             pulumi.StringOutput `pulumi:"dataStoreId"`
	DiscoveryEngineSchemaId pulumi.StringOutput `pulumi:"discoveryEngineSchemaId"`
	// The JSON representation of the schema.
	JsonSchema pulumi.StringPtrOutput `pulumi:"jsonSchema"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location pulumi.StringOutput `pulumi:"location"`
	// The unique full resource name of the schema. Values are of the format
	// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/schemas/{schema_id}'.
	// This field must be a UTF-8 encoded string with a length limit of 1024 characters.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The unique id of the schema.
	SchemaId pulumi.StringOutput                    `pulumi:"schemaId"`
	Timeouts DiscoveryEngineSchemaTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDiscoveryEngineSchema registers a new resource with the given unique name, arguments, and options.
func NewDiscoveryEngineSchema(ctx *pulumi.Context,
	name string, args *DiscoveryEngineSchemaArgs, opts ...pulumi.ResourceOption) (*DiscoveryEngineSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataStoreId == nil {
		return nil, errors.New("invalid value for required argument 'DataStoreId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.SchemaId == nil {
		return nil, errors.New("invalid value for required argument 'SchemaId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DiscoveryEngineSchema
	err = ctx.RegisterPackageResource("google-beta:index/discoveryEngineSchema:DiscoveryEngineSchema", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiscoveryEngineSchema gets an existing DiscoveryEngineSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiscoveryEngineSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiscoveryEngineSchemaState, opts ...pulumi.ResourceOption) (*DiscoveryEngineSchema, error) {
	var resource DiscoveryEngineSchema
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/discoveryEngineSchema:DiscoveryEngineSchema", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiscoveryEngineSchema resources.
type discoveryEngineSchemaState struct {
	// The unique id of the data store.
	DataStoreId             *string `pulumi:"dataStoreId"`
	DiscoveryEngineSchemaId *string `pulumi:"discoveryEngineSchemaId"`
	// The JSON representation of the schema.
	JsonSchema *string `pulumi:"jsonSchema"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location *string `pulumi:"location"`
	// The unique full resource name of the schema. Values are of the format
	// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/schemas/{schema_id}'.
	// This field must be a UTF-8 encoded string with a length limit of 1024 characters.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The unique id of the schema.
	SchemaId *string                        `pulumi:"schemaId"`
	Timeouts *DiscoveryEngineSchemaTimeouts `pulumi:"timeouts"`
}

type DiscoveryEngineSchemaState struct {
	// The unique id of the data store.
	DataStoreId             pulumi.StringPtrInput
	DiscoveryEngineSchemaId pulumi.StringPtrInput
	// The JSON representation of the schema.
	JsonSchema pulumi.StringPtrInput
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location pulumi.StringPtrInput
	// The unique full resource name of the schema. Values are of the format
	// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/schemas/{schema_id}'.
	// This field must be a UTF-8 encoded string with a length limit of 1024 characters.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The unique id of the schema.
	SchemaId pulumi.StringPtrInput
	Timeouts DiscoveryEngineSchemaTimeoutsPtrInput
}

func (DiscoveryEngineSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*discoveryEngineSchemaState)(nil)).Elem()
}

type discoveryEngineSchemaArgs struct {
	// The unique id of the data store.
	DataStoreId             string  `pulumi:"dataStoreId"`
	DiscoveryEngineSchemaId *string `pulumi:"discoveryEngineSchemaId"`
	// The JSON representation of the schema.
	JsonSchema *string `pulumi:"jsonSchema"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// The unique id of the schema.
	SchemaId string                         `pulumi:"schemaId"`
	Timeouts *DiscoveryEngineSchemaTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DiscoveryEngineSchema resource.
type DiscoveryEngineSchemaArgs struct {
	// The unique id of the data store.
	DataStoreId             pulumi.StringInput
	DiscoveryEngineSchemaId pulumi.StringPtrInput
	// The JSON representation of the schema.
	JsonSchema pulumi.StringPtrInput
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// The unique id of the schema.
	SchemaId pulumi.StringInput
	Timeouts DiscoveryEngineSchemaTimeoutsPtrInput
}

func (DiscoveryEngineSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*discoveryEngineSchemaArgs)(nil)).Elem()
}

type DiscoveryEngineSchemaInput interface {
	pulumi.Input

	ToDiscoveryEngineSchemaOutput() DiscoveryEngineSchemaOutput
	ToDiscoveryEngineSchemaOutputWithContext(ctx context.Context) DiscoveryEngineSchemaOutput
}

func (*DiscoveryEngineSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**DiscoveryEngineSchema)(nil)).Elem()
}

func (i *DiscoveryEngineSchema) ToDiscoveryEngineSchemaOutput() DiscoveryEngineSchemaOutput {
	return i.ToDiscoveryEngineSchemaOutputWithContext(context.Background())
}

func (i *DiscoveryEngineSchema) ToDiscoveryEngineSchemaOutputWithContext(ctx context.Context) DiscoveryEngineSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiscoveryEngineSchemaOutput)
}

type DiscoveryEngineSchemaOutput struct{ *pulumi.OutputState }

func (DiscoveryEngineSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiscoveryEngineSchema)(nil)).Elem()
}

func (o DiscoveryEngineSchemaOutput) ToDiscoveryEngineSchemaOutput() DiscoveryEngineSchemaOutput {
	return o
}

func (o DiscoveryEngineSchemaOutput) ToDiscoveryEngineSchemaOutputWithContext(ctx context.Context) DiscoveryEngineSchemaOutput {
	return o
}

// The unique id of the data store.
func (o DiscoveryEngineSchemaOutput) DataStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) pulumi.StringOutput { return v.DataStoreId }).(pulumi.StringOutput)
}

func (o DiscoveryEngineSchemaOutput) DiscoveryEngineSchemaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) pulumi.StringOutput { return v.DiscoveryEngineSchemaId }).(pulumi.StringOutput)
}

// The JSON representation of the schema.
func (o DiscoveryEngineSchemaOutput) JsonSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) pulumi.StringPtrOutput { return v.JsonSchema }).(pulumi.StringPtrOutput)
}

// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
func (o DiscoveryEngineSchemaOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The unique full resource name of the schema. Values are of the format
// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/schemas/{schema_id}'.
// This field must be a UTF-8 encoded string with a length limit of 1024 characters.
func (o DiscoveryEngineSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiscoveryEngineSchemaOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The unique id of the schema.
func (o DiscoveryEngineSchemaOutput) SchemaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) pulumi.StringOutput { return v.SchemaId }).(pulumi.StringOutput)
}

func (o DiscoveryEngineSchemaOutput) Timeouts() DiscoveryEngineSchemaTimeoutsPtrOutput {
	return o.ApplyT(func(v *DiscoveryEngineSchema) DiscoveryEngineSchemaTimeoutsPtrOutput { return v.Timeouts }).(DiscoveryEngineSchemaTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiscoveryEngineSchemaInput)(nil)).Elem(), &DiscoveryEngineSchema{})
	pulumi.RegisterOutputType(DiscoveryEngineSchemaOutput{})
}
