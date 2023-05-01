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

// QosPolicyRateControl describes the resource data model.
type QosPolicyRateControl struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafQosPolicyRateControlDescrIPtion types.String `tfsdk:"description"`
	LeafQosPolicyRateControlBandwIDth   types.String `tfsdk:"bandwidth"`
	LeafQosPolicyRateControlBurst       types.String `tfsdk:"burst"`
	LeafQosPolicyRateControlLatency     types.String `tfsdk:"latency"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyRateControl) GetVyosPath() []string {
	return []string{
		"qos",
		"policy",
		"rate-control",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyRateControl) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "rate-control"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyRateControlDescrIPtion.IsNull() || o.LeafQosPolicyRateControlDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafQosPolicyRateControlDescrIPtion.ValueString()
	}
	if !(o.LeafQosPolicyRateControlBandwIDth.IsNull() || o.LeafQosPolicyRateControlBandwIDth.IsUnknown()) {
		vyosData["bandwidth"] = o.LeafQosPolicyRateControlBandwIDth.ValueString()
	}
	if !(o.LeafQosPolicyRateControlBurst.IsNull() || o.LeafQosPolicyRateControlBurst.IsUnknown()) {
		vyosData["burst"] = o.LeafQosPolicyRateControlBurst.ValueString()
	}
	if !(o.LeafQosPolicyRateControlLatency.IsNull() || o.LeafQosPolicyRateControlLatency.IsUnknown()) {
		vyosData["latency"] = o.LeafQosPolicyRateControlLatency.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyRateControl) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "rate-control"}})

	// Leafs
	if value, ok := vyosData["description"]; ok {
		o.LeafQosPolicyRateControlDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRateControlDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["bandwidth"]; ok {
		o.LeafQosPolicyRateControlBandwIDth = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRateControlBandwIDth = basetypes.NewStringNull()
	}
	if value, ok := vyosData["burst"]; ok {
		o.LeafQosPolicyRateControlBurst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRateControlBurst = basetypes.NewStringNull()
	}
	if value, ok := vyosData["latency"]; ok {
		o.LeafQosPolicyRateControlLatency = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRateControlLatency = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "rate-control"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyRateControl) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"description": types.StringType,
		"bandwidth":   types.StringType,
		"burst":       types.StringType,
		"latency":     types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRateControl) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rate limiting policy (Token Bucket Filter)

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bits per second  |
|  <number>bit  |  Bits per second  |
|  <number>kbit  |  Kilobits per second  |
|  <number>mbit  |  Megabits per second  |
|  <number>gbit  |  Gigabits per second  |
|  <number>tbit  |  Terabits per second  |
|  <number>%  |  Percentage of interface link speed  |

`,
		},

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Burst size for this class

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bytes  |
|  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |

`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"latency": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum latency

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Time in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`50`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}