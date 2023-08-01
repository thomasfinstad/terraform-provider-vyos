// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyPrefixListsix describes the resource data model.
type PolicyPrefixListsix struct {
	SelfIdentifier types.String `tfsdk:"prefix_list6_id" vyos:",self-id"`

	// LeafNodes
	LeafPolicyPrefixListsixDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagPolicyPrefixListsixRule bool `tfsdk:"rule" vyos:"rule,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyPrefixListsix) GetVyosPath() []string {
	return []string{
		"policy",

		"prefix-list6",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyPrefixListsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `prefix_list6_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"prefix_list6_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 prefix-list filter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of IPv6 prefix-list  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyPrefixListsix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyPrefixListsix) UnmarshalJSON(_ []byte) error {
	return nil
}
