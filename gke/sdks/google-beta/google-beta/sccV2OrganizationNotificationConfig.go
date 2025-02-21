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

type SccV2OrganizationNotificationConfig struct {
	pulumi.CustomResourceState

	// This must be unique within the organization.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// location Id is provided by organization. If not provided, Use global as default.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name of this notification config, in the format
	// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                           pulumi.StringOutput `pulumi:"pubsubTopic"`
	SccV2OrganizationNotificationConfigId pulumi.StringOutput `pulumi:"sccV2OrganizationNotificationConfigId"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2OrganizationNotificationConfigStreamingConfigOutput `pulumi:"streamingConfig"`
	Timeouts        SccV2OrganizationNotificationConfigTimeoutsPtrOutput     `pulumi:"timeouts"`
}

// NewSccV2OrganizationNotificationConfig registers a new resource with the given unique name, arguments, and options.
func NewSccV2OrganizationNotificationConfig(ctx *pulumi.Context,
	name string, args *SccV2OrganizationNotificationConfigArgs, opts ...pulumi.ResourceOption) (*SccV2OrganizationNotificationConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.PubsubTopic == nil {
		return nil, errors.New("invalid value for required argument 'PubsubTopic'")
	}
	if args.StreamingConfig == nil {
		return nil, errors.New("invalid value for required argument 'StreamingConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccV2OrganizationNotificationConfig
	err = ctx.RegisterPackageResource("google-beta:index/sccV2OrganizationNotificationConfig:SccV2OrganizationNotificationConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccV2OrganizationNotificationConfig gets an existing SccV2OrganizationNotificationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccV2OrganizationNotificationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccV2OrganizationNotificationConfigState, opts ...pulumi.ResourceOption) (*SccV2OrganizationNotificationConfig, error) {
	var resource SccV2OrganizationNotificationConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/sccV2OrganizationNotificationConfig:SccV2OrganizationNotificationConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccV2OrganizationNotificationConfig resources.
type sccV2OrganizationNotificationConfigState struct {
	// This must be unique within the organization.
	ConfigId *string `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// location Id is provided by organization. If not provided, Use global as default.
	Location *string `pulumi:"location"`
	// The resource name of this notification config, in the format
	// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
	Name *string `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization *string `pulumi:"organization"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                           *string `pulumi:"pubsubTopic"`
	SccV2OrganizationNotificationConfigId *string `pulumi:"sccV2OrganizationNotificationConfigId"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig *SccV2OrganizationNotificationConfigStreamingConfig `pulumi:"streamingConfig"`
	Timeouts        *SccV2OrganizationNotificationConfigTimeouts        `pulumi:"timeouts"`
}

type SccV2OrganizationNotificationConfigState struct {
	// This must be unique within the organization.
	ConfigId pulumi.StringPtrInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// location Id is provided by organization. If not provided, Use global as default.
	Location pulumi.StringPtrInput
	// The resource name of this notification config, in the format
	// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
	Name pulumi.StringPtrInput
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization pulumi.StringPtrInput
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                           pulumi.StringPtrInput
	SccV2OrganizationNotificationConfigId pulumi.StringPtrInput
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount pulumi.StringPtrInput
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2OrganizationNotificationConfigStreamingConfigPtrInput
	Timeouts        SccV2OrganizationNotificationConfigTimeoutsPtrInput
}

func (SccV2OrganizationNotificationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2OrganizationNotificationConfigState)(nil)).Elem()
}

type sccV2OrganizationNotificationConfigArgs struct {
	// This must be unique within the organization.
	ConfigId string `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// location Id is provided by organization. If not provided, Use global as default.
	Location *string `pulumi:"location"`
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization string `pulumi:"organization"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                           string  `pulumi:"pubsubTopic"`
	SccV2OrganizationNotificationConfigId *string `pulumi:"sccV2OrganizationNotificationConfigId"`
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2OrganizationNotificationConfigStreamingConfig `pulumi:"streamingConfig"`
	Timeouts        *SccV2OrganizationNotificationConfigTimeouts       `pulumi:"timeouts"`
}

// The set of arguments for constructing a SccV2OrganizationNotificationConfig resource.
type SccV2OrganizationNotificationConfigArgs struct {
	// This must be unique within the organization.
	ConfigId pulumi.StringInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// location Id is provided by organization. If not provided, Use global as default.
	Location pulumi.StringPtrInput
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization pulumi.StringInput
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                           pulumi.StringInput
	SccV2OrganizationNotificationConfigId pulumi.StringPtrInput
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2OrganizationNotificationConfigStreamingConfigInput
	Timeouts        SccV2OrganizationNotificationConfigTimeoutsPtrInput
}

func (SccV2OrganizationNotificationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2OrganizationNotificationConfigArgs)(nil)).Elem()
}

type SccV2OrganizationNotificationConfigInput interface {
	pulumi.Input

	ToSccV2OrganizationNotificationConfigOutput() SccV2OrganizationNotificationConfigOutput
	ToSccV2OrganizationNotificationConfigOutputWithContext(ctx context.Context) SccV2OrganizationNotificationConfigOutput
}

func (*SccV2OrganizationNotificationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2OrganizationNotificationConfig)(nil)).Elem()
}

func (i *SccV2OrganizationNotificationConfig) ToSccV2OrganizationNotificationConfigOutput() SccV2OrganizationNotificationConfigOutput {
	return i.ToSccV2OrganizationNotificationConfigOutputWithContext(context.Background())
}

func (i *SccV2OrganizationNotificationConfig) ToSccV2OrganizationNotificationConfigOutputWithContext(ctx context.Context) SccV2OrganizationNotificationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccV2OrganizationNotificationConfigOutput)
}

type SccV2OrganizationNotificationConfigOutput struct{ *pulumi.OutputState }

func (SccV2OrganizationNotificationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2OrganizationNotificationConfig)(nil)).Elem()
}

func (o SccV2OrganizationNotificationConfigOutput) ToSccV2OrganizationNotificationConfigOutput() SccV2OrganizationNotificationConfigOutput {
	return o
}

func (o SccV2OrganizationNotificationConfigOutput) ToSccV2OrganizationNotificationConfigOutputWithContext(ctx context.Context) SccV2OrganizationNotificationConfigOutput {
	return o
}

// This must be unique within the organization.
func (o SccV2OrganizationNotificationConfigOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// The description of the notification config (max of 1024 characters).
func (o SccV2OrganizationNotificationConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// location Id is provided by organization. If not provided, Use global as default.
func (o SccV2OrganizationNotificationConfigOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name of this notification config, in the format
// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
func (o SccV2OrganizationNotificationConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization whose Cloud Security Command Center the Notification Config lives in.
func (o SccV2OrganizationNotificationConfigOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
func (o SccV2OrganizationNotificationConfigOutput) PubsubTopic() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringOutput { return v.PubsubTopic }).(pulumi.StringOutput)
}

func (o SccV2OrganizationNotificationConfigOutput) SccV2OrganizationNotificationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringOutput {
		return v.SccV2OrganizationNotificationConfigId
	}).(pulumi.StringOutput)
}

// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
func (o SccV2OrganizationNotificationConfigOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The config for triggering streaming-based notifications.
func (o SccV2OrganizationNotificationConfigOutput) StreamingConfig() SccV2OrganizationNotificationConfigStreamingConfigOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) SccV2OrganizationNotificationConfigStreamingConfigOutput {
		return v.StreamingConfig
	}).(SccV2OrganizationNotificationConfigStreamingConfigOutput)
}

func (o SccV2OrganizationNotificationConfigOutput) Timeouts() SccV2OrganizationNotificationConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccV2OrganizationNotificationConfig) SccV2OrganizationNotificationConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(SccV2OrganizationNotificationConfigTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccV2OrganizationNotificationConfigInput)(nil)).Elem(), &SccV2OrganizationNotificationConfig{})
	pulumi.RegisterOutputType(SccV2OrganizationNotificationConfigOutput{})
}
