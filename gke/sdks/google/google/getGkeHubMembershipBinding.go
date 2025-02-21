// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGkeHubMembershipBinding(ctx *pulumi.Context, args *LookupGkeHubMembershipBindingArgs, opts ...pulumi.InvokeOption) (*LookupGkeHubMembershipBindingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupGkeHubMembershipBindingResult
	err = ctx.InvokePackage("google:index/getGkeHubMembershipBinding:getGkeHubMembershipBinding", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGkeHubMembershipBinding.
type LookupGkeHubMembershipBindingArgs struct {
	Id                  *string `pulumi:"id"`
	Location            string  `pulumi:"location"`
	MembershipBindingId string  `pulumi:"membershipBindingId"`
	MembershipId        string  `pulumi:"membershipId"`
	Project             *string `pulumi:"project"`
}

// A collection of values returned by getGkeHubMembershipBinding.
type LookupGkeHubMembershipBindingResult struct {
	CreateTime          string                            `pulumi:"createTime"`
	DeleteTime          string                            `pulumi:"deleteTime"`
	EffectiveLabels     map[string]string                 `pulumi:"effectiveLabels"`
	Id                  string                            `pulumi:"id"`
	Labels              map[string]string                 `pulumi:"labels"`
	Location            string                            `pulumi:"location"`
	MembershipBindingId string                            `pulumi:"membershipBindingId"`
	MembershipId        string                            `pulumi:"membershipId"`
	Name                string                            `pulumi:"name"`
	Project             *string                           `pulumi:"project"`
	Scope               string                            `pulumi:"scope"`
	States              []GetGkeHubMembershipBindingState `pulumi:"states"`
	TerraformLabels     map[string]string                 `pulumi:"terraformLabels"`
	Uid                 string                            `pulumi:"uid"`
	UpdateTime          string                            `pulumi:"updateTime"`
}

func LookupGkeHubMembershipBindingOutput(ctx *pulumi.Context, args LookupGkeHubMembershipBindingOutputArgs, opts ...pulumi.InvokeOption) LookupGkeHubMembershipBindingResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGkeHubMembershipBindingResultOutput, error) {
			args := v.(LookupGkeHubMembershipBindingArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupGkeHubMembershipBindingResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getGkeHubMembershipBinding:getGkeHubMembershipBinding", args, LookupGkeHubMembershipBindingResultOutput{}, options).(LookupGkeHubMembershipBindingResultOutput), nil
		}).(LookupGkeHubMembershipBindingResultOutput)
}

// A collection of arguments for invoking getGkeHubMembershipBinding.
type LookupGkeHubMembershipBindingOutputArgs struct {
	Id                  pulumi.StringPtrInput `pulumi:"id"`
	Location            pulumi.StringInput    `pulumi:"location"`
	MembershipBindingId pulumi.StringInput    `pulumi:"membershipBindingId"`
	MembershipId        pulumi.StringInput    `pulumi:"membershipId"`
	Project             pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGkeHubMembershipBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubMembershipBindingArgs)(nil)).Elem()
}

// A collection of values returned by getGkeHubMembershipBinding.
type LookupGkeHubMembershipBindingResultOutput struct{ *pulumi.OutputState }

func (LookupGkeHubMembershipBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubMembershipBindingResult)(nil)).Elem()
}

func (o LookupGkeHubMembershipBindingResultOutput) ToLookupGkeHubMembershipBindingResultOutput() LookupGkeHubMembershipBindingResultOutput {
	return o
}

func (o LookupGkeHubMembershipBindingResultOutput) ToLookupGkeHubMembershipBindingResultOutputWithContext(ctx context.Context) LookupGkeHubMembershipBindingResultOutput {
	return o
}

func (o LookupGkeHubMembershipBindingResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) MembershipBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.MembershipBindingId }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) MembershipId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.MembershipId }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) States() GetGkeHubMembershipBindingStateArrayOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) []GetGkeHubMembershipBindingState { return v.States }).(GetGkeHubMembershipBindingStateArrayOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) map[string]string { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.Uid }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipBindingResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipBindingResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGkeHubMembershipBindingResultOutput{})
}
