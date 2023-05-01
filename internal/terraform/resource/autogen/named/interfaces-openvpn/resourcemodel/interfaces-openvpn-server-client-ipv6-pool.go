// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesOpenvpnServerClientIPvsixPool describes the resource data model.
type InterfacesOpenvpnServerClientIPvsixPool struct {
	// LeafNodes
	LeafInterfacesOpenvpnServerClientIPvsixPoolBase    types.String `tfsdk:"base" json:"base,omitempty"`
	LeafInterfacesOpenvpnServerClientIPvsixPoolDisable types.String `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerClientIPvsixPool) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"base": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client IPv6 pool base address with optional prefix length

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Client IPv6 pool base address with optional prefix length (defaults: base = server subnet + 0x1000, prefix length = server prefix length)  |

`,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesOpenvpnServerClientIPvsixPool) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesOpenvpnServerClientIPvsixPoolBase.IsNull() && !o.LeafInterfacesOpenvpnServerClientIPvsixPoolBase.IsUnknown() {
		jsonData["base"] = o.LeafInterfacesOpenvpnServerClientIPvsixPoolBase.ValueString()
	}

	if !o.LeafInterfacesOpenvpnServerClientIPvsixPoolDisable.IsNull() && !o.LeafInterfacesOpenvpnServerClientIPvsixPoolDisable.IsUnknown() {
		jsonData["disable"] = o.LeafInterfacesOpenvpnServerClientIPvsixPoolDisable.ValueString()
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
func (o *InterfacesOpenvpnServerClientIPvsixPool) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["base"]; ok {
		o.LeafInterfacesOpenvpnServerClientIPvsixPoolBase = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnServerClientIPvsixPoolBase = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafInterfacesOpenvpnServerClientIPvsixPoolDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnServerClientIPvsixPoolDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
