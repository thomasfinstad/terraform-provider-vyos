// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ProtocolsBgpNeighborAddressFamily describes the resource data model.
type ProtocolsBgpNeighborAddressFamily struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsBgpNeighborAddressFamilyIPvfourUnicast        *ProtocolsBgpNeighborAddressFamilyIPvfourUnicast        `tfsdk:"ipv4_unicast" json:"ipv4-unicast,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixUnicast         *ProtocolsBgpNeighborAddressFamilyIPvsixUnicast         `tfsdk:"ipv6_unicast" json:"ipv6-unicast,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast *ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast `tfsdk:"ipv4_labeled_unicast" json:"ipv4-labeled-unicast,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast  *ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast  `tfsdk:"ipv6_labeled_unicast" json:"ipv6-labeled-unicast,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpn            *ProtocolsBgpNeighborAddressFamilyIPvfourVpn            `tfsdk:"ipv4_vpn" json:"ipv4-vpn,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixVpn             *ProtocolsBgpNeighborAddressFamilyIPvsixVpn             `tfsdk:"ipv6_vpn" json:"ipv6-vpn,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourFlowspec       *ProtocolsBgpNeighborAddressFamilyIPvfourFlowspec       `tfsdk:"ipv4_flowspec" json:"ipv4-flowspec,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspec        *ProtocolsBgpNeighborAddressFamilyIPvsixFlowspec        `tfsdk:"ipv6_flowspec" json:"ipv6-flowspec,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourMulticast      *ProtocolsBgpNeighborAddressFamilyIPvfourMulticast      `tfsdk:"ipv4_multicast" json:"ipv4-multicast,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixMulticast       *ProtocolsBgpNeighborAddressFamilyIPvsixMulticast       `tfsdk:"ipv6_multicast" json:"ipv6-multicast,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpn           *ProtocolsBgpNeighborAddressFamilyLtwovpnEvpn           `tfsdk:"l2vpn_evpn" json:"l2vpn-evpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamily) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 BGP neighbor parameters

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 BGP neighbor parameters

`,
		},

		"ipv4_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv6_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv4_vpn": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 VPN BGP neighbor parameters

`,
		},

		"ipv6_vpn": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 VPN BGP neighbor parameters

`,
		},

		"ipv4_flowspec": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourFlowspec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 Flow Specification BGP neighbor parameters

`,
		},

		"ipv6_flowspec": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixFlowspec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 Flow Specification BGP neighbor parameters

`,
		},

		"ipv4_multicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourMulticast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 Multicast BGP neighbor parameters

`,
		},

		"ipv6_multicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 Multicast BGP neighbor parameters

`,
		},

		"l2vpn_evpn": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyLtwovpnEvpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `L2VPN EVPN BGP settings

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamily) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourUnicast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourUnicast)
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

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixUnicast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixUnicast)
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

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv4-labeled-unicast"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6-labeled-unicast"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpn).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpn)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv4-vpn"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixVpn).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixVpn)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6-vpn"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourFlowspec).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourFlowspec)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv4-flowspec"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspec).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspec)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6-flowspec"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourMulticast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourMulticast)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv4-multicast"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixMulticast).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixMulticast)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6-multicast"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpn).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpn)
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
func (o *ProtocolsBgpNeighborAddressFamily) UnmarshalJSON(jsonStr []byte) error {
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

		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourUnicast = &ProtocolsBgpNeighborAddressFamilyIPvfourUnicast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvfourUnicast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6-unicast"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvsixUnicast = &ProtocolsBgpNeighborAddressFamilyIPvsixUnicast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvsixUnicast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv4-labeled-unicast"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast = &ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6-labeled-unicast"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast = &ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv4-vpn"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpn = &ProtocolsBgpNeighborAddressFamilyIPvfourVpn{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpn)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6-vpn"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvsixVpn = &ProtocolsBgpNeighborAddressFamilyIPvsixVpn{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvsixVpn)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv4-flowspec"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourFlowspec = &ProtocolsBgpNeighborAddressFamilyIPvfourFlowspec{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvfourFlowspec)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6-flowspec"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspec = &ProtocolsBgpNeighborAddressFamilyIPvsixFlowspec{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspec)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv4-multicast"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourMulticast = &ProtocolsBgpNeighborAddressFamilyIPvfourMulticast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvfourMulticast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6-multicast"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyIPvsixMulticast = &ProtocolsBgpNeighborAddressFamilyIPvsixMulticast{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyIPvsixMulticast)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["l2vpn-evpn"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpn = &ProtocolsBgpNeighborAddressFamilyLtwovpnEvpn{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpn)
		if err != nil {
			return err
		}
	}

	return nil
}
