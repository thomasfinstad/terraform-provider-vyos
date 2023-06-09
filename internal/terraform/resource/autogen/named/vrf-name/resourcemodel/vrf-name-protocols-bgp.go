// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgp describes the resource data model.
type VrfNameProtocolsBgp struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpSystemAs types.String `tfsdk:"system_as" json:"system-as,omitempty"`
	LeafVrfNameProtocolsBgpRouteMap types.String `tfsdk:"route_map" json:"route-map,omitempty"`

	// TagNodes
	TagVrfNameProtocolsBgpNeighbor  *map[string]VrfNameProtocolsBgpNeighbor  `tfsdk:"neighbor" json:"neighbor,omitempty"`
	TagVrfNameProtocolsBgpPeerGroup *map[string]VrfNameProtocolsBgpPeerGroup `tfsdk:"peer_group" json:"peer-group,omitempty"`

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamily *VrfNameProtocolsBgpAddressFamily `tfsdk:"address_family" json:"address-family,omitempty"`
	NodeVrfNameProtocolsBgpListen        *VrfNameProtocolsBgpListen        `tfsdk:"listen" json:"listen,omitempty"`
	NodeVrfNameProtocolsBgpParameters    *VrfNameProtocolsBgpParameters    `tfsdk:"parameters" json:"parameters,omitempty"`
	NodeVrfNameProtocolsBgpTimers        *VrfNameProtocolsBgpTimers        `tfsdk:"timers" json:"timers,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"system_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Autonomous System Number (ASN)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967294  |  Autonomous System Number  |

`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		"neighbor": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpNeighbor{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `BGP neighbor

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  BGP neighbor IP address  |
|  ipv6  |  BGP neighbor IPv6 address  |
|  txt  |  Interface name  |

`,
		},

		"peer_group": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpPeerGroup{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Name of peer-group

`,
		},

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamily{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP address-family parameters

`,
		},

		"listen": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpListen{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Listen for and accept BGP dynamic neighbors from range

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParameters{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP parameters

`,
		},

		"timers": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpTimers{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP protocol timers

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgp) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpSystemAs.IsNull() && !o.LeafVrfNameProtocolsBgpSystemAs.IsUnknown() {
		jsonData["system-as"] = o.LeafVrfNameProtocolsBgpSystemAs.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpRouteMap.IsNull() && !o.LeafVrfNameProtocolsBgpRouteMap.IsUnknown() {
		jsonData["route-map"] = o.LeafVrfNameProtocolsBgpRouteMap.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagVrfNameProtocolsBgpNeighbor).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsBgpNeighbor)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["neighbor"] = subData
	}

	if !reflect.ValueOf(o.TagVrfNameProtocolsBgpPeerGroup).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsBgpPeerGroup)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["peer-group"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpAddressFamily).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpAddressFamily)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["address-family"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpListen).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpListen)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["listen"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpParameters).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpParameters)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["parameters"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpTimers).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpTimers)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["timers"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgp) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["system-as"]; ok {
		o.LeafVrfNameProtocolsBgpSystemAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpSystemAs = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-map"]; ok {
		o.LeafVrfNameProtocolsBgpRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpRouteMap = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["neighbor"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsBgpNeighbor = &map[string]VrfNameProtocolsBgpNeighbor{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsBgpNeighbor)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["peer-group"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsBgpPeerGroup = &map[string]VrfNameProtocolsBgpPeerGroup{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsBgpPeerGroup)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["address-family"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpAddressFamily = &VrfNameProtocolsBgpAddressFamily{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpAddressFamily)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["listen"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpListen = &VrfNameProtocolsBgpListen{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpListen)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["parameters"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpParameters = &VrfNameProtocolsBgpParameters{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpParameters)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["timers"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpTimers = &VrfNameProtocolsBgpTimers{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpTimers)
		if err != nil {
			return err
		}
	}

	return nil
}
