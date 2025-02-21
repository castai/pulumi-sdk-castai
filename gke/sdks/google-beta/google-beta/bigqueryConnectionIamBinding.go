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

type BigqueryConnectionIamBinding struct {
	pulumi.CustomResourceState

	BigqueryConnectionIamBindingId pulumi.StringOutput                            `pulumi:"bigqueryConnectionIamBindingId"`
	Condition                      BigqueryConnectionIamBindingConditionPtrOutput `pulumi:"condition"`
	ConnectionId                   pulumi.StringOutput                            `pulumi:"connectionId"`
	Etag                           pulumi.StringOutput                            `pulumi:"etag"`
	Location                       pulumi.StringOutput                            `pulumi:"location"`
	Members                        pulumi.StringArrayOutput                       `pulumi:"members"`
	Project                        pulumi.StringOutput                            `pulumi:"project"`
	Role                           pulumi.StringOutput                            `pulumi:"role"`
}

// NewBigqueryConnectionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewBigqueryConnectionIamBinding(ctx *pulumi.Context,
	name string, args *BigqueryConnectionIamBindingArgs, opts ...pulumi.ResourceOption) (*BigqueryConnectionIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigqueryConnectionIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/bigqueryConnectionIamBinding:BigqueryConnectionIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryConnectionIamBinding gets an existing BigqueryConnectionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryConnectionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryConnectionIamBindingState, opts ...pulumi.ResourceOption) (*BigqueryConnectionIamBinding, error) {
	var resource BigqueryConnectionIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/bigqueryConnectionIamBinding:BigqueryConnectionIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryConnectionIamBinding resources.
type bigqueryConnectionIamBindingState struct {
	BigqueryConnectionIamBindingId *string                                `pulumi:"bigqueryConnectionIamBindingId"`
	Condition                      *BigqueryConnectionIamBindingCondition `pulumi:"condition"`
	ConnectionId                   *string                                `pulumi:"connectionId"`
	Etag                           *string                                `pulumi:"etag"`
	Location                       *string                                `pulumi:"location"`
	Members                        []string                               `pulumi:"members"`
	Project                        *string                                `pulumi:"project"`
	Role                           *string                                `pulumi:"role"`
}

type BigqueryConnectionIamBindingState struct {
	BigqueryConnectionIamBindingId pulumi.StringPtrInput
	Condition                      BigqueryConnectionIamBindingConditionPtrInput
	ConnectionId                   pulumi.StringPtrInput
	Etag                           pulumi.StringPtrInput
	Location                       pulumi.StringPtrInput
	Members                        pulumi.StringArrayInput
	Project                        pulumi.StringPtrInput
	Role                           pulumi.StringPtrInput
}

func (BigqueryConnectionIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryConnectionIamBindingState)(nil)).Elem()
}

type bigqueryConnectionIamBindingArgs struct {
	BigqueryConnectionIamBindingId *string                                `pulumi:"bigqueryConnectionIamBindingId"`
	Condition                      *BigqueryConnectionIamBindingCondition `pulumi:"condition"`
	ConnectionId                   string                                 `pulumi:"connectionId"`
	Location                       *string                                `pulumi:"location"`
	Members                        []string                               `pulumi:"members"`
	Project                        *string                                `pulumi:"project"`
	Role                           string                                 `pulumi:"role"`
}

// The set of arguments for constructing a BigqueryConnectionIamBinding resource.
type BigqueryConnectionIamBindingArgs struct {
	BigqueryConnectionIamBindingId pulumi.StringPtrInput
	Condition                      BigqueryConnectionIamBindingConditionPtrInput
	ConnectionId                   pulumi.StringInput
	Location                       pulumi.StringPtrInput
	Members                        pulumi.StringArrayInput
	Project                        pulumi.StringPtrInput
	Role                           pulumi.StringInput
}

func (BigqueryConnectionIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryConnectionIamBindingArgs)(nil)).Elem()
}

type BigqueryConnectionIamBindingInput interface {
	pulumi.Input

	ToBigqueryConnectionIamBindingOutput() BigqueryConnectionIamBindingOutput
	ToBigqueryConnectionIamBindingOutputWithContext(ctx context.Context) BigqueryConnectionIamBindingOutput
}

func (*BigqueryConnectionIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryConnectionIamBinding)(nil)).Elem()
}

func (i *BigqueryConnectionIamBinding) ToBigqueryConnectionIamBindingOutput() BigqueryConnectionIamBindingOutput {
	return i.ToBigqueryConnectionIamBindingOutputWithContext(context.Background())
}

func (i *BigqueryConnectionIamBinding) ToBigqueryConnectionIamBindingOutputWithContext(ctx context.Context) BigqueryConnectionIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryConnectionIamBindingOutput)
}

type BigqueryConnectionIamBindingOutput struct{ *pulumi.OutputState }

func (BigqueryConnectionIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryConnectionIamBinding)(nil)).Elem()
}

func (o BigqueryConnectionIamBindingOutput) ToBigqueryConnectionIamBindingOutput() BigqueryConnectionIamBindingOutput {
	return o
}

func (o BigqueryConnectionIamBindingOutput) ToBigqueryConnectionIamBindingOutputWithContext(ctx context.Context) BigqueryConnectionIamBindingOutput {
	return o
}

func (o BigqueryConnectionIamBindingOutput) BigqueryConnectionIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) pulumi.StringOutput { return v.BigqueryConnectionIamBindingId }).(pulumi.StringOutput)
}

func (o BigqueryConnectionIamBindingOutput) Condition() BigqueryConnectionIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) BigqueryConnectionIamBindingConditionPtrOutput {
		return v.Condition
	}).(BigqueryConnectionIamBindingConditionPtrOutput)
}

func (o BigqueryConnectionIamBindingOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

func (o BigqueryConnectionIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BigqueryConnectionIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BigqueryConnectionIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o BigqueryConnectionIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o BigqueryConnectionIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryConnectionIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryConnectionIamBindingInput)(nil)).Elem(), &BigqueryConnectionIamBinding{})
	pulumi.RegisterOutputType(BigqueryConnectionIamBindingOutput{})
}
