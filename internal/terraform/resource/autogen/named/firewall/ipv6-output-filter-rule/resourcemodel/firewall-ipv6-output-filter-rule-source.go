// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixOutputFilterRuleSource{}

// FirewallIPvsixOutputFilterRuleSource describes the resource data model.
type FirewallIPvsixOutputFilterRuleSource struct {
	// LeafNodes
	LeafFirewallIPvsixOutputFilterRuleSourceAddress     types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleSourceAddressMask types.String `tfsdk:"address_mask" vyos:"address-mask,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleSourceFqdn        types.String `tfsdk:"fqdn" vyos:"fqdn,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleSourceMacAddress  types.String `tfsdk:"mac_address" vyos:"mac-address,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleSourcePort        types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixOutputFilterRuleSourceGeoIP *FirewallIPvsixOutputFilterRuleSourceGeoIP `tfsdk:"geoip" vyos:"geoip,omitempty"`
	NodeFirewallIPvsixOutputFilterRuleSourceGroup *FirewallIPvsixOutputFilterRuleSourceGroup `tfsdk:"group" vyos:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleSource) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv6        |  IP address to match                            |
    |  ipv6net     |  Subnet to match                                |
    |  ipv6range   |  IP range to match                              |
    |  !ipv6       |  Match everything except the specified address  |
    |  !ipv6net    |  Match everything except the specified prefix   |
    |  !ipv6range  |  Match everything except the specified range    |
`,
			Description: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv6        |  IP address to match                            |
    |  ipv6net     |  Subnet to match                                |
    |  ipv6range   |  IP range to match                              |
    |  !ipv6       |  Match everything except the specified address  |
    |  !ipv6net    |  Match everything except the specified prefix   |
    |  !ipv6range  |  Match everything except the specified range    |
`,
		},

		"address_mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP mask

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv6    |  IP mask to apply  |
`,
			Description: `IP mask

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv6    |  IP mask to apply  |
`,
		},

		"fqdn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Fully qualified domain name

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  <fqdn>  |  Fully qualified domain name  |
`,
			Description: `Fully qualified domain name

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  <fqdn>  |  Fully qualified domain name  |
`,
		},

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MAC address

    |  Format    |  Description                                        |
    |------------|-----------------------------------------------------|
    |  macaddr   |  MAC address to match                               |
    |  !macaddr  |  Match everything except the specified MAC address  |
`,
			Description: `MAC address

    |  Format    |  Description                                        |
    |------------|-----------------------------------------------------|
    |  macaddr   |  MAC address to match                               |
    |  !macaddr  |  Match everything except the specified MAC address  |
`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port

    |  Format       |  Description                                                                                                               |
    |---------------|----------------------------------------------------------------------------------------------------------------------------|
    |  txt          |  Named port (any name in /etc/services, e.g., http)                                                                        |
    |  1-65535      |  Numbered port                                                                                                             |
    |  <start-end>  |  Numbered port range (e.g. 1001-1005)                                                                                      |
    |               |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |
`,
			Description: `Port

    |  Format       |  Description                                                                                                               |
    |---------------|----------------------------------------------------------------------------------------------------------------------------|
    |  txt          |  Named port (any name in /etc/services, e.g., http)                                                                        |
    |  1-65535      |  Numbered port                                                                                                             |
    |  <start-end>  |  Numbered port range (e.g. 1001-1005)                                                                                      |
    |               |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |
`,
		},

		// Nodes

		"geoip": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixOutputFilterRuleSourceGeoIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `GeoIP options - Data provided by DB-IP.com

`,
			Description: `GeoIP options - Data provided by DB-IP.com

`,
		},

		"group": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixOutputFilterRuleSourceGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Group

`,
			Description: `Group

`,
		},
	}
}
