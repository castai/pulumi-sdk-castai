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

type HealthcareDataset struct {
	pulumi.CustomResourceState

	// A nested object resource.
	EncryptionSpec      HealthcareDatasetEncryptionSpecPtrOutput `pulumi:"encryptionSpec"`
	HealthcareDatasetId pulumi.StringOutput                      `pulumi:"healthcareDatasetId"`
	// The location for the Dataset.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the Dataset.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The fully qualified name of this dataset
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or
	// empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7 messages) where no explicit
	// timezone is specified.
	TimeZone pulumi.StringOutput                `pulumi:"timeZone"`
	Timeouts HealthcareDatasetTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewHealthcareDataset registers a new resource with the given unique name, arguments, and options.
func NewHealthcareDataset(ctx *pulumi.Context,
	name string, args *HealthcareDatasetArgs, opts ...pulumi.ResourceOption) (*HealthcareDataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource HealthcareDataset
	err = ctx.RegisterPackageResource("google:index/healthcareDataset:HealthcareDataset", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthcareDataset gets an existing HealthcareDataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthcareDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthcareDatasetState, opts ...pulumi.ResourceOption) (*HealthcareDataset, error) {
	var resource HealthcareDataset
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/healthcareDataset:HealthcareDataset", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthcareDataset resources.
type healthcareDatasetState struct {
	// A nested object resource.
	EncryptionSpec      *HealthcareDatasetEncryptionSpec `pulumi:"encryptionSpec"`
	HealthcareDatasetId *string                          `pulumi:"healthcareDatasetId"`
	// The location for the Dataset.
	Location *string `pulumi:"location"`
	// The resource name for the Dataset.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The fully qualified name of this dataset
	SelfLink *string `pulumi:"selfLink"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or
	// empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7 messages) where no explicit
	// timezone is specified.
	TimeZone *string                    `pulumi:"timeZone"`
	Timeouts *HealthcareDatasetTimeouts `pulumi:"timeouts"`
}

type HealthcareDatasetState struct {
	// A nested object resource.
	EncryptionSpec      HealthcareDatasetEncryptionSpecPtrInput
	HealthcareDatasetId pulumi.StringPtrInput
	// The location for the Dataset.
	Location pulumi.StringPtrInput
	// The resource name for the Dataset.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The fully qualified name of this dataset
	SelfLink pulumi.StringPtrInput
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or
	// empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7 messages) where no explicit
	// timezone is specified.
	TimeZone pulumi.StringPtrInput
	Timeouts HealthcareDatasetTimeoutsPtrInput
}

func (HealthcareDatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcareDatasetState)(nil)).Elem()
}

type healthcareDatasetArgs struct {
	// A nested object resource.
	EncryptionSpec      *HealthcareDatasetEncryptionSpec `pulumi:"encryptionSpec"`
	HealthcareDatasetId *string                          `pulumi:"healthcareDatasetId"`
	// The location for the Dataset.
	Location string `pulumi:"location"`
	// The resource name for the Dataset.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or
	// empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7 messages) where no explicit
	// timezone is specified.
	TimeZone *string                    `pulumi:"timeZone"`
	Timeouts *HealthcareDatasetTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a HealthcareDataset resource.
type HealthcareDatasetArgs struct {
	// A nested object resource.
	EncryptionSpec      HealthcareDatasetEncryptionSpecPtrInput
	HealthcareDatasetId pulumi.StringPtrInput
	// The location for the Dataset.
	Location pulumi.StringInput
	// The resource name for the Dataset.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or
	// empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7 messages) where no explicit
	// timezone is specified.
	TimeZone pulumi.StringPtrInput
	Timeouts HealthcareDatasetTimeoutsPtrInput
}

func (HealthcareDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthcareDatasetArgs)(nil)).Elem()
}

type HealthcareDatasetInput interface {
	pulumi.Input

	ToHealthcareDatasetOutput() HealthcareDatasetOutput
	ToHealthcareDatasetOutputWithContext(ctx context.Context) HealthcareDatasetOutput
}

func (*HealthcareDataset) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcareDataset)(nil)).Elem()
}

func (i *HealthcareDataset) ToHealthcareDatasetOutput() HealthcareDatasetOutput {
	return i.ToHealthcareDatasetOutputWithContext(context.Background())
}

func (i *HealthcareDataset) ToHealthcareDatasetOutputWithContext(ctx context.Context) HealthcareDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthcareDatasetOutput)
}

type HealthcareDatasetOutput struct{ *pulumi.OutputState }

func (HealthcareDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthcareDataset)(nil)).Elem()
}

func (o HealthcareDatasetOutput) ToHealthcareDatasetOutput() HealthcareDatasetOutput {
	return o
}

func (o HealthcareDatasetOutput) ToHealthcareDatasetOutputWithContext(ctx context.Context) HealthcareDatasetOutput {
	return o
}

// A nested object resource.
func (o HealthcareDatasetOutput) EncryptionSpec() HealthcareDatasetEncryptionSpecPtrOutput {
	return o.ApplyT(func(v *HealthcareDataset) HealthcareDatasetEncryptionSpecPtrOutput { return v.EncryptionSpec }).(HealthcareDatasetEncryptionSpecPtrOutput)
}

func (o HealthcareDatasetOutput) HealthcareDatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareDataset) pulumi.StringOutput { return v.HealthcareDatasetId }).(pulumi.StringOutput)
}

// The location for the Dataset.
func (o HealthcareDatasetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareDataset) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for the Dataset.
func (o HealthcareDatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareDataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HealthcareDatasetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareDataset) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The fully qualified name of this dataset
func (o HealthcareDatasetOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareDataset) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or
// empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7 messages) where no explicit
// timezone is specified.
func (o HealthcareDatasetOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthcareDataset) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

func (o HealthcareDatasetOutput) Timeouts() HealthcareDatasetTimeoutsPtrOutput {
	return o.ApplyT(func(v *HealthcareDataset) HealthcareDatasetTimeoutsPtrOutput { return v.Timeouts }).(HealthcareDatasetTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthcareDatasetInput)(nil)).Elem(), &HealthcareDataset{})
	pulumi.RegisterOutputType(HealthcareDatasetOutput{})
}
