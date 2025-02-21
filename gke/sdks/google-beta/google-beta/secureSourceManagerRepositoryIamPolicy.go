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

type SecureSourceManagerRepositoryIamPolicy struct {
	pulumi.CustomResourceState

	Etag                                     pulumi.StringOutput `pulumi:"etag"`
	Location                                 pulumi.StringOutput `pulumi:"location"`
	PolicyData                               pulumi.StringOutput `pulumi:"policyData"`
	Project                                  pulumi.StringOutput `pulumi:"project"`
	RepositoryId                             pulumi.StringOutput `pulumi:"repositoryId"`
	SecureSourceManagerRepositoryIamPolicyId pulumi.StringOutput `pulumi:"secureSourceManagerRepositoryIamPolicyId"`
}

// NewSecureSourceManagerRepositoryIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecureSourceManagerRepositoryIamPolicy(ctx *pulumi.Context,
	name string, args *SecureSourceManagerRepositoryIamPolicyArgs, opts ...pulumi.ResourceOption) (*SecureSourceManagerRepositoryIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SecureSourceManagerRepositoryIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/secureSourceManagerRepositoryIamPolicy:SecureSourceManagerRepositoryIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecureSourceManagerRepositoryIamPolicy gets an existing SecureSourceManagerRepositoryIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecureSourceManagerRepositoryIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecureSourceManagerRepositoryIamPolicyState, opts ...pulumi.ResourceOption) (*SecureSourceManagerRepositoryIamPolicy, error) {
	var resource SecureSourceManagerRepositoryIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/secureSourceManagerRepositoryIamPolicy:SecureSourceManagerRepositoryIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecureSourceManagerRepositoryIamPolicy resources.
type secureSourceManagerRepositoryIamPolicyState struct {
	Etag                                     *string `pulumi:"etag"`
	Location                                 *string `pulumi:"location"`
	PolicyData                               *string `pulumi:"policyData"`
	Project                                  *string `pulumi:"project"`
	RepositoryId                             *string `pulumi:"repositoryId"`
	SecureSourceManagerRepositoryIamPolicyId *string `pulumi:"secureSourceManagerRepositoryIamPolicyId"`
}

type SecureSourceManagerRepositoryIamPolicyState struct {
	Etag                                     pulumi.StringPtrInput
	Location                                 pulumi.StringPtrInput
	PolicyData                               pulumi.StringPtrInput
	Project                                  pulumi.StringPtrInput
	RepositoryId                             pulumi.StringPtrInput
	SecureSourceManagerRepositoryIamPolicyId pulumi.StringPtrInput
}

func (SecureSourceManagerRepositoryIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerRepositoryIamPolicyState)(nil)).Elem()
}

type secureSourceManagerRepositoryIamPolicyArgs struct {
	Location                                 *string `pulumi:"location"`
	PolicyData                               string  `pulumi:"policyData"`
	Project                                  *string `pulumi:"project"`
	RepositoryId                             string  `pulumi:"repositoryId"`
	SecureSourceManagerRepositoryIamPolicyId *string `pulumi:"secureSourceManagerRepositoryIamPolicyId"`
}

// The set of arguments for constructing a SecureSourceManagerRepositoryIamPolicy resource.
type SecureSourceManagerRepositoryIamPolicyArgs struct {
	Location                                 pulumi.StringPtrInput
	PolicyData                               pulumi.StringInput
	Project                                  pulumi.StringPtrInput
	RepositoryId                             pulumi.StringInput
	SecureSourceManagerRepositoryIamPolicyId pulumi.StringPtrInput
}

func (SecureSourceManagerRepositoryIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerRepositoryIamPolicyArgs)(nil)).Elem()
}

type SecureSourceManagerRepositoryIamPolicyInput interface {
	pulumi.Input

	ToSecureSourceManagerRepositoryIamPolicyOutput() SecureSourceManagerRepositoryIamPolicyOutput
	ToSecureSourceManagerRepositoryIamPolicyOutputWithContext(ctx context.Context) SecureSourceManagerRepositoryIamPolicyOutput
}

func (*SecureSourceManagerRepositoryIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerRepositoryIamPolicy)(nil)).Elem()
}

func (i *SecureSourceManagerRepositoryIamPolicy) ToSecureSourceManagerRepositoryIamPolicyOutput() SecureSourceManagerRepositoryIamPolicyOutput {
	return i.ToSecureSourceManagerRepositoryIamPolicyOutputWithContext(context.Background())
}

func (i *SecureSourceManagerRepositoryIamPolicy) ToSecureSourceManagerRepositoryIamPolicyOutputWithContext(ctx context.Context) SecureSourceManagerRepositoryIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureSourceManagerRepositoryIamPolicyOutput)
}

type SecureSourceManagerRepositoryIamPolicyOutput struct{ *pulumi.OutputState }

func (SecureSourceManagerRepositoryIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerRepositoryIamPolicy)(nil)).Elem()
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) ToSecureSourceManagerRepositoryIamPolicyOutput() SecureSourceManagerRepositoryIamPolicyOutput {
	return o
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) ToSecureSourceManagerRepositoryIamPolicyOutputWithContext(ctx context.Context) SecureSourceManagerRepositoryIamPolicyOutput {
	return o
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepositoryIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepositoryIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepositoryIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepositoryIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepositoryIamPolicy) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryIamPolicyOutput) SecureSourceManagerRepositoryIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepositoryIamPolicy) pulumi.StringOutput {
		return v.SecureSourceManagerRepositoryIamPolicyId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecureSourceManagerRepositoryIamPolicyInput)(nil)).Elem(), &SecureSourceManagerRepositoryIamPolicy{})
	pulumi.RegisterOutputType(SecureSourceManagerRepositoryIamPolicyOutput{})
}
