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

type BigtableAppProfile struct {
	pulumi.CustomResourceState

	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId         pulumi.StringOutput `pulumi:"appProfileId"`
	BigtableAppProfileId pulumi.StringOutput `pulumi:"bigtableAppProfileId"`
	// Specifies that this app profile is intended for read-only usage via the Data Boost feature.
	DataBoostIsolationReadOnly BigtableAppProfileDataBoostIsolationReadOnlyPtrOutput `pulumi:"dataBoostIsolationReadOnly"`
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrOutput `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrOutput `pulumi:"instance"`
	// The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
	// clusters are eligible.
	MultiClusterRoutingClusterIds pulumi.StringArrayOutput `pulumi:"multiClusterRoutingClusterIds"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrOutput `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Must be used with multi-cluster routing. If true, then this app profile will use row affinity sticky routing. With row
	// affinity, Bigtable will route single row key requests based on the row key, rather than randomly. Instead, each row key
	// will be assigned to a cluster by Cloud Bigtable, and will stick to that cluster. Choosing this option improves
	// read-your-writes consistency for most requests under most circumstances, without sacrificing availability. Consistency
	// is not guaranteed, as requests may still fail over between clusters in the event of errors or latency.
	RowAffinity pulumi.BoolPtrOutput `pulumi:"rowAffinity"`
	// Use a single-cluster routing policy.
	SingleClusterRouting BigtableAppProfileSingleClusterRoutingPtrOutput `pulumi:"singleClusterRouting"`
	// The standard options used for isolating this app profile's traffic from other use cases.
	StandardIsolation BigtableAppProfileStandardIsolationPtrOutput `pulumi:"standardIsolation"`
	Timeouts          BigtableAppProfileTimeoutsPtrOutput          `pulumi:"timeouts"`
}

// NewBigtableAppProfile registers a new resource with the given unique name, arguments, and options.
func NewBigtableAppProfile(ctx *pulumi.Context,
	name string, args *BigtableAppProfileArgs, opts ...pulumi.ResourceOption) (*BigtableAppProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppProfileId == nil {
		return nil, errors.New("invalid value for required argument 'AppProfileId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigtableAppProfile
	err = ctx.RegisterPackageResource("google:index/bigtableAppProfile:BigtableAppProfile", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigtableAppProfile gets an existing BigtableAppProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigtableAppProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigtableAppProfileState, opts ...pulumi.ResourceOption) (*BigtableAppProfile, error) {
	var resource BigtableAppProfile
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/bigtableAppProfile:BigtableAppProfile", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigtableAppProfile resources.
type bigtableAppProfileState struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId         *string `pulumi:"appProfileId"`
	BigtableAppProfileId *string `pulumi:"bigtableAppProfileId"`
	// Specifies that this app profile is intended for read-only usage via the Data Boost feature.
	DataBoostIsolationReadOnly *BigtableAppProfileDataBoostIsolationReadOnly `pulumi:"dataBoostIsolationReadOnly"`
	// Long form description of the use case for this app profile.
	Description *string `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance *string `pulumi:"instance"`
	// The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
	// clusters are eligible.
	MultiClusterRoutingClusterIds []string `pulumi:"multiClusterRoutingClusterIds"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny *bool `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Must be used with multi-cluster routing. If true, then this app profile will use row affinity sticky routing. With row
	// affinity, Bigtable will route single row key requests based on the row key, rather than randomly. Instead, each row key
	// will be assigned to a cluster by Cloud Bigtable, and will stick to that cluster. Choosing this option improves
	// read-your-writes consistency for most requests under most circumstances, without sacrificing availability. Consistency
	// is not guaranteed, as requests may still fail over between clusters in the event of errors or latency.
	RowAffinity *bool `pulumi:"rowAffinity"`
	// Use a single-cluster routing policy.
	SingleClusterRouting *BigtableAppProfileSingleClusterRouting `pulumi:"singleClusterRouting"`
	// The standard options used for isolating this app profile's traffic from other use cases.
	StandardIsolation *BigtableAppProfileStandardIsolation `pulumi:"standardIsolation"`
	Timeouts          *BigtableAppProfileTimeouts          `pulumi:"timeouts"`
}

type BigtableAppProfileState struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId         pulumi.StringPtrInput
	BigtableAppProfileId pulumi.StringPtrInput
	// Specifies that this app profile is intended for read-only usage via the Data Boost feature.
	DataBoostIsolationReadOnly BigtableAppProfileDataBoostIsolationReadOnlyPtrInput
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrInput
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrInput
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrInput
	// The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
	// clusters are eligible.
	MultiClusterRoutingClusterIds pulumi.StringArrayInput
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrInput
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Must be used with multi-cluster routing. If true, then this app profile will use row affinity sticky routing. With row
	// affinity, Bigtable will route single row key requests based on the row key, rather than randomly. Instead, each row key
	// will be assigned to a cluster by Cloud Bigtable, and will stick to that cluster. Choosing this option improves
	// read-your-writes consistency for most requests under most circumstances, without sacrificing availability. Consistency
	// is not guaranteed, as requests may still fail over between clusters in the event of errors or latency.
	RowAffinity pulumi.BoolPtrInput
	// Use a single-cluster routing policy.
	SingleClusterRouting BigtableAppProfileSingleClusterRoutingPtrInput
	// The standard options used for isolating this app profile's traffic from other use cases.
	StandardIsolation BigtableAppProfileStandardIsolationPtrInput
	Timeouts          BigtableAppProfileTimeoutsPtrInput
}

func (BigtableAppProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigtableAppProfileState)(nil)).Elem()
}

type bigtableAppProfileArgs struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId         string  `pulumi:"appProfileId"`
	BigtableAppProfileId *string `pulumi:"bigtableAppProfileId"`
	// Specifies that this app profile is intended for read-only usage via the Data Boost feature.
	DataBoostIsolationReadOnly *BigtableAppProfileDataBoostIsolationReadOnly `pulumi:"dataBoostIsolationReadOnly"`
	// Long form description of the use case for this app profile.
	Description *string `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance *string `pulumi:"instance"`
	// The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
	// clusters are eligible.
	MultiClusterRoutingClusterIds []string `pulumi:"multiClusterRoutingClusterIds"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny *bool   `pulumi:"multiClusterRoutingUseAny"`
	Project                   *string `pulumi:"project"`
	// Must be used with multi-cluster routing. If true, then this app profile will use row affinity sticky routing. With row
	// affinity, Bigtable will route single row key requests based on the row key, rather than randomly. Instead, each row key
	// will be assigned to a cluster by Cloud Bigtable, and will stick to that cluster. Choosing this option improves
	// read-your-writes consistency for most requests under most circumstances, without sacrificing availability. Consistency
	// is not guaranteed, as requests may still fail over between clusters in the event of errors or latency.
	RowAffinity *bool `pulumi:"rowAffinity"`
	// Use a single-cluster routing policy.
	SingleClusterRouting *BigtableAppProfileSingleClusterRouting `pulumi:"singleClusterRouting"`
	// The standard options used for isolating this app profile's traffic from other use cases.
	StandardIsolation *BigtableAppProfileStandardIsolation `pulumi:"standardIsolation"`
	Timeouts          *BigtableAppProfileTimeouts          `pulumi:"timeouts"`
}

// The set of arguments for constructing a BigtableAppProfile resource.
type BigtableAppProfileArgs struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId         pulumi.StringInput
	BigtableAppProfileId pulumi.StringPtrInput
	// Specifies that this app profile is intended for read-only usage via the Data Boost feature.
	DataBoostIsolationReadOnly BigtableAppProfileDataBoostIsolationReadOnlyPtrInput
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrInput
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrInput
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrInput
	// The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
	// clusters are eligible.
	MultiClusterRoutingClusterIds pulumi.StringArrayInput
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrInput
	Project                   pulumi.StringPtrInput
	// Must be used with multi-cluster routing. If true, then this app profile will use row affinity sticky routing. With row
	// affinity, Bigtable will route single row key requests based on the row key, rather than randomly. Instead, each row key
	// will be assigned to a cluster by Cloud Bigtable, and will stick to that cluster. Choosing this option improves
	// read-your-writes consistency for most requests under most circumstances, without sacrificing availability. Consistency
	// is not guaranteed, as requests may still fail over between clusters in the event of errors or latency.
	RowAffinity pulumi.BoolPtrInput
	// Use a single-cluster routing policy.
	SingleClusterRouting BigtableAppProfileSingleClusterRoutingPtrInput
	// The standard options used for isolating this app profile's traffic from other use cases.
	StandardIsolation BigtableAppProfileStandardIsolationPtrInput
	Timeouts          BigtableAppProfileTimeoutsPtrInput
}

func (BigtableAppProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigtableAppProfileArgs)(nil)).Elem()
}

type BigtableAppProfileInput interface {
	pulumi.Input

	ToBigtableAppProfileOutput() BigtableAppProfileOutput
	ToBigtableAppProfileOutputWithContext(ctx context.Context) BigtableAppProfileOutput
}

func (*BigtableAppProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**BigtableAppProfile)(nil)).Elem()
}

func (i *BigtableAppProfile) ToBigtableAppProfileOutput() BigtableAppProfileOutput {
	return i.ToBigtableAppProfileOutputWithContext(context.Background())
}

func (i *BigtableAppProfile) ToBigtableAppProfileOutputWithContext(ctx context.Context) BigtableAppProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigtableAppProfileOutput)
}

type BigtableAppProfileOutput struct{ *pulumi.OutputState }

func (BigtableAppProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigtableAppProfile)(nil)).Elem()
}

func (o BigtableAppProfileOutput) ToBigtableAppProfileOutput() BigtableAppProfileOutput {
	return o
}

func (o BigtableAppProfileOutput) ToBigtableAppProfileOutputWithContext(ctx context.Context) BigtableAppProfileOutput {
	return o
}

// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
func (o BigtableAppProfileOutput) AppProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.StringOutput { return v.AppProfileId }).(pulumi.StringOutput)
}

func (o BigtableAppProfileOutput) BigtableAppProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.StringOutput { return v.BigtableAppProfileId }).(pulumi.StringOutput)
}

// Specifies that this app profile is intended for read-only usage via the Data Boost feature.
func (o BigtableAppProfileOutput) DataBoostIsolationReadOnly() BigtableAppProfileDataBoostIsolationReadOnlyPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) BigtableAppProfileDataBoostIsolationReadOnlyPtrOutput {
		return v.DataBoostIsolationReadOnly
	}).(BigtableAppProfileDataBoostIsolationReadOnlyPtrOutput)
}

// Long form description of the use case for this app profile.
func (o BigtableAppProfileOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If true, ignore safety checks when deleting/updating the app profile.
func (o BigtableAppProfileOutput) IgnoreWarnings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.BoolPtrOutput { return v.IgnoreWarnings }).(pulumi.BoolPtrOutput)
}

// The name of the instance to create the app profile within.
func (o BigtableAppProfileOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.StringPtrOutput { return v.Instance }).(pulumi.StringPtrOutput)
}

// The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
// clusters are eligible.
func (o BigtableAppProfileOutput) MultiClusterRoutingClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.StringArrayOutput { return v.MultiClusterRoutingClusterIds }).(pulumi.StringArrayOutput)
}

// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
// Choosing this option sacrifices read-your-writes consistency to improve availability.
func (o BigtableAppProfileOutput) MultiClusterRoutingUseAny() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.BoolPtrOutput { return v.MultiClusterRoutingUseAny }).(pulumi.BoolPtrOutput)
}

// The unique name of the requested app profile. Values are of the form
// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
func (o BigtableAppProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BigtableAppProfileOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Must be used with multi-cluster routing. If true, then this app profile will use row affinity sticky routing. With row
// affinity, Bigtable will route single row key requests based on the row key, rather than randomly. Instead, each row key
// will be assigned to a cluster by Cloud Bigtable, and will stick to that cluster. Choosing this option improves
// read-your-writes consistency for most requests under most circumstances, without sacrificing availability. Consistency
// is not guaranteed, as requests may still fail over between clusters in the event of errors or latency.
func (o BigtableAppProfileOutput) RowAffinity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) pulumi.BoolPtrOutput { return v.RowAffinity }).(pulumi.BoolPtrOutput)
}

// Use a single-cluster routing policy.
func (o BigtableAppProfileOutput) SingleClusterRouting() BigtableAppProfileSingleClusterRoutingPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) BigtableAppProfileSingleClusterRoutingPtrOutput {
		return v.SingleClusterRouting
	}).(BigtableAppProfileSingleClusterRoutingPtrOutput)
}

// The standard options used for isolating this app profile's traffic from other use cases.
func (o BigtableAppProfileOutput) StandardIsolation() BigtableAppProfileStandardIsolationPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) BigtableAppProfileStandardIsolationPtrOutput { return v.StandardIsolation }).(BigtableAppProfileStandardIsolationPtrOutput)
}

func (o BigtableAppProfileOutput) Timeouts() BigtableAppProfileTimeoutsPtrOutput {
	return o.ApplyT(func(v *BigtableAppProfile) BigtableAppProfileTimeoutsPtrOutput { return v.Timeouts }).(BigtableAppProfileTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigtableAppProfileInput)(nil)).Elem(), &BigtableAppProfile{})
	pulumi.RegisterOutputType(BigtableAppProfileOutput{})
}
