// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyLimiterDefault describes the resource data model.
type QosPolicyLimiterDefault struct {
	// LeafNodes
	LeafQosPolicyLimiterDefaultBandwIDth types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyLimiterDefaultBurst     types.String `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafQosPolicyLimiterDefaultExceed    types.String `tfsdk:"exceed" vyos:"exceed,omitempty"`
	LeafQosPolicyLimiterDefaultNotExceed types.String `tfsdk:"not_exceed" vyos:"not-exceed,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiterDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Bits per second  |
    |  <number>bit  &emsp; |  Bits per second  |
    |  <number>kbit  &emsp; |  Kilobits per second  |
    |  <number>mbit  &emsp; |  Megabits per second  |
    |  <number>gbit  &emsp; |  Gigabits per second  |
    |  <number>tbit  &emsp; |  Terabits per second  |
    |  <number>%%  &emsp; |  Percentage of interface link speed  |

`,
		},

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Burst size for this class

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Bytes  |
    |  <number><suffix>  &emsp; |  Bytes with scaling suffix (kb, mb, gb)  |

`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"exceed": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default action for packets exceeding the limiter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  continue  &emsp; |  Do not do anything, just continue with the next action in line  |
    |  drop  &emsp; |  Drop the packet immediately  |
    |  ok  &emsp; |  Accept the packet  |
    |  reclassify  &emsp; |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    |  pipe  &emsp; |  Pass the packet to the next action in line  |

`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"not_exceed": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default action for packets not exceeding the limiter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  continue  &emsp; |  Do not do anything, just continue with the next action in line  |
    |  drop  &emsp; |  Drop the packet immediately  |
    |  ok  &emsp; |  Accept the packet  |
    |  reclassify  &emsp; |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    |  pipe  &emsp; |  Pass the packet to the next action in line  |

`,

			// Default:          stringdefault.StaticString(`ok`),
			Computed: true,
		},

		// Nodes

	}
}