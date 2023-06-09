// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PkiOpenvpnSharedSecret describes the resource data model.
type PkiOpenvpnSharedSecret struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafPkiOpenvpnSharedSecretKey     types.String `tfsdk:"key" json:"key,omitempty"`
	LeafPkiOpenvpnSharedSecretVersion types.String `tfsdk:"version" json:"version,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiOpenvpnSharedSecret) GetVyosPath() []string {
	return []string{
		"pki",
		"openvpn",
		"shared-secret",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiOpenvpnSharedSecret) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OpenVPN shared secret key

`,
		},

		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OpenVPN shared secret key data

`,
		},

		"version": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OpenVPN shared secret key version

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PkiOpenvpnSharedSecret) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPkiOpenvpnSharedSecretKey.IsNull() && !o.LeafPkiOpenvpnSharedSecretKey.IsUnknown() {
		jsonData["key"] = o.LeafPkiOpenvpnSharedSecretKey.ValueString()
	}

	if !o.LeafPkiOpenvpnSharedSecretVersion.IsNull() && !o.LeafPkiOpenvpnSharedSecretVersion.IsUnknown() {
		jsonData["version"] = o.LeafPkiOpenvpnSharedSecretVersion.ValueString()
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
func (o *PkiOpenvpnSharedSecret) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["key"]; ok {
		o.LeafPkiOpenvpnSharedSecretKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiOpenvpnSharedSecretKey = basetypes.NewStringNull()
	}

	if value, ok := jsonData["version"]; ok {
		o.LeafPkiOpenvpnSharedSecretVersion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiOpenvpnSharedSecretVersion = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
