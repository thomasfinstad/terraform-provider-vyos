// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBondingVifSIPvsix describes the resource data model.
type InterfacesBondingVifSIPvsix struct {
	// LeafNodes
	LeafInterfacesBondingVifSIPvsixAdjustMss              types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesBondingVifSIPvsixDisableForwarding      types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesBondingVifSIPvsixDupAddrDetectTransmits types.String `tfsdk:"dup_addr_detect_transmits" json:"dup-addr-detect-transmits,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesBondingVifSIPvsixAddress *InterfacesBondingVifSIPvsixAddress `tfsdk:"address" json:"address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: InterfacesBondingVifSIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBondingVifSIPvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBondingVifSIPvsixAdjustMss.IsNull() && !o.LeafInterfacesBondingVifSIPvsixAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesBondingVifSIPvsixAdjustMss.ValueString()
	}

	if !o.LeafInterfacesBondingVifSIPvsixDisableForwarding.IsNull() && !o.LeafInterfacesBondingVifSIPvsixDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesBondingVifSIPvsixDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesBondingVifSIPvsixDupAddrDetectTransmits.IsNull() && !o.LeafInterfacesBondingVifSIPvsixDupAddrDetectTransmits.IsUnknown() {
		jsonData["dup-addr-detect-transmits"] = o.LeafInterfacesBondingVifSIPvsixDupAddrDetectTransmits.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesBondingVifSIPvsixAddress).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesBondingVifSIPvsixAddress)
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
func (o *InterfacesBondingVifSIPvsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesBondingVifSIPvsixAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSIPvsixAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesBondingVifSIPvsixDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSIPvsixDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dup-addr-detect-transmits"]; ok {
		o.LeafInterfacesBondingVifSIPvsixDupAddrDetectTransmits = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSIPvsixDupAddrDetectTransmits = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["address"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesBondingVifSIPvsixAddress = &InterfacesBondingVifSIPvsixAddress{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesBondingVifSIPvsixAddress)
		if err != nil {
			return err
		}
	}

	return nil
}
