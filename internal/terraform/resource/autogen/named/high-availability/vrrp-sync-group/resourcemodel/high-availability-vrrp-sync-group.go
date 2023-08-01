// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HighAvailabilityVrrpSyncGroup describes the resource data model.
type HighAvailabilityVrrpSyncGroup struct {
	SelfIdentifier types.String `tfsdk:"sync_group_id" vyos:",self-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpSyncGroupMember types.List `tfsdk:"member" vyos:"member,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeHighAvailabilityVrrpSyncGroupTransitionScrIPt *HighAvailabilityVrrpSyncGroupTransitionScrIPt `tfsdk:"transition_script" vyos:"transition-script,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpSyncGroup) GetVyosPath() []string {
	return []string{
		"high-availability",

		"vrrp",

		"sync-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpSyncGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `sync_group_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"sync_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VRRP sync group

`,
		},

		// LeafNodes

		"member": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Sync group member

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRRP group name  |

`,
		},

		// Nodes

		"transition_script": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpSyncGroupTransitionScrIPt{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VRRP transition scripts

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *HighAvailabilityVrrpSyncGroup) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *HighAvailabilityVrrpSyncGroup) UnmarshalJSON(_ []byte) error {
	return nil
}
