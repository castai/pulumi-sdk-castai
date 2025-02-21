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

type SecretManagerRegionalSecret struct {
	pulumi.CustomResourceState

	// Custom metadata about the regional secret. Annotations are distinct from various forms of labels. Annotations exist to
	// allow client tools to store their own state information without requiring a database. Annotation keys must be between 1
	// and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, begin and end with an alphanumeric character
	// ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and alphanumerics in between these symbols. The total
	// size of annotation keys and values must be less than 16KiB. An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the annotations
	// present on the resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// The time at which the regional secret was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The customer-managed encryption configuration of the regional secret.
	CustomerManagedEncryption SecretManagerRegionalSecretCustomerManagedEncryptionPtrOutput `pulumi:"customerManagedEncryption"`
	EffectiveAnnotations      pulumi.StringMapOutput                                        `pulumi:"effectiveAnnotations"`
	EffectiveLabels           pulumi.StringMapOutput                                        `pulumi:"effectiveLabels"`
	// Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on output, regardless of what
	// was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be
	// provided.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The labels assigned to this regional secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more
	// than 64 labels can be assigned to a given resource. An object containing a list of "key": value pairs. Example: {
	// "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location of the regional secret. eg us-central1
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the regional secret. Format: 'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}'
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub
	// notification to the topics configured on the Secret. 'topics' must be set to configure rotation.
	Rotation SecretManagerRegionalSecretRotationPtrOutput `pulumi:"rotation"`
	// This must be unique within the project.
	SecretId                      pulumi.StringOutput `pulumi:"secretId"`
	SecretManagerRegionalSecretId pulumi.StringOutput `pulumi:"secretManagerRegionalSecretId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                       `pulumi:"terraformLabels"`
	Timeouts        SecretManagerRegionalSecretTimeoutsPtrOutput `pulumi:"timeouts"`
	// A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the
	// regional secret or its versions.
	Topics SecretManagerRegionalSecretTopicArrayOutput `pulumi:"topics"`
	// The TTL for the regional secret. A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
	// "3.5s". Only one of 'ttl' or 'expire_time' can be provided.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can
	// contain uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_') characters. An alias string
	// must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given
	// secret. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	VersionAliases pulumi.StringMapOutput `pulumi:"versionAliases"`
	// Secret Version TTL after destruction request. This is a part of the delayed delete feature on Secret Version. For secret
	// with versionDestroyTtl>0, version destruction doesn't happen immediately on calling destroy instead the version goes to
	// a disabled state and the actual destruction happens after this TTL expires. It must be atleast 24h.
	VersionDestroyTtl pulumi.StringPtrOutput `pulumi:"versionDestroyTtl"`
}

// NewSecretManagerRegionalSecret registers a new resource with the given unique name, arguments, and options.
func NewSecretManagerRegionalSecret(ctx *pulumi.Context,
	name string, args *SecretManagerRegionalSecretArgs, opts ...pulumi.ResourceOption) (*SecretManagerRegionalSecret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.SecretId == nil {
		return nil, errors.New("invalid value for required argument 'SecretId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SecretManagerRegionalSecret
	err = ctx.RegisterPackageResource("google:index/secretManagerRegionalSecret:SecretManagerRegionalSecret", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretManagerRegionalSecret gets an existing SecretManagerRegionalSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretManagerRegionalSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretManagerRegionalSecretState, opts ...pulumi.ResourceOption) (*SecretManagerRegionalSecret, error) {
	var resource SecretManagerRegionalSecret
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/secretManagerRegionalSecret:SecretManagerRegionalSecret", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretManagerRegionalSecret resources.
type secretManagerRegionalSecretState struct {
	// Custom metadata about the regional secret. Annotations are distinct from various forms of labels. Annotations exist to
	// allow client tools to store their own state information without requiring a database. Annotation keys must be between 1
	// and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, begin and end with an alphanumeric character
	// ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and alphanumerics in between these symbols. The total
	// size of annotation keys and values must be less than 16KiB. An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the annotations
	// present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// The time at which the regional secret was created.
	CreateTime *string `pulumi:"createTime"`
	// The customer-managed encryption configuration of the regional secret.
	CustomerManagedEncryption *SecretManagerRegionalSecretCustomerManagedEncryption `pulumi:"customerManagedEncryption"`
	EffectiveAnnotations      map[string]string                                     `pulumi:"effectiveAnnotations"`
	EffectiveLabels           map[string]string                                     `pulumi:"effectiveLabels"`
	// Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on output, regardless of what
	// was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be
	// provided.
	ExpireTime *string `pulumi:"expireTime"`
	// The labels assigned to this regional secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more
	// than 64 labels can be assigned to a given resource. An object containing a list of "key": value pairs. Example: {
	// "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the regional secret. eg us-central1
	Location *string `pulumi:"location"`
	// The resource name of the regional secret. Format: 'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}'
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub
	// notification to the topics configured on the Secret. 'topics' must be set to configure rotation.
	Rotation *SecretManagerRegionalSecretRotation `pulumi:"rotation"`
	// This must be unique within the project.
	SecretId                      *string `pulumi:"secretId"`
	SecretManagerRegionalSecretId *string `pulumi:"secretManagerRegionalSecretId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                    `pulumi:"terraformLabels"`
	Timeouts        *SecretManagerRegionalSecretTimeouts `pulumi:"timeouts"`
	// A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the
	// regional secret or its versions.
	Topics []SecretManagerRegionalSecretTopic `pulumi:"topics"`
	// The TTL for the regional secret. A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
	// "3.5s". Only one of 'ttl' or 'expire_time' can be provided.
	Ttl *string `pulumi:"ttl"`
	// Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can
	// contain uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_') characters. An alias string
	// must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given
	// secret. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	VersionAliases map[string]string `pulumi:"versionAliases"`
	// Secret Version TTL after destruction request. This is a part of the delayed delete feature on Secret Version. For secret
	// with versionDestroyTtl>0, version destruction doesn't happen immediately on calling destroy instead the version goes to
	// a disabled state and the actual destruction happens after this TTL expires. It must be atleast 24h.
	VersionDestroyTtl *string `pulumi:"versionDestroyTtl"`
}

type SecretManagerRegionalSecretState struct {
	// Custom metadata about the regional secret. Annotations are distinct from various forms of labels. Annotations exist to
	// allow client tools to store their own state information without requiring a database. Annotation keys must be between 1
	// and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, begin and end with an alphanumeric character
	// ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and alphanumerics in between these symbols. The total
	// size of annotation keys and values must be less than 16KiB. An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the annotations
	// present on the resource.
	Annotations pulumi.StringMapInput
	// The time at which the regional secret was created.
	CreateTime pulumi.StringPtrInput
	// The customer-managed encryption configuration of the regional secret.
	CustomerManagedEncryption SecretManagerRegionalSecretCustomerManagedEncryptionPtrInput
	EffectiveAnnotations      pulumi.StringMapInput
	EffectiveLabels           pulumi.StringMapInput
	// Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on output, regardless of what
	// was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be
	// provided.
	ExpireTime pulumi.StringPtrInput
	// The labels assigned to this regional secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more
	// than 64 labels can be assigned to a given resource. An object containing a list of "key": value pairs. Example: {
	// "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location of the regional secret. eg us-central1
	Location pulumi.StringPtrInput
	// The resource name of the regional secret. Format: 'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}'
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub
	// notification to the topics configured on the Secret. 'topics' must be set to configure rotation.
	Rotation SecretManagerRegionalSecretRotationPtrInput
	// This must be unique within the project.
	SecretId                      pulumi.StringPtrInput
	SecretManagerRegionalSecretId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        SecretManagerRegionalSecretTimeoutsPtrInput
	// A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the
	// regional secret or its versions.
	Topics SecretManagerRegionalSecretTopicArrayInput
	// The TTL for the regional secret. A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
	// "3.5s". Only one of 'ttl' or 'expire_time' can be provided.
	Ttl pulumi.StringPtrInput
	// Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can
	// contain uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_') characters. An alias string
	// must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given
	// secret. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	VersionAliases pulumi.StringMapInput
	// Secret Version TTL after destruction request. This is a part of the delayed delete feature on Secret Version. For secret
	// with versionDestroyTtl>0, version destruction doesn't happen immediately on calling destroy instead the version goes to
	// a disabled state and the actual destruction happens after this TTL expires. It must be atleast 24h.
	VersionDestroyTtl pulumi.StringPtrInput
}

func (SecretManagerRegionalSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretManagerRegionalSecretState)(nil)).Elem()
}

type secretManagerRegionalSecretArgs struct {
	// Custom metadata about the regional secret. Annotations are distinct from various forms of labels. Annotations exist to
	// allow client tools to store their own state information without requiring a database. Annotation keys must be between 1
	// and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, begin and end with an alphanumeric character
	// ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and alphanumerics in between these symbols. The total
	// size of annotation keys and values must be less than 16KiB. An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the annotations
	// present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// The customer-managed encryption configuration of the regional secret.
	CustomerManagedEncryption *SecretManagerRegionalSecretCustomerManagedEncryption `pulumi:"customerManagedEncryption"`
	// Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on output, regardless of what
	// was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be
	// provided.
	ExpireTime *string `pulumi:"expireTime"`
	// The labels assigned to this regional secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more
	// than 64 labels can be assigned to a given resource. An object containing a list of "key": value pairs. Example: {
	// "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location of the regional secret. eg us-central1
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub
	// notification to the topics configured on the Secret. 'topics' must be set to configure rotation.
	Rotation *SecretManagerRegionalSecretRotation `pulumi:"rotation"`
	// This must be unique within the project.
	SecretId                      string                               `pulumi:"secretId"`
	SecretManagerRegionalSecretId *string                              `pulumi:"secretManagerRegionalSecretId"`
	Timeouts                      *SecretManagerRegionalSecretTimeouts `pulumi:"timeouts"`
	// A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the
	// regional secret or its versions.
	Topics []SecretManagerRegionalSecretTopic `pulumi:"topics"`
	// The TTL for the regional secret. A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
	// "3.5s". Only one of 'ttl' or 'expire_time' can be provided.
	Ttl *string `pulumi:"ttl"`
	// Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can
	// contain uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_') characters. An alias string
	// must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given
	// secret. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	VersionAliases map[string]string `pulumi:"versionAliases"`
	// Secret Version TTL after destruction request. This is a part of the delayed delete feature on Secret Version. For secret
	// with versionDestroyTtl>0, version destruction doesn't happen immediately on calling destroy instead the version goes to
	// a disabled state and the actual destruction happens after this TTL expires. It must be atleast 24h.
	VersionDestroyTtl *string `pulumi:"versionDestroyTtl"`
}

// The set of arguments for constructing a SecretManagerRegionalSecret resource.
type SecretManagerRegionalSecretArgs struct {
	// Custom metadata about the regional secret. Annotations are distinct from various forms of labels. Annotations exist to
	// allow client tools to store their own state information without requiring a database. Annotation keys must be between 1
	// and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, begin and end with an alphanumeric character
	// ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and alphanumerics in between these symbols. The total
	// size of annotation keys and values must be less than 16KiB. An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the annotations
	// present on the resource.
	Annotations pulumi.StringMapInput
	// The customer-managed encryption configuration of the regional secret.
	CustomerManagedEncryption SecretManagerRegionalSecretCustomerManagedEncryptionPtrInput
	// Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on output, regardless of what
	// was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be
	// provided.
	ExpireTime pulumi.StringPtrInput
	// The labels assigned to this regional secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more
	// than 64 labels can be assigned to a given resource. An object containing a list of "key": value pairs. Example: {
	// "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location of the regional secret. eg us-central1
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub
	// notification to the topics configured on the Secret. 'topics' must be set to configure rotation.
	Rotation SecretManagerRegionalSecretRotationPtrInput
	// This must be unique within the project.
	SecretId                      pulumi.StringInput
	SecretManagerRegionalSecretId pulumi.StringPtrInput
	Timeouts                      SecretManagerRegionalSecretTimeoutsPtrInput
	// A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the
	// regional secret or its versions.
	Topics SecretManagerRegionalSecretTopicArrayInput
	// The TTL for the regional secret. A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
	// "3.5s". Only one of 'ttl' or 'expire_time' can be provided.
	Ttl pulumi.StringPtrInput
	// Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can
	// contain uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_') characters. An alias string
	// must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given
	// secret. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	VersionAliases pulumi.StringMapInput
	// Secret Version TTL after destruction request. This is a part of the delayed delete feature on Secret Version. For secret
	// with versionDestroyTtl>0, version destruction doesn't happen immediately on calling destroy instead the version goes to
	// a disabled state and the actual destruction happens after this TTL expires. It must be atleast 24h.
	VersionDestroyTtl pulumi.StringPtrInput
}

func (SecretManagerRegionalSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretManagerRegionalSecretArgs)(nil)).Elem()
}

type SecretManagerRegionalSecretInput interface {
	pulumi.Input

	ToSecretManagerRegionalSecretOutput() SecretManagerRegionalSecretOutput
	ToSecretManagerRegionalSecretOutputWithContext(ctx context.Context) SecretManagerRegionalSecretOutput
}

func (*SecretManagerRegionalSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretManagerRegionalSecret)(nil)).Elem()
}

func (i *SecretManagerRegionalSecret) ToSecretManagerRegionalSecretOutput() SecretManagerRegionalSecretOutput {
	return i.ToSecretManagerRegionalSecretOutputWithContext(context.Background())
}

func (i *SecretManagerRegionalSecret) ToSecretManagerRegionalSecretOutputWithContext(ctx context.Context) SecretManagerRegionalSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretManagerRegionalSecretOutput)
}

type SecretManagerRegionalSecretOutput struct{ *pulumi.OutputState }

func (SecretManagerRegionalSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretManagerRegionalSecret)(nil)).Elem()
}

func (o SecretManagerRegionalSecretOutput) ToSecretManagerRegionalSecretOutput() SecretManagerRegionalSecretOutput {
	return o
}

func (o SecretManagerRegionalSecretOutput) ToSecretManagerRegionalSecretOutputWithContext(ctx context.Context) SecretManagerRegionalSecretOutput {
	return o
}

// Custom metadata about the regional secret. Annotations are distinct from various forms of labels. Annotations exist to
// allow client tools to store their own state information without requiring a database. Annotation keys must be between 1
// and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, begin and end with an alphanumeric character
// ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and alphanumerics in between these symbols. The total
// size of annotation keys and values must be less than 16KiB. An object containing a list of "key": value pairs. Example:
// { "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
// annotations present in your configuration. Please refer to the field 'effective_annotations' for all of the annotations
// present on the resource.
func (o SecretManagerRegionalSecretOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// The time at which the regional secret was created.
func (o SecretManagerRegionalSecretOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The customer-managed encryption configuration of the regional secret.
func (o SecretManagerRegionalSecretOutput) CustomerManagedEncryption() SecretManagerRegionalSecretCustomerManagedEncryptionPtrOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) SecretManagerRegionalSecretCustomerManagedEncryptionPtrOutput {
		return v.CustomerManagedEncryption
	}).(SecretManagerRegionalSecretCustomerManagedEncryptionPtrOutput)
}

func (o SecretManagerRegionalSecretOutput) EffectiveAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringMapOutput { return v.EffectiveAnnotations }).(pulumi.StringMapOutput)
}

func (o SecretManagerRegionalSecretOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on output, regardless of what
// was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be
// provided.
func (o SecretManagerRegionalSecretOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// The labels assigned to this regional secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
// of maximum 128 bytes, and must conform to the following PCRE regular expression:
// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a UTF-8 encoding
// of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more
// than 64 labels can be assigned to a given resource. An object containing a list of "key": value pairs. Example: {
// "name": "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o SecretManagerRegionalSecretOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location of the regional secret. eg us-central1
func (o SecretManagerRegionalSecretOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the regional secret. Format: 'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}'
func (o SecretManagerRegionalSecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub
// notification to the topics configured on the Secret. 'topics' must be set to configure rotation.
func (o SecretManagerRegionalSecretOutput) Rotation() SecretManagerRegionalSecretRotationPtrOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) SecretManagerRegionalSecretRotationPtrOutput { return v.Rotation }).(SecretManagerRegionalSecretRotationPtrOutput)
}

// This must be unique within the project.
func (o SecretManagerRegionalSecretOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringOutput { return v.SecretId }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretOutput) SecretManagerRegionalSecretId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringOutput { return v.SecretManagerRegionalSecretId }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o SecretManagerRegionalSecretOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o SecretManagerRegionalSecretOutput) Timeouts() SecretManagerRegionalSecretTimeoutsPtrOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) SecretManagerRegionalSecretTimeoutsPtrOutput { return v.Timeouts }).(SecretManagerRegionalSecretTimeoutsPtrOutput)
}

// A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the
// regional secret or its versions.
func (o SecretManagerRegionalSecretOutput) Topics() SecretManagerRegionalSecretTopicArrayOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) SecretManagerRegionalSecretTopicArrayOutput { return v.Topics }).(SecretManagerRegionalSecretTopicArrayOutput)
}

// The TTL for the regional secret. A duration in seconds with up to nine fractional digits, terminated by 's'. Example:
// "3.5s". Only one of 'ttl' or 'expire_time' can be provided.
func (o SecretManagerRegionalSecretOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can
// contain uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_') characters. An alias string
// must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given
// secret. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
func (o SecretManagerRegionalSecretOutput) VersionAliases() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringMapOutput { return v.VersionAliases }).(pulumi.StringMapOutput)
}

// Secret Version TTL after destruction request. This is a part of the delayed delete feature on Secret Version. For secret
// with versionDestroyTtl>0, version destruction doesn't happen immediately on calling destroy instead the version goes to
// a disabled state and the actual destruction happens after this TTL expires. It must be atleast 24h.
func (o SecretManagerRegionalSecretOutput) VersionDestroyTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecret) pulumi.StringPtrOutput { return v.VersionDestroyTtl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretManagerRegionalSecretInput)(nil)).Elem(), &SecretManagerRegionalSecret{})
	pulumi.RegisterOutputType(SecretManagerRegionalSecretOutput{})
}
