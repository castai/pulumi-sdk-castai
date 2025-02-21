// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataplexAspectType struct {
	pulumi.CustomResourceState

	// The aspect type id of the aspect type.
	AspectTypeId pulumi.StringPtrOutput `pulumi:"aspectTypeId"`
	// The time when the AspectType was created.
	CreateTime           pulumi.StringOutput `pulumi:"createTime"`
	DataplexAspectTypeId pulumi.StringOutput `pulumi:"dataplexAspectTypeId"`
	// Description of the AspectType.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User friendly display name.
	DisplayName     pulumi.StringPtrOutput `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// User-defined labels for the AspectType. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where aspect type will be created in.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// MetadataTemplate of the Aspect.
	MetadataTemplate pulumi.StringPtrOutput `pulumi:"metadataTemplate"`
	// The relative resource name of the AspectType, of the form:
	// projects/{project_number}/locations/{location_id}/aspectTypes/{aspect_type_id}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput              `pulumi:"terraformLabels"`
	Timeouts        DataplexAspectTypeTimeoutsPtrOutput `pulumi:"timeouts"`
	// Denotes the transfer status of the Aspect Type. It is unspecified for Aspect Type created from Dataplex API.
	TransferStatus pulumi.StringOutput `pulumi:"transferStatus"`
	// System generated globally unique ID for the AspectType. This ID will be different if the AspectType is deleted and
	// re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the AspectType was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDataplexAspectType registers a new resource with the given unique name, arguments, and options.
func NewDataplexAspectType(ctx *pulumi.Context,
	name string, args *DataplexAspectTypeArgs, opts ...pulumi.ResourceOption) (*DataplexAspectType, error) {
	if args == nil {
		args = &DataplexAspectTypeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexAspectType
	err = ctx.RegisterPackageResource("google:index/dataplexAspectType:DataplexAspectType", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexAspectType gets an existing DataplexAspectType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexAspectType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexAspectTypeState, opts ...pulumi.ResourceOption) (*DataplexAspectType, error) {
	var resource DataplexAspectType
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataplexAspectType:DataplexAspectType", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexAspectType resources.
type dataplexAspectTypeState struct {
	// The aspect type id of the aspect type.
	AspectTypeId *string `pulumi:"aspectTypeId"`
	// The time when the AspectType was created.
	CreateTime           *string `pulumi:"createTime"`
	DataplexAspectTypeId *string `pulumi:"dataplexAspectTypeId"`
	// Description of the AspectType.
	Description *string `pulumi:"description"`
	// User friendly display name.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// User-defined labels for the AspectType. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where aspect type will be created in.
	Location *string `pulumi:"location"`
	// MetadataTemplate of the Aspect.
	MetadataTemplate *string `pulumi:"metadataTemplate"`
	// The relative resource name of the AspectType, of the form:
	// projects/{project_number}/locations/{location_id}/aspectTypes/{aspect_type_id}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string           `pulumi:"terraformLabels"`
	Timeouts        *DataplexAspectTypeTimeouts `pulumi:"timeouts"`
	// Denotes the transfer status of the Aspect Type. It is unspecified for Aspect Type created from Dataplex API.
	TransferStatus *string `pulumi:"transferStatus"`
	// System generated globally unique ID for the AspectType. This ID will be different if the AspectType is deleted and
	// re-created with the same name.
	Uid *string `pulumi:"uid"`
	// The time when the AspectType was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type DataplexAspectTypeState struct {
	// The aspect type id of the aspect type.
	AspectTypeId pulumi.StringPtrInput
	// The time when the AspectType was created.
	CreateTime           pulumi.StringPtrInput
	DataplexAspectTypeId pulumi.StringPtrInput
	// Description of the AspectType.
	Description pulumi.StringPtrInput
	// User friendly display name.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// User-defined labels for the AspectType. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where aspect type will be created in.
	Location pulumi.StringPtrInput
	// MetadataTemplate of the Aspect.
	MetadataTemplate pulumi.StringPtrInput
	// The relative resource name of the AspectType, of the form:
	// projects/{project_number}/locations/{location_id}/aspectTypes/{aspect_type_id}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DataplexAspectTypeTimeoutsPtrInput
	// Denotes the transfer status of the Aspect Type. It is unspecified for Aspect Type created from Dataplex API.
	TransferStatus pulumi.StringPtrInput
	// System generated globally unique ID for the AspectType. This ID will be different if the AspectType is deleted and
	// re-created with the same name.
	Uid pulumi.StringPtrInput
	// The time when the AspectType was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (DataplexAspectTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexAspectTypeState)(nil)).Elem()
}

type dataplexAspectTypeArgs struct {
	// The aspect type id of the aspect type.
	AspectTypeId         *string `pulumi:"aspectTypeId"`
	DataplexAspectTypeId *string `pulumi:"dataplexAspectTypeId"`
	// Description of the AspectType.
	Description *string `pulumi:"description"`
	// User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// User-defined labels for the AspectType. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where aspect type will be created in.
	Location *string `pulumi:"location"`
	// MetadataTemplate of the Aspect.
	MetadataTemplate *string                     `pulumi:"metadataTemplate"`
	Project          *string                     `pulumi:"project"`
	Timeouts         *DataplexAspectTypeTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DataplexAspectType resource.
type DataplexAspectTypeArgs struct {
	// The aspect type id of the aspect type.
	AspectTypeId         pulumi.StringPtrInput
	DataplexAspectTypeId pulumi.StringPtrInput
	// Description of the AspectType.
	Description pulumi.StringPtrInput
	// User friendly display name.
	DisplayName pulumi.StringPtrInput
	// User-defined labels for the AspectType. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where aspect type will be created in.
	Location pulumi.StringPtrInput
	// MetadataTemplate of the Aspect.
	MetadataTemplate pulumi.StringPtrInput
	Project          pulumi.StringPtrInput
	Timeouts         DataplexAspectTypeTimeoutsPtrInput
}

func (DataplexAspectTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexAspectTypeArgs)(nil)).Elem()
}

type DataplexAspectTypeInput interface {
	pulumi.Input

	ToDataplexAspectTypeOutput() DataplexAspectTypeOutput
	ToDataplexAspectTypeOutputWithContext(ctx context.Context) DataplexAspectTypeOutput
}

func (*DataplexAspectType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexAspectType)(nil)).Elem()
}

func (i *DataplexAspectType) ToDataplexAspectTypeOutput() DataplexAspectTypeOutput {
	return i.ToDataplexAspectTypeOutputWithContext(context.Background())
}

func (i *DataplexAspectType) ToDataplexAspectTypeOutputWithContext(ctx context.Context) DataplexAspectTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexAspectTypeOutput)
}

type DataplexAspectTypeOutput struct{ *pulumi.OutputState }

func (DataplexAspectTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexAspectType)(nil)).Elem()
}

func (o DataplexAspectTypeOutput) ToDataplexAspectTypeOutput() DataplexAspectTypeOutput {
	return o
}

func (o DataplexAspectTypeOutput) ToDataplexAspectTypeOutputWithContext(ctx context.Context) DataplexAspectTypeOutput {
	return o
}

// The aspect type id of the aspect type.
func (o DataplexAspectTypeOutput) AspectTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringPtrOutput { return v.AspectTypeId }).(pulumi.StringPtrOutput)
}

// The time when the AspectType was created.
func (o DataplexAspectTypeOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o DataplexAspectTypeOutput) DataplexAspectTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringOutput { return v.DataplexAspectTypeId }).(pulumi.StringOutput)
}

// Description of the AspectType.
func (o DataplexAspectTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User friendly display name.
func (o DataplexAspectTypeOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o DataplexAspectTypeOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// User-defined labels for the AspectType. **Note**: This field is non-authoritative, and will only manage the labels
// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o DataplexAspectTypeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where aspect type will be created in.
func (o DataplexAspectTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// MetadataTemplate of the Aspect.
func (o DataplexAspectTypeOutput) MetadataTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringPtrOutput { return v.MetadataTemplate }).(pulumi.StringPtrOutput)
}

// The relative resource name of the AspectType, of the form:
// projects/{project_number}/locations/{location_id}/aspectTypes/{aspect_type_id}
func (o DataplexAspectTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataplexAspectTypeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataplexAspectTypeOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DataplexAspectTypeOutput) Timeouts() DataplexAspectTypeTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataplexAspectType) DataplexAspectTypeTimeoutsPtrOutput { return v.Timeouts }).(DataplexAspectTypeTimeoutsPtrOutput)
}

// Denotes the transfer status of the Aspect Type. It is unspecified for Aspect Type created from Dataplex API.
func (o DataplexAspectTypeOutput) TransferStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringOutput { return v.TransferStatus }).(pulumi.StringOutput)
}

// System generated globally unique ID for the AspectType. This ID will be different if the AspectType is deleted and
// re-created with the same name.
func (o DataplexAspectTypeOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the AspectType was last updated.
func (o DataplexAspectTypeOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectType) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexAspectTypeInput)(nil)).Elem(), &DataplexAspectType{})
	pulumi.RegisterOutputType(DataplexAspectTypeOutput{})
}
