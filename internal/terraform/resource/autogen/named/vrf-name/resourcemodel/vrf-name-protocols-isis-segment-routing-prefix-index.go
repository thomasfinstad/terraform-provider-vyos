// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsIsisSegmentRoutingPrefixIndex describes the resource data model.
type VrfNameProtocolsIsisSegmentRoutingPrefixIndex struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexValue        types.String `tfsdk:"value" json:"value,omitempty"`
	LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexExplicitNull types.String `tfsdk:"explicit_null" json:"explicit-null,omitempty"`
	LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag    types.String `tfsdk:"no_php_flag" json:"no-php-flag,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRoutingPrefixIndex) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify the index value of prefix segment/label ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  The index segment/label ID value  |

`,
		},

		"explicit_null": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Request upstream neighbor to replace segment/label with explicit null label

`,
		},

		"no_php_flag": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Do not request penultimate hop popping for segment/label

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisSegmentRoutingPrefixIndex) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexValue.IsNull() && !o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexValue.IsUnknown() {
		jsonData["value"] = o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexValue.ValueString()
	}

	if !o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexExplicitNull.IsNull() && !o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexExplicitNull.IsUnknown() {
		jsonData["explicit-null"] = o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexExplicitNull.ValueString()
	}

	if !o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag.IsNull() && !o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag.IsUnknown() {
		jsonData["no-php-flag"] = o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag.ValueString()
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
func (o *VrfNameProtocolsIsisSegmentRoutingPrefixIndex) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["value"]; ok {
		o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexValue = basetypes.NewStringNull()
	}

	if value, ok := jsonData["explicit-null"]; ok {
		o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexExplicitNull = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexExplicitNull = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-php-flag"]; ok {
		o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
