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

// ProtocolsBgpAddressFamilyLtwovpnEvpnVni describes the resource data model.
type ProtocolsBgpAddressFamilyLtwovpnEvpnVni struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"vni_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw types.Bool   `tfsdk:"advertise_default_gw" vyos:"advertise-default-gw,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP     types.Bool   `tfsdk:"advertise_svi_ip" vyos:"advertise-svi-ip,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd                 types.String `tfsdk:"rd" vyos:"rd,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget *ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget `tfsdk:"route_target" vyos:"route-target,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",

		"address-family",

		"l2vpn-evpn",

		"vni",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpnVni) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"vni_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VXLAN Network Identifier

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-16777215  &emsp; |  VNI number  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"advertise_default_gw": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_svi_ip": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  &emsp; |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		// Nodes

		"route_target": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route Target

`,
		},
	}
}
