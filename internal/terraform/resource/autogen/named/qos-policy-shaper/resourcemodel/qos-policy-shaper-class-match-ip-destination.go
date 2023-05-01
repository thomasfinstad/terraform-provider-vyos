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

// QosPolicyShaperClassMatchIPDestination describes the resource data model.
type QosPolicyShaperClassMatchIPDestination struct {
	// LeafNodes
	LeafQosPolicyShaperClassMatchIPDestinationAddress types.String `tfsdk:"address"`
	LeafQosPolicyShaperClassMatchIPDestinationPort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperClassMatchIPDestination) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match", "ip", "destination"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperClassMatchIPDestinationAddress.IsNull() || o.LeafQosPolicyShaperClassMatchIPDestinationAddress.IsUnknown()) {
		vyosData["address"] = o.LeafQosPolicyShaperClassMatchIPDestinationAddress.ValueString()
	}
	if !(o.LeafQosPolicyShaperClassMatchIPDestinationPort.IsNull() || o.LeafQosPolicyShaperClassMatchIPDestinationPort.IsUnknown()) {
		vyosData["port"] = o.LeafQosPolicyShaperClassMatchIPDestinationPort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperClassMatchIPDestination) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match", "ip", "destination"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafQosPolicyShaperClassMatchIPDestinationAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPDestinationAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafQosPolicyShaperClassMatchIPDestinationPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPDestinationPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match", "ip", "destination"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperClassMatchIPDestination) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatchIPDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 destination address for this match

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
|  ipv4net  |  IPv4 prefix  |

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
