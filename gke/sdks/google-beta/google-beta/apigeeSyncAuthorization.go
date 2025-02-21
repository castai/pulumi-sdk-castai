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

type ApigeeSyncAuthorization struct {
	pulumi.CustomResourceState

	ApigeeSyncAuthorizationId pulumi.StringOutput `pulumi:"apigeeSyncAuthorizationId"`
	// Entity tag (ETag) used for optimistic concurrency control as a way to help prevent simultaneous updates from overwriting
	// each other. Used internally during updates.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Array of service accounts to grant access to control plane resources, each specified using the following format:
	// 'serviceAccount:service-account-name'. The 'service-account-name' is formatted like an email address. For example:
	// my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com You might specify multiple service
	// accounts, for example, if you have multiple environments and wish to assign a unique service account to each one. The
	// service accounts must have **Apigee Synchronizer Manager** role. See also [Create service
	// accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
	Identities pulumi.StringArrayOutput `pulumi:"identities"`
	// Name of the Apigee organization.
	Name     pulumi.StringOutput                      `pulumi:"name"`
	Timeouts ApigeeSyncAuthorizationTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewApigeeSyncAuthorization registers a new resource with the given unique name, arguments, and options.
func NewApigeeSyncAuthorization(ctx *pulumi.Context,
	name string, args *ApigeeSyncAuthorizationArgs, opts ...pulumi.ResourceOption) (*ApigeeSyncAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identities == nil {
		return nil, errors.New("invalid value for required argument 'Identities'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApigeeSyncAuthorization
	err = ctx.RegisterPackageResource("google-beta:index/apigeeSyncAuthorization:ApigeeSyncAuthorization", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeSyncAuthorization gets an existing ApigeeSyncAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeSyncAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeSyncAuthorizationState, opts ...pulumi.ResourceOption) (*ApigeeSyncAuthorization, error) {
	var resource ApigeeSyncAuthorization
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/apigeeSyncAuthorization:ApigeeSyncAuthorization", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeSyncAuthorization resources.
type apigeeSyncAuthorizationState struct {
	ApigeeSyncAuthorizationId *string `pulumi:"apigeeSyncAuthorizationId"`
	// Entity tag (ETag) used for optimistic concurrency control as a way to help prevent simultaneous updates from overwriting
	// each other. Used internally during updates.
	Etag *string `pulumi:"etag"`
	// Array of service accounts to grant access to control plane resources, each specified using the following format:
	// 'serviceAccount:service-account-name'. The 'service-account-name' is formatted like an email address. For example:
	// my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com You might specify multiple service
	// accounts, for example, if you have multiple environments and wish to assign a unique service account to each one. The
	// service accounts must have **Apigee Synchronizer Manager** role. See also [Create service
	// accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
	Identities []string `pulumi:"identities"`
	// Name of the Apigee organization.
	Name     *string                          `pulumi:"name"`
	Timeouts *ApigeeSyncAuthorizationTimeouts `pulumi:"timeouts"`
}

type ApigeeSyncAuthorizationState struct {
	ApigeeSyncAuthorizationId pulumi.StringPtrInput
	// Entity tag (ETag) used for optimistic concurrency control as a way to help prevent simultaneous updates from overwriting
	// each other. Used internally during updates.
	Etag pulumi.StringPtrInput
	// Array of service accounts to grant access to control plane resources, each specified using the following format:
	// 'serviceAccount:service-account-name'. The 'service-account-name' is formatted like an email address. For example:
	// my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com You might specify multiple service
	// accounts, for example, if you have multiple environments and wish to assign a unique service account to each one. The
	// service accounts must have **Apigee Synchronizer Manager** role. See also [Create service
	// accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
	Identities pulumi.StringArrayInput
	// Name of the Apigee organization.
	Name     pulumi.StringPtrInput
	Timeouts ApigeeSyncAuthorizationTimeoutsPtrInput
}

func (ApigeeSyncAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeSyncAuthorizationState)(nil)).Elem()
}

type apigeeSyncAuthorizationArgs struct {
	ApigeeSyncAuthorizationId *string `pulumi:"apigeeSyncAuthorizationId"`
	// Array of service accounts to grant access to control plane resources, each specified using the following format:
	// 'serviceAccount:service-account-name'. The 'service-account-name' is formatted like an email address. For example:
	// my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com You might specify multiple service
	// accounts, for example, if you have multiple environments and wish to assign a unique service account to each one. The
	// service accounts must have **Apigee Synchronizer Manager** role. See also [Create service
	// accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
	Identities []string `pulumi:"identities"`
	// Name of the Apigee organization.
	Name     *string                          `pulumi:"name"`
	Timeouts *ApigeeSyncAuthorizationTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApigeeSyncAuthorization resource.
type ApigeeSyncAuthorizationArgs struct {
	ApigeeSyncAuthorizationId pulumi.StringPtrInput
	// Array of service accounts to grant access to control plane resources, each specified using the following format:
	// 'serviceAccount:service-account-name'. The 'service-account-name' is formatted like an email address. For example:
	// my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com You might specify multiple service
	// accounts, for example, if you have multiple environments and wish to assign a unique service account to each one. The
	// service accounts must have **Apigee Synchronizer Manager** role. See also [Create service
	// accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
	Identities pulumi.StringArrayInput
	// Name of the Apigee organization.
	Name     pulumi.StringPtrInput
	Timeouts ApigeeSyncAuthorizationTimeoutsPtrInput
}

func (ApigeeSyncAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeSyncAuthorizationArgs)(nil)).Elem()
}

type ApigeeSyncAuthorizationInput interface {
	pulumi.Input

	ToApigeeSyncAuthorizationOutput() ApigeeSyncAuthorizationOutput
	ToApigeeSyncAuthorizationOutputWithContext(ctx context.Context) ApigeeSyncAuthorizationOutput
}

func (*ApigeeSyncAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeSyncAuthorization)(nil)).Elem()
}

func (i *ApigeeSyncAuthorization) ToApigeeSyncAuthorizationOutput() ApigeeSyncAuthorizationOutput {
	return i.ToApigeeSyncAuthorizationOutputWithContext(context.Background())
}

func (i *ApigeeSyncAuthorization) ToApigeeSyncAuthorizationOutputWithContext(ctx context.Context) ApigeeSyncAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeSyncAuthorizationOutput)
}

type ApigeeSyncAuthorizationOutput struct{ *pulumi.OutputState }

func (ApigeeSyncAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeSyncAuthorization)(nil)).Elem()
}

func (o ApigeeSyncAuthorizationOutput) ToApigeeSyncAuthorizationOutput() ApigeeSyncAuthorizationOutput {
	return o
}

func (o ApigeeSyncAuthorizationOutput) ToApigeeSyncAuthorizationOutputWithContext(ctx context.Context) ApigeeSyncAuthorizationOutput {
	return o
}

func (o ApigeeSyncAuthorizationOutput) ApigeeSyncAuthorizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSyncAuthorization) pulumi.StringOutput { return v.ApigeeSyncAuthorizationId }).(pulumi.StringOutput)
}

// Entity tag (ETag) used for optimistic concurrency control as a way to help prevent simultaneous updates from overwriting
// each other. Used internally during updates.
func (o ApigeeSyncAuthorizationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSyncAuthorization) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Array of service accounts to grant access to control plane resources, each specified using the following format:
// 'serviceAccount:service-account-name'. The 'service-account-name' is formatted like an email address. For example:
// my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com You might specify multiple service
// accounts, for example, if you have multiple environments and wish to assign a unique service account to each one. The
// service accounts must have **Apigee Synchronizer Manager** role. See also [Create service
// accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
func (o ApigeeSyncAuthorizationOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigeeSyncAuthorization) pulumi.StringArrayOutput { return v.Identities }).(pulumi.StringArrayOutput)
}

// Name of the Apigee organization.
func (o ApigeeSyncAuthorizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeSyncAuthorization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApigeeSyncAuthorizationOutput) Timeouts() ApigeeSyncAuthorizationTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApigeeSyncAuthorization) ApigeeSyncAuthorizationTimeoutsPtrOutput { return v.Timeouts }).(ApigeeSyncAuthorizationTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeSyncAuthorizationInput)(nil)).Elem(), &ApigeeSyncAuthorization{})
	pulumi.RegisterOutputType(ApigeeSyncAuthorizationOutput{})
}
