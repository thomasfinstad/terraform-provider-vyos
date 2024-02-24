// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBabelDistributeListIPvfourAccessList describes the resource data model.
type ProtocolsBabelDistributeListIPvfourAccessList struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsBabelDistributeListIPvfourAccessListIn  types.Number `tfsdk:"in" vyos:"in,omitempty"`
	LeafProtocolsBabelDistributeListIPvfourAccessListOut types.Number `tfsdk:"out" vyos:"out,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsBabelDistributeListIPvfourAccessList) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelDistributeListIPvfourAccessList) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"babel",

		"distribute-list",

		"ipv4",

		"access-list",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelDistributeListIPvfourAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"in": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to input packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Access list to apply to input packets  |

`,
		},

		"out": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to output packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Access list to apply to output packets  |

`,
		},
	}
}
