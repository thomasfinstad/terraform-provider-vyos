// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourInputFilterRuleDestinationGroup{}

// FirewallIPvfourInputFilterRuleDestinationGroup describes the resource data model.
type FirewallIPvfourInputFilterRuleDestinationGroup struct {
	// LeafNodes
	LeafFirewallIPvfourInputFilterRuleDestinationGroupAddressGroup        types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafFirewallIPvfourInputFilterRuleDestinationGroupDomainGroup         types.String `tfsdk:"domain_group" vyos:"domain-group,omitempty"`
	LeafFirewallIPvfourInputFilterRuleDestinationGroupMacGroup            types.String `tfsdk:"mac_group" vyos:"mac-group,omitempty"`
	LeafFirewallIPvfourInputFilterRuleDestinationGroupNetworkGroup        types.String `tfsdk:"network_group" vyos:"network-group,omitempty"`
	LeafFirewallIPvfourInputFilterRuleDestinationGroupPortGroup           types.String `tfsdk:"port_group" vyos:"port-group,omitempty"`
	LeafFirewallIPvfourInputFilterRuleDestinationGroupDynamicAddressGroup types.String `tfsdk:"dynamic_address_group" vyos:"dynamic-address-group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourInputFilterRuleDestinationGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of addresses

`,
			Description: `Group of addresses

`,
		},

		"domain_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of domains

`,
			Description: `Group of domains

`,
		},

		"mac_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of MAC addresses

`,
			Description: `Group of MAC addresses

`,
		},

		"network_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of networks

`,
			Description: `Group of networks

`,
		},

		"port_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of ports

`,
			Description: `Group of ports

`,
		},

		"dynamic_address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of dynamic addresses

`,
			Description: `Group of dynamic addresses

`,
		},

		// Nodes

	}
}
