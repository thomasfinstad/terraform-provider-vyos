// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseAllVni    types.Bool   `tfsdk:"advertise_all_vni" vyos:"advertise-all-vni,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseDefaultGw types.Bool   `tfsdk:"advertise_default_gw" vyos:"advertise-default-gw,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseSviIP     types.Bool   `tfsdk:"advertise_svi_ip" vyos:"advertise-svi-ip,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRd                 types.String `tfsdk:"rd" vyos:"rd,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertisePIP       types.String `tfsdk:"advertise_pip" vyos:"advertise-pip,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRtAutoDerive       types.Bool   `tfsdk:"rt_auto_derive" vyos:"rt-auto-derive,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDisableEadEviRx    types.Bool   `tfsdk:"disable_ead_evi_rx" vyos:"disable-ead-evi-rx,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDisableEadEviTx    types.Bool   `tfsdk:"disable_ead_evi_tx" vyos:"disable-ead-evi-tx,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni bool `tfsdk:"vni" vyos:"vni,child"`

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise        *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise        `tfsdk:"advertise" vyos:"advertise,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget      *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget      `tfsdk:"route_target" vyos:"route-target,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate `tfsdk:"default_originate" vyos:"default-originate,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag        *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag        `tfsdk:"ead_es_frag" vyos:"ead-es-frag,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget `tfsdk:"ead_es_route_target" vyos:"ead-es-route-target,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding         *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding         `tfsdk:"flooding" vyos:"flooding,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf           *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf           `tfsdk:"mac_vrf" vyos:"mac-vrf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise_all_vni": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All local VNIs

`,
			Description: `Advertise All local VNIs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_default_gw": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
			Description: `Advertise All default g/w mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_svi_ip": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
			Description: `Advertise svi mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
			Description: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
		},

		"advertise_pip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `EVPN system primary IP

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4    |  IP address   |
`,
			Description: `EVPN system primary IP

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4    |  IP address   |
`,
		},

		"rt_auto_derive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Auto derivation of Route Target (RFC8365)

`,
			Description: `Auto derivation of Route Target (RFC8365)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_ead_evi_rx": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Activate PE on EAD-ES even if EAD-EVI is not received

`,
			Description: `Activate PE on EAD-ES even if EAD-EVI is not received

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_ead_evi_tx": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not advertise EAD-EVI for local ESs

`,
			Description: `Do not advertise EAD-EVI for local ESs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise prefix routes

`,
			Description: `Advertise prefix routes

`,
		},

		"route_target": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route Target

`,
			Description: `Route Target

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Originate a default route

`,
			Description: `Originate a default route

`,
		},

		"ead_es_frag": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `EAD ES fragment config

`,
			Description: `EAD ES fragment config

`,
		},

		"ead_es_route_target": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `EAD ES Route Target

`,
			Description: `EAD ES Route Target

`,
		},

		"flooding": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify handling for BUM packets

`,
			Description: `Specify handling for BUM packets

`,
		},

		"mac_vrf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `EVPN MAC-VRF

`,
			Description: `EVPN MAC-VRF

`,
		},
	}
}
