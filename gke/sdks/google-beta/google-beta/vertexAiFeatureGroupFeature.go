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

type VertexAiFeatureGroupFeature struct {
	pulumi.CustomResourceState

	// The timestamp of when the FeatureGroup was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the FeatureGroup.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The name of the Feature Group.
	FeatureGroup pulumi.StringOutput `pulumi:"featureGroup"`
	// The labels with user-defined metadata to organize your FeatureGroup. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name of the Feature Group Feature.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The region for the resource. It should be the same as the feature group's region.
	Region pulumi.StringOutput `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                       `pulumi:"terraformLabels"`
	Timeouts        VertexAiFeatureGroupFeatureTimeoutsPtrOutput `pulumi:"timeouts"`
	// The timestamp of when the FeatureGroup was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The name of the BigQuery Table/View column hosting data for this version. If no value is provided, will use featureId.
	VersionColumnName             pulumi.StringOutput `pulumi:"versionColumnName"`
	VertexAiFeatureGroupFeatureId pulumi.StringOutput `pulumi:"vertexAiFeatureGroupFeatureId"`
}

// NewVertexAiFeatureGroupFeature registers a new resource with the given unique name, arguments, and options.
func NewVertexAiFeatureGroupFeature(ctx *pulumi.Context,
	name string, args *VertexAiFeatureGroupFeatureArgs, opts ...pulumi.ResourceOption) (*VertexAiFeatureGroupFeature, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeatureGroup == nil {
		return nil, errors.New("invalid value for required argument 'FeatureGroup'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource VertexAiFeatureGroupFeature
	err = ctx.RegisterPackageResource("google-beta:index/vertexAiFeatureGroupFeature:VertexAiFeatureGroupFeature", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVertexAiFeatureGroupFeature gets an existing VertexAiFeatureGroupFeature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVertexAiFeatureGroupFeature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VertexAiFeatureGroupFeatureState, opts ...pulumi.ResourceOption) (*VertexAiFeatureGroupFeature, error) {
	var resource VertexAiFeatureGroupFeature
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/vertexAiFeatureGroupFeature:VertexAiFeatureGroupFeature", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VertexAiFeatureGroupFeature resources.
type vertexAiFeatureGroupFeatureState struct {
	// The timestamp of when the FeatureGroup was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime *string `pulumi:"createTime"`
	// The description of the FeatureGroup.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The name of the Feature Group.
	FeatureGroup *string `pulumi:"featureGroup"`
	// The labels with user-defined metadata to organize your FeatureGroup. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the Feature Group Feature.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The region for the resource. It should be the same as the feature group's region.
	Region *string `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                    `pulumi:"terraformLabels"`
	Timeouts        *VertexAiFeatureGroupFeatureTimeouts `pulumi:"timeouts"`
	// The timestamp of when the FeatureGroup was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime *string `pulumi:"updateTime"`
	// The name of the BigQuery Table/View column hosting data for this version. If no value is provided, will use featureId.
	VersionColumnName             *string `pulumi:"versionColumnName"`
	VertexAiFeatureGroupFeatureId *string `pulumi:"vertexAiFeatureGroupFeatureId"`
}

type VertexAiFeatureGroupFeatureState struct {
	// The timestamp of when the FeatureGroup was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime pulumi.StringPtrInput
	// The description of the FeatureGroup.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The name of the Feature Group.
	FeatureGroup pulumi.StringPtrInput
	// The labels with user-defined metadata to organize your FeatureGroup. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The resource name of the Feature Group Feature.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region for the resource. It should be the same as the feature group's region.
	Region pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        VertexAiFeatureGroupFeatureTimeoutsPtrInput
	// The timestamp of when the FeatureGroup was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime pulumi.StringPtrInput
	// The name of the BigQuery Table/View column hosting data for this version. If no value is provided, will use featureId.
	VersionColumnName             pulumi.StringPtrInput
	VertexAiFeatureGroupFeatureId pulumi.StringPtrInput
}

func (VertexAiFeatureGroupFeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiFeatureGroupFeatureState)(nil)).Elem()
}

type vertexAiFeatureGroupFeatureArgs struct {
	// The description of the FeatureGroup.
	Description *string `pulumi:"description"`
	// The name of the Feature Group.
	FeatureGroup string `pulumi:"featureGroup"`
	// The labels with user-defined metadata to organize your FeatureGroup. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the Feature Group Feature.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The region for the resource. It should be the same as the feature group's region.
	Region   string                               `pulumi:"region"`
	Timeouts *VertexAiFeatureGroupFeatureTimeouts `pulumi:"timeouts"`
	// The name of the BigQuery Table/View column hosting data for this version. If no value is provided, will use featureId.
	VersionColumnName             *string `pulumi:"versionColumnName"`
	VertexAiFeatureGroupFeatureId *string `pulumi:"vertexAiFeatureGroupFeatureId"`
}

// The set of arguments for constructing a VertexAiFeatureGroupFeature resource.
type VertexAiFeatureGroupFeatureArgs struct {
	// The description of the FeatureGroup.
	Description pulumi.StringPtrInput
	// The name of the Feature Group.
	FeatureGroup pulumi.StringInput
	// The labels with user-defined metadata to organize your FeatureGroup. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The resource name of the Feature Group Feature.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region for the resource. It should be the same as the feature group's region.
	Region   pulumi.StringInput
	Timeouts VertexAiFeatureGroupFeatureTimeoutsPtrInput
	// The name of the BigQuery Table/View column hosting data for this version. If no value is provided, will use featureId.
	VersionColumnName             pulumi.StringPtrInput
	VertexAiFeatureGroupFeatureId pulumi.StringPtrInput
}

func (VertexAiFeatureGroupFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiFeatureGroupFeatureArgs)(nil)).Elem()
}

type VertexAiFeatureGroupFeatureInput interface {
	pulumi.Input

	ToVertexAiFeatureGroupFeatureOutput() VertexAiFeatureGroupFeatureOutput
	ToVertexAiFeatureGroupFeatureOutputWithContext(ctx context.Context) VertexAiFeatureGroupFeatureOutput
}

func (*VertexAiFeatureGroupFeature) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiFeatureGroupFeature)(nil)).Elem()
}

func (i *VertexAiFeatureGroupFeature) ToVertexAiFeatureGroupFeatureOutput() VertexAiFeatureGroupFeatureOutput {
	return i.ToVertexAiFeatureGroupFeatureOutputWithContext(context.Background())
}

func (i *VertexAiFeatureGroupFeature) ToVertexAiFeatureGroupFeatureOutputWithContext(ctx context.Context) VertexAiFeatureGroupFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VertexAiFeatureGroupFeatureOutput)
}

type VertexAiFeatureGroupFeatureOutput struct{ *pulumi.OutputState }

func (VertexAiFeatureGroupFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiFeatureGroupFeature)(nil)).Elem()
}

func (o VertexAiFeatureGroupFeatureOutput) ToVertexAiFeatureGroupFeatureOutput() VertexAiFeatureGroupFeatureOutput {
	return o
}

func (o VertexAiFeatureGroupFeatureOutput) ToVertexAiFeatureGroupFeatureOutputWithContext(ctx context.Context) VertexAiFeatureGroupFeatureOutput {
	return o
}

// The timestamp of when the FeatureGroup was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
// nine fractional digits.
func (o VertexAiFeatureGroupFeatureOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the FeatureGroup.
func (o VertexAiFeatureGroupFeatureOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VertexAiFeatureGroupFeatureOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The name of the Feature Group.
func (o VertexAiFeatureGroupFeatureOutput) FeatureGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.FeatureGroup }).(pulumi.StringOutput)
}

// The labels with user-defined metadata to organize your FeatureGroup. **Note**: This field is non-authoritative, and will
// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o VertexAiFeatureGroupFeatureOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name of the Feature Group Feature.
func (o VertexAiFeatureGroupFeatureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VertexAiFeatureGroupFeatureOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region for the resource. It should be the same as the feature group's region.
func (o VertexAiFeatureGroupFeatureOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o VertexAiFeatureGroupFeatureOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o VertexAiFeatureGroupFeatureOutput) Timeouts() VertexAiFeatureGroupFeatureTimeoutsPtrOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) VertexAiFeatureGroupFeatureTimeoutsPtrOutput { return v.Timeouts }).(VertexAiFeatureGroupFeatureTimeoutsPtrOutput)
}

// The timestamp of when the FeatureGroup was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
// to nine fractional digits.
func (o VertexAiFeatureGroupFeatureOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The name of the BigQuery Table/View column hosting data for this version. If no value is provided, will use featureId.
func (o VertexAiFeatureGroupFeatureOutput) VersionColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.VersionColumnName }).(pulumi.StringOutput)
}

func (o VertexAiFeatureGroupFeatureOutput) VertexAiFeatureGroupFeatureId() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeatureGroupFeature) pulumi.StringOutput { return v.VertexAiFeatureGroupFeatureId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VertexAiFeatureGroupFeatureInput)(nil)).Elem(), &VertexAiFeatureGroupFeature{})
	pulumi.RegisterOutputType(VertexAiFeatureGroupFeatureOutput{})
}
