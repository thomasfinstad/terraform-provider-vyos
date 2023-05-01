// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallGroupDomainGroup describes the resource data model.
type FirewallGroupDomainGroup struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafFirewallGroupDomainGroupAddress     types.String `tfsdk:"address" json:"address,omitempty"`
	LeafFirewallGroupDomainGroupDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupDomainGroup) GetVyosPath() []string {
	return []string{
		"firewall",
		"group",
		"domain-group",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupDomainGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall domain-group

`,
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Domain-group member

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Domain address to match  |

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
func (o *FirewallGroupDomainGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallGroupDomainGroupAddress.IsNull() && !o.LeafFirewallGroupDomainGroupAddress.IsUnknown() {
		jsonData["address"] = o.LeafFirewallGroupDomainGroupAddress.ValueString()
	}

	if !o.LeafFirewallGroupDomainGroupDescrIPtion.IsNull() && !o.LeafFirewallGroupDomainGroupDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafFirewallGroupDomainGroupDescrIPtion.ValueString()
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
func (o *FirewallGroupDomainGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafFirewallGroupDomainGroupAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupDomainGroupAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafFirewallGroupDomainGroupDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupDomainGroupDescrIPtion = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
