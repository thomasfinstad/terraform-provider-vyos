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

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetExtcommunity{}

// PolicyRouteMapRuleSetExtcommunity describes the resource data model.
type PolicyRouteMapRuleSetExtcommunity struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetExtcommunityBandwIDth              types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafPolicyRouteMapRuleSetExtcommunityBandwIDthNonTransitive types.Bool   `tfsdk:"bandwidth_non_transitive" vyos:"bandwidth-non-transitive,omitempty"`
	LeafPolicyRouteMapRuleSetExtcommunityRt                     types.List   `tfsdk:"rt" vyos:"rt,omitempty"`
	LeafPolicyRouteMapRuleSetExtcommunitySoo                    types.List   `tfsdk:"soo" vyos:"soo,omitempty"`
	LeafPolicyRouteMapRuleSetExtcommunityNone                   types.Bool   `tfsdk:"none" vyos:"none,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetExtcommunity) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bandwidth value in Mbps

    |  Format          |  Description                                                                  |
    |------------------|-------------------------------------------------------------------------------|
    |  1-25600         |  Bandwidth value in Mbps                                                      |
    |  cumulative      |  Cumulative bandwidth of all multipaths (outbound-only)                       |
    |  num-multipaths  |  Internally computed bandwidth based on number of multipaths (outbound-only)  |
`,
			Description: `Bandwidth value in Mbps

    |  Format          |  Description                                                                  |
    |------------------|-------------------------------------------------------------------------------|
    |  1-25600         |  Bandwidth value in Mbps                                                      |
    |  cumulative      |  Cumulative bandwidth of all multipaths (outbound-only)                       |
    |  num-multipaths  |  Internally computed bandwidth based on number of multipaths (outbound-only)  |
`,
		},

		"bandwidth_non_transitive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `The link bandwidth extended community is encoded as non-transitive

`,
			Description: `The link bandwidth extended community is encoded as non-transitive

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rt": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Set route target value

    |  Format  |  Description                                                         |
    |----------|----------------------------------------------------------------------|
    |  ASN:NN  |  based on autonomous system number in format <0-65535:0-4294967295>  |
    |  IP:NN   |  Based on a router-id IP address in format <IP:0-65535>              |
`,
			Description: `Set route target value

    |  Format  |  Description                                                         |
    |----------|----------------------------------------------------------------------|
    |  ASN:NN  |  based on autonomous system number in format <0-65535:0-4294967295>  |
    |  IP:NN   |  Based on a router-id IP address in format <IP:0-65535>              |
`,
		},

		"soo": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Set Site of Origin value

    |  Format  |  Description                                                         |
    |----------|----------------------------------------------------------------------|
    |  ASN:NN  |  based on autonomous system number in format <0-65535:0-4294967295>  |
    |  IP:NN   |  Based on a router-id IP address in format <IP:0-65535>              |
`,
			Description: `Set Site of Origin value

    |  Format  |  Description                                                         |
    |----------|----------------------------------------------------------------------|
    |  ASN:NN  |  based on autonomous system number in format <0-65535:0-4294967295>  |
    |  IP:NN   |  Based on a router-id IP address in format <IP:0-65535>              |
`,
		},

		"none": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Completely remove communities attribute from a prefix

`,
			Description: `Completely remove communities attribute from a prefix

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
