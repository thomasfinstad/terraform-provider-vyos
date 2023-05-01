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

// VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneAlways   types.String `tfsdk:"always"`
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneMetric   types.String `tfsdk:"metric"`
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "default-information", "originate", "ipv6", "level-1"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneAlways.IsNull() || o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneAlways.IsUnknown()) {
		vyosData["always"] = o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneAlways.ValueString()
	}
	if !(o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneMetric.IsNull() || o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneMetric.IsUnknown()) {
		vyosData["metric"] = o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneMetric.ValueString()
	}
	if !(o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneRouteMap.IsNull() || o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "default-information", "originate", "ipv6", "level-1"}})

	// Leafs
	if value, ok := vyosData["always"]; ok {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneAlways = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneAlways = basetypes.NewStringNull()
	}
	if value, ok := vyosData["metric"]; ok {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneMetric = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOneRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "default-information", "originate", "ipv6", "level-1"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"always":    types.StringType,
		"metric":    types.StringType,
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"always": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Always advertise default route

`,
		},

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