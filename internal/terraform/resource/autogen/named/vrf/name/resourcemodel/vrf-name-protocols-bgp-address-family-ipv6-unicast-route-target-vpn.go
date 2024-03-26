// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTargetVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTargetVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTargetVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTargetVpnBoth   types.String `tfsdk:"both" vyos:"both,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTargetVpnImport types.String `tfsdk:"import" vyos:"import,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTargetVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTargetVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"both": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target both import and export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target both import and export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target import

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target import

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		// Nodes

	}
}
