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

type SccNotificationConfig struct {
	pulumi.CustomResourceState

	// This must be unique within the organization.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The resource name of this notification config, in the format
	// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic             pulumi.StringOutput `pulumi:"pubsubTopic"`
	SccNotificationConfigId pulumi.StringOutput `pulumi:"sccNotificationConfigId"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig SccNotificationConfigStreamingConfigOutput `pulumi:"streamingConfig"`
	Timeouts        SccNotificationConfigTimeoutsPtrOutput     `pulumi:"timeouts"`
}

// NewSccNotificationConfig registers a new resource with the given unique name, arguments, and options.
func NewSccNotificationConfig(ctx *pulumi.Context,
	name string, args *SccNotificationConfigArgs, opts ...pulumi.ResourceOption) (*SccNotificationConfig, error) {
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
	var resource SccNotificationConfig
	err = ctx.RegisterPackageResource("google:index/sccNotificationConfig:SccNotificationConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccNotificationConfig gets an existing SccNotificationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccNotificationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccNotificationConfigState, opts ...pulumi.ResourceOption) (*SccNotificationConfig, error) {
	var resource SccNotificationConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/sccNotificationConfig:SccNotificationConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccNotificationConfig resources.
type sccNotificationConfigState struct {
	// This must be unique within the organization.
	ConfigId *string `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// The resource name of this notification config, in the format
	// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
	Name *string `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization *string `pulumi:"organization"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic             *string `pulumi:"pubsubTopic"`
	SccNotificationConfigId *string `pulumi:"sccNotificationConfigId"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig *SccNotificationConfigStreamingConfig `pulumi:"streamingConfig"`
	Timeouts        *SccNotificationConfigTimeouts        `pulumi:"timeouts"`
}

type SccNotificationConfigState struct {
	// This must be unique within the organization.
	ConfigId pulumi.StringPtrInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// The resource name of this notification config, in the format
	// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
	Name pulumi.StringPtrInput
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization pulumi.StringPtrInput
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic             pulumi.StringPtrInput
	SccNotificationConfigId pulumi.StringPtrInput
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount pulumi.StringPtrInput
	// The config for triggering streaming-based notifications.
	StreamingConfig SccNotificationConfigStreamingConfigPtrInput
	Timeouts        SccNotificationConfigTimeoutsPtrInput
}

func (SccNotificationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccNotificationConfigState)(nil)).Elem()
}

type sccNotificationConfigArgs struct {
	// This must be unique within the organization.
	ConfigId string `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization string `pulumi:"organization"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic             string  `pulumi:"pubsubTopic"`
	SccNotificationConfigId *string `pulumi:"sccNotificationConfigId"`
	// The config for triggering streaming-based notifications.
	StreamingConfig SccNotificationConfigStreamingConfig `pulumi:"streamingConfig"`
	Timeouts        *SccNotificationConfigTimeouts       `pulumi:"timeouts"`
}

// The set of arguments for constructing a SccNotificationConfig resource.
type SccNotificationConfigArgs struct {
	// This must be unique within the organization.
	ConfigId pulumi.StringInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// The organization whose Cloud Security Command Center the Notification Config lives in.
	Organization pulumi.StringInput
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic             pulumi.StringInput
	SccNotificationConfigId pulumi.StringPtrInput
	// The config for triggering streaming-based notifications.
	StreamingConfig SccNotificationConfigStreamingConfigInput
	Timeouts        SccNotificationConfigTimeoutsPtrInput
}

func (SccNotificationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccNotificationConfigArgs)(nil)).Elem()
}

type SccNotificationConfigInput interface {
	pulumi.Input

	ToSccNotificationConfigOutput() SccNotificationConfigOutput
	ToSccNotificationConfigOutputWithContext(ctx context.Context) SccNotificationConfigOutput
}

func (*SccNotificationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SccNotificationConfig)(nil)).Elem()
}

func (i *SccNotificationConfig) ToSccNotificationConfigOutput() SccNotificationConfigOutput {
	return i.ToSccNotificationConfigOutputWithContext(context.Background())
}

func (i *SccNotificationConfig) ToSccNotificationConfigOutputWithContext(ctx context.Context) SccNotificationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccNotificationConfigOutput)
}

type SccNotificationConfigOutput struct{ *pulumi.OutputState }

func (SccNotificationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccNotificationConfig)(nil)).Elem()
}

func (o SccNotificationConfigOutput) ToSccNotificationConfigOutput() SccNotificationConfigOutput {
	return o
}

func (o SccNotificationConfigOutput) ToSccNotificationConfigOutputWithContext(ctx context.Context) SccNotificationConfigOutput {
	return o
}

// This must be unique within the organization.
func (o SccNotificationConfigOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccNotificationConfig) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// The description of the notification config (max of 1024 characters).
func (o SccNotificationConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccNotificationConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The resource name of this notification config, in the format
// 'organizations/{{organization}}/notificationConfigs/{{config_id}}'.
func (o SccNotificationConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccNotificationConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization whose Cloud Security Command Center the Notification Config lives in.
func (o SccNotificationConfigOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccNotificationConfig) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
func (o SccNotificationConfigOutput) PubsubTopic() pulumi.StringOutput {
	return o.ApplyT(func(v *SccNotificationConfig) pulumi.StringOutput { return v.PubsubTopic }).(pulumi.StringOutput)
}

func (o SccNotificationConfigOutput) SccNotificationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccNotificationConfig) pulumi.StringOutput { return v.SccNotificationConfigId }).(pulumi.StringOutput)
}

// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
func (o SccNotificationConfigOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *SccNotificationConfig) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The config for triggering streaming-based notifications.
func (o SccNotificationConfigOutput) StreamingConfig() SccNotificationConfigStreamingConfigOutput {
	return o.ApplyT(func(v *SccNotificationConfig) SccNotificationConfigStreamingConfigOutput { return v.StreamingConfig }).(SccNotificationConfigStreamingConfigOutput)
}

func (o SccNotificationConfigOutput) Timeouts() SccNotificationConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccNotificationConfig) SccNotificationConfigTimeoutsPtrOutput { return v.Timeouts }).(SccNotificationConfigTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccNotificationConfigInput)(nil)).Elem(), &SccNotificationConfig{})
	pulumi.RegisterOutputType(SccNotificationConfigOutput{})
}
