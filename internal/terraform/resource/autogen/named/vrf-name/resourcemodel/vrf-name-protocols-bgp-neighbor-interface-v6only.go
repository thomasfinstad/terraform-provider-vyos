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

// VrfNameProtocolsBgpNeighborInterfaceVsixonly describes the resource data model.
type VrfNameProtocolsBgpNeighborInterfaceVsixonly struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyPeerGroup types.String `tfsdk:"peer_group"`
	LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyRemoteAs  types.String `tfsdk:"remote_as"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborInterfaceVsixonly) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "interface", "v6only"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyPeerGroup.IsNull() || o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyPeerGroup.IsUnknown()) {
		vyosData["peer-group"] = o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyPeerGroup.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyRemoteAs.IsNull() || o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyRemoteAs.IsUnknown()) {
		vyosData["remote-as"] = o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyRemoteAs.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborInterfaceVsixonly) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "interface", "v6only"}})

	// Leafs
	if value, ok := vyosData["peer-group"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyPeerGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyPeerGroup = basetypes.NewStringNull()
	}
	if value, ok := vyosData["remote-as"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyRemoteAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborInterfaceVsixonlyRemoteAs = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "interface", "v6only"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborInterfaceVsixonly) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"peer_group": types.StringType,
		"remote_as":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborInterfaceVsixonly) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer group for this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Peer-group name  |

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967294  |  Neighbor AS number  |
|  external  |  Any AS different from the local AS  |
|  internal  |  Neighbor AS number  |

`,
		},

		// TagNodes

		// Nodes

	}
}
