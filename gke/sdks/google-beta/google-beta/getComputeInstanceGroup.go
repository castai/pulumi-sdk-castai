// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeInstanceGroup(ctx *pulumi.Context, args *LookupComputeInstanceGroupArgs, opts ...pulumi.InvokeOption) (*LookupComputeInstanceGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeInstanceGroupResult
	err = ctx.InvokePackage("google-beta:index/getComputeInstanceGroup:getComputeInstanceGroup", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeInstanceGroup.
type LookupComputeInstanceGroupArgs struct {
	Id       *string `pulumi:"id"`
	Name     *string `pulumi:"name"`
	Project  *string `pulumi:"project"`
	SelfLink *string `pulumi:"selfLink"`
	Zone     *string `pulumi:"zone"`
}

// A collection of values returned by getComputeInstanceGroup.
type LookupComputeInstanceGroupResult struct {
	Description string                                 `pulumi:"description"`
	Id          string                                 `pulumi:"id"`
	Instances   []string                               `pulumi:"instances"`
	Name        *string                                `pulumi:"name"`
	NamedPorts  []GetComputeInstanceGroupNamedPortType `pulumi:"namedPorts"`
	Network     string                                 `pulumi:"network"`
	Project     string                                 `pulumi:"project"`
	SelfLink    string                                 `pulumi:"selfLink"`
	Size        float64                                `pulumi:"size"`
	Zone        string                                 `pulumi:"zone"`
}

func LookupComputeInstanceGroupOutput(ctx *pulumi.Context, args LookupComputeInstanceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupComputeInstanceGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeInstanceGroupResultOutput, error) {
			args := v.(LookupComputeInstanceGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeInstanceGroupResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeInstanceGroup:getComputeInstanceGroup", args, LookupComputeInstanceGroupResultOutput{}, options).(LookupComputeInstanceGroupResultOutput), nil
		}).(LookupComputeInstanceGroupResultOutput)
}

// A collection of arguments for invoking getComputeInstanceGroup.
type LookupComputeInstanceGroupOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	SelfLink pulumi.StringPtrInput `pulumi:"selfLink"`
	Zone     pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupComputeInstanceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeInstanceGroupArgs)(nil)).Elem()
}

// A collection of values returned by getComputeInstanceGroup.
type LookupComputeInstanceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupComputeInstanceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeInstanceGroupResult)(nil)).Elem()
}

func (o LookupComputeInstanceGroupResultOutput) ToLookupComputeInstanceGroupResultOutput() LookupComputeInstanceGroupResultOutput {
	return o
}

func (o LookupComputeInstanceGroupResultOutput) ToLookupComputeInstanceGroupResultOutputWithContext(ctx context.Context) LookupComputeInstanceGroupResultOutput {
	return o
}

func (o LookupComputeInstanceGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceGroupResultOutput) Instances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) []string { return v.Instances }).(pulumi.StringArrayOutput)
}

func (o LookupComputeInstanceGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupComputeInstanceGroupResultOutput) NamedPorts() GetComputeInstanceGroupNamedPortTypeArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) []GetComputeInstanceGroupNamedPortType { return v.NamedPorts }).(GetComputeInstanceGroupNamedPortTypeArrayOutput)
}

func (o LookupComputeInstanceGroupResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceGroupResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceGroupResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceGroupResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) float64 { return v.Size }).(pulumi.Float64Output)
}

func (o LookupComputeInstanceGroupResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceGroupResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeInstanceGroupResultOutput{})
}
