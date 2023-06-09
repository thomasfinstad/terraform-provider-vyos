// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsecRemoteAccessConnectionAuthentication describes the resource data model.
type VpnIPsecRemoteAccessConnectionAuthentication struct {
	// LeafNodes
	LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID         types.String `tfsdk:"local_id" json:"local-id,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode      types.String `tfsdk:"client_mode" json:"client-mode,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode      types.String `tfsdk:"server_mode" json:"server-mode,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret types.String `tfsdk:"pre_shared_secret" json:"pre-shared-secret,omitempty"`

	// TagNodes

	// Nodes
	NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine *VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine `tfsdk:"x509" json:"x509,omitempty"`
	NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers    *VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers    `tfsdk:"local_users" json:"local-users,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessConnectionAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local ID for peer authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Local ID used for peer authentication  |

`,
		},

		"client_mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client authentication mode

|  Format  |  Description  |
|----------|---------------|
|  eap-tls  |  Use EAP-TLS authentication  |
|  eap-mschapv2  |  Use EAP-MSCHAPv2 authentication  |
|  eap-radius  |  Use EAP-RADIUS authentication  |

`,

			// Default:          stringdefault.StaticString(`eap-mschapv2`),
			Computed: true,
		},

		"server_mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Server authentication mode

|  Format  |  Description  |
|----------|---------------|
|  pre-shared-secret  |  Use a pre-shared secret key  |
|  x509  |  Use x.509 certificate  |

`,

			// Default:          stringdefault.StaticString(`x509`),
			Computed: true,
		},

		"pre_shared_secret": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pre-shared secret key

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Pre-shared secret key  |

`,
		},

		// TagNodes

		// Nodes

		"x509": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `X.509 certificate

`,
		},

		"local_users": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Local user authentication

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecRemoteAccessConnectionAuthentication) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID.IsUnknown() {
		jsonData["local-id"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode.IsUnknown() {
		jsonData["client-mode"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode.IsUnknown() {
		jsonData["server-mode"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret.IsUnknown() {
		jsonData["pre-shared-secret"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["x509"] = subData
	}

	if !reflect.ValueOf(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["local-users"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecRemoteAccessConnectionAuthentication) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["local-id"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["client-mode"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode = basetypes.NewStringNull()
	}

	if value, ok := jsonData["server-mode"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode = basetypes.NewStringNull()
	}

	if value, ok := jsonData["pre-shared-secret"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["x509"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine = &VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine{}

		err = json.Unmarshal(subJSONStr, o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["local-users"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers = &VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers{}

		err = json.Unmarshal(subJSONStr, o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers)
		if err != nil {
			return err
		}
	}

	return nil
}
