// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringAppEngineService(ctx *pulumi.Context, args *GetMonitoringAppEngineServiceArgs, opts ...pulumi.InvokeOption) (*GetMonitoringAppEngineServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetMonitoringAppEngineServiceResult
	err = ctx.InvokePackage("google-beta:index/getMonitoringAppEngineService:getMonitoringAppEngineService", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMonitoringAppEngineService.
type GetMonitoringAppEngineServiceArgs struct {
	Id       *string `pulumi:"id"`
	ModuleId string  `pulumi:"moduleId"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getMonitoringAppEngineService.
type GetMonitoringAppEngineServiceResult struct {
	DisplayName string                                   `pulumi:"displayName"`
	Id          string                                   `pulumi:"id"`
	ModuleId    string                                   `pulumi:"moduleId"`
	Name        string                                   `pulumi:"name"`
	Project     *string                                  `pulumi:"project"`
	ServiceId   string                                   `pulumi:"serviceId"`
	Telemetries []GetMonitoringAppEngineServiceTelemetry `pulumi:"telemetries"`
	UserLabels  map[string]string                        `pulumi:"userLabels"`
}

func GetMonitoringAppEngineServiceOutput(ctx *pulumi.Context, args GetMonitoringAppEngineServiceOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringAppEngineServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetMonitoringAppEngineServiceResultOutput, error) {
			args := v.(GetMonitoringAppEngineServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetMonitoringAppEngineServiceResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getMonitoringAppEngineService:getMonitoringAppEngineService", args, GetMonitoringAppEngineServiceResultOutput{}, options).(GetMonitoringAppEngineServiceResultOutput), nil
		}).(GetMonitoringAppEngineServiceResultOutput)
}

// A collection of arguments for invoking getMonitoringAppEngineService.
type GetMonitoringAppEngineServiceOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	ModuleId pulumi.StringInput    `pulumi:"moduleId"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (GetMonitoringAppEngineServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringAppEngineServiceArgs)(nil)).Elem()
}

// A collection of values returned by getMonitoringAppEngineService.
type GetMonitoringAppEngineServiceResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringAppEngineServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringAppEngineServiceResult)(nil)).Elem()
}

func (o GetMonitoringAppEngineServiceResultOutput) ToGetMonitoringAppEngineServiceResultOutput() GetMonitoringAppEngineServiceResultOutput {
	return o
}

func (o GetMonitoringAppEngineServiceResultOutput) ToGetMonitoringAppEngineServiceResultOutputWithContext(ctx context.Context) GetMonitoringAppEngineServiceResultOutput {
	return o
}

func (o GetMonitoringAppEngineServiceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GetMonitoringAppEngineServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMonitoringAppEngineServiceResultOutput) ModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) string { return v.ModuleId }).(pulumi.StringOutput)
}

func (o GetMonitoringAppEngineServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetMonitoringAppEngineServiceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o GetMonitoringAppEngineServiceResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func (o GetMonitoringAppEngineServiceResultOutput) Telemetries() GetMonitoringAppEngineServiceTelemetryArrayOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) []GetMonitoringAppEngineServiceTelemetry {
		return v.Telemetries
	}).(GetMonitoringAppEngineServiceTelemetryArrayOutput)
}

func (o GetMonitoringAppEngineServiceResultOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetMonitoringAppEngineServiceResult) map[string]string { return v.UserLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringAppEngineServiceResultOutput{})
}
