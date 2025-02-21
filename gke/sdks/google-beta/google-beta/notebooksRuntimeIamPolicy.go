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

type NotebooksRuntimeIamPolicy struct {
	pulumi.CustomResourceState

	Etag                        pulumi.StringOutput `pulumi:"etag"`
	Location                    pulumi.StringOutput `pulumi:"location"`
	NotebooksRuntimeIamPolicyId pulumi.StringOutput `pulumi:"notebooksRuntimeIamPolicyId"`
	PolicyData                  pulumi.StringOutput `pulumi:"policyData"`
	Project                     pulumi.StringOutput `pulumi:"project"`
	RuntimeName                 pulumi.StringOutput `pulumi:"runtimeName"`
}

// NewNotebooksRuntimeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewNotebooksRuntimeIamPolicy(ctx *pulumi.Context,
	name string, args *NotebooksRuntimeIamPolicyArgs, opts ...pulumi.ResourceOption) (*NotebooksRuntimeIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.RuntimeName == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NotebooksRuntimeIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/notebooksRuntimeIamPolicy:NotebooksRuntimeIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebooksRuntimeIamPolicy gets an existing NotebooksRuntimeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebooksRuntimeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebooksRuntimeIamPolicyState, opts ...pulumi.ResourceOption) (*NotebooksRuntimeIamPolicy, error) {
	var resource NotebooksRuntimeIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/notebooksRuntimeIamPolicy:NotebooksRuntimeIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebooksRuntimeIamPolicy resources.
type notebooksRuntimeIamPolicyState struct {
	Etag                        *string `pulumi:"etag"`
	Location                    *string `pulumi:"location"`
	NotebooksRuntimeIamPolicyId *string `pulumi:"notebooksRuntimeIamPolicyId"`
	PolicyData                  *string `pulumi:"policyData"`
	Project                     *string `pulumi:"project"`
	RuntimeName                 *string `pulumi:"runtimeName"`
}

type NotebooksRuntimeIamPolicyState struct {
	Etag                        pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	NotebooksRuntimeIamPolicyId pulumi.StringPtrInput
	PolicyData                  pulumi.StringPtrInput
	Project                     pulumi.StringPtrInput
	RuntimeName                 pulumi.StringPtrInput
}

func (NotebooksRuntimeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksRuntimeIamPolicyState)(nil)).Elem()
}

type notebooksRuntimeIamPolicyArgs struct {
	Location                    *string `pulumi:"location"`
	NotebooksRuntimeIamPolicyId *string `pulumi:"notebooksRuntimeIamPolicyId"`
	PolicyData                  string  `pulumi:"policyData"`
	Project                     *string `pulumi:"project"`
	RuntimeName                 string  `pulumi:"runtimeName"`
}

// The set of arguments for constructing a NotebooksRuntimeIamPolicy resource.
type NotebooksRuntimeIamPolicyArgs struct {
	Location                    pulumi.StringPtrInput
	NotebooksRuntimeIamPolicyId pulumi.StringPtrInput
	PolicyData                  pulumi.StringInput
	Project                     pulumi.StringPtrInput
	RuntimeName                 pulumi.StringInput
}

func (NotebooksRuntimeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksRuntimeIamPolicyArgs)(nil)).Elem()
}

type NotebooksRuntimeIamPolicyInput interface {
	pulumi.Input

	ToNotebooksRuntimeIamPolicyOutput() NotebooksRuntimeIamPolicyOutput
	ToNotebooksRuntimeIamPolicyOutputWithContext(ctx context.Context) NotebooksRuntimeIamPolicyOutput
}

func (*NotebooksRuntimeIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksRuntimeIamPolicy)(nil)).Elem()
}

func (i *NotebooksRuntimeIamPolicy) ToNotebooksRuntimeIamPolicyOutput() NotebooksRuntimeIamPolicyOutput {
	return i.ToNotebooksRuntimeIamPolicyOutputWithContext(context.Background())
}

func (i *NotebooksRuntimeIamPolicy) ToNotebooksRuntimeIamPolicyOutputWithContext(ctx context.Context) NotebooksRuntimeIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebooksRuntimeIamPolicyOutput)
}

type NotebooksRuntimeIamPolicyOutput struct{ *pulumi.OutputState }

func (NotebooksRuntimeIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksRuntimeIamPolicy)(nil)).Elem()
}

func (o NotebooksRuntimeIamPolicyOutput) ToNotebooksRuntimeIamPolicyOutput() NotebooksRuntimeIamPolicyOutput {
	return o
}

func (o NotebooksRuntimeIamPolicyOutput) ToNotebooksRuntimeIamPolicyOutputWithContext(ctx context.Context) NotebooksRuntimeIamPolicyOutput {
	return o
}

func (o NotebooksRuntimeIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksRuntimeIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NotebooksRuntimeIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksRuntimeIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NotebooksRuntimeIamPolicyOutput) NotebooksRuntimeIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksRuntimeIamPolicy) pulumi.StringOutput { return v.NotebooksRuntimeIamPolicyId }).(pulumi.StringOutput)
}

func (o NotebooksRuntimeIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksRuntimeIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o NotebooksRuntimeIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksRuntimeIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NotebooksRuntimeIamPolicyOutput) RuntimeName() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksRuntimeIamPolicy) pulumi.StringOutput { return v.RuntimeName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotebooksRuntimeIamPolicyInput)(nil)).Elem(), &NotebooksRuntimeIamPolicy{})
	pulumi.RegisterOutputType(NotebooksRuntimeIamPolicyOutput{})
}
