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

// QosPolicyShaperHfscClassMatchIPDestination describes the resource data model.
type QosPolicyShaperHfscClassMatchIPDestination struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchIPDestinationAddress types.String `tfsdk:"address"`
	LeafQosPolicyShaperHfscClassMatchIPDestinationPort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperHfscClassMatchIPDestination) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ip", "destination"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperHfscClassMatchIPDestinationAddress.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPDestinationAddress.IsUnknown()) {
		vyosData["address"] = o.LeafQosPolicyShaperHfscClassMatchIPDestinationAddress.ValueString()
	}
	if !(o.LeafQosPolicyShaperHfscClassMatchIPDestinationPort.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPDestinationPort.IsUnknown()) {
		vyosData["port"] = o.LeafQosPolicyShaperHfscClassMatchIPDestinationPort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperHfscClassMatchIPDestination) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ip", "destination"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPDestinationAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPDestinationAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPDestinationPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPDestinationPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ip", "destination"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPDestination) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
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
