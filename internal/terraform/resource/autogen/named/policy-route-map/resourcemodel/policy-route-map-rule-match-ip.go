// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteMapRuleMatchIP describes the resource data model.
type PolicyRouteMapRuleMatchIP struct {
	// LeafNodes

	// TagNodes

	// Nodes
	PolicyRouteMapRuleMatchIPAddress     types.Object `tfsdk:"address" json:"address,omitempty"`
	PolicyRouteMapRuleMatchIPNexthop     types.Object `tfsdk:"nexthop" json:"nexthop,omitempty"`
	PolicyRouteMapRuleMatchIPRouteSource types.Object `tfsdk:"route_source" json:"route-source,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIP) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPAddress{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IP address of route to match

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPNexthop{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IP next-hop of route to match

`,
		},

		"route_source": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPRouteSource{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Match advertising source address of route

`,
		},
	}
}
