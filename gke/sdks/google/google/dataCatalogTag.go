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

type DataCatalogTag struct {
	pulumi.CustomResourceState

	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an individual
	// column based on that schema. For attaching a tag to a nested column, use '.' to separate the column names. Example:
	// 'outer_column.inner_column'
	Column           pulumi.StringPtrOutput `pulumi:"column"`
	DataCatalogTagId pulumi.StringOutput    `pulumi:"dataCatalogTagId"`
	// This maps the ID of a tag field to the value of and additional information about that field. Valid field IDs are defined
	// by the tag's template. A tag must have at least 1 field and at most 500 fields.
	Fields DataCatalogTagFieldArrayOutput `pulumi:"fields"`
	// The resource name of the tag in URL format. Example:
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated
	// identifier. Note that this Tag may not actually be stored in the location in this name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group,
	// the tag will be attached to all entries in that group.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId} This field cannot be modified after creation.
	Template pulumi.StringOutput `pulumi:"template"`
	// The display name of the tag template.
	TemplateDisplayname pulumi.StringOutput             `pulumi:"templateDisplayname"`
	Timeouts            DataCatalogTagTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDataCatalogTag registers a new resource with the given unique name, arguments, and options.
func NewDataCatalogTag(ctx *pulumi.Context,
	name string, args *DataCatalogTagArgs, opts ...pulumi.ResourceOption) (*DataCatalogTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataCatalogTag
	err = ctx.RegisterPackageResource("google:index/dataCatalogTag:DataCatalogTag", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCatalogTag gets an existing DataCatalogTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCatalogTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCatalogTagState, opts ...pulumi.ResourceOption) (*DataCatalogTag, error) {
	var resource DataCatalogTag
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataCatalogTag:DataCatalogTag", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCatalogTag resources.
type dataCatalogTagState struct {
	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an individual
	// column based on that schema. For attaching a tag to a nested column, use '.' to separate the column names. Example:
	// 'outer_column.inner_column'
	Column           *string `pulumi:"column"`
	DataCatalogTagId *string `pulumi:"dataCatalogTagId"`
	// This maps the ID of a tag field to the value of and additional information about that field. Valid field IDs are defined
	// by the tag's template. A tag must have at least 1 field and at most 500 fields.
	Fields []DataCatalogTagField `pulumi:"fields"`
	// The resource name of the tag in URL format. Example:
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated
	// identifier. Note that this Tag may not actually be stored in the location in this name.
	Name *string `pulumi:"name"`
	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group,
	// the tag will be attached to all entries in that group.
	Parent *string `pulumi:"parent"`
	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId} This field cannot be modified after creation.
	Template *string `pulumi:"template"`
	// The display name of the tag template.
	TemplateDisplayname *string                 `pulumi:"templateDisplayname"`
	Timeouts            *DataCatalogTagTimeouts `pulumi:"timeouts"`
}

type DataCatalogTagState struct {
	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an individual
	// column based on that schema. For attaching a tag to a nested column, use '.' to separate the column names. Example:
	// 'outer_column.inner_column'
	Column           pulumi.StringPtrInput
	DataCatalogTagId pulumi.StringPtrInput
	// This maps the ID of a tag field to the value of and additional information about that field. Valid field IDs are defined
	// by the tag's template. A tag must have at least 1 field and at most 500 fields.
	Fields DataCatalogTagFieldArrayInput
	// The resource name of the tag in URL format. Example:
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated
	// identifier. Note that this Tag may not actually be stored in the location in this name.
	Name pulumi.StringPtrInput
	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group,
	// the tag will be attached to all entries in that group.
	Parent pulumi.StringPtrInput
	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId} This field cannot be modified after creation.
	Template pulumi.StringPtrInput
	// The display name of the tag template.
	TemplateDisplayname pulumi.StringPtrInput
	Timeouts            DataCatalogTagTimeoutsPtrInput
}

func (DataCatalogTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogTagState)(nil)).Elem()
}

type dataCatalogTagArgs struct {
	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an individual
	// column based on that schema. For attaching a tag to a nested column, use '.' to separate the column names. Example:
	// 'outer_column.inner_column'
	Column           *string `pulumi:"column"`
	DataCatalogTagId *string `pulumi:"dataCatalogTagId"`
	// This maps the ID of a tag field to the value of and additional information about that field. Valid field IDs are defined
	// by the tag's template. A tag must have at least 1 field and at most 500 fields.
	Fields []DataCatalogTagField `pulumi:"fields"`
	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group,
	// the tag will be attached to all entries in that group.
	Parent *string `pulumi:"parent"`
	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId} This field cannot be modified after creation.
	Template string                  `pulumi:"template"`
	Timeouts *DataCatalogTagTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DataCatalogTag resource.
type DataCatalogTagArgs struct {
	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an individual
	// column based on that schema. For attaching a tag to a nested column, use '.' to separate the column names. Example:
	// 'outer_column.inner_column'
	Column           pulumi.StringPtrInput
	DataCatalogTagId pulumi.StringPtrInput
	// This maps the ID of a tag field to the value of and additional information about that field. Valid field IDs are defined
	// by the tag's template. A tag must have at least 1 field and at most 500 fields.
	Fields DataCatalogTagFieldArrayInput
	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group,
	// the tag will be attached to all entries in that group.
	Parent pulumi.StringPtrInput
	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId} This field cannot be modified after creation.
	Template pulumi.StringInput
	Timeouts DataCatalogTagTimeoutsPtrInput
}

func (DataCatalogTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogTagArgs)(nil)).Elem()
}

type DataCatalogTagInput interface {
	pulumi.Input

	ToDataCatalogTagOutput() DataCatalogTagOutput
	ToDataCatalogTagOutputWithContext(ctx context.Context) DataCatalogTagOutput
}

func (*DataCatalogTag) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogTag)(nil)).Elem()
}

func (i *DataCatalogTag) ToDataCatalogTagOutput() DataCatalogTagOutput {
	return i.ToDataCatalogTagOutputWithContext(context.Background())
}

func (i *DataCatalogTag) ToDataCatalogTagOutputWithContext(ctx context.Context) DataCatalogTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCatalogTagOutput)
}

type DataCatalogTagOutput struct{ *pulumi.OutputState }

func (DataCatalogTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogTag)(nil)).Elem()
}

func (o DataCatalogTagOutput) ToDataCatalogTagOutput() DataCatalogTagOutput {
	return o
}

func (o DataCatalogTagOutput) ToDataCatalogTagOutputWithContext(ctx context.Context) DataCatalogTagOutput {
	return o
}

// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an individual
// column based on that schema. For attaching a tag to a nested column, use '.' to separate the column names. Example:
// 'outer_column.inner_column'
func (o DataCatalogTagOutput) Column() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCatalogTag) pulumi.StringPtrOutput { return v.Column }).(pulumi.StringPtrOutput)
}

func (o DataCatalogTagOutput) DataCatalogTagId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTag) pulumi.StringOutput { return v.DataCatalogTagId }).(pulumi.StringOutput)
}

// This maps the ID of a tag field to the value of and additional information about that field. Valid field IDs are defined
// by the tag's template. A tag must have at least 1 field and at most 500 fields.
func (o DataCatalogTagOutput) Fields() DataCatalogTagFieldArrayOutput {
	return o.ApplyT(func(v *DataCatalogTag) DataCatalogTagFieldArrayOutput { return v.Fields }).(DataCatalogTagFieldArrayOutput)
}

// The resource name of the tag in URL format. Example:
// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated
// identifier. Note that this Tag may not actually be stored in the location in this name.
func (o DataCatalogTagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTag) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group,
// the tag will be attached to all entries in that group.
func (o DataCatalogTagOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCatalogTag) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// The resource name of the tag template that this tag uses. Example:
// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId} This field cannot be modified after creation.
func (o DataCatalogTagOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTag) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

// The display name of the tag template.
func (o DataCatalogTagOutput) TemplateDisplayname() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTag) pulumi.StringOutput { return v.TemplateDisplayname }).(pulumi.StringOutput)
}

func (o DataCatalogTagOutput) Timeouts() DataCatalogTagTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataCatalogTag) DataCatalogTagTimeoutsPtrOutput { return v.Timeouts }).(DataCatalogTagTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCatalogTagInput)(nil)).Elem(), &DataCatalogTag{})
	pulumi.RegisterOutputType(DataCatalogTagOutput{})
}
