// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourNameRuleRecent{}

// FirewallIPvfourNameRuleRecent describes the resource data model.
type FirewallIPvfourNameRuleRecent struct {
	// LeafNodes
	LeafFirewallIPvfourNameRuleRecentCount types.Number `tfsdk:"count" vyos:"count,omitempty"`
	LeafFirewallIPvfourNameRuleRecentTime  types.String `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRuleRecent) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"count": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen more than N times

    |  Format  &emsp;|  Description                              |
    |----------------|-------------------------------------------|
    |  1-255   &emsp;|  Source addresses seen more than N times  |
`,
			Description: `Source addresses seen more than N times

    |  Format  |  Description                              |
    |----------------|-------------------------------------------|
    |  1-255   |  Source addresses seen more than N times  |
`,
		},

		"time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen in the last second/minute/hour

    |  Format  &emsp;|  Description                                           |
    |----------------|--------------------------------------------------------|
    |  second  &emsp;|  Source addresses seen COUNT times in the last second  |
    |  minute  &emsp;|  Source addresses seen COUNT times in the last minute  |
    |  hour    &emsp;|  Source addresses seen COUNT times in the last hour    |
`,
			Description: `Source addresses seen in the last second/minute/hour

    |  Format  |  Description                                           |
    |----------------|--------------------------------------------------------|
    |  second  |  Source addresses seen COUNT times in the last second  |
    |  minute  |  Source addresses seen COUNT times in the last minute  |
    |  hour    |  Source addresses seen COUNT times in the last hour    |
`,
		},

		// Nodes

	}
}
