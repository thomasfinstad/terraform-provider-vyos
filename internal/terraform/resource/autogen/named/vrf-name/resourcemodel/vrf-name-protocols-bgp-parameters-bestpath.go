// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpParametersBestpath describes the resource data model.
type VrfNameProtocolsBgpParametersBestpath struct {
	// LeafNodes
	VrfNameProtocolsBgpParametersBestpathBandwIDth       customtypes.CustomStringValue `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
	VrfNameProtocolsBgpParametersBestpathCompareRouterID customtypes.CustomStringValue `tfsdk:"compare_routerid" json:"compare-routerid,omitempty"`

	// TagNodes

	// Nodes
	VrfNameProtocolsBgpParametersBestpathAsPath   types.Object `tfsdk:"as_path" json:"as-path,omitempty"`
	VrfNameProtocolsBgpParametersBestpathMed      types.Object `tfsdk:"med" json:"med,omitempty"`
	VrfNameProtocolsBgpParametersBestpathPeerType types.Object `tfsdk:"peer_type" json:"peer-type,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersBestpath) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Link Bandwidth attribute

|  Format  |  Description  |
|----------|---------------|
|  default-weight-for-missing  |  Assign low default weight (1) to paths not having link bandwidth  |
|  ignore  |  Ignore link bandwidth (do regular ECMP, not weighted)  |
|  skip-missing  |  Ignore paths without link bandwidth for ECMP (if other paths have it)  |
`,
		},

		"compare_routerid": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Compare the router-id for identical EBGP paths

`,
		},

		// TagNodes

		// Nodes

		"as_path": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersBestpathAsPath{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `AS-path attribute comparison parameters

`,
		},

		"med": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersBestpathMed{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `MED attribute comparison parameters

`,
		},

		"peer_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersBestpathPeerType{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Peer type

`,
		},
	}
}
