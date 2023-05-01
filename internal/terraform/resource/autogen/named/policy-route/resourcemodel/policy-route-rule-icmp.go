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

// PolicyRouteRuleIcmp describes the resource data model.
type PolicyRouteRuleIcmp struct {
	// LeafNodes
	LeafPolicyRouteRuleIcmpCode     types.String `tfsdk:"code"`
	LeafPolicyRouteRuleIcmpType     types.String `tfsdk:"type"`
	LeafPolicyRouteRuleIcmpTypeName types.String `tfsdk:"type_name"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteRuleIcmp) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route", "rule", "icmp"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteRuleIcmpCode.IsNull() || o.LeafPolicyRouteRuleIcmpCode.IsUnknown()) {
		vyosData["code"] = o.LeafPolicyRouteRuleIcmpCode.ValueString()
	}
	if !(o.LeafPolicyRouteRuleIcmpType.IsNull() || o.LeafPolicyRouteRuleIcmpType.IsUnknown()) {
		vyosData["type"] = o.LeafPolicyRouteRuleIcmpType.ValueString()
	}
	if !(o.LeafPolicyRouteRuleIcmpTypeName.IsNull() || o.LeafPolicyRouteRuleIcmpTypeName.IsUnknown()) {
		vyosData["type-name"] = o.LeafPolicyRouteRuleIcmpTypeName.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteRuleIcmp) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route", "rule", "icmp"}})

	// Leafs
	if value, ok := vyosData["code"]; ok {
		o.LeafPolicyRouteRuleIcmpCode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleIcmpCode = basetypes.NewStringNull()
	}
	if value, ok := vyosData["type"]; ok {
		o.LeafPolicyRouteRuleIcmpType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleIcmpType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["type-name"]; ok {
		o.LeafPolicyRouteRuleIcmpTypeName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleIcmpTypeName = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route", "rule", "icmp"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteRuleIcmp) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"code":      types.StringType,
		"type":      types.StringType,
		"type_name": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleIcmp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"code": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMP code (0-255)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMP code (0-255)  |

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMP type (0-255)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMP type (0-255)  |

`,
		},

		"type_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMP type-name

|  Format  |  Description  |
|----------|---------------|
|  echo-reply  |  ICMP type 0: echo-reply  |
|  destination-unreachable  |  ICMP type 3: destination-unreachable  |
|  source-quench  |  ICMP type 4: source-quench  |
|  redirect  |  ICMP type 5: redirect  |
|  echo-request  |  ICMP type 8: echo-request  |
|  router-advertisement  |  ICMP type 9: router-advertisement  |
|  router-solicitation  |  ICMP type 10: router-solicitation  |
|  time-exceeded  |  ICMP type 11: time-exceeded  |
|  parameter-problem  |  ICMP type 12: parameter-problem  |
|  timestamp-request  |  ICMP type 13: timestamp-request  |
|  timestamp-reply  |  ICMP type 14: timestamp-reply  |
|  info-request  |  ICMP type 15: info-request  |
|  info-reply  |  ICMP type 16: info-reply  |
|  address-mask-request  |  ICMP type 17: address-mask-request  |
|  address-mask-reply  |  ICMP type 18: address-mask-reply  |

`,
		},

		// TagNodes

		// Nodes

	}
}