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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxAll         types.Bool   `tfsdk:"addpath_tx_all" vyos:"addpath-tx-all,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxPerAs       types.Bool   `tfsdk:"addpath_tx_per_as" vyos:"addpath-tx-per-as,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAsOverrIDe           types.Bool   `tfsdk:"as_override" vyos:"as-override,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefix        types.Number `tfsdk:"maximum_prefix" vyos:"maximum-prefix,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefixOut     types.Number `tfsdk:"maximum_prefix_out" vyos:"maximum-prefix-out,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteReflectorClient types.Bool   `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteServerClient    types.Bool   `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnUnsuppressMap        types.String `tfsdk:"unsuppress_map" vyos:"unsuppress-map,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnWeight               types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal           *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal           `tfsdk:"nexthop_local" vyos:"nexthop-local,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList             *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList             `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise `tfsdk:"conditionally_advertise" vyos:"conditionally-advertise,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn              *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn              `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged     *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged     `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity   *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity   `tfsdk:"disable_send_community" vyos:"disable-send-community,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList         *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList         `tfsdk:"distribute_list" vyos:"distribute-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList             *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList             `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf            *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf            `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs        *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs        `tfsdk:"remove_private_as" vyos:"remove-private-as,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap               *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap               `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration    *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration    `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
			Description: `Use addpath to advertise all paths to a neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"addpath_tx_per_as": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
			Description: `Use addpath to advertise the bestpath per each neighboring AS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"as_override": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
			Description: `Override ASN in outbound updates to configured neighbor local-as

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"maximum_prefix": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
			Description: `Maximum number of prefixes to accept from this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
		},

		"maximum_prefix_out": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
			Description: `Maximum number of prefixes to be sent to this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
		},

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

		"unsuppress_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to selectively unsuppress suppressed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default weight for routes from this peer

    |  Format   |  Description     |
    |-----------|------------------|
    |  1-65535  |  Default weight  |
`,
			Description: `Default weight for routes from this peer

    |  Format   |  Description     |
    |-----------|------------------|
    |  1-65535  |  Default weight  |
`,
		},

		// Nodes

		"nexthop_local": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Nexthop attributes

`,
			Description: `Nexthop attributes

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
			Description: `Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
			Description: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
			Description: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
			Description: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
			Description: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
			Description: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
			Description: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
			Description: `Disable the next hop calculation for this peer

`,
		},

		"remove_private_as": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
			Description: `Remove private AS numbers from AS path in outbound route updates

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},
	}
}
