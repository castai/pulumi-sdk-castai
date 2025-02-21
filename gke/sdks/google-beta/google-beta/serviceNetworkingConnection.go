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

type ServiceNetworkingConnection struct {
	pulumi.CustomResourceState

	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// Name of VPC network connected with service producers using VPC peering.
	Network pulumi.StringOutput `pulumi:"network"`
	Peering pulumi.StringOutput `pulumi:"peering"`
	// Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a
	// different range when connection is already established will not reallocate already provisioned service producer
	// subnetworks.
	ReservedPeeringRanges pulumi.StringArrayOutput `pulumi:"reservedPeeringRanges"`
	// Provider peering service that is managing peering connectivity for a service provider organization. For Google services
	// that support this functionality it is 'servicenetworking.googleapis.com'.
	Service                       pulumi.StringOutput                          `pulumi:"service"`
	ServiceNetworkingConnectionId pulumi.StringOutput                          `pulumi:"serviceNetworkingConnectionId"`
	Timeouts                      ServiceNetworkingConnectionTimeoutsPtrOutput `pulumi:"timeouts"`
	// When set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case
	// of a new connection creation failure.
	UpdateOnCreationFail pulumi.BoolPtrOutput `pulumi:"updateOnCreationFail"`
}

// NewServiceNetworkingConnection registers a new resource with the given unique name, arguments, and options.
func NewServiceNetworkingConnection(ctx *pulumi.Context,
	name string, args *ServiceNetworkingConnectionArgs, opts ...pulumi.ResourceOption) (*ServiceNetworkingConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.ReservedPeeringRanges == nil {
		return nil, errors.New("invalid value for required argument 'ReservedPeeringRanges'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ServiceNetworkingConnection
	err = ctx.RegisterPackageResource("google-beta:index/serviceNetworkingConnection:ServiceNetworkingConnection", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceNetworkingConnection gets an existing ServiceNetworkingConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceNetworkingConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceNetworkingConnectionState, opts ...pulumi.ResourceOption) (*ServiceNetworkingConnection, error) {
	var resource ServiceNetworkingConnection
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/serviceNetworkingConnection:ServiceNetworkingConnection", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceNetworkingConnection resources.
type serviceNetworkingConnectionState struct {
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// Name of VPC network connected with service producers using VPC peering.
	Network *string `pulumi:"network"`
	Peering *string `pulumi:"peering"`
	// Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a
	// different range when connection is already established will not reallocate already provisioned service producer
	// subnetworks.
	ReservedPeeringRanges []string `pulumi:"reservedPeeringRanges"`
	// Provider peering service that is managing peering connectivity for a service provider organization. For Google services
	// that support this functionality it is 'servicenetworking.googleapis.com'.
	Service                       *string                              `pulumi:"service"`
	ServiceNetworkingConnectionId *string                              `pulumi:"serviceNetworkingConnectionId"`
	Timeouts                      *ServiceNetworkingConnectionTimeouts `pulumi:"timeouts"`
	// When set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case
	// of a new connection creation failure.
	UpdateOnCreationFail *bool `pulumi:"updateOnCreationFail"`
}

type ServiceNetworkingConnectionState struct {
	DeletionPolicy pulumi.StringPtrInput
	// Name of VPC network connected with service producers using VPC peering.
	Network pulumi.StringPtrInput
	Peering pulumi.StringPtrInput
	// Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a
	// different range when connection is already established will not reallocate already provisioned service producer
	// subnetworks.
	ReservedPeeringRanges pulumi.StringArrayInput
	// Provider peering service that is managing peering connectivity for a service provider organization. For Google services
	// that support this functionality it is 'servicenetworking.googleapis.com'.
	Service                       pulumi.StringPtrInput
	ServiceNetworkingConnectionId pulumi.StringPtrInput
	Timeouts                      ServiceNetworkingConnectionTimeoutsPtrInput
	// When set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case
	// of a new connection creation failure.
	UpdateOnCreationFail pulumi.BoolPtrInput
}

func (ServiceNetworkingConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceNetworkingConnectionState)(nil)).Elem()
}

type serviceNetworkingConnectionArgs struct {
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// Name of VPC network connected with service producers using VPC peering.
	Network string `pulumi:"network"`
	// Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a
	// different range when connection is already established will not reallocate already provisioned service producer
	// subnetworks.
	ReservedPeeringRanges []string `pulumi:"reservedPeeringRanges"`
	// Provider peering service that is managing peering connectivity for a service provider organization. For Google services
	// that support this functionality it is 'servicenetworking.googleapis.com'.
	Service                       string                               `pulumi:"service"`
	ServiceNetworkingConnectionId *string                              `pulumi:"serviceNetworkingConnectionId"`
	Timeouts                      *ServiceNetworkingConnectionTimeouts `pulumi:"timeouts"`
	// When set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case
	// of a new connection creation failure.
	UpdateOnCreationFail *bool `pulumi:"updateOnCreationFail"`
}

// The set of arguments for constructing a ServiceNetworkingConnection resource.
type ServiceNetworkingConnectionArgs struct {
	DeletionPolicy pulumi.StringPtrInput
	// Name of VPC network connected with service producers using VPC peering.
	Network pulumi.StringInput
	// Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a
	// different range when connection is already established will not reallocate already provisioned service producer
	// subnetworks.
	ReservedPeeringRanges pulumi.StringArrayInput
	// Provider peering service that is managing peering connectivity for a service provider organization. For Google services
	// that support this functionality it is 'servicenetworking.googleapis.com'.
	Service                       pulumi.StringInput
	ServiceNetworkingConnectionId pulumi.StringPtrInput
	Timeouts                      ServiceNetworkingConnectionTimeoutsPtrInput
	// When set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case
	// of a new connection creation failure.
	UpdateOnCreationFail pulumi.BoolPtrInput
}

func (ServiceNetworkingConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceNetworkingConnectionArgs)(nil)).Elem()
}

type ServiceNetworkingConnectionInput interface {
	pulumi.Input

	ToServiceNetworkingConnectionOutput() ServiceNetworkingConnectionOutput
	ToServiceNetworkingConnectionOutputWithContext(ctx context.Context) ServiceNetworkingConnectionOutput
}

func (*ServiceNetworkingConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceNetworkingConnection)(nil)).Elem()
}

func (i *ServiceNetworkingConnection) ToServiceNetworkingConnectionOutput() ServiceNetworkingConnectionOutput {
	return i.ToServiceNetworkingConnectionOutputWithContext(context.Background())
}

func (i *ServiceNetworkingConnection) ToServiceNetworkingConnectionOutputWithContext(ctx context.Context) ServiceNetworkingConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceNetworkingConnectionOutput)
}

type ServiceNetworkingConnectionOutput struct{ *pulumi.OutputState }

func (ServiceNetworkingConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceNetworkingConnection)(nil)).Elem()
}

func (o ServiceNetworkingConnectionOutput) ToServiceNetworkingConnectionOutput() ServiceNetworkingConnectionOutput {
	return o
}

func (o ServiceNetworkingConnectionOutput) ToServiceNetworkingConnectionOutputWithContext(ctx context.Context) ServiceNetworkingConnectionOutput {
	return o
}

func (o ServiceNetworkingConnectionOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

// Name of VPC network connected with service producers using VPC peering.
func (o ServiceNetworkingConnectionOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o ServiceNetworkingConnectionOutput) Peering() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) pulumi.StringOutput { return v.Peering }).(pulumi.StringOutput)
}

// Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a
// different range when connection is already established will not reallocate already provisioned service producer
// subnetworks.
func (o ServiceNetworkingConnectionOutput) ReservedPeeringRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) pulumi.StringArrayOutput { return v.ReservedPeeringRanges }).(pulumi.StringArrayOutput)
}

// Provider peering service that is managing peering connectivity for a service provider organization. For Google services
// that support this functionality it is 'servicenetworking.googleapis.com'.
func (o ServiceNetworkingConnectionOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

func (o ServiceNetworkingConnectionOutput) ServiceNetworkingConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) pulumi.StringOutput { return v.ServiceNetworkingConnectionId }).(pulumi.StringOutput)
}

func (o ServiceNetworkingConnectionOutput) Timeouts() ServiceNetworkingConnectionTimeoutsPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) ServiceNetworkingConnectionTimeoutsPtrOutput { return v.Timeouts }).(ServiceNetworkingConnectionTimeoutsPtrOutput)
}

// When set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case
// of a new connection creation failure.
func (o ServiceNetworkingConnectionOutput) UpdateOnCreationFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkingConnection) pulumi.BoolPtrOutput { return v.UpdateOnCreationFail }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceNetworkingConnectionInput)(nil)).Elem(), &ServiceNetworkingConnection{})
	pulumi.RegisterOutputType(ServiceNetworkingConnectionOutput{})
}
