// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesOpenvpnAuthentication describes the resource data model.
type InterfacesOpenvpnAuthentication struct {
	// LeafNodes
	LeafInterfacesOpenvpnAuthenticationUsername types.String `tfsdk:"username" json:"username,omitempty"`
	LeafInterfacesOpenvpnAuthenticationPassword types.String `tfsdk:"password" json:"password,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"username": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username  |

`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Password  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesOpenvpnAuthentication) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesOpenvpnAuthenticationUsername.IsNull() && !o.LeafInterfacesOpenvpnAuthenticationUsername.IsUnknown() {
		jsonData["username"] = o.LeafInterfacesOpenvpnAuthenticationUsername.ValueString()
	}

	if !o.LeafInterfacesOpenvpnAuthenticationPassword.IsNull() && !o.LeafInterfacesOpenvpnAuthenticationPassword.IsUnknown() {
		jsonData["password"] = o.LeafInterfacesOpenvpnAuthenticationPassword.ValueString()
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
func (o *InterfacesOpenvpnAuthentication) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["username"]; ok {
		o.LeafInterfacesOpenvpnAuthenticationUsername = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnAuthenticationUsername = basetypes.NewStringNull()
	}

	if value, ok := jsonData["password"]; ok {
		o.LeafInterfacesOpenvpnAuthenticationPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnAuthenticationPassword = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
