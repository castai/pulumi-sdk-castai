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

type ClouddeployCustomTargetType struct {
	pulumi.CustomResourceState

	// User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and size limitations. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations                   pulumi.StringMapOutput `pulumi:"annotations"`
	ClouddeployCustomTargetTypeId pulumi.StringOutput    `pulumi:"clouddeployCustomTargetTypeId"`
	// Time at which the 'CustomTargetType' was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Configures render and deploy for the 'CustomTargetType' using Skaffold custom actions.
	CustomActions ClouddeployCustomTargetTypeCustomActionsPtrOutput `pulumi:"customActions"`
	// Resource id of the 'CustomTargetType'.
	CustomTargetTypeId pulumi.StringOutput `pulumi:"customTargetTypeId"`
	// Description of the 'CustomTargetType'. Max length is 255 characters.
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveAnnotations pulumi.StringMapOutput `pulumi:"effectiveAnnotations"`
	EffectiveLabels      pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The weak etag of the 'CustomTargetType' resource. This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following
	// constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All
	// characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter
	// or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally
	// constrained to be <= 128 bytes. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location of the source.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the 'CustomTargetType'.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                       `pulumi:"terraformLabels"`
	Timeouts        ClouddeployCustomTargetTypeTimeoutsPtrOutput `pulumi:"timeouts"`
	// Unique identifier of the 'CustomTargetType'.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Time at which the 'CustomTargetType' was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewClouddeployCustomTargetType registers a new resource with the given unique name, arguments, and options.
func NewClouddeployCustomTargetType(ctx *pulumi.Context,
	name string, args *ClouddeployCustomTargetTypeArgs, opts ...pulumi.ResourceOption) (*ClouddeployCustomTargetType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ClouddeployCustomTargetType
	err = ctx.RegisterPackageResource("google:index/clouddeployCustomTargetType:ClouddeployCustomTargetType", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClouddeployCustomTargetType gets an existing ClouddeployCustomTargetType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClouddeployCustomTargetType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClouddeployCustomTargetTypeState, opts ...pulumi.ResourceOption) (*ClouddeployCustomTargetType, error) {
	var resource ClouddeployCustomTargetType
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/clouddeployCustomTargetType:ClouddeployCustomTargetType", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClouddeployCustomTargetType resources.
type clouddeployCustomTargetTypeState struct {
	// User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and size limitations. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations                   map[string]string `pulumi:"annotations"`
	ClouddeployCustomTargetTypeId *string           `pulumi:"clouddeployCustomTargetTypeId"`
	// Time at which the 'CustomTargetType' was created.
	CreateTime *string `pulumi:"createTime"`
	// Configures render and deploy for the 'CustomTargetType' using Skaffold custom actions.
	CustomActions *ClouddeployCustomTargetTypeCustomActions `pulumi:"customActions"`
	// Resource id of the 'CustomTargetType'.
	CustomTargetTypeId *string `pulumi:"customTargetTypeId"`
	// Description of the 'CustomTargetType'. Max length is 255 characters.
	Description          *string           `pulumi:"description"`
	EffectiveAnnotations map[string]string `pulumi:"effectiveAnnotations"`
	EffectiveLabels      map[string]string `pulumi:"effectiveLabels"`
	// The weak etag of the 'CustomTargetType' resource. This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following
	// constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All
	// characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter
	// or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally
	// constrained to be <= 128 bytes. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the source.
	Location *string `pulumi:"location"`
	// Name of the 'CustomTargetType'.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                    `pulumi:"terraformLabels"`
	Timeouts        *ClouddeployCustomTargetTypeTimeouts `pulumi:"timeouts"`
	// Unique identifier of the 'CustomTargetType'.
	Uid *string `pulumi:"uid"`
	// Time at which the 'CustomTargetType' was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type ClouddeployCustomTargetTypeState struct {
	// User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and size limitations. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations                   pulumi.StringMapInput
	ClouddeployCustomTargetTypeId pulumi.StringPtrInput
	// Time at which the 'CustomTargetType' was created.
	CreateTime pulumi.StringPtrInput
	// Configures render and deploy for the 'CustomTargetType' using Skaffold custom actions.
	CustomActions ClouddeployCustomTargetTypeCustomActionsPtrInput
	// Resource id of the 'CustomTargetType'.
	CustomTargetTypeId pulumi.StringPtrInput
	// Description of the 'CustomTargetType'. Max length is 255 characters.
	Description          pulumi.StringPtrInput
	EffectiveAnnotations pulumi.StringMapInput
	EffectiveLabels      pulumi.StringMapInput
	// The weak etag of the 'CustomTargetType' resource. This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following
	// constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All
	// characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter
	// or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally
	// constrained to be <= 128 bytes. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location of the source.
	Location pulumi.StringPtrInput
	// Name of the 'CustomTargetType'.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ClouddeployCustomTargetTypeTimeoutsPtrInput
	// Unique identifier of the 'CustomTargetType'.
	Uid pulumi.StringPtrInput
	// Time at which the 'CustomTargetType' was updated.
	UpdateTime pulumi.StringPtrInput
}

func (ClouddeployCustomTargetTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*clouddeployCustomTargetTypeState)(nil)).Elem()
}

type clouddeployCustomTargetTypeArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and size limitations. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations                   map[string]string `pulumi:"annotations"`
	ClouddeployCustomTargetTypeId *string           `pulumi:"clouddeployCustomTargetTypeId"`
	// Configures render and deploy for the 'CustomTargetType' using Skaffold custom actions.
	CustomActions *ClouddeployCustomTargetTypeCustomActions `pulumi:"customActions"`
	// Description of the 'CustomTargetType'. Max length is 255 characters.
	Description *string `pulumi:"description"`
	// Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following
	// constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All
	// characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter
	// or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally
	// constrained to be <= 128 bytes. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the source.
	Location string `pulumi:"location"`
	// Name of the 'CustomTargetType'.
	Name     *string                              `pulumi:"name"`
	Project  *string                              `pulumi:"project"`
	Timeouts *ClouddeployCustomTargetTypeTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ClouddeployCustomTargetType resource.
type ClouddeployCustomTargetTypeArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and size limitations. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations                   pulumi.StringMapInput
	ClouddeployCustomTargetTypeId pulumi.StringPtrInput
	// Configures render and deploy for the 'CustomTargetType' using Skaffold custom actions.
	CustomActions ClouddeployCustomTargetTypeCustomActionsPtrInput
	// Description of the 'CustomTargetType'. Max length is 255 characters.
	Description pulumi.StringPtrInput
	// Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following
	// constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All
	// characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter
	// or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally
	// constrained to be <= 128 bytes. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location of the source.
	Location pulumi.StringInput
	// Name of the 'CustomTargetType'.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ClouddeployCustomTargetTypeTimeoutsPtrInput
}

func (ClouddeployCustomTargetTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clouddeployCustomTargetTypeArgs)(nil)).Elem()
}

type ClouddeployCustomTargetTypeInput interface {
	pulumi.Input

	ToClouddeployCustomTargetTypeOutput() ClouddeployCustomTargetTypeOutput
	ToClouddeployCustomTargetTypeOutputWithContext(ctx context.Context) ClouddeployCustomTargetTypeOutput
}

func (*ClouddeployCustomTargetType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClouddeployCustomTargetType)(nil)).Elem()
}

func (i *ClouddeployCustomTargetType) ToClouddeployCustomTargetTypeOutput() ClouddeployCustomTargetTypeOutput {
	return i.ToClouddeployCustomTargetTypeOutputWithContext(context.Background())
}

func (i *ClouddeployCustomTargetType) ToClouddeployCustomTargetTypeOutputWithContext(ctx context.Context) ClouddeployCustomTargetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClouddeployCustomTargetTypeOutput)
}

type ClouddeployCustomTargetTypeOutput struct{ *pulumi.OutputState }

func (ClouddeployCustomTargetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClouddeployCustomTargetType)(nil)).Elem()
}

func (o ClouddeployCustomTargetTypeOutput) ToClouddeployCustomTargetTypeOutput() ClouddeployCustomTargetTypeOutput {
	return o
}

func (o ClouddeployCustomTargetTypeOutput) ToClouddeployCustomTargetTypeOutputWithContext(ctx context.Context) ClouddeployCustomTargetTypeOutput {
	return o
}

// User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See
// https://google.aip.dev/128#annotations for more details such as format and size limitations. **Note**: This field is
// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
// 'effective_annotations' for all of the annotations present on the resource.
func (o ClouddeployCustomTargetTypeOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

func (o ClouddeployCustomTargetTypeOutput) ClouddeployCustomTargetTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.ClouddeployCustomTargetTypeId }).(pulumi.StringOutput)
}

// Time at which the 'CustomTargetType' was created.
func (o ClouddeployCustomTargetTypeOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Configures render and deploy for the 'CustomTargetType' using Skaffold custom actions.
func (o ClouddeployCustomTargetTypeOutput) CustomActions() ClouddeployCustomTargetTypeCustomActionsPtrOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) ClouddeployCustomTargetTypeCustomActionsPtrOutput {
		return v.CustomActions
	}).(ClouddeployCustomTargetTypeCustomActionsPtrOutput)
}

// Resource id of the 'CustomTargetType'.
func (o ClouddeployCustomTargetTypeOutput) CustomTargetTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.CustomTargetTypeId }).(pulumi.StringOutput)
}

// Description of the 'CustomTargetType'. Max length is 255 characters.
func (o ClouddeployCustomTargetTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClouddeployCustomTargetTypeOutput) EffectiveAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringMapOutput { return v.EffectiveAnnotations }).(pulumi.StringMapOutput)
}

func (o ClouddeployCustomTargetTypeOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The weak etag of the 'CustomTargetType' resource. This checksum is computed by the server based on the value of other
// fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o ClouddeployCustomTargetTypeOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following
// constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All
// characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter
// or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally
// constrained to be <= 128 bytes. **Note**: This field is non-authoritative, and will only manage the labels present in
// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o ClouddeployCustomTargetTypeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location of the source.
func (o ClouddeployCustomTargetTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the 'CustomTargetType'.
func (o ClouddeployCustomTargetTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClouddeployCustomTargetTypeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ClouddeployCustomTargetTypeOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ClouddeployCustomTargetTypeOutput) Timeouts() ClouddeployCustomTargetTypeTimeoutsPtrOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) ClouddeployCustomTargetTypeTimeoutsPtrOutput { return v.Timeouts }).(ClouddeployCustomTargetTypeTimeoutsPtrOutput)
}

// Unique identifier of the 'CustomTargetType'.
func (o ClouddeployCustomTargetTypeOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Time at which the 'CustomTargetType' was updated.
func (o ClouddeployCustomTargetTypeOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployCustomTargetType) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClouddeployCustomTargetTypeInput)(nil)).Elem(), &ClouddeployCustomTargetType{})
	pulumi.RegisterOutputType(ClouddeployCustomTargetTypeOutput{})
}
