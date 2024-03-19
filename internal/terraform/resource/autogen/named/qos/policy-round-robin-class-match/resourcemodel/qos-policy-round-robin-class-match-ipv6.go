// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyRoundRobinClassMatchIPvsix{}

// QosPolicyRoundRobinClassMatchIPvsix describes the resource data model.
type QosPolicyRoundRobinClassMatchIPvsix struct {
	// LeafNodes
	LeafQosPolicyRoundRobinClassMatchIPvsixDscp      types.String `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafQosPolicyRoundRobinClassMatchIPvsixMaxLength types.Number `tfsdk:"max_length" vyos:"max-length,omitempty"`
	LeafQosPolicyRoundRobinClassMatchIPvsixProtocol  types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyRoundRobinClassMatchIPvsixDestination *QosPolicyRoundRobinClassMatchIPvsixDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeQosPolicyRoundRobinClassMatchIPvsixSource      *QosPolicyRoundRobinClassMatchIPvsixSource      `tfsdk:"source" vyos:"source,omitempty"`
	NodeQosPolicyRoundRobinClassMatchIPvsixTCP         *QosPolicyRoundRobinClassMatchIPvsixTCP         `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinClassMatchIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on Differentiated Services Codepoint (DSCP)

    |  Format          &emsp;|  Description                                      |
    |------------------------|---------------------------------------------------|
    |  0-63            &emsp;|  Differentiated Services Codepoint (DSCP) value   |
    |  default         &emsp;|  match DSCP (000000)                              |
    |  reliability     &emsp;|  match DSCP (000001)                              |
    |  throughput      &emsp;|  match DSCP (000010)                              |
    |  lowdelay        &emsp;|  match DSCP (000100)                              |
    |  priority        &emsp;|  match DSCP (001000)                              |
    |  immediate       &emsp;|  match DSCP (010000)                              |
    |  flash           &emsp;|  match DSCP (011000)                              |
    |  flash-override  &emsp;|  match DSCP (100000)                              |
    |  critical        &emsp;|  match DSCP (101000)                              |
    |  internet        &emsp;|  match DSCP (110000)                              |
    |  network         &emsp;|  match DSCP (111000)                              |
    |  AF11            &emsp;|  High-throughput data                             |
    |  AF12            &emsp;|  High-throughput data                             |
    |  AF13            &emsp;|  High-throughput data                             |
    |  AF21            &emsp;|  Low-latency data                                 |
    |  AF22            &emsp;|  Low-latency data                                 |
    |  AF23            &emsp;|  Low-latency data                                 |
    |  AF31            &emsp;|  Multimedia streaming                             |
    |  AF32            &emsp;|  Multimedia streaming                             |
    |  AF33            &emsp;|  Multimedia streaming                             |
    |  AF41            &emsp;|  Multimedia conferencing                          |
    |  AF42            &emsp;|  Multimedia conferencing                          |
    |  AF43            &emsp;|  Multimedia conferencing                          |
    |  CS1             &emsp;|  Low-priority data                                |
    |  CS2             &emsp;|  OAM                                              |
    |  CS3             &emsp;|  Broadcast video                                  |
    |  CS4             &emsp;|  Real-time interactive                            |
    |  CS5             &emsp;|  Signaling                                        |
    |  CS6             &emsp;|  Network control                                  |
    |  CS7             &emsp;|  N/A                                              |
    |  EF              &emsp;|  Expedited Forwarding                             |
`,
			Description: `Match on Differentiated Services Codepoint (DSCP)

    |  Format          |  Description                                      |
    |------------------------|---------------------------------------------------|
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

    |  Format   &emsp;|  Description                    |
    |-----------------|---------------------------------|
    |  1-65535  &emsp;|  Maximum packet/payload length  |
`,
			Description: `Maximum packet length

    |  Format   |  Description                    |
    |-----------------|---------------------------------|
    |  1-65535  |  Maximum packet/payload length  |
`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

    |  Format  &emsp;|  Description    |
    |----------------|-----------------|
    |  txt     &emsp;|  Protocol name  |
`,
			Description: `Protocol

    |  Format  |  Description    |
    |----------------|-----------------|
    |  txt     |  Protocol name  |
`,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchIPvsixDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
			Description: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchIPvsixSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
			Description: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchIPvsixTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
			Description: `TCP Flags matching

`,
		},
	}
}
