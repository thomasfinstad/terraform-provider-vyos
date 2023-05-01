// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesWirelessDhcpvsixOptionsPdInterface struct {
	// LeafNodes
	LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" json:"address,omitempty"`
	LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceSLAID   types.String `tfsdk:"sla_id" json:"sla-id,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
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

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceAddress.IsNull() && !o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceAddress.IsUnknown() {
		jsonData["address"] = o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceAddress.ValueString()
	}

	if !o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceSLAID.IsNull() && !o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceSLAID.IsUnknown() {
		jsonData["sla-id"] = o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceSLAID.ValueString()
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
func (o *InterfacesWirelessDhcpvsixOptionsPdInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["sla-id"]; ok {
		o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceSLAID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceSLAID = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
