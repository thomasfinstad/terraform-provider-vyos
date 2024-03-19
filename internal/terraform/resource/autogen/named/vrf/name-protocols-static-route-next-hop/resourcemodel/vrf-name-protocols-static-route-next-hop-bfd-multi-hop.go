// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsStaticRouteNextHopBfdMultiHop{}

// VrfNameProtocolsStaticRouteNextHopBfdMultiHop describes the resource data model.
type VrfNameProtocolsStaticRouteNextHopBfdMultiHop struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsStaticRouteNextHopBfdMultiHopSource bool `tfsdk:"source" vyos:"source,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRouteNextHopBfdMultiHop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// Nodes

	}
}
