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

type AccessContextManagerAccessLevel struct {
	pulumi.CustomResourceState

	AccessContextManagerAccessLevelId pulumi.StringOutput `pulumi:"accessContextManagerAccessLevelId"`
	// A set of predefined conditions for the access level and a combining function.
	Basic AccessContextManagerAccessLevelBasicPtrOutput `pulumi:"basic"`
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions
	// for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec.
	Custom AccessContextManagerAccessLevelCustomPtrOutput `pulumi:"custom"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringOutput `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent   pulumi.StringOutput                              `pulumi:"parent"`
	Timeouts AccessContextManagerAccessLevelTimeoutsPtrOutput `pulumi:"timeouts"`
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewAccessContextManagerAccessLevel registers a new resource with the given unique name, arguments, and options.
func NewAccessContextManagerAccessLevel(ctx *pulumi.Context,
	name string, args *AccessContextManagerAccessLevelArgs, opts ...pulumi.ResourceOption) (*AccessContextManagerAccessLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource AccessContextManagerAccessLevel
	err = ctx.RegisterPackageResource("google:index/accessContextManagerAccessLevel:AccessContextManagerAccessLevel", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessContextManagerAccessLevel gets an existing AccessContextManagerAccessLevel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessContextManagerAccessLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessContextManagerAccessLevelState, opts ...pulumi.ResourceOption) (*AccessContextManagerAccessLevel, error) {
	var resource AccessContextManagerAccessLevel
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/accessContextManagerAccessLevel:AccessContextManagerAccessLevel", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessContextManagerAccessLevel resources.
type accessContextManagerAccessLevelState struct {
	AccessContextManagerAccessLevelId *string `pulumi:"accessContextManagerAccessLevelId"`
	// A set of predefined conditions for the access level and a combining function.
	Basic *AccessContextManagerAccessLevelBasic `pulumi:"basic"`
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions
	// for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec.
	Custom *AccessContextManagerAccessLevelCustom `pulumi:"custom"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description *string `pulumi:"description"`
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent   *string                                  `pulumi:"parent"`
	Timeouts *AccessContextManagerAccessLevelTimeouts `pulumi:"timeouts"`
	// Human readable title. Must be unique within the Policy.
	Title *string `pulumi:"title"`
}

type AccessContextManagerAccessLevelState struct {
	AccessContextManagerAccessLevelId pulumi.StringPtrInput
	// A set of predefined conditions for the access level and a combining function.
	Basic AccessContextManagerAccessLevelBasicPtrInput
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions
	// for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec.
	Custom AccessContextManagerAccessLevelCustomPtrInput
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description pulumi.StringPtrInput
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent   pulumi.StringPtrInput
	Timeouts AccessContextManagerAccessLevelTimeoutsPtrInput
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringPtrInput
}

func (AccessContextManagerAccessLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerAccessLevelState)(nil)).Elem()
}

type accessContextManagerAccessLevelArgs struct {
	AccessContextManagerAccessLevelId *string `pulumi:"accessContextManagerAccessLevelId"`
	// A set of predefined conditions for the access level and a combining function.
	Basic *AccessContextManagerAccessLevelBasic `pulumi:"basic"`
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions
	// for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec.
	Custom *AccessContextManagerAccessLevelCustom `pulumi:"custom"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description *string `pulumi:"description"`
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent   string                                   `pulumi:"parent"`
	Timeouts *AccessContextManagerAccessLevelTimeouts `pulumi:"timeouts"`
	// Human readable title. Must be unique within the Policy.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a AccessContextManagerAccessLevel resource.
type AccessContextManagerAccessLevelArgs struct {
	AccessContextManagerAccessLevelId pulumi.StringPtrInput
	// A set of predefined conditions for the access level and a combining function.
	Basic AccessContextManagerAccessLevelBasicPtrInput
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions
	// for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec.
	Custom AccessContextManagerAccessLevelCustomPtrInput
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description pulumi.StringPtrInput
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent   pulumi.StringInput
	Timeouts AccessContextManagerAccessLevelTimeoutsPtrInput
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringInput
}

func (AccessContextManagerAccessLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerAccessLevelArgs)(nil)).Elem()
}

type AccessContextManagerAccessLevelInput interface {
	pulumi.Input

	ToAccessContextManagerAccessLevelOutput() AccessContextManagerAccessLevelOutput
	ToAccessContextManagerAccessLevelOutputWithContext(ctx context.Context) AccessContextManagerAccessLevelOutput
}

func (*AccessContextManagerAccessLevel) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerAccessLevel)(nil)).Elem()
}

func (i *AccessContextManagerAccessLevel) ToAccessContextManagerAccessLevelOutput() AccessContextManagerAccessLevelOutput {
	return i.ToAccessContextManagerAccessLevelOutputWithContext(context.Background())
}

func (i *AccessContextManagerAccessLevel) ToAccessContextManagerAccessLevelOutputWithContext(ctx context.Context) AccessContextManagerAccessLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessContextManagerAccessLevelOutput)
}

type AccessContextManagerAccessLevelOutput struct{ *pulumi.OutputState }

func (AccessContextManagerAccessLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerAccessLevel)(nil)).Elem()
}

func (o AccessContextManagerAccessLevelOutput) ToAccessContextManagerAccessLevelOutput() AccessContextManagerAccessLevelOutput {
	return o
}

func (o AccessContextManagerAccessLevelOutput) ToAccessContextManagerAccessLevelOutputWithContext(ctx context.Context) AccessContextManagerAccessLevelOutput {
	return o
}

func (o AccessContextManagerAccessLevelOutput) AccessContextManagerAccessLevelId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) pulumi.StringOutput {
		return v.AccessContextManagerAccessLevelId
	}).(pulumi.StringOutput)
}

// A set of predefined conditions for the access level and a combining function.
func (o AccessContextManagerAccessLevelOutput) Basic() AccessContextManagerAccessLevelBasicPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) AccessContextManagerAccessLevelBasicPtrOutput { return v.Basic }).(AccessContextManagerAccessLevelBasicPtrOutput)
}

// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions
// for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec.
func (o AccessContextManagerAccessLevelOutput) Custom() AccessContextManagerAccessLevelCustomPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) AccessContextManagerAccessLevelCustomPtrOutput {
		return v.Custom
	}).(AccessContextManagerAccessLevelCustomPtrOutput)
}

// Description of the AccessLevel and its use. Does not affect behavior.
func (o AccessContextManagerAccessLevelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
func (o AccessContextManagerAccessLevelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
func (o AccessContextManagerAccessLevelOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

func (o AccessContextManagerAccessLevelOutput) Timeouts() AccessContextManagerAccessLevelTimeoutsPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) AccessContextManagerAccessLevelTimeoutsPtrOutput {
		return v.Timeouts
	}).(AccessContextManagerAccessLevelTimeoutsPtrOutput)
}

// Human readable title. Must be unique within the Policy.
func (o AccessContextManagerAccessLevelOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerAccessLevel) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessContextManagerAccessLevelInput)(nil)).Elem(), &AccessContextManagerAccessLevel{})
	pulumi.RegisterOutputType(AccessContextManagerAccessLevelOutput{})
}
