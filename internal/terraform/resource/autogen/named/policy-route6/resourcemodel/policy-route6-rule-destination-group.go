// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRoutesixRuleDestinationGroup describes the resource data model.
type PolicyRoutesixRuleDestinationGroup struct {
	// LeafNodes
	LeafPolicyRoutesixRuleDestinationGroupAddressGroup types.String `tfsdk:"address_group" json:"address-group,omitempty"`
	LeafPolicyRoutesixRuleDestinationGroupDomainGroup  types.String `tfsdk:"domain_group" json:"domain-group,omitempty"`
	LeafPolicyRoutesixRuleDestinationGroupMacGroup     types.String `tfsdk:"mac_group" json:"mac-group,omitempty"`
	LeafPolicyRoutesixRuleDestinationGroupNetworkGroup types.String `tfsdk:"network_group" json:"network-group,omitempty"`
	LeafPolicyRoutesixRuleDestinationGroupPortGroup    types.String `tfsdk:"port_group" json:"port-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleDestinationGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *PolicyRoutesixRuleDestinationGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRoutesixRuleDestinationGroupAddressGroup.IsNull() && !o.LeafPolicyRoutesixRuleDestinationGroupAddressGroup.IsUnknown() {
		jsonData["address-group"] = o.LeafPolicyRoutesixRuleDestinationGroupAddressGroup.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleDestinationGroupDomainGroup.IsNull() && !o.LeafPolicyRoutesixRuleDestinationGroupDomainGroup.IsUnknown() {
		jsonData["domain-group"] = o.LeafPolicyRoutesixRuleDestinationGroupDomainGroup.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleDestinationGroupMacGroup.IsNull() && !o.LeafPolicyRoutesixRuleDestinationGroupMacGroup.IsUnknown() {
		jsonData["mac-group"] = o.LeafPolicyRoutesixRuleDestinationGroupMacGroup.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleDestinationGroupNetworkGroup.IsNull() && !o.LeafPolicyRoutesixRuleDestinationGroupNetworkGroup.IsUnknown() {
		jsonData["network-group"] = o.LeafPolicyRoutesixRuleDestinationGroupNetworkGroup.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleDestinationGroupPortGroup.IsNull() && !o.LeafPolicyRoutesixRuleDestinationGroupPortGroup.IsUnknown() {
		jsonData["port-group"] = o.LeafPolicyRoutesixRuleDestinationGroupPortGroup.ValueString()
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
func (o *PolicyRoutesixRuleDestinationGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address-group"]; ok {
		o.LeafPolicyRoutesixRuleDestinationGroupAddressGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDestinationGroupAddressGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["domain-group"]; ok {
		o.LeafPolicyRoutesixRuleDestinationGroupDomainGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDestinationGroupDomainGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-group"]; ok {
		o.LeafPolicyRoutesixRuleDestinationGroupMacGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDestinationGroupMacGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network-group"]; ok {
		o.LeafPolicyRoutesixRuleDestinationGroupNetworkGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDestinationGroupNetworkGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port-group"]; ok {
		o.LeafPolicyRoutesixRuleDestinationGroupPortGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDestinationGroupPortGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
