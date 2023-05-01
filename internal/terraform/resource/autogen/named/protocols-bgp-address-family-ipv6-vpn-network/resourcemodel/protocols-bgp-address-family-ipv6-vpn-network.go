// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpAddressFamilyIPvsixVpnNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixVpnNetwork struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd    types.String `tfsdk:"rd" json:"rd,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel types.String `tfsdk:"label" json:"label,omitempty"`

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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyIPvsixVpnNetwork) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd.IsNull() && !o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd.IsUnknown() {
		jsonData["rd"] = o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd.ValueString()
	}

	if !o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel.IsNull() && !o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel.IsUnknown() {
		jsonData["label"] = o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpAddressFamilyIPvsixVpnNetwork) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["rd"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkRd = basetypes.NewStringNull()
	}

	if value, ok := jsonData["label"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
