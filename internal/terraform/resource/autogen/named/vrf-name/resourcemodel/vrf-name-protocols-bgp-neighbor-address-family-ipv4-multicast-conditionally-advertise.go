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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseAdvertiseMap types.String `tfsdk:"advertise_map"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseExistMap     types.String `tfsdk:"exist_map"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseNonExistMap  types.String `tfsdk:"non_exist_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-multicast", "conditionally-advertise"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseAdvertiseMap.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseAdvertiseMap.IsUnknown()) {
		vyosData["advertise-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseAdvertiseMap.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseExistMap.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseExistMap.IsUnknown()) {
		vyosData["exist-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseExistMap.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseNonExistMap.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseNonExistMap.IsUnknown()) {
		vyosData["non-exist-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseNonExistMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-multicast", "conditionally-advertise"}})

	// Leafs
	if value, ok := vyosData["advertise-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseAdvertiseMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseAdvertiseMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["exist-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseExistMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseExistMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["non-exist-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseNonExistMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseNonExistMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-multicast", "conditionally-advertise"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"advertise_map": types.StringType,
		"exist_map":     types.StringType,
		"non_exist_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to conditionally advertise routes

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"exist_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"non_exist_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		// Nodes

	}
}