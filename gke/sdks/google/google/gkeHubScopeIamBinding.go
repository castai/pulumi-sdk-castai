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

type GkeHubScopeIamBinding struct {
	pulumi.CustomResourceState

	Condition               GkeHubScopeIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                    pulumi.StringOutput                     `pulumi:"etag"`
	GkeHubScopeIamBindingId pulumi.StringOutput                     `pulumi:"gkeHubScopeIamBindingId"`
	Members                 pulumi.StringArrayOutput                `pulumi:"members"`
	Project                 pulumi.StringOutput                     `pulumi:"project"`
	Role                    pulumi.StringOutput                     `pulumi:"role"`
	ScopeId                 pulumi.StringOutput                     `pulumi:"scopeId"`
}

// NewGkeHubScopeIamBinding registers a new resource with the given unique name, arguments, and options.
func NewGkeHubScopeIamBinding(ctx *pulumi.Context,
	name string, args *GkeHubScopeIamBindingArgs, opts ...pulumi.ResourceOption) (*GkeHubScopeIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GkeHubScopeIamBinding
	err = ctx.RegisterPackageResource("google:index/gkeHubScopeIamBinding:GkeHubScopeIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeHubScopeIamBinding gets an existing GkeHubScopeIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeHubScopeIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeHubScopeIamBindingState, opts ...pulumi.ResourceOption) (*GkeHubScopeIamBinding, error) {
	var resource GkeHubScopeIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/gkeHubScopeIamBinding:GkeHubScopeIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeHubScopeIamBinding resources.
type gkeHubScopeIamBindingState struct {
	Condition               *GkeHubScopeIamBindingCondition `pulumi:"condition"`
	Etag                    *string                         `pulumi:"etag"`
	GkeHubScopeIamBindingId *string                         `pulumi:"gkeHubScopeIamBindingId"`
	Members                 []string                        `pulumi:"members"`
	Project                 *string                         `pulumi:"project"`
	Role                    *string                         `pulumi:"role"`
	ScopeId                 *string                         `pulumi:"scopeId"`
}

type GkeHubScopeIamBindingState struct {
	Condition               GkeHubScopeIamBindingConditionPtrInput
	Etag                    pulumi.StringPtrInput
	GkeHubScopeIamBindingId pulumi.StringPtrInput
	Members                 pulumi.StringArrayInput
	Project                 pulumi.StringPtrInput
	Role                    pulumi.StringPtrInput
	ScopeId                 pulumi.StringPtrInput
}

func (GkeHubScopeIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubScopeIamBindingState)(nil)).Elem()
}

type gkeHubScopeIamBindingArgs struct {
	Condition               *GkeHubScopeIamBindingCondition `pulumi:"condition"`
	GkeHubScopeIamBindingId *string                         `pulumi:"gkeHubScopeIamBindingId"`
	Members                 []string                        `pulumi:"members"`
	Project                 *string                         `pulumi:"project"`
	Role                    string                          `pulumi:"role"`
	ScopeId                 string                          `pulumi:"scopeId"`
}

// The set of arguments for constructing a GkeHubScopeIamBinding resource.
type GkeHubScopeIamBindingArgs struct {
	Condition               GkeHubScopeIamBindingConditionPtrInput
	GkeHubScopeIamBindingId pulumi.StringPtrInput
	Members                 pulumi.StringArrayInput
	Project                 pulumi.StringPtrInput
	Role                    pulumi.StringInput
	ScopeId                 pulumi.StringInput
}

func (GkeHubScopeIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubScopeIamBindingArgs)(nil)).Elem()
}

type GkeHubScopeIamBindingInput interface {
	pulumi.Input

	ToGkeHubScopeIamBindingOutput() GkeHubScopeIamBindingOutput
	ToGkeHubScopeIamBindingOutputWithContext(ctx context.Context) GkeHubScopeIamBindingOutput
}

func (*GkeHubScopeIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubScopeIamBinding)(nil)).Elem()
}

func (i *GkeHubScopeIamBinding) ToGkeHubScopeIamBindingOutput() GkeHubScopeIamBindingOutput {
	return i.ToGkeHubScopeIamBindingOutputWithContext(context.Background())
}

func (i *GkeHubScopeIamBinding) ToGkeHubScopeIamBindingOutputWithContext(ctx context.Context) GkeHubScopeIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeHubScopeIamBindingOutput)
}

type GkeHubScopeIamBindingOutput struct{ *pulumi.OutputState }

func (GkeHubScopeIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubScopeIamBinding)(nil)).Elem()
}

func (o GkeHubScopeIamBindingOutput) ToGkeHubScopeIamBindingOutput() GkeHubScopeIamBindingOutput {
	return o
}

func (o GkeHubScopeIamBindingOutput) ToGkeHubScopeIamBindingOutputWithContext(ctx context.Context) GkeHubScopeIamBindingOutput {
	return o
}

func (o GkeHubScopeIamBindingOutput) Condition() GkeHubScopeIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *GkeHubScopeIamBinding) GkeHubScopeIamBindingConditionPtrOutput { return v.Condition }).(GkeHubScopeIamBindingConditionPtrOutput)
}

func (o GkeHubScopeIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamBindingOutput) GkeHubScopeIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamBinding) pulumi.StringOutput { return v.GkeHubScopeIamBindingId }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GkeHubScopeIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o GkeHubScopeIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o GkeHubScopeIamBindingOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubScopeIamBinding) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeHubScopeIamBindingInput)(nil)).Elem(), &GkeHubScopeIamBinding{})
	pulumi.RegisterOutputType(GkeHubScopeIamBindingOutput{})
}
