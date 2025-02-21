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

type TagsTagBinding struct {
	pulumi.CustomResourceState

	// The generated id for the TagBinding. This is a string of the form: 'tagBindings/{full-resource-name}/{tag-value-name}'
	Name pulumi.StringOutput `pulumi:"name"`
	// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
	Parent pulumi.StringOutput `pulumi:"parent"`
	// The TagValue of the TagBinding. Must be of the form tagValues/456.
	TagValue         pulumi.StringOutput             `pulumi:"tagValue"`
	TagsTagBindingId pulumi.StringOutput             `pulumi:"tagsTagBindingId"`
	Timeouts         TagsTagBindingTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewTagsTagBinding registers a new resource with the given unique name, arguments, and options.
func NewTagsTagBinding(ctx *pulumi.Context,
	name string, args *TagsTagBindingArgs, opts ...pulumi.ResourceOption) (*TagsTagBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.TagValue == nil {
		return nil, errors.New("invalid value for required argument 'TagValue'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource TagsTagBinding
	err = ctx.RegisterPackageResource("google:index/tagsTagBinding:TagsTagBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagsTagBinding gets an existing TagsTagBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagsTagBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagsTagBindingState, opts ...pulumi.ResourceOption) (*TagsTagBinding, error) {
	var resource TagsTagBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/tagsTagBinding:TagsTagBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagsTagBinding resources.
type tagsTagBindingState struct {
	// The generated id for the TagBinding. This is a string of the form: 'tagBindings/{full-resource-name}/{tag-value-name}'
	Name *string `pulumi:"name"`
	// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
	Parent *string `pulumi:"parent"`
	// The TagValue of the TagBinding. Must be of the form tagValues/456.
	TagValue         *string                 `pulumi:"tagValue"`
	TagsTagBindingId *string                 `pulumi:"tagsTagBindingId"`
	Timeouts         *TagsTagBindingTimeouts `pulumi:"timeouts"`
}

type TagsTagBindingState struct {
	// The generated id for the TagBinding. This is a string of the form: 'tagBindings/{full-resource-name}/{tag-value-name}'
	Name pulumi.StringPtrInput
	// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
	Parent pulumi.StringPtrInput
	// The TagValue of the TagBinding. Must be of the form tagValues/456.
	TagValue         pulumi.StringPtrInput
	TagsTagBindingId pulumi.StringPtrInput
	Timeouts         TagsTagBindingTimeoutsPtrInput
}

func (TagsTagBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsTagBindingState)(nil)).Elem()
}

type tagsTagBindingArgs struct {
	// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
	Parent string `pulumi:"parent"`
	// The TagValue of the TagBinding. Must be of the form tagValues/456.
	TagValue         string                  `pulumi:"tagValue"`
	TagsTagBindingId *string                 `pulumi:"tagsTagBindingId"`
	Timeouts         *TagsTagBindingTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a TagsTagBinding resource.
type TagsTagBindingArgs struct {
	// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
	Parent pulumi.StringInput
	// The TagValue of the TagBinding. Must be of the form tagValues/456.
	TagValue         pulumi.StringInput
	TagsTagBindingId pulumi.StringPtrInput
	Timeouts         TagsTagBindingTimeoutsPtrInput
}

func (TagsTagBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsTagBindingArgs)(nil)).Elem()
}

type TagsTagBindingInput interface {
	pulumi.Input

	ToTagsTagBindingOutput() TagsTagBindingOutput
	ToTagsTagBindingOutputWithContext(ctx context.Context) TagsTagBindingOutput
}

func (*TagsTagBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsTagBinding)(nil)).Elem()
}

func (i *TagsTagBinding) ToTagsTagBindingOutput() TagsTagBindingOutput {
	return i.ToTagsTagBindingOutputWithContext(context.Background())
}

func (i *TagsTagBinding) ToTagsTagBindingOutputWithContext(ctx context.Context) TagsTagBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsTagBindingOutput)
}

type TagsTagBindingOutput struct{ *pulumi.OutputState }

func (TagsTagBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsTagBinding)(nil)).Elem()
}

func (o TagsTagBindingOutput) ToTagsTagBindingOutput() TagsTagBindingOutput {
	return o
}

func (o TagsTagBindingOutput) ToTagsTagBindingOutputWithContext(ctx context.Context) TagsTagBindingOutput {
	return o
}

// The generated id for the TagBinding. This is a string of the form: 'tagBindings/{full-resource-name}/{tag-value-name}'
func (o TagsTagBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
func (o TagsTagBindingOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagBinding) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// The TagValue of the TagBinding. Must be of the form tagValues/456.
func (o TagsTagBindingOutput) TagValue() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagBinding) pulumi.StringOutput { return v.TagValue }).(pulumi.StringOutput)
}

func (o TagsTagBindingOutput) TagsTagBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagBinding) pulumi.StringOutput { return v.TagsTagBindingId }).(pulumi.StringOutput)
}

func (o TagsTagBindingOutput) Timeouts() TagsTagBindingTimeoutsPtrOutput {
	return o.ApplyT(func(v *TagsTagBinding) TagsTagBindingTimeoutsPtrOutput { return v.Timeouts }).(TagsTagBindingTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagsTagBindingInput)(nil)).Elem(), &TagsTagBinding{})
	pulumi.RegisterOutputType(TagsTagBindingOutput{})
}
