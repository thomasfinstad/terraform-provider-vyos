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

// QosPolicyPriorityQueueClassMatchIPvsixTCP describes the resource data model.
type QosPolicyPriorityQueueClassMatchIPvsixTCP struct {
	// LeafNodes
	LeafQosPolicyPriorityQueueClassMatchIPvsixTCPAck types.String `tfsdk:"ack"`
	LeafQosPolicyPriorityQueueClassMatchIPvsixTCPSyn types.String `tfsdk:"syn"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyPriorityQueueClassMatchIPvsixTCP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "priority-queue", "class", "match", "ipv6", "tcp"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPAck.IsNull() || o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPAck.IsUnknown()) {
		vyosData["ack"] = o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPAck.ValueString()
	}
	if !(o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPSyn.IsNull() || o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPSyn.IsUnknown()) {
		vyosData["syn"] = o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPSyn.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyPriorityQueueClassMatchIPvsixTCP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "priority-queue", "class", "match", "ipv6", "tcp"}})

	// Leafs
	if value, ok := vyosData["ack"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPAck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPAck = basetypes.NewStringNull()
	}
	if value, ok := vyosData["syn"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPSyn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixTCPSyn = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "priority-queue", "class", "match", "ipv6", "tcp"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyPriorityQueueClassMatchIPvsixTCP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"ack": types.StringType,
		"syn": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueClassMatchIPvsixTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
