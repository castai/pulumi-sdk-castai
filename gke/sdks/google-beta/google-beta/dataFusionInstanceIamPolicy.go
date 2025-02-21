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

type DataFusionInstanceIamPolicy struct {
	pulumi.CustomResourceState

	DataFusionInstanceIamPolicyId pulumi.StringOutput `pulumi:"dataFusionInstanceIamPolicyId"`
	Etag                          pulumi.StringOutput `pulumi:"etag"`
	Name                          pulumi.StringOutput `pulumi:"name"`
	PolicyData                    pulumi.StringOutput `pulumi:"policyData"`
	Project                       pulumi.StringOutput `pulumi:"project"`
	Region                        pulumi.StringOutput `pulumi:"region"`
}

// NewDataFusionInstanceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataFusionInstanceIamPolicy(ctx *pulumi.Context,
	name string, args *DataFusionInstanceIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataFusionInstanceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataFusionInstanceIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/dataFusionInstanceIamPolicy:DataFusionInstanceIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataFusionInstanceIamPolicy gets an existing DataFusionInstanceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataFusionInstanceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataFusionInstanceIamPolicyState, opts ...pulumi.ResourceOption) (*DataFusionInstanceIamPolicy, error) {
	var resource DataFusionInstanceIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataFusionInstanceIamPolicy:DataFusionInstanceIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataFusionInstanceIamPolicy resources.
type dataFusionInstanceIamPolicyState struct {
	DataFusionInstanceIamPolicyId *string `pulumi:"dataFusionInstanceIamPolicyId"`
	Etag                          *string `pulumi:"etag"`
	Name                          *string `pulumi:"name"`
	PolicyData                    *string `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
	Region                        *string `pulumi:"region"`
}

type DataFusionInstanceIamPolicyState struct {
	DataFusionInstanceIamPolicyId pulumi.StringPtrInput
	Etag                          pulumi.StringPtrInput
	Name                          pulumi.StringPtrInput
	PolicyData                    pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Region                        pulumi.StringPtrInput
}

func (DataFusionInstanceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFusionInstanceIamPolicyState)(nil)).Elem()
}

type dataFusionInstanceIamPolicyArgs struct {
	DataFusionInstanceIamPolicyId *string `pulumi:"dataFusionInstanceIamPolicyId"`
	Name                          *string `pulumi:"name"`
	PolicyData                    string  `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
	Region                        *string `pulumi:"region"`
}

// The set of arguments for constructing a DataFusionInstanceIamPolicy resource.
type DataFusionInstanceIamPolicyArgs struct {
	DataFusionInstanceIamPolicyId pulumi.StringPtrInput
	Name                          pulumi.StringPtrInput
	PolicyData                    pulumi.StringInput
	Project                       pulumi.StringPtrInput
	Region                        pulumi.StringPtrInput
}

func (DataFusionInstanceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFusionInstanceIamPolicyArgs)(nil)).Elem()
}

type DataFusionInstanceIamPolicyInput interface {
	pulumi.Input

	ToDataFusionInstanceIamPolicyOutput() DataFusionInstanceIamPolicyOutput
	ToDataFusionInstanceIamPolicyOutputWithContext(ctx context.Context) DataFusionInstanceIamPolicyOutput
}

func (*DataFusionInstanceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFusionInstanceIamPolicy)(nil)).Elem()
}

func (i *DataFusionInstanceIamPolicy) ToDataFusionInstanceIamPolicyOutput() DataFusionInstanceIamPolicyOutput {
	return i.ToDataFusionInstanceIamPolicyOutputWithContext(context.Background())
}

func (i *DataFusionInstanceIamPolicy) ToDataFusionInstanceIamPolicyOutputWithContext(ctx context.Context) DataFusionInstanceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFusionInstanceIamPolicyOutput)
}

type DataFusionInstanceIamPolicyOutput struct{ *pulumi.OutputState }

func (DataFusionInstanceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFusionInstanceIamPolicy)(nil)).Elem()
}

func (o DataFusionInstanceIamPolicyOutput) ToDataFusionInstanceIamPolicyOutput() DataFusionInstanceIamPolicyOutput {
	return o
}

func (o DataFusionInstanceIamPolicyOutput) ToDataFusionInstanceIamPolicyOutputWithContext(ctx context.Context) DataFusionInstanceIamPolicyOutput {
	return o
}

func (o DataFusionInstanceIamPolicyOutput) DataFusionInstanceIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFusionInstanceIamPolicy) pulumi.StringOutput { return v.DataFusionInstanceIamPolicyId }).(pulumi.StringOutput)
}

func (o DataFusionInstanceIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFusionInstanceIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataFusionInstanceIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFusionInstanceIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataFusionInstanceIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFusionInstanceIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o DataFusionInstanceIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFusionInstanceIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataFusionInstanceIamPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFusionInstanceIamPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataFusionInstanceIamPolicyInput)(nil)).Elem(), &DataFusionInstanceIamPolicy{})
	pulumi.RegisterOutputType(DataFusionInstanceIamPolicyOutput{})
}
