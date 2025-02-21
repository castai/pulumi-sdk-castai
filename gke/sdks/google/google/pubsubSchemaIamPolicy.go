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

type PubsubSchemaIamPolicy struct {
	pulumi.CustomResourceState

	Etag                    pulumi.StringOutput `pulumi:"etag"`
	PolicyData              pulumi.StringOutput `pulumi:"policyData"`
	Project                 pulumi.StringOutput `pulumi:"project"`
	PubsubSchemaIamPolicyId pulumi.StringOutput `pulumi:"pubsubSchemaIamPolicyId"`
	Schema                  pulumi.StringOutput `pulumi:"schema"`
}

// NewPubsubSchemaIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewPubsubSchemaIamPolicy(ctx *pulumi.Context,
	name string, args *PubsubSchemaIamPolicyArgs, opts ...pulumi.ResourceOption) (*PubsubSchemaIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource PubsubSchemaIamPolicy
	err = ctx.RegisterPackageResource("google:index/pubsubSchemaIamPolicy:PubsubSchemaIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPubsubSchemaIamPolicy gets an existing PubsubSchemaIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPubsubSchemaIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PubsubSchemaIamPolicyState, opts ...pulumi.ResourceOption) (*PubsubSchemaIamPolicy, error) {
	var resource PubsubSchemaIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/pubsubSchemaIamPolicy:PubsubSchemaIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PubsubSchemaIamPolicy resources.
type pubsubSchemaIamPolicyState struct {
	Etag                    *string `pulumi:"etag"`
	PolicyData              *string `pulumi:"policyData"`
	Project                 *string `pulumi:"project"`
	PubsubSchemaIamPolicyId *string `pulumi:"pubsubSchemaIamPolicyId"`
	Schema                  *string `pulumi:"schema"`
}

type PubsubSchemaIamPolicyState struct {
	Etag                    pulumi.StringPtrInput
	PolicyData              pulumi.StringPtrInput
	Project                 pulumi.StringPtrInput
	PubsubSchemaIamPolicyId pulumi.StringPtrInput
	Schema                  pulumi.StringPtrInput
}

func (PubsubSchemaIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*pubsubSchemaIamPolicyState)(nil)).Elem()
}

type pubsubSchemaIamPolicyArgs struct {
	PolicyData              string  `pulumi:"policyData"`
	Project                 *string `pulumi:"project"`
	PubsubSchemaIamPolicyId *string `pulumi:"pubsubSchemaIamPolicyId"`
	Schema                  string  `pulumi:"schema"`
}

// The set of arguments for constructing a PubsubSchemaIamPolicy resource.
type PubsubSchemaIamPolicyArgs struct {
	PolicyData              pulumi.StringInput
	Project                 pulumi.StringPtrInput
	PubsubSchemaIamPolicyId pulumi.StringPtrInput
	Schema                  pulumi.StringInput
}

func (PubsubSchemaIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pubsubSchemaIamPolicyArgs)(nil)).Elem()
}

type PubsubSchemaIamPolicyInput interface {
	pulumi.Input

	ToPubsubSchemaIamPolicyOutput() PubsubSchemaIamPolicyOutput
	ToPubsubSchemaIamPolicyOutputWithContext(ctx context.Context) PubsubSchemaIamPolicyOutput
}

func (*PubsubSchemaIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PubsubSchemaIamPolicy)(nil)).Elem()
}

func (i *PubsubSchemaIamPolicy) ToPubsubSchemaIamPolicyOutput() PubsubSchemaIamPolicyOutput {
	return i.ToPubsubSchemaIamPolicyOutputWithContext(context.Background())
}

func (i *PubsubSchemaIamPolicy) ToPubsubSchemaIamPolicyOutputWithContext(ctx context.Context) PubsubSchemaIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PubsubSchemaIamPolicyOutput)
}

type PubsubSchemaIamPolicyOutput struct{ *pulumi.OutputState }

func (PubsubSchemaIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PubsubSchemaIamPolicy)(nil)).Elem()
}

func (o PubsubSchemaIamPolicyOutput) ToPubsubSchemaIamPolicyOutput() PubsubSchemaIamPolicyOutput {
	return o
}

func (o PubsubSchemaIamPolicyOutput) ToPubsubSchemaIamPolicyOutputWithContext(ctx context.Context) PubsubSchemaIamPolicyOutput {
	return o
}

func (o PubsubSchemaIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSchemaIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PubsubSchemaIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSchemaIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o PubsubSchemaIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSchemaIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o PubsubSchemaIamPolicyOutput) PubsubSchemaIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSchemaIamPolicy) pulumi.StringOutput { return v.PubsubSchemaIamPolicyId }).(pulumi.StringOutput)
}

func (o PubsubSchemaIamPolicyOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSchemaIamPolicy) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PubsubSchemaIamPolicyInput)(nil)).Elem(), &PubsubSchemaIamPolicy{})
	pulumi.RegisterOutputType(PubsubSchemaIamPolicyOutput{})
}
