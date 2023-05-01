// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsStaticRoute describes the resource data model.
type ProtocolsStaticRoute struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsStaticRouteDhcpInterface types.String `tfsdk:"dhcp_interface" json:"dhcp-interface,omitempty"`
	LeafProtocolsStaticRouteDescrIPtion   types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes
	TagProtocolsStaticRouteInterface *map[string]ProtocolsStaticRouteInterface `tfsdk:"interface" json:"interface,omitempty"`
	TagProtocolsStaticRouteNextHop   *map[string]ProtocolsStaticRouteNextHop   `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// Nodes
	NodeProtocolsStaticRouteBlackhole *ProtocolsStaticRouteBlackhole `tfsdk:"blackhole" json:"blackhole,omitempty"`
	NodeProtocolsStaticRouteReject    *ProtocolsStaticRouteReject    `tfsdk:"reject" json:"reject,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticRoute) GetVyosPath() []string {
	return []string{
		"protocols",
		"static",
		"route",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 route

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 static route  |

`,
		},

		// LeafNodes

		"dhcp_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP interface supplying next-hop IP address

|  Format  |  Description  |
|----------|---------------|
|  txt  |  DHCP interface name  |

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

		// TagNodes

		"interface": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsStaticRouteInterface{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Next-hop IPv4 router interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Gateway interface name  |

`,
		},

		"next_hop": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsStaticRouteNextHop{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Next-hop IPv4 router address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Next-hop router address  |

`,
		},

		// Nodes

		"blackhole": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRouteBlackhole{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Silently discard pkts when matched

`,
		},

		"reject": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRouteReject{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Emit an ICMP unreachable when matched

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticRoute) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsStaticRouteDhcpInterface.IsNull() && !o.LeafProtocolsStaticRouteDhcpInterface.IsUnknown() {
		jsonData["dhcp-interface"] = o.LeafProtocolsStaticRouteDhcpInterface.ValueString()
	}

	if !o.LeafProtocolsStaticRouteDescrIPtion.IsNull() && !o.LeafProtocolsStaticRouteDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafProtocolsStaticRouteDescrIPtion.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagProtocolsStaticRouteInterface).IsZero() {
		subJSONStr, err := json.Marshal(o.TagProtocolsStaticRouteInterface)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["interface"] = subData
	}

	if !reflect.ValueOf(o.TagProtocolsStaticRouteNextHop).IsZero() {
		subJSONStr, err := json.Marshal(o.TagProtocolsStaticRouteNextHop)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["next-hop"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsStaticRouteBlackhole).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsStaticRouteBlackhole)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["blackhole"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsStaticRouteReject).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsStaticRouteReject)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["reject"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsStaticRoute) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["dhcp-interface"]; ok {
		o.LeafProtocolsStaticRouteDhcpInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticRouteDhcpInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafProtocolsStaticRouteDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticRouteDescrIPtion = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["interface"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagProtocolsStaticRouteInterface = &map[string]ProtocolsStaticRouteInterface{}

		err = json.Unmarshal(subJSONStr, o.TagProtocolsStaticRouteInterface)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["next-hop"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagProtocolsStaticRouteNextHop = &map[string]ProtocolsStaticRouteNextHop{}

		err = json.Unmarshal(subJSONStr, o.TagProtocolsStaticRouteNextHop)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["blackhole"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsStaticRouteBlackhole = &ProtocolsStaticRouteBlackhole{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsStaticRouteBlackhole)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["reject"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsStaticRouteReject = &ProtocolsStaticRouteReject{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsStaticRouteReject)
		if err != nil {
			return err
		}
	}

	return nil
}
