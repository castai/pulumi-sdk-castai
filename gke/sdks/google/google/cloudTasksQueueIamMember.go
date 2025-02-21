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

type CloudTasksQueueIamMember struct {
	pulumi.CustomResourceState

	CloudTasksQueueIamMemberId pulumi.StringOutput                        `pulumi:"cloudTasksQueueIamMemberId"`
	Condition                  CloudTasksQueueIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                       pulumi.StringOutput                        `pulumi:"etag"`
	Location                   pulumi.StringOutput                        `pulumi:"location"`
	Member                     pulumi.StringOutput                        `pulumi:"member"`
	Name                       pulumi.StringOutput                        `pulumi:"name"`
	Project                    pulumi.StringOutput                        `pulumi:"project"`
	Role                       pulumi.StringOutput                        `pulumi:"role"`
}

// NewCloudTasksQueueIamMember registers a new resource with the given unique name, arguments, and options.
func NewCloudTasksQueueIamMember(ctx *pulumi.Context,
	name string, args *CloudTasksQueueIamMemberArgs, opts ...pulumi.ResourceOption) (*CloudTasksQueueIamMember, error) {
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
	var resource CloudTasksQueueIamMember
	err = ctx.RegisterPackageResource("google:index/cloudTasksQueueIamMember:CloudTasksQueueIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudTasksQueueIamMember gets an existing CloudTasksQueueIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudTasksQueueIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudTasksQueueIamMemberState, opts ...pulumi.ResourceOption) (*CloudTasksQueueIamMember, error) {
	var resource CloudTasksQueueIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/cloudTasksQueueIamMember:CloudTasksQueueIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudTasksQueueIamMember resources.
type cloudTasksQueueIamMemberState struct {
	CloudTasksQueueIamMemberId *string                            `pulumi:"cloudTasksQueueIamMemberId"`
	Condition                  *CloudTasksQueueIamMemberCondition `pulumi:"condition"`
	Etag                       *string                            `pulumi:"etag"`
	Location                   *string                            `pulumi:"location"`
	Member                     *string                            `pulumi:"member"`
	Name                       *string                            `pulumi:"name"`
	Project                    *string                            `pulumi:"project"`
	Role                       *string                            `pulumi:"role"`
}

type CloudTasksQueueIamMemberState struct {
	CloudTasksQueueIamMemberId pulumi.StringPtrInput
	Condition                  CloudTasksQueueIamMemberConditionPtrInput
	Etag                       pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	Member                     pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	Role                       pulumi.StringPtrInput
}

func (CloudTasksQueueIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudTasksQueueIamMemberState)(nil)).Elem()
}

type cloudTasksQueueIamMemberArgs struct {
	CloudTasksQueueIamMemberId *string                            `pulumi:"cloudTasksQueueIamMemberId"`
	Condition                  *CloudTasksQueueIamMemberCondition `pulumi:"condition"`
	Location                   *string                            `pulumi:"location"`
	Member                     string                             `pulumi:"member"`
	Name                       *string                            `pulumi:"name"`
	Project                    *string                            `pulumi:"project"`
	Role                       string                             `pulumi:"role"`
}

// The set of arguments for constructing a CloudTasksQueueIamMember resource.
type CloudTasksQueueIamMemberArgs struct {
	CloudTasksQueueIamMemberId pulumi.StringPtrInput
	Condition                  CloudTasksQueueIamMemberConditionPtrInput
	Location                   pulumi.StringPtrInput
	Member                     pulumi.StringInput
	Name                       pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	Role                       pulumi.StringInput
}

func (CloudTasksQueueIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudTasksQueueIamMemberArgs)(nil)).Elem()
}

type CloudTasksQueueIamMemberInput interface {
	pulumi.Input

	ToCloudTasksQueueIamMemberOutput() CloudTasksQueueIamMemberOutput
	ToCloudTasksQueueIamMemberOutputWithContext(ctx context.Context) CloudTasksQueueIamMemberOutput
}

func (*CloudTasksQueueIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTasksQueueIamMember)(nil)).Elem()
}

func (i *CloudTasksQueueIamMember) ToCloudTasksQueueIamMemberOutput() CloudTasksQueueIamMemberOutput {
	return i.ToCloudTasksQueueIamMemberOutputWithContext(context.Background())
}

func (i *CloudTasksQueueIamMember) ToCloudTasksQueueIamMemberOutputWithContext(ctx context.Context) CloudTasksQueueIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTasksQueueIamMemberOutput)
}

type CloudTasksQueueIamMemberOutput struct{ *pulumi.OutputState }

func (CloudTasksQueueIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTasksQueueIamMember)(nil)).Elem()
}

func (o CloudTasksQueueIamMemberOutput) ToCloudTasksQueueIamMemberOutput() CloudTasksQueueIamMemberOutput {
	return o
}

func (o CloudTasksQueueIamMemberOutput) ToCloudTasksQueueIamMemberOutputWithContext(ctx context.Context) CloudTasksQueueIamMemberOutput {
	return o
}

func (o CloudTasksQueueIamMemberOutput) CloudTasksQueueIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) pulumi.StringOutput { return v.CloudTasksQueueIamMemberId }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamMemberOutput) Condition() CloudTasksQueueIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) CloudTasksQueueIamMemberConditionPtrOutput { return v.Condition }).(CloudTasksQueueIamMemberConditionPtrOutput)
}

func (o CloudTasksQueueIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o CloudTasksQueueIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudTasksQueueIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudTasksQueueIamMemberInput)(nil)).Elem(), &CloudTasksQueueIamMember{})
	pulumi.RegisterOutputType(CloudTasksQueueIamMemberOutput{})
}
