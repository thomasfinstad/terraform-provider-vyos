// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListExport customtypes.CustomStringValue `tfsdk:"export" json:"export,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListImport customtypes.CustomStringValue `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
		},

		"import": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
		},

		// TagNodes

		// Nodes

	}
}
