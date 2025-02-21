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

type Cloudfunctions2FunctionIamMember struct {
	pulumi.CustomResourceState

	CloudFunction                      pulumi.StringOutput                                `pulumi:"cloudFunction"`
	Cloudfunctions2FunctionIamMemberId pulumi.StringOutput                                `pulumi:"cloudfunctions2FunctionIamMemberId"`
	Condition                          Cloudfunctions2FunctionIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                               pulumi.StringOutput                                `pulumi:"etag"`
	Location                           pulumi.StringOutput                                `pulumi:"location"`
	Member                             pulumi.StringOutput                                `pulumi:"member"`
	Project                            pulumi.StringOutput                                `pulumi:"project"`
	Role                               pulumi.StringOutput                                `pulumi:"role"`
}

// NewCloudfunctions2FunctionIamMember registers a new resource with the given unique name, arguments, and options.
func NewCloudfunctions2FunctionIamMember(ctx *pulumi.Context,
	name string, args *Cloudfunctions2FunctionIamMemberArgs, opts ...pulumi.ResourceOption) (*Cloudfunctions2FunctionIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudFunction == nil {
		return nil, errors.New("invalid value for required argument 'CloudFunction'")
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
	var resource Cloudfunctions2FunctionIamMember
	err = ctx.RegisterPackageResource("google-beta:index/cloudfunctions2FunctionIamMember:Cloudfunctions2FunctionIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudfunctions2FunctionIamMember gets an existing Cloudfunctions2FunctionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudfunctions2FunctionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Cloudfunctions2FunctionIamMemberState, opts ...pulumi.ResourceOption) (*Cloudfunctions2FunctionIamMember, error) {
	var resource Cloudfunctions2FunctionIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/cloudfunctions2FunctionIamMember:Cloudfunctions2FunctionIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cloudfunctions2FunctionIamMember resources.
type cloudfunctions2FunctionIamMemberState struct {
	CloudFunction                      *string                                    `pulumi:"cloudFunction"`
	Cloudfunctions2FunctionIamMemberId *string                                    `pulumi:"cloudfunctions2FunctionIamMemberId"`
	Condition                          *Cloudfunctions2FunctionIamMemberCondition `pulumi:"condition"`
	Etag                               *string                                    `pulumi:"etag"`
	Location                           *string                                    `pulumi:"location"`
	Member                             *string                                    `pulumi:"member"`
	Project                            *string                                    `pulumi:"project"`
	Role                               *string                                    `pulumi:"role"`
}

type Cloudfunctions2FunctionIamMemberState struct {
	CloudFunction                      pulumi.StringPtrInput
	Cloudfunctions2FunctionIamMemberId pulumi.StringPtrInput
	Condition                          Cloudfunctions2FunctionIamMemberConditionPtrInput
	Etag                               pulumi.StringPtrInput
	Location                           pulumi.StringPtrInput
	Member                             pulumi.StringPtrInput
	Project                            pulumi.StringPtrInput
	Role                               pulumi.StringPtrInput
}

func (Cloudfunctions2FunctionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudfunctions2FunctionIamMemberState)(nil)).Elem()
}

type cloudfunctions2FunctionIamMemberArgs struct {
	CloudFunction                      string                                     `pulumi:"cloudFunction"`
	Cloudfunctions2FunctionIamMemberId *string                                    `pulumi:"cloudfunctions2FunctionIamMemberId"`
	Condition                          *Cloudfunctions2FunctionIamMemberCondition `pulumi:"condition"`
	Location                           *string                                    `pulumi:"location"`
	Member                             string                                     `pulumi:"member"`
	Project                            *string                                    `pulumi:"project"`
	Role                               string                                     `pulumi:"role"`
}

// The set of arguments for constructing a Cloudfunctions2FunctionIamMember resource.
type Cloudfunctions2FunctionIamMemberArgs struct {
	CloudFunction                      pulumi.StringInput
	Cloudfunctions2FunctionIamMemberId pulumi.StringPtrInput
	Condition                          Cloudfunctions2FunctionIamMemberConditionPtrInput
	Location                           pulumi.StringPtrInput
	Member                             pulumi.StringInput
	Project                            pulumi.StringPtrInput
	Role                               pulumi.StringInput
}

func (Cloudfunctions2FunctionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudfunctions2FunctionIamMemberArgs)(nil)).Elem()
}

type Cloudfunctions2FunctionIamMemberInput interface {
	pulumi.Input

	ToCloudfunctions2FunctionIamMemberOutput() Cloudfunctions2FunctionIamMemberOutput
	ToCloudfunctions2FunctionIamMemberOutputWithContext(ctx context.Context) Cloudfunctions2FunctionIamMemberOutput
}

func (*Cloudfunctions2FunctionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloudfunctions2FunctionIamMember)(nil)).Elem()
}

func (i *Cloudfunctions2FunctionIamMember) ToCloudfunctions2FunctionIamMemberOutput() Cloudfunctions2FunctionIamMemberOutput {
	return i.ToCloudfunctions2FunctionIamMemberOutputWithContext(context.Background())
}

func (i *Cloudfunctions2FunctionIamMember) ToCloudfunctions2FunctionIamMemberOutputWithContext(ctx context.Context) Cloudfunctions2FunctionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Cloudfunctions2FunctionIamMemberOutput)
}

type Cloudfunctions2FunctionIamMemberOutput struct{ *pulumi.OutputState }

func (Cloudfunctions2FunctionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloudfunctions2FunctionIamMember)(nil)).Elem()
}

func (o Cloudfunctions2FunctionIamMemberOutput) ToCloudfunctions2FunctionIamMemberOutput() Cloudfunctions2FunctionIamMemberOutput {
	return o
}

func (o Cloudfunctions2FunctionIamMemberOutput) ToCloudfunctions2FunctionIamMemberOutputWithContext(ctx context.Context) Cloudfunctions2FunctionIamMemberOutput {
	return o
}

func (o Cloudfunctions2FunctionIamMemberOutput) CloudFunction() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) pulumi.StringOutput { return v.CloudFunction }).(pulumi.StringOutput)
}

func (o Cloudfunctions2FunctionIamMemberOutput) Cloudfunctions2FunctionIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) pulumi.StringOutput {
		return v.Cloudfunctions2FunctionIamMemberId
	}).(pulumi.StringOutput)
}

func (o Cloudfunctions2FunctionIamMemberOutput) Condition() Cloudfunctions2FunctionIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) Cloudfunctions2FunctionIamMemberConditionPtrOutput {
		return v.Condition
	}).(Cloudfunctions2FunctionIamMemberConditionPtrOutput)
}

func (o Cloudfunctions2FunctionIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o Cloudfunctions2FunctionIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o Cloudfunctions2FunctionIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o Cloudfunctions2FunctionIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o Cloudfunctions2FunctionIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudfunctions2FunctionIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Cloudfunctions2FunctionIamMemberInput)(nil)).Elem(), &Cloudfunctions2FunctionIamMember{})
	pulumi.RegisterOutputType(Cloudfunctions2FunctionIamMemberOutput{})
}
