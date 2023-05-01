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

// ProtocolsRIPngInterfaceSplitHorizon describes the resource data model.
type ProtocolsRIPngInterfaceSplitHorizon struct {
	// LeafNodes
	LeafProtocolsRIPngInterfaceSplitHorizonDisable       types.String `tfsdk:"disable"`
	LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse types.String `tfsdk:"poison_reverse"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsRIPngInterfaceSplitHorizon) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "ripng", "interface", "split-horizon"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsRIPngInterfaceSplitHorizonDisable.IsNull() || o.LeafProtocolsRIPngInterfaceSplitHorizonDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafProtocolsRIPngInterfaceSplitHorizonDisable.ValueString()
	}
	if !(o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse.IsNull() || o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse.IsUnknown()) {
		vyosData["poison-reverse"] = o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsRIPngInterfaceSplitHorizon) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "ripng", "interface", "split-horizon"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafProtocolsRIPngInterfaceSplitHorizonDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPngInterfaceSplitHorizonDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["poison-reverse"]; ok {
		o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "ripng", "interface", "split-horizon"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsRIPngInterfaceSplitHorizon) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":        types.StringType,
		"poison_reverse": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngInterfaceSplitHorizon) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable split horizon on specified interface

`,
		},

		"poison_reverse": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable split horizon on specified interface

`,
		},

		// TagNodes

		// Nodes

	}
}
