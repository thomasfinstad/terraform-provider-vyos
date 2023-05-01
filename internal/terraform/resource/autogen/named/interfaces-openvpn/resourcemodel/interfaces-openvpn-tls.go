// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesOpenvpnTLS describes the resource data model.
type InterfacesOpenvpnTLS struct {
	// LeafNodes
	LeafInterfacesOpenvpnTLSAuthKey       types.String `tfsdk:"auth_key" json:"auth-key,omitempty"`
	LeafInterfacesOpenvpnTLSCertificate   types.String `tfsdk:"certificate" json:"certificate,omitempty"`
	LeafInterfacesOpenvpnTLSCaCertificate types.String `tfsdk:"ca_certificate" json:"ca-certificate,omitempty"`
	LeafInterfacesOpenvpnTLSDhParams      types.String `tfsdk:"dh_params" json:"dh-params,omitempty"`
	LeafInterfacesOpenvpnTLSCryptKey      types.String `tfsdk:"crypt_key" json:"crypt-key,omitempty"`
	LeafInterfacesOpenvpnTLSTLSVersionMin types.String `tfsdk:"tls_version_min" json:"tls-version-min,omitempty"`
	LeafInterfacesOpenvpnTLSRole          types.String `tfsdk:"role" json:"role,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnTLS) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"auth_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TLS shared secret key for tls-auth

`,
		},

		"certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of certificate in PKI configuration  |

`,
		},

		"ca_certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate Authority chain in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of CA in PKI configuration  |

`,
		},

		"dh_params": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Diffie Hellman parameters (server only)

`,
		},

		"crypt_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Static key to use to authenticate control channel

`,
		},

		"tls_version_min": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify the minimum required TLS version

|  Format  |  Description  |
|----------|---------------|
|  1.0  |  TLS v1.0  |
|  1.1  |  TLS v1.1  |
|  1.2  |  TLS v1.2  |
|  1.3  |  TLS v1.3  |

`,
		},

		"role": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TLS negotiation role

|  Format  |  Description  |
|----------|---------------|
|  active  |  Initiate TLS negotiation actively  |
|  passive  |  Wait for incoming TLS connection  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesOpenvpnTLS) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesOpenvpnTLSAuthKey.IsNull() && !o.LeafInterfacesOpenvpnTLSAuthKey.IsUnknown() {
		jsonData["auth-key"] = o.LeafInterfacesOpenvpnTLSAuthKey.ValueString()
	}

	if !o.LeafInterfacesOpenvpnTLSCertificate.IsNull() && !o.LeafInterfacesOpenvpnTLSCertificate.IsUnknown() {
		jsonData["certificate"] = o.LeafInterfacesOpenvpnTLSCertificate.ValueString()
	}

	if !o.LeafInterfacesOpenvpnTLSCaCertificate.IsNull() && !o.LeafInterfacesOpenvpnTLSCaCertificate.IsUnknown() {
		jsonData["ca-certificate"] = o.LeafInterfacesOpenvpnTLSCaCertificate.ValueString()
	}

	if !o.LeafInterfacesOpenvpnTLSDhParams.IsNull() && !o.LeafInterfacesOpenvpnTLSDhParams.IsUnknown() {
		jsonData["dh-params"] = o.LeafInterfacesOpenvpnTLSDhParams.ValueString()
	}

	if !o.LeafInterfacesOpenvpnTLSCryptKey.IsNull() && !o.LeafInterfacesOpenvpnTLSCryptKey.IsUnknown() {
		jsonData["crypt-key"] = o.LeafInterfacesOpenvpnTLSCryptKey.ValueString()
	}

	if !o.LeafInterfacesOpenvpnTLSTLSVersionMin.IsNull() && !o.LeafInterfacesOpenvpnTLSTLSVersionMin.IsUnknown() {
		jsonData["tls-version-min"] = o.LeafInterfacesOpenvpnTLSTLSVersionMin.ValueString()
	}

	if !o.LeafInterfacesOpenvpnTLSRole.IsNull() && !o.LeafInterfacesOpenvpnTLSRole.IsUnknown() {
		jsonData["role"] = o.LeafInterfacesOpenvpnTLSRole.ValueString()
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
func (o *InterfacesOpenvpnTLS) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["auth-key"]; ok {
		o.LeafInterfacesOpenvpnTLSAuthKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnTLSAuthKey = basetypes.NewStringNull()
	}

	if value, ok := jsonData["certificate"]; ok {
		o.LeafInterfacesOpenvpnTLSCertificate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnTLSCertificate = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ca-certificate"]; ok {
		o.LeafInterfacesOpenvpnTLSCaCertificate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnTLSCaCertificate = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dh-params"]; ok {
		o.LeafInterfacesOpenvpnTLSDhParams = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnTLSDhParams = basetypes.NewStringNull()
	}

	if value, ok := jsonData["crypt-key"]; ok {
		o.LeafInterfacesOpenvpnTLSCryptKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnTLSCryptKey = basetypes.NewStringNull()
	}

	if value, ok := jsonData["tls-version-min"]; ok {
		o.LeafInterfacesOpenvpnTLSTLSVersionMin = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnTLSTLSVersionMin = basetypes.NewStringNull()
	}

	if value, ok := jsonData["role"]; ok {
		o.LeafInterfacesOpenvpnTLSRole = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnTLSRole = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
