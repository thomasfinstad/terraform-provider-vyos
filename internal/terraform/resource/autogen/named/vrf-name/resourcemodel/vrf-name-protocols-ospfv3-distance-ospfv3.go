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

// VrfNameProtocolsOspfvthreeDistanceOspfvthree describes the resource data model.
type VrfNameProtocolsOspfvthreeDistanceOspfvthree struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal  types.String `tfsdk:"external"`
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea types.String `tfsdk:"inter_area"`
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea types.String `tfsdk:"intra_area"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfvthreeDistanceOspfvthree) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "distance", "ospfv3"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal.IsNull() || o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal.IsUnknown()) {
		vyosData["external"] = o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea.IsNull() || o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea.IsUnknown()) {
		vyosData["inter-area"] = o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea.IsNull() || o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea.IsUnknown()) {
		vyosData["intra-area"] = o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfvthreeDistanceOspfvthree) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "distance", "ospfv3"}})

	// Leafs
	if value, ok := vyosData["external"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal = basetypes.NewStringNull()
	}
	if value, ok := vyosData["inter-area"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea = basetypes.NewStringNull()
	}
	if value, ok := vyosData["intra-area"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "distance", "ospfv3"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfvthreeDistanceOspfvthree) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"external":   types.StringType,
		"inter_area": types.StringType,
		"intra_area": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeDistanceOspfvthree) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance for external routes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for external routes  |

`,
		},

		"inter_area": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance for inter-area routes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for inter-area routes  |

`,
		},

		"intra_area": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance for intra-area routes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for intra-area routes  |

`,
		},

		// TagNodes

		// Nodes

	}
}
