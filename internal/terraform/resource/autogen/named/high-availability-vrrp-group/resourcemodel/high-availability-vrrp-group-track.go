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

// HighAvailabilityVrrpGroupTrack describes the resource data model.
type HighAvailabilityVrrpGroupTrack struct {
	// LeafNodes
	LeafHighAvailabilityVrrpGroupTrackExcludeVrrpInterface types.String `tfsdk:"exclude_vrrp_interface"`
	LeafHighAvailabilityVrrpGroupTrackInterface            types.String `tfsdk:"interface"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *HighAvailabilityVrrpGroupTrack) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"high-availability", "vrrp", "group", "track"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafHighAvailabilityVrrpGroupTrackExcludeVrrpInterface.IsNull() || o.LeafHighAvailabilityVrrpGroupTrackExcludeVrrpInterface.IsUnknown()) {
		vyosData["exclude-vrrp-interface"] = o.LeafHighAvailabilityVrrpGroupTrackExcludeVrrpInterface.ValueString()
	}
	if !(o.LeafHighAvailabilityVrrpGroupTrackInterface.IsNull() || o.LeafHighAvailabilityVrrpGroupTrackInterface.IsUnknown()) {
		vyosData["interface"] = o.LeafHighAvailabilityVrrpGroupTrackInterface.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *HighAvailabilityVrrpGroupTrack) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"high-availability", "vrrp", "group", "track"}})

	// Leafs
	if value, ok := vyosData["exclude-vrrp-interface"]; ok {
		o.LeafHighAvailabilityVrrpGroupTrackExcludeVrrpInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVrrpGroupTrackExcludeVrrpInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interface"]; ok {
		o.LeafHighAvailabilityVrrpGroupTrackInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVrrpGroupTrackInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"high-availability", "vrrp", "group", "track"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o HighAvailabilityVrrpGroupTrack) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"exclude_vrrp_interface": types.StringType,
		"interface":              types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupTrack) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"exclude_vrrp_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable track state of main interface

`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface name state check

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		// TagNodes

		// Nodes

	}
}