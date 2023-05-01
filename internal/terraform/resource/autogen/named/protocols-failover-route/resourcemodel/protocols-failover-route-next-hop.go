// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ProtocolsFailoverRouteNextHop describes the resource data model.
type ProtocolsFailoverRouteNextHop struct {
	// LeafNodes
	LeafProtocolsFailoverRouteNextHopInterface types.String `tfsdk:"interface"`
	LeafProtocolsFailoverRouteNextHopMetric    types.String `tfsdk:"metric"`

	// TagNodes

	// Nodes
	NodeProtocolsFailoverRouteNextHopCheck types.Object `tfsdk:"check"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsFailoverRouteNextHop) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "failover", "route", "next-hop"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsFailoverRouteNextHopInterface.IsNull() || o.LeafProtocolsFailoverRouteNextHopInterface.IsUnknown()) {
		vyosData["interface"] = o.LeafProtocolsFailoverRouteNextHopInterface.ValueString()
	}
	if !(o.LeafProtocolsFailoverRouteNextHopMetric.IsNull() || o.LeafProtocolsFailoverRouteNextHopMetric.IsUnknown()) {
		vyosData["metric"] = o.LeafProtocolsFailoverRouteNextHopMetric.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeProtocolsFailoverRouteNextHopCheck.IsNull() || o.NodeProtocolsFailoverRouteNextHopCheck.IsUnknown()) {
		var subModel ProtocolsFailoverRouteNextHopCheck
		diags.Append(o.NodeProtocolsFailoverRouteNextHopCheck.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["check"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsFailoverRouteNextHop) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "failover", "route", "next-hop"}})

	// Leafs
	if value, ok := vyosData["interface"]; ok {
		o.LeafProtocolsFailoverRouteNextHopInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsFailoverRouteNextHopInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["metric"]; ok {
		o.LeafProtocolsFailoverRouteNextHopMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsFailoverRouteNextHopMetric = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["check"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsFailoverRouteNextHopCheck{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsFailoverRouteNextHopCheck = data

	} else {
		o.NodeProtocolsFailoverRouteNextHopCheck = basetypes.NewObjectNull(ProtocolsFailoverRouteNextHopCheck{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "failover", "route", "next-hop"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsFailoverRouteNextHop) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"interface": types.StringType,
		"metric":    types.StringType,

		// Tags

		// Nodes
		"check": types.ObjectType{AttrTypes: ProtocolsFailoverRouteNextHopCheck{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsFailoverRouteNextHop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway interface name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Gateway interface name  |

`,
		},

		"metric": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route metric for this gateway

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Route metric  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"check": schema.SingleNestedAttribute{
			Attributes: ProtocolsFailoverRouteNextHopCheck{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Check target options

`,
		},
	}
}