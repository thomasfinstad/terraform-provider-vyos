// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddress{}

// FirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddress describes the resource data model.
type FirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddress struct {
	// LeafNodes
	LeafFirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddressAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddressTimeout      types.String `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourForwardFilterRuleAddAddressToGroupSourceAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Dynamic address-group

`,
			Description: `Dynamic address-group

`,
		},

		"timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set timeout

    |  Format           &emsp;|  Description               |
    |-------------------------|----------------------------|
    |  &lt;number&gt;s  &emsp;|  Timeout value in seconds  |
    |  &lt;number&gt;m  &emsp;|  Timeout value in minutes  |
    |  &lt;number&gt;h  &emsp;|  Timeout value in hours    |
    |  &lt;number&gt;d  &emsp;|  Timeout value in days     |
`,
			Description: `Set timeout

    |  Format           |  Description               |
    |-------------------------|----------------------------|
    |  <number>s  |  Timeout value in seconds  |
    |  <number>m  |  Timeout value in minutes  |
    |  <number>h  |  Timeout value in hours    |
    |  <number>d  |  Timeout value in days     |
`,
		},

		// Nodes

	}
}
