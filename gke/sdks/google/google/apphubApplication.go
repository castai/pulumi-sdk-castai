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

type ApphubApplication struct {
	pulumi.CustomResourceState

	ApphubApplicationId pulumi.StringOutput `pulumi:"apphubApplicationId"`
	// Required. The Application identifier.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Consumer provided attributes.
	Attributes ApphubApplicationAttributesPtrOutput `pulumi:"attributes"`
	// Output only. Create time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. User-defined description of an Application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Optional. User-defined name for the Application.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Part of 'parent'. See documentation of 'projectsId'.
	Location pulumi.StringOutput `pulumi:"location"`
	// Identifier. The resource name of an Application. Format:
	// "projects/{host-project-id}/locations/{location}/applications/{application-id}"
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Scope of an application.
	Scope ApphubApplicationScopeOutput `pulumi:"scope"`
	// Output only. Application state. Possible values: STATE_UNSPECIFIED CREATING ACTIVE DELETING
	State    pulumi.StringOutput                `pulumi:"state"`
	Timeouts ApphubApplicationTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. A universally unique identifier (in UUID4 format) for the 'Application'.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. Update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewApphubApplication registers a new resource with the given unique name, arguments, and options.
func NewApphubApplication(ctx *pulumi.Context,
	name string, args *ApphubApplicationArgs, opts ...pulumi.ResourceOption) (*ApphubApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApphubApplication
	err = ctx.RegisterPackageResource("google:index/apphubApplication:ApphubApplication", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApphubApplication gets an existing ApphubApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApphubApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApphubApplicationState, opts ...pulumi.ResourceOption) (*ApphubApplication, error) {
	var resource ApphubApplication
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apphubApplication:ApphubApplication", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApphubApplication resources.
type apphubApplicationState struct {
	ApphubApplicationId *string `pulumi:"apphubApplicationId"`
	// Required. The Application identifier.
	ApplicationId *string `pulumi:"applicationId"`
	// Consumer provided attributes.
	Attributes *ApphubApplicationAttributes `pulumi:"attributes"`
	// Output only. Create time.
	CreateTime *string `pulumi:"createTime"`
	// Optional. User-defined description of an Application.
	Description *string `pulumi:"description"`
	// Optional. User-defined name for the Application.
	DisplayName *string `pulumi:"displayName"`
	// Part of 'parent'. See documentation of 'projectsId'.
	Location *string `pulumi:"location"`
	// Identifier. The resource name of an Application. Format:
	// "projects/{host-project-id}/locations/{location}/applications/{application-id}"
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Scope of an application.
	Scope *ApphubApplicationScope `pulumi:"scope"`
	// Output only. Application state. Possible values: STATE_UNSPECIFIED CREATING ACTIVE DELETING
	State    *string                    `pulumi:"state"`
	Timeouts *ApphubApplicationTimeouts `pulumi:"timeouts"`
	// Output only. A universally unique identifier (in UUID4 format) for the 'Application'.
	Uid *string `pulumi:"uid"`
	// Output only. Update time.
	UpdateTime *string `pulumi:"updateTime"`
}

type ApphubApplicationState struct {
	ApphubApplicationId pulumi.StringPtrInput
	// Required. The Application identifier.
	ApplicationId pulumi.StringPtrInput
	// Consumer provided attributes.
	Attributes ApphubApplicationAttributesPtrInput
	// Output only. Create time.
	CreateTime pulumi.StringPtrInput
	// Optional. User-defined description of an Application.
	Description pulumi.StringPtrInput
	// Optional. User-defined name for the Application.
	DisplayName pulumi.StringPtrInput
	// Part of 'parent'. See documentation of 'projectsId'.
	Location pulumi.StringPtrInput
	// Identifier. The resource name of an Application. Format:
	// "projects/{host-project-id}/locations/{location}/applications/{application-id}"
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Scope of an application.
	Scope ApphubApplicationScopePtrInput
	// Output only. Application state. Possible values: STATE_UNSPECIFIED CREATING ACTIVE DELETING
	State    pulumi.StringPtrInput
	Timeouts ApphubApplicationTimeoutsPtrInput
	// Output only. A universally unique identifier (in UUID4 format) for the 'Application'.
	Uid pulumi.StringPtrInput
	// Output only. Update time.
	UpdateTime pulumi.StringPtrInput
}

func (ApphubApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*apphubApplicationState)(nil)).Elem()
}

type apphubApplicationArgs struct {
	ApphubApplicationId *string `pulumi:"apphubApplicationId"`
	// Required. The Application identifier.
	ApplicationId string `pulumi:"applicationId"`
	// Consumer provided attributes.
	Attributes *ApphubApplicationAttributes `pulumi:"attributes"`
	// Optional. User-defined description of an Application.
	Description *string `pulumi:"description"`
	// Optional. User-defined name for the Application.
	DisplayName *string `pulumi:"displayName"`
	// Part of 'parent'. See documentation of 'projectsId'.
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Scope of an application.
	Scope    ApphubApplicationScope     `pulumi:"scope"`
	Timeouts *ApphubApplicationTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApphubApplication resource.
type ApphubApplicationArgs struct {
	ApphubApplicationId pulumi.StringPtrInput
	// Required. The Application identifier.
	ApplicationId pulumi.StringInput
	// Consumer provided attributes.
	Attributes ApphubApplicationAttributesPtrInput
	// Optional. User-defined description of an Application.
	Description pulumi.StringPtrInput
	// Optional. User-defined name for the Application.
	DisplayName pulumi.StringPtrInput
	// Part of 'parent'. See documentation of 'projectsId'.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// Scope of an application.
	Scope    ApphubApplicationScopeInput
	Timeouts ApphubApplicationTimeoutsPtrInput
}

func (ApphubApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apphubApplicationArgs)(nil)).Elem()
}

type ApphubApplicationInput interface {
	pulumi.Input

	ToApphubApplicationOutput() ApphubApplicationOutput
	ToApphubApplicationOutputWithContext(ctx context.Context) ApphubApplicationOutput
}

func (*ApphubApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**ApphubApplication)(nil)).Elem()
}

func (i *ApphubApplication) ToApphubApplicationOutput() ApphubApplicationOutput {
	return i.ToApphubApplicationOutputWithContext(context.Background())
}

func (i *ApphubApplication) ToApphubApplicationOutputWithContext(ctx context.Context) ApphubApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApphubApplicationOutput)
}

type ApphubApplicationOutput struct{ *pulumi.OutputState }

func (ApphubApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApphubApplication)(nil)).Elem()
}

func (o ApphubApplicationOutput) ToApphubApplicationOutput() ApphubApplicationOutput {
	return o
}

func (o ApphubApplicationOutput) ToApphubApplicationOutputWithContext(ctx context.Context) ApphubApplicationOutput {
	return o
}

func (o ApphubApplicationOutput) ApphubApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.ApphubApplicationId }).(pulumi.StringOutput)
}

// Required. The Application identifier.
func (o ApphubApplicationOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// Consumer provided attributes.
func (o ApphubApplicationOutput) Attributes() ApphubApplicationAttributesPtrOutput {
	return o.ApplyT(func(v *ApphubApplication) ApphubApplicationAttributesPtrOutput { return v.Attributes }).(ApphubApplicationAttributesPtrOutput)
}

// Output only. Create time.
func (o ApphubApplicationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. User-defined description of an Application.
func (o ApphubApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Optional. User-defined name for the Application.
func (o ApphubApplicationOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Part of 'parent'. See documentation of 'projectsId'.
func (o ApphubApplicationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Identifier. The resource name of an Application. Format:
// "projects/{host-project-id}/locations/{location}/applications/{application-id}"
func (o ApphubApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApphubApplicationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Scope of an application.
func (o ApphubApplicationOutput) Scope() ApphubApplicationScopeOutput {
	return o.ApplyT(func(v *ApphubApplication) ApphubApplicationScopeOutput { return v.Scope }).(ApphubApplicationScopeOutput)
}

// Output only. Application state. Possible values: STATE_UNSPECIFIED CREATING ACTIVE DELETING
func (o ApphubApplicationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ApphubApplicationOutput) Timeouts() ApphubApplicationTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApphubApplication) ApphubApplicationTimeoutsPtrOutput { return v.Timeouts }).(ApphubApplicationTimeoutsPtrOutput)
}

// Output only. A universally unique identifier (in UUID4 format) for the 'Application'.
func (o ApphubApplicationOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. Update time.
func (o ApphubApplicationOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubApplication) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApphubApplicationInput)(nil)).Elem(), &ApphubApplication{})
	pulumi.RegisterOutputType(ApphubApplicationOutput{})
}
