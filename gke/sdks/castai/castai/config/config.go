// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The token used to connect to CAST AI API.
func GetApiToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "castai:apiToken")
}

// CAST.AI API url.
func GetApiUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "castai:apiUrl")
}
