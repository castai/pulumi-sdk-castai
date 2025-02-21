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

type ComputeRegionBackendServiceIamPolicy struct {
	pulumi.CustomResourceState

	ComputeRegionBackendServiceIamPolicyId pulumi.StringOutput `pulumi:"computeRegionBackendServiceIamPolicyId"`
	Etag                                   pulumi.StringOutput `pulumi:"etag"`
	Name                                   pulumi.StringOutput `pulumi:"name"`
	PolicyData                             pulumi.StringOutput `pulumi:"policyData"`
	Project                                pulumi.StringOutput `pulumi:"project"`
	Region                                 pulumi.StringOutput `pulumi:"region"`
}

// NewComputeRegionBackendServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewComputeRegionBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, args *ComputeRegionBackendServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*ComputeRegionBackendServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeRegionBackendServiceIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/computeRegionBackendServiceIamPolicy:ComputeRegionBackendServiceIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRegionBackendServiceIamPolicy gets an existing ComputeRegionBackendServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRegionBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRegionBackendServiceIamPolicyState, opts ...pulumi.ResourceOption) (*ComputeRegionBackendServiceIamPolicy, error) {
	var resource ComputeRegionBackendServiceIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeRegionBackendServiceIamPolicy:ComputeRegionBackendServiceIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRegionBackendServiceIamPolicy resources.
type computeRegionBackendServiceIamPolicyState struct {
	ComputeRegionBackendServiceIamPolicyId *string `pulumi:"computeRegionBackendServiceIamPolicyId"`
	Etag                                   *string `pulumi:"etag"`
	Name                                   *string `pulumi:"name"`
	PolicyData                             *string `pulumi:"policyData"`
	Project                                *string `pulumi:"project"`
	Region                                 *string `pulumi:"region"`
}

type ComputeRegionBackendServiceIamPolicyState struct {
	ComputeRegionBackendServiceIamPolicyId pulumi.StringPtrInput
	Etag                                   pulumi.StringPtrInput
	Name                                   pulumi.StringPtrInput
	PolicyData                             pulumi.StringPtrInput
	Project                                pulumi.StringPtrInput
	Region                                 pulumi.StringPtrInput
}

func (ComputeRegionBackendServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionBackendServiceIamPolicyState)(nil)).Elem()
}

type computeRegionBackendServiceIamPolicyArgs struct {
	ComputeRegionBackendServiceIamPolicyId *string `pulumi:"computeRegionBackendServiceIamPolicyId"`
	Name                                   *string `pulumi:"name"`
	PolicyData                             string  `pulumi:"policyData"`
	Project                                *string `pulumi:"project"`
	Region                                 *string `pulumi:"region"`
}

// The set of arguments for constructing a ComputeRegionBackendServiceIamPolicy resource.
type ComputeRegionBackendServiceIamPolicyArgs struct {
	ComputeRegionBackendServiceIamPolicyId pulumi.StringPtrInput
	Name                                   pulumi.StringPtrInput
	PolicyData                             pulumi.StringInput
	Project                                pulumi.StringPtrInput
	Region                                 pulumi.StringPtrInput
}

func (ComputeRegionBackendServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionBackendServiceIamPolicyArgs)(nil)).Elem()
}

type ComputeRegionBackendServiceIamPolicyInput interface {
	pulumi.Input

	ToComputeRegionBackendServiceIamPolicyOutput() ComputeRegionBackendServiceIamPolicyOutput
	ToComputeRegionBackendServiceIamPolicyOutputWithContext(ctx context.Context) ComputeRegionBackendServiceIamPolicyOutput
}

func (*ComputeRegionBackendServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionBackendServiceIamPolicy)(nil)).Elem()
}

func (i *ComputeRegionBackendServiceIamPolicy) ToComputeRegionBackendServiceIamPolicyOutput() ComputeRegionBackendServiceIamPolicyOutput {
	return i.ToComputeRegionBackendServiceIamPolicyOutputWithContext(context.Background())
}

func (i *ComputeRegionBackendServiceIamPolicy) ToComputeRegionBackendServiceIamPolicyOutputWithContext(ctx context.Context) ComputeRegionBackendServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRegionBackendServiceIamPolicyOutput)
}

type ComputeRegionBackendServiceIamPolicyOutput struct{ *pulumi.OutputState }

func (ComputeRegionBackendServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionBackendServiceIamPolicy)(nil)).Elem()
}

func (o ComputeRegionBackendServiceIamPolicyOutput) ToComputeRegionBackendServiceIamPolicyOutput() ComputeRegionBackendServiceIamPolicyOutput {
	return o
}

func (o ComputeRegionBackendServiceIamPolicyOutput) ToComputeRegionBackendServiceIamPolicyOutputWithContext(ctx context.Context) ComputeRegionBackendServiceIamPolicyOutput {
	return o
}

func (o ComputeRegionBackendServiceIamPolicyOutput) ComputeRegionBackendServiceIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionBackendServiceIamPolicy) pulumi.StringOutput {
		return v.ComputeRegionBackendServiceIamPolicyId
	}).(pulumi.StringOutput)
}

func (o ComputeRegionBackendServiceIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionBackendServiceIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ComputeRegionBackendServiceIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionBackendServiceIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeRegionBackendServiceIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionBackendServiceIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o ComputeRegionBackendServiceIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionBackendServiceIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeRegionBackendServiceIamPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionBackendServiceIamPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRegionBackendServiceIamPolicyInput)(nil)).Elem(), &ComputeRegionBackendServiceIamPolicy{})
	pulumi.RegisterOutputType(ComputeRegionBackendServiceIamPolicyOutput{})
}
