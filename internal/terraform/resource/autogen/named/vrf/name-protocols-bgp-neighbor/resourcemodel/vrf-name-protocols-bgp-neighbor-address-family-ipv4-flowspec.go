// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteReflectorClient types.Bool `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteServerClient    types.Bool `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecPrefixList          *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecPrefixList          `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecFilterList          *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecFilterList          `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap            *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap            `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecSoftReconfiguration *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecSoftReconfiguration `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_reflector_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
			Description: `Peer is a route reflector client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_server_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
			Description: `Peer is a route server client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
			Description: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecFilterList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
			Description: `as-path-list to filter route updates to/from this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},
	}
}
