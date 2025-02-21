// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeGlobalForwardingRule(ctx *pulumi.Context, args *LookupComputeGlobalForwardingRuleArgs, opts ...pulumi.InvokeOption) (*LookupComputeGlobalForwardingRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeGlobalForwardingRuleResult
	err = ctx.InvokePackage("google-beta:index/getComputeGlobalForwardingRule:getComputeGlobalForwardingRule", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeGlobalForwardingRule.
type LookupComputeGlobalForwardingRuleArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getComputeGlobalForwardingRule.
type LookupComputeGlobalForwardingRuleResult struct {
	AllowPscGlobalAccess          bool                                                         `pulumi:"allowPscGlobalAccess"`
	BaseForwardingRule            string                                                       `pulumi:"baseForwardingRule"`
	Description                   string                                                       `pulumi:"description"`
	EffectiveLabels               map[string]string                                            `pulumi:"effectiveLabels"`
	ForwardingRuleId              float64                                                      `pulumi:"forwardingRuleId"`
	Id                            string                                                       `pulumi:"id"`
	IpAddress                     string                                                       `pulumi:"ipAddress"`
	IpProtocol                    string                                                       `pulumi:"ipProtocol"`
	IpVersion                     string                                                       `pulumi:"ipVersion"`
	LabelFingerprint              string                                                       `pulumi:"labelFingerprint"`
	Labels                        map[string]string                                            `pulumi:"labels"`
	LoadBalancingScheme           string                                                       `pulumi:"loadBalancingScheme"`
	MetadataFilters               []GetComputeGlobalForwardingRuleMetadataFilter               `pulumi:"metadataFilters"`
	Name                          string                                                       `pulumi:"name"`
	Network                       string                                                       `pulumi:"network"`
	NetworkTier                   string                                                       `pulumi:"networkTier"`
	NoAutomateDnsZone             bool                                                         `pulumi:"noAutomateDnsZone"`
	PortRange                     string                                                       `pulumi:"portRange"`
	Project                       *string                                                      `pulumi:"project"`
	PscConnectionId               string                                                       `pulumi:"pscConnectionId"`
	PscConnectionStatus           string                                                       `pulumi:"pscConnectionStatus"`
	SelfLink                      string                                                       `pulumi:"selfLink"`
	ServiceDirectoryRegistrations []GetComputeGlobalForwardingRuleServiceDirectoryRegistration `pulumi:"serviceDirectoryRegistrations"`
	SourceIpRanges                []string                                                     `pulumi:"sourceIpRanges"`
	Subnetwork                    string                                                       `pulumi:"subnetwork"`
	Target                        string                                                       `pulumi:"target"`
	TerraformLabels               map[string]string                                            `pulumi:"terraformLabels"`
}

func LookupComputeGlobalForwardingRuleOutput(ctx *pulumi.Context, args LookupComputeGlobalForwardingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupComputeGlobalForwardingRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeGlobalForwardingRuleResultOutput, error) {
			args := v.(LookupComputeGlobalForwardingRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeGlobalForwardingRuleResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeGlobalForwardingRule:getComputeGlobalForwardingRule", args, LookupComputeGlobalForwardingRuleResultOutput{}, options).(LookupComputeGlobalForwardingRuleResultOutput), nil
		}).(LookupComputeGlobalForwardingRuleResultOutput)
}

// A collection of arguments for invoking getComputeGlobalForwardingRule.
type LookupComputeGlobalForwardingRuleOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupComputeGlobalForwardingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeGlobalForwardingRuleArgs)(nil)).Elem()
}

// A collection of values returned by getComputeGlobalForwardingRule.
type LookupComputeGlobalForwardingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupComputeGlobalForwardingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeGlobalForwardingRuleResult)(nil)).Elem()
}

func (o LookupComputeGlobalForwardingRuleResultOutput) ToLookupComputeGlobalForwardingRuleResultOutput() LookupComputeGlobalForwardingRuleResultOutput {
	return o
}

func (o LookupComputeGlobalForwardingRuleResultOutput) ToLookupComputeGlobalForwardingRuleResultOutputWithContext(ctx context.Context) LookupComputeGlobalForwardingRuleResultOutput {
	return o
}

func (o LookupComputeGlobalForwardingRuleResultOutput) AllowPscGlobalAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) bool { return v.AllowPscGlobalAccess }).(pulumi.BoolOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) BaseForwardingRule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.BaseForwardingRule }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) ForwardingRuleId() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) float64 { return v.ForwardingRuleId }).(pulumi.Float64Output)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.IpProtocol }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) LoadBalancingScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.LoadBalancingScheme }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) MetadataFilters() GetComputeGlobalForwardingRuleMetadataFilterArrayOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) []GetComputeGlobalForwardingRuleMetadataFilter {
		return v.MetadataFilters
	}).(GetComputeGlobalForwardingRuleMetadataFilterArrayOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) NetworkTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.NetworkTier }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) NoAutomateDnsZone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) bool { return v.NoAutomateDnsZone }).(pulumi.BoolOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) PortRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.PortRange }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) PscConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.PscConnectionId }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) PscConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.PscConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) ServiceDirectoryRegistrations() GetComputeGlobalForwardingRuleServiceDirectoryRegistrationArrayOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) []GetComputeGlobalForwardingRuleServiceDirectoryRegistration {
		return v.ServiceDirectoryRegistrations
	}).(GetComputeGlobalForwardingRuleServiceDirectoryRegistrationArrayOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) SourceIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) []string { return v.SourceIpRanges }).(pulumi.StringArrayOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) string { return v.Target }).(pulumi.StringOutput)
}

func (o LookupComputeGlobalForwardingRuleResultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeGlobalForwardingRuleResult) map[string]string { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeGlobalForwardingRuleResultOutput{})
}
