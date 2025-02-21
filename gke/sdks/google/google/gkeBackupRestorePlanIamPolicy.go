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

type GkeBackupRestorePlanIamPolicy struct {
	pulumi.CustomResourceState

	Etag                            pulumi.StringOutput `pulumi:"etag"`
	GkeBackupRestorePlanIamPolicyId pulumi.StringOutput `pulumi:"gkeBackupRestorePlanIamPolicyId"`
	Location                        pulumi.StringOutput `pulumi:"location"`
	Name                            pulumi.StringOutput `pulumi:"name"`
	PolicyData                      pulumi.StringOutput `pulumi:"policyData"`
	Project                         pulumi.StringOutput `pulumi:"project"`
}

// NewGkeBackupRestorePlanIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewGkeBackupRestorePlanIamPolicy(ctx *pulumi.Context,
	name string, args *GkeBackupRestorePlanIamPolicyArgs, opts ...pulumi.ResourceOption) (*GkeBackupRestorePlanIamPolicy, error) {
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
	var resource GkeBackupRestorePlanIamPolicy
	err = ctx.RegisterPackageResource("google:index/gkeBackupRestorePlanIamPolicy:GkeBackupRestorePlanIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeBackupRestorePlanIamPolicy gets an existing GkeBackupRestorePlanIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeBackupRestorePlanIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeBackupRestorePlanIamPolicyState, opts ...pulumi.ResourceOption) (*GkeBackupRestorePlanIamPolicy, error) {
	var resource GkeBackupRestorePlanIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/gkeBackupRestorePlanIamPolicy:GkeBackupRestorePlanIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeBackupRestorePlanIamPolicy resources.
type gkeBackupRestorePlanIamPolicyState struct {
	Etag                            *string `pulumi:"etag"`
	GkeBackupRestorePlanIamPolicyId *string `pulumi:"gkeBackupRestorePlanIamPolicyId"`
	Location                        *string `pulumi:"location"`
	Name                            *string `pulumi:"name"`
	PolicyData                      *string `pulumi:"policyData"`
	Project                         *string `pulumi:"project"`
}

type GkeBackupRestorePlanIamPolicyState struct {
	Etag                            pulumi.StringPtrInput
	GkeBackupRestorePlanIamPolicyId pulumi.StringPtrInput
	Location                        pulumi.StringPtrInput
	Name                            pulumi.StringPtrInput
	PolicyData                      pulumi.StringPtrInput
	Project                         pulumi.StringPtrInput
}

func (GkeBackupRestorePlanIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeBackupRestorePlanIamPolicyState)(nil)).Elem()
}

type gkeBackupRestorePlanIamPolicyArgs struct {
	GkeBackupRestorePlanIamPolicyId *string `pulumi:"gkeBackupRestorePlanIamPolicyId"`
	Location                        *string `pulumi:"location"`
	Name                            *string `pulumi:"name"`
	PolicyData                      string  `pulumi:"policyData"`
	Project                         *string `pulumi:"project"`
}

// The set of arguments for constructing a GkeBackupRestorePlanIamPolicy resource.
type GkeBackupRestorePlanIamPolicyArgs struct {
	GkeBackupRestorePlanIamPolicyId pulumi.StringPtrInput
	Location                        pulumi.StringPtrInput
	Name                            pulumi.StringPtrInput
	PolicyData                      pulumi.StringInput
	Project                         pulumi.StringPtrInput
}

func (GkeBackupRestorePlanIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeBackupRestorePlanIamPolicyArgs)(nil)).Elem()
}

type GkeBackupRestorePlanIamPolicyInput interface {
	pulumi.Input

	ToGkeBackupRestorePlanIamPolicyOutput() GkeBackupRestorePlanIamPolicyOutput
	ToGkeBackupRestorePlanIamPolicyOutputWithContext(ctx context.Context) GkeBackupRestorePlanIamPolicyOutput
}

func (*GkeBackupRestorePlanIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeBackupRestorePlanIamPolicy)(nil)).Elem()
}

func (i *GkeBackupRestorePlanIamPolicy) ToGkeBackupRestorePlanIamPolicyOutput() GkeBackupRestorePlanIamPolicyOutput {
	return i.ToGkeBackupRestorePlanIamPolicyOutputWithContext(context.Background())
}

func (i *GkeBackupRestorePlanIamPolicy) ToGkeBackupRestorePlanIamPolicyOutputWithContext(ctx context.Context) GkeBackupRestorePlanIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeBackupRestorePlanIamPolicyOutput)
}

type GkeBackupRestorePlanIamPolicyOutput struct{ *pulumi.OutputState }

func (GkeBackupRestorePlanIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeBackupRestorePlanIamPolicy)(nil)).Elem()
}

func (o GkeBackupRestorePlanIamPolicyOutput) ToGkeBackupRestorePlanIamPolicyOutput() GkeBackupRestorePlanIamPolicyOutput {
	return o
}

func (o GkeBackupRestorePlanIamPolicyOutput) ToGkeBackupRestorePlanIamPolicyOutputWithContext(ctx context.Context) GkeBackupRestorePlanIamPolicyOutput {
	return o
}

func (o GkeBackupRestorePlanIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupRestorePlanIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GkeBackupRestorePlanIamPolicyOutput) GkeBackupRestorePlanIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupRestorePlanIamPolicy) pulumi.StringOutput { return v.GkeBackupRestorePlanIamPolicyId }).(pulumi.StringOutput)
}

func (o GkeBackupRestorePlanIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupRestorePlanIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GkeBackupRestorePlanIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupRestorePlanIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GkeBackupRestorePlanIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupRestorePlanIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o GkeBackupRestorePlanIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeBackupRestorePlanIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeBackupRestorePlanIamPolicyInput)(nil)).Elem(), &GkeBackupRestorePlanIamPolicy{})
	pulumi.RegisterOutputType(GkeBackupRestorePlanIamPolicyOutput{})
}
