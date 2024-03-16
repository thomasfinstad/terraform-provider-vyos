// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeNameRuleSource{}

// FirewallBrIDgeNameRuleSource describes the resource data model.
type FirewallBrIDgeNameRuleSource struct {
	// LeafNodes
	LeafFirewallBrIDgeNameRuleSourceMacAddress types.String `tfsdk:"mac_address" vyos:"mac-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeNameRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MAC address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  MAC address to match  |
    |  !macaddr  &emsp; |  Match everything except the specified MAC address  |

`,
		},

		// Nodes

	}
}
