// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceIDsDdosProtection describes the resource data model.
type ServiceIDsDdosProtection struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceIDsDdosProtectionAlertScrIPt     types.String `tfsdk:"alert_script" vyos:"alert-script,omitempty"`
	LeafServiceIDsDdosProtectionBanTime         types.Number `tfsdk:"ban_time" vyos:"ban-time,omitempty"`
	LeafServiceIDsDdosProtectionDirection       types.List   `tfsdk:"direction" vyos:"direction,omitempty"`
	LeafServiceIDsDdosProtectionExcludedNetwork types.List   `tfsdk:"excluded_network" vyos:"excluded-network,omitempty"`
	LeafServiceIDsDdosProtectionListenInterface types.List   `tfsdk:"listen_interface" vyos:"listen-interface,omitempty"`
	LeafServiceIDsDdosProtectionNetwork         types.List   `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceIDsDdosProtectionMode      bool `tfsdk:"-" vyos:"mode,ignore,omitempty"`
	ExistsNodeServiceIDsDdosProtectionThreshold bool `tfsdk:"-" vyos:"threshold,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceIDsDdosProtection) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIDsDdosProtection) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"ids",

		"ddos-protection",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIDsDdosProtection) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"alert_script": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Path to fastnetmon alert script

`,
		},

		"ban_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `How long we should keep an IP in blocked state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Time in seconds  |

`,

			// Default:          stringdefault.StaticString(`1900`),
			Computed: true,
		},

		"direction": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Direction for processing traffic

`,
		},

		"excluded_network": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Specify IPv4 and IPv6 networks which are going to be excluded from protection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 prefix(es) to exclude  |
    |  ipv6net  &emsp; |  IPv6 prefix(es) to exclude  |

`,
		},

		"listen_interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Listen interface for mirroring traffic

`,
		},

		"network": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Specify IPv4 and IPv6 networks which belong to you

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Your IPv4 prefix(es)  |
    |  ipv6net  &emsp; |  Your IPv6 prefix(es)  |

`,
		},
	}
}
