// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivilegedAccessManagerEntitlement(ctx *pulumi.Context, args *LookupPrivilegedAccessManagerEntitlementArgs, opts ...pulumi.InvokeOption) (*LookupPrivilegedAccessManagerEntitlementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupPrivilegedAccessManagerEntitlementResult
	err = ctx.InvokePackage("google-beta:index/getPrivilegedAccessManagerEntitlement:getPrivilegedAccessManagerEntitlement", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivilegedAccessManagerEntitlement.
type LookupPrivilegedAccessManagerEntitlementArgs struct {
	EntitlementId *string `pulumi:"entitlementId"`
	Id            *string `pulumi:"id"`
	Location      *string `pulumi:"location"`
	Parent        *string `pulumi:"parent"`
}

// A collection of values returned by getPrivilegedAccessManagerEntitlement.
type LookupPrivilegedAccessManagerEntitlementResult struct {
	AdditionalNotificationTargets []GetPrivilegedAccessManagerEntitlementAdditionalNotificationTarget `pulumi:"additionalNotificationTargets"`
	ApprovalWorkflows             []GetPrivilegedAccessManagerEntitlementApprovalWorkflow             `pulumi:"approvalWorkflows"`
	CreateTime                    string                                                              `pulumi:"createTime"`
	EligibleUsers                 []GetPrivilegedAccessManagerEntitlementEligibleUser                 `pulumi:"eligibleUsers"`
	EntitlementId                 *string                                                             `pulumi:"entitlementId"`
	Etag                          string                                                              `pulumi:"etag"`
	Id                            string                                                              `pulumi:"id"`
	Location                      *string                                                             `pulumi:"location"`
	MaxRequestDuration            string                                                              `pulumi:"maxRequestDuration"`
	Name                          string                                                              `pulumi:"name"`
	Parent                        *string                                                             `pulumi:"parent"`
	PrivilegedAccesses            []GetPrivilegedAccessManagerEntitlementPrivilegedAccess             `pulumi:"privilegedAccesses"`
	RequesterJustificationConfigs []GetPrivilegedAccessManagerEntitlementRequesterJustificationConfig `pulumi:"requesterJustificationConfigs"`
	State                         string                                                              `pulumi:"state"`
	UpdateTime                    string                                                              `pulumi:"updateTime"`
}

func LookupPrivilegedAccessManagerEntitlementOutput(ctx *pulumi.Context, args LookupPrivilegedAccessManagerEntitlementOutputArgs, opts ...pulumi.InvokeOption) LookupPrivilegedAccessManagerEntitlementResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPrivilegedAccessManagerEntitlementResultOutput, error) {
			args := v.(LookupPrivilegedAccessManagerEntitlementArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupPrivilegedAccessManagerEntitlementResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getPrivilegedAccessManagerEntitlement:getPrivilegedAccessManagerEntitlement", args, LookupPrivilegedAccessManagerEntitlementResultOutput{}, options).(LookupPrivilegedAccessManagerEntitlementResultOutput), nil
		}).(LookupPrivilegedAccessManagerEntitlementResultOutput)
}

// A collection of arguments for invoking getPrivilegedAccessManagerEntitlement.
type LookupPrivilegedAccessManagerEntitlementOutputArgs struct {
	EntitlementId pulumi.StringPtrInput `pulumi:"entitlementId"`
	Id            pulumi.StringPtrInput `pulumi:"id"`
	Location      pulumi.StringPtrInput `pulumi:"location"`
	Parent        pulumi.StringPtrInput `pulumi:"parent"`
}

func (LookupPrivilegedAccessManagerEntitlementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivilegedAccessManagerEntitlementArgs)(nil)).Elem()
}

// A collection of values returned by getPrivilegedAccessManagerEntitlement.
type LookupPrivilegedAccessManagerEntitlementResultOutput struct{ *pulumi.OutputState }

func (LookupPrivilegedAccessManagerEntitlementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivilegedAccessManagerEntitlementResult)(nil)).Elem()
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) ToLookupPrivilegedAccessManagerEntitlementResultOutput() LookupPrivilegedAccessManagerEntitlementResultOutput {
	return o
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) ToLookupPrivilegedAccessManagerEntitlementResultOutputWithContext(ctx context.Context) LookupPrivilegedAccessManagerEntitlementResultOutput {
	return o
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) AdditionalNotificationTargets() GetPrivilegedAccessManagerEntitlementAdditionalNotificationTargetArrayOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) []GetPrivilegedAccessManagerEntitlementAdditionalNotificationTarget {
		return v.AdditionalNotificationTargets
	}).(GetPrivilegedAccessManagerEntitlementAdditionalNotificationTargetArrayOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) ApprovalWorkflows() GetPrivilegedAccessManagerEntitlementApprovalWorkflowArrayOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) []GetPrivilegedAccessManagerEntitlementApprovalWorkflow {
		return v.ApprovalWorkflows
	}).(GetPrivilegedAccessManagerEntitlementApprovalWorkflowArrayOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) EligibleUsers() GetPrivilegedAccessManagerEntitlementEligibleUserArrayOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) []GetPrivilegedAccessManagerEntitlementEligibleUser {
		return v.EligibleUsers
	}).(GetPrivilegedAccessManagerEntitlementEligibleUserArrayOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) EntitlementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) *string { return v.EntitlementId }).(pulumi.StringPtrOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) MaxRequestDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) string { return v.MaxRequestDuration }).(pulumi.StringOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) *string { return v.Parent }).(pulumi.StringPtrOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) PrivilegedAccesses() GetPrivilegedAccessManagerEntitlementPrivilegedAccessArrayOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) []GetPrivilegedAccessManagerEntitlementPrivilegedAccess {
		return v.PrivilegedAccesses
	}).(GetPrivilegedAccessManagerEntitlementPrivilegedAccessArrayOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) RequesterJustificationConfigs() GetPrivilegedAccessManagerEntitlementRequesterJustificationConfigArrayOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) []GetPrivilegedAccessManagerEntitlementRequesterJustificationConfig {
		return v.RequesterJustificationConfigs
	}).(GetPrivilegedAccessManagerEntitlementRequesterJustificationConfigArrayOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupPrivilegedAccessManagerEntitlementResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegedAccessManagerEntitlementResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivilegedAccessManagerEntitlementResultOutput{})
}
