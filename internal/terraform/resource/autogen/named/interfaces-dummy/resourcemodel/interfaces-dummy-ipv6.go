// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesDummyIPvsix describes the resource data model.
type InterfacesDummyIPvsix struct {
	// LeafNodes
	LeafInterfacesDummyIPvsixDisableForwarding types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`

	// TagNodes

	// Nodes
	NodeInterfacesDummyIPvsixAddress *InterfacesDummyIPvsixAddress `tfsdk:"address" json:"address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesDummyIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable_forwarding": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		// TagNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: InterfacesDummyIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesDummyIPvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesDummyIPvsixDisableForwarding.IsNull() && !o.LeafInterfacesDummyIPvsixDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesDummyIPvsixDisableForwarding.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeInterfacesDummyIPvsixAddress).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeInterfacesDummyIPvsixAddress)
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
func (o *InterfacesDummyIPvsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesDummyIPvsixDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesDummyIPvsixDisableForwarding = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["address"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeInterfacesDummyIPvsixAddress = &InterfacesDummyIPvsixAddress{}

		err = json.Unmarshal(subJSONStr, o.NodeInterfacesDummyIPvsixAddress)
		if err != nil {
			return err
		}
	}

	return nil
}
