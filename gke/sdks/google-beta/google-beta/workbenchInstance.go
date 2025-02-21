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

type WorkbenchInstance struct {
	pulumi.CustomResourceState

	// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
	// optional.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. Email address of entity that sent original CreateInstance request.
	Creator pulumi.StringOutput `pulumi:"creator"`
	// Desired state of the Workbench Instance. Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the
	// Instance.
	DesiredState pulumi.StringPtrOutput `pulumi:"desiredState"`
	// Optional. If true, the workbench instance will not register with the proxy.
	DisableProxyAccess pulumi.BoolPtrOutput   `pulumi:"disableProxyAccess"`
	EffectiveLabels    pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Flag that specifies that a notebook can be accessed with third party identity provider.
	EnableThirdPartyIdentity pulumi.BoolPtrOutput `pulumi:"enableThirdPartyIdentity"`
	// The definition of how to configure a VM instance outside of Resources and Identity.
	GceSetup WorkbenchInstanceGceSetupPtrOutput `pulumi:"gceSetup"`
	// 'Output only. Additional information about instance health. Example: healthInfo": { "docker_proxy_agent_status": "1",
	// "docker_status": "1", "jupyterlab_api_status": "-1", "jupyterlab_status": "-1", "updated": "2020-10-18 09:40:03.573409"
	// }'
	HealthInfos WorkbenchInstanceHealthInfoArrayOutput `pulumi:"healthInfos"`
	// Output only. Instance health_state.
	HealthState pulumi.StringOutput `pulumi:"healthState"`
	// Required. User-defined unique ID of this instance.
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// 'Optional. Input only. The owner of this instance after creation. Format: 'alias@example.com' Currently supports one
	// owner only. If not specified, all of the service account users of your VM instance''s service account can use the
	// instance. If specified, sets the access mode to 'Single user'. For more details, see
	// https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'
	InstanceOwners pulumi.StringArrayOutput `pulumi:"instanceOwners"`
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Part of 'parent'. See documentation of 'projectsId'.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. The proxy endpoint that is used to access the Jupyter notebook.
	ProxyUri pulumi.StringOutput `pulumi:"proxyUri"`
	// Output only. The state of this instance.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput             `pulumi:"terraformLabels"`
	Timeouts        WorkbenchInstanceTimeoutsPtrOutput `pulumi:"timeouts"`
	// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
	// optional.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Output only. The upgrade history of this instance.
	UpgradeHistories    WorkbenchInstanceUpgradeHistoryArrayOutput `pulumi:"upgradeHistories"`
	WorkbenchInstanceId pulumi.StringOutput                        `pulumi:"workbenchInstanceId"`
}

// NewWorkbenchInstance registers a new resource with the given unique name, arguments, and options.
func NewWorkbenchInstance(ctx *pulumi.Context,
	name string, args *WorkbenchInstanceArgs, opts ...pulumi.ResourceOption) (*WorkbenchInstance, error) {
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
	var resource WorkbenchInstance
	err = ctx.RegisterPackageResource("google-beta:index/workbenchInstance:WorkbenchInstance", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkbenchInstance gets an existing WorkbenchInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkbenchInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbenchInstanceState, opts ...pulumi.ResourceOption) (*WorkbenchInstance, error) {
	var resource WorkbenchInstance
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/workbenchInstance:WorkbenchInstance", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkbenchInstance resources.
type workbenchInstanceState struct {
	// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
	// optional.
	CreateTime *string `pulumi:"createTime"`
	// Output only. Email address of entity that sent original CreateInstance request.
	Creator *string `pulumi:"creator"`
	// Desired state of the Workbench Instance. Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the
	// Instance.
	DesiredState *string `pulumi:"desiredState"`
	// Optional. If true, the workbench instance will not register with the proxy.
	DisableProxyAccess *bool             `pulumi:"disableProxyAccess"`
	EffectiveLabels    map[string]string `pulumi:"effectiveLabels"`
	// Flag that specifies that a notebook can be accessed with third party identity provider.
	EnableThirdPartyIdentity *bool `pulumi:"enableThirdPartyIdentity"`
	// The definition of how to configure a VM instance outside of Resources and Identity.
	GceSetup *WorkbenchInstanceGceSetup `pulumi:"gceSetup"`
	// 'Output only. Additional information about instance health. Example: healthInfo": { "docker_proxy_agent_status": "1",
	// "docker_status": "1", "jupyterlab_api_status": "-1", "jupyterlab_status": "-1", "updated": "2020-10-18 09:40:03.573409"
	// }'
	HealthInfos []WorkbenchInstanceHealthInfo `pulumi:"healthInfos"`
	// Output only. Instance health_state.
	HealthState *string `pulumi:"healthState"`
	// Required. User-defined unique ID of this instance.
	InstanceId *string `pulumi:"instanceId"`
	// 'Optional. Input only. The owner of this instance after creation. Format: 'alias@example.com' Currently supports one
	// owner only. If not specified, all of the service account users of your VM instance''s service account can use the
	// instance. If specified, sets the access mode to 'Single user'. For more details, see
	// https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'
	InstanceOwners []string `pulumi:"instanceOwners"`
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Part of 'parent'. See documentation of 'projectsId'.
	Location *string `pulumi:"location"`
	// The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Output only. The proxy endpoint that is used to access the Jupyter notebook.
	ProxyUri *string `pulumi:"proxyUri"`
	// Output only. The state of this instance.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string          `pulumi:"terraformLabels"`
	Timeouts        *WorkbenchInstanceTimeouts `pulumi:"timeouts"`
	// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
	// optional.
	UpdateTime *string `pulumi:"updateTime"`
	// Output only. The upgrade history of this instance.
	UpgradeHistories    []WorkbenchInstanceUpgradeHistory `pulumi:"upgradeHistories"`
	WorkbenchInstanceId *string                           `pulumi:"workbenchInstanceId"`
}

type WorkbenchInstanceState struct {
	// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
	// optional.
	CreateTime pulumi.StringPtrInput
	// Output only. Email address of entity that sent original CreateInstance request.
	Creator pulumi.StringPtrInput
	// Desired state of the Workbench Instance. Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the
	// Instance.
	DesiredState pulumi.StringPtrInput
	// Optional. If true, the workbench instance will not register with the proxy.
	DisableProxyAccess pulumi.BoolPtrInput
	EffectiveLabels    pulumi.StringMapInput
	// Flag that specifies that a notebook can be accessed with third party identity provider.
	EnableThirdPartyIdentity pulumi.BoolPtrInput
	// The definition of how to configure a VM instance outside of Resources and Identity.
	GceSetup WorkbenchInstanceGceSetupPtrInput
	// 'Output only. Additional information about instance health. Example: healthInfo": { "docker_proxy_agent_status": "1",
	// "docker_status": "1", "jupyterlab_api_status": "-1", "jupyterlab_status": "-1", "updated": "2020-10-18 09:40:03.573409"
	// }'
	HealthInfos WorkbenchInstanceHealthInfoArrayInput
	// Output only. Instance health_state.
	HealthState pulumi.StringPtrInput
	// Required. User-defined unique ID of this instance.
	InstanceId pulumi.StringPtrInput
	// 'Optional. Input only. The owner of this instance after creation. Format: 'alias@example.com' Currently supports one
	// owner only. If not specified, all of the service account users of your VM instance''s service account can use the
	// instance. If specified, sets the access mode to 'Single user'. For more details, see
	// https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'
	InstanceOwners pulumi.StringArrayInput
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Part of 'parent'. See documentation of 'projectsId'.
	Location pulumi.StringPtrInput
	// The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Output only. The proxy endpoint that is used to access the Jupyter notebook.
	ProxyUri pulumi.StringPtrInput
	// Output only. The state of this instance.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        WorkbenchInstanceTimeoutsPtrInput
	// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
	// optional.
	UpdateTime pulumi.StringPtrInput
	// Output only. The upgrade history of this instance.
	UpgradeHistories    WorkbenchInstanceUpgradeHistoryArrayInput
	WorkbenchInstanceId pulumi.StringPtrInput
}

func (WorkbenchInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workbenchInstanceState)(nil)).Elem()
}

type workbenchInstanceArgs struct {
	// Desired state of the Workbench Instance. Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the
	// Instance.
	DesiredState *string `pulumi:"desiredState"`
	// Optional. If true, the workbench instance will not register with the proxy.
	DisableProxyAccess *bool `pulumi:"disableProxyAccess"`
	// Flag that specifies that a notebook can be accessed with third party identity provider.
	EnableThirdPartyIdentity *bool `pulumi:"enableThirdPartyIdentity"`
	// The definition of how to configure a VM instance outside of Resources and Identity.
	GceSetup *WorkbenchInstanceGceSetup `pulumi:"gceSetup"`
	// Required. User-defined unique ID of this instance.
	InstanceId *string `pulumi:"instanceId"`
	// 'Optional. Input only. The owner of this instance after creation. Format: 'alias@example.com' Currently supports one
	// owner only. If not specified, all of the service account users of your VM instance''s service account can use the
	// instance. If specified, sets the access mode to 'Single user'. For more details, see
	// https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'
	InstanceOwners []string `pulumi:"instanceOwners"`
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Part of 'parent'. See documentation of 'projectsId'.
	Location string `pulumi:"location"`
	// The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'
	Name                *string                    `pulumi:"name"`
	Project             *string                    `pulumi:"project"`
	Timeouts            *WorkbenchInstanceTimeouts `pulumi:"timeouts"`
	WorkbenchInstanceId *string                    `pulumi:"workbenchInstanceId"`
}

// The set of arguments for constructing a WorkbenchInstance resource.
type WorkbenchInstanceArgs struct {
	// Desired state of the Workbench Instance. Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the
	// Instance.
	DesiredState pulumi.StringPtrInput
	// Optional. If true, the workbench instance will not register with the proxy.
	DisableProxyAccess pulumi.BoolPtrInput
	// Flag that specifies that a notebook can be accessed with third party identity provider.
	EnableThirdPartyIdentity pulumi.BoolPtrInput
	// The definition of how to configure a VM instance outside of Resources and Identity.
	GceSetup WorkbenchInstanceGceSetupPtrInput
	// Required. User-defined unique ID of this instance.
	InstanceId pulumi.StringPtrInput
	// 'Optional. Input only. The owner of this instance after creation. Format: 'alias@example.com' Currently supports one
	// owner only. If not specified, all of the service account users of your VM instance''s service account can use the
	// instance. If specified, sets the access mode to 'Single user'. For more details, see
	// https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'
	InstanceOwners pulumi.StringArrayInput
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Part of 'parent'. See documentation of 'projectsId'.
	Location pulumi.StringInput
	// The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'
	Name                pulumi.StringPtrInput
	Project             pulumi.StringPtrInput
	Timeouts            WorkbenchInstanceTimeoutsPtrInput
	WorkbenchInstanceId pulumi.StringPtrInput
}

func (WorkbenchInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workbenchInstanceArgs)(nil)).Elem()
}

type WorkbenchInstanceInput interface {
	pulumi.Input

	ToWorkbenchInstanceOutput() WorkbenchInstanceOutput
	ToWorkbenchInstanceOutputWithContext(ctx context.Context) WorkbenchInstanceOutput
}

func (*WorkbenchInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbenchInstance)(nil)).Elem()
}

func (i *WorkbenchInstance) ToWorkbenchInstanceOutput() WorkbenchInstanceOutput {
	return i.ToWorkbenchInstanceOutputWithContext(context.Background())
}

func (i *WorkbenchInstance) ToWorkbenchInstanceOutputWithContext(ctx context.Context) WorkbenchInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbenchInstanceOutput)
}

type WorkbenchInstanceOutput struct{ *pulumi.OutputState }

func (WorkbenchInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbenchInstance)(nil)).Elem()
}

func (o WorkbenchInstanceOutput) ToWorkbenchInstanceOutput() WorkbenchInstanceOutput {
	return o
}

func (o WorkbenchInstanceOutput) ToWorkbenchInstanceOutputWithContext(ctx context.Context) WorkbenchInstanceOutput {
	return o
}

// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
// optional.
func (o WorkbenchInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. Email address of entity that sent original CreateInstance request.
func (o WorkbenchInstanceOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.Creator }).(pulumi.StringOutput)
}

// Desired state of the Workbench Instance. Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the
// Instance.
func (o WorkbenchInstanceOutput) DesiredState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringPtrOutput { return v.DesiredState }).(pulumi.StringPtrOutput)
}

// Optional. If true, the workbench instance will not register with the proxy.
func (o WorkbenchInstanceOutput) DisableProxyAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.BoolPtrOutput { return v.DisableProxyAccess }).(pulumi.BoolPtrOutput)
}

func (o WorkbenchInstanceOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Flag that specifies that a notebook can be accessed with third party identity provider.
func (o WorkbenchInstanceOutput) EnableThirdPartyIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.BoolPtrOutput { return v.EnableThirdPartyIdentity }).(pulumi.BoolPtrOutput)
}

// The definition of how to configure a VM instance outside of Resources and Identity.
func (o WorkbenchInstanceOutput) GceSetup() WorkbenchInstanceGceSetupPtrOutput {
	return o.ApplyT(func(v *WorkbenchInstance) WorkbenchInstanceGceSetupPtrOutput { return v.GceSetup }).(WorkbenchInstanceGceSetupPtrOutput)
}

// 'Output only. Additional information about instance health. Example: healthInfo": { "docker_proxy_agent_status": "1",
// "docker_status": "1", "jupyterlab_api_status": "-1", "jupyterlab_status": "-1", "updated": "2020-10-18 09:40:03.573409"
// }'
func (o WorkbenchInstanceOutput) HealthInfos() WorkbenchInstanceHealthInfoArrayOutput {
	return o.ApplyT(func(v *WorkbenchInstance) WorkbenchInstanceHealthInfoArrayOutput { return v.HealthInfos }).(WorkbenchInstanceHealthInfoArrayOutput)
}

// Output only. Instance health_state.
func (o WorkbenchInstanceOutput) HealthState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.HealthState }).(pulumi.StringOutput)
}

// Required. User-defined unique ID of this instance.
func (o WorkbenchInstanceOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringPtrOutput { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// 'Optional. Input only. The owner of this instance after creation. Format: 'alias@example.com' Currently supports one
// owner only. If not specified, all of the service account users of your VM instance”s service account can use the
// instance. If specified, sets the access mode to 'Single user'. For more details, see
// https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'
func (o WorkbenchInstanceOutput) InstanceOwners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringArrayOutput { return v.InstanceOwners }).(pulumi.StringArrayOutput)
}

// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method. **Note**: This
// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
// 'effective_labels' for all of the labels present on the resource.
func (o WorkbenchInstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Part of 'parent'. See documentation of 'projectsId'.
func (o WorkbenchInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'
func (o WorkbenchInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkbenchInstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. The proxy endpoint that is used to access the Jupyter notebook.
func (o WorkbenchInstanceOutput) ProxyUri() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.ProxyUri }).(pulumi.StringOutput)
}

// Output only. The state of this instance.
func (o WorkbenchInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o WorkbenchInstanceOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o WorkbenchInstanceOutput) Timeouts() WorkbenchInstanceTimeoutsPtrOutput {
	return o.ApplyT(func(v *WorkbenchInstance) WorkbenchInstanceTimeoutsPtrOutput { return v.Timeouts }).(WorkbenchInstanceTimeoutsPtrOutput)
}

// An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ. The milliseconds portion (".SSS") is
// optional.
func (o WorkbenchInstanceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Output only. The upgrade history of this instance.
func (o WorkbenchInstanceOutput) UpgradeHistories() WorkbenchInstanceUpgradeHistoryArrayOutput {
	return o.ApplyT(func(v *WorkbenchInstance) WorkbenchInstanceUpgradeHistoryArrayOutput { return v.UpgradeHistories }).(WorkbenchInstanceUpgradeHistoryArrayOutput)
}

func (o WorkbenchInstanceOutput) WorkbenchInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbenchInstance) pulumi.StringOutput { return v.WorkbenchInstanceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbenchInstanceInput)(nil)).Elem(), &WorkbenchInstance{})
	pulumi.RegisterOutputType(WorkbenchInstanceOutput{})
}
