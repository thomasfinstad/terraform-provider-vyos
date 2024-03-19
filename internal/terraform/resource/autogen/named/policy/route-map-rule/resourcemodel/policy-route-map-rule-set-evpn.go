// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetEvpn{}

// PolicyRouteMapRuleSetEvpn describes the resource data model.
type PolicyRouteMapRuleSetEvpn struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRouteMapRuleSetEvpnGateway *PolicyRouteMapRuleSetEvpnGateway `tfsdk:"gateway" vyos:"gateway,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetEvpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"gateway": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetEvpnGateway{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Set gateway IP for prefix advertisement route

`,
			Description: `Set gateway IP for prefix advertisement route

`,
		},
	}
}
