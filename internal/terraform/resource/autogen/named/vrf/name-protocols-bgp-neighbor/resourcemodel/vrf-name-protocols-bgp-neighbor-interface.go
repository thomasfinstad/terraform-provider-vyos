// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborInterface{}

// VrfNameProtocolsBgpNeighborInterface describes the resource data model.
type VrfNameProtocolsBgpNeighborInterface struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborInterfacePeerGroup       types.String `tfsdk:"peer_group" vyos:"peer-group,omitempty"`
	LeafVrfNameProtocolsBgpNeighborInterfaceRemoteAs        types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafVrfNameProtocolsBgpNeighborInterfaceSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborInterfaceVsixonly *VrfNameProtocolsBgpNeighborInterfaceVsixonly `tfsdk:"v6only" vyos:"v6only,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer group for this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Peer-group name  |

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Neighbor AS number  |
    |  external  &emsp; |  Any AS different from the local AS  |
    |  internal  &emsp; |  Neighbor AS number  |

`,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  interface  &emsp; |  Interface name  |

`,
		},

		// Nodes

		"v6only": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborInterfaceVsixonly{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable BGP with v6 link-local only

`,
		},
	}
}
