// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfRedistributeConnected{}

// VrfNameProtocolsOspfRedistributeConnected describes the resource data model.
type VrfNameProtocolsOspfRedistributeConnected struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfRedistributeConnectedMetric     types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsOspfRedistributeConnectedMetricType types.Number `tfsdk:"metric_type" vyos:"metric-type,omitempty"`
	LeafVrfNameProtocolsOspfRedistributeConnectedRouteMap   types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfRedistributeConnected) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `OSPF default metric

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777214  &emsp; |  Default metric  |

`,
		},

		"metric_type": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `OSPF metric type for default routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2  &emsp; |  Set OSPF External Type 1/2 metrics  |

`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		// Nodes

	}
}
