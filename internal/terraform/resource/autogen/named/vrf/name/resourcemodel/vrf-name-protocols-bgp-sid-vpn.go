// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpSIDVpn{}

// VrfNameProtocolsBgpSIDVpn describes the resource data model.
type VrfNameProtocolsBgpSIDVpn struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpSIDVpnPerVrf *VrfNameProtocolsBgpSIDVpnPerVrf `tfsdk:"per_vrf" vyos:"per-vrf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpSIDVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"per_vrf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpSIDVpnPerVrf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `SID per-VRF (both IPv4 and IPv6 address families)

`,
			Description: `SID per-VRF (both IPv4 and IPv6 address families)

`,
		},
	}
}
