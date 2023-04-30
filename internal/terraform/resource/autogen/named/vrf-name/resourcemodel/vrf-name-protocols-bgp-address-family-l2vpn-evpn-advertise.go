// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise struct {
	// LeafNodes

	// TagNodes

	// Nodes
	VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour types.Object `tfsdk:"ipv4" json:"ipv4,omitempty"`
	VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix  types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"ipv4": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 address family

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address family

`,
		},
	}
}
