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

type BigqueryReservation struct {
	pulumi.CustomResourceState

	// The configuration parameters for the auto scaling feature.
	Autoscale             BigqueryReservationAutoscalePtrOutput `pulumi:"autoscale"`
	BigqueryReservationId pulumi.StringOutput                   `pulumi:"bigqueryReservationId"`
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to
	// asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that
	// concurrency will be automatically set based on the reservation size.
	Concurrency pulumi.Float64PtrOutput `pulumi:"concurrency"`
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition pulumi.StringOutput `pulumi:"edition"`
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrOutput `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.Float64Output                 `pulumi:"slotCapacity"`
	Timeouts     BigqueryReservationTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewBigqueryReservation registers a new resource with the given unique name, arguments, and options.
func NewBigqueryReservation(ctx *pulumi.Context,
	name string, args *BigqueryReservationArgs, opts ...pulumi.ResourceOption) (*BigqueryReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SlotCapacity == nil {
		return nil, errors.New("invalid value for required argument 'SlotCapacity'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigqueryReservation
	err = ctx.RegisterPackageResource("google-beta:index/bigqueryReservation:BigqueryReservation", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryReservation gets an existing BigqueryReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryReservationState, opts ...pulumi.ResourceOption) (*BigqueryReservation, error) {
	var resource BigqueryReservation
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/bigqueryReservation:BigqueryReservation", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryReservation resources.
type bigqueryReservationState struct {
	// The configuration parameters for the auto scaling feature.
	Autoscale             *BigqueryReservationAutoscale `pulumi:"autoscale"`
	BigqueryReservationId *string                       `pulumi:"bigqueryReservationId"`
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to
	// asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that
	// concurrency will be automatically set based on the reservation size.
	Concurrency *float64 `pulumi:"concurrency"`
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition *string `pulumi:"edition"`
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots *bool `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location *string `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity *float64                     `pulumi:"slotCapacity"`
	Timeouts     *BigqueryReservationTimeouts `pulumi:"timeouts"`
}

type BigqueryReservationState struct {
	// The configuration parameters for the auto scaling feature.
	Autoscale             BigqueryReservationAutoscalePtrInput
	BigqueryReservationId pulumi.StringPtrInput
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to
	// asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that
	// concurrency will be automatically set based on the reservation size.
	Concurrency pulumi.Float64PtrInput
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition pulumi.StringPtrInput
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrInput
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrInput
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.Float64PtrInput
	Timeouts     BigqueryReservationTimeoutsPtrInput
}

func (BigqueryReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryReservationState)(nil)).Elem()
}

type bigqueryReservationArgs struct {
	// The configuration parameters for the auto scaling feature.
	Autoscale             *BigqueryReservationAutoscale `pulumi:"autoscale"`
	BigqueryReservationId *string                       `pulumi:"bigqueryReservationId"`
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to
	// asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that
	// concurrency will be automatically set based on the reservation size.
	Concurrency *float64 `pulumi:"concurrency"`
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition *string `pulumi:"edition"`
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots *bool `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location *string `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity float64                      `pulumi:"slotCapacity"`
	Timeouts     *BigqueryReservationTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a BigqueryReservation resource.
type BigqueryReservationArgs struct {
	// The configuration parameters for the auto scaling feature.
	Autoscale             BigqueryReservationAutoscalePtrInput
	BigqueryReservationId pulumi.StringPtrInput
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to
	// asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that
	// concurrency will be automatically set based on the reservation size.
	Concurrency pulumi.Float64PtrInput
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition pulumi.StringPtrInput
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrInput
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrInput
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.Float64Input
	Timeouts     BigqueryReservationTimeoutsPtrInput
}

func (BigqueryReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryReservationArgs)(nil)).Elem()
}

type BigqueryReservationInput interface {
	pulumi.Input

	ToBigqueryReservationOutput() BigqueryReservationOutput
	ToBigqueryReservationOutputWithContext(ctx context.Context) BigqueryReservationOutput
}

func (*BigqueryReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryReservation)(nil)).Elem()
}

func (i *BigqueryReservation) ToBigqueryReservationOutput() BigqueryReservationOutput {
	return i.ToBigqueryReservationOutputWithContext(context.Background())
}

func (i *BigqueryReservation) ToBigqueryReservationOutputWithContext(ctx context.Context) BigqueryReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryReservationOutput)
}

type BigqueryReservationOutput struct{ *pulumi.OutputState }

func (BigqueryReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryReservation)(nil)).Elem()
}

func (o BigqueryReservationOutput) ToBigqueryReservationOutput() BigqueryReservationOutput {
	return o
}

func (o BigqueryReservationOutput) ToBigqueryReservationOutputWithContext(ctx context.Context) BigqueryReservationOutput {
	return o
}

// The configuration parameters for the auto scaling feature.
func (o BigqueryReservationOutput) Autoscale() BigqueryReservationAutoscalePtrOutput {
	return o.ApplyT(func(v *BigqueryReservation) BigqueryReservationAutoscalePtrOutput { return v.Autoscale }).(BigqueryReservationAutoscalePtrOutput)
}

func (o BigqueryReservationOutput) BigqueryReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.StringOutput { return v.BigqueryReservationId }).(pulumi.StringOutput)
}

// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to
// asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that
// concurrency will be automatically set based on the reservation size.
func (o BigqueryReservationOutput) Concurrency() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.Float64PtrOutput { return v.Concurrency }).(pulumi.Float64PtrOutput)
}

// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
func (o BigqueryReservationOutput) Edition() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.StringOutput { return v.Edition }).(pulumi.StringOutput)
}

// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
// true, a query using this reservation will execute with the slot capacity specified above at most.
func (o BigqueryReservationOutput) IgnoreIdleSlots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.BoolPtrOutput { return v.IgnoreIdleSlots }).(pulumi.BoolPtrOutput)
}

// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
// US.
func (o BigqueryReservationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the reservation. This field must only contain alphanumeric characters or dash.
func (o BigqueryReservationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BigqueryReservationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
func (o BigqueryReservationOutput) SlotCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v *BigqueryReservation) pulumi.Float64Output { return v.SlotCapacity }).(pulumi.Float64Output)
}

func (o BigqueryReservationOutput) Timeouts() BigqueryReservationTimeoutsPtrOutput {
	return o.ApplyT(func(v *BigqueryReservation) BigqueryReservationTimeoutsPtrOutput { return v.Timeouts }).(BigqueryReservationTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryReservationInput)(nil)).Elem(), &BigqueryReservation{})
	pulumi.RegisterOutputType(BigqueryReservationOutput{})
}
