// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteMapRuleSetEvpn describes the resource data model.
type PolicyRouteMapRuleSetEvpn struct {
	// LeafNodes

	// TagNodes

	// Nodes
	PolicyRouteMapRuleSetEvpnGateway types.Object `tfsdk:"gateway" json:"gateway,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteMapRuleSetEvpn) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"gateway": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetEvpnGateway{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Set gateway IP for prefix advertisement route

`,
		},
	}
}
