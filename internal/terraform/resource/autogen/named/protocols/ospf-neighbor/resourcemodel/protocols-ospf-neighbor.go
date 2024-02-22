// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsOspfNeighbor describes the resource data model.
type ProtocolsOspfNeighbor struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"neighbor_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsOspfNeighborPollInterval types.Number `tfsdk:"poll_interval" vyos:"poll-interval,omitempty"`
	LeafProtocolsOspfNeighborPriority     types.Number `tfsdk:"priority" vyos:"priority,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsOspfNeighbor) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfNeighbor) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"ospf",

		"neighbor",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfNeighbor) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"neighbor_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Specify neighbor router

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Neighbor IP address  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"poll_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Dead neighbor polling interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Seconds between dead neighbor polling interval  |

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor priority in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  Neighbor priority  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// Nodes

	}
}
