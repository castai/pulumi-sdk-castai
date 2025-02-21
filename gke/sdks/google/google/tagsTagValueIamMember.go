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

type TagsTagValueIamMember struct {
	pulumi.CustomResourceState

	Condition               TagsTagValueIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                    pulumi.StringOutput                     `pulumi:"etag"`
	Member                  pulumi.StringOutput                     `pulumi:"member"`
	Role                    pulumi.StringOutput                     `pulumi:"role"`
	TagValue                pulumi.StringOutput                     `pulumi:"tagValue"`
	TagsTagValueIamMemberId pulumi.StringOutput                     `pulumi:"tagsTagValueIamMemberId"`
}

// NewTagsTagValueIamMember registers a new resource with the given unique name, arguments, and options.
func NewTagsTagValueIamMember(ctx *pulumi.Context,
	name string, args *TagsTagValueIamMemberArgs, opts ...pulumi.ResourceOption) (*TagsTagValueIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
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
	var resource TagsTagValueIamMember
	err = ctx.RegisterPackageResource("google:index/tagsTagValueIamMember:TagsTagValueIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagsTagValueIamMember gets an existing TagsTagValueIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagsTagValueIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagsTagValueIamMemberState, opts ...pulumi.ResourceOption) (*TagsTagValueIamMember, error) {
	var resource TagsTagValueIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/tagsTagValueIamMember:TagsTagValueIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagsTagValueIamMember resources.
type tagsTagValueIamMemberState struct {
	Condition               *TagsTagValueIamMemberCondition `pulumi:"condition"`
	Etag                    *string                         `pulumi:"etag"`
	Member                  *string                         `pulumi:"member"`
	Role                    *string                         `pulumi:"role"`
	TagValue                *string                         `pulumi:"tagValue"`
	TagsTagValueIamMemberId *string                         `pulumi:"tagsTagValueIamMemberId"`
}

type TagsTagValueIamMemberState struct {
	Condition               TagsTagValueIamMemberConditionPtrInput
	Etag                    pulumi.StringPtrInput
	Member                  pulumi.StringPtrInput
	Role                    pulumi.StringPtrInput
	TagValue                pulumi.StringPtrInput
	TagsTagValueIamMemberId pulumi.StringPtrInput
}

func (TagsTagValueIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsTagValueIamMemberState)(nil)).Elem()
}

type tagsTagValueIamMemberArgs struct {
	Condition               *TagsTagValueIamMemberCondition `pulumi:"condition"`
	Member                  string                          `pulumi:"member"`
	Role                    string                          `pulumi:"role"`
	TagValue                string                          `pulumi:"tagValue"`
	TagsTagValueIamMemberId *string                         `pulumi:"tagsTagValueIamMemberId"`
}

// The set of arguments for constructing a TagsTagValueIamMember resource.
type TagsTagValueIamMemberArgs struct {
	Condition               TagsTagValueIamMemberConditionPtrInput
	Member                  pulumi.StringInput
	Role                    pulumi.StringInput
	TagValue                pulumi.StringInput
	TagsTagValueIamMemberId pulumi.StringPtrInput
}

func (TagsTagValueIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsTagValueIamMemberArgs)(nil)).Elem()
}

type TagsTagValueIamMemberInput interface {
	pulumi.Input

	ToTagsTagValueIamMemberOutput() TagsTagValueIamMemberOutput
	ToTagsTagValueIamMemberOutputWithContext(ctx context.Context) TagsTagValueIamMemberOutput
}

func (*TagsTagValueIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsTagValueIamMember)(nil)).Elem()
}

func (i *TagsTagValueIamMember) ToTagsTagValueIamMemberOutput() TagsTagValueIamMemberOutput {
	return i.ToTagsTagValueIamMemberOutputWithContext(context.Background())
}

func (i *TagsTagValueIamMember) ToTagsTagValueIamMemberOutputWithContext(ctx context.Context) TagsTagValueIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsTagValueIamMemberOutput)
}

type TagsTagValueIamMemberOutput struct{ *pulumi.OutputState }

func (TagsTagValueIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsTagValueIamMember)(nil)).Elem()
}

func (o TagsTagValueIamMemberOutput) ToTagsTagValueIamMemberOutput() TagsTagValueIamMemberOutput {
	return o
}

func (o TagsTagValueIamMemberOutput) ToTagsTagValueIamMemberOutputWithContext(ctx context.Context) TagsTagValueIamMemberOutput {
	return o
}

func (o TagsTagValueIamMemberOutput) Condition() TagsTagValueIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *TagsTagValueIamMember) TagsTagValueIamMemberConditionPtrOutput { return v.Condition }).(TagsTagValueIamMemberConditionPtrOutput)
}

func (o TagsTagValueIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TagsTagValueIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o TagsTagValueIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o TagsTagValueIamMemberOutput) TagValue() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamMember) pulumi.StringOutput { return v.TagValue }).(pulumi.StringOutput)
}

func (o TagsTagValueIamMemberOutput) TagsTagValueIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagsTagValueIamMember) pulumi.StringOutput { return v.TagsTagValueIamMemberId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagsTagValueIamMemberInput)(nil)).Elem(), &TagsTagValueIamMember{})
	pulumi.RegisterOutputType(TagsTagValueIamMemberOutput{})
}
