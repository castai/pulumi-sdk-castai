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

type NotebooksEnvironment struct {
	pulumi.CustomResourceState

	// Use a container image to start the notebook instance.
	ContainerImage NotebooksEnvironmentContainerImagePtrOutput `pulumi:"containerImage"`
	// Instance creation time
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A brief description of this environment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name specified for the Environment instance. Format:
	// projects/{project_id}/locations/{location}/environments/{environmentId}
	Name                   pulumi.StringOutput `pulumi:"name"`
	NotebooksEnvironmentId pulumi.StringOutput `pulumi:"notebooksEnvironmentId"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud
	// Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrOutput                `pulumi:"postStartupScript"`
	Project           pulumi.StringOutput                   `pulumi:"project"`
	Timeouts          NotebooksEnvironmentTimeoutsPtrOutput `pulumi:"timeouts"`
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage NotebooksEnvironmentVmImagePtrOutput `pulumi:"vmImage"`
}

// NewNotebooksEnvironment registers a new resource with the given unique name, arguments, and options.
func NewNotebooksEnvironment(ctx *pulumi.Context,
	name string, args *NotebooksEnvironmentArgs, opts ...pulumi.ResourceOption) (*NotebooksEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NotebooksEnvironment
	err = ctx.RegisterPackageResource("google:index/notebooksEnvironment:NotebooksEnvironment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebooksEnvironment gets an existing NotebooksEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebooksEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebooksEnvironmentState, opts ...pulumi.ResourceOption) (*NotebooksEnvironment, error) {
	var resource NotebooksEnvironment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/notebooksEnvironment:NotebooksEnvironment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebooksEnvironment resources.
type notebooksEnvironmentState struct {
	// Use a container image to start the notebook instance.
	ContainerImage *NotebooksEnvironmentContainerImage `pulumi:"containerImage"`
	// Instance creation time
	CreateTime *string `pulumi:"createTime"`
	// A brief description of this environment.
	Description *string `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName *string `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location *string `pulumi:"location"`
	// The name specified for the Environment instance. Format:
	// projects/{project_id}/locations/{location}/environments/{environmentId}
	Name                   *string `pulumi:"name"`
	NotebooksEnvironmentId *string `pulumi:"notebooksEnvironmentId"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud
	// Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript *string                       `pulumi:"postStartupScript"`
	Project           *string                       `pulumi:"project"`
	Timeouts          *NotebooksEnvironmentTimeouts `pulumi:"timeouts"`
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage *NotebooksEnvironmentVmImage `pulumi:"vmImage"`
}

type NotebooksEnvironmentState struct {
	// Use a container image to start the notebook instance.
	ContainerImage NotebooksEnvironmentContainerImagePtrInput
	// Instance creation time
	CreateTime pulumi.StringPtrInput
	// A brief description of this environment.
	Description pulumi.StringPtrInput
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringPtrInput
	// The name specified for the Environment instance. Format:
	// projects/{project_id}/locations/{location}/environments/{environmentId}
	Name                   pulumi.StringPtrInput
	NotebooksEnvironmentId pulumi.StringPtrInput
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud
	// Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrInput
	Project           pulumi.StringPtrInput
	Timeouts          NotebooksEnvironmentTimeoutsPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage NotebooksEnvironmentVmImagePtrInput
}

func (NotebooksEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksEnvironmentState)(nil)).Elem()
}

type notebooksEnvironmentArgs struct {
	// Use a container image to start the notebook instance.
	ContainerImage *NotebooksEnvironmentContainerImage `pulumi:"containerImage"`
	// A brief description of this environment.
	Description *string `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName *string `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location string `pulumi:"location"`
	// The name specified for the Environment instance. Format:
	// projects/{project_id}/locations/{location}/environments/{environmentId}
	Name                   *string `pulumi:"name"`
	NotebooksEnvironmentId *string `pulumi:"notebooksEnvironmentId"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud
	// Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript *string                       `pulumi:"postStartupScript"`
	Project           *string                       `pulumi:"project"`
	Timeouts          *NotebooksEnvironmentTimeouts `pulumi:"timeouts"`
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage *NotebooksEnvironmentVmImage `pulumi:"vmImage"`
}

// The set of arguments for constructing a NotebooksEnvironment resource.
type NotebooksEnvironmentArgs struct {
	// Use a container image to start the notebook instance.
	ContainerImage NotebooksEnvironmentContainerImagePtrInput
	// A brief description of this environment.
	Description pulumi.StringPtrInput
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringInput
	// The name specified for the Environment instance. Format:
	// projects/{project_id}/locations/{location}/environments/{environmentId}
	Name                   pulumi.StringPtrInput
	NotebooksEnvironmentId pulumi.StringPtrInput
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud
	// Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrInput
	Project           pulumi.StringPtrInput
	Timeouts          NotebooksEnvironmentTimeoutsPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage NotebooksEnvironmentVmImagePtrInput
}

func (NotebooksEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksEnvironmentArgs)(nil)).Elem()
}

type NotebooksEnvironmentInput interface {
	pulumi.Input

	ToNotebooksEnvironmentOutput() NotebooksEnvironmentOutput
	ToNotebooksEnvironmentOutputWithContext(ctx context.Context) NotebooksEnvironmentOutput
}

func (*NotebooksEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksEnvironment)(nil)).Elem()
}

func (i *NotebooksEnvironment) ToNotebooksEnvironmentOutput() NotebooksEnvironmentOutput {
	return i.ToNotebooksEnvironmentOutputWithContext(context.Background())
}

func (i *NotebooksEnvironment) ToNotebooksEnvironmentOutputWithContext(ctx context.Context) NotebooksEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebooksEnvironmentOutput)
}

type NotebooksEnvironmentOutput struct{ *pulumi.OutputState }

func (NotebooksEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksEnvironment)(nil)).Elem()
}

func (o NotebooksEnvironmentOutput) ToNotebooksEnvironmentOutput() NotebooksEnvironmentOutput {
	return o
}

func (o NotebooksEnvironmentOutput) ToNotebooksEnvironmentOutputWithContext(ctx context.Context) NotebooksEnvironmentOutput {
	return o
}

// Use a container image to start the notebook instance.
func (o NotebooksEnvironmentOutput) ContainerImage() NotebooksEnvironmentContainerImagePtrOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) NotebooksEnvironmentContainerImagePtrOutput { return v.ContainerImage }).(NotebooksEnvironmentContainerImagePtrOutput)
}

// Instance creation time
func (o NotebooksEnvironmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A brief description of this environment.
func (o NotebooksEnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of this environment for the UI.
func (o NotebooksEnvironmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// A reference to the zone where the machine resides.
func (o NotebooksEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name specified for the Environment instance. Format:
// projects/{project_id}/locations/{location}/environments/{environmentId}
func (o NotebooksEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NotebooksEnvironmentOutput) NotebooksEnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringOutput { return v.NotebooksEnvironmentId }).(pulumi.StringOutput)
}

// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud
// Storage path. Example: "gs://path-to-file/file-name"
func (o NotebooksEnvironmentOutput) PostStartupScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringPtrOutput { return v.PostStartupScript }).(pulumi.StringPtrOutput)
}

func (o NotebooksEnvironmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NotebooksEnvironmentOutput) Timeouts() NotebooksEnvironmentTimeoutsPtrOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) NotebooksEnvironmentTimeoutsPtrOutput { return v.Timeouts }).(NotebooksEnvironmentTimeoutsPtrOutput)
}

// Use a Compute Engine VM image to start the notebook instance.
func (o NotebooksEnvironmentOutput) VmImage() NotebooksEnvironmentVmImagePtrOutput {
	return o.ApplyT(func(v *NotebooksEnvironment) NotebooksEnvironmentVmImagePtrOutput { return v.VmImage }).(NotebooksEnvironmentVmImagePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotebooksEnvironmentInput)(nil)).Elem(), &NotebooksEnvironment{})
	pulumi.RegisterOutputType(NotebooksEnvironmentOutput{})
}
