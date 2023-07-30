// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBondingDhcpvsixOptions describes the resource data model.
type InterfacesBondingDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesBondingDhcpvsixOptionsDuID           types.String `tfsdk:"duid" vyos:"duid,omitempty"`
	LeafInterfacesBondingDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only" vyos:"parameters-only,omitempty"`
	LeafInterfacesBondingDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit" vyos:"rapid-commit,omitempty"`
	LeafInterfacesBondingDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary" vyos:"temporary,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesBondingDhcpvsixOptionsPd bool `tfsdk:"pd" vyos:"pd,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBondingDhcpvsixOptions) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBondingDhcpvsixOptionsDuID.IsNull() && !o.LeafInterfacesBondingDhcpvsixOptionsDuID.IsUnknown() {
		jsonData["duid"] = o.LeafInterfacesBondingDhcpvsixOptionsDuID.ValueString()
	}

	if !o.LeafInterfacesBondingDhcpvsixOptionsParametersOnly.IsNull() && !o.LeafInterfacesBondingDhcpvsixOptionsParametersOnly.IsUnknown() {
		jsonData["parameters-only"] = o.LeafInterfacesBondingDhcpvsixOptionsParametersOnly.ValueString()
	}

	if !o.LeafInterfacesBondingDhcpvsixOptionsRAPIDCommit.IsNull() && !o.LeafInterfacesBondingDhcpvsixOptionsRAPIDCommit.IsUnknown() {
		jsonData["rapid-commit"] = o.LeafInterfacesBondingDhcpvsixOptionsRAPIDCommit.ValueString()
	}

	if !o.LeafInterfacesBondingDhcpvsixOptionsTemporary.IsNull() && !o.LeafInterfacesBondingDhcpvsixOptionsTemporary.IsUnknown() {
		jsonData["temporary"] = o.LeafInterfacesBondingDhcpvsixOptionsTemporary.ValueString()
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
func (o *InterfacesBondingDhcpvsixOptions) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["duid"]; ok {
		o.LeafInterfacesBondingDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["parameters-only"]; ok {
		o.LeafInterfacesBondingDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rapid-commit"]; ok {
		o.LeafInterfacesBondingDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}

	if value, ok := jsonData["temporary"]; ok {
		o.LeafInterfacesBondingDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
