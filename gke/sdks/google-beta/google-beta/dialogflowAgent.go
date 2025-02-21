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

type DialogflowAgent struct {
	pulumi.CustomResourceState

	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query different
	// service endpoints for different API versions. However, bots connectors and webhook calls will follow the specified API
	// version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1 API. Possible
	// values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into this
	// field, the Dialogflow will save the image in the backend. The address of the backend image returned from the API will be
	// shown in the [avatarUriBackend] field.
	AvatarUri pulumi.StringPtrOutput `pulumi:"avatarUri"`
	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar, the
	// [avatarUri] field can be used.
	AvatarUriBackend pulumi.StringOutput `pulumi:"avatarUriBackend"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can
	// tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a
	// fallback intent will be triggered or, if there are no fallback intents defined, no intent will be triggered. The score
	// values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold pulumi.Float64PtrOutput `pulumi:"classificationThreshold"`
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringOutput `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description       pulumi.StringPtrOutput `pulumi:"description"`
	DialogflowAgentId pulumi.StringOutput    `pulumi:"dialogflowAgentId"`
	// The name of this agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Determines whether this agent should log conversation queries.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
	// examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used for
	// agents with a large number of examples in intents, especially the ones using @sys.any or very large developer entities.
	// Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
	MatchMode pulumi.StringOutput `pulumi:"matchMode"`
	Project   pulumi.StringOutput `pulumi:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes pulumi.StringArrayOutput `pulumi:"supportedLanguageCodes"`
	Tier                   pulumi.StringPtrOutput   `pulumi:"tier"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringOutput              `pulumi:"timeZone"`
	Timeouts DialogflowAgentTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDialogflowAgent registers a new resource with the given unique name, arguments, and options.
func NewDialogflowAgent(ctx *pulumi.Context,
	name string, args *DialogflowAgentArgs, opts ...pulumi.ResourceOption) (*DialogflowAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultLanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'DefaultLanguageCode'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DialogflowAgent
	err = ctx.RegisterPackageResource("google-beta:index/dialogflowAgent:DialogflowAgent", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDialogflowAgent gets an existing DialogflowAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDialogflowAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DialogflowAgentState, opts ...pulumi.ResourceOption) (*DialogflowAgent, error) {
	var resource DialogflowAgent
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dialogflowAgent:DialogflowAgent", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DialogflowAgent resources.
type dialogflowAgentState struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query different
	// service endpoints for different API versions. However, bots connectors and webhook calls will follow the specified API
	// version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1 API. Possible
	// values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
	ApiVersion *string `pulumi:"apiVersion"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into this
	// field, the Dialogflow will save the image in the backend. The address of the backend image returned from the API will be
	// shown in the [avatarUriBackend] field.
	AvatarUri *string `pulumi:"avatarUri"`
	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar, the
	// [avatarUri] field can be used.
	AvatarUriBackend *string `pulumi:"avatarUriBackend"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can
	// tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a
	// fallback intent will be triggered or, if there are no fallback intents defined, no intent will be triggered. The score
	// values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold *float64 `pulumi:"classificationThreshold"`
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode *string `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description       *string `pulumi:"description"`
	DialogflowAgentId *string `pulumi:"dialogflowAgentId"`
	// The name of this agent.
	DisplayName *string `pulumi:"displayName"`
	// Determines whether this agent should log conversation queries.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
	// examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used for
	// agents with a large number of examples in intents, especially the ones using @sys.any or very large developer entities.
	// Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
	MatchMode *string `pulumi:"matchMode"`
	Project   *string `pulumi:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	Tier                   *string  `pulumi:"tier"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone *string                  `pulumi:"timeZone"`
	Timeouts *DialogflowAgentTimeouts `pulumi:"timeouts"`
}

type DialogflowAgentState struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query different
	// service endpoints for different API versions. However, bots connectors and webhook calls will follow the specified API
	// version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1 API. Possible
	// values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
	ApiVersion pulumi.StringPtrInput
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into this
	// field, the Dialogflow will save the image in the backend. The address of the backend image returned from the API will be
	// shown in the [avatarUriBackend] field.
	AvatarUri pulumi.StringPtrInput
	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar, the
	// [avatarUri] field can be used.
	AvatarUriBackend pulumi.StringPtrInput
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can
	// tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a
	// fallback intent will be triggered or, if there are no fallback intents defined, no intent will be triggered. The score
	// values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold pulumi.Float64PtrInput
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringPtrInput
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description       pulumi.StringPtrInput
	DialogflowAgentId pulumi.StringPtrInput
	// The name of this agent.
	DisplayName pulumi.StringPtrInput
	// Determines whether this agent should log conversation queries.
	EnableLogging pulumi.BoolPtrInput
	// Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
	// examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used for
	// agents with a large number of examples in intents, especially the ones using @sys.any or very large developer entities.
	// Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
	MatchMode pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes pulumi.StringArrayInput
	Tier                   pulumi.StringPtrInput
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringPtrInput
	Timeouts DialogflowAgentTimeoutsPtrInput
}

func (DialogflowAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowAgentState)(nil)).Elem()
}

type dialogflowAgentArgs struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query different
	// service endpoints for different API versions. However, bots connectors and webhook calls will follow the specified API
	// version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1 API. Possible
	// values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
	ApiVersion *string `pulumi:"apiVersion"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into this
	// field, the Dialogflow will save the image in the backend. The address of the backend image returned from the API will be
	// shown in the [avatarUriBackend] field.
	AvatarUri *string `pulumi:"avatarUri"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can
	// tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a
	// fallback intent will be triggered or, if there are no fallback intents defined, no intent will be triggered. The score
	// values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold *float64 `pulumi:"classificationThreshold"`
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode string `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description       *string `pulumi:"description"`
	DialogflowAgentId *string `pulumi:"dialogflowAgentId"`
	// The name of this agent.
	DisplayName string `pulumi:"displayName"`
	// Determines whether this agent should log conversation queries.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
	// examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used for
	// agents with a large number of examples in intents, especially the ones using @sys.any or very large developer entities.
	// Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
	MatchMode *string `pulumi:"matchMode"`
	Project   *string `pulumi:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	Tier                   *string  `pulumi:"tier"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone string                   `pulumi:"timeZone"`
	Timeouts *DialogflowAgentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DialogflowAgent resource.
type DialogflowAgentArgs struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query different
	// service endpoints for different API versions. However, bots connectors and webhook calls will follow the specified API
	// version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1 API. Possible
	// values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
	ApiVersion pulumi.StringPtrInput
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into this
	// field, the Dialogflow will save the image in the backend. The address of the backend image returned from the API will be
	// shown in the [avatarUriBackend] field.
	AvatarUri pulumi.StringPtrInput
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can
	// tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a
	// fallback intent will be triggered or, if there are no fallback intents defined, no intent will be triggered. The score
	// values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold pulumi.Float64PtrInput
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringInput
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description       pulumi.StringPtrInput
	DialogflowAgentId pulumi.StringPtrInput
	// The name of this agent.
	DisplayName pulumi.StringInput
	// Determines whether this agent should log conversation queries.
	EnableLogging pulumi.BoolPtrInput
	// Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
	// examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used for
	// agents with a large number of examples in intents, especially the ones using @sys.any or very large developer entities.
	// Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
	MatchMode pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes pulumi.StringArrayInput
	Tier                   pulumi.StringPtrInput
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringInput
	Timeouts DialogflowAgentTimeoutsPtrInput
}

func (DialogflowAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowAgentArgs)(nil)).Elem()
}

type DialogflowAgentInput interface {
	pulumi.Input

	ToDialogflowAgentOutput() DialogflowAgentOutput
	ToDialogflowAgentOutputWithContext(ctx context.Context) DialogflowAgentOutput
}

func (*DialogflowAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowAgent)(nil)).Elem()
}

func (i *DialogflowAgent) ToDialogflowAgentOutput() DialogflowAgentOutput {
	return i.ToDialogflowAgentOutputWithContext(context.Background())
}

func (i *DialogflowAgent) ToDialogflowAgentOutputWithContext(ctx context.Context) DialogflowAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DialogflowAgentOutput)
}

type DialogflowAgentOutput struct{ *pulumi.OutputState }

func (DialogflowAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowAgent)(nil)).Elem()
}

func (o DialogflowAgentOutput) ToDialogflowAgentOutput() DialogflowAgentOutput {
	return o
}

func (o DialogflowAgentOutput) ToDialogflowAgentOutputWithContext(ctx context.Context) DialogflowAgentOutput {
	return o
}

// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query different
// service endpoints for different API versions. However, bots connectors and webhook calls will follow the specified API
// version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1 API. Possible
// values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
func (o DialogflowAgentOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into this
// field, the Dialogflow will save the image in the backend. The address of the backend image returned from the API will be
// shown in the [avatarUriBackend] field.
func (o DialogflowAgentOutput) AvatarUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringPtrOutput { return v.AvatarUri }).(pulumi.StringPtrOutput)
}

// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar, the
// [avatarUri] field can be used.
func (o DialogflowAgentOutput) AvatarUriBackend() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.AvatarUriBackend }).(pulumi.StringOutput)
}

// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can
// tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a
// fallback intent will be triggered or, if there are no fallback intents defined, no intent will be triggered. The score
// values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
func (o DialogflowAgentOutput) ClassificationThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.Float64PtrOutput { return v.ClassificationThreshold }).(pulumi.Float64PtrOutput)
}

// The default language of the agent as a language tag. [See Language
// Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
// codes. This field cannot be updated after creation.
func (o DialogflowAgentOutput) DefaultLanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.DefaultLanguageCode }).(pulumi.StringOutput)
}

// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
func (o DialogflowAgentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DialogflowAgentOutput) DialogflowAgentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.DialogflowAgentId }).(pulumi.StringOutput)
}

// The name of this agent.
func (o DialogflowAgentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this agent should log conversation queries.
func (o DialogflowAgentOutput) EnableLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.BoolPtrOutput { return v.EnableLogging }).(pulumi.BoolPtrOutput)
}

// Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
// examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used for
// agents with a large number of examples in intents, especially the ones using @sys.any or very large developer entities.
// Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
func (o DialogflowAgentOutput) MatchMode() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.MatchMode }).(pulumi.StringOutput)
}

func (o DialogflowAgentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The list of all languages supported by this agent (except for the defaultLanguageCode).
func (o DialogflowAgentOutput) SupportedLanguageCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringArrayOutput { return v.SupportedLanguageCodes }).(pulumi.StringArrayOutput)
}

func (o DialogflowAgentOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringPtrOutput { return v.Tier }).(pulumi.StringPtrOutput)
}

// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
// Europe/Paris.
func (o DialogflowAgentOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowAgent) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

func (o DialogflowAgentOutput) Timeouts() DialogflowAgentTimeoutsPtrOutput {
	return o.ApplyT(func(v *DialogflowAgent) DialogflowAgentTimeoutsPtrOutput { return v.Timeouts }).(DialogflowAgentTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DialogflowAgentInput)(nil)).Elem(), &DialogflowAgent{})
	pulumi.RegisterOutputType(DialogflowAgentOutput{})
}
