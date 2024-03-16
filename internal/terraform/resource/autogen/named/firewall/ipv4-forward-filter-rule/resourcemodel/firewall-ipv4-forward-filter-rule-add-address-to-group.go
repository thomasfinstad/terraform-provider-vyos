// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourForwardFilterRuleAddAddressToGroup{}

// FirewallIPvfourForwardFilterRuleAddAddressToGroup describes the resource data model.
type FirewallIPvfourForwardFilterRuleAddAddressToGroup struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddress      *FirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddress      `tfsdk:"source_address" vyos:"source-address,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleAddAddressToGroupDestinationAddress *FirewallIPvfourForwardFilterRuleAddAddressToGroupDestinationAddress `tfsdk:"destination_address" vyos:"destination-address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourForwardFilterRuleAddAddressToGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"source_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add source ip addresses to dynamic address-group

`,
		},

		"destination_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleAddAddressToGroupDestinationAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add destination ip addresses to dynamic address-group

`,
		},
	}
}
