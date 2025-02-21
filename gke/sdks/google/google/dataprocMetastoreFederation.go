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

type DataprocMetastoreFederation struct {
	pulumi.CustomResourceState

	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time.
	// The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time
	// and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a
	// BackendMetastore with a higher number.
	BackendMetastores             DataprocMetastoreFederationBackendMetastoreArrayOutput `pulumi:"backendMetastores"`
	DataprocMetastoreFederationId pulumi.StringOutput                                    `pulumi:"dataprocMetastoreFederationId"`
	EffectiveLabels               pulumi.StringMapOutput                                 `pulumi:"effectiveLabels"`
	// The URI of the endpoint used to access the metastore federation.
	EndpointUri pulumi.StringOutput `pulumi:"endpointUri"`
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	FederationId pulumi.StringOutput `pulumi:"federationId"`
	// User-defined labels for the metastore federation. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where the metastore federation should reside.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The relative resource name of the metastore federation.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The current state of the metastore federation.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of the metastore federation, if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                       `pulumi:"terraformLabels"`
	Timeouts        DataprocMetastoreFederationTimeoutsPtrOutput `pulumi:"timeouts"`
	// The globally unique resource identifier of the metastore federation.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the
	// federation version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewDataprocMetastoreFederation registers a new resource with the given unique name, arguments, and options.
func NewDataprocMetastoreFederation(ctx *pulumi.Context,
	name string, args *DataprocMetastoreFederationArgs, opts ...pulumi.ResourceOption) (*DataprocMetastoreFederation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendMetastores == nil {
		return nil, errors.New("invalid value for required argument 'BackendMetastores'")
	}
	if args.FederationId == nil {
		return nil, errors.New("invalid value for required argument 'FederationId'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataprocMetastoreFederation
	err = ctx.RegisterPackageResource("google:index/dataprocMetastoreFederation:DataprocMetastoreFederation", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocMetastoreFederation gets an existing DataprocMetastoreFederation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocMetastoreFederation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocMetastoreFederationState, opts ...pulumi.ResourceOption) (*DataprocMetastoreFederation, error) {
	var resource DataprocMetastoreFederation
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataprocMetastoreFederation:DataprocMetastoreFederation", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocMetastoreFederation resources.
type dataprocMetastoreFederationState struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time.
	// The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time
	// and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a
	// BackendMetastore with a higher number.
	BackendMetastores             []DataprocMetastoreFederationBackendMetastore `pulumi:"backendMetastores"`
	DataprocMetastoreFederationId *string                                       `pulumi:"dataprocMetastoreFederationId"`
	EffectiveLabels               map[string]string                             `pulumi:"effectiveLabels"`
	// The URI of the endpoint used to access the metastore federation.
	EndpointUri *string `pulumi:"endpointUri"`
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	FederationId *string `pulumi:"federationId"`
	// User-defined labels for the metastore federation. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the metastore federation should reside.
	Location *string `pulumi:"location"`
	// The relative resource name of the metastore federation.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The current state of the metastore federation.
	State *string `pulumi:"state"`
	// Additional information about the current state of the metastore federation, if available.
	StateMessage *string `pulumi:"stateMessage"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                    `pulumi:"terraformLabels"`
	Timeouts        *DataprocMetastoreFederationTimeouts `pulumi:"timeouts"`
	// The globally unique resource identifier of the metastore federation.
	Uid *string `pulumi:"uid"`
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the
	// federation version.
	Version *string `pulumi:"version"`
}

type DataprocMetastoreFederationState struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time.
	// The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time
	// and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a
	// BackendMetastore with a higher number.
	BackendMetastores             DataprocMetastoreFederationBackendMetastoreArrayInput
	DataprocMetastoreFederationId pulumi.StringPtrInput
	EffectiveLabels               pulumi.StringMapInput
	// The URI of the endpoint used to access the metastore federation.
	EndpointUri pulumi.StringPtrInput
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	FederationId pulumi.StringPtrInput
	// User-defined labels for the metastore federation. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where the metastore federation should reside.
	Location pulumi.StringPtrInput
	// The relative resource name of the metastore federation.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The current state of the metastore federation.
	State pulumi.StringPtrInput
	// Additional information about the current state of the metastore federation, if available.
	StateMessage pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DataprocMetastoreFederationTimeoutsPtrInput
	// The globally unique resource identifier of the metastore federation.
	Uid pulumi.StringPtrInput
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the
	// federation version.
	Version pulumi.StringPtrInput
}

func (DataprocMetastoreFederationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocMetastoreFederationState)(nil)).Elem()
}

type dataprocMetastoreFederationArgs struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time.
	// The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time
	// and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a
	// BackendMetastore with a higher number.
	BackendMetastores             []DataprocMetastoreFederationBackendMetastore `pulumi:"backendMetastores"`
	DataprocMetastoreFederationId *string                                       `pulumi:"dataprocMetastoreFederationId"`
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	FederationId string `pulumi:"federationId"`
	// User-defined labels for the metastore federation. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the metastore federation should reside.
	Location *string                              `pulumi:"location"`
	Project  *string                              `pulumi:"project"`
	Timeouts *DataprocMetastoreFederationTimeouts `pulumi:"timeouts"`
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the
	// federation version.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a DataprocMetastoreFederation resource.
type DataprocMetastoreFederationArgs struct {
	// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time.
	// The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time
	// and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a
	// BackendMetastore with a higher number.
	BackendMetastores             DataprocMetastoreFederationBackendMetastoreArrayInput
	DataprocMetastoreFederationId pulumi.StringPtrInput
	// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	FederationId pulumi.StringInput
	// User-defined labels for the metastore federation. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where the metastore federation should reside.
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts DataprocMetastoreFederationTimeoutsPtrInput
	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the
	// federation version.
	Version pulumi.StringInput
}

func (DataprocMetastoreFederationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocMetastoreFederationArgs)(nil)).Elem()
}

type DataprocMetastoreFederationInput interface {
	pulumi.Input

	ToDataprocMetastoreFederationOutput() DataprocMetastoreFederationOutput
	ToDataprocMetastoreFederationOutputWithContext(ctx context.Context) DataprocMetastoreFederationOutput
}

func (*DataprocMetastoreFederation) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocMetastoreFederation)(nil)).Elem()
}

func (i *DataprocMetastoreFederation) ToDataprocMetastoreFederationOutput() DataprocMetastoreFederationOutput {
	return i.ToDataprocMetastoreFederationOutputWithContext(context.Background())
}

func (i *DataprocMetastoreFederation) ToDataprocMetastoreFederationOutputWithContext(ctx context.Context) DataprocMetastoreFederationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocMetastoreFederationOutput)
}

type DataprocMetastoreFederationOutput struct{ *pulumi.OutputState }

func (DataprocMetastoreFederationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocMetastoreFederation)(nil)).Elem()
}

func (o DataprocMetastoreFederationOutput) ToDataprocMetastoreFederationOutput() DataprocMetastoreFederationOutput {
	return o
}

func (o DataprocMetastoreFederationOutput) ToDataprocMetastoreFederationOutputWithContext(ctx context.Context) DataprocMetastoreFederationOutput {
	return o
}

// A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time.
// The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time
// and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a
// BackendMetastore with a higher number.
func (o DataprocMetastoreFederationOutput) BackendMetastores() DataprocMetastoreFederationBackendMetastoreArrayOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) DataprocMetastoreFederationBackendMetastoreArrayOutput {
		return v.BackendMetastores
	}).(DataprocMetastoreFederationBackendMetastoreArrayOutput)
}

func (o DataprocMetastoreFederationOutput) DataprocMetastoreFederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.DataprocMetastoreFederationId }).(pulumi.StringOutput)
}

func (o DataprocMetastoreFederationOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The URI of the endpoint used to access the metastore federation.
func (o DataprocMetastoreFederationOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.EndpointUri }).(pulumi.StringOutput)
}

// The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
func (o DataprocMetastoreFederationOutput) FederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.FederationId }).(pulumi.StringOutput)
}

// User-defined labels for the metastore federation. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o DataprocMetastoreFederationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where the metastore federation should reside.
func (o DataprocMetastoreFederationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The relative resource name of the metastore federation.
func (o DataprocMetastoreFederationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataprocMetastoreFederationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The current state of the metastore federation.
func (o DataprocMetastoreFederationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current state of the metastore federation, if available.
func (o DataprocMetastoreFederationOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.StateMessage }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataprocMetastoreFederationOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DataprocMetastoreFederationOutput) Timeouts() DataprocMetastoreFederationTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) DataprocMetastoreFederationTimeoutsPtrOutput { return v.Timeouts }).(DataprocMetastoreFederationTimeoutsPtrOutput)
}

// The globally unique resource identifier of the metastore federation.
func (o DataprocMetastoreFederationOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the
// federation version.
func (o DataprocMetastoreFederationOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreFederation) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocMetastoreFederationInput)(nil)).Elem(), &DataprocMetastoreFederation{})
	pulumi.RegisterOutputType(DataprocMetastoreFederationOutput{})
}
