// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersDistance{}

// VrfNameProtocolsBgpParametersDistance describes the resource data model.
type VrfNameProtocolsBgpParametersDistance struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpParametersDistancePrefix bool `tfsdk:"prefix" vyos:"prefix,child"`

	// Nodes
	NodeVrfNameProtocolsBgpParametersDistanceGlobal *VrfNameProtocolsBgpParametersDistanceGlobal `tfsdk:"global" vyos:"global,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersDistance) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"global": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDistanceGlobal{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Global administratives distances for BGP routes

`,
			Description: `Global administratives distances for BGP routes

`,
		},
	}
}
