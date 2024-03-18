// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &PolicyAccessListRule{}

// PolicyAccessListRule describes the resource data model.
type PolicyAccessListRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	ParentIDPolicyAccessList types.Number `tfsdk:"access_list_id" vyos:"access-list,parent-id"`

	// LeafNodes
	LeafPolicyAccessListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyAccessListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyAccessListRuleDestination *PolicyAccessListRuleDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodePolicyAccessListRuleSource      *PolicyAccessListRuleSource      `tfsdk:"source" vyos:"source,omitempty"`
}

// SetID configures the resource ID
func (o *PolicyAccessListRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyAccessListRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAccessListRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *PolicyAccessListRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"access-list",
		o.ParentIDPolicyAccessList.ValueBigFloat().String(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyAccessListRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"access-list",
		o.ParentIDPolicyAccessList.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Rule for this access-list

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Access-list rule number  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"access_list_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `IP access-list filter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-99  &emsp; |  IP standard access list  |
    |  number: 100-199  &emsp; |  IP extended access list  |
    |  number: 1300-1999  &emsp; |  IP standard access list (expanded range)  |
    |  number: 2000-2699  &emsp; |  IP extended access list (expanded range)  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
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

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: PolicyAccessListRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination network or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyAccessListRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source network or address to match

`,
		},
	}
}
