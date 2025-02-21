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

type ComputeImageIamBinding struct {
	pulumi.CustomResourceState

	ComputeImageIamBindingId pulumi.StringOutput                      `pulumi:"computeImageIamBindingId"`
	Condition                ComputeImageIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                     pulumi.StringOutput                      `pulumi:"etag"`
	Image                    pulumi.StringOutput                      `pulumi:"image"`
	Members                  pulumi.StringArrayOutput                 `pulumi:"members"`
	Project                  pulumi.StringOutput                      `pulumi:"project"`
	Role                     pulumi.StringOutput                      `pulumi:"role"`
}

// NewComputeImageIamBinding registers a new resource with the given unique name, arguments, and options.
func NewComputeImageIamBinding(ctx *pulumi.Context,
	name string, args *ComputeImageIamBindingArgs, opts ...pulumi.ResourceOption) (*ComputeImageIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
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
	var resource ComputeImageIamBinding
	err = ctx.RegisterPackageResource("google:index/computeImageIamBinding:ComputeImageIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeImageIamBinding gets an existing ComputeImageIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeImageIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeImageIamBindingState, opts ...pulumi.ResourceOption) (*ComputeImageIamBinding, error) {
	var resource ComputeImageIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeImageIamBinding:ComputeImageIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeImageIamBinding resources.
type computeImageIamBindingState struct {
	ComputeImageIamBindingId *string                          `pulumi:"computeImageIamBindingId"`
	Condition                *ComputeImageIamBindingCondition `pulumi:"condition"`
	Etag                     *string                          `pulumi:"etag"`
	Image                    *string                          `pulumi:"image"`
	Members                  []string                         `pulumi:"members"`
	Project                  *string                          `pulumi:"project"`
	Role                     *string                          `pulumi:"role"`
}

type ComputeImageIamBindingState struct {
	ComputeImageIamBindingId pulumi.StringPtrInput
	Condition                ComputeImageIamBindingConditionPtrInput
	Etag                     pulumi.StringPtrInput
	Image                    pulumi.StringPtrInput
	Members                  pulumi.StringArrayInput
	Project                  pulumi.StringPtrInput
	Role                     pulumi.StringPtrInput
}

func (ComputeImageIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeImageIamBindingState)(nil)).Elem()
}

type computeImageIamBindingArgs struct {
	ComputeImageIamBindingId *string                          `pulumi:"computeImageIamBindingId"`
	Condition                *ComputeImageIamBindingCondition `pulumi:"condition"`
	Image                    string                           `pulumi:"image"`
	Members                  []string                         `pulumi:"members"`
	Project                  *string                          `pulumi:"project"`
	Role                     string                           `pulumi:"role"`
}

// The set of arguments for constructing a ComputeImageIamBinding resource.
type ComputeImageIamBindingArgs struct {
	ComputeImageIamBindingId pulumi.StringPtrInput
	Condition                ComputeImageIamBindingConditionPtrInput
	Image                    pulumi.StringInput
	Members                  pulumi.StringArrayInput
	Project                  pulumi.StringPtrInput
	Role                     pulumi.StringInput
}

func (ComputeImageIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeImageIamBindingArgs)(nil)).Elem()
}

type ComputeImageIamBindingInput interface {
	pulumi.Input

	ToComputeImageIamBindingOutput() ComputeImageIamBindingOutput
	ToComputeImageIamBindingOutputWithContext(ctx context.Context) ComputeImageIamBindingOutput
}

func (*ComputeImageIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeImageIamBinding)(nil)).Elem()
}

func (i *ComputeImageIamBinding) ToComputeImageIamBindingOutput() ComputeImageIamBindingOutput {
	return i.ToComputeImageIamBindingOutputWithContext(context.Background())
}

func (i *ComputeImageIamBinding) ToComputeImageIamBindingOutputWithContext(ctx context.Context) ComputeImageIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeImageIamBindingOutput)
}

type ComputeImageIamBindingOutput struct{ *pulumi.OutputState }

func (ComputeImageIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeImageIamBinding)(nil)).Elem()
}

func (o ComputeImageIamBindingOutput) ToComputeImageIamBindingOutput() ComputeImageIamBindingOutput {
	return o
}

func (o ComputeImageIamBindingOutput) ToComputeImageIamBindingOutputWithContext(ctx context.Context) ComputeImageIamBindingOutput {
	return o
}

func (o ComputeImageIamBindingOutput) ComputeImageIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImageIamBinding) pulumi.StringOutput { return v.ComputeImageIamBindingId }).(pulumi.StringOutput)
}

func (o ComputeImageIamBindingOutput) Condition() ComputeImageIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ComputeImageIamBinding) ComputeImageIamBindingConditionPtrOutput { return v.Condition }).(ComputeImageIamBindingConditionPtrOutput)
}

func (o ComputeImageIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImageIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ComputeImageIamBindingOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImageIamBinding) pulumi.StringOutput { return v.Image }).(pulumi.StringOutput)
}

func (o ComputeImageIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeImageIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o ComputeImageIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImageIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeImageIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImageIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeImageIamBindingInput)(nil)).Elem(), &ComputeImageIamBinding{})
	pulumi.RegisterOutputType(ComputeImageIamBindingOutput{})
}
