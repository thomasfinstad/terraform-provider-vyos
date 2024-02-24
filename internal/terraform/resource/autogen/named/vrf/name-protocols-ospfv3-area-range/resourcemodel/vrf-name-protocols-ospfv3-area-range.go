// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsOspfvthreeAreaRange describes the resource data model.
type VrfNameProtocolsOspfvthreeAreaRange struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"range_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	ParentIDVrfNameProtocolsOspfvthreeArea types.String `tfsdk:"area" vyos:"area,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeAreaRangeAdvertise    types.Bool `tfsdk:"advertise" vyos:"advertise,omitempty"`
	LeafVrfNameProtocolsOspfvthreeAreaRangeNotAdvertise types.Bool `tfsdk:"not_advertise" vyos:"not-advertise,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameProtocolsOspfvthreeAreaRange) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfvthreeAreaRange) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospfv3",

		"area",
		o.ParentIDVrfNameProtocolsOspfvthreeArea.ValueString(),

		"range",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeAreaRange) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"range_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Specify IPv6 prefix (border routers only)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  Specify IPv6 prefix (border routers only)  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
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
			},
		},

		"area_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OSPFv3 Area

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Area ID as a decimal value  |
    |  ipv4  &emsp; |  Area ID in IP address forma  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"advertise": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise this range

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"not_advertise": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not advertise this range

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
