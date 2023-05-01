// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRoutesixRuleSource describes the resource data model.
type PolicyRoutesixRuleSource struct {
	// LeafNodes
	LeafPolicyRoutesixRuleSourceAddress    types.String `tfsdk:"address" json:"address,omitempty"`
	LeafPolicyRoutesixRuleSourcePort       types.String `tfsdk:"port" json:"port,omitempty"`
	LeafPolicyRoutesixRuleSourceMacAddress types.String `tfsdk:"mac_address" json:"mac-address,omitempty"`

	// TagNodes

	// Nodes
	NodePolicyRoutesixRuleSourceGroup *PolicyRoutesixRuleSourceGroup `tfsdk:"group" json:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IP address to match  |
|  ipv6net  |  Subnet to match  |
|  ipv6range  |  IP range to match  |
|  !ipv6  |  Match everything except the specified address  |
|  !ipv6net  |  Match everything except the specified prefix  |
|  !ipv6range  |  Match everything except the specified range  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numbered port  |
|  <start-end>  |  Numbered port range (e.g. 1001-1005)  |
|     |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |

`,
		},

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MAC address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |
|  !macaddr  |  Match everything except the specified MAC address  |

`,
		},

		// TagNodes

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSourceGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRoutesixRuleSource) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRoutesixRuleSourceAddress.IsNull() && !o.LeafPolicyRoutesixRuleSourceAddress.IsUnknown() {
		jsonData["address"] = o.LeafPolicyRoutesixRuleSourceAddress.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleSourcePort.IsNull() && !o.LeafPolicyRoutesixRuleSourcePort.IsUnknown() {
		jsonData["port"] = o.LeafPolicyRoutesixRuleSourcePort.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleSourceMacAddress.IsNull() && !o.LeafPolicyRoutesixRuleSourceMacAddress.IsUnknown() {
		jsonData["mac-address"] = o.LeafPolicyRoutesixRuleSourceMacAddress.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePolicyRoutesixRuleSourceGroup).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRoutesixRuleSourceGroup)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["group"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRoutesixRuleSource) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafPolicyRoutesixRuleSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSourceAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafPolicyRoutesixRuleSourcePort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSourcePort = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-address"]; ok {
		o.LeafPolicyRoutesixRuleSourceMacAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSourceMacAddress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["group"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRoutesixRuleSourceGroup = &PolicyRoutesixRuleSourceGroup{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRoutesixRuleSourceGroup)
		if err != nil {
			return err
		}
	}

	return nil
}
