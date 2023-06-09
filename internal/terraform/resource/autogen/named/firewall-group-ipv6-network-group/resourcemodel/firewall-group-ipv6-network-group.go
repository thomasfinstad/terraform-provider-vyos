// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallGroupIPvsixNetworkGroup describes the resource data model.
type FirewallGroupIPvsixNetworkGroup struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafFirewallGroupIPvsixNetworkGroupDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`
	LeafFirewallGroupIPvsixNetworkGroupNetwork     types.String `tfsdk:"network" json:"network,omitempty"`
	LeafFirewallGroupIPvsixNetworkGroupInclude     types.String `tfsdk:"include" json:"include,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupIPvsixNetworkGroup) GetVyosPath() []string {
	return []string{
		"firewall",
		"group",
		"ipv6-network-group",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupIPvsixNetworkGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall ipv6-network-group

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network-group member

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address to match  |

`,
		},

		"include": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Include another ipv6-network-group

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallGroupIPvsixNetworkGroup) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallGroupIPvsixNetworkGroupDescrIPtion.IsNull() && !o.LeafFirewallGroupIPvsixNetworkGroupDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafFirewallGroupIPvsixNetworkGroupDescrIPtion.ValueString()
	}

	if !o.LeafFirewallGroupIPvsixNetworkGroupNetwork.IsNull() && !o.LeafFirewallGroupIPvsixNetworkGroupNetwork.IsUnknown() {
		jsonData["network"] = o.LeafFirewallGroupIPvsixNetworkGroupNetwork.ValueString()
	}

	if !o.LeafFirewallGroupIPvsixNetworkGroupInclude.IsNull() && !o.LeafFirewallGroupIPvsixNetworkGroupInclude.IsUnknown() {
		jsonData["include"] = o.LeafFirewallGroupIPvsixNetworkGroupInclude.ValueString()
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
func (o *FirewallGroupIPvsixNetworkGroup) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafFirewallGroupIPvsixNetworkGroupDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupIPvsixNetworkGroupDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network"]; ok {
		o.LeafFirewallGroupIPvsixNetworkGroupNetwork = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupIPvsixNetworkGroupNetwork = basetypes.NewStringNull()
	}

	if value, ok := jsonData["include"]; ok {
		o.LeafFirewallGroupIPvsixNetworkGroupInclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallGroupIPvsixNetworkGroupInclude = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
