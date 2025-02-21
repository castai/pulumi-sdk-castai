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

type StorageBucketIamMember struct {
	pulumi.CustomResourceState

	Bucket                   pulumi.StringOutput                      `pulumi:"bucket"`
	Condition                StorageBucketIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                     pulumi.StringOutput                      `pulumi:"etag"`
	Member                   pulumi.StringOutput                      `pulumi:"member"`
	Role                     pulumi.StringOutput                      `pulumi:"role"`
	StorageBucketIamMemberId pulumi.StringOutput                      `pulumi:"storageBucketIamMemberId"`
}

// NewStorageBucketIamMember registers a new resource with the given unique name, arguments, and options.
func NewStorageBucketIamMember(ctx *pulumi.Context,
	name string, args *StorageBucketIamMemberArgs, opts ...pulumi.ResourceOption) (*StorageBucketIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageBucketIamMember
	err = ctx.RegisterPackageResource("google-beta:index/storageBucketIamMember:StorageBucketIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucketIamMember gets an existing StorageBucketIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucketIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketIamMemberState, opts ...pulumi.ResourceOption) (*StorageBucketIamMember, error) {
	var resource StorageBucketIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/storageBucketIamMember:StorageBucketIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucketIamMember resources.
type storageBucketIamMemberState struct {
	Bucket                   *string                          `pulumi:"bucket"`
	Condition                *StorageBucketIamMemberCondition `pulumi:"condition"`
	Etag                     *string                          `pulumi:"etag"`
	Member                   *string                          `pulumi:"member"`
	Role                     *string                          `pulumi:"role"`
	StorageBucketIamMemberId *string                          `pulumi:"storageBucketIamMemberId"`
}

type StorageBucketIamMemberState struct {
	Bucket                   pulumi.StringPtrInput
	Condition                StorageBucketIamMemberConditionPtrInput
	Etag                     pulumi.StringPtrInput
	Member                   pulumi.StringPtrInput
	Role                     pulumi.StringPtrInput
	StorageBucketIamMemberId pulumi.StringPtrInput
}

func (StorageBucketIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketIamMemberState)(nil)).Elem()
}

type storageBucketIamMemberArgs struct {
	Bucket                   string                           `pulumi:"bucket"`
	Condition                *StorageBucketIamMemberCondition `pulumi:"condition"`
	Member                   string                           `pulumi:"member"`
	Role                     string                           `pulumi:"role"`
	StorageBucketIamMemberId *string                          `pulumi:"storageBucketIamMemberId"`
}

// The set of arguments for constructing a StorageBucketIamMember resource.
type StorageBucketIamMemberArgs struct {
	Bucket                   pulumi.StringInput
	Condition                StorageBucketIamMemberConditionPtrInput
	Member                   pulumi.StringInput
	Role                     pulumi.StringInput
	StorageBucketIamMemberId pulumi.StringPtrInput
}

func (StorageBucketIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketIamMemberArgs)(nil)).Elem()
}

type StorageBucketIamMemberInput interface {
	pulumi.Input

	ToStorageBucketIamMemberOutput() StorageBucketIamMemberOutput
	ToStorageBucketIamMemberOutputWithContext(ctx context.Context) StorageBucketIamMemberOutput
}

func (*StorageBucketIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketIamMember)(nil)).Elem()
}

func (i *StorageBucketIamMember) ToStorageBucketIamMemberOutput() StorageBucketIamMemberOutput {
	return i.ToStorageBucketIamMemberOutputWithContext(context.Background())
}

func (i *StorageBucketIamMember) ToStorageBucketIamMemberOutputWithContext(ctx context.Context) StorageBucketIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketIamMemberOutput)
}

type StorageBucketIamMemberOutput struct{ *pulumi.OutputState }

func (StorageBucketIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketIamMember)(nil)).Elem()
}

func (o StorageBucketIamMemberOutput) ToStorageBucketIamMemberOutput() StorageBucketIamMemberOutput {
	return o
}

func (o StorageBucketIamMemberOutput) ToStorageBucketIamMemberOutputWithContext(ctx context.Context) StorageBucketIamMemberOutput {
	return o
}

func (o StorageBucketIamMemberOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamMember) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o StorageBucketIamMemberOutput) Condition() StorageBucketIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *StorageBucketIamMember) StorageBucketIamMemberConditionPtrOutput { return v.Condition }).(StorageBucketIamMemberConditionPtrOutput)
}

func (o StorageBucketIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o StorageBucketIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o StorageBucketIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o StorageBucketIamMemberOutput) StorageBucketIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamMember) pulumi.StringOutput { return v.StorageBucketIamMemberId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketIamMemberInput)(nil)).Elem(), &StorageBucketIamMember{})
	pulumi.RegisterOutputType(StorageBucketIamMemberOutput{})
}
