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

type IapWebTypeAppEngineIamBinding struct {
	pulumi.CustomResourceState

	AppId                           pulumi.StringOutput                             `pulumi:"appId"`
	Condition                       IapWebTypeAppEngineIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                            pulumi.StringOutput                             `pulumi:"etag"`
	IapWebTypeAppEngineIamBindingId pulumi.StringOutput                             `pulumi:"iapWebTypeAppEngineIamBindingId"`
	Members                         pulumi.StringArrayOutput                        `pulumi:"members"`
	Project                         pulumi.StringOutput                             `pulumi:"project"`
	Role                            pulumi.StringOutput                             `pulumi:"role"`
}

// NewIapWebTypeAppEngineIamBinding registers a new resource with the given unique name, arguments, and options.
func NewIapWebTypeAppEngineIamBinding(ctx *pulumi.Context,
	name string, args *IapWebTypeAppEngineIamBindingArgs, opts ...pulumi.ResourceOption) (*IapWebTypeAppEngineIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
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
	var resource IapWebTypeAppEngineIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/iapWebTypeAppEngineIamBinding:IapWebTypeAppEngineIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIapWebTypeAppEngineIamBinding gets an existing IapWebTypeAppEngineIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIapWebTypeAppEngineIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IapWebTypeAppEngineIamBindingState, opts ...pulumi.ResourceOption) (*IapWebTypeAppEngineIamBinding, error) {
	var resource IapWebTypeAppEngineIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/iapWebTypeAppEngineIamBinding:IapWebTypeAppEngineIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IapWebTypeAppEngineIamBinding resources.
type iapWebTypeAppEngineIamBindingState struct {
	AppId                           *string                                 `pulumi:"appId"`
	Condition                       *IapWebTypeAppEngineIamBindingCondition `pulumi:"condition"`
	Etag                            *string                                 `pulumi:"etag"`
	IapWebTypeAppEngineIamBindingId *string                                 `pulumi:"iapWebTypeAppEngineIamBindingId"`
	Members                         []string                                `pulumi:"members"`
	Project                         *string                                 `pulumi:"project"`
	Role                            *string                                 `pulumi:"role"`
}

type IapWebTypeAppEngineIamBindingState struct {
	AppId                           pulumi.StringPtrInput
	Condition                       IapWebTypeAppEngineIamBindingConditionPtrInput
	Etag                            pulumi.StringPtrInput
	IapWebTypeAppEngineIamBindingId pulumi.StringPtrInput
	Members                         pulumi.StringArrayInput
	Project                         pulumi.StringPtrInput
	Role                            pulumi.StringPtrInput
}

func (IapWebTypeAppEngineIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iapWebTypeAppEngineIamBindingState)(nil)).Elem()
}

type iapWebTypeAppEngineIamBindingArgs struct {
	AppId                           string                                  `pulumi:"appId"`
	Condition                       *IapWebTypeAppEngineIamBindingCondition `pulumi:"condition"`
	IapWebTypeAppEngineIamBindingId *string                                 `pulumi:"iapWebTypeAppEngineIamBindingId"`
	Members                         []string                                `pulumi:"members"`
	Project                         *string                                 `pulumi:"project"`
	Role                            string                                  `pulumi:"role"`
}

// The set of arguments for constructing a IapWebTypeAppEngineIamBinding resource.
type IapWebTypeAppEngineIamBindingArgs struct {
	AppId                           pulumi.StringInput
	Condition                       IapWebTypeAppEngineIamBindingConditionPtrInput
	IapWebTypeAppEngineIamBindingId pulumi.StringPtrInput
	Members                         pulumi.StringArrayInput
	Project                         pulumi.StringPtrInput
	Role                            pulumi.StringInput
}

func (IapWebTypeAppEngineIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iapWebTypeAppEngineIamBindingArgs)(nil)).Elem()
}

type IapWebTypeAppEngineIamBindingInput interface {
	pulumi.Input

	ToIapWebTypeAppEngineIamBindingOutput() IapWebTypeAppEngineIamBindingOutput
	ToIapWebTypeAppEngineIamBindingOutputWithContext(ctx context.Context) IapWebTypeAppEngineIamBindingOutput
}

func (*IapWebTypeAppEngineIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**IapWebTypeAppEngineIamBinding)(nil)).Elem()
}

func (i *IapWebTypeAppEngineIamBinding) ToIapWebTypeAppEngineIamBindingOutput() IapWebTypeAppEngineIamBindingOutput {
	return i.ToIapWebTypeAppEngineIamBindingOutputWithContext(context.Background())
}

func (i *IapWebTypeAppEngineIamBinding) ToIapWebTypeAppEngineIamBindingOutputWithContext(ctx context.Context) IapWebTypeAppEngineIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IapWebTypeAppEngineIamBindingOutput)
}

type IapWebTypeAppEngineIamBindingOutput struct{ *pulumi.OutputState }

func (IapWebTypeAppEngineIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IapWebTypeAppEngineIamBinding)(nil)).Elem()
}

func (o IapWebTypeAppEngineIamBindingOutput) ToIapWebTypeAppEngineIamBindingOutput() IapWebTypeAppEngineIamBindingOutput {
	return o
}

func (o IapWebTypeAppEngineIamBindingOutput) ToIapWebTypeAppEngineIamBindingOutputWithContext(ctx context.Context) IapWebTypeAppEngineIamBindingOutput {
	return o
}

func (o IapWebTypeAppEngineIamBindingOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebTypeAppEngineIamBinding) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o IapWebTypeAppEngineIamBindingOutput) Condition() IapWebTypeAppEngineIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *IapWebTypeAppEngineIamBinding) IapWebTypeAppEngineIamBindingConditionPtrOutput {
		return v.Condition
	}).(IapWebTypeAppEngineIamBindingConditionPtrOutput)
}

func (o IapWebTypeAppEngineIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebTypeAppEngineIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IapWebTypeAppEngineIamBindingOutput) IapWebTypeAppEngineIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebTypeAppEngineIamBinding) pulumi.StringOutput { return v.IapWebTypeAppEngineIamBindingId }).(pulumi.StringOutput)
}

func (o IapWebTypeAppEngineIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IapWebTypeAppEngineIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o IapWebTypeAppEngineIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebTypeAppEngineIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o IapWebTypeAppEngineIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IapWebTypeAppEngineIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IapWebTypeAppEngineIamBindingInput)(nil)).Elem(), &IapWebTypeAppEngineIamBinding{})
	pulumi.RegisterOutputType(IapWebTypeAppEngineIamBindingOutput{})
}
