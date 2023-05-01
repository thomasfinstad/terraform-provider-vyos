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

// QosPolicyRoundRobinClassMatchIPvsixSource describes the resource data model.
type QosPolicyRoundRobinClassMatchIPvsixSource struct {
	// LeafNodes
	LeafQosPolicyRoundRobinClassMatchIPvsixSourceAddress types.String `tfsdk:"address"`
	LeafQosPolicyRoundRobinClassMatchIPvsixSourcePort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyRoundRobinClassMatchIPvsixSource) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "round-robin", "class", "match", "ipv6", "source"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyRoundRobinClassMatchIPvsixSourceAddress.IsNull() || o.LeafQosPolicyRoundRobinClassMatchIPvsixSourceAddress.IsUnknown()) {
		vyosData["address"] = o.LeafQosPolicyRoundRobinClassMatchIPvsixSourceAddress.ValueString()
	}
	if !(o.LeafQosPolicyRoundRobinClassMatchIPvsixSourcePort.IsNull() || o.LeafQosPolicyRoundRobinClassMatchIPvsixSourcePort.IsUnknown()) {
		vyosData["port"] = o.LeafQosPolicyRoundRobinClassMatchIPvsixSourcePort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyRoundRobinClassMatchIPvsixSource) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "round-robin", "class", "match", "ipv6", "source"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafQosPolicyRoundRobinClassMatchIPvsixSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRoundRobinClassMatchIPvsixSourceAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafQosPolicyRoundRobinClassMatchIPvsixSourcePort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRoundRobinClassMatchIPvsixSourcePort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "round-robin", "class", "match", "ipv6", "source"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyRoundRobinClassMatchIPvsixSource) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinClassMatchIPvsixSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 destination address for this match

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,
		},

		// TagNodes

		// Nodes

	}
}