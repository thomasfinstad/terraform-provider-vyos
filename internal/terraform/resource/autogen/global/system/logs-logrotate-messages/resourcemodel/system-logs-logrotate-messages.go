// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemLogsLogrotateMessages describes the resource data model.
type SystemLogsLogrotateMessages struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafSystemLogsLogrotateMessagesMaxSize types.Number `tfsdk:"max_size" vyos:"max-size,omitempty"`
	LeafSystemLogsLogrotateMessagesRotate  types.Number `tfsdk:"rotate" vyos:"rotate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemLogsLogrotateMessages) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemLogsLogrotateMessages) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"logs",

		"logrotate",

		"messages",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLogsLogrotateMessages) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"max_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Size of a single log file that triggers rotation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-1024  &emsp; |  Size in MB  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"rotate": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Count of rotations before old logs will be deleted

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-100  &emsp; |  Rotations  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},
	}
}
