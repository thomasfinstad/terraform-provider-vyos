// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesPseudoEthernetVifSVifC describes the resource data model.
type InterfacesPseudoEthernetVifSVifC struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetVifSVifCDescrIPtion       types.String `tfsdk:"description" json:"description,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCAddress           types.String `tfsdk:"address" json:"address,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCDisableLinkDetect types.String `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCDisable           types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCMac               types.String `tfsdk:"mac" json:"mac,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCMtu               types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCRedirect          types.String `tfsdk:"redirect" json:"redirect,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCVrf               types.String `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesPseudoEthernetVifSVifCDhcpOptions     *InterfacesPseudoEthernetVifSVifCDhcpOptions     `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	NodeInterfacesPseudoEthernetVifSVifCDhcpvsixOptions *InterfacesPseudoEthernetVifSVifCDhcpvsixOptions `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	NodeInterfacesPseudoEthernetVifSVifCIP              *InterfacesPseudoEthernetVifSVifCIP              `tfsdk:"ip" json:"ip,omitempty"`
	NodeInterfacesPseudoEthernetVifSVifCIPvsix          *InterfacesPseudoEthernetVifSVifCIPvsix          `tfsdk:"ipv6" json:"ipv6,omitempty"`
	NodeInterfacesPseudoEthernetVifSVifCMirror          *InterfacesPseudoEthernetVifSVifCMirror          `tfsdk:"mirror" json:"mirror,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVifSVifC) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifSVifCDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifSVifCDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifSVifCIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifSVifCIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifSVifCMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesPseudoEthernetVifSVifC) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesPseudoEthernetVifSVifCDescrIPtion.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafInterfacesPseudoEthernetVifSVifCDescrIPtion.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCAddress.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesPseudoEthernetVifSVifCAddress.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCDisableLinkDetect.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCDisableLinkDetect.IsUnknown() {
		jsonData["disable-link-detect"] = o.LeafInterfacesPseudoEthernetVifSVifCDisableLinkDetect.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCDisable.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesPseudoEthernetVifSVifCDisable.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCMac.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCMac.IsUnknown() {
		jsonData["mac"] = o.LeafInterfacesPseudoEthernetVifSVifCMac.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCMtu.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesPseudoEthernetVifSVifCMtu.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCRedirect.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCRedirect.IsUnknown() {
		jsonData["redirect"] = o.LeafInterfacesPseudoEthernetVifSVifCRedirect.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCVrf.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafInterfacesPseudoEthernetVifSVifCVrf.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetVifSVifCDhcpOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetVifSVifCDhcpOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetVifSVifCDhcpvsixOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetVifSVifCDhcpvsixOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetVifSVifCIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetVifSVifCIP)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetVifSVifCIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetVifSVifCIPvsix)
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

	if !reflect.ValueOf(o.NodeInterfacesPseudoEthernetVifSVifCMirror).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesPseudoEthernetVifSVifCMirror)
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
func (o *InterfacesPseudoEthernetVifSVifC) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-link-detect"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCDisableLinkDetect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCMac = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCMac = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCRedirect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["dhcp-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetVifSVifCDhcpOptions = &InterfacesPseudoEthernetVifSVifCDhcpOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetVifSVifCDhcpOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["dhcpv6-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetVifSVifCDhcpvsixOptions = &InterfacesPseudoEthernetVifSVifCDhcpvsixOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetVifSVifCDhcpvsixOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetVifSVifCIP = &InterfacesPseudoEthernetVifSVifCIP{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetVifSVifCIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetVifSVifCIPvsix = &InterfacesPseudoEthernetVifSVifCIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetVifSVifCIPvsix)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["mirror"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesPseudoEthernetVifSVifCMirror = &InterfacesPseudoEthernetVifSVifCMirror{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesPseudoEthernetVifSVifCMirror)
		if err != nil {
			return err
		}
	}

	return nil
}
