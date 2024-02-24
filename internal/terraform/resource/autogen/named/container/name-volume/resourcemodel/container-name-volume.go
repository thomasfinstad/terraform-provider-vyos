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

// ContainerNameVolume describes the resource data model.
type ContainerNameVolume struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"volume_id" vyos:"-,self-id"`

	ParentIDContainerName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafContainerNameVolumeSource      types.String `tfsdk:"source" vyos:"source,omitempty"`
	LeafContainerNameVolumeDestination types.String `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafContainerNameVolumeMode        types.String `tfsdk:"mode" vyos:"mode,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ContainerNameVolume) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerNameVolume) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"container",

		"name",
		o.ParentIDContainerName.ValueString(),

		"volume",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNameVolume) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"volume_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Mount a volume into the container

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Container name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source host directory

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Source host directory  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination container directory

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination container directory  |

`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Volume access mode ro/rw

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ro  &emsp; |  Volume mounted into the container as read-only  |
    |  rw  &emsp; |  Volume mounted into the container as read-write  |

`,

			// Default:          stringdefault.StaticString(`rw`),
			Computed: true,
		},

		// Nodes

	}
}
