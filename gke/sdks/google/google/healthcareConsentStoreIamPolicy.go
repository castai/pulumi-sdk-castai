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

type HealthcareConsentStoreIamPolicy struct {
	pulumi.CustomResourceState

	ConsentStoreId                    pulumi.StringOutput `pulumi:"consentStoreId"`
	Dataset                           pulumi.StringOutput `pulumi:"dataset"`
	Etag                              pulumi.StringOutput `pulumi:"etag"`
	HealthcareConsentStoreIamPolicyId pulumi.StringOutput `pulumi:"healthcareConsentStoreIamPolicyId"`
	PolicyData                        pulumi.StringOutput `pulumi:"policyData"`
}

// NewHealthcareConsentStoreIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewHealthcareConsentStoreIamPolicy(ctx *pulumi.Context,
	name string, args *HealthcareConsentStoreIamPolicyArgs, opts ...pulumi.ResourceOption) (*HealthcareConsentStoreIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource HealthcareConsentStoreIamPolicy
	err = ctx.RegisterPackageResource("google:index/healthcareConsentStoreIamPolicy:HealthcareConsentStoreIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthcareConsentStoreIamPolicy gets an existing HealthcareConsentStoreIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthcareConsentStoreIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthcareConsentStoreIamPolicyState, opts ...pulumi.ResourceOption) (*HealthcareConsentStoreIamPolicy, error) {
	var resource HealthcareConsentStoreIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/healthcareConsentStoreIamPolicy:HealthcareConsentStoreIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthcareConsentStoreIamPolicy resources.
type healthcareConsentStoreIamPolicyState struct {
	ConsentStoreId                    *string `pulumi:"consentStoreId"`
	Dataset                           *string `pulumi:"dataset"`
	Etag                              *string `pulumi:"etag"`
	HealthcareConsentStoreIamPolicyId *string `pulumi:"healthcareConsentStoreIamPolicyId"`
	PolicyData                        *string `pulumi:"policyData"`
}

type HealthcareConsentStoreIamPolicyState struct {
	ConsentStoreId                    pulumi.StringPtrInput
	Dataset                           pulumi.StringPtrInput
	Etag                              pulumi.StringPtrInput
	HealthcareConsentStoreIamPolicyId pulumi.StringPtrInput
	PolicyData                        pulumi.StringPtrInput
}

func (HealthcareConsentStoreIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcareConsentStoreIamPolicyState)(nil)).Elem()
}

type healthcareConsentStoreIamPolicyArgs struct {
	ConsentStoreId                    string  `pulumi:"consentStoreId"`
	Dataset                           string  `pulumi:"dataset"`
	HealthcareConsentStoreIamPolicyId *string `pulumi:"healthcareConsentStoreIamPolicyId"`
	PolicyData                        string  `pulumi:"policyData"`
}

// The set of arguments for constructing a HealthcareConsentStoreIamPolicy resource.
type HealthcareConsentStoreIamPolicyArgs struct {
	ConsentStoreId                    pulumi.StringInput
	Dataset                           pulumi.StringInput
	HealthcareConsentStoreIamPolicyId pulumi.StringPtrInput
	PolicyData                        pulumi.StringInput
}

func (HealthcareConsentStoreIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcareConsentStoreIamPolicyArgs)(nil)).Elem()
}

type HealthcareConsentStoreIamPolicyInput interface {
	pulumi.Input

	ToHealthcareConsentStoreIamPolicyOutput() HealthcareConsentStoreIamPolicyOutput
	ToHealthcareConsentStoreIamPolicyOutputWithContext(ctx context.Context) HealthcareConsentStoreIamPolicyOutput
}

func (*HealthcareConsentStoreIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcareConsentStoreIamPolicy)(nil)).Elem()
}

func (i *HealthcareConsentStoreIamPolicy) ToHealthcareConsentStoreIamPolicyOutput() HealthcareConsentStoreIamPolicyOutput {
	return i.ToHealthcareConsentStoreIamPolicyOutputWithContext(context.Background())
}

func (i *HealthcareConsentStoreIamPolicy) ToHealthcareConsentStoreIamPolicyOutputWithContext(ctx context.Context) HealthcareConsentStoreIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthcareConsentStoreIamPolicyOutput)
}

type HealthcareConsentStoreIamPolicyOutput struct{ *pulumi.OutputState }

func (HealthcareConsentStoreIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcareConsentStoreIamPolicy)(nil)).Elem()
}

func (o HealthcareConsentStoreIamPolicyOutput) ToHealthcareConsentStoreIamPolicyOutput() HealthcareConsentStoreIamPolicyOutput {
	return o
}

func (o HealthcareConsentStoreIamPolicyOutput) ToHealthcareConsentStoreIamPolicyOutputWithContext(ctx context.Context) HealthcareConsentStoreIamPolicyOutput {
	return o
}

func (o HealthcareConsentStoreIamPolicyOutput) ConsentStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamPolicy) pulumi.StringOutput { return v.ConsentStoreId }).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamPolicyOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamPolicy) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamPolicyOutput) HealthcareConsentStoreIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamPolicy) pulumi.StringOutput {
		return v.HealthcareConsentStoreIamPolicyId
	}).(pulumi.StringOutput)
}

func (o HealthcareConsentStoreIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareConsentStoreIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthcareConsentStoreIamPolicyInput)(nil)).Elem(), &HealthcareConsentStoreIamPolicy{})
	pulumi.RegisterOutputType(HealthcareConsentStoreIamPolicyOutput{})
}
