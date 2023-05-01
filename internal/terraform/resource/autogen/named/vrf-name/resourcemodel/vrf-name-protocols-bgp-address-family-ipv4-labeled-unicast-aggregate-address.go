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

// VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressAsSet       types.String `tfsdk:"as_set"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressRouteMap    types.String `tfsdk:"route_map"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressSummaryOnly types.String `tfsdk:"summary_only"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-labeled-unicast", "aggregate-address"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressAsSet.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressAsSet.IsUnknown()) {
		vyosData["as-set"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressAsSet.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressRouteMap.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressRouteMap.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressSummaryOnly.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressSummaryOnly.IsUnknown()) {
		vyosData["summary-only"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressSummaryOnly.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-labeled-unicast", "aggregate-address"}})

	// Leafs
	if value, ok := vyosData["as-set"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressAsSet = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressAsSet = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressRouteMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["summary-only"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressSummaryOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddressSummaryOnly = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-labeled-unicast", "aggregate-address"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"as_set":       types.StringType,
		"route_map":    types.StringType,
		"summary_only": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_set": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Generate AS-set path information for this aggregate address

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

		"summary_only": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Announce the aggregate summary network only

`,
		},

		// TagNodes

		// Nodes

	}
}