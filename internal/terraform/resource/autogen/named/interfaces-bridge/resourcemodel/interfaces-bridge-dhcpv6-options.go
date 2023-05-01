// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBrIDgeDhcpvsixOptions describes the resource data model.
type InterfacesBrIDgeDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesBrIDgeDhcpvsixOptionsDuID           types.String `tfsdk:"duid" json:"duid,omitempty"`
	LeafInterfacesBrIDgeDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only" json:"parameters-only,omitempty"`
	LeafInterfacesBrIDgeDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit" json:"rapid-commit,omitempty"`
	LeafInterfacesBrIDgeDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary" json:"temporary,omitempty"`

	// TagNodes
	TagInterfacesBrIDgeDhcpvsixOptionsPd *map[string]InterfacesBrIDgeDhcpvsixOptionsPd `tfsdk:"pd" json:"pd,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesBrIDgeDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
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
func (o *InterfacesBrIDgeDhcpvsixOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBrIDgeDhcpvsixOptionsDuID.IsNull() && !o.LeafInterfacesBrIDgeDhcpvsixOptionsDuID.IsUnknown() {
		jsonData["duid"] = o.LeafInterfacesBrIDgeDhcpvsixOptionsDuID.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpvsixOptionsParametersOnly.IsNull() && !o.LeafInterfacesBrIDgeDhcpvsixOptionsParametersOnly.IsUnknown() {
		jsonData["parameters-only"] = o.LeafInterfacesBrIDgeDhcpvsixOptionsParametersOnly.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpvsixOptionsRAPIDCommit.IsNull() && !o.LeafInterfacesBrIDgeDhcpvsixOptionsRAPIDCommit.IsUnknown() {
		jsonData["rapid-commit"] = o.LeafInterfacesBrIDgeDhcpvsixOptionsRAPIDCommit.ValueString()
	}

	if !o.LeafInterfacesBrIDgeDhcpvsixOptionsTemporary.IsNull() && !o.LeafInterfacesBrIDgeDhcpvsixOptionsTemporary.IsUnknown() {
		jsonData["temporary"] = o.LeafInterfacesBrIDgeDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagInterfacesBrIDgeDhcpvsixOptionsPd).IsZero() {
		subJSONStr, err := json.Marshal(o.TagInterfacesBrIDgeDhcpvsixOptionsPd)
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
func (o *InterfacesBrIDgeDhcpvsixOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["duid"]; ok {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["parameters-only"]; ok {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rapid-commit"]; ok {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}

	if value, ok := jsonData["temporary"]; ok {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["pd"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagInterfacesBrIDgeDhcpvsixOptionsPd = &map[string]InterfacesBrIDgeDhcpvsixOptionsPd{}

		err = json.Unmarshal(subJSONStr, o.TagInterfacesBrIDgeDhcpvsixOptionsPd)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
