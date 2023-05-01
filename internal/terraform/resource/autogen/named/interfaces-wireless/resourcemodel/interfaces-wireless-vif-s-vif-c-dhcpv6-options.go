// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesWirelessVifSVifCDhcpvsixOptions describes the resource data model.
type InterfacesWirelessVifSVifCDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesWirelessVifSVifCDhcpvsixOptionsDuID           types.String `tfsdk:"duid" json:"duid,omitempty"`
	LeafInterfacesWirelessVifSVifCDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only" json:"parameters-only,omitempty"`
	LeafInterfacesWirelessVifSVifCDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit" json:"rapid-commit,omitempty"`
	LeafInterfacesWirelessVifSVifCDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary" json:"temporary,omitempty"`

	// TagNodes
	TagInterfacesWirelessVifSVifCDhcpvsixOptionsPd *map[string]InterfacesWirelessVifSVifCDhcpvsixOptionsPd `tfsdk:"pd" json:"pd,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSVifCDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"duid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

|  Format  |  Description  |
|----------|---------------|
|  duid  |  DHCP unique identifier (DUID)  |

`,
		},

		"parameters_only": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Acquire only config parameters, no address

`,
		},

		"rapid_commit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
		},

		"temporary": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 temporary address

`,
		},

		// TagNodes

		"pd": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesWirelessVifSVifCDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

|  Format  |  Description  |
|----------|---------------|
|  instance number  |  Prefix delegation instance (>= 0)  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessVifSVifCDhcpvsixOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsDuID.IsNull() && !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsDuID.IsUnknown() {
		jsonData["duid"] = o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsDuID.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsParametersOnly.IsNull() && !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsParametersOnly.IsUnknown() {
		jsonData["parameters-only"] = o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsParametersOnly.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsRAPIDCommit.IsNull() && !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsRAPIDCommit.IsUnknown() {
		jsonData["rapid-commit"] = o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsRAPIDCommit.ValueString()
	}

	if !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsTemporary.IsNull() && !o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsTemporary.IsUnknown() {
		jsonData["temporary"] = o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagInterfacesWirelessVifSVifCDhcpvsixOptionsPd).IsZero() {
		subJSONStr, err := json.Marshal(o.TagInterfacesWirelessVifSVifCDhcpvsixOptionsPd)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["pd"] = subData
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
func (o *InterfacesWirelessVifSVifCDhcpvsixOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["duid"]; ok {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["parameters-only"]; ok {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rapid-commit"]; ok {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}

	if value, ok := jsonData["temporary"]; ok {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["pd"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagInterfacesWirelessVifSVifCDhcpvsixOptionsPd = &map[string]InterfacesWirelessVifSVifCDhcpvsixOptionsPd{}

		err = json.Unmarshal(subJSONStr, o.TagInterfacesWirelessVifSVifCDhcpvsixOptionsPd)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
