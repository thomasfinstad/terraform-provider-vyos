// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// HighAvailabilityVrrpGroupAuthentication describes the resource data model.
type HighAvailabilityVrrpGroupAuthentication struct {
	// LeafNodes
	HighAvailabilityVrrpGroupAuthenticationPassword customtypes.CustomStringValue `tfsdk:"password" json:"password,omitempty"`
	HighAvailabilityVrrpGroupAuthenticationType     customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o HighAvailabilityVrrpGroupAuthentication) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VRRP password

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Password string (up to 8 characters)  |
`,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Authentication type

|  Format  |  Description  |
|----------|---------------|
|  plaintext-password  |  Simple password string  |
|  ah  |  AH - IPSEC (not recommended)  |
`,
		},

		// TagNodes

		// Nodes

	}
}
