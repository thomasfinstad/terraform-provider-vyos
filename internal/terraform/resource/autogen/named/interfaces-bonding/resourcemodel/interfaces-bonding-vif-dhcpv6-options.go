// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesBondingVifDhcpvsixOptions describes the resource data model.
type InterfacesBondingVifDhcpvsixOptions struct {
	// LeafNodes
	InterfacesBondingVifDhcpvsixOptionsDuID           customtypes.CustomStringValue `tfsdk:"duid" json:"duid,omitempty"`
	InterfacesBondingVifDhcpvsixOptionsParametersOnly customtypes.CustomStringValue `tfsdk:"parameters_only" json:"parameters-only,omitempty"`
	InterfacesBondingVifDhcpvsixOptionsRAPIDCommit    customtypes.CustomStringValue `tfsdk:"rapid_commit" json:"rapid-commit,omitempty"`
	InterfacesBondingVifDhcpvsixOptionsTemporary      customtypes.CustomStringValue `tfsdk:"temporary" json:"temporary,omitempty"`

	// TagNodes
	InterfacesBondingVifDhcpvsixOptionsPd types.Map `tfsdk:"pd" json:"pd,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesBondingVifDhcpvsixOptions) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"duid": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

|  Format  |  Description  |
|----------|---------------|
|  duid  |  DHCP unique identifier (DUID)  |
`,
		},

		"parameters_only": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Acquire only config parameters, no address

`,
		},

		"rapid_commit": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
		},

		"temporary": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv6 temporary address

`,
		},

		// TagNodes

		"pd": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesBondingVifDhcpvsixOptionsPd{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

|  Format  |  Description  |
|----------|---------------|
|  instance number  |  Prefix delegation instance (>= 0)  |
`,
		},

		// Nodes

	}
}
