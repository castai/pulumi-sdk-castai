// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGkeHubMembershipIamPolicy(ctx *pulumi.Context, args *LookupGkeHubMembershipIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGkeHubMembershipIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupGkeHubMembershipIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getGkeHubMembershipIamPolicy:getGkeHubMembershipIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGkeHubMembershipIamPolicy.
type LookupGkeHubMembershipIamPolicyArgs struct {
	Id           *string `pulumi:"id"`
	Location     *string `pulumi:"location"`
	MembershipId string  `pulumi:"membershipId"`
	Project      *string `pulumi:"project"`
}

// A collection of values returned by getGkeHubMembershipIamPolicy.
type LookupGkeHubMembershipIamPolicyResult struct {
	Etag         string `pulumi:"etag"`
	Id           string `pulumi:"id"`
	Location     string `pulumi:"location"`
	MembershipId string `pulumi:"membershipId"`
	PolicyData   string `pulumi:"policyData"`
	Project      string `pulumi:"project"`
}

func LookupGkeHubMembershipIamPolicyOutput(ctx *pulumi.Context, args LookupGkeHubMembershipIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGkeHubMembershipIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGkeHubMembershipIamPolicyResultOutput, error) {
			args := v.(LookupGkeHubMembershipIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupGkeHubMembershipIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getGkeHubMembershipIamPolicy:getGkeHubMembershipIamPolicy", args, LookupGkeHubMembershipIamPolicyResultOutput{}, options).(LookupGkeHubMembershipIamPolicyResultOutput), nil
		}).(LookupGkeHubMembershipIamPolicyResultOutput)
}

// A collection of arguments for invoking getGkeHubMembershipIamPolicy.
type LookupGkeHubMembershipIamPolicyOutputArgs struct {
	Id           pulumi.StringPtrInput `pulumi:"id"`
	Location     pulumi.StringPtrInput `pulumi:"location"`
	MembershipId pulumi.StringInput    `pulumi:"membershipId"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGkeHubMembershipIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubMembershipIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getGkeHubMembershipIamPolicy.
type LookupGkeHubMembershipIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGkeHubMembershipIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubMembershipIamPolicyResult)(nil)).Elem()
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) ToLookupGkeHubMembershipIamPolicyResultOutput() LookupGkeHubMembershipIamPolicyResultOutput {
	return o
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) ToLookupGkeHubMembershipIamPolicyResultOutputWithContext(ctx context.Context) LookupGkeHubMembershipIamPolicyResultOutput {
	return o
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) MembershipId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipIamPolicyResult) string { return v.MembershipId }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupGkeHubMembershipIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubMembershipIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGkeHubMembershipIamPolicyResultOutput{})
}
