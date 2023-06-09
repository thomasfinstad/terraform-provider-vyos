// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsecProfileAuthentication describes the resource data model.
type VpnIPsecProfileAuthentication struct {
	// LeafNodes
	LeafVpnIPsecProfileAuthenticationMode            types.String `tfsdk:"mode" json:"mode,omitempty"`
	LeafVpnIPsecProfileAuthenticationPreSharedSecret types.String `tfsdk:"pre_shared_secret" json:"pre-shared-secret,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecProfileAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication mode

|  Format  |  Description  |
|----------|---------------|
|  pre-shared-secret  |  Use a pre-shared secret key  |

`,
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

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecProfileAuthentication) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVpnIPsecProfileAuthenticationMode.IsNull() && !o.LeafVpnIPsecProfileAuthenticationMode.IsUnknown() {
		jsonData["mode"] = o.LeafVpnIPsecProfileAuthenticationMode.ValueString()
	}

	if !o.LeafVpnIPsecProfileAuthenticationPreSharedSecret.IsNull() && !o.LeafVpnIPsecProfileAuthenticationPreSharedSecret.IsUnknown() {
		jsonData["pre-shared-secret"] = o.LeafVpnIPsecProfileAuthenticationPreSharedSecret.ValueString()
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
func (o *VpnIPsecProfileAuthentication) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["mode"]; ok {
		o.LeafVpnIPsecProfileAuthenticationMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecProfileAuthenticationMode = basetypes.NewStringNull()
	}

	if value, ok := jsonData["pre-shared-secret"]; ok {
		o.LeafVpnIPsecProfileAuthenticationPreSharedSecret = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecProfileAuthenticationPreSharedSecret = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
