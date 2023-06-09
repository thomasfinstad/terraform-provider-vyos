// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWwan describes the resource data model.
type InterfacesWwan struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafInterfacesWwanAddress           types.String `tfsdk:"address" json:"address,omitempty"`
	LeafInterfacesWwanApn               types.String `tfsdk:"apn" json:"apn,omitempty"`
	LeafInterfacesWwanDescrIPtion       types.String `tfsdk:"description" json:"description,omitempty"`
	LeafInterfacesWwanDisable           types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafInterfacesWwanDisableLinkDetect types.String `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	LeafInterfacesWwanMtu               types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesWwanConnectOnDemand   types.String `tfsdk:"connect_on_demand" json:"connect-on-demand,omitempty"`
	LeafInterfacesWwanRedirect          types.String `tfsdk:"redirect" json:"redirect,omitempty"`
	LeafInterfacesWwanVrf               types.String `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesWwanDhcpOptions     *InterfacesWwanDhcpOptions     `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	NodeInterfacesWwanDhcpvsixOptions *InterfacesWwanDhcpvsixOptions `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	NodeInterfacesWwanAuthentication  *InterfacesWwanAuthentication  `tfsdk:"authentication" json:"authentication,omitempty"`
	NodeInterfacesWwanMirror          *InterfacesWwanMirror          `tfsdk:"mirror" json:"mirror,omitempty"`
	NodeInterfacesWwanIP              *InterfacesWwanIP              `tfsdk:"ip" json:"ip,omitempty"`
	NodeInterfacesWwanIPvsix          *InterfacesWwanIPvsix          `tfsdk:"ipv6" json:"ipv6,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWwan) GetVyosPath() []string {
	return []string{
		"interfaces",
		"wwan",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWwan) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Wireless Modem (WWAN) Interface

|  Format  |  Description  |
|----------|---------------|
|  wwanN  |  Wireless Wide Area Network interface name  |

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

		"apn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access Point Name (APN)

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

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"disable_link_detect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-1500  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1430`),
			Computed: true,
		},

		"connect_on_demand": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Establishment connection automatically when traffic is sent

`,
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
			Attributes: InterfacesWwanDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWwan) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWwanAddress.IsNull() && !o.LeafInterfacesWwanAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesWwanAddress.ValueString()
	}

	if !o.LeafInterfacesWwanApn.IsNull() && !o.LeafInterfacesWwanApn.IsUnknown() {
		jsonData["apn"] = o.LeafInterfacesWwanApn.ValueString()
	}

	if !o.LeafInterfacesWwanDescrIPtion.IsNull() && !o.LeafInterfacesWwanDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafInterfacesWwanDescrIPtion.ValueString()
	}

	if !o.LeafInterfacesWwanDisable.IsNull() && !o.LeafInterfacesWwanDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesWwanDisable.ValueString()
	}

	if !o.LeafInterfacesWwanDisableLinkDetect.IsNull() && !o.LeafInterfacesWwanDisableLinkDetect.IsUnknown() {
		jsonData["disable-link-detect"] = o.LeafInterfacesWwanDisableLinkDetect.ValueString()
	}

	if !o.LeafInterfacesWwanMtu.IsNull() && !o.LeafInterfacesWwanMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesWwanMtu.ValueString()
	}

	if !o.LeafInterfacesWwanConnectOnDemand.IsNull() && !o.LeafInterfacesWwanConnectOnDemand.IsUnknown() {
		jsonData["connect-on-demand"] = o.LeafInterfacesWwanConnectOnDemand.ValueString()
	}

	if !o.LeafInterfacesWwanRedirect.IsNull() && !o.LeafInterfacesWwanRedirect.IsUnknown() {
		jsonData["redirect"] = o.LeafInterfacesWwanRedirect.ValueString()
	}

	if !o.LeafInterfacesWwanVrf.IsNull() && !o.LeafInterfacesWwanVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafInterfacesWwanVrf.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesWwanDhcpOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWwanDhcpOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesWwanDhcpvsixOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWwanDhcpvsixOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesWwanAuthentication).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWwanAuthentication)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["authentication"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesWwanMirror).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWwanMirror)
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

	if !reflect.ValueOf(o.NodeInterfacesWwanIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWwanIP)
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

	if !reflect.ValueOf(o.NodeInterfacesWwanIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesWwanIPvsix)
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

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWwan) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesWwanAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["apn"]; ok {
		o.LeafInterfacesWwanApn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanApn = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafInterfacesWwanDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesWwanDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-link-detect"]; ok {
		o.LeafInterfacesWwanDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanDisableLinkDetect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesWwanMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["connect-on-demand"]; ok {
		o.LeafInterfacesWwanConnectOnDemand = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanConnectOnDemand = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect"]; ok {
		o.LeafInterfacesWwanRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanRedirect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafInterfacesWwanVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["dhcp-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWwanDhcpOptions = &InterfacesWwanDhcpOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWwanDhcpOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["dhcpv6-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWwanDhcpvsixOptions = &InterfacesWwanDhcpvsixOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWwanDhcpvsixOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["authentication"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWwanAuthentication = &InterfacesWwanAuthentication{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWwanAuthentication)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["mirror"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWwanMirror = &InterfacesWwanMirror{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWwanMirror)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWwanIP = &InterfacesWwanIP{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWwanIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesWwanIPvsix = &InterfacesWwanIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesWwanIPvsix)
		if err != nil {
			return err
		}
	}

	return nil
}
