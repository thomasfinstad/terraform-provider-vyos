// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NatDestinationRuleTranSLAtion describes the resource data model.
type NatDestinationRuleTranSLAtion struct {
	// LeafNodes
	LeafNatDestinationRuleTranSLAtionAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafNatDestinationRuleTranSLAtionPort    types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeNatDestinationRuleTranSLAtionOptions  *NatDestinationRuleTranSLAtionOptions  `tfsdk:"options" vyos:"options,omitempty"`
	NodeNatDestinationRuleTranSLAtionRedirect *NatDestinationRuleTranSLAtionRedirect `tfsdk:"redirect" vyos:"redirect,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleTranSLAtion) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to match  |
    |  ipv4net  &emsp; |  IPv4 prefix to match  |
    |  ipv4range  &emsp; |  IPv4 address range to match  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |
    |  range  &emsp; |  Numbered port range (e.g., 1001-1005)  |

`,
		},

		// Nodes

		"options": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleTranSLAtionOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Translation options

`,
		},

		"redirect": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleTranSLAtionRedirect{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redirect to local host

`,
		},
	}
}