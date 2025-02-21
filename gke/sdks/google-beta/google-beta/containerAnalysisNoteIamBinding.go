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

type ContainerAnalysisNoteIamBinding struct {
	pulumi.CustomResourceState

	Condition                         ContainerAnalysisNoteIamBindingConditionPtrOutput `pulumi:"condition"`
	ContainerAnalysisNoteIamBindingId pulumi.StringOutput                               `pulumi:"containerAnalysisNoteIamBindingId"`
	Etag                              pulumi.StringOutput                               `pulumi:"etag"`
	Members                           pulumi.StringArrayOutput                          `pulumi:"members"`
	Note                              pulumi.StringOutput                               `pulumi:"note"`
	Project                           pulumi.StringOutput                               `pulumi:"project"`
	Role                              pulumi.StringOutput                               `pulumi:"role"`
}

// NewContainerAnalysisNoteIamBinding registers a new resource with the given unique name, arguments, and options.
func NewContainerAnalysisNoteIamBinding(ctx *pulumi.Context,
	name string, args *ContainerAnalysisNoteIamBindingArgs, opts ...pulumi.ResourceOption) (*ContainerAnalysisNoteIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Note == nil {
		return nil, errors.New("invalid value for required argument 'Note'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ContainerAnalysisNoteIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/containerAnalysisNoteIamBinding:ContainerAnalysisNoteIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerAnalysisNoteIamBinding gets an existing ContainerAnalysisNoteIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerAnalysisNoteIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAnalysisNoteIamBindingState, opts ...pulumi.ResourceOption) (*ContainerAnalysisNoteIamBinding, error) {
	var resource ContainerAnalysisNoteIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/containerAnalysisNoteIamBinding:ContainerAnalysisNoteIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerAnalysisNoteIamBinding resources.
type containerAnalysisNoteIamBindingState struct {
	Condition                         *ContainerAnalysisNoteIamBindingCondition `pulumi:"condition"`
	ContainerAnalysisNoteIamBindingId *string                                   `pulumi:"containerAnalysisNoteIamBindingId"`
	Etag                              *string                                   `pulumi:"etag"`
	Members                           []string                                  `pulumi:"members"`
	Note                              *string                                   `pulumi:"note"`
	Project                           *string                                   `pulumi:"project"`
	Role                              *string                                   `pulumi:"role"`
}

type ContainerAnalysisNoteIamBindingState struct {
	Condition                         ContainerAnalysisNoteIamBindingConditionPtrInput
	ContainerAnalysisNoteIamBindingId pulumi.StringPtrInput
	Etag                              pulumi.StringPtrInput
	Members                           pulumi.StringArrayInput
	Note                              pulumi.StringPtrInput
	Project                           pulumi.StringPtrInput
	Role                              pulumi.StringPtrInput
}

func (ContainerAnalysisNoteIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAnalysisNoteIamBindingState)(nil)).Elem()
}

type containerAnalysisNoteIamBindingArgs struct {
	Condition                         *ContainerAnalysisNoteIamBindingCondition `pulumi:"condition"`
	ContainerAnalysisNoteIamBindingId *string                                   `pulumi:"containerAnalysisNoteIamBindingId"`
	Members                           []string                                  `pulumi:"members"`
	Note                              string                                    `pulumi:"note"`
	Project                           *string                                   `pulumi:"project"`
	Role                              string                                    `pulumi:"role"`
}

// The set of arguments for constructing a ContainerAnalysisNoteIamBinding resource.
type ContainerAnalysisNoteIamBindingArgs struct {
	Condition                         ContainerAnalysisNoteIamBindingConditionPtrInput
	ContainerAnalysisNoteIamBindingId pulumi.StringPtrInput
	Members                           pulumi.StringArrayInput
	Note                              pulumi.StringInput
	Project                           pulumi.StringPtrInput
	Role                              pulumi.StringInput
}

func (ContainerAnalysisNoteIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAnalysisNoteIamBindingArgs)(nil)).Elem()
}

type ContainerAnalysisNoteIamBindingInput interface {
	pulumi.Input

	ToContainerAnalysisNoteIamBindingOutput() ContainerAnalysisNoteIamBindingOutput
	ToContainerAnalysisNoteIamBindingOutputWithContext(ctx context.Context) ContainerAnalysisNoteIamBindingOutput
}

func (*ContainerAnalysisNoteIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAnalysisNoteIamBinding)(nil)).Elem()
}

func (i *ContainerAnalysisNoteIamBinding) ToContainerAnalysisNoteIamBindingOutput() ContainerAnalysisNoteIamBindingOutput {
	return i.ToContainerAnalysisNoteIamBindingOutputWithContext(context.Background())
}

func (i *ContainerAnalysisNoteIamBinding) ToContainerAnalysisNoteIamBindingOutputWithContext(ctx context.Context) ContainerAnalysisNoteIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAnalysisNoteIamBindingOutput)
}

type ContainerAnalysisNoteIamBindingOutput struct{ *pulumi.OutputState }

func (ContainerAnalysisNoteIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAnalysisNoteIamBinding)(nil)).Elem()
}

func (o ContainerAnalysisNoteIamBindingOutput) ToContainerAnalysisNoteIamBindingOutput() ContainerAnalysisNoteIamBindingOutput {
	return o
}

func (o ContainerAnalysisNoteIamBindingOutput) ToContainerAnalysisNoteIamBindingOutputWithContext(ctx context.Context) ContainerAnalysisNoteIamBindingOutput {
	return o
}

func (o ContainerAnalysisNoteIamBindingOutput) Condition() ContainerAnalysisNoteIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ContainerAnalysisNoteIamBinding) ContainerAnalysisNoteIamBindingConditionPtrOutput {
		return v.Condition
	}).(ContainerAnalysisNoteIamBindingConditionPtrOutput)
}

func (o ContainerAnalysisNoteIamBindingOutput) ContainerAnalysisNoteIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAnalysisNoteIamBinding) pulumi.StringOutput {
		return v.ContainerAnalysisNoteIamBindingId
	}).(pulumi.StringOutput)
}

func (o ContainerAnalysisNoteIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAnalysisNoteIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ContainerAnalysisNoteIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerAnalysisNoteIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o ContainerAnalysisNoteIamBindingOutput) Note() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAnalysisNoteIamBinding) pulumi.StringOutput { return v.Note }).(pulumi.StringOutput)
}

func (o ContainerAnalysisNoteIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAnalysisNoteIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ContainerAnalysisNoteIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAnalysisNoteIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerAnalysisNoteIamBindingInput)(nil)).Elem(), &ContainerAnalysisNoteIamBinding{})
	pulumi.RegisterOutputType(ContainerAnalysisNoteIamBindingOutput{})
}
