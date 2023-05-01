// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteRuleDestinationGroup describes the resource data model.
type PolicyRouteRuleDestinationGroup struct {
	// LeafNodes
	LeafPolicyRouteRuleDestinationGroupAddressGroup types.String `tfsdk:"address_group" json:"address-group,omitempty"`
	LeafPolicyRouteRuleDestinationGroupDomainGroup  types.String `tfsdk:"domain_group" json:"domain-group,omitempty"`
	LeafPolicyRouteRuleDestinationGroupMacGroup     types.String `tfsdk:"mac_group" json:"mac-group,omitempty"`
	LeafPolicyRouteRuleDestinationGroupNetworkGroup types.String `tfsdk:"network_group" json:"network-group,omitempty"`
	LeafPolicyRouteRuleDestinationGroupPortGroup    types.String `tfsdk:"port_group" json:"port-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleDestinationGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *PolicyRouteRuleDestinationGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteRuleDestinationGroupAddressGroup.IsNull() && !o.LeafPolicyRouteRuleDestinationGroupAddressGroup.IsUnknown() {
		jsonData["address-group"] = o.LeafPolicyRouteRuleDestinationGroupAddressGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleDestinationGroupDomainGroup.IsNull() && !o.LeafPolicyRouteRuleDestinationGroupDomainGroup.IsUnknown() {
		jsonData["domain-group"] = o.LeafPolicyRouteRuleDestinationGroupDomainGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleDestinationGroupMacGroup.IsNull() && !o.LeafPolicyRouteRuleDestinationGroupMacGroup.IsUnknown() {
		jsonData["mac-group"] = o.LeafPolicyRouteRuleDestinationGroupMacGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleDestinationGroupNetworkGroup.IsNull() && !o.LeafPolicyRouteRuleDestinationGroupNetworkGroup.IsUnknown() {
		jsonData["network-group"] = o.LeafPolicyRouteRuleDestinationGroupNetworkGroup.ValueString()
	}

	if !o.LeafPolicyRouteRuleDestinationGroupPortGroup.IsNull() && !o.LeafPolicyRouteRuleDestinationGroupPortGroup.IsUnknown() {
		jsonData["port-group"] = o.LeafPolicyRouteRuleDestinationGroupPortGroup.ValueString()
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
func (o *PolicyRouteRuleDestinationGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address-group"]; ok {
		o.LeafPolicyRouteRuleDestinationGroupAddressGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDestinationGroupAddressGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["domain-group"]; ok {
		o.LeafPolicyRouteRuleDestinationGroupDomainGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDestinationGroupDomainGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-group"]; ok {
		o.LeafPolicyRouteRuleDestinationGroupMacGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDestinationGroupMacGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network-group"]; ok {
		o.LeafPolicyRouteRuleDestinationGroupNetworkGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDestinationGroupNetworkGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port-group"]; ok {
		o.LeafPolicyRouteRuleDestinationGroupPortGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDestinationGroupPortGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
