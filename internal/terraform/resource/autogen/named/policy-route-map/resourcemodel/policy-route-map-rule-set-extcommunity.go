// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRouteMapRuleSetExtcommunity describes the resource data model.
type PolicyRouteMapRuleSetExtcommunity struct {
	// LeafNodes
	PolicyRouteMapRuleSetExtcommunityBandwIDth              customtypes.CustomStringValue `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
	PolicyRouteMapRuleSetExtcommunityBandwIDthNonTransitive customtypes.CustomStringValue `tfsdk:"bandwidth_non_transitive" json:"bandwidth-non-transitive,omitempty"`
	PolicyRouteMapRuleSetExtcommunityRt                     customtypes.CustomStringValue `tfsdk:"rt" json:"rt,omitempty"`
	PolicyRouteMapRuleSetExtcommunitySoo                    customtypes.CustomStringValue `tfsdk:"soo" json:"soo,omitempty"`
	PolicyRouteMapRuleSetExtcommunityNone                   customtypes.CustomStringValue `tfsdk:"none" json:"none,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteMapRuleSetExtcommunity) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Bandwidth value in Mbps

|  Format  |  Description  |
|----------|---------------|
|  u32:1-25600  |  Bandwidth value in Mbps  |
|  cumulative  |  Cumulative bandwidth of all multipaths (outbound-only)  |
|  num-multipaths  |  Internally computed bandwidth based on number of multipaths (outbound-only)  |
`,
		},

		"bandwidth_non_transitive": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `The link bandwidth extended community is encoded as non-transitive

`,
		},

		"rt": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set route target value

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN  |  based on autonomous system number in format <0-65535:0-4294967295>  |
|  IP:NN  |  Based on a router-id IP address in format <IP:0-65535>  |
`,
		},

		"soo": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set Site of Origin value

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN  |  based on autonomous system number in format <0-65535:0-4294967295>  |
|  IP:NN  |  Based on a router-id IP address in format <IP:0-65535>  |
`,
		},

		"none": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Completely remove communities attribute from a prefix

`,
		},

		// TagNodes

		// Nodes

	}
}
