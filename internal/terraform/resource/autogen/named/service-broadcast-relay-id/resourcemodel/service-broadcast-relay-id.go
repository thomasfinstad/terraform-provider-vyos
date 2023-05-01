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

// ServiceBroadcastRelayID describes the resource data model.
type ServiceBroadcastRelayID struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServiceBroadcastRelayIDDisable     types.String `tfsdk:"disable"`
	LeafServiceBroadcastRelayIDAddress     types.String `tfsdk:"address"`
	LeafServiceBroadcastRelayIDDescrIPtion types.String `tfsdk:"description"`
	LeafServiceBroadcastRelayIDInterface   types.String `tfsdk:"interface"`
	LeafServiceBroadcastRelayIDPort        types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceBroadcastRelayID) GetVyosPath() []string {
	return []string{
		"service",
		"broadcast-relay",
		"id",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceBroadcastRelayID) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "broadcast-relay", "id"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceBroadcastRelayIDDisable.IsNull() || o.LeafServiceBroadcastRelayIDDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafServiceBroadcastRelayIDDisable.ValueString()
	}
	if !(o.LeafServiceBroadcastRelayIDAddress.IsNull() || o.LeafServiceBroadcastRelayIDAddress.IsUnknown()) {
		vyosData["address"] = o.LeafServiceBroadcastRelayIDAddress.ValueString()
	}
	if !(o.LeafServiceBroadcastRelayIDDescrIPtion.IsNull() || o.LeafServiceBroadcastRelayIDDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafServiceBroadcastRelayIDDescrIPtion.ValueString()
	}
	if !(o.LeafServiceBroadcastRelayIDInterface.IsNull() || o.LeafServiceBroadcastRelayIDInterface.IsUnknown()) {
		vyosData["interface"] = o.LeafServiceBroadcastRelayIDInterface.ValueString()
	}
	if !(o.LeafServiceBroadcastRelayIDPort.IsNull() || o.LeafServiceBroadcastRelayIDPort.IsUnknown()) {
		vyosData["port"] = o.LeafServiceBroadcastRelayIDPort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceBroadcastRelayID) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "broadcast-relay", "id"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafServiceBroadcastRelayIDDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceBroadcastRelayIDDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["address"]; ok {
		o.LeafServiceBroadcastRelayIDAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceBroadcastRelayIDAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafServiceBroadcastRelayIDDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceBroadcastRelayIDDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interface"]; ok {
		o.LeafServiceBroadcastRelayIDInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceBroadcastRelayIDInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafServiceBroadcastRelayIDPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceBroadcastRelayIDPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "broadcast-relay", "id"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceBroadcastRelayID) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":     types.StringType,
		"address":     types.StringType,
		"description": types.StringType,
		"interface":   types.StringType,
		"port":        types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceBroadcastRelayID) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Unique ID for each UDP port to forward

|  Format  |  Description  |
|----------|---------------|
|  u32:1-99  |  Broadcast relay instance ID  |

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set source IP of forwarded packets, otherwise original senders address is used

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Optional source address for forwarded packets  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

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
