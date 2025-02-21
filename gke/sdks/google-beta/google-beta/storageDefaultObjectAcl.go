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

type StorageDefaultObjectAcl struct {
	pulumi.CustomResourceState

	Bucket                    pulumi.StringOutput      `pulumi:"bucket"`
	RoleEntities              pulumi.StringArrayOutput `pulumi:"roleEntities"`
	StorageDefaultObjectAclId pulumi.StringOutput      `pulumi:"storageDefaultObjectAclId"`
}

// NewStorageDefaultObjectAcl registers a new resource with the given unique name, arguments, and options.
func NewStorageDefaultObjectAcl(ctx *pulumi.Context,
	name string, args *StorageDefaultObjectAclArgs, opts ...pulumi.ResourceOption) (*StorageDefaultObjectAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageDefaultObjectAcl
	err = ctx.RegisterPackageResource("google-beta:index/storageDefaultObjectAcl:StorageDefaultObjectAcl", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageDefaultObjectAcl gets an existing StorageDefaultObjectAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageDefaultObjectAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageDefaultObjectAclState, opts ...pulumi.ResourceOption) (*StorageDefaultObjectAcl, error) {
	var resource StorageDefaultObjectAcl
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/storageDefaultObjectAcl:StorageDefaultObjectAcl", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageDefaultObjectAcl resources.
type storageDefaultObjectAclState struct {
	Bucket                    *string  `pulumi:"bucket"`
	RoleEntities              []string `pulumi:"roleEntities"`
	StorageDefaultObjectAclId *string  `pulumi:"storageDefaultObjectAclId"`
}

type StorageDefaultObjectAclState struct {
	Bucket                    pulumi.StringPtrInput
	RoleEntities              pulumi.StringArrayInput
	StorageDefaultObjectAclId pulumi.StringPtrInput
}

func (StorageDefaultObjectAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDefaultObjectAclState)(nil)).Elem()
}

type storageDefaultObjectAclArgs struct {
	Bucket                    string   `pulumi:"bucket"`
	RoleEntities              []string `pulumi:"roleEntities"`
	StorageDefaultObjectAclId *string  `pulumi:"storageDefaultObjectAclId"`
}

// The set of arguments for constructing a StorageDefaultObjectAcl resource.
type StorageDefaultObjectAclArgs struct {
	Bucket                    pulumi.StringInput
	RoleEntities              pulumi.StringArrayInput
	StorageDefaultObjectAclId pulumi.StringPtrInput
}

func (StorageDefaultObjectAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDefaultObjectAclArgs)(nil)).Elem()
}

type StorageDefaultObjectAclInput interface {
	pulumi.Input

	ToStorageDefaultObjectAclOutput() StorageDefaultObjectAclOutput
	ToStorageDefaultObjectAclOutputWithContext(ctx context.Context) StorageDefaultObjectAclOutput
}

func (*StorageDefaultObjectAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageDefaultObjectAcl)(nil)).Elem()
}

func (i *StorageDefaultObjectAcl) ToStorageDefaultObjectAclOutput() StorageDefaultObjectAclOutput {
	return i.ToStorageDefaultObjectAclOutputWithContext(context.Background())
}

func (i *StorageDefaultObjectAcl) ToStorageDefaultObjectAclOutputWithContext(ctx context.Context) StorageDefaultObjectAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageDefaultObjectAclOutput)
}

type StorageDefaultObjectAclOutput struct{ *pulumi.OutputState }

func (StorageDefaultObjectAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageDefaultObjectAcl)(nil)).Elem()
}

func (o StorageDefaultObjectAclOutput) ToStorageDefaultObjectAclOutput() StorageDefaultObjectAclOutput {
	return o
}

func (o StorageDefaultObjectAclOutput) ToStorageDefaultObjectAclOutputWithContext(ctx context.Context) StorageDefaultObjectAclOutput {
	return o
}

func (o StorageDefaultObjectAclOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAcl) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o StorageDefaultObjectAclOutput) RoleEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAcl) pulumi.StringArrayOutput { return v.RoleEntities }).(pulumi.StringArrayOutput)
}

func (o StorageDefaultObjectAclOutput) StorageDefaultObjectAclId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAcl) pulumi.StringOutput { return v.StorageDefaultObjectAclId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageDefaultObjectAclInput)(nil)).Elem(), &StorageDefaultObjectAcl{})
	pulumi.RegisterOutputType(StorageDefaultObjectAclOutput{})
}
