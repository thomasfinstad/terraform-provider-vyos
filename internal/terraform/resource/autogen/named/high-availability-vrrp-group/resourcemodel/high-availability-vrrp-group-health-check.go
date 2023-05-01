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

// HighAvailabilityVrrpGroupHealthCheck describes the resource data model.
type HighAvailabilityVrrpGroupHealthCheck struct {
	// LeafNodes
	LeafHighAvailabilityVrrpGroupHealthCheckFailureCount types.String `tfsdk:"failure_count"`
	LeafHighAvailabilityVrrpGroupHealthCheckInterval     types.String `tfsdk:"interval"`
	LeafHighAvailabilityVrrpGroupHealthCheckScrIPt       types.String `tfsdk:"script"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *HighAvailabilityVrrpGroupHealthCheck) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"high-availability", "vrrp", "group", "health-check"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafHighAvailabilityVrrpGroupHealthCheckFailureCount.IsNull() || o.LeafHighAvailabilityVrrpGroupHealthCheckFailureCount.IsUnknown()) {
		vyosData["failure-count"] = o.LeafHighAvailabilityVrrpGroupHealthCheckFailureCount.ValueString()
	}
	if !(o.LeafHighAvailabilityVrrpGroupHealthCheckInterval.IsNull() || o.LeafHighAvailabilityVrrpGroupHealthCheckInterval.IsUnknown()) {
		vyosData["interval"] = o.LeafHighAvailabilityVrrpGroupHealthCheckInterval.ValueString()
	}
	if !(o.LeafHighAvailabilityVrrpGroupHealthCheckScrIPt.IsNull() || o.LeafHighAvailabilityVrrpGroupHealthCheckScrIPt.IsUnknown()) {
		vyosData["script"] = o.LeafHighAvailabilityVrrpGroupHealthCheckScrIPt.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *HighAvailabilityVrrpGroupHealthCheck) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"high-availability", "vrrp", "group", "health-check"}})

	// Leafs
	if value, ok := vyosData["failure-count"]; ok {
		o.LeafHighAvailabilityVrrpGroupHealthCheckFailureCount = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVrrpGroupHealthCheckFailureCount = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interval"]; ok {
		o.LeafHighAvailabilityVrrpGroupHealthCheckInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVrrpGroupHealthCheckInterval = basetypes.NewStringNull()
	}
	if value, ok := vyosData["script"]; ok {
		o.LeafHighAvailabilityVrrpGroupHealthCheckScrIPt = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVrrpGroupHealthCheckScrIPt = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"high-availability", "vrrp", "group", "health-check"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o HighAvailabilityVrrpGroupHealthCheck) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"failure_count": types.StringType,
		"interval":      types.StringType,
		"script":        types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupHealthCheck) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"failure_count": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check failure count required for transition to fault

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check execution interval in seconds

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"script": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check script file

`,
		},

		// TagNodes

		// Nodes

	}
}
