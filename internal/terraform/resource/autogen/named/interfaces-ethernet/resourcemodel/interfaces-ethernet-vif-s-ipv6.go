// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesEthernetVifSIPvsix describes the resource data model.
type InterfacesEthernetVifSIPvsix struct {
	// LeafNodes
	LeafInterfacesEthernetVifSIPvsixAdjustMss              types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesEthernetVifSIPvsixDisableForwarding      types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesEthernetVifSIPvsixDupAddrDetectTransmits types.String `tfsdk:"dup_addr_detect_transmits" json:"dup-addr-detect-transmits,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesEthernetVifSIPvsixAddress *InterfacesEthernetVifSIPvsixAddress `tfsdk:"address" json:"address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: InterfacesEthernetVifSIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesEthernetVifSIPvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesEthernetVifSIPvsixAdjustMss.IsNull() && !o.LeafInterfacesEthernetVifSIPvsixAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesEthernetVifSIPvsixAdjustMss.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSIPvsixDisableForwarding.IsNull() && !o.LeafInterfacesEthernetVifSIPvsixDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesEthernetVifSIPvsixDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesEthernetVifSIPvsixDupAddrDetectTransmits.IsNull() && !o.LeafInterfacesEthernetVifSIPvsixDupAddrDetectTransmits.IsUnknown() {
		jsonData["dup-addr-detect-transmits"] = o.LeafInterfacesEthernetVifSIPvsixDupAddrDetectTransmits.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesEthernetVifSIPvsixAddress).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesEthernetVifSIPvsixAddress)
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
func (o *InterfacesEthernetVifSIPvsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesEthernetVifSIPvsixAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSIPvsixAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesEthernetVifSIPvsixDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSIPvsixDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dup-addr-detect-transmits"]; ok {
		o.LeafInterfacesEthernetVifSIPvsixDupAddrDetectTransmits = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSIPvsixDupAddrDetectTransmits = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["address"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesEthernetVifSIPvsixAddress = &InterfacesEthernetVifSIPvsixAddress{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesEthernetVifSIPvsixAddress)
		if err != nil {
			return err
		}
	}

	return nil
}
