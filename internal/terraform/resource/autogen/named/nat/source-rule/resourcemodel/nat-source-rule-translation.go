// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatSourceRuleTranSLAtion{}

// NatSourceRuleTranSLAtion describes the resource data model.
type NatSourceRuleTranSLAtion struct {
	// LeafNodes
	LeafNatSourceRuleTranSLAtionAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafNatSourceRuleTranSLAtionPort    types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeNatSourceRuleTranSLAtionOptions *NatSourceRuleTranSLAtionOptions `tfsdk:"options" vyos:"options,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatSourceRuleTranSLAtion) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format      |  Description                                       |
    |--------------|----------------------------------------------------|
    |  ipv4        |  IPv4 address to match                             |
    |  ipv4net     |  IPv4 prefix to match                              |
    |  ipv4range   |  IPv4 address range to match                       |
    |  masquerade  |  NAT to the primary address of outbound-interface  |
`,
			Description: `IP address, subnet, or range

    |  Format      |  Description                                       |
    |--------------|----------------------------------------------------|
    |  ipv4        |  IPv4 address to match                             |
    |  ipv4net     |  IPv4 prefix to match                              |
    |  ipv4range   |  IPv4 address range to match                       |
    |  masquerade  |  NAT to the primary address of outbound-interface  |
`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

    |  Format   |  Description                            |
    |-----------|-----------------------------------------|
    |  1-65535  |  Numeric IP port                        |
    |  range    |  Numbered port range (e.g., 1001-1005)  |
`,
			Description: `Port number

    |  Format   |  Description                            |
    |-----------|-----------------------------------------|
    |  1-65535  |  Numeric IP port                        |
    |  range    |  Numbered port range (e.g., 1001-1005)  |
`,
		},

		// Nodes

		"options": schema.SingleNestedAttribute{
			Attributes: NatSourceRuleTranSLAtionOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Translation options

`,
			Description: `Translation options

`,
		},
	}
}
