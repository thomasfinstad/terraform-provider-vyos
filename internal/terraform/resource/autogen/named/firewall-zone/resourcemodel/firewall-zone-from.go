// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// FirewallZoneFrom describes the resource data model.
type FirewallZoneFrom struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeFirewallZoneFromFirewall *FirewallZoneFromFirewall `tfsdk:"firewall" json:"firewall,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneFrom) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"firewall": schema.SingleNestedAttribute{
			Attributes: FirewallZoneFromFirewall{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Firewall options

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallZoneFrom) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeFirewallZoneFromFirewall).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallZoneFromFirewall)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["firewall"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallZoneFrom) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["firewall"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallZoneFromFirewall = &FirewallZoneFromFirewall{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallZoneFromFirewall)
		if err != nil {
			return err
		}
	}

	return nil
}
