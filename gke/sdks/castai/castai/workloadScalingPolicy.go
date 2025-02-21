// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package castai

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadScalingPolicy struct {
	pulumi.CustomResourceState

	AntiAffinity WorkloadScalingPolicyAntiAffinityPtrOutput `pulumi:"antiAffinity"`
	// Recommendation apply type. - IMMEDIATE - pods are restarted immediately when new recommendation is generated. - DEFERRED
	// - pods are not restarted and recommendation values are applied during natural restarts only (new deployment, etc.)
	ApplyType pulumi.StringOutput `pulumi:"applyType"`
	// CAST AI cluster id
	ClusterId   pulumi.StringOutput                       `pulumi:"clusterId"`
	Cpu         WorkloadScalingPolicyCpuOutput            `pulumi:"cpu"`
	Downscaling WorkloadScalingPolicyDownscalingPtrOutput `pulumi:"downscaling"`
	// Defines possible options for workload management. - READ_ONLY - workload watched (metrics collected), but no actions
	// performed by CAST AI. - MANAGED - workload watched (metrics collected), CAST AI may perform actions on the workload.
	ManagementOption pulumi.StringOutput                       `pulumi:"managementOption"`
	Memory           WorkloadScalingPolicyMemoryOutput         `pulumi:"memory"`
	MemoryEvent      WorkloadScalingPolicyMemoryEventPtrOutput `pulumi:"memoryEvent"`
	// Scaling policy name
	Name                    pulumi.StringOutput                    `pulumi:"name"`
	Startup                 WorkloadScalingPolicyStartupPtrOutput  `pulumi:"startup"`
	Timeouts                WorkloadScalingPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
	WorkloadScalingPolicyId pulumi.StringOutput                    `pulumi:"workloadScalingPolicyId"`
}

// NewWorkloadScalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewWorkloadScalingPolicy(ctx *pulumi.Context,
	name string, args *WorkloadScalingPolicyArgs, opts ...pulumi.ResourceOption) (*WorkloadScalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplyType == nil {
		return nil, errors.New("invalid value for required argument 'ApplyType'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Cpu == nil {
		return nil, errors.New("invalid value for required argument 'Cpu'")
	}
	if args.ManagementOption == nil {
		return nil, errors.New("invalid value for required argument 'ManagementOption'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource WorkloadScalingPolicy
	err = ctx.RegisterPackageResource("castai:index/workloadScalingPolicy:WorkloadScalingPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadScalingPolicy gets an existing WorkloadScalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadScalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadScalingPolicyState, opts ...pulumi.ResourceOption) (*WorkloadScalingPolicy, error) {
	var resource WorkloadScalingPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("castai:index/workloadScalingPolicy:WorkloadScalingPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadScalingPolicy resources.
type workloadScalingPolicyState struct {
	AntiAffinity *WorkloadScalingPolicyAntiAffinity `pulumi:"antiAffinity"`
	// Recommendation apply type. - IMMEDIATE - pods are restarted immediately when new recommendation is generated. - DEFERRED
	// - pods are not restarted and recommendation values are applied during natural restarts only (new deployment, etc.)
	ApplyType *string `pulumi:"applyType"`
	// CAST AI cluster id
	ClusterId   *string                           `pulumi:"clusterId"`
	Cpu         *WorkloadScalingPolicyCpu         `pulumi:"cpu"`
	Downscaling *WorkloadScalingPolicyDownscaling `pulumi:"downscaling"`
	// Defines possible options for workload management. - READ_ONLY - workload watched (metrics collected), but no actions
	// performed by CAST AI. - MANAGED - workload watched (metrics collected), CAST AI may perform actions on the workload.
	ManagementOption *string                           `pulumi:"managementOption"`
	Memory           *WorkloadScalingPolicyMemory      `pulumi:"memory"`
	MemoryEvent      *WorkloadScalingPolicyMemoryEvent `pulumi:"memoryEvent"`
	// Scaling policy name
	Name                    *string                        `pulumi:"name"`
	Startup                 *WorkloadScalingPolicyStartup  `pulumi:"startup"`
	Timeouts                *WorkloadScalingPolicyTimeouts `pulumi:"timeouts"`
	WorkloadScalingPolicyId *string                        `pulumi:"workloadScalingPolicyId"`
}

type WorkloadScalingPolicyState struct {
	AntiAffinity WorkloadScalingPolicyAntiAffinityPtrInput
	// Recommendation apply type. - IMMEDIATE - pods are restarted immediately when new recommendation is generated. - DEFERRED
	// - pods are not restarted and recommendation values are applied during natural restarts only (new deployment, etc.)
	ApplyType pulumi.StringPtrInput
	// CAST AI cluster id
	ClusterId   pulumi.StringPtrInput
	Cpu         WorkloadScalingPolicyCpuPtrInput
	Downscaling WorkloadScalingPolicyDownscalingPtrInput
	// Defines possible options for workload management. - READ_ONLY - workload watched (metrics collected), but no actions
	// performed by CAST AI. - MANAGED - workload watched (metrics collected), CAST AI may perform actions on the workload.
	ManagementOption pulumi.StringPtrInput
	Memory           WorkloadScalingPolicyMemoryPtrInput
	MemoryEvent      WorkloadScalingPolicyMemoryEventPtrInput
	// Scaling policy name
	Name                    pulumi.StringPtrInput
	Startup                 WorkloadScalingPolicyStartupPtrInput
	Timeouts                WorkloadScalingPolicyTimeoutsPtrInput
	WorkloadScalingPolicyId pulumi.StringPtrInput
}

func (WorkloadScalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadScalingPolicyState)(nil)).Elem()
}

type workloadScalingPolicyArgs struct {
	AntiAffinity *WorkloadScalingPolicyAntiAffinity `pulumi:"antiAffinity"`
	// Recommendation apply type. - IMMEDIATE - pods are restarted immediately when new recommendation is generated. - DEFERRED
	// - pods are not restarted and recommendation values are applied during natural restarts only (new deployment, etc.)
	ApplyType string `pulumi:"applyType"`
	// CAST AI cluster id
	ClusterId   string                            `pulumi:"clusterId"`
	Cpu         WorkloadScalingPolicyCpu          `pulumi:"cpu"`
	Downscaling *WorkloadScalingPolicyDownscaling `pulumi:"downscaling"`
	// Defines possible options for workload management. - READ_ONLY - workload watched (metrics collected), but no actions
	// performed by CAST AI. - MANAGED - workload watched (metrics collected), CAST AI may perform actions on the workload.
	ManagementOption string                            `pulumi:"managementOption"`
	Memory           WorkloadScalingPolicyMemory       `pulumi:"memory"`
	MemoryEvent      *WorkloadScalingPolicyMemoryEvent `pulumi:"memoryEvent"`
	// Scaling policy name
	Name                    *string                        `pulumi:"name"`
	Startup                 *WorkloadScalingPolicyStartup  `pulumi:"startup"`
	Timeouts                *WorkloadScalingPolicyTimeouts `pulumi:"timeouts"`
	WorkloadScalingPolicyId *string                        `pulumi:"workloadScalingPolicyId"`
}

// The set of arguments for constructing a WorkloadScalingPolicy resource.
type WorkloadScalingPolicyArgs struct {
	AntiAffinity WorkloadScalingPolicyAntiAffinityPtrInput
	// Recommendation apply type. - IMMEDIATE - pods are restarted immediately when new recommendation is generated. - DEFERRED
	// - pods are not restarted and recommendation values are applied during natural restarts only (new deployment, etc.)
	ApplyType pulumi.StringInput
	// CAST AI cluster id
	ClusterId   pulumi.StringInput
	Cpu         WorkloadScalingPolicyCpuInput
	Downscaling WorkloadScalingPolicyDownscalingPtrInput
	// Defines possible options for workload management. - READ_ONLY - workload watched (metrics collected), but no actions
	// performed by CAST AI. - MANAGED - workload watched (metrics collected), CAST AI may perform actions on the workload.
	ManagementOption pulumi.StringInput
	Memory           WorkloadScalingPolicyMemoryInput
	MemoryEvent      WorkloadScalingPolicyMemoryEventPtrInput
	// Scaling policy name
	Name                    pulumi.StringPtrInput
	Startup                 WorkloadScalingPolicyStartupPtrInput
	Timeouts                WorkloadScalingPolicyTimeoutsPtrInput
	WorkloadScalingPolicyId pulumi.StringPtrInput
}

func (WorkloadScalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadScalingPolicyArgs)(nil)).Elem()
}

type WorkloadScalingPolicyInput interface {
	pulumi.Input

	ToWorkloadScalingPolicyOutput() WorkloadScalingPolicyOutput
	ToWorkloadScalingPolicyOutputWithContext(ctx context.Context) WorkloadScalingPolicyOutput
}

func (*WorkloadScalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadScalingPolicy)(nil)).Elem()
}

func (i *WorkloadScalingPolicy) ToWorkloadScalingPolicyOutput() WorkloadScalingPolicyOutput {
	return i.ToWorkloadScalingPolicyOutputWithContext(context.Background())
}

func (i *WorkloadScalingPolicy) ToWorkloadScalingPolicyOutputWithContext(ctx context.Context) WorkloadScalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadScalingPolicyOutput)
}

type WorkloadScalingPolicyOutput struct{ *pulumi.OutputState }

func (WorkloadScalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadScalingPolicy)(nil)).Elem()
}

func (o WorkloadScalingPolicyOutput) ToWorkloadScalingPolicyOutput() WorkloadScalingPolicyOutput {
	return o
}

func (o WorkloadScalingPolicyOutput) ToWorkloadScalingPolicyOutputWithContext(ctx context.Context) WorkloadScalingPolicyOutput {
	return o
}

func (o WorkloadScalingPolicyOutput) AntiAffinity() WorkloadScalingPolicyAntiAffinityPtrOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) WorkloadScalingPolicyAntiAffinityPtrOutput { return v.AntiAffinity }).(WorkloadScalingPolicyAntiAffinityPtrOutput)
}

// Recommendation apply type. - IMMEDIATE - pods are restarted immediately when new recommendation is generated. - DEFERRED
// - pods are not restarted and recommendation values are applied during natural restarts only (new deployment, etc.)
func (o WorkloadScalingPolicyOutput) ApplyType() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) pulumi.StringOutput { return v.ApplyType }).(pulumi.StringOutput)
}

// CAST AI cluster id
func (o WorkloadScalingPolicyOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o WorkloadScalingPolicyOutput) Cpu() WorkloadScalingPolicyCpuOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) WorkloadScalingPolicyCpuOutput { return v.Cpu }).(WorkloadScalingPolicyCpuOutput)
}

func (o WorkloadScalingPolicyOutput) Downscaling() WorkloadScalingPolicyDownscalingPtrOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) WorkloadScalingPolicyDownscalingPtrOutput { return v.Downscaling }).(WorkloadScalingPolicyDownscalingPtrOutput)
}

// Defines possible options for workload management. - READ_ONLY - workload watched (metrics collected), but no actions
// performed by CAST AI. - MANAGED - workload watched (metrics collected), CAST AI may perform actions on the workload.
func (o WorkloadScalingPolicyOutput) ManagementOption() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) pulumi.StringOutput { return v.ManagementOption }).(pulumi.StringOutput)
}

func (o WorkloadScalingPolicyOutput) Memory() WorkloadScalingPolicyMemoryOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) WorkloadScalingPolicyMemoryOutput { return v.Memory }).(WorkloadScalingPolicyMemoryOutput)
}

func (o WorkloadScalingPolicyOutput) MemoryEvent() WorkloadScalingPolicyMemoryEventPtrOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) WorkloadScalingPolicyMemoryEventPtrOutput { return v.MemoryEvent }).(WorkloadScalingPolicyMemoryEventPtrOutput)
}

// Scaling policy name
func (o WorkloadScalingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadScalingPolicyOutput) Startup() WorkloadScalingPolicyStartupPtrOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) WorkloadScalingPolicyStartupPtrOutput { return v.Startup }).(WorkloadScalingPolicyStartupPtrOutput)
}

func (o WorkloadScalingPolicyOutput) Timeouts() WorkloadScalingPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) WorkloadScalingPolicyTimeoutsPtrOutput { return v.Timeouts }).(WorkloadScalingPolicyTimeoutsPtrOutput)
}

func (o WorkloadScalingPolicyOutput) WorkloadScalingPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadScalingPolicy) pulumi.StringOutput { return v.WorkloadScalingPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadScalingPolicyInput)(nil)).Elem(), &WorkloadScalingPolicy{})
	pulumi.RegisterOutputType(WorkloadScalingPolicyOutput{})
}
