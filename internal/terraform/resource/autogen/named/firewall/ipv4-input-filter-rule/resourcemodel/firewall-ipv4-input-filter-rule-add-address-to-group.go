// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourInputFilterRuleAddAddressToGroup{}

// FirewallIPvfourInputFilterRuleAddAddressToGroup describes the resource data model.
type FirewallIPvfourInputFilterRuleAddAddressToGroup struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourInputFilterRuleAddAddressToGroupSourceAddress      *FirewallIPvfourInputFilterRuleAddAddressToGroupSourceAddress      `tfsdk:"source_address" vyos:"source-address,omitempty"`
	NodeFirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddress *FirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddress `tfsdk:"destination_address" vyos:"destination-address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourInputFilterRuleAddAddressToGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"source_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourInputFilterRuleAddAddressToGroupSourceAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add source ip addresses to dynamic address-group

`,
			Description: `Add source ip addresses to dynamic address-group

`,
		},

		"destination_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add destination ip addresses to dynamic address-group

`,
			Description: `Add destination ip addresses to dynamic address-group

`,
		},
	}
}
