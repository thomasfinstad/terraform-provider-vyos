// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyAccessListsixRuleSource describes the resource data model.
type PolicyAccessListsixRuleSource struct {
	// LeafNodes
	PolicyAccessListsixRuleSourceAny        customtypes.CustomStringValue `tfsdk:"any" json:"any,omitempty"`
	PolicyAccessListsixRuleSourceExactMatch customtypes.CustomStringValue `tfsdk:"exact_match" json:"exact-match,omitempty"`
	PolicyAccessListsixRuleSourceNetwork    customtypes.CustomStringValue `tfsdk:"network" json:"network,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyAccessListsixRuleSource) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"any": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Any IP address to match

`,
		},

		"exact_match": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Exact match of the network prefixes

`,
		},

		"network": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Network/netmask to match

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		// TagNodes

		// Nodes

	}
}
