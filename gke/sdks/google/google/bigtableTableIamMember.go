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

type BigtableTableIamMember struct {
	pulumi.CustomResourceState

	BigtableTableIamMemberId pulumi.StringOutput                      `pulumi:"bigtableTableIamMemberId"`
	Condition                BigtableTableIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                     pulumi.StringOutput                      `pulumi:"etag"`
	Instance                 pulumi.StringOutput                      `pulumi:"instance"`
	Member                   pulumi.StringOutput                      `pulumi:"member"`
	Project                  pulumi.StringOutput                      `pulumi:"project"`
	Role                     pulumi.StringOutput                      `pulumi:"role"`
	Table                    pulumi.StringOutput                      `pulumi:"table"`
}

// NewBigtableTableIamMember registers a new resource with the given unique name, arguments, and options.
func NewBigtableTableIamMember(ctx *pulumi.Context,
	name string, args *BigtableTableIamMemberArgs, opts ...pulumi.ResourceOption) (*BigtableTableIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Table == nil {
		return nil, errors.New("invalid value for required argument 'Table'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigtableTableIamMember
	err = ctx.RegisterPackageResource("google:index/bigtableTableIamMember:BigtableTableIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigtableTableIamMember gets an existing BigtableTableIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigtableTableIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigtableTableIamMemberState, opts ...pulumi.ResourceOption) (*BigtableTableIamMember, error) {
	var resource BigtableTableIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/bigtableTableIamMember:BigtableTableIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigtableTableIamMember resources.
type bigtableTableIamMemberState struct {
	BigtableTableIamMemberId *string                          `pulumi:"bigtableTableIamMemberId"`
	Condition                *BigtableTableIamMemberCondition `pulumi:"condition"`
	Etag                     *string                          `pulumi:"etag"`
	Instance                 *string                          `pulumi:"instance"`
	Member                   *string                          `pulumi:"member"`
	Project                  *string                          `pulumi:"project"`
	Role                     *string                          `pulumi:"role"`
	Table                    *string                          `pulumi:"table"`
}

type BigtableTableIamMemberState struct {
	BigtableTableIamMemberId pulumi.StringPtrInput
	Condition                BigtableTableIamMemberConditionPtrInput
	Etag                     pulumi.StringPtrInput
	Instance                 pulumi.StringPtrInput
	Member                   pulumi.StringPtrInput
	Project                  pulumi.StringPtrInput
	Role                     pulumi.StringPtrInput
	Table                    pulumi.StringPtrInput
}

func (BigtableTableIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigtableTableIamMemberState)(nil)).Elem()
}

type bigtableTableIamMemberArgs struct {
	BigtableTableIamMemberId *string                          `pulumi:"bigtableTableIamMemberId"`
	Condition                *BigtableTableIamMemberCondition `pulumi:"condition"`
	Instance                 string                           `pulumi:"instance"`
	Member                   string                           `pulumi:"member"`
	Project                  *string                          `pulumi:"project"`
	Role                     string                           `pulumi:"role"`
	Table                    string                           `pulumi:"table"`
}

// The set of arguments for constructing a BigtableTableIamMember resource.
type BigtableTableIamMemberArgs struct {
	BigtableTableIamMemberId pulumi.StringPtrInput
	Condition                BigtableTableIamMemberConditionPtrInput
	Instance                 pulumi.StringInput
	Member                   pulumi.StringInput
	Project                  pulumi.StringPtrInput
	Role                     pulumi.StringInput
	Table                    pulumi.StringInput
}

func (BigtableTableIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigtableTableIamMemberArgs)(nil)).Elem()
}

type BigtableTableIamMemberInput interface {
	pulumi.Input

	ToBigtableTableIamMemberOutput() BigtableTableIamMemberOutput
	ToBigtableTableIamMemberOutputWithContext(ctx context.Context) BigtableTableIamMemberOutput
}

func (*BigtableTableIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BigtableTableIamMember)(nil)).Elem()
}

func (i *BigtableTableIamMember) ToBigtableTableIamMemberOutput() BigtableTableIamMemberOutput {
	return i.ToBigtableTableIamMemberOutputWithContext(context.Background())
}

func (i *BigtableTableIamMember) ToBigtableTableIamMemberOutputWithContext(ctx context.Context) BigtableTableIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigtableTableIamMemberOutput)
}

type BigtableTableIamMemberOutput struct{ *pulumi.OutputState }

func (BigtableTableIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigtableTableIamMember)(nil)).Elem()
}

func (o BigtableTableIamMemberOutput) ToBigtableTableIamMemberOutput() BigtableTableIamMemberOutput {
	return o
}

func (o BigtableTableIamMemberOutput) ToBigtableTableIamMemberOutputWithContext(ctx context.Context) BigtableTableIamMemberOutput {
	return o
}

func (o BigtableTableIamMemberOutput) BigtableTableIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) pulumi.StringOutput { return v.BigtableTableIamMemberId }).(pulumi.StringOutput)
}

func (o BigtableTableIamMemberOutput) Condition() BigtableTableIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) BigtableTableIamMemberConditionPtrOutput { return v.Condition }).(BigtableTableIamMemberConditionPtrOutput)
}

func (o BigtableTableIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BigtableTableIamMemberOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

func (o BigtableTableIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o BigtableTableIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o BigtableTableIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o BigtableTableIamMemberOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableTableIamMember) pulumi.StringOutput { return v.Table }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigtableTableIamMemberInput)(nil)).Elem(), &BigtableTableIamMember{})
	pulumi.RegisterOutputType(BigtableTableIamMemberOutput{})
}
