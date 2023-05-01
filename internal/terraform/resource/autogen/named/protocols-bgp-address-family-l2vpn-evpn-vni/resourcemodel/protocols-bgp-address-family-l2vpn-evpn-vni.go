// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ProtocolsBgpAddressFamilyLtwovpnEvpnVni describes the resource data model.
type ProtocolsBgpAddressFamilyLtwovpnEvpnVni struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw types.String `tfsdk:"advertise_default_gw"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP     types.String `tfsdk:"advertise_svi_ip"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd                 types.String `tfsdk:"rd"`

	// TagNodes

	// Nodes
	NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget types.Object `tfsdk:"route_target"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosPath() []string {
	return []string{
		"protocols",
		"bgp",
		"address-family",
		"l2vpn-evpn",
		"vni",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "l2vpn-evpn", "vni"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw.IsNull() || o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw.IsUnknown()) {
		vyosData["advertise-default-gw"] = o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw.ValueString()
	}
	if !(o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP.IsNull() || o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP.IsUnknown()) {
		vyosData["advertise-svi-ip"] = o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP.ValueString()
	}
	if !(o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd.IsNull() || o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd.IsUnknown()) {
		vyosData["rd"] = o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget.IsNull() || o.NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget.IsUnknown()) {
		var subModel ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget
		diags.Append(o.NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["route-target"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "l2vpn-evpn", "vni"}})

	// Leafs
	if value, ok := vyosData["advertise-default-gw"]; ok {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw = basetypes.NewStringNull()
	}
	if value, ok := vyosData["advertise-svi-ip"]; ok {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rd"]; ok {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["route-target"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget = data

	} else {
		o.NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget = basetypes.NewObjectNull(ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "l2vpn-evpn", "vni"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpnVni) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"advertise_default_gw": types.StringType,
		"advertise_svi_ip":     types.StringType,
		"rd":                   types.StringType,

		// Tags

		// Nodes
		"route_target": types.ObjectType{AttrTypes: ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpnVni) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VXLAN Network Identifier

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16777215  |  VNI number  |

`,
		},

		// LeafNodes

		"advertise_default_gw": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
		},

		"advertise_svi_ip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		// TagNodes

		// Nodes

		"route_target": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route Target

`,
		},
	}
}