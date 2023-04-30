// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// FirewallInterfaceLocal describes the resource data model.
type FirewallInterfaceLocal struct {
	// LeafNodes
	FirewallInterfaceLocalName       customtypes.CustomStringValue `tfsdk:"name" json:"name,omitempty"`
	FirewallInterfaceLocalIPvsixName customtypes.CustomStringValue `tfsdk:"ipv6_name" json:"ipv6-name,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o FirewallInterfaceLocal) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Local IPv4 firewall ruleset name for interface

`,
		},

		"ipv6_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Local IPv6 firewall ruleset name for interface

`,
		},

		// TagNodes

		// Nodes

	}
}
