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

type DataCatalogTaxonomyIamBinding struct {
	pulumi.CustomResourceState

	Condition                       DataCatalogTaxonomyIamBindingConditionPtrOutput `pulumi:"condition"`
	DataCatalogTaxonomyIamBindingId pulumi.StringOutput                             `pulumi:"dataCatalogTaxonomyIamBindingId"`
	Etag                            pulumi.StringOutput                             `pulumi:"etag"`
	Members                         pulumi.StringArrayOutput                        `pulumi:"members"`
	Project                         pulumi.StringOutput                             `pulumi:"project"`
	Region                          pulumi.StringOutput                             `pulumi:"region"`
	Role                            pulumi.StringOutput                             `pulumi:"role"`
	Taxonomy                        pulumi.StringOutput                             `pulumi:"taxonomy"`
}

// NewDataCatalogTaxonomyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataCatalogTaxonomyIamBinding(ctx *pulumi.Context,
	name string, args *DataCatalogTaxonomyIamBindingArgs, opts ...pulumi.ResourceOption) (*DataCatalogTaxonomyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Taxonomy == nil {
		return nil, errors.New("invalid value for required argument 'Taxonomy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataCatalogTaxonomyIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/dataCatalogTaxonomyIamBinding:DataCatalogTaxonomyIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCatalogTaxonomyIamBinding gets an existing DataCatalogTaxonomyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCatalogTaxonomyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCatalogTaxonomyIamBindingState, opts ...pulumi.ResourceOption) (*DataCatalogTaxonomyIamBinding, error) {
	var resource DataCatalogTaxonomyIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataCatalogTaxonomyIamBinding:DataCatalogTaxonomyIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCatalogTaxonomyIamBinding resources.
type dataCatalogTaxonomyIamBindingState struct {
	Condition                       *DataCatalogTaxonomyIamBindingCondition `pulumi:"condition"`
	DataCatalogTaxonomyIamBindingId *string                                 `pulumi:"dataCatalogTaxonomyIamBindingId"`
	Etag                            *string                                 `pulumi:"etag"`
	Members                         []string                                `pulumi:"members"`
	Project                         *string                                 `pulumi:"project"`
	Region                          *string                                 `pulumi:"region"`
	Role                            *string                                 `pulumi:"role"`
	Taxonomy                        *string                                 `pulumi:"taxonomy"`
}

type DataCatalogTaxonomyIamBindingState struct {
	Condition                       DataCatalogTaxonomyIamBindingConditionPtrInput
	DataCatalogTaxonomyIamBindingId pulumi.StringPtrInput
	Etag                            pulumi.StringPtrInput
	Members                         pulumi.StringArrayInput
	Project                         pulumi.StringPtrInput
	Region                          pulumi.StringPtrInput
	Role                            pulumi.StringPtrInput
	Taxonomy                        pulumi.StringPtrInput
}

func (DataCatalogTaxonomyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogTaxonomyIamBindingState)(nil)).Elem()
}

type dataCatalogTaxonomyIamBindingArgs struct {
	Condition                       *DataCatalogTaxonomyIamBindingCondition `pulumi:"condition"`
	DataCatalogTaxonomyIamBindingId *string                                 `pulumi:"dataCatalogTaxonomyIamBindingId"`
	Members                         []string                                `pulumi:"members"`
	Project                         *string                                 `pulumi:"project"`
	Region                          *string                                 `pulumi:"region"`
	Role                            string                                  `pulumi:"role"`
	Taxonomy                        string                                  `pulumi:"taxonomy"`
}

// The set of arguments for constructing a DataCatalogTaxonomyIamBinding resource.
type DataCatalogTaxonomyIamBindingArgs struct {
	Condition                       DataCatalogTaxonomyIamBindingConditionPtrInput
	DataCatalogTaxonomyIamBindingId pulumi.StringPtrInput
	Members                         pulumi.StringArrayInput
	Project                         pulumi.StringPtrInput
	Region                          pulumi.StringPtrInput
	Role                            pulumi.StringInput
	Taxonomy                        pulumi.StringInput
}

func (DataCatalogTaxonomyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogTaxonomyIamBindingArgs)(nil)).Elem()
}

type DataCatalogTaxonomyIamBindingInput interface {
	pulumi.Input

	ToDataCatalogTaxonomyIamBindingOutput() DataCatalogTaxonomyIamBindingOutput
	ToDataCatalogTaxonomyIamBindingOutputWithContext(ctx context.Context) DataCatalogTaxonomyIamBindingOutput
}

func (*DataCatalogTaxonomyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogTaxonomyIamBinding)(nil)).Elem()
}

func (i *DataCatalogTaxonomyIamBinding) ToDataCatalogTaxonomyIamBindingOutput() DataCatalogTaxonomyIamBindingOutput {
	return i.ToDataCatalogTaxonomyIamBindingOutputWithContext(context.Background())
}

func (i *DataCatalogTaxonomyIamBinding) ToDataCatalogTaxonomyIamBindingOutputWithContext(ctx context.Context) DataCatalogTaxonomyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCatalogTaxonomyIamBindingOutput)
}

type DataCatalogTaxonomyIamBindingOutput struct{ *pulumi.OutputState }

func (DataCatalogTaxonomyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogTaxonomyIamBinding)(nil)).Elem()
}

func (o DataCatalogTaxonomyIamBindingOutput) ToDataCatalogTaxonomyIamBindingOutput() DataCatalogTaxonomyIamBindingOutput {
	return o
}

func (o DataCatalogTaxonomyIamBindingOutput) ToDataCatalogTaxonomyIamBindingOutputWithContext(ctx context.Context) DataCatalogTaxonomyIamBindingOutput {
	return o
}

func (o DataCatalogTaxonomyIamBindingOutput) Condition() DataCatalogTaxonomyIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) DataCatalogTaxonomyIamBindingConditionPtrOutput {
		return v.Condition
	}).(DataCatalogTaxonomyIamBindingConditionPtrOutput)
}

func (o DataCatalogTaxonomyIamBindingOutput) DataCatalogTaxonomyIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) pulumi.StringOutput { return v.DataCatalogTaxonomyIamBindingId }).(pulumi.StringOutput)
}

func (o DataCatalogTaxonomyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataCatalogTaxonomyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o DataCatalogTaxonomyIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataCatalogTaxonomyIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o DataCatalogTaxonomyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o DataCatalogTaxonomyIamBindingOutput) Taxonomy() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTaxonomyIamBinding) pulumi.StringOutput { return v.Taxonomy }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCatalogTaxonomyIamBindingInput)(nil)).Elem(), &DataCatalogTaxonomyIamBinding{})
	pulumi.RegisterOutputType(DataCatalogTaxonomyIamBindingOutput{})
}
