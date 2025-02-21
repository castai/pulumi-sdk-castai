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

type HealthcarePipelineJob struct {
	pulumi.CustomResourceState

	// Specifies the backfill configuration.
	BackfillPipelineJob HealthcarePipelineJobBackfillPipelineJobPtrOutput `pulumi:"backfillPipelineJob"`
	// Healthcare Dataset under which the Pipeline Job is to run
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// If true, disables writing lineage for the pipeline.
	DisableLineage          pulumi.BoolPtrOutput   `pulumi:"disableLineage"`
	EffectiveLabels         pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	HealthcarePipelineJobId pulumi.StringOutput    `pulumi:"healthcarePipelineJobId"`
	// User-supplied key-value pairs used to organize Pipeline Jobs. Label keys must be between 1 and 63 characters long, have
	// a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given pipeline. An object containing a list
	// of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location where the Pipeline Job is to run
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies mapping configuration.
	MappingPipelineJob HealthcarePipelineJobMappingPipelineJobPtrOutput `pulumi:"mappingPipelineJob"`
	// Specifies the name of the pipeline job. This field is user-assigned.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies reconciliation configuration.
	ReconciliationPipelineJob HealthcarePipelineJobReconciliationPipelineJobPtrOutput `pulumi:"reconciliationPipelineJob"`
	// The fully qualified name of this dataset
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                 `pulumi:"terraformLabels"`
	Timeouts        HealthcarePipelineJobTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewHealthcarePipelineJob registers a new resource with the given unique name, arguments, and options.
func NewHealthcarePipelineJob(ctx *pulumi.Context,
	name string, args *HealthcarePipelineJobArgs, opts ...pulumi.ResourceOption) (*HealthcarePipelineJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource HealthcarePipelineJob
	err = ctx.RegisterPackageResource("google:index/healthcarePipelineJob:HealthcarePipelineJob", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthcarePipelineJob gets an existing HealthcarePipelineJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthcarePipelineJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthcarePipelineJobState, opts ...pulumi.ResourceOption) (*HealthcarePipelineJob, error) {
	var resource HealthcarePipelineJob
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/healthcarePipelineJob:HealthcarePipelineJob", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthcarePipelineJob resources.
type healthcarePipelineJobState struct {
	// Specifies the backfill configuration.
	BackfillPipelineJob *HealthcarePipelineJobBackfillPipelineJob `pulumi:"backfillPipelineJob"`
	// Healthcare Dataset under which the Pipeline Job is to run
	Dataset *string `pulumi:"dataset"`
	// If true, disables writing lineage for the pipeline.
	DisableLineage          *bool             `pulumi:"disableLineage"`
	EffectiveLabels         map[string]string `pulumi:"effectiveLabels"`
	HealthcarePipelineJobId *string           `pulumi:"healthcarePipelineJobId"`
	// User-supplied key-value pairs used to organize Pipeline Jobs. Label keys must be between 1 and 63 characters long, have
	// a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given pipeline. An object containing a list
	// of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location where the Pipeline Job is to run
	Location *string `pulumi:"location"`
	// Specifies mapping configuration.
	MappingPipelineJob *HealthcarePipelineJobMappingPipelineJob `pulumi:"mappingPipelineJob"`
	// Specifies the name of the pipeline job. This field is user-assigned.
	Name *string `pulumi:"name"`
	// Specifies reconciliation configuration.
	ReconciliationPipelineJob *HealthcarePipelineJobReconciliationPipelineJob `pulumi:"reconciliationPipelineJob"`
	// The fully qualified name of this dataset
	SelfLink *string `pulumi:"selfLink"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string              `pulumi:"terraformLabels"`
	Timeouts        *HealthcarePipelineJobTimeouts `pulumi:"timeouts"`
}

type HealthcarePipelineJobState struct {
	// Specifies the backfill configuration.
	BackfillPipelineJob HealthcarePipelineJobBackfillPipelineJobPtrInput
	// Healthcare Dataset under which the Pipeline Job is to run
	Dataset pulumi.StringPtrInput
	// If true, disables writing lineage for the pipeline.
	DisableLineage          pulumi.BoolPtrInput
	EffectiveLabels         pulumi.StringMapInput
	HealthcarePipelineJobId pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize Pipeline Jobs. Label keys must be between 1 and 63 characters long, have
	// a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given pipeline. An object containing a list
	// of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location where the Pipeline Job is to run
	Location pulumi.StringPtrInput
	// Specifies mapping configuration.
	MappingPipelineJob HealthcarePipelineJobMappingPipelineJobPtrInput
	// Specifies the name of the pipeline job. This field is user-assigned.
	Name pulumi.StringPtrInput
	// Specifies reconciliation configuration.
	ReconciliationPipelineJob HealthcarePipelineJobReconciliationPipelineJobPtrInput
	// The fully qualified name of this dataset
	SelfLink pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        HealthcarePipelineJobTimeoutsPtrInput
}

func (HealthcarePipelineJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcarePipelineJobState)(nil)).Elem()
}

type healthcarePipelineJobArgs struct {
	// Specifies the backfill configuration.
	BackfillPipelineJob *HealthcarePipelineJobBackfillPipelineJob `pulumi:"backfillPipelineJob"`
	// Healthcare Dataset under which the Pipeline Job is to run
	Dataset string `pulumi:"dataset"`
	// If true, disables writing lineage for the pipeline.
	DisableLineage          *bool   `pulumi:"disableLineage"`
	HealthcarePipelineJobId *string `pulumi:"healthcarePipelineJobId"`
	// User-supplied key-value pairs used to organize Pipeline Jobs. Label keys must be between 1 and 63 characters long, have
	// a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given pipeline. An object containing a list
	// of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location where the Pipeline Job is to run
	Location string `pulumi:"location"`
	// Specifies mapping configuration.
	MappingPipelineJob *HealthcarePipelineJobMappingPipelineJob `pulumi:"mappingPipelineJob"`
	// Specifies the name of the pipeline job. This field is user-assigned.
	Name *string `pulumi:"name"`
	// Specifies reconciliation configuration.
	ReconciliationPipelineJob *HealthcarePipelineJobReconciliationPipelineJob `pulumi:"reconciliationPipelineJob"`
	Timeouts                  *HealthcarePipelineJobTimeouts                  `pulumi:"timeouts"`
}

// The set of arguments for constructing a HealthcarePipelineJob resource.
type HealthcarePipelineJobArgs struct {
	// Specifies the backfill configuration.
	BackfillPipelineJob HealthcarePipelineJobBackfillPipelineJobPtrInput
	// Healthcare Dataset under which the Pipeline Job is to run
	Dataset pulumi.StringInput
	// If true, disables writing lineage for the pipeline.
	DisableLineage          pulumi.BoolPtrInput
	HealthcarePipelineJobId pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize Pipeline Jobs. Label keys must be between 1 and 63 characters long, have
	// a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given pipeline. An object containing a list
	// of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location where the Pipeline Job is to run
	Location pulumi.StringInput
	// Specifies mapping configuration.
	MappingPipelineJob HealthcarePipelineJobMappingPipelineJobPtrInput
	// Specifies the name of the pipeline job. This field is user-assigned.
	Name pulumi.StringPtrInput
	// Specifies reconciliation configuration.
	ReconciliationPipelineJob HealthcarePipelineJobReconciliationPipelineJobPtrInput
	Timeouts                  HealthcarePipelineJobTimeoutsPtrInput
}

func (HealthcarePipelineJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcarePipelineJobArgs)(nil)).Elem()
}

type HealthcarePipelineJobInput interface {
	pulumi.Input

	ToHealthcarePipelineJobOutput() HealthcarePipelineJobOutput
	ToHealthcarePipelineJobOutputWithContext(ctx context.Context) HealthcarePipelineJobOutput
}

func (*HealthcarePipelineJob) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcarePipelineJob)(nil)).Elem()
}

func (i *HealthcarePipelineJob) ToHealthcarePipelineJobOutput() HealthcarePipelineJobOutput {
	return i.ToHealthcarePipelineJobOutputWithContext(context.Background())
}

func (i *HealthcarePipelineJob) ToHealthcarePipelineJobOutputWithContext(ctx context.Context) HealthcarePipelineJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthcarePipelineJobOutput)
}

type HealthcarePipelineJobOutput struct{ *pulumi.OutputState }

func (HealthcarePipelineJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcarePipelineJob)(nil)).Elem()
}

func (o HealthcarePipelineJobOutput) ToHealthcarePipelineJobOutput() HealthcarePipelineJobOutput {
	return o
}

func (o HealthcarePipelineJobOutput) ToHealthcarePipelineJobOutputWithContext(ctx context.Context) HealthcarePipelineJobOutput {
	return o
}

// Specifies the backfill configuration.
func (o HealthcarePipelineJobOutput) BackfillPipelineJob() HealthcarePipelineJobBackfillPipelineJobPtrOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) HealthcarePipelineJobBackfillPipelineJobPtrOutput {
		return v.BackfillPipelineJob
	}).(HealthcarePipelineJobBackfillPipelineJobPtrOutput)
}

// Healthcare Dataset under which the Pipeline Job is to run
func (o HealthcarePipelineJobOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

// If true, disables writing lineage for the pipeline.
func (o HealthcarePipelineJobOutput) DisableLineage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.BoolPtrOutput { return v.DisableLineage }).(pulumi.BoolPtrOutput)
}

func (o HealthcarePipelineJobOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o HealthcarePipelineJobOutput) HealthcarePipelineJobId() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringOutput { return v.HealthcarePipelineJobId }).(pulumi.StringOutput)
}

// User-supplied key-value pairs used to organize Pipeline Jobs. Label keys must be between 1 and 63 characters long, have
// a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given pipeline. An object containing a list
// of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is
// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
// 'effective_labels' for all of the labels present on the resource.
func (o HealthcarePipelineJobOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Location where the Pipeline Job is to run
func (o HealthcarePipelineJobOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Specifies mapping configuration.
func (o HealthcarePipelineJobOutput) MappingPipelineJob() HealthcarePipelineJobMappingPipelineJobPtrOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) HealthcarePipelineJobMappingPipelineJobPtrOutput {
		return v.MappingPipelineJob
	}).(HealthcarePipelineJobMappingPipelineJobPtrOutput)
}

// Specifies the name of the pipeline job. This field is user-assigned.
func (o HealthcarePipelineJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies reconciliation configuration.
func (o HealthcarePipelineJobOutput) ReconciliationPipelineJob() HealthcarePipelineJobReconciliationPipelineJobPtrOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) HealthcarePipelineJobReconciliationPipelineJobPtrOutput {
		return v.ReconciliationPipelineJob
	}).(HealthcarePipelineJobReconciliationPipelineJobPtrOutput)
}

// The fully qualified name of this dataset
func (o HealthcarePipelineJobOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o HealthcarePipelineJobOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o HealthcarePipelineJobOutput) Timeouts() HealthcarePipelineJobTimeoutsPtrOutput {
	return o.ApplyT(func(v *HealthcarePipelineJob) HealthcarePipelineJobTimeoutsPtrOutput { return v.Timeouts }).(HealthcarePipelineJobTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthcarePipelineJobInput)(nil)).Elem(), &HealthcarePipelineJob{})
	pulumi.RegisterOutputType(HealthcarePipelineJobOutput{})
}
