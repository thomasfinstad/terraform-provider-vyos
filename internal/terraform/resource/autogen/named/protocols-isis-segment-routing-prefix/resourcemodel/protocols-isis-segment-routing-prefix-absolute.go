// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsIsisSegmentRoutingPrefixAbsolute describes the resource data model.
type ProtocolsIsisSegmentRoutingPrefixAbsolute struct {
	// LeafNodes
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue        types.String `tfsdk:"value" json:"value,omitempty"`
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull types.String `tfsdk:"explicit_null" json:"explicit-null,omitempty"`
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag    types.String `tfsdk:"no_php_flag" json:"no-php-flag,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsisSegmentRoutingPrefixAbsolute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify the absolute value of prefix segment/label ID

|  Format  |  Description  |
|----------|---------------|
|  u32:16-1048575  |  The absolute segment/label ID value  |

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
func (o *ProtocolsIsisSegmentRoutingPrefixAbsolute) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue.IsNull() && !o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue.IsUnknown() {
		jsonData["value"] = o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue.ValueString()
	}

	if !o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull.IsNull() && !o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull.IsUnknown() {
		jsonData["explicit-null"] = o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull.ValueString()
	}

	if !o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag.IsNull() && !o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag.IsUnknown() {
		jsonData["no-php-flag"] = o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag.ValueString()
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
func (o *ProtocolsIsisSegmentRoutingPrefixAbsolute) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["value"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue = basetypes.NewStringNull()
	}

	if value, ok := jsonData["explicit-null"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-php-flag"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
