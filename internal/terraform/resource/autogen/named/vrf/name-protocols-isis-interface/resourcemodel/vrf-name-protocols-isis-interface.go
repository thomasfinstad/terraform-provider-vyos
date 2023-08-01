// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisInterface describes the resource data model.
type VrfNameProtocolsIsisInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceCircuitType         types.String `tfsdk:"circuit_type" vyos:"circuit-type,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceHelloPadding        types.Bool   `tfsdk:"hello_padding" vyos:"hello-padding,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceHelloInterval       types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceHelloMultIPlier     types.Number `tfsdk:"hello_multiplier" vyos:"hello-multiplier,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceMetric              types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsIsisInterfacePassive             types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafVrfNameProtocolsIsisInterfacePriority            types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafVrfNameProtocolsIsisInterfacePsnpInterval        types.Number `tfsdk:"psnp_interval" vyos:"psnp-interval,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceNoThreeWayHandshake types.Bool   `tfsdk:"no_three_way_handshake" vyos:"no-three-way-handshake,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisInterfaceBfd      *VrfNameProtocolsIsisInterfaceBfd      `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeVrfNameProtocolsIsisInterfaceNetwork  *VrfNameProtocolsIsisInterfaceNetwork  `tfsdk:"network" vyos:"network,omitempty"`
	NodeVrfNameProtocolsIsisInterfacePassword *VrfNameProtocolsIsisInterfacePassword `tfsdk:"password" vyos:"password,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsIsisInterface) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"isis",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `interface_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface params

`,
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		// LeafNodes

		"circuit_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Configure circuit type for interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  level-1  &emsp; |  Level-1 only adjacencies are formed  |
    |  level-1-2  &emsp; |  Level-1-2 adjacencies are formed  |
    |  level-2-only  &emsp; |  Level-2 only adjacencies are formed  |

`,
		},

		"hello_padding": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Add padding to IS-IS hello packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set Hello interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-600  &emsp; |  Set Hello interval  |

`,
		},

		"hello_multiplier": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set Hello interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 2-100  &emsp; |  Set multiplier for Hello holding time  |

`,
		},

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set default metric for circuit

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777215  &emsp; |  Default metric value  |

`,
		},

		"passive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Configure passive mode for interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set priority for Designated Router election

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-127  &emsp; |  Priority value  |

`,
		},

		"psnp_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set PSNP interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-127  &emsp; |  PSNP interval in seconds  |

`,
		},

		"no_three_way_handshake": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable three-way handshake

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisInterfaceBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},

		"network": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisInterfaceNetwork{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Set network type

`,
		},

		"password": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisInterfacePassword{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Configure the authentication password for a circuit

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
