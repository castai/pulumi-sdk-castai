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

type BeyondcorpSecurityGateway struct {
	pulumi.CustomResourceState

	BeyondcorpSecurityGatewayId pulumi.StringOutput `pulumi:"beyondcorpSecurityGatewayId"`
	// Output only. Timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. An arbitrary user-provided name for the SecurityGateway. Cannot exceed 64 characters.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Output only. IP addresses that will be used for establishing connection to the endpoints.
	ExternalIps pulumi.StringArrayOutput `pulumi:"externalIps"`
	// Optional. Map of Hubs that represents regional data path deployment with GCP region as a key.
	Hubs BeyondcorpSecurityGatewayHubArrayOutput `pulumi:"hubs"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. Must be omitted or set to 'global'.
	//
	// Deprecated: Deprecated
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Identifier. Name of the resource.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. User-settable SecurityGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters
	// from '/a-z-/'. * Must end with a number or letter.
	SecurityGatewayId pulumi.StringOutput `pulumi:"securityGatewayId"`
	// Output only. The operational state of the SecurityGateway. Possible values: STATE_UNSPECIFIED CREATING UPDATING DELETING
	// RUNNING DOWN ERROR
	State    pulumi.StringOutput                        `pulumi:"state"`
	Timeouts BeyondcorpSecurityGatewayTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. Timestamp when the resource was last modified.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewBeyondcorpSecurityGateway registers a new resource with the given unique name, arguments, and options.
func NewBeyondcorpSecurityGateway(ctx *pulumi.Context,
	name string, args *BeyondcorpSecurityGatewayArgs, opts ...pulumi.ResourceOption) (*BeyondcorpSecurityGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGatewayId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BeyondcorpSecurityGateway
	err = ctx.RegisterPackageResource("google-beta:index/beyondcorpSecurityGateway:BeyondcorpSecurityGateway", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBeyondcorpSecurityGateway gets an existing BeyondcorpSecurityGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBeyondcorpSecurityGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BeyondcorpSecurityGatewayState, opts ...pulumi.ResourceOption) (*BeyondcorpSecurityGateway, error) {
	var resource BeyondcorpSecurityGateway
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/beyondcorpSecurityGateway:BeyondcorpSecurityGateway", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BeyondcorpSecurityGateway resources.
type beyondcorpSecurityGatewayState struct {
	BeyondcorpSecurityGatewayId *string `pulumi:"beyondcorpSecurityGatewayId"`
	// Output only. Timestamp when the resource was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. An arbitrary user-provided name for the SecurityGateway. Cannot exceed 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// Output only. IP addresses that will be used for establishing connection to the endpoints.
	ExternalIps []string `pulumi:"externalIps"`
	// Optional. Map of Hubs that represents regional data path deployment with GCP region as a key.
	Hubs []BeyondcorpSecurityGatewayHub `pulumi:"hubs"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. Must be omitted or set to 'global'.
	//
	// Deprecated: Deprecated
	Location *string `pulumi:"location"`
	// Identifier. Name of the resource.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. User-settable SecurityGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters
	// from '/a-z-/'. * Must end with a number or letter.
	SecurityGatewayId *string `pulumi:"securityGatewayId"`
	// Output only. The operational state of the SecurityGateway. Possible values: STATE_UNSPECIFIED CREATING UPDATING DELETING
	// RUNNING DOWN ERROR
	State    *string                            `pulumi:"state"`
	Timeouts *BeyondcorpSecurityGatewayTimeouts `pulumi:"timeouts"`
	// Output only. Timestamp when the resource was last modified.
	UpdateTime *string `pulumi:"updateTime"`
}

type BeyondcorpSecurityGatewayState struct {
	BeyondcorpSecurityGatewayId pulumi.StringPtrInput
	// Output only. Timestamp when the resource was created.
	CreateTime pulumi.StringPtrInput
	// Optional. An arbitrary user-provided name for the SecurityGateway. Cannot exceed 64 characters.
	DisplayName pulumi.StringPtrInput
	// Output only. IP addresses that will be used for establishing connection to the endpoints.
	ExternalIps pulumi.StringArrayInput
	// Optional. Map of Hubs that represents regional data path deployment with GCP region as a key.
	Hubs BeyondcorpSecurityGatewayHubArrayInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. Must be omitted or set to 'global'.
	//
	// Deprecated: Deprecated
	Location pulumi.StringPtrInput
	// Identifier. Name of the resource.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. User-settable SecurityGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters
	// from '/a-z-/'. * Must end with a number or letter.
	SecurityGatewayId pulumi.StringPtrInput
	// Output only. The operational state of the SecurityGateway. Possible values: STATE_UNSPECIFIED CREATING UPDATING DELETING
	// RUNNING DOWN ERROR
	State    pulumi.StringPtrInput
	Timeouts BeyondcorpSecurityGatewayTimeoutsPtrInput
	// Output only. Timestamp when the resource was last modified.
	UpdateTime pulumi.StringPtrInput
}

func (BeyondcorpSecurityGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*beyondcorpSecurityGatewayState)(nil)).Elem()
}

type beyondcorpSecurityGatewayArgs struct {
	BeyondcorpSecurityGatewayId *string `pulumi:"beyondcorpSecurityGatewayId"`
	// Optional. An arbitrary user-provided name for the SecurityGateway. Cannot exceed 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Map of Hubs that represents regional data path deployment with GCP region as a key.
	Hubs []BeyondcorpSecurityGatewayHub `pulumi:"hubs"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. Must be omitted or set to 'global'.
	//
	// Deprecated: Deprecated
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Optional. User-settable SecurityGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters
	// from '/a-z-/'. * Must end with a number or letter.
	SecurityGatewayId string                             `pulumi:"securityGatewayId"`
	Timeouts          *BeyondcorpSecurityGatewayTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a BeyondcorpSecurityGateway resource.
type BeyondcorpSecurityGatewayArgs struct {
	BeyondcorpSecurityGatewayId pulumi.StringPtrInput
	// Optional. An arbitrary user-provided name for the SecurityGateway. Cannot exceed 64 characters.
	DisplayName pulumi.StringPtrInput
	// Optional. Map of Hubs that represents regional data path deployment with GCP region as a key.
	Hubs BeyondcorpSecurityGatewayHubArrayInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. Must be omitted or set to 'global'.
	//
	// Deprecated: Deprecated
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Optional. User-settable SecurityGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters
	// from '/a-z-/'. * Must end with a number or letter.
	SecurityGatewayId pulumi.StringInput
	Timeouts          BeyondcorpSecurityGatewayTimeoutsPtrInput
}

func (BeyondcorpSecurityGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*beyondcorpSecurityGatewayArgs)(nil)).Elem()
}

type BeyondcorpSecurityGatewayInput interface {
	pulumi.Input

	ToBeyondcorpSecurityGatewayOutput() BeyondcorpSecurityGatewayOutput
	ToBeyondcorpSecurityGatewayOutputWithContext(ctx context.Context) BeyondcorpSecurityGatewayOutput
}

func (*BeyondcorpSecurityGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**BeyondcorpSecurityGateway)(nil)).Elem()
}

func (i *BeyondcorpSecurityGateway) ToBeyondcorpSecurityGatewayOutput() BeyondcorpSecurityGatewayOutput {
	return i.ToBeyondcorpSecurityGatewayOutputWithContext(context.Background())
}

func (i *BeyondcorpSecurityGateway) ToBeyondcorpSecurityGatewayOutputWithContext(ctx context.Context) BeyondcorpSecurityGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BeyondcorpSecurityGatewayOutput)
}

type BeyondcorpSecurityGatewayOutput struct{ *pulumi.OutputState }

func (BeyondcorpSecurityGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BeyondcorpSecurityGateway)(nil)).Elem()
}

func (o BeyondcorpSecurityGatewayOutput) ToBeyondcorpSecurityGatewayOutput() BeyondcorpSecurityGatewayOutput {
	return o
}

func (o BeyondcorpSecurityGatewayOutput) ToBeyondcorpSecurityGatewayOutputWithContext(ctx context.Context) BeyondcorpSecurityGatewayOutput {
	return o
}

func (o BeyondcorpSecurityGatewayOutput) BeyondcorpSecurityGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringOutput { return v.BeyondcorpSecurityGatewayId }).(pulumi.StringOutput)
}

// Output only. Timestamp when the resource was created.
func (o BeyondcorpSecurityGatewayOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. An arbitrary user-provided name for the SecurityGateway. Cannot exceed 64 characters.
func (o BeyondcorpSecurityGatewayOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Output only. IP addresses that will be used for establishing connection to the endpoints.
func (o BeyondcorpSecurityGatewayOutput) ExternalIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringArrayOutput { return v.ExternalIps }).(pulumi.StringArrayOutput)
}

// Optional. Map of Hubs that represents regional data path deployment with GCP region as a key.
func (o BeyondcorpSecurityGatewayOutput) Hubs() BeyondcorpSecurityGatewayHubArrayOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) BeyondcorpSecurityGatewayHubArrayOutput { return v.Hubs }).(BeyondcorpSecurityGatewayHubArrayOutput)
}

// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
// https://google.aip.dev/122. Must be omitted or set to 'global'.
//
// Deprecated: Deprecated
func (o BeyondcorpSecurityGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Identifier. Name of the resource.
func (o BeyondcorpSecurityGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BeyondcorpSecurityGatewayOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. User-settable SecurityGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters
// from '/a-z-/'. * Must end with a number or letter.
func (o BeyondcorpSecurityGatewayOutput) SecurityGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringOutput { return v.SecurityGatewayId }).(pulumi.StringOutput)
}

// Output only. The operational state of the SecurityGateway. Possible values: STATE_UNSPECIFIED CREATING UPDATING DELETING
// RUNNING DOWN ERROR
func (o BeyondcorpSecurityGatewayOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o BeyondcorpSecurityGatewayOutput) Timeouts() BeyondcorpSecurityGatewayTimeoutsPtrOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) BeyondcorpSecurityGatewayTimeoutsPtrOutput { return v.Timeouts }).(BeyondcorpSecurityGatewayTimeoutsPtrOutput)
}

// Output only. Timestamp when the resource was last modified.
func (o BeyondcorpSecurityGatewayOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BeyondcorpSecurityGateway) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BeyondcorpSecurityGatewayInput)(nil)).Elem(), &BeyondcorpSecurityGateway{})
	pulumi.RegisterOutputType(BeyondcorpSecurityGatewayOutput{})
}
