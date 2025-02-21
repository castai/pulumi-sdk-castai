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

type DataplexAspectTypeIamPolicy struct {
	pulumi.CustomResourceState

	AspectTypeId                  pulumi.StringOutput `pulumi:"aspectTypeId"`
	DataplexAspectTypeIamPolicyId pulumi.StringOutput `pulumi:"dataplexAspectTypeIamPolicyId"`
	Etag                          pulumi.StringOutput `pulumi:"etag"`
	Location                      pulumi.StringOutput `pulumi:"location"`
	PolicyData                    pulumi.StringOutput `pulumi:"policyData"`
	Project                       pulumi.StringOutput `pulumi:"project"`
}

// NewDataplexAspectTypeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataplexAspectTypeIamPolicy(ctx *pulumi.Context,
	name string, args *DataplexAspectTypeIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataplexAspectTypeIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AspectTypeId == nil {
		return nil, errors.New("invalid value for required argument 'AspectTypeId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexAspectTypeIamPolicy
	err = ctx.RegisterPackageResource("google:index/dataplexAspectTypeIamPolicy:DataplexAspectTypeIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexAspectTypeIamPolicy gets an existing DataplexAspectTypeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexAspectTypeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexAspectTypeIamPolicyState, opts ...pulumi.ResourceOption) (*DataplexAspectTypeIamPolicy, error) {
	var resource DataplexAspectTypeIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataplexAspectTypeIamPolicy:DataplexAspectTypeIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexAspectTypeIamPolicy resources.
type dataplexAspectTypeIamPolicyState struct {
	AspectTypeId                  *string `pulumi:"aspectTypeId"`
	DataplexAspectTypeIamPolicyId *string `pulumi:"dataplexAspectTypeIamPolicyId"`
	Etag                          *string `pulumi:"etag"`
	Location                      *string `pulumi:"location"`
	PolicyData                    *string `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
}

type DataplexAspectTypeIamPolicyState struct {
	AspectTypeId                  pulumi.StringPtrInput
	DataplexAspectTypeIamPolicyId pulumi.StringPtrInput
	Etag                          pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	PolicyData                    pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
}

func (DataplexAspectTypeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexAspectTypeIamPolicyState)(nil)).Elem()
}

type dataplexAspectTypeIamPolicyArgs struct {
	AspectTypeId                  string  `pulumi:"aspectTypeId"`
	DataplexAspectTypeIamPolicyId *string `pulumi:"dataplexAspectTypeIamPolicyId"`
	Location                      *string `pulumi:"location"`
	PolicyData                    string  `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
}

// The set of arguments for constructing a DataplexAspectTypeIamPolicy resource.
type DataplexAspectTypeIamPolicyArgs struct {
	AspectTypeId                  pulumi.StringInput
	DataplexAspectTypeIamPolicyId pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	PolicyData                    pulumi.StringInput
	Project                       pulumi.StringPtrInput
}

func (DataplexAspectTypeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexAspectTypeIamPolicyArgs)(nil)).Elem()
}

type DataplexAspectTypeIamPolicyInput interface {
	pulumi.Input

	ToDataplexAspectTypeIamPolicyOutput() DataplexAspectTypeIamPolicyOutput
	ToDataplexAspectTypeIamPolicyOutputWithContext(ctx context.Context) DataplexAspectTypeIamPolicyOutput
}

func (*DataplexAspectTypeIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexAspectTypeIamPolicy)(nil)).Elem()
}

func (i *DataplexAspectTypeIamPolicy) ToDataplexAspectTypeIamPolicyOutput() DataplexAspectTypeIamPolicyOutput {
	return i.ToDataplexAspectTypeIamPolicyOutputWithContext(context.Background())
}

func (i *DataplexAspectTypeIamPolicy) ToDataplexAspectTypeIamPolicyOutputWithContext(ctx context.Context) DataplexAspectTypeIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexAspectTypeIamPolicyOutput)
}

type DataplexAspectTypeIamPolicyOutput struct{ *pulumi.OutputState }

func (DataplexAspectTypeIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexAspectTypeIamPolicy)(nil)).Elem()
}

func (o DataplexAspectTypeIamPolicyOutput) ToDataplexAspectTypeIamPolicyOutput() DataplexAspectTypeIamPolicyOutput {
	return o
}

func (o DataplexAspectTypeIamPolicyOutput) ToDataplexAspectTypeIamPolicyOutputWithContext(ctx context.Context) DataplexAspectTypeIamPolicyOutput {
	return o
}

func (o DataplexAspectTypeIamPolicyOutput) AspectTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectTypeIamPolicy) pulumi.StringOutput { return v.AspectTypeId }).(pulumi.StringOutput)
}

func (o DataplexAspectTypeIamPolicyOutput) DataplexAspectTypeIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectTypeIamPolicy) pulumi.StringOutput { return v.DataplexAspectTypeIamPolicyId }).(pulumi.StringOutput)
}

func (o DataplexAspectTypeIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectTypeIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataplexAspectTypeIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectTypeIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataplexAspectTypeIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectTypeIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o DataplexAspectTypeIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAspectTypeIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexAspectTypeIamPolicyInput)(nil)).Elem(), &DataplexAspectTypeIamPolicy{})
	pulumi.RegisterOutputType(DataplexAspectTypeIamPolicyOutput{})
}
