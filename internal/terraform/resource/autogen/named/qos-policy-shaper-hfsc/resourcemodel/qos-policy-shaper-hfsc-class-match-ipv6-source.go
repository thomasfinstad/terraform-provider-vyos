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

// QosPolicyShaperHfscClassMatchIPvsixSource describes the resource data model.
type QosPolicyShaperHfscClassMatchIPvsixSource struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchIPvsixSourceAddress types.String `tfsdk:"address"`
	LeafQosPolicyShaperHfscClassMatchIPvsixSourcePort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperHfscClassMatchIPvsixSource) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ipv6", "source"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperHfscClassMatchIPvsixSourceAddress.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPvsixSourceAddress.IsUnknown()) {
		vyosData["address"] = o.LeafQosPolicyShaperHfscClassMatchIPvsixSourceAddress.ValueString()
	}
	if !(o.LeafQosPolicyShaperHfscClassMatchIPvsixSourcePort.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPvsixSourcePort.IsUnknown()) {
		vyosData["port"] = o.LeafQosPolicyShaperHfscClassMatchIPvsixSourcePort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperHfscClassMatchIPvsixSource) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ipv6", "source"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixSourceAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixSourcePort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixSourcePort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ipv6", "source"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPvsixSource) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPvsixSource) ResourceSchemaAttributes() map[string]schema.Attribute {
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
