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

// VrfNameProtocolsOspfvthreeRedistributeBgp describes the resource data model.
type VrfNameProtocolsOspfvthreeRedistributeBgp struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeRedistributeBgpRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfvthreeRedistributeBgp) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "redistribute", "bgp"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfvthreeRedistributeBgpRouteMap.IsNull() || o.LeafVrfNameProtocolsOspfvthreeRedistributeBgpRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsOspfvthreeRedistributeBgpRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfvthreeRedistributeBgp) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "redistribute", "bgp"}})

	// Leafs
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeRedistributeBgpRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeRedistributeBgpRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "redistribute", "bgp"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfvthreeRedistributeBgp) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeRedistributeBgp) ResourceSchemaAttributes() map[string]schema.Attribute {
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
