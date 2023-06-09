// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteRuleSourceGroup describes the resource data model.
type PolicyRouteRuleSourceGroup struct {
	// LeafNodes
	LeafPolicyRouteRuleSourceGroupAddressGroup types.String `tfsdk:"address_group" json:"address-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupDomainGroup  types.String `tfsdk:"domain_group" json:"domain-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupMacGroup     types.String `tfsdk:"mac_group" json:"mac-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupNetworkGroup types.String `tfsdk:"network_group" json:"network-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupPortGroup    types.String `tfsdk:"port_group" json:"port-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleSourceGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of addresses

`,
		},

		"domain_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of domains

`,
		},

		"mac_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of MAC addresses

`,
		},

		"network_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of networks

`,
		},

		"port_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of ports

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteRuleSourceGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteRuleSourceGroupAddressGroup.IsNull() && !o.LeafPolicyRouteRuleSourceGroupAddressGroup.IsUnknown() {
		jsonData["address-group"] = o.LeafPolicyRouteRuleSourceGroupAddressGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleSourceGroupDomainGroup.IsNull() && !o.LeafPolicyRouteRuleSourceGroupDomainGroup.IsUnknown() {
		jsonData["domain-group"] = o.LeafPolicyRouteRuleSourceGroupDomainGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleSourceGroupMacGroup.IsNull() && !o.LeafPolicyRouteRuleSourceGroupMacGroup.IsUnknown() {
		jsonData["mac-group"] = o.LeafPolicyRouteRuleSourceGroupMacGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleSourceGroupNetworkGroup.IsNull() && !o.LeafPolicyRouteRuleSourceGroupNetworkGroup.IsUnknown() {
		jsonData["network-group"] = o.LeafPolicyRouteRuleSourceGroupNetworkGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleSourceGroupPortGroup.IsNull() && !o.LeafPolicyRouteRuleSourceGroupPortGroup.IsUnknown() {
		jsonData["port-group"] = o.LeafPolicyRouteRuleSourceGroupPortGroup.ValueString()
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
func (o *PolicyRouteRuleSourceGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address-group"]; ok {
		o.LeafPolicyRouteRuleSourceGroupAddressGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourceGroupAddressGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["domain-group"]; ok {
		o.LeafPolicyRouteRuleSourceGroupDomainGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourceGroupDomainGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-group"]; ok {
		o.LeafPolicyRouteRuleSourceGroupMacGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourceGroupMacGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network-group"]; ok {
		o.LeafPolicyRouteRuleSourceGroupNetworkGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourceGroupNetworkGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port-group"]; ok {
		o.LeafPolicyRouteRuleSourceGroupPortGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSourceGroupPortGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
