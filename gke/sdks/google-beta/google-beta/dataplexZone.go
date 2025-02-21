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

type DataplexZone struct {
	pulumi.CustomResourceState

	// Output only. Aggregated status of the underlying assets of the zone.
	AssetStatuses DataplexZoneAssetStatusArrayOutput `pulumi:"assetStatuses"`
	// Output only. The time when the zone was created.
	CreateTime     pulumi.StringOutput `pulumi:"createTime"`
	DataplexZoneId pulumi.StringOutput `pulumi:"dataplexZoneId"`
	// Optional. Description of the zone.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec DataplexZoneDiscoverySpecOutput `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName     pulumi.StringPtrOutput `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Optional. User defined labels for the zone. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The lake for the resource
	Lake pulumi.StringOutput `pulumi:"lake"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec DataplexZoneResourceSpecOutput `pulumi:"resourceSpec"`
	// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput        `pulumi:"terraformLabels"`
	Timeouts        DataplexZoneTimeoutsPtrOutput `pulumi:"timeouts"`
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type pulumi.StringOutput `pulumi:"type"`
	// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and
	// re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time when the zone was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDataplexZone registers a new resource with the given unique name, arguments, and options.
func NewDataplexZone(ctx *pulumi.Context,
	name string, args *DataplexZoneArgs, opts ...pulumi.ResourceOption) (*DataplexZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiscoverySpec == nil {
		return nil, errors.New("invalid value for required argument 'DiscoverySpec'")
	}
	if args.Lake == nil {
		return nil, errors.New("invalid value for required argument 'Lake'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceSpec == nil {
		return nil, errors.New("invalid value for required argument 'ResourceSpec'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexZone
	err = ctx.RegisterPackageResource("google-beta:index/dataplexZone:DataplexZone", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexZone gets an existing DataplexZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexZoneState, opts ...pulumi.ResourceOption) (*DataplexZone, error) {
	var resource DataplexZone
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataplexZone:DataplexZone", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexZone resources.
type dataplexZoneState struct {
	// Output only. Aggregated status of the underlying assets of the zone.
	AssetStatuses []DataplexZoneAssetStatus `pulumi:"assetStatuses"`
	// Output only. The time when the zone was created.
	CreateTime     *string `pulumi:"createTime"`
	DataplexZoneId *string `pulumi:"dataplexZoneId"`
	// Optional. Description of the zone.
	Description *string `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec *DataplexZoneDiscoverySpec `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Optional. User defined labels for the zone. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake for the resource
	Lake *string `pulumi:"lake"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// The name of the zone.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec *DataplexZoneResourceSpec `pulumi:"resourceSpec"`
	// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string     `pulumi:"terraformLabels"`
	Timeouts        *DataplexZoneTimeouts `pulumi:"timeouts"`
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type *string `pulumi:"type"`
	// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and
	// re-created with the same name.
	Uid *string `pulumi:"uid"`
	// Output only. The time when the zone was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type DataplexZoneState struct {
	// Output only. Aggregated status of the underlying assets of the zone.
	AssetStatuses DataplexZoneAssetStatusArrayInput
	// Output only. The time when the zone was created.
	CreateTime     pulumi.StringPtrInput
	DataplexZoneId pulumi.StringPtrInput
	// Optional. Description of the zone.
	Description pulumi.StringPtrInput
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec DataplexZoneDiscoverySpecPtrInput
	// Optional. User friendly display name.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Optional. User defined labels for the zone. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The lake for the resource
	Lake pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// The name of the zone.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec DataplexZoneResourceSpecPtrInput
	// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DataplexZoneTimeoutsPtrInput
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type pulumi.StringPtrInput
	// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and
	// re-created with the same name.
	Uid pulumi.StringPtrInput
	// Output only. The time when the zone was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (DataplexZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexZoneState)(nil)).Elem()
}

type dataplexZoneArgs struct {
	DataplexZoneId *string `pulumi:"dataplexZoneId"`
	// Optional. Description of the zone.
	Description *string `pulumi:"description"`
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec DataplexZoneDiscoverySpec `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// Optional. User defined labels for the zone. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake for the resource
	Lake string `pulumi:"lake"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The name of the zone.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec DataplexZoneResourceSpec `pulumi:"resourceSpec"`
	Timeouts     *DataplexZoneTimeouts    `pulumi:"timeouts"`
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DataplexZone resource.
type DataplexZoneArgs struct {
	DataplexZoneId pulumi.StringPtrInput
	// Optional. Description of the zone.
	Description pulumi.StringPtrInput
	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec DataplexZoneDiscoverySpecInput
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrInput
	// Optional. User defined labels for the zone. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The lake for the resource
	Lake pulumi.StringInput
	// The location for the resource
	Location pulumi.StringInput
	// The name of the zone.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec DataplexZoneResourceSpecInput
	Timeouts     DataplexZoneTimeoutsPtrInput
	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type pulumi.StringInput
}

func (DataplexZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexZoneArgs)(nil)).Elem()
}

type DataplexZoneInput interface {
	pulumi.Input

	ToDataplexZoneOutput() DataplexZoneOutput
	ToDataplexZoneOutputWithContext(ctx context.Context) DataplexZoneOutput
}

func (*DataplexZone) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexZone)(nil)).Elem()
}

func (i *DataplexZone) ToDataplexZoneOutput() DataplexZoneOutput {
	return i.ToDataplexZoneOutputWithContext(context.Background())
}

func (i *DataplexZone) ToDataplexZoneOutputWithContext(ctx context.Context) DataplexZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexZoneOutput)
}

type DataplexZoneOutput struct{ *pulumi.OutputState }

func (DataplexZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexZone)(nil)).Elem()
}

func (o DataplexZoneOutput) ToDataplexZoneOutput() DataplexZoneOutput {
	return o
}

func (o DataplexZoneOutput) ToDataplexZoneOutputWithContext(ctx context.Context) DataplexZoneOutput {
	return o
}

// Output only. Aggregated status of the underlying assets of the zone.
func (o DataplexZoneOutput) AssetStatuses() DataplexZoneAssetStatusArrayOutput {
	return o.ApplyT(func(v *DataplexZone) DataplexZoneAssetStatusArrayOutput { return v.AssetStatuses }).(DataplexZoneAssetStatusArrayOutput)
}

// Output only. The time when the zone was created.
func (o DataplexZoneOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o DataplexZoneOutput) DataplexZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.DataplexZoneId }).(pulumi.StringOutput)
}

// Optional. Description of the zone.
func (o DataplexZoneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Required. Specification of the discovery feature applied to data in this zone.
func (o DataplexZoneOutput) DiscoverySpec() DataplexZoneDiscoverySpecOutput {
	return o.ApplyT(func(v *DataplexZone) DataplexZoneDiscoverySpecOutput { return v.DiscoverySpec }).(DataplexZoneDiscoverySpecOutput)
}

// Optional. User friendly display name.
func (o DataplexZoneOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o DataplexZoneOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Optional. User defined labels for the zone. **Note**: This field is non-authoritative, and will only manage the labels
// present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the
// resource.
func (o DataplexZoneOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The lake for the resource
func (o DataplexZoneOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

// The location for the resource
func (o DataplexZoneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the zone.
func (o DataplexZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o DataplexZoneOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
func (o DataplexZoneOutput) ResourceSpec() DataplexZoneResourceSpecOutput {
	return o.ApplyT(func(v *DataplexZone) DataplexZoneResourceSpecOutput { return v.ResourceSpec }).(DataplexZoneResourceSpecOutput)
}

// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
func (o DataplexZoneOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataplexZoneOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DataplexZoneOutput) Timeouts() DataplexZoneTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataplexZone) DataplexZoneTimeoutsPtrOutput { return v.Timeouts }).(DataplexZoneTimeoutsPtrOutput)
}

// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
func (o DataplexZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and
// re-created with the same name.
func (o DataplexZoneOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The time when the zone was last updated.
func (o DataplexZoneOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexZone) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexZoneInput)(nil)).Elem(), &DataplexZone{})
	pulumi.RegisterOutputType(DataplexZoneOutput{})
}
