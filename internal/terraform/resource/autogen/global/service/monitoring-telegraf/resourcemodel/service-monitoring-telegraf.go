// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceMonitoringTelegraf describes the resource data model.
type ServiceMonitoringTelegraf struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceMonitoringTelegrafSource types.List   `tfsdk:"source" vyos:"source,omitempty"`
	LeafServiceMonitoringTelegrafVrf    types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceMonitoringTelegrafInfluxdb          bool `tfsdk:"-" vyos:"influxdb,omitempty"`
	ExistsNodeServiceMonitoringTelegrafAzureDataExplorer bool `tfsdk:"-" vyos:"azure-data-explorer,omitempty"`
	ExistsNodeServiceMonitoringTelegrafPrometheusClient  bool `tfsdk:"-" vyos:"prometheus-client,omitempty"`
	ExistsNodeServiceMonitoringTelegrafSplunk            bool `tfsdk:"-" vyos:"splunk,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceMonitoringTelegraf) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceMonitoringTelegraf) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"monitoring",

		"telegraf",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceMonitoringTelegraf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"source": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Source parameters for monitoring

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  all  &emsp; |  All parameters  |
    |  hardware-utilization  &emsp; |  Hardware-utilization parameters (CPU, disk, memory)  |
    |  logs  &emsp; |  Logs parameters  |
    |  network  &emsp; |  Network parameters (net, netstat, nftables)  |
    |  system  &emsp; |  System parameters (system, processes, interrupts)  |
    |  telegraf  &emsp; |  Telegraf internal statistics  |

`,

			// Default:          stringdefault.StaticString(`all`),
			Computed: true,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},
	}
}
