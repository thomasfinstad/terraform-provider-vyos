// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetAggregator{}

// PolicyRouteMapRuleSetAggregator describes the resource data model.
type PolicyRouteMapRuleSetAggregator struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetAggregatorAs types.Number `tfsdk:"as" vyos:"as,omitempty"`
	LeafPolicyRouteMapRuleSetAggregatorIP types.String `tfsdk:"ip" vyos:"ip,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetAggregator) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `AS number of an aggregation

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-4294967295  |  Rule number  |
`,
			Description: `AS number of an aggregation

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-4294967295  |  Rule number  |
`,
		},

		"ip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of an aggregation

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4    |  IP address   |
`,
			Description: `IP address of an aggregation

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4    |  IP address   |
`,
		},

		// Nodes

	}
}
