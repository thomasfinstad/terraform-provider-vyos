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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound types.String `tfsdk:"inbound"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "soft-reconfiguration"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound.IsUnknown()) {
		vyosData["inbound"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "soft-reconfiguration"}})

	// Leafs
	if value, ok := vyosData["inbound"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "soft-reconfiguration"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"inbound": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
		},

		// TagNodes

		// Nodes

	}
}
