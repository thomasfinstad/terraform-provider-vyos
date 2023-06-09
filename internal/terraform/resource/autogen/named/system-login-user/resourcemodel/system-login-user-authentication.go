// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemLoginUserAuthentication describes the resource data model.
type SystemLoginUserAuthentication struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationEncryptedPassword types.String `tfsdk:"encrypted_password" json:"encrypted-password,omitempty"`
	LeafSystemLoginUserAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password" json:"plaintext-password,omitempty"`

	// TagNodes
	TagSystemLoginUserAuthenticationPublicKeys *map[string]SystemLoginUserAuthenticationPublicKeys `tfsdk:"public_keys" json:"public-keys,omitempty"`

	// Nodes
	NodeSystemLoginUserAuthenticationOtp *SystemLoginUserAuthenticationOtp `tfsdk:"otp" json:"otp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encrypted password

`,

			// Default:          stringdefault.StaticString(`!`),
			Computed: true,
		},

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plaintext password used for encryption

`,
		},

		// TagNodes

		"public_keys": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: SystemLoginUserAuthenticationPublicKeys{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Remote access public keys

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Key identifier used by ssh-keygen (usually of form user@host)  |

`,
		},

		// Nodes

		"otp": schema.SingleNestedAttribute{
			Attributes: SystemLoginUserAuthenticationOtp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `One-Time-Pad (two-factor) authentication parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemLoginUserAuthentication) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafSystemLoginUserAuthenticationEncryptedPassword.IsNull() && !o.LeafSystemLoginUserAuthenticationEncryptedPassword.IsUnknown() {
		jsonData["encrypted-password"] = o.LeafSystemLoginUserAuthenticationEncryptedPassword.ValueString()
	}

	if !o.LeafSystemLoginUserAuthenticationPlaintextPassword.IsNull() && !o.LeafSystemLoginUserAuthenticationPlaintextPassword.IsUnknown() {
		jsonData["plaintext-password"] = o.LeafSystemLoginUserAuthenticationPlaintextPassword.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagSystemLoginUserAuthenticationPublicKeys).IsZero() {
		subJSONStr, err := json.Marshal(o.TagSystemLoginUserAuthenticationPublicKeys)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["public-keys"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeSystemLoginUserAuthenticationOtp).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeSystemLoginUserAuthenticationOtp)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["otp"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemLoginUserAuthentication) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["encrypted-password"]; ok {
		o.LeafSystemLoginUserAuthenticationEncryptedPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationEncryptedPassword = basetypes.NewStringNull()
	}

	if value, ok := jsonData["plaintext-password"]; ok {
		o.LeafSystemLoginUserAuthenticationPlaintextPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationPlaintextPassword = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["public-keys"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagSystemLoginUserAuthenticationPublicKeys = &map[string]SystemLoginUserAuthenticationPublicKeys{}

		err = json.Unmarshal(subJSONStr, o.TagSystemLoginUserAuthenticationPublicKeys)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["otp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeSystemLoginUserAuthenticationOtp = &SystemLoginUserAuthenticationOtp{}

		err = json.Unmarshal(subJSONStr, o.NodeSystemLoginUserAuthenticationOtp)
		if err != nil {
			return err
		}
	}

	return nil
}
