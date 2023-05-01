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

// ProtocolsIsisSegmentRoutingPrefixIndex describes the resource data model.
type ProtocolsIsisSegmentRoutingPrefixIndex struct {
	// LeafNodes
	LeafProtocolsIsisSegmentRoutingPrefixIndexValue        types.String `tfsdk:"value"`
	LeafProtocolsIsisSegmentRoutingPrefixIndexExplicitNull types.String `tfsdk:"explicit_null"`
	LeafProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag    types.String `tfsdk:"no_php_flag"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsIsisSegmentRoutingPrefixIndex) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "isis", "segment-routing", "prefix", "index"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsIsisSegmentRoutingPrefixIndexValue.IsNull() || o.LeafProtocolsIsisSegmentRoutingPrefixIndexValue.IsUnknown()) {
		vyosData["value"] = o.LeafProtocolsIsisSegmentRoutingPrefixIndexValue.ValueString()
	}
	if !(o.LeafProtocolsIsisSegmentRoutingPrefixIndexExplicitNull.IsNull() || o.LeafProtocolsIsisSegmentRoutingPrefixIndexExplicitNull.IsUnknown()) {
		vyosData["explicit-null"] = o.LeafProtocolsIsisSegmentRoutingPrefixIndexExplicitNull.ValueString()
	}
	if !(o.LeafProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag.IsNull() || o.LeafProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag.IsUnknown()) {
		vyosData["no-php-flag"] = o.LeafProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsIsisSegmentRoutingPrefixIndex) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "isis", "segment-routing", "prefix", "index"}})

	// Leafs
	if value, ok := vyosData["value"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixIndexValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixIndexValue = basetypes.NewStringNull()
	}
	if value, ok := vyosData["explicit-null"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixIndexExplicitNull = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixIndexExplicitNull = basetypes.NewStringNull()
	}
	if value, ok := vyosData["no-php-flag"]; ok {
		o.LeafProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "isis", "segment-routing", "prefix", "index"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsIsisSegmentRoutingPrefixIndex) AttributeTypes() map[string]attr.Type {
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
func (o ProtocolsIsisSegmentRoutingPrefixIndex) ResourceSchemaAttributes() map[string]schema.Attribute {
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
