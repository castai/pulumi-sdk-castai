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

type ProjectDefaultServiceAccounts struct {
	pulumi.CustomResourceState

	// The action to be performed in the default service accounts. Valid values are: DEPRIVILEGE, DELETE, DISABLE. Note that
	// DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.
	Action pulumi.StringOutput `pulumi:"action"`
	// The project ID where service accounts are created.
	Project                         pulumi.StringOutput `pulumi:"project"`
	ProjectDefaultServiceAccountsId pulumi.StringOutput `pulumi:"projectDefaultServiceAccountsId"`
	// The action to be performed in the default service accounts on the resource destroy. Valid values are NONE, REVERT and
	// REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	RestorePolicy pulumi.StringPtrOutput `pulumi:"restorePolicy"`
	// The Service Accounts changed by this resource. It is used for revert the action on the destroy.
	ServiceAccounts pulumi.StringMapOutput                         `pulumi:"serviceAccounts"`
	Timeouts        ProjectDefaultServiceAccountsTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewProjectDefaultServiceAccounts registers a new resource with the given unique name, arguments, and options.
func NewProjectDefaultServiceAccounts(ctx *pulumi.Context,
	name string, args *ProjectDefaultServiceAccountsArgs, opts ...pulumi.ResourceOption) (*ProjectDefaultServiceAccounts, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ProjectDefaultServiceAccounts
	err = ctx.RegisterPackageResource("google:index/projectDefaultServiceAccounts:ProjectDefaultServiceAccounts", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectDefaultServiceAccounts gets an existing ProjectDefaultServiceAccounts resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectDefaultServiceAccounts(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectDefaultServiceAccountsState, opts ...pulumi.ResourceOption) (*ProjectDefaultServiceAccounts, error) {
	var resource ProjectDefaultServiceAccounts
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/projectDefaultServiceAccounts:ProjectDefaultServiceAccounts", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectDefaultServiceAccounts resources.
type projectDefaultServiceAccountsState struct {
	// The action to be performed in the default service accounts. Valid values are: DEPRIVILEGE, DELETE, DISABLE. Note that
	// DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.
	Action *string `pulumi:"action"`
	// The project ID where service accounts are created.
	Project                         *string `pulumi:"project"`
	ProjectDefaultServiceAccountsId *string `pulumi:"projectDefaultServiceAccountsId"`
	// The action to be performed in the default service accounts on the resource destroy. Valid values are NONE, REVERT and
	// REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	RestorePolicy *string `pulumi:"restorePolicy"`
	// The Service Accounts changed by this resource. It is used for revert the action on the destroy.
	ServiceAccounts map[string]string                      `pulumi:"serviceAccounts"`
	Timeouts        *ProjectDefaultServiceAccountsTimeouts `pulumi:"timeouts"`
}

type ProjectDefaultServiceAccountsState struct {
	// The action to be performed in the default service accounts. Valid values are: DEPRIVILEGE, DELETE, DISABLE. Note that
	// DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.
	Action pulumi.StringPtrInput
	// The project ID where service accounts are created.
	Project                         pulumi.StringPtrInput
	ProjectDefaultServiceAccountsId pulumi.StringPtrInput
	// The action to be performed in the default service accounts on the resource destroy. Valid values are NONE, REVERT and
	// REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	RestorePolicy pulumi.StringPtrInput
	// The Service Accounts changed by this resource. It is used for revert the action on the destroy.
	ServiceAccounts pulumi.StringMapInput
	Timeouts        ProjectDefaultServiceAccountsTimeoutsPtrInput
}

func (ProjectDefaultServiceAccountsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectDefaultServiceAccountsState)(nil)).Elem()
}

type projectDefaultServiceAccountsArgs struct {
	// The action to be performed in the default service accounts. Valid values are: DEPRIVILEGE, DELETE, DISABLE. Note that
	// DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.
	Action string `pulumi:"action"`
	// The project ID where service accounts are created.
	Project                         string  `pulumi:"project"`
	ProjectDefaultServiceAccountsId *string `pulumi:"projectDefaultServiceAccountsId"`
	// The action to be performed in the default service accounts on the resource destroy. Valid values are NONE, REVERT and
	// REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	RestorePolicy *string                                `pulumi:"restorePolicy"`
	Timeouts      *ProjectDefaultServiceAccountsTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ProjectDefaultServiceAccounts resource.
type ProjectDefaultServiceAccountsArgs struct {
	// The action to be performed in the default service accounts. Valid values are: DEPRIVILEGE, DELETE, DISABLE. Note that
	// DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.
	Action pulumi.StringInput
	// The project ID where service accounts are created.
	Project                         pulumi.StringInput
	ProjectDefaultServiceAccountsId pulumi.StringPtrInput
	// The action to be performed in the default service accounts on the resource destroy. Valid values are NONE, REVERT and
	// REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	RestorePolicy pulumi.StringPtrInput
	Timeouts      ProjectDefaultServiceAccountsTimeoutsPtrInput
}

func (ProjectDefaultServiceAccountsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectDefaultServiceAccountsArgs)(nil)).Elem()
}

type ProjectDefaultServiceAccountsInput interface {
	pulumi.Input

	ToProjectDefaultServiceAccountsOutput() ProjectDefaultServiceAccountsOutput
	ToProjectDefaultServiceAccountsOutputWithContext(ctx context.Context) ProjectDefaultServiceAccountsOutput
}

func (*ProjectDefaultServiceAccounts) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectDefaultServiceAccounts)(nil)).Elem()
}

func (i *ProjectDefaultServiceAccounts) ToProjectDefaultServiceAccountsOutput() ProjectDefaultServiceAccountsOutput {
	return i.ToProjectDefaultServiceAccountsOutputWithContext(context.Background())
}

func (i *ProjectDefaultServiceAccounts) ToProjectDefaultServiceAccountsOutputWithContext(ctx context.Context) ProjectDefaultServiceAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectDefaultServiceAccountsOutput)
}

type ProjectDefaultServiceAccountsOutput struct{ *pulumi.OutputState }

func (ProjectDefaultServiceAccountsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectDefaultServiceAccounts)(nil)).Elem()
}

func (o ProjectDefaultServiceAccountsOutput) ToProjectDefaultServiceAccountsOutput() ProjectDefaultServiceAccountsOutput {
	return o
}

func (o ProjectDefaultServiceAccountsOutput) ToProjectDefaultServiceAccountsOutputWithContext(ctx context.Context) ProjectDefaultServiceAccountsOutput {
	return o
}

// The action to be performed in the default service accounts. Valid values are: DEPRIVILEGE, DELETE, DISABLE. Note that
// DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.
func (o ProjectDefaultServiceAccountsOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDefaultServiceAccounts) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The project ID where service accounts are created.
func (o ProjectDefaultServiceAccountsOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDefaultServiceAccounts) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ProjectDefaultServiceAccountsOutput) ProjectDefaultServiceAccountsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDefaultServiceAccounts) pulumi.StringOutput { return v.ProjectDefaultServiceAccountsId }).(pulumi.StringOutput)
}

// The action to be performed in the default service accounts on the resource destroy. Valid values are NONE, REVERT and
// REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
func (o ProjectDefaultServiceAccountsOutput) RestorePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectDefaultServiceAccounts) pulumi.StringPtrOutput { return v.RestorePolicy }).(pulumi.StringPtrOutput)
}

// The Service Accounts changed by this resource. It is used for revert the action on the destroy.
func (o ProjectDefaultServiceAccountsOutput) ServiceAccounts() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProjectDefaultServiceAccounts) pulumi.StringMapOutput { return v.ServiceAccounts }).(pulumi.StringMapOutput)
}

func (o ProjectDefaultServiceAccountsOutput) Timeouts() ProjectDefaultServiceAccountsTimeoutsPtrOutput {
	return o.ApplyT(func(v *ProjectDefaultServiceAccounts) ProjectDefaultServiceAccountsTimeoutsPtrOutput {
		return v.Timeouts
	}).(ProjectDefaultServiceAccountsTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectDefaultServiceAccountsInput)(nil)).Elem(), &ProjectDefaultServiceAccounts{})
	pulumi.RegisterOutputType(ProjectDefaultServiceAccountsOutput{})
}
