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

// ProtocolsStaticRoutesixNextHop describes the resource data model.
type ProtocolsStaticRoutesixNextHop struct {
	// LeafNodes
	LeafProtocolsStaticRoutesixNextHopDisable   types.String `tfsdk:"disable"`
	LeafProtocolsStaticRoutesixNextHopDistance  types.String `tfsdk:"distance"`
	LeafProtocolsStaticRoutesixNextHopInterface types.String `tfsdk:"interface"`
	LeafProtocolsStaticRoutesixNextHopVrf       types.String `tfsdk:"vrf"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsStaticRoutesixNextHop) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "static", "route6", "next-hop"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsStaticRoutesixNextHopDisable.IsNull() || o.LeafProtocolsStaticRoutesixNextHopDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafProtocolsStaticRoutesixNextHopDisable.ValueString()
	}
	if !(o.LeafProtocolsStaticRoutesixNextHopDistance.IsNull() || o.LeafProtocolsStaticRoutesixNextHopDistance.IsUnknown()) {
		vyosData["distance"] = o.LeafProtocolsStaticRoutesixNextHopDistance.ValueString()
	}
	if !(o.LeafProtocolsStaticRoutesixNextHopInterface.IsNull() || o.LeafProtocolsStaticRoutesixNextHopInterface.IsUnknown()) {
		vyosData["interface"] = o.LeafProtocolsStaticRoutesixNextHopInterface.ValueString()
	}
	if !(o.LeafProtocolsStaticRoutesixNextHopVrf.IsNull() || o.LeafProtocolsStaticRoutesixNextHopVrf.IsUnknown()) {
		vyosData["vrf"] = o.LeafProtocolsStaticRoutesixNextHopVrf.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsStaticRoutesixNextHop) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "static", "route6", "next-hop"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafProtocolsStaticRoutesixNextHopDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticRoutesixNextHopDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["distance"]; ok {
		o.LeafProtocolsStaticRoutesixNextHopDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticRoutesixNextHopDistance = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interface"]; ok {
		o.LeafProtocolsStaticRoutesixNextHopInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticRoutesixNextHopInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vrf"]; ok {
		o.LeafProtocolsStaticRoutesixNextHopVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticRoutesixNextHopVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "static", "route6", "next-hop"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsStaticRoutesixNextHop) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":   types.StringType,
		"distance":  types.StringType,
		"interface": types.StringType,
		"vrf":       types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRoutesixNextHop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"distance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for this route  |

`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway interface name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Gateway interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to leak route

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of VRF to leak to  |

`,
		},

		// TagNodes

		// Nodes

	}
}
