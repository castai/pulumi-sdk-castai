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

type DataplexEntryGroupIamPolicy struct {
	pulumi.CustomResourceState

	DataplexEntryGroupIamPolicyId pulumi.StringOutput `pulumi:"dataplexEntryGroupIamPolicyId"`
	EntryGroupId                  pulumi.StringOutput `pulumi:"entryGroupId"`
	Etag                          pulumi.StringOutput `pulumi:"etag"`
	Location                      pulumi.StringOutput `pulumi:"location"`
	PolicyData                    pulumi.StringOutput `pulumi:"policyData"`
	Project                       pulumi.StringOutput `pulumi:"project"`
}

// NewDataplexEntryGroupIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataplexEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, args *DataplexEntryGroupIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataplexEntryGroupIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EntryGroupId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexEntryGroupIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/dataplexEntryGroupIamPolicy:DataplexEntryGroupIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexEntryGroupIamPolicy gets an existing DataplexEntryGroupIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexEntryGroupIamPolicyState, opts ...pulumi.ResourceOption) (*DataplexEntryGroupIamPolicy, error) {
	var resource DataplexEntryGroupIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataplexEntryGroupIamPolicy:DataplexEntryGroupIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexEntryGroupIamPolicy resources.
type dataplexEntryGroupIamPolicyState struct {
	DataplexEntryGroupIamPolicyId *string `pulumi:"dataplexEntryGroupIamPolicyId"`
	EntryGroupId                  *string `pulumi:"entryGroupId"`
	Etag                          *string `pulumi:"etag"`
	Location                      *string `pulumi:"location"`
	PolicyData                    *string `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
}

type DataplexEntryGroupIamPolicyState struct {
	DataplexEntryGroupIamPolicyId pulumi.StringPtrInput
	EntryGroupId                  pulumi.StringPtrInput
	Etag                          pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	PolicyData                    pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
}

func (DataplexEntryGroupIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexEntryGroupIamPolicyState)(nil)).Elem()
}

type dataplexEntryGroupIamPolicyArgs struct {
	DataplexEntryGroupIamPolicyId *string `pulumi:"dataplexEntryGroupIamPolicyId"`
	EntryGroupId                  string  `pulumi:"entryGroupId"`
	Location                      *string `pulumi:"location"`
	PolicyData                    string  `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
}

// The set of arguments for constructing a DataplexEntryGroupIamPolicy resource.
type DataplexEntryGroupIamPolicyArgs struct {
	DataplexEntryGroupIamPolicyId pulumi.StringPtrInput
	EntryGroupId                  pulumi.StringInput
	Location                      pulumi.StringPtrInput
	PolicyData                    pulumi.StringInput
	Project                       pulumi.StringPtrInput
}

func (DataplexEntryGroupIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexEntryGroupIamPolicyArgs)(nil)).Elem()
}

type DataplexEntryGroupIamPolicyInput interface {
	pulumi.Input

	ToDataplexEntryGroupIamPolicyOutput() DataplexEntryGroupIamPolicyOutput
	ToDataplexEntryGroupIamPolicyOutputWithContext(ctx context.Context) DataplexEntryGroupIamPolicyOutput
}

func (*DataplexEntryGroupIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexEntryGroupIamPolicy)(nil)).Elem()
}

func (i *DataplexEntryGroupIamPolicy) ToDataplexEntryGroupIamPolicyOutput() DataplexEntryGroupIamPolicyOutput {
	return i.ToDataplexEntryGroupIamPolicyOutputWithContext(context.Background())
}

func (i *DataplexEntryGroupIamPolicy) ToDataplexEntryGroupIamPolicyOutputWithContext(ctx context.Context) DataplexEntryGroupIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexEntryGroupIamPolicyOutput)
}

type DataplexEntryGroupIamPolicyOutput struct{ *pulumi.OutputState }

func (DataplexEntryGroupIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexEntryGroupIamPolicy)(nil)).Elem()
}

func (o DataplexEntryGroupIamPolicyOutput) ToDataplexEntryGroupIamPolicyOutput() DataplexEntryGroupIamPolicyOutput {
	return o
}

func (o DataplexEntryGroupIamPolicyOutput) ToDataplexEntryGroupIamPolicyOutputWithContext(ctx context.Context) DataplexEntryGroupIamPolicyOutput {
	return o
}

func (o DataplexEntryGroupIamPolicyOutput) DataplexEntryGroupIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroupIamPolicy) pulumi.StringOutput { return v.DataplexEntryGroupIamPolicyId }).(pulumi.StringOutput)
}

func (o DataplexEntryGroupIamPolicyOutput) EntryGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroupIamPolicy) pulumi.StringOutput { return v.EntryGroupId }).(pulumi.StringOutput)
}

func (o DataplexEntryGroupIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroupIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataplexEntryGroupIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroupIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataplexEntryGroupIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroupIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o DataplexEntryGroupIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryGroupIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexEntryGroupIamPolicyInput)(nil)).Elem(), &DataplexEntryGroupIamPolicy{})
	pulumi.RegisterOutputType(DataplexEntryGroupIamPolicyOutput{})
}
