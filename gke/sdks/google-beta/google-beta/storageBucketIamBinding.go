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

type StorageBucketIamBinding struct {
	pulumi.CustomResourceState

	Bucket                    pulumi.StringOutput                       `pulumi:"bucket"`
	Condition                 StorageBucketIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                      pulumi.StringOutput                       `pulumi:"etag"`
	Members                   pulumi.StringArrayOutput                  `pulumi:"members"`
	Role                      pulumi.StringOutput                       `pulumi:"role"`
	StorageBucketIamBindingId pulumi.StringOutput                       `pulumi:"storageBucketIamBindingId"`
}

// NewStorageBucketIamBinding registers a new resource with the given unique name, arguments, and options.
func NewStorageBucketIamBinding(ctx *pulumi.Context,
	name string, args *StorageBucketIamBindingArgs, opts ...pulumi.ResourceOption) (*StorageBucketIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageBucketIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/storageBucketIamBinding:StorageBucketIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucketIamBinding gets an existing StorageBucketIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucketIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketIamBindingState, opts ...pulumi.ResourceOption) (*StorageBucketIamBinding, error) {
	var resource StorageBucketIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/storageBucketIamBinding:StorageBucketIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucketIamBinding resources.
type storageBucketIamBindingState struct {
	Bucket                    *string                           `pulumi:"bucket"`
	Condition                 *StorageBucketIamBindingCondition `pulumi:"condition"`
	Etag                      *string                           `pulumi:"etag"`
	Members                   []string                          `pulumi:"members"`
	Role                      *string                           `pulumi:"role"`
	StorageBucketIamBindingId *string                           `pulumi:"storageBucketIamBindingId"`
}

type StorageBucketIamBindingState struct {
	Bucket                    pulumi.StringPtrInput
	Condition                 StorageBucketIamBindingConditionPtrInput
	Etag                      pulumi.StringPtrInput
	Members                   pulumi.StringArrayInput
	Role                      pulumi.StringPtrInput
	StorageBucketIamBindingId pulumi.StringPtrInput
}

func (StorageBucketIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketIamBindingState)(nil)).Elem()
}

type storageBucketIamBindingArgs struct {
	Bucket                    string                            `pulumi:"bucket"`
	Condition                 *StorageBucketIamBindingCondition `pulumi:"condition"`
	Members                   []string                          `pulumi:"members"`
	Role                      string                            `pulumi:"role"`
	StorageBucketIamBindingId *string                           `pulumi:"storageBucketIamBindingId"`
}

// The set of arguments for constructing a StorageBucketIamBinding resource.
type StorageBucketIamBindingArgs struct {
	Bucket                    pulumi.StringInput
	Condition                 StorageBucketIamBindingConditionPtrInput
	Members                   pulumi.StringArrayInput
	Role                      pulumi.StringInput
	StorageBucketIamBindingId pulumi.StringPtrInput
}

func (StorageBucketIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketIamBindingArgs)(nil)).Elem()
}

type StorageBucketIamBindingInput interface {
	pulumi.Input

	ToStorageBucketIamBindingOutput() StorageBucketIamBindingOutput
	ToStorageBucketIamBindingOutputWithContext(ctx context.Context) StorageBucketIamBindingOutput
}

func (*StorageBucketIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketIamBinding)(nil)).Elem()
}

func (i *StorageBucketIamBinding) ToStorageBucketIamBindingOutput() StorageBucketIamBindingOutput {
	return i.ToStorageBucketIamBindingOutputWithContext(context.Background())
}

func (i *StorageBucketIamBinding) ToStorageBucketIamBindingOutputWithContext(ctx context.Context) StorageBucketIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketIamBindingOutput)
}

type StorageBucketIamBindingOutput struct{ *pulumi.OutputState }

func (StorageBucketIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketIamBinding)(nil)).Elem()
}

func (o StorageBucketIamBindingOutput) ToStorageBucketIamBindingOutput() StorageBucketIamBindingOutput {
	return o
}

func (o StorageBucketIamBindingOutput) ToStorageBucketIamBindingOutputWithContext(ctx context.Context) StorageBucketIamBindingOutput {
	return o
}

func (o StorageBucketIamBindingOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamBinding) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o StorageBucketIamBindingOutput) Condition() StorageBucketIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *StorageBucketIamBinding) StorageBucketIamBindingConditionPtrOutput { return v.Condition }).(StorageBucketIamBindingConditionPtrOutput)
}

func (o StorageBucketIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o StorageBucketIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageBucketIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o StorageBucketIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o StorageBucketIamBindingOutput) StorageBucketIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamBinding) pulumi.StringOutput { return v.StorageBucketIamBindingId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketIamBindingInput)(nil)).Elem(), &StorageBucketIamBinding{})
	pulumi.RegisterOutputType(StorageBucketIamBindingOutput{})
}
