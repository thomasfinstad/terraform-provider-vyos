// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsOspfInterfaceAuthentication describes the resource data model.
type VrfNameProtocolsOspfInterfaceAuthentication struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfInterfaceAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password" json:"plaintext-password,omitempty"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsOspfInterfaceAuthenticationMdfive *VrfNameProtocolsOspfInterfaceAuthenticationMdfive `tfsdk:"md5" json:"md5,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfInterfaceAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain text password

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Plain text password (8 characters or less)  |

`,
		},

		// TagNodes

		// Nodes

		"md5": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceAuthenticationMdfive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `MD5 key id

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfInterfaceAuthentication) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsOspfInterfaceAuthenticationPlaintextPassword.IsNull() && !o.LeafVrfNameProtocolsOspfInterfaceAuthenticationPlaintextPassword.IsUnknown() {
		jsonData["plaintext-password"] = o.LeafVrfNameProtocolsOspfInterfaceAuthenticationPlaintextPassword.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfInterfaceAuthenticationMdfive).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfInterfaceAuthenticationMdfive)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["md5"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfInterfaceAuthentication) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["plaintext-password"]; ok {
		o.LeafVrfNameProtocolsOspfInterfaceAuthenticationPlaintextPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfInterfaceAuthenticationPlaintextPassword = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["md5"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfInterfaceAuthenticationMdfive = &VrfNameProtocolsOspfInterfaceAuthenticationMdfive{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfInterfaceAuthenticationMdfive)
		if err != nil {
			return err
		}
	}

	return nil
}
