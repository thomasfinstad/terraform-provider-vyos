// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceDhcpRelayRelayOptions describes the resource data model.
type ServiceDhcpRelayRelayOptions struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceDhcpRelayRelayOptionsHopCount           types.Number `tfsdk:"hop_count" vyos:"hop-count,omitempty"`
	LeafServiceDhcpRelayRelayOptionsMaxSize            types.Number `tfsdk:"max_size" vyos:"max-size,omitempty"`
	LeafServiceDhcpRelayRelayOptionsRelayAgentsPackets types.String `tfsdk:"relay_agents_packets" vyos:"relay-agents-packets,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceDhcpRelayRelayOptions) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpRelayRelayOptions) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"dhcp-relay",

		"relay-options",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpRelayRelayOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"hop_count": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Policy to discard packets that have reached specified hop-count

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Hop count  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"max_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum packet size to send to a DHCPv4/BOOTP server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 64-1400  &emsp; |  Maximum packet size  |

`,

			// Default:          stringdefault.StaticString(`576`),
			Computed: true,
		},

		"relay_agents_packets": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy to handle incoming DHCPv4 packets which already contain relay agent options

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  append  &emsp; |  append own relay options to packet  |
    |  replace  &emsp; |  replace existing agent option field  |
    |  forward  &emsp; |  forward packet unchanged  |
    |  discard  &emsp; |  discard packet (default action if giaddr not set in packet)  |

`,

			// Default:          stringdefault.StaticString(`forward`),
			Computed: true,
		},
	}
}
