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

type ClouddeployDeliveryPipelineIamMember struct {
	pulumi.CustomResourceState

	ClouddeployDeliveryPipelineIamMemberId pulumi.StringOutput                                    `pulumi:"clouddeployDeliveryPipelineIamMemberId"`
	Condition                              ClouddeployDeliveryPipelineIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                                   pulumi.StringOutput                                    `pulumi:"etag"`
	Location                               pulumi.StringOutput                                    `pulumi:"location"`
	Member                                 pulumi.StringOutput                                    `pulumi:"member"`
	Name                                   pulumi.StringOutput                                    `pulumi:"name"`
	Project                                pulumi.StringOutput                                    `pulumi:"project"`
	Role                                   pulumi.StringOutput                                    `pulumi:"role"`
}

// NewClouddeployDeliveryPipelineIamMember registers a new resource with the given unique name, arguments, and options.
func NewClouddeployDeliveryPipelineIamMember(ctx *pulumi.Context,
	name string, args *ClouddeployDeliveryPipelineIamMemberArgs, opts ...pulumi.ResourceOption) (*ClouddeployDeliveryPipelineIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ClouddeployDeliveryPipelineIamMember
	err = ctx.RegisterPackageResource("google:index/clouddeployDeliveryPipelineIamMember:ClouddeployDeliveryPipelineIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClouddeployDeliveryPipelineIamMember gets an existing ClouddeployDeliveryPipelineIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClouddeployDeliveryPipelineIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClouddeployDeliveryPipelineIamMemberState, opts ...pulumi.ResourceOption) (*ClouddeployDeliveryPipelineIamMember, error) {
	var resource ClouddeployDeliveryPipelineIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/clouddeployDeliveryPipelineIamMember:ClouddeployDeliveryPipelineIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClouddeployDeliveryPipelineIamMember resources.
type clouddeployDeliveryPipelineIamMemberState struct {
	ClouddeployDeliveryPipelineIamMemberId *string                                        `pulumi:"clouddeployDeliveryPipelineIamMemberId"`
	Condition                              *ClouddeployDeliveryPipelineIamMemberCondition `pulumi:"condition"`
	Etag                                   *string                                        `pulumi:"etag"`
	Location                               *string                                        `pulumi:"location"`
	Member                                 *string                                        `pulumi:"member"`
	Name                                   *string                                        `pulumi:"name"`
	Project                                *string                                        `pulumi:"project"`
	Role                                   *string                                        `pulumi:"role"`
}

type ClouddeployDeliveryPipelineIamMemberState struct {
	ClouddeployDeliveryPipelineIamMemberId pulumi.StringPtrInput
	Condition                              ClouddeployDeliveryPipelineIamMemberConditionPtrInput
	Etag                                   pulumi.StringPtrInput
	Location                               pulumi.StringPtrInput
	Member                                 pulumi.StringPtrInput
	Name                                   pulumi.StringPtrInput
	Project                                pulumi.StringPtrInput
	Role                                   pulumi.StringPtrInput
}

func (ClouddeployDeliveryPipelineIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*clouddeployDeliveryPipelineIamMemberState)(nil)).Elem()
}

type clouddeployDeliveryPipelineIamMemberArgs struct {
	ClouddeployDeliveryPipelineIamMemberId *string                                        `pulumi:"clouddeployDeliveryPipelineIamMemberId"`
	Condition                              *ClouddeployDeliveryPipelineIamMemberCondition `pulumi:"condition"`
	Location                               *string                                        `pulumi:"location"`
	Member                                 string                                         `pulumi:"member"`
	Name                                   *string                                        `pulumi:"name"`
	Project                                *string                                        `pulumi:"project"`
	Role                                   string                                         `pulumi:"role"`
}

// The set of arguments for constructing a ClouddeployDeliveryPipelineIamMember resource.
type ClouddeployDeliveryPipelineIamMemberArgs struct {
	ClouddeployDeliveryPipelineIamMemberId pulumi.StringPtrInput
	Condition                              ClouddeployDeliveryPipelineIamMemberConditionPtrInput
	Location                               pulumi.StringPtrInput
	Member                                 pulumi.StringInput
	Name                                   pulumi.StringPtrInput
	Project                                pulumi.StringPtrInput
	Role                                   pulumi.StringInput
}

func (ClouddeployDeliveryPipelineIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clouddeployDeliveryPipelineIamMemberArgs)(nil)).Elem()
}

type ClouddeployDeliveryPipelineIamMemberInput interface {
	pulumi.Input

	ToClouddeployDeliveryPipelineIamMemberOutput() ClouddeployDeliveryPipelineIamMemberOutput
	ToClouddeployDeliveryPipelineIamMemberOutputWithContext(ctx context.Context) ClouddeployDeliveryPipelineIamMemberOutput
}

func (*ClouddeployDeliveryPipelineIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ClouddeployDeliveryPipelineIamMember)(nil)).Elem()
}

func (i *ClouddeployDeliveryPipelineIamMember) ToClouddeployDeliveryPipelineIamMemberOutput() ClouddeployDeliveryPipelineIamMemberOutput {
	return i.ToClouddeployDeliveryPipelineIamMemberOutputWithContext(context.Background())
}

func (i *ClouddeployDeliveryPipelineIamMember) ToClouddeployDeliveryPipelineIamMemberOutputWithContext(ctx context.Context) ClouddeployDeliveryPipelineIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClouddeployDeliveryPipelineIamMemberOutput)
}

type ClouddeployDeliveryPipelineIamMemberOutput struct{ *pulumi.OutputState }

func (ClouddeployDeliveryPipelineIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClouddeployDeliveryPipelineIamMember)(nil)).Elem()
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) ToClouddeployDeliveryPipelineIamMemberOutput() ClouddeployDeliveryPipelineIamMemberOutput {
	return o
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) ToClouddeployDeliveryPipelineIamMemberOutputWithContext(ctx context.Context) ClouddeployDeliveryPipelineIamMemberOutput {
	return o
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) ClouddeployDeliveryPipelineIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) pulumi.StringOutput {
		return v.ClouddeployDeliveryPipelineIamMemberId
	}).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) Condition() ClouddeployDeliveryPipelineIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) ClouddeployDeliveryPipelineIamMemberConditionPtrOutput {
		return v.Condition
	}).(ClouddeployDeliveryPipelineIamMemberConditionPtrOutput)
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClouddeployDeliveryPipelineIamMemberInput)(nil)).Elem(), &ClouddeployDeliveryPipelineIamMember{})
	pulumi.RegisterOutputType(ClouddeployDeliveryPipelineIamMemberOutput{})
}
