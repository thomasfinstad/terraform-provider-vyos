// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersTCPKeepalive{}

// VrfNameProtocolsBgpParametersTCPKeepalive describes the resource data model.
type VrfNameProtocolsBgpParametersTCPKeepalive struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersTCPKeepaliveIDle     types.Number `tfsdk:"idle" vyos:"idle,omitempty"`
	LeafVrfNameProtocolsBgpParametersTCPKeepaliveInterval types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafVrfNameProtocolsBgpParametersTCPKeepaliveProbes   types.Number `tfsdk:"probes" vyos:"probes,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersTCPKeepalive) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"idle": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP keepalive idle time

    |  Format   &emsp;|  Description           |
    |-----------------|------------------------|
    |  1-65535  &emsp;|  Idle time in seconds  |
`,
			Description: `TCP keepalive idle time

    |  Format   |  Description           |
    |-----------------|------------------------|
    |  1-65535  |  Idle time in seconds  |
`,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP keepalive interval

    |  Format   &emsp;|  Description          |
    |-----------------|-----------------------|
    |  1-65535  &emsp;|  Interval in seconds  |
`,
			Description: `TCP keepalive interval

    |  Format   |  Description          |
    |-----------------|-----------------------|
    |  1-65535  |  Interval in seconds  |
`,
		},

		"probes": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP keepalive maximum probes

    |  Format  &emsp;|  Description     |
    |----------------|------------------|
    |  1-30    &emsp;|  Maximum probes  |
`,
			Description: `TCP keepalive maximum probes

    |  Format  |  Description     |
    |----------------|------------------|
    |  1-30    |  Maximum probes  |
`,
		},

		// Nodes

	}
}
