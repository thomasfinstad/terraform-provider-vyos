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

// PolicyPrefixListsix describes the resource data model.
type PolicyPrefixListsix struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"prefix_list6_id" vyos:"-,self-id"`

	// LeafNodes
	LeafPolicyPrefixListsixDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagPolicyPrefixListsixRule bool `tfsdk:"-" vyos:"rule,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyPrefixListsix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyPrefixListsix) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

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
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"prefix_list6_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 prefix-list filter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of IPv6 prefix-list  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
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