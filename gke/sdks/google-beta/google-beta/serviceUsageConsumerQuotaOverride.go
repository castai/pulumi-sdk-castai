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

type ServiceUsageConsumerQuotaOverride struct {
	pulumi.CustomResourceState

	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions pulumi.StringMapOutput `pulumi:"dimensions"`
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The limit on the metric, e.g. '/project/region'. > Make sure that 'limit' is in a format that doesn't start with '1/' or
	// contain curly braces. E.g. use '/project/user' instead of '1/{project}/{user}'.
	Limit pulumi.StringOutput `pulumi:"limit"`
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric pulumi.StringOutput `pulumi:"metric"`
	// The server-generated name of the quota override.
	Name pulumi.StringOutput `pulumi:"name"`
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue pulumi.StringOutput `pulumi:"overrideValue"`
	Project       pulumi.StringOutput `pulumi:"project"`
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service                             pulumi.StringOutput                                `pulumi:"service"`
	ServiceUsageConsumerQuotaOverrideId pulumi.StringOutput                                `pulumi:"serviceUsageConsumerQuotaOverrideId"`
	Timeouts                            ServiceUsageConsumerQuotaOverrideTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewServiceUsageConsumerQuotaOverride registers a new resource with the given unique name, arguments, and options.
func NewServiceUsageConsumerQuotaOverride(ctx *pulumi.Context,
	name string, args *ServiceUsageConsumerQuotaOverrideArgs, opts ...pulumi.ResourceOption) (*ServiceUsageConsumerQuotaOverride, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Limit == nil {
		return nil, errors.New("invalid value for required argument 'Limit'")
	}
	if args.Metric == nil {
		return nil, errors.New("invalid value for required argument 'Metric'")
	}
	if args.OverrideValue == nil {
		return nil, errors.New("invalid value for required argument 'OverrideValue'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ServiceUsageConsumerQuotaOverride
	err = ctx.RegisterPackageResource("google-beta:index/serviceUsageConsumerQuotaOverride:ServiceUsageConsumerQuotaOverride", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceUsageConsumerQuotaOverride gets an existing ServiceUsageConsumerQuotaOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceUsageConsumerQuotaOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceUsageConsumerQuotaOverrideState, opts ...pulumi.ResourceOption) (*ServiceUsageConsumerQuotaOverride, error) {
	var resource ServiceUsageConsumerQuotaOverride
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/serviceUsageConsumerQuotaOverride:ServiceUsageConsumerQuotaOverride", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceUsageConsumerQuotaOverride resources.
type serviceUsageConsumerQuotaOverrideState struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions map[string]string `pulumi:"dimensions"`
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force *bool `pulumi:"force"`
	// The limit on the metric, e.g. '/project/region'. > Make sure that 'limit' is in a format that doesn't start with '1/' or
	// contain curly braces. E.g. use '/project/user' instead of '1/{project}/{user}'.
	Limit *string `pulumi:"limit"`
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric *string `pulumi:"metric"`
	// The server-generated name of the quota override.
	Name *string `pulumi:"name"`
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue *string `pulumi:"overrideValue"`
	Project       *string `pulumi:"project"`
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service                             *string                                    `pulumi:"service"`
	ServiceUsageConsumerQuotaOverrideId *string                                    `pulumi:"serviceUsageConsumerQuotaOverrideId"`
	Timeouts                            *ServiceUsageConsumerQuotaOverrideTimeouts `pulumi:"timeouts"`
}

type ServiceUsageConsumerQuotaOverrideState struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions pulumi.StringMapInput
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force pulumi.BoolPtrInput
	// The limit on the metric, e.g. '/project/region'. > Make sure that 'limit' is in a format that doesn't start with '1/' or
	// contain curly braces. E.g. use '/project/user' instead of '1/{project}/{user}'.
	Limit pulumi.StringPtrInput
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric pulumi.StringPtrInput
	// The server-generated name of the quota override.
	Name pulumi.StringPtrInput
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue pulumi.StringPtrInput
	Project       pulumi.StringPtrInput
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service                             pulumi.StringPtrInput
	ServiceUsageConsumerQuotaOverrideId pulumi.StringPtrInput
	Timeouts                            ServiceUsageConsumerQuotaOverrideTimeoutsPtrInput
}

func (ServiceUsageConsumerQuotaOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceUsageConsumerQuotaOverrideState)(nil)).Elem()
}

type serviceUsageConsumerQuotaOverrideArgs struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions map[string]string `pulumi:"dimensions"`
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force *bool `pulumi:"force"`
	// The limit on the metric, e.g. '/project/region'. > Make sure that 'limit' is in a format that doesn't start with '1/' or
	// contain curly braces. E.g. use '/project/user' instead of '1/{project}/{user}'.
	Limit string `pulumi:"limit"`
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric string `pulumi:"metric"`
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue string  `pulumi:"overrideValue"`
	Project       *string `pulumi:"project"`
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service                             string                                     `pulumi:"service"`
	ServiceUsageConsumerQuotaOverrideId *string                                    `pulumi:"serviceUsageConsumerQuotaOverrideId"`
	Timeouts                            *ServiceUsageConsumerQuotaOverrideTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ServiceUsageConsumerQuotaOverride resource.
type ServiceUsageConsumerQuotaOverrideArgs struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions pulumi.StringMapInput
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force pulumi.BoolPtrInput
	// The limit on the metric, e.g. '/project/region'. > Make sure that 'limit' is in a format that doesn't start with '1/' or
	// contain curly braces. E.g. use '/project/user' instead of '1/{project}/{user}'.
	Limit pulumi.StringInput
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric pulumi.StringInput
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue pulumi.StringInput
	Project       pulumi.StringPtrInput
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service                             pulumi.StringInput
	ServiceUsageConsumerQuotaOverrideId pulumi.StringPtrInput
	Timeouts                            ServiceUsageConsumerQuotaOverrideTimeoutsPtrInput
}

func (ServiceUsageConsumerQuotaOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceUsageConsumerQuotaOverrideArgs)(nil)).Elem()
}

type ServiceUsageConsumerQuotaOverrideInput interface {
	pulumi.Input

	ToServiceUsageConsumerQuotaOverrideOutput() ServiceUsageConsumerQuotaOverrideOutput
	ToServiceUsageConsumerQuotaOverrideOutputWithContext(ctx context.Context) ServiceUsageConsumerQuotaOverrideOutput
}

func (*ServiceUsageConsumerQuotaOverride) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUsageConsumerQuotaOverride)(nil)).Elem()
}

func (i *ServiceUsageConsumerQuotaOverride) ToServiceUsageConsumerQuotaOverrideOutput() ServiceUsageConsumerQuotaOverrideOutput {
	return i.ToServiceUsageConsumerQuotaOverrideOutputWithContext(context.Background())
}

func (i *ServiceUsageConsumerQuotaOverride) ToServiceUsageConsumerQuotaOverrideOutputWithContext(ctx context.Context) ServiceUsageConsumerQuotaOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUsageConsumerQuotaOverrideOutput)
}

type ServiceUsageConsumerQuotaOverrideOutput struct{ *pulumi.OutputState }

func (ServiceUsageConsumerQuotaOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUsageConsumerQuotaOverride)(nil)).Elem()
}

func (o ServiceUsageConsumerQuotaOverrideOutput) ToServiceUsageConsumerQuotaOverrideOutput() ServiceUsageConsumerQuotaOverrideOutput {
	return o
}

func (o ServiceUsageConsumerQuotaOverrideOutput) ToServiceUsageConsumerQuotaOverrideOutputWithContext(ctx context.Context) ServiceUsageConsumerQuotaOverrideOutput {
	return o
}

// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
func (o ServiceUsageConsumerQuotaOverrideOutput) Dimensions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringMapOutput { return v.Dimensions }).(pulumi.StringMapOutput)
}

// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
// safety check is ignored.
func (o ServiceUsageConsumerQuotaOverrideOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// The limit on the metric, e.g. '/project/region'. > Make sure that 'limit' is in a format that doesn't start with '1/' or
// contain curly braces. E.g. use '/project/user' instead of '1/{project}/{user}'.
func (o ServiceUsageConsumerQuotaOverrideOutput) Limit() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringOutput { return v.Limit }).(pulumi.StringOutput)
}

// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
func (o ServiceUsageConsumerQuotaOverrideOutput) Metric() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringOutput { return v.Metric }).(pulumi.StringOutput)
}

// The server-generated name of the quota override.
func (o ServiceUsageConsumerQuotaOverrideOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
func (o ServiceUsageConsumerQuotaOverrideOutput) OverrideValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringOutput { return v.OverrideValue }).(pulumi.StringOutput)
}

func (o ServiceUsageConsumerQuotaOverrideOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
func (o ServiceUsageConsumerQuotaOverrideOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

func (o ServiceUsageConsumerQuotaOverrideOutput) ServiceUsageConsumerQuotaOverrideId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) pulumi.StringOutput {
		return v.ServiceUsageConsumerQuotaOverrideId
	}).(pulumi.StringOutput)
}

func (o ServiceUsageConsumerQuotaOverrideOutput) Timeouts() ServiceUsageConsumerQuotaOverrideTimeoutsPtrOutput {
	return o.ApplyT(func(v *ServiceUsageConsumerQuotaOverride) ServiceUsageConsumerQuotaOverrideTimeoutsPtrOutput {
		return v.Timeouts
	}).(ServiceUsageConsumerQuotaOverrideTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceUsageConsumerQuotaOverrideInput)(nil)).Elem(), &ServiceUsageConsumerQuotaOverride{})
	pulumi.RegisterOutputType(ServiceUsageConsumerQuotaOverrideOutput{})
}
