// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatSourceRuleLoadBalance{}

// NatSourceRuleLoadBalance describes the resource data model.
type NatSourceRuleLoadBalance struct {
	// LeafNodes
	LeafNatSourceRuleLoadBalanceHash types.List `tfsdk:"hash" vyos:"hash,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagNatSourceRuleLoadBalanceBackend bool `tfsdk:"backend" vyos:"backend,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatSourceRuleLoadBalance) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hash": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Define the parameters of the packet header to apply the hashing

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  source-address  &emsp; |  Use source IP address for hashing  |
    |  destination-address  &emsp; |  Use destination IP address for hashing  |
    |  source-port  &emsp; |  Use source port for hashing  |
    |  destination-port  &emsp; |  Use destination port for hashing  |
    |  random  &emsp; |  Do not use information from ip header. Use random value.  |

`,

			// Default:          stringdefault.StaticString(`random`),
			Computed: true,
		},

		// Nodes

	}
}
