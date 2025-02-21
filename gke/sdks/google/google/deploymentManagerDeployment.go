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

type DeploymentManagerDeployment struct {
	pulumi.CustomResourceState

	// Set the policy to use for creating new resources. Only used on create and update. Valid values are 'CREATE_OR_ACQUIRE'
	// (default) or 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist, the deployment will fail. Note that
	// updating this field does not actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE"
	// Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
	CreatePolicy pulumi.StringPtrOutput `pulumi:"createPolicy"`
	// Set the policy to use for deleting new resources on update/delete. Valid values are 'DELETE' (default) or 'ABANDON'. If
	// 'DELETE', resource is deleted after removal from Deployment Manager. If 'ABANDON', the resource is only removed from
	// Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
	// just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
	DeletePolicy pulumi.StringPtrOutput `pulumi:"deletePolicy"`
	// Unique identifier for deployment. Output only.
	DeploymentId                  pulumi.StringOutput `pulumi:"deploymentId"`
	DeploymentManagerDeploymentId pulumi.StringOutput `pulumi:"deploymentManagerDeploymentId"`
	// Optional user-provided description of deployment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Key-value pairs to apply to this labels.
	Labels DeploymentManagerDeploymentLabelArrayOutput `pulumi:"labels"`
	// Output only. URL of the manifest representing the last manifest that was successfully deployed.
	Manifest pulumi.StringOutput `pulumi:"manifest"`
	// Unique name for the deployment
	Name    pulumi.StringOutput  `pulumi:"name"`
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	Project pulumi.StringOutput  `pulumi:"project"`
	// Output only. Server defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Parameters that define your deployment, including the deployment configuration and relevant templates.
	Target   DeploymentManagerDeploymentTargetOutput      `pulumi:"target"`
	Timeouts DeploymentManagerDeploymentTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDeploymentManagerDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeploymentManagerDeployment(ctx *pulumi.Context,
	name string, args *DeploymentManagerDeploymentArgs, opts ...pulumi.ResourceOption) (*DeploymentManagerDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DeploymentManagerDeployment
	err = ctx.RegisterPackageResource("google:index/deploymentManagerDeployment:DeploymentManagerDeployment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentManagerDeployment gets an existing DeploymentManagerDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentManagerDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentManagerDeploymentState, opts ...pulumi.ResourceOption) (*DeploymentManagerDeployment, error) {
	var resource DeploymentManagerDeployment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/deploymentManagerDeployment:DeploymentManagerDeployment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentManagerDeployment resources.
type deploymentManagerDeploymentState struct {
	// Set the policy to use for creating new resources. Only used on create and update. Valid values are 'CREATE_OR_ACQUIRE'
	// (default) or 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist, the deployment will fail. Note that
	// updating this field does not actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE"
	// Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
	CreatePolicy *string `pulumi:"createPolicy"`
	// Set the policy to use for deleting new resources on update/delete. Valid values are 'DELETE' (default) or 'ABANDON'. If
	// 'DELETE', resource is deleted after removal from Deployment Manager. If 'ABANDON', the resource is only removed from
	// Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
	// just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
	DeletePolicy *string `pulumi:"deletePolicy"`
	// Unique identifier for deployment. Output only.
	DeploymentId                  *string `pulumi:"deploymentId"`
	DeploymentManagerDeploymentId *string `pulumi:"deploymentManagerDeploymentId"`
	// Optional user-provided description of deployment.
	Description *string `pulumi:"description"`
	// Key-value pairs to apply to this labels.
	Labels []DeploymentManagerDeploymentLabel `pulumi:"labels"`
	// Output only. URL of the manifest representing the last manifest that was successfully deployed.
	Manifest *string `pulumi:"manifest"`
	// Unique name for the deployment
	Name    *string `pulumi:"name"`
	Preview *bool   `pulumi:"preview"`
	Project *string `pulumi:"project"`
	// Output only. Server defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Parameters that define your deployment, including the deployment configuration and relevant templates.
	Target   *DeploymentManagerDeploymentTarget   `pulumi:"target"`
	Timeouts *DeploymentManagerDeploymentTimeouts `pulumi:"timeouts"`
}

type DeploymentManagerDeploymentState struct {
	// Set the policy to use for creating new resources. Only used on create and update. Valid values are 'CREATE_OR_ACQUIRE'
	// (default) or 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist, the deployment will fail. Note that
	// updating this field does not actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE"
	// Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
	CreatePolicy pulumi.StringPtrInput
	// Set the policy to use for deleting new resources on update/delete. Valid values are 'DELETE' (default) or 'ABANDON'. If
	// 'DELETE', resource is deleted after removal from Deployment Manager. If 'ABANDON', the resource is only removed from
	// Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
	// just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
	DeletePolicy pulumi.StringPtrInput
	// Unique identifier for deployment. Output only.
	DeploymentId                  pulumi.StringPtrInput
	DeploymentManagerDeploymentId pulumi.StringPtrInput
	// Optional user-provided description of deployment.
	Description pulumi.StringPtrInput
	// Key-value pairs to apply to this labels.
	Labels DeploymentManagerDeploymentLabelArrayInput
	// Output only. URL of the manifest representing the last manifest that was successfully deployed.
	Manifest pulumi.StringPtrInput
	// Unique name for the deployment
	Name    pulumi.StringPtrInput
	Preview pulumi.BoolPtrInput
	Project pulumi.StringPtrInput
	// Output only. Server defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Parameters that define your deployment, including the deployment configuration and relevant templates.
	Target   DeploymentManagerDeploymentTargetPtrInput
	Timeouts DeploymentManagerDeploymentTimeoutsPtrInput
}

func (DeploymentManagerDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentManagerDeploymentState)(nil)).Elem()
}

type deploymentManagerDeploymentArgs struct {
	// Set the policy to use for creating new resources. Only used on create and update. Valid values are 'CREATE_OR_ACQUIRE'
	// (default) or 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist, the deployment will fail. Note that
	// updating this field does not actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE"
	// Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
	CreatePolicy *string `pulumi:"createPolicy"`
	// Set the policy to use for deleting new resources on update/delete. Valid values are 'DELETE' (default) or 'ABANDON'. If
	// 'DELETE', resource is deleted after removal from Deployment Manager. If 'ABANDON', the resource is only removed from
	// Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
	// just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
	DeletePolicy                  *string `pulumi:"deletePolicy"`
	DeploymentManagerDeploymentId *string `pulumi:"deploymentManagerDeploymentId"`
	// Optional user-provided description of deployment.
	Description *string `pulumi:"description"`
	// Key-value pairs to apply to this labels.
	Labels []DeploymentManagerDeploymentLabel `pulumi:"labels"`
	// Unique name for the deployment
	Name    *string `pulumi:"name"`
	Preview *bool   `pulumi:"preview"`
	Project *string `pulumi:"project"`
	// Parameters that define your deployment, including the deployment configuration and relevant templates.
	Target   DeploymentManagerDeploymentTarget    `pulumi:"target"`
	Timeouts *DeploymentManagerDeploymentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DeploymentManagerDeployment resource.
type DeploymentManagerDeploymentArgs struct {
	// Set the policy to use for creating new resources. Only used on create and update. Valid values are 'CREATE_OR_ACQUIRE'
	// (default) or 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist, the deployment will fail. Note that
	// updating this field does not actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE"
	// Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
	CreatePolicy pulumi.StringPtrInput
	// Set the policy to use for deleting new resources on update/delete. Valid values are 'DELETE' (default) or 'ABANDON'. If
	// 'DELETE', resource is deleted after removal from Deployment Manager. If 'ABANDON', the resource is only removed from
	// Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
	// just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
	DeletePolicy                  pulumi.StringPtrInput
	DeploymentManagerDeploymentId pulumi.StringPtrInput
	// Optional user-provided description of deployment.
	Description pulumi.StringPtrInput
	// Key-value pairs to apply to this labels.
	Labels DeploymentManagerDeploymentLabelArrayInput
	// Unique name for the deployment
	Name    pulumi.StringPtrInput
	Preview pulumi.BoolPtrInput
	Project pulumi.StringPtrInput
	// Parameters that define your deployment, including the deployment configuration and relevant templates.
	Target   DeploymentManagerDeploymentTargetInput
	Timeouts DeploymentManagerDeploymentTimeoutsPtrInput
}

func (DeploymentManagerDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentManagerDeploymentArgs)(nil)).Elem()
}

type DeploymentManagerDeploymentInput interface {
	pulumi.Input

	ToDeploymentManagerDeploymentOutput() DeploymentManagerDeploymentOutput
	ToDeploymentManagerDeploymentOutputWithContext(ctx context.Context) DeploymentManagerDeploymentOutput
}

func (*DeploymentManagerDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentManagerDeployment)(nil)).Elem()
}

func (i *DeploymentManagerDeployment) ToDeploymentManagerDeploymentOutput() DeploymentManagerDeploymentOutput {
	return i.ToDeploymentManagerDeploymentOutputWithContext(context.Background())
}

func (i *DeploymentManagerDeployment) ToDeploymentManagerDeploymentOutputWithContext(ctx context.Context) DeploymentManagerDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentManagerDeploymentOutput)
}

type DeploymentManagerDeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentManagerDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentManagerDeployment)(nil)).Elem()
}

func (o DeploymentManagerDeploymentOutput) ToDeploymentManagerDeploymentOutput() DeploymentManagerDeploymentOutput {
	return o
}

func (o DeploymentManagerDeploymentOutput) ToDeploymentManagerDeploymentOutputWithContext(ctx context.Context) DeploymentManagerDeploymentOutput {
	return o
}

// Set the policy to use for creating new resources. Only used on create and update. Valid values are 'CREATE_OR_ACQUIRE'
// (default) or 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist, the deployment will fail. Note that
// updating this field does not actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE"
// Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
func (o DeploymentManagerDeploymentOutput) CreatePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringPtrOutput { return v.CreatePolicy }).(pulumi.StringPtrOutput)
}

// Set the policy to use for deleting new resources on update/delete. Valid values are 'DELETE' (default) or 'ABANDON'. If
// 'DELETE', resource is deleted after removal from Deployment Manager. If 'ABANDON', the resource is only removed from
// Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
// just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
func (o DeploymentManagerDeploymentOutput) DeletePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringPtrOutput { return v.DeletePolicy }).(pulumi.StringPtrOutput)
}

// Unique identifier for deployment. Output only.
func (o DeploymentManagerDeploymentOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o DeploymentManagerDeploymentOutput) DeploymentManagerDeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringOutput { return v.DeploymentManagerDeploymentId }).(pulumi.StringOutput)
}

// Optional user-provided description of deployment.
func (o DeploymentManagerDeploymentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Key-value pairs to apply to this labels.
func (o DeploymentManagerDeploymentOutput) Labels() DeploymentManagerDeploymentLabelArrayOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) DeploymentManagerDeploymentLabelArrayOutput { return v.Labels }).(DeploymentManagerDeploymentLabelArrayOutput)
}

// Output only. URL of the manifest representing the last manifest that was successfully deployed.
func (o DeploymentManagerDeploymentOutput) Manifest() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringOutput { return v.Manifest }).(pulumi.StringOutput)
}

// Unique name for the deployment
func (o DeploymentManagerDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentManagerDeploymentOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.BoolPtrOutput { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o DeploymentManagerDeploymentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Server defined URL for the resource.
func (o DeploymentManagerDeploymentOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Parameters that define your deployment, including the deployment configuration and relevant templates.
func (o DeploymentManagerDeploymentOutput) Target() DeploymentManagerDeploymentTargetOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) DeploymentManagerDeploymentTargetOutput { return v.Target }).(DeploymentManagerDeploymentTargetOutput)
}

func (o DeploymentManagerDeploymentOutput) Timeouts() DeploymentManagerDeploymentTimeoutsPtrOutput {
	return o.ApplyT(func(v *DeploymentManagerDeployment) DeploymentManagerDeploymentTimeoutsPtrOutput { return v.Timeouts }).(DeploymentManagerDeploymentTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentManagerDeploymentInput)(nil)).Elem(), &DeploymentManagerDeployment{})
	pulumi.RegisterOutputType(DeploymentManagerDeploymentOutput{})
}
