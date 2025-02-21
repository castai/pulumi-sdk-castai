// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirebaseHostingSite struct {
	pulumi.CustomResourceState

	// Optional. The [ID of a Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
	// associated with the Hosting site.
	AppId pulumi.StringPtrOutput `pulumi:"appId"`
	// The default URL for the site in the form of https://{name}.web.app
	DefaultUrl            pulumi.StringOutput `pulumi:"defaultUrl"`
	FirebaseHostingSiteId pulumi.StringOutput `pulumi:"firebaseHostingSiteId"`
	// Output only. The fully-qualified resource name of the Hosting site, in the format:
	// projects/PROJECT_IDENTIFIER/sites/SITE_ID PROJECT_IDENTIFIER: the Firebase project's
	// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// ***(recommended)*** or its
	// ['ProjectId'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_id).
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. Immutable. A globally unique identifier for the Hosting site. This identifier is used to construct the
	// Firebase-provisioned subdomains for the site, so it must also be a valid domain name label.
	SiteId   pulumi.StringPtrOutput               `pulumi:"siteId"`
	Timeouts FirebaseHostingSiteTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewFirebaseHostingSite registers a new resource with the given unique name, arguments, and options.
func NewFirebaseHostingSite(ctx *pulumi.Context,
	name string, args *FirebaseHostingSiteArgs, opts ...pulumi.ResourceOption) (*FirebaseHostingSite, error) {
	if args == nil {
		args = &FirebaseHostingSiteArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseHostingSite
	err = ctx.RegisterPackageResource("google-beta:index/firebaseHostingSite:FirebaseHostingSite", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseHostingSite gets an existing FirebaseHostingSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseHostingSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseHostingSiteState, opts ...pulumi.ResourceOption) (*FirebaseHostingSite, error) {
	var resource FirebaseHostingSite
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/firebaseHostingSite:FirebaseHostingSite", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseHostingSite resources.
type firebaseHostingSiteState struct {
	// Optional. The [ID of a Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
	// associated with the Hosting site.
	AppId *string `pulumi:"appId"`
	// The default URL for the site in the form of https://{name}.web.app
	DefaultUrl            *string `pulumi:"defaultUrl"`
	FirebaseHostingSiteId *string `pulumi:"firebaseHostingSiteId"`
	// Output only. The fully-qualified resource name of the Hosting site, in the format:
	// projects/PROJECT_IDENTIFIER/sites/SITE_ID PROJECT_IDENTIFIER: the Firebase project's
	// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// ***(recommended)*** or its
	// ['ProjectId'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_id).
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. Immutable. A globally unique identifier for the Hosting site. This identifier is used to construct the
	// Firebase-provisioned subdomains for the site, so it must also be a valid domain name label.
	SiteId   *string                      `pulumi:"siteId"`
	Timeouts *FirebaseHostingSiteTimeouts `pulumi:"timeouts"`
}

type FirebaseHostingSiteState struct {
	// Optional. The [ID of a Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
	// associated with the Hosting site.
	AppId pulumi.StringPtrInput
	// The default URL for the site in the form of https://{name}.web.app
	DefaultUrl            pulumi.StringPtrInput
	FirebaseHostingSiteId pulumi.StringPtrInput
	// Output only. The fully-qualified resource name of the Hosting site, in the format:
	// projects/PROJECT_IDENTIFIER/sites/SITE_ID PROJECT_IDENTIFIER: the Firebase project's
	// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// ***(recommended)*** or its
	// ['ProjectId'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_id).
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. Immutable. A globally unique identifier for the Hosting site. This identifier is used to construct the
	// Firebase-provisioned subdomains for the site, so it must also be a valid domain name label.
	SiteId   pulumi.StringPtrInput
	Timeouts FirebaseHostingSiteTimeoutsPtrInput
}

func (FirebaseHostingSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseHostingSiteState)(nil)).Elem()
}

type firebaseHostingSiteArgs struct {
	// Optional. The [ID of a Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
	// associated with the Hosting site.
	AppId                 *string `pulumi:"appId"`
	FirebaseHostingSiteId *string `pulumi:"firebaseHostingSiteId"`
	Project               *string `pulumi:"project"`
	// Required. Immutable. A globally unique identifier for the Hosting site. This identifier is used to construct the
	// Firebase-provisioned subdomains for the site, so it must also be a valid domain name label.
	SiteId   *string                      `pulumi:"siteId"`
	Timeouts *FirebaseHostingSiteTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a FirebaseHostingSite resource.
type FirebaseHostingSiteArgs struct {
	// Optional. The [ID of a Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
	// associated with the Hosting site.
	AppId                 pulumi.StringPtrInput
	FirebaseHostingSiteId pulumi.StringPtrInput
	Project               pulumi.StringPtrInput
	// Required. Immutable. A globally unique identifier for the Hosting site. This identifier is used to construct the
	// Firebase-provisioned subdomains for the site, so it must also be a valid domain name label.
	SiteId   pulumi.StringPtrInput
	Timeouts FirebaseHostingSiteTimeoutsPtrInput
}

func (FirebaseHostingSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseHostingSiteArgs)(nil)).Elem()
}

type FirebaseHostingSiteInput interface {
	pulumi.Input

	ToFirebaseHostingSiteOutput() FirebaseHostingSiteOutput
	ToFirebaseHostingSiteOutputWithContext(ctx context.Context) FirebaseHostingSiteOutput
}

func (*FirebaseHostingSite) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseHostingSite)(nil)).Elem()
}

func (i *FirebaseHostingSite) ToFirebaseHostingSiteOutput() FirebaseHostingSiteOutput {
	return i.ToFirebaseHostingSiteOutputWithContext(context.Background())
}

func (i *FirebaseHostingSite) ToFirebaseHostingSiteOutputWithContext(ctx context.Context) FirebaseHostingSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseHostingSiteOutput)
}

type FirebaseHostingSiteOutput struct{ *pulumi.OutputState }

func (FirebaseHostingSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseHostingSite)(nil)).Elem()
}

func (o FirebaseHostingSiteOutput) ToFirebaseHostingSiteOutput() FirebaseHostingSiteOutput {
	return o
}

func (o FirebaseHostingSiteOutput) ToFirebaseHostingSiteOutputWithContext(ctx context.Context) FirebaseHostingSiteOutput {
	return o
}

// Optional. The [ID of a Web
// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
// associated with the Hosting site.
func (o FirebaseHostingSiteOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirebaseHostingSite) pulumi.StringPtrOutput { return v.AppId }).(pulumi.StringPtrOutput)
}

// The default URL for the site in the form of https://{name}.web.app
func (o FirebaseHostingSiteOutput) DefaultUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingSite) pulumi.StringOutput { return v.DefaultUrl }).(pulumi.StringOutput)
}

func (o FirebaseHostingSiteOutput) FirebaseHostingSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingSite) pulumi.StringOutput { return v.FirebaseHostingSiteId }).(pulumi.StringOutput)
}

// Output only. The fully-qualified resource name of the Hosting site, in the format:
// projects/PROJECT_IDENTIFIER/sites/SITE_ID PROJECT_IDENTIFIER: the Firebase project's
// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
// ***(recommended)*** or its
// ['ProjectId'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_id).
// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
func (o FirebaseHostingSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirebaseHostingSiteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseHostingSite) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. Immutable. A globally unique identifier for the Hosting site. This identifier is used to construct the
// Firebase-provisioned subdomains for the site, so it must also be a valid domain name label.
func (o FirebaseHostingSiteOutput) SiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirebaseHostingSite) pulumi.StringPtrOutput { return v.SiteId }).(pulumi.StringPtrOutput)
}

func (o FirebaseHostingSiteOutput) Timeouts() FirebaseHostingSiteTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseHostingSite) FirebaseHostingSiteTimeoutsPtrOutput { return v.Timeouts }).(FirebaseHostingSiteTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseHostingSiteInput)(nil)).Elem(), &FirebaseHostingSite{})
	pulumi.RegisterOutputType(FirebaseHostingSiteOutput{})
}
