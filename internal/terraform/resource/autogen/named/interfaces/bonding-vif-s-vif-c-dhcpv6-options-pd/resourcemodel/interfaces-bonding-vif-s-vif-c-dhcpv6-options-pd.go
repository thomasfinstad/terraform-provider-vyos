// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBondingVifSVifCDhcpvsixOptionsPd describes the resource data model.
type InterfacesBondingVifSVifCDhcpvsixOptionsPd struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesBonding any `tfsdk:"bonding" vyos:"bonding,parent-id"`

	ParentIDInterfacesBondingVifS any `tfsdk:"vif_s" vyos:"vif-s,parent-id"`

	ParentIDInterfacesBondingVifSVifC any `tfsdk:"vif_c" vyos:"vif-c,parent-id"`

	// LeafNodes
	LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdLength types.String `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesBondingVifSVifCDhcpvsixOptionsPdInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPd) GetVyosPath() []string {
	return []string{
		"interfaces",
		"bonding",
		"vif-s",
		"vif-c",
		"dhcpv6-options",
		"pd",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSVifCDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format  |  Description  |
    |----------|---------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |

`,
		},

		// LeafNodes

		"length": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Request IPv6 prefix length from peer

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:32-64  |  Length of delegated prefix  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPd) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdLength.IsNull() && !o.LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdLength.IsUnknown() {
		jsonData["length"] = o.LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdLength.ValueString()
	}

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPd) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["length"]; ok {
		o.LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdLength = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
