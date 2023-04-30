// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeList struct {
	// LeafNodes
	VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeListExport customtypes.CustomStringValue `tfsdk:"export" json:"export,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeListImport customtypes.CustomStringValue `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeList) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list to filter outgoing route updates to this peer-group  |
`,
		},

		"import": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list to filter incoming route updates from this peer-group  |
`,
		},

		// TagNodes

		// Nodes

	}
}
