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

type IntegrationConnectorsManagedZone struct {
	pulumi.CustomResourceState

	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// DNS Name of the resource.
	Dns                                pulumi.StringOutput    `pulumi:"dns"`
	EffectiveLabels                    pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	IntegrationConnectorsManagedZoneId pulumi.StringOutput    `pulumi:"integrationConnectorsManagedZoneId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of Managed Zone needs to be created.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the Target Project.
	TargetProject pulumi.StringOutput `pulumi:"targetProject"`
	// The name of the Target Project VPC Network.
	TargetVpc pulumi.StringOutput `pulumi:"targetVpc"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                            `pulumi:"terraformLabels"`
	Timeouts        IntegrationConnectorsManagedZoneTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewIntegrationConnectorsManagedZone registers a new resource with the given unique name, arguments, and options.
func NewIntegrationConnectorsManagedZone(ctx *pulumi.Context,
	name string, args *IntegrationConnectorsManagedZoneArgs, opts ...pulumi.ResourceOption) (*IntegrationConnectorsManagedZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dns == nil {
		return nil, errors.New("invalid value for required argument 'Dns'")
	}
	if args.TargetProject == nil {
		return nil, errors.New("invalid value for required argument 'TargetProject'")
	}
	if args.TargetVpc == nil {
		return nil, errors.New("invalid value for required argument 'TargetVpc'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IntegrationConnectorsManagedZone
	err = ctx.RegisterPackageResource("google-beta:index/integrationConnectorsManagedZone:IntegrationConnectorsManagedZone", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationConnectorsManagedZone gets an existing IntegrationConnectorsManagedZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationConnectorsManagedZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationConnectorsManagedZoneState, opts ...pulumi.ResourceOption) (*IntegrationConnectorsManagedZone, error) {
	var resource IntegrationConnectorsManagedZone
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/integrationConnectorsManagedZone:IntegrationConnectorsManagedZone", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationConnectorsManagedZone resources.
type integrationConnectorsManagedZoneState struct {
	// Time the Namespace was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// DNS Name of the resource.
	Dns                                *string           `pulumi:"dns"`
	EffectiveLabels                    map[string]string `pulumi:"effectiveLabels"`
	IntegrationConnectorsManagedZoneId *string           `pulumi:"integrationConnectorsManagedZoneId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of Managed Zone needs to be created.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The name of the Target Project.
	TargetProject *string `pulumi:"targetProject"`
	// The name of the Target Project VPC Network.
	TargetVpc *string `pulumi:"targetVpc"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                         `pulumi:"terraformLabels"`
	Timeouts        *IntegrationConnectorsManagedZoneTimeouts `pulumi:"timeouts"`
	// Time the Namespace was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type IntegrationConnectorsManagedZoneState struct {
	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// DNS Name of the resource.
	Dns                                pulumi.StringPtrInput
	EffectiveLabels                    pulumi.StringMapInput
	IntegrationConnectorsManagedZoneId pulumi.StringPtrInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Name of Managed Zone needs to be created.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The name of the Target Project.
	TargetProject pulumi.StringPtrInput
	// The name of the Target Project VPC Network.
	TargetVpc pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        IntegrationConnectorsManagedZoneTimeoutsPtrInput
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (IntegrationConnectorsManagedZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationConnectorsManagedZoneState)(nil)).Elem()
}

type integrationConnectorsManagedZoneArgs struct {
	// Description of the resource.
	Description *string `pulumi:"description"`
	// DNS Name of the resource.
	Dns                                string  `pulumi:"dns"`
	IntegrationConnectorsManagedZoneId *string `pulumi:"integrationConnectorsManagedZoneId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of Managed Zone needs to be created.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The name of the Target Project.
	TargetProject string `pulumi:"targetProject"`
	// The name of the Target Project VPC Network.
	TargetVpc string                                    `pulumi:"targetVpc"`
	Timeouts  *IntegrationConnectorsManagedZoneTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a IntegrationConnectorsManagedZone resource.
type IntegrationConnectorsManagedZoneArgs struct {
	// Description of the resource.
	Description pulumi.StringPtrInput
	// DNS Name of the resource.
	Dns                                pulumi.StringInput
	IntegrationConnectorsManagedZoneId pulumi.StringPtrInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Name of Managed Zone needs to be created.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The name of the Target Project.
	TargetProject pulumi.StringInput
	// The name of the Target Project VPC Network.
	TargetVpc pulumi.StringInput
	Timeouts  IntegrationConnectorsManagedZoneTimeoutsPtrInput
}

func (IntegrationConnectorsManagedZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationConnectorsManagedZoneArgs)(nil)).Elem()
}

type IntegrationConnectorsManagedZoneInput interface {
	pulumi.Input

	ToIntegrationConnectorsManagedZoneOutput() IntegrationConnectorsManagedZoneOutput
	ToIntegrationConnectorsManagedZoneOutputWithContext(ctx context.Context) IntegrationConnectorsManagedZoneOutput
}

func (*IntegrationConnectorsManagedZone) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationConnectorsManagedZone)(nil)).Elem()
}

func (i *IntegrationConnectorsManagedZone) ToIntegrationConnectorsManagedZoneOutput() IntegrationConnectorsManagedZoneOutput {
	return i.ToIntegrationConnectorsManagedZoneOutputWithContext(context.Background())
}

func (i *IntegrationConnectorsManagedZone) ToIntegrationConnectorsManagedZoneOutputWithContext(ctx context.Context) IntegrationConnectorsManagedZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationConnectorsManagedZoneOutput)
}

type IntegrationConnectorsManagedZoneOutput struct{ *pulumi.OutputState }

func (IntegrationConnectorsManagedZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationConnectorsManagedZone)(nil)).Elem()
}

func (o IntegrationConnectorsManagedZoneOutput) ToIntegrationConnectorsManagedZoneOutput() IntegrationConnectorsManagedZoneOutput {
	return o
}

func (o IntegrationConnectorsManagedZoneOutput) ToIntegrationConnectorsManagedZoneOutputWithContext(ctx context.Context) IntegrationConnectorsManagedZoneOutput {
	return o
}

// Time the Namespace was created in UTC.
func (o IntegrationConnectorsManagedZoneOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the resource.
func (o IntegrationConnectorsManagedZoneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// DNS Name of the resource.
func (o IntegrationConnectorsManagedZoneOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o IntegrationConnectorsManagedZoneOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o IntegrationConnectorsManagedZoneOutput) IntegrationConnectorsManagedZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput {
		return v.IntegrationConnectorsManagedZoneId
	}).(pulumi.StringOutput)
}

// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o IntegrationConnectorsManagedZoneOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of Managed Zone needs to be created.
func (o IntegrationConnectorsManagedZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IntegrationConnectorsManagedZoneOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the Target Project.
func (o IntegrationConnectorsManagedZoneOutput) TargetProject() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput { return v.TargetProject }).(pulumi.StringOutput)
}

// The name of the Target Project VPC Network.
func (o IntegrationConnectorsManagedZoneOutput) TargetVpc() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput { return v.TargetVpc }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o IntegrationConnectorsManagedZoneOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o IntegrationConnectorsManagedZoneOutput) Timeouts() IntegrationConnectorsManagedZoneTimeoutsPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) IntegrationConnectorsManagedZoneTimeoutsPtrOutput {
		return v.Timeouts
	}).(IntegrationConnectorsManagedZoneTimeoutsPtrOutput)
}

// Time the Namespace was updated in UTC.
func (o IntegrationConnectorsManagedZoneOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsManagedZone) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationConnectorsManagedZoneInput)(nil)).Elem(), &IntegrationConnectorsManagedZone{})
	pulumi.RegisterOutputType(IntegrationConnectorsManagedZoneOutput{})
}
