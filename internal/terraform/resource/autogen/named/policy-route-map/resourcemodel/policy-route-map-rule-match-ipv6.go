// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// PolicyRouteMapRuleMatchIPvsix describes the resource data model.
type PolicyRouteMapRuleMatchIPvsix struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodePolicyRouteMapRuleMatchIPvsixAddress *PolicyRouteMapRuleMatchIPvsixAddress `tfsdk:"address" json:"address,omitempty"`
	NodePolicyRouteMapRuleMatchIPvsixNexthop *PolicyRouteMapRuleMatchIPvsixNexthop `tfsdk:"nexthop" json:"nexthop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address of route to match

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPvsixNexthop{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 next-hop of route to match

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleMatchIPvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePolicyRouteMapRuleMatchIPvsixAddress).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteMapRuleMatchIPvsixAddress)
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

	if !reflect.ValueOf(o.NodePolicyRouteMapRuleMatchIPvsixNexthop).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteMapRuleMatchIPvsixNexthop)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["nexthop"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRouteMapRuleMatchIPvsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["address"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteMapRuleMatchIPvsixAddress = &PolicyRouteMapRuleMatchIPvsixAddress{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteMapRuleMatchIPvsixAddress)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["nexthop"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteMapRuleMatchIPvsixNexthop = &PolicyRouteMapRuleMatchIPvsixNexthop{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteMapRuleMatchIPvsixNexthop)
		if err != nil {
			return err
		}
	}

	return nil
}
