// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBrIDge describes the resource data model.
type InterfacesBrIDge struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafInterfacesBrIDgeAddress           types.String `tfsdk:"address" json:"address,omitempty"`
	LeafInterfacesBrIDgeAging             types.String `tfsdk:"aging" json:"aging,omitempty"`
	LeafInterfacesBrIDgeDescrIPtion       types.String `tfsdk:"description" json:"description,omitempty"`
	LeafInterfacesBrIDgeDisableLinkDetect types.String `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	LeafInterfacesBrIDgeDisable           types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafInterfacesBrIDgeVrf               types.String `tfsdk:"vrf" json:"vrf,omitempty"`
	LeafInterfacesBrIDgeMtu               types.String `tfsdk:"mtu" json:"mtu,omitempty"`
	LeafInterfacesBrIDgeForwardingDelay   types.String `tfsdk:"forwarding_delay" json:"forwarding-delay,omitempty"`
	LeafInterfacesBrIDgeHelloTime         types.String `tfsdk:"hello_time" json:"hello-time,omitempty"`
	LeafInterfacesBrIDgeMac               types.String `tfsdk:"mac" json:"mac,omitempty"`
	LeafInterfacesBrIDgeEnableVlan        types.String `tfsdk:"enable_vlan" json:"enable-vlan,omitempty"`
	LeafInterfacesBrIDgeMaxAge            types.String `tfsdk:"max_age" json:"max-age,omitempty"`
	LeafInterfacesBrIDgePriority          types.String `tfsdk:"priority" json:"priority,omitempty"`
	LeafInterfacesBrIDgeStp               types.String `tfsdk:"stp" json:"stp,omitempty"`
	LeafInterfacesBrIDgeRedirect          types.String `tfsdk:"redirect" json:"redirect,omitempty"`

	// TagNodes
	TagInterfacesBrIDgeVif *map[string]InterfacesBrIDgeVif `tfsdk:"vif" json:"vif,omitempty"`

	// Nodes
	NodeInterfacesBrIDgeDhcpOptions     *InterfacesBrIDgeDhcpOptions     `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	NodeInterfacesBrIDgeDhcpvsixOptions *InterfacesBrIDgeDhcpvsixOptions `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	NodeInterfacesBrIDgeIgmp            *InterfacesBrIDgeIgmp            `tfsdk:"igmp" json:"igmp,omitempty"`
	NodeInterfacesBrIDgeIP              *InterfacesBrIDgeIP              `tfsdk:"ip" json:"ip,omitempty"`
	NodeInterfacesBrIDgeIPvsix          *InterfacesBrIDgeIPvsix          `tfsdk:"ipv6" json:"ipv6,omitempty"`
	NodeInterfacesBrIDgeMirror          *InterfacesBrIDgeMirror          `tfsdk:"mirror" json:"mirror,omitempty"`
	NodeInterfacesBrIDgeMember          *InterfacesBrIDgeMember          `tfsdk:"member" json:"member,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBrIDge) GetVyosPath() []string {
	return []string{
		"interfaces",
		"bridge",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDge) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bridge Interface

|  Format  |  Description  |
|----------|---------------|
|  brN  |  Bridge interface name  |

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

		"aging": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MAC address aging interval

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable MAC address learning (always flood)  |
|  u32:10-1000000  |  MAC address aging time in seconds  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
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

		"forwarding_delay": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Forwarding delay

|  Format  |  Description  |
|----------|---------------|
|  u32:0-200  |  Spanning Tree Protocol forwarding delay in seconds  |

`,

			// Default:          stringdefault.StaticString(`14`),
			Computed: true,
		},

		"hello_time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hello packet advertisement interval

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Spanning Tree Protocol hello advertisement interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`2`),
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

		"enable_vlan": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable VLAN aware bridge

`,
		},

		"max_age": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval at which neighbor bridges are removed

|  Format  |  Description  |
|----------|---------------|
|  u32:1-40  |  Bridge maximum aging time in seconds  |

`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
		},

		"priority": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Priority for this bridge

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Bridge priority  |

`,

			// Default:          stringdefault.StaticString(`32768`),
			Computed: true,
		},

		"stp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable spanning tree protocol

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

		// TagNodes

		"vif": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesBrIDgeVif{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4094  |  Virtual Local Area Network (VLAN) ID  |

`,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"igmp": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIgmp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Internet Group Management Protocol (IGMP) and Multicast Listener Discovery (MLD) settings

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"member": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeMember{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Bridge member interfaces

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBrIDge) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBrIDgeAddress.IsNull() && !o.LeafInterfacesBrIDgeAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesBrIDgeAddress.ValueString()
	}

	if !o.LeafInterfacesBrIDgeAging.IsNull() && !o.LeafInterfacesBrIDgeAging.IsUnknown() {
		jsonData["aging"] = o.LeafInterfacesBrIDgeAging.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDescrIPtion.IsNull() && !o.LeafInterfacesBrIDgeDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafInterfacesBrIDgeDescrIPtion.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDisableLinkDetect.IsNull() && !o.LeafInterfacesBrIDgeDisableLinkDetect.IsUnknown() {
		jsonData["disable-link-detect"] = o.LeafInterfacesBrIDgeDisableLinkDetect.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDisable.IsNull() && !o.LeafInterfacesBrIDgeDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesBrIDgeDisable.ValueString()
	}

	if !o.LeafInterfacesBrIDgeVrf.IsNull() && !o.LeafInterfacesBrIDgeVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafInterfacesBrIDgeVrf.ValueString()
	}

	if !o.LeafInterfacesBrIDgeMtu.IsNull() && !o.LeafInterfacesBrIDgeMtu.IsUnknown() {
		jsonData["mtu"] = o.LeafInterfacesBrIDgeMtu.ValueString()
	}

	if !o.LeafInterfacesBrIDgeForwardingDelay.IsNull() && !o.LeafInterfacesBrIDgeForwardingDelay.IsUnknown() {
		jsonData["forwarding-delay"] = o.LeafInterfacesBrIDgeForwardingDelay.ValueString()
	}

	if !o.LeafInterfacesBrIDgeHelloTime.IsNull() && !o.LeafInterfacesBrIDgeHelloTime.IsUnknown() {
		jsonData["hello-time"] = o.LeafInterfacesBrIDgeHelloTime.ValueString()
	}

	if !o.LeafInterfacesBrIDgeMac.IsNull() && !o.LeafInterfacesBrIDgeMac.IsUnknown() {
		jsonData["mac"] = o.LeafInterfacesBrIDgeMac.ValueString()
	}

	if !o.LeafInterfacesBrIDgeEnableVlan.IsNull() && !o.LeafInterfacesBrIDgeEnableVlan.IsUnknown() {
		jsonData["enable-vlan"] = o.LeafInterfacesBrIDgeEnableVlan.ValueString()
	}

	if !o.LeafInterfacesBrIDgeMaxAge.IsNull() && !o.LeafInterfacesBrIDgeMaxAge.IsUnknown() {
		jsonData["max-age"] = o.LeafInterfacesBrIDgeMaxAge.ValueString()
	}

	if !o.LeafInterfacesBrIDgePriority.IsNull() && !o.LeafInterfacesBrIDgePriority.IsUnknown() {
		jsonData["priority"] = o.LeafInterfacesBrIDgePriority.ValueString()
	}

	if !o.LeafInterfacesBrIDgeStp.IsNull() && !o.LeafInterfacesBrIDgeStp.IsUnknown() {
		jsonData["stp"] = o.LeafInterfacesBrIDgeStp.ValueString()
	}

	if !o.LeafInterfacesBrIDgeRedirect.IsNull() && !o.LeafInterfacesBrIDgeRedirect.IsUnknown() {
		jsonData["redirect"] = o.LeafInterfacesBrIDgeRedirect.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagInterfacesBrIDgeVif).IsZero() {
		subJSONStr, err := json.Marshal(o.TagInterfacesBrIDgeVif)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["vif"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesBrIDgeDhcpOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBrIDgeDhcpOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesBrIDgeDhcpvsixOptions).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBrIDgeDhcpvsixOptions)
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

	if !reflect.ValueOf(o.NodeInterfacesBrIDgeIgmp).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBrIDgeIgmp)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["igmp"] = subData
	}

	if !reflect.ValueOf(o.NodeInterfacesBrIDgeIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBrIDgeIP)
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

	if !reflect.ValueOf(o.NodeInterfacesBrIDgeIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBrIDgeIPvsix)
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

	if !reflect.ValueOf(o.NodeInterfacesBrIDgeMirror).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBrIDgeMirror)
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

	if !reflect.ValueOf(o.NodeInterfacesBrIDgeMember).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBrIDgeMember)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["member"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBrIDge) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesBrIDgeAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["aging"]; ok {
		o.LeafInterfacesBrIDgeAging = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeAging = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafInterfacesBrIDgeDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-link-detect"]; ok {
		o.LeafInterfacesBrIDgeDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDisableLinkDetect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesBrIDgeDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafInterfacesBrIDgeVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVrf = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mtu"]; ok {
		o.LeafInterfacesBrIDgeMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["forwarding-delay"]; ok {
		o.LeafInterfacesBrIDgeForwardingDelay = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeForwardingDelay = basetypes.NewStringNull()
	}

	if value, ok := jsonData["hello-time"]; ok {
		o.LeafInterfacesBrIDgeHelloTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeHelloTime = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac"]; ok {
		o.LeafInterfacesBrIDgeMac = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeMac = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-vlan"]; ok {
		o.LeafInterfacesBrIDgeEnableVlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeEnableVlan = basetypes.NewStringNull()
	}

	if value, ok := jsonData["max-age"]; ok {
		o.LeafInterfacesBrIDgeMaxAge = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeMaxAge = basetypes.NewStringNull()
	}

	if value, ok := jsonData["priority"]; ok {
		o.LeafInterfacesBrIDgePriority = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgePriority = basetypes.NewStringNull()
	}

	if value, ok := jsonData["stp"]; ok {
		o.LeafInterfacesBrIDgeStp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeStp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["redirect"]; ok {
		o.LeafInterfacesBrIDgeRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeRedirect = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["vif"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagInterfacesBrIDgeVif = &map[string]InterfacesBrIDgeVif{}

		err = json.Unmarshal(subJSONStr, o.TagInterfacesBrIDgeVif)
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

		o.NodeInterfacesBrIDgeDhcpOptions = &InterfacesBrIDgeDhcpOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBrIDgeDhcpOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["dhcpv6-options"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBrIDgeDhcpvsixOptions = &InterfacesBrIDgeDhcpvsixOptions{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBrIDgeDhcpvsixOptions)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["igmp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBrIDgeIgmp = &InterfacesBrIDgeIgmp{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBrIDgeIgmp)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBrIDgeIP = &InterfacesBrIDgeIP{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBrIDgeIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBrIDgeIPvsix = &InterfacesBrIDgeIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBrIDgeIPvsix)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["mirror"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBrIDgeMirror = &InterfacesBrIDgeMirror{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBrIDgeMirror)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["member"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBrIDgeMember = &InterfacesBrIDgeMember{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBrIDgeMember)
		if err != nil {
			return err
		}
	}

	return nil
}
