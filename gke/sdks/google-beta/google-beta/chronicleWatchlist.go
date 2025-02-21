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

type ChronicleWatchlist struct {
	pulumi.CustomResourceState

	ChronicleWatchlistId pulumi.StringOutput `pulumi:"chronicleWatchlistId"`
	// Output only. Time the watchlist was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Description of the watchlist.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters
	// (https://google.aip.dev/148).
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Count of different types of entities in the watchlist.
	EntityCounts ChronicleWatchlistEntityCountArrayOutput `pulumi:"entityCounts"`
	// Mechanism to populate entities in the watchlist.
	EntityPopulationMechanism ChronicleWatchlistEntityPopulationMechanismOutput `pulumi:"entityPopulationMechanism"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringOutput `pulumi:"location"`
	// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
	MultiplyingFactor pulumi.Float64PtrOutput `pulumi:"multiplyingFactor"`
	// Identifier. Resource name of the watchlist. This unique identifier is generated using values provided for the URL
	// parameters. Format: projects/{project}/locations/{location}/instances/{instance}/watchlists/{watchlist}
	Name     pulumi.StringOutput                 `pulumi:"name"`
	Project  pulumi.StringOutput                 `pulumi:"project"`
	Timeouts ChronicleWatchlistTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. Time the watchlist was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Optional. The ID to use for the watchlist, which will become the final component of the watchlist's resource name. This
	// value should be 4-63 characters, and valid characters are /a-z-/.
	WatchlistId pulumi.StringOutput `pulumi:"watchlistId"`
	// A collection of user preferences for watchlist UI configuration.
	WatchlistUserPreferences ChronicleWatchlistWatchlistUserPreferencesPtrOutput `pulumi:"watchlistUserPreferences"`
}

// NewChronicleWatchlist registers a new resource with the given unique name, arguments, and options.
func NewChronicleWatchlist(ctx *pulumi.Context,
	name string, args *ChronicleWatchlistArgs, opts ...pulumi.ResourceOption) (*ChronicleWatchlist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.EntityPopulationMechanism == nil {
		return nil, errors.New("invalid value for required argument 'EntityPopulationMechanism'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ChronicleWatchlist
	err = ctx.RegisterPackageResource("google-beta:index/chronicleWatchlist:ChronicleWatchlist", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChronicleWatchlist gets an existing ChronicleWatchlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChronicleWatchlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChronicleWatchlistState, opts ...pulumi.ResourceOption) (*ChronicleWatchlist, error) {
	var resource ChronicleWatchlist
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/chronicleWatchlist:ChronicleWatchlist", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChronicleWatchlist resources.
type chronicleWatchlistState struct {
	ChronicleWatchlistId *string `pulumi:"chronicleWatchlistId"`
	// Output only. Time the watchlist was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. Description of the watchlist.
	Description *string `pulumi:"description"`
	// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters
	// (https://google.aip.dev/148).
	DisplayName *string `pulumi:"displayName"`
	// Count of different types of entities in the watchlist.
	EntityCounts []ChronicleWatchlistEntityCount `pulumi:"entityCounts"`
	// Mechanism to populate entities in the watchlist.
	EntityPopulationMechanism *ChronicleWatchlistEntityPopulationMechanism `pulumi:"entityPopulationMechanism"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance *string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location *string `pulumi:"location"`
	// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
	MultiplyingFactor *float64 `pulumi:"multiplyingFactor"`
	// Identifier. Resource name of the watchlist. This unique identifier is generated using values provided for the URL
	// parameters. Format: projects/{project}/locations/{location}/instances/{instance}/watchlists/{watchlist}
	Name     *string                     `pulumi:"name"`
	Project  *string                     `pulumi:"project"`
	Timeouts *ChronicleWatchlistTimeouts `pulumi:"timeouts"`
	// Output only. Time the watchlist was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Optional. The ID to use for the watchlist, which will become the final component of the watchlist's resource name. This
	// value should be 4-63 characters, and valid characters are /a-z-/.
	WatchlistId *string `pulumi:"watchlistId"`
	// A collection of user preferences for watchlist UI configuration.
	WatchlistUserPreferences *ChronicleWatchlistWatchlistUserPreferences `pulumi:"watchlistUserPreferences"`
}

type ChronicleWatchlistState struct {
	ChronicleWatchlistId pulumi.StringPtrInput
	// Output only. Time the watchlist was created.
	CreateTime pulumi.StringPtrInput
	// Optional. Description of the watchlist.
	Description pulumi.StringPtrInput
	// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters
	// (https://google.aip.dev/148).
	DisplayName pulumi.StringPtrInput
	// Count of different types of entities in the watchlist.
	EntityCounts ChronicleWatchlistEntityCountArrayInput
	// Mechanism to populate entities in the watchlist.
	EntityPopulationMechanism ChronicleWatchlistEntityPopulationMechanismPtrInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringPtrInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringPtrInput
	// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
	MultiplyingFactor pulumi.Float64PtrInput
	// Identifier. Resource name of the watchlist. This unique identifier is generated using values provided for the URL
	// parameters. Format: projects/{project}/locations/{location}/instances/{instance}/watchlists/{watchlist}
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ChronicleWatchlistTimeoutsPtrInput
	// Output only. Time the watchlist was last updated.
	UpdateTime pulumi.StringPtrInput
	// Optional. The ID to use for the watchlist, which will become the final component of the watchlist's resource name. This
	// value should be 4-63 characters, and valid characters are /a-z-/.
	WatchlistId pulumi.StringPtrInput
	// A collection of user preferences for watchlist UI configuration.
	WatchlistUserPreferences ChronicleWatchlistWatchlistUserPreferencesPtrInput
}

func (ChronicleWatchlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleWatchlistState)(nil)).Elem()
}

type chronicleWatchlistArgs struct {
	ChronicleWatchlistId *string `pulumi:"chronicleWatchlistId"`
	// Optional. Description of the watchlist.
	Description *string `pulumi:"description"`
	// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters
	// (https://google.aip.dev/148).
	DisplayName string `pulumi:"displayName"`
	// Mechanism to populate entities in the watchlist.
	EntityPopulationMechanism ChronicleWatchlistEntityPopulationMechanism `pulumi:"entityPopulationMechanism"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location string `pulumi:"location"`
	// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
	MultiplyingFactor *float64                    `pulumi:"multiplyingFactor"`
	Project           *string                     `pulumi:"project"`
	Timeouts          *ChronicleWatchlistTimeouts `pulumi:"timeouts"`
	// Optional. The ID to use for the watchlist, which will become the final component of the watchlist's resource name. This
	// value should be 4-63 characters, and valid characters are /a-z-/.
	WatchlistId *string `pulumi:"watchlistId"`
	// A collection of user preferences for watchlist UI configuration.
	WatchlistUserPreferences *ChronicleWatchlistWatchlistUserPreferences `pulumi:"watchlistUserPreferences"`
}

// The set of arguments for constructing a ChronicleWatchlist resource.
type ChronicleWatchlistArgs struct {
	ChronicleWatchlistId pulumi.StringPtrInput
	// Optional. Description of the watchlist.
	Description pulumi.StringPtrInput
	// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters
	// (https://google.aip.dev/148).
	DisplayName pulumi.StringInput
	// Mechanism to populate entities in the watchlist.
	EntityPopulationMechanism ChronicleWatchlistEntityPopulationMechanismInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringInput
	// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
	MultiplyingFactor pulumi.Float64PtrInput
	Project           pulumi.StringPtrInput
	Timeouts          ChronicleWatchlistTimeoutsPtrInput
	// Optional. The ID to use for the watchlist, which will become the final component of the watchlist's resource name. This
	// value should be 4-63 characters, and valid characters are /a-z-/.
	WatchlistId pulumi.StringPtrInput
	// A collection of user preferences for watchlist UI configuration.
	WatchlistUserPreferences ChronicleWatchlistWatchlistUserPreferencesPtrInput
}

func (ChronicleWatchlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleWatchlistArgs)(nil)).Elem()
}

type ChronicleWatchlistInput interface {
	pulumi.Input

	ToChronicleWatchlistOutput() ChronicleWatchlistOutput
	ToChronicleWatchlistOutputWithContext(ctx context.Context) ChronicleWatchlistOutput
}

func (*ChronicleWatchlist) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleWatchlist)(nil)).Elem()
}

func (i *ChronicleWatchlist) ToChronicleWatchlistOutput() ChronicleWatchlistOutput {
	return i.ToChronicleWatchlistOutputWithContext(context.Background())
}

func (i *ChronicleWatchlist) ToChronicleWatchlistOutputWithContext(ctx context.Context) ChronicleWatchlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChronicleWatchlistOutput)
}

type ChronicleWatchlistOutput struct{ *pulumi.OutputState }

func (ChronicleWatchlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleWatchlist)(nil)).Elem()
}

func (o ChronicleWatchlistOutput) ToChronicleWatchlistOutput() ChronicleWatchlistOutput {
	return o
}

func (o ChronicleWatchlistOutput) ToChronicleWatchlistOutputWithContext(ctx context.Context) ChronicleWatchlistOutput {
	return o
}

func (o ChronicleWatchlistOutput) ChronicleWatchlistId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.ChronicleWatchlistId }).(pulumi.StringOutput)
}

// Output only. Time the watchlist was created.
func (o ChronicleWatchlistOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the watchlist.
func (o ChronicleWatchlistOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters
// (https://google.aip.dev/148).
func (o ChronicleWatchlistOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Count of different types of entities in the watchlist.
func (o ChronicleWatchlistOutput) EntityCounts() ChronicleWatchlistEntityCountArrayOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) ChronicleWatchlistEntityCountArrayOutput { return v.EntityCounts }).(ChronicleWatchlistEntityCountArrayOutput)
}

// Mechanism to populate entities in the watchlist.
func (o ChronicleWatchlistOutput) EntityPopulationMechanism() ChronicleWatchlistEntityPopulationMechanismOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) ChronicleWatchlistEntityPopulationMechanismOutput {
		return v.EntityPopulationMechanism
	}).(ChronicleWatchlistEntityPopulationMechanismOutput)
}

// The unique identifier for the Chronicle instance, which is the same as the customer ID.
func (o ChronicleWatchlistOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
// "europe-west2".
func (o ChronicleWatchlistOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
func (o ChronicleWatchlistOutput) MultiplyingFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.Float64PtrOutput { return v.MultiplyingFactor }).(pulumi.Float64PtrOutput)
}

// Identifier. Resource name of the watchlist. This unique identifier is generated using values provided for the URL
// parameters. Format: projects/{project}/locations/{location}/instances/{instance}/watchlists/{watchlist}
func (o ChronicleWatchlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ChronicleWatchlistOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ChronicleWatchlistOutput) Timeouts() ChronicleWatchlistTimeoutsPtrOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) ChronicleWatchlistTimeoutsPtrOutput { return v.Timeouts }).(ChronicleWatchlistTimeoutsPtrOutput)
}

// Output only. Time the watchlist was last updated.
func (o ChronicleWatchlistOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. The ID to use for the watchlist, which will become the final component of the watchlist's resource name. This
// value should be 4-63 characters, and valid characters are /a-z-/.
func (o ChronicleWatchlistOutput) WatchlistId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) pulumi.StringOutput { return v.WatchlistId }).(pulumi.StringOutput)
}

// A collection of user preferences for watchlist UI configuration.
func (o ChronicleWatchlistOutput) WatchlistUserPreferences() ChronicleWatchlistWatchlistUserPreferencesPtrOutput {
	return o.ApplyT(func(v *ChronicleWatchlist) ChronicleWatchlistWatchlistUserPreferencesPtrOutput {
		return v.WatchlistUserPreferences
	}).(ChronicleWatchlistWatchlistUserPreferencesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChronicleWatchlistInput)(nil)).Elem(), &ChronicleWatchlist{})
	pulumi.RegisterOutputType(ChronicleWatchlistOutput{})
}
