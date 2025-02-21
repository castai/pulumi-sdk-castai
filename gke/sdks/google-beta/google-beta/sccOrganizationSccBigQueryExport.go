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

type SccOrganizationSccBigQueryExport struct {
	pulumi.CustomResourceState

	// This must be unique within the organization.
	BigQueryExportId pulumi.StringOutput `pulumi:"bigQueryExportId"`
	// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
	// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset pulumi.StringPtrOutput `pulumi:"dataset"`
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * \>, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are:
	// * string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	//   [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	//   information on how to write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
	// provided on export creation or update.
	MostRecentEditor pulumi.StringOutput `pulumi:"mostRecentEditor"`
	// The resource name of this export, in the format
	// 'organizations/{{organization}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in responses, and is
	// ignored when provided in create requests.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Big Query Export Config lives in.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The service account that needs permission to create table and upload data to the BigQuery dataset.
	Principal                          pulumi.StringOutput                               `pulumi:"principal"`
	SccOrganizationSccBigQueryExportId pulumi.StringOutput                               `pulumi:"sccOrganizationSccBigQueryExportId"`
	Timeouts                           SccOrganizationSccBigQueryExportTimeoutsPtrOutput `pulumi:"timeouts"`
	// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
	// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSccOrganizationSccBigQueryExport registers a new resource with the given unique name, arguments, and options.
func NewSccOrganizationSccBigQueryExport(ctx *pulumi.Context,
	name string, args *SccOrganizationSccBigQueryExportArgs, opts ...pulumi.ResourceOption) (*SccOrganizationSccBigQueryExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BigQueryExportId == nil {
		return nil, errors.New("invalid value for required argument 'BigQueryExportId'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccOrganizationSccBigQueryExport
	err = ctx.RegisterPackageResource("google-beta:index/sccOrganizationSccBigQueryExport:SccOrganizationSccBigQueryExport", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccOrganizationSccBigQueryExport gets an existing SccOrganizationSccBigQueryExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccOrganizationSccBigQueryExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccOrganizationSccBigQueryExportState, opts ...pulumi.ResourceOption) (*SccOrganizationSccBigQueryExport, error) {
	var resource SccOrganizationSccBigQueryExport
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/sccOrganizationSccBigQueryExport:SccOrganizationSccBigQueryExport", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccOrganizationSccBigQueryExport resources.
type sccOrganizationSccBigQueryExportState struct {
	// This must be unique within the organization.
	BigQueryExportId *string `pulumi:"bigQueryExportId"`
	// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
	// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset *string `pulumi:"dataset"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * \>, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are:
	// * string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	//   [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	//   information on how to write a filter.
	Filter *string `pulumi:"filter"`
	// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
	// provided on export creation or update.
	MostRecentEditor *string `pulumi:"mostRecentEditor"`
	// The resource name of this export, in the format
	// 'organizations/{{organization}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in responses, and is
	// ignored when provided in create requests.
	Name *string `pulumi:"name"`
	// The organization whose Cloud Security Command Center the Big Query Export Config lives in.
	Organization *string `pulumi:"organization"`
	// The service account that needs permission to create table and upload data to the BigQuery dataset.
	Principal                          *string                                   `pulumi:"principal"`
	SccOrganizationSccBigQueryExportId *string                                   `pulumi:"sccOrganizationSccBigQueryExportId"`
	Timeouts                           *SccOrganizationSccBigQueryExportTimeouts `pulumi:"timeouts"`
	// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
	// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type SccOrganizationSccBigQueryExportState struct {
	// This must be unique within the organization.
	BigQueryExportId pulumi.StringPtrInput
	// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
	// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset pulumi.StringPtrInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * \>, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are:
	// * string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	//   [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	//   information on how to write a filter.
	Filter pulumi.StringPtrInput
	// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
	// provided on export creation or update.
	MostRecentEditor pulumi.StringPtrInput
	// The resource name of this export, in the format
	// 'organizations/{{organization}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in responses, and is
	// ignored when provided in create requests.
	Name pulumi.StringPtrInput
	// The organization whose Cloud Security Command Center the Big Query Export Config lives in.
	Organization pulumi.StringPtrInput
	// The service account that needs permission to create table and upload data to the BigQuery dataset.
	Principal                          pulumi.StringPtrInput
	SccOrganizationSccBigQueryExportId pulumi.StringPtrInput
	Timeouts                           SccOrganizationSccBigQueryExportTimeoutsPtrInput
	// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
	// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (SccOrganizationSccBigQueryExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccOrganizationSccBigQueryExportState)(nil)).Elem()
}

type sccOrganizationSccBigQueryExportArgs struct {
	// This must be unique within the organization.
	BigQueryExportId string `pulumi:"bigQueryExportId"`
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset *string `pulumi:"dataset"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * \>, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are:
	// * string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	//   [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	//   information on how to write a filter.
	Filter *string `pulumi:"filter"`
	// The organization whose Cloud Security Command Center the Big Query Export Config lives in.
	Organization                       string                                    `pulumi:"organization"`
	SccOrganizationSccBigQueryExportId *string                                   `pulumi:"sccOrganizationSccBigQueryExportId"`
	Timeouts                           *SccOrganizationSccBigQueryExportTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a SccOrganizationSccBigQueryExport resource.
type SccOrganizationSccBigQueryExportArgs struct {
	// This must be unique within the organization.
	BigQueryExportId pulumi.StringInput
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset pulumi.StringPtrInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * \>, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are:
	// * string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	//   [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	//   information on how to write a filter.
	Filter pulumi.StringPtrInput
	// The organization whose Cloud Security Command Center the Big Query Export Config lives in.
	Organization                       pulumi.StringInput
	SccOrganizationSccBigQueryExportId pulumi.StringPtrInput
	Timeouts                           SccOrganizationSccBigQueryExportTimeoutsPtrInput
}

func (SccOrganizationSccBigQueryExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccOrganizationSccBigQueryExportArgs)(nil)).Elem()
}

type SccOrganizationSccBigQueryExportInput interface {
	pulumi.Input

	ToSccOrganizationSccBigQueryExportOutput() SccOrganizationSccBigQueryExportOutput
	ToSccOrganizationSccBigQueryExportOutputWithContext(ctx context.Context) SccOrganizationSccBigQueryExportOutput
}

func (*SccOrganizationSccBigQueryExport) ElementType() reflect.Type {
	return reflect.TypeOf((**SccOrganizationSccBigQueryExport)(nil)).Elem()
}

func (i *SccOrganizationSccBigQueryExport) ToSccOrganizationSccBigQueryExportOutput() SccOrganizationSccBigQueryExportOutput {
	return i.ToSccOrganizationSccBigQueryExportOutputWithContext(context.Background())
}

func (i *SccOrganizationSccBigQueryExport) ToSccOrganizationSccBigQueryExportOutputWithContext(ctx context.Context) SccOrganizationSccBigQueryExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccOrganizationSccBigQueryExportOutput)
}

type SccOrganizationSccBigQueryExportOutput struct{ *pulumi.OutputState }

func (SccOrganizationSccBigQueryExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccOrganizationSccBigQueryExport)(nil)).Elem()
}

func (o SccOrganizationSccBigQueryExportOutput) ToSccOrganizationSccBigQueryExportOutput() SccOrganizationSccBigQueryExportOutput {
	return o
}

func (o SccOrganizationSccBigQueryExportOutput) ToSccOrganizationSccBigQueryExportOutputWithContext(ctx context.Context) SccOrganizationSccBigQueryExportOutput {
	return o
}

// This must be unique within the organization.
func (o SccOrganizationSccBigQueryExportOutput) BigQueryExportId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput { return v.BigQueryExportId }).(pulumi.StringOutput)
}

// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccOrganizationSccBigQueryExportOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
func (o SccOrganizationSccBigQueryExportOutput) Dataset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringPtrOutput { return v.Dataset }).(pulumi.StringPtrOutput)
}

// The description of the notification config (max of 1024 characters).
func (o SccOrganizationSccBigQueryExportOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
// types. * \>, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are:
//   - string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
//     [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
//     information on how to write a filter.
func (o SccOrganizationSccBigQueryExportOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringPtrOutput { return v.Filter }).(pulumi.StringPtrOutput)
}

// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
// provided on export creation or update.
func (o SccOrganizationSccBigQueryExportOutput) MostRecentEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput { return v.MostRecentEditor }).(pulumi.StringOutput)
}

// The resource name of this export, in the format
// 'organizations/{{organization}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in responses, and is
// ignored when provided in create requests.
func (o SccOrganizationSccBigQueryExportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization whose Cloud Security Command Center the Big Query Export Config lives in.
func (o SccOrganizationSccBigQueryExportOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The service account that needs permission to create table and upload data to the BigQuery dataset.
func (o SccOrganizationSccBigQueryExportOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

func (o SccOrganizationSccBigQueryExportOutput) SccOrganizationSccBigQueryExportId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput {
		return v.SccOrganizationSccBigQueryExportId
	}).(pulumi.StringOutput)
}

func (o SccOrganizationSccBigQueryExportOutput) Timeouts() SccOrganizationSccBigQueryExportTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) SccOrganizationSccBigQueryExportTimeoutsPtrOutput {
		return v.Timeouts
	}).(SccOrganizationSccBigQueryExportTimeoutsPtrOutput)
}

// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccOrganizationSccBigQueryExportOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccOrganizationSccBigQueryExport) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccOrganizationSccBigQueryExportInput)(nil)).Elem(), &SccOrganizationSccBigQueryExport{})
	pulumi.RegisterOutputType(SccOrganizationSccBigQueryExportOutput{})
}
