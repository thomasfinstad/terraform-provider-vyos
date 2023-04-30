// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborInterface describes the resource data model.
type ProtocolsBgpNeighborInterface struct {
	// LeafNodes
	ProtocolsBgpNeighborInterfacePeerGroup       customtypes.CustomStringValue `tfsdk:"peer_group" json:"peer-group,omitempty"`
	ProtocolsBgpNeighborInterfaceRemoteAs        customtypes.CustomStringValue `tfsdk:"remote_as" json:"remote-as,omitempty"`
	ProtocolsBgpNeighborInterfaceSourceInterface customtypes.CustomStringValue `tfsdk:"source_interface" json:"source-interface,omitempty"`

	// TagNodes

	// Nodes
	ProtocolsBgpNeighborInterfaceVsixonly types.Object `tfsdk:"v6only" json:"v6only,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer_group": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Peer group for this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Peer-group name  |
`,
		},

		"remote_as": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Neighbor BGP AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967294  |  Neighbor AS number  |
|  external  |  Any AS different from the local AS  |
|  internal  |  Neighbor AS number  |
`,
		},

		"source_interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface used to establish connection

|  Format  |  Description  |
|----------|---------------|
|  interface  |  Interface name  |
`,
		},

		// TagNodes

		// Nodes

		"v6only": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborInterfaceVsixonly{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable BGP with v6 link-local only

`,
		},
	}
}
