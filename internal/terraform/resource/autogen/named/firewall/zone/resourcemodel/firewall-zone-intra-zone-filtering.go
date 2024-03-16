// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallZoneIntraZoneFiltering{}

// FirewallZoneIntraZoneFiltering describes the resource data model.
type FirewallZoneIntraZoneFiltering struct {
	// LeafNodes
	LeafFirewallZoneIntraZoneFilteringAction types.String `tfsdk:"action" vyos:"action,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallZoneIntraZoneFilteringFirewall *FirewallZoneIntraZoneFilteringFirewall `tfsdk:"firewall" vyos:"firewall,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneIntraZoneFiltering) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action for intra-zone traffic

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  accept  &emsp; |  Accept traffic  |
    |  drop  &emsp; |  Drop silently  |

`,
		},

		// Nodes

		"firewall": schema.SingleNestedAttribute{
			Attributes: FirewallZoneIntraZoneFilteringFirewall{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use the specified firewall chain

`,
		},
	}
}
