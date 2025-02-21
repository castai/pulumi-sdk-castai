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

type CloudTasksQueueIamPolicy struct {
	pulumi.CustomResourceState

	CloudTasksQueueIamPolicyId pulumi.StringOutput `pulumi:"cloudTasksQueueIamPolicyId"`
	Etag                       pulumi.StringOutput `pulumi:"etag"`
	Location                   pulumi.StringOutput `pulumi:"location"`
	Name                       pulumi.StringOutput `pulumi:"name"`
	PolicyData                 pulumi.StringOutput `pulumi:"policyData"`
	Project                    pulumi.StringOutput `pulumi:"project"`
}

// NewCloudTasksQueueIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewCloudTasksQueueIamPolicy(ctx *pulumi.Context,
	name string, args *CloudTasksQueueIamPolicyArgs, opts ...pulumi.ResourceOption) (*CloudTasksQueueIamPolicy, error) {
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
	var resource CloudTasksQueueIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/cloudTasksQueueIamPolicy:CloudTasksQueueIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudTasksQueueIamPolicy gets an existing CloudTasksQueueIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudTasksQueueIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudTasksQueueIamPolicyState, opts ...pulumi.ResourceOption) (*CloudTasksQueueIamPolicy, error) {
	var resource CloudTasksQueueIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/cloudTasksQueueIamPolicy:CloudTasksQueueIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudTasksQueueIamPolicy resources.
type cloudTasksQueueIamPolicyState struct {
	CloudTasksQueueIamPolicyId *string `pulumi:"cloudTasksQueueIamPolicyId"`
	Etag                       *string `pulumi:"etag"`
	Location                   *string `pulumi:"location"`
	Name                       *string `pulumi:"name"`
	PolicyData                 *string `pulumi:"policyData"`
	Project                    *string `pulumi:"project"`
}

type CloudTasksQueueIamPolicyState struct {
	CloudTasksQueueIamPolicyId pulumi.StringPtrInput
	Etag                       pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	PolicyData                 pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
}

func (CloudTasksQueueIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudTasksQueueIamPolicyState)(nil)).Elem()
}

type cloudTasksQueueIamPolicyArgs struct {
	CloudTasksQueueIamPolicyId *string `pulumi:"cloudTasksQueueIamPolicyId"`
	Location                   *string `pulumi:"location"`
	Name                       *string `pulumi:"name"`
	PolicyData                 string  `pulumi:"policyData"`
	Project                    *string `pulumi:"project"`
}

// The set of arguments for constructing a CloudTasksQueueIamPolicy resource.
type CloudTasksQueueIamPolicyArgs struct {
	CloudTasksQueueIamPolicyId pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	PolicyData                 pulumi.StringInput
	Project                    pulumi.StringPtrInput
}

func (CloudTasksQueueIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudTasksQueueIamPolicyArgs)(nil)).Elem()
}

type CloudTasksQueueIamPolicyInput interface {
	pulumi.Input

	ToCloudTasksQueueIamPolicyOutput() CloudTasksQueueIamPolicyOutput
	ToCloudTasksQueueIamPolicyOutputWithContext(ctx context.Context) CloudTasksQueueIamPolicyOutput
}

func (*CloudTasksQueueIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTasksQueueIamPolicy)(nil)).Elem()
}

func (i *CloudTasksQueueIamPolicy) ToCloudTasksQueueIamPolicyOutput() CloudTasksQueueIamPolicyOutput {
	return i.ToCloudTasksQueueIamPolicyOutputWithContext(context.Background())
}

func (i *CloudTasksQueueIamPolicy) ToCloudTasksQueueIamPolicyOutputWithContext(ctx context.Context) CloudTasksQueueIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTasksQueueIamPolicyOutput)
}

type CloudTasksQueueIamPolicyOutput struct{ *pulumi.OutputState }

func (CloudTasksQueueIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTasksQueueIamPolicy)(nil)).Elem()
}

func (o CloudTasksQueueIamPolicyOutput) ToCloudTasksQueueIamPolicyOutput() CloudTasksQueueIamPolicyOutput {
	return o
}

func (o CloudTasksQueueIamPolicyOutput) ToCloudTasksQueueIamPolicyOutputWithContext(ctx context.Context) CloudTasksQueueIamPolicyOutput {
	return o
}

func (o CloudTasksQueueIamPolicyOutput) CloudTasksQueueIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamPolicy) pulumi.StringOutput { return v.CloudTasksQueueIamPolicyId }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudTasksQueueIamPolicyInput)(nil)).Elem(), &CloudTasksQueueIamPolicy{})
	pulumi.RegisterOutputType(CloudTasksQueueIamPolicyOutput{})
}
