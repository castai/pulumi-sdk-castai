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

type ChronicleReferenceList struct {
	pulumi.CustomResourceState

	ChronicleReferenceListId pulumi.StringOutput `pulumi:"chronicleReferenceListId"`
	// Required. A user-provided description of the reference list.
	Description pulumi.StringOutput `pulumi:"description"`
	// Output only. The unique display name of the reference list.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Required. The entries of the reference list. When listed, they are returned in the order that was specified at creation
	// or update. The combined size of the values of the reference list may not exceed 6MB. This is returned only when the view
	// is REFERENCE_LIST_VIEW_FULL.
	Entries ChronicleReferenceListEntryArrayOutput `pulumi:"entries"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. The resource name of the reference list. Format:
	// projects/{project}/locations/{location}/instances/{instance}/referenceLists/{reference_list}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. The ID to use for the reference list. This is also the display name for the reference list. It must satisfy
	// the following requirements: - Starts with letter. - Contains only letters, numbers and underscore. - Has length < 256. -
	// Must be unique.
	ReferenceListId pulumi.StringOutput `pulumi:"referenceListId"`
	// Output only. The timestamp when the reference list was last updated.
	RevisionCreateTime pulumi.StringOutput `pulumi:"revisionCreateTime"`
	// Output only. The count of self-authored rules using the reference list.
	RuleAssociationsCount pulumi.Float64Output `pulumi:"ruleAssociationsCount"`
	// Output only. The resource names for the associated self-authored Rules that use this reference list. This is returned
	// only when the view is REFERENCE_LIST_VIEW_FULL.
	Rules pulumi.StringArrayOutput `pulumi:"rules"`
	// ScopeInfo specifies the scope info of the reference list.
	ScopeInfos ChronicleReferenceListScopeInfoArrayOutput `pulumi:"scopeInfos"`
	// Possible values: REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING REFERENCE_LIST_SYNTAX_TYPE_REGEX
	// REFERENCE_LIST_SYNTAX_TYPE_CIDR
	SyntaxType pulumi.StringOutput                     `pulumi:"syntaxType"`
	Timeouts   ChronicleReferenceListTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewChronicleReferenceList registers a new resource with the given unique name, arguments, and options.
func NewChronicleReferenceList(ctx *pulumi.Context,
	name string, args *ChronicleReferenceListArgs, opts ...pulumi.ResourceOption) (*ChronicleReferenceList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Entries == nil {
		return nil, errors.New("invalid value for required argument 'Entries'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ReferenceListId == nil {
		return nil, errors.New("invalid value for required argument 'ReferenceListId'")
	}
	if args.SyntaxType == nil {
		return nil, errors.New("invalid value for required argument 'SyntaxType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ChronicleReferenceList
	err = ctx.RegisterPackageResource("google-beta:index/chronicleReferenceList:ChronicleReferenceList", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChronicleReferenceList gets an existing ChronicleReferenceList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChronicleReferenceList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChronicleReferenceListState, opts ...pulumi.ResourceOption) (*ChronicleReferenceList, error) {
	var resource ChronicleReferenceList
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/chronicleReferenceList:ChronicleReferenceList", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChronicleReferenceList resources.
type chronicleReferenceListState struct {
	ChronicleReferenceListId *string `pulumi:"chronicleReferenceListId"`
	// Required. A user-provided description of the reference list.
	Description *string `pulumi:"description"`
	// Output only. The unique display name of the reference list.
	DisplayName *string `pulumi:"displayName"`
	// Required. The entries of the reference list. When listed, they are returned in the order that was specified at creation
	// or update. The combined size of the values of the reference list may not exceed 6MB. This is returned only when the view
	// is REFERENCE_LIST_VIEW_FULL.
	Entries []ChronicleReferenceListEntry `pulumi:"entries"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance *string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location *string `pulumi:"location"`
	// Output only. The resource name of the reference list. Format:
	// projects/{project}/locations/{location}/instances/{instance}/referenceLists/{reference_list}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. The ID to use for the reference list. This is also the display name for the reference list. It must satisfy
	// the following requirements: - Starts with letter. - Contains only letters, numbers and underscore. - Has length < 256. -
	// Must be unique.
	ReferenceListId *string `pulumi:"referenceListId"`
	// Output only. The timestamp when the reference list was last updated.
	RevisionCreateTime *string `pulumi:"revisionCreateTime"`
	// Output only. The count of self-authored rules using the reference list.
	RuleAssociationsCount *float64 `pulumi:"ruleAssociationsCount"`
	// Output only. The resource names for the associated self-authored Rules that use this reference list. This is returned
	// only when the view is REFERENCE_LIST_VIEW_FULL.
	Rules []string `pulumi:"rules"`
	// ScopeInfo specifies the scope info of the reference list.
	ScopeInfos []ChronicleReferenceListScopeInfo `pulumi:"scopeInfos"`
	// Possible values: REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING REFERENCE_LIST_SYNTAX_TYPE_REGEX
	// REFERENCE_LIST_SYNTAX_TYPE_CIDR
	SyntaxType *string                         `pulumi:"syntaxType"`
	Timeouts   *ChronicleReferenceListTimeouts `pulumi:"timeouts"`
}

type ChronicleReferenceListState struct {
	ChronicleReferenceListId pulumi.StringPtrInput
	// Required. A user-provided description of the reference list.
	Description pulumi.StringPtrInput
	// Output only. The unique display name of the reference list.
	DisplayName pulumi.StringPtrInput
	// Required. The entries of the reference list. When listed, they are returned in the order that was specified at creation
	// or update. The combined size of the values of the reference list may not exceed 6MB. This is returned only when the view
	// is REFERENCE_LIST_VIEW_FULL.
	Entries ChronicleReferenceListEntryArrayInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringPtrInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringPtrInput
	// Output only. The resource name of the reference list. Format:
	// projects/{project}/locations/{location}/instances/{instance}/referenceLists/{reference_list}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. The ID to use for the reference list. This is also the display name for the reference list. It must satisfy
	// the following requirements: - Starts with letter. - Contains only letters, numbers and underscore. - Has length < 256. -
	// Must be unique.
	ReferenceListId pulumi.StringPtrInput
	// Output only. The timestamp when the reference list was last updated.
	RevisionCreateTime pulumi.StringPtrInput
	// Output only. The count of self-authored rules using the reference list.
	RuleAssociationsCount pulumi.Float64PtrInput
	// Output only. The resource names for the associated self-authored Rules that use this reference list. This is returned
	// only when the view is REFERENCE_LIST_VIEW_FULL.
	Rules pulumi.StringArrayInput
	// ScopeInfo specifies the scope info of the reference list.
	ScopeInfos ChronicleReferenceListScopeInfoArrayInput
	// Possible values: REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING REFERENCE_LIST_SYNTAX_TYPE_REGEX
	// REFERENCE_LIST_SYNTAX_TYPE_CIDR
	SyntaxType pulumi.StringPtrInput
	Timeouts   ChronicleReferenceListTimeoutsPtrInput
}

func (ChronicleReferenceListState) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleReferenceListState)(nil)).Elem()
}

type chronicleReferenceListArgs struct {
	ChronicleReferenceListId *string `pulumi:"chronicleReferenceListId"`
	// Required. A user-provided description of the reference list.
	Description string `pulumi:"description"`
	// Required. The entries of the reference list. When listed, they are returned in the order that was specified at creation
	// or update. The combined size of the values of the reference list may not exceed 6MB. This is returned only when the view
	// is REFERENCE_LIST_VIEW_FULL.
	Entries []ChronicleReferenceListEntry `pulumi:"entries"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Required. The ID to use for the reference list. This is also the display name for the reference list. It must satisfy
	// the following requirements: - Starts with letter. - Contains only letters, numbers and underscore. - Has length < 256. -
	// Must be unique.
	ReferenceListId string `pulumi:"referenceListId"`
	// Possible values: REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING REFERENCE_LIST_SYNTAX_TYPE_REGEX
	// REFERENCE_LIST_SYNTAX_TYPE_CIDR
	SyntaxType string                          `pulumi:"syntaxType"`
	Timeouts   *ChronicleReferenceListTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ChronicleReferenceList resource.
type ChronicleReferenceListArgs struct {
	ChronicleReferenceListId pulumi.StringPtrInput
	// Required. A user-provided description of the reference list.
	Description pulumi.StringInput
	// Required. The entries of the reference list. When listed, they are returned in the order that was specified at creation
	// or update. The combined size of the values of the reference list may not exceed 6MB. This is returned only when the view
	// is REFERENCE_LIST_VIEW_FULL.
	Entries ChronicleReferenceListEntryArrayInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// Required. The ID to use for the reference list. This is also the display name for the reference list. It must satisfy
	// the following requirements: - Starts with letter. - Contains only letters, numbers and underscore. - Has length < 256. -
	// Must be unique.
	ReferenceListId pulumi.StringInput
	// Possible values: REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING REFERENCE_LIST_SYNTAX_TYPE_REGEX
	// REFERENCE_LIST_SYNTAX_TYPE_CIDR
	SyntaxType pulumi.StringInput
	Timeouts   ChronicleReferenceListTimeoutsPtrInput
}

func (ChronicleReferenceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleReferenceListArgs)(nil)).Elem()
}

type ChronicleReferenceListInput interface {
	pulumi.Input

	ToChronicleReferenceListOutput() ChronicleReferenceListOutput
	ToChronicleReferenceListOutputWithContext(ctx context.Context) ChronicleReferenceListOutput
}

func (*ChronicleReferenceList) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleReferenceList)(nil)).Elem()
}

func (i *ChronicleReferenceList) ToChronicleReferenceListOutput() ChronicleReferenceListOutput {
	return i.ToChronicleReferenceListOutputWithContext(context.Background())
}

func (i *ChronicleReferenceList) ToChronicleReferenceListOutputWithContext(ctx context.Context) ChronicleReferenceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChronicleReferenceListOutput)
}

type ChronicleReferenceListOutput struct{ *pulumi.OutputState }

func (ChronicleReferenceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleReferenceList)(nil)).Elem()
}

func (o ChronicleReferenceListOutput) ToChronicleReferenceListOutput() ChronicleReferenceListOutput {
	return o
}

func (o ChronicleReferenceListOutput) ToChronicleReferenceListOutputWithContext(ctx context.Context) ChronicleReferenceListOutput {
	return o
}

func (o ChronicleReferenceListOutput) ChronicleReferenceListId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.ChronicleReferenceListId }).(pulumi.StringOutput)
}

// Required. A user-provided description of the reference list.
func (o ChronicleReferenceListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Output only. The unique display name of the reference list.
func (o ChronicleReferenceListOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Required. The entries of the reference list. When listed, they are returned in the order that was specified at creation
// or update. The combined size of the values of the reference list may not exceed 6MB. This is returned only when the view
// is REFERENCE_LIST_VIEW_FULL.
func (o ChronicleReferenceListOutput) Entries() ChronicleReferenceListEntryArrayOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) ChronicleReferenceListEntryArrayOutput { return v.Entries }).(ChronicleReferenceListEntryArrayOutput)
}

// The unique identifier for the Chronicle instance, which is the same as the customer ID.
func (o ChronicleReferenceListOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
// "europe-west2".
func (o ChronicleReferenceListOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. The resource name of the reference list. Format:
// projects/{project}/locations/{location}/instances/{instance}/referenceLists/{reference_list}
func (o ChronicleReferenceListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ChronicleReferenceListOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. The ID to use for the reference list. This is also the display name for the reference list. It must satisfy
// the following requirements: - Starts with letter. - Contains only letters, numbers and underscore. - Has length < 256. -
// Must be unique.
func (o ChronicleReferenceListOutput) ReferenceListId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.ReferenceListId }).(pulumi.StringOutput)
}

// Output only. The timestamp when the reference list was last updated.
func (o ChronicleReferenceListOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// Output only. The count of self-authored rules using the reference list.
func (o ChronicleReferenceListOutput) RuleAssociationsCount() pulumi.Float64Output {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.Float64Output { return v.RuleAssociationsCount }).(pulumi.Float64Output)
}

// Output only. The resource names for the associated self-authored Rules that use this reference list. This is returned
// only when the view is REFERENCE_LIST_VIEW_FULL.
func (o ChronicleReferenceListOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringArrayOutput { return v.Rules }).(pulumi.StringArrayOutput)
}

// ScopeInfo specifies the scope info of the reference list.
func (o ChronicleReferenceListOutput) ScopeInfos() ChronicleReferenceListScopeInfoArrayOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) ChronicleReferenceListScopeInfoArrayOutput { return v.ScopeInfos }).(ChronicleReferenceListScopeInfoArrayOutput)
}

// Possible values: REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING REFERENCE_LIST_SYNTAX_TYPE_REGEX
// REFERENCE_LIST_SYNTAX_TYPE_CIDR
func (o ChronicleReferenceListOutput) SyntaxType() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) pulumi.StringOutput { return v.SyntaxType }).(pulumi.StringOutput)
}

func (o ChronicleReferenceListOutput) Timeouts() ChronicleReferenceListTimeoutsPtrOutput {
	return o.ApplyT(func(v *ChronicleReferenceList) ChronicleReferenceListTimeoutsPtrOutput { return v.Timeouts }).(ChronicleReferenceListTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChronicleReferenceListInput)(nil)).Elem(), &ChronicleReferenceList{})
	pulumi.RegisterOutputType(ChronicleReferenceListOutput{})
}
