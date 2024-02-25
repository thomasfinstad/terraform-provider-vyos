// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsOspf describes the resource data model.
type ProtocolsOspf struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsOspfDefaultMetric    types.Number `tfsdk:"default_metric" vyos:"default-metric,omitempty"`
	LeafProtocolsOspfMaximumPaths     types.Number `tfsdk:"maximum_paths" vyos:"maximum-paths,omitempty"`
	LeafProtocolsOspfPassiveInterface types.String `tfsdk:"passive_interface" vyos:"passive-interface,omitempty"`
	LeafProtocolsOspfRouteMap         types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsOspfAccessList bool `tfsdk:"-" vyos:"access-list,child"`
	ExistsTagProtocolsOspfArea       bool `tfsdk:"-" vyos:"area,child"`
	ExistsTagProtocolsOspfInterface  bool `tfsdk:"-" vyos:"interface,child"`
	ExistsTagProtocolsOspfNeighbor   bool `tfsdk:"-" vyos:"neighbor,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsOspfAutoCost            bool `tfsdk:"-" vyos:"auto-cost,omitempty"`
	ExistsNodeProtocolsOspfDefaultInformation  bool `tfsdk:"-" vyos:"default-information,omitempty"`
	ExistsNodeProtocolsOspfDistance            bool `tfsdk:"-" vyos:"distance,omitempty"`
	ExistsNodeProtocolsOspfLogAdjacencyChanges bool `tfsdk:"-" vyos:"log-adjacency-changes,omitempty"`
	ExistsNodeProtocolsOspfMaxMetric           bool `tfsdk:"-" vyos:"max-metric,omitempty"`
	ExistsNodeProtocolsOspfMplsTe              bool `tfsdk:"-" vyos:"mpls-te,omitempty"`
	ExistsNodeProtocolsOspfParameters          bool `tfsdk:"-" vyos:"parameters,omitempty"`
	ExistsNodeProtocolsOspfSegmentRouting      bool `tfsdk:"-" vyos:"segment-routing,omitempty"`
	ExistsNodeProtocolsOspfRedistribute        bool `tfsdk:"-" vyos:"redistribute,omitempty"`
	ExistsNodeProtocolsOspfRefresh             bool `tfsdk:"-" vyos:"refresh,omitempty"`
	ExistsNodeProtocolsOspfTimers              bool `tfsdk:"-" vyos:"timers,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsOspf) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspf) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"ospf",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"default_metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric of redistributed routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777214  &emsp; |  Metric of redistributed routes  |

`,
		},

		"maximum_paths": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum multiple paths (ECMP)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-64  &emsp; |  Maximum multiple paths (ECMP)  |

`,
		},

		"passive_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Suppress routing updates on an interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  default  &emsp; |  Default to suppress routing updates on all interfaces  |

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
