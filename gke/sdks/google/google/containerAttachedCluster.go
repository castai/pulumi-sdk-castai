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

type ContainerAttachedCluster struct {
	pulumi.CustomResourceState

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization ContainerAttachedClusterAuthorizationPtrOutput `pulumi:"authorization"`
	// Binary Authorization configuration.
	BinaryAuthorization ContainerAttachedClusterBinaryAuthorizationPtrOutput `pulumi:"binaryAuthorization"`
	// Output only. The region where this cluster runs. For EKS clusters, this is an AWS region. For AKS clusters, this is an
	// Azure region.
	ClusterRegion              pulumi.StringOutput `pulumi:"clusterRegion"`
	ContainerAttachedClusterId pulumi.StringOutput `pulumi:"containerAttachedClusterId"`
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Policy to determine what flags to send on delete. Possible values: DELETE, DELETE_IGNORE_ERRORS
	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks", "generic". The generic
	// distribution provides the ability to register or migrate any CNCF conformant cluster.
	Distribution         pulumi.StringOutput    `pulumi:"distribution"`
	EffectiveAnnotations pulumi.StringMapOutput `pulumi:"effectiveAnnotations"`
	// A set of errors found in the cluster.
	Errors ContainerAttachedClusterErrorArrayOutput `pulumi:"errors"`
	// Fleet configuration.
	Fleet ContainerAttachedClusterFleetOutput `pulumi:"fleet"`
	// The Kubernetes version of the cluster.
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Logging configuration.
	LoggingConfig ContainerAttachedClusterLoggingConfigPtrOutput `pulumi:"loggingConfig"`
	// Monitoring configuration.
	MonitoringConfig ContainerAttachedClusterMonitoringConfigPtrOutput `pulumi:"monitoringConfig"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// OIDC discovery information of the target cluster. Kubernetes Service Account (KSA) tokens are JWT tokens signed by the
	// cluster API server. This fields indicates how GCP services validate KSA tokens in order to allow system workloads (such
	// as GKE Connect and telemetry agents) to authenticate back to GCP. Both clusters with public and private issuer URLs are
	// supported. Clusters with public issuers only need to specify the 'issuer_url' field while clusters with private issuers
	// need to provide both 'issuer_url' and 'jwks'.
	OidcConfig ContainerAttachedClusterOidcConfigOutput `pulumi:"oidcConfig"`
	// The platform version for the cluster (e.g. '1.23.0-gke.1').
	PlatformVersion pulumi.StringOutput `pulumi:"platformVersion"`
	Project         pulumi.StringOutput `pulumi:"project"`
	// Support for proxy configuration.
	ProxyConfig ContainerAttachedClusterProxyConfigPtrOutput `pulumi:"proxyConfig"`
	// If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Enable/Disable Security Posture API features for the cluster.
	//
	// Deprecated: Deprecated
	SecurityPostureConfig ContainerAttachedClusterSecurityPostureConfigPtrOutput `pulumi:"securityPostureConfig"`
	// The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING,
	// ERROR, DEGRADED
	State    pulumi.StringOutput                       `pulumi:"state"`
	Timeouts ContainerAttachedClusterTimeoutsPtrOutput `pulumi:"timeouts"`
	// A globally unique identifier for the cluster.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time at which this cluster was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Workload Identity settings.
	WorkloadIdentityConfigs ContainerAttachedClusterWorkloadIdentityConfigArrayOutput `pulumi:"workloadIdentityConfigs"`
}

// NewContainerAttachedCluster registers a new resource with the given unique name, arguments, and options.
func NewContainerAttachedCluster(ctx *pulumi.Context,
	name string, args *ContainerAttachedClusterArgs, opts ...pulumi.ResourceOption) (*ContainerAttachedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Distribution == nil {
		return nil, errors.New("invalid value for required argument 'Distribution'")
	}
	if args.Fleet == nil {
		return nil, errors.New("invalid value for required argument 'Fleet'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.OidcConfig == nil {
		return nil, errors.New("invalid value for required argument 'OidcConfig'")
	}
	if args.PlatformVersion == nil {
		return nil, errors.New("invalid value for required argument 'PlatformVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ContainerAttachedCluster
	err = ctx.RegisterPackageResource("google:index/containerAttachedCluster:ContainerAttachedCluster", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerAttachedCluster gets an existing ContainerAttachedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerAttachedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAttachedClusterState, opts ...pulumi.ResourceOption) (*ContainerAttachedCluster, error) {
	var resource ContainerAttachedCluster
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/containerAttachedCluster:ContainerAttachedCluster", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerAttachedCluster resources.
type containerAttachedClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization *ContainerAttachedClusterAuthorization `pulumi:"authorization"`
	// Binary Authorization configuration.
	BinaryAuthorization *ContainerAttachedClusterBinaryAuthorization `pulumi:"binaryAuthorization"`
	// Output only. The region where this cluster runs. For EKS clusters, this is an AWS region. For AKS clusters, this is an
	// Azure region.
	ClusterRegion              *string `pulumi:"clusterRegion"`
	ContainerAttachedClusterId *string `pulumi:"containerAttachedClusterId"`
	// Output only. The time at which this cluster was created.
	CreateTime *string `pulumi:"createTime"`
	// Policy to determine what flags to send on delete. Possible values: DELETE, DELETE_IGNORE_ERRORS
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks", "generic". The generic
	// distribution provides the ability to register or migrate any CNCF conformant cluster.
	Distribution         *string           `pulumi:"distribution"`
	EffectiveAnnotations map[string]string `pulumi:"effectiveAnnotations"`
	// A set of errors found in the cluster.
	Errors []ContainerAttachedClusterError `pulumi:"errors"`
	// Fleet configuration.
	Fleet *ContainerAttachedClusterFleet `pulumi:"fleet"`
	// The Kubernetes version of the cluster.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Logging configuration.
	LoggingConfig *ContainerAttachedClusterLoggingConfig `pulumi:"loggingConfig"`
	// Monitoring configuration.
	MonitoringConfig *ContainerAttachedClusterMonitoringConfig `pulumi:"monitoringConfig"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// OIDC discovery information of the target cluster. Kubernetes Service Account (KSA) tokens are JWT tokens signed by the
	// cluster API server. This fields indicates how GCP services validate KSA tokens in order to allow system workloads (such
	// as GKE Connect and telemetry agents) to authenticate back to GCP. Both clusters with public and private issuer URLs are
	// supported. Clusters with public issuers only need to specify the 'issuer_url' field while clusters with private issuers
	// need to provide both 'issuer_url' and 'jwks'.
	OidcConfig *ContainerAttachedClusterOidcConfig `pulumi:"oidcConfig"`
	// The platform version for the cluster (e.g. '1.23.0-gke.1').
	PlatformVersion *string `pulumi:"platformVersion"`
	Project         *string `pulumi:"project"`
	// Support for proxy configuration.
	ProxyConfig *ContainerAttachedClusterProxyConfig `pulumi:"proxyConfig"`
	// If set, there are currently changes in flight to the cluster.
	Reconciling *bool `pulumi:"reconciling"`
	// Enable/Disable Security Posture API features for the cluster.
	//
	// Deprecated: Deprecated
	SecurityPostureConfig *ContainerAttachedClusterSecurityPostureConfig `pulumi:"securityPostureConfig"`
	// The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING,
	// ERROR, DEGRADED
	State    *string                           `pulumi:"state"`
	Timeouts *ContainerAttachedClusterTimeouts `pulumi:"timeouts"`
	// A globally unique identifier for the cluster.
	Uid *string `pulumi:"uid"`
	// The time at which this cluster was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Workload Identity settings.
	WorkloadIdentityConfigs []ContainerAttachedClusterWorkloadIdentityConfig `pulumi:"workloadIdentityConfigs"`
}

type ContainerAttachedClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	Authorization ContainerAttachedClusterAuthorizationPtrInput
	// Binary Authorization configuration.
	BinaryAuthorization ContainerAttachedClusterBinaryAuthorizationPtrInput
	// Output only. The region where this cluster runs. For EKS clusters, this is an AWS region. For AKS clusters, this is an
	// Azure region.
	ClusterRegion              pulumi.StringPtrInput
	ContainerAttachedClusterId pulumi.StringPtrInput
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringPtrInput
	// Policy to determine what flags to send on delete. Possible values: DELETE, DELETE_IGNORE_ERRORS
	DeletionPolicy pulumi.StringPtrInput
	// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks", "generic". The generic
	// distribution provides the ability to register or migrate any CNCF conformant cluster.
	Distribution         pulumi.StringPtrInput
	EffectiveAnnotations pulumi.StringMapInput
	// A set of errors found in the cluster.
	Errors ContainerAttachedClusterErrorArrayInput
	// Fleet configuration.
	Fleet ContainerAttachedClusterFleetPtrInput
	// The Kubernetes version of the cluster.
	KubernetesVersion pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Logging configuration.
	LoggingConfig ContainerAttachedClusterLoggingConfigPtrInput
	// Monitoring configuration.
	MonitoringConfig ContainerAttachedClusterMonitoringConfigPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// OIDC discovery information of the target cluster. Kubernetes Service Account (KSA) tokens are JWT tokens signed by the
	// cluster API server. This fields indicates how GCP services validate KSA tokens in order to allow system workloads (such
	// as GKE Connect and telemetry agents) to authenticate back to GCP. Both clusters with public and private issuer URLs are
	// supported. Clusters with public issuers only need to specify the 'issuer_url' field while clusters with private issuers
	// need to provide both 'issuer_url' and 'jwks'.
	OidcConfig ContainerAttachedClusterOidcConfigPtrInput
	// The platform version for the cluster (e.g. '1.23.0-gke.1').
	PlatformVersion pulumi.StringPtrInput
	Project         pulumi.StringPtrInput
	// Support for proxy configuration.
	ProxyConfig ContainerAttachedClusterProxyConfigPtrInput
	// If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolPtrInput
	// Enable/Disable Security Posture API features for the cluster.
	//
	// Deprecated: Deprecated
	SecurityPostureConfig ContainerAttachedClusterSecurityPostureConfigPtrInput
	// The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING,
	// ERROR, DEGRADED
	State    pulumi.StringPtrInput
	Timeouts ContainerAttachedClusterTimeoutsPtrInput
	// A globally unique identifier for the cluster.
	Uid pulumi.StringPtrInput
	// The time at which this cluster was last updated.
	UpdateTime pulumi.StringPtrInput
	// Workload Identity settings.
	WorkloadIdentityConfigs ContainerAttachedClusterWorkloadIdentityConfigArrayInput
}

func (ContainerAttachedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAttachedClusterState)(nil)).Elem()
}

type containerAttachedClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization *ContainerAttachedClusterAuthorization `pulumi:"authorization"`
	// Binary Authorization configuration.
	BinaryAuthorization        *ContainerAttachedClusterBinaryAuthorization `pulumi:"binaryAuthorization"`
	ContainerAttachedClusterId *string                                      `pulumi:"containerAttachedClusterId"`
	// Policy to determine what flags to send on delete. Possible values: DELETE, DELETE_IGNORE_ERRORS
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks", "generic". The generic
	// distribution provides the ability to register or migrate any CNCF conformant cluster.
	Distribution string `pulumi:"distribution"`
	// Fleet configuration.
	Fleet ContainerAttachedClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Logging configuration.
	LoggingConfig *ContainerAttachedClusterLoggingConfig `pulumi:"loggingConfig"`
	// Monitoring configuration.
	MonitoringConfig *ContainerAttachedClusterMonitoringConfig `pulumi:"monitoringConfig"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// OIDC discovery information of the target cluster. Kubernetes Service Account (KSA) tokens are JWT tokens signed by the
	// cluster API server. This fields indicates how GCP services validate KSA tokens in order to allow system workloads (such
	// as GKE Connect and telemetry agents) to authenticate back to GCP. Both clusters with public and private issuer URLs are
	// supported. Clusters with public issuers only need to specify the 'issuer_url' field while clusters with private issuers
	// need to provide both 'issuer_url' and 'jwks'.
	OidcConfig ContainerAttachedClusterOidcConfig `pulumi:"oidcConfig"`
	// The platform version for the cluster (e.g. '1.23.0-gke.1').
	PlatformVersion string  `pulumi:"platformVersion"`
	Project         *string `pulumi:"project"`
	// Support for proxy configuration.
	ProxyConfig *ContainerAttachedClusterProxyConfig `pulumi:"proxyConfig"`
	// Enable/Disable Security Posture API features for the cluster.
	//
	// Deprecated: Deprecated
	SecurityPostureConfig *ContainerAttachedClusterSecurityPostureConfig `pulumi:"securityPostureConfig"`
	Timeouts              *ContainerAttachedClusterTimeouts              `pulumi:"timeouts"`
}

// The set of arguments for constructing a ContainerAttachedCluster resource.
type ContainerAttachedClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	Authorization ContainerAttachedClusterAuthorizationPtrInput
	// Binary Authorization configuration.
	BinaryAuthorization        ContainerAttachedClusterBinaryAuthorizationPtrInput
	ContainerAttachedClusterId pulumi.StringPtrInput
	// Policy to determine what flags to send on delete. Possible values: DELETE, DELETE_IGNORE_ERRORS
	DeletionPolicy pulumi.StringPtrInput
	// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks", "generic". The generic
	// distribution provides the ability to register or migrate any CNCF conformant cluster.
	Distribution pulumi.StringInput
	// Fleet configuration.
	Fleet ContainerAttachedClusterFleetInput
	// The location for the resource
	Location pulumi.StringInput
	// Logging configuration.
	LoggingConfig ContainerAttachedClusterLoggingConfigPtrInput
	// Monitoring configuration.
	MonitoringConfig ContainerAttachedClusterMonitoringConfigPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// OIDC discovery information of the target cluster. Kubernetes Service Account (KSA) tokens are JWT tokens signed by the
	// cluster API server. This fields indicates how GCP services validate KSA tokens in order to allow system workloads (such
	// as GKE Connect and telemetry agents) to authenticate back to GCP. Both clusters with public and private issuer URLs are
	// supported. Clusters with public issuers only need to specify the 'issuer_url' field while clusters with private issuers
	// need to provide both 'issuer_url' and 'jwks'.
	OidcConfig ContainerAttachedClusterOidcConfigInput
	// The platform version for the cluster (e.g. '1.23.0-gke.1').
	PlatformVersion pulumi.StringInput
	Project         pulumi.StringPtrInput
	// Support for proxy configuration.
	ProxyConfig ContainerAttachedClusterProxyConfigPtrInput
	// Enable/Disable Security Posture API features for the cluster.
	//
	// Deprecated: Deprecated
	SecurityPostureConfig ContainerAttachedClusterSecurityPostureConfigPtrInput
	Timeouts              ContainerAttachedClusterTimeoutsPtrInput
}

func (ContainerAttachedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAttachedClusterArgs)(nil)).Elem()
}

type ContainerAttachedClusterInput interface {
	pulumi.Input

	ToContainerAttachedClusterOutput() ContainerAttachedClusterOutput
	ToContainerAttachedClusterOutputWithContext(ctx context.Context) ContainerAttachedClusterOutput
}

func (*ContainerAttachedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAttachedCluster)(nil)).Elem()
}

func (i *ContainerAttachedCluster) ToContainerAttachedClusterOutput() ContainerAttachedClusterOutput {
	return i.ToContainerAttachedClusterOutputWithContext(context.Background())
}

func (i *ContainerAttachedCluster) ToContainerAttachedClusterOutputWithContext(ctx context.Context) ContainerAttachedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAttachedClusterOutput)
}

type ContainerAttachedClusterOutput struct{ *pulumi.OutputState }

func (ContainerAttachedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAttachedCluster)(nil)).Elem()
}

func (o ContainerAttachedClusterOutput) ToContainerAttachedClusterOutput() ContainerAttachedClusterOutput {
	return o
}

func (o ContainerAttachedClusterOutput) ToContainerAttachedClusterOutputWithContext(ctx context.Context) ContainerAttachedClusterOutput {
	return o
}

// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
// all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required),
// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
// 'effective_annotations' for all of the annotations present on the resource.
func (o ContainerAttachedClusterOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Configuration related to the cluster RBAC settings.
func (o ContainerAttachedClusterOutput) Authorization() ContainerAttachedClusterAuthorizationPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterAuthorizationPtrOutput {
		return v.Authorization
	}).(ContainerAttachedClusterAuthorizationPtrOutput)
}

// Binary Authorization configuration.
func (o ContainerAttachedClusterOutput) BinaryAuthorization() ContainerAttachedClusterBinaryAuthorizationPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterBinaryAuthorizationPtrOutput {
		return v.BinaryAuthorization
	}).(ContainerAttachedClusterBinaryAuthorizationPtrOutput)
}

// Output only. The region where this cluster runs. For EKS clusters, this is an AWS region. For AKS clusters, this is an
// Azure region.
func (o ContainerAttachedClusterOutput) ClusterRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.ClusterRegion }).(pulumi.StringOutput)
}

func (o ContainerAttachedClusterOutput) ContainerAttachedClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.ContainerAttachedClusterId }).(pulumi.StringOutput)
}

// Output only. The time at which this cluster was created.
func (o ContainerAttachedClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Policy to determine what flags to send on delete. Possible values: DELETE, DELETE_IGNORE_ERRORS
func (o ContainerAttachedClusterOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
func (o ContainerAttachedClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks", "generic". The generic
// distribution provides the ability to register or migrate any CNCF conformant cluster.
func (o ContainerAttachedClusterOutput) Distribution() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.Distribution }).(pulumi.StringOutput)
}

func (o ContainerAttachedClusterOutput) EffectiveAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringMapOutput { return v.EffectiveAnnotations }).(pulumi.StringMapOutput)
}

// A set of errors found in the cluster.
func (o ContainerAttachedClusterOutput) Errors() ContainerAttachedClusterErrorArrayOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterErrorArrayOutput { return v.Errors }).(ContainerAttachedClusterErrorArrayOutput)
}

// Fleet configuration.
func (o ContainerAttachedClusterOutput) Fleet() ContainerAttachedClusterFleetOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterFleetOutput { return v.Fleet }).(ContainerAttachedClusterFleetOutput)
}

// The Kubernetes version of the cluster.
func (o ContainerAttachedClusterOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The location for the resource
func (o ContainerAttachedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Logging configuration.
func (o ContainerAttachedClusterOutput) LoggingConfig() ContainerAttachedClusterLoggingConfigPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterLoggingConfigPtrOutput {
		return v.LoggingConfig
	}).(ContainerAttachedClusterLoggingConfigPtrOutput)
}

// Monitoring configuration.
func (o ContainerAttachedClusterOutput) MonitoringConfig() ContainerAttachedClusterMonitoringConfigPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterMonitoringConfigPtrOutput {
		return v.MonitoringConfig
	}).(ContainerAttachedClusterMonitoringConfigPtrOutput)
}

// The name of this resource.
func (o ContainerAttachedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OIDC discovery information of the target cluster. Kubernetes Service Account (KSA) tokens are JWT tokens signed by the
// cluster API server. This fields indicates how GCP services validate KSA tokens in order to allow system workloads (such
// as GKE Connect and telemetry agents) to authenticate back to GCP. Both clusters with public and private issuer URLs are
// supported. Clusters with public issuers only need to specify the 'issuer_url' field while clusters with private issuers
// need to provide both 'issuer_url' and 'jwks'.
func (o ContainerAttachedClusterOutput) OidcConfig() ContainerAttachedClusterOidcConfigOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterOidcConfigOutput { return v.OidcConfig }).(ContainerAttachedClusterOidcConfigOutput)
}

// The platform version for the cluster (e.g. '1.23.0-gke.1').
func (o ContainerAttachedClusterOutput) PlatformVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.PlatformVersion }).(pulumi.StringOutput)
}

func (o ContainerAttachedClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Support for proxy configuration.
func (o ContainerAttachedClusterOutput) ProxyConfig() ContainerAttachedClusterProxyConfigPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterProxyConfigPtrOutput { return v.ProxyConfig }).(ContainerAttachedClusterProxyConfigPtrOutput)
}

// If set, there are currently changes in flight to the cluster.
func (o ContainerAttachedClusterOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Enable/Disable Security Posture API features for the cluster.
//
// Deprecated: Deprecated
func (o ContainerAttachedClusterOutput) SecurityPostureConfig() ContainerAttachedClusterSecurityPostureConfigPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterSecurityPostureConfigPtrOutput {
		return v.SecurityPostureConfig
	}).(ContainerAttachedClusterSecurityPostureConfigPtrOutput)
}

// The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING,
// ERROR, DEGRADED
func (o ContainerAttachedClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ContainerAttachedClusterOutput) Timeouts() ContainerAttachedClusterTimeoutsPtrOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterTimeoutsPtrOutput { return v.Timeouts }).(ContainerAttachedClusterTimeoutsPtrOutput)
}

// A globally unique identifier for the cluster.
func (o ContainerAttachedClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time at which this cluster was last updated.
func (o ContainerAttachedClusterOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Workload Identity settings.
func (o ContainerAttachedClusterOutput) WorkloadIdentityConfigs() ContainerAttachedClusterWorkloadIdentityConfigArrayOutput {
	return o.ApplyT(func(v *ContainerAttachedCluster) ContainerAttachedClusterWorkloadIdentityConfigArrayOutput {
		return v.WorkloadIdentityConfigs
	}).(ContainerAttachedClusterWorkloadIdentityConfigArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerAttachedClusterInput)(nil)).Elem(), &ContainerAttachedCluster{})
	pulumi.RegisterOutputType(ContainerAttachedClusterOutput{})
}
