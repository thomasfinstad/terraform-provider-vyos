// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// PolicyRouteMapRuleSetLthreevpnNexthop describes the resource data model.
type PolicyRouteMapRuleSetLthreevpnNexthop struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodePolicyRouteMapRuleSetLthreevpnNexthopEncapsulation *PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation `tfsdk:"encapsulation" json:"encapsulation,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetLthreevpnNexthop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"encapsulation": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Encapsulation options (for BGP only)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleSetLthreevpnNexthop) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePolicyRouteMapRuleSetLthreevpnNexthopEncapsulation).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteMapRuleSetLthreevpnNexthopEncapsulation)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["encapsulation"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRouteMapRuleSetLthreevpnNexthop) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["encapsulation"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteMapRuleSetLthreevpnNexthopEncapsulation = &PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteMapRuleSetLthreevpnNexthopEncapsulation)
		if err != nil {
			return err
		}
	}

	return nil
}
