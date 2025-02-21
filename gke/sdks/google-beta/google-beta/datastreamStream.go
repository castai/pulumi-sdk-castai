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

type DatastreamStream struct {
	pulumi.CustomResourceState

	// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
	BackfillAll DatastreamStreamBackfillAllPtrOutput `pulumi:"backfillAll"`
	// Backfill strategy to disable automatic backfill for the Stream's objects.
	BackfillNone DatastreamStreamBackfillNonePtrOutput `pulumi:"backfillNone"`
	// Create the stream without validating it.
	CreateWithoutValidation pulumi.BoolPtrOutput `pulumi:"createWithoutValidation"`
	// A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be
	// encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey pulumi.StringPtrOutput `pulumi:"customerManagedEncryptionKey"`
	DatastreamStreamId           pulumi.StringOutput    `pulumi:"datastreamStreamId"`
	// Desired state of the Stream. Set this field to 'RUNNING' to start the stream, 'NOT_STARTED' to create the stream without
	// starting and 'PAUSED' to pause the stream from a 'RUNNING' state. Possible values: NOT_STARTED, RUNNING, PAUSED.
	// Default: NOT_STARTED
	DesiredState pulumi.StringPtrOutput `pulumi:"desiredState"`
	// Destination connection profile configuration.
	DestinationConfig DatastreamStreamDestinationConfigOutput `pulumi:"destinationConfig"`
	// Display name.
	DisplayName     pulumi.StringOutput    `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Labels. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the location this stream is located in.
	Location pulumi.StringOutput `pulumi:"location"`
	// The stream's name.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Source connection profile configuration.
	SourceConfig DatastreamStreamSourceConfigOutput `pulumi:"sourceConfig"`
	// The state of the stream.
	State pulumi.StringOutput `pulumi:"state"`
	// The stream identifier.
	StreamId pulumi.StringOutput `pulumi:"streamId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput            `pulumi:"terraformLabels"`
	Timeouts        DatastreamStreamTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDatastreamStream registers a new resource with the given unique name, arguments, and options.
func NewDatastreamStream(ctx *pulumi.Context,
	name string, args *DatastreamStreamArgs, opts ...pulumi.ResourceOption) (*DatastreamStream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationConfig == nil {
		return nil, errors.New("invalid value for required argument 'DestinationConfig'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.SourceConfig == nil {
		return nil, errors.New("invalid value for required argument 'SourceConfig'")
	}
	if args.StreamId == nil {
		return nil, errors.New("invalid value for required argument 'StreamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DatastreamStream
	err = ctx.RegisterPackageResource("google-beta:index/datastreamStream:DatastreamStream", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatastreamStream gets an existing DatastreamStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatastreamStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastreamStreamState, opts ...pulumi.ResourceOption) (*DatastreamStream, error) {
	var resource DatastreamStream
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/datastreamStream:DatastreamStream", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatastreamStream resources.
type datastreamStreamState struct {
	// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
	BackfillAll *DatastreamStreamBackfillAll `pulumi:"backfillAll"`
	// Backfill strategy to disable automatic backfill for the Stream's objects.
	BackfillNone *DatastreamStreamBackfillNone `pulumi:"backfillNone"`
	// Create the stream without validating it.
	CreateWithoutValidation *bool `pulumi:"createWithoutValidation"`
	// A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be
	// encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey *string `pulumi:"customerManagedEncryptionKey"`
	DatastreamStreamId           *string `pulumi:"datastreamStreamId"`
	// Desired state of the Stream. Set this field to 'RUNNING' to start the stream, 'NOT_STARTED' to create the stream without
	// starting and 'PAUSED' to pause the stream from a 'RUNNING' state. Possible values: NOT_STARTED, RUNNING, PAUSED.
	// Default: NOT_STARTED
	DesiredState *string `pulumi:"desiredState"`
	// Destination connection profile configuration.
	DestinationConfig *DatastreamStreamDestinationConfig `pulumi:"destinationConfig"`
	// Display name.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Labels. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this stream is located in.
	Location *string `pulumi:"location"`
	// The stream's name.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Source connection profile configuration.
	SourceConfig *DatastreamStreamSourceConfig `pulumi:"sourceConfig"`
	// The state of the stream.
	State *string `pulumi:"state"`
	// The stream identifier.
	StreamId *string `pulumi:"streamId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string         `pulumi:"terraformLabels"`
	Timeouts        *DatastreamStreamTimeouts `pulumi:"timeouts"`
}

type DatastreamStreamState struct {
	// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
	BackfillAll DatastreamStreamBackfillAllPtrInput
	// Backfill strategy to disable automatic backfill for the Stream's objects.
	BackfillNone DatastreamStreamBackfillNonePtrInput
	// Create the stream without validating it.
	CreateWithoutValidation pulumi.BoolPtrInput
	// A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be
	// encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey pulumi.StringPtrInput
	DatastreamStreamId           pulumi.StringPtrInput
	// Desired state of the Stream. Set this field to 'RUNNING' to start the stream, 'NOT_STARTED' to create the stream without
	// starting and 'PAUSED' to pause the stream from a 'RUNNING' state. Possible values: NOT_STARTED, RUNNING, PAUSED.
	// Default: NOT_STARTED
	DesiredState pulumi.StringPtrInput
	// Destination connection profile configuration.
	DestinationConfig DatastreamStreamDestinationConfigPtrInput
	// Display name.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Labels. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the location this stream is located in.
	Location pulumi.StringPtrInput
	// The stream's name.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Source connection profile configuration.
	SourceConfig DatastreamStreamSourceConfigPtrInput
	// The state of the stream.
	State pulumi.StringPtrInput
	// The stream identifier.
	StreamId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DatastreamStreamTimeoutsPtrInput
}

func (DatastreamStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastreamStreamState)(nil)).Elem()
}

type datastreamStreamArgs struct {
	// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
	BackfillAll *DatastreamStreamBackfillAll `pulumi:"backfillAll"`
	// Backfill strategy to disable automatic backfill for the Stream's objects.
	BackfillNone *DatastreamStreamBackfillNone `pulumi:"backfillNone"`
	// Create the stream without validating it.
	CreateWithoutValidation *bool `pulumi:"createWithoutValidation"`
	// A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be
	// encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey *string `pulumi:"customerManagedEncryptionKey"`
	DatastreamStreamId           *string `pulumi:"datastreamStreamId"`
	// Desired state of the Stream. Set this field to 'RUNNING' to start the stream, 'NOT_STARTED' to create the stream without
	// starting and 'PAUSED' to pause the stream from a 'RUNNING' state. Possible values: NOT_STARTED, RUNNING, PAUSED.
	// Default: NOT_STARTED
	DesiredState *string `pulumi:"desiredState"`
	// Destination connection profile configuration.
	DestinationConfig DatastreamStreamDestinationConfig `pulumi:"destinationConfig"`
	// Display name.
	DisplayName string `pulumi:"displayName"`
	// Labels. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this stream is located in.
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Source connection profile configuration.
	SourceConfig DatastreamStreamSourceConfig `pulumi:"sourceConfig"`
	// The stream identifier.
	StreamId string                    `pulumi:"streamId"`
	Timeouts *DatastreamStreamTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DatastreamStream resource.
type DatastreamStreamArgs struct {
	// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
	BackfillAll DatastreamStreamBackfillAllPtrInput
	// Backfill strategy to disable automatic backfill for the Stream's objects.
	BackfillNone DatastreamStreamBackfillNonePtrInput
	// Create the stream without validating it.
	CreateWithoutValidation pulumi.BoolPtrInput
	// A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be
	// encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey pulumi.StringPtrInput
	DatastreamStreamId           pulumi.StringPtrInput
	// Desired state of the Stream. Set this field to 'RUNNING' to start the stream, 'NOT_STARTED' to create the stream without
	// starting and 'PAUSED' to pause the stream from a 'RUNNING' state. Possible values: NOT_STARTED, RUNNING, PAUSED.
	// Default: NOT_STARTED
	DesiredState pulumi.StringPtrInput
	// Destination connection profile configuration.
	DestinationConfig DatastreamStreamDestinationConfigInput
	// Display name.
	DisplayName pulumi.StringInput
	// Labels. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the location this stream is located in.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// Source connection profile configuration.
	SourceConfig DatastreamStreamSourceConfigInput
	// The stream identifier.
	StreamId pulumi.StringInput
	Timeouts DatastreamStreamTimeoutsPtrInput
}

func (DatastreamStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastreamStreamArgs)(nil)).Elem()
}

type DatastreamStreamInput interface {
	pulumi.Input

	ToDatastreamStreamOutput() DatastreamStreamOutput
	ToDatastreamStreamOutputWithContext(ctx context.Context) DatastreamStreamOutput
}

func (*DatastreamStream) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastreamStream)(nil)).Elem()
}

func (i *DatastreamStream) ToDatastreamStreamOutput() DatastreamStreamOutput {
	return i.ToDatastreamStreamOutputWithContext(context.Background())
}

func (i *DatastreamStream) ToDatastreamStreamOutputWithContext(ctx context.Context) DatastreamStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastreamStreamOutput)
}

type DatastreamStreamOutput struct{ *pulumi.OutputState }

func (DatastreamStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastreamStream)(nil)).Elem()
}

func (o DatastreamStreamOutput) ToDatastreamStreamOutput() DatastreamStreamOutput {
	return o
}

func (o DatastreamStreamOutput) ToDatastreamStreamOutputWithContext(ctx context.Context) DatastreamStreamOutput {
	return o
}

// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
func (o DatastreamStreamOutput) BackfillAll() DatastreamStreamBackfillAllPtrOutput {
	return o.ApplyT(func(v *DatastreamStream) DatastreamStreamBackfillAllPtrOutput { return v.BackfillAll }).(DatastreamStreamBackfillAllPtrOutput)
}

// Backfill strategy to disable automatic backfill for the Stream's objects.
func (o DatastreamStreamOutput) BackfillNone() DatastreamStreamBackfillNonePtrOutput {
	return o.ApplyT(func(v *DatastreamStream) DatastreamStreamBackfillNonePtrOutput { return v.BackfillNone }).(DatastreamStreamBackfillNonePtrOutput)
}

// Create the stream without validating it.
func (o DatastreamStreamOutput) CreateWithoutValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.BoolPtrOutput { return v.CreateWithoutValidation }).(pulumi.BoolPtrOutput)
}

// A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be
// encrypted using an internal Stream-specific encryption key provisioned through KMS.
func (o DatastreamStreamOutput) CustomerManagedEncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringPtrOutput { return v.CustomerManagedEncryptionKey }).(pulumi.StringPtrOutput)
}

func (o DatastreamStreamOutput) DatastreamStreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringOutput { return v.DatastreamStreamId }).(pulumi.StringOutput)
}

// Desired state of the Stream. Set this field to 'RUNNING' to start the stream, 'NOT_STARTED' to create the stream without
// starting and 'PAUSED' to pause the stream from a 'RUNNING' state. Possible values: NOT_STARTED, RUNNING, PAUSED.
// Default: NOT_STARTED
func (o DatastreamStreamOutput) DesiredState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringPtrOutput { return v.DesiredState }).(pulumi.StringPtrOutput)
}

// Destination connection profile configuration.
func (o DatastreamStreamOutput) DestinationConfig() DatastreamStreamDestinationConfigOutput {
	return o.ApplyT(func(v *DatastreamStream) DatastreamStreamDestinationConfigOutput { return v.DestinationConfig }).(DatastreamStreamDestinationConfigOutput)
}

// Display name.
func (o DatastreamStreamOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o DatastreamStreamOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Labels. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
// refer to the field 'effective_labels' for all of the labels present on the resource.
func (o DatastreamStreamOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the location this stream is located in.
func (o DatastreamStreamOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The stream's name.
func (o DatastreamStreamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatastreamStreamOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Source connection profile configuration.
func (o DatastreamStreamOutput) SourceConfig() DatastreamStreamSourceConfigOutput {
	return o.ApplyT(func(v *DatastreamStream) DatastreamStreamSourceConfigOutput { return v.SourceConfig }).(DatastreamStreamSourceConfigOutput)
}

// The state of the stream.
func (o DatastreamStreamOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The stream identifier.
func (o DatastreamStreamOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringOutput { return v.StreamId }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DatastreamStreamOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatastreamStream) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DatastreamStreamOutput) Timeouts() DatastreamStreamTimeoutsPtrOutput {
	return o.ApplyT(func(v *DatastreamStream) DatastreamStreamTimeoutsPtrOutput { return v.Timeouts }).(DatastreamStreamTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatastreamStreamInput)(nil)).Elem(), &DatastreamStream{})
	pulumi.RegisterOutputType(DatastreamStreamOutput{})
}
