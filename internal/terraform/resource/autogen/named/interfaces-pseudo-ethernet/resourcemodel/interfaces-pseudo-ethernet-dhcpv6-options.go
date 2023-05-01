// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesPseudoEthernetDhcpvsixOptions describes the resource data model.
type InterfacesPseudoEthernetDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetDhcpvsixOptionsDuID           types.String `tfsdk:"duid" json:"duid,omitempty"`
	LeafInterfacesPseudoEthernetDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only" json:"parameters-only,omitempty"`
	LeafInterfacesPseudoEthernetDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit" json:"rapid-commit,omitempty"`
	LeafInterfacesPseudoEthernetDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary" json:"temporary,omitempty"`

	// TagNodes
	TagInterfacesPseudoEthernetDhcpvsixOptionsPd *map[string]InterfacesPseudoEthernetDhcpvsixOptionsPd `tfsdk:"pd" json:"pd,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesPseudoEthernetDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
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
func (o *InterfacesPseudoEthernetDhcpvsixOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsDuID.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsDuID.IsUnknown() {
		jsonData["duid"] = o.LeafInterfacesPseudoEthernetDhcpvsixOptionsDuID.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsParametersOnly.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsParametersOnly.IsUnknown() {
		jsonData["parameters-only"] = o.LeafInterfacesPseudoEthernetDhcpvsixOptionsParametersOnly.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsRAPIDCommit.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsRAPIDCommit.IsUnknown() {
		jsonData["rapid-commit"] = o.LeafInterfacesPseudoEthernetDhcpvsixOptionsRAPIDCommit.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsTemporary.IsNull() && !o.LeafInterfacesPseudoEthernetDhcpvsixOptionsTemporary.IsUnknown() {
		jsonData["temporary"] = o.LeafInterfacesPseudoEthernetDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagInterfacesPseudoEthernetDhcpvsixOptionsPd).IsZero() {
		subJSONStr, err := json.Marshal(o.TagInterfacesPseudoEthernetDhcpvsixOptionsPd)
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
func (o *InterfacesPseudoEthernetDhcpvsixOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["duid"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["parameters-only"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rapid-commit"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}

	if value, ok := jsonData["temporary"]; ok {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["pd"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagInterfacesPseudoEthernetDhcpvsixOptionsPd = &map[string]InterfacesPseudoEthernetDhcpvsixOptionsPd{}

		err = json.Unmarshal(subJSONStr, o.TagInterfacesPseudoEthernetDhcpvsixOptionsPd)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
