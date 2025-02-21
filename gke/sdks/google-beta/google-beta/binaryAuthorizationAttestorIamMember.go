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

type BinaryAuthorizationAttestorIamMember struct {
	pulumi.CustomResourceState

	Attestor                               pulumi.StringOutput                                    `pulumi:"attestor"`
	BinaryAuthorizationAttestorIamMemberId pulumi.StringOutput                                    `pulumi:"binaryAuthorizationAttestorIamMemberId"`
	Condition                              BinaryAuthorizationAttestorIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                                   pulumi.StringOutput                                    `pulumi:"etag"`
	Member                                 pulumi.StringOutput                                    `pulumi:"member"`
	Project                                pulumi.StringOutput                                    `pulumi:"project"`
	Role                                   pulumi.StringOutput                                    `pulumi:"role"`
}

// NewBinaryAuthorizationAttestorIamMember registers a new resource with the given unique name, arguments, and options.
func NewBinaryAuthorizationAttestorIamMember(ctx *pulumi.Context,
	name string, args *BinaryAuthorizationAttestorIamMemberArgs, opts ...pulumi.ResourceOption) (*BinaryAuthorizationAttestorIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Attestor == nil {
		return nil, errors.New("invalid value for required argument 'Attestor'")
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
	var resource BinaryAuthorizationAttestorIamMember
	err = ctx.RegisterPackageResource("google-beta:index/binaryAuthorizationAttestorIamMember:BinaryAuthorizationAttestorIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinaryAuthorizationAttestorIamMember gets an existing BinaryAuthorizationAttestorIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinaryAuthorizationAttestorIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BinaryAuthorizationAttestorIamMemberState, opts ...pulumi.ResourceOption) (*BinaryAuthorizationAttestorIamMember, error) {
	var resource BinaryAuthorizationAttestorIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/binaryAuthorizationAttestorIamMember:BinaryAuthorizationAttestorIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BinaryAuthorizationAttestorIamMember resources.
type binaryAuthorizationAttestorIamMemberState struct {
	Attestor                               *string                                        `pulumi:"attestor"`
	BinaryAuthorizationAttestorIamMemberId *string                                        `pulumi:"binaryAuthorizationAttestorIamMemberId"`
	Condition                              *BinaryAuthorizationAttestorIamMemberCondition `pulumi:"condition"`
	Etag                                   *string                                        `pulumi:"etag"`
	Member                                 *string                                        `pulumi:"member"`
	Project                                *string                                        `pulumi:"project"`
	Role                                   *string                                        `pulumi:"role"`
}

type BinaryAuthorizationAttestorIamMemberState struct {
	Attestor                               pulumi.StringPtrInput
	BinaryAuthorizationAttestorIamMemberId pulumi.StringPtrInput
	Condition                              BinaryAuthorizationAttestorIamMemberConditionPtrInput
	Etag                                   pulumi.StringPtrInput
	Member                                 pulumi.StringPtrInput
	Project                                pulumi.StringPtrInput
	Role                                   pulumi.StringPtrInput
}

func (BinaryAuthorizationAttestorIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*binaryAuthorizationAttestorIamMemberState)(nil)).Elem()
}

type binaryAuthorizationAttestorIamMemberArgs struct {
	Attestor                               string                                         `pulumi:"attestor"`
	BinaryAuthorizationAttestorIamMemberId *string                                        `pulumi:"binaryAuthorizationAttestorIamMemberId"`
	Condition                              *BinaryAuthorizationAttestorIamMemberCondition `pulumi:"condition"`
	Member                                 string                                         `pulumi:"member"`
	Project                                *string                                        `pulumi:"project"`
	Role                                   string                                         `pulumi:"role"`
}

// The set of arguments for constructing a BinaryAuthorizationAttestorIamMember resource.
type BinaryAuthorizationAttestorIamMemberArgs struct {
	Attestor                               pulumi.StringInput
	BinaryAuthorizationAttestorIamMemberId pulumi.StringPtrInput
	Condition                              BinaryAuthorizationAttestorIamMemberConditionPtrInput
	Member                                 pulumi.StringInput
	Project                                pulumi.StringPtrInput
	Role                                   pulumi.StringInput
}

func (BinaryAuthorizationAttestorIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*binaryAuthorizationAttestorIamMemberArgs)(nil)).Elem()
}

type BinaryAuthorizationAttestorIamMemberInput interface {
	pulumi.Input

	ToBinaryAuthorizationAttestorIamMemberOutput() BinaryAuthorizationAttestorIamMemberOutput
	ToBinaryAuthorizationAttestorIamMemberOutputWithContext(ctx context.Context) BinaryAuthorizationAttestorIamMemberOutput
}

func (*BinaryAuthorizationAttestorIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BinaryAuthorizationAttestorIamMember)(nil)).Elem()
}

func (i *BinaryAuthorizationAttestorIamMember) ToBinaryAuthorizationAttestorIamMemberOutput() BinaryAuthorizationAttestorIamMemberOutput {
	return i.ToBinaryAuthorizationAttestorIamMemberOutputWithContext(context.Background())
}

func (i *BinaryAuthorizationAttestorIamMember) ToBinaryAuthorizationAttestorIamMemberOutputWithContext(ctx context.Context) BinaryAuthorizationAttestorIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BinaryAuthorizationAttestorIamMemberOutput)
}

type BinaryAuthorizationAttestorIamMemberOutput struct{ *pulumi.OutputState }

func (BinaryAuthorizationAttestorIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BinaryAuthorizationAttestorIamMember)(nil)).Elem()
}

func (o BinaryAuthorizationAttestorIamMemberOutput) ToBinaryAuthorizationAttestorIamMemberOutput() BinaryAuthorizationAttestorIamMemberOutput {
	return o
}

func (o BinaryAuthorizationAttestorIamMemberOutput) ToBinaryAuthorizationAttestorIamMemberOutputWithContext(ctx context.Context) BinaryAuthorizationAttestorIamMemberOutput {
	return o
}

func (o BinaryAuthorizationAttestorIamMemberOutput) Attestor() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamMember) pulumi.StringOutput { return v.Attestor }).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamMemberOutput) BinaryAuthorizationAttestorIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamMember) pulumi.StringOutput {
		return v.BinaryAuthorizationAttestorIamMemberId
	}).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamMemberOutput) Condition() BinaryAuthorizationAttestorIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamMember) BinaryAuthorizationAttestorIamMemberConditionPtrOutput {
		return v.Condition
	}).(BinaryAuthorizationAttestorIamMemberConditionPtrOutput)
}

func (o BinaryAuthorizationAttestorIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BinaryAuthorizationAttestorIamMemberInput)(nil)).Elem(), &BinaryAuthorizationAttestorIamMember{})
	pulumi.RegisterOutputType(BinaryAuthorizationAttestorIamMemberOutput{})
}
