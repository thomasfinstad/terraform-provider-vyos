// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeListExport types.Number `tfsdk:"export" vyos:"export,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeListImport types.Number `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

    |  Format   |  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter outgoing route updates to this peer-group  |
`,
			Description: `Access-list to filter outgoing route updates to this peer-group

    |  Format   |  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter outgoing route updates to this peer-group  |
`,
		},

		"import": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter incoming route updates from this peer-group  |
`,
			Description: `Access-list to filter incoming route updates from this peer-group

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter incoming route updates from this peer-group  |
`,
		},

		// Nodes

	}
}
