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

// VrfNameProtocolsStaticRouteInterface describes the resource data model.
type VrfNameProtocolsStaticRouteInterface struct {
	// LeafNodes
	LeafVrfNameProtocolsStaticRouteInterfaceDisable  types.String `tfsdk:"disable"`
	LeafVrfNameProtocolsStaticRouteInterfaceDistance types.String `tfsdk:"distance"`
	LeafVrfNameProtocolsStaticRouteInterfaceVrf      types.String `tfsdk:"vrf"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsStaticRouteInterface) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "static", "route", "interface"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsStaticRouteInterfaceDisable.IsNull() || o.LeafVrfNameProtocolsStaticRouteInterfaceDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafVrfNameProtocolsStaticRouteInterfaceDisable.ValueString()
	}
	if !(o.LeafVrfNameProtocolsStaticRouteInterfaceDistance.IsNull() || o.LeafVrfNameProtocolsStaticRouteInterfaceDistance.IsUnknown()) {
		vyosData["distance"] = o.LeafVrfNameProtocolsStaticRouteInterfaceDistance.ValueString()
	}
	if !(o.LeafVrfNameProtocolsStaticRouteInterfaceVrf.IsNull() || o.LeafVrfNameProtocolsStaticRouteInterfaceVrf.IsUnknown()) {
		vyosData["vrf"] = o.LeafVrfNameProtocolsStaticRouteInterfaceVrf.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsStaticRouteInterface) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "static", "route", "interface"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafVrfNameProtocolsStaticRouteInterfaceDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsStaticRouteInterfaceDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["distance"]; ok {
		o.LeafVrfNameProtocolsStaticRouteInterfaceDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsStaticRouteInterfaceDistance = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vrf"]; ok {
		o.LeafVrfNameProtocolsStaticRouteInterfaceVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsStaticRouteInterfaceVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "static", "route", "interface"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsStaticRouteInterface) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":  types.StringType,
		"distance": types.StringType,
		"vrf":      types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRouteInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
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