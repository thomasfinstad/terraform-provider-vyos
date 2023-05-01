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

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasInNumber types.String `tfsdk:"number"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "allowas-in"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasInNumber.IsNull() || o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasInNumber.IsUnknown()) {
		vyosData["number"] = o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasInNumber.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "allowas-in"}})

	// Leafs
	if value, ok := vyosData["number"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasInNumber = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasInNumber = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "allowas-in"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"number": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of times AS is allowed in path  |

`,
		},

		// TagNodes

		// Nodes

	}
}
