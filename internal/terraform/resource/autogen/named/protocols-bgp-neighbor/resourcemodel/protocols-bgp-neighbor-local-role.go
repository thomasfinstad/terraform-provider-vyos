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

// ProtocolsBgpNeighborLocalRole describes the resource data model.
type ProtocolsBgpNeighborLocalRole struct {
	// LeafNodes
	LeafProtocolsBgpNeighborLocalRoleStrict types.String `tfsdk:"strict"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborLocalRole) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "local-role"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborLocalRoleStrict.IsNull() || o.LeafProtocolsBgpNeighborLocalRoleStrict.IsUnknown()) {
		vyosData["strict"] = o.LeafProtocolsBgpNeighborLocalRoleStrict.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborLocalRole) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "local-role"}})

	// Leafs
	if value, ok := vyosData["strict"]; ok {
		o.LeafProtocolsBgpNeighborLocalRoleStrict = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborLocalRoleStrict = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "local-role"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborLocalRole) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"strict": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborLocalRole) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"strict": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor must send this exact capability, otherwise a role missmatch notification will be sent

`,
		},

		// TagNodes

		// Nodes

	}
}
