// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID describes the resource data model.
type ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDProtocolsOspfArea any `tfsdk:"area" vyos:"area,parent-id"`

	ParentIDProtocolsOspfAreaVirtualLink any `tfsdk:"virtual_link" vyos:"virtual-link,parent-id"`

	// LeafNodes
	LeafProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey types.String `tfsdk:"md5_key" vyos:"md5-key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) GetVyosPath() []string {
	return []string{
		"protocols",
		"ospf",
		"area",
		"virtual-link",
		"authentication",
		"md5",
		"key-id",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey.IsNull() && !o.LeafProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey.IsUnknown() {
		jsonData["md5-key"] = o.LeafProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey.ValueString()
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
func (o *ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["md5-key"]; ok {
		o.LeafProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
