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

// PolicyRoutesixRuleSet describes the resource data model.
type PolicyRoutesixRuleSet struct {
	// LeafNodes
	LeafPolicyRoutesixRuleSetConnectionMark types.String `tfsdk:"connection_mark"`
	LeafPolicyRoutesixRuleSetDscp           types.String `tfsdk:"dscp"`
	LeafPolicyRoutesixRuleSetMark           types.String `tfsdk:"mark"`
	LeafPolicyRoutesixRuleSetTable          types.String `tfsdk:"table"`
	LeafPolicyRoutesixRuleSetTCPMss         types.String `tfsdk:"tcp_mss"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRoutesixRuleSet) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "set"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRoutesixRuleSetConnectionMark.IsNull() || o.LeafPolicyRoutesixRuleSetConnectionMark.IsUnknown()) {
		vyosData["connection-mark"] = o.LeafPolicyRoutesixRuleSetConnectionMark.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleSetDscp.IsNull() || o.LeafPolicyRoutesixRuleSetDscp.IsUnknown()) {
		vyosData["dscp"] = o.LeafPolicyRoutesixRuleSetDscp.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleSetMark.IsNull() || o.LeafPolicyRoutesixRuleSetMark.IsUnknown()) {
		vyosData["mark"] = o.LeafPolicyRoutesixRuleSetMark.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleSetTable.IsNull() || o.LeafPolicyRoutesixRuleSetTable.IsUnknown()) {
		vyosData["table"] = o.LeafPolicyRoutesixRuleSetTable.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleSetTCPMss.IsNull() || o.LeafPolicyRoutesixRuleSetTCPMss.IsUnknown()) {
		vyosData["tcp-mss"] = o.LeafPolicyRoutesixRuleSetTCPMss.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRoutesixRuleSet) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "set"}})

	// Leafs
	if value, ok := vyosData["connection-mark"]; ok {
		o.LeafPolicyRoutesixRuleSetConnectionMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSetConnectionMark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dscp"]; ok {
		o.LeafPolicyRoutesixRuleSetDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSetDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mark"]; ok {
		o.LeafPolicyRoutesixRuleSetMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSetMark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["table"]; ok {
		o.LeafPolicyRoutesixRuleSetTable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSetTable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["tcp-mss"]; ok {
		o.LeafPolicyRoutesixRuleSetTCPMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleSetTCPMss = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "set"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRoutesixRuleSet) AttributeTypes() map[string]attr.Type {
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
func (o PolicyRoutesixRuleSet) ResourceSchemaAttributes() map[string]schema.Attribute {
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