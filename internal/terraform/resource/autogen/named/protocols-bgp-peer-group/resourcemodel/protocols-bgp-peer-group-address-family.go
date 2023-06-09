// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ProtocolsBgpPeerGroupAddressFamily describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamily struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast `tfsdk:"ipv4_unicast" json:"ipv4-unicast,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast  *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast  `tfsdk:"ipv6_unicast" json:"ipv6-unicast,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn    *ProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn    `tfsdk:"l2vpn_evpn" json:"l2vpn-evpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamily) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 BGP neighbor parameters

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 BGP neighbor parameters

`,
		},

		"l2vpn_evpn": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `L2VPN EVPN BGP settings

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroupAddressFamily) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv4-unicast"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6-unicast"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["l2vpn-evpn"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroupAddressFamily) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["ipv4-unicast"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast = &ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6-unicast"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast = &ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["l2vpn-evpn"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn = &ProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn)
		if err != nil {
			return err
		}
	}

	return nil
}
