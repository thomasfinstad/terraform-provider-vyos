// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddress{}

// FirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddress describes the resource data model.
type FirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddress struct {
	// LeafNodes
	LeafFirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddressAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafFirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddressTimeout      types.String `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

    |  Format     |  Description               |
    |-------------|----------------------------|
    |  <number>s  |  Timeout value in seconds  |
    |  <number>m  |  Timeout value in minutes  |
    |  <number>h  |  Timeout value in hours    |
    |  <number>d  |  Timeout value in days     |
`,
			Description: `Set timeout

    |  Format     |  Description               |
    |-------------|----------------------------|
    |  <number>s  |  Timeout value in seconds  |
    |  <number>m  |  Timeout value in minutes  |
    |  <number>h  |  Timeout value in hours    |
    |  <number>d  |  Timeout value in days     |
`,
		},

		// Nodes

	}
}
