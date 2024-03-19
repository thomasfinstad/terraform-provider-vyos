// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixInputFilterRuleSynproxy{}

// FirewallIPvsixInputFilterRuleSynproxy describes the resource data model.
type FirewallIPvsixInputFilterRuleSynproxy struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixInputFilterRuleSynproxyTCP *FirewallIPvsixInputFilterRuleSynproxyTCP `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixInputFilterRuleSynproxy) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleSynproxyTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP synproxy options

`,
			Description: `TCP synproxy options

`,
		},
	}
}
