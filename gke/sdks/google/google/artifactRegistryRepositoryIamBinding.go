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

type ArtifactRegistryRepositoryIamBinding struct {
	pulumi.CustomResourceState

	ArtifactRegistryRepositoryIamBindingId pulumi.StringOutput                                    `pulumi:"artifactRegistryRepositoryIamBindingId"`
	Condition                              ArtifactRegistryRepositoryIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                                   pulumi.StringOutput                                    `pulumi:"etag"`
	Location                               pulumi.StringOutput                                    `pulumi:"location"`
	Members                                pulumi.StringArrayOutput                               `pulumi:"members"`
	Project                                pulumi.StringOutput                                    `pulumi:"project"`
	Repository                             pulumi.StringOutput                                    `pulumi:"repository"`
	Role                                   pulumi.StringOutput                                    `pulumi:"role"`
}

// NewArtifactRegistryRepositoryIamBinding registers a new resource with the given unique name, arguments, and options.
func NewArtifactRegistryRepositoryIamBinding(ctx *pulumi.Context,
	name string, args *ArtifactRegistryRepositoryIamBindingArgs, opts ...pulumi.ResourceOption) (*ArtifactRegistryRepositoryIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ArtifactRegistryRepositoryIamBinding
	err = ctx.RegisterPackageResource("google:index/artifactRegistryRepositoryIamBinding:ArtifactRegistryRepositoryIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactRegistryRepositoryIamBinding gets an existing ArtifactRegistryRepositoryIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactRegistryRepositoryIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactRegistryRepositoryIamBindingState, opts ...pulumi.ResourceOption) (*ArtifactRegistryRepositoryIamBinding, error) {
	var resource ArtifactRegistryRepositoryIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/artifactRegistryRepositoryIamBinding:ArtifactRegistryRepositoryIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactRegistryRepositoryIamBinding resources.
type artifactRegistryRepositoryIamBindingState struct {
	ArtifactRegistryRepositoryIamBindingId *string                                        `pulumi:"artifactRegistryRepositoryIamBindingId"`
	Condition                              *ArtifactRegistryRepositoryIamBindingCondition `pulumi:"condition"`
	Etag                                   *string                                        `pulumi:"etag"`
	Location                               *string                                        `pulumi:"location"`
	Members                                []string                                       `pulumi:"members"`
	Project                                *string                                        `pulumi:"project"`
	Repository                             *string                                        `pulumi:"repository"`
	Role                                   *string                                        `pulumi:"role"`
}

type ArtifactRegistryRepositoryIamBindingState struct {
	ArtifactRegistryRepositoryIamBindingId pulumi.StringPtrInput
	Condition                              ArtifactRegistryRepositoryIamBindingConditionPtrInput
	Etag                                   pulumi.StringPtrInput
	Location                               pulumi.StringPtrInput
	Members                                pulumi.StringArrayInput
	Project                                pulumi.StringPtrInput
	Repository                             pulumi.StringPtrInput
	Role                                   pulumi.StringPtrInput
}

func (ArtifactRegistryRepositoryIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactRegistryRepositoryIamBindingState)(nil)).Elem()
}

type artifactRegistryRepositoryIamBindingArgs struct {
	ArtifactRegistryRepositoryIamBindingId *string                                        `pulumi:"artifactRegistryRepositoryIamBindingId"`
	Condition                              *ArtifactRegistryRepositoryIamBindingCondition `pulumi:"condition"`
	Location                               *string                                        `pulumi:"location"`
	Members                                []string                                       `pulumi:"members"`
	Project                                *string                                        `pulumi:"project"`
	Repository                             string                                         `pulumi:"repository"`
	Role                                   string                                         `pulumi:"role"`
}

// The set of arguments for constructing a ArtifactRegistryRepositoryIamBinding resource.
type ArtifactRegistryRepositoryIamBindingArgs struct {
	ArtifactRegistryRepositoryIamBindingId pulumi.StringPtrInput
	Condition                              ArtifactRegistryRepositoryIamBindingConditionPtrInput
	Location                               pulumi.StringPtrInput
	Members                                pulumi.StringArrayInput
	Project                                pulumi.StringPtrInput
	Repository                             pulumi.StringInput
	Role                                   pulumi.StringInput
}

func (ArtifactRegistryRepositoryIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactRegistryRepositoryIamBindingArgs)(nil)).Elem()
}

type ArtifactRegistryRepositoryIamBindingInput interface {
	pulumi.Input

	ToArtifactRegistryRepositoryIamBindingOutput() ArtifactRegistryRepositoryIamBindingOutput
	ToArtifactRegistryRepositoryIamBindingOutputWithContext(ctx context.Context) ArtifactRegistryRepositoryIamBindingOutput
}

func (*ArtifactRegistryRepositoryIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactRegistryRepositoryIamBinding)(nil)).Elem()
}

func (i *ArtifactRegistryRepositoryIamBinding) ToArtifactRegistryRepositoryIamBindingOutput() ArtifactRegistryRepositoryIamBindingOutput {
	return i.ToArtifactRegistryRepositoryIamBindingOutputWithContext(context.Background())
}

func (i *ArtifactRegistryRepositoryIamBinding) ToArtifactRegistryRepositoryIamBindingOutputWithContext(ctx context.Context) ArtifactRegistryRepositoryIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactRegistryRepositoryIamBindingOutput)
}

type ArtifactRegistryRepositoryIamBindingOutput struct{ *pulumi.OutputState }

func (ArtifactRegistryRepositoryIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactRegistryRepositoryIamBinding)(nil)).Elem()
}

func (o ArtifactRegistryRepositoryIamBindingOutput) ToArtifactRegistryRepositoryIamBindingOutput() ArtifactRegistryRepositoryIamBindingOutput {
	return o
}

func (o ArtifactRegistryRepositoryIamBindingOutput) ToArtifactRegistryRepositoryIamBindingOutputWithContext(ctx context.Context) ArtifactRegistryRepositoryIamBindingOutput {
	return o
}

func (o ArtifactRegistryRepositoryIamBindingOutput) ArtifactRegistryRepositoryIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) pulumi.StringOutput {
		return v.ArtifactRegistryRepositoryIamBindingId
	}).(pulumi.StringOutput)
}

func (o ArtifactRegistryRepositoryIamBindingOutput) Condition() ArtifactRegistryRepositoryIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) ArtifactRegistryRepositoryIamBindingConditionPtrOutput {
		return v.Condition
	}).(ArtifactRegistryRepositoryIamBindingConditionPtrOutput)
}

func (o ArtifactRegistryRepositoryIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ArtifactRegistryRepositoryIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ArtifactRegistryRepositoryIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o ArtifactRegistryRepositoryIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ArtifactRegistryRepositoryIamBindingOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

func (o ArtifactRegistryRepositoryIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactRegistryRepositoryIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactRegistryRepositoryIamBindingInput)(nil)).Elem(), &ArtifactRegistryRepositoryIamBinding{})
	pulumi.RegisterOutputType(ArtifactRegistryRepositoryIamBindingOutput{})
}
