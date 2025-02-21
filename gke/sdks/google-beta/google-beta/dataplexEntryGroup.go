// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataplexEntryGroup struct {
	pulumi.CustomResourceState

	// The time when the EntryGroup was created.
	CreateTime           pulumi.StringOutput `pulumi:"createTime"`
	DataplexEntryGroupId pulumi.StringOutput `pulumi:"dataplexEntryGroupId"`
	// Description of the EntryGroup.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User friendly display name.
	DisplayName     pulumi.StringPtrOutput `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The entry group id of the entry group.
	EntryGroupId pulumi.StringPtrOutput `pulumi:"entryGroupId"`
	// User-defined labels for the EntryGroup. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where entry group will be created in.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The relative resource name of the EntryGroup, of the form:
	// projects/{project_number}/locations/{location_id}/entryGroups/{entry_group_id}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput              `pulumi:"terraformLabels"`
	Timeouts        DataplexEntryGroupTimeoutsPtrOutput `pulumi:"timeouts"`
	// Denotes the transfer status of the Entry Group. It is unspecified for Entry Group created from Dataplex API.
	TransferStatus pulumi.StringOutput `pulumi:"transferStatus"`
	// System generated globally unique ID for the EntryGroup. This ID will be different if the EntryGroup is deleted and
	// re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the EntryGroup was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDataplexEntryGroup registers a new resource with the given unique name, arguments, and options.
func NewDataplexEntryGroup(ctx *pulumi.Context,
	name string, args *DataplexEntryGroupArgs, opts ...pulumi.ResourceOption) (*DataplexEntryGroup, error) {
	if args == nil {
		args = &DataplexEntryGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexEntryGroup
	err = ctx.RegisterPackageResource("google-beta:index/dataplexEntryGroup:DataplexEntryGroup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexEntryGroup gets an existing DataplexEntryGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexEntryGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexEntryGroupState, opts ...pulumi.ResourceOption) (*DataplexEntryGroup, error) {
	var resource DataplexEntryGroup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataplexEntryGroup:DataplexEntryGroup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexEntryGroup resources.
type dataplexEntryGroupState struct {
	// The time when the EntryGroup was created.
	CreateTime           *string `pulumi:"createTime"`
	DataplexEntryGroupId *string `pulumi:"dataplexEntryGroupId"`
	// Description of the EntryGroup.
	Description *string `pulumi:"description"`
	// User friendly display name.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The entry group id of the entry group.
	EntryGroupId *string `pulumi:"entryGroupId"`
	// User-defined labels for the EntryGroup. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where entry group will be created in.
	Location *string `pulumi:"location"`
	// The relative resource name of the EntryGroup, of the form:
	// projects/{project_number}/locations/{location_id}/entryGroups/{entry_group_id}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string           `pulumi:"terraformLabels"`
	Timeouts        *DataplexEntryGroupTimeouts `pulumi:"timeouts"`
	// Denotes the transfer status of the Entry Group. It is unspecified for Entry Group created from Dataplex API.
	TransferStatus *string `pulumi:"transferStatus"`
	// System generated globally unique ID for the EntryGroup. This ID will be different if the EntryGroup is deleted and
	// re-created with the same name.
	Uid *string `pulumi:"uid"`
	// The time when the EntryGroup was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type DataplexEntryGroupState struct {
	// The time when the EntryGroup was created.
	CreateTime           pulumi.StringPtrInput
	DataplexEntryGroupId pulumi.StringPtrInput
	// Description of the EntryGroup.
	Description pulumi.StringPtrInput
	// User friendly display name.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The entry group id of the entry group.
	EntryGroupId pulumi.StringPtrInput
	// User-defined labels for the EntryGroup. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where entry group will be created in.
	Location pulumi.StringPtrInput
	// The relative resource name of the EntryGroup, of the form:
	// projects/{project_number}/locations/{location_id}/entryGroups/{entry_group_id}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DataplexEntryGroupTimeoutsPtrInput
	// Denotes the transfer status of the Entry Group. It is unspecified for Entry Group created from Dataplex API.
	TransferStatus pulumi.StringPtrInput
	// System generated globally unique ID for the EntryGroup. This ID will be different if the EntryGroup is deleted and
	// re-created with the same name.
	Uid pulumi.StringPtrInput
	// The time when the EntryGroup was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (DataplexEntryGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexEntryGroupState)(nil)).Elem()
}

type dataplexEntryGroupArgs struct {
	DataplexEntryGroupId *string `pulumi:"dataplexEntryGroupId"`
	// Description of the EntryGroup.
	Description *string `pulumi:"description"`
	// User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// The entry group id of the entry group.
	EntryGroupId *string `pulumi:"entryGroupId"`
	// User-defined labels for the EntryGroup. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where entry group will be created in.
	Location *string                     `pulumi:"location"`
	Project  *string                     `pulumi:"project"`
	Timeouts *DataplexEntryGroupTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DataplexEntryGroup resource.
type DataplexEntryGroupArgs struct {
	DataplexEntryGroupId pulumi.StringPtrInput
	// Description of the EntryGroup.
	Description pulumi.StringPtrInput
	// User friendly display name.
	DisplayName pulumi.StringPtrInput
	// The entry group id of the entry group.
	EntryGroupId pulumi.StringPtrInput
	// User-defined labels for the EntryGroup. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where entry group will be created in.
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts DataplexEntryGroupTimeoutsPtrInput
}

func (DataplexEntryGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexEntryGroupArgs)(nil)).Elem()
}

type DataplexEntryGroupInput interface {
	pulumi.Input

	ToDataplexEntryGroupOutput() DataplexEntryGroupOutput
	ToDataplexEntryGroupOutputWithContext(ctx context.Context) DataplexEntryGroupOutput
}

func (*DataplexEntryGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexEntryGroup)(nil)).Elem()
}

func (i *DataplexEntryGroup) ToDataplexEntryGroupOutput() DataplexEntryGroupOutput {
	return i.ToDataplexEntryGroupOutputWithContext(context.Background())
}

func (i *DataplexEntryGroup) ToDataplexEntryGroupOutputWithContext(ctx context.Context) DataplexEntryGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexEntryGroupOutput)
}

type DataplexEntryGroupOutput struct{ *pulumi.OutputState }

func (DataplexEntryGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexEntryGroup)(nil)).Elem()
}

func (o DataplexEntryGroupOutput) ToDataplexEntryGroupOutput() DataplexEntryGroupOutput {
	return o
}

func (o DataplexEntryGroupOutput) ToDataplexEntryGroupOutputWithContext(ctx context.Context) DataplexEntryGroupOutput {
	return o
}

// The time when the EntryGroup was created.
func (o DataplexEntryGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o DataplexEntryGroupOutput) DataplexEntryGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringOutput { return v.DataplexEntryGroupId }).(pulumi.StringOutput)
}

// Description of the EntryGroup.
func (o DataplexEntryGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User friendly display name.
func (o DataplexEntryGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o DataplexEntryGroupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The entry group id of the entry group.
func (o DataplexEntryGroupOutput) EntryGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringPtrOutput { return v.EntryGroupId }).(pulumi.StringPtrOutput)
}

// User-defined labels for the EntryGroup. **Note**: This field is non-authoritative, and will only manage the labels
// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o DataplexEntryGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where entry group will be created in.
func (o DataplexEntryGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The relative resource name of the EntryGroup, of the form:
// projects/{project_number}/locations/{location_id}/entryGroups/{entry_group_id}
func (o DataplexEntryGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataplexEntryGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataplexEntryGroupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DataplexEntryGroupOutput) Timeouts() DataplexEntryGroupTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) DataplexEntryGroupTimeoutsPtrOutput { return v.Timeouts }).(DataplexEntryGroupTimeoutsPtrOutput)
}

// Denotes the transfer status of the Entry Group. It is unspecified for Entry Group created from Dataplex API.
func (o DataplexEntryGroupOutput) TransferStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringOutput { return v.TransferStatus }).(pulumi.StringOutput)
}

// System generated globally unique ID for the EntryGroup. This ID will be different if the EntryGroup is deleted and
// re-created with the same name.
func (o DataplexEntryGroupOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the EntryGroup was last updated.
func (o DataplexEntryGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexEntryGroupInput)(nil)).Elem(), &DataplexEntryGroup{})
	pulumi.RegisterOutputType(DataplexEntryGroupOutput{})
}
