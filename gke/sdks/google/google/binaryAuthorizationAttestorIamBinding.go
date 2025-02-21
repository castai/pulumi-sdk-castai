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

type BinaryAuthorizationAttestorIamBinding struct {
	pulumi.CustomResourceState

	Attestor                                pulumi.StringOutput                                     `pulumi:"attestor"`
	BinaryAuthorizationAttestorIamBindingId pulumi.StringOutput                                     `pulumi:"binaryAuthorizationAttestorIamBindingId"`
	Condition                               BinaryAuthorizationAttestorIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                                    pulumi.StringOutput                                     `pulumi:"etag"`
	Members                                 pulumi.StringArrayOutput                                `pulumi:"members"`
	Project                                 pulumi.StringOutput                                     `pulumi:"project"`
	Role                                    pulumi.StringOutput                                     `pulumi:"role"`
}

// NewBinaryAuthorizationAttestorIamBinding registers a new resource with the given unique name, arguments, and options.
func NewBinaryAuthorizationAttestorIamBinding(ctx *pulumi.Context,
	name string, args *BinaryAuthorizationAttestorIamBindingArgs, opts ...pulumi.ResourceOption) (*BinaryAuthorizationAttestorIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Attestor == nil {
		return nil, errors.New("invalid value for required argument 'Attestor'")
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
	var resource BinaryAuthorizationAttestorIamBinding
	err = ctx.RegisterPackageResource("google:index/binaryAuthorizationAttestorIamBinding:BinaryAuthorizationAttestorIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinaryAuthorizationAttestorIamBinding gets an existing BinaryAuthorizationAttestorIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinaryAuthorizationAttestorIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BinaryAuthorizationAttestorIamBindingState, opts ...pulumi.ResourceOption) (*BinaryAuthorizationAttestorIamBinding, error) {
	var resource BinaryAuthorizationAttestorIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/binaryAuthorizationAttestorIamBinding:BinaryAuthorizationAttestorIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BinaryAuthorizationAttestorIamBinding resources.
type binaryAuthorizationAttestorIamBindingState struct {
	Attestor                                *string                                         `pulumi:"attestor"`
	BinaryAuthorizationAttestorIamBindingId *string                                         `pulumi:"binaryAuthorizationAttestorIamBindingId"`
	Condition                               *BinaryAuthorizationAttestorIamBindingCondition `pulumi:"condition"`
	Etag                                    *string                                         `pulumi:"etag"`
	Members                                 []string                                        `pulumi:"members"`
	Project                                 *string                                         `pulumi:"project"`
	Role                                    *string                                         `pulumi:"role"`
}

type BinaryAuthorizationAttestorIamBindingState struct {
	Attestor                                pulumi.StringPtrInput
	BinaryAuthorizationAttestorIamBindingId pulumi.StringPtrInput
	Condition                               BinaryAuthorizationAttestorIamBindingConditionPtrInput
	Etag                                    pulumi.StringPtrInput
	Members                                 pulumi.StringArrayInput
	Project                                 pulumi.StringPtrInput
	Role                                    pulumi.StringPtrInput
}

func (BinaryAuthorizationAttestorIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*binaryAuthorizationAttestorIamBindingState)(nil)).Elem()
}

type binaryAuthorizationAttestorIamBindingArgs struct {
	Attestor                                string                                          `pulumi:"attestor"`
	BinaryAuthorizationAttestorIamBindingId *string                                         `pulumi:"binaryAuthorizationAttestorIamBindingId"`
	Condition                               *BinaryAuthorizationAttestorIamBindingCondition `pulumi:"condition"`
	Members                                 []string                                        `pulumi:"members"`
	Project                                 *string                                         `pulumi:"project"`
	Role                                    string                                          `pulumi:"role"`
}

// The set of arguments for constructing a BinaryAuthorizationAttestorIamBinding resource.
type BinaryAuthorizationAttestorIamBindingArgs struct {
	Attestor                                pulumi.StringInput
	BinaryAuthorizationAttestorIamBindingId pulumi.StringPtrInput
	Condition                               BinaryAuthorizationAttestorIamBindingConditionPtrInput
	Members                                 pulumi.StringArrayInput
	Project                                 pulumi.StringPtrInput
	Role                                    pulumi.StringInput
}

func (BinaryAuthorizationAttestorIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*binaryAuthorizationAttestorIamBindingArgs)(nil)).Elem()
}

type BinaryAuthorizationAttestorIamBindingInput interface {
	pulumi.Input

	ToBinaryAuthorizationAttestorIamBindingOutput() BinaryAuthorizationAttestorIamBindingOutput
	ToBinaryAuthorizationAttestorIamBindingOutputWithContext(ctx context.Context) BinaryAuthorizationAttestorIamBindingOutput
}

func (*BinaryAuthorizationAttestorIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**BinaryAuthorizationAttestorIamBinding)(nil)).Elem()
}

func (i *BinaryAuthorizationAttestorIamBinding) ToBinaryAuthorizationAttestorIamBindingOutput() BinaryAuthorizationAttestorIamBindingOutput {
	return i.ToBinaryAuthorizationAttestorIamBindingOutputWithContext(context.Background())
}

func (i *BinaryAuthorizationAttestorIamBinding) ToBinaryAuthorizationAttestorIamBindingOutputWithContext(ctx context.Context) BinaryAuthorizationAttestorIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BinaryAuthorizationAttestorIamBindingOutput)
}

type BinaryAuthorizationAttestorIamBindingOutput struct{ *pulumi.OutputState }

func (BinaryAuthorizationAttestorIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BinaryAuthorizationAttestorIamBinding)(nil)).Elem()
}

func (o BinaryAuthorizationAttestorIamBindingOutput) ToBinaryAuthorizationAttestorIamBindingOutput() BinaryAuthorizationAttestorIamBindingOutput {
	return o
}

func (o BinaryAuthorizationAttestorIamBindingOutput) ToBinaryAuthorizationAttestorIamBindingOutputWithContext(ctx context.Context) BinaryAuthorizationAttestorIamBindingOutput {
	return o
}

func (o BinaryAuthorizationAttestorIamBindingOutput) Attestor() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamBinding) pulumi.StringOutput { return v.Attestor }).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamBindingOutput) BinaryAuthorizationAttestorIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamBinding) pulumi.StringOutput {
		return v.BinaryAuthorizationAttestorIamBindingId
	}).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamBindingOutput) Condition() BinaryAuthorizationAttestorIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamBinding) BinaryAuthorizationAttestorIamBindingConditionPtrOutput {
		return v.Condition
	}).(BinaryAuthorizationAttestorIamBindingConditionPtrOutput)
}

func (o BinaryAuthorizationAttestorIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o BinaryAuthorizationAttestorIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o BinaryAuthorizationAttestorIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BinaryAuthorizationAttestorIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BinaryAuthorizationAttestorIamBindingInput)(nil)).Elem(), &BinaryAuthorizationAttestorIamBinding{})
	pulumi.RegisterOutputType(BinaryAuthorizationAttestorIamBindingOutput{})
}
