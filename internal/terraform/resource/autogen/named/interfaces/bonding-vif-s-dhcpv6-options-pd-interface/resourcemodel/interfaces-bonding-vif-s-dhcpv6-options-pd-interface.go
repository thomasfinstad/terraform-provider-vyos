// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBondingVifSDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesBondingVifSDhcpvsixOptionsPdInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesBonding any `tfsdk:"bonding" vyos:"bonding,parent-id"`

	ParentIDInterfacesBondingVifS any `tfsdk:"vif_s" vyos:"vif-s,parent-id"`

	ParentIDInterfacesBondingVifSDhcpvsixOptionsPd any `tfsdk:"pd" vyos:"pd,parent-id"`

	// LeafNodes
	LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceSLAID   types.String `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBondingVifSDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	return []string{
		"interfaces",
		"bonding",
		"vif-s",
		"dhcpv6-options",
		"pd",
		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

    |  Format  |  Description  |
    |----------|---------------|
    |  >0  |  Used to form IPv6 interface address  |

`,
		},

		"sla_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface site-Level aggregator (SLA)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-65535  |  Decimal integer which fits in the length of SLA IDs  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBondingVifSDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceAddress.IsNull() && !o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceAddress.ValueString()
	}

	if !o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceSLAID.IsNull() && !o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceSLAID.IsUnknown() {
		jsonData["sla-id"] = o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceSLAID.ValueString()
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
func (o *InterfacesBondingVifSDhcpvsixOptionsPdInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["sla-id"]; ok {
		o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceSLAID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSDhcpvsixOptionsPdInterfaceSLAID = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
