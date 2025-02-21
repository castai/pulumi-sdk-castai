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

type BigqueryTableIamMember struct {
	pulumi.CustomResourceState

	BigqueryTableIamMemberId pulumi.StringOutput                      `pulumi:"bigqueryTableIamMemberId"`
	Condition                BigqueryTableIamMemberConditionPtrOutput `pulumi:"condition"`
	DatasetId                pulumi.StringOutput                      `pulumi:"datasetId"`
	Etag                     pulumi.StringOutput                      `pulumi:"etag"`
	Member                   pulumi.StringOutput                      `pulumi:"member"`
	Project                  pulumi.StringOutput                      `pulumi:"project"`
	Role                     pulumi.StringOutput                      `pulumi:"role"`
	TableId                  pulumi.StringOutput                      `pulumi:"tableId"`
}

// NewBigqueryTableIamMember registers a new resource with the given unique name, arguments, and options.
func NewBigqueryTableIamMember(ctx *pulumi.Context,
	name string, args *BigqueryTableIamMemberArgs, opts ...pulumi.ResourceOption) (*BigqueryTableIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TableId == nil {
		return nil, errors.New("invalid value for required argument 'TableId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigqueryTableIamMember
	err = ctx.RegisterPackageResource("google:index/bigqueryTableIamMember:BigqueryTableIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryTableIamMember gets an existing BigqueryTableIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryTableIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryTableIamMemberState, opts ...pulumi.ResourceOption) (*BigqueryTableIamMember, error) {
	var resource BigqueryTableIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/bigqueryTableIamMember:BigqueryTableIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryTableIamMember resources.
type bigqueryTableIamMemberState struct {
	BigqueryTableIamMemberId *string                          `pulumi:"bigqueryTableIamMemberId"`
	Condition                *BigqueryTableIamMemberCondition `pulumi:"condition"`
	DatasetId                *string                          `pulumi:"datasetId"`
	Etag                     *string                          `pulumi:"etag"`
	Member                   *string                          `pulumi:"member"`
	Project                  *string                          `pulumi:"project"`
	Role                     *string                          `pulumi:"role"`
	TableId                  *string                          `pulumi:"tableId"`
}

type BigqueryTableIamMemberState struct {
	BigqueryTableIamMemberId pulumi.StringPtrInput
	Condition                BigqueryTableIamMemberConditionPtrInput
	DatasetId                pulumi.StringPtrInput
	Etag                     pulumi.StringPtrInput
	Member                   pulumi.StringPtrInput
	Project                  pulumi.StringPtrInput
	Role                     pulumi.StringPtrInput
	TableId                  pulumi.StringPtrInput
}

func (BigqueryTableIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryTableIamMemberState)(nil)).Elem()
}

type bigqueryTableIamMemberArgs struct {
	BigqueryTableIamMemberId *string                          `pulumi:"bigqueryTableIamMemberId"`
	Condition                *BigqueryTableIamMemberCondition `pulumi:"condition"`
	DatasetId                string                           `pulumi:"datasetId"`
	Member                   string                           `pulumi:"member"`
	Project                  *string                          `pulumi:"project"`
	Role                     string                           `pulumi:"role"`
	TableId                  string                           `pulumi:"tableId"`
}

// The set of arguments for constructing a BigqueryTableIamMember resource.
type BigqueryTableIamMemberArgs struct {
	BigqueryTableIamMemberId pulumi.StringPtrInput
	Condition                BigqueryTableIamMemberConditionPtrInput
	DatasetId                pulumi.StringInput
	Member                   pulumi.StringInput
	Project                  pulumi.StringPtrInput
	Role                     pulumi.StringInput
	TableId                  pulumi.StringInput
}

func (BigqueryTableIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryTableIamMemberArgs)(nil)).Elem()
}

type BigqueryTableIamMemberInput interface {
	pulumi.Input

	ToBigqueryTableIamMemberOutput() BigqueryTableIamMemberOutput
	ToBigqueryTableIamMemberOutputWithContext(ctx context.Context) BigqueryTableIamMemberOutput
}

func (*BigqueryTableIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryTableIamMember)(nil)).Elem()
}

func (i *BigqueryTableIamMember) ToBigqueryTableIamMemberOutput() BigqueryTableIamMemberOutput {
	return i.ToBigqueryTableIamMemberOutputWithContext(context.Background())
}

func (i *BigqueryTableIamMember) ToBigqueryTableIamMemberOutputWithContext(ctx context.Context) BigqueryTableIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryTableIamMemberOutput)
}

type BigqueryTableIamMemberOutput struct{ *pulumi.OutputState }

func (BigqueryTableIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryTableIamMember)(nil)).Elem()
}

func (o BigqueryTableIamMemberOutput) ToBigqueryTableIamMemberOutput() BigqueryTableIamMemberOutput {
	return o
}

func (o BigqueryTableIamMemberOutput) ToBigqueryTableIamMemberOutputWithContext(ctx context.Context) BigqueryTableIamMemberOutput {
	return o
}

func (o BigqueryTableIamMemberOutput) BigqueryTableIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) pulumi.StringOutput { return v.BigqueryTableIamMemberId }).(pulumi.StringOutput)
}

func (o BigqueryTableIamMemberOutput) Condition() BigqueryTableIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) BigqueryTableIamMemberConditionPtrOutput { return v.Condition }).(BigqueryTableIamMemberConditionPtrOutput)
}

func (o BigqueryTableIamMemberOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

func (o BigqueryTableIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BigqueryTableIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o BigqueryTableIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o BigqueryTableIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o BigqueryTableIamMemberOutput) TableId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryTableIamMember) pulumi.StringOutput { return v.TableId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryTableIamMemberInput)(nil)).Elem(), &BigqueryTableIamMember{})
	pulumi.RegisterOutputType(BigqueryTableIamMemberOutput{})
}
