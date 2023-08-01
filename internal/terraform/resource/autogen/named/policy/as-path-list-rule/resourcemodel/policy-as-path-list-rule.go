// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyAsPathListRule describes the resource data model.
type PolicyAsPathListRule struct {
	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:",self-id"`

	ParentIDPolicyAsPathList types.String `tfsdk:"as_path_list" vyos:"as-path-list,parent-id"`

	// LeafNodes
	LeafPolicyAsPathListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyAsPathListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyAsPathListRuleRegex       types.String `tfsdk:"regex" vyos:"regex,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAsPathListRule) GetVyosPath() []string {
	return []string{
		"policy",

		"as-path-list",
		o.ParentIDPolicyAsPathList.ValueString(),

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAsPathListRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `rule_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule for this as-path-list

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  AS path list rule number  |

`,
		},

		"as_path_list_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Add a BGP autonomous system path filter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  AS path list name  |

`,
		},

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take on entries matching this rule

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  permit  &emsp; |  Permit matching entries  |
    |  deny  &emsp; |  Deny matching entries  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"regex": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Regular expression to match against an AS path

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  AS path regular expression (ex: "64501 64502")  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyAsPathListRule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyAsPathListRule) UnmarshalJSON(_ []byte) error {
	return nil
}
