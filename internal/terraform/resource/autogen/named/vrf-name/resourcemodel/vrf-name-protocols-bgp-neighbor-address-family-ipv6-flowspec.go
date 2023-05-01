// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspec describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspec struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteReflectorClient types.String `tfsdk:"route_reflector_client" json:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteServerClient    types.String `tfsdk:"route_server_client" json:"route-server-client,omitempty"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList          *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList          `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList          *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList          `tfsdk:"filter_list" json:"filter-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap            *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap            `tfsdk:"route_map" json:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration `tfsdk:"soft_reconfiguration" json:"soft-reconfiguration,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_reflector_client": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
		},

		"route_server_client": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
		},

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspec) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteReflectorClient.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteReflectorClient.IsUnknown() {
		jsonData["route-reflector-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteReflectorClient.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteServerClient.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteServerClient.IsUnknown() {
		jsonData["route-server-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteServerClient.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["prefix-list"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["filter-list"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["route-map"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["soft-reconfiguration"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspec) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["route-reflector-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteReflectorClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteReflectorClient = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-server-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteServerClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteServerClient = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["prefix-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["filter-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["route-map"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["soft-reconfiguration"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration)
		if err != nil {
			return err
		}
	}

	return nil
}
