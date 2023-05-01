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

// VrfNameProtocolsOspfvthreeRedistributeStatic describes the resource data model.
type VrfNameProtocolsOspfvthreeRedistributeStatic struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeRedistributeStaticRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfvthreeRedistributeStatic) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "redistribute", "static"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfvthreeRedistributeStaticRouteMap.IsNull() || o.LeafVrfNameProtocolsOspfvthreeRedistributeStaticRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsOspfvthreeRedistributeStaticRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfvthreeRedistributeStatic) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "redistribute", "static"}})

	// Leafs
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeRedistributeStaticRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeRedistributeStaticRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "redistribute", "static"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfvthreeRedistributeStatic) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeRedistributeStatic) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		// Nodes

	}
}