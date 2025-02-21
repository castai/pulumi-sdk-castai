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

type DataformRepositoryIamPolicy struct {
	pulumi.CustomResourceState

	DataformRepositoryIamPolicyId pulumi.StringOutput `pulumi:"dataformRepositoryIamPolicyId"`
	Etag                          pulumi.StringOutput `pulumi:"etag"`
	PolicyData                    pulumi.StringOutput `pulumi:"policyData"`
	Project                       pulumi.StringOutput `pulumi:"project"`
	Region                        pulumi.StringOutput `pulumi:"region"`
	Repository                    pulumi.StringOutput `pulumi:"repository"`
}

// NewDataformRepositoryIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataformRepositoryIamPolicy(ctx *pulumi.Context,
	name string, args *DataformRepositoryIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataformRepositoryIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataformRepositoryIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/dataformRepositoryIamPolicy:DataformRepositoryIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataformRepositoryIamPolicy gets an existing DataformRepositoryIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataformRepositoryIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataformRepositoryIamPolicyState, opts ...pulumi.ResourceOption) (*DataformRepositoryIamPolicy, error) {
	var resource DataformRepositoryIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataformRepositoryIamPolicy:DataformRepositoryIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataformRepositoryIamPolicy resources.
type dataformRepositoryIamPolicyState struct {
	DataformRepositoryIamPolicyId *string `pulumi:"dataformRepositoryIamPolicyId"`
	Etag                          *string `pulumi:"etag"`
	PolicyData                    *string `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
	Region                        *string `pulumi:"region"`
	Repository                    *string `pulumi:"repository"`
}

type DataformRepositoryIamPolicyState struct {
	DataformRepositoryIamPolicyId pulumi.StringPtrInput
	Etag                          pulumi.StringPtrInput
	PolicyData                    pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Region                        pulumi.StringPtrInput
	Repository                    pulumi.StringPtrInput
}

func (DataformRepositoryIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataformRepositoryIamPolicyState)(nil)).Elem()
}

type dataformRepositoryIamPolicyArgs struct {
	DataformRepositoryIamPolicyId *string `pulumi:"dataformRepositoryIamPolicyId"`
	PolicyData                    string  `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
	Region                        *string `pulumi:"region"`
	Repository                    string  `pulumi:"repository"`
}

// The set of arguments for constructing a DataformRepositoryIamPolicy resource.
type DataformRepositoryIamPolicyArgs struct {
	DataformRepositoryIamPolicyId pulumi.StringPtrInput
	PolicyData                    pulumi.StringInput
	Project                       pulumi.StringPtrInput
	Region                        pulumi.StringPtrInput
	Repository                    pulumi.StringInput
}

func (DataformRepositoryIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataformRepositoryIamPolicyArgs)(nil)).Elem()
}

type DataformRepositoryIamPolicyInput interface {
	pulumi.Input

	ToDataformRepositoryIamPolicyOutput() DataformRepositoryIamPolicyOutput
	ToDataformRepositoryIamPolicyOutputWithContext(ctx context.Context) DataformRepositoryIamPolicyOutput
}

func (*DataformRepositoryIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataformRepositoryIamPolicy)(nil)).Elem()
}

func (i *DataformRepositoryIamPolicy) ToDataformRepositoryIamPolicyOutput() DataformRepositoryIamPolicyOutput {
	return i.ToDataformRepositoryIamPolicyOutputWithContext(context.Background())
}

func (i *DataformRepositoryIamPolicy) ToDataformRepositoryIamPolicyOutputWithContext(ctx context.Context) DataformRepositoryIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataformRepositoryIamPolicyOutput)
}

type DataformRepositoryIamPolicyOutput struct{ *pulumi.OutputState }

func (DataformRepositoryIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataformRepositoryIamPolicy)(nil)).Elem()
}

func (o DataformRepositoryIamPolicyOutput) ToDataformRepositoryIamPolicyOutput() DataformRepositoryIamPolicyOutput {
	return o
}

func (o DataformRepositoryIamPolicyOutput) ToDataformRepositoryIamPolicyOutputWithContext(ctx context.Context) DataformRepositoryIamPolicyOutput {
	return o
}

func (o DataformRepositoryIamPolicyOutput) DataformRepositoryIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataformRepositoryIamPolicy) pulumi.StringOutput { return v.DataformRepositoryIamPolicyId }).(pulumi.StringOutput)
}

func (o DataformRepositoryIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataformRepositoryIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataformRepositoryIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataformRepositoryIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o DataformRepositoryIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataformRepositoryIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataformRepositoryIamPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DataformRepositoryIamPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o DataformRepositoryIamPolicyOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *DataformRepositoryIamPolicy) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataformRepositoryIamPolicyInput)(nil)).Elem(), &DataformRepositoryIamPolicy{})
	pulumi.RegisterOutputType(DataformRepositoryIamPolicyOutput{})
}
