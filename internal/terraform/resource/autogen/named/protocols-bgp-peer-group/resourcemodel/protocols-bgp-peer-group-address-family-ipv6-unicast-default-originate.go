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

// ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginateRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "default-originate"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginateRouteMap.IsNull() || o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginateRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginateRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "default-originate"}})

	// Leafs
	if value, ok := vyosData["route-map"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginateRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginateRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "default-originate"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate) ResourceSchemaAttributes() map[string]schema.Attribute {
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
