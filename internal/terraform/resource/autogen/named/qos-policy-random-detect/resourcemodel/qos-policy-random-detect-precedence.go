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

// QosPolicyRandomDetectPrecedence describes the resource data model.
type QosPolicyRandomDetectPrecedence struct {
	// LeafNodes
	LeafQosPolicyRandomDetectPrecedenceQueueLimit       types.String `tfsdk:"queue_limit"`
	LeafQosPolicyRandomDetectPrecedenceAveragePacket    types.String `tfsdk:"average_packet"`
	LeafQosPolicyRandomDetectPrecedenceMarkProbability  types.String `tfsdk:"mark_probability"`
	LeafQosPolicyRandomDetectPrecedenceMaximumThreshold types.String `tfsdk:"maximum_threshold"`
	LeafQosPolicyRandomDetectPrecedenceMinimumThreshold types.String `tfsdk:"minimum_threshold"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyRandomDetectPrecedence) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "random-detect", "precedence"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyRandomDetectPrecedenceQueueLimit.IsNull() || o.LeafQosPolicyRandomDetectPrecedenceQueueLimit.IsUnknown()) {
		vyosData["queue-limit"] = o.LeafQosPolicyRandomDetectPrecedenceQueueLimit.ValueString()
	}
	if !(o.LeafQosPolicyRandomDetectPrecedenceAveragePacket.IsNull() || o.LeafQosPolicyRandomDetectPrecedenceAveragePacket.IsUnknown()) {
		vyosData["average-packet"] = o.LeafQosPolicyRandomDetectPrecedenceAveragePacket.ValueString()
	}
	if !(o.LeafQosPolicyRandomDetectPrecedenceMarkProbability.IsNull() || o.LeafQosPolicyRandomDetectPrecedenceMarkProbability.IsUnknown()) {
		vyosData["mark-probability"] = o.LeafQosPolicyRandomDetectPrecedenceMarkProbability.ValueString()
	}
	if !(o.LeafQosPolicyRandomDetectPrecedenceMaximumThreshold.IsNull() || o.LeafQosPolicyRandomDetectPrecedenceMaximumThreshold.IsUnknown()) {
		vyosData["maximum-threshold"] = o.LeafQosPolicyRandomDetectPrecedenceMaximumThreshold.ValueString()
	}
	if !(o.LeafQosPolicyRandomDetectPrecedenceMinimumThreshold.IsNull() || o.LeafQosPolicyRandomDetectPrecedenceMinimumThreshold.IsUnknown()) {
		vyosData["minimum-threshold"] = o.LeafQosPolicyRandomDetectPrecedenceMinimumThreshold.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyRandomDetectPrecedence) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "random-detect", "precedence"}})

	// Leafs
	if value, ok := vyosData["queue-limit"]; ok {
		o.LeafQosPolicyRandomDetectPrecedenceQueueLimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRandomDetectPrecedenceQueueLimit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["average-packet"]; ok {
		o.LeafQosPolicyRandomDetectPrecedenceAveragePacket = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRandomDetectPrecedenceAveragePacket = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mark-probability"]; ok {
		o.LeafQosPolicyRandomDetectPrecedenceMarkProbability = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRandomDetectPrecedenceMarkProbability = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-threshold"]; ok {
		o.LeafQosPolicyRandomDetectPrecedenceMaximumThreshold = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRandomDetectPrecedenceMaximumThreshold = basetypes.NewStringNull()
	}
	if value, ok := vyosData["minimum-threshold"]; ok {
		o.LeafQosPolicyRandomDetectPrecedenceMinimumThreshold = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRandomDetectPrecedenceMinimumThreshold = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "random-detect", "precedence"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyRandomDetectPrecedence) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"queue_limit":       types.StringType,
		"average_packet":    types.StringType,
		"mark_probability":  types.StringType,
		"maximum_threshold": types.StringType,
		"minimum_threshold": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRandomDetectPrecedence) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"queue_limit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Queue size in packets  |

`,
		},

		"average_packet": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Average packet size (bytes)

|  Format  |  Description  |
|----------|---------------|
|  u32:16-10240  |  Average packet size in bytes  |

`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"mark_probability": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mark probability for this precedence

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Numeric value (1/N)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"maximum_threshold": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum threshold for random detection

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4096  |  Maximum Threshold in packets  |

`,

			// Default:          stringdefault.StaticString(`18`),
			Computed: true,
		},

		"minimum_threshold": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Minimum  threshold for random detection

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4096  |  Maximum Threshold in packets  |

`,
		},

		// TagNodes

		// Nodes

	}
}
