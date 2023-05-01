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

// VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneMetric   types.String `tfsdk:"metric"`
	LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv4", "rip", "level-1"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneMetric.IsNull() || o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneMetric.IsUnknown()) {
		vyosData["metric"] = o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneMetric.ValueString()
	}
	if !(o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneRouteMap.IsNull() || o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv4", "rip", "level-1"}})

	// Leafs
	if value, ok := vyosData["metric"]; ok {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneMetric = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourRIPLevelOneRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv4", "rip", "level-1"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"metric":    types.StringType,
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set default metric for circuit

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Default metric value  |

`,
		},

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