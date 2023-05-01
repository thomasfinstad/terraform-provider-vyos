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

// VrfNameProtocolsOspfTimersThroTTLeSpf describes the resource data model.
type VrfNameProtocolsOspfTimersThroTTLeSpf struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay           types.String `tfsdk:"delay"`
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime types.String `tfsdk:"initial_holdtime"`
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime     types.String `tfsdk:"max_holdtime"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfTimersThroTTLeSpf) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "timers", "throttle", "spf"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay.IsNull() || o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay.IsUnknown()) {
		vyosData["delay"] = o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime.IsNull() || o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime.IsUnknown()) {
		vyosData["initial-holdtime"] = o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime.IsNull() || o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime.IsUnknown()) {
		vyosData["max-holdtime"] = o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfTimersThroTTLeSpf) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "timers", "throttle", "spf"}})

	// Leafs
	if value, ok := vyosData["delay"]; ok {
		o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay = basetypes.NewStringNull()
	}
	if value, ok := vyosData["initial-holdtime"]; ok {
		o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime = basetypes.NewStringNull()
	}
	if value, ok := vyosData["max-holdtime"]; ok {
		o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "timers", "throttle", "spf"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfTimersThroTTLeSpf) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"delay":            types.StringType,
		"initial_holdtime": types.StringType,
		"max_holdtime":     types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfTimersThroTTLeSpf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"delay": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Delay from the first change received to SPF calculation

|  Format  |  Description  |
|----------|---------------|
|  u32:0-600000  |  Delay in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`200`),
			Computed: true,
		},

		"initial_holdtime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Initial hold time between consecutive SPF calculations

|  Format  |  Description  |
|----------|---------------|
|  u32:0-600000  |  Initial hold time in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`1000`),
			Computed: true,
		},

		"max_holdtime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum hold time

|  Format  |  Description  |
|----------|---------------|
|  u32:0-600000  |  Max hold time in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`10000`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}