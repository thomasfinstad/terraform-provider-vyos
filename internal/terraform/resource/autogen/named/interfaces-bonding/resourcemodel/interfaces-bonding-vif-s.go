// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBondingVifS describes the resource data model.
type InterfacesBondingVifS struct {
	// LeafNodes
	LeafInterfacesBondingVifSDescrIPtion       types.String `tfsdk:"description" json:"description,omitempty"`
	LeafInterfacesBondingVifSAddress           types.String `tfsdk:"address" json:"address,omitempty"`
	LeafInterfacesBondingVifSDisableLinkDetect types.String `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	LeafInterfacesBondingVifSDisable           types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafInterfacesBondingVifSProtocol          types.String `tfsdk:"protocol" json:"protocol,omitempty"`
	LeafInterfacesBondingVifSMac               types.String `tfsdk:"mac" json:"mac,omitempty"`
	LeafInterfacesBondingVifSMtu               types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesBondingVifSRedirect          types.String `tfsdk:"redirect" json:"redirect,omitempty"`
	LeafInterfacesBondingVifSVrf               types.String `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes
	TagInterfacesBondingVifSVifC *map[string]InterfacesBondingVifSVifC `tfsdk:"vif_c" json:"vif-c,omitempty"`

	// Nodes
	NodeInterfacesBondingVifSDhcpOptions     *InterfacesBondingVifSDhcpOptions     `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	NodeInterfacesBondingVifSDhcpvsixOptions *InterfacesBondingVifSDhcpvsixOptions `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	NodeInterfacesBondingVifSIP              *InterfacesBondingVifSIP              `tfsdk:"ip" json:"ip,omitempty"`
	NodeInterfacesBondingVifSIPvsix          *InterfacesBondingVifSIPvsix          `tfsdk:"ipv6" json:"ipv6,omitempty"`
	NodeInterfacesBondingVifSMirror          *InterfacesBondingVifSMirror          `tfsdk:"mirror" json:"mirror,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifS) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesBondingVifSVifC{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

`,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBondingVifS) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBondingVifSDescrIPtion.IsNull() && !o.LeafInterfacesBondingVifSDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafInterfacesBondingVifSDescrIPtion.ValueString()
	}

	if !o.LeafInterfacesBondingVifSAddress.IsNull() && !o.LeafInterfacesBondingVifSAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesBondingVifSAddress.ValueString()
	}

	if !o.LeafInterfacesBondingVifSDisableLinkDetect.IsNull() && !o.LeafInterfacesBondingVifSDisableLinkDetect.IsUnknown() {
		jsonData["disable-link-detect"] = o.LeafInterfacesBondingVifSDisableLinkDetect.ValueString()
	}

	if !o.LeafInterfacesBondingVifSDisable.IsNull() && !o.LeafInterfacesBondingVifSDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesBondingVifSDisable.ValueString()
	}

	if !o.LeafInterfacesBondingVifSProtocol.IsNull() && !o.LeafInterfacesBondingVifSProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafInterfacesBondingVifSProtocol.ValueString()
	}

	if !o.LeafInterfacesBondingVifSMac.IsNull() && !o.LeafInterfacesBondingVifSMac.IsUnknown() {
		jsonData["mac"] = o.LeafInterfacesBondingVifSMac.ValueString()
	}

	if !o.LeafInterfacesBondingVifSMtu.IsNull() && !o.LeafInterfacesBondingVifSMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesBondingVifSMtu.ValueString()
	}

	if !o.LeafInterfacesBondingVifSRedirect.IsNull() && !o.LeafInterfacesBondingVifSRedirect.IsUnknown() {
		jsonData["redirect"] = o.LeafInterfacesBondingVifSRedirect.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVrf.IsNull() && !o.LeafInterfacesBondingVifSVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafInterfacesBondingVifSVrf.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagInterfacesBondingVifSVifC).IsZero() {
		subJSONStr, err := json.Marshal(o.TagInterfacesBondingVifSVifC)
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

	if !reflect.ValueOf(o.NodeInterfacesBondingVifSDhcpOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBondingVifSDhcpOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesBondingVifSDhcpvsixOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBondingVifSDhcpvsixOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesBondingVifSIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBondingVifSIP)
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

	if !reflect.ValueOf(o.NodeInterfacesBondingVifSIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBondingVifSIPvsix)
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

	if !reflect.ValueOf(o.NodeInterfacesBondingVifSMirror).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBondingVifSMirror)
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
func (o *InterfacesBondingVifS) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafInterfacesBondingVifSDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesBondingVifSAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-link-detect"]; ok {
		o.LeafInterfacesBondingVifSDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSDisableLinkDetect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesBondingVifSDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafInterfacesBondingVifSProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSProtocol = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac"]; ok {
		o.LeafInterfacesBondingVifSMac = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSMac = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesBondingVifSMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect"]; ok {
		o.LeafInterfacesBondingVifSRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSRedirect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafInterfacesBondingVifSVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVrf = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["vif-c"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagInterfacesBondingVifSVifC = &map[string]InterfacesBondingVifSVifC{}

		err = json.Unmarshal(subJSONStr, o.TagInterfacesBondingVifSVifC)
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

		o.NodeInterfacesBondingVifSDhcpOptions = &InterfacesBondingVifSDhcpOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBondingVifSDhcpOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["dhcpv6-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBondingVifSDhcpvsixOptions = &InterfacesBondingVifSDhcpvsixOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBondingVifSDhcpvsixOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBondingVifSIP = &InterfacesBondingVifSIP{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBondingVifSIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBondingVifSIPvsix = &InterfacesBondingVifSIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBondingVifSIPvsix)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["mirror"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBondingVifSMirror = &InterfacesBondingVifSMirror{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBondingVifSMirror)
		if err != nil {
			return err
		}
	}

	return nil
}
