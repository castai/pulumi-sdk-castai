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

type PubsubSubscriptionIamBinding struct {
	pulumi.CustomResourceState

	Condition                      PubsubSubscriptionIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                           pulumi.StringOutput                            `pulumi:"etag"`
	Members                        pulumi.StringArrayOutput                       `pulumi:"members"`
	Project                        pulumi.StringOutput                            `pulumi:"project"`
	PubsubSubscriptionIamBindingId pulumi.StringOutput                            `pulumi:"pubsubSubscriptionIamBindingId"`
	Role                           pulumi.StringOutput                            `pulumi:"role"`
	Subscription                   pulumi.StringOutput                            `pulumi:"subscription"`
}

// NewPubsubSubscriptionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewPubsubSubscriptionIamBinding(ctx *pulumi.Context,
	name string, args *PubsubSubscriptionIamBindingArgs, opts ...pulumi.ResourceOption) (*PubsubSubscriptionIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Subscription == nil {
		return nil, errors.New("invalid value for required argument 'Subscription'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource PubsubSubscriptionIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/pubsubSubscriptionIamBinding:PubsubSubscriptionIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPubsubSubscriptionIamBinding gets an existing PubsubSubscriptionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPubsubSubscriptionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PubsubSubscriptionIamBindingState, opts ...pulumi.ResourceOption) (*PubsubSubscriptionIamBinding, error) {
	var resource PubsubSubscriptionIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/pubsubSubscriptionIamBinding:PubsubSubscriptionIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PubsubSubscriptionIamBinding resources.
type pubsubSubscriptionIamBindingState struct {
	Condition                      *PubsubSubscriptionIamBindingCondition `pulumi:"condition"`
	Etag                           *string                                `pulumi:"etag"`
	Members                        []string                               `pulumi:"members"`
	Project                        *string                                `pulumi:"project"`
	PubsubSubscriptionIamBindingId *string                                `pulumi:"pubsubSubscriptionIamBindingId"`
	Role                           *string                                `pulumi:"role"`
	Subscription                   *string                                `pulumi:"subscription"`
}

type PubsubSubscriptionIamBindingState struct {
	Condition                      PubsubSubscriptionIamBindingConditionPtrInput
	Etag                           pulumi.StringPtrInput
	Members                        pulumi.StringArrayInput
	Project                        pulumi.StringPtrInput
	PubsubSubscriptionIamBindingId pulumi.StringPtrInput
	Role                           pulumi.StringPtrInput
	Subscription                   pulumi.StringPtrInput
}

func (PubsubSubscriptionIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*pubsubSubscriptionIamBindingState)(nil)).Elem()
}

type pubsubSubscriptionIamBindingArgs struct {
	Condition                      *PubsubSubscriptionIamBindingCondition `pulumi:"condition"`
	Members                        []string                               `pulumi:"members"`
	Project                        *string                                `pulumi:"project"`
	PubsubSubscriptionIamBindingId *string                                `pulumi:"pubsubSubscriptionIamBindingId"`
	Role                           string                                 `pulumi:"role"`
	Subscription                   string                                 `pulumi:"subscription"`
}

// The set of arguments for constructing a PubsubSubscriptionIamBinding resource.
type PubsubSubscriptionIamBindingArgs struct {
	Condition                      PubsubSubscriptionIamBindingConditionPtrInput
	Members                        pulumi.StringArrayInput
	Project                        pulumi.StringPtrInput
	PubsubSubscriptionIamBindingId pulumi.StringPtrInput
	Role                           pulumi.StringInput
	Subscription                   pulumi.StringInput
}

func (PubsubSubscriptionIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pubsubSubscriptionIamBindingArgs)(nil)).Elem()
}

type PubsubSubscriptionIamBindingInput interface {
	pulumi.Input

	ToPubsubSubscriptionIamBindingOutput() PubsubSubscriptionIamBindingOutput
	ToPubsubSubscriptionIamBindingOutputWithContext(ctx context.Context) PubsubSubscriptionIamBindingOutput
}

func (*PubsubSubscriptionIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**PubsubSubscriptionIamBinding)(nil)).Elem()
}

func (i *PubsubSubscriptionIamBinding) ToPubsubSubscriptionIamBindingOutput() PubsubSubscriptionIamBindingOutput {
	return i.ToPubsubSubscriptionIamBindingOutputWithContext(context.Background())
}

func (i *PubsubSubscriptionIamBinding) ToPubsubSubscriptionIamBindingOutputWithContext(ctx context.Context) PubsubSubscriptionIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PubsubSubscriptionIamBindingOutput)
}

type PubsubSubscriptionIamBindingOutput struct{ *pulumi.OutputState }

func (PubsubSubscriptionIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PubsubSubscriptionIamBinding)(nil)).Elem()
}

func (o PubsubSubscriptionIamBindingOutput) ToPubsubSubscriptionIamBindingOutput() PubsubSubscriptionIamBindingOutput {
	return o
}

func (o PubsubSubscriptionIamBindingOutput) ToPubsubSubscriptionIamBindingOutputWithContext(ctx context.Context) PubsubSubscriptionIamBindingOutput {
	return o
}

func (o PubsubSubscriptionIamBindingOutput) Condition() PubsubSubscriptionIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamBinding) PubsubSubscriptionIamBindingConditionPtrOutput {
		return v.Condition
	}).(PubsubSubscriptionIamBindingConditionPtrOutput)
}

func (o PubsubSubscriptionIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o PubsubSubscriptionIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamBindingOutput) PubsubSubscriptionIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamBinding) pulumi.StringOutput { return v.PubsubSubscriptionIamBindingId }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamBindingOutput) Subscription() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamBinding) pulumi.StringOutput { return v.Subscription }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PubsubSubscriptionIamBindingInput)(nil)).Elem(), &PubsubSubscriptionIamBinding{})
	pulumi.RegisterOutputType(PubsubSubscriptionIamBindingOutput{})
}
