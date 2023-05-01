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

// ProtocolsBgpNeighborInterface describes the resource data model.
type ProtocolsBgpNeighborInterface struct {
	// LeafNodes
	LeafProtocolsBgpNeighborInterfacePeerGroup       types.String `tfsdk:"peer_group"`
	LeafProtocolsBgpNeighborInterfaceRemoteAs        types.String `tfsdk:"remote_as"`
	LeafProtocolsBgpNeighborInterfaceSourceInterface types.String `tfsdk:"source_interface"`

	// TagNodes

	// Nodes
	NodeProtocolsBgpNeighborInterfaceVsixonly types.Object `tfsdk:"v6only"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborInterface) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "interface"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborInterfacePeerGroup.IsNull() || o.LeafProtocolsBgpNeighborInterfacePeerGroup.IsUnknown()) {
		vyosData["peer-group"] = o.LeafProtocolsBgpNeighborInterfacePeerGroup.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborInterfaceRemoteAs.IsNull() || o.LeafProtocolsBgpNeighborInterfaceRemoteAs.IsUnknown()) {
		vyosData["remote-as"] = o.LeafProtocolsBgpNeighborInterfaceRemoteAs.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborInterfaceSourceInterface.IsNull() || o.LeafProtocolsBgpNeighborInterfaceSourceInterface.IsUnknown()) {
		vyosData["source-interface"] = o.LeafProtocolsBgpNeighborInterfaceSourceInterface.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeProtocolsBgpNeighborInterfaceVsixonly.IsNull() || o.NodeProtocolsBgpNeighborInterfaceVsixonly.IsUnknown()) {
		var subModel ProtocolsBgpNeighborInterfaceVsixonly
		diags.Append(o.NodeProtocolsBgpNeighborInterfaceVsixonly.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["v6only"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborInterface) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "interface"}})

	// Leafs
	if value, ok := vyosData["peer-group"]; ok {
		o.LeafProtocolsBgpNeighborInterfacePeerGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborInterfacePeerGroup = basetypes.NewStringNull()
	}
	if value, ok := vyosData["remote-as"]; ok {
		o.LeafProtocolsBgpNeighborInterfaceRemoteAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborInterfaceRemoteAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source-interface"]; ok {
		o.LeafProtocolsBgpNeighborInterfaceSourceInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborInterfaceSourceInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["v6only"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborInterfaceVsixonly{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborInterfaceVsixonly = data

	} else {
		o.NodeProtocolsBgpNeighborInterfaceVsixonly = basetypes.NewObjectNull(ProtocolsBgpNeighborInterfaceVsixonly{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "interface"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborInterface) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"peer_group":       types.StringType,
		"remote_as":        types.StringType,
		"source_interface": types.StringType,

		// Tags

		// Nodes
		"v6only": types.ObjectType{AttrTypes: ProtocolsBgpNeighborInterfaceVsixonly{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

|  Format  |  Description  |
|----------|---------------|
|  interface  |  Interface name  |

`,
		},

		// TagNodes

		// Nodes

		"v6only": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborInterfaceVsixonly{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable BGP with v6 link-local only

`,
		},
	}
}
