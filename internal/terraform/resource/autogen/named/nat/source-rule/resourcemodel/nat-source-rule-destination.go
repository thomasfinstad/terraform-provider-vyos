// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatSourceRuleDestination{}

// NatSourceRuleDestination describes the resource data model.
type NatSourceRuleDestination struct {
	// LeafNodes
	LeafNatSourceRuleDestinationAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafNatSourceRuleDestinationPort    types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeNatSourceRuleDestinationGroup *NatSourceRuleDestinationGroup `tfsdk:"group" vyos:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatSourceRuleDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv4        |  IPv4 address to match                          |
    |  ipv4net     |  IPv4 prefix to match                           |
    |  ipv4range   |  IPv4 address range to match                    |
    |  !ipv4       |  Match everything except the specified address  |
    |  !ipv4net    |  Match everything except the specified prefix   |
    |  !ipv4range  |  Match everything except the specified range    |
`,
			Description: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv4        |  IPv4 address to match                          |
    |  ipv4net     |  IPv4 prefix to match                           |
    |  ipv4range   |  IPv4 address range to match                    |
    |  !ipv4       |  Match everything except the specified address  |
    |  !ipv4net    |  Match everything except the specified prefix   |
    |  !ipv4range  |  Match everything except the specified range    |
`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

    |  Format     |  Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |  txt        |  Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    |  1-65535    |  Numeric IP port                                                                                                                                                          |
    |  start-end  |  Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    |             |  \n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using '!'.\nFor example: '!22,telnet,http,123,1001-1005'  |
`,
			Description: `Port number

    |  Format     |  Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |  txt        |  Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    |  1-65535    |  Numeric IP port                                                                                                                                                          |
    |  start-end  |  Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    |             |  \n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using '!'.\nFor example: '!22,telnet,http,123,1001-1005'  |
`,
		},

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: NatSourceRuleDestinationGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Group

`,
			Description: `Group

`,
		},
	}
}
