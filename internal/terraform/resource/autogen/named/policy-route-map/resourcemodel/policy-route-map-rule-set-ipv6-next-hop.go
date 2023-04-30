// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRouteMapRuleSetIPvsixNextHop describes the resource data model.
type PolicyRouteMapRuleSetIPvsixNextHop struct {
	// LeafNodes
	PolicyRouteMapRuleSetIPvsixNextHopGlobal       customtypes.CustomStringValue `tfsdk:"global" json:"global,omitempty"`
	PolicyRouteMapRuleSetIPvsixNextHopLocal        customtypes.CustomStringValue `tfsdk:"local" json:"local,omitempty"`
	PolicyRouteMapRuleSetIPvsixNextHopPeerAddress  customtypes.CustomStringValue `tfsdk:"peer_address" json:"peer-address,omitempty"`
	PolicyRouteMapRuleSetIPvsixNextHopPreferGlobal customtypes.CustomStringValue `tfsdk:"prefer_global" json:"prefer-global,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteMapRuleSetIPvsixNextHop) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"global": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Nexthop IPv6 global address

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IPv6 address and prefix length  |
`,
		},

		"local": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Nexthop IPv6 local address

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IPv6 address and prefix length  |
`,
		},

		"peer_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use peer address (for BGP only)

`,
		},

		"prefer_global": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Prefer global address as the nexthop

`,
		},

		// TagNodes

		// Nodes

	}
}
