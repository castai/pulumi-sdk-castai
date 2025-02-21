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

type LoggingProjectExclusion struct {
	pulumi.CustomResourceState

	// A human-readable description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                    pulumi.StringOutput `pulumi:"filter"`
	LoggingProjectExclusionId pulumi.StringOutput `pulumi:"loggingProjectExclusionId"`
	// The name of the logging exclusion.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewLoggingProjectExclusion registers a new resource with the given unique name, arguments, and options.
func NewLoggingProjectExclusion(ctx *pulumi.Context,
	name string, args *LoggingProjectExclusionArgs, opts ...pulumi.ResourceOption) (*LoggingProjectExclusion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource LoggingProjectExclusion
	err = ctx.RegisterPackageResource("google:index/loggingProjectExclusion:LoggingProjectExclusion", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingProjectExclusion gets an existing LoggingProjectExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingProjectExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingProjectExclusionState, opts ...pulumi.ResourceOption) (*LoggingProjectExclusion, error) {
	var resource LoggingProjectExclusion
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/loggingProjectExclusion:LoggingProjectExclusion", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingProjectExclusion resources.
type loggingProjectExclusionState struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                    *string `pulumi:"filter"`
	LoggingProjectExclusionId *string `pulumi:"loggingProjectExclusionId"`
	// The name of the logging exclusion.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

type LoggingProjectExclusionState struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                    pulumi.StringPtrInput
	LoggingProjectExclusionId pulumi.StringPtrInput
	// The name of the logging exclusion.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (LoggingProjectExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingProjectExclusionState)(nil)).Elem()
}

type loggingProjectExclusionArgs struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                    string  `pulumi:"filter"`
	LoggingProjectExclusionId *string `pulumi:"loggingProjectExclusionId"`
	// The name of the logging exclusion.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a LoggingProjectExclusion resource.
type LoggingProjectExclusionArgs struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                    pulumi.StringInput
	LoggingProjectExclusionId pulumi.StringPtrInput
	// The name of the logging exclusion.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (LoggingProjectExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingProjectExclusionArgs)(nil)).Elem()
}

type LoggingProjectExclusionInput interface {
	pulumi.Input

	ToLoggingProjectExclusionOutput() LoggingProjectExclusionOutput
	ToLoggingProjectExclusionOutputWithContext(ctx context.Context) LoggingProjectExclusionOutput
}

func (*LoggingProjectExclusion) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingProjectExclusion)(nil)).Elem()
}

func (i *LoggingProjectExclusion) ToLoggingProjectExclusionOutput() LoggingProjectExclusionOutput {
	return i.ToLoggingProjectExclusionOutputWithContext(context.Background())
}

func (i *LoggingProjectExclusion) ToLoggingProjectExclusionOutputWithContext(ctx context.Context) LoggingProjectExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingProjectExclusionOutput)
}

type LoggingProjectExclusionOutput struct{ *pulumi.OutputState }

func (LoggingProjectExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingProjectExclusion)(nil)).Elem()
}

func (o LoggingProjectExclusionOutput) ToLoggingProjectExclusionOutput() LoggingProjectExclusionOutput {
	return o
}

func (o LoggingProjectExclusionOutput) ToLoggingProjectExclusionOutputWithContext(ctx context.Context) LoggingProjectExclusionOutput {
	return o
}

// A human-readable description.
func (o LoggingProjectExclusionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingProjectExclusion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether this exclusion rule should be disabled or not. This defaults to false.
func (o LoggingProjectExclusionOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoggingProjectExclusion) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
func (o LoggingProjectExclusionOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingProjectExclusion) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LoggingProjectExclusionOutput) LoggingProjectExclusionId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingProjectExclusion) pulumi.StringOutput { return v.LoggingProjectExclusionId }).(pulumi.StringOutput)
}

// The name of the logging exclusion.
func (o LoggingProjectExclusionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingProjectExclusion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LoggingProjectExclusionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingProjectExclusion) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingProjectExclusionInput)(nil)).Elem(), &LoggingProjectExclusion{})
	pulumi.RegisterOutputType(LoggingProjectExclusionOutput{})
}
