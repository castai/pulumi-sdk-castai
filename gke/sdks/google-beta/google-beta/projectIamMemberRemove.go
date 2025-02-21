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

type ProjectIamMemberRemove struct {
	pulumi.CustomResourceState

	// The IAM principal that should not have the target role.
	Member pulumi.StringOutput `pulumi:"member"`
	// The project id of the target project.
	Project                  pulumi.StringOutput `pulumi:"project"`
	ProjectIamMemberRemoveId pulumi.StringOutput `pulumi:"projectIamMemberRemoveId"`
	// The target role that should be removed.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewProjectIamMemberRemove registers a new resource with the given unique name, arguments, and options.
func NewProjectIamMemberRemove(ctx *pulumi.Context,
	name string, args *ProjectIamMemberRemoveArgs, opts ...pulumi.ResourceOption) (*ProjectIamMemberRemove, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ProjectIamMemberRemove
	err = ctx.RegisterPackageResource("google-beta:index/projectIamMemberRemove:ProjectIamMemberRemove", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectIamMemberRemove gets an existing ProjectIamMemberRemove resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectIamMemberRemove(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectIamMemberRemoveState, opts ...pulumi.ResourceOption) (*ProjectIamMemberRemove, error) {
	var resource ProjectIamMemberRemove
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/projectIamMemberRemove:ProjectIamMemberRemove", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectIamMemberRemove resources.
type projectIamMemberRemoveState struct {
	// The IAM principal that should not have the target role.
	Member *string `pulumi:"member"`
	// The project id of the target project.
	Project                  *string `pulumi:"project"`
	ProjectIamMemberRemoveId *string `pulumi:"projectIamMemberRemoveId"`
	// The target role that should be removed.
	Role *string `pulumi:"role"`
}

type ProjectIamMemberRemoveState struct {
	// The IAM principal that should not have the target role.
	Member pulumi.StringPtrInput
	// The project id of the target project.
	Project                  pulumi.StringPtrInput
	ProjectIamMemberRemoveId pulumi.StringPtrInput
	// The target role that should be removed.
	Role pulumi.StringPtrInput
}

func (ProjectIamMemberRemoveState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIamMemberRemoveState)(nil)).Elem()
}

type projectIamMemberRemoveArgs struct {
	// The IAM principal that should not have the target role.
	Member string `pulumi:"member"`
	// The project id of the target project.
	Project                  string  `pulumi:"project"`
	ProjectIamMemberRemoveId *string `pulumi:"projectIamMemberRemoveId"`
	// The target role that should be removed.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ProjectIamMemberRemove resource.
type ProjectIamMemberRemoveArgs struct {
	// The IAM principal that should not have the target role.
	Member pulumi.StringInput
	// The project id of the target project.
	Project                  pulumi.StringInput
	ProjectIamMemberRemoveId pulumi.StringPtrInput
	// The target role that should be removed.
	Role pulumi.StringInput
}

func (ProjectIamMemberRemoveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIamMemberRemoveArgs)(nil)).Elem()
}

type ProjectIamMemberRemoveInput interface {
	pulumi.Input

	ToProjectIamMemberRemoveOutput() ProjectIamMemberRemoveOutput
	ToProjectIamMemberRemoveOutputWithContext(ctx context.Context) ProjectIamMemberRemoveOutput
}

func (*ProjectIamMemberRemove) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIamMemberRemove)(nil)).Elem()
}

func (i *ProjectIamMemberRemove) ToProjectIamMemberRemoveOutput() ProjectIamMemberRemoveOutput {
	return i.ToProjectIamMemberRemoveOutputWithContext(context.Background())
}

func (i *ProjectIamMemberRemove) ToProjectIamMemberRemoveOutputWithContext(ctx context.Context) ProjectIamMemberRemoveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIamMemberRemoveOutput)
}

type ProjectIamMemberRemoveOutput struct{ *pulumi.OutputState }

func (ProjectIamMemberRemoveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIamMemberRemove)(nil)).Elem()
}

func (o ProjectIamMemberRemoveOutput) ToProjectIamMemberRemoveOutput() ProjectIamMemberRemoveOutput {
	return o
}

func (o ProjectIamMemberRemoveOutput) ToProjectIamMemberRemoveOutputWithContext(ctx context.Context) ProjectIamMemberRemoveOutput {
	return o
}

// The IAM principal that should not have the target role.
func (o ProjectIamMemberRemoveOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIamMemberRemove) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The project id of the target project.
func (o ProjectIamMemberRemoveOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIamMemberRemove) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ProjectIamMemberRemoveOutput) ProjectIamMemberRemoveId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIamMemberRemove) pulumi.StringOutput { return v.ProjectIamMemberRemoveId }).(pulumi.StringOutput)
}

// The target role that should be removed.
func (o ProjectIamMemberRemoveOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIamMemberRemove) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIamMemberRemoveInput)(nil)).Elem(), &ProjectIamMemberRemove{})
	pulumi.RegisterOutputType(ProjectIamMemberRemoveOutput{})
}
