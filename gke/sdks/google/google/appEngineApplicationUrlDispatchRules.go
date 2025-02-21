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

type AppEngineApplicationUrlDispatchRules struct {
	pulumi.CustomResourceState

	AppEngineApplicationUrlDispatchRulesId pulumi.StringOutput `pulumi:"appEngineApplicationUrlDispatchRulesId"`
	// Rules to match an HTTP request and dispatch that request to a service.
	DispatchRules AppEngineApplicationUrlDispatchRulesDispatchRuleArrayOutput `pulumi:"dispatchRules"`
	Project       pulumi.StringOutput                                         `pulumi:"project"`
	Timeouts      AppEngineApplicationUrlDispatchRulesTimeoutsPtrOutput       `pulumi:"timeouts"`
}

// NewAppEngineApplicationUrlDispatchRules registers a new resource with the given unique name, arguments, and options.
func NewAppEngineApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, args *AppEngineApplicationUrlDispatchRulesArgs, opts ...pulumi.ResourceOption) (*AppEngineApplicationUrlDispatchRules, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DispatchRules == nil {
		return nil, errors.New("invalid value for required argument 'DispatchRules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource AppEngineApplicationUrlDispatchRules
	err = ctx.RegisterPackageResource("google:index/appEngineApplicationUrlDispatchRules:AppEngineApplicationUrlDispatchRules", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineApplicationUrlDispatchRules gets an existing AppEngineApplicationUrlDispatchRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineApplicationUrlDispatchRulesState, opts ...pulumi.ResourceOption) (*AppEngineApplicationUrlDispatchRules, error) {
	var resource AppEngineApplicationUrlDispatchRules
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/appEngineApplicationUrlDispatchRules:AppEngineApplicationUrlDispatchRules", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineApplicationUrlDispatchRules resources.
type appEngineApplicationUrlDispatchRulesState struct {
	AppEngineApplicationUrlDispatchRulesId *string `pulumi:"appEngineApplicationUrlDispatchRulesId"`
	// Rules to match an HTTP request and dispatch that request to a service.
	DispatchRules []AppEngineApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	Project       *string                                            `pulumi:"project"`
	Timeouts      *AppEngineApplicationUrlDispatchRulesTimeouts      `pulumi:"timeouts"`
}

type AppEngineApplicationUrlDispatchRulesState struct {
	AppEngineApplicationUrlDispatchRulesId pulumi.StringPtrInput
	// Rules to match an HTTP request and dispatch that request to a service.
	DispatchRules AppEngineApplicationUrlDispatchRulesDispatchRuleArrayInput
	Project       pulumi.StringPtrInput
	Timeouts      AppEngineApplicationUrlDispatchRulesTimeoutsPtrInput
}

func (AppEngineApplicationUrlDispatchRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineApplicationUrlDispatchRulesState)(nil)).Elem()
}

type appEngineApplicationUrlDispatchRulesArgs struct {
	AppEngineApplicationUrlDispatchRulesId *string `pulumi:"appEngineApplicationUrlDispatchRulesId"`
	// Rules to match an HTTP request and dispatch that request to a service.
	DispatchRules []AppEngineApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	Project       *string                                            `pulumi:"project"`
	Timeouts      *AppEngineApplicationUrlDispatchRulesTimeouts      `pulumi:"timeouts"`
}

// The set of arguments for constructing a AppEngineApplicationUrlDispatchRules resource.
type AppEngineApplicationUrlDispatchRulesArgs struct {
	AppEngineApplicationUrlDispatchRulesId pulumi.StringPtrInput
	// Rules to match an HTTP request and dispatch that request to a service.
	DispatchRules AppEngineApplicationUrlDispatchRulesDispatchRuleArrayInput
	Project       pulumi.StringPtrInput
	Timeouts      AppEngineApplicationUrlDispatchRulesTimeoutsPtrInput
}

func (AppEngineApplicationUrlDispatchRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineApplicationUrlDispatchRulesArgs)(nil)).Elem()
}

type AppEngineApplicationUrlDispatchRulesInput interface {
	pulumi.Input

	ToAppEngineApplicationUrlDispatchRulesOutput() AppEngineApplicationUrlDispatchRulesOutput
	ToAppEngineApplicationUrlDispatchRulesOutputWithContext(ctx context.Context) AppEngineApplicationUrlDispatchRulesOutput
}

func (*AppEngineApplicationUrlDispatchRules) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineApplicationUrlDispatchRules)(nil)).Elem()
}

func (i *AppEngineApplicationUrlDispatchRules) ToAppEngineApplicationUrlDispatchRulesOutput() AppEngineApplicationUrlDispatchRulesOutput {
	return i.ToAppEngineApplicationUrlDispatchRulesOutputWithContext(context.Background())
}

func (i *AppEngineApplicationUrlDispatchRules) ToAppEngineApplicationUrlDispatchRulesOutputWithContext(ctx context.Context) AppEngineApplicationUrlDispatchRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineApplicationUrlDispatchRulesOutput)
}

type AppEngineApplicationUrlDispatchRulesOutput struct{ *pulumi.OutputState }

func (AppEngineApplicationUrlDispatchRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineApplicationUrlDispatchRules)(nil)).Elem()
}

func (o AppEngineApplicationUrlDispatchRulesOutput) ToAppEngineApplicationUrlDispatchRulesOutput() AppEngineApplicationUrlDispatchRulesOutput {
	return o
}

func (o AppEngineApplicationUrlDispatchRulesOutput) ToAppEngineApplicationUrlDispatchRulesOutputWithContext(ctx context.Context) AppEngineApplicationUrlDispatchRulesOutput {
	return o
}

func (o AppEngineApplicationUrlDispatchRulesOutput) AppEngineApplicationUrlDispatchRulesId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineApplicationUrlDispatchRules) pulumi.StringOutput {
		return v.AppEngineApplicationUrlDispatchRulesId
	}).(pulumi.StringOutput)
}

// Rules to match an HTTP request and dispatch that request to a service.
func (o AppEngineApplicationUrlDispatchRulesOutput) DispatchRules() AppEngineApplicationUrlDispatchRulesDispatchRuleArrayOutput {
	return o.ApplyT(func(v *AppEngineApplicationUrlDispatchRules) AppEngineApplicationUrlDispatchRulesDispatchRuleArrayOutput {
		return v.DispatchRules
	}).(AppEngineApplicationUrlDispatchRulesDispatchRuleArrayOutput)
}

func (o AppEngineApplicationUrlDispatchRulesOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppEngineApplicationUrlDispatchRules) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o AppEngineApplicationUrlDispatchRulesOutput) Timeouts() AppEngineApplicationUrlDispatchRulesTimeoutsPtrOutput {
	return o.ApplyT(func(v *AppEngineApplicationUrlDispatchRules) AppEngineApplicationUrlDispatchRulesTimeoutsPtrOutput {
		return v.Timeouts
	}).(AppEngineApplicationUrlDispatchRulesTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineApplicationUrlDispatchRulesInput)(nil)).Elem(), &AppEngineApplicationUrlDispatchRules{})
	pulumi.RegisterOutputType(AppEngineApplicationUrlDispatchRulesOutput{})
}
