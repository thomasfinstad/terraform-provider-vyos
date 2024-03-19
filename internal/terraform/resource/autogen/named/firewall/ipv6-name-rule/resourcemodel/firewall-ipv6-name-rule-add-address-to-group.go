// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixNameRuleAddAddressToGroup{}

// FirewallIPvsixNameRuleAddAddressToGroup describes the resource data model.
type FirewallIPvsixNameRuleAddAddressToGroup struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixNameRuleAddAddressToGroupSourceAddress      *FirewallIPvsixNameRuleAddAddressToGroupSourceAddress      `tfsdk:"source_address" vyos:"source-address,omitempty"`
	NodeFirewallIPvsixNameRuleAddAddressToGroupDestinationAddress *FirewallIPvsixNameRuleAddAddressToGroupDestinationAddress `tfsdk:"destination_address" vyos:"destination-address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleAddAddressToGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"source_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleAddAddressToGroupSourceAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add source ipv6 addresses to dynamic ipv6-address-group

`,
			Description: `Add source ipv6 addresses to dynamic ipv6-address-group

`,
		},

		"destination_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleAddAddressToGroupDestinationAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add destination ipv6 addresses to dynamic ipv6-address-group

`,
			Description: `Add destination ipv6 addresses to dynamic ipv6-address-group

`,
		},
	}
}
