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

type NetworkManagementVpcFlowLogsConfig struct {
	pulumi.CustomResourceState

	// Optional. The aggregation interval for the logs. Default value is INTERVAL_5_SEC. Possible values:
	// AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN
	// INTERVAL_15_MIN"
	AggregationInterval pulumi.StringOutput `pulumi:"aggregationInterval"`
	// Output only. The time the config was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Optional. Export filter used to define which VPC Flow Logs should be logged.
	FilterExpr pulumi.StringPtrOutput `pulumi:"filterExpr"`
	// Optional. The value of the field must be in (0, 1]. The sampling rate of VPC Flow Logs where 1.0 means all collected
	// logs are reported. Setting the sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use the state
	// field instead. Default value is 1.0.
	FlowSampling pulumi.Float64Output `pulumi:"flowSampling"`
	// Traffic will be logged from the Interconnect Attachment. Format:
	// projects/{project_id}/regions/{region}/interconnectAttachments/{name}
	InterconnectAttachment pulumi.StringPtrOutput `pulumi:"interconnectAttachment"`
	// Optional. Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.
	Location pulumi.StringOutput `pulumi:"location"`
	// Optional. Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs.
	// Default value is INCLUDE_ALL_METADATA. Possible values: METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA
	// CUSTOM_METADATA
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// Optional. Custom metadata fields to include in the reported VPC flow logs. Can only be specified if \"metadata\" was set
	// to CUSTOM_METADATA.
	MetadataFields pulumi.StringArrayOutput `pulumi:"metadataFields"`
	// Identifier. Unique name of the configuration using the form:
	// 'projects/{project_id}/locations/global/vpcFlowLogsConfigs/{vpc_flow_logs_config_id}'
	Name                                 pulumi.StringOutput `pulumi:"name"`
	NetworkManagementVpcFlowLogsConfigId pulumi.StringOutput `pulumi:"networkManagementVpcFlowLogsConfigId"`
	Project                              pulumi.StringOutput `pulumi:"project"`
	// Optional. The state of the VPC Flow Log configuration. Default value is ENABLED. When creating a new configuration, it
	// must be enabled. Possible
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                              `pulumi:"terraformLabels"`
	Timeouts        NetworkManagementVpcFlowLogsConfigTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. The time the config was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Required. ID of the 'VpcFlowLogsConfig'.
	VpcFlowLogsConfigId pulumi.StringOutput `pulumi:"vpcFlowLogsConfigId"`
	// Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}
	VpnTunnel pulumi.StringPtrOutput `pulumi:"vpnTunnel"`
}

// NewNetworkManagementVpcFlowLogsConfig registers a new resource with the given unique name, arguments, and options.
func NewNetworkManagementVpcFlowLogsConfig(ctx *pulumi.Context,
	name string, args *NetworkManagementVpcFlowLogsConfigArgs, opts ...pulumi.ResourceOption) (*NetworkManagementVpcFlowLogsConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.VpcFlowLogsConfigId == nil {
		return nil, errors.New("invalid value for required argument 'VpcFlowLogsConfigId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkManagementVpcFlowLogsConfig
	err = ctx.RegisterPackageResource("google-beta:index/networkManagementVpcFlowLogsConfig:NetworkManagementVpcFlowLogsConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkManagementVpcFlowLogsConfig gets an existing NetworkManagementVpcFlowLogsConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkManagementVpcFlowLogsConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkManagementVpcFlowLogsConfigState, opts ...pulumi.ResourceOption) (*NetworkManagementVpcFlowLogsConfig, error) {
	var resource NetworkManagementVpcFlowLogsConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkManagementVpcFlowLogsConfig:NetworkManagementVpcFlowLogsConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkManagementVpcFlowLogsConfig resources.
type networkManagementVpcFlowLogsConfigState struct {
	// Optional. The aggregation interval for the logs. Default value is INTERVAL_5_SEC. Possible values:
	// AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN
	// INTERVAL_15_MIN"
	AggregationInterval *string `pulumi:"aggregationInterval"`
	// Output only. The time the config was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Optional. Export filter used to define which VPC Flow Logs should be logged.
	FilterExpr *string `pulumi:"filterExpr"`
	// Optional. The value of the field must be in (0, 1]. The sampling rate of VPC Flow Logs where 1.0 means all collected
	// logs are reported. Setting the sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use the state
	// field instead. Default value is 1.0.
	FlowSampling *float64 `pulumi:"flowSampling"`
	// Traffic will be logged from the Interconnect Attachment. Format:
	// projects/{project_id}/regions/{region}/interconnectAttachments/{name}
	InterconnectAttachment *string `pulumi:"interconnectAttachment"`
	// Optional. Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.
	Location *string `pulumi:"location"`
	// Optional. Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs.
	// Default value is INCLUDE_ALL_METADATA. Possible values: METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA
	// CUSTOM_METADATA
	Metadata *string `pulumi:"metadata"`
	// Optional. Custom metadata fields to include in the reported VPC flow logs. Can only be specified if \"metadata\" was set
	// to CUSTOM_METADATA.
	MetadataFields []string `pulumi:"metadataFields"`
	// Identifier. Unique name of the configuration using the form:
	// 'projects/{project_id}/locations/global/vpcFlowLogsConfigs/{vpc_flow_logs_config_id}'
	Name                                 *string `pulumi:"name"`
	NetworkManagementVpcFlowLogsConfigId *string `pulumi:"networkManagementVpcFlowLogsConfigId"`
	Project                              *string `pulumi:"project"`
	// Optional. The state of the VPC Flow Log configuration. Default value is ENABLED. When creating a new configuration, it
	// must be enabled. Possible
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                           `pulumi:"terraformLabels"`
	Timeouts        *NetworkManagementVpcFlowLogsConfigTimeouts `pulumi:"timeouts"`
	// Output only. The time the config was updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Required. ID of the 'VpcFlowLogsConfig'.
	VpcFlowLogsConfigId *string `pulumi:"vpcFlowLogsConfigId"`
	// Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}
	VpnTunnel *string `pulumi:"vpnTunnel"`
}

type NetworkManagementVpcFlowLogsConfigState struct {
	// Optional. The aggregation interval for the logs. Default value is INTERVAL_5_SEC. Possible values:
	// AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN
	// INTERVAL_15_MIN"
	AggregationInterval pulumi.StringPtrInput
	// Output only. The time the config was created.
	CreateTime pulumi.StringPtrInput
	// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Optional. Export filter used to define which VPC Flow Logs should be logged.
	FilterExpr pulumi.StringPtrInput
	// Optional. The value of the field must be in (0, 1]. The sampling rate of VPC Flow Logs where 1.0 means all collected
	// logs are reported. Setting the sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use the state
	// field instead. Default value is 1.0.
	FlowSampling pulumi.Float64PtrInput
	// Traffic will be logged from the Interconnect Attachment. Format:
	// projects/{project_id}/regions/{region}/interconnectAttachments/{name}
	InterconnectAttachment pulumi.StringPtrInput
	// Optional. Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.
	Location pulumi.StringPtrInput
	// Optional. Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs.
	// Default value is INCLUDE_ALL_METADATA. Possible values: METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA
	// CUSTOM_METADATA
	Metadata pulumi.StringPtrInput
	// Optional. Custom metadata fields to include in the reported VPC flow logs. Can only be specified if \"metadata\" was set
	// to CUSTOM_METADATA.
	MetadataFields pulumi.StringArrayInput
	// Identifier. Unique name of the configuration using the form:
	// 'projects/{project_id}/locations/global/vpcFlowLogsConfigs/{vpc_flow_logs_config_id}'
	Name                                 pulumi.StringPtrInput
	NetworkManagementVpcFlowLogsConfigId pulumi.StringPtrInput
	Project                              pulumi.StringPtrInput
	// Optional. The state of the VPC Flow Log configuration. Default value is ENABLED. When creating a new configuration, it
	// must be enabled. Possible
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkManagementVpcFlowLogsConfigTimeoutsPtrInput
	// Output only. The time the config was updated.
	UpdateTime pulumi.StringPtrInput
	// Required. ID of the 'VpcFlowLogsConfig'.
	VpcFlowLogsConfigId pulumi.StringPtrInput
	// Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}
	VpnTunnel pulumi.StringPtrInput
}

func (NetworkManagementVpcFlowLogsConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkManagementVpcFlowLogsConfigState)(nil)).Elem()
}

type networkManagementVpcFlowLogsConfigArgs struct {
	// Optional. The aggregation interval for the logs. Default value is INTERVAL_5_SEC. Possible values:
	// AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN
	// INTERVAL_15_MIN"
	AggregationInterval *string `pulumi:"aggregationInterval"`
	// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Optional. Export filter used to define which VPC Flow Logs should be logged.
	FilterExpr *string `pulumi:"filterExpr"`
	// Optional. The value of the field must be in (0, 1]. The sampling rate of VPC Flow Logs where 1.0 means all collected
	// logs are reported. Setting the sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use the state
	// field instead. Default value is 1.0.
	FlowSampling *float64 `pulumi:"flowSampling"`
	// Traffic will be logged from the Interconnect Attachment. Format:
	// projects/{project_id}/regions/{region}/interconnectAttachments/{name}
	InterconnectAttachment *string `pulumi:"interconnectAttachment"`
	// Optional. Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.
	Location string `pulumi:"location"`
	// Optional. Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs.
	// Default value is INCLUDE_ALL_METADATA. Possible values: METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA
	// CUSTOM_METADATA
	Metadata *string `pulumi:"metadata"`
	// Optional. Custom metadata fields to include in the reported VPC flow logs. Can only be specified if \"metadata\" was set
	// to CUSTOM_METADATA.
	MetadataFields                       []string `pulumi:"metadataFields"`
	NetworkManagementVpcFlowLogsConfigId *string  `pulumi:"networkManagementVpcFlowLogsConfigId"`
	Project                              *string  `pulumi:"project"`
	// Optional. The state of the VPC Flow Log configuration. Default value is ENABLED. When creating a new configuration, it
	// must be enabled. Possible
	State    *string                                     `pulumi:"state"`
	Timeouts *NetworkManagementVpcFlowLogsConfigTimeouts `pulumi:"timeouts"`
	// Required. ID of the 'VpcFlowLogsConfig'.
	VpcFlowLogsConfigId string `pulumi:"vpcFlowLogsConfigId"`
	// Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}
	VpnTunnel *string `pulumi:"vpnTunnel"`
}

// The set of arguments for constructing a NetworkManagementVpcFlowLogsConfig resource.
type NetworkManagementVpcFlowLogsConfigArgs struct {
	// Optional. The aggregation interval for the logs. Default value is INTERVAL_5_SEC. Possible values:
	// AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN
	// INTERVAL_15_MIN"
	AggregationInterval pulumi.StringPtrInput
	// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Optional. Export filter used to define which VPC Flow Logs should be logged.
	FilterExpr pulumi.StringPtrInput
	// Optional. The value of the field must be in (0, 1]. The sampling rate of VPC Flow Logs where 1.0 means all collected
	// logs are reported. Setting the sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use the state
	// field instead. Default value is 1.0.
	FlowSampling pulumi.Float64PtrInput
	// Traffic will be logged from the Interconnect Attachment. Format:
	// projects/{project_id}/regions/{region}/interconnectAttachments/{name}
	InterconnectAttachment pulumi.StringPtrInput
	// Optional. Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.
	Location pulumi.StringInput
	// Optional. Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs.
	// Default value is INCLUDE_ALL_METADATA. Possible values: METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA
	// CUSTOM_METADATA
	Metadata pulumi.StringPtrInput
	// Optional. Custom metadata fields to include in the reported VPC flow logs. Can only be specified if \"metadata\" was set
	// to CUSTOM_METADATA.
	MetadataFields                       pulumi.StringArrayInput
	NetworkManagementVpcFlowLogsConfigId pulumi.StringPtrInput
	Project                              pulumi.StringPtrInput
	// Optional. The state of the VPC Flow Log configuration. Default value is ENABLED. When creating a new configuration, it
	// must be enabled. Possible
	State    pulumi.StringPtrInput
	Timeouts NetworkManagementVpcFlowLogsConfigTimeoutsPtrInput
	// Required. ID of the 'VpcFlowLogsConfig'.
	VpcFlowLogsConfigId pulumi.StringInput
	// Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}
	VpnTunnel pulumi.StringPtrInput
}

func (NetworkManagementVpcFlowLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkManagementVpcFlowLogsConfigArgs)(nil)).Elem()
}

type NetworkManagementVpcFlowLogsConfigInput interface {
	pulumi.Input

	ToNetworkManagementVpcFlowLogsConfigOutput() NetworkManagementVpcFlowLogsConfigOutput
	ToNetworkManagementVpcFlowLogsConfigOutputWithContext(ctx context.Context) NetworkManagementVpcFlowLogsConfigOutput
}

func (*NetworkManagementVpcFlowLogsConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManagementVpcFlowLogsConfig)(nil)).Elem()
}

func (i *NetworkManagementVpcFlowLogsConfig) ToNetworkManagementVpcFlowLogsConfigOutput() NetworkManagementVpcFlowLogsConfigOutput {
	return i.ToNetworkManagementVpcFlowLogsConfigOutputWithContext(context.Background())
}

func (i *NetworkManagementVpcFlowLogsConfig) ToNetworkManagementVpcFlowLogsConfigOutputWithContext(ctx context.Context) NetworkManagementVpcFlowLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagementVpcFlowLogsConfigOutput)
}

type NetworkManagementVpcFlowLogsConfigOutput struct{ *pulumi.OutputState }

func (NetworkManagementVpcFlowLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManagementVpcFlowLogsConfig)(nil)).Elem()
}

func (o NetworkManagementVpcFlowLogsConfigOutput) ToNetworkManagementVpcFlowLogsConfigOutput() NetworkManagementVpcFlowLogsConfigOutput {
	return o
}

func (o NetworkManagementVpcFlowLogsConfigOutput) ToNetworkManagementVpcFlowLogsConfigOutputWithContext(ctx context.Context) NetworkManagementVpcFlowLogsConfigOutput {
	return o
}

// Optional. The aggregation interval for the logs. Default value is INTERVAL_5_SEC. Possible values:
// AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN
// INTERVAL_15_MIN"
func (o NetworkManagementVpcFlowLogsConfigOutput) AggregationInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.AggregationInterval }).(pulumi.StringOutput)
}

// Output only. The time the config was created.
func (o NetworkManagementVpcFlowLogsConfigOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
func (o NetworkManagementVpcFlowLogsConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkManagementVpcFlowLogsConfigOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Optional. Export filter used to define which VPC Flow Logs should be logged.
func (o NetworkManagementVpcFlowLogsConfigOutput) FilterExpr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringPtrOutput { return v.FilterExpr }).(pulumi.StringPtrOutput)
}

// Optional. The value of the field must be in (0, 1]. The sampling rate of VPC Flow Logs where 1.0 means all collected
// logs are reported. Setting the sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use the state
// field instead. Default value is 1.0.
func (o NetworkManagementVpcFlowLogsConfigOutput) FlowSampling() pulumi.Float64Output {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.Float64Output { return v.FlowSampling }).(pulumi.Float64Output)
}

// Traffic will be logged from the Interconnect Attachment. Format:
// projects/{project_id}/regions/{region}/interconnectAttachments/{name}
func (o NetworkManagementVpcFlowLogsConfigOutput) InterconnectAttachment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringPtrOutput { return v.InterconnectAttachment }).(pulumi.StringPtrOutput)
}

// Optional. Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkManagementVpcFlowLogsConfigOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
// https://google.aip.dev/122. See documentation for resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.
func (o NetworkManagementVpcFlowLogsConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs.
// Default value is INCLUDE_ALL_METADATA. Possible values: METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA
// CUSTOM_METADATA
func (o NetworkManagementVpcFlowLogsConfigOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.Metadata }).(pulumi.StringOutput)
}

// Optional. Custom metadata fields to include in the reported VPC flow logs. Can only be specified if \"metadata\" was set
// to CUSTOM_METADATA.
func (o NetworkManagementVpcFlowLogsConfigOutput) MetadataFields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringArrayOutput { return v.MetadataFields }).(pulumi.StringArrayOutput)
}

// Identifier. Unique name of the configuration using the form:
// 'projects/{project_id}/locations/global/vpcFlowLogsConfigs/{vpc_flow_logs_config_id}'
func (o NetworkManagementVpcFlowLogsConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkManagementVpcFlowLogsConfigOutput) NetworkManagementVpcFlowLogsConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput {
		return v.NetworkManagementVpcFlowLogsConfigId
	}).(pulumi.StringOutput)
}

func (o NetworkManagementVpcFlowLogsConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. The state of the VPC Flow Log configuration. Default value is ENABLED. When creating a new configuration, it
// must be enabled. Possible
func (o NetworkManagementVpcFlowLogsConfigOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkManagementVpcFlowLogsConfigOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkManagementVpcFlowLogsConfigOutput) Timeouts() NetworkManagementVpcFlowLogsConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) NetworkManagementVpcFlowLogsConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkManagementVpcFlowLogsConfigTimeoutsPtrOutput)
}

// Output only. The time the config was updated.
func (o NetworkManagementVpcFlowLogsConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Required. ID of the 'VpcFlowLogsConfig'.
func (o NetworkManagementVpcFlowLogsConfigOutput) VpcFlowLogsConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringOutput { return v.VpcFlowLogsConfigId }).(pulumi.StringOutput)
}

// Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}
func (o NetworkManagementVpcFlowLogsConfigOutput) VpnTunnel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkManagementVpcFlowLogsConfig) pulumi.StringPtrOutput { return v.VpnTunnel }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkManagementVpcFlowLogsConfigInput)(nil)).Elem(), &NetworkManagementVpcFlowLogsConfig{})
	pulumi.RegisterOutputType(NetworkManagementVpcFlowLogsConfigOutput{})
}
