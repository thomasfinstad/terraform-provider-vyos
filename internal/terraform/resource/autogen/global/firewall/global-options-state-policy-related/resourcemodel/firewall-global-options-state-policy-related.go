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
var _ helpers.VyosTopResourceDataModel = &FirewallGlobalOptionsStatePolicyRelated{}

// FirewallGlobalOptionsStatePolicyRelated describes the resource data model.
type FirewallGlobalOptionsStatePolicyRelated struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafFirewallGlobalOptionsStatePolicyRelatedAction   types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallGlobalOptionsStatePolicyRelatedLog      types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallGlobalOptionsStatePolicyRelatedLogLevel types.String `tfsdk:"log_level" vyos:"log-level,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *FirewallGlobalOptionsStatePolicyRelated) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallGlobalOptionsStatePolicyRelated) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGlobalOptionsStatePolicyRelated) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"related",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallGlobalOptionsStatePolicyRelated) GetVyosParentPath() []string {
	return []string{
		"firewall",

		"global-options",

		"state-policy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *FirewallGlobalOptionsStatePolicyRelated) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGlobalOptionsStatePolicyRelated) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action for packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  accept  &emsp; |  Action to accept  |
    |  drop  &emsp; |  Action to drop  |
    |  reject  &emsp; |  Action to reject  |

`,
		},

		"log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"log_level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set log-level. Log must be enable.

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  emerg  &emsp; |  Emerg log level  |
    |  alert  &emsp; |  Alert log level  |
    |  crit  &emsp; |  Critical log level  |
    |  err  &emsp; |  Error log level  |
    |  warn  &emsp; |  Warning log level  |
    |  notice  &emsp; |  Notice log level  |
    |  info  &emsp; |  Info log level  |
    |  debug  &emsp; |  Debug log level  |

`,
		},
	}
}
