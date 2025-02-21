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

type DataprocJobIamMember struct {
	pulumi.CustomResourceState

	Condition              DataprocJobIamMemberConditionPtrOutput `pulumi:"condition"`
	DataprocJobIamMemberId pulumi.StringOutput                    `pulumi:"dataprocJobIamMemberId"`
	Etag                   pulumi.StringOutput                    `pulumi:"etag"`
	JobId                  pulumi.StringOutput                    `pulumi:"jobId"`
	Member                 pulumi.StringOutput                    `pulumi:"member"`
	Project                pulumi.StringOutput                    `pulumi:"project"`
	Region                 pulumi.StringOutput                    `pulumi:"region"`
	Role                   pulumi.StringOutput                    `pulumi:"role"`
}

// NewDataprocJobIamMember registers a new resource with the given unique name, arguments, and options.
func NewDataprocJobIamMember(ctx *pulumi.Context,
	name string, args *DataprocJobIamMemberArgs, opts ...pulumi.ResourceOption) (*DataprocJobIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
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
	var resource DataprocJobIamMember
	err = ctx.RegisterPackageResource("google:index/dataprocJobIamMember:DataprocJobIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocJobIamMember gets an existing DataprocJobIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocJobIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocJobIamMemberState, opts ...pulumi.ResourceOption) (*DataprocJobIamMember, error) {
	var resource DataprocJobIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataprocJobIamMember:DataprocJobIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocJobIamMember resources.
type dataprocJobIamMemberState struct {
	Condition              *DataprocJobIamMemberCondition `pulumi:"condition"`
	DataprocJobIamMemberId *string                        `pulumi:"dataprocJobIamMemberId"`
	Etag                   *string                        `pulumi:"etag"`
	JobId                  *string                        `pulumi:"jobId"`
	Member                 *string                        `pulumi:"member"`
	Project                *string                        `pulumi:"project"`
	Region                 *string                        `pulumi:"region"`
	Role                   *string                        `pulumi:"role"`
}

type DataprocJobIamMemberState struct {
	Condition              DataprocJobIamMemberConditionPtrInput
	DataprocJobIamMemberId pulumi.StringPtrInput
	Etag                   pulumi.StringPtrInput
	JobId                  pulumi.StringPtrInput
	Member                 pulumi.StringPtrInput
	Project                pulumi.StringPtrInput
	Region                 pulumi.StringPtrInput
	Role                   pulumi.StringPtrInput
}

func (DataprocJobIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocJobIamMemberState)(nil)).Elem()
}

type dataprocJobIamMemberArgs struct {
	Condition              *DataprocJobIamMemberCondition `pulumi:"condition"`
	DataprocJobIamMemberId *string                        `pulumi:"dataprocJobIamMemberId"`
	JobId                  string                         `pulumi:"jobId"`
	Member                 string                         `pulumi:"member"`
	Project                *string                        `pulumi:"project"`
	Region                 *string                        `pulumi:"region"`
	Role                   string                         `pulumi:"role"`
}

// The set of arguments for constructing a DataprocJobIamMember resource.
type DataprocJobIamMemberArgs struct {
	Condition              DataprocJobIamMemberConditionPtrInput
	DataprocJobIamMemberId pulumi.StringPtrInput
	JobId                  pulumi.StringInput
	Member                 pulumi.StringInput
	Project                pulumi.StringPtrInput
	Region                 pulumi.StringPtrInput
	Role                   pulumi.StringInput
}

func (DataprocJobIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocJobIamMemberArgs)(nil)).Elem()
}

type DataprocJobIamMemberInput interface {
	pulumi.Input

	ToDataprocJobIamMemberOutput() DataprocJobIamMemberOutput
	ToDataprocJobIamMemberOutputWithContext(ctx context.Context) DataprocJobIamMemberOutput
}

func (*DataprocJobIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocJobIamMember)(nil)).Elem()
}

func (i *DataprocJobIamMember) ToDataprocJobIamMemberOutput() DataprocJobIamMemberOutput {
	return i.ToDataprocJobIamMemberOutputWithContext(context.Background())
}

func (i *DataprocJobIamMember) ToDataprocJobIamMemberOutputWithContext(ctx context.Context) DataprocJobIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocJobIamMemberOutput)
}

type DataprocJobIamMemberOutput struct{ *pulumi.OutputState }

func (DataprocJobIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocJobIamMember)(nil)).Elem()
}

func (o DataprocJobIamMemberOutput) ToDataprocJobIamMemberOutput() DataprocJobIamMemberOutput {
	return o
}

func (o DataprocJobIamMemberOutput) ToDataprocJobIamMemberOutputWithContext(ctx context.Context) DataprocJobIamMemberOutput {
	return o
}

func (o DataprocJobIamMemberOutput) Condition() DataprocJobIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) DataprocJobIamMemberConditionPtrOutput { return v.Condition }).(DataprocJobIamMemberConditionPtrOutput)
}

func (o DataprocJobIamMemberOutput) DataprocJobIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) pulumi.StringOutput { return v.DataprocJobIamMemberId }).(pulumi.StringOutput)
}

func (o DataprocJobIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataprocJobIamMemberOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

func (o DataprocJobIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o DataprocJobIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataprocJobIamMemberOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o DataprocJobIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocJobIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocJobIamMemberInput)(nil)).Elem(), &DataprocJobIamMember{})
	pulumi.RegisterOutputType(DataprocJobIamMemberOutput{})
}
