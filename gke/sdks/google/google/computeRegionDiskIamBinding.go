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

type ComputeRegionDiskIamBinding struct {
	pulumi.CustomResourceState

	ComputeRegionDiskIamBindingId pulumi.StringOutput                           `pulumi:"computeRegionDiskIamBindingId"`
	Condition                     ComputeRegionDiskIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                          pulumi.StringOutput                           `pulumi:"etag"`
	Members                       pulumi.StringArrayOutput                      `pulumi:"members"`
	Name                          pulumi.StringOutput                           `pulumi:"name"`
	Project                       pulumi.StringOutput                           `pulumi:"project"`
	Region                        pulumi.StringOutput                           `pulumi:"region"`
	Role                          pulumi.StringOutput                           `pulumi:"role"`
}

// NewComputeRegionDiskIamBinding registers a new resource with the given unique name, arguments, and options.
func NewComputeRegionDiskIamBinding(ctx *pulumi.Context,
	name string, args *ComputeRegionDiskIamBindingArgs, opts ...pulumi.ResourceOption) (*ComputeRegionDiskIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource ComputeRegionDiskIamBinding
	err = ctx.RegisterPackageResource("google:index/computeRegionDiskIamBinding:ComputeRegionDiskIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRegionDiskIamBinding gets an existing ComputeRegionDiskIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRegionDiskIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRegionDiskIamBindingState, opts ...pulumi.ResourceOption) (*ComputeRegionDiskIamBinding, error) {
	var resource ComputeRegionDiskIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeRegionDiskIamBinding:ComputeRegionDiskIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRegionDiskIamBinding resources.
type computeRegionDiskIamBindingState struct {
	ComputeRegionDiskIamBindingId *string                               `pulumi:"computeRegionDiskIamBindingId"`
	Condition                     *ComputeRegionDiskIamBindingCondition `pulumi:"condition"`
	Etag                          *string                               `pulumi:"etag"`
	Members                       []string                              `pulumi:"members"`
	Name                          *string                               `pulumi:"name"`
	Project                       *string                               `pulumi:"project"`
	Region                        *string                               `pulumi:"region"`
	Role                          *string                               `pulumi:"role"`
}

type ComputeRegionDiskIamBindingState struct {
	ComputeRegionDiskIamBindingId pulumi.StringPtrInput
	Condition                     ComputeRegionDiskIamBindingConditionPtrInput
	Etag                          pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	Name                          pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Region                        pulumi.StringPtrInput
	Role                          pulumi.StringPtrInput
}

func (ComputeRegionDiskIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionDiskIamBindingState)(nil)).Elem()
}

type computeRegionDiskIamBindingArgs struct {
	ComputeRegionDiskIamBindingId *string                               `pulumi:"computeRegionDiskIamBindingId"`
	Condition                     *ComputeRegionDiskIamBindingCondition `pulumi:"condition"`
	Members                       []string                              `pulumi:"members"`
	Name                          *string                               `pulumi:"name"`
	Project                       *string                               `pulumi:"project"`
	Region                        *string                               `pulumi:"region"`
	Role                          string                                `pulumi:"role"`
}

// The set of arguments for constructing a ComputeRegionDiskIamBinding resource.
type ComputeRegionDiskIamBindingArgs struct {
	ComputeRegionDiskIamBindingId pulumi.StringPtrInput
	Condition                     ComputeRegionDiskIamBindingConditionPtrInput
	Members                       pulumi.StringArrayInput
	Name                          pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Region                        pulumi.StringPtrInput
	Role                          pulumi.StringInput
}

func (ComputeRegionDiskIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionDiskIamBindingArgs)(nil)).Elem()
}

type ComputeRegionDiskIamBindingInput interface {
	pulumi.Input

	ToComputeRegionDiskIamBindingOutput() ComputeRegionDiskIamBindingOutput
	ToComputeRegionDiskIamBindingOutputWithContext(ctx context.Context) ComputeRegionDiskIamBindingOutput
}

func (*ComputeRegionDiskIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionDiskIamBinding)(nil)).Elem()
}

func (i *ComputeRegionDiskIamBinding) ToComputeRegionDiskIamBindingOutput() ComputeRegionDiskIamBindingOutput {
	return i.ToComputeRegionDiskIamBindingOutputWithContext(context.Background())
}

func (i *ComputeRegionDiskIamBinding) ToComputeRegionDiskIamBindingOutputWithContext(ctx context.Context) ComputeRegionDiskIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRegionDiskIamBindingOutput)
}

type ComputeRegionDiskIamBindingOutput struct{ *pulumi.OutputState }

func (ComputeRegionDiskIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionDiskIamBinding)(nil)).Elem()
}

func (o ComputeRegionDiskIamBindingOutput) ToComputeRegionDiskIamBindingOutput() ComputeRegionDiskIamBindingOutput {
	return o
}

func (o ComputeRegionDiskIamBindingOutput) ToComputeRegionDiskIamBindingOutputWithContext(ctx context.Context) ComputeRegionDiskIamBindingOutput {
	return o
}

func (o ComputeRegionDiskIamBindingOutput) ComputeRegionDiskIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) pulumi.StringOutput { return v.ComputeRegionDiskIamBindingId }).(pulumi.StringOutput)
}

func (o ComputeRegionDiskIamBindingOutput) Condition() ComputeRegionDiskIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) ComputeRegionDiskIamBindingConditionPtrOutput { return v.Condition }).(ComputeRegionDiskIamBindingConditionPtrOutput)
}

func (o ComputeRegionDiskIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ComputeRegionDiskIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o ComputeRegionDiskIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeRegionDiskIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeRegionDiskIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeRegionDiskIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRegionDiskIamBindingInput)(nil)).Elem(), &ComputeRegionDiskIamBinding{})
	pulumi.RegisterOutputType(ComputeRegionDiskIamBindingOutput{})
}
