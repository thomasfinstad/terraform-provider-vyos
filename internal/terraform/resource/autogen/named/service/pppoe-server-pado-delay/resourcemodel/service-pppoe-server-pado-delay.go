// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServicePppoeServerPadoDelay describes the resource data model.
type ServicePppoeServerPadoDelay struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"pado_delay_id" vyos:",self-id"`

	// LeafNodes
	LeafServicePppoeServerPadoDelaySessions types.Number `tfsdk:"sessions" vyos:"sessions,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ServicePppoeServerPadoDelay) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServicePppoeServerPadoDelay) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"pppoe-server",

		"pado-delay",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServicePppoeServerPadoDelay) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"pado_delay_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `PADO delays

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999999  &emsp; |  Number in ms  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"sessions": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of sessions

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999999  &emsp; |  Number of sessions  |

`,
		},

		// Nodes

	}
}
