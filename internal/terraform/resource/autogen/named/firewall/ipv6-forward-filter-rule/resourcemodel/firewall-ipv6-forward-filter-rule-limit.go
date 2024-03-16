// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleLimit{}

// FirewallIPvsixForwardFilterRuleLimit describes the resource data model.
type FirewallIPvsixForwardFilterRuleLimit struct {
	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleLimitBurst types.Number `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleLimitRate  types.String `tfsdk:"rate" vyos:"rate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleLimit) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"burst": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of packets to allow in excess of rate

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4294967295  &emsp; |  Maximum number of packets to allow in excess of rate  |

`,
		},

		"rate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum average matching rate

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  integer/unit (Example: 5/minute)  |

`,
		},

		// Nodes

	}
}
