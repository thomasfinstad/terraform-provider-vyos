// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourNameRuleTTL{}

// FirewallIPvfourNameRuleTTL describes the resource data model.
type FirewallIPvfourNameRuleTTL struct {
	// LeafNodes
	LeafFirewallIPvfourNameRuleTTLEq types.Number `tfsdk:"eq" vyos:"eq,omitempty"`
	LeafFirewallIPvfourNameRuleTTLGt types.Number `tfsdk:"gt" vyos:"gt,omitempty"`
	LeafFirewallIPvfourNameRuleTTLLt types.Number `tfsdk:"lt" vyos:"lt,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRuleTTL) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"eq": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on equal value

    |  Format  &emsp;|  Description     |
    |----------------|------------------|
    |  0-255   &emsp;|  Equal to value  |
`,
			Description: `Match on equal value

    |  Format  |  Description     |
    |----------------|------------------|
    |  0-255   |  Equal to value  |
`,
		},

		"gt": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on greater then value

    |  Format  &emsp;|  Description         |
    |----------------|----------------------|
    |  0-255   &emsp;|  Greater then value  |
`,
			Description: `Match on greater then value

    |  Format  |  Description         |
    |----------------|----------------------|
    |  0-255   |  Greater then value  |
`,
		},

		"lt": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on less then value

    |  Format  &emsp;|  Description      |
    |----------------|-------------------|
    |  0-255   &emsp;|  Less then value  |
`,
			Description: `Match on less then value

    |  Format  |  Description      |
    |----------------|-------------------|
    |  0-255   |  Less then value  |
`,
		},

		// Nodes

	}
}
