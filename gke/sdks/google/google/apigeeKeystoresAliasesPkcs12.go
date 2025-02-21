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

type ApigeeKeystoresAliasesPkcs12 struct {
	pulumi.CustomResourceState

	// Alias Name
	Alias                          pulumi.StringOutput `pulumi:"alias"`
	ApigeeKeystoresAliasesPkcs12Id pulumi.StringOutput `pulumi:"apigeeKeystoresAliasesPkcs12Id"`
	// Chain of certificates under this alias.
	CertsInfos ApigeeKeystoresAliasesPkcs12CertsInfoArrayOutput `pulumi:"certsInfos"`
	// Environment associated with the alias
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Cert content
	File pulumi.StringOutput `pulumi:"file"`
	// Hash of the pkcs file
	Filehash pulumi.StringOutput `pulumi:"filehash"`
	// Keystore Name
	Keystore pulumi.StringOutput `pulumi:"keystore"`
	// Organization ID associated with the alias
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Password for the Private Key if it's encrypted
	Password pulumi.StringOutput                           `pulumi:"password"`
	Timeouts ApigeeKeystoresAliasesPkcs12TimeoutsPtrOutput `pulumi:"timeouts"`
	// Optional.Type of Alias
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApigeeKeystoresAliasesPkcs12 registers a new resource with the given unique name, arguments, and options.
func NewApigeeKeystoresAliasesPkcs12(ctx *pulumi.Context,
	name string, args *ApigeeKeystoresAliasesPkcs12Args, opts ...pulumi.ResourceOption) (*ApigeeKeystoresAliasesPkcs12, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.File == nil {
		return nil, errors.New("invalid value for required argument 'File'")
	}
	if args.Filehash == nil {
		return nil, errors.New("invalid value for required argument 'Filehash'")
	}
	if args.Keystore == nil {
		return nil, errors.New("invalid value for required argument 'Keystore'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApigeeKeystoresAliasesPkcs12
	err = ctx.RegisterPackageResource("google:index/apigeeKeystoresAliasesPkcs12:ApigeeKeystoresAliasesPkcs12", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeKeystoresAliasesPkcs12 gets an existing ApigeeKeystoresAliasesPkcs12 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeKeystoresAliasesPkcs12(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeKeystoresAliasesPkcs12State, opts ...pulumi.ResourceOption) (*ApigeeKeystoresAliasesPkcs12, error) {
	var resource ApigeeKeystoresAliasesPkcs12
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apigeeKeystoresAliasesPkcs12:ApigeeKeystoresAliasesPkcs12", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeKeystoresAliasesPkcs12 resources.
type apigeeKeystoresAliasesPkcs12State struct {
	// Alias Name
	Alias                          *string `pulumi:"alias"`
	ApigeeKeystoresAliasesPkcs12Id *string `pulumi:"apigeeKeystoresAliasesPkcs12Id"`
	// Chain of certificates under this alias.
	CertsInfos []ApigeeKeystoresAliasesPkcs12CertsInfo `pulumi:"certsInfos"`
	// Environment associated with the alias
	Environment *string `pulumi:"environment"`
	// Cert content
	File *string `pulumi:"file"`
	// Hash of the pkcs file
	Filehash *string `pulumi:"filehash"`
	// Keystore Name
	Keystore *string `pulumi:"keystore"`
	// Organization ID associated with the alias
	OrgId *string `pulumi:"orgId"`
	// Password for the Private Key if it's encrypted
	Password *string                               `pulumi:"password"`
	Timeouts *ApigeeKeystoresAliasesPkcs12Timeouts `pulumi:"timeouts"`
	// Optional.Type of Alias
	Type *string `pulumi:"type"`
}

type ApigeeKeystoresAliasesPkcs12State struct {
	// Alias Name
	Alias                          pulumi.StringPtrInput
	ApigeeKeystoresAliasesPkcs12Id pulumi.StringPtrInput
	// Chain of certificates under this alias.
	CertsInfos ApigeeKeystoresAliasesPkcs12CertsInfoArrayInput
	// Environment associated with the alias
	Environment pulumi.StringPtrInput
	// Cert content
	File pulumi.StringPtrInput
	// Hash of the pkcs file
	Filehash pulumi.StringPtrInput
	// Keystore Name
	Keystore pulumi.StringPtrInput
	// Organization ID associated with the alias
	OrgId pulumi.StringPtrInput
	// Password for the Private Key if it's encrypted
	Password pulumi.StringPtrInput
	Timeouts ApigeeKeystoresAliasesPkcs12TimeoutsPtrInput
	// Optional.Type of Alias
	Type pulumi.StringPtrInput
}

func (ApigeeKeystoresAliasesPkcs12State) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeKeystoresAliasesPkcs12State)(nil)).Elem()
}

type apigeeKeystoresAliasesPkcs12Args struct {
	// Alias Name
	Alias                          string  `pulumi:"alias"`
	ApigeeKeystoresAliasesPkcs12Id *string `pulumi:"apigeeKeystoresAliasesPkcs12Id"`
	// Environment associated with the alias
	Environment string `pulumi:"environment"`
	// Cert content
	File string `pulumi:"file"`
	// Hash of the pkcs file
	Filehash string `pulumi:"filehash"`
	// Keystore Name
	Keystore string `pulumi:"keystore"`
	// Organization ID associated with the alias
	OrgId string `pulumi:"orgId"`
	// Password for the Private Key if it's encrypted
	Password *string                               `pulumi:"password"`
	Timeouts *ApigeeKeystoresAliasesPkcs12Timeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApigeeKeystoresAliasesPkcs12 resource.
type ApigeeKeystoresAliasesPkcs12Args struct {
	// Alias Name
	Alias                          pulumi.StringInput
	ApigeeKeystoresAliasesPkcs12Id pulumi.StringPtrInput
	// Environment associated with the alias
	Environment pulumi.StringInput
	// Cert content
	File pulumi.StringInput
	// Hash of the pkcs file
	Filehash pulumi.StringInput
	// Keystore Name
	Keystore pulumi.StringInput
	// Organization ID associated with the alias
	OrgId pulumi.StringInput
	// Password for the Private Key if it's encrypted
	Password pulumi.StringPtrInput
	Timeouts ApigeeKeystoresAliasesPkcs12TimeoutsPtrInput
}

func (ApigeeKeystoresAliasesPkcs12Args) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeKeystoresAliasesPkcs12Args)(nil)).Elem()
}

type ApigeeKeystoresAliasesPkcs12Input interface {
	pulumi.Input

	ToApigeeKeystoresAliasesPkcs12Output() ApigeeKeystoresAliasesPkcs12Output
	ToApigeeKeystoresAliasesPkcs12OutputWithContext(ctx context.Context) ApigeeKeystoresAliasesPkcs12Output
}

func (*ApigeeKeystoresAliasesPkcs12) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeKeystoresAliasesPkcs12)(nil)).Elem()
}

func (i *ApigeeKeystoresAliasesPkcs12) ToApigeeKeystoresAliasesPkcs12Output() ApigeeKeystoresAliasesPkcs12Output {
	return i.ToApigeeKeystoresAliasesPkcs12OutputWithContext(context.Background())
}

func (i *ApigeeKeystoresAliasesPkcs12) ToApigeeKeystoresAliasesPkcs12OutputWithContext(ctx context.Context) ApigeeKeystoresAliasesPkcs12Output {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeKeystoresAliasesPkcs12Output)
}

type ApigeeKeystoresAliasesPkcs12Output struct{ *pulumi.OutputState }

func (ApigeeKeystoresAliasesPkcs12Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeKeystoresAliasesPkcs12)(nil)).Elem()
}

func (o ApigeeKeystoresAliasesPkcs12Output) ToApigeeKeystoresAliasesPkcs12Output() ApigeeKeystoresAliasesPkcs12Output {
	return o
}

func (o ApigeeKeystoresAliasesPkcs12Output) ToApigeeKeystoresAliasesPkcs12OutputWithContext(ctx context.Context) ApigeeKeystoresAliasesPkcs12Output {
	return o
}

// Alias Name
func (o ApigeeKeystoresAliasesPkcs12Output) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

func (o ApigeeKeystoresAliasesPkcs12Output) ApigeeKeystoresAliasesPkcs12Id() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.ApigeeKeystoresAliasesPkcs12Id }).(pulumi.StringOutput)
}

// Chain of certificates under this alias.
func (o ApigeeKeystoresAliasesPkcs12Output) CertsInfos() ApigeeKeystoresAliasesPkcs12CertsInfoArrayOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) ApigeeKeystoresAliasesPkcs12CertsInfoArrayOutput {
		return v.CertsInfos
	}).(ApigeeKeystoresAliasesPkcs12CertsInfoArrayOutput)
}

// Environment associated with the alias
func (o ApigeeKeystoresAliasesPkcs12Output) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Cert content
func (o ApigeeKeystoresAliasesPkcs12Output) File() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.File }).(pulumi.StringOutput)
}

// Hash of the pkcs file
func (o ApigeeKeystoresAliasesPkcs12Output) Filehash() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.Filehash }).(pulumi.StringOutput)
}

// Keystore Name
func (o ApigeeKeystoresAliasesPkcs12Output) Keystore() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.Keystore }).(pulumi.StringOutput)
}

// Organization ID associated with the alias
func (o ApigeeKeystoresAliasesPkcs12Output) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Password for the Private Key if it's encrypted
func (o ApigeeKeystoresAliasesPkcs12Output) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o ApigeeKeystoresAliasesPkcs12Output) Timeouts() ApigeeKeystoresAliasesPkcs12TimeoutsPtrOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) ApigeeKeystoresAliasesPkcs12TimeoutsPtrOutput { return v.Timeouts }).(ApigeeKeystoresAliasesPkcs12TimeoutsPtrOutput)
}

// Optional.Type of Alias
func (o ApigeeKeystoresAliasesPkcs12Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeKeystoresAliasesPkcs12) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeKeystoresAliasesPkcs12Input)(nil)).Elem(), &ApigeeKeystoresAliasesPkcs12{})
	pulumi.RegisterOutputType(ApigeeKeystoresAliasesPkcs12Output{})
}
