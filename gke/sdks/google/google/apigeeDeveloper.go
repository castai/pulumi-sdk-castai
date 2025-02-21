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

type ApigeeDeveloper struct {
	pulumi.CustomResourceState

	ApigeeDeveloperId pulumi.StringOutput `pulumi:"apigeeDeveloperId"`
	// Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes ApigeeDeveloperAttributeArrayOutput `pulumi:"attributes"`
	// Time at which the developer was created in milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the
	// email address has to be in lowercase only..
	Email pulumi.StringOutput `pulumi:"email"`
	// First name of the developer.
	FirstName pulumi.StringOutput `pulumi:"firstName"`
	// Time at which the developer was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Last name of the developer.
	LastName pulumi.StringOutput `pulumi:"lastName"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Name of the Apigee organization in which the developer resides.
	OrganizatioName pulumi.StringOutput `pulumi:"organizatioName"`
	// Status of the developer. Valid values are active and inactive.
	Status   pulumi.StringOutput              `pulumi:"status"`
	Timeouts ApigeeDeveloperTimeoutsPtrOutput `pulumi:"timeouts"`
	// User name of the developer. Not used by Apigee hybrid.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewApigeeDeveloper registers a new resource with the given unique name, arguments, and options.
func NewApigeeDeveloper(ctx *pulumi.Context,
	name string, args *ApigeeDeveloperArgs, opts ...pulumi.ResourceOption) (*ApigeeDeveloper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApigeeDeveloper
	err = ctx.RegisterPackageResource("google:index/apigeeDeveloper:ApigeeDeveloper", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeDeveloper gets an existing ApigeeDeveloper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeDeveloper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeDeveloperState, opts ...pulumi.ResourceOption) (*ApigeeDeveloper, error) {
	var resource ApigeeDeveloper
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apigeeDeveloper:ApigeeDeveloper", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeDeveloper resources.
type apigeeDeveloperState struct {
	ApigeeDeveloperId *string `pulumi:"apigeeDeveloperId"`
	// Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes []ApigeeDeveloperAttribute `pulumi:"attributes"`
	// Time at which the developer was created in milliseconds since epoch.
	CreatedAt *string `pulumi:"createdAt"`
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the
	// email address has to be in lowercase only..
	Email *string `pulumi:"email"`
	// First name of the developer.
	FirstName *string `pulumi:"firstName"`
	// Time at which the developer was last modified in milliseconds since epoch.
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// Last name of the developer.
	LastName *string `pulumi:"lastName"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	OrgId *string `pulumi:"orgId"`
	// Name of the Apigee organization in which the developer resides.
	OrganizatioName *string `pulumi:"organizatioName"`
	// Status of the developer. Valid values are active and inactive.
	Status   *string                  `pulumi:"status"`
	Timeouts *ApigeeDeveloperTimeouts `pulumi:"timeouts"`
	// User name of the developer. Not used by Apigee hybrid.
	UserName *string `pulumi:"userName"`
}

type ApigeeDeveloperState struct {
	ApigeeDeveloperId pulumi.StringPtrInput
	// Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes ApigeeDeveloperAttributeArrayInput
	// Time at which the developer was created in milliseconds since epoch.
	CreatedAt pulumi.StringPtrInput
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the
	// email address has to be in lowercase only..
	Email pulumi.StringPtrInput
	// First name of the developer.
	FirstName pulumi.StringPtrInput
	// Time at which the developer was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringPtrInput
	// Last name of the developer.
	LastName pulumi.StringPtrInput
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	OrgId pulumi.StringPtrInput
	// Name of the Apigee organization in which the developer resides.
	OrganizatioName pulumi.StringPtrInput
	// Status of the developer. Valid values are active and inactive.
	Status   pulumi.StringPtrInput
	Timeouts ApigeeDeveloperTimeoutsPtrInput
	// User name of the developer. Not used by Apigee hybrid.
	UserName pulumi.StringPtrInput
}

func (ApigeeDeveloperState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeDeveloperState)(nil)).Elem()
}

type apigeeDeveloperArgs struct {
	ApigeeDeveloperId *string `pulumi:"apigeeDeveloperId"`
	// Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes []ApigeeDeveloperAttribute `pulumi:"attributes"`
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the
	// email address has to be in lowercase only..
	Email string `pulumi:"email"`
	// First name of the developer.
	FirstName string `pulumi:"firstName"`
	// Last name of the developer.
	LastName string `pulumi:"lastName"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	OrgId    string                   `pulumi:"orgId"`
	Timeouts *ApigeeDeveloperTimeouts `pulumi:"timeouts"`
	// User name of the developer. Not used by Apigee hybrid.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a ApigeeDeveloper resource.
type ApigeeDeveloperArgs struct {
	ApigeeDeveloperId pulumi.StringPtrInput
	// Developer attributes (name/value pairs). The custom attribute limit is 18.
	Attributes ApigeeDeveloperAttributeArrayInput
	// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the
	// email address has to be in lowercase only..
	Email pulumi.StringInput
	// First name of the developer.
	FirstName pulumi.StringInput
	// Last name of the developer.
	LastName pulumi.StringInput
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	OrgId    pulumi.StringInput
	Timeouts ApigeeDeveloperTimeoutsPtrInput
	// User name of the developer. Not used by Apigee hybrid.
	UserName pulumi.StringInput
}

func (ApigeeDeveloperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeDeveloperArgs)(nil)).Elem()
}

type ApigeeDeveloperInput interface {
	pulumi.Input

	ToApigeeDeveloperOutput() ApigeeDeveloperOutput
	ToApigeeDeveloperOutputWithContext(ctx context.Context) ApigeeDeveloperOutput
}

func (*ApigeeDeveloper) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeDeveloper)(nil)).Elem()
}

func (i *ApigeeDeveloper) ToApigeeDeveloperOutput() ApigeeDeveloperOutput {
	return i.ToApigeeDeveloperOutputWithContext(context.Background())
}

func (i *ApigeeDeveloper) ToApigeeDeveloperOutputWithContext(ctx context.Context) ApigeeDeveloperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeDeveloperOutput)
}

type ApigeeDeveloperOutput struct{ *pulumi.OutputState }

func (ApigeeDeveloperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeDeveloper)(nil)).Elem()
}

func (o ApigeeDeveloperOutput) ToApigeeDeveloperOutput() ApigeeDeveloperOutput {
	return o
}

func (o ApigeeDeveloperOutput) ToApigeeDeveloperOutputWithContext(ctx context.Context) ApigeeDeveloperOutput {
	return o
}

func (o ApigeeDeveloperOutput) ApigeeDeveloperId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.ApigeeDeveloperId }).(pulumi.StringOutput)
}

// Developer attributes (name/value pairs). The custom attribute limit is 18.
func (o ApigeeDeveloperOutput) Attributes() ApigeeDeveloperAttributeArrayOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) ApigeeDeveloperAttributeArrayOutput { return v.Attributes }).(ApigeeDeveloperAttributeArrayOutput)
}

// Time at which the developer was created in milliseconds since epoch.
func (o ApigeeDeveloperOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the
// email address has to be in lowercase only..
func (o ApigeeDeveloperOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// First name of the developer.
func (o ApigeeDeveloperOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.FirstName }).(pulumi.StringOutput)
}

// Time at which the developer was last modified in milliseconds since epoch.
func (o ApigeeDeveloperOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.LastModifiedAt }).(pulumi.StringOutput)
}

// Last name of the developer.
func (o ApigeeDeveloperOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.LastName }).(pulumi.StringOutput)
}

// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
func (o ApigeeDeveloperOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Name of the Apigee organization in which the developer resides.
func (o ApigeeDeveloperOutput) OrganizatioName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.OrganizatioName }).(pulumi.StringOutput)
}

// Status of the developer. Valid values are active and inactive.
func (o ApigeeDeveloperOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ApigeeDeveloperOutput) Timeouts() ApigeeDeveloperTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) ApigeeDeveloperTimeoutsPtrOutput { return v.Timeouts }).(ApigeeDeveloperTimeoutsPtrOutput)
}

// User name of the developer. Not used by Apigee hybrid.
func (o ApigeeDeveloperOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeDeveloper) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeDeveloperInput)(nil)).Elem(), &ApigeeDeveloper{})
	pulumi.RegisterOutputType(ApigeeDeveloperOutput{})
}
