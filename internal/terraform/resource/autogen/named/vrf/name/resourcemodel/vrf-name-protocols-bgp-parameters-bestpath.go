// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersBestpath{}

// VrfNameProtocolsBgpParametersBestpath describes the resource data model.
type VrfNameProtocolsBgpParametersBestpath struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersBestpathBandwIDth       types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafVrfNameProtocolsBgpParametersBestpathCompareRouterID types.Bool   `tfsdk:"compare_routerid" vyos:"compare-routerid,omitempty"`
	LeafVrfNameProtocolsBgpParametersBestpathMed             types.List   `tfsdk:"med" vyos:"med,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpParametersBestpathAsPath   *VrfNameProtocolsBgpParametersBestpathAsPath   `tfsdk:"as_path" vyos:"as-path,omitempty"`
	NodeVrfNameProtocolsBgpParametersBestpathPeerType *VrfNameProtocolsBgpParametersBestpathPeerType `tfsdk:"peer_type" vyos:"peer-type,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersBestpath) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Link Bandwidth attribute

    |  Format                      |  Description                                                            |
    |------------------------------|-------------------------------------------------------------------------|
    |  default-weight-for-missing  |  Assign low default weight (1) to paths not having link bandwidth       |
    |  ignore                      |  Ignore link bandwidth (do regular ECMP, not weighted)                  |
    |  skip-missing                |  Ignore paths without link bandwidth for ECMP (if other paths have it)  |
`,
			Description: `Link Bandwidth attribute

    |  Format                      |  Description                                                            |
    |------------------------------|-------------------------------------------------------------------------|
    |  default-weight-for-missing  |  Assign low default weight (1) to paths not having link bandwidth       |
    |  ignore                      |  Ignore link bandwidth (do regular ECMP, not weighted)                  |
    |  skip-missing                |  Ignore paths without link bandwidth for ECMP (if other paths have it)  |
`,
		},

		"compare_routerid": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Compare the router-id for identical EBGP paths

`,
			Description: `Compare the router-id for identical EBGP paths

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"med": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `MED attribute comparison parameters

    |  Format            |  Description                                              |
    |--------------------|-----------------------------------------------------------|
    |  confed            |  Compare MEDs among confederation paths                   |
    |  missing-as-worst  |  Treat missing route as a MED as the least preferred one  |
`,
			Description: `MED attribute comparison parameters

    |  Format            |  Description                                              |
    |--------------------|-----------------------------------------------------------|
    |  confed            |  Compare MEDs among confederation paths                   |
    |  missing-as-worst  |  Treat missing route as a MED as the least preferred one  |
`,
		},

		// Nodes

		"as_path": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersBestpathAsPath{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `AS-path attribute comparison parameters

`,
			Description: `AS-path attribute comparison parameters

`,
		},

		"peer_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersBestpathPeerType{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Peer type

`,
			Description: `Peer type

`,
		},
	}
}
