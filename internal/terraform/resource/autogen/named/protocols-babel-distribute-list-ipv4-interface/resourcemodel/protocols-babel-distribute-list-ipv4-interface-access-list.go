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

// ProtocolsBabelDistributeListIPvfourInterfaceAccessList describes the resource data model.
type ProtocolsBabelDistributeListIPvfourInterfaceAccessList struct {
	// LeafNodes
	LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListIn  types.String `tfsdk:"in"`
	LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListOut types.String `tfsdk:"out"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBabelDistributeListIPvfourInterfaceAccessList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "babel", "distribute-list", "ipv4", "interface", "access-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListIn.IsNull() || o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListIn.IsUnknown()) {
		vyosData["in"] = o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListIn.ValueString()
	}
	if !(o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListOut.IsNull() || o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListOut.IsUnknown()) {
		vyosData["out"] = o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListOut.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBabelDistributeListIPvfourInterfaceAccessList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "babel", "distribute-list", "ipv4", "interface", "access-list"}})

	// Leafs
	if value, ok := vyosData["in"]; ok {
		o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListIn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListIn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["out"]; ok {
		o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelDistributeListIPvfourInterfaceAccessListOut = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "babel", "distribute-list", "ipv4", "interface", "access-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBabelDistributeListIPvfourInterfaceAccessList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"in":  types.StringType,
		"out": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelDistributeListIPvfourInterfaceAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
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
