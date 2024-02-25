// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgp describes the resource data model.
type ProtocolsBgp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsBgpSystemAs types.Number `tfsdk:"system_as" vyos:"system-as,omitempty"`
	LeafProtocolsBgpRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsBgpNeighbor  bool `tfsdk:"-" vyos:"neighbor,child"`
	ExistsTagProtocolsBgpPeerGroup bool `tfsdk:"-" vyos:"peer-group,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsBgpAddressFamily bool `tfsdk:"-" vyos:"address-family,omitempty"`
	ExistsNodeProtocolsBgpListen        bool `tfsdk:"-" vyos:"listen,omitempty"`
	ExistsNodeProtocolsBgpParameters    bool `tfsdk:"-" vyos:"parameters,omitempty"`
	ExistsNodeProtocolsBgpTimers        bool `tfsdk:"-" vyos:"timers,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"system_as": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Autonomous System Number (ASN)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Autonomous System Number  |

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
