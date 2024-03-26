// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyRoundRobinClassMatchIPvsixTCP{}

// QosPolicyRoundRobinClassMatchIPvsixTCP describes the resource data model.
type QosPolicyRoundRobinClassMatchIPvsixTCP struct {
	// LeafNodes
	LeafQosPolicyRoundRobinClassMatchIPvsixTCPAck types.Bool `tfsdk:"ack" vyos:"ack,omitempty"`
	LeafQosPolicyRoundRobinClassMatchIPvsixTCPSyn types.Bool `tfsdk:"syn" vyos:"syn,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinClassMatchIPvsixTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ack": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP ACK

`,
			Description: `Match TCP ACK

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"syn": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP SYN

`,
			Description: `Match TCP SYN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
