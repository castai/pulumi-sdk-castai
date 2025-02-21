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

type ApphubServiceProjectAttachment struct {
	pulumi.CustomResourceState

	ApphubServiceProjectAttachmentId pulumi.StringOutput `pulumi:"apphubServiceProjectAttachmentId"`
	// Output only. Create time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// "Identifier. The resource name of a ServiceProjectAttachment.
	// Format:\"projects/{host-project-id}/locations/global/serviceProjectAttachments/{service-project-id}.\""
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// "Immutable. Service project name in the format: \"projects/abc\" or \"projects/123\". As input, project name with either
	// project id or number are accepted. As output, this field will contain project number."
	ServiceProject pulumi.StringPtrOutput `pulumi:"serviceProject"`
	// Required. The service project attachment identifier must contain the project_id of the service project specified in the
	// service_project_attachment.service_project field. Hint: "projects/{project_id}"
	ServiceProjectAttachmentId pulumi.StringOutput `pulumi:"serviceProjectAttachmentId"`
	// ServiceProjectAttachment state.
	State    pulumi.StringOutput                             `pulumi:"state"`
	Timeouts ApphubServiceProjectAttachmentTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. A globally unique identifier (in UUID4 format) for the 'ServiceProjectAttachment'.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewApphubServiceProjectAttachment registers a new resource with the given unique name, arguments, and options.
func NewApphubServiceProjectAttachment(ctx *pulumi.Context,
	name string, args *ApphubServiceProjectAttachmentArgs, opts ...pulumi.ResourceOption) (*ApphubServiceProjectAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceProjectAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceProjectAttachmentId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApphubServiceProjectAttachment
	err = ctx.RegisterPackageResource("google:index/apphubServiceProjectAttachment:ApphubServiceProjectAttachment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApphubServiceProjectAttachment gets an existing ApphubServiceProjectAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApphubServiceProjectAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApphubServiceProjectAttachmentState, opts ...pulumi.ResourceOption) (*ApphubServiceProjectAttachment, error) {
	var resource ApphubServiceProjectAttachment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apphubServiceProjectAttachment:ApphubServiceProjectAttachment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApphubServiceProjectAttachment resources.
type apphubServiceProjectAttachmentState struct {
	ApphubServiceProjectAttachmentId *string `pulumi:"apphubServiceProjectAttachmentId"`
	// Output only. Create time.
	CreateTime *string `pulumi:"createTime"`
	// "Identifier. The resource name of a ServiceProjectAttachment.
	// Format:\"projects/{host-project-id}/locations/global/serviceProjectAttachments/{service-project-id}.\""
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// "Immutable. Service project name in the format: \"projects/abc\" or \"projects/123\". As input, project name with either
	// project id or number are accepted. As output, this field will contain project number."
	ServiceProject *string `pulumi:"serviceProject"`
	// Required. The service project attachment identifier must contain the project_id of the service project specified in the
	// service_project_attachment.service_project field. Hint: "projects/{project_id}"
	ServiceProjectAttachmentId *string `pulumi:"serviceProjectAttachmentId"`
	// ServiceProjectAttachment state.
	State    *string                                 `pulumi:"state"`
	Timeouts *ApphubServiceProjectAttachmentTimeouts `pulumi:"timeouts"`
	// Output only. A globally unique identifier (in UUID4 format) for the 'ServiceProjectAttachment'.
	Uid *string `pulumi:"uid"`
}

type ApphubServiceProjectAttachmentState struct {
	ApphubServiceProjectAttachmentId pulumi.StringPtrInput
	// Output only. Create time.
	CreateTime pulumi.StringPtrInput
	// "Identifier. The resource name of a ServiceProjectAttachment.
	// Format:\"projects/{host-project-id}/locations/global/serviceProjectAttachments/{service-project-id}.\""
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// "Immutable. Service project name in the format: \"projects/abc\" or \"projects/123\". As input, project name with either
	// project id or number are accepted. As output, this field will contain project number."
	ServiceProject pulumi.StringPtrInput
	// Required. The service project attachment identifier must contain the project_id of the service project specified in the
	// service_project_attachment.service_project field. Hint: "projects/{project_id}"
	ServiceProjectAttachmentId pulumi.StringPtrInput
	// ServiceProjectAttachment state.
	State    pulumi.StringPtrInput
	Timeouts ApphubServiceProjectAttachmentTimeoutsPtrInput
	// Output only. A globally unique identifier (in UUID4 format) for the 'ServiceProjectAttachment'.
	Uid pulumi.StringPtrInput
}

func (ApphubServiceProjectAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apphubServiceProjectAttachmentState)(nil)).Elem()
}

type apphubServiceProjectAttachmentArgs struct {
	ApphubServiceProjectAttachmentId *string `pulumi:"apphubServiceProjectAttachmentId"`
	Project                          *string `pulumi:"project"`
	// "Immutable. Service project name in the format: \"projects/abc\" or \"projects/123\". As input, project name with either
	// project id or number are accepted. As output, this field will contain project number."
	ServiceProject *string `pulumi:"serviceProject"`
	// Required. The service project attachment identifier must contain the project_id of the service project specified in the
	// service_project_attachment.service_project field. Hint: "projects/{project_id}"
	ServiceProjectAttachmentId string                                  `pulumi:"serviceProjectAttachmentId"`
	Timeouts                   *ApphubServiceProjectAttachmentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApphubServiceProjectAttachment resource.
type ApphubServiceProjectAttachmentArgs struct {
	ApphubServiceProjectAttachmentId pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
	// "Immutable. Service project name in the format: \"projects/abc\" or \"projects/123\". As input, project name with either
	// project id or number are accepted. As output, this field will contain project number."
	ServiceProject pulumi.StringPtrInput
	// Required. The service project attachment identifier must contain the project_id of the service project specified in the
	// service_project_attachment.service_project field. Hint: "projects/{project_id}"
	ServiceProjectAttachmentId pulumi.StringInput
	Timeouts                   ApphubServiceProjectAttachmentTimeoutsPtrInput
}

func (ApphubServiceProjectAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apphubServiceProjectAttachmentArgs)(nil)).Elem()
}

type ApphubServiceProjectAttachmentInput interface {
	pulumi.Input

	ToApphubServiceProjectAttachmentOutput() ApphubServiceProjectAttachmentOutput
	ToApphubServiceProjectAttachmentOutputWithContext(ctx context.Context) ApphubServiceProjectAttachmentOutput
}

func (*ApphubServiceProjectAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApphubServiceProjectAttachment)(nil)).Elem()
}

func (i *ApphubServiceProjectAttachment) ToApphubServiceProjectAttachmentOutput() ApphubServiceProjectAttachmentOutput {
	return i.ToApphubServiceProjectAttachmentOutputWithContext(context.Background())
}

func (i *ApphubServiceProjectAttachment) ToApphubServiceProjectAttachmentOutputWithContext(ctx context.Context) ApphubServiceProjectAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApphubServiceProjectAttachmentOutput)
}

type ApphubServiceProjectAttachmentOutput struct{ *pulumi.OutputState }

func (ApphubServiceProjectAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApphubServiceProjectAttachment)(nil)).Elem()
}

func (o ApphubServiceProjectAttachmentOutput) ToApphubServiceProjectAttachmentOutput() ApphubServiceProjectAttachmentOutput {
	return o
}

func (o ApphubServiceProjectAttachmentOutput) ToApphubServiceProjectAttachmentOutputWithContext(ctx context.Context) ApphubServiceProjectAttachmentOutput {
	return o
}

func (o ApphubServiceProjectAttachmentOutput) ApphubServiceProjectAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringOutput { return v.ApphubServiceProjectAttachmentId }).(pulumi.StringOutput)
}

// Output only. Create time.
func (o ApphubServiceProjectAttachmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// "Identifier. The resource name of a ServiceProjectAttachment.
// Format:\"projects/{host-project-id}/locations/global/serviceProjectAttachments/{service-project-id}.\""
func (o ApphubServiceProjectAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApphubServiceProjectAttachmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// "Immutable. Service project name in the format: \"projects/abc\" or \"projects/123\". As input, project name with either
// project id or number are accepted. As output, this field will contain project number."
func (o ApphubServiceProjectAttachmentOutput) ServiceProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringPtrOutput { return v.ServiceProject }).(pulumi.StringPtrOutput)
}

// Required. The service project attachment identifier must contain the project_id of the service project specified in the
// service_project_attachment.service_project field. Hint: "projects/{project_id}"
func (o ApphubServiceProjectAttachmentOutput) ServiceProjectAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringOutput { return v.ServiceProjectAttachmentId }).(pulumi.StringOutput)
}

// ServiceProjectAttachment state.
func (o ApphubServiceProjectAttachmentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ApphubServiceProjectAttachmentOutput) Timeouts() ApphubServiceProjectAttachmentTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) ApphubServiceProjectAttachmentTimeoutsPtrOutput {
		return v.Timeouts
	}).(ApphubServiceProjectAttachmentTimeoutsPtrOutput)
}

// Output only. A globally unique identifier (in UUID4 format) for the 'ServiceProjectAttachment'.
func (o ApphubServiceProjectAttachmentOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ApphubServiceProjectAttachment) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApphubServiceProjectAttachmentInput)(nil)).Elem(), &ApphubServiceProjectAttachment{})
	pulumi.RegisterOutputType(ApphubServiceProjectAttachmentOutput{})
}
