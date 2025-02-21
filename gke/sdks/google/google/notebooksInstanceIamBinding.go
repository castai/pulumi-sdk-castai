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

type NotebooksInstanceIamBinding struct {
	pulumi.CustomResourceState

	Condition                     NotebooksInstanceIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                          pulumi.StringOutput                           `pulumi:"etag"`
	InstanceName                  pulumi.StringOutput                           `pulumi:"instanceName"`
	Location                      pulumi.StringOutput                           `pulumi:"location"`
	Members                       pulumi.StringArrayOutput                      `pulumi:"members"`
	NotebooksInstanceIamBindingId pulumi.StringOutput                           `pulumi:"notebooksInstanceIamBindingId"`
	Project                       pulumi.StringOutput                           `pulumi:"project"`
	Role                          pulumi.StringOutput                           `pulumi:"role"`
}

// NewNotebooksInstanceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewNotebooksInstanceIamBinding(ctx *pulumi.Context,
	name string, args *NotebooksInstanceIamBindingArgs, opts ...pulumi.ResourceOption) (*NotebooksInstanceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
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
	var resource NotebooksInstanceIamBinding
	err = ctx.RegisterPackageResource("google:index/notebooksInstanceIamBinding:NotebooksInstanceIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebooksInstanceIamBinding gets an existing NotebooksInstanceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebooksInstanceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebooksInstanceIamBindingState, opts ...pulumi.ResourceOption) (*NotebooksInstanceIamBinding, error) {
	var resource NotebooksInstanceIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/notebooksInstanceIamBinding:NotebooksInstanceIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebooksInstanceIamBinding resources.
type notebooksInstanceIamBindingState struct {
	Condition                     *NotebooksInstanceIamBindingCondition `pulumi:"condition"`
	Etag                          *string                               `pulumi:"etag"`
	InstanceName                  *string                               `pulumi:"instanceName"`
	Location                      *string                               `pulumi:"location"`
	Members                       []string                              `pulumi:"members"`
	NotebooksInstanceIamBindingId *string                               `pulumi:"notebooksInstanceIamBindingId"`
	Project                       *string                               `pulumi:"project"`
	Role                          *string                               `pulumi:"role"`
}

type NotebooksInstanceIamBindingState struct {
	Condition                     NotebooksInstanceIamBindingConditionPtrInput
	Etag                          pulumi.StringPtrInput
	InstanceName                  pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	NotebooksInstanceIamBindingId pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Role                          pulumi.StringPtrInput
}

func (NotebooksInstanceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksInstanceIamBindingState)(nil)).Elem()
}

type notebooksInstanceIamBindingArgs struct {
	Condition                     *NotebooksInstanceIamBindingCondition `pulumi:"condition"`
	InstanceName                  string                                `pulumi:"instanceName"`
	Location                      *string                               `pulumi:"location"`
	Members                       []string                              `pulumi:"members"`
	NotebooksInstanceIamBindingId *string                               `pulumi:"notebooksInstanceIamBindingId"`
	Project                       *string                               `pulumi:"project"`
	Role                          string                                `pulumi:"role"`
}

// The set of arguments for constructing a NotebooksInstanceIamBinding resource.
type NotebooksInstanceIamBindingArgs struct {
	Condition                     NotebooksInstanceIamBindingConditionPtrInput
	InstanceName                  pulumi.StringInput
	Location                      pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	NotebooksInstanceIamBindingId pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Role                          pulumi.StringInput
}

func (NotebooksInstanceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksInstanceIamBindingArgs)(nil)).Elem()
}

type NotebooksInstanceIamBindingInput interface {
	pulumi.Input

	ToNotebooksInstanceIamBindingOutput() NotebooksInstanceIamBindingOutput
	ToNotebooksInstanceIamBindingOutputWithContext(ctx context.Context) NotebooksInstanceIamBindingOutput
}

func (*NotebooksInstanceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksInstanceIamBinding)(nil)).Elem()
}

func (i *NotebooksInstanceIamBinding) ToNotebooksInstanceIamBindingOutput() NotebooksInstanceIamBindingOutput {
	return i.ToNotebooksInstanceIamBindingOutputWithContext(context.Background())
}

func (i *NotebooksInstanceIamBinding) ToNotebooksInstanceIamBindingOutputWithContext(ctx context.Context) NotebooksInstanceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebooksInstanceIamBindingOutput)
}

type NotebooksInstanceIamBindingOutput struct{ *pulumi.OutputState }

func (NotebooksInstanceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksInstanceIamBinding)(nil)).Elem()
}

func (o NotebooksInstanceIamBindingOutput) ToNotebooksInstanceIamBindingOutput() NotebooksInstanceIamBindingOutput {
	return o
}

func (o NotebooksInstanceIamBindingOutput) ToNotebooksInstanceIamBindingOutputWithContext(ctx context.Context) NotebooksInstanceIamBindingOutput {
	return o
}

func (o NotebooksInstanceIamBindingOutput) Condition() NotebooksInstanceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) NotebooksInstanceIamBindingConditionPtrOutput { return v.Condition }).(NotebooksInstanceIamBindingConditionPtrOutput)
}

func (o NotebooksInstanceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamBindingOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o NotebooksInstanceIamBindingOutput) NotebooksInstanceIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) pulumi.StringOutput { return v.NotebooksInstanceIamBindingId }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotebooksInstanceIamBindingInput)(nil)).Elem(), &NotebooksInstanceIamBinding{})
	pulumi.RegisterOutputType(NotebooksInstanceIamBindingOutput{})
}
