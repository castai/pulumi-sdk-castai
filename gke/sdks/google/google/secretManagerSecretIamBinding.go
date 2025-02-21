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

type SecretManagerSecretIamBinding struct {
	pulumi.CustomResourceState

	Condition                       SecretManagerSecretIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                            pulumi.StringOutput                             `pulumi:"etag"`
	Members                         pulumi.StringArrayOutput                        `pulumi:"members"`
	Project                         pulumi.StringOutput                             `pulumi:"project"`
	Role                            pulumi.StringOutput                             `pulumi:"role"`
	SecretId                        pulumi.StringOutput                             `pulumi:"secretId"`
	SecretManagerSecretIamBindingId pulumi.StringOutput                             `pulumi:"secretManagerSecretIamBindingId"`
}

// NewSecretManagerSecretIamBinding registers a new resource with the given unique name, arguments, and options.
func NewSecretManagerSecretIamBinding(ctx *pulumi.Context,
	name string, args *SecretManagerSecretIamBindingArgs, opts ...pulumi.ResourceOption) (*SecretManagerSecretIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.SecretId == nil {
		return nil, errors.New("invalid value for required argument 'SecretId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SecretManagerSecretIamBinding
	err = ctx.RegisterPackageResource("google:index/secretManagerSecretIamBinding:SecretManagerSecretIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretManagerSecretIamBinding gets an existing SecretManagerSecretIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretManagerSecretIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretManagerSecretIamBindingState, opts ...pulumi.ResourceOption) (*SecretManagerSecretIamBinding, error) {
	var resource SecretManagerSecretIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/secretManagerSecretIamBinding:SecretManagerSecretIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretManagerSecretIamBinding resources.
type secretManagerSecretIamBindingState struct {
	Condition                       *SecretManagerSecretIamBindingCondition `pulumi:"condition"`
	Etag                            *string                                 `pulumi:"etag"`
	Members                         []string                                `pulumi:"members"`
	Project                         *string                                 `pulumi:"project"`
	Role                            *string                                 `pulumi:"role"`
	SecretId                        *string                                 `pulumi:"secretId"`
	SecretManagerSecretIamBindingId *string                                 `pulumi:"secretManagerSecretIamBindingId"`
}

type SecretManagerSecretIamBindingState struct {
	Condition                       SecretManagerSecretIamBindingConditionPtrInput
	Etag                            pulumi.StringPtrInput
	Members                         pulumi.StringArrayInput
	Project                         pulumi.StringPtrInput
	Role                            pulumi.StringPtrInput
	SecretId                        pulumi.StringPtrInput
	SecretManagerSecretIamBindingId pulumi.StringPtrInput
}

func (SecretManagerSecretIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretManagerSecretIamBindingState)(nil)).Elem()
}

type secretManagerSecretIamBindingArgs struct {
	Condition                       *SecretManagerSecretIamBindingCondition `pulumi:"condition"`
	Members                         []string                                `pulumi:"members"`
	Project                         *string                                 `pulumi:"project"`
	Role                            string                                  `pulumi:"role"`
	SecretId                        string                                  `pulumi:"secretId"`
	SecretManagerSecretIamBindingId *string                                 `pulumi:"secretManagerSecretIamBindingId"`
}

// The set of arguments for constructing a SecretManagerSecretIamBinding resource.
type SecretManagerSecretIamBindingArgs struct {
	Condition                       SecretManagerSecretIamBindingConditionPtrInput
	Members                         pulumi.StringArrayInput
	Project                         pulumi.StringPtrInput
	Role                            pulumi.StringInput
	SecretId                        pulumi.StringInput
	SecretManagerSecretIamBindingId pulumi.StringPtrInput
}

func (SecretManagerSecretIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretManagerSecretIamBindingArgs)(nil)).Elem()
}

type SecretManagerSecretIamBindingInput interface {
	pulumi.Input

	ToSecretManagerSecretIamBindingOutput() SecretManagerSecretIamBindingOutput
	ToSecretManagerSecretIamBindingOutputWithContext(ctx context.Context) SecretManagerSecretIamBindingOutput
}

func (*SecretManagerSecretIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretManagerSecretIamBinding)(nil)).Elem()
}

func (i *SecretManagerSecretIamBinding) ToSecretManagerSecretIamBindingOutput() SecretManagerSecretIamBindingOutput {
	return i.ToSecretManagerSecretIamBindingOutputWithContext(context.Background())
}

func (i *SecretManagerSecretIamBinding) ToSecretManagerSecretIamBindingOutputWithContext(ctx context.Context) SecretManagerSecretIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretManagerSecretIamBindingOutput)
}

type SecretManagerSecretIamBindingOutput struct{ *pulumi.OutputState }

func (SecretManagerSecretIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretManagerSecretIamBinding)(nil)).Elem()
}

func (o SecretManagerSecretIamBindingOutput) ToSecretManagerSecretIamBindingOutput() SecretManagerSecretIamBindingOutput {
	return o
}

func (o SecretManagerSecretIamBindingOutput) ToSecretManagerSecretIamBindingOutputWithContext(ctx context.Context) SecretManagerSecretIamBindingOutput {
	return o
}

func (o SecretManagerSecretIamBindingOutput) Condition() SecretManagerSecretIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *SecretManagerSecretIamBinding) SecretManagerSecretIamBindingConditionPtrOutput {
		return v.Condition
	}).(SecretManagerSecretIamBindingConditionPtrOutput)
}

func (o SecretManagerSecretIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerSecretIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SecretManagerSecretIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretManagerSecretIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o SecretManagerSecretIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerSecretIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SecretManagerSecretIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerSecretIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o SecretManagerSecretIamBindingOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerSecretIamBinding) pulumi.StringOutput { return v.SecretId }).(pulumi.StringOutput)
}

func (o SecretManagerSecretIamBindingOutput) SecretManagerSecretIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerSecretIamBinding) pulumi.StringOutput { return v.SecretManagerSecretIamBindingId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretManagerSecretIamBindingInput)(nil)).Elem(), &SecretManagerSecretIamBinding{})
	pulumi.RegisterOutputType(SecretManagerSecretIamBindingOutput{})
}
