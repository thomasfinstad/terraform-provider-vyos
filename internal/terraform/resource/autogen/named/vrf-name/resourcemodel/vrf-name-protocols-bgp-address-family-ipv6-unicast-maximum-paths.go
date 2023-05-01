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

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsEbgp types.String `tfsdk:"ebgp"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsIbgp types.String `tfsdk:"ibgp"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "maximum-paths"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsEbgp.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsEbgp.IsUnknown()) {
		vyosData["ebgp"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsEbgp.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsIbgp.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsIbgp.IsUnknown()) {
		vyosData["ibgp"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsIbgp.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "maximum-paths"}})

	// Leafs
	if value, ok := vyosData["ebgp"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsEbgp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsEbgp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ibgp"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsIbgp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsIbgp = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "maximum-paths"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"ebgp": types.StringType,
		"ibgp": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ebgp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `eBGP maximum paths

|  Format  |  Description  |
|----------|---------------|
|  u32:1-256  |  Number of paths to consider  |

`,
		},

		"ibgp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `iBGP maximum paths

|  Format  |  Description  |
|----------|---------------|
|  u32:1-256  |  Number of paths to consider  |

`,
		},

		// TagNodes

		// Nodes

	}
}