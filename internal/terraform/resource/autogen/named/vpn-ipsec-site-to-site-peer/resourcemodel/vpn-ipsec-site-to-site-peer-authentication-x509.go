// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VpnIPsecSiteToSitePeerAuthenticationXfivezeronine describes the resource data model.
type VpnIPsecSiteToSitePeerAuthenticationXfivezeronine struct {
	// LeafNodes
	VpnIPsecSiteToSitePeerAuthenticationXfivezeronineCertificate   customtypes.CustomStringValue `tfsdk:"certificate" json:"certificate,omitempty"`
	VpnIPsecSiteToSitePeerAuthenticationXfivezeroninePassphrase    customtypes.CustomStringValue `tfsdk:"passphrase" json:"passphrase,omitempty"`
	VpnIPsecSiteToSitePeerAuthenticationXfivezeronineCaCertificate customtypes.CustomStringValue `tfsdk:"ca_certificate" json:"ca-certificate,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerAuthenticationXfivezeronine) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"certificate": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Certificate in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of certificate in PKI configuration  |
`,
		},

		"passphrase": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Private key passphrase

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Passphrase to decrypt the private key  |
`,
		},

		"ca_certificate": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Certificate Authority in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of CA in PKI configuration  |
`,
		},

		// TagNodes

		// Nodes

	}
}
