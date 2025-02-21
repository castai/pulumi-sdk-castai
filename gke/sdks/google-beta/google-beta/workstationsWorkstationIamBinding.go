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

type WorkstationsWorkstationIamBinding struct {
	pulumi.CustomResourceState

	Condition                           WorkstationsWorkstationIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                                pulumi.StringOutput                                 `pulumi:"etag"`
	Location                            pulumi.StringOutput                                 `pulumi:"location"`
	Members                             pulumi.StringArrayOutput                            `pulumi:"members"`
	Project                             pulumi.StringOutput                                 `pulumi:"project"`
	Role                                pulumi.StringOutput                                 `pulumi:"role"`
	WorkstationClusterId                pulumi.StringOutput                                 `pulumi:"workstationClusterId"`
	WorkstationConfigId                 pulumi.StringOutput                                 `pulumi:"workstationConfigId"`
	WorkstationId                       pulumi.StringOutput                                 `pulumi:"workstationId"`
	WorkstationsWorkstationIamBindingId pulumi.StringOutput                                 `pulumi:"workstationsWorkstationIamBindingId"`
}

// NewWorkstationsWorkstationIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWorkstationsWorkstationIamBinding(ctx *pulumi.Context,
	name string, args *WorkstationsWorkstationIamBindingArgs, opts ...pulumi.ResourceOption) (*WorkstationsWorkstationIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WorkstationClusterId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationClusterId'")
	}
	if args.WorkstationConfigId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationConfigId'")
	}
	if args.WorkstationId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource WorkstationsWorkstationIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/workstationsWorkstationIamBinding:WorkstationsWorkstationIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstationsWorkstationIamBinding gets an existing WorkstationsWorkstationIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstationsWorkstationIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationsWorkstationIamBindingState, opts ...pulumi.ResourceOption) (*WorkstationsWorkstationIamBinding, error) {
	var resource WorkstationsWorkstationIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/workstationsWorkstationIamBinding:WorkstationsWorkstationIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkstationsWorkstationIamBinding resources.
type workstationsWorkstationIamBindingState struct {
	Condition                           *WorkstationsWorkstationIamBindingCondition `pulumi:"condition"`
	Etag                                *string                                     `pulumi:"etag"`
	Location                            *string                                     `pulumi:"location"`
	Members                             []string                                    `pulumi:"members"`
	Project                             *string                                     `pulumi:"project"`
	Role                                *string                                     `pulumi:"role"`
	WorkstationClusterId                *string                                     `pulumi:"workstationClusterId"`
	WorkstationConfigId                 *string                                     `pulumi:"workstationConfigId"`
	WorkstationId                       *string                                     `pulumi:"workstationId"`
	WorkstationsWorkstationIamBindingId *string                                     `pulumi:"workstationsWorkstationIamBindingId"`
}

type WorkstationsWorkstationIamBindingState struct {
	Condition                           WorkstationsWorkstationIamBindingConditionPtrInput
	Etag                                pulumi.StringPtrInput
	Location                            pulumi.StringPtrInput
	Members                             pulumi.StringArrayInput
	Project                             pulumi.StringPtrInput
	Role                                pulumi.StringPtrInput
	WorkstationClusterId                pulumi.StringPtrInput
	WorkstationConfigId                 pulumi.StringPtrInput
	WorkstationId                       pulumi.StringPtrInput
	WorkstationsWorkstationIamBindingId pulumi.StringPtrInput
}

func (WorkstationsWorkstationIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationsWorkstationIamBindingState)(nil)).Elem()
}

type workstationsWorkstationIamBindingArgs struct {
	Condition                           *WorkstationsWorkstationIamBindingCondition `pulumi:"condition"`
	Location                            *string                                     `pulumi:"location"`
	Members                             []string                                    `pulumi:"members"`
	Project                             *string                                     `pulumi:"project"`
	Role                                string                                      `pulumi:"role"`
	WorkstationClusterId                string                                      `pulumi:"workstationClusterId"`
	WorkstationConfigId                 string                                      `pulumi:"workstationConfigId"`
	WorkstationId                       string                                      `pulumi:"workstationId"`
	WorkstationsWorkstationIamBindingId *string                                     `pulumi:"workstationsWorkstationIamBindingId"`
}

// The set of arguments for constructing a WorkstationsWorkstationIamBinding resource.
type WorkstationsWorkstationIamBindingArgs struct {
	Condition                           WorkstationsWorkstationIamBindingConditionPtrInput
	Location                            pulumi.StringPtrInput
	Members                             pulumi.StringArrayInput
	Project                             pulumi.StringPtrInput
	Role                                pulumi.StringInput
	WorkstationClusterId                pulumi.StringInput
	WorkstationConfigId                 pulumi.StringInput
	WorkstationId                       pulumi.StringInput
	WorkstationsWorkstationIamBindingId pulumi.StringPtrInput
}

func (WorkstationsWorkstationIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationsWorkstationIamBindingArgs)(nil)).Elem()
}

type WorkstationsWorkstationIamBindingInput interface {
	pulumi.Input

	ToWorkstationsWorkstationIamBindingOutput() WorkstationsWorkstationIamBindingOutput
	ToWorkstationsWorkstationIamBindingOutputWithContext(ctx context.Context) WorkstationsWorkstationIamBindingOutput
}

func (*WorkstationsWorkstationIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationsWorkstationIamBinding)(nil)).Elem()
}

func (i *WorkstationsWorkstationIamBinding) ToWorkstationsWorkstationIamBindingOutput() WorkstationsWorkstationIamBindingOutput {
	return i.ToWorkstationsWorkstationIamBindingOutputWithContext(context.Background())
}

func (i *WorkstationsWorkstationIamBinding) ToWorkstationsWorkstationIamBindingOutputWithContext(ctx context.Context) WorkstationsWorkstationIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationsWorkstationIamBindingOutput)
}

type WorkstationsWorkstationIamBindingOutput struct{ *pulumi.OutputState }

func (WorkstationsWorkstationIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationsWorkstationIamBinding)(nil)).Elem()
}

func (o WorkstationsWorkstationIamBindingOutput) ToWorkstationsWorkstationIamBindingOutput() WorkstationsWorkstationIamBindingOutput {
	return o
}

func (o WorkstationsWorkstationIamBindingOutput) ToWorkstationsWorkstationIamBindingOutputWithContext(ctx context.Context) WorkstationsWorkstationIamBindingOutput {
	return o
}

func (o WorkstationsWorkstationIamBindingOutput) Condition() WorkstationsWorkstationIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) WorkstationsWorkstationIamBindingConditionPtrOutput {
		return v.Condition
	}).(WorkstationsWorkstationIamBindingConditionPtrOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) WorkstationId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput { return v.WorkstationId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamBindingOutput) WorkstationsWorkstationIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamBinding) pulumi.StringOutput {
		return v.WorkstationsWorkstationIamBindingId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationsWorkstationIamBindingInput)(nil)).Elem(), &WorkstationsWorkstationIamBinding{})
	pulumi.RegisterOutputType(WorkstationsWorkstationIamBindingOutput{})
}
