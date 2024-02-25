// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsPimRp describes the resource data model.
type ProtocolsPimRp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsPimRpKeepAliveTimer types.Number `tfsdk:"keep_alive_timer" vyos:"keep-alive-timer,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsPimRpAddress bool `tfsdk:"-" vyos:"address,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsPimRp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsPimRp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"pim",

		"rp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsPimRp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"keep_alive_timer": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Keep alive Timer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 31-60000  &emsp; |  Keep alive Timer in seconds  |

`,
		},
	}
}
