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

type GeminiCodeRepositoryIndex struct {
	pulumi.CustomResourceState

	// Required. Id of the Code Repository Index.
	CodeRepositoryIndexId pulumi.StringOutput `pulumi:"codeRepositoryIndexId"`
	// Output only. Create time stamp.
	CreateTime      pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the
	// resource. These RepositoryGroups will also be deleted.
	ForceDestroy                pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	GeminiCodeRepositoryIndexId pulumi.StringOutput  `pulumi:"geminiCodeRepositoryIndexId"`
	// Optional. Immutable. Customer-managed encryption key name, in the format
	// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.
	KmsKey pulumi.StringPtrOutput `pulumi:"kmsKey"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location of the Code Repository Index, for example 'us-central1'.
	Location pulumi.StringOutput `pulumi:"location"`
	// Immutable. Identifier. Name of Code Repository Index.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. Code Repository Index instance State. Possible values are: 'STATE_UNSPECIFIED', 'CREATING', 'ACTIVE',
	// 'DELETING', 'SUSPENDED'.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                     `pulumi:"terraformLabels"`
	Timeouts        GeminiCodeRepositoryIndexTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. Update time stamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGeminiCodeRepositoryIndex registers a new resource with the given unique name, arguments, and options.
func NewGeminiCodeRepositoryIndex(ctx *pulumi.Context,
	name string, args *GeminiCodeRepositoryIndexArgs, opts ...pulumi.ResourceOption) (*GeminiCodeRepositoryIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CodeRepositoryIndexId == nil {
		return nil, errors.New("invalid value for required argument 'CodeRepositoryIndexId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GeminiCodeRepositoryIndex
	err = ctx.RegisterPackageResource("google:index/geminiCodeRepositoryIndex:GeminiCodeRepositoryIndex", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeminiCodeRepositoryIndex gets an existing GeminiCodeRepositoryIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeminiCodeRepositoryIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeminiCodeRepositoryIndexState, opts ...pulumi.ResourceOption) (*GeminiCodeRepositoryIndex, error) {
	var resource GeminiCodeRepositoryIndex
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/geminiCodeRepositoryIndex:GeminiCodeRepositoryIndex", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeminiCodeRepositoryIndex resources.
type geminiCodeRepositoryIndexState struct {
	// Required. Id of the Code Repository Index.
	CodeRepositoryIndexId *string `pulumi:"codeRepositoryIndexId"`
	// Output only. Create time stamp.
	CreateTime      *string           `pulumi:"createTime"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the
	// resource. These RepositoryGroups will also be deleted.
	ForceDestroy                *bool   `pulumi:"forceDestroy"`
	GeminiCodeRepositoryIndexId *string `pulumi:"geminiCodeRepositoryIndexId"`
	// Optional. Immutable. Customer-managed encryption key name, in the format
	// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.
	KmsKey *string `pulumi:"kmsKey"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the Code Repository Index, for example 'us-central1'.
	Location *string `pulumi:"location"`
	// Immutable. Identifier. Name of Code Repository Index.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Output only. Code Repository Index instance State. Possible values are: 'STATE_UNSPECIFIED', 'CREATING', 'ACTIVE',
	// 'DELETING', 'SUSPENDED'.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                  `pulumi:"terraformLabels"`
	Timeouts        *GeminiCodeRepositoryIndexTimeouts `pulumi:"timeouts"`
	// Output only. Update time stamp.
	UpdateTime *string `pulumi:"updateTime"`
}

type GeminiCodeRepositoryIndexState struct {
	// Required. Id of the Code Repository Index.
	CodeRepositoryIndexId pulumi.StringPtrInput
	// Output only. Create time stamp.
	CreateTime      pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the
	// resource. These RepositoryGroups will also be deleted.
	ForceDestroy                pulumi.BoolPtrInput
	GeminiCodeRepositoryIndexId pulumi.StringPtrInput
	// Optional. Immutable. Customer-managed encryption key name, in the format
	// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.
	KmsKey pulumi.StringPtrInput
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location of the Code Repository Index, for example 'us-central1'.
	Location pulumi.StringPtrInput
	// Immutable. Identifier. Name of Code Repository Index.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Output only. Code Repository Index instance State. Possible values are: 'STATE_UNSPECIFIED', 'CREATING', 'ACTIVE',
	// 'DELETING', 'SUSPENDED'.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        GeminiCodeRepositoryIndexTimeoutsPtrInput
	// Output only. Update time stamp.
	UpdateTime pulumi.StringPtrInput
}

func (GeminiCodeRepositoryIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*geminiCodeRepositoryIndexState)(nil)).Elem()
}

type geminiCodeRepositoryIndexArgs struct {
	// Required. Id of the Code Repository Index.
	CodeRepositoryIndexId string `pulumi:"codeRepositoryIndexId"`
	// If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the
	// resource. These RepositoryGroups will also be deleted.
	ForceDestroy                *bool   `pulumi:"forceDestroy"`
	GeminiCodeRepositoryIndexId *string `pulumi:"geminiCodeRepositoryIndexId"`
	// Optional. Immutable. Customer-managed encryption key name, in the format
	// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.
	KmsKey *string `pulumi:"kmsKey"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the Code Repository Index, for example 'us-central1'.
	Location string                             `pulumi:"location"`
	Project  *string                            `pulumi:"project"`
	Timeouts *GeminiCodeRepositoryIndexTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a GeminiCodeRepositoryIndex resource.
type GeminiCodeRepositoryIndexArgs struct {
	// Required. Id of the Code Repository Index.
	CodeRepositoryIndexId pulumi.StringInput
	// If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the
	// resource. These RepositoryGroups will also be deleted.
	ForceDestroy                pulumi.BoolPtrInput
	GeminiCodeRepositoryIndexId pulumi.StringPtrInput
	// Optional. Immutable. Customer-managed encryption key name, in the format
	// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.
	KmsKey pulumi.StringPtrInput
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location of the Code Repository Index, for example 'us-central1'.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	Timeouts GeminiCodeRepositoryIndexTimeoutsPtrInput
}

func (GeminiCodeRepositoryIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geminiCodeRepositoryIndexArgs)(nil)).Elem()
}

type GeminiCodeRepositoryIndexInput interface {
	pulumi.Input

	ToGeminiCodeRepositoryIndexOutput() GeminiCodeRepositoryIndexOutput
	ToGeminiCodeRepositoryIndexOutputWithContext(ctx context.Context) GeminiCodeRepositoryIndexOutput
}

func (*GeminiCodeRepositoryIndex) ElementType() reflect.Type {
	return reflect.TypeOf((**GeminiCodeRepositoryIndex)(nil)).Elem()
}

func (i *GeminiCodeRepositoryIndex) ToGeminiCodeRepositoryIndexOutput() GeminiCodeRepositoryIndexOutput {
	return i.ToGeminiCodeRepositoryIndexOutputWithContext(context.Background())
}

func (i *GeminiCodeRepositoryIndex) ToGeminiCodeRepositoryIndexOutputWithContext(ctx context.Context) GeminiCodeRepositoryIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeminiCodeRepositoryIndexOutput)
}

type GeminiCodeRepositoryIndexOutput struct{ *pulumi.OutputState }

func (GeminiCodeRepositoryIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeminiCodeRepositoryIndex)(nil)).Elem()
}

func (o GeminiCodeRepositoryIndexOutput) ToGeminiCodeRepositoryIndexOutput() GeminiCodeRepositoryIndexOutput {
	return o
}

func (o GeminiCodeRepositoryIndexOutput) ToGeminiCodeRepositoryIndexOutputWithContext(ctx context.Context) GeminiCodeRepositoryIndexOutput {
	return o
}

// Required. Id of the Code Repository Index.
func (o GeminiCodeRepositoryIndexOutput) CodeRepositoryIndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.CodeRepositoryIndexId }).(pulumi.StringOutput)
}

// Output only. Create time stamp.
func (o GeminiCodeRepositoryIndexOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GeminiCodeRepositoryIndexOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the
// resource. These RepositoryGroups will also be deleted.
func (o GeminiCodeRepositoryIndexOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

func (o GeminiCodeRepositoryIndexOutput) GeminiCodeRepositoryIndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.GeminiCodeRepositoryIndexId }).(pulumi.StringOutput)
}

// Optional. Immutable. Customer-managed encryption key name, in the format
// 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.
func (o GeminiCodeRepositoryIndexOutput) KmsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringPtrOutput { return v.KmsKey }).(pulumi.StringPtrOutput)
}

// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o GeminiCodeRepositoryIndexOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location of the Code Repository Index, for example 'us-central1'.
func (o GeminiCodeRepositoryIndexOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. Identifier. Name of Code Repository Index.
func (o GeminiCodeRepositoryIndexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GeminiCodeRepositoryIndexOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Code Repository Index instance State. Possible values are: 'STATE_UNSPECIFIED', 'CREATING', 'ACTIVE',
// 'DELETING', 'SUSPENDED'.
func (o GeminiCodeRepositoryIndexOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o GeminiCodeRepositoryIndexOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o GeminiCodeRepositoryIndexOutput) Timeouts() GeminiCodeRepositoryIndexTimeoutsPtrOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) GeminiCodeRepositoryIndexTimeoutsPtrOutput { return v.Timeouts }).(GeminiCodeRepositoryIndexTimeoutsPtrOutput)
}

// Output only. Update time stamp.
func (o GeminiCodeRepositoryIndexOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiCodeRepositoryIndex) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GeminiCodeRepositoryIndexInput)(nil)).Elem(), &GeminiCodeRepositoryIndex{})
	pulumi.RegisterOutputType(GeminiCodeRepositoryIndexOutput{})
}
