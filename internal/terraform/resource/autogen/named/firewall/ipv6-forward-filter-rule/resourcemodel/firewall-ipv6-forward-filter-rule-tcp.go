// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleTCP{}

// FirewallIPvsixForwardFilterRuleTCP describes the resource data model.
type FirewallIPvsixForwardFilterRuleTCP struct {
	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleTCPMss types.String `tfsdk:"mss" vyos:"mss,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixForwardFilterRuleTCPFlags *FirewallIPvsixForwardFilterRuleTCPFlags `tfsdk:"flags" vyos:"flags,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum segment size (MSS)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-16384  &emsp; |  Maximum segment size  |
    |  <min>-<max>  &emsp; |  TCP MSS range (use '-' as delimiter)  |

`,
		},

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleTCPFlags{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},
	}
}
