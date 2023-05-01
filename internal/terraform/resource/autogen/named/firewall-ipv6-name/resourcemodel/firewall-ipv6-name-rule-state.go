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

// FirewallIPvsixNameRuleState describes the resource data model.
type FirewallIPvsixNameRuleState struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleStateEstablished types.String `tfsdk:"established"`
	LeafFirewallIPvsixNameRuleStateInvalID     types.String `tfsdk:"invalid"`
	LeafFirewallIPvsixNameRuleStateNew         types.String `tfsdk:"new"`
	LeafFirewallIPvsixNameRuleStateRelated     types.String `tfsdk:"related"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallIPvsixNameRuleState) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "state"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallIPvsixNameRuleStateEstablished.IsNull() || o.LeafFirewallIPvsixNameRuleStateEstablished.IsUnknown()) {
		vyosData["established"] = o.LeafFirewallIPvsixNameRuleStateEstablished.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleStateInvalID.IsNull() || o.LeafFirewallIPvsixNameRuleStateInvalID.IsUnknown()) {
		vyosData["invalid"] = o.LeafFirewallIPvsixNameRuleStateInvalID.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleStateNew.IsNull() || o.LeafFirewallIPvsixNameRuleStateNew.IsUnknown()) {
		vyosData["new"] = o.LeafFirewallIPvsixNameRuleStateNew.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleStateRelated.IsNull() || o.LeafFirewallIPvsixNameRuleStateRelated.IsUnknown()) {
		vyosData["related"] = o.LeafFirewallIPvsixNameRuleStateRelated.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallIPvsixNameRuleState) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "state"}})

	// Leafs
	if value, ok := vyosData["established"]; ok {
		o.LeafFirewallIPvsixNameRuleStateEstablished = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleStateEstablished = basetypes.NewStringNull()
	}
	if value, ok := vyosData["invalid"]; ok {
		o.LeafFirewallIPvsixNameRuleStateInvalID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleStateInvalID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["new"]; ok {
		o.LeafFirewallIPvsixNameRuleStateNew = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleStateNew = basetypes.NewStringNull()
	}
	if value, ok := vyosData["related"]; ok {
		o.LeafFirewallIPvsixNameRuleStateRelated = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleStateRelated = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "state"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallIPvsixNameRuleState) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"established": types.StringType,
		"invalid":     types.StringType,
		"new":         types.StringType,
		"related":     types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleState) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"established": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Established state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"invalid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Invalid state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"new": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `New state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"related": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Related state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		// TagNodes

		// Nodes

	}
}