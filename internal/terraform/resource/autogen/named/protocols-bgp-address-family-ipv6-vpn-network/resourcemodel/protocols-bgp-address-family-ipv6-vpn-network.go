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

// ProtocolsBgpAddressFamilyIPvsixVpnNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixVpnNetwork struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd    types.String `tfsdk:"rd"`
	LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel types.String `tfsdk:"label"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixVpnNetwork) GetVyosPath() []string {
	return []string{
		"protocols",
		"bgp",
		"address-family",
		"ipv6-vpn",
		"network",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpAddressFamilyIPvsixVpnNetwork) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-vpn", "network"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd.IsUnknown()) {
		vyosData["rd"] = o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd.ValueString()
	}
	if !(o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel.IsUnknown()) {
		vyosData["label"] = o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpAddressFamilyIPvsixVpnNetwork) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-vpn", "network"}})

	// Leafs
	if value, ok := vyosData["rd"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd = basetypes.NewStringNull()
	}
	if value, ok := vyosData["label"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-vpn", "network"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixVpnNetwork) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"rd":    types.StringType,
		"label": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixVpnNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Import BGP network/prefix into unicast VPN IPv6 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Unicast VPN IPv6 BGP network/prefix  |

`,
		},

		// LeafNodes

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		"label": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label value assigned to route

|  Format  |  Description  |
|----------|---------------|
|  u32:0-1048575  |  MPLS label value  |

`,
		},

		// TagNodes

		// Nodes

	}
}
