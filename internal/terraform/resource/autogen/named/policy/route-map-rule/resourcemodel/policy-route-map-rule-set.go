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

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSet{}

// PolicyRouteMapRuleSet describes the resource data model.
type PolicyRouteMapRuleSet struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetAtomicAggregate types.Bool   `tfsdk:"atomic_aggregate" vyos:"atomic-aggregate,omitempty"`
	LeafPolicyRouteMapRuleSetDistance        types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafPolicyRouteMapRuleSetIPNextHop       types.String `tfsdk:"ip_next_hop" vyos:"ip-next-hop,omitempty"`
	LeafPolicyRouteMapRuleSetLocalPreference types.Number `tfsdk:"local_preference" vyos:"local-preference,omitempty"`
	LeafPolicyRouteMapRuleSetMetric          types.String `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafPolicyRouteMapRuleSetMetricType      types.String `tfsdk:"metric_type" vyos:"metric-type,omitempty"`
	LeafPolicyRouteMapRuleSetOrigin          types.String `tfsdk:"origin" vyos:"origin,omitempty"`
	LeafPolicyRouteMapRuleSetOriginatorID    types.String `tfsdk:"originator_id" vyos:"originator-id,omitempty"`
	LeafPolicyRouteMapRuleSetSrc             types.String `tfsdk:"src" vyos:"src,omitempty"`
	LeafPolicyRouteMapRuleSetTable           types.Number `tfsdk:"table" vyos:"table,omitempty"`
	LeafPolicyRouteMapRuleSetTag             types.Number `tfsdk:"tag" vyos:"tag,omitempty"`
	LeafPolicyRouteMapRuleSetWeight          types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRouteMapRuleSetAggregator       *PolicyRouteMapRuleSetAggregator       `tfsdk:"aggregator" vyos:"aggregator,omitempty"`
	NodePolicyRouteMapRuleSetAsPath           *PolicyRouteMapRuleSetAsPath           `tfsdk:"as_path" vyos:"as-path,omitempty"`
	NodePolicyRouteMapRuleSetCommunity        *PolicyRouteMapRuleSetCommunity        `tfsdk:"community" vyos:"community,omitempty"`
	NodePolicyRouteMapRuleSetLargeCommunity   *PolicyRouteMapRuleSetLargeCommunity   `tfsdk:"large_community" vyos:"large-community,omitempty"`
	NodePolicyRouteMapRuleSetExtcommunity     *PolicyRouteMapRuleSetExtcommunity     `tfsdk:"extcommunity" vyos:"extcommunity,omitempty"`
	NodePolicyRouteMapRuleSetEvpn             *PolicyRouteMapRuleSetEvpn             `tfsdk:"evpn" vyos:"evpn,omitempty"`
	NodePolicyRouteMapRuleSetIPvsixNextHop    *PolicyRouteMapRuleSetIPvsixNextHop    `tfsdk:"ipv6_next_hop" vyos:"ipv6-next-hop,omitempty"`
	NodePolicyRouteMapRuleSetLthreevpnNexthop *PolicyRouteMapRuleSetLthreevpnNexthop `tfsdk:"l3vpn_nexthop" vyos:"l3vpn-nexthop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"atomic_aggregate": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `BGP atomic aggregate attribute

`,
			Description: `BGP atomic aggregate attribute

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Locally significant administrative distance

    |  Format  |  Description     |
    |----------|------------------|
    |  0-255   |  Distance value  |
`,
			Description: `Locally significant administrative distance

    |  Format  |  Description     |
    |----------|------------------|
    |  0-255   |  Distance value  |
`,
		},

		"ip_next_hop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Nexthop IP address

    |  Format        |  Description                                             |
    |----------------|----------------------------------------------------------|
    |  ipv4          |  IP address                                              |
    |  unchanged     |  Set the BGP nexthop address as unchanged                |
    |  peer-address  |  Set the BGP nexthop address to the address of the peer  |
`,
			Description: `Nexthop IP address

    |  Format        |  Description                                             |
    |----------------|----------------------------------------------------------|
    |  ipv4          |  IP address                                              |
    |  unchanged     |  Set the BGP nexthop address as unchanged                |
    |  peer-address  |  Set the BGP nexthop address to the address of the peer  |
`,
		},

		"local_preference": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP local preference attribute

    |  Format        |  Description             |
    |----------------|--------------------------|
    |  0-4294967295  |  Local preference value  |
`,
			Description: `BGP local preference attribute

    |  Format        |  Description             |
    |----------------|--------------------------|
    |  0-4294967295  |  Local preference value  |
`,
		},

		"metric": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination routing protocol metric

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  <+/-metric>   |  Add or subtract metric           |
    |  0-4294967295  |  Metric value                     |
    |  <+/-rtt>      |  Add or subtract round trip time  |
    |  <rtt>         |  Round trip time                  |
`,
			Description: `Destination routing protocol metric

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  <+/-metric>   |  Add or subtract metric           |
    |  0-4294967295  |  Metric value                     |
    |  <+/-rtt>      |  Add or subtract round trip time  |
    |  <rtt>         |  Round trip time                  |
`,
		},

		"metric_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Open Shortest Path First (OSPF) external metric-type

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  type-1  |  OSPF external type 1 metric  |
    |  type-2  |  OSPF external type 2 metric  |
`,
			Description: `Open Shortest Path First (OSPF) external metric-type

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  type-1  |  OSPF external type 1 metric  |
    |  type-2  |  OSPF external type 2 metric  |
`,
		},

		"origin": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Border Gateway Protocl (BGP) origin code

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  igp         |  Interior gateway protocol origin  |
    |  egp         |  Exterior gateway protocol origin  |
    |  incomplete  |  Incomplete origin                 |
`,
			Description: `Border Gateway Protocl (BGP) origin code

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  igp         |  Interior gateway protocol origin  |
    |  egp         |  Exterior gateway protocol origin  |
    |  incomplete  |  Incomplete origin                 |
`,
		},

		"originator_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP originator ID attribute

    |  Format  |  Description           |
    |----------|------------------------|
    |  ipv4    |  Orignator IP address  |
`,
			Description: `BGP originator ID attribute

    |  Format  |  Description           |
    |----------|------------------------|
    |  ipv4    |  Orignator IP address  |
`,
		},

		"src": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source address for route

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
			Description: `Source address for route

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
		},

		"table": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set prefixes to table

    |  Format  |  Description  |
    |----------|---------------|
    |  1-200   |  Table value  |
`,
			Description: `Set prefixes to table

    |  Format  |  Description  |
    |----------|---------------|
    |  1-200   |  Table value  |
`,
		},

		"tag": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Route tag value

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  Route tag    |
`,
			Description: `Route tag value

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  Route tag    |
`,
		},

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP weight attribute

    |  Format        |  Description  |
    |----------------|---------------|
    |  0-4294967295  |  BGP weight   |
`,
			Description: `BGP weight attribute

    |  Format        |  Description  |
    |----------------|---------------|
    |  0-4294967295  |  BGP weight   |
`,
		},

		// Nodes

		"aggregator": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetAggregator{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP aggregator attribute

`,
			Description: `BGP aggregator attribute

`,
		},

		"as_path": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetAsPath{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Transform BGP AS_PATH attribute

`,
			Description: `Transform BGP AS_PATH attribute

`,
		},

		"community": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP community attribute

`,
			Description: `BGP community attribute

`,
		},

		"large_community": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetLargeCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP large community attribute

`,
			Description: `BGP large community attribute

`,
		},

		"extcommunity": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetExtcommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP extended community attribute

`,
			Description: `BGP extended community attribute

`,
		},

		"evpn": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetEvpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Ethernet Virtual Private Network

`,
			Description: `Ethernet Virtual Private Network

`,
		},

		"ipv6_next_hop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetIPvsixNextHop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Nexthop IPv6 address

`,
			Description: `Nexthop IPv6 address

`,
		},

		"l3vpn_nexthop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetLthreevpnNexthop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Next hop Information

`,
			Description: `Next hop Information

`,
		},
	}
}
