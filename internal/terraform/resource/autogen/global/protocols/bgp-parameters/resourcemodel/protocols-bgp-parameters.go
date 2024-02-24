// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpParameters describes the resource data model.
type ProtocolsBgpParameters struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsBgpParametersAlwaysCompareMed                  types.Bool   `tfsdk:"always_compare_med" vyos:"always-compare-med,omitempty"`
	LeafProtocolsBgpParametersClusterID                         types.String `tfsdk:"cluster_id" vyos:"cluster-id,omitempty"`
	LeafProtocolsBgpParametersDeterministicMed                  types.Bool   `tfsdk:"deterministic_med" vyos:"deterministic-med,omitempty"`
	LeafProtocolsBgpParametersEbgpRequiresPolicy                types.Bool   `tfsdk:"ebgp_requires_policy" vyos:"ebgp-requires-policy,omitempty"`
	LeafProtocolsBgpParametersFastConvergence                   types.Bool   `tfsdk:"fast_convergence" vyos:"fast-convergence,omitempty"`
	LeafProtocolsBgpParametersGracefulShutdown                  types.Bool   `tfsdk:"graceful_shutdown" vyos:"graceful-shutdown,omitempty"`
	LeafProtocolsBgpParametersLogNeighborChanges                types.Bool   `tfsdk:"log_neighbor_changes" vyos:"log-neighbor-changes,omitempty"`
	LeafProtocolsBgpParametersMinimumHoldtime                   types.Number `tfsdk:"minimum_holdtime" vyos:"minimum-holdtime,omitempty"`
	LeafProtocolsBgpParametersNetworkImportCheck                types.Bool   `tfsdk:"network_import_check" vyos:"network-import-check,omitempty"`
	LeafProtocolsBgpParametersRouteReflectorAllowOutboundPolicy types.Bool   `tfsdk:"route_reflector_allow_outbound_policy" vyos:"route-reflector-allow-outbound-policy,omitempty"`
	LeafProtocolsBgpParametersNoClientToClientReflection        types.Bool   `tfsdk:"no_client_to_client_reflection" vyos:"no-client-to-client-reflection,omitempty"`
	LeafProtocolsBgpParametersNoFastExternalFailover            types.Bool   `tfsdk:"no_fast_external_failover" vyos:"no-fast-external-failover,omitempty"`
	LeafProtocolsBgpParametersNoSuppressDuplicates              types.Bool   `tfsdk:"no_suppress_duplicates" vyos:"no-suppress-duplicates,omitempty"`
	LeafProtocolsBgpParametersRejectAsSets                      types.Bool   `tfsdk:"reject_as_sets" vyos:"reject-as-sets,omitempty"`
	LeafProtocolsBgpParametersShutdown                          types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafProtocolsBgpParametersSuppressFibPending                types.Bool   `tfsdk:"suppress_fib_pending" vyos:"suppress-fib-pending,omitempty"`
	LeafProtocolsBgpParametersRouterID                          types.String `tfsdk:"router_id" vyos:"router-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsBgpParametersBestpath                 bool `tfsdk:"-" vyos:"bestpath,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersConfederation            bool `tfsdk:"-" vyos:"confederation,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersConditionalAdvertisement bool `tfsdk:"-" vyos:"conditional-advertisement,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersDampening                bool `tfsdk:"-" vyos:"dampening,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersDefault                  bool `tfsdk:"-" vyos:"default,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersDistance                 bool `tfsdk:"-" vyos:"distance,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersGracefulRestart          bool `tfsdk:"-" vyos:"graceful-restart,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpParameters) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpParameters) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",

		"parameters",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpParameters) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"always_compare_med": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Always compare MEDs from different neighbors

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"cluster_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-reflector cluster-id

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Route-reflector cluster-id  |

`,
		},

		"deterministic_med": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Compare MEDs between different peers in the same AS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ebgp_requires_policy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Require in and out policy for eBGP peers (RFC8212)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fast_convergence": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Teardown sessions immediately whenever peer becomes unreachable

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"graceful_shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Graceful shutdown

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"log_neighbor_changes": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log neighbor up/down changes and reset reason

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"minimum_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP minimum holdtime

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Minimum holdtime in seconds  |

`,
		},

		"network_import_check": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable IGP route check for network statements

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_reflector_allow_outbound_policy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Route reflector client allow policy outbound

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_client_to_client_reflection": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable client to client route reflection

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_fast_external_failover": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable immediate session reset on peer link down event

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_suppress_duplicates": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable suppress duplicate updates if the route actually not changed

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"reject_as_sets": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Reject routes with AS_SET or AS_CONFED_SET flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administrative shutdown of the BGP instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"suppress_fib_pending": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise only routes that are programmed in kernel to peers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"router_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override default router identifier

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Router-ID in IP address format  |

`,
		},
	}
}
