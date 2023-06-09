// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemConntrackIgnoreRuleSource describes the resource data model.
type SystemConntrackIgnoreRuleSource struct {
	// LeafNodes
	LeafSystemConntrackIgnoreRuleSourceAddress types.String `tfsdk:"address" json:"address,omitempty"`
	LeafSystemConntrackIgnoreRuleSourcePort    types.String `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackIgnoreRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			MarkdownDescription: `Port number

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numeric IP port  |
|  start-end  |  Numbered port range (e.g. 1001-1005)  |
|   |   |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemConntrackIgnoreRuleSource) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafSystemConntrackIgnoreRuleSourceAddress.IsNull() && !o.LeafSystemConntrackIgnoreRuleSourceAddress.IsUnknown() {
		jsonData["address"] = o.LeafSystemConntrackIgnoreRuleSourceAddress.ValueString()
	}

	if !o.LeafSystemConntrackIgnoreRuleSourcePort.IsNull() && !o.LeafSystemConntrackIgnoreRuleSourcePort.IsUnknown() {
		jsonData["port"] = o.LeafSystemConntrackIgnoreRuleSourcePort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemConntrackIgnoreRuleSource) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafSystemConntrackIgnoreRuleSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemConntrackIgnoreRuleSourceAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafSystemConntrackIgnoreRuleSourcePort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemConntrackIgnoreRuleSourcePort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
