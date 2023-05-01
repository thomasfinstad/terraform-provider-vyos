// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallGroupInterfaceGroup describes the resource data model.
type FirewallGroupInterfaceGroup struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafFirewallGroupInterfaceGroupInterface   types.String `tfsdk:"interface" json:"interface,omitempty"`
	LeafFirewallGroupInterfaceGroupInclude     types.String `tfsdk:"include" json:"include,omitempty"`
	LeafFirewallGroupInterfaceGroupDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupInterfaceGroup) GetVyosPath() []string {
	return []string{
		"firewall",
		"group",
		"interface-group",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupInterfaceGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall interface-group

`,
		},

		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface-group member

`,
		},

		"include": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Include another interface-group

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
func (o *FirewallGroupInterfaceGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallGroupInterfaceGroupInterface.IsNull() && !o.LeafFirewallGroupInterfaceGroupInterface.IsUnknown() {
		jsonData["interface"] = o.LeafFirewallGroupInterfaceGroupInterface.ValueString()
	}

	if !o.LeafFirewallGroupInterfaceGroupInclude.IsNull() && !o.LeafFirewallGroupInterfaceGroupInclude.IsUnknown() {
		jsonData["include"] = o.LeafFirewallGroupInterfaceGroupInclude.ValueString()
	}

	if !o.LeafFirewallGroupInterfaceGroupDescrIPtion.IsNull() && !o.LeafFirewallGroupInterfaceGroupDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafFirewallGroupInterfaceGroupDescrIPtion.ValueString()
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
func (o *FirewallGroupInterfaceGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["interface"]; ok {
		o.LeafFirewallGroupInterfaceGroupInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupInterfaceGroupInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["include"]; ok {
		o.LeafFirewallGroupInterfaceGroupInclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupInterfaceGroupInclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafFirewallGroupInterfaceGroupDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupInterfaceGroupDescrIPtion = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
