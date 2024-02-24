// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsMplsLdpExportIPvsixExportFilter describes the resource data model.
type ProtocolsMplsLdpExportIPvsixExportFilter struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsMplsLdpExportIPvsixExportFilterFilterAccessListsix   types.Number `tfsdk:"filter_access_list6" vyos:"filter-access-list6,omitempty"`
	LeafProtocolsMplsLdpExportIPvsixExportFilterNeighborAccessListsix types.Number `tfsdk:"neighbor_access_list6" vyos:"neighbor-access-list6,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"mpls",

		"ldp",

		"export",

		"ipv6",

		"export-filter",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpExportIPvsixExportFilter) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"filter_access_list6": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list6 number to apply FEC filtering

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2699  &emsp; |  Access list number  |

`,
		},

		"neighbor_access_list6": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list6 number for IPv6 neighbor selection to apply filtering

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2699  &emsp; |  Access list number  |

`,
		},
	}
}
