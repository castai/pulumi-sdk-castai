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

type DatabaseMigrationServiceConnectionProfile struct {
	pulumi.CustomResourceState

	// Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
	Alloydb DatabaseMigrationServiceConnectionProfileAlloydbPtrOutput `pulumi:"alloydb"`
	// Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination
	// database instance.
	Cloudsql DatabaseMigrationServiceConnectionProfileCloudsqlPtrOutput `pulumi:"cloudsql"`
	// The ID of the connection profile.
	ConnectionProfileId pulumi.StringOutput `pulumi:"connectionProfileId"`
	// Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC 'Zulu' format, accurate to
	// nanoseconds. Example: '2014-10-02T15:01:23.045123456Z'.
	CreateTime                                  pulumi.StringOutput `pulumi:"createTime"`
	DatabaseMigrationServiceConnectionProfileId pulumi.StringOutput `pulumi:"databaseMigrationServiceConnectionProfileId"`
	// The database provider.
	Dbprovider pulumi.StringOutput `pulumi:"dbprovider"`
	// The connection profile display name.
	DisplayName     pulumi.StringPtrOutput `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Output only. The error details in case of state FAILED.
	Errors DatabaseMigrationServiceConnectionProfileErrorArrayOutput `pulumi:"errors"`
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine
	// VMs. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where the connection profile should reside.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies connection parameters required specifically for MySQL databases.
	Mysql DatabaseMigrationServiceConnectionProfileMysqlPtrOutput `pulumi:"mysql"`
	// The name of this connection profile resource in the form of
	// projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies connection parameters required specifically for Oracle databases.
	Oracle DatabaseMigrationServiceConnectionProfileOraclePtrOutput `pulumi:"oracle"`
	// Specifies connection parameters required specifically for PostgreSQL databases.
	Postgresql DatabaseMigrationServiceConnectionProfilePostgresqlPtrOutput `pulumi:"postgresql"`
	Project    pulumi.StringOutput                                          `pulumi:"project"`
	// The current connection profile state.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                                     `pulumi:"terraformLabels"`
	Timeouts        DatabaseMigrationServiceConnectionProfileTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDatabaseMigrationServiceConnectionProfile registers a new resource with the given unique name, arguments, and options.
func NewDatabaseMigrationServiceConnectionProfile(ctx *pulumi.Context,
	name string, args *DatabaseMigrationServiceConnectionProfileArgs, opts ...pulumi.ResourceOption) (*DatabaseMigrationServiceConnectionProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionProfileId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DatabaseMigrationServiceConnectionProfile
	err = ctx.RegisterPackageResource("google-beta:index/databaseMigrationServiceConnectionProfile:DatabaseMigrationServiceConnectionProfile", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseMigrationServiceConnectionProfile gets an existing DatabaseMigrationServiceConnectionProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseMigrationServiceConnectionProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseMigrationServiceConnectionProfileState, opts ...pulumi.ResourceOption) (*DatabaseMigrationServiceConnectionProfile, error) {
	var resource DatabaseMigrationServiceConnectionProfile
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/databaseMigrationServiceConnectionProfile:DatabaseMigrationServiceConnectionProfile", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseMigrationServiceConnectionProfile resources.
type databaseMigrationServiceConnectionProfileState struct {
	// Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
	Alloydb *DatabaseMigrationServiceConnectionProfileAlloydb `pulumi:"alloydb"`
	// Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination
	// database instance.
	Cloudsql *DatabaseMigrationServiceConnectionProfileCloudsql `pulumi:"cloudsql"`
	// The ID of the connection profile.
	ConnectionProfileId *string `pulumi:"connectionProfileId"`
	// Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC 'Zulu' format, accurate to
	// nanoseconds. Example: '2014-10-02T15:01:23.045123456Z'.
	CreateTime                                  *string `pulumi:"createTime"`
	DatabaseMigrationServiceConnectionProfileId *string `pulumi:"databaseMigrationServiceConnectionProfileId"`
	// The database provider.
	Dbprovider *string `pulumi:"dbprovider"`
	// The connection profile display name.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Output only. The error details in case of state FAILED.
	Errors []DatabaseMigrationServiceConnectionProfileError `pulumi:"errors"`
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine
	// VMs. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the connection profile should reside.
	Location *string `pulumi:"location"`
	// Specifies connection parameters required specifically for MySQL databases.
	Mysql *DatabaseMigrationServiceConnectionProfileMysql `pulumi:"mysql"`
	// The name of this connection profile resource in the form of
	// projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
	Name *string `pulumi:"name"`
	// Specifies connection parameters required specifically for Oracle databases.
	Oracle *DatabaseMigrationServiceConnectionProfileOracle `pulumi:"oracle"`
	// Specifies connection parameters required specifically for PostgreSQL databases.
	Postgresql *DatabaseMigrationServiceConnectionProfilePostgresql `pulumi:"postgresql"`
	Project    *string                                              `pulumi:"project"`
	// The current connection profile state.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                                  `pulumi:"terraformLabels"`
	Timeouts        *DatabaseMigrationServiceConnectionProfileTimeouts `pulumi:"timeouts"`
}

type DatabaseMigrationServiceConnectionProfileState struct {
	// Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
	Alloydb DatabaseMigrationServiceConnectionProfileAlloydbPtrInput
	// Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination
	// database instance.
	Cloudsql DatabaseMigrationServiceConnectionProfileCloudsqlPtrInput
	// The ID of the connection profile.
	ConnectionProfileId pulumi.StringPtrInput
	// Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC 'Zulu' format, accurate to
	// nanoseconds. Example: '2014-10-02T15:01:23.045123456Z'.
	CreateTime                                  pulumi.StringPtrInput
	DatabaseMigrationServiceConnectionProfileId pulumi.StringPtrInput
	// The database provider.
	Dbprovider pulumi.StringPtrInput
	// The connection profile display name.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Output only. The error details in case of state FAILED.
	Errors DatabaseMigrationServiceConnectionProfileErrorArrayInput
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine
	// VMs. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location where the connection profile should reside.
	Location pulumi.StringPtrInput
	// Specifies connection parameters required specifically for MySQL databases.
	Mysql DatabaseMigrationServiceConnectionProfileMysqlPtrInput
	// The name of this connection profile resource in the form of
	// projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
	Name pulumi.StringPtrInput
	// Specifies connection parameters required specifically for Oracle databases.
	Oracle DatabaseMigrationServiceConnectionProfileOraclePtrInput
	// Specifies connection parameters required specifically for PostgreSQL databases.
	Postgresql DatabaseMigrationServiceConnectionProfilePostgresqlPtrInput
	Project    pulumi.StringPtrInput
	// The current connection profile state.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DatabaseMigrationServiceConnectionProfileTimeoutsPtrInput
}

func (DatabaseMigrationServiceConnectionProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseMigrationServiceConnectionProfileState)(nil)).Elem()
}

type databaseMigrationServiceConnectionProfileArgs struct {
	// Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
	Alloydb *DatabaseMigrationServiceConnectionProfileAlloydb `pulumi:"alloydb"`
	// Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination
	// database instance.
	Cloudsql *DatabaseMigrationServiceConnectionProfileCloudsql `pulumi:"cloudsql"`
	// The ID of the connection profile.
	ConnectionProfileId                         string  `pulumi:"connectionProfileId"`
	DatabaseMigrationServiceConnectionProfileId *string `pulumi:"databaseMigrationServiceConnectionProfileId"`
	// The connection profile display name.
	DisplayName *string `pulumi:"displayName"`
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine
	// VMs. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the connection profile should reside.
	Location *string `pulumi:"location"`
	// Specifies connection parameters required specifically for MySQL databases.
	Mysql *DatabaseMigrationServiceConnectionProfileMysql `pulumi:"mysql"`
	// Specifies connection parameters required specifically for Oracle databases.
	Oracle *DatabaseMigrationServiceConnectionProfileOracle `pulumi:"oracle"`
	// Specifies connection parameters required specifically for PostgreSQL databases.
	Postgresql *DatabaseMigrationServiceConnectionProfilePostgresql `pulumi:"postgresql"`
	Project    *string                                              `pulumi:"project"`
	Timeouts   *DatabaseMigrationServiceConnectionProfileTimeouts   `pulumi:"timeouts"`
}

// The set of arguments for constructing a DatabaseMigrationServiceConnectionProfile resource.
type DatabaseMigrationServiceConnectionProfileArgs struct {
	// Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
	Alloydb DatabaseMigrationServiceConnectionProfileAlloydbPtrInput
	// Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination
	// database instance.
	Cloudsql DatabaseMigrationServiceConnectionProfileCloudsqlPtrInput
	// The ID of the connection profile.
	ConnectionProfileId                         pulumi.StringInput
	DatabaseMigrationServiceConnectionProfileId pulumi.StringPtrInput
	// The connection profile display name.
	DisplayName pulumi.StringPtrInput
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine
	// VMs. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
	// refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location where the connection profile should reside.
	Location pulumi.StringPtrInput
	// Specifies connection parameters required specifically for MySQL databases.
	Mysql DatabaseMigrationServiceConnectionProfileMysqlPtrInput
	// Specifies connection parameters required specifically for Oracle databases.
	Oracle DatabaseMigrationServiceConnectionProfileOraclePtrInput
	// Specifies connection parameters required specifically for PostgreSQL databases.
	Postgresql DatabaseMigrationServiceConnectionProfilePostgresqlPtrInput
	Project    pulumi.StringPtrInput
	Timeouts   DatabaseMigrationServiceConnectionProfileTimeoutsPtrInput
}

func (DatabaseMigrationServiceConnectionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseMigrationServiceConnectionProfileArgs)(nil)).Elem()
}

type DatabaseMigrationServiceConnectionProfileInput interface {
	pulumi.Input

	ToDatabaseMigrationServiceConnectionProfileOutput() DatabaseMigrationServiceConnectionProfileOutput
	ToDatabaseMigrationServiceConnectionProfileOutputWithContext(ctx context.Context) DatabaseMigrationServiceConnectionProfileOutput
}

func (*DatabaseMigrationServiceConnectionProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseMigrationServiceConnectionProfile)(nil)).Elem()
}

func (i *DatabaseMigrationServiceConnectionProfile) ToDatabaseMigrationServiceConnectionProfileOutput() DatabaseMigrationServiceConnectionProfileOutput {
	return i.ToDatabaseMigrationServiceConnectionProfileOutputWithContext(context.Background())
}

func (i *DatabaseMigrationServiceConnectionProfile) ToDatabaseMigrationServiceConnectionProfileOutputWithContext(ctx context.Context) DatabaseMigrationServiceConnectionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMigrationServiceConnectionProfileOutput)
}

type DatabaseMigrationServiceConnectionProfileOutput struct{ *pulumi.OutputState }

func (DatabaseMigrationServiceConnectionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseMigrationServiceConnectionProfile)(nil)).Elem()
}

func (o DatabaseMigrationServiceConnectionProfileOutput) ToDatabaseMigrationServiceConnectionProfileOutput() DatabaseMigrationServiceConnectionProfileOutput {
	return o
}

func (o DatabaseMigrationServiceConnectionProfileOutput) ToDatabaseMigrationServiceConnectionProfileOutputWithContext(ctx context.Context) DatabaseMigrationServiceConnectionProfileOutput {
	return o
}

// Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
func (o DatabaseMigrationServiceConnectionProfileOutput) Alloydb() DatabaseMigrationServiceConnectionProfileAlloydbPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) DatabaseMigrationServiceConnectionProfileAlloydbPtrOutput {
		return v.Alloydb
	}).(DatabaseMigrationServiceConnectionProfileAlloydbPtrOutput)
}

// Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination
// database instance.
func (o DatabaseMigrationServiceConnectionProfileOutput) Cloudsql() DatabaseMigrationServiceConnectionProfileCloudsqlPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) DatabaseMigrationServiceConnectionProfileCloudsqlPtrOutput {
		return v.Cloudsql
	}).(DatabaseMigrationServiceConnectionProfileCloudsqlPtrOutput)
}

// The ID of the connection profile.
func (o DatabaseMigrationServiceConnectionProfileOutput) ConnectionProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringOutput { return v.ConnectionProfileId }).(pulumi.StringOutput)
}

// Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC 'Zulu' format, accurate to
// nanoseconds. Example: '2014-10-02T15:01:23.045123456Z'.
func (o DatabaseMigrationServiceConnectionProfileOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o DatabaseMigrationServiceConnectionProfileOutput) DatabaseMigrationServiceConnectionProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringOutput {
		return v.DatabaseMigrationServiceConnectionProfileId
	}).(pulumi.StringOutput)
}

// The database provider.
func (o DatabaseMigrationServiceConnectionProfileOutput) Dbprovider() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringOutput { return v.Dbprovider }).(pulumi.StringOutput)
}

// The connection profile display name.
func (o DatabaseMigrationServiceConnectionProfileOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationServiceConnectionProfileOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Output only. The error details in case of state FAILED.
func (o DatabaseMigrationServiceConnectionProfileOutput) Errors() DatabaseMigrationServiceConnectionProfileErrorArrayOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) DatabaseMigrationServiceConnectionProfileErrorArrayOutput {
		return v.Errors
	}).(DatabaseMigrationServiceConnectionProfileErrorArrayOutput)
}

// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine
// VMs. **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please
// refer to the field 'effective_labels' for all of the labels present on the resource.
func (o DatabaseMigrationServiceConnectionProfileOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where the connection profile should reside.
func (o DatabaseMigrationServiceConnectionProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies connection parameters required specifically for MySQL databases.
func (o DatabaseMigrationServiceConnectionProfileOutput) Mysql() DatabaseMigrationServiceConnectionProfileMysqlPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) DatabaseMigrationServiceConnectionProfileMysqlPtrOutput {
		return v.Mysql
	}).(DatabaseMigrationServiceConnectionProfileMysqlPtrOutput)
}

// The name of this connection profile resource in the form of
// projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
func (o DatabaseMigrationServiceConnectionProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies connection parameters required specifically for Oracle databases.
func (o DatabaseMigrationServiceConnectionProfileOutput) Oracle() DatabaseMigrationServiceConnectionProfileOraclePtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) DatabaseMigrationServiceConnectionProfileOraclePtrOutput {
		return v.Oracle
	}).(DatabaseMigrationServiceConnectionProfileOraclePtrOutput)
}

// Specifies connection parameters required specifically for PostgreSQL databases.
func (o DatabaseMigrationServiceConnectionProfileOutput) Postgresql() DatabaseMigrationServiceConnectionProfilePostgresqlPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) DatabaseMigrationServiceConnectionProfilePostgresqlPtrOutput {
		return v.Postgresql
	}).(DatabaseMigrationServiceConnectionProfilePostgresqlPtrOutput)
}

func (o DatabaseMigrationServiceConnectionProfileOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The current connection profile state.
func (o DatabaseMigrationServiceConnectionProfileOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DatabaseMigrationServiceConnectionProfileOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DatabaseMigrationServiceConnectionProfileOutput) Timeouts() DatabaseMigrationServiceConnectionProfileTimeoutsPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationServiceConnectionProfile) DatabaseMigrationServiceConnectionProfileTimeoutsPtrOutput {
		return v.Timeouts
	}).(DatabaseMigrationServiceConnectionProfileTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMigrationServiceConnectionProfileInput)(nil)).Elem(), &DatabaseMigrationServiceConnectionProfile{})
	pulumi.RegisterOutputType(DatabaseMigrationServiceConnectionProfileOutput{})
}
