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

type ApigeeSharedflowDeployment struct {
	pulumi.CustomResourceState

	ApigeeSharedflowDeploymentId pulumi.StringOutput `pulumi:"apigeeSharedflowDeploymentId"`
	// The resource ID of the environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The Apigee Organization associated with the Apigee instance
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Revision of the Sharedflow to be deployed.
	Revision pulumi.StringOutput `pulumi:"revision"`
	// The service account represents the identity of the deployed proxy, and determines what permissions it has. The format
	// must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
	ServiceAccount pulumi.StringPtrOutput `pulumi:"serviceAccount"`
	// Id of the Sharedflow to be deployed.
	SharedflowId pulumi.StringOutput                         `pulumi:"sharedflowId"`
	Timeouts     ApigeeSharedflowDeploymentTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewApigeeSharedflowDeployment registers a new resource with the given unique name, arguments, and options.
func NewApigeeSharedflowDeployment(ctx *pulumi.Context,
	name string, args *ApigeeSharedflowDeploymentArgs, opts ...pulumi.ResourceOption) (*ApigeeSharedflowDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Revision == nil {
		return nil, errors.New("invalid value for required argument 'Revision'")
	}
	if args.SharedflowId == nil {
		return nil, errors.New("invalid value for required argument 'SharedflowId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApigeeSharedflowDeployment
	err = ctx.RegisterPackageResource("google:index/apigeeSharedflowDeployment:ApigeeSharedflowDeployment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeSharedflowDeployment gets an existing ApigeeSharedflowDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeSharedflowDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeSharedflowDeploymentState, opts ...pulumi.ResourceOption) (*ApigeeSharedflowDeployment, error) {
	var resource ApigeeSharedflowDeployment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apigeeSharedflowDeployment:ApigeeSharedflowDeployment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeSharedflowDeployment resources.
type apigeeSharedflowDeploymentState struct {
	ApigeeSharedflowDeploymentId *string `pulumi:"apigeeSharedflowDeploymentId"`
	// The resource ID of the environment.
	Environment *string `pulumi:"environment"`
	// The Apigee Organization associated with the Apigee instance
	OrgId *string `pulumi:"orgId"`
	// Revision of the Sharedflow to be deployed.
	Revision *string `pulumi:"revision"`
	// The service account represents the identity of the deployed proxy, and determines what permissions it has. The format
	// must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Id of the Sharedflow to be deployed.
	SharedflowId *string                             `pulumi:"sharedflowId"`
	Timeouts     *ApigeeSharedflowDeploymentTimeouts `pulumi:"timeouts"`
}

type ApigeeSharedflowDeploymentState struct {
	ApigeeSharedflowDeploymentId pulumi.StringPtrInput
	// The resource ID of the environment.
	Environment pulumi.StringPtrInput
	// The Apigee Organization associated with the Apigee instance
	OrgId pulumi.StringPtrInput
	// Revision of the Sharedflow to be deployed.
	Revision pulumi.StringPtrInput
	// The service account represents the identity of the deployed proxy, and determines what permissions it has. The format
	// must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
	ServiceAccount pulumi.StringPtrInput
	// Id of the Sharedflow to be deployed.
	SharedflowId pulumi.StringPtrInput
	Timeouts     ApigeeSharedflowDeploymentTimeoutsPtrInput
}

func (ApigeeSharedflowDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeSharedflowDeploymentState)(nil)).Elem()
}

type apigeeSharedflowDeploymentArgs struct {
	ApigeeSharedflowDeploymentId *string `pulumi:"apigeeSharedflowDeploymentId"`
	// The resource ID of the environment.
	Environment string `pulumi:"environment"`
	// The Apigee Organization associated with the Apigee instance
	OrgId string `pulumi:"orgId"`
	// Revision of the Sharedflow to be deployed.
	Revision string `pulumi:"revision"`
	// The service account represents the identity of the deployed proxy, and determines what permissions it has. The format
	// must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Id of the Sharedflow to be deployed.
	SharedflowId string                              `pulumi:"sharedflowId"`
	Timeouts     *ApigeeSharedflowDeploymentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApigeeSharedflowDeployment resource.
type ApigeeSharedflowDeploymentArgs struct {
	ApigeeSharedflowDeploymentId pulumi.StringPtrInput
	// The resource ID of the environment.
	Environment pulumi.StringInput
	// The Apigee Organization associated with the Apigee instance
	OrgId pulumi.StringInput
	// Revision of the Sharedflow to be deployed.
	Revision pulumi.StringInput
	// The service account represents the identity of the deployed proxy, and determines what permissions it has. The format
	// must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
	ServiceAccount pulumi.StringPtrInput
	// Id of the Sharedflow to be deployed.
	SharedflowId pulumi.StringInput
	Timeouts     ApigeeSharedflowDeploymentTimeoutsPtrInput
}

func (ApigeeSharedflowDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeSharedflowDeploymentArgs)(nil)).Elem()
}

type ApigeeSharedflowDeploymentInput interface {
	pulumi.Input

	ToApigeeSharedflowDeploymentOutput() ApigeeSharedflowDeploymentOutput
	ToApigeeSharedflowDeploymentOutputWithContext(ctx context.Context) ApigeeSharedflowDeploymentOutput
}

func (*ApigeeSharedflowDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeSharedflowDeployment)(nil)).Elem()
}

func (i *ApigeeSharedflowDeployment) ToApigeeSharedflowDeploymentOutput() ApigeeSharedflowDeploymentOutput {
	return i.ToApigeeSharedflowDeploymentOutputWithContext(context.Background())
}

func (i *ApigeeSharedflowDeployment) ToApigeeSharedflowDeploymentOutputWithContext(ctx context.Context) ApigeeSharedflowDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeSharedflowDeploymentOutput)
}

type ApigeeSharedflowDeploymentOutput struct{ *pulumi.OutputState }

func (ApigeeSharedflowDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeSharedflowDeployment)(nil)).Elem()
}

func (o ApigeeSharedflowDeploymentOutput) ToApigeeSharedflowDeploymentOutput() ApigeeSharedflowDeploymentOutput {
	return o
}

func (o ApigeeSharedflowDeploymentOutput) ToApigeeSharedflowDeploymentOutputWithContext(ctx context.Context) ApigeeSharedflowDeploymentOutput {
	return o
}

func (o ApigeeSharedflowDeploymentOutput) ApigeeSharedflowDeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSharedflowDeployment) pulumi.StringOutput { return v.ApigeeSharedflowDeploymentId }).(pulumi.StringOutput)
}

// The resource ID of the environment.
func (o ApigeeSharedflowDeploymentOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSharedflowDeployment) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The Apigee Organization associated with the Apigee instance
func (o ApigeeSharedflowDeploymentOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSharedflowDeployment) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Revision of the Sharedflow to be deployed.
func (o ApigeeSharedflowDeploymentOutput) Revision() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSharedflowDeployment) pulumi.StringOutput { return v.Revision }).(pulumi.StringOutput)
}

// The service account represents the identity of the deployed proxy, and determines what permissions it has. The format
// must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
func (o ApigeeSharedflowDeploymentOutput) ServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigeeSharedflowDeployment) pulumi.StringPtrOutput { return v.ServiceAccount }).(pulumi.StringPtrOutput)
}

// Id of the Sharedflow to be deployed.
func (o ApigeeSharedflowDeploymentOutput) SharedflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSharedflowDeployment) pulumi.StringOutput { return v.SharedflowId }).(pulumi.StringOutput)
}

func (o ApigeeSharedflowDeploymentOutput) Timeouts() ApigeeSharedflowDeploymentTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApigeeSharedflowDeployment) ApigeeSharedflowDeploymentTimeoutsPtrOutput { return v.Timeouts }).(ApigeeSharedflowDeploymentTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeSharedflowDeploymentInput)(nil)).Elem(), &ApigeeSharedflowDeployment{})
	pulumi.RegisterOutputType(ApigeeSharedflowDeploymentOutput{})
}
