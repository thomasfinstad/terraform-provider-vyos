// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallIPvsixNameRuleInboundInterface describes the resource data model.
type FirewallIPvsixNameRuleInboundInterface struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceName  types.String `tfsdk:"interface_name" json:"interface-name,omitempty"`
	LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceGroup types.String `tfsdk:"interface_group" json:"interface-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleInboundInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface

`,
		},

		"interface_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface-group

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallIPvsixNameRuleInboundInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceName.IsNull() && !o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceName.IsUnknown() {
		jsonData["interface-name"] = o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceName.ValueString()
	}

	if !o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceGroup.IsNull() && !o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceGroup.IsUnknown() {
		jsonData["interface-group"] = o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceGroup.ValueString()
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
func (o *FirewallIPvsixNameRuleInboundInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["interface-name"]; ok {
		o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["interface-group"]; ok {
		o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
