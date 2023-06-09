// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesEthernetVifSVifCIPvsix describes the resource data model.
type InterfacesEthernetVifSVifCIPvsix struct {
	// LeafNodes
	LeafInterfacesEthernetVifSVifCIPvsixAdjustMss              types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesEthernetVifSVifCIPvsixDisableForwarding      types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesEthernetVifSVifCIPvsixDupAddrDetectTransmits types.String `tfsdk:"dup_addr_detect_transmits" json:"dup-addr-detect-transmits,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesEthernetVifSVifCIPvsixAddress *InterfacesEthernetVifSVifCIPvsixAddress `tfsdk:"address" json:"address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSVifCIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |

`,
		},

		"disable_forwarding": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		"dup_addr_detect_transmits": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable Duplicate Address Dectection (DAD)  |
|  u32:1-n  |  Number of NS messages to send while performing DAD  |

`,
		},

		// TagNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifSVifCIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesEthernetVifSVifCIPvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesEthernetVifSVifCIPvsixAdjustMss.IsNull() && !o.LeafInterfacesEthernetVifSVifCIPvsixAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesEthernetVifSVifCIPvsixAdjustMss.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSVifCIPvsixDisableForwarding.IsNull() && !o.LeafInterfacesEthernetVifSVifCIPvsixDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesEthernetVifSVifCIPvsixDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSVifCIPvsixDupAddrDetectTransmits.IsNull() && !o.LeafInterfacesEthernetVifSVifCIPvsixDupAddrDetectTransmits.IsUnknown() {
		jsonData["dup-addr-detect-transmits"] = o.LeafInterfacesEthernetVifSVifCIPvsixDupAddrDetectTransmits.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesEthernetVifSVifCIPvsixAddress).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesEthernetVifSVifCIPvsixAddress)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["address"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesEthernetVifSVifCIPvsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesEthernetVifSVifCIPvsixAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSVifCIPvsixAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesEthernetVifSVifCIPvsixDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSVifCIPvsixDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dup-addr-detect-transmits"]; ok {
		o.LeafInterfacesEthernetVifSVifCIPvsixDupAddrDetectTransmits = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSVifCIPvsixDupAddrDetectTransmits = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["address"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesEthernetVifSVifCIPvsixAddress = &InterfacesEthernetVifSVifCIPvsixAddress{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesEthernetVifSVifCIPvsixAddress)
		if err != nil {
			return err
		}
	}

	return nil
}
