// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesMacsecDhcpvsixOptions describes the resource data model.
type InterfacesMacsecDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesMacsecDhcpvsixOptionsDuID           types.String `tfsdk:"duid" json:"duid,omitempty"`
	LeafInterfacesMacsecDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only" json:"parameters-only,omitempty"`
	LeafInterfacesMacsecDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit" json:"rapid-commit,omitempty"`
	LeafInterfacesMacsecDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary" json:"temporary,omitempty"`

	// TagNodes
	TagInterfacesMacsecDhcpvsixOptionsPd *map[string]InterfacesMacsecDhcpvsixOptionsPd `tfsdk:"pd" json:"pd,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesMacsecDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
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
func (o *InterfacesMacsecDhcpvsixOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesMacsecDhcpvsixOptionsDuID.IsNull() && !o.LeafInterfacesMacsecDhcpvsixOptionsDuID.IsUnknown() {
		jsonData["duid"] = o.LeafInterfacesMacsecDhcpvsixOptionsDuID.ValueString()
	}

	if !o.LeafInterfacesMacsecDhcpvsixOptionsParametersOnly.IsNull() && !o.LeafInterfacesMacsecDhcpvsixOptionsParametersOnly.IsUnknown() {
		jsonData["parameters-only"] = o.LeafInterfacesMacsecDhcpvsixOptionsParametersOnly.ValueString()
	}

	if !o.LeafInterfacesMacsecDhcpvsixOptionsRAPIDCommit.IsNull() && !o.LeafInterfacesMacsecDhcpvsixOptionsRAPIDCommit.IsUnknown() {
		jsonData["rapid-commit"] = o.LeafInterfacesMacsecDhcpvsixOptionsRAPIDCommit.ValueString()
	}

	if !o.LeafInterfacesMacsecDhcpvsixOptionsTemporary.IsNull() && !o.LeafInterfacesMacsecDhcpvsixOptionsTemporary.IsUnknown() {
		jsonData["temporary"] = o.LeafInterfacesMacsecDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagInterfacesMacsecDhcpvsixOptionsPd).IsZero() {
		subJSONStr, err := json.Marshal(o.TagInterfacesMacsecDhcpvsixOptionsPd)
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
func (o *InterfacesMacsecDhcpvsixOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["duid"]; ok {
		o.LeafInterfacesMacsecDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["parameters-only"]; ok {
		o.LeafInterfacesMacsecDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rapid-commit"]; ok {
		o.LeafInterfacesMacsecDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}

	if value, ok := jsonData["temporary"]; ok {
		o.LeafInterfacesMacsecDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["pd"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagInterfacesMacsecDhcpvsixOptionsPd = &map[string]InterfacesMacsecDhcpvsixOptionsPd{}

		err = json.Unmarshal(subJSONStr, o.TagInterfacesMacsecDhcpvsixOptionsPd)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
