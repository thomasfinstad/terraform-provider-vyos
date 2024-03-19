// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyLimiterClassMatchIPvsix{}

// QosPolicyLimiterClassMatchIPvsix describes the resource data model.
type QosPolicyLimiterClassMatchIPvsix struct {
	// LeafNodes
	LeafQosPolicyLimiterClassMatchIPvsixDscp      types.String `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafQosPolicyLimiterClassMatchIPvsixMaxLength types.Number `tfsdk:"max_length" vyos:"max-length,omitempty"`
	LeafQosPolicyLimiterClassMatchIPvsixProtocol  types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyLimiterClassMatchIPvsixDestination *QosPolicyLimiterClassMatchIPvsixDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeQosPolicyLimiterClassMatchIPvsixSource      *QosPolicyLimiterClassMatchIPvsixSource      `tfsdk:"source" vyos:"source,omitempty"`
	NodeQosPolicyLimiterClassMatchIPvsixTCP         *QosPolicyLimiterClassMatchIPvsixTCP         `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiterClassMatchIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on Differentiated Services Codepoint (DSCP)

    |  Format          |  Description                                      |
    |------------------|---------------------------------------------------|
    |  0-63            |  Differentiated Services Codepoint (DSCP) value   |
    |  default         |  match DSCP (000000)                              |
    |  reliability     |  match DSCP (000001)                              |
    |  throughput      |  match DSCP (000010)                              |
    |  lowdelay        |  match DSCP (000100)                              |
    |  priority        |  match DSCP (001000)                              |
    |  immediate       |  match DSCP (010000)                              |
    |  flash           |  match DSCP (011000)                              |
    |  flash-override  |  match DSCP (100000)                              |
    |  critical        |  match DSCP (101000)                              |
    |  internet        |  match DSCP (110000)                              |
    |  network         |  match DSCP (111000)                              |
    |  AF11            |  High-throughput data                             |
    |  AF12            |  High-throughput data                             |
    |  AF13            |  High-throughput data                             |
    |  AF21            |  Low-latency data                                 |
    |  AF22            |  Low-latency data                                 |
    |  AF23            |  Low-latency data                                 |
    |  AF31            |  Multimedia streaming                             |
    |  AF32            |  Multimedia streaming                             |
    |  AF33            |  Multimedia streaming                             |
    |  AF41            |  Multimedia conferencing                          |
    |  AF42            |  Multimedia conferencing                          |
    |  AF43            |  Multimedia conferencing                          |
    |  CS1             |  Low-priority data                                |
    |  CS2             |  OAM                                              |
    |  CS3             |  Broadcast video                                  |
    |  CS4             |  Real-time interactive                            |
    |  CS5             |  Signaling                                        |
    |  CS6             |  Network control                                  |
    |  CS7             |  N/A                                              |
    |  EF              |  Expedited Forwarding                             |
`,
			Description: `Match on Differentiated Services Codepoint (DSCP)

    |  Format          |  Description                                      |
    |------------------|---------------------------------------------------|
    |  0-63            |  Differentiated Services Codepoint (DSCP) value   |
    |  default         |  match DSCP (000000)                              |
    |  reliability     |  match DSCP (000001)                              |
    |  throughput      |  match DSCP (000010)                              |
    |  lowdelay        |  match DSCP (000100)                              |
    |  priority        |  match DSCP (001000)                              |
    |  immediate       |  match DSCP (010000)                              |
    |  flash           |  match DSCP (011000)                              |
    |  flash-override  |  match DSCP (100000)                              |
    |  critical        |  match DSCP (101000)                              |
    |  internet        |  match DSCP (110000)                              |
    |  network         |  match DSCP (111000)                              |
    |  AF11            |  High-throughput data                             |
    |  AF12            |  High-throughput data                             |
    |  AF13            |  High-throughput data                             |
    |  AF21            |  Low-latency data                                 |
    |  AF22            |  Low-latency data                                 |
    |  AF23            |  Low-latency data                                 |
    |  AF31            |  Multimedia streaming                             |
    |  AF32            |  Multimedia streaming                             |
    |  AF33            |  Multimedia streaming                             |
    |  AF41            |  Multimedia conferencing                          |
    |  AF42            |  Multimedia conferencing                          |
    |  AF43            |  Multimedia conferencing                          |
    |  CS1             |  Low-priority data                                |
    |  CS2             |  OAM                                              |
    |  CS3             |  Broadcast video                                  |
    |  CS4             |  Real-time interactive                            |
    |  CS5             |  Signaling                                        |
    |  CS6             |  Network control                                  |
    |  CS7             |  N/A                                              |
    |  EF              |  Expedited Forwarding                             |
`,
		},

		"max_length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum packet length

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Maximum packet/payload length  |
`,
			Description: `Maximum packet length

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Maximum packet/payload length  |
`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

    |  Format  |  Description    |
    |----------|-----------------|
    |  txt     |  Protocol name  |
`,
			Description: `Protocol

    |  Format  |  Description    |
    |----------|-----------------|
    |  txt     |  Protocol name  |
`,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: QosPolicyLimiterClassMatchIPvsixDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
			Description: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyLimiterClassMatchIPvsixSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
			Description: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyLimiterClassMatchIPvsixTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
			Description: `TCP Flags matching

`,
		},
	}
}
