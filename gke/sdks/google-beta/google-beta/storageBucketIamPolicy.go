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

type StorageBucketIamPolicy struct {
	pulumi.CustomResourceState

	Bucket                   pulumi.StringOutput `pulumi:"bucket"`
	Etag                     pulumi.StringOutput `pulumi:"etag"`
	PolicyData               pulumi.StringOutput `pulumi:"policyData"`
	StorageBucketIamPolicyId pulumi.StringOutput `pulumi:"storageBucketIamPolicyId"`
}

// NewStorageBucketIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewStorageBucketIamPolicy(ctx *pulumi.Context,
	name string, args *StorageBucketIamPolicyArgs, opts ...pulumi.ResourceOption) (*StorageBucketIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageBucketIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/storageBucketIamPolicy:StorageBucketIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucketIamPolicy gets an existing StorageBucketIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucketIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketIamPolicyState, opts ...pulumi.ResourceOption) (*StorageBucketIamPolicy, error) {
	var resource StorageBucketIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/storageBucketIamPolicy:StorageBucketIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucketIamPolicy resources.
type storageBucketIamPolicyState struct {
	Bucket                   *string `pulumi:"bucket"`
	Etag                     *string `pulumi:"etag"`
	PolicyData               *string `pulumi:"policyData"`
	StorageBucketIamPolicyId *string `pulumi:"storageBucketIamPolicyId"`
}

type StorageBucketIamPolicyState struct {
	Bucket                   pulumi.StringPtrInput
	Etag                     pulumi.StringPtrInput
	PolicyData               pulumi.StringPtrInput
	StorageBucketIamPolicyId pulumi.StringPtrInput
}

func (StorageBucketIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketIamPolicyState)(nil)).Elem()
}

type storageBucketIamPolicyArgs struct {
	Bucket                   string  `pulumi:"bucket"`
	PolicyData               string  `pulumi:"policyData"`
	StorageBucketIamPolicyId *string `pulumi:"storageBucketIamPolicyId"`
}

// The set of arguments for constructing a StorageBucketIamPolicy resource.
type StorageBucketIamPolicyArgs struct {
	Bucket                   pulumi.StringInput
	PolicyData               pulumi.StringInput
	StorageBucketIamPolicyId pulumi.StringPtrInput
}

func (StorageBucketIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketIamPolicyArgs)(nil)).Elem()
}

type StorageBucketIamPolicyInput interface {
	pulumi.Input

	ToStorageBucketIamPolicyOutput() StorageBucketIamPolicyOutput
	ToStorageBucketIamPolicyOutputWithContext(ctx context.Context) StorageBucketIamPolicyOutput
}

func (*StorageBucketIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketIamPolicy)(nil)).Elem()
}

func (i *StorageBucketIamPolicy) ToStorageBucketIamPolicyOutput() StorageBucketIamPolicyOutput {
	return i.ToStorageBucketIamPolicyOutputWithContext(context.Background())
}

func (i *StorageBucketIamPolicy) ToStorageBucketIamPolicyOutputWithContext(ctx context.Context) StorageBucketIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketIamPolicyOutput)
}

type StorageBucketIamPolicyOutput struct{ *pulumi.OutputState }

func (StorageBucketIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketIamPolicy)(nil)).Elem()
}

func (o StorageBucketIamPolicyOutput) ToStorageBucketIamPolicyOutput() StorageBucketIamPolicyOutput {
	return o
}

func (o StorageBucketIamPolicyOutput) ToStorageBucketIamPolicyOutputWithContext(ctx context.Context) StorageBucketIamPolicyOutput {
	return o
}

func (o StorageBucketIamPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o StorageBucketIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o StorageBucketIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o StorageBucketIamPolicyOutput) StorageBucketIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketIamPolicy) pulumi.StringOutput { return v.StorageBucketIamPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketIamPolicyInput)(nil)).Elem(), &StorageBucketIamPolicy{})
	pulumi.RegisterOutputType(StorageBucketIamPolicyOutput{})
}
