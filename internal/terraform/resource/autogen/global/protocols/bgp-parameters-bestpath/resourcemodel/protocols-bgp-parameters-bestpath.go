// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpParametersBestpath describes the resource data model.
type ProtocolsBgpParametersBestpath struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsBgpParametersBestpathBandwIDth       types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafProtocolsBgpParametersBestpathCompareRouterID types.Bool   `tfsdk:"compare_routerid" vyos:"compare-routerid,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsBgpParametersBestpathAsPath   bool `tfsdk:"-" vyos:"as-path,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersBestpathMed      bool `tfsdk:"-" vyos:"med,ignore,omitempty"`
	ExistsNodeProtocolsBgpParametersBestpathPeerType bool `tfsdk:"-" vyos:"peer-type,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpParametersBestpath) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpParametersBestpath) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",

		"parameters",

		"bestpath",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpParametersBestpath) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Link Bandwidth attribute

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  default-weight-for-missing  &emsp; |  Assign low default weight (1) to paths not having link bandwidth  |
    |  ignore  &emsp; |  Ignore link bandwidth (do regular ECMP, not weighted)  |
    |  skip-missing  &emsp; |  Ignore paths without link bandwidth for ECMP (if other paths have it)  |

`,
		},

		"compare_routerid": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Compare the router-id for identical EBGP paths

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
