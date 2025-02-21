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

type VertexAiEndpoint struct {
	pulumi.CustomResourceState

	// Output only. Timestamp when this Endpoint was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. DNS of the dedicated endpoint. Will only be populated if dedicatedEndpointEnabled is true. Format:
	// 'https://{endpointId}.{region}-{projectNumber}.prediction.vertexai.goog'.
	DedicatedEndpointDns pulumi.StringOutput `pulumi:"dedicatedEndpointDns"`
	// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the
	// dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once
	// you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com.
	// The limitation will be removed soon.
	DedicatedEndpointEnabled pulumi.BoolPtrOutput `pulumi:"dedicatedEndpointEnabled"`
	// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and
	// EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud
	// Console](https://console.cloud.google.com/vertex-ai/).
	DeployedModels VertexAiEndpointDeployedModelArrayOutput `pulumi:"deployedModels"`
	// The description of the Endpoint.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName     pulumi.StringOutput    `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will
	// be secured by this key.
	EncryptionSpec VertexAiEndpointEncryptionSpecPtrOutput `pulumi:"encryptionSpec"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64
	// characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes.
	// International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. **Note**:
	// This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the
	// field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by
	// CreateModelDeploymentMonitoringJob. Format:
	// 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'
	ModelDeploymentMonitoringJob pulumi.StringOutput `pulumi:"modelDeploymentMonitoringJob"`
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name pulumi.StringOutput `pulumi:"name"`
	// The full name of the Google Compute Engine
	// [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be
	// peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not
	// peered with any network. Only one of the fields, network or enable_private_service_connect, can be set.
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	// 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is
	// network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// Configures the request-response logging for online prediction.
	PredictRequestResponseLoggingConfig VertexAiEndpointPredictRequestResponseLoggingConfigPtrOutput `pulumi:"predictRequestResponseLoggingConfig"`
	// Configuration for private service connect. 'network' and 'privateServiceConnectConfig' are mutually exclusive.
	PrivateServiceConnectConfig VertexAiEndpointPrivateServiceConnectConfigPtrOutput `pulumi:"privateServiceConnectConfig"`
	Project                     pulumi.StringOutput                                  `pulumi:"project"`
	// The region for the resource
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput            `pulumi:"terraformLabels"`
	Timeouts        VertexAiEndpointTimeoutsPtrOutput `pulumi:"timeouts"`
	// A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that
	// DeployedModel. If a DeployedModel's id is not listed in this map, then it receives no traffic. The traffic percentage
	// values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See the
	// 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and
	// [documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel)
	// for more information. > **Note:** To set the map to empty, set '"{}"', apply, and then remove the field from your
	// config.
	TrafficSplit pulumi.StringOutput `pulumi:"trafficSplit"`
	// Output only. Timestamp when this Endpoint was last updated.
	UpdateTime         pulumi.StringOutput `pulumi:"updateTime"`
	VertexAiEndpointId pulumi.StringOutput `pulumi:"vertexAiEndpointId"`
}

// NewVertexAiEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVertexAiEndpoint(ctx *pulumi.Context,
	name string, args *VertexAiEndpointArgs, opts ...pulumi.ResourceOption) (*VertexAiEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource VertexAiEndpoint
	err = ctx.RegisterPackageResource("google-beta:index/vertexAiEndpoint:VertexAiEndpoint", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVertexAiEndpoint gets an existing VertexAiEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVertexAiEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VertexAiEndpointState, opts ...pulumi.ResourceOption) (*VertexAiEndpoint, error) {
	var resource VertexAiEndpoint
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/vertexAiEndpoint:VertexAiEndpoint", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VertexAiEndpoint resources.
type vertexAiEndpointState struct {
	// Output only. Timestamp when this Endpoint was created.
	CreateTime *string `pulumi:"createTime"`
	// Output only. DNS of the dedicated endpoint. Will only be populated if dedicatedEndpointEnabled is true. Format:
	// 'https://{endpointId}.{region}-{projectNumber}.prediction.vertexai.goog'.
	DedicatedEndpointDns *string `pulumi:"dedicatedEndpointDns"`
	// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the
	// dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once
	// you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com.
	// The limitation will be removed soon.
	DedicatedEndpointEnabled *bool `pulumi:"dedicatedEndpointEnabled"`
	// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and
	// EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud
	// Console](https://console.cloud.google.com/vertex-ai/).
	DeployedModels []VertexAiEndpointDeployedModel `pulumi:"deployedModels"`
	// The description of the Endpoint.
	Description *string `pulumi:"description"`
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will
	// be secured by this key.
	EncryptionSpec *VertexAiEndpointEncryptionSpec `pulumi:"encryptionSpec"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag *string `pulumi:"etag"`
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64
	// characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes.
	// International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. **Note**:
	// This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the
	// field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by
	// CreateModelDeploymentMonitoringJob. Format:
	// 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'
	ModelDeploymentMonitoringJob *string `pulumi:"modelDeploymentMonitoringJob"`
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name *string `pulumi:"name"`
	// The full name of the Google Compute Engine
	// [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be
	// peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not
	// peered with any network. Only one of the fields, network or enable_private_service_connect, can be set.
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	// 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is
	// network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.
	Network *string `pulumi:"network"`
	// Configures the request-response logging for online prediction.
	PredictRequestResponseLoggingConfig *VertexAiEndpointPredictRequestResponseLoggingConfig `pulumi:"predictRequestResponseLoggingConfig"`
	// Configuration for private service connect. 'network' and 'privateServiceConnectConfig' are mutually exclusive.
	PrivateServiceConnectConfig *VertexAiEndpointPrivateServiceConnectConfig `pulumi:"privateServiceConnectConfig"`
	Project                     *string                                      `pulumi:"project"`
	// The region for the resource
	Region *string `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string         `pulumi:"terraformLabels"`
	Timeouts        *VertexAiEndpointTimeouts `pulumi:"timeouts"`
	// A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that
	// DeployedModel. If a DeployedModel's id is not listed in this map, then it receives no traffic. The traffic percentage
	// values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See the
	// 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and
	// [documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel)
	// for more information. > **Note:** To set the map to empty, set '"{}"', apply, and then remove the field from your
	// config.
	TrafficSplit *string `pulumi:"trafficSplit"`
	// Output only. Timestamp when this Endpoint was last updated.
	UpdateTime         *string `pulumi:"updateTime"`
	VertexAiEndpointId *string `pulumi:"vertexAiEndpointId"`
}

type VertexAiEndpointState struct {
	// Output only. Timestamp when this Endpoint was created.
	CreateTime pulumi.StringPtrInput
	// Output only. DNS of the dedicated endpoint. Will only be populated if dedicatedEndpointEnabled is true. Format:
	// 'https://{endpointId}.{region}-{projectNumber}.prediction.vertexai.goog'.
	DedicatedEndpointDns pulumi.StringPtrInput
	// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the
	// dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once
	// you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com.
	// The limitation will be removed soon.
	DedicatedEndpointEnabled pulumi.BoolPtrInput
	// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and
	// EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud
	// Console](https://console.cloud.google.com/vertex-ai/).
	DeployedModels VertexAiEndpointDeployedModelArrayInput
	// The description of the Endpoint.
	Description pulumi.StringPtrInput
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will
	// be secured by this key.
	EncryptionSpec VertexAiEndpointEncryptionSpecPtrInput
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag pulumi.StringPtrInput
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64
	// characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes.
	// International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. **Note**:
	// This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the
	// field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by
	// CreateModelDeploymentMonitoringJob. Format:
	// 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'
	ModelDeploymentMonitoringJob pulumi.StringPtrInput
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name pulumi.StringPtrInput
	// The full name of the Google Compute Engine
	// [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be
	// peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not
	// peered with any network. Only one of the fields, network or enable_private_service_connect, can be set.
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	// 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is
	// network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.
	Network pulumi.StringPtrInput
	// Configures the request-response logging for online prediction.
	PredictRequestResponseLoggingConfig VertexAiEndpointPredictRequestResponseLoggingConfigPtrInput
	// Configuration for private service connect. 'network' and 'privateServiceConnectConfig' are mutually exclusive.
	PrivateServiceConnectConfig VertexAiEndpointPrivateServiceConnectConfigPtrInput
	Project                     pulumi.StringPtrInput
	// The region for the resource
	Region pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        VertexAiEndpointTimeoutsPtrInput
	// A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that
	// DeployedModel. If a DeployedModel's id is not listed in this map, then it receives no traffic. The traffic percentage
	// values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See the
	// 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and
	// [documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel)
	// for more information. > **Note:** To set the map to empty, set '"{}"', apply, and then remove the field from your
	// config.
	TrafficSplit pulumi.StringPtrInput
	// Output only. Timestamp when this Endpoint was last updated.
	UpdateTime         pulumi.StringPtrInput
	VertexAiEndpointId pulumi.StringPtrInput
}

func (VertexAiEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiEndpointState)(nil)).Elem()
}

type vertexAiEndpointArgs struct {
	// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the
	// dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once
	// you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com.
	// The limitation will be removed soon.
	DedicatedEndpointEnabled *bool `pulumi:"dedicatedEndpointEnabled"`
	// The description of the Endpoint.
	Description *string `pulumi:"description"`
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName string `pulumi:"displayName"`
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will
	// be secured by this key.
	EncryptionSpec *VertexAiEndpointEncryptionSpec `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64
	// characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes.
	// International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. **Note**:
	// This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the
	// field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name *string `pulumi:"name"`
	// The full name of the Google Compute Engine
	// [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be
	// peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not
	// peered with any network. Only one of the fields, network or enable_private_service_connect, can be set.
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	// 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is
	// network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.
	Network *string `pulumi:"network"`
	// Configures the request-response logging for online prediction.
	PredictRequestResponseLoggingConfig *VertexAiEndpointPredictRequestResponseLoggingConfig `pulumi:"predictRequestResponseLoggingConfig"`
	// Configuration for private service connect. 'network' and 'privateServiceConnectConfig' are mutually exclusive.
	PrivateServiceConnectConfig *VertexAiEndpointPrivateServiceConnectConfig `pulumi:"privateServiceConnectConfig"`
	Project                     *string                                      `pulumi:"project"`
	// The region for the resource
	Region   *string                   `pulumi:"region"`
	Timeouts *VertexAiEndpointTimeouts `pulumi:"timeouts"`
	// A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that
	// DeployedModel. If a DeployedModel's id is not listed in this map, then it receives no traffic. The traffic percentage
	// values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See the
	// 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and
	// [documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel)
	// for more information. > **Note:** To set the map to empty, set '"{}"', apply, and then remove the field from your
	// config.
	TrafficSplit       *string `pulumi:"trafficSplit"`
	VertexAiEndpointId *string `pulumi:"vertexAiEndpointId"`
}

// The set of arguments for constructing a VertexAiEndpoint resource.
type VertexAiEndpointArgs struct {
	// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the
	// dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once
	// you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com.
	// The limitation will be removed soon.
	DedicatedEndpointEnabled pulumi.BoolPtrInput
	// The description of the Endpoint.
	Description pulumi.StringPtrInput
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName pulumi.StringInput
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will
	// be secured by this key.
	EncryptionSpec VertexAiEndpointEncryptionSpecPtrInput
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64
	// characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes.
	// International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. **Note**:
	// This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the
	// field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringInput
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name pulumi.StringPtrInput
	// The full name of the Google Compute Engine
	// [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be
	// peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not
	// peered with any network. Only one of the fields, network or enable_private_service_connect, can be set.
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	// 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is
	// network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.
	Network pulumi.StringPtrInput
	// Configures the request-response logging for online prediction.
	PredictRequestResponseLoggingConfig VertexAiEndpointPredictRequestResponseLoggingConfigPtrInput
	// Configuration for private service connect. 'network' and 'privateServiceConnectConfig' are mutually exclusive.
	PrivateServiceConnectConfig VertexAiEndpointPrivateServiceConnectConfigPtrInput
	Project                     pulumi.StringPtrInput
	// The region for the resource
	Region   pulumi.StringPtrInput
	Timeouts VertexAiEndpointTimeoutsPtrInput
	// A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that
	// DeployedModel. If a DeployedModel's id is not listed in this map, then it receives no traffic. The traffic percentage
	// values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See the
	// 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and
	// [documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel)
	// for more information. > **Note:** To set the map to empty, set '"{}"', apply, and then remove the field from your
	// config.
	TrafficSplit       pulumi.StringPtrInput
	VertexAiEndpointId pulumi.StringPtrInput
}

func (VertexAiEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiEndpointArgs)(nil)).Elem()
}

type VertexAiEndpointInput interface {
	pulumi.Input

	ToVertexAiEndpointOutput() VertexAiEndpointOutput
	ToVertexAiEndpointOutputWithContext(ctx context.Context) VertexAiEndpointOutput
}

func (*VertexAiEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiEndpoint)(nil)).Elem()
}

func (i *VertexAiEndpoint) ToVertexAiEndpointOutput() VertexAiEndpointOutput {
	return i.ToVertexAiEndpointOutputWithContext(context.Background())
}

func (i *VertexAiEndpoint) ToVertexAiEndpointOutputWithContext(ctx context.Context) VertexAiEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VertexAiEndpointOutput)
}

type VertexAiEndpointOutput struct{ *pulumi.OutputState }

func (VertexAiEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiEndpoint)(nil)).Elem()
}

func (o VertexAiEndpointOutput) ToVertexAiEndpointOutput() VertexAiEndpointOutput {
	return o
}

func (o VertexAiEndpointOutput) ToVertexAiEndpointOutputWithContext(ctx context.Context) VertexAiEndpointOutput {
	return o
}

// Output only. Timestamp when this Endpoint was created.
func (o VertexAiEndpointOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. DNS of the dedicated endpoint. Will only be populated if dedicatedEndpointEnabled is true. Format:
// 'https://{endpointId}.{region}-{projectNumber}.prediction.vertexai.goog'.
func (o VertexAiEndpointOutput) DedicatedEndpointDns() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.DedicatedEndpointDns }).(pulumi.StringOutput)
}

// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the
// dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once
// you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com.
// The limitation will be removed soon.
func (o VertexAiEndpointOutput) DedicatedEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.BoolPtrOutput { return v.DedicatedEndpointEnabled }).(pulumi.BoolPtrOutput)
}

// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and
// EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud
// Console](https://console.cloud.google.com/vertex-ai/).
func (o VertexAiEndpointOutput) DeployedModels() VertexAiEndpointDeployedModelArrayOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) VertexAiEndpointDeployedModelArrayOutput { return v.DeployedModels }).(VertexAiEndpointDeployedModelArrayOutput)
}

// The description of the Endpoint.
func (o VertexAiEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8
// characters.
func (o VertexAiEndpointOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o VertexAiEndpointOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will
// be secured by this key.
func (o VertexAiEndpointOutput) EncryptionSpec() VertexAiEndpointEncryptionSpecPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) VertexAiEndpointEncryptionSpecPtrOutput { return v.EncryptionSpec }).(VertexAiEndpointEncryptionSpecPtrOutput)
}

// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
func (o VertexAiEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64
// characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes.
// International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. **Note**:
// This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the
// field 'effective_labels' for all of the labels present on the resource.
func (o VertexAiEndpointOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
func (o VertexAiEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by
// CreateModelDeploymentMonitoringJob. Format:
// 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'
func (o VertexAiEndpointOutput) ModelDeploymentMonitoringJob() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.ModelDeploymentMonitoringJob }).(pulumi.StringOutput)
}

// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
func (o VertexAiEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The full name of the Google Compute Engine
// [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be
// peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not
// peered with any network. Only one of the fields, network or enable_private_service_connect, can be set.
// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
// 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is
// network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.
func (o VertexAiEndpointOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// Configures the request-response logging for online prediction.
func (o VertexAiEndpointOutput) PredictRequestResponseLoggingConfig() VertexAiEndpointPredictRequestResponseLoggingConfigPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) VertexAiEndpointPredictRequestResponseLoggingConfigPtrOutput {
		return v.PredictRequestResponseLoggingConfig
	}).(VertexAiEndpointPredictRequestResponseLoggingConfigPtrOutput)
}

// Configuration for private service connect. 'network' and 'privateServiceConnectConfig' are mutually exclusive.
func (o VertexAiEndpointOutput) PrivateServiceConnectConfig() VertexAiEndpointPrivateServiceConnectConfigPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) VertexAiEndpointPrivateServiceConnectConfigPtrOutput {
		return v.PrivateServiceConnectConfig
	}).(VertexAiEndpointPrivateServiceConnectConfigPtrOutput)
}

func (o VertexAiEndpointOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region for the resource
func (o VertexAiEndpointOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o VertexAiEndpointOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o VertexAiEndpointOutput) Timeouts() VertexAiEndpointTimeoutsPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) VertexAiEndpointTimeoutsPtrOutput { return v.Timeouts }).(VertexAiEndpointTimeoutsPtrOutput)
}

// A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that
// DeployedModel. If a DeployedModel's id is not listed in this map, then it receives no traffic. The traffic percentage
// values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See the
// 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and
// [documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel)
// for more information. > **Note:** To set the map to empty, set '"{}"', apply, and then remove the field from your
// config.
func (o VertexAiEndpointOutput) TrafficSplit() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.TrafficSplit }).(pulumi.StringOutput)
}

// Output only. Timestamp when this Endpoint was last updated.
func (o VertexAiEndpointOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func (o VertexAiEndpointOutput) VertexAiEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpoint) pulumi.StringOutput { return v.VertexAiEndpointId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VertexAiEndpointInput)(nil)).Elem(), &VertexAiEndpoint{})
	pulumi.RegisterOutputType(VertexAiEndpointOutput{})
}
