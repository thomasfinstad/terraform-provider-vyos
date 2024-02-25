// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsis describes the resource data model.
type VrfNameProtocolsIsis struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisDynamicHostname     types.Bool   `tfsdk:"dynamic_hostname" vyos:"dynamic-hostname,omitempty"`
	LeafVrfNameProtocolsIsisLevel               types.String `tfsdk:"level" vyos:"level,omitempty"`
	LeafVrfNameProtocolsIsisLogAdjacencyChanges types.Bool   `tfsdk:"log_adjacency_changes" vyos:"log-adjacency-changes,omitempty"`
	LeafVrfNameProtocolsIsisLspGenInterval      types.Number `tfsdk:"lsp_gen_interval" vyos:"lsp-gen-interval,omitempty"`
	LeafVrfNameProtocolsIsisLspMtu              types.Number `tfsdk:"lsp_mtu" vyos:"lsp-mtu,omitempty"`
	LeafVrfNameProtocolsIsisLspRefreshInterval  types.Number `tfsdk:"lsp_refresh_interval" vyos:"lsp-refresh-interval,omitempty"`
	LeafVrfNameProtocolsIsisMaxLspLifetime      types.Number `tfsdk:"max_lsp_lifetime" vyos:"max-lsp-lifetime,omitempty"`
	LeafVrfNameProtocolsIsisMetricStyle         types.String `tfsdk:"metric_style" vyos:"metric-style,omitempty"`
	LeafVrfNameProtocolsIsisNet                 types.String `tfsdk:"net" vyos:"net,omitempty"`
	LeafVrfNameProtocolsIsisPurgeOriginator     types.Bool   `tfsdk:"purge_originator" vyos:"purge-originator,omitempty"`
	LeafVrfNameProtocolsIsisSetAttachedBit      types.Bool   `tfsdk:"set_attached_bit" vyos:"set-attached-bit,omitempty"`
	LeafVrfNameProtocolsIsisSetOverloadBit      types.Bool   `tfsdk:"set_overload_bit" vyos:"set-overload-bit,omitempty"`
	LeafVrfNameProtocolsIsisSpfInterval         types.Number `tfsdk:"spf_interval" vyos:"spf-interval,omitempty"`
	LeafVrfNameProtocolsIsisRouteMap            types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsIsisInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
	NodeVrfNameProtocolsIsisAreaPassword       *VrfNameProtocolsIsisAreaPassword       `tfsdk:"area_password" vyos:"area-password,omitempty"`
	NodeVrfNameProtocolsIsisDefaultInformation *VrfNameProtocolsIsisDefaultInformation `tfsdk:"default_information" vyos:"default-information,omitempty"`
	NodeVrfNameProtocolsIsisDomainPassword     *VrfNameProtocolsIsisDomainPassword     `tfsdk:"domain_password" vyos:"domain-password,omitempty"`
	NodeVrfNameProtocolsIsisTrafficEngineering *VrfNameProtocolsIsisTrafficEngineering `tfsdk:"traffic_engineering" vyos:"traffic-engineering,omitempty"`
	NodeVrfNameProtocolsIsisSegmentRouting     *VrfNameProtocolsIsisSegmentRouting     `tfsdk:"segment_routing" vyos:"segment-routing,omitempty"`
	NodeVrfNameProtocolsIsisRedistribute       *VrfNameProtocolsIsisRedistribute       `tfsdk:"redistribute" vyos:"redistribute,omitempty"`
	NodeVrfNameProtocolsIsisSpfDelayIetf       *VrfNameProtocolsIsisSpfDelayIetf       `tfsdk:"spf_delay_ietf" vyos:"spf-delay-ietf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsis) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dynamic_hostname": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Dynamic hostname for IS-IS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IS-IS level number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  level-1  &emsp; |  Act as a station router  |
    |  level-1-2  &emsp; |  Act as both a station and an area router  |
    |  level-2  &emsp; |  Act as an area router  |

`,
		},

		"log_adjacency_changes": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log adjacency state changes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lsp_gen_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between regenerating same LSP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-120  &emsp; |  Minimum interval in seconds  |

`,
		},

		"lsp_mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Configure the maximum size of generated LSPs

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 128-4352  &emsp; |  Maximum size of generated LSPs  |

`,

			// Default:          stringdefault.StaticString(`1497`),
			Computed: true,
		},

		"lsp_refresh_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `LSP refresh interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65235  &emsp; |  LSP refresh interval in seconds  |

`,
		},

		"max_lsp_lifetime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum LSP lifetime

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 350-65535  &emsp; |  LSP lifetime in seconds  |

`,
		},

		"metric_style": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use old-style (ISO 10589) or new-style packet formats

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  narrow  &emsp; |  Use old style of TLVs with narrow metric  |
    |  transition  &emsp; |  Send and accept both styles of TLVs during transition  |
    |  wide  &emsp; |  Use new style of TLVs to carry wider metric  |

`,
		},

		"net": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `A Network Entity Title for this process (ISO only)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  XX.XXXX. ... .XXX.XX  &emsp; |  Network entity title (NET)  |

`,
		},

		"purge_originator": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use the RFC 6232 purge-originator

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"set_attached_bit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set attached bit to identify as L1/L2 router for inter-area traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"set_overload_bit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set overload bit to avoid any transit traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"spf_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between SPF calculations

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-120  &emsp; |  Interval in seconds  |

`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		// Nodes

		"area_password": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisAreaPassword{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Configure the authentication password for an area

`,
		},

		"default_information": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformation{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Control distribution of default information

`,
		},

		"domain_password": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDomainPassword{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Set the authentication password for a routing domain

`,
		},

		"traffic_engineering": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisTrafficEngineering{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Show IS-IS neighbor adjacencies

`,
		},

		"segment_routing": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRouting{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Segment-Routing (SPRING) settings

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistribute{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute information from another routing protocol

`,
		},

		"spf_delay_ietf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSpfDelayIetf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IETF SPF delay algorithm

`,
		},
	}
}
