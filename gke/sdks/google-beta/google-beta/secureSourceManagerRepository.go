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

type SecureSourceManagerRepository struct {
	pulumi.CustomResourceState

	// Time the repository was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the repository, which cannot exceed 500 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Initial configurations for the repository.
	InitialConfig SecureSourceManagerRepositoryInitialConfigPtrOutput `pulumi:"initialConfig"`
	// The name of the instance in which the repository is hosted.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The location for the Repository.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the Repository.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The ID for the Repository.
	RepositoryId                    pulumi.StringOutput                            `pulumi:"repositoryId"`
	SecureSourceManagerRepositoryId pulumi.StringOutput                            `pulumi:"secureSourceManagerRepositoryId"`
	Timeouts                        SecureSourceManagerRepositoryTimeoutsPtrOutput `pulumi:"timeouts"`
	// Unique identifier of the repository.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Time the repository was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// URIs for the repository.
	Uris SecureSourceManagerRepositoryUriArrayOutput `pulumi:"uris"`
}

// NewSecureSourceManagerRepository registers a new resource with the given unique name, arguments, and options.
func NewSecureSourceManagerRepository(ctx *pulumi.Context,
	name string, args *SecureSourceManagerRepositoryArgs, opts ...pulumi.ResourceOption) (*SecureSourceManagerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SecureSourceManagerRepository
	err = ctx.RegisterPackageResource("google-beta:index/secureSourceManagerRepository:SecureSourceManagerRepository", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecureSourceManagerRepository gets an existing SecureSourceManagerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecureSourceManagerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecureSourceManagerRepositoryState, opts ...pulumi.ResourceOption) (*SecureSourceManagerRepository, error) {
	var resource SecureSourceManagerRepository
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/secureSourceManagerRepository:SecureSourceManagerRepository", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecureSourceManagerRepository resources.
type secureSourceManagerRepositoryState struct {
	// Time the repository was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Description of the repository, which cannot exceed 500 characters.
	Description *string `pulumi:"description"`
	// Initial configurations for the repository.
	InitialConfig *SecureSourceManagerRepositoryInitialConfig `pulumi:"initialConfig"`
	// The name of the instance in which the repository is hosted.
	Instance *string `pulumi:"instance"`
	// The location for the Repository.
	Location *string `pulumi:"location"`
	// The resource name for the Repository.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The ID for the Repository.
	RepositoryId                    *string                                `pulumi:"repositoryId"`
	SecureSourceManagerRepositoryId *string                                `pulumi:"secureSourceManagerRepositoryId"`
	Timeouts                        *SecureSourceManagerRepositoryTimeouts `pulumi:"timeouts"`
	// Unique identifier of the repository.
	Uid *string `pulumi:"uid"`
	// Time the repository was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
	// URIs for the repository.
	Uris []SecureSourceManagerRepositoryUri `pulumi:"uris"`
}

type SecureSourceManagerRepositoryState struct {
	// Time the repository was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Description of the repository, which cannot exceed 500 characters.
	Description pulumi.StringPtrInput
	// Initial configurations for the repository.
	InitialConfig SecureSourceManagerRepositoryInitialConfigPtrInput
	// The name of the instance in which the repository is hosted.
	Instance pulumi.StringPtrInput
	// The location for the Repository.
	Location pulumi.StringPtrInput
	// The resource name for the Repository.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The ID for the Repository.
	RepositoryId                    pulumi.StringPtrInput
	SecureSourceManagerRepositoryId pulumi.StringPtrInput
	Timeouts                        SecureSourceManagerRepositoryTimeoutsPtrInput
	// Unique identifier of the repository.
	Uid pulumi.StringPtrInput
	// Time the repository was updated in UTC.
	UpdateTime pulumi.StringPtrInput
	// URIs for the repository.
	Uris SecureSourceManagerRepositoryUriArrayInput
}

func (SecureSourceManagerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerRepositoryState)(nil)).Elem()
}

type secureSourceManagerRepositoryArgs struct {
	// Description of the repository, which cannot exceed 500 characters.
	Description *string `pulumi:"description"`
	// Initial configurations for the repository.
	InitialConfig *SecureSourceManagerRepositoryInitialConfig `pulumi:"initialConfig"`
	// The name of the instance in which the repository is hosted.
	Instance string `pulumi:"instance"`
	// The location for the Repository.
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// The ID for the Repository.
	RepositoryId                    string                                 `pulumi:"repositoryId"`
	SecureSourceManagerRepositoryId *string                                `pulumi:"secureSourceManagerRepositoryId"`
	Timeouts                        *SecureSourceManagerRepositoryTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a SecureSourceManagerRepository resource.
type SecureSourceManagerRepositoryArgs struct {
	// Description of the repository, which cannot exceed 500 characters.
	Description pulumi.StringPtrInput
	// Initial configurations for the repository.
	InitialConfig SecureSourceManagerRepositoryInitialConfigPtrInput
	// The name of the instance in which the repository is hosted.
	Instance pulumi.StringInput
	// The location for the Repository.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// The ID for the Repository.
	RepositoryId                    pulumi.StringInput
	SecureSourceManagerRepositoryId pulumi.StringPtrInput
	Timeouts                        SecureSourceManagerRepositoryTimeoutsPtrInput
}

func (SecureSourceManagerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secureSourceManagerRepositoryArgs)(nil)).Elem()
}

type SecureSourceManagerRepositoryInput interface {
	pulumi.Input

	ToSecureSourceManagerRepositoryOutput() SecureSourceManagerRepositoryOutput
	ToSecureSourceManagerRepositoryOutputWithContext(ctx context.Context) SecureSourceManagerRepositoryOutput
}

func (*SecureSourceManagerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerRepository)(nil)).Elem()
}

func (i *SecureSourceManagerRepository) ToSecureSourceManagerRepositoryOutput() SecureSourceManagerRepositoryOutput {
	return i.ToSecureSourceManagerRepositoryOutputWithContext(context.Background())
}

func (i *SecureSourceManagerRepository) ToSecureSourceManagerRepositoryOutputWithContext(ctx context.Context) SecureSourceManagerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureSourceManagerRepositoryOutput)
}

type SecureSourceManagerRepositoryOutput struct{ *pulumi.OutputState }

func (SecureSourceManagerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureSourceManagerRepository)(nil)).Elem()
}

func (o SecureSourceManagerRepositoryOutput) ToSecureSourceManagerRepositoryOutput() SecureSourceManagerRepositoryOutput {
	return o
}

func (o SecureSourceManagerRepositoryOutput) ToSecureSourceManagerRepositoryOutputWithContext(ctx context.Context) SecureSourceManagerRepositoryOutput {
	return o
}

// Time the repository was created in UTC.
func (o SecureSourceManagerRepositoryOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the repository, which cannot exceed 500 characters.
func (o SecureSourceManagerRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Initial configurations for the repository.
func (o SecureSourceManagerRepositoryOutput) InitialConfig() SecureSourceManagerRepositoryInitialConfigPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) SecureSourceManagerRepositoryInitialConfigPtrOutput {
		return v.InitialConfig
	}).(SecureSourceManagerRepositoryInitialConfigPtrOutput)
}

// The name of the instance in which the repository is hosted.
func (o SecureSourceManagerRepositoryOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The location for the Repository.
func (o SecureSourceManagerRepositoryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for the Repository.
func (o SecureSourceManagerRepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The ID for the Repository.
func (o SecureSourceManagerRepositoryOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryOutput) SecureSourceManagerRepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.SecureSourceManagerRepositoryId }).(pulumi.StringOutput)
}

func (o SecureSourceManagerRepositoryOutput) Timeouts() SecureSourceManagerRepositoryTimeoutsPtrOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) SecureSourceManagerRepositoryTimeoutsPtrOutput {
		return v.Timeouts
	}).(SecureSourceManagerRepositoryTimeoutsPtrOutput)
}

// Unique identifier of the repository.
func (o SecureSourceManagerRepositoryOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Time the repository was updated in UTC.
func (o SecureSourceManagerRepositoryOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// URIs for the repository.
func (o SecureSourceManagerRepositoryOutput) Uris() SecureSourceManagerRepositoryUriArrayOutput {
	return o.ApplyT(func(v *SecureSourceManagerRepository) SecureSourceManagerRepositoryUriArrayOutput { return v.Uris }).(SecureSourceManagerRepositoryUriArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecureSourceManagerRepositoryInput)(nil)).Elem(), &SecureSourceManagerRepository{})
	pulumi.RegisterOutputType(SecureSourceManagerRepositoryOutput{})
}
