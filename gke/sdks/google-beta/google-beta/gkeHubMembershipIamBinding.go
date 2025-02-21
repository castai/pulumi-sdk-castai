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

type GkeHubMembershipIamBinding struct {
	pulumi.CustomResourceState

	Condition                    GkeHubMembershipIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                         pulumi.StringOutput                          `pulumi:"etag"`
	GkeHubMembershipIamBindingId pulumi.StringOutput                          `pulumi:"gkeHubMembershipIamBindingId"`
	Location                     pulumi.StringOutput                          `pulumi:"location"`
	Members                      pulumi.StringArrayOutput                     `pulumi:"members"`
	MembershipId                 pulumi.StringOutput                          `pulumi:"membershipId"`
	Project                      pulumi.StringOutput                          `pulumi:"project"`
	Role                         pulumi.StringOutput                          `pulumi:"role"`
}

// NewGkeHubMembershipIamBinding registers a new resource with the given unique name, arguments, and options.
func NewGkeHubMembershipIamBinding(ctx *pulumi.Context,
	name string, args *GkeHubMembershipIamBindingArgs, opts ...pulumi.ResourceOption) (*GkeHubMembershipIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.MembershipId == nil {
		return nil, errors.New("invalid value for required argument 'MembershipId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GkeHubMembershipIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/gkeHubMembershipIamBinding:GkeHubMembershipIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeHubMembershipIamBinding gets an existing GkeHubMembershipIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeHubMembershipIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeHubMembershipIamBindingState, opts ...pulumi.ResourceOption) (*GkeHubMembershipIamBinding, error) {
	var resource GkeHubMembershipIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/gkeHubMembershipIamBinding:GkeHubMembershipIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeHubMembershipIamBinding resources.
type gkeHubMembershipIamBindingState struct {
	Condition                    *GkeHubMembershipIamBindingCondition `pulumi:"condition"`
	Etag                         *string                              `pulumi:"etag"`
	GkeHubMembershipIamBindingId *string                              `pulumi:"gkeHubMembershipIamBindingId"`
	Location                     *string                              `pulumi:"location"`
	Members                      []string                             `pulumi:"members"`
	MembershipId                 *string                              `pulumi:"membershipId"`
	Project                      *string                              `pulumi:"project"`
	Role                         *string                              `pulumi:"role"`
}

type GkeHubMembershipIamBindingState struct {
	Condition                    GkeHubMembershipIamBindingConditionPtrInput
	Etag                         pulumi.StringPtrInput
	GkeHubMembershipIamBindingId pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Members                      pulumi.StringArrayInput
	MembershipId                 pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringPtrInput
}

func (GkeHubMembershipIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubMembershipIamBindingState)(nil)).Elem()
}

type gkeHubMembershipIamBindingArgs struct {
	Condition                    *GkeHubMembershipIamBindingCondition `pulumi:"condition"`
	GkeHubMembershipIamBindingId *string                              `pulumi:"gkeHubMembershipIamBindingId"`
	Location                     *string                              `pulumi:"location"`
	Members                      []string                             `pulumi:"members"`
	MembershipId                 string                               `pulumi:"membershipId"`
	Project                      *string                              `pulumi:"project"`
	Role                         string                               `pulumi:"role"`
}

// The set of arguments for constructing a GkeHubMembershipIamBinding resource.
type GkeHubMembershipIamBindingArgs struct {
	Condition                    GkeHubMembershipIamBindingConditionPtrInput
	GkeHubMembershipIamBindingId pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Members                      pulumi.StringArrayInput
	MembershipId                 pulumi.StringInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringInput
}

func (GkeHubMembershipIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubMembershipIamBindingArgs)(nil)).Elem()
}

type GkeHubMembershipIamBindingInput interface {
	pulumi.Input

	ToGkeHubMembershipIamBindingOutput() GkeHubMembershipIamBindingOutput
	ToGkeHubMembershipIamBindingOutputWithContext(ctx context.Context) GkeHubMembershipIamBindingOutput
}

func (*GkeHubMembershipIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubMembershipIamBinding)(nil)).Elem()
}

func (i *GkeHubMembershipIamBinding) ToGkeHubMembershipIamBindingOutput() GkeHubMembershipIamBindingOutput {
	return i.ToGkeHubMembershipIamBindingOutputWithContext(context.Background())
}

func (i *GkeHubMembershipIamBinding) ToGkeHubMembershipIamBindingOutputWithContext(ctx context.Context) GkeHubMembershipIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeHubMembershipIamBindingOutput)
}

type GkeHubMembershipIamBindingOutput struct{ *pulumi.OutputState }

func (GkeHubMembershipIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubMembershipIamBinding)(nil)).Elem()
}

func (o GkeHubMembershipIamBindingOutput) ToGkeHubMembershipIamBindingOutput() GkeHubMembershipIamBindingOutput {
	return o
}

func (o GkeHubMembershipIamBindingOutput) ToGkeHubMembershipIamBindingOutputWithContext(ctx context.Context) GkeHubMembershipIamBindingOutput {
	return o
}

func (o GkeHubMembershipIamBindingOutput) Condition() GkeHubMembershipIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) GkeHubMembershipIamBindingConditionPtrOutput { return v.Condition }).(GkeHubMembershipIamBindingConditionPtrOutput)
}

func (o GkeHubMembershipIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GkeHubMembershipIamBindingOutput) GkeHubMembershipIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) pulumi.StringOutput { return v.GkeHubMembershipIamBindingId }).(pulumi.StringOutput)
}

func (o GkeHubMembershipIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GkeHubMembershipIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o GkeHubMembershipIamBindingOutput) MembershipId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) pulumi.StringOutput { return v.MembershipId }).(pulumi.StringOutput)
}

func (o GkeHubMembershipIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o GkeHubMembershipIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubMembershipIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeHubMembershipIamBindingInput)(nil)).Elem(), &GkeHubMembershipIamBinding{})
	pulumi.RegisterOutputType(GkeHubMembershipIamBindingOutput{})
}
