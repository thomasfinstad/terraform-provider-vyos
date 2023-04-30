// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix struct {
	// LeafNodes

	// TagNodes

	// Nodes
	VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsixUnicast types.Object `tfsdk:"unicast" json:"unicast,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsixUnicast{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 address family

`,
		},
	}
}
