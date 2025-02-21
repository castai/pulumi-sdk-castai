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

type SecureSourceManagerInstanceIamBinding struct {
	pulumi.CustomResourceState

	Condition                               SecureSourceManagerInstanceIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                                    pulumi.StringOutput                                     `pulumi:"etag"`
	InstanceId                              pulumi.StringOutput                                     `pulumi:"instanceId"`
	Location                                pulumi.StringOutput                                     `pulumi:"location"`
	Members                                 pulumi.StringArrayOutput                                `pulumi:"members"`
	Project                                 pulumi.StringOutput                                     `pulumi:"project"`
	Role                                    pulumi.StringOutput                                     `pulumi:"role"`
	SecureSourceManagerInstanceIamBindingId pulumi.StringOutput                                     `pulumi:"secureSourceManagerInstanceIamBindingId"`
}

// NewSecureSourceManagerInstanceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewSecureSourceManagerInstanceIamBinding(ctx *pulumi.Context,
	name string, args *SecureSourceManagerInstanceIamBindingArgs, opts ...pulumi.ResourceOption) (*SecureSourceManagerInstanceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
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
	var resource SecureSourceManagerInstanceIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/secureSourceManagerInstanceIamBinding:SecureSourceManagerInstanceIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecureSourceManagerInstanceIamBinding gets an existing SecureSourceManagerInstanceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecureSourceManagerInstanceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecureSourceManagerInstanceIamBindingState, opts ...pulumi.ResourceOption) (*SecureSourceManagerInstanceIamBinding, error) {
	var resource SecureSourceManagerInstanceIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/secureSourceManagerInstanceIamBinding:SecureSourceManagerInstanceIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecureSourceManagerInstanceIamBinding resources.
type secureSourceManagerInstanceIamBindingState struct {
	Condition                               *SecureSourceManagerInstanceIamBindingCondition `pulumi:"condition"`
	Etag                                    *string                                         `pulumi:"etag"`
	InstanceId                              *string                                         `pulumi:"instanceId"`
	Location                                *string                                         `pulumi:"location"`
	Members                                 []string                                        `pulumi:"members"`
	Project                                 *string                                         `pulumi:"project"`
	Role                                    *string                                         `pulumi:"role"`
	SecureSourceManagerInstanceIamBindingId *string                                         `pulumi:"secureSourceManagerInstanceIamBindingId"`
}

type SecureSourceManagerInstanceIamBindingState struct {
	Condition                               SecureSourceManagerInstanceIamBindingConditionPtrInput
	Etag                                    pulumi.StringPtrInput
	InstanceId                              pulumi.StringPtrInput
	Location                                pulumi.StringPtrInput
	Members                                 pulumi.StringArrayInput
	Project                                 pulumi.StringPtrInput
	Role                                    pulumi.StringPtrInput
	SecureSourceManagerInstanceIamBindingId pulumi.StringPtrInput
}

func (SecureSourceManagerInstanceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerInstanceIamBindingState)(nil)).Elem()
}

type secureSourceManagerInstanceIamBindingArgs struct {
	Condition                               *SecureSourceManagerInstanceIamBindingCondition `pulumi:"condition"`
	InstanceId                              string                                          `pulumi:"instanceId"`
	Location                                *string                                         `pulumi:"location"`
	Members                                 []string                                        `pulumi:"members"`
	Project                                 *string                                         `pulumi:"project"`
	Role                                    string                                          `pulumi:"role"`
	SecureSourceManagerInstanceIamBindingId *string                                         `pulumi:"secureSourceManagerInstanceIamBindingId"`
}

// The set of arguments for constructing a SecureSourceManagerInstanceIamBinding resource.
type SecureSourceManagerInstanceIamBindingArgs struct {
	Condition                               SecureSourceManagerInstanceIamBindingConditionPtrInput
	InstanceId                              pulumi.StringInput
	Location                                pulumi.StringPtrInput
	Members                                 pulumi.StringArrayInput
	Project                                 pulumi.StringPtrInput
	Role                                    pulumi.StringInput
	SecureSourceManagerInstanceIamBindingId pulumi.StringPtrInput
}

func (SecureSourceManagerInstanceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerInstanceIamBindingArgs)(nil)).Elem()
}

type SecureSourceManagerInstanceIamBindingInput interface {
	pulumi.Input

	ToSecureSourceManagerInstanceIamBindingOutput() SecureSourceManagerInstanceIamBindingOutput
	ToSecureSourceManagerInstanceIamBindingOutputWithContext(ctx context.Context) SecureSourceManagerInstanceIamBindingOutput
}

func (*SecureSourceManagerInstanceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerInstanceIamBinding)(nil)).Elem()
}

func (i *SecureSourceManagerInstanceIamBinding) ToSecureSourceManagerInstanceIamBindingOutput() SecureSourceManagerInstanceIamBindingOutput {
	return i.ToSecureSourceManagerInstanceIamBindingOutputWithContext(context.Background())
}

func (i *SecureSourceManagerInstanceIamBinding) ToSecureSourceManagerInstanceIamBindingOutputWithContext(ctx context.Context) SecureSourceManagerInstanceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureSourceManagerInstanceIamBindingOutput)
}

type SecureSourceManagerInstanceIamBindingOutput struct{ *pulumi.OutputState }

func (SecureSourceManagerInstanceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerInstanceIamBinding)(nil)).Elem()
}

func (o SecureSourceManagerInstanceIamBindingOutput) ToSecureSourceManagerInstanceIamBindingOutput() SecureSourceManagerInstanceIamBindingOutput {
	return o
}

func (o SecureSourceManagerInstanceIamBindingOutput) ToSecureSourceManagerInstanceIamBindingOutputWithContext(ctx context.Context) SecureSourceManagerInstanceIamBindingOutput {
	return o
}

func (o SecureSourceManagerInstanceIamBindingOutput) Condition() SecureSourceManagerInstanceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) SecureSourceManagerInstanceIamBindingConditionPtrOutput {
		return v.Condition
	}).(SecureSourceManagerInstanceIamBindingConditionPtrOutput)
}

func (o SecureSourceManagerInstanceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SecureSourceManagerInstanceIamBindingOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o SecureSourceManagerInstanceIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SecureSourceManagerInstanceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o SecureSourceManagerInstanceIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SecureSourceManagerInstanceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o SecureSourceManagerInstanceIamBindingOutput) SecureSourceManagerInstanceIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerInstanceIamBinding) pulumi.StringOutput {
		return v.SecureSourceManagerInstanceIamBindingId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecureSourceManagerInstanceIamBindingInput)(nil)).Elem(), &SecureSourceManagerInstanceIamBinding{})
	pulumi.RegisterOutputType(SecureSourceManagerInstanceIamBindingOutput{})
}
