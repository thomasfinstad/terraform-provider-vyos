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

// PolicyLocalRoutesixRule describes the resource data model.
type PolicyLocalRoutesixRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	// LeafNodes
	LeafPolicyLocalRoutesixRuleFwmark           types.Number `tfsdk:"fwmark" vyos:"fwmark,omitempty"`
	LeafPolicyLocalRoutesixRuleSource           types.List   `tfsdk:"source" vyos:"source,omitempty"`
	LeafPolicyLocalRoutesixRuleDestination      types.List   `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafPolicyLocalRoutesixRuleInboundInterface types.String `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyLocalRoutesixRuleSet *PolicyLocalRoutesixRuleSet `tfsdk:"set" vyos:"set,omitempty"`
}

// SetID configures the resource ID
func (o *PolicyLocalRoutesixRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyLocalRoutesixRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"policy",

		"local-route6",

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 policy local-route rule set number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-32765  &emsp; |  Local-route rule number (1-32765)  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"fwmark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2147483647  &emsp; |  Address to match against  |

`,
		},

		"source": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Source address or prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  Address to match against  |
    |  ipv6net  &emsp; |  Prefix to match against  |

`,
		},

		"destination": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Destination address or prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  Address to match against  |
    |  ipv6net  &emsp; |  Prefix to match against  |

`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound Interface

`,
		},

		// Nodes

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRoutesixRuleSet{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},
	}
}
