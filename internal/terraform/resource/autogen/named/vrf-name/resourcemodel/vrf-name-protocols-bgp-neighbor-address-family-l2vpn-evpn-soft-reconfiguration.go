// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration struct {
	// LeafNodes
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfigurationInbound customtypes.CustomStringValue `tfsdk:"inbound" json:"inbound,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
		},

		// TagNodes

		// Nodes

	}
}
