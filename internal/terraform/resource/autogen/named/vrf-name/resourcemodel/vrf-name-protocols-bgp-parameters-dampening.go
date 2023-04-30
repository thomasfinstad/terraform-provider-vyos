// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpParametersDampening describes the resource data model.
type VrfNameProtocolsBgpParametersDampening struct {
	// LeafNodes
	VrfNameProtocolsBgpParametersDampeningHalfLife          customtypes.CustomStringValue `tfsdk:"half_life" json:"half-life,omitempty"`
	VrfNameProtocolsBgpParametersDampeningMaxSuppressTime   customtypes.CustomStringValue `tfsdk:"max_suppress_time" json:"max-suppress-time,omitempty"`
	VrfNameProtocolsBgpParametersDampeningReUse             customtypes.CustomStringValue `tfsdk:"re_use" json:"re-use,omitempty"`
	VrfNameProtocolsBgpParametersDampeningStartSuppressTime customtypes.CustomStringValue `tfsdk:"start_suppress_time" json:"start-suppress-time,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersDampening) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"half_life": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Half-life time for dampening

|  Format  |  Description  |
|----------|---------------|
|  u32:1-45  |  Half-life penalty in minutes  |
`,
		},

		"max_suppress_time": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum duration to suppress a stable route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Maximum suppress duration in minutes  |
`,
		},

		"re_use": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Threshold to start reusing a route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-20000  |  Re-use penalty points  |
`,
		},

		"start_suppress_time": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `When to start suppressing a route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-20000  |  Start-suppress penalty points  |
`,
		},

		// TagNodes

		// Nodes

	}
}
