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

type IamPrincipalAccessBoundaryPolicy struct {
	pulumi.CustomResourceState

	// User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size
	// limitations **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Output only. The time when the principal access boundary policy was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Principal access boundary policy details
	Details IamPrincipalAccessBoundaryPolicyDetailsPtrOutput `pulumi:"details"`
	// The description of the principal access boundary policy. Must be less than or equal to 63 characters.
	DisplayName          pulumi.StringPtrOutput `pulumi:"displayName"`
	EffectiveAnnotations pulumi.StringMapOutput `pulumi:"effectiveAnnotations"`
	// The etag for the principal access boundary. If this is provided on update, it must match the server's etag.
	Etag                               pulumi.StringOutput `pulumi:"etag"`
	IamPrincipalAccessBoundaryPolicyId pulumi.StringOutput `pulumi:"iamPrincipalAccessBoundaryPolicyId"`
	// The location the principal access boundary policy is in.
	Location pulumi.StringOutput `pulumi:"location"`
	// Identifier. The resource name of the principal access boundary policy. The following format is supported:
	// 'organizations/{organization_id}/locations/{location}/principalAccessBoundaryPolicies/{policy_id}'
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent organization of the principal access boundary policy.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The ID to use to create the principal access boundary policy. This value must start with a lowercase letter followed by
	// up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.
	PrincipalAccessBoundaryPolicyId pulumi.StringOutput                               `pulumi:"principalAccessBoundaryPolicyId"`
	Timeouts                        IamPrincipalAccessBoundaryPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. The globally unique ID of the principal access boundary policy.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time when the principal access boundary policy was most recently updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewIamPrincipalAccessBoundaryPolicy registers a new resource with the given unique name, arguments, and options.
func NewIamPrincipalAccessBoundaryPolicy(ctx *pulumi.Context,
	name string, args *IamPrincipalAccessBoundaryPolicyArgs, opts ...pulumi.ResourceOption) (*IamPrincipalAccessBoundaryPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.PrincipalAccessBoundaryPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalAccessBoundaryPolicyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IamPrincipalAccessBoundaryPolicy
	err = ctx.RegisterPackageResource("google-beta:index/iamPrincipalAccessBoundaryPolicy:IamPrincipalAccessBoundaryPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamPrincipalAccessBoundaryPolicy gets an existing IamPrincipalAccessBoundaryPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamPrincipalAccessBoundaryPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamPrincipalAccessBoundaryPolicyState, opts ...pulumi.ResourceOption) (*IamPrincipalAccessBoundaryPolicy, error) {
	var resource IamPrincipalAccessBoundaryPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/iamPrincipalAccessBoundaryPolicy:IamPrincipalAccessBoundaryPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamPrincipalAccessBoundaryPolicy resources.
type iamPrincipalAccessBoundaryPolicyState struct {
	// User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size
	// limitations **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Output only. The time when the principal access boundary policy was created.
	CreateTime *string `pulumi:"createTime"`
	// Principal access boundary policy details
	Details *IamPrincipalAccessBoundaryPolicyDetails `pulumi:"details"`
	// The description of the principal access boundary policy. Must be less than or equal to 63 characters.
	DisplayName          *string           `pulumi:"displayName"`
	EffectiveAnnotations map[string]string `pulumi:"effectiveAnnotations"`
	// The etag for the principal access boundary. If this is provided on update, it must match the server's etag.
	Etag                               *string `pulumi:"etag"`
	IamPrincipalAccessBoundaryPolicyId *string `pulumi:"iamPrincipalAccessBoundaryPolicyId"`
	// The location the principal access boundary policy is in.
	Location *string `pulumi:"location"`
	// Identifier. The resource name of the principal access boundary policy. The following format is supported:
	// 'organizations/{organization_id}/locations/{location}/principalAccessBoundaryPolicies/{policy_id}'
	Name *string `pulumi:"name"`
	// The parent organization of the principal access boundary policy.
	Organization *string `pulumi:"organization"`
	// The ID to use to create the principal access boundary policy. This value must start with a lowercase letter followed by
	// up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.
	PrincipalAccessBoundaryPolicyId *string                                   `pulumi:"principalAccessBoundaryPolicyId"`
	Timeouts                        *IamPrincipalAccessBoundaryPolicyTimeouts `pulumi:"timeouts"`
	// Output only. The globally unique ID of the principal access boundary policy.
	Uid *string `pulumi:"uid"`
	// Output only. The time when the principal access boundary policy was most recently updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type IamPrincipalAccessBoundaryPolicyState struct {
	// User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size
	// limitations **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Output only. The time when the principal access boundary policy was created.
	CreateTime pulumi.StringPtrInput
	// Principal access boundary policy details
	Details IamPrincipalAccessBoundaryPolicyDetailsPtrInput
	// The description of the principal access boundary policy. Must be less than or equal to 63 characters.
	DisplayName          pulumi.StringPtrInput
	EffectiveAnnotations pulumi.StringMapInput
	// The etag for the principal access boundary. If this is provided on update, it must match the server's etag.
	Etag                               pulumi.StringPtrInput
	IamPrincipalAccessBoundaryPolicyId pulumi.StringPtrInput
	// The location the principal access boundary policy is in.
	Location pulumi.StringPtrInput
	// Identifier. The resource name of the principal access boundary policy. The following format is supported:
	// 'organizations/{organization_id}/locations/{location}/principalAccessBoundaryPolicies/{policy_id}'
	Name pulumi.StringPtrInput
	// The parent organization of the principal access boundary policy.
	Organization pulumi.StringPtrInput
	// The ID to use to create the principal access boundary policy. This value must start with a lowercase letter followed by
	// up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.
	PrincipalAccessBoundaryPolicyId pulumi.StringPtrInput
	Timeouts                        IamPrincipalAccessBoundaryPolicyTimeoutsPtrInput
	// Output only. The globally unique ID of the principal access boundary policy.
	Uid pulumi.StringPtrInput
	// Output only. The time when the principal access boundary policy was most recently updated.
	UpdateTime pulumi.StringPtrInput
}

func (IamPrincipalAccessBoundaryPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPrincipalAccessBoundaryPolicyState)(nil)).Elem()
}

type iamPrincipalAccessBoundaryPolicyArgs struct {
	// User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size
	// limitations **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Principal access boundary policy details
	Details *IamPrincipalAccessBoundaryPolicyDetails `pulumi:"details"`
	// The description of the principal access boundary policy. Must be less than or equal to 63 characters.
	DisplayName                        *string `pulumi:"displayName"`
	IamPrincipalAccessBoundaryPolicyId *string `pulumi:"iamPrincipalAccessBoundaryPolicyId"`
	// The location the principal access boundary policy is in.
	Location string `pulumi:"location"`
	// The parent organization of the principal access boundary policy.
	Organization string `pulumi:"organization"`
	// The ID to use to create the principal access boundary policy. This value must start with a lowercase letter followed by
	// up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.
	PrincipalAccessBoundaryPolicyId string                                    `pulumi:"principalAccessBoundaryPolicyId"`
	Timeouts                        *IamPrincipalAccessBoundaryPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a IamPrincipalAccessBoundaryPolicy resource.
type IamPrincipalAccessBoundaryPolicyArgs struct {
	// User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size
	// limitations **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Principal access boundary policy details
	Details IamPrincipalAccessBoundaryPolicyDetailsPtrInput
	// The description of the principal access boundary policy. Must be less than or equal to 63 characters.
	DisplayName                        pulumi.StringPtrInput
	IamPrincipalAccessBoundaryPolicyId pulumi.StringPtrInput
	// The location the principal access boundary policy is in.
	Location pulumi.StringInput
	// The parent organization of the principal access boundary policy.
	Organization pulumi.StringInput
	// The ID to use to create the principal access boundary policy. This value must start with a lowercase letter followed by
	// up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.
	PrincipalAccessBoundaryPolicyId pulumi.StringInput
	Timeouts                        IamPrincipalAccessBoundaryPolicyTimeoutsPtrInput
}

func (IamPrincipalAccessBoundaryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPrincipalAccessBoundaryPolicyArgs)(nil)).Elem()
}

type IamPrincipalAccessBoundaryPolicyInput interface {
	pulumi.Input

	ToIamPrincipalAccessBoundaryPolicyOutput() IamPrincipalAccessBoundaryPolicyOutput
	ToIamPrincipalAccessBoundaryPolicyOutputWithContext(ctx context.Context) IamPrincipalAccessBoundaryPolicyOutput
}

func (*IamPrincipalAccessBoundaryPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPrincipalAccessBoundaryPolicy)(nil)).Elem()
}

func (i *IamPrincipalAccessBoundaryPolicy) ToIamPrincipalAccessBoundaryPolicyOutput() IamPrincipalAccessBoundaryPolicyOutput {
	return i.ToIamPrincipalAccessBoundaryPolicyOutputWithContext(context.Background())
}

func (i *IamPrincipalAccessBoundaryPolicy) ToIamPrincipalAccessBoundaryPolicyOutputWithContext(ctx context.Context) IamPrincipalAccessBoundaryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPrincipalAccessBoundaryPolicyOutput)
}

type IamPrincipalAccessBoundaryPolicyOutput struct{ *pulumi.OutputState }

func (IamPrincipalAccessBoundaryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPrincipalAccessBoundaryPolicy)(nil)).Elem()
}

func (o IamPrincipalAccessBoundaryPolicyOutput) ToIamPrincipalAccessBoundaryPolicyOutput() IamPrincipalAccessBoundaryPolicyOutput {
	return o
}

func (o IamPrincipalAccessBoundaryPolicyOutput) ToIamPrincipalAccessBoundaryPolicyOutputWithContext(ctx context.Context) IamPrincipalAccessBoundaryPolicyOutput {
	return o
}

// User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size
// limitations **Note**: This field is non-authoritative, and will only manage the annotations present in your
// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
func (o IamPrincipalAccessBoundaryPolicyOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Output only. The time when the principal access boundary policy was created.
func (o IamPrincipalAccessBoundaryPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Principal access boundary policy details
func (o IamPrincipalAccessBoundaryPolicyOutput) Details() IamPrincipalAccessBoundaryPolicyDetailsPtrOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) IamPrincipalAccessBoundaryPolicyDetailsPtrOutput {
		return v.Details
	}).(IamPrincipalAccessBoundaryPolicyDetailsPtrOutput)
}

// The description of the principal access boundary policy. Must be less than or equal to 63 characters.
func (o IamPrincipalAccessBoundaryPolicyOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o IamPrincipalAccessBoundaryPolicyOutput) EffectiveAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringMapOutput { return v.EffectiveAnnotations }).(pulumi.StringMapOutput)
}

// The etag for the principal access boundary. If this is provided on update, it must match the server's etag.
func (o IamPrincipalAccessBoundaryPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IamPrincipalAccessBoundaryPolicyOutput) IamPrincipalAccessBoundaryPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput {
		return v.IamPrincipalAccessBoundaryPolicyId
	}).(pulumi.StringOutput)
}

// The location the principal access boundary policy is in.
func (o IamPrincipalAccessBoundaryPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Identifier. The resource name of the principal access boundary policy. The following format is supported:
// 'organizations/{organization_id}/locations/{location}/principalAccessBoundaryPolicies/{policy_id}'
func (o IamPrincipalAccessBoundaryPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent organization of the principal access boundary policy.
func (o IamPrincipalAccessBoundaryPolicyOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The ID to use to create the principal access boundary policy. This value must start with a lowercase letter followed by
// up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.
func (o IamPrincipalAccessBoundaryPolicyOutput) PrincipalAccessBoundaryPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput {
		return v.PrincipalAccessBoundaryPolicyId
	}).(pulumi.StringOutput)
}

func (o IamPrincipalAccessBoundaryPolicyOutput) Timeouts() IamPrincipalAccessBoundaryPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) IamPrincipalAccessBoundaryPolicyTimeoutsPtrOutput {
		return v.Timeouts
	}).(IamPrincipalAccessBoundaryPolicyTimeoutsPtrOutput)
}

// Output only. The globally unique ID of the principal access boundary policy.
func (o IamPrincipalAccessBoundaryPolicyOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The time when the principal access boundary policy was most recently updated.
func (o IamPrincipalAccessBoundaryPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPrincipalAccessBoundaryPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamPrincipalAccessBoundaryPolicyInput)(nil)).Elem(), &IamPrincipalAccessBoundaryPolicy{})
	pulumi.RegisterOutputType(IamPrincipalAccessBoundaryPolicyOutput{})
}
