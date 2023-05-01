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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive types.String `tfsdk:"receive"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend    types.String `tfsdk:"send"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "capability", "orf", "prefix-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive.IsUnknown()) {
		vyosData["receive"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend.IsUnknown()) {
		vyosData["send"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "capability", "orf", "prefix-list"}})

	// Leafs
	if value, ok := vyosData["receive"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive = basetypes.NewStringNull()
	}
	if value, ok := vyosData["send"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "capability", "orf", "prefix-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"receive": types.StringType,
		"send":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"receive": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Capability to receive the ORF

`,
		},

		"send": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Capability to send the ORF

`,
		},

		// TagNodes

		// Nodes

	}
}