// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceRouterAdvertInterface describes the resource data model.
type ServiceRouterAdvertInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServiceRouterAdvertInterfaceHopLimit           types.String `tfsdk:"hop_limit" json:"hop-limit,omitempty"`
	LeafServiceRouterAdvertInterfaceDefaultLifetime    types.String `tfsdk:"default_lifetime" json:"default-lifetime,omitempty"`
	LeafServiceRouterAdvertInterfaceDefaultPreference  types.String `tfsdk:"default_preference" json:"default-preference,omitempty"`
	LeafServiceRouterAdvertInterfaceDNSsl              types.String `tfsdk:"dnssl" json:"dnssl,omitempty"`
	LeafServiceRouterAdvertInterfaceLinkMtu            types.String `tfsdk:"link_mtu" json:"link-mtu,omitempty"`
	LeafServiceRouterAdvertInterfaceManagedFlag        types.String `tfsdk:"managed_flag" json:"managed-flag,omitempty"`
	LeafServiceRouterAdvertInterfaceNameServer         types.String `tfsdk:"name_server" json:"name-server,omitempty"`
	LeafServiceRouterAdvertInterfaceNameServerLifetime types.String `tfsdk:"name_server_lifetime" json:"name-server-lifetime,omitempty"`
	LeafServiceRouterAdvertInterfaceOtherConfigFlag    types.String `tfsdk:"other_config_flag" json:"other-config-flag,omitempty"`
	LeafServiceRouterAdvertInterfaceSourceAddress      types.String `tfsdk:"source_address" json:"source-address,omitempty"`
	LeafServiceRouterAdvertInterfaceReachableTime      types.String `tfsdk:"reachable_time" json:"reachable-time,omitempty"`
	LeafServiceRouterAdvertInterfaceRetransTimer       types.String `tfsdk:"retrans_timer" json:"retrans-timer,omitempty"`
	LeafServiceRouterAdvertInterfaceNoSendAdvert       types.String `tfsdk:"no_send_advert" json:"no-send-advert,omitempty"`

	// TagNodes
	TagServiceRouterAdvertInterfaceRoute  *map[string]ServiceRouterAdvertInterfaceRoute  `tfsdk:"route" json:"route,omitempty"`
	TagServiceRouterAdvertInterfacePrefix *map[string]ServiceRouterAdvertInterfacePrefix `tfsdk:"prefix" json:"prefix,omitempty"`

	// Nodes
	NodeServiceRouterAdvertInterfaceInterval *ServiceRouterAdvertInterfaceInterval `tfsdk:"interval" json:"interval,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceRouterAdvertInterface) GetVyosPath() []string {
	return []string{
		"service",
		"router-advert",
		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceRouterAdvertInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface to send RA on

`,
		},

		// LeafNodes

		"hop_limit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set Hop Count field of the IP header for outgoing packets

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Unspecified (by this router)  |
|  u32:1-255  |  Value should represent current diameter of the Internet  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		"default_lifetime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Lifetime associated with the default router in units of seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:4-9000  |  Router Lifetime in seconds  |
|  0  |  Not a default router  |

`,
		},

		"default_preference": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Preference associated with the default router,

|  Format  |  Description  |
|----------|---------------|
|  low  |  Default router has low preference  |
|  medium  |  Default router has medium preference  |
|  high  |  Default router has high preference  |

`,

			// Default:          stringdefault.StaticString(`medium`),
			Computed: true,
		},

		"dnssl": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DNS search list

`,
		},

		"link_mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Link MTU value placed in RAs, exluded in RAs if unset

|  Format  |  Description  |
|----------|---------------|
|  u32:1280-9000  |  Link MTU value in RAs  |

`,
		},

		"managed_flag": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hosts use the administered (stateful) protocol for address autoconfiguration in addition to any addresses autoconfigured using SLAAC

`,
		},

		"name_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Domain Name Server (DNS) IPv6 address  |

`,
		},

		"name_server_lifetime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum duration how long the RDNSS entries are used

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Name-servers should no longer be used  |
|  u32:1-7200  |  Maximum interval in seconds  |

`,
		},

		"other_config_flag": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hosts use the administered (stateful) protocol for autoconfiguration of other (non-address) information

`,
		},

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use IPv6 address as source address. Useful with VRRP.

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IPv6 address to be advertized (must be configured on interface)  |

`,
		},

		"reachable_time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time, in milliseconds, that a node assumes a neighbor is reachable after having received a reachability confirmation

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Reachable Time unspecified by this router  |
|  u32:1-3600000  |  Reachable Time value in RAs (in milliseconds)  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"retrans_timer": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time in milliseconds between retransmitted Neighbor Solicitation messages

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Time, in milliseconds, between retransmitted Neighbor Solicitation messages  |
|  u32:1-4294967295  |  Minimum interval in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"no_send_advert": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Do not send router adverts

`,
		},

		// TagNodes

		"route": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceRouterAdvertInterfaceRoute{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IPv6 route to be advertised in Router Advertisements (RAs)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 route to be advertized  |

`,
		},

		"prefix": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceRouterAdvertInterfacePrefix{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IPv6 prefix to be advertised in Router Advertisements (RAs)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 prefix to be advertized  |

`,
		},

		// Nodes

		"interval": schema.SingleNestedAttribute{
			Attributes: ServiceRouterAdvertInterfaceInterval{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Set interval between unsolicited multicast RAs

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceRouterAdvertInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceRouterAdvertInterfaceHopLimit.IsNull() && !o.LeafServiceRouterAdvertInterfaceHopLimit.IsUnknown() {
		jsonData["hop-limit"] = o.LeafServiceRouterAdvertInterfaceHopLimit.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceDefaultLifetime.IsNull() && !o.LeafServiceRouterAdvertInterfaceDefaultLifetime.IsUnknown() {
		jsonData["default-lifetime"] = o.LeafServiceRouterAdvertInterfaceDefaultLifetime.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceDefaultPreference.IsNull() && !o.LeafServiceRouterAdvertInterfaceDefaultPreference.IsUnknown() {
		jsonData["default-preference"] = o.LeafServiceRouterAdvertInterfaceDefaultPreference.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceDNSsl.IsNull() && !o.LeafServiceRouterAdvertInterfaceDNSsl.IsUnknown() {
		jsonData["dnssl"] = o.LeafServiceRouterAdvertInterfaceDNSsl.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceLinkMtu.IsNull() && !o.LeafServiceRouterAdvertInterfaceLinkMtu.IsUnknown() {
		jsonData["link-mtu"] = o.LeafServiceRouterAdvertInterfaceLinkMtu.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceManagedFlag.IsNull() && !o.LeafServiceRouterAdvertInterfaceManagedFlag.IsUnknown() {
		jsonData["managed-flag"] = o.LeafServiceRouterAdvertInterfaceManagedFlag.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceNameServer.IsNull() && !o.LeafServiceRouterAdvertInterfaceNameServer.IsUnknown() {
		jsonData["name-server"] = o.LeafServiceRouterAdvertInterfaceNameServer.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceNameServerLifetime.IsNull() && !o.LeafServiceRouterAdvertInterfaceNameServerLifetime.IsUnknown() {
		jsonData["name-server-lifetime"] = o.LeafServiceRouterAdvertInterfaceNameServerLifetime.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceOtherConfigFlag.IsNull() && !o.LeafServiceRouterAdvertInterfaceOtherConfigFlag.IsUnknown() {
		jsonData["other-config-flag"] = o.LeafServiceRouterAdvertInterfaceOtherConfigFlag.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceSourceAddress.IsNull() && !o.LeafServiceRouterAdvertInterfaceSourceAddress.IsUnknown() {
		jsonData["source-address"] = o.LeafServiceRouterAdvertInterfaceSourceAddress.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceReachableTime.IsNull() && !o.LeafServiceRouterAdvertInterfaceReachableTime.IsUnknown() {
		jsonData["reachable-time"] = o.LeafServiceRouterAdvertInterfaceReachableTime.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceRetransTimer.IsNull() && !o.LeafServiceRouterAdvertInterfaceRetransTimer.IsUnknown() {
		jsonData["retrans-timer"] = o.LeafServiceRouterAdvertInterfaceRetransTimer.ValueString()
	}

	if !o.LeafServiceRouterAdvertInterfaceNoSendAdvert.IsNull() && !o.LeafServiceRouterAdvertInterfaceNoSendAdvert.IsUnknown() {
		jsonData["no-send-advert"] = o.LeafServiceRouterAdvertInterfaceNoSendAdvert.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagServiceRouterAdvertInterfaceRoute).IsZero() {
		subJSONStr, err := json.Marshal(o.TagServiceRouterAdvertInterfaceRoute)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["route"] = subData
	}

	if !reflect.ValueOf(o.TagServiceRouterAdvertInterfacePrefix).IsZero() {
		subJSONStr, err := json.Marshal(o.TagServiceRouterAdvertInterfacePrefix)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["prefix"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeServiceRouterAdvertInterfaceInterval).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeServiceRouterAdvertInterfaceInterval)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["interval"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceRouterAdvertInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["hop-limit"]; ok {
		o.LeafServiceRouterAdvertInterfaceHopLimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceHopLimit = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-lifetime"]; ok {
		o.LeafServiceRouterAdvertInterfaceDefaultLifetime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceDefaultLifetime = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-preference"]; ok {
		o.LeafServiceRouterAdvertInterfaceDefaultPreference = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceDefaultPreference = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dnssl"]; ok {
		o.LeafServiceRouterAdvertInterfaceDNSsl = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceDNSsl = basetypes.NewStringNull()
	}

	if value, ok := jsonData["link-mtu"]; ok {
		o.LeafServiceRouterAdvertInterfaceLinkMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceLinkMtu = basetypes.NewStringNull()
	}

	if value, ok := jsonData["managed-flag"]; ok {
		o.LeafServiceRouterAdvertInterfaceManagedFlag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceManagedFlag = basetypes.NewStringNull()
	}

	if value, ok := jsonData["name-server"]; ok {
		o.LeafServiceRouterAdvertInterfaceNameServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceNameServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["name-server-lifetime"]; ok {
		o.LeafServiceRouterAdvertInterfaceNameServerLifetime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceNameServerLifetime = basetypes.NewStringNull()
	}

	if value, ok := jsonData["other-config-flag"]; ok {
		o.LeafServiceRouterAdvertInterfaceOtherConfigFlag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceOtherConfigFlag = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-address"]; ok {
		o.LeafServiceRouterAdvertInterfaceSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceSourceAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["reachable-time"]; ok {
		o.LeafServiceRouterAdvertInterfaceReachableTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceReachableTime = basetypes.NewStringNull()
	}

	if value, ok := jsonData["retrans-timer"]; ok {
		o.LeafServiceRouterAdvertInterfaceRetransTimer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceRetransTimer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-send-advert"]; ok {
		o.LeafServiceRouterAdvertInterfaceNoSendAdvert = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceNoSendAdvert = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["route"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagServiceRouterAdvertInterfaceRoute = &map[string]ServiceRouterAdvertInterfaceRoute{}

		err = json.Unmarshal(subJSONStr, o.TagServiceRouterAdvertInterfaceRoute)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["prefix"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagServiceRouterAdvertInterfacePrefix = &map[string]ServiceRouterAdvertInterfacePrefix{}

		err = json.Unmarshal(subJSONStr, o.TagServiceRouterAdvertInterfacePrefix)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["interval"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeServiceRouterAdvertInterfaceInterval = &ServiceRouterAdvertInterfaceInterval{}

		err = json.Unmarshal(subJSONStr, o.NodeServiceRouterAdvertInterfaceInterval)
		if err != nil {
			return err
		}
	}

	return nil
}
