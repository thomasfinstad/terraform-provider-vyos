// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallGroupIPvsixAddressGroup describes the resource data model.
type FirewallGroupIPvsixAddressGroup struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafFirewallGroupIPvsixAddressGroupAddress     types.String `tfsdk:"address" json:"address,omitempty"`
	LeafFirewallGroupIPvsixAddressGroupInclude     types.String `tfsdk:"include" json:"include,omitempty"`
	LeafFirewallGroupIPvsixAddressGroupDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupIPvsixAddressGroup) GetVyosPath() []string {
	return []string{
		"firewall",
		"group",
		"ipv6-address-group",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupIPvsixAddressGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall ipv6-address-group

`,
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address-group member

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IPv6 address to match  |
|  ipv6range  |  IPv6 range to match (e.g. 2002::1-2002::ff)  |

`,
		},

		"include": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Include another ipv6-address-group

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
func (o *FirewallGroupIPvsixAddressGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallGroupIPvsixAddressGroupAddress.IsNull() && !o.LeafFirewallGroupIPvsixAddressGroupAddress.IsUnknown() {
		jsonData["address"] = o.LeafFirewallGroupIPvsixAddressGroupAddress.ValueString()
	}

	if !o.LeafFirewallGroupIPvsixAddressGroupInclude.IsNull() && !o.LeafFirewallGroupIPvsixAddressGroupInclude.IsUnknown() {
		jsonData["include"] = o.LeafFirewallGroupIPvsixAddressGroupInclude.ValueString()
	}

	if !o.LeafFirewallGroupIPvsixAddressGroupDescrIPtion.IsNull() && !o.LeafFirewallGroupIPvsixAddressGroupDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafFirewallGroupIPvsixAddressGroupDescrIPtion.ValueString()
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
func (o *FirewallGroupIPvsixAddressGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafFirewallGroupIPvsixAddressGroupAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupIPvsixAddressGroupAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["include"]; ok {
		o.LeafFirewallGroupIPvsixAddressGroupInclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupIPvsixAddressGroupInclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafFirewallGroupIPvsixAddressGroupDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupIPvsixAddressGroupDescrIPtion = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
