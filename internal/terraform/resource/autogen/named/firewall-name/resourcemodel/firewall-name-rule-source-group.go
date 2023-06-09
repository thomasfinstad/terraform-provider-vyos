// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallNameRuleSourceGroup describes the resource data model.
type FirewallNameRuleSourceGroup struct {
	// LeafNodes
	LeafFirewallNameRuleSourceGroupAddressGroup types.String `tfsdk:"address_group" json:"address-group,omitempty"`
	LeafFirewallNameRuleSourceGroupDomainGroup  types.String `tfsdk:"domain_group" json:"domain-group,omitempty"`
	LeafFirewallNameRuleSourceGroupMacGroup     types.String `tfsdk:"mac_group" json:"mac-group,omitempty"`
	LeafFirewallNameRuleSourceGroupNetworkGroup types.String `tfsdk:"network_group" json:"network-group,omitempty"`
	LeafFirewallNameRuleSourceGroupPortGroup    types.String `tfsdk:"port_group" json:"port-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRuleSourceGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *FirewallNameRuleSourceGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallNameRuleSourceGroupAddressGroup.IsNull() && !o.LeafFirewallNameRuleSourceGroupAddressGroup.IsUnknown() {
		jsonData["address-group"] = o.LeafFirewallNameRuleSourceGroupAddressGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleSourceGroupDomainGroup.IsNull() && !o.LeafFirewallNameRuleSourceGroupDomainGroup.IsUnknown() {
		jsonData["domain-group"] = o.LeafFirewallNameRuleSourceGroupDomainGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleSourceGroupMacGroup.IsNull() && !o.LeafFirewallNameRuleSourceGroupMacGroup.IsUnknown() {
		jsonData["mac-group"] = o.LeafFirewallNameRuleSourceGroupMacGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleSourceGroupNetworkGroup.IsNull() && !o.LeafFirewallNameRuleSourceGroupNetworkGroup.IsUnknown() {
		jsonData["network-group"] = o.LeafFirewallNameRuleSourceGroupNetworkGroup.ValueString()
	}

	if !o.LeafFirewallNameRuleSourceGroupPortGroup.IsNull() && !o.LeafFirewallNameRuleSourceGroupPortGroup.IsUnknown() {
		jsonData["port-group"] = o.LeafFirewallNameRuleSourceGroupPortGroup.ValueString()
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
func (o *FirewallNameRuleSourceGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address-group"]; ok {
		o.LeafFirewallNameRuleSourceGroupAddressGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleSourceGroupAddressGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["domain-group"]; ok {
		o.LeafFirewallNameRuleSourceGroupDomainGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleSourceGroupDomainGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mac-group"]; ok {
		o.LeafFirewallNameRuleSourceGroupMacGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleSourceGroupMacGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network-group"]; ok {
		o.LeafFirewallNameRuleSourceGroupNetworkGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleSourceGroupNetworkGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port-group"]; ok {
		o.LeafFirewallNameRuleSourceGroupPortGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleSourceGroupPortGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
