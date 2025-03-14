// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package castai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Commitments struct {
	pulumi.CustomResourceState

	// List of Azure reservations.
	AzureReservations CommitmentsAzureReservationArrayOutput `pulumi:"azureReservations"`
	// CSV file containing reservations exported from Azure.
	AzureReservationsCsv pulumi.StringPtrOutput `pulumi:"azureReservationsCsv"`
	// List of commitment configurations.
	CommitmentConfigs CommitmentsCommitmentConfigArrayOutput `pulumi:"commitmentConfigs"`
	CommitmentsId     pulumi.StringOutput                    `pulumi:"commitmentsId"`
	// List of GCP CUDs.
	GcpCuds CommitmentsGcpCudArrayOutput `pulumi:"gcpCuds"`
	// JSON file containing CUDs exported from GCP.
	GcpCudsJson pulumi.StringPtrOutput       `pulumi:"gcpCudsJson"`
	Timeouts    CommitmentsTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewCommitments registers a new resource with the given unique name, arguments, and options.
func NewCommitments(ctx *pulumi.Context,
	name string, args *CommitmentsArgs, opts ...pulumi.ResourceOption) (*Commitments, error) {
	if args == nil {
		args = &CommitmentsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource Commitments
	err = ctx.RegisterPackageResource("castai:index/commitments:Commitments", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommitments gets an existing Commitments resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommitments(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommitmentsState, opts ...pulumi.ResourceOption) (*Commitments, error) {
	var resource Commitments
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("castai:index/commitments:Commitments", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Commitments resources.
type commitmentsState struct {
	// List of Azure reservations.
	AzureReservations []CommitmentsAzureReservation `pulumi:"azureReservations"`
	// CSV file containing reservations exported from Azure.
	AzureReservationsCsv *string `pulumi:"azureReservationsCsv"`
	// List of commitment configurations.
	CommitmentConfigs []CommitmentsCommitmentConfig `pulumi:"commitmentConfigs"`
	CommitmentsId     *string                       `pulumi:"commitmentsId"`
	// List of GCP CUDs.
	GcpCuds []CommitmentsGcpCud `pulumi:"gcpCuds"`
	// JSON file containing CUDs exported from GCP.
	GcpCudsJson *string              `pulumi:"gcpCudsJson"`
	Timeouts    *CommitmentsTimeouts `pulumi:"timeouts"`
}

type CommitmentsState struct {
	// List of Azure reservations.
	AzureReservations CommitmentsAzureReservationArrayInput
	// CSV file containing reservations exported from Azure.
	AzureReservationsCsv pulumi.StringPtrInput
	// List of commitment configurations.
	CommitmentConfigs CommitmentsCommitmentConfigArrayInput
	CommitmentsId     pulumi.StringPtrInput
	// List of GCP CUDs.
	GcpCuds CommitmentsGcpCudArrayInput
	// JSON file containing CUDs exported from GCP.
	GcpCudsJson pulumi.StringPtrInput
	Timeouts    CommitmentsTimeoutsPtrInput
}

func (CommitmentsState) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentsState)(nil)).Elem()
}

type commitmentsArgs struct {
	// CSV file containing reservations exported from Azure.
	AzureReservationsCsv *string `pulumi:"azureReservationsCsv"`
	// List of commitment configurations.
	CommitmentConfigs []CommitmentsCommitmentConfig `pulumi:"commitmentConfigs"`
	CommitmentsId     *string                       `pulumi:"commitmentsId"`
	// JSON file containing CUDs exported from GCP.
	GcpCudsJson *string              `pulumi:"gcpCudsJson"`
	Timeouts    *CommitmentsTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a Commitments resource.
type CommitmentsArgs struct {
	// CSV file containing reservations exported from Azure.
	AzureReservationsCsv pulumi.StringPtrInput
	// List of commitment configurations.
	CommitmentConfigs CommitmentsCommitmentConfigArrayInput
	CommitmentsId     pulumi.StringPtrInput
	// JSON file containing CUDs exported from GCP.
	GcpCudsJson pulumi.StringPtrInput
	Timeouts    CommitmentsTimeoutsPtrInput
}

func (CommitmentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentsArgs)(nil)).Elem()
}

type CommitmentsInput interface {
	pulumi.Input

	ToCommitmentsOutput() CommitmentsOutput
	ToCommitmentsOutputWithContext(ctx context.Context) CommitmentsOutput
}

func (*Commitments) ElementType() reflect.Type {
	return reflect.TypeOf((**Commitments)(nil)).Elem()
}

func (i *Commitments) ToCommitmentsOutput() CommitmentsOutput {
	return i.ToCommitmentsOutputWithContext(context.Background())
}

func (i *Commitments) ToCommitmentsOutputWithContext(ctx context.Context) CommitmentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentsOutput)
}

type CommitmentsOutput struct{ *pulumi.OutputState }

func (CommitmentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Commitments)(nil)).Elem()
}

func (o CommitmentsOutput) ToCommitmentsOutput() CommitmentsOutput {
	return o
}

func (o CommitmentsOutput) ToCommitmentsOutputWithContext(ctx context.Context) CommitmentsOutput {
	return o
}

// List of Azure reservations.
func (o CommitmentsOutput) AzureReservations() CommitmentsAzureReservationArrayOutput {
	return o.ApplyT(func(v *Commitments) CommitmentsAzureReservationArrayOutput { return v.AzureReservations }).(CommitmentsAzureReservationArrayOutput)
}

// CSV file containing reservations exported from Azure.
func (o CommitmentsOutput) AzureReservationsCsv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Commitments) pulumi.StringPtrOutput { return v.AzureReservationsCsv }).(pulumi.StringPtrOutput)
}

// List of commitment configurations.
func (o CommitmentsOutput) CommitmentConfigs() CommitmentsCommitmentConfigArrayOutput {
	return o.ApplyT(func(v *Commitments) CommitmentsCommitmentConfigArrayOutput { return v.CommitmentConfigs }).(CommitmentsCommitmentConfigArrayOutput)
}

func (o CommitmentsOutput) CommitmentsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Commitments) pulumi.StringOutput { return v.CommitmentsId }).(pulumi.StringOutput)
}

// List of GCP CUDs.
func (o CommitmentsOutput) GcpCuds() CommitmentsGcpCudArrayOutput {
	return o.ApplyT(func(v *Commitments) CommitmentsGcpCudArrayOutput { return v.GcpCuds }).(CommitmentsGcpCudArrayOutput)
}

// JSON file containing CUDs exported from GCP.
func (o CommitmentsOutput) GcpCudsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Commitments) pulumi.StringPtrOutput { return v.GcpCudsJson }).(pulumi.StringPtrOutput)
}

func (o CommitmentsOutput) Timeouts() CommitmentsTimeoutsPtrOutput {
	return o.ApplyT(func(v *Commitments) CommitmentsTimeoutsPtrOutput { return v.Timeouts }).(CommitmentsTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommitmentsInput)(nil)).Elem(), &Commitments{})
	pulumi.RegisterOutputType(CommitmentsOutput{})
}
