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

// VrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetwork describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetwork struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkBackdoor types.String `tfsdk:"backdoor"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetwork) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-multicast", "network"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkBackdoor.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkBackdoor.IsUnknown()) {
		vyosData["backdoor"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkBackdoor.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkRouteMap.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetwork) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-multicast", "network"}})

	// Leafs
	if value, ok := vyosData["backdoor"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkBackdoor = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkBackdoor = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetworkRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-multicast", "network"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetwork) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"backdoor":  types.StringType,
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
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