// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIPRouteSource{}

// PolicyRouteMapRuleMatchIPRouteSource describes the resource data model.
type PolicyRouteMapRuleMatchIPRouteSource struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPRouteSourceAccessList types.Number `tfsdk:"access_list" vyos:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPRouteSourcePrefixList types.String `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPRouteSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"access_list": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IP access-list to match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-99  &emsp; |  IP standard access list  |
    |  number: 100-199  &emsp; |  IP extended access list  |
    |  number: 1300-1999  &emsp; |  IP standard access list (expanded range)  |
    |  number: 2000-2699  &emsp; |  IP extended access list (expanded range)  |

`,
		},

		"prefix_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-list to match

`,
		},

		// Nodes

	}
}
