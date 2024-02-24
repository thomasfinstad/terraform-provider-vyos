// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsMplsLdpTargetedNeighborIPvfour describes the resource data model.
type ProtocolsMplsLdpTargetedNeighborIPvfour struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsMplsLdpTargetedNeighborIPvfourAddress       types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafProtocolsMplsLdpTargetedNeighborIPvfourEnable        types.Bool   `tfsdk:"enable" vyos:"enable,omitempty"`
	LeafProtocolsMplsLdpTargetedNeighborIPvfourHelloInterval types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsMplsLdpTargetedNeighborIPvfourHelloHoldtime types.Number `tfsdk:"hello_holdtime" vyos:"hello-holdtime,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"mpls",

		"ldp",

		"targeted-neighbor",

		"ipv4",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpTargetedNeighborIPvfour) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Neighbor/session address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Neighbor/session address  |

`,
		},

		"enable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Accept and respond to targeted hellos

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |

`,
		},

		"hello_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |

`,
		},
	}
}
