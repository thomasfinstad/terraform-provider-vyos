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

// ProtocolsNhrpTunnelMap describes the resource data model.
type ProtocolsNhrpTunnelMap struct {
	// LeafNodes
	LeafProtocolsNhrpTunnelMapCisco       types.String `tfsdk:"cisco"`
	LeafProtocolsNhrpTunnelMapNbmaAddress types.String `tfsdk:"nbma_address"`
	LeafProtocolsNhrpTunnelMapRegister    types.String `tfsdk:"register"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsNhrpTunnelMap) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "nhrp", "tunnel", "map"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsNhrpTunnelMapCisco.IsNull() || o.LeafProtocolsNhrpTunnelMapCisco.IsUnknown()) {
		vyosData["cisco"] = o.LeafProtocolsNhrpTunnelMapCisco.ValueString()
	}
	if !(o.LeafProtocolsNhrpTunnelMapNbmaAddress.IsNull() || o.LeafProtocolsNhrpTunnelMapNbmaAddress.IsUnknown()) {
		vyosData["nbma-address"] = o.LeafProtocolsNhrpTunnelMapNbmaAddress.ValueString()
	}
	if !(o.LeafProtocolsNhrpTunnelMapRegister.IsNull() || o.LeafProtocolsNhrpTunnelMapRegister.IsUnknown()) {
		vyosData["register"] = o.LeafProtocolsNhrpTunnelMapRegister.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsNhrpTunnelMap) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "nhrp", "tunnel", "map"}})

	// Leafs
	if value, ok := vyosData["cisco"]; ok {
		o.LeafProtocolsNhrpTunnelMapCisco = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsNhrpTunnelMapCisco = basetypes.NewStringNull()
	}
	if value, ok := vyosData["nbma-address"]; ok {
		o.LeafProtocolsNhrpTunnelMapNbmaAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsNhrpTunnelMapNbmaAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["register"]; ok {
		o.LeafProtocolsNhrpTunnelMapRegister = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsNhrpTunnelMapRegister = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "nhrp", "tunnel", "map"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsNhrpTunnelMap) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"cisco":        types.StringType,
		"nbma_address": types.StringType,
		"register":     types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsNhrpTunnelMap) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cisco": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `If the statically mapped peer is running Cisco IOS, specify this

`,
		},

		"nbma_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set HUB address (nbma-address - external hub address or fqdn)

`,
		},

		"register": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies that Registration Request should be sent to this peer on startup

`,
		},

		// TagNodes

		// Nodes

	}
}
