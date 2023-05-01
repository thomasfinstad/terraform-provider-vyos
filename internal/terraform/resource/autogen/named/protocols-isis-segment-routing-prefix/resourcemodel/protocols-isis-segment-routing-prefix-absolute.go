// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ProtocolsIsisSegmentRoutingPrefixAbsolute describes the resource data model.
type ProtocolsIsisSegmentRoutingPrefixAbsolute struct {
	// LeafNodes
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue        types.String `tfsdk:"value"`
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull types.String `tfsdk:"explicit_null"`
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag    types.String `tfsdk:"no_php_flag"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsIsisSegmentRoutingPrefixAbsolute) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "isis", "segment-routing", "prefix", "absolute"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue.IsNull() || o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue.IsUnknown()) {
		vyosData["value"] = o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue.ValueString()
	}
	if !(o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull.IsNull() || o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull.IsUnknown()) {
		vyosData["explicit-null"] = o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull.ValueString()
	}
	if !(o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag.IsNull() || o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag.IsUnknown()) {
		vyosData["no-php-flag"] = o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsIsisSegmentRoutingPrefixAbsolute) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "isis", "segment-routing", "prefix", "absolute"}})

	// Leafs
	if value, ok := vyosData["value"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue = basetypes.NewStringNull()
	}
	if value, ok := vyosData["explicit-null"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull = basetypes.NewStringNull()
	}
	if value, ok := vyosData["no-php-flag"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "isis", "segment-routing", "prefix", "absolute"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsIsisSegmentRoutingPrefixAbsolute) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"value":         types.StringType,
		"explicit_null": types.StringType,
		"no_php_flag":   types.StringType,

		// Tags

		// Nodes

	}
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
