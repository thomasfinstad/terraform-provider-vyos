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

// QosPolicyLimiterClassMatchIPTCP describes the resource data model.
type QosPolicyLimiterClassMatchIPTCP struct {
	// LeafNodes
	LeafQosPolicyLimiterClassMatchIPTCPAck types.String `tfsdk:"ack"`
	LeafQosPolicyLimiterClassMatchIPTCPSyn types.String `tfsdk:"syn"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyLimiterClassMatchIPTCP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "limiter", "class", "match", "ip", "tcp"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyLimiterClassMatchIPTCPAck.IsNull() || o.LeafQosPolicyLimiterClassMatchIPTCPAck.IsUnknown()) {
		vyosData["ack"] = o.LeafQosPolicyLimiterClassMatchIPTCPAck.ValueString()
	}
	if !(o.LeafQosPolicyLimiterClassMatchIPTCPSyn.IsNull() || o.LeafQosPolicyLimiterClassMatchIPTCPSyn.IsUnknown()) {
		vyosData["syn"] = o.LeafQosPolicyLimiterClassMatchIPTCPSyn.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyLimiterClassMatchIPTCP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "limiter", "class", "match", "ip", "tcp"}})

	// Leafs
	if value, ok := vyosData["ack"]; ok {
		o.LeafQosPolicyLimiterClassMatchIPTCPAck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyLimiterClassMatchIPTCPAck = basetypes.NewStringNull()
	}
	if value, ok := vyosData["syn"]; ok {
		o.LeafQosPolicyLimiterClassMatchIPTCPSyn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyLimiterClassMatchIPTCPSyn = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "limiter", "class", "match", "ip", "tcp"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyLimiterClassMatchIPTCP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"ack": types.StringType,
		"syn": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiterClassMatchIPTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ack": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP ACK

`,
		},

		"syn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP SYN

`,
		},

		// TagNodes

		// Nodes

	}
}