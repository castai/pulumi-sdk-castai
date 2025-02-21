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

type EndpointsServiceIamMember struct {
	pulumi.CustomResourceState

	Condition                   EndpointsServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	EndpointsServiceIamMemberId pulumi.StringOutput                         `pulumi:"endpointsServiceIamMemberId"`
	Etag                        pulumi.StringOutput                         `pulumi:"etag"`
	Member                      pulumi.StringOutput                         `pulumi:"member"`
	Role                        pulumi.StringOutput                         `pulumi:"role"`
	ServiceName                 pulumi.StringOutput                         `pulumi:"serviceName"`
}

// NewEndpointsServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewEndpointsServiceIamMember(ctx *pulumi.Context,
	name string, args *EndpointsServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*EndpointsServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource EndpointsServiceIamMember
	err = ctx.RegisterPackageResource("google-beta:index/endpointsServiceIamMember:EndpointsServiceIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointsServiceIamMember gets an existing EndpointsServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointsServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointsServiceIamMemberState, opts ...pulumi.ResourceOption) (*EndpointsServiceIamMember, error) {
	var resource EndpointsServiceIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/endpointsServiceIamMember:EndpointsServiceIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointsServiceIamMember resources.
type endpointsServiceIamMemberState struct {
	Condition                   *EndpointsServiceIamMemberCondition `pulumi:"condition"`
	EndpointsServiceIamMemberId *string                             `pulumi:"endpointsServiceIamMemberId"`
	Etag                        *string                             `pulumi:"etag"`
	Member                      *string                             `pulumi:"member"`
	Role                        *string                             `pulumi:"role"`
	ServiceName                 *string                             `pulumi:"serviceName"`
}

type EndpointsServiceIamMemberState struct {
	Condition                   EndpointsServiceIamMemberConditionPtrInput
	EndpointsServiceIamMemberId pulumi.StringPtrInput
	Etag                        pulumi.StringPtrInput
	Member                      pulumi.StringPtrInput
	Role                        pulumi.StringPtrInput
	ServiceName                 pulumi.StringPtrInput
}

func (EndpointsServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointsServiceIamMemberState)(nil)).Elem()
}

type endpointsServiceIamMemberArgs struct {
	Condition                   *EndpointsServiceIamMemberCondition `pulumi:"condition"`
	EndpointsServiceIamMemberId *string                             `pulumi:"endpointsServiceIamMemberId"`
	Member                      string                              `pulumi:"member"`
	Role                        string                              `pulumi:"role"`
	ServiceName                 string                              `pulumi:"serviceName"`
}

// The set of arguments for constructing a EndpointsServiceIamMember resource.
type EndpointsServiceIamMemberArgs struct {
	Condition                   EndpointsServiceIamMemberConditionPtrInput
	EndpointsServiceIamMemberId pulumi.StringPtrInput
	Member                      pulumi.StringInput
	Role                        pulumi.StringInput
	ServiceName                 pulumi.StringInput
}

func (EndpointsServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointsServiceIamMemberArgs)(nil)).Elem()
}

type EndpointsServiceIamMemberInput interface {
	pulumi.Input

	ToEndpointsServiceIamMemberOutput() EndpointsServiceIamMemberOutput
	ToEndpointsServiceIamMemberOutputWithContext(ctx context.Context) EndpointsServiceIamMemberOutput
}

func (*EndpointsServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsServiceIamMember)(nil)).Elem()
}

func (i *EndpointsServiceIamMember) ToEndpointsServiceIamMemberOutput() EndpointsServiceIamMemberOutput {
	return i.ToEndpointsServiceIamMemberOutputWithContext(context.Background())
}

func (i *EndpointsServiceIamMember) ToEndpointsServiceIamMemberOutputWithContext(ctx context.Context) EndpointsServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsServiceIamMemberOutput)
}

type EndpointsServiceIamMemberOutput struct{ *pulumi.OutputState }

func (EndpointsServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsServiceIamMember)(nil)).Elem()
}

func (o EndpointsServiceIamMemberOutput) ToEndpointsServiceIamMemberOutput() EndpointsServiceIamMemberOutput {
	return o
}

func (o EndpointsServiceIamMemberOutput) ToEndpointsServiceIamMemberOutputWithContext(ctx context.Context) EndpointsServiceIamMemberOutput {
	return o
}

func (o EndpointsServiceIamMemberOutput) Condition() EndpointsServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *EndpointsServiceIamMember) EndpointsServiceIamMemberConditionPtrOutput { return v.Condition }).(EndpointsServiceIamMemberConditionPtrOutput)
}

func (o EndpointsServiceIamMemberOutput) EndpointsServiceIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointsServiceIamMember) pulumi.StringOutput { return v.EndpointsServiceIamMemberId }).(pulumi.StringOutput)
}

func (o EndpointsServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointsServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o EndpointsServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointsServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o EndpointsServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointsServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o EndpointsServiceIamMemberOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointsServiceIamMember) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointsServiceIamMemberInput)(nil)).Elem(), &EndpointsServiceIamMember{})
	pulumi.RegisterOutputType(EndpointsServiceIamMemberOutput{})
}
