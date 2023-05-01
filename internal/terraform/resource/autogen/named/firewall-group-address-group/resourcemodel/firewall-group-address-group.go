// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallGroupAddressGroup describes the resource data model.
type FirewallGroupAddressGroup struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafFirewallGroupAddressGroupAddress     types.String `tfsdk:"address" json:"address,omitempty"`
	LeafFirewallGroupAddressGroupInclude     types.String `tfsdk:"include" json:"include,omitempty"`
	LeafFirewallGroupAddressGroupDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupAddressGroup) GetVyosPath() []string {
	return []string{
		"firewall",
		"group",
		"address-group",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupAddressGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall address-group

`,
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address-group member

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4range  |  IPv4 range to match (e.g. 10.0.0.1-10.0.0.200)  |

`,
		},

		"include": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Include another address-group

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallGroupAddressGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallGroupAddressGroupAddress.IsNull() && !o.LeafFirewallGroupAddressGroupAddress.IsUnknown() {
		jsonData["address"] = o.LeafFirewallGroupAddressGroupAddress.ValueString()
	}

	if !o.LeafFirewallGroupAddressGroupInclude.IsNull() && !o.LeafFirewallGroupAddressGroupInclude.IsUnknown() {
		jsonData["include"] = o.LeafFirewallGroupAddressGroupInclude.ValueString()
	}

	if !o.LeafFirewallGroupAddressGroupDescrIPtion.IsNull() && !o.LeafFirewallGroupAddressGroupDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafFirewallGroupAddressGroupDescrIPtion.ValueString()
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
func (o *FirewallGroupAddressGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafFirewallGroupAddressGroupAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupAddressGroupAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["include"]; ok {
		o.LeafFirewallGroupAddressGroupInclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupAddressGroupInclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafFirewallGroupAddressGroupDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupAddressGroupDescrIPtion = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
