// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixListExport customtypes.CustomStringValue `tfsdk:"export" json:"export,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixListImport customtypes.CustomStringValue `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv4 prefix-list  |
`,
		},

		"import": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv4 prefix-list  |
`,
		},

		// TagNodes

		// Nodes

	}
}
