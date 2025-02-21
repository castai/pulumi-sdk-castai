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

type AccessContextManagerIngressPolicy struct {
	pulumi.CustomResourceState

	AccessContextManagerIngressPolicyId pulumi.StringOutput `pulumi:"accessContextManagerIngressPolicyId"`
	// The name of the Access Policy this resource belongs to.
	AccessPolicyId pulumi.StringOutput `pulumi:"accessPolicyId"`
	// The name of the Service Perimeter to add this resource to.
	IngressPolicyName pulumi.StringOutput `pulumi:"ingressPolicyName"`
	// A GCP resource that is inside of the service perimeter.
	Resource pulumi.StringOutput                                `pulumi:"resource"`
	Timeouts AccessContextManagerIngressPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewAccessContextManagerIngressPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessContextManagerIngressPolicy(ctx *pulumi.Context,
	name string, args *AccessContextManagerIngressPolicyArgs, opts ...pulumi.ResourceOption) (*AccessContextManagerIngressPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IngressPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'IngressPolicyName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource AccessContextManagerIngressPolicy
	err = ctx.RegisterPackageResource("google-beta:index/accessContextManagerIngressPolicy:AccessContextManagerIngressPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessContextManagerIngressPolicy gets an existing AccessContextManagerIngressPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessContextManagerIngressPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessContextManagerIngressPolicyState, opts ...pulumi.ResourceOption) (*AccessContextManagerIngressPolicy, error) {
	var resource AccessContextManagerIngressPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/accessContextManagerIngressPolicy:AccessContextManagerIngressPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessContextManagerIngressPolicy resources.
type accessContextManagerIngressPolicyState struct {
	AccessContextManagerIngressPolicyId *string `pulumi:"accessContextManagerIngressPolicyId"`
	// The name of the Access Policy this resource belongs to.
	AccessPolicyId *string `pulumi:"accessPolicyId"`
	// The name of the Service Perimeter to add this resource to.
	IngressPolicyName *string `pulumi:"ingressPolicyName"`
	// A GCP resource that is inside of the service perimeter.
	Resource *string                                    `pulumi:"resource"`
	Timeouts *AccessContextManagerIngressPolicyTimeouts `pulumi:"timeouts"`
}

type AccessContextManagerIngressPolicyState struct {
	AccessContextManagerIngressPolicyId pulumi.StringPtrInput
	// The name of the Access Policy this resource belongs to.
	AccessPolicyId pulumi.StringPtrInput
	// The name of the Service Perimeter to add this resource to.
	IngressPolicyName pulumi.StringPtrInput
	// A GCP resource that is inside of the service perimeter.
	Resource pulumi.StringPtrInput
	Timeouts AccessContextManagerIngressPolicyTimeoutsPtrInput
}

func (AccessContextManagerIngressPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerIngressPolicyState)(nil)).Elem()
}

type accessContextManagerIngressPolicyArgs struct {
	AccessContextManagerIngressPolicyId *string `pulumi:"accessContextManagerIngressPolicyId"`
	// The name of the Service Perimeter to add this resource to.
	IngressPolicyName string `pulumi:"ingressPolicyName"`
	// A GCP resource that is inside of the service perimeter.
	Resource string                                     `pulumi:"resource"`
	Timeouts *AccessContextManagerIngressPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a AccessContextManagerIngressPolicy resource.
type AccessContextManagerIngressPolicyArgs struct {
	AccessContextManagerIngressPolicyId pulumi.StringPtrInput
	// The name of the Service Perimeter to add this resource to.
	IngressPolicyName pulumi.StringInput
	// A GCP resource that is inside of the service perimeter.
	Resource pulumi.StringInput
	Timeouts AccessContextManagerIngressPolicyTimeoutsPtrInput
}

func (AccessContextManagerIngressPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerIngressPolicyArgs)(nil)).Elem()
}

type AccessContextManagerIngressPolicyInput interface {
	pulumi.Input

	ToAccessContextManagerIngressPolicyOutput() AccessContextManagerIngressPolicyOutput
	ToAccessContextManagerIngressPolicyOutputWithContext(ctx context.Context) AccessContextManagerIngressPolicyOutput
}

func (*AccessContextManagerIngressPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerIngressPolicy)(nil)).Elem()
}

func (i *AccessContextManagerIngressPolicy) ToAccessContextManagerIngressPolicyOutput() AccessContextManagerIngressPolicyOutput {
	return i.ToAccessContextManagerIngressPolicyOutputWithContext(context.Background())
}

func (i *AccessContextManagerIngressPolicy) ToAccessContextManagerIngressPolicyOutputWithContext(ctx context.Context) AccessContextManagerIngressPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessContextManagerIngressPolicyOutput)
}

type AccessContextManagerIngressPolicyOutput struct{ *pulumi.OutputState }

func (AccessContextManagerIngressPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerIngressPolicy)(nil)).Elem()
}

func (o AccessContextManagerIngressPolicyOutput) ToAccessContextManagerIngressPolicyOutput() AccessContextManagerIngressPolicyOutput {
	return o
}

func (o AccessContextManagerIngressPolicyOutput) ToAccessContextManagerIngressPolicyOutputWithContext(ctx context.Context) AccessContextManagerIngressPolicyOutput {
	return o
}

func (o AccessContextManagerIngressPolicyOutput) AccessContextManagerIngressPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerIngressPolicy) pulumi.StringOutput {
		return v.AccessContextManagerIngressPolicyId
	}).(pulumi.StringOutput)
}

// The name of the Access Policy this resource belongs to.
func (o AccessContextManagerIngressPolicyOutput) AccessPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerIngressPolicy) pulumi.StringOutput { return v.AccessPolicyId }).(pulumi.StringOutput)
}

// The name of the Service Perimeter to add this resource to.
func (o AccessContextManagerIngressPolicyOutput) IngressPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerIngressPolicy) pulumi.StringOutput { return v.IngressPolicyName }).(pulumi.StringOutput)
}

// A GCP resource that is inside of the service perimeter.
func (o AccessContextManagerIngressPolicyOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerIngressPolicy) pulumi.StringOutput { return v.Resource }).(pulumi.StringOutput)
}

func (o AccessContextManagerIngressPolicyOutput) Timeouts() AccessContextManagerIngressPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerIngressPolicy) AccessContextManagerIngressPolicyTimeoutsPtrOutput {
		return v.Timeouts
	}).(AccessContextManagerIngressPolicyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessContextManagerIngressPolicyInput)(nil)).Elem(), &AccessContextManagerIngressPolicy{})
	pulumi.RegisterOutputType(AccessContextManagerIngressPolicyOutput{})
}
