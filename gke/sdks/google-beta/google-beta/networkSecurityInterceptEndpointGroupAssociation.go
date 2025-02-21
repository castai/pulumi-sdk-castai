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

type NetworkSecurityInterceptEndpointGroupAssociation struct {
	pulumi.CustomResourceState

	// Create time stamp.
	CreateTime      pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Immutable. The Intercept Endpoint Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/interceptEndpointGroups/{interceptEndpointGroup}'.
	InterceptEndpointGroup pulumi.StringOutput `pulumi:"interceptEndpointGroup"`
	// ID of the Intercept Endpoint Group Association.
	InterceptEndpointGroupAssociationId pulumi.StringPtrOutput `pulumi:"interceptEndpointGroupAssociationId"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location of the Intercept Endpoint Group Association, currently restricted to 'global'.
	Location pulumi.StringOutput `pulumi:"location"`
	// The list of locations that are currently supported by the associated Intercept Deployment Group and their state.
	LocationsDetails NetworkSecurityInterceptEndpointGroupAssociationLocationsDetailArrayOutput `pulumi:"locationsDetails"`
	// Identifier. The name of the Intercept Endpoint Group Association.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The VPC network associated. Format: 'projects/{project}/global/networks/{network}'.
	Network                                            pulumi.StringOutput `pulumi:"network"`
	NetworkSecurityInterceptEndpointGroupAssociationId pulumi.StringOutput `pulumi:"networkSecurityInterceptEndpointGroupAssociationId"`
	Project                                            pulumi.StringOutput `pulumi:"project"`
	// Whether reconciling is in progress.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Current state of the Intercept Endpoint Group Association. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	// CLOSED OUT_OF_SYNC DELETE_FAILED
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                                            `pulumi:"terraformLabels"`
	Timeouts        NetworkSecurityInterceptEndpointGroupAssociationTimeoutsPtrOutput `pulumi:"timeouts"`
	// Update time stamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecurityInterceptEndpointGroupAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityInterceptEndpointGroupAssociation(ctx *pulumi.Context,
	name string, args *NetworkSecurityInterceptEndpointGroupAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityInterceptEndpointGroupAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterceptEndpointGroup == nil {
		return nil, errors.New("invalid value for required argument 'InterceptEndpointGroup'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkSecurityInterceptEndpointGroupAssociation
	err = ctx.RegisterPackageResource("google-beta:index/networkSecurityInterceptEndpointGroupAssociation:NetworkSecurityInterceptEndpointGroupAssociation", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityInterceptEndpointGroupAssociation gets an existing NetworkSecurityInterceptEndpointGroupAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityInterceptEndpointGroupAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityInterceptEndpointGroupAssociationState, opts ...pulumi.ResourceOption) (*NetworkSecurityInterceptEndpointGroupAssociation, error) {
	var resource NetworkSecurityInterceptEndpointGroupAssociation
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkSecurityInterceptEndpointGroupAssociation:NetworkSecurityInterceptEndpointGroupAssociation", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityInterceptEndpointGroupAssociation resources.
type networkSecurityInterceptEndpointGroupAssociationState struct {
	// Create time stamp.
	CreateTime      *string           `pulumi:"createTime"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Immutable. The Intercept Endpoint Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/interceptEndpointGroups/{interceptEndpointGroup}'.
	InterceptEndpointGroup *string `pulumi:"interceptEndpointGroup"`
	// ID of the Intercept Endpoint Group Association.
	InterceptEndpointGroupAssociationId *string `pulumi:"interceptEndpointGroupAssociationId"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the Intercept Endpoint Group Association, currently restricted to 'global'.
	Location *string `pulumi:"location"`
	// The list of locations that are currently supported by the associated Intercept Deployment Group and their state.
	LocationsDetails []NetworkSecurityInterceptEndpointGroupAssociationLocationsDetail `pulumi:"locationsDetails"`
	// Identifier. The name of the Intercept Endpoint Group Association.
	Name *string `pulumi:"name"`
	// Immutable. The VPC network associated. Format: 'projects/{project}/global/networks/{network}'.
	Network                                            *string `pulumi:"network"`
	NetworkSecurityInterceptEndpointGroupAssociationId *string `pulumi:"networkSecurityInterceptEndpointGroupAssociationId"`
	Project                                            *string `pulumi:"project"`
	// Whether reconciling is in progress.
	Reconciling *bool `pulumi:"reconciling"`
	// Current state of the Intercept Endpoint Group Association. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	// CLOSED OUT_OF_SYNC DELETE_FAILED
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                                         `pulumi:"terraformLabels"`
	Timeouts        *NetworkSecurityInterceptEndpointGroupAssociationTimeouts `pulumi:"timeouts"`
	// Update time stamp.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecurityInterceptEndpointGroupAssociationState struct {
	// Create time stamp.
	CreateTime      pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Immutable. The Intercept Endpoint Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/interceptEndpointGroups/{interceptEndpointGroup}'.
	InterceptEndpointGroup pulumi.StringPtrInput
	// ID of the Intercept Endpoint Group Association.
	InterceptEndpointGroupAssociationId pulumi.StringPtrInput
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location of the Intercept Endpoint Group Association, currently restricted to 'global'.
	Location pulumi.StringPtrInput
	// The list of locations that are currently supported by the associated Intercept Deployment Group and their state.
	LocationsDetails NetworkSecurityInterceptEndpointGroupAssociationLocationsDetailArrayInput
	// Identifier. The name of the Intercept Endpoint Group Association.
	Name pulumi.StringPtrInput
	// Immutable. The VPC network associated. Format: 'projects/{project}/global/networks/{network}'.
	Network                                            pulumi.StringPtrInput
	NetworkSecurityInterceptEndpointGroupAssociationId pulumi.StringPtrInput
	Project                                            pulumi.StringPtrInput
	// Whether reconciling is in progress.
	Reconciling pulumi.BoolPtrInput
	// Current state of the Intercept Endpoint Group Association. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	// CLOSED OUT_OF_SYNC DELETE_FAILED
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkSecurityInterceptEndpointGroupAssociationTimeoutsPtrInput
	// Update time stamp.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecurityInterceptEndpointGroupAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityInterceptEndpointGroupAssociationState)(nil)).Elem()
}

type networkSecurityInterceptEndpointGroupAssociationArgs struct {
	// Immutable. The Intercept Endpoint Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/interceptEndpointGroups/{interceptEndpointGroup}'.
	InterceptEndpointGroup string `pulumi:"interceptEndpointGroup"`
	// ID of the Intercept Endpoint Group Association.
	InterceptEndpointGroupAssociationId *string `pulumi:"interceptEndpointGroupAssociationId"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the Intercept Endpoint Group Association, currently restricted to 'global'.
	Location string `pulumi:"location"`
	// Immutable. The VPC network associated. Format: 'projects/{project}/global/networks/{network}'.
	Network                                            string                                                    `pulumi:"network"`
	NetworkSecurityInterceptEndpointGroupAssociationId *string                                                   `pulumi:"networkSecurityInterceptEndpointGroupAssociationId"`
	Project                                            *string                                                   `pulumi:"project"`
	Timeouts                                           *NetworkSecurityInterceptEndpointGroupAssociationTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkSecurityInterceptEndpointGroupAssociation resource.
type NetworkSecurityInterceptEndpointGroupAssociationArgs struct {
	// Immutable. The Intercept Endpoint Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/interceptEndpointGroups/{interceptEndpointGroup}'.
	InterceptEndpointGroup pulumi.StringInput
	// ID of the Intercept Endpoint Group Association.
	InterceptEndpointGroupAssociationId pulumi.StringPtrInput
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location of the Intercept Endpoint Group Association, currently restricted to 'global'.
	Location pulumi.StringInput
	// Immutable. The VPC network associated. Format: 'projects/{project}/global/networks/{network}'.
	Network                                            pulumi.StringInput
	NetworkSecurityInterceptEndpointGroupAssociationId pulumi.StringPtrInput
	Project                                            pulumi.StringPtrInput
	Timeouts                                           NetworkSecurityInterceptEndpointGroupAssociationTimeoutsPtrInput
}

func (NetworkSecurityInterceptEndpointGroupAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityInterceptEndpointGroupAssociationArgs)(nil)).Elem()
}

type NetworkSecurityInterceptEndpointGroupAssociationInput interface {
	pulumi.Input

	ToNetworkSecurityInterceptEndpointGroupAssociationOutput() NetworkSecurityInterceptEndpointGroupAssociationOutput
	ToNetworkSecurityInterceptEndpointGroupAssociationOutputWithContext(ctx context.Context) NetworkSecurityInterceptEndpointGroupAssociationOutput
}

func (*NetworkSecurityInterceptEndpointGroupAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityInterceptEndpointGroupAssociation)(nil)).Elem()
}

func (i *NetworkSecurityInterceptEndpointGroupAssociation) ToNetworkSecurityInterceptEndpointGroupAssociationOutput() NetworkSecurityInterceptEndpointGroupAssociationOutput {
	return i.ToNetworkSecurityInterceptEndpointGroupAssociationOutputWithContext(context.Background())
}

func (i *NetworkSecurityInterceptEndpointGroupAssociation) ToNetworkSecurityInterceptEndpointGroupAssociationOutputWithContext(ctx context.Context) NetworkSecurityInterceptEndpointGroupAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityInterceptEndpointGroupAssociationOutput)
}

type NetworkSecurityInterceptEndpointGroupAssociationOutput struct{ *pulumi.OutputState }

func (NetworkSecurityInterceptEndpointGroupAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityInterceptEndpointGroupAssociation)(nil)).Elem()
}

func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) ToNetworkSecurityInterceptEndpointGroupAssociationOutput() NetworkSecurityInterceptEndpointGroupAssociationOutput {
	return o
}

func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) ToNetworkSecurityInterceptEndpointGroupAssociationOutputWithContext(ctx context.Context) NetworkSecurityInterceptEndpointGroupAssociationOutput {
	return o
}

// Create time stamp.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringMapOutput {
		return v.EffectiveLabels
	}).(pulumi.StringMapOutput)
}

// Immutable. The Intercept Endpoint Group that this resource is connected to. Format is:
// 'projects/{project}/locations/global/interceptEndpointGroups/{interceptEndpointGroup}'.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) InterceptEndpointGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput {
		return v.InterceptEndpointGroup
	}).(pulumi.StringOutput)
}

// ID of the Intercept Endpoint Group Association.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) InterceptEndpointGroupAssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringPtrOutput {
		return v.InterceptEndpointGroupAssociationId
	}).(pulumi.StringPtrOutput)
}

// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location of the Intercept Endpoint Group Association, currently restricted to 'global'.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The list of locations that are currently supported by the associated Intercept Deployment Group and their state.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) LocationsDetails() NetworkSecurityInterceptEndpointGroupAssociationLocationsDetailArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) NetworkSecurityInterceptEndpointGroupAssociationLocationsDetailArrayOutput {
		return v.LocationsDetails
	}).(NetworkSecurityInterceptEndpointGroupAssociationLocationsDetailArrayOutput)
}

// Identifier. The name of the Intercept Endpoint Group Association.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The VPC network associated. Format: 'projects/{project}/global/networks/{network}'.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) NetworkSecurityInterceptEndpointGroupAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput {
		return v.NetworkSecurityInterceptEndpointGroupAssociationId
	}).(pulumi.StringOutput)
}

func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Whether reconciling is in progress.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Current state of the Intercept Endpoint Group Association. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
// CLOSED OUT_OF_SYNC DELETE_FAILED
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringMapOutput {
		return v.TerraformLabels
	}).(pulumi.StringMapOutput)
}

func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) Timeouts() NetworkSecurityInterceptEndpointGroupAssociationTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) NetworkSecurityInterceptEndpointGroupAssociationTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecurityInterceptEndpointGroupAssociationTimeoutsPtrOutput)
}

// Update time stamp.
func (o NetworkSecurityInterceptEndpointGroupAssociationOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptEndpointGroupAssociation) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityInterceptEndpointGroupAssociationInput)(nil)).Elem(), &NetworkSecurityInterceptEndpointGroupAssociation{})
	pulumi.RegisterOutputType(NetworkSecurityInterceptEndpointGroupAssociationOutput{})
}
