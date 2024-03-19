// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance{}

// VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal types.Number `tfsdk:"internal" vyos:"internal,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal    types.Number `tfsdk:"local" vyos:"local,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix bool `tfsdk:"prefix" vyos:"prefix,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------------|---------------------------------------|
    |  1-255   &emsp;|  eBGP routes administrative distance  |
`,
			Description: `eBGP routes administrative distance

    |  Format  |  Description                          |
    |----------------|---------------------------------------|
    |  1-255   |  eBGP routes administrative distance  |
`,
		},

		"internal": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------------|---------------------------------------|
    |  1-255   &emsp;|  iBGP routes administrative distance  |
`,
			Description: `iBGP routes administrative distance

    |  Format  |  Description                          |
    |----------------|---------------------------------------|
    |  1-255   |  iBGP routes administrative distance  |
`,
		},

		"local": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Locally originated BGP routes administrative distance

    |  Format  &emsp;|  Description                                            |
    |----------------|---------------------------------------------------------|
    |  1-255   &emsp;|  Locally originated BGP routes administrative distance  |
`,
			Description: `Locally originated BGP routes administrative distance

    |  Format  |  Description                                            |
    |----------------|---------------------------------------------------------|
    |  1-255   |  Locally originated BGP routes administrative distance  |
`,
		},

		// Nodes

	}
}
