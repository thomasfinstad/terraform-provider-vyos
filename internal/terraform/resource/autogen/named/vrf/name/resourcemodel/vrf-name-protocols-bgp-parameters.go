// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParameters{}

// VrfNameProtocolsBgpParameters describes the resource data model.
type VrfNameProtocolsBgpParameters struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersAllowMartianNexthop               types.Bool   `tfsdk:"allow_martian_nexthop" vyos:"allow-martian-nexthop,omitempty"`
	LeafVrfNameProtocolsBgpParametersAlwaysCompareMed                  types.Bool   `tfsdk:"always_compare_med" vyos:"always-compare-med,omitempty"`
	LeafVrfNameProtocolsBgpParametersClusterID                         types.String `tfsdk:"cluster_id" vyos:"cluster-id,omitempty"`
	LeafVrfNameProtocolsBgpParametersDeterministicMed                  types.Bool   `tfsdk:"deterministic_med" vyos:"deterministic-med,omitempty"`
	LeafVrfNameProtocolsBgpParametersEbgpRequiresPolicy                types.Bool   `tfsdk:"ebgp_requires_policy" vyos:"ebgp-requires-policy,omitempty"`
	LeafVrfNameProtocolsBgpParametersFastConvergence                   types.Bool   `tfsdk:"fast_convergence" vyos:"fast-convergence,omitempty"`
	LeafVrfNameProtocolsBgpParametersGracefulShutdown                  types.Bool   `tfsdk:"graceful_shutdown" vyos:"graceful-shutdown,omitempty"`
	LeafVrfNameProtocolsBgpParametersNoHardAdministrativeReset         types.Bool   `tfsdk:"no_hard_administrative_reset" vyos:"no-hard-administrative-reset,omitempty"`
	LeafVrfNameProtocolsBgpParametersLabeledUnicast                    types.String `tfsdk:"labeled_unicast" vyos:"labeled-unicast,omitempty"`
	LeafVrfNameProtocolsBgpParametersLogNeighborChanges                types.Bool   `tfsdk:"log_neighbor_changes" vyos:"log-neighbor-changes,omitempty"`
	LeafVrfNameProtocolsBgpParametersMinimumHoldtime                   types.Number `tfsdk:"minimum_holdtime" vyos:"minimum-holdtime,omitempty"`
	LeafVrfNameProtocolsBgpParametersNetworkImportCheck                types.Bool   `tfsdk:"network_import_check" vyos:"network-import-check,omitempty"`
	LeafVrfNameProtocolsBgpParametersRouteReflectorAllowOutboundPolicy types.Bool   `tfsdk:"route_reflector_allow_outbound_policy" vyos:"route-reflector-allow-outbound-policy,omitempty"`
	LeafVrfNameProtocolsBgpParametersNoClientToClientReflection        types.Bool   `tfsdk:"no_client_to_client_reflection" vyos:"no-client-to-client-reflection,omitempty"`
	LeafVrfNameProtocolsBgpParametersNoFastExternalFailover            types.Bool   `tfsdk:"no_fast_external_failover" vyos:"no-fast-external-failover,omitempty"`
	LeafVrfNameProtocolsBgpParametersNoSuppressDuplicates              types.Bool   `tfsdk:"no_suppress_duplicates" vyos:"no-suppress-duplicates,omitempty"`
	LeafVrfNameProtocolsBgpParametersRejectAsSets                      types.Bool   `tfsdk:"reject_as_sets" vyos:"reject-as-sets,omitempty"`
	LeafVrfNameProtocolsBgpParametersShutdown                          types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafVrfNameProtocolsBgpParametersSuppressFibPending                types.Bool   `tfsdk:"suppress_fib_pending" vyos:"suppress-fib-pending,omitempty"`
	LeafVrfNameProtocolsBgpParametersRouterID                          types.String `tfsdk:"router_id" vyos:"router-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpParametersBestpath                 *VrfNameProtocolsBgpParametersBestpath                 `tfsdk:"bestpath" vyos:"bestpath,omitempty"`
	NodeVrfNameProtocolsBgpParametersConfederation            *VrfNameProtocolsBgpParametersConfederation            `tfsdk:"confederation" vyos:"confederation,omitempty"`
	NodeVrfNameProtocolsBgpParametersConditionalAdvertisement *VrfNameProtocolsBgpParametersConditionalAdvertisement `tfsdk:"conditional_advertisement" vyos:"conditional-advertisement,omitempty"`
	NodeVrfNameProtocolsBgpParametersDampening                *VrfNameProtocolsBgpParametersDampening                `tfsdk:"dampening" vyos:"dampening,omitempty"`
	NodeVrfNameProtocolsBgpParametersDefault                  *VrfNameProtocolsBgpParametersDefault                  `tfsdk:"default" vyos:"default,omitempty"`
	NodeVrfNameProtocolsBgpParametersDistance                 *VrfNameProtocolsBgpParametersDistance                 `tfsdk:"distance" vyos:"distance,omitempty"`
	NodeVrfNameProtocolsBgpParametersGracefulRestart          *VrfNameProtocolsBgpParametersGracefulRestart          `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	NodeVrfNameProtocolsBgpParametersTCPKeepalive             *VrfNameProtocolsBgpParametersTCPKeepalive             `tfsdk:"tcp_keepalive" vyos:"tcp-keepalive,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParameters) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"allow_martian_nexthop": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow Martian nexthops to be received in the NLRI from a peer

`,
			Description: `Allow Martian nexthops to be received in the NLRI from a peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"always_compare_med": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Always compare MEDs from different neighbors

`,
			Description: `Always compare MEDs from different neighbors

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"cluster_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-reflector cluster-id

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Route-reflector cluster-id  |
`,
			Description: `Route-reflector cluster-id

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Route-reflector cluster-id  |
`,
		},

		"deterministic_med": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Compare MEDs between different peers in the same AS

`,
			Description: `Compare MEDs between different peers in the same AS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ebgp_requires_policy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Require in and out policy for eBGP peers (RFC8212)

`,
			Description: `Require in and out policy for eBGP peers (RFC8212)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fast_convergence": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Teardown sessions immediately whenever peer becomes unreachable

`,
			Description: `Teardown sessions immediately whenever peer becomes unreachable

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"graceful_shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Graceful shutdown

`,
			Description: `Graceful shutdown

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_hard_administrative_reset": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not send hard reset CEASE Notification for 'Administrative Reset'

`,
			Description: `Do not send hard reset CEASE Notification for 'Administrative Reset'

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"labeled_unicast": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP Labeled-unicast options

    |  Format              |  Description                                                 |
    |----------------------|--------------------------------------------------------------|
    |  explicit-null       |  Use explicit-null label values for all local prefixes       |
    |  ipv4-explicit-null  |  Use IPv4 explicit-null label value for IPv4 local prefixes  |
    |  ipv6-explicit-null  |  Use IPv6 explicit-null label value for IPv4 local prefixes  |
`,
			Description: `BGP Labeled-unicast options

    |  Format              |  Description                                                 |
    |----------------------|--------------------------------------------------------------|
    |  explicit-null       |  Use explicit-null label values for all local prefixes       |
    |  ipv4-explicit-null  |  Use IPv4 explicit-null label value for IPv4 local prefixes  |
    |  ipv6-explicit-null  |  Use IPv6 explicit-null label value for IPv4 local prefixes  |
`,
		},

		"log_neighbor_changes": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log neighbor up/down changes and reset reason

`,
			Description: `Log neighbor up/down changes and reset reason

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"minimum_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP minimum holdtime

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  1-65535  |  Minimum holdtime in seconds  |
`,
			Description: `BGP minimum holdtime

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  1-65535  |  Minimum holdtime in seconds  |
`,
		},

		"network_import_check": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable IGP route check for network statements

`,
			Description: `Enable IGP route check for network statements

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_reflector_allow_outbound_policy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Route reflector client allow policy outbound

`,
			Description: `Route reflector client allow policy outbound

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_client_to_client_reflection": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable client to client route reflection

`,
			Description: `Disable client to client route reflection

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_fast_external_failover": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable immediate session reset on peer link down event

`,
			Description: `Disable immediate session reset on peer link down event

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_suppress_duplicates": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable suppress duplicate updates if the route actually not changed

`,
			Description: `Disable suppress duplicate updates if the route actually not changed

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"reject_as_sets": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Reject routes with AS_SET or AS_CONFED_SET flag

`,
			Description: `Reject routes with AS_SET or AS_CONFED_SET flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administrative shutdown of the BGP instance

`,
			Description: `Administrative shutdown of the BGP instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"suppress_fib_pending": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise only routes that are programmed in kernel to peers

`,
			Description: `Advertise only routes that are programmed in kernel to peers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"router_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override default router identifier

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  Router-ID in IP address format  |
`,
			Description: `Override default router identifier

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  Router-ID in IP address format  |
`,
		},

		// Nodes

		"bestpath": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersBestpath{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Default bestpath selection mechanism

`,
			Description: `Default bestpath selection mechanism

`,
		},

		"confederation": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersConfederation{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `AS confederation parameters

`,
			Description: `AS confederation parameters

`,
		},

		"conditional_advertisement": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersConditionalAdvertisement{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Conditional advertisement settings

`,
			Description: `Conditional advertisement settings

`,
		},

		"dampening": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDampening{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable route-flap dampening

`,
			Description: `Enable route-flap dampening

`,
		},

		"default": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDefault{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP defaults

`,
			Description: `BGP defaults

`,
		},

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDistance{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Administratives distances for BGP routes

`,
			Description: `Administratives distances for BGP routes

`,
		},

		"graceful_restart": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersGracefulRestart{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Graceful restart capability parameters

`,
			Description: `Graceful restart capability parameters

`,
		},

		"tcp_keepalive": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersTCPKeepalive{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP keepalive parameters

`,
			Description: `TCP keepalive parameters

`,
		},
	}
}
