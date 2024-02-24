// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsRIPRedistributeBgp describes the resource data model.
type ProtocolsRIPRedistributeBgp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsRIPRedistributeBgpMetric   types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafProtocolsRIPRedistributeBgpRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsRIPRedistributeBgp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPRedistributeBgp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"rip",

		"redistribute",

		"bgp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPRedistributeBgp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric for redistributed routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-16  &emsp; |  Redistribute route metric  |

`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},
	}
}
