// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpTimers describes the resource data model.
type VrfNameProtocolsBgpTimers struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpTimersHoldtime  types.String `tfsdk:"holdtime" vyos:"holdtime,omitempty"`
	LeafVrfNameProtocolsBgpTimersKeepalive types.Number `tfsdk:"keepalive" vyos:"keepalive,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpTimers) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"holdtime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hold timer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Hold timer in seconds  |
    |  0  &emsp; |  Disable hold timer  |

`,
		},

		"keepalive": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP keepalive interval for this neighbor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Keepalive interval in seconds  |

`,
		},

		// Nodes

	}
}