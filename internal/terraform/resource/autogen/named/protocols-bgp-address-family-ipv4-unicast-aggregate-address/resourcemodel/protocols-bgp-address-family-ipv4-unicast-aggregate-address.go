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

// ProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress describes the resource data model.
type ProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressAsSet       types.String `tfsdk:"as_set"`
	LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressRouteMap    types.String `tfsdk:"route_map"`
	LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressSummaryOnly types.String `tfsdk:"summary_only"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress) GetVyosPath() []string {
	return []string{
		"protocols",
		"bgp",
		"address-family",
		"ipv4-unicast",
		"aggregate-address",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv4-unicast", "aggregate-address"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressAsSet.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressAsSet.IsUnknown()) {
		vyosData["as-set"] = o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressAsSet.ValueString()
	}
	if !(o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressRouteMap.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressRouteMap.ValueString()
	}
	if !(o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressSummaryOnly.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressSummaryOnly.IsUnknown()) {
		vyosData["summary-only"] = o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressSummaryOnly.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv4-unicast", "aggregate-address"}})

	// Leafs
	if value, ok := vyosData["as-set"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressAsSet = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressAsSet = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressRouteMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["summary-only"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressSummaryOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddressSummaryOnly = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv4-unicast", "aggregate-address"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress) AttributeTypes() map[string]attr.Type {
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
func (o ProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP aggregate network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  BGP aggregate network  |

`,
		},

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