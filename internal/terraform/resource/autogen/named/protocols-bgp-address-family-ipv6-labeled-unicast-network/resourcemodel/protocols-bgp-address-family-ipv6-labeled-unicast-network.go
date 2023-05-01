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

// ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor types.String `tfsdk:"backdoor"`
	LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetVyosPath() []string {
	return []string{
		"protocols",
		"bgp",
		"address-family",
		"ipv6-labeled-unicast",
		"network",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-labeled-unicast", "network"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor.IsUnknown()) {
		vyosData["backdoor"] = o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor.ValueString()
	}
	if !(o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-labeled-unicast", "network"}})

	// Leafs
	if value, ok := vyosData["backdoor"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-labeled-unicast", "network"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"backdoor":  types.StringType,
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Import BGP network/prefix into labeled unicast IPv6 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Labeled Unicast IPv6 BGP network/prefix  |

`,
		},

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