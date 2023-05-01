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

// ProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance types.String `tfsdk:"distance"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) GetVyosPath() []string {
	return []string{
		"protocols",
		"bgp",
		"address-family",
		"ipv6-unicast",
		"distance",
		"prefix",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-unicast", "distance", "prefix"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance.IsNull() || o.LeafProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance.IsUnknown()) {
		vyosData["distance"] = o.LeafProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-unicast", "distance", "prefix"}})

	// Leafs
	if value, ok := vyosData["distance"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefixDistance = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "address-family", "ipv6-unicast", "distance", "prefix"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"distance": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Administrative distance for a specific BGP prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Administrative distance for a specific BGP prefix  |

`,
		},

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