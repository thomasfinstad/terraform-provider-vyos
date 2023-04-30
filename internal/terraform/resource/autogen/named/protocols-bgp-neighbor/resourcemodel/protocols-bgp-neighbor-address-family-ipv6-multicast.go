// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixMulticast describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixMulticast struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastAddpathTxAll         customtypes.CustomStringValue `tfsdk:"addpath_tx_all" json:"addpath-tx-all,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastAddpathTxPerAs       customtypes.CustomStringValue `tfsdk:"addpath_tx_per_as" json:"addpath-tx-per-as,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastAsOverrIDe           customtypes.CustomStringValue `tfsdk:"as_override" json:"as-override,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastMaximumPrefix        customtypes.CustomStringValue `tfsdk:"maximum_prefix" json:"maximum-prefix,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastMaximumPrefixOut     customtypes.CustomStringValue `tfsdk:"maximum_prefix_out" json:"maximum-prefix-out,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRemovePrivateAs      customtypes.CustomStringValue `tfsdk:"remove_private_as" json:"remove-private-as,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteReflectorClient customtypes.CustomStringValue `tfsdk:"route_reflector_client" json:"route-reflector-client,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteServerClient    customtypes.CustomStringValue `tfsdk:"route_server_client" json:"route-server-client,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastUnsuppressMap        customtypes.CustomStringValue `tfsdk:"unsuppress_map" json:"unsuppress-map,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastWeight               customtypes.CustomStringValue `tfsdk:"weight" json:"weight,omitempty"`

	// TagNodes

	// Nodes
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal           types.Object `tfsdk:"nexthop_local" json:"nexthop-local,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastPrefixList             types.Object `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastConditionallyAdvertise types.Object `tfsdk:"conditionally_advertise" json:"conditionally-advertise,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastAllowasIn              types.Object `tfsdk:"allowas_in" json:"allowas-in,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchanged     types.Object `tfsdk:"attribute_unchanged" json:"attribute-unchanged,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDisableSendCommunity   types.Object `tfsdk:"disable_send_community" json:"disable-send-community,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeList         types.Object `tfsdk:"distribute_list" json:"distribute-list,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastFilterList             types.Object `tfsdk:"filter_list" json:"filter-list,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopSelf            types.Object `tfsdk:"nexthop_self" json:"nexthop-self,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap               types.Object `tfsdk:"route_map" json:"route-map,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfiguration    types.Object `tfsdk:"soft_reconfiguration" json:"soft-reconfiguration,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDefaultOriginate       types.Object `tfsdk:"default_originate" json:"default-originate,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixMulticast) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
		},

		"addpath_tx_per_as": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
		},

		"as_override": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
		},

		"maximum_prefix": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |
`,
		},

		"maximum_prefix_out": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |
`,
		},

		"remove_private_as": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
		},

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

		"unsuppress_map": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
		},

		"weight": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Default weight for routes from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Default weight  |
`,
		},

		// TagNodes

		// Nodes

		"nexthop_local": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Nexthop attributes

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastPrefixList{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastConditionallyAdvertise{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastAllowasIn{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchanged{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDisableSendCommunity{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeList{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastFilterList{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopSelf{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfiguration{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDefaultOriginate{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
		},
	}
}
