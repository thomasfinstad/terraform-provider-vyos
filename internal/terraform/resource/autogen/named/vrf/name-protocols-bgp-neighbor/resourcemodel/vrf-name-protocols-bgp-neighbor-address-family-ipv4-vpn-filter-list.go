// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterListExport types.String `tfsdk:"export" vyos:"export,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterListImport types.String `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
			Description: `As-path-list to filter outgoing route updates to this peer

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
			Description: `As-path-list to filter incoming route updates from this peer

`,
		},

		// Nodes

	}
}
