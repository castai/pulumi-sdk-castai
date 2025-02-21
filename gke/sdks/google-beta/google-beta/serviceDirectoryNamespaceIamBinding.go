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

type ServiceDirectoryNamespaceIamBinding struct {
	pulumi.CustomResourceState

	Condition                             ServiceDirectoryNamespaceIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                                  pulumi.StringOutput                                   `pulumi:"etag"`
	Members                               pulumi.StringArrayOutput                              `pulumi:"members"`
	Name                                  pulumi.StringOutput                                   `pulumi:"name"`
	Role                                  pulumi.StringOutput                                   `pulumi:"role"`
	ServiceDirectoryNamespaceIamBindingId pulumi.StringOutput                                   `pulumi:"serviceDirectoryNamespaceIamBindingId"`
}

// NewServiceDirectoryNamespaceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewServiceDirectoryNamespaceIamBinding(ctx *pulumi.Context,
	name string, args *ServiceDirectoryNamespaceIamBindingArgs, opts ...pulumi.ResourceOption) (*ServiceDirectoryNamespaceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ServiceDirectoryNamespaceIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/serviceDirectoryNamespaceIamBinding:ServiceDirectoryNamespaceIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceDirectoryNamespaceIamBinding gets an existing ServiceDirectoryNamespaceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceDirectoryNamespaceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceDirectoryNamespaceIamBindingState, opts ...pulumi.ResourceOption) (*ServiceDirectoryNamespaceIamBinding, error) {
	var resource ServiceDirectoryNamespaceIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/serviceDirectoryNamespaceIamBinding:ServiceDirectoryNamespaceIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceDirectoryNamespaceIamBinding resources.
type serviceDirectoryNamespaceIamBindingState struct {
	Condition                             *ServiceDirectoryNamespaceIamBindingCondition `pulumi:"condition"`
	Etag                                  *string                                       `pulumi:"etag"`
	Members                               []string                                      `pulumi:"members"`
	Name                                  *string                                       `pulumi:"name"`
	Role                                  *string                                       `pulumi:"role"`
	ServiceDirectoryNamespaceIamBindingId *string                                       `pulumi:"serviceDirectoryNamespaceIamBindingId"`
}

type ServiceDirectoryNamespaceIamBindingState struct {
	Condition                             ServiceDirectoryNamespaceIamBindingConditionPtrInput
	Etag                                  pulumi.StringPtrInput
	Members                               pulumi.StringArrayInput
	Name                                  pulumi.StringPtrInput
	Role                                  pulumi.StringPtrInput
	ServiceDirectoryNamespaceIamBindingId pulumi.StringPtrInput
}

func (ServiceDirectoryNamespaceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDirectoryNamespaceIamBindingState)(nil)).Elem()
}

type serviceDirectoryNamespaceIamBindingArgs struct {
	Condition                             *ServiceDirectoryNamespaceIamBindingCondition `pulumi:"condition"`
	Members                               []string                                      `pulumi:"members"`
	Name                                  *string                                       `pulumi:"name"`
	Role                                  string                                        `pulumi:"role"`
	ServiceDirectoryNamespaceIamBindingId *string                                       `pulumi:"serviceDirectoryNamespaceIamBindingId"`
}

// The set of arguments for constructing a ServiceDirectoryNamespaceIamBinding resource.
type ServiceDirectoryNamespaceIamBindingArgs struct {
	Condition                             ServiceDirectoryNamespaceIamBindingConditionPtrInput
	Members                               pulumi.StringArrayInput
	Name                                  pulumi.StringPtrInput
	Role                                  pulumi.StringInput
	ServiceDirectoryNamespaceIamBindingId pulumi.StringPtrInput
}

func (ServiceDirectoryNamespaceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDirectoryNamespaceIamBindingArgs)(nil)).Elem()
}

type ServiceDirectoryNamespaceIamBindingInput interface {
	pulumi.Input

	ToServiceDirectoryNamespaceIamBindingOutput() ServiceDirectoryNamespaceIamBindingOutput
	ToServiceDirectoryNamespaceIamBindingOutputWithContext(ctx context.Context) ServiceDirectoryNamespaceIamBindingOutput
}

func (*ServiceDirectoryNamespaceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDirectoryNamespaceIamBinding)(nil)).Elem()
}

func (i *ServiceDirectoryNamespaceIamBinding) ToServiceDirectoryNamespaceIamBindingOutput() ServiceDirectoryNamespaceIamBindingOutput {
	return i.ToServiceDirectoryNamespaceIamBindingOutputWithContext(context.Background())
}

func (i *ServiceDirectoryNamespaceIamBinding) ToServiceDirectoryNamespaceIamBindingOutputWithContext(ctx context.Context) ServiceDirectoryNamespaceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDirectoryNamespaceIamBindingOutput)
}

type ServiceDirectoryNamespaceIamBindingOutput struct{ *pulumi.OutputState }

func (ServiceDirectoryNamespaceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDirectoryNamespaceIamBinding)(nil)).Elem()
}

func (o ServiceDirectoryNamespaceIamBindingOutput) ToServiceDirectoryNamespaceIamBindingOutput() ServiceDirectoryNamespaceIamBindingOutput {
	return o
}

func (o ServiceDirectoryNamespaceIamBindingOutput) ToServiceDirectoryNamespaceIamBindingOutputWithContext(ctx context.Context) ServiceDirectoryNamespaceIamBindingOutput {
	return o
}

func (o ServiceDirectoryNamespaceIamBindingOutput) Condition() ServiceDirectoryNamespaceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ServiceDirectoryNamespaceIamBinding) ServiceDirectoryNamespaceIamBindingConditionPtrOutput {
		return v.Condition
	}).(ServiceDirectoryNamespaceIamBindingConditionPtrOutput)
}

func (o ServiceDirectoryNamespaceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryNamespaceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ServiceDirectoryNamespaceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceDirectoryNamespaceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o ServiceDirectoryNamespaceIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryNamespaceIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceDirectoryNamespaceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryNamespaceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ServiceDirectoryNamespaceIamBindingOutput) ServiceDirectoryNamespaceIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryNamespaceIamBinding) pulumi.StringOutput {
		return v.ServiceDirectoryNamespaceIamBindingId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceDirectoryNamespaceIamBindingInput)(nil)).Elem(), &ServiceDirectoryNamespaceIamBinding{})
	pulumi.RegisterOutputType(ServiceDirectoryNamespaceIamBindingOutput{})
}
