// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyLocalRoutesixRule describes the resource data model.
type PolicyLocalRoutesixRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafPolicyLocalRoutesixRuleFwmark           types.String `tfsdk:"fwmark" json:"fwmark,omitempty"`
	LeafPolicyLocalRoutesixRuleSource           types.String `tfsdk:"source" json:"source,omitempty"`
	LeafPolicyLocalRoutesixRuleDestination      types.String `tfsdk:"destination" json:"destination,omitempty"`
	LeafPolicyLocalRoutesixRuleInboundInterface types.String `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`

	// TagNodes

	// Nodes
	NodePolicyLocalRoutesixRuleSet *PolicyLocalRoutesixRuleSet `tfsdk:"set" json:"set,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyLocalRoutesixRule) GetVyosPath() []string {
	return []string{
		"policy",
		"local-route6",
		"rule",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 policy local-route rule set number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-32765  |  Local-route rule number (1-32765)  |

`,
		},

		// LeafNodes

		"fwmark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Address to match against  |

`,
		},

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Address to match against  |
|  ipv6net  |  Prefix to match against  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Address to match against  |
|  ipv6net  |  Prefix to match against  |

`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound Interface

`,
		},

		// TagNodes

		// Nodes

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRoutesixRuleSet{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyLocalRoutesixRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyLocalRoutesixRuleFwmark.IsNull() && !o.LeafPolicyLocalRoutesixRuleFwmark.IsUnknown() {
		jsonData["fwmark"] = o.LeafPolicyLocalRoutesixRuleFwmark.ValueString()
	}

	if !o.LeafPolicyLocalRoutesixRuleSource.IsNull() && !o.LeafPolicyLocalRoutesixRuleSource.IsUnknown() {
		jsonData["source"] = o.LeafPolicyLocalRoutesixRuleSource.ValueString()
	}

	if !o.LeafPolicyLocalRoutesixRuleDestination.IsNull() && !o.LeafPolicyLocalRoutesixRuleDestination.IsUnknown() {
		jsonData["destination"] = o.LeafPolicyLocalRoutesixRuleDestination.ValueString()
	}

	if !o.LeafPolicyLocalRoutesixRuleInboundInterface.IsNull() && !o.LeafPolicyLocalRoutesixRuleInboundInterface.IsUnknown() {
		jsonData["inbound-interface"] = o.LeafPolicyLocalRoutesixRuleInboundInterface.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePolicyLocalRoutesixRuleSet).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyLocalRoutesixRuleSet)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["set"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyLocalRoutesixRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["fwmark"]; ok {
		o.LeafPolicyLocalRoutesixRuleFwmark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleFwmark = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source"]; ok {
		o.LeafPolicyLocalRoutesixRuleSource = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleSource = basetypes.NewStringNull()
	}

	if value, ok := jsonData["destination"]; ok {
		o.LeafPolicyLocalRoutesixRuleDestination = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleDestination = basetypes.NewStringNull()
	}

	if value, ok := jsonData["inbound-interface"]; ok {
		o.LeafPolicyLocalRoutesixRuleInboundInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleInboundInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["set"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyLocalRoutesixRuleSet = &PolicyLocalRoutesixRuleSet{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyLocalRoutesixRuleSet)
		if err != nil {
			return err
		}
	}

	return nil
}
