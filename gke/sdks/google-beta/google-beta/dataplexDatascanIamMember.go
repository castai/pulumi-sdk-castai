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

type DataplexDatascanIamMember struct {
	pulumi.CustomResourceState

	Condition                   DataplexDatascanIamMemberConditionPtrOutput `pulumi:"condition"`
	DataScanId                  pulumi.StringOutput                         `pulumi:"dataScanId"`
	DataplexDatascanIamMemberId pulumi.StringOutput                         `pulumi:"dataplexDatascanIamMemberId"`
	Etag                        pulumi.StringOutput                         `pulumi:"etag"`
	Location                    pulumi.StringOutput                         `pulumi:"location"`
	Member                      pulumi.StringOutput                         `pulumi:"member"`
	Project                     pulumi.StringOutput                         `pulumi:"project"`
	Role                        pulumi.StringOutput                         `pulumi:"role"`
}

// NewDataplexDatascanIamMember registers a new resource with the given unique name, arguments, and options.
func NewDataplexDatascanIamMember(ctx *pulumi.Context,
	name string, args *DataplexDatascanIamMemberArgs, opts ...pulumi.ResourceOption) (*DataplexDatascanIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataScanId == nil {
		return nil, errors.New("invalid value for required argument 'DataScanId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexDatascanIamMember
	err = ctx.RegisterPackageResource("google-beta:index/dataplexDatascanIamMember:DataplexDatascanIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexDatascanIamMember gets an existing DataplexDatascanIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexDatascanIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexDatascanIamMemberState, opts ...pulumi.ResourceOption) (*DataplexDatascanIamMember, error) {
	var resource DataplexDatascanIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataplexDatascanIamMember:DataplexDatascanIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexDatascanIamMember resources.
type dataplexDatascanIamMemberState struct {
	Condition                   *DataplexDatascanIamMemberCondition `pulumi:"condition"`
	DataScanId                  *string                             `pulumi:"dataScanId"`
	DataplexDatascanIamMemberId *string                             `pulumi:"dataplexDatascanIamMemberId"`
	Etag                        *string                             `pulumi:"etag"`
	Location                    *string                             `pulumi:"location"`
	Member                      *string                             `pulumi:"member"`
	Project                     *string                             `pulumi:"project"`
	Role                        *string                             `pulumi:"role"`
}

type DataplexDatascanIamMemberState struct {
	Condition                   DataplexDatascanIamMemberConditionPtrInput
	DataScanId                  pulumi.StringPtrInput
	DataplexDatascanIamMemberId pulumi.StringPtrInput
	Etag                        pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	Member                      pulumi.StringPtrInput
	Project                     pulumi.StringPtrInput
	Role                        pulumi.StringPtrInput
}

func (DataplexDatascanIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexDatascanIamMemberState)(nil)).Elem()
}

type dataplexDatascanIamMemberArgs struct {
	Condition                   *DataplexDatascanIamMemberCondition `pulumi:"condition"`
	DataScanId                  string                              `pulumi:"dataScanId"`
	DataplexDatascanIamMemberId *string                             `pulumi:"dataplexDatascanIamMemberId"`
	Location                    *string                             `pulumi:"location"`
	Member                      string                              `pulumi:"member"`
	Project                     *string                             `pulumi:"project"`
	Role                        string                              `pulumi:"role"`
}

// The set of arguments for constructing a DataplexDatascanIamMember resource.
type DataplexDatascanIamMemberArgs struct {
	Condition                   DataplexDatascanIamMemberConditionPtrInput
	DataScanId                  pulumi.StringInput
	DataplexDatascanIamMemberId pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	Member                      pulumi.StringInput
	Project                     pulumi.StringPtrInput
	Role                        pulumi.StringInput
}

func (DataplexDatascanIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexDatascanIamMemberArgs)(nil)).Elem()
}

type DataplexDatascanIamMemberInput interface {
	pulumi.Input

	ToDataplexDatascanIamMemberOutput() DataplexDatascanIamMemberOutput
	ToDataplexDatascanIamMemberOutputWithContext(ctx context.Context) DataplexDatascanIamMemberOutput
}

func (*DataplexDatascanIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexDatascanIamMember)(nil)).Elem()
}

func (i *DataplexDatascanIamMember) ToDataplexDatascanIamMemberOutput() DataplexDatascanIamMemberOutput {
	return i.ToDataplexDatascanIamMemberOutputWithContext(context.Background())
}

func (i *DataplexDatascanIamMember) ToDataplexDatascanIamMemberOutputWithContext(ctx context.Context) DataplexDatascanIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexDatascanIamMemberOutput)
}

type DataplexDatascanIamMemberOutput struct{ *pulumi.OutputState }

func (DataplexDatascanIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexDatascanIamMember)(nil)).Elem()
}

func (o DataplexDatascanIamMemberOutput) ToDataplexDatascanIamMemberOutput() DataplexDatascanIamMemberOutput {
	return o
}

func (o DataplexDatascanIamMemberOutput) ToDataplexDatascanIamMemberOutputWithContext(ctx context.Context) DataplexDatascanIamMemberOutput {
	return o
}

func (o DataplexDatascanIamMemberOutput) Condition() DataplexDatascanIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) DataplexDatascanIamMemberConditionPtrOutput { return v.Condition }).(DataplexDatascanIamMemberConditionPtrOutput)
}

func (o DataplexDatascanIamMemberOutput) DataScanId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) pulumi.StringOutput { return v.DataScanId }).(pulumi.StringOutput)
}

func (o DataplexDatascanIamMemberOutput) DataplexDatascanIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) pulumi.StringOutput { return v.DataplexDatascanIamMemberId }).(pulumi.StringOutput)
}

func (o DataplexDatascanIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataplexDatascanIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataplexDatascanIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o DataplexDatascanIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataplexDatascanIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexDatascanIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexDatascanIamMemberInput)(nil)).Elem(), &DataplexDatascanIamMember{})
	pulumi.RegisterOutputType(DataplexDatascanIamMemberOutput{})
}
