// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessVifS describes the resource data model.
type InterfacesWirelessVifS struct {
	// LeafNodes
	LeafInterfacesWirelessVifSDescrIPtion       types.String `tfsdk:"description" json:"description,omitempty"`
	LeafInterfacesWirelessVifSAddress           types.String `tfsdk:"address" json:"address,omitempty"`
	LeafInterfacesWirelessVifSDisableLinkDetect types.String `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	LeafInterfacesWirelessVifSDisable           types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafInterfacesWirelessVifSProtocol          types.String `tfsdk:"protocol" json:"protocol,omitempty"`
	LeafInterfacesWirelessVifSMac               types.String `tfsdk:"mac" json:"mac,omitempty"`
	LeafInterfacesWirelessVifSMtu               types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesWirelessVifSRedirect          types.String `tfsdk:"redirect" json:"redirect,omitempty"`
	LeafInterfacesWirelessVifSVrf               types.String `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes
	TagInterfacesWirelessVifSVifC *map[string]InterfacesWirelessVifSVifC `tfsdk:"vif_c" json:"vif-c,omitempty"`

	// Nodes
	NodeInterfacesWirelessVifSDhcpOptions     *InterfacesWirelessVifSDhcpOptions     `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	NodeInterfacesWirelessVifSDhcpvsixOptions *InterfacesWirelessVifSDhcpvsixOptions `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	NodeInterfacesWirelessVifSIP              *InterfacesWirelessVifSIP              `tfsdk:"ip" json:"ip,omitempty"`
	NodeInterfacesWirelessVifSIPvsix          *InterfacesWirelessVifSIPvsix          `tfsdk:"ipv6" json:"ipv6,omitempty"`
	NodeInterfacesWirelessVifSMirror          *InterfacesWirelessVifSMirror          `tfsdk:"mirror" json:"mirror,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifS) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
|  dhcp  |  Dynamic Host Configuration Protocol  |
|  dhcpv6  |  Dynamic Host Configuration Protocol for IPv6  |

`,
		},

		"disable_link_detect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol used for service VLAN (default: 802.1ad)

|  Format  |  Description  |
|----------|---------------|
|  802.1ad  |  Provider Bridging (IEEE 802.1ad, Q-inQ), ethertype 0x88a8  |
|  802.1q  |  VLAN-tagged frame (IEEE 802.1q), ethertype 0x8100  |

`,

			// Default:          stringdefault.StaticString(`802.1ad`),
			Computed: true,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

`,
		},

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		// TagNodes

		"vif_c": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesWirelessVifSVifC{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

`,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessVifS) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessVifSDescrIPtion.IsNull() && !o.LeafInterfacesWirelessVifSDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafInterfacesWirelessVifSDescrIPtion.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSAddress.IsNull() && !o.LeafInterfacesWirelessVifSAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesWirelessVifSAddress.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSDisableLinkDetect.IsNull() && !o.LeafInterfacesWirelessVifSDisableLinkDetect.IsUnknown() {
		jsonData["disable-link-detect"] = o.LeafInterfacesWirelessVifSDisableLinkDetect.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSDisable.IsNull() && !o.LeafInterfacesWirelessVifSDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesWirelessVifSDisable.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSProtocol.IsNull() && !o.LeafInterfacesWirelessVifSProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafInterfacesWirelessVifSProtocol.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSMac.IsNull() && !o.LeafInterfacesWirelessVifSMac.IsUnknown() {
		jsonData["mac"] = o.LeafInterfacesWirelessVifSMac.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSMtu.IsNull() && !o.LeafInterfacesWirelessVifSMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesWirelessVifSMtu.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSRedirect.IsNull() && !o.LeafInterfacesWirelessVifSRedirect.IsUnknown() {
		jsonData["redirect"] = o.LeafInterfacesWirelessVifSRedirect.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSVrf.IsNull() && !o.LeafInterfacesWirelessVifSVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafInterfacesWirelessVifSVrf.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagInterfacesWirelessVifSVifC).IsZero() {
		subJSONStr, err := json.Marshal(o.TagInterfacesWirelessVifSVifC)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["vif-c"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesWirelessVifSDhcpOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWirelessVifSDhcpOptions)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["dhcp-options"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesWirelessVifSDhcpvsixOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWirelessVifSDhcpvsixOptions)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["dhcpv6-options"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesWirelessVifSIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWirelessVifSIP)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ip"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesWirelessVifSIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWirelessVifSIPvsix)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesWirelessVifSMirror).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWirelessVifSMirror)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["mirror"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessVifS) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafInterfacesWirelessVifSDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesWirelessVifSAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-link-detect"]; ok {
		o.LeafInterfacesWirelessVifSDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSDisableLinkDetect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesWirelessVifSDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafInterfacesWirelessVifSProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSProtocol = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac"]; ok {
		o.LeafInterfacesWirelessVifSMac = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSMac = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesWirelessVifSMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect"]; ok {
		o.LeafInterfacesWirelessVifSRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSRedirect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafInterfacesWirelessVifSVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVrf = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["vif-c"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagInterfacesWirelessVifSVifC = &map[string]InterfacesWirelessVifSVifC{}

		err = json.Unmarshal(subJSONStr, o.TagInterfacesWirelessVifSVifC)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["dhcp-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWirelessVifSDhcpOptions = &InterfacesWirelessVifSDhcpOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWirelessVifSDhcpOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["dhcpv6-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWirelessVifSDhcpvsixOptions = &InterfacesWirelessVifSDhcpvsixOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWirelessVifSDhcpvsixOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWirelessVifSIP = &InterfacesWirelessVifSIP{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWirelessVifSIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWirelessVifSIPvsix = &InterfacesWirelessVifSIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWirelessVifSIPvsix)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["mirror"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWirelessVifSMirror = &InterfacesWirelessVifSMirror{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWirelessVifSMirror)
		if err != nil {
			return err
		}
	}

	return nil
}
