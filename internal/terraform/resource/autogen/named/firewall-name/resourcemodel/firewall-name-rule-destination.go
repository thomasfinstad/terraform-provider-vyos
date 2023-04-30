// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// FirewallNameRuleDestination describes the resource data model.
type FirewallNameRuleDestination struct {
	// LeafNodes
	FirewallNameRuleDestinationAddress     customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	FirewallNameRuleDestinationFqdn        customtypes.CustomStringValue `tfsdk:"fqdn" json:"fqdn,omitempty"`
	FirewallNameRuleDestinationPort        customtypes.CustomStringValue `tfsdk:"port" json:"port,omitempty"`
	FirewallNameRuleDestinationAddressMask customtypes.CustomStringValue `tfsdk:"address_mask" json:"address-mask,omitempty"`
	FirewallNameRuleDestinationMacAddress  customtypes.CustomStringValue `tfsdk:"mac_address" json:"mac-address,omitempty"`

	// TagNodes

	// Nodes
	FirewallNameRuleDestinationGeoIP types.Object `tfsdk:"geoip" json:"geoip,omitempty"`
	FirewallNameRuleDestinationGroup types.Object `tfsdk:"group" json:"group,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o FirewallNameRuleDestination) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
|  ipv4range  |  IPv4 address range to match  |
|  !ipv4  |  Match everything except the specified address  |
|  !ipv4net  |  Match everything except the specified prefix  |
|  !ipv4range  |  Match everything except the specified range  |
`,
		},

		"fqdn": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Fully qualified domain name

|  Format  |  Description  |
|----------|---------------|
|  <fqdn>  |  Fully qualified domain name  |
`,
		},

		"port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Port

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numbered port  |
|  <start-end>  |  Numbered port range (e.g. 1001-1005)  |
|     |  \n\n Multiple destination ports can be specified as a
                          comma-separated list.\n For example: 'telnet,http,123,1001-1005'  |
`,
		},

		"address_mask": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP mask

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 mask to apply  |
`,
		},

		"mac_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `MAC address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |
|  !macaddr  |  Match everything except the specified MAC address  |
`,
		},

		// TagNodes

		// Nodes

		"geoip": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleDestinationGeoIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `GeoIP options - Data provided by DB-IP.com

`,
		},

		"group": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleDestinationGroup{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}
