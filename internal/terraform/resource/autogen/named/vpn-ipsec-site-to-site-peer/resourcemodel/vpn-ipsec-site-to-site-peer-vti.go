// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VpnIPsecSiteToSitePeerVti describes the resource data model.
type VpnIPsecSiteToSitePeerVti struct {
	// LeafNodes
	VpnIPsecSiteToSitePeerVtiBind     customtypes.CustomStringValue `tfsdk:"bind" json:"bind,omitempty"`
	VpnIPsecSiteToSitePeerVtiEspGroup customtypes.CustomStringValue `tfsdk:"esp_group" json:"esp-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerVti) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bind": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VTI tunnel interface associated with this configuration

`,
		},

		"esp_group": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
		},

		// TagNodes

		// Nodes

	}
}
