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

type SccV2OrganizationSourceIamPolicy struct {
	pulumi.CustomResourceState

	Etag                               pulumi.StringOutput `pulumi:"etag"`
	Organization                       pulumi.StringOutput `pulumi:"organization"`
	PolicyData                         pulumi.StringOutput `pulumi:"policyData"`
	SccV2OrganizationSourceIamPolicyId pulumi.StringOutput `pulumi:"sccV2OrganizationSourceIamPolicyId"`
	Source                             pulumi.StringOutput `pulumi:"source"`
}

// NewSccV2OrganizationSourceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewSccV2OrganizationSourceIamPolicy(ctx *pulumi.Context,
	name string, args *SccV2OrganizationSourceIamPolicyArgs, opts ...pulumi.ResourceOption) (*SccV2OrganizationSourceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccV2OrganizationSourceIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/sccV2OrganizationSourceIamPolicy:SccV2OrganizationSourceIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccV2OrganizationSourceIamPolicy gets an existing SccV2OrganizationSourceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccV2OrganizationSourceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccV2OrganizationSourceIamPolicyState, opts ...pulumi.ResourceOption) (*SccV2OrganizationSourceIamPolicy, error) {
	var resource SccV2OrganizationSourceIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/sccV2OrganizationSourceIamPolicy:SccV2OrganizationSourceIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccV2OrganizationSourceIamPolicy resources.
type sccV2OrganizationSourceIamPolicyState struct {
	Etag                               *string `pulumi:"etag"`
	Organization                       *string `pulumi:"organization"`
	PolicyData                         *string `pulumi:"policyData"`
	SccV2OrganizationSourceIamPolicyId *string `pulumi:"sccV2OrganizationSourceIamPolicyId"`
	Source                             *string `pulumi:"source"`
}

type SccV2OrganizationSourceIamPolicyState struct {
	Etag                               pulumi.StringPtrInput
	Organization                       pulumi.StringPtrInput
	PolicyData                         pulumi.StringPtrInput
	SccV2OrganizationSourceIamPolicyId pulumi.StringPtrInput
	Source                             pulumi.StringPtrInput
}

func (SccV2OrganizationSourceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2OrganizationSourceIamPolicyState)(nil)).Elem()
}

type sccV2OrganizationSourceIamPolicyArgs struct {
	Organization                       string  `pulumi:"organization"`
	PolicyData                         string  `pulumi:"policyData"`
	SccV2OrganizationSourceIamPolicyId *string `pulumi:"sccV2OrganizationSourceIamPolicyId"`
	Source                             string  `pulumi:"source"`
}

// The set of arguments for constructing a SccV2OrganizationSourceIamPolicy resource.
type SccV2OrganizationSourceIamPolicyArgs struct {
	Organization                       pulumi.StringInput
	PolicyData                         pulumi.StringInput
	SccV2OrganizationSourceIamPolicyId pulumi.StringPtrInput
	Source                             pulumi.StringInput
}

func (SccV2OrganizationSourceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2OrganizationSourceIamPolicyArgs)(nil)).Elem()
}

type SccV2OrganizationSourceIamPolicyInput interface {
	pulumi.Input

	ToSccV2OrganizationSourceIamPolicyOutput() SccV2OrganizationSourceIamPolicyOutput
	ToSccV2OrganizationSourceIamPolicyOutputWithContext(ctx context.Context) SccV2OrganizationSourceIamPolicyOutput
}

func (*SccV2OrganizationSourceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2OrganizationSourceIamPolicy)(nil)).Elem()
}

func (i *SccV2OrganizationSourceIamPolicy) ToSccV2OrganizationSourceIamPolicyOutput() SccV2OrganizationSourceIamPolicyOutput {
	return i.ToSccV2OrganizationSourceIamPolicyOutputWithContext(context.Background())
}

func (i *SccV2OrganizationSourceIamPolicy) ToSccV2OrganizationSourceIamPolicyOutputWithContext(ctx context.Context) SccV2OrganizationSourceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccV2OrganizationSourceIamPolicyOutput)
}

type SccV2OrganizationSourceIamPolicyOutput struct{ *pulumi.OutputState }

func (SccV2OrganizationSourceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2OrganizationSourceIamPolicy)(nil)).Elem()
}

func (o SccV2OrganizationSourceIamPolicyOutput) ToSccV2OrganizationSourceIamPolicyOutput() SccV2OrganizationSourceIamPolicyOutput {
	return o
}

func (o SccV2OrganizationSourceIamPolicyOutput) ToSccV2OrganizationSourceIamPolicyOutputWithContext(ctx context.Context) SccV2OrganizationSourceIamPolicyOutput {
	return o
}

func (o SccV2OrganizationSourceIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationSourceIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SccV2OrganizationSourceIamPolicyOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationSourceIamPolicy) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

func (o SccV2OrganizationSourceIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationSourceIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o SccV2OrganizationSourceIamPolicyOutput) SccV2OrganizationSourceIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationSourceIamPolicy) pulumi.StringOutput {
		return v.SccV2OrganizationSourceIamPolicyId
	}).(pulumi.StringOutput)
}

func (o SccV2OrganizationSourceIamPolicyOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationSourceIamPolicy) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccV2OrganizationSourceIamPolicyInput)(nil)).Elem(), &SccV2OrganizationSourceIamPolicy{})
	pulumi.RegisterOutputType(SccV2OrganizationSourceIamPolicyOutput{})
}
