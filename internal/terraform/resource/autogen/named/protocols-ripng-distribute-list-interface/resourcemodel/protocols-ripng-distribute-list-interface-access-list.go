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

// ProtocolsRIPngDistributeListInterfaceAccessList describes the resource data model.
type ProtocolsRIPngDistributeListInterfaceAccessList struct {
	// LeafNodes
	LeafProtocolsRIPngDistributeListInterfaceAccessListIn  types.String `tfsdk:"in"`
	LeafProtocolsRIPngDistributeListInterfaceAccessListOut types.String `tfsdk:"out"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsRIPngDistributeListInterfaceAccessList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "ripng", "distribute-list", "interface", "access-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsRIPngDistributeListInterfaceAccessListIn.IsNull() || o.LeafProtocolsRIPngDistributeListInterfaceAccessListIn.IsUnknown()) {
		vyosData["in"] = o.LeafProtocolsRIPngDistributeListInterfaceAccessListIn.ValueString()
	}
	if !(o.LeafProtocolsRIPngDistributeListInterfaceAccessListOut.IsNull() || o.LeafProtocolsRIPngDistributeListInterfaceAccessListOut.IsUnknown()) {
		vyosData["out"] = o.LeafProtocolsRIPngDistributeListInterfaceAccessListOut.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsRIPngDistributeListInterfaceAccessList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "ripng", "distribute-list", "interface", "access-list"}})

	// Leafs
	if value, ok := vyosData["in"]; ok {
		o.LeafProtocolsRIPngDistributeListInterfaceAccessListIn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPngDistributeListInterfaceAccessListIn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["out"]; ok {
		o.LeafProtocolsRIPngDistributeListInterfaceAccessListOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPngDistributeListInterfaceAccessListOut = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "ripng", "distribute-list", "interface", "access-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsRIPngDistributeListInterfaceAccessList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"in":  types.StringType,
		"out": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngDistributeListInterfaceAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"in": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to input packets

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access list to apply to input packets  |

`,
		},

		"out": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to output packets

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access list to apply to output packets  |

`,
		},

		// TagNodes

		// Nodes

	}
}