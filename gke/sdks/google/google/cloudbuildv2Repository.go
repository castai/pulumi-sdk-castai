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

type Cloudbuildv2Repository struct {
	pulumi.CustomResourceState

	// Allows clients to store small amounts of arbitrary data. **Note**: This field is non-authoritative, and will only manage
	// the annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the
	// annotations present on the resource.
	Annotations              pulumi.StringMapOutput `pulumi:"annotations"`
	Cloudbuildv2RepositoryId pulumi.StringOutput    `pulumi:"cloudbuildv2RepositoryId"`
	// Output only. Server assigned timestamp for when the connection was created.
	CreateTime           pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveAnnotations pulumi.StringMapOutput `pulumi:"effectiveAnnotations"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
	// requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the repository.
	Name pulumi.StringOutput `pulumi:"name"`
	// The connection for the resource
	ParentConnection pulumi.StringOutput `pulumi:"parentConnection"`
	Project          pulumi.StringOutput `pulumi:"project"`
	// Required. Git Clone HTTPS URI.
	RemoteUri pulumi.StringOutput                     `pulumi:"remoteUri"`
	Timeouts  Cloudbuildv2RepositoryTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. Server assigned timestamp for when the connection was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCloudbuildv2Repository registers a new resource with the given unique name, arguments, and options.
func NewCloudbuildv2Repository(ctx *pulumi.Context,
	name string, args *Cloudbuildv2RepositoryArgs, opts ...pulumi.ResourceOption) (*Cloudbuildv2Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentConnection == nil {
		return nil, errors.New("invalid value for required argument 'ParentConnection'")
	}
	if args.RemoteUri == nil {
		return nil, errors.New("invalid value for required argument 'RemoteUri'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource Cloudbuildv2Repository
	err = ctx.RegisterPackageResource("google:index/cloudbuildv2Repository:Cloudbuildv2Repository", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudbuildv2Repository gets an existing Cloudbuildv2Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudbuildv2Repository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Cloudbuildv2RepositoryState, opts ...pulumi.ResourceOption) (*Cloudbuildv2Repository, error) {
	var resource Cloudbuildv2Repository
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/cloudbuildv2Repository:Cloudbuildv2Repository", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cloudbuildv2Repository resources.
type cloudbuildv2RepositoryState struct {
	// Allows clients to store small amounts of arbitrary data. **Note**: This field is non-authoritative, and will only manage
	// the annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the
	// annotations present on the resource.
	Annotations              map[string]string `pulumi:"annotations"`
	Cloudbuildv2RepositoryId *string           `pulumi:"cloudbuildv2RepositoryId"`
	// Output only. Server assigned timestamp for when the connection was created.
	CreateTime           *string           `pulumi:"createTime"`
	EffectiveAnnotations map[string]string `pulumi:"effectiveAnnotations"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
	// requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Name of the repository.
	Name *string `pulumi:"name"`
	// The connection for the resource
	ParentConnection *string `pulumi:"parentConnection"`
	Project          *string `pulumi:"project"`
	// Required. Git Clone HTTPS URI.
	RemoteUri *string                         `pulumi:"remoteUri"`
	Timeouts  *Cloudbuildv2RepositoryTimeouts `pulumi:"timeouts"`
	// Output only. Server assigned timestamp for when the connection was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type Cloudbuildv2RepositoryState struct {
	// Allows clients to store small amounts of arbitrary data. **Note**: This field is non-authoritative, and will only manage
	// the annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the
	// annotations present on the resource.
	Annotations              pulumi.StringMapInput
	Cloudbuildv2RepositoryId pulumi.StringPtrInput
	// Output only. Server assigned timestamp for when the connection was created.
	CreateTime           pulumi.StringPtrInput
	EffectiveAnnotations pulumi.StringMapInput
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
	// requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Name of the repository.
	Name pulumi.StringPtrInput
	// The connection for the resource
	ParentConnection pulumi.StringPtrInput
	Project          pulumi.StringPtrInput
	// Required. Git Clone HTTPS URI.
	RemoteUri pulumi.StringPtrInput
	Timeouts  Cloudbuildv2RepositoryTimeoutsPtrInput
	// Output only. Server assigned timestamp for when the connection was updated.
	UpdateTime pulumi.StringPtrInput
}

func (Cloudbuildv2RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudbuildv2RepositoryState)(nil)).Elem()
}

type cloudbuildv2RepositoryArgs struct {
	// Allows clients to store small amounts of arbitrary data. **Note**: This field is non-authoritative, and will only manage
	// the annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the
	// annotations present on the resource.
	Annotations              map[string]string `pulumi:"annotations"`
	Cloudbuildv2RepositoryId *string           `pulumi:"cloudbuildv2RepositoryId"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Name of the repository.
	Name *string `pulumi:"name"`
	// The connection for the resource
	ParentConnection string  `pulumi:"parentConnection"`
	Project          *string `pulumi:"project"`
	// Required. Git Clone HTTPS URI.
	RemoteUri string                          `pulumi:"remoteUri"`
	Timeouts  *Cloudbuildv2RepositoryTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a Cloudbuildv2Repository resource.
type Cloudbuildv2RepositoryArgs struct {
	// Allows clients to store small amounts of arbitrary data. **Note**: This field is non-authoritative, and will only manage
	// the annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the
	// annotations present on the resource.
	Annotations              pulumi.StringMapInput
	Cloudbuildv2RepositoryId pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Name of the repository.
	Name pulumi.StringPtrInput
	// The connection for the resource
	ParentConnection pulumi.StringInput
	Project          pulumi.StringPtrInput
	// Required. Git Clone HTTPS URI.
	RemoteUri pulumi.StringInput
	Timeouts  Cloudbuildv2RepositoryTimeoutsPtrInput
}

func (Cloudbuildv2RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudbuildv2RepositoryArgs)(nil)).Elem()
}

type Cloudbuildv2RepositoryInput interface {
	pulumi.Input

	ToCloudbuildv2RepositoryOutput() Cloudbuildv2RepositoryOutput
	ToCloudbuildv2RepositoryOutputWithContext(ctx context.Context) Cloudbuildv2RepositoryOutput
}

func (*Cloudbuildv2Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloudbuildv2Repository)(nil)).Elem()
}

func (i *Cloudbuildv2Repository) ToCloudbuildv2RepositoryOutput() Cloudbuildv2RepositoryOutput {
	return i.ToCloudbuildv2RepositoryOutputWithContext(context.Background())
}

func (i *Cloudbuildv2Repository) ToCloudbuildv2RepositoryOutputWithContext(ctx context.Context) Cloudbuildv2RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Cloudbuildv2RepositoryOutput)
}

type Cloudbuildv2RepositoryOutput struct{ *pulumi.OutputState }

func (Cloudbuildv2RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloudbuildv2Repository)(nil)).Elem()
}

func (o Cloudbuildv2RepositoryOutput) ToCloudbuildv2RepositoryOutput() Cloudbuildv2RepositoryOutput {
	return o
}

func (o Cloudbuildv2RepositoryOutput) ToCloudbuildv2RepositoryOutputWithContext(ctx context.Context) Cloudbuildv2RepositoryOutput {
	return o
}

// Allows clients to store small amounts of arbitrary data. **Note**: This field is non-authoritative, and will only manage
// the annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the
// annotations present on the resource.
func (o Cloudbuildv2RepositoryOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

func (o Cloudbuildv2RepositoryOutput) Cloudbuildv2RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.Cloudbuildv2RepositoryId }).(pulumi.StringOutput)
}

// Output only. Server assigned timestamp for when the connection was created.
func (o Cloudbuildv2RepositoryOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o Cloudbuildv2RepositoryOutput) EffectiveAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringMapOutput { return v.EffectiveAnnotations }).(pulumi.StringMapOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete
// requests to ensure the client has an up-to-date value before proceeding.
func (o Cloudbuildv2RepositoryOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location for the resource
func (o Cloudbuildv2RepositoryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the repository.
func (o Cloudbuildv2RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The connection for the resource
func (o Cloudbuildv2RepositoryOutput) ParentConnection() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.ParentConnection }).(pulumi.StringOutput)
}

func (o Cloudbuildv2RepositoryOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. Git Clone HTTPS URI.
func (o Cloudbuildv2RepositoryOutput) RemoteUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.RemoteUri }).(pulumi.StringOutput)
}

func (o Cloudbuildv2RepositoryOutput) Timeouts() Cloudbuildv2RepositoryTimeoutsPtrOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) Cloudbuildv2RepositoryTimeoutsPtrOutput { return v.Timeouts }).(Cloudbuildv2RepositoryTimeoutsPtrOutput)
}

// Output only. Server assigned timestamp for when the connection was updated.
func (o Cloudbuildv2RepositoryOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2Repository) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Cloudbuildv2RepositoryInput)(nil)).Elem(), &Cloudbuildv2Repository{})
	pulumi.RegisterOutputType(Cloudbuildv2RepositoryOutput{})
}
