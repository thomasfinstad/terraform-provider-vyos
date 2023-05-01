// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteRuleSource describes the resource data model.
type PolicyRouteRuleSource struct {
	// LeafNodes
	LeafPolicyRouteRuleSourceAddress    types.String `tfsdk:"address" json:"address,omitempty"`
	LeafPolicyRouteRuleSourcePort       types.String `tfsdk:"port" json:"port,omitempty"`
	LeafPolicyRouteRuleSourceMacAddress types.String `tfsdk:"mac_address" json:"mac-address,omitempty"`

	// TagNodes

	// Nodes
	NodePolicyRouteRuleSourceGroup *PolicyRouteRuleSourceGroup `tfsdk:"group" json:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
|  ipv4range  |  IPv4 address range to match  |
|  !ipv4  |  Match everything except the specified address  |
|  !ipv4net  |  Match everything except the specified prefix  |
|  !ipv4range  |  Match everything except the specified range  |

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
			Attributes: PolicyRouteRuleSourceGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteRuleSource) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteRuleSourceAddress.IsNull() && !o.LeafPolicyRouteRuleSourceAddress.IsUnknown() {
		jsonData["address"] = o.LeafPolicyRouteRuleSourceAddress.ValueString()
	}

	if !o.LeafPolicyRouteRuleSourcePort.IsNull() && !o.LeafPolicyRouteRuleSourcePort.IsUnknown() {
		jsonData["port"] = o.LeafPolicyRouteRuleSourcePort.ValueString()
	}

	if !o.LeafPolicyRouteRuleSourceMacAddress.IsNull() && !o.LeafPolicyRouteRuleSourceMacAddress.IsUnknown() {
		jsonData["mac-address"] = o.LeafPolicyRouteRuleSourceMacAddress.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePolicyRouteRuleSourceGroup).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleSourceGroup)
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
func (o *PolicyRouteRuleSource) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafPolicyRouteRuleSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourceAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafPolicyRouteRuleSourcePort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourcePort = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-address"]; ok {
		o.LeafPolicyRouteRuleSourceMacAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourceMacAddress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["group"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleSourceGroup = &PolicyRouteRuleSourceGroup{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleSourceGroup)
		if err != nil {
			return err
		}
	}

	return nil
}
