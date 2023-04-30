// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// FirewallIPvsixNameRuleInboundInterface describes the resource data model.
type FirewallIPvsixNameRuleInboundInterface struct {
	// LeafNodes
	FirewallIPvsixNameRuleInboundInterfaceInterfaceName  customtypes.CustomStringValue `tfsdk:"interface_name" json:"interface-name,omitempty"`
	FirewallIPvsixNameRuleInboundInterfaceInterfaceGroup customtypes.CustomStringValue `tfsdk:"interface_group" json:"interface-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o FirewallIPvsixNameRuleInboundInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Match interface

`,
		},

		"interface_group": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Match interface-group

`,
		},

		// TagNodes

		// Nodes

	}
}
