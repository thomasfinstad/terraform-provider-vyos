// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// NatStaticRule describes the resource data model.
type NatStaticRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafNatStaticRuleDescrIPtion      types.String `tfsdk:"description" json:"description,omitempty"`
	LeafNatStaticRuleInboundInterface types.String `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`

	// TagNodes

	// Nodes
	NodeNatStaticRuleDestination *NatStaticRuleDestination `tfsdk:"destination" json:"destination,omitempty"`
	NodeNatStaticRuleTranSLAtion *NatStaticRuleTranSLAtion `tfsdk:"translation" json:"translation,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NatStaticRule) GetVyosPath() []string {
	return []string{
		"nat",
		"static",
		"rule",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatStaticRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule number for NAT

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound interface of NAT traffic

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `NAT destination parameters

`,
		},

		"translation": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleTranSLAtion{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Translation address or prefix

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *NatStaticRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafNatStaticRuleDescrIPtion.IsNull() && !o.LeafNatStaticRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafNatStaticRuleDescrIPtion.ValueString()
	}

	if !o.LeafNatStaticRuleInboundInterface.IsNull() && !o.LeafNatStaticRuleInboundInterface.IsUnknown() {
		jsonData["inbound-interface"] = o.LeafNatStaticRuleInboundInterface.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeNatStaticRuleDestination).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeNatStaticRuleDestination)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["destination"] = subData
	}

	if !reflect.ValueOf(o.NodeNatStaticRuleTranSLAtion).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeNatStaticRuleTranSLAtion)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["translation"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *NatStaticRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafNatStaticRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatStaticRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["inbound-interface"]; ok {
		o.LeafNatStaticRuleInboundInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatStaticRuleInboundInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["destination"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeNatStaticRuleDestination = &NatStaticRuleDestination{}

		err = json.Unmarshal(subJSONStr, o.NodeNatStaticRuleDestination)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["translation"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeNatStaticRuleTranSLAtion = &NatStaticRuleTranSLAtion{}

		err = json.Unmarshal(subJSONStr, o.NodeNatStaticRuleTranSLAtion)
		if err != nil {
			return err
		}
	}

	return nil
}
