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

type SiteVerificationOwner struct {
	pulumi.CustomResourceState

	// The email address of the owner.
	Email                   pulumi.StringOutput                    `pulumi:"email"`
	SiteVerificationOwnerId pulumi.StringOutput                    `pulumi:"siteVerificationOwnerId"`
	Timeouts                SiteVerificationOwnerTimeoutsPtrOutput `pulumi:"timeouts"`
	// The id of the Web Resource to add this owner to, in the form "webResource/<web-resource-id>".
	WebResourceId pulumi.StringOutput `pulumi:"webResourceId"`
}

// NewSiteVerificationOwner registers a new resource with the given unique name, arguments, and options.
func NewSiteVerificationOwner(ctx *pulumi.Context,
	name string, args *SiteVerificationOwnerArgs, opts ...pulumi.ResourceOption) (*SiteVerificationOwner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.WebResourceId == nil {
		return nil, errors.New("invalid value for required argument 'WebResourceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SiteVerificationOwner
	err = ctx.RegisterPackageResource("google-beta:index/siteVerificationOwner:SiteVerificationOwner", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteVerificationOwner gets an existing SiteVerificationOwner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteVerificationOwner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteVerificationOwnerState, opts ...pulumi.ResourceOption) (*SiteVerificationOwner, error) {
	var resource SiteVerificationOwner
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/siteVerificationOwner:SiteVerificationOwner", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteVerificationOwner resources.
type siteVerificationOwnerState struct {
	// The email address of the owner.
	Email                   *string                        `pulumi:"email"`
	SiteVerificationOwnerId *string                        `pulumi:"siteVerificationOwnerId"`
	Timeouts                *SiteVerificationOwnerTimeouts `pulumi:"timeouts"`
	// The id of the Web Resource to add this owner to, in the form "webResource/<web-resource-id>".
	WebResourceId *string `pulumi:"webResourceId"`
}

type SiteVerificationOwnerState struct {
	// The email address of the owner.
	Email                   pulumi.StringPtrInput
	SiteVerificationOwnerId pulumi.StringPtrInput
	Timeouts                SiteVerificationOwnerTimeoutsPtrInput
	// The id of the Web Resource to add this owner to, in the form "webResource/<web-resource-id>".
	WebResourceId pulumi.StringPtrInput
}

func (SiteVerificationOwnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteVerificationOwnerState)(nil)).Elem()
}

type siteVerificationOwnerArgs struct {
	// The email address of the owner.
	Email                   string                         `pulumi:"email"`
	SiteVerificationOwnerId *string                        `pulumi:"siteVerificationOwnerId"`
	Timeouts                *SiteVerificationOwnerTimeouts `pulumi:"timeouts"`
	// The id of the Web Resource to add this owner to, in the form "webResource/<web-resource-id>".
	WebResourceId string `pulumi:"webResourceId"`
}

// The set of arguments for constructing a SiteVerificationOwner resource.
type SiteVerificationOwnerArgs struct {
	// The email address of the owner.
	Email                   pulumi.StringInput
	SiteVerificationOwnerId pulumi.StringPtrInput
	Timeouts                SiteVerificationOwnerTimeoutsPtrInput
	// The id of the Web Resource to add this owner to, in the form "webResource/<web-resource-id>".
	WebResourceId pulumi.StringInput
}

func (SiteVerificationOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteVerificationOwnerArgs)(nil)).Elem()
}

type SiteVerificationOwnerInput interface {
	pulumi.Input

	ToSiteVerificationOwnerOutput() SiteVerificationOwnerOutput
	ToSiteVerificationOwnerOutputWithContext(ctx context.Context) SiteVerificationOwnerOutput
}

func (*SiteVerificationOwner) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteVerificationOwner)(nil)).Elem()
}

func (i *SiteVerificationOwner) ToSiteVerificationOwnerOutput() SiteVerificationOwnerOutput {
	return i.ToSiteVerificationOwnerOutputWithContext(context.Background())
}

func (i *SiteVerificationOwner) ToSiteVerificationOwnerOutputWithContext(ctx context.Context) SiteVerificationOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteVerificationOwnerOutput)
}

type SiteVerificationOwnerOutput struct{ *pulumi.OutputState }

func (SiteVerificationOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteVerificationOwner)(nil)).Elem()
}

func (o SiteVerificationOwnerOutput) ToSiteVerificationOwnerOutput() SiteVerificationOwnerOutput {
	return o
}

func (o SiteVerificationOwnerOutput) ToSiteVerificationOwnerOutputWithContext(ctx context.Context) SiteVerificationOwnerOutput {
	return o
}

// The email address of the owner.
func (o SiteVerificationOwnerOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteVerificationOwner) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o SiteVerificationOwnerOutput) SiteVerificationOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteVerificationOwner) pulumi.StringOutput { return v.SiteVerificationOwnerId }).(pulumi.StringOutput)
}

func (o SiteVerificationOwnerOutput) Timeouts() SiteVerificationOwnerTimeoutsPtrOutput {
	return o.ApplyT(func(v *SiteVerificationOwner) SiteVerificationOwnerTimeoutsPtrOutput { return v.Timeouts }).(SiteVerificationOwnerTimeoutsPtrOutput)
}

// The id of the Web Resource to add this owner to, in the form "webResource/<web-resource-id>".
func (o SiteVerificationOwnerOutput) WebResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteVerificationOwner) pulumi.StringOutput { return v.WebResourceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteVerificationOwnerInput)(nil)).Elem(), &SiteVerificationOwner{})
	pulumi.RegisterOutputType(SiteVerificationOwnerOutput{})
}
