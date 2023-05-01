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

// PolicyRouteRuleSet describes the resource data model.
type PolicyRouteRuleSet struct {
	// LeafNodes
	LeafPolicyRouteRuleSetConnectionMark types.String `tfsdk:"connection_mark"`
	LeafPolicyRouteRuleSetDscp           types.String `tfsdk:"dscp"`
	LeafPolicyRouteRuleSetMark           types.String `tfsdk:"mark"`
	LeafPolicyRouteRuleSetTable          types.String `tfsdk:"table"`
	LeafPolicyRouteRuleSetTCPMss         types.String `tfsdk:"tcp_mss"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteRuleSet) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route", "rule", "set"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteRuleSetConnectionMark.IsNull() || o.LeafPolicyRouteRuleSetConnectionMark.IsUnknown()) {
		vyosData["connection-mark"] = o.LeafPolicyRouteRuleSetConnectionMark.ValueString()
	}
	if !(o.LeafPolicyRouteRuleSetDscp.IsNull() || o.LeafPolicyRouteRuleSetDscp.IsUnknown()) {
		vyosData["dscp"] = o.LeafPolicyRouteRuleSetDscp.ValueString()
	}
	if !(o.LeafPolicyRouteRuleSetMark.IsNull() || o.LeafPolicyRouteRuleSetMark.IsUnknown()) {
		vyosData["mark"] = o.LeafPolicyRouteRuleSetMark.ValueString()
	}
	if !(o.LeafPolicyRouteRuleSetTable.IsNull() || o.LeafPolicyRouteRuleSetTable.IsUnknown()) {
		vyosData["table"] = o.LeafPolicyRouteRuleSetTable.ValueString()
	}
	if !(o.LeafPolicyRouteRuleSetTCPMss.IsNull() || o.LeafPolicyRouteRuleSetTCPMss.IsUnknown()) {
		vyosData["tcp-mss"] = o.LeafPolicyRouteRuleSetTCPMss.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteRuleSet) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route", "rule", "set"}})

	// Leafs
	if value, ok := vyosData["connection-mark"]; ok {
		o.LeafPolicyRouteRuleSetConnectionMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSetConnectionMark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dscp"]; ok {
		o.LeafPolicyRouteRuleSetDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSetDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mark"]; ok {
		o.LeafPolicyRouteRuleSetMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSetMark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["table"]; ok {
		o.LeafPolicyRouteRuleSetTable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSetTable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["tcp-mss"]; ok {
		o.LeafPolicyRouteRuleSetTCPMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleSetTCPMss = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route", "rule", "set"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteRuleSet) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"connection_mark": types.StringType,
		"dscp":            types.StringType,
		"mark":            types.StringType,
		"table":           types.StringType,
		"tcp_mss":         types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleSet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"connection_mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Connection marking

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection marking  |

`,
		},

		"dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Packet Differentiated Services Codepoint (DSCP)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP number  |

`,
		},

		"mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Packet marking

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Packet marking  |

`,
		},

		"table": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Routing table to forward packet with

|  Format  |  Description  |
|----------|---------------|
|  u32:1-200  |  Table number  |
|  main  |  Main table  |

`,
		},

		"tcp_mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TCP Maximum Segment Size

|  Format  |  Description  |
|----------|---------------|
|  u32:500-1460  |  Explicitly set TCP MSS value  |

`,
		},

		// TagNodes

		// Nodes

	}
}