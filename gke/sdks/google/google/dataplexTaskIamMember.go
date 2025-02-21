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

type DataplexTaskIamMember struct {
	pulumi.CustomResourceState

	Condition               DataplexTaskIamMemberConditionPtrOutput `pulumi:"condition"`
	DataplexTaskIamMemberId pulumi.StringOutput                     `pulumi:"dataplexTaskIamMemberId"`
	Etag                    pulumi.StringOutput                     `pulumi:"etag"`
	Lake                    pulumi.StringOutput                     `pulumi:"lake"`
	Location                pulumi.StringOutput                     `pulumi:"location"`
	Member                  pulumi.StringOutput                     `pulumi:"member"`
	Project                 pulumi.StringOutput                     `pulumi:"project"`
	Role                    pulumi.StringOutput                     `pulumi:"role"`
	TaskId                  pulumi.StringOutput                     `pulumi:"taskId"`
}

// NewDataplexTaskIamMember registers a new resource with the given unique name, arguments, and options.
func NewDataplexTaskIamMember(ctx *pulumi.Context,
	name string, args *DataplexTaskIamMemberArgs, opts ...pulumi.ResourceOption) (*DataplexTaskIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Lake == nil {
		return nil, errors.New("invalid value for required argument 'Lake'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TaskId == nil {
		return nil, errors.New("invalid value for required argument 'TaskId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexTaskIamMember
	err = ctx.RegisterPackageResource("google:index/dataplexTaskIamMember:DataplexTaskIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexTaskIamMember gets an existing DataplexTaskIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexTaskIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexTaskIamMemberState, opts ...pulumi.ResourceOption) (*DataplexTaskIamMember, error) {
	var resource DataplexTaskIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataplexTaskIamMember:DataplexTaskIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexTaskIamMember resources.
type dataplexTaskIamMemberState struct {
	Condition               *DataplexTaskIamMemberCondition `pulumi:"condition"`
	DataplexTaskIamMemberId *string                         `pulumi:"dataplexTaskIamMemberId"`
	Etag                    *string                         `pulumi:"etag"`
	Lake                    *string                         `pulumi:"lake"`
	Location                *string                         `pulumi:"location"`
	Member                  *string                         `pulumi:"member"`
	Project                 *string                         `pulumi:"project"`
	Role                    *string                         `pulumi:"role"`
	TaskId                  *string                         `pulumi:"taskId"`
}

type DataplexTaskIamMemberState struct {
	Condition               DataplexTaskIamMemberConditionPtrInput
	DataplexTaskIamMemberId pulumi.StringPtrInput
	Etag                    pulumi.StringPtrInput
	Lake                    pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	Member                  pulumi.StringPtrInput
	Project                 pulumi.StringPtrInput
	Role                    pulumi.StringPtrInput
	TaskId                  pulumi.StringPtrInput
}

func (DataplexTaskIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexTaskIamMemberState)(nil)).Elem()
}

type dataplexTaskIamMemberArgs struct {
	Condition               *DataplexTaskIamMemberCondition `pulumi:"condition"`
	DataplexTaskIamMemberId *string                         `pulumi:"dataplexTaskIamMemberId"`
	Lake                    string                          `pulumi:"lake"`
	Location                *string                         `pulumi:"location"`
	Member                  string                          `pulumi:"member"`
	Project                 *string                         `pulumi:"project"`
	Role                    string                          `pulumi:"role"`
	TaskId                  string                          `pulumi:"taskId"`
}

// The set of arguments for constructing a DataplexTaskIamMember resource.
type DataplexTaskIamMemberArgs struct {
	Condition               DataplexTaskIamMemberConditionPtrInput
	DataplexTaskIamMemberId pulumi.StringPtrInput
	Lake                    pulumi.StringInput
	Location                pulumi.StringPtrInput
	Member                  pulumi.StringInput
	Project                 pulumi.StringPtrInput
	Role                    pulumi.StringInput
	TaskId                  pulumi.StringInput
}

func (DataplexTaskIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexTaskIamMemberArgs)(nil)).Elem()
}

type DataplexTaskIamMemberInput interface {
	pulumi.Input

	ToDataplexTaskIamMemberOutput() DataplexTaskIamMemberOutput
	ToDataplexTaskIamMemberOutputWithContext(ctx context.Context) DataplexTaskIamMemberOutput
}

func (*DataplexTaskIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexTaskIamMember)(nil)).Elem()
}

func (i *DataplexTaskIamMember) ToDataplexTaskIamMemberOutput() DataplexTaskIamMemberOutput {
	return i.ToDataplexTaskIamMemberOutputWithContext(context.Background())
}

func (i *DataplexTaskIamMember) ToDataplexTaskIamMemberOutputWithContext(ctx context.Context) DataplexTaskIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexTaskIamMemberOutput)
}

type DataplexTaskIamMemberOutput struct{ *pulumi.OutputState }

func (DataplexTaskIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexTaskIamMember)(nil)).Elem()
}

func (o DataplexTaskIamMemberOutput) ToDataplexTaskIamMemberOutput() DataplexTaskIamMemberOutput {
	return o
}

func (o DataplexTaskIamMemberOutput) ToDataplexTaskIamMemberOutputWithContext(ctx context.Context) DataplexTaskIamMemberOutput {
	return o
}

func (o DataplexTaskIamMemberOutput) Condition() DataplexTaskIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) DataplexTaskIamMemberConditionPtrOutput { return v.Condition }).(DataplexTaskIamMemberConditionPtrOutput)
}

func (o DataplexTaskIamMemberOutput) DataplexTaskIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.DataplexTaskIamMemberId }).(pulumi.StringOutput)
}

func (o DataplexTaskIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataplexTaskIamMemberOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

func (o DataplexTaskIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataplexTaskIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o DataplexTaskIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataplexTaskIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o DataplexTaskIamMemberOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTaskIamMember) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexTaskIamMemberInput)(nil)).Elem(), &DataplexTaskIamMember{})
	pulumi.RegisterOutputType(DataplexTaskIamMemberOutput{})
}
