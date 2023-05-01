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

// VrfNameProtocolsBgpParametersDefault describes the resource data model.
type VrfNameProtocolsBgpParametersDefault struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersDefaultLocalPref types.String `tfsdk:"local_pref"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpParametersDefault) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "parameters", "default"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpParametersDefaultLocalPref.IsNull() || o.LeafVrfNameProtocolsBgpParametersDefaultLocalPref.IsUnknown()) {
		vyosData["local-pref"] = o.LeafVrfNameProtocolsBgpParametersDefaultLocalPref.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpParametersDefault) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "parameters", "default"}})

	// Leafs
	if value, ok := vyosData["local-pref"]; ok {
		o.LeafVrfNameProtocolsBgpParametersDefaultLocalPref = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpParametersDefaultLocalPref = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "parameters", "default"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpParametersDefault) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"local_pref": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_pref": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default local preference

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Local preference  |

`,
		},

		// TagNodes

		// Nodes

	}
}