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

type DataprocAutoscalingPolicy struct {
	pulumi.CustomResourceState

	// Basic algorithm for autoscaling.
	BasicAlgorithm              DataprocAutoscalingPolicyBasicAlgorithmPtrOutput `pulumi:"basicAlgorithm"`
	DataprocAutoscalingPolicyId pulumi.StringOutput                              `pulumi:"dataprocAutoscalingPolicyId"`
	// The location where the autoscaling policy should reside. The default value is 'global'.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The "resource name" of the autoscaling policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig DataprocAutoscalingPolicySecondaryWorkerConfigPtrOutput `pulumi:"secondaryWorkerConfig"`
	Timeouts              DataprocAutoscalingPolicyTimeoutsPtrOutput              `pulumi:"timeouts"`
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig DataprocAutoscalingPolicyWorkerConfigPtrOutput `pulumi:"workerConfig"`
}

// NewDataprocAutoscalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataprocAutoscalingPolicy(ctx *pulumi.Context,
	name string, args *DataprocAutoscalingPolicyArgs, opts ...pulumi.ResourceOption) (*DataprocAutoscalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataprocAutoscalingPolicy
	err = ctx.RegisterPackageResource("google-beta:index/dataprocAutoscalingPolicy:DataprocAutoscalingPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocAutoscalingPolicy gets an existing DataprocAutoscalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocAutoscalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocAutoscalingPolicyState, opts ...pulumi.ResourceOption) (*DataprocAutoscalingPolicy, error) {
	var resource DataprocAutoscalingPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataprocAutoscalingPolicy:DataprocAutoscalingPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocAutoscalingPolicy resources.
type dataprocAutoscalingPolicyState struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm              *DataprocAutoscalingPolicyBasicAlgorithm `pulumi:"basicAlgorithm"`
	DataprocAutoscalingPolicyId *string                                  `pulumi:"dataprocAutoscalingPolicyId"`
	// The location where the autoscaling policy should reside. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The "resource name" of the autoscaling policy.
	Name *string `pulumi:"name"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId *string `pulumi:"policyId"`
	Project  *string `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig *DataprocAutoscalingPolicySecondaryWorkerConfig `pulumi:"secondaryWorkerConfig"`
	Timeouts              *DataprocAutoscalingPolicyTimeouts              `pulumi:"timeouts"`
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig *DataprocAutoscalingPolicyWorkerConfig `pulumi:"workerConfig"`
}

type DataprocAutoscalingPolicyState struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm              DataprocAutoscalingPolicyBasicAlgorithmPtrInput
	DataprocAutoscalingPolicyId pulumi.StringPtrInput
	// The location where the autoscaling policy should reside. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The "resource name" of the autoscaling policy.
	Name pulumi.StringPtrInput
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig DataprocAutoscalingPolicySecondaryWorkerConfigPtrInput
	Timeouts              DataprocAutoscalingPolicyTimeoutsPtrInput
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig DataprocAutoscalingPolicyWorkerConfigPtrInput
}

func (DataprocAutoscalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocAutoscalingPolicyState)(nil)).Elem()
}

type dataprocAutoscalingPolicyArgs struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm              *DataprocAutoscalingPolicyBasicAlgorithm `pulumi:"basicAlgorithm"`
	DataprocAutoscalingPolicyId *string                                  `pulumi:"dataprocAutoscalingPolicyId"`
	// The location where the autoscaling policy should reside. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId string  `pulumi:"policyId"`
	Project  *string `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig *DataprocAutoscalingPolicySecondaryWorkerConfig `pulumi:"secondaryWorkerConfig"`
	Timeouts              *DataprocAutoscalingPolicyTimeouts              `pulumi:"timeouts"`
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig *DataprocAutoscalingPolicyWorkerConfig `pulumi:"workerConfig"`
}

// The set of arguments for constructing a DataprocAutoscalingPolicy resource.
type DataprocAutoscalingPolicyArgs struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm              DataprocAutoscalingPolicyBasicAlgorithmPtrInput
	DataprocAutoscalingPolicyId pulumi.StringPtrInput
	// The location where the autoscaling policy should reside. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId pulumi.StringInput
	Project  pulumi.StringPtrInput
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig DataprocAutoscalingPolicySecondaryWorkerConfigPtrInput
	Timeouts              DataprocAutoscalingPolicyTimeoutsPtrInput
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig DataprocAutoscalingPolicyWorkerConfigPtrInput
}

func (DataprocAutoscalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocAutoscalingPolicyArgs)(nil)).Elem()
}

type DataprocAutoscalingPolicyInput interface {
	pulumi.Input

	ToDataprocAutoscalingPolicyOutput() DataprocAutoscalingPolicyOutput
	ToDataprocAutoscalingPolicyOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyOutput
}

func (*DataprocAutoscalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocAutoscalingPolicy)(nil)).Elem()
}

func (i *DataprocAutoscalingPolicy) ToDataprocAutoscalingPolicyOutput() DataprocAutoscalingPolicyOutput {
	return i.ToDataprocAutoscalingPolicyOutputWithContext(context.Background())
}

func (i *DataprocAutoscalingPolicy) ToDataprocAutoscalingPolicyOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocAutoscalingPolicyOutput)
}

type DataprocAutoscalingPolicyOutput struct{ *pulumi.OutputState }

func (DataprocAutoscalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocAutoscalingPolicy)(nil)).Elem()
}

func (o DataprocAutoscalingPolicyOutput) ToDataprocAutoscalingPolicyOutput() DataprocAutoscalingPolicyOutput {
	return o
}

func (o DataprocAutoscalingPolicyOutput) ToDataprocAutoscalingPolicyOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyOutput {
	return o
}

// Basic algorithm for autoscaling.
func (o DataprocAutoscalingPolicyOutput) BasicAlgorithm() DataprocAutoscalingPolicyBasicAlgorithmPtrOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) DataprocAutoscalingPolicyBasicAlgorithmPtrOutput {
		return v.BasicAlgorithm
	}).(DataprocAutoscalingPolicyBasicAlgorithmPtrOutput)
}

func (o DataprocAutoscalingPolicyOutput) DataprocAutoscalingPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) pulumi.StringOutput { return v.DataprocAutoscalingPolicyId }).(pulumi.StringOutput)
}

// The location where the autoscaling policy should reside. The default value is 'global'.
func (o DataprocAutoscalingPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The "resource name" of the autoscaling policy.
func (o DataprocAutoscalingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
func (o DataprocAutoscalingPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Describes how the autoscaler will operate for secondary workers.
func (o DataprocAutoscalingPolicyOutput) SecondaryWorkerConfig() DataprocAutoscalingPolicySecondaryWorkerConfigPtrOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) DataprocAutoscalingPolicySecondaryWorkerConfigPtrOutput {
		return v.SecondaryWorkerConfig
	}).(DataprocAutoscalingPolicySecondaryWorkerConfigPtrOutput)
}

func (o DataprocAutoscalingPolicyOutput) Timeouts() DataprocAutoscalingPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) DataprocAutoscalingPolicyTimeoutsPtrOutput { return v.Timeouts }).(DataprocAutoscalingPolicyTimeoutsPtrOutput)
}

// Describes how the autoscaler will operate for primary workers.
func (o DataprocAutoscalingPolicyOutput) WorkerConfig() DataprocAutoscalingPolicyWorkerConfigPtrOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicy) DataprocAutoscalingPolicyWorkerConfigPtrOutput {
		return v.WorkerConfig
	}).(DataprocAutoscalingPolicyWorkerConfigPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocAutoscalingPolicyInput)(nil)).Elem(), &DataprocAutoscalingPolicy{})
	pulumi.RegisterOutputType(DataprocAutoscalingPolicyOutput{})
}
