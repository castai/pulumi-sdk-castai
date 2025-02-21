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

type TagsTagValueIamBinding struct {
	pulumi.CustomResourceState

	Condition                TagsTagValueIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                     pulumi.StringOutput                      `pulumi:"etag"`
	Members                  pulumi.StringArrayOutput                 `pulumi:"members"`
	Role                     pulumi.StringOutput                      `pulumi:"role"`
	TagValue                 pulumi.StringOutput                      `pulumi:"tagValue"`
	TagsTagValueIamBindingId pulumi.StringOutput                      `pulumi:"tagsTagValueIamBindingId"`
}

// NewTagsTagValueIamBinding registers a new resource with the given unique name, arguments, and options.
func NewTagsTagValueIamBinding(ctx *pulumi.Context,
	name string, args *TagsTagValueIamBindingArgs, opts ...pulumi.ResourceOption) (*TagsTagValueIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TagValue == nil {
		return nil, errors.New("invalid value for required argument 'TagValue'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource TagsTagValueIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/tagsTagValueIamBinding:TagsTagValueIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagsTagValueIamBinding gets an existing TagsTagValueIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagsTagValueIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagsTagValueIamBindingState, opts ...pulumi.ResourceOption) (*TagsTagValueIamBinding, error) {
	var resource TagsTagValueIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/tagsTagValueIamBinding:TagsTagValueIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagsTagValueIamBinding resources.
type tagsTagValueIamBindingState struct {
	Condition                *TagsTagValueIamBindingCondition `pulumi:"condition"`
	Etag                     *string                          `pulumi:"etag"`
	Members                  []string                         `pulumi:"members"`
	Role                     *string                          `pulumi:"role"`
	TagValue                 *string                          `pulumi:"tagValue"`
	TagsTagValueIamBindingId *string                          `pulumi:"tagsTagValueIamBindingId"`
}

type TagsTagValueIamBindingState struct {
	Condition                TagsTagValueIamBindingConditionPtrInput
	Etag                     pulumi.StringPtrInput
	Members                  pulumi.StringArrayInput
	Role                     pulumi.StringPtrInput
	TagValue                 pulumi.StringPtrInput
	TagsTagValueIamBindingId pulumi.StringPtrInput
}

func (TagsTagValueIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsTagValueIamBindingState)(nil)).Elem()
}

type tagsTagValueIamBindingArgs struct {
	Condition                *TagsTagValueIamBindingCondition `pulumi:"condition"`
	Members                  []string                         `pulumi:"members"`
	Role                     string                           `pulumi:"role"`
	TagValue                 string                           `pulumi:"tagValue"`
	TagsTagValueIamBindingId *string                          `pulumi:"tagsTagValueIamBindingId"`
}

// The set of arguments for constructing a TagsTagValueIamBinding resource.
type TagsTagValueIamBindingArgs struct {
	Condition                TagsTagValueIamBindingConditionPtrInput
	Members                  pulumi.StringArrayInput
	Role                     pulumi.StringInput
	TagValue                 pulumi.StringInput
	TagsTagValueIamBindingId pulumi.StringPtrInput
}

func (TagsTagValueIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsTagValueIamBindingArgs)(nil)).Elem()
}

type TagsTagValueIamBindingInput interface {
	pulumi.Input

	ToTagsTagValueIamBindingOutput() TagsTagValueIamBindingOutput
	ToTagsTagValueIamBindingOutputWithContext(ctx context.Context) TagsTagValueIamBindingOutput
}

func (*TagsTagValueIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsTagValueIamBinding)(nil)).Elem()
}

func (i *TagsTagValueIamBinding) ToTagsTagValueIamBindingOutput() TagsTagValueIamBindingOutput {
	return i.ToTagsTagValueIamBindingOutputWithContext(context.Background())
}

func (i *TagsTagValueIamBinding) ToTagsTagValueIamBindingOutputWithContext(ctx context.Context) TagsTagValueIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsTagValueIamBindingOutput)
}

type TagsTagValueIamBindingOutput struct{ *pulumi.OutputState }

func (TagsTagValueIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsTagValueIamBinding)(nil)).Elem()
}

func (o TagsTagValueIamBindingOutput) ToTagsTagValueIamBindingOutput() TagsTagValueIamBindingOutput {
	return o
}

func (o TagsTagValueIamBindingOutput) ToTagsTagValueIamBindingOutputWithContext(ctx context.Context) TagsTagValueIamBindingOutput {
	return o
}

func (o TagsTagValueIamBindingOutput) Condition() TagsTagValueIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *TagsTagValueIamBinding) TagsTagValueIamBindingConditionPtrOutput { return v.Condition }).(TagsTagValueIamBindingConditionPtrOutput)
}

func (o TagsTagValueIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TagsTagValueIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TagsTagValueIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o TagsTagValueIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o TagsTagValueIamBindingOutput) TagValue() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamBinding) pulumi.StringOutput { return v.TagValue }).(pulumi.StringOutput)
}

func (o TagsTagValueIamBindingOutput) TagsTagValueIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamBinding) pulumi.StringOutput { return v.TagsTagValueIamBindingId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagsTagValueIamBindingInput)(nil)).Elem(), &TagsTagValueIamBinding{})
	pulumi.RegisterOutputType(TagsTagValueIamBindingOutput{})
}
