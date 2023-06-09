// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemStaticHostMappingHostName describes the resource data model.
type SystemStaticHostMappingHostName struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafSystemStaticHostMappingHostNameAlias types.String `tfsdk:"alias" json:"alias,omitempty"`
	LeafSystemStaticHostMappingHostNameInet  types.String `tfsdk:"inet" json:"inet,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemStaticHostMappingHostName) GetVyosPath() []string {
	return []string{
		"system",
		"static-host-mapping",
		"host-name",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemStaticHostMappingHostName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Host name for static address mapping

`,
		},

		// LeafNodes

		"alias": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Alias for this address

`,
		},

		"inet": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP Address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
|  ipv6  |  IPv6 address  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemStaticHostMappingHostName) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafSystemStaticHostMappingHostNameAlias.IsNull() && !o.LeafSystemStaticHostMappingHostNameAlias.IsUnknown() {
		jsonData["alias"] = o.LeafSystemStaticHostMappingHostNameAlias.ValueString()
	}

	if !o.LeafSystemStaticHostMappingHostNameInet.IsNull() && !o.LeafSystemStaticHostMappingHostNameInet.IsUnknown() {
		jsonData["inet"] = o.LeafSystemStaticHostMappingHostNameInet.ValueString()
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
func (o *SystemStaticHostMappingHostName) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["alias"]; ok {
		o.LeafSystemStaticHostMappingHostNameAlias = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemStaticHostMappingHostNameAlias = basetypes.NewStringNull()
	}

	if value, ok := jsonData["inet"]; ok {
		o.LeafSystemStaticHostMappingHostNameInet = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemStaticHostMappingHostNameInet = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
