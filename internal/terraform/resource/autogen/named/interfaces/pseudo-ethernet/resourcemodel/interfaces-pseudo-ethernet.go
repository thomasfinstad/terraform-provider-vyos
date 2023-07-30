// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesPseudoEthernet describes the resource data model.
type InterfacesPseudoEthernet struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesPseudoEthernetAddress           types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesPseudoEthernetDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesPseudoEthernetDisableLinkDetect types.String `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesPseudoEthernetDisable           types.String `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesPseudoEthernetVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesPseudoEthernetSourceInterface   types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesPseudoEthernetMac               types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesPseudoEthernetMode              types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafInterfacesPseudoEthernetMtu               types.String `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesPseudoEthernetRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesPseudoEthernetVifS bool `tfsdk:"vif_s" vyos:"vif-s,child"`
	ExistsTagInterfacesPseudoEthernetVif  bool `tfsdk:"vif" vyos:"vif,child"`

	// Nodes
	NodeInterfacesPseudoEthernetDhcpOptions     *InterfacesPseudoEthernetDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesPseudoEthernetDhcpvsixOptions *InterfacesPseudoEthernetDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesPseudoEthernetIP              *InterfacesPseudoEthernetIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesPseudoEthernetIPvsix          *InterfacesPseudoEthernetIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesPseudoEthernetMirror          *InterfacesPseudoEthernetMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPseudoEthernet) GetVyosPath() []string {
	return []string{
		"interfaces",
		"pseudo-ethernet",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pseudo Ethernet Interface (Macvlan)

    |  Format  |  Description  |
    |----------|---------------|
    |  pethN  |  Pseudo Ethernet interface name  |

`,
		},

		// LeafNodes

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

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Description  |

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

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  VRF instance name  |

`,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Physical interface the traffic will go through

    |  Format  |  Description  |
    |----------|---------------|
    |  interface  |  Physical interface used for traffic forwarding  |

`,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format  |  Description  |
    |----------|---------------|
    |  macaddr  |  Hardware (MAC) address  |

`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Receive mode (default: private)

    |  Format  |  Description  |
    |----------|---------------|
    |  private  |  No communication with other pseudo-devices  |
    |  vepa  |  Virtual Ethernet Port Aggregator reflective relay  |
    |  bridge  |  Simple bridge between pseudo-devices  |
    |  passthru  |  Promicious mode passthrough of underlying device  |

`,

			// Default:          stringdefault.StaticString(`private`),
			Computed: true,
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

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesPseudoEthernet) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesPseudoEthernetAddress.IsNull() && !o.LeafInterfacesPseudoEthernetAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesPseudoEthernetAddress.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDescrIPtion.IsNull() && !o.LeafInterfacesPseudoEthernetDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafInterfacesPseudoEthernetDescrIPtion.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDisableLinkDetect.IsNull() && !o.LeafInterfacesPseudoEthernetDisableLinkDetect.IsUnknown() {
		jsonData["disable-link-detect"] = o.LeafInterfacesPseudoEthernetDisableLinkDetect.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDisable.IsNull() && !o.LeafInterfacesPseudoEthernetDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesPseudoEthernetDisable.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVrf.IsNull() && !o.LeafInterfacesPseudoEthernetVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafInterfacesPseudoEthernetVrf.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetSourceInterface.IsNull() && !o.LeafInterfacesPseudoEthernetSourceInterface.IsUnknown() {
		jsonData["source-interface"] = o.LeafInterfacesPseudoEthernetSourceInterface.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetMac.IsNull() && !o.LeafInterfacesPseudoEthernetMac.IsUnknown() {
		jsonData["mac"] = o.LeafInterfacesPseudoEthernetMac.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetMode.IsNull() && !o.LeafInterfacesPseudoEthernetMode.IsUnknown() {
		jsonData["mode"] = o.LeafInterfacesPseudoEthernetMode.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetMtu.IsNull() && !o.LeafInterfacesPseudoEthernetMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesPseudoEthernetMtu.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetRedirect.IsNull() && !o.LeafInterfacesPseudoEthernetRedirect.IsUnknown() {
		jsonData["redirect"] = o.LeafInterfacesPseudoEthernetRedirect.ValueString()
	}

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetDhcpOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetDhcpOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetDhcpvsixOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetDhcpvsixOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetIP)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetIPvsix)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetMirror).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetMirror)
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
func (o *InterfacesPseudoEthernet) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesPseudoEthernetAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafInterfacesPseudoEthernetDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-link-detect"]; ok {
		o.LeafInterfacesPseudoEthernetDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDisableLinkDetect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesPseudoEthernetDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafInterfacesPseudoEthernetVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVrf = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-interface"]; ok {
		o.LeafInterfacesPseudoEthernetSourceInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetSourceInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac"]; ok {
		o.LeafInterfacesPseudoEthernetMac = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetMac = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mode"]; ok {
		o.LeafInterfacesPseudoEthernetMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetMode = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesPseudoEthernetMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect"]; ok {
		o.LeafInterfacesPseudoEthernetRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetRedirect = basetypes.NewStringNull()
	}

	// Nodes
	if value, ok := jsonData["dhcp-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetDhcpOptions = &InterfacesPseudoEthernetDhcpOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetDhcpOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["dhcpv6-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetDhcpvsixOptions = &InterfacesPseudoEthernetDhcpvsixOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetDhcpvsixOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetIP = &InterfacesPseudoEthernetIP{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetIPvsix = &InterfacesPseudoEthernetIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetIPvsix)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["mirror"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetMirror = &InterfacesPseudoEthernetMirror{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetMirror)
		if err != nil {
			return err
		}
	}

	return nil
}
