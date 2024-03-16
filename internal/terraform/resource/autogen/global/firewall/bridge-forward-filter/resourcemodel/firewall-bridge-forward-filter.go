// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &FirewallBrIDgeForwardFilter{}

// FirewallBrIDgeForwardFilter describes the resource data model.
type FirewallBrIDgeForwardFilter struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafFirewallBrIDgeForwardFilterDefaultAction types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafFirewallBrIDgeForwardFilterDefaultLog    types.Bool   `tfsdk:"default_log" vyos:"default-log,omitempty"`
	LeafFirewallBrIDgeForwardFilterDescrIPtion   types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagFirewallBrIDgeForwardFilterRule bool `tfsdk:"-" vyos:"rule,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *FirewallBrIDgeForwardFilter) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallBrIDgeForwardFilter) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"bridge",

		"forward",

		"filter",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeForwardFilter) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"default_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default-action for rule-set

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  drop  &emsp; |  Drop if no prior rules are hit  |
    |  accept  &emsp; |  Accept if no prior rules are hit  |

`,

			// Default:          stringdefault.StaticString(`accept`),
			Computed: true,
		},

		"default_log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting default-action

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},
	}
}
