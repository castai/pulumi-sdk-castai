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

type FirebaseHostingChannel struct {
	pulumi.CustomResourceState

	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId       pulumi.StringOutput    `pulumi:"channelId"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted.
	// This field is present in the output whether it's set directly or via the 'ttl' field.
	ExpireTime               pulumi.StringOutput `pulumi:"expireTime"`
	FirebaseHostingChannelId pulumi.StringOutput `pulumi:"firebaseHostingChannelId"`
	// Text labels used for extra metadata and/or filtering **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format: sites/SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100.
	// Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.Float64Output `pulumi:"retainedReleaseCount"`
	// Required. The ID of the site in which to create this channel.
	SiteId pulumi.StringOutput `pulumi:"siteId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                  `pulumi:"terraformLabels"`
	Timeouts        FirebaseHostingChannelTimeoutsPtrOutput `pulumi:"timeouts"`
	// Input only. A time-to-live for this channel. Sets 'expire_time' to the provided duration past the time of the request. A
	// duration in seconds with up to nine fractional digits, terminated by 's'. Example: "86400s" (one day).
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
}

// NewFirebaseHostingChannel registers a new resource with the given unique name, arguments, and options.
func NewFirebaseHostingChannel(ctx *pulumi.Context,
	name string, args *FirebaseHostingChannelArgs, opts ...pulumi.ResourceOption) (*FirebaseHostingChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChannelId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelId'")
	}
	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseHostingChannel
	err = ctx.RegisterPackageResource("google-beta:index/firebaseHostingChannel:FirebaseHostingChannel", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseHostingChannel gets an existing FirebaseHostingChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseHostingChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseHostingChannelState, opts ...pulumi.ResourceOption) (*FirebaseHostingChannel, error) {
	var resource FirebaseHostingChannel
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/firebaseHostingChannel:FirebaseHostingChannel", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseHostingChannel resources.
type firebaseHostingChannelState struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId       *string           `pulumi:"channelId"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted.
	// This field is present in the output whether it's set directly or via the 'ttl' field.
	ExpireTime               *string `pulumi:"expireTime"`
	FirebaseHostingChannelId *string `pulumi:"firebaseHostingChannelId"`
	// Text labels used for extra metadata and/or filtering **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format: sites/SITE_ID/channels/CHANNEL_ID
	Name *string `pulumi:"name"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100.
	// Defaults to 10 for new channels.
	RetainedReleaseCount *float64 `pulumi:"retainedReleaseCount"`
	// Required. The ID of the site in which to create this channel.
	SiteId *string `pulumi:"siteId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string               `pulumi:"terraformLabels"`
	Timeouts        *FirebaseHostingChannelTimeouts `pulumi:"timeouts"`
	// Input only. A time-to-live for this channel. Sets 'expire_time' to the provided duration past the time of the request. A
	// duration in seconds with up to nine fractional digits, terminated by 's'. Example: "86400s" (one day).
	Ttl *string `pulumi:"ttl"`
}

type FirebaseHostingChannelState struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId       pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted.
	// This field is present in the output whether it's set directly or via the 'ttl' field.
	ExpireTime               pulumi.StringPtrInput
	FirebaseHostingChannelId pulumi.StringPtrInput
	// Text labels used for extra metadata and/or filtering **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The fully-qualified resource name for the channel, in the format: sites/SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringPtrInput
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100.
	// Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.Float64PtrInput
	// Required. The ID of the site in which to create this channel.
	SiteId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        FirebaseHostingChannelTimeoutsPtrInput
	// Input only. A time-to-live for this channel. Sets 'expire_time' to the provided duration past the time of the request. A
	// duration in seconds with up to nine fractional digits, terminated by 's'. Example: "86400s" (one day).
	Ttl pulumi.StringPtrInput
}

func (FirebaseHostingChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseHostingChannelState)(nil)).Elem()
}

type firebaseHostingChannelArgs struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId string `pulumi:"channelId"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted.
	// This field is present in the output whether it's set directly or via the 'ttl' field.
	ExpireTime               *string `pulumi:"expireTime"`
	FirebaseHostingChannelId *string `pulumi:"firebaseHostingChannelId"`
	// Text labels used for extra metadata and/or filtering **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100.
	// Defaults to 10 for new channels.
	RetainedReleaseCount *float64 `pulumi:"retainedReleaseCount"`
	// Required. The ID of the site in which to create this channel.
	SiteId   string                          `pulumi:"siteId"`
	Timeouts *FirebaseHostingChannelTimeouts `pulumi:"timeouts"`
	// Input only. A time-to-live for this channel. Sets 'expire_time' to the provided duration past the time of the request. A
	// duration in seconds with up to nine fractional digits, terminated by 's'. Example: "86400s" (one day).
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a FirebaseHostingChannel resource.
type FirebaseHostingChannelArgs struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId pulumi.StringInput
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted.
	// This field is present in the output whether it's set directly or via the 'ttl' field.
	ExpireTime               pulumi.StringPtrInput
	FirebaseHostingChannelId pulumi.StringPtrInput
	// Text labels used for extra metadata and/or filtering **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100.
	// Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.Float64PtrInput
	// Required. The ID of the site in which to create this channel.
	SiteId   pulumi.StringInput
	Timeouts FirebaseHostingChannelTimeoutsPtrInput
	// Input only. A time-to-live for this channel. Sets 'expire_time' to the provided duration past the time of the request. A
	// duration in seconds with up to nine fractional digits, terminated by 's'. Example: "86400s" (one day).
	Ttl pulumi.StringPtrInput
}

func (FirebaseHostingChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseHostingChannelArgs)(nil)).Elem()
}

type FirebaseHostingChannelInput interface {
	pulumi.Input

	ToFirebaseHostingChannelOutput() FirebaseHostingChannelOutput
	ToFirebaseHostingChannelOutputWithContext(ctx context.Context) FirebaseHostingChannelOutput
}

func (*FirebaseHostingChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseHostingChannel)(nil)).Elem()
}

func (i *FirebaseHostingChannel) ToFirebaseHostingChannelOutput() FirebaseHostingChannelOutput {
	return i.ToFirebaseHostingChannelOutputWithContext(context.Background())
}

func (i *FirebaseHostingChannel) ToFirebaseHostingChannelOutputWithContext(ctx context.Context) FirebaseHostingChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseHostingChannelOutput)
}

type FirebaseHostingChannelOutput struct{ *pulumi.OutputState }

func (FirebaseHostingChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseHostingChannel)(nil)).Elem()
}

func (o FirebaseHostingChannelOutput) ToFirebaseHostingChannelOutput() FirebaseHostingChannelOutput {
	return o
}

func (o FirebaseHostingChannelOutput) ToFirebaseHostingChannelOutputWithContext(ctx context.Context) FirebaseHostingChannelOutput {
	return o
}

// Required. Immutable. A unique ID within the site that identifies the channel.
func (o FirebaseHostingChannelOutput) ChannelId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringOutput { return v.ChannelId }).(pulumi.StringOutput)
}

func (o FirebaseHostingChannelOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted.
// This field is present in the output whether it's set directly or via the 'ttl' field.
func (o FirebaseHostingChannelOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

func (o FirebaseHostingChannelOutput) FirebaseHostingChannelId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringOutput { return v.FirebaseHostingChannelId }).(pulumi.StringOutput)
}

// Text labels used for extra metadata and/or filtering **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o FirebaseHostingChannelOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The fully-qualified resource name for the channel, in the format: sites/SITE_ID/channels/CHANNEL_ID
func (o FirebaseHostingChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100.
// Defaults to 10 for new channels.
func (o FirebaseHostingChannelOutput) RetainedReleaseCount() pulumi.Float64Output {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.Float64Output { return v.RetainedReleaseCount }).(pulumi.Float64Output)
}

// Required. The ID of the site in which to create this channel.
func (o FirebaseHostingChannelOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringOutput { return v.SiteId }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o FirebaseHostingChannelOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o FirebaseHostingChannelOutput) Timeouts() FirebaseHostingChannelTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) FirebaseHostingChannelTimeoutsPtrOutput { return v.Timeouts }).(FirebaseHostingChannelTimeoutsPtrOutput)
}

// Input only. A time-to-live for this channel. Sets 'expire_time' to the provided duration past the time of the request. A
// duration in seconds with up to nine fractional digits, terminated by 's'. Example: "86400s" (one day).
func (o FirebaseHostingChannelOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirebaseHostingChannel) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseHostingChannelInput)(nil)).Elem(), &FirebaseHostingChannel{})
	pulumi.RegisterOutputType(FirebaseHostingChannelOutput{})
}
