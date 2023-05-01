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

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance types.String `tfsdk:"distance"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "distance", "prefix"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance.IsUnknown()) {
		vyosData["distance"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "distance", "prefix"}})

	// Leafs
	if value, ok := vyosData["distance"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "distance", "prefix"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"distance": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"distance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance for prefix

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Administrative distance for external BGP routes  |

`,
		},

		// TagNodes

		// Nodes

	}
}
