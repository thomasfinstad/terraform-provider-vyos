// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBabelDistributeListIPvsixInterface describes the resource data model.
type ProtocolsBabelDistributeListIPvsixInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList *ProtocolsBabelDistributeListIPvsixInterfaceAccessList `tfsdk:"access_list" json:"access-list,omitempty"`
	NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList *ProtocolsBabelDistributeListIPvsixInterfacePrefixList `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelDistributeListIPvsixInterface) GetVyosPath() []string {
	return []string{
		"protocols",
		"babel",
		"distribute-list",
		"ipv6",
		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelDistributeListIPvsixInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Apply filtering to an interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Apply filtering to an interface  |

`,
		},

		// LeafNodes

		// TagNodes

		// Nodes

		"access_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBabelDistributeListIPvsixInterfaceAccessList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBabelDistributeListIPvsixInterfacePrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBabelDistributeListIPvsixInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["access-list"] = subData
	}

	if !reflect.ValueOf(o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList)
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

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBabelDistributeListIPvsixInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["access-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList = &ProtocolsBabelDistributeListIPvsixInterfaceAccessList{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["prefix-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList = &ProtocolsBabelDistributeListIPvsixInterfacePrefixList{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList)
		if err != nil {
			return err
		}
	}

	return nil
}
