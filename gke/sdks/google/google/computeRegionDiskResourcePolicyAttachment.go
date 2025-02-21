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

type ComputeRegionDiskResourcePolicyAttachment struct {
	pulumi.CustomResourceState

	ComputeRegionDiskResourcePolicyAttachmentId pulumi.StringOutput `pulumi:"computeRegionDiskResourcePolicyAttachmentId"`
	// The name of the regional disk in which the resource policies are attached to.
	Disk pulumi.StringOutput `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region   pulumi.StringOutput                                        `pulumi:"region"`
	Timeouts ComputeRegionDiskResourcePolicyAttachmentTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeRegionDiskResourcePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewComputeRegionDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, args *ComputeRegionDiskResourcePolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*ComputeRegionDiskResourcePolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disk == nil {
		return nil, errors.New("invalid value for required argument 'Disk'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeRegionDiskResourcePolicyAttachment
	err = ctx.RegisterPackageResource("google:index/computeRegionDiskResourcePolicyAttachment:ComputeRegionDiskResourcePolicyAttachment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRegionDiskResourcePolicyAttachment gets an existing ComputeRegionDiskResourcePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRegionDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRegionDiskResourcePolicyAttachmentState, opts ...pulumi.ResourceOption) (*ComputeRegionDiskResourcePolicyAttachment, error) {
	var resource ComputeRegionDiskResourcePolicyAttachment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeRegionDiskResourcePolicyAttachment:ComputeRegionDiskResourcePolicyAttachment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRegionDiskResourcePolicyAttachment resources.
type computeRegionDiskResourcePolicyAttachmentState struct {
	ComputeRegionDiskResourcePolicyAttachmentId *string `pulumi:"computeRegionDiskResourcePolicyAttachmentId"`
	// The name of the regional disk in which the resource policies are attached to.
	Disk *string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region   *string                                            `pulumi:"region"`
	Timeouts *ComputeRegionDiskResourcePolicyAttachmentTimeouts `pulumi:"timeouts"`
}

type ComputeRegionDiskResourcePolicyAttachmentState struct {
	ComputeRegionDiskResourcePolicyAttachmentId pulumi.StringPtrInput
	// The name of the regional disk in which the resource policies are attached to.
	Disk pulumi.StringPtrInput
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region   pulumi.StringPtrInput
	Timeouts ComputeRegionDiskResourcePolicyAttachmentTimeoutsPtrInput
}

func (ComputeRegionDiskResourcePolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionDiskResourcePolicyAttachmentState)(nil)).Elem()
}

type computeRegionDiskResourcePolicyAttachmentArgs struct {
	ComputeRegionDiskResourcePolicyAttachmentId *string `pulumi:"computeRegionDiskResourcePolicyAttachmentId"`
	// The name of the regional disk in which the resource policies are attached to.
	Disk string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region   *string                                            `pulumi:"region"`
	Timeouts *ComputeRegionDiskResourcePolicyAttachmentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeRegionDiskResourcePolicyAttachment resource.
type ComputeRegionDiskResourcePolicyAttachmentArgs struct {
	ComputeRegionDiskResourcePolicyAttachmentId pulumi.StringPtrInput
	// The name of the regional disk in which the resource policies are attached to.
	Disk pulumi.StringInput
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region   pulumi.StringPtrInput
	Timeouts ComputeRegionDiskResourcePolicyAttachmentTimeoutsPtrInput
}

func (ComputeRegionDiskResourcePolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionDiskResourcePolicyAttachmentArgs)(nil)).Elem()
}

type ComputeRegionDiskResourcePolicyAttachmentInput interface {
	pulumi.Input

	ToComputeRegionDiskResourcePolicyAttachmentOutput() ComputeRegionDiskResourcePolicyAttachmentOutput
	ToComputeRegionDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) ComputeRegionDiskResourcePolicyAttachmentOutput
}

func (*ComputeRegionDiskResourcePolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (i *ComputeRegionDiskResourcePolicyAttachment) ToComputeRegionDiskResourcePolicyAttachmentOutput() ComputeRegionDiskResourcePolicyAttachmentOutput {
	return i.ToComputeRegionDiskResourcePolicyAttachmentOutputWithContext(context.Background())
}

func (i *ComputeRegionDiskResourcePolicyAttachment) ToComputeRegionDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) ComputeRegionDiskResourcePolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRegionDiskResourcePolicyAttachmentOutput)
}

type ComputeRegionDiskResourcePolicyAttachmentOutput struct{ *pulumi.OutputState }

func (ComputeRegionDiskResourcePolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionDiskResourcePolicyAttachment)(nil)).Elem()
}

func (o ComputeRegionDiskResourcePolicyAttachmentOutput) ToComputeRegionDiskResourcePolicyAttachmentOutput() ComputeRegionDiskResourcePolicyAttachmentOutput {
	return o
}

func (o ComputeRegionDiskResourcePolicyAttachmentOutput) ToComputeRegionDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) ComputeRegionDiskResourcePolicyAttachmentOutput {
	return o
}

func (o ComputeRegionDiskResourcePolicyAttachmentOutput) ComputeRegionDiskResourcePolicyAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskResourcePolicyAttachment) pulumi.StringOutput {
		return v.ComputeRegionDiskResourcePolicyAttachmentId
	}).(pulumi.StringOutput)
}

// The name of the regional disk in which the resource policies are attached to.
func (o ComputeRegionDiskResourcePolicyAttachmentOutput) Disk() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Disk }).(pulumi.StringOutput)
}

// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
func (o ComputeRegionDiskResourcePolicyAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeRegionDiskResourcePolicyAttachmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A reference to the region where the disk resides.
func (o ComputeRegionDiskResourcePolicyAttachmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeRegionDiskResourcePolicyAttachmentOutput) Timeouts() ComputeRegionDiskResourcePolicyAttachmentTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeRegionDiskResourcePolicyAttachment) ComputeRegionDiskResourcePolicyAttachmentTimeoutsPtrOutput {
		return v.Timeouts
	}).(ComputeRegionDiskResourcePolicyAttachmentTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRegionDiskResourcePolicyAttachmentInput)(nil)).Elem(), &ComputeRegionDiskResourcePolicyAttachment{})
	pulumi.RegisterOutputType(ComputeRegionDiskResourcePolicyAttachmentOutput{})
}
