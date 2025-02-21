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

type DataCatalogPolicyTagIamPolicy struct {
	pulumi.CustomResourceState

	DataCatalogPolicyTagIamPolicyId pulumi.StringOutput `pulumi:"dataCatalogPolicyTagIamPolicyId"`
	Etag                            pulumi.StringOutput `pulumi:"etag"`
	PolicyData                      pulumi.StringOutput `pulumi:"policyData"`
	PolicyTag                       pulumi.StringOutput `pulumi:"policyTag"`
}

// NewDataCatalogPolicyTagIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataCatalogPolicyTagIamPolicy(ctx *pulumi.Context,
	name string, args *DataCatalogPolicyTagIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataCatalogPolicyTagIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.PolicyTag == nil {
		return nil, errors.New("invalid value for required argument 'PolicyTag'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataCatalogPolicyTagIamPolicy
	err = ctx.RegisterPackageResource("google:index/dataCatalogPolicyTagIamPolicy:DataCatalogPolicyTagIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCatalogPolicyTagIamPolicy gets an existing DataCatalogPolicyTagIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCatalogPolicyTagIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCatalogPolicyTagIamPolicyState, opts ...pulumi.ResourceOption) (*DataCatalogPolicyTagIamPolicy, error) {
	var resource DataCatalogPolicyTagIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataCatalogPolicyTagIamPolicy:DataCatalogPolicyTagIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCatalogPolicyTagIamPolicy resources.
type dataCatalogPolicyTagIamPolicyState struct {
	DataCatalogPolicyTagIamPolicyId *string `pulumi:"dataCatalogPolicyTagIamPolicyId"`
	Etag                            *string `pulumi:"etag"`
	PolicyData                      *string `pulumi:"policyData"`
	PolicyTag                       *string `pulumi:"policyTag"`
}

type DataCatalogPolicyTagIamPolicyState struct {
	DataCatalogPolicyTagIamPolicyId pulumi.StringPtrInput
	Etag                            pulumi.StringPtrInput
	PolicyData                      pulumi.StringPtrInput
	PolicyTag                       pulumi.StringPtrInput
}

func (DataCatalogPolicyTagIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogPolicyTagIamPolicyState)(nil)).Elem()
}

type dataCatalogPolicyTagIamPolicyArgs struct {
	DataCatalogPolicyTagIamPolicyId *string `pulumi:"dataCatalogPolicyTagIamPolicyId"`
	PolicyData                      string  `pulumi:"policyData"`
	PolicyTag                       string  `pulumi:"policyTag"`
}

// The set of arguments for constructing a DataCatalogPolicyTagIamPolicy resource.
type DataCatalogPolicyTagIamPolicyArgs struct {
	DataCatalogPolicyTagIamPolicyId pulumi.StringPtrInput
	PolicyData                      pulumi.StringInput
	PolicyTag                       pulumi.StringInput
}

func (DataCatalogPolicyTagIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogPolicyTagIamPolicyArgs)(nil)).Elem()
}

type DataCatalogPolicyTagIamPolicyInput interface {
	pulumi.Input

	ToDataCatalogPolicyTagIamPolicyOutput() DataCatalogPolicyTagIamPolicyOutput
	ToDataCatalogPolicyTagIamPolicyOutputWithContext(ctx context.Context) DataCatalogPolicyTagIamPolicyOutput
}

func (*DataCatalogPolicyTagIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogPolicyTagIamPolicy)(nil)).Elem()
}

func (i *DataCatalogPolicyTagIamPolicy) ToDataCatalogPolicyTagIamPolicyOutput() DataCatalogPolicyTagIamPolicyOutput {
	return i.ToDataCatalogPolicyTagIamPolicyOutputWithContext(context.Background())
}

func (i *DataCatalogPolicyTagIamPolicy) ToDataCatalogPolicyTagIamPolicyOutputWithContext(ctx context.Context) DataCatalogPolicyTagIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCatalogPolicyTagIamPolicyOutput)
}

type DataCatalogPolicyTagIamPolicyOutput struct{ *pulumi.OutputState }

func (DataCatalogPolicyTagIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogPolicyTagIamPolicy)(nil)).Elem()
}

func (o DataCatalogPolicyTagIamPolicyOutput) ToDataCatalogPolicyTagIamPolicyOutput() DataCatalogPolicyTagIamPolicyOutput {
	return o
}

func (o DataCatalogPolicyTagIamPolicyOutput) ToDataCatalogPolicyTagIamPolicyOutputWithContext(ctx context.Context) DataCatalogPolicyTagIamPolicyOutput {
	return o
}

func (o DataCatalogPolicyTagIamPolicyOutput) DataCatalogPolicyTagIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogPolicyTagIamPolicy) pulumi.StringOutput { return v.DataCatalogPolicyTagIamPolicyId }).(pulumi.StringOutput)
}

func (o DataCatalogPolicyTagIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogPolicyTagIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataCatalogPolicyTagIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogPolicyTagIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o DataCatalogPolicyTagIamPolicyOutput) PolicyTag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogPolicyTagIamPolicy) pulumi.StringOutput { return v.PolicyTag }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCatalogPolicyTagIamPolicyInput)(nil)).Elem(), &DataCatalogPolicyTagIamPolicy{})
	pulumi.RegisterOutputType(DataCatalogPolicyTagIamPolicyOutput{})
}
