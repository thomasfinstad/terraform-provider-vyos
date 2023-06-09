// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesSstpcAuthentication describes the resource data model.
type InterfacesSstpcAuthentication struct {
	// LeafNodes
	LeafInterfacesSstpcAuthenticationUsername types.String `tfsdk:"username" json:"username,omitempty"`
	LeafInterfacesSstpcAuthenticationPassword types.String `tfsdk:"password" json:"password,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesSstpcAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesSstpcAuthentication) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesSstpcAuthenticationUsername.IsNull() && !o.LeafInterfacesSstpcAuthenticationUsername.IsUnknown() {
		jsonData["username"] = o.LeafInterfacesSstpcAuthenticationUsername.ValueString()
	}

	if !o.LeafInterfacesSstpcAuthenticationPassword.IsNull() && !o.LeafInterfacesSstpcAuthenticationPassword.IsUnknown() {
		jsonData["password"] = o.LeafInterfacesSstpcAuthenticationPassword.ValueString()
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
func (o *InterfacesSstpcAuthentication) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["username"]; ok {
		o.LeafInterfacesSstpcAuthenticationUsername = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesSstpcAuthenticationUsername = basetypes.NewStringNull()
	}

	if value, ok := jsonData["password"]; ok {
		o.LeafInterfacesSstpcAuthenticationPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesSstpcAuthenticationPassword = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
