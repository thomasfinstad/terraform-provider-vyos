// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyShaperHfscClassMatch describes the resource data model.
type QosPolicyShaperHfscClassMatch struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`
	LeafQosPolicyShaperHfscClassMatchInterface   types.String `tfsdk:"interface" json:"interface,omitempty"`
	LeafQosPolicyShaperHfscClassMatchMark        types.String `tfsdk:"mark" json:"mark,omitempty"`
	LeafQosPolicyShaperHfscClassMatchVif         types.String `tfsdk:"vif" json:"vif,omitempty"`

	// TagNodes

	// Nodes
	NodeQosPolicyShaperHfscClassMatchEther  *QosPolicyShaperHfscClassMatchEther  `tfsdk:"ether" json:"ether,omitempty"`
	NodeQosPolicyShaperHfscClassMatchIP     *QosPolicyShaperHfscClassMatchIP     `tfsdk:"ip" json:"ip,omitempty"`
	NodeQosPolicyShaperHfscClassMatchIPvsix *QosPolicyShaperHfscClassMatchIPvsix `tfsdk:"ipv6" json:"ipv6,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatch) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		"mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on mark applied by firewall

|  Format  |  Description  |
|----------|---------------|
|  u32  |  FW mark to match  |

`,
		},

		"vif": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID for this match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4095  |  Virtual Local Area Network (VLAN) tag   |

`,
		},

		// TagNodes

		// Nodes

		"ether": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchEther{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperHfscClassMatch) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyShaperHfscClassMatchDescrIPtion.IsNull() && !o.LeafQosPolicyShaperHfscClassMatchDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafQosPolicyShaperHfscClassMatchDescrIPtion.ValueString()
	}

	if !o.LeafQosPolicyShaperHfscClassMatchInterface.IsNull() && !o.LeafQosPolicyShaperHfscClassMatchInterface.IsUnknown() {
		jsonData["interface"] = o.LeafQosPolicyShaperHfscClassMatchInterface.ValueString()
	}

	if !o.LeafQosPolicyShaperHfscClassMatchMark.IsNull() && !o.LeafQosPolicyShaperHfscClassMatchMark.IsUnknown() {
		jsonData["mark"] = o.LeafQosPolicyShaperHfscClassMatchMark.ValueString()
	}

	if !o.LeafQosPolicyShaperHfscClassMatchVif.IsNull() && !o.LeafQosPolicyShaperHfscClassMatchVif.IsUnknown() {
		jsonData["vif"] = o.LeafQosPolicyShaperHfscClassMatchVif.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeQosPolicyShaperHfscClassMatchEther).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeQosPolicyShaperHfscClassMatchEther)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ether"] = subData
	}

	if !reflect.ValueOf(o.NodeQosPolicyShaperHfscClassMatchIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeQosPolicyShaperHfscClassMatchIP)
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

	if !reflect.ValueOf(o.NodeQosPolicyShaperHfscClassMatchIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeQosPolicyShaperHfscClassMatchIPvsix)
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

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyShaperHfscClassMatch) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["interface"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mark"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchMark = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vif"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchVif = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchVif = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["ether"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeQosPolicyShaperHfscClassMatchEther = &QosPolicyShaperHfscClassMatchEther{}

		err = json.Unmarshal(subJSONStr, o.NodeQosPolicyShaperHfscClassMatchEther)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeQosPolicyShaperHfscClassMatchIP = &QosPolicyShaperHfscClassMatchIP{}

		err = json.Unmarshal(subJSONStr, o.NodeQosPolicyShaperHfscClassMatchIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeQosPolicyShaperHfscClassMatchIPvsix = &QosPolicyShaperHfscClassMatchIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeQosPolicyShaperHfscClassMatchIPvsix)
		if err != nil {
			return err
		}
	}

	return nil
}
