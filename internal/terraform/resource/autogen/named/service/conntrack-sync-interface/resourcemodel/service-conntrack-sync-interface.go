// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceConntrackSyncInterface describes the resource data model.
type ServiceConntrackSyncInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceConntrackSyncInterfacePeer types.String `tfsdk:"peer" vyos:"peer,omitempty"`
	LeafServiceConntrackSyncInterfacePort types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceConntrackSyncInterface) GetVyosPath() []string {
	return []string{
		"service",

		"conntrack-sync",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConntrackSyncInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `interface_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface to use for syncing conntrack entries

`,
		},

		// LeafNodes

		"peer": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of the peer to send the UDP conntrack info too. This disable multicast.

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address to listen for incoming connections  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceConntrackSyncInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceConntrackSyncInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
