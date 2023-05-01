// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesOpenvpnEncryption describes the resource data model.
type InterfacesOpenvpnEncryption struct {
	// LeafNodes
	LeafInterfacesOpenvpnEncryptionCIPher     types.String `tfsdk:"cipher" json:"cipher,omitempty"`
	LeafInterfacesOpenvpnEncryptionNcpCIPhers types.String `tfsdk:"ncp_ciphers" json:"ncp-ciphers,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnEncryption) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cipher": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Standard Data Encryption Algorithm

|  Format  |  Description  |
|----------|---------------|
|  none  |  Disable encryption  |
|  des  |  DES algorithm  |
|  3des  |  DES algorithm with triple encryption  |
|  bf128  |  Blowfish algorithm with 128-bit key  |
|  bf256  |  Blowfish algorithm with 256-bit key  |
|  aes128  |  AES algorithm with 128-bit key CBC  |
|  aes128gcm  |  AES algorithm with 128-bit key GCM  |
|  aes192  |  AES algorithm with 192-bit key CBC  |
|  aes192gcm  |  AES algorithm with 192-bit key GCM  |
|  aes256  |  AES algorithm with 256-bit key CBC  |
|  aes256gcm  |  AES algorithm with 256-bit key GCM  |

`,
		},

		"ncp_ciphers": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cipher negotiation list for use in server or client mode

|  Format  |  Description  |
|----------|---------------|
|  none  |  Disable encryption  |
|  des  |  DES algorithm  |
|  3des  |  DES algorithm with triple encryption  |
|  aes128  |  AES algorithm with 128-bit key CBC  |
|  aes128gcm  |  AES algorithm with 128-bit key GCM  |
|  aes192  |  AES algorithm with 192-bit key CBC  |
|  aes192gcm  |  AES algorithm with 192-bit key GCM  |
|  aes256  |  AES algorithm with 256-bit key CBC  |
|  aes256gcm  |  AES algorithm with 256-bit key GCM  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesOpenvpnEncryption) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesOpenvpnEncryptionCIPher.IsNull() && !o.LeafInterfacesOpenvpnEncryptionCIPher.IsUnknown() {
		jsonData["cipher"] = o.LeafInterfacesOpenvpnEncryptionCIPher.ValueString()
	}

	if !o.LeafInterfacesOpenvpnEncryptionNcpCIPhers.IsNull() && !o.LeafInterfacesOpenvpnEncryptionNcpCIPhers.IsUnknown() {
		jsonData["ncp-ciphers"] = o.LeafInterfacesOpenvpnEncryptionNcpCIPhers.ValueString()
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
func (o *InterfacesOpenvpnEncryption) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["cipher"]; ok {
		o.LeafInterfacesOpenvpnEncryptionCIPher = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnEncryptionCIPher = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ncp-ciphers"]; ok {
		o.LeafInterfacesOpenvpnEncryptionNcpCIPhers = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnEncryptionNcpCIPhers = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
