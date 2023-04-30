// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget struct {
	// LeafNodes

	// TagNodes

	// Nodes
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn types.Object `tfsdk:"vpn" json:"vpn,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Between current address-family and VPN

`,
		},
	}
}
