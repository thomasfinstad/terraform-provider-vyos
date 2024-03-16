// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourNameRuleOutboundInterface{}

// FirewallIPvfourNameRuleOutboundInterface describes the resource data model.
type FirewallIPvfourNameRuleOutboundInterface struct {
	// LeafNodes
	LeafFirewallIPvfourNameRuleOutboundInterfaceName  types.String `tfsdk:"name" vyos:"name,omitempty"`
	LeafFirewallIPvfourNameRuleOutboundInterfaceGroup types.String `tfsdk:"group" vyos:"group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRuleOutboundInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |
    |  txt&  &emsp; |  Interface name with wildcard  |
    |  !txt  &emsp; |  Inverted interface name to match  |

`,
		},

		"group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface-group

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface-group name to match  |
    |  !txt  &emsp; |  Inverted interface-group name to match  |

`,
		},

		// Nodes

	}
}
