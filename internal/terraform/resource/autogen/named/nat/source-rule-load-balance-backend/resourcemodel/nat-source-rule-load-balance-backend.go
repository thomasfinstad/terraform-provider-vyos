// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &NatSourceRuleLoadBalanceBackend{}

// NatSourceRuleLoadBalanceBackend describes the resource data model.
type NatSourceRuleLoadBalanceBackend struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"backend_id" vyos:"-,self-id"`

	ParentIDNatSourceRule types.Number `tfsdk:"rule_id" vyos:"rule,parent-id"`

	// LeafNodes
	LeafNatSourceRuleLoadBalanceBackendWeight types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *NatSourceRuleLoadBalanceBackend) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *NatSourceRuleLoadBalanceBackend) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NatSourceRuleLoadBalanceBackend) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"backend",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *NatSourceRuleLoadBalanceBackend) GetVyosParentPath() []string {
	return []string{
		"nat",

		"source",

		"rule",
		o.ParentIDNatSourceRule.ValueBigFloat().String(),

		"load-balance",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *NatSourceRuleLoadBalanceBackend) GetVyosNamedParentPath() []string {
	return []string{
		"nat",

		"source",

		"rule",
		o.ParentIDNatSourceRule.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatSourceRuleLoadBalanceBackend) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"backend_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Translated IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to match  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in backend_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  backend_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Rule number for NAT

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999999  &emsp; |  Number of NAT rule  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set probability for this output value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-100  &emsp; |  Set probability for this output value  |

`,
		},

		// Nodes

	}
}
