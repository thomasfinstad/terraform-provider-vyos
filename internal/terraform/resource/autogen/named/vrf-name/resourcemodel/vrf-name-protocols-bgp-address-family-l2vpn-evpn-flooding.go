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

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable            types.String `tfsdk:"disable"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication types.String `tfsdk:"head_end_replication"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "l2vpn-evpn", "flooding"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication.IsUnknown()) {
		vyosData["head-end-replication"] = o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "l2vpn-evpn", "flooding"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["head-end-replication"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "l2vpn-evpn", "flooding"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":              types.StringType,
		"head_end_replication": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Do not flood any BUM packets

`,
		},

		"head_end_replication": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Flood BUM packets using head-end replication

`,
		},

		// TagNodes

		// Nodes

	}
}