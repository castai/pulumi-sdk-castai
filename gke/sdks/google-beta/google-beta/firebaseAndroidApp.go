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

type FirebaseAndroidApp struct {
	pulumi.CustomResourceState

	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp. If
	// apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp. This
	// auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	ApiKeyId pulumi.StringOutput `pulumi:"apiKeyId"`
	// The globally unique, Firebase-assigned identifier of the AndroidApp. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId          pulumi.StringOutput    `pulumi:"appId"`
	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// The user-assigned display name of the AndroidApp.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// This checksum is computed by the server based on the value of other fields, and it may be sent with update requests to
	// ensure the client has an up-to-date value before proceeding.
	Etag                 pulumi.StringOutput `pulumi:"etag"`
	FirebaseAndroidAppId pulumi.StringOutput `pulumi:"firebaseAndroidAppId"`
	// The fully qualified resource name of the AndroidApp, for example: projects/projectId/androidApps/appId
	Name pulumi.StringOutput `pulumi:"name"`
	// The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName pulumi.StringOutput `pulumi:"packageName"`
	Project     pulumi.StringOutput `pulumi:"project"`
	// The SHA1 certificate hashes for the AndroidApp.
	Sha1Hashes pulumi.StringArrayOutput `pulumi:"sha1Hashes"`
	// The SHA256 certificate hashes for the AndroidApp.
	Sha256Hashes pulumi.StringArrayOutput            `pulumi:"sha256Hashes"`
	Timeouts     FirebaseAndroidAppTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewFirebaseAndroidApp registers a new resource with the given unique name, arguments, and options.
func NewFirebaseAndroidApp(ctx *pulumi.Context,
	name string, args *FirebaseAndroidAppArgs, opts ...pulumi.ResourceOption) (*FirebaseAndroidApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.PackageName == nil {
		return nil, errors.New("invalid value for required argument 'PackageName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseAndroidApp
	err = ctx.RegisterPackageResource("google-beta:index/firebaseAndroidApp:FirebaseAndroidApp", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseAndroidApp gets an existing FirebaseAndroidApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseAndroidApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseAndroidAppState, opts ...pulumi.ResourceOption) (*FirebaseAndroidApp, error) {
	var resource FirebaseAndroidApp
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/firebaseAndroidApp:FirebaseAndroidApp", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseAndroidApp resources.
type firebaseAndroidAppState struct {
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp. If
	// apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp. This
	// auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	ApiKeyId *string `pulumi:"apiKeyId"`
	// The globally unique, Firebase-assigned identifier of the AndroidApp. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId          *string `pulumi:"appId"`
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The user-assigned display name of the AndroidApp.
	DisplayName *string `pulumi:"displayName"`
	// This checksum is computed by the server based on the value of other fields, and it may be sent with update requests to
	// ensure the client has an up-to-date value before proceeding.
	Etag                 *string `pulumi:"etag"`
	FirebaseAndroidAppId *string `pulumi:"firebaseAndroidAppId"`
	// The fully qualified resource name of the AndroidApp, for example: projects/projectId/androidApps/appId
	Name *string `pulumi:"name"`
	// The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName *string `pulumi:"packageName"`
	Project     *string `pulumi:"project"`
	// The SHA1 certificate hashes for the AndroidApp.
	Sha1Hashes []string `pulumi:"sha1Hashes"`
	// The SHA256 certificate hashes for the AndroidApp.
	Sha256Hashes []string                    `pulumi:"sha256Hashes"`
	Timeouts     *FirebaseAndroidAppTimeouts `pulumi:"timeouts"`
}

type FirebaseAndroidAppState struct {
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp. If
	// apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp. This
	// auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	ApiKeyId pulumi.StringPtrInput
	// The globally unique, Firebase-assigned identifier of the AndroidApp. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId          pulumi.StringPtrInput
	DeletionPolicy pulumi.StringPtrInput
	// The user-assigned display name of the AndroidApp.
	DisplayName pulumi.StringPtrInput
	// This checksum is computed by the server based on the value of other fields, and it may be sent with update requests to
	// ensure the client has an up-to-date value before proceeding.
	Etag                 pulumi.StringPtrInput
	FirebaseAndroidAppId pulumi.StringPtrInput
	// The fully qualified resource name of the AndroidApp, for example: projects/projectId/androidApps/appId
	Name pulumi.StringPtrInput
	// The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName pulumi.StringPtrInput
	Project     pulumi.StringPtrInput
	// The SHA1 certificate hashes for the AndroidApp.
	Sha1Hashes pulumi.StringArrayInput
	// The SHA256 certificate hashes for the AndroidApp.
	Sha256Hashes pulumi.StringArrayInput
	Timeouts     FirebaseAndroidAppTimeoutsPtrInput
}

func (FirebaseAndroidAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAndroidAppState)(nil)).Elem()
}

type firebaseAndroidAppArgs struct {
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp. If
	// apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp. This
	// auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	ApiKeyId       *string `pulumi:"apiKeyId"`
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The user-assigned display name of the AndroidApp.
	DisplayName          string  `pulumi:"displayName"`
	FirebaseAndroidAppId *string `pulumi:"firebaseAndroidAppId"`
	// The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName string  `pulumi:"packageName"`
	Project     *string `pulumi:"project"`
	// The SHA1 certificate hashes for the AndroidApp.
	Sha1Hashes []string `pulumi:"sha1Hashes"`
	// The SHA256 certificate hashes for the AndroidApp.
	Sha256Hashes []string                    `pulumi:"sha256Hashes"`
	Timeouts     *FirebaseAndroidAppTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a FirebaseAndroidApp resource.
type FirebaseAndroidAppArgs struct {
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp. If
	// apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp. This
	// auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	ApiKeyId       pulumi.StringPtrInput
	DeletionPolicy pulumi.StringPtrInput
	// The user-assigned display name of the AndroidApp.
	DisplayName          pulumi.StringInput
	FirebaseAndroidAppId pulumi.StringPtrInput
	// The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName pulumi.StringInput
	Project     pulumi.StringPtrInput
	// The SHA1 certificate hashes for the AndroidApp.
	Sha1Hashes pulumi.StringArrayInput
	// The SHA256 certificate hashes for the AndroidApp.
	Sha256Hashes pulumi.StringArrayInput
	Timeouts     FirebaseAndroidAppTimeoutsPtrInput
}

func (FirebaseAndroidAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAndroidAppArgs)(nil)).Elem()
}

type FirebaseAndroidAppInput interface {
	pulumi.Input

	ToFirebaseAndroidAppOutput() FirebaseAndroidAppOutput
	ToFirebaseAndroidAppOutputWithContext(ctx context.Context) FirebaseAndroidAppOutput
}

func (*FirebaseAndroidApp) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAndroidApp)(nil)).Elem()
}

func (i *FirebaseAndroidApp) ToFirebaseAndroidAppOutput() FirebaseAndroidAppOutput {
	return i.ToFirebaseAndroidAppOutputWithContext(context.Background())
}

func (i *FirebaseAndroidApp) ToFirebaseAndroidAppOutputWithContext(ctx context.Context) FirebaseAndroidAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseAndroidAppOutput)
}

type FirebaseAndroidAppOutput struct{ *pulumi.OutputState }

func (FirebaseAndroidAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAndroidApp)(nil)).Elem()
}

func (o FirebaseAndroidAppOutput) ToFirebaseAndroidAppOutput() FirebaseAndroidAppOutput {
	return o
}

func (o FirebaseAndroidAppOutput) ToFirebaseAndroidAppOutputWithContext(ctx context.Context) FirebaseAndroidAppOutput {
	return o
}

// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp. If
// apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp. This
// auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
func (o FirebaseAndroidAppOutput) ApiKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.ApiKeyId }).(pulumi.StringOutput)
}

// The globally unique, Firebase-assigned identifier of the AndroidApp. This identifier should be treated as an opaque
// token, as the data format is not specified.
func (o FirebaseAndroidAppOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o FirebaseAndroidAppOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

// The user-assigned display name of the AndroidApp.
func (o FirebaseAndroidAppOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and it may be sent with update requests to
// ensure the client has an up-to-date value before proceeding.
func (o FirebaseAndroidAppOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FirebaseAndroidAppOutput) FirebaseAndroidAppId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.FirebaseAndroidAppId }).(pulumi.StringOutput)
}

// The fully qualified resource name of the AndroidApp, for example: projects/projectId/androidApps/appId
func (o FirebaseAndroidAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The canonical package name of the Android app as would appear in the Google Play Developer Console.
func (o FirebaseAndroidAppOutput) PackageName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.PackageName }).(pulumi.StringOutput)
}

func (o FirebaseAndroidAppOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The SHA1 certificate hashes for the AndroidApp.
func (o FirebaseAndroidAppOutput) Sha1Hashes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringArrayOutput { return v.Sha1Hashes }).(pulumi.StringArrayOutput)
}

// The SHA256 certificate hashes for the AndroidApp.
func (o FirebaseAndroidAppOutput) Sha256Hashes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) pulumi.StringArrayOutput { return v.Sha256Hashes }).(pulumi.StringArrayOutput)
}

func (o FirebaseAndroidAppOutput) Timeouts() FirebaseAndroidAppTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseAndroidApp) FirebaseAndroidAppTimeoutsPtrOutput { return v.Timeouts }).(FirebaseAndroidAppTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseAndroidAppInput)(nil)).Elem(), &FirebaseAndroidApp{})
	pulumi.RegisterOutputType(FirebaseAndroidAppOutput{})
}
