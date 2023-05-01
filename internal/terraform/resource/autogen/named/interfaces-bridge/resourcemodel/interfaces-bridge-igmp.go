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

// InterfacesBrIDgeIgmp describes the resource data model.
type InterfacesBrIDgeIgmp struct {
	// LeafNodes
	LeafInterfacesBrIDgeIgmpQuerier  types.String `tfsdk:"querier"`
	LeafInterfacesBrIDgeIgmpSnooping types.String `tfsdk:"snooping"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBrIDgeIgmp) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bridge", "igmp"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBrIDgeIgmpQuerier.IsNull() || o.LeafInterfacesBrIDgeIgmpQuerier.IsUnknown()) {
		vyosData["querier"] = o.LeafInterfacesBrIDgeIgmpQuerier.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeIgmpSnooping.IsNull() || o.LeafInterfacesBrIDgeIgmpSnooping.IsUnknown()) {
		vyosData["snooping"] = o.LeafInterfacesBrIDgeIgmpSnooping.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBrIDgeIgmp) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bridge", "igmp"}})

	// Leafs
	if value, ok := vyosData["querier"]; ok {
		o.LeafInterfacesBrIDgeIgmpQuerier = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeIgmpQuerier = basetypes.NewStringNull()
	}
	if value, ok := vyosData["snooping"]; ok {
		o.LeafInterfacesBrIDgeIgmpSnooping = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeIgmpSnooping = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bridge", "igmp"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBrIDgeIgmp) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"querier":  types.StringType,
		"snooping": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeIgmp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"querier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable IGMP/MLD querier

`,
		},

		"snooping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable IGMP/MLD snooping

`,
		},

		// TagNodes

		// Nodes

	}
}
