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

type SccV2ProjectNotificationConfig struct {
	pulumi.CustomResourceState

	// This must be unique within the project.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Location ID of the parent organization. Only global is supported at the moment.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name of this notification config, in the format
	// 'projects/{{projectId}}/locations/{{location}}/notificationConfigs/{{config_id}}'.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                      pulumi.StringPtrOutput `pulumi:"pubsubTopic"`
	SccV2ProjectNotificationConfigId pulumi.StringOutput    `pulumi:"sccV2ProjectNotificationConfigId"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2ProjectNotificationConfigStreamingConfigOutput `pulumi:"streamingConfig"`
	Timeouts        SccV2ProjectNotificationConfigTimeoutsPtrOutput     `pulumi:"timeouts"`
}

// NewSccV2ProjectNotificationConfig registers a new resource with the given unique name, arguments, and options.
func NewSccV2ProjectNotificationConfig(ctx *pulumi.Context,
	name string, args *SccV2ProjectNotificationConfigArgs, opts ...pulumi.ResourceOption) (*SccV2ProjectNotificationConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.StreamingConfig == nil {
		return nil, errors.New("invalid value for required argument 'StreamingConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccV2ProjectNotificationConfig
	err = ctx.RegisterPackageResource("google-beta:index/sccV2ProjectNotificationConfig:SccV2ProjectNotificationConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccV2ProjectNotificationConfig gets an existing SccV2ProjectNotificationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccV2ProjectNotificationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccV2ProjectNotificationConfigState, opts ...pulumi.ResourceOption) (*SccV2ProjectNotificationConfig, error) {
	var resource SccV2ProjectNotificationConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/sccV2ProjectNotificationConfig:SccV2ProjectNotificationConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccV2ProjectNotificationConfig resources.
type sccV2ProjectNotificationConfigState struct {
	// This must be unique within the project.
	ConfigId *string `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// Location ID of the parent organization. Only global is supported at the moment.
	Location *string `pulumi:"location"`
	// The resource name of this notification config, in the format
	// 'projects/{{projectId}}/locations/{{location}}/notificationConfigs/{{config_id}}'.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                      *string `pulumi:"pubsubTopic"`
	SccV2ProjectNotificationConfigId *string `pulumi:"sccV2ProjectNotificationConfigId"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig *SccV2ProjectNotificationConfigStreamingConfig `pulumi:"streamingConfig"`
	Timeouts        *SccV2ProjectNotificationConfigTimeouts        `pulumi:"timeouts"`
}

type SccV2ProjectNotificationConfigState struct {
	// This must be unique within the project.
	ConfigId pulumi.StringPtrInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// Location ID of the parent organization. Only global is supported at the moment.
	Location pulumi.StringPtrInput
	// The resource name of this notification config, in the format
	// 'projects/{{projectId}}/locations/{{location}}/notificationConfigs/{{config_id}}'.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                      pulumi.StringPtrInput
	SccV2ProjectNotificationConfigId pulumi.StringPtrInput
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount pulumi.StringPtrInput
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2ProjectNotificationConfigStreamingConfigPtrInput
	Timeouts        SccV2ProjectNotificationConfigTimeoutsPtrInput
}

func (SccV2ProjectNotificationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2ProjectNotificationConfigState)(nil)).Elem()
}

type sccV2ProjectNotificationConfigArgs struct {
	// This must be unique within the project.
	ConfigId string `pulumi:"configId"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// Location ID of the parent organization. Only global is supported at the moment.
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                      *string `pulumi:"pubsubTopic"`
	SccV2ProjectNotificationConfigId *string `pulumi:"sccV2ProjectNotificationConfigId"`
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2ProjectNotificationConfigStreamingConfig `pulumi:"streamingConfig"`
	Timeouts        *SccV2ProjectNotificationConfigTimeouts       `pulumi:"timeouts"`
}

// The set of arguments for constructing a SccV2ProjectNotificationConfig resource.
type SccV2ProjectNotificationConfigArgs struct {
	// This must be unique within the project.
	ConfigId pulumi.StringInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// Location ID of the parent organization. Only global is supported at the moment.
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic                      pulumi.StringPtrInput
	SccV2ProjectNotificationConfigId pulumi.StringPtrInput
	// The config for triggering streaming-based notifications.
	StreamingConfig SccV2ProjectNotificationConfigStreamingConfigInput
	Timeouts        SccV2ProjectNotificationConfigTimeoutsPtrInput
}

func (SccV2ProjectNotificationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2ProjectNotificationConfigArgs)(nil)).Elem()
}

type SccV2ProjectNotificationConfigInput interface {
	pulumi.Input

	ToSccV2ProjectNotificationConfigOutput() SccV2ProjectNotificationConfigOutput
	ToSccV2ProjectNotificationConfigOutputWithContext(ctx context.Context) SccV2ProjectNotificationConfigOutput
}

func (*SccV2ProjectNotificationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2ProjectNotificationConfig)(nil)).Elem()
}

func (i *SccV2ProjectNotificationConfig) ToSccV2ProjectNotificationConfigOutput() SccV2ProjectNotificationConfigOutput {
	return i.ToSccV2ProjectNotificationConfigOutputWithContext(context.Background())
}

func (i *SccV2ProjectNotificationConfig) ToSccV2ProjectNotificationConfigOutputWithContext(ctx context.Context) SccV2ProjectNotificationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccV2ProjectNotificationConfigOutput)
}

type SccV2ProjectNotificationConfigOutput struct{ *pulumi.OutputState }

func (SccV2ProjectNotificationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2ProjectNotificationConfig)(nil)).Elem()
}

func (o SccV2ProjectNotificationConfigOutput) ToSccV2ProjectNotificationConfigOutput() SccV2ProjectNotificationConfigOutput {
	return o
}

func (o SccV2ProjectNotificationConfigOutput) ToSccV2ProjectNotificationConfigOutputWithContext(ctx context.Context) SccV2ProjectNotificationConfigOutput {
	return o
}

// This must be unique within the project.
func (o SccV2ProjectNotificationConfigOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// The description of the notification config (max of 1024 characters).
func (o SccV2ProjectNotificationConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Location ID of the parent organization. Only global is supported at the moment.
func (o SccV2ProjectNotificationConfigOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name of this notification config, in the format
// 'projects/{{projectId}}/locations/{{location}}/notificationConfigs/{{config_id}}'.
func (o SccV2ProjectNotificationConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SccV2ProjectNotificationConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
func (o SccV2ProjectNotificationConfigOutput) PubsubTopic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringPtrOutput { return v.PubsubTopic }).(pulumi.StringPtrOutput)
}

func (o SccV2ProjectNotificationConfigOutput) SccV2ProjectNotificationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringOutput { return v.SccV2ProjectNotificationConfigId }).(pulumi.StringOutput)
}

// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
func (o SccV2ProjectNotificationConfigOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The config for triggering streaming-based notifications.
func (o SccV2ProjectNotificationConfigOutput) StreamingConfig() SccV2ProjectNotificationConfigStreamingConfigOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) SccV2ProjectNotificationConfigStreamingConfigOutput {
		return v.StreamingConfig
	}).(SccV2ProjectNotificationConfigStreamingConfigOutput)
}

func (o SccV2ProjectNotificationConfigOutput) Timeouts() SccV2ProjectNotificationConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccV2ProjectNotificationConfig) SccV2ProjectNotificationConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(SccV2ProjectNotificationConfigTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccV2ProjectNotificationConfigInput)(nil)).Elem(), &SccV2ProjectNotificationConfig{})
	pulumi.RegisterOutputType(SccV2ProjectNotificationConfigOutput{})
}
