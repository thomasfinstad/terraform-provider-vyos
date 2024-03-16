// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourInputFilterRuleSynproxy{}

// FirewallIPvfourInputFilterRuleSynproxy describes the resource data model.
type FirewallIPvfourInputFilterRuleSynproxy struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourInputFilterRuleSynproxyTCP *FirewallIPvfourInputFilterRuleSynproxyTCP `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourInputFilterRuleSynproxy) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourInputFilterRuleSynproxyTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP synproxy options

`,
		},
	}
}
