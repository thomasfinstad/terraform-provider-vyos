// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsIsisInterface{}

// VrfNameProtocolsIsisInterface describes the resource data model.
type VrfNameProtocolsIsisInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

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
	NodeVrfNameProtocolsIsisInterfaceLdpSync  *VrfNameProtocolsIsisInterfaceLdpSync  `tfsdk:"ldp_sync" vyos:"ldp-sync,omitempty"`
	NodeVrfNameProtocolsIsisInterfaceNetwork  *VrfNameProtocolsIsisInterfaceNetwork  `tfsdk:"network" vyos:"network,omitempty"`
	NodeVrfNameProtocolsIsisInterfacePassword *VrfNameProtocolsIsisInterfacePassword `tfsdk:"password" vyos:"password,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsIsisInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsIsisInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsIsisInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsIsisInterface) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"isis",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsIsisInterface) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface params

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in interface_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  interface_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in name_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
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

		"ldp_sync": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisInterfaceLdpSync{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `LDP-IGP synchronization configuration for interface

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
