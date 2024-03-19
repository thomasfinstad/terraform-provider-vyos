// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &HighAvailabilityVrrpGlobalParameters{}

// HighAvailabilityVrrpGlobalParameters describes the resource data model.
type HighAvailabilityVrrpGlobalParameters struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafHighAvailabilityVrrpGlobalParametersStartupDelay types.Number `tfsdk:"startup_delay" vyos:"startup-delay,omitempty"`
	LeafHighAvailabilityVrrpGlobalParametersVersion      types.String `tfsdk:"version" vyos:"version,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeHighAvailabilityVrrpGlobalParametersGarp bool `tfsdk:"-" vyos:"garp,child"`
}

// SetID configures the resource ID
func (o *HighAvailabilityVrrpGlobalParameters) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *HighAvailabilityVrrpGlobalParameters) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *HighAvailabilityVrrpGlobalParameters) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpGlobalParameters) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"global-parameters",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *HighAvailabilityVrrpGlobalParameters) GetVyosParentPath() []string {
	return []string{
		"high-availability",

		"vrrp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *HighAvailabilityVrrpGlobalParameters) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGlobalParameters) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"startup_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time VRRP startup process (in seconds)

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-600   |  Interval in seconds  |
`,
			Description: `Time VRRP startup process (in seconds)

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-600   |  Interval in seconds  |
`,
		},

		"version": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default VRRP version to use, IPv6 always uses VRRP version 3

    |  Format  |  Description     |
    |----------|------------------|
    |  2       |  VRRP version 2  |
    |  3       |  VRRP version 3  |
`,
			Description: `Default VRRP version to use, IPv6 always uses VRRP version 3

    |  Format  |  Description     |
    |----------|------------------|
    |  2       |  VRRP version 2  |
    |  3       |  VRRP version 3  |
`,
		},
	}
}
