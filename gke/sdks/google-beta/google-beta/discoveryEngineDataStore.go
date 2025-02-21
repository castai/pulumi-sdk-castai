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

type DiscoveryEngineDataStore struct {
	pulumi.CustomResourceState

	// Configuration data for advance site search.
	AdvancedSiteSearchConfig DiscoveryEngineDataStoreAdvancedSiteSearchConfigPtrOutput `pulumi:"advancedSiteSearchConfig"`
	// The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"]
	ContentConfig pulumi.StringOutput `pulumi:"contentConfig"`
	// If true, an advanced data store for site search will be created. If the data store is not configured as site search
	// (GENERIC vertical and PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch pulumi.BoolPtrOutput `pulumi:"createAdvancedSiteSearch"`
	// Timestamp when the DataStore was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The unique id of the data store.
	DataStoreId pulumi.StringOutput `pulumi:"dataStoreId"`
	// The id of the default Schema associated with this data store.
	DefaultSchemaId            pulumi.StringOutput `pulumi:"defaultSchemaId"`
	DiscoveryEngineDataStoreId pulumi.StringOutput `pulumi:"discoveryEngineDataStoreId"`
	// The display name of the data store. This field must be a UTF-8 encoded string with a length limit of 128 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Configuration for Document understanding and enrichment.
	DocumentProcessingConfig DiscoveryEngineDataStoreDocumentProcessingConfigPtrOutput `pulumi:"documentProcessingConfig"`
	// The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]
	IndustryVertical pulumi.StringOutput `pulumi:"industryVertical"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location pulumi.StringOutput `pulumi:"location"`
	// The unique full resource name of the data store. Values are of the format
	// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}'. This field must be a
	// UTF-8 encoded string with a length limit of 1024 characters.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A boolean flag indicating whether to skip the default schema creation for the data store. Only enable this flag if you
	// are certain that the default schema is incompatible with your use case. If set to true, you must manually create a
	// schema for the data store before any documents can be ingested. This flag cannot be specified if
	// 'data_store.starting_schema' is specified.
	SkipDefaultSchemaCreation pulumi.BoolPtrOutput `pulumi:"skipDefaultSchemaCreation"`
	// The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH",
	// "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"]
	SolutionTypes pulumi.StringArrayOutput                  `pulumi:"solutionTypes"`
	Timeouts      DiscoveryEngineDataStoreTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDiscoveryEngineDataStore registers a new resource with the given unique name, arguments, and options.
func NewDiscoveryEngineDataStore(ctx *pulumi.Context,
	name string, args *DiscoveryEngineDataStoreArgs, opts ...pulumi.ResourceOption) (*DiscoveryEngineDataStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentConfig == nil {
		return nil, errors.New("invalid value for required argument 'ContentConfig'")
	}
	if args.DataStoreId == nil {
		return nil, errors.New("invalid value for required argument 'DataStoreId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IndustryVertical == nil {
		return nil, errors.New("invalid value for required argument 'IndustryVertical'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DiscoveryEngineDataStore
	err = ctx.RegisterPackageResource("google-beta:index/discoveryEngineDataStore:DiscoveryEngineDataStore", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiscoveryEngineDataStore gets an existing DiscoveryEngineDataStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiscoveryEngineDataStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiscoveryEngineDataStoreState, opts ...pulumi.ResourceOption) (*DiscoveryEngineDataStore, error) {
	var resource DiscoveryEngineDataStore
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/discoveryEngineDataStore:DiscoveryEngineDataStore", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiscoveryEngineDataStore resources.
type discoveryEngineDataStoreState struct {
	// Configuration data for advance site search.
	AdvancedSiteSearchConfig *DiscoveryEngineDataStoreAdvancedSiteSearchConfig `pulumi:"advancedSiteSearchConfig"`
	// The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"]
	ContentConfig *string `pulumi:"contentConfig"`
	// If true, an advanced data store for site search will be created. If the data store is not configured as site search
	// (GENERIC vertical and PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch *bool `pulumi:"createAdvancedSiteSearch"`
	// Timestamp when the DataStore was created.
	CreateTime *string `pulumi:"createTime"`
	// The unique id of the data store.
	DataStoreId *string `pulumi:"dataStoreId"`
	// The id of the default Schema associated with this data store.
	DefaultSchemaId            *string `pulumi:"defaultSchemaId"`
	DiscoveryEngineDataStoreId *string `pulumi:"discoveryEngineDataStoreId"`
	// The display name of the data store. This field must be a UTF-8 encoded string with a length limit of 128 characters.
	DisplayName *string `pulumi:"displayName"`
	// Configuration for Document understanding and enrichment.
	DocumentProcessingConfig *DiscoveryEngineDataStoreDocumentProcessingConfig `pulumi:"documentProcessingConfig"`
	// The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]
	IndustryVertical *string `pulumi:"industryVertical"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location *string `pulumi:"location"`
	// The unique full resource name of the data store. Values are of the format
	// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}'. This field must be a
	// UTF-8 encoded string with a length limit of 1024 characters.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A boolean flag indicating whether to skip the default schema creation for the data store. Only enable this flag if you
	// are certain that the default schema is incompatible with your use case. If set to true, you must manually create a
	// schema for the data store before any documents can be ingested. This flag cannot be specified if
	// 'data_store.starting_schema' is specified.
	SkipDefaultSchemaCreation *bool `pulumi:"skipDefaultSchemaCreation"`
	// The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH",
	// "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"]
	SolutionTypes []string                          `pulumi:"solutionTypes"`
	Timeouts      *DiscoveryEngineDataStoreTimeouts `pulumi:"timeouts"`
}

type DiscoveryEngineDataStoreState struct {
	// Configuration data for advance site search.
	AdvancedSiteSearchConfig DiscoveryEngineDataStoreAdvancedSiteSearchConfigPtrInput
	// The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"]
	ContentConfig pulumi.StringPtrInput
	// If true, an advanced data store for site search will be created. If the data store is not configured as site search
	// (GENERIC vertical and PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch pulumi.BoolPtrInput
	// Timestamp when the DataStore was created.
	CreateTime pulumi.StringPtrInput
	// The unique id of the data store.
	DataStoreId pulumi.StringPtrInput
	// The id of the default Schema associated with this data store.
	DefaultSchemaId            pulumi.StringPtrInput
	DiscoveryEngineDataStoreId pulumi.StringPtrInput
	// The display name of the data store. This field must be a UTF-8 encoded string with a length limit of 128 characters.
	DisplayName pulumi.StringPtrInput
	// Configuration for Document understanding and enrichment.
	DocumentProcessingConfig DiscoveryEngineDataStoreDocumentProcessingConfigPtrInput
	// The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]
	IndustryVertical pulumi.StringPtrInput
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location pulumi.StringPtrInput
	// The unique full resource name of the data store. Values are of the format
	// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}'. This field must be a
	// UTF-8 encoded string with a length limit of 1024 characters.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A boolean flag indicating whether to skip the default schema creation for the data store. Only enable this flag if you
	// are certain that the default schema is incompatible with your use case. If set to true, you must manually create a
	// schema for the data store before any documents can be ingested. This flag cannot be specified if
	// 'data_store.starting_schema' is specified.
	SkipDefaultSchemaCreation pulumi.BoolPtrInput
	// The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH",
	// "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"]
	SolutionTypes pulumi.StringArrayInput
	Timeouts      DiscoveryEngineDataStoreTimeoutsPtrInput
}

func (DiscoveryEngineDataStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*discoveryEngineDataStoreState)(nil)).Elem()
}

type discoveryEngineDataStoreArgs struct {
	// Configuration data for advance site search.
	AdvancedSiteSearchConfig *DiscoveryEngineDataStoreAdvancedSiteSearchConfig `pulumi:"advancedSiteSearchConfig"`
	// The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"]
	ContentConfig string `pulumi:"contentConfig"`
	// If true, an advanced data store for site search will be created. If the data store is not configured as site search
	// (GENERIC vertical and PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch *bool `pulumi:"createAdvancedSiteSearch"`
	// The unique id of the data store.
	DataStoreId                string  `pulumi:"dataStoreId"`
	DiscoveryEngineDataStoreId *string `pulumi:"discoveryEngineDataStoreId"`
	// The display name of the data store. This field must be a UTF-8 encoded string with a length limit of 128 characters.
	DisplayName string `pulumi:"displayName"`
	// Configuration for Document understanding and enrichment.
	DocumentProcessingConfig *DiscoveryEngineDataStoreDocumentProcessingConfig `pulumi:"documentProcessingConfig"`
	// The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]
	IndustryVertical string `pulumi:"industryVertical"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// A boolean flag indicating whether to skip the default schema creation for the data store. Only enable this flag if you
	// are certain that the default schema is incompatible with your use case. If set to true, you must manually create a
	// schema for the data store before any documents can be ingested. This flag cannot be specified if
	// 'data_store.starting_schema' is specified.
	SkipDefaultSchemaCreation *bool `pulumi:"skipDefaultSchemaCreation"`
	// The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH",
	// "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"]
	SolutionTypes []string                          `pulumi:"solutionTypes"`
	Timeouts      *DiscoveryEngineDataStoreTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DiscoveryEngineDataStore resource.
type DiscoveryEngineDataStoreArgs struct {
	// Configuration data for advance site search.
	AdvancedSiteSearchConfig DiscoveryEngineDataStoreAdvancedSiteSearchConfigPtrInput
	// The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"]
	ContentConfig pulumi.StringInput
	// If true, an advanced data store for site search will be created. If the data store is not configured as site search
	// (GENERIC vertical and PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch pulumi.BoolPtrInput
	// The unique id of the data store.
	DataStoreId                pulumi.StringInput
	DiscoveryEngineDataStoreId pulumi.StringPtrInput
	// The display name of the data store. This field must be a UTF-8 encoded string with a length limit of 128 characters.
	DisplayName pulumi.StringInput
	// Configuration for Document understanding and enrichment.
	DocumentProcessingConfig DiscoveryEngineDataStoreDocumentProcessingConfigPtrInput
	// The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]
	IndustryVertical pulumi.StringInput
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// A boolean flag indicating whether to skip the default schema creation for the data store. Only enable this flag if you
	// are certain that the default schema is incompatible with your use case. If set to true, you must manually create a
	// schema for the data store before any documents can be ingested. This flag cannot be specified if
	// 'data_store.starting_schema' is specified.
	SkipDefaultSchemaCreation pulumi.BoolPtrInput
	// The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH",
	// "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"]
	SolutionTypes pulumi.StringArrayInput
	Timeouts      DiscoveryEngineDataStoreTimeoutsPtrInput
}

func (DiscoveryEngineDataStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*discoveryEngineDataStoreArgs)(nil)).Elem()
}

type DiscoveryEngineDataStoreInput interface {
	pulumi.Input

	ToDiscoveryEngineDataStoreOutput() DiscoveryEngineDataStoreOutput
	ToDiscoveryEngineDataStoreOutputWithContext(ctx context.Context) DiscoveryEngineDataStoreOutput
}

func (*DiscoveryEngineDataStore) ElementType() reflect.Type {
	return reflect.TypeOf((**DiscoveryEngineDataStore)(nil)).Elem()
}

func (i *DiscoveryEngineDataStore) ToDiscoveryEngineDataStoreOutput() DiscoveryEngineDataStoreOutput {
	return i.ToDiscoveryEngineDataStoreOutputWithContext(context.Background())
}

func (i *DiscoveryEngineDataStore) ToDiscoveryEngineDataStoreOutputWithContext(ctx context.Context) DiscoveryEngineDataStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiscoveryEngineDataStoreOutput)
}

type DiscoveryEngineDataStoreOutput struct{ *pulumi.OutputState }

func (DiscoveryEngineDataStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiscoveryEngineDataStore)(nil)).Elem()
}

func (o DiscoveryEngineDataStoreOutput) ToDiscoveryEngineDataStoreOutput() DiscoveryEngineDataStoreOutput {
	return o
}

func (o DiscoveryEngineDataStoreOutput) ToDiscoveryEngineDataStoreOutputWithContext(ctx context.Context) DiscoveryEngineDataStoreOutput {
	return o
}

// Configuration data for advance site search.
func (o DiscoveryEngineDataStoreOutput) AdvancedSiteSearchConfig() DiscoveryEngineDataStoreAdvancedSiteSearchConfigPtrOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) DiscoveryEngineDataStoreAdvancedSiteSearchConfigPtrOutput {
		return v.AdvancedSiteSearchConfig
	}).(DiscoveryEngineDataStoreAdvancedSiteSearchConfigPtrOutput)
}

// The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"]
func (o DiscoveryEngineDataStoreOutput) ContentConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.ContentConfig }).(pulumi.StringOutput)
}

// If true, an advanced data store for site search will be created. If the data store is not configured as site search
// (GENERIC vertical and PUBLIC_WEBSITE contentConfig), this flag will be ignored.
func (o DiscoveryEngineDataStoreOutput) CreateAdvancedSiteSearch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.BoolPtrOutput { return v.CreateAdvancedSiteSearch }).(pulumi.BoolPtrOutput)
}

// Timestamp when the DataStore was created.
func (o DiscoveryEngineDataStoreOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The unique id of the data store.
func (o DiscoveryEngineDataStoreOutput) DataStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.DataStoreId }).(pulumi.StringOutput)
}

// The id of the default Schema associated with this data store.
func (o DiscoveryEngineDataStoreOutput) DefaultSchemaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.DefaultSchemaId }).(pulumi.StringOutput)
}

func (o DiscoveryEngineDataStoreOutput) DiscoveryEngineDataStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.DiscoveryEngineDataStoreId }).(pulumi.StringOutput)
}

// The display name of the data store. This field must be a UTF-8 encoded string with a length limit of 128 characters.
func (o DiscoveryEngineDataStoreOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Configuration for Document understanding and enrichment.
func (o DiscoveryEngineDataStoreOutput) DocumentProcessingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigPtrOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) DiscoveryEngineDataStoreDocumentProcessingConfigPtrOutput {
		return v.DocumentProcessingConfig
	}).(DiscoveryEngineDataStoreDocumentProcessingConfigPtrOutput)
}

// The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"]
func (o DiscoveryEngineDataStoreOutput) IndustryVertical() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.IndustryVertical }).(pulumi.StringOutput)
}

// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
func (o DiscoveryEngineDataStoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The unique full resource name of the data store. Values are of the format
// 'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}'. This field must be a
// UTF-8 encoded string with a length limit of 1024 characters.
func (o DiscoveryEngineDataStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiscoveryEngineDataStoreOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A boolean flag indicating whether to skip the default schema creation for the data store. Only enable this flag if you
// are certain that the default schema is incompatible with your use case. If set to true, you must manually create a
// schema for the data store before any documents can be ingested. This flag cannot be specified if
// 'data_store.starting_schema' is specified.
func (o DiscoveryEngineDataStoreOutput) SkipDefaultSchemaCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.BoolPtrOutput { return v.SkipDefaultSchemaCreation }).(pulumi.BoolPtrOutput)
}

// The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH",
// "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"]
func (o DiscoveryEngineDataStoreOutput) SolutionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) pulumi.StringArrayOutput { return v.SolutionTypes }).(pulumi.StringArrayOutput)
}

func (o DiscoveryEngineDataStoreOutput) Timeouts() DiscoveryEngineDataStoreTimeoutsPtrOutput {
	return o.ApplyT(func(v *DiscoveryEngineDataStore) DiscoveryEngineDataStoreTimeoutsPtrOutput { return v.Timeouts }).(DiscoveryEngineDataStoreTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiscoveryEngineDataStoreInput)(nil)).Elem(), &DiscoveryEngineDataStore{})
	pulumi.RegisterOutputType(DiscoveryEngineDataStoreOutput{})
}
