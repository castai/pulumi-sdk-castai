// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudSchedulerJob struct {
	pulumi.CustomResourceState

	// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the service
	// instance
	AppEngineHttpTarget CloudSchedulerJobAppEngineHttpTargetPtrOutput `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled
	// and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in execution logs. Cloud
	// Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP
	// targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * **Note**:
	// For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff. A duration in seconds with
	// up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline     pulumi.StringPtrOutput `pulumi:"attemptDeadline"`
	CloudSchedulerJobId pulumi.StringOutput    `pulumi:"cloudSchedulerJobId"`
	// A human-readable description for the job. This string must not contain more than 500 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
	HttpTarget CloudSchedulerJobHttpTargetPtrOutput `pulumi:"httpTarget"`
	// The name of the job.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused  pulumi.BoolOutput   `pulumi:"paused"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
	PubsubTarget CloudSchedulerJobPubsubTargetPtrOutput `pulumi:"pubsubTarget"`
	Region       pulumi.StringOutput                    `pulumi:"region"`
	// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	RetryConfig CloudSchedulerJobRetryConfigPtrOutput `pulumi:"retryConfig"`
	// Describes the schedule on which the job will be executed.
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// State of the job.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the
	// tz database.
	TimeZone pulumi.StringPtrOutput             `pulumi:"timeZone"`
	Timeouts CloudSchedulerJobTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewCloudSchedulerJob registers a new resource with the given unique name, arguments, and options.
func NewCloudSchedulerJob(ctx *pulumi.Context,
	name string, args *CloudSchedulerJobArgs, opts ...pulumi.ResourceOption) (*CloudSchedulerJob, error) {
	if args == nil {
		args = &CloudSchedulerJobArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudSchedulerJob
	err = ctx.RegisterPackageResource("google:index/cloudSchedulerJob:CloudSchedulerJob", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudSchedulerJob gets an existing CloudSchedulerJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudSchedulerJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudSchedulerJobState, opts ...pulumi.ResourceOption) (*CloudSchedulerJob, error) {
	var resource CloudSchedulerJob
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/cloudSchedulerJob:CloudSchedulerJob", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudSchedulerJob resources.
type cloudSchedulerJobState struct {
	// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the service
	// instance
	AppEngineHttpTarget *CloudSchedulerJobAppEngineHttpTarget `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled
	// and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in execution logs. Cloud
	// Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP
	// targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * **Note**:
	// For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff. A duration in seconds with
	// up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline     *string `pulumi:"attemptDeadline"`
	CloudSchedulerJobId *string `pulumi:"cloudSchedulerJobId"`
	// A human-readable description for the job. This string must not contain more than 500 characters.
	Description *string `pulumi:"description"`
	// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
	HttpTarget *CloudSchedulerJobHttpTarget `pulumi:"httpTarget"`
	// The name of the job.
	Name *string `pulumi:"name"`
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused  *bool   `pulumi:"paused"`
	Project *string `pulumi:"project"`
	// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
	PubsubTarget *CloudSchedulerJobPubsubTarget `pulumi:"pubsubTarget"`
	Region       *string                        `pulumi:"region"`
	// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	RetryConfig *CloudSchedulerJobRetryConfig `pulumi:"retryConfig"`
	// Describes the schedule on which the job will be executed.
	Schedule *string `pulumi:"schedule"`
	// State of the job.
	State *string `pulumi:"state"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the
	// tz database.
	TimeZone *string                    `pulumi:"timeZone"`
	Timeouts *CloudSchedulerJobTimeouts `pulumi:"timeouts"`
}

type CloudSchedulerJobState struct {
	// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the service
	// instance
	AppEngineHttpTarget CloudSchedulerJobAppEngineHttpTargetPtrInput
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled
	// and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in execution logs. Cloud
	// Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP
	// targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * **Note**:
	// For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff. A duration in seconds with
	// up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline     pulumi.StringPtrInput
	CloudSchedulerJobId pulumi.StringPtrInput
	// A human-readable description for the job. This string must not contain more than 500 characters.
	Description pulumi.StringPtrInput
	// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
	HttpTarget CloudSchedulerJobHttpTargetPtrInput
	// The name of the job.
	Name pulumi.StringPtrInput
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused  pulumi.BoolPtrInput
	Project pulumi.StringPtrInput
	// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
	PubsubTarget CloudSchedulerJobPubsubTargetPtrInput
	Region       pulumi.StringPtrInput
	// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	RetryConfig CloudSchedulerJobRetryConfigPtrInput
	// Describes the schedule on which the job will be executed.
	Schedule pulumi.StringPtrInput
	// State of the job.
	State pulumi.StringPtrInput
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the
	// tz database.
	TimeZone pulumi.StringPtrInput
	Timeouts CloudSchedulerJobTimeoutsPtrInput
}

func (CloudSchedulerJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudSchedulerJobState)(nil)).Elem()
}

type cloudSchedulerJobArgs struct {
	// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the service
	// instance
	AppEngineHttpTarget *CloudSchedulerJobAppEngineHttpTarget `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled
	// and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in execution logs. Cloud
	// Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP
	// targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * **Note**:
	// For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff. A duration in seconds with
	// up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline     *string `pulumi:"attemptDeadline"`
	CloudSchedulerJobId *string `pulumi:"cloudSchedulerJobId"`
	// A human-readable description for the job. This string must not contain more than 500 characters.
	Description *string `pulumi:"description"`
	// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
	HttpTarget *CloudSchedulerJobHttpTarget `pulumi:"httpTarget"`
	// The name of the job.
	Name *string `pulumi:"name"`
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused  *bool   `pulumi:"paused"`
	Project *string `pulumi:"project"`
	// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
	PubsubTarget *CloudSchedulerJobPubsubTarget `pulumi:"pubsubTarget"`
	Region       *string                        `pulumi:"region"`
	// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	RetryConfig *CloudSchedulerJobRetryConfig `pulumi:"retryConfig"`
	// Describes the schedule on which the job will be executed.
	Schedule *string `pulumi:"schedule"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the
	// tz database.
	TimeZone *string                    `pulumi:"timeZone"`
	Timeouts *CloudSchedulerJobTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a CloudSchedulerJob resource.
type CloudSchedulerJobArgs struct {
	// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the service
	// instance
	AppEngineHttpTarget CloudSchedulerJobAppEngineHttpTargetPtrInput
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled
	// and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in execution logs. Cloud
	// Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP
	// targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * **Note**:
	// For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff. A duration in seconds with
	// up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline     pulumi.StringPtrInput
	CloudSchedulerJobId pulumi.StringPtrInput
	// A human-readable description for the job. This string must not contain more than 500 characters.
	Description pulumi.StringPtrInput
	// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
	HttpTarget CloudSchedulerJobHttpTargetPtrInput
	// The name of the job.
	Name pulumi.StringPtrInput
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused  pulumi.BoolPtrInput
	Project pulumi.StringPtrInput
	// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
	PubsubTarget CloudSchedulerJobPubsubTargetPtrInput
	Region       pulumi.StringPtrInput
	// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	RetryConfig CloudSchedulerJobRetryConfigPtrInput
	// Describes the schedule on which the job will be executed.
	Schedule pulumi.StringPtrInput
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the
	// tz database.
	TimeZone pulumi.StringPtrInput
	Timeouts CloudSchedulerJobTimeoutsPtrInput
}

func (CloudSchedulerJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudSchedulerJobArgs)(nil)).Elem()
}

type CloudSchedulerJobInput interface {
	pulumi.Input

	ToCloudSchedulerJobOutput() CloudSchedulerJobOutput
	ToCloudSchedulerJobOutputWithContext(ctx context.Context) CloudSchedulerJobOutput
}

func (*CloudSchedulerJob) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudSchedulerJob)(nil)).Elem()
}

func (i *CloudSchedulerJob) ToCloudSchedulerJobOutput() CloudSchedulerJobOutput {
	return i.ToCloudSchedulerJobOutputWithContext(context.Background())
}

func (i *CloudSchedulerJob) ToCloudSchedulerJobOutputWithContext(ctx context.Context) CloudSchedulerJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudSchedulerJobOutput)
}

type CloudSchedulerJobOutput struct{ *pulumi.OutputState }

func (CloudSchedulerJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudSchedulerJob)(nil)).Elem()
}

func (o CloudSchedulerJobOutput) ToCloudSchedulerJobOutput() CloudSchedulerJobOutput {
	return o
}

func (o CloudSchedulerJobOutput) ToCloudSchedulerJobOutputWithContext(ctx context.Context) CloudSchedulerJobOutput {
	return o
}

// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the service
// instance
func (o CloudSchedulerJobOutput) AppEngineHttpTarget() CloudSchedulerJobAppEngineHttpTargetPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) CloudSchedulerJobAppEngineHttpTargetPtrOutput { return v.AppEngineHttpTarget }).(CloudSchedulerJobAppEngineHttpTargetPtrOutput)
}

// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled
// and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in execution logs. Cloud
// Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP
// targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. * **Note**:
// For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff. A duration in seconds with
// up to nine fractional digits, terminated by 's'. Example: "3.5s"
func (o CloudSchedulerJobOutput) AttemptDeadline() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringPtrOutput { return v.AttemptDeadline }).(pulumi.StringPtrOutput)
}

func (o CloudSchedulerJobOutput) CloudSchedulerJobId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringOutput { return v.CloudSchedulerJobId }).(pulumi.StringOutput)
}

// A human-readable description for the job. This string must not contain more than 500 characters.
func (o CloudSchedulerJobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
func (o CloudSchedulerJobOutput) HttpTarget() CloudSchedulerJobHttpTargetPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) CloudSchedulerJobHttpTargetPtrOutput { return v.HttpTarget }).(CloudSchedulerJobHttpTargetPtrOutput)
}

// The name of the job.
func (o CloudSchedulerJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
func (o CloudSchedulerJobOutput) Paused() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.BoolOutput { return v.Paused }).(pulumi.BoolOutput)
}

func (o CloudSchedulerJobOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
func (o CloudSchedulerJobOutput) PubsubTarget() CloudSchedulerJobPubsubTargetPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) CloudSchedulerJobPubsubTargetPtrOutput { return v.PubsubTarget }).(CloudSchedulerJobPubsubTargetPtrOutput)
}

func (o CloudSchedulerJobOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from the handler,
// then it will be retried with exponential backoff according to the settings
func (o CloudSchedulerJobOutput) RetryConfig() CloudSchedulerJobRetryConfigPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) CloudSchedulerJobRetryConfigPtrOutput { return v.RetryConfig }).(CloudSchedulerJobRetryConfigPtrOutput)
}

// Describes the schedule on which the job will be executed.
func (o CloudSchedulerJobOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringPtrOutput { return v.Schedule }).(pulumi.StringPtrOutput)
}

// State of the job.
func (o CloudSchedulerJobOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the
// tz database.
func (o CloudSchedulerJobOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o CloudSchedulerJobOutput) Timeouts() CloudSchedulerJobTimeoutsPtrOutput {
	return o.ApplyT(func(v *CloudSchedulerJob) CloudSchedulerJobTimeoutsPtrOutput { return v.Timeouts }).(CloudSchedulerJobTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudSchedulerJobInput)(nil)).Elem(), &CloudSchedulerJob{})
	pulumi.RegisterOutputType(CloudSchedulerJobOutput{})
}
