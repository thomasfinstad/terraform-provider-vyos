// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemSflow describes the resource data model.
type SystemSflow struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafSystemSflowAgentAddress     types.String `tfsdk:"agent_address" vyos:"agent-address,omitempty"`
	LeafSystemSflowAgentInterface   types.String `tfsdk:"agent_interface" vyos:"agent-interface,omitempty"`
	LeafSystemSflowDropMonitorLimit types.Number `tfsdk:"drop_monitor_limit" vyos:"drop-monitor-limit,omitempty"`
	LeafSystemSflowInterface        types.List   `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafSystemSflowPolling          types.Number `tfsdk:"polling" vyos:"polling,omitempty"`
	LeafSystemSflowSamplingRate     types.Number `tfsdk:"sampling_rate" vyos:"sampling-rate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemSflowServer bool `tfsdk:"-" vyos:"server,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemSflow) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSflow) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"sflow",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSflow) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"agent_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `sFlow agent IPv4 or IPv6 address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  sFlow IPv4 agent address  |
    |  ipv6  &emsp; |  sFlow IPv6 agent address  |

`,
		},

		"agent_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address associated with this interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		"drop_monitor_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Export headers of dropped by kernel packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Maximum rate limit of N drops per second send out in the sFlow datagrams  |

`,
		},

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		"polling": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Schedule counter-polling in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-600  &emsp; |  Polling rate in seconds  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"sampling_rate": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `sFlow sampling-rate

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Sampling rate (1 in N packets)  |

`,

			// Default:          stringdefault.StaticString(`1000`),
			Computed: true,
		},
	}
}
