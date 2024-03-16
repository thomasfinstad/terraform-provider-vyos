// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthreeDefaultInformationOriginate{}

// VrfNameProtocolsOspfvthreeDefaultInformationOriginate describes the resource data model.
type VrfNameProtocolsOspfvthreeDefaultInformationOriginate struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeDefaultInformationOriginateAlways     types.Bool   `tfsdk:"always" vyos:"always,omitempty"`
	LeafVrfNameProtocolsOspfvthreeDefaultInformationOriginateMetric     types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsOspfvthreeDefaultInformationOriginateMetricType types.Number `tfsdk:"metric_type" vyos:"metric-type,omitempty"`
	LeafVrfNameProtocolsOspfvthreeDefaultInformationOriginateRouteMap   types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeDefaultInformationOriginate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"always": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Always advertise a default route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

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
