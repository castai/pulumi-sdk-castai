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

type CloudRunDomainMapping struct {
	pulumi.CustomResourceState

	CloudRunDomainMappingId pulumi.StringOutput `pulumi:"cloudRunDomainMappingId"`
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata associated with this DomainMapping.
	Metadata CloudRunDomainMappingMetadataPtrOutput `pulumi:"metadata"`
	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The spec for this DomainMapping.
	Spec CloudRunDomainMappingSpecOutput `pulumi:"spec"`
	// The current status of the DomainMapping.
	Statuses CloudRunDomainMappingStatusArrayOutput `pulumi:"statuses"`
	Timeouts CloudRunDomainMappingTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewCloudRunDomainMapping registers a new resource with the given unique name, arguments, and options.
func NewCloudRunDomainMapping(ctx *pulumi.Context,
	name string, args *CloudRunDomainMappingArgs, opts ...pulumi.ResourceOption) (*CloudRunDomainMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudRunDomainMapping
	err = ctx.RegisterPackageResource("google-beta:index/cloudRunDomainMapping:CloudRunDomainMapping", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudRunDomainMapping gets an existing CloudRunDomainMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudRunDomainMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudRunDomainMappingState, opts ...pulumi.ResourceOption) (*CloudRunDomainMapping, error) {
	var resource CloudRunDomainMapping
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/cloudRunDomainMapping:CloudRunDomainMapping", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudRunDomainMapping resources.
type cloudRunDomainMappingState struct {
	CloudRunDomainMappingId *string `pulumi:"cloudRunDomainMappingId"`
	// The location of the cloud run instance. eg us-central1
	Location *string `pulumi:"location"`
	// Metadata associated with this DomainMapping.
	Metadata *CloudRunDomainMappingMetadata `pulumi:"metadata"`
	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The spec for this DomainMapping.
	Spec *CloudRunDomainMappingSpec `pulumi:"spec"`
	// The current status of the DomainMapping.
	Statuses []CloudRunDomainMappingStatus  `pulumi:"statuses"`
	Timeouts *CloudRunDomainMappingTimeouts `pulumi:"timeouts"`
}

type CloudRunDomainMappingState struct {
	CloudRunDomainMappingId pulumi.StringPtrInput
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringPtrInput
	// Metadata associated with this DomainMapping.
	Metadata CloudRunDomainMappingMetadataPtrInput
	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The spec for this DomainMapping.
	Spec CloudRunDomainMappingSpecPtrInput
	// The current status of the DomainMapping.
	Statuses CloudRunDomainMappingStatusArrayInput
	Timeouts CloudRunDomainMappingTimeoutsPtrInput
}

func (CloudRunDomainMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRunDomainMappingState)(nil)).Elem()
}

type cloudRunDomainMappingArgs struct {
	CloudRunDomainMappingId *string `pulumi:"cloudRunDomainMappingId"`
	// The location of the cloud run instance. eg us-central1
	Location string `pulumi:"location"`
	// Metadata associated with this DomainMapping.
	Metadata *CloudRunDomainMappingMetadata `pulumi:"metadata"`
	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The spec for this DomainMapping.
	Spec     CloudRunDomainMappingSpec      `pulumi:"spec"`
	Timeouts *CloudRunDomainMappingTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a CloudRunDomainMapping resource.
type CloudRunDomainMappingArgs struct {
	CloudRunDomainMappingId pulumi.StringPtrInput
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringInput
	// Metadata associated with this DomainMapping.
	Metadata CloudRunDomainMappingMetadataPtrInput
	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The spec for this DomainMapping.
	Spec     CloudRunDomainMappingSpecInput
	Timeouts CloudRunDomainMappingTimeoutsPtrInput
}

func (CloudRunDomainMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRunDomainMappingArgs)(nil)).Elem()
}

type CloudRunDomainMappingInput interface {
	pulumi.Input

	ToCloudRunDomainMappingOutput() CloudRunDomainMappingOutput
	ToCloudRunDomainMappingOutputWithContext(ctx context.Context) CloudRunDomainMappingOutput
}

func (*CloudRunDomainMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRunDomainMapping)(nil)).Elem()
}

func (i *CloudRunDomainMapping) ToCloudRunDomainMappingOutput() CloudRunDomainMappingOutput {
	return i.ToCloudRunDomainMappingOutputWithContext(context.Background())
}

func (i *CloudRunDomainMapping) ToCloudRunDomainMappingOutputWithContext(ctx context.Context) CloudRunDomainMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudRunDomainMappingOutput)
}

type CloudRunDomainMappingOutput struct{ *pulumi.OutputState }

func (CloudRunDomainMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRunDomainMapping)(nil)).Elem()
}

func (o CloudRunDomainMappingOutput) ToCloudRunDomainMappingOutput() CloudRunDomainMappingOutput {
	return o
}

func (o CloudRunDomainMappingOutput) ToCloudRunDomainMappingOutputWithContext(ctx context.Context) CloudRunDomainMappingOutput {
	return o
}

func (o CloudRunDomainMappingOutput) CloudRunDomainMappingId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) pulumi.StringOutput { return v.CloudRunDomainMappingId }).(pulumi.StringOutput)
}

// The location of the cloud run instance. eg us-central1
func (o CloudRunDomainMappingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Metadata associated with this DomainMapping.
func (o CloudRunDomainMappingOutput) Metadata() CloudRunDomainMappingMetadataPtrOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) CloudRunDomainMappingMetadataPtrOutput { return v.Metadata }).(CloudRunDomainMappingMetadataPtrOutput)
}

// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
func (o CloudRunDomainMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudRunDomainMappingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The spec for this DomainMapping.
func (o CloudRunDomainMappingOutput) Spec() CloudRunDomainMappingSpecOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) CloudRunDomainMappingSpecOutput { return v.Spec }).(CloudRunDomainMappingSpecOutput)
}

// The current status of the DomainMapping.
func (o CloudRunDomainMappingOutput) Statuses() CloudRunDomainMappingStatusArrayOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) CloudRunDomainMappingStatusArrayOutput { return v.Statuses }).(CloudRunDomainMappingStatusArrayOutput)
}

func (o CloudRunDomainMappingOutput) Timeouts() CloudRunDomainMappingTimeoutsPtrOutput {
	return o.ApplyT(func(v *CloudRunDomainMapping) CloudRunDomainMappingTimeoutsPtrOutput { return v.Timeouts }).(CloudRunDomainMappingTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudRunDomainMappingInput)(nil)).Elem(), &CloudRunDomainMapping{})
	pulumi.RegisterOutputType(CloudRunDomainMappingOutput{})
}
