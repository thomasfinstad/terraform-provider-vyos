// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallNameRuleDestinationGroup describes the resource data model.
type FirewallNameRuleDestinationGroup struct {
	// LeafNodes
	LeafFirewallNameRuleDestinationGroupAddressGroup types.String `tfsdk:"address_group" json:"address-group,omitempty"`
	LeafFirewallNameRuleDestinationGroupDomainGroup  types.String `tfsdk:"domain_group" json:"domain-group,omitempty"`
	LeafFirewallNameRuleDestinationGroupMacGroup     types.String `tfsdk:"mac_group" json:"mac-group,omitempty"`
	LeafFirewallNameRuleDestinationGroupNetworkGroup types.String `tfsdk:"network_group" json:"network-group,omitempty"`
	LeafFirewallNameRuleDestinationGroupPortGroup    types.String `tfsdk:"port_group" json:"port-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRuleDestinationGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *FirewallNameRuleDestinationGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallNameRuleDestinationGroupAddressGroup.IsNull() && !o.LeafFirewallNameRuleDestinationGroupAddressGroup.IsUnknown() {
		jsonData["address-group"] = o.LeafFirewallNameRuleDestinationGroupAddressGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleDestinationGroupDomainGroup.IsNull() && !o.LeafFirewallNameRuleDestinationGroupDomainGroup.IsUnknown() {
		jsonData["domain-group"] = o.LeafFirewallNameRuleDestinationGroupDomainGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleDestinationGroupMacGroup.IsNull() && !o.LeafFirewallNameRuleDestinationGroupMacGroup.IsUnknown() {
		jsonData["mac-group"] = o.LeafFirewallNameRuleDestinationGroupMacGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleDestinationGroupNetworkGroup.IsNull() && !o.LeafFirewallNameRuleDestinationGroupNetworkGroup.IsUnknown() {
		jsonData["network-group"] = o.LeafFirewallNameRuleDestinationGroupNetworkGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleDestinationGroupPortGroup.IsNull() && !o.LeafFirewallNameRuleDestinationGroupPortGroup.IsUnknown() {
		jsonData["port-group"] = o.LeafFirewallNameRuleDestinationGroupPortGroup.ValueString()
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
func (o *FirewallNameRuleDestinationGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address-group"]; ok {
		o.LeafFirewallNameRuleDestinationGroupAddressGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDestinationGroupAddressGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["domain-group"]; ok {
		o.LeafFirewallNameRuleDestinationGroupDomainGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDestinationGroupDomainGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-group"]; ok {
		o.LeafFirewallNameRuleDestinationGroupMacGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDestinationGroupMacGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network-group"]; ok {
		o.LeafFirewallNameRuleDestinationGroupNetworkGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDestinationGroupNetworkGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port-group"]; ok {
		o.LeafFirewallNameRuleDestinationGroupPortGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDestinationGroupPortGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
