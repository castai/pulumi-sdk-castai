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

type AlloydbUser struct {
	pulumi.CustomResourceState

	AlloydbUserId pulumi.StringOutput `pulumi:"alloydbUserId"`
	// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// List of database roles this database user has.
	DatabaseRoles pulumi.StringArrayOutput `pulumi:"databaseRoles"`
	// Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for this database user.
	Password pulumi.StringPtrOutput       `pulumi:"password"`
	Timeouts AlloydbUserTimeoutsPtrOutput `pulumi:"timeouts"`
	// The database role name of the user.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]
	UserType pulumi.StringOutput `pulumi:"userType"`
}

// NewAlloydbUser registers a new resource with the given unique name, arguments, and options.
func NewAlloydbUser(ctx *pulumi.Context,
	name string, args *AlloydbUserArgs, opts ...pulumi.ResourceOption) (*AlloydbUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	if args.UserType == nil {
		return nil, errors.New("invalid value for required argument 'UserType'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource AlloydbUser
	err = ctx.RegisterPackageResource("google-beta:index/alloydbUser:AlloydbUser", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlloydbUser gets an existing AlloydbUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlloydbUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlloydbUserState, opts ...pulumi.ResourceOption) (*AlloydbUser, error) {
	var resource AlloydbUser
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/alloydbUser:AlloydbUser", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlloydbUser resources.
type alloydbUserState struct {
	AlloydbUserId *string `pulumi:"alloydbUserId"`
	// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	Cluster *string `pulumi:"cluster"`
	// List of database roles this database user has.
	DatabaseRoles []string `pulumi:"databaseRoles"`
	// Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.
	Name *string `pulumi:"name"`
	// Password for this database user.
	Password *string              `pulumi:"password"`
	Timeouts *AlloydbUserTimeouts `pulumi:"timeouts"`
	// The database role name of the user.
	UserId *string `pulumi:"userId"`
	// The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]
	UserType *string `pulumi:"userType"`
}

type AlloydbUserState struct {
	AlloydbUserId pulumi.StringPtrInput
	// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	Cluster pulumi.StringPtrInput
	// List of database roles this database user has.
	DatabaseRoles pulumi.StringArrayInput
	// Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.
	Name pulumi.StringPtrInput
	// Password for this database user.
	Password pulumi.StringPtrInput
	Timeouts AlloydbUserTimeoutsPtrInput
	// The database role name of the user.
	UserId pulumi.StringPtrInput
	// The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]
	UserType pulumi.StringPtrInput
}

func (AlloydbUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*alloydbUserState)(nil)).Elem()
}

type alloydbUserArgs struct {
	AlloydbUserId *string `pulumi:"alloydbUserId"`
	// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	Cluster string `pulumi:"cluster"`
	// List of database roles this database user has.
	DatabaseRoles []string `pulumi:"databaseRoles"`
	// Password for this database user.
	Password *string              `pulumi:"password"`
	Timeouts *AlloydbUserTimeouts `pulumi:"timeouts"`
	// The database role name of the user.
	UserId string `pulumi:"userId"`
	// The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]
	UserType string `pulumi:"userType"`
}

// The set of arguments for constructing a AlloydbUser resource.
type AlloydbUserArgs struct {
	AlloydbUserId pulumi.StringPtrInput
	// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	Cluster pulumi.StringInput
	// List of database roles this database user has.
	DatabaseRoles pulumi.StringArrayInput
	// Password for this database user.
	Password pulumi.StringPtrInput
	Timeouts AlloydbUserTimeoutsPtrInput
	// The database role name of the user.
	UserId pulumi.StringInput
	// The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]
	UserType pulumi.StringInput
}

func (AlloydbUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alloydbUserArgs)(nil)).Elem()
}

type AlloydbUserInput interface {
	pulumi.Input

	ToAlloydbUserOutput() AlloydbUserOutput
	ToAlloydbUserOutputWithContext(ctx context.Context) AlloydbUserOutput
}

func (*AlloydbUser) ElementType() reflect.Type {
	return reflect.TypeOf((**AlloydbUser)(nil)).Elem()
}

func (i *AlloydbUser) ToAlloydbUserOutput() AlloydbUserOutput {
	return i.ToAlloydbUserOutputWithContext(context.Background())
}

func (i *AlloydbUser) ToAlloydbUserOutputWithContext(ctx context.Context) AlloydbUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlloydbUserOutput)
}

type AlloydbUserOutput struct{ *pulumi.OutputState }

func (AlloydbUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlloydbUser)(nil)).Elem()
}

func (o AlloydbUserOutput) ToAlloydbUserOutput() AlloydbUserOutput {
	return o
}

func (o AlloydbUserOutput) ToAlloydbUserOutputWithContext(ctx context.Context) AlloydbUserOutput {
	return o
}

func (o AlloydbUserOutput) AlloydbUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbUser) pulumi.StringOutput { return v.AlloydbUserId }).(pulumi.StringOutput)
}

// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'
func (o AlloydbUserOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbUser) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

// List of database roles this database user has.
func (o AlloydbUserOutput) DatabaseRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlloydbUser) pulumi.StringArrayOutput { return v.DatabaseRoles }).(pulumi.StringArrayOutput)
}

// Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.
func (o AlloydbUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for this database user.
func (o AlloydbUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlloydbUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AlloydbUserOutput) Timeouts() AlloydbUserTimeoutsPtrOutput {
	return o.ApplyT(func(v *AlloydbUser) AlloydbUserTimeoutsPtrOutput { return v.Timeouts }).(AlloydbUserTimeoutsPtrOutput)
}

// The database role name of the user.
func (o AlloydbUserOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbUser) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]
func (o AlloydbUserOutput) UserType() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbUser) pulumi.StringOutput { return v.UserType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlloydbUserInput)(nil)).Elem(), &AlloydbUser{})
	pulumi.RegisterOutputType(AlloydbUserOutput{})
}
