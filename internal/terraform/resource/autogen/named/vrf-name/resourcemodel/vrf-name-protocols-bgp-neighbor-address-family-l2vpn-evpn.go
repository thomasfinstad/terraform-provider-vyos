// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpn struct {
	// LeafNodes
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteReflectorClient customtypes.CustomStringValue `tfsdk:"route_reflector_client" json:"route-reflector-client,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteServerClient    customtypes.CustomStringValue `tfsdk:"route_server_client" json:"route-server-client,omitempty"`

	// TagNodes

	// Nodes
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn           types.Object `tfsdk:"allowas_in" json:"allowas-in,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged  types.Object `tfsdk:"attribute_unchanged" json:"attribute-unchanged,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnNexthopSelf         types.Object `tfsdk:"nexthop_self" json:"nexthop-self,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap            types.Object `tfsdk:"route_map" json:"route-map,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration types.Object `tfsdk:"soft_reconfiguration" json:"soft-reconfiguration,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpn) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_reflector_client": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Peer is a route reflector client

`,
		},

		"route_server_client": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Peer is a route server client

`,
		},

		// TagNodes

		// Nodes

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnNexthopSelf{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},
	}
}
