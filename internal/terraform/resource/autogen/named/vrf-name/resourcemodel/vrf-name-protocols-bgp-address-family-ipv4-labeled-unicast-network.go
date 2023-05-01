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

// VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkBackdoor types.String `tfsdk:"backdoor"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-labeled-unicast", "network"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkBackdoor.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkBackdoor.IsUnknown()) {
		vyosData["backdoor"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkBackdoor.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkRouteMap.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-labeled-unicast", "network"}})

	// Leafs
	if value, ok := vyosData["backdoor"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkBackdoor = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkBackdoor = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetworkRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-labeled-unicast", "network"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"backdoor":  types.StringType,
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"backdoor": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use BGP network/prefix as a backdoor route

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