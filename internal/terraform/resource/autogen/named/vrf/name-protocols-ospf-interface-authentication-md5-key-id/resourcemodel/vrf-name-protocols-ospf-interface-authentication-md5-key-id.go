// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID describes the resource data model.
type VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDVrfName any `tfsdk:"name" vyos:"name,parent-id"`

	ParentIDVrfNameProtocolsOspfInterface any `tfsdk:"interface" vyos:"interface,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyIDMdfiveKey types.String `tfsdk:"md5_key" vyos:"md5-key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) GetVyosPath() []string {
	return []string{
		"vrf",
		"name",
		"protocols",
		"ospf",
		"interface",
		"authentication",
		"md5",
		"key-id",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `MD5 key id

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  MD5 key id  |

`,
		},

		// LeafNodes

		"md5_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MD5 authentication type

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  MD5 Key (16 characters or less)  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyIDMdfiveKey.IsNull() && !o.LeafVrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyIDMdfiveKey.IsUnknown() {
		jsonData["md5-key"] = o.LeafVrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyIDMdfiveKey.ValueString()
	}

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["md5-key"]; ok {
		o.LeafVrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyIDMdfiveKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyIDMdfiveKey = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
