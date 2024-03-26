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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnRouteReflectorClient types.Bool `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnRouteServerClient    types.Bool `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnAllowasIn           *VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnAllowasIn           `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnAttributeUnchanged  *VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnAttributeUnchanged  `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnNexthopSelf         *VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnNexthopSelf         `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnRouteMap            *VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnRouteMap            `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnSoftReconfiguration *VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnSoftReconfiguration `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnAllowasIn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
			Description: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnAttributeUnchanged{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
			Description: `BGP attributes are sent unchanged

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnNexthopSelf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
			Description: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpnSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},
	}
}
