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

// PolicyPrefixListRule describes the resource data model.
type PolicyPrefixListRule struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:",self-id"`

	ParentIDPolicyPrefixList types.String `tfsdk:"prefix_list" vyos:"prefix-list,parent-id"`

	// LeafNodes
	LeafPolicyPrefixListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyPrefixListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyPrefixListRuleGe          types.Number `tfsdk:"ge" vyos:"ge,omitempty"`
	LeafPolicyPrefixListRuleLe          types.Number `tfsdk:"le" vyos:"le,omitempty"`
	LeafPolicyPrefixListRulePrefix      types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyPrefixListRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyPrefixListRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"policy",

		"prefix-list",
		o.ParentIDPolicyPrefixList.ValueString(),

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyPrefixListRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule for this prefix-list

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Prefix-list rule number  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"prefix_list_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IP prefix-list filter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of IPv4 prefix-list  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
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

		"ge": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length to match a netmask greater than or equal to it

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-32  &emsp; |  Netmask greater than length  |

`,
		},

		"le": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length to match a netmask less than or equal to it

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-32  &emsp; |  Netmask less than length  |

`,
		},

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix to match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Prefix to match against  |

`,
		},

		// Nodes

	}
}
