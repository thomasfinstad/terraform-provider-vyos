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

// VrfNameProtocolsStaticRoutesixReject describes the resource data model.
type VrfNameProtocolsStaticRoutesixReject struct {
	// LeafNodes
	LeafVrfNameProtocolsStaticRoutesixRejectDistance types.String `tfsdk:"distance"`
	LeafVrfNameProtocolsStaticRoutesixRejectTag      types.String `tfsdk:"tag"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsStaticRoutesixReject) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "static", "route6", "reject"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsStaticRoutesixRejectDistance.IsNull() || o.LeafVrfNameProtocolsStaticRoutesixRejectDistance.IsUnknown()) {
		vyosData["distance"] = o.LeafVrfNameProtocolsStaticRoutesixRejectDistance.ValueString()
	}
	if !(o.LeafVrfNameProtocolsStaticRoutesixRejectTag.IsNull() || o.LeafVrfNameProtocolsStaticRoutesixRejectTag.IsUnknown()) {
		vyosData["tag"] = o.LeafVrfNameProtocolsStaticRoutesixRejectTag.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsStaticRoutesixReject) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "static", "route6", "reject"}})

	// Leafs
	if value, ok := vyosData["distance"]; ok {
		o.LeafVrfNameProtocolsStaticRoutesixRejectDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsStaticRoutesixRejectDistance = basetypes.NewStringNull()
	}
	if value, ok := vyosData["tag"]; ok {
		o.LeafVrfNameProtocolsStaticRoutesixRejectTag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsStaticRoutesixRejectTag = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "static", "route6", "reject"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsStaticRoutesixReject) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"distance": types.StringType,
		"tag":      types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoutesixReject) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"distance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for this route  |

`,
		},

		"tag": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Tag value for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Tag value for this route  |

`,
		},

		// TagNodes

		// Nodes

	}
}
