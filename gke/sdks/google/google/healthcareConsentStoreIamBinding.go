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

type HealthcareConsentStoreIamBinding struct {
	pulumi.CustomResourceState

	Condition                          HealthcareConsentStoreIamBindingConditionPtrOutput `pulumi:"condition"`
	ConsentStoreId                     pulumi.StringOutput                                `pulumi:"consentStoreId"`
	Dataset                            pulumi.StringOutput                                `pulumi:"dataset"`
	Etag                               pulumi.StringOutput                                `pulumi:"etag"`
	HealthcareConsentStoreIamBindingId pulumi.StringOutput                                `pulumi:"healthcareConsentStoreIamBindingId"`
	Members                            pulumi.StringArrayOutput                           `pulumi:"members"`
	Role                               pulumi.StringOutput                                `pulumi:"role"`
}

// NewHealthcareConsentStoreIamBinding registers a new resource with the given unique name, arguments, and options.
func NewHealthcareConsentStoreIamBinding(ctx *pulumi.Context,
	name string, args *HealthcareConsentStoreIamBindingArgs, opts ...pulumi.ResourceOption) (*HealthcareConsentStoreIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
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
	var resource HealthcareConsentStoreIamBinding
	err = ctx.RegisterPackageResource("google:index/healthcareConsentStoreIamBinding:HealthcareConsentStoreIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthcareConsentStoreIamBinding gets an existing HealthcareConsentStoreIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthcareConsentStoreIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthcareConsentStoreIamBindingState, opts ...pulumi.ResourceOption) (*HealthcareConsentStoreIamBinding, error) {
	var resource HealthcareConsentStoreIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/healthcareConsentStoreIamBinding:HealthcareConsentStoreIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthcareConsentStoreIamBinding resources.
type healthcareConsentStoreIamBindingState struct {
	Condition                          *HealthcareConsentStoreIamBindingCondition `pulumi:"condition"`
	ConsentStoreId                     *string                                    `pulumi:"consentStoreId"`
	Dataset                            *string                                    `pulumi:"dataset"`
	Etag                               *string                                    `pulumi:"etag"`
	HealthcareConsentStoreIamBindingId *string                                    `pulumi:"healthcareConsentStoreIamBindingId"`
	Members                            []string                                   `pulumi:"members"`
	Role                               *string                                    `pulumi:"role"`
}

type HealthcareConsentStoreIamBindingState struct {
	Condition                          HealthcareConsentStoreIamBindingConditionPtrInput
	ConsentStoreId                     pulumi.StringPtrInput
	Dataset                            pulumi.StringPtrInput
	Etag                               pulumi.StringPtrInput
	HealthcareConsentStoreIamBindingId pulumi.StringPtrInput
	Members                            pulumi.StringArrayInput
	Role                               pulumi.StringPtrInput
}

func (HealthcareConsentStoreIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcareConsentStoreIamBindingState)(nil)).Elem()
}

type healthcareConsentStoreIamBindingArgs struct {
	Condition                          *HealthcareConsentStoreIamBindingCondition `pulumi:"condition"`
	ConsentStoreId                     string                                     `pulumi:"consentStoreId"`
	Dataset                            string                                     `pulumi:"dataset"`
	HealthcareConsentStoreIamBindingId *string                                    `pulumi:"healthcareConsentStoreIamBindingId"`
	Members                            []string                                   `pulumi:"members"`
	Role                               string                                     `pulumi:"role"`
}

// The set of arguments for constructing a HealthcareConsentStoreIamBinding resource.
type HealthcareConsentStoreIamBindingArgs struct {
	Condition                          HealthcareConsentStoreIamBindingConditionPtrInput
	ConsentStoreId                     pulumi.StringInput
	Dataset                            pulumi.StringInput
	HealthcareConsentStoreIamBindingId pulumi.StringPtrInput
	Members                            pulumi.StringArrayInput
	Role                               pulumi.StringInput
}

func (HealthcareConsentStoreIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcareConsentStoreIamBindingArgs)(nil)).Elem()
}

type HealthcareConsentStoreIamBindingInput interface {
	pulumi.Input

	ToHealthcareConsentStoreIamBindingOutput() HealthcareConsentStoreIamBindingOutput
	ToHealthcareConsentStoreIamBindingOutputWithContext(ctx context.Context) HealthcareConsentStoreIamBindingOutput
}

func (*HealthcareConsentStoreIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcareConsentStoreIamBinding)(nil)).Elem()
}

func (i *HealthcareConsentStoreIamBinding) ToHealthcareConsentStoreIamBindingOutput() HealthcareConsentStoreIamBindingOutput {
	return i.ToHealthcareConsentStoreIamBindingOutputWithContext(context.Background())
}

func (i *HealthcareConsentStoreIamBinding) ToHealthcareConsentStoreIamBindingOutputWithContext(ctx context.Context) HealthcareConsentStoreIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthcareConsentStoreIamBindingOutput)
}

type HealthcareConsentStoreIamBindingOutput struct{ *pulumi.OutputState }

func (HealthcareConsentStoreIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcareConsentStoreIamBinding)(nil)).Elem()
}

func (o HealthcareConsentStoreIamBindingOutput) ToHealthcareConsentStoreIamBindingOutput() HealthcareConsentStoreIamBindingOutput {
	return o
}

func (o HealthcareConsentStoreIamBindingOutput) ToHealthcareConsentStoreIamBindingOutputWithContext(ctx context.Context) HealthcareConsentStoreIamBindingOutput {
	return o
}

func (o HealthcareConsentStoreIamBindingOutput) Condition() HealthcareConsentStoreIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamBinding) HealthcareConsentStoreIamBindingConditionPtrOutput {
		return v.Condition
	}).(HealthcareConsentStoreIamBindingConditionPtrOutput)
}

func (o HealthcareConsentStoreIamBindingOutput) ConsentStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamBinding) pulumi.StringOutput { return v.ConsentStoreId }).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamBindingOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamBinding) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamBindingOutput) HealthcareConsentStoreIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamBinding) pulumi.StringOutput {
		return v.HealthcareConsentStoreIamBindingId
	}).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o HealthcareConsentStoreIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthcareConsentStoreIamBindingInput)(nil)).Elem(), &HealthcareConsentStoreIamBinding{})
	pulumi.RegisterOutputType(HealthcareConsentStoreIamBindingOutput{})
}
