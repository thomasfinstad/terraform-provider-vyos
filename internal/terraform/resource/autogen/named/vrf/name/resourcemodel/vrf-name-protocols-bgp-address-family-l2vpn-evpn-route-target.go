// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTargetBoth   types.List `tfsdk:"both" vyos:"both,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTargetImport types.List `tfsdk:"import" vyos:"import,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTargetExport types.List `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"both": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route Target both import and export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target both import and export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"import": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route Target import

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target import

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"export": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route Target export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		// Nodes

	}
}
