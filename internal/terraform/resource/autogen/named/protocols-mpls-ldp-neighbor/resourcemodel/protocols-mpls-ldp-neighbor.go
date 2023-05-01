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

// ProtocolsMplsLdpNeighbor describes the resource data model.
type ProtocolsMplsLdpNeighbor struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsMplsLdpNeighborPassword        types.String `tfsdk:"password"`
	LeafProtocolsMplsLdpNeighborTTLSecURIty     types.String `tfsdk:"ttl_security"`
	LeafProtocolsMplsLdpNeighborSessionHoldtime types.String `tfsdk:"session_holdtime"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpNeighbor) GetVyosPath() []string {
	return []string{
		"protocols",
		"mpls",
		"ldp",
		"neighbor",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsMplsLdpNeighbor) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "mpls", "ldp", "neighbor"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsMplsLdpNeighborPassword.IsNull() || o.LeafProtocolsMplsLdpNeighborPassword.IsUnknown()) {
		vyosData["password"] = o.LeafProtocolsMplsLdpNeighborPassword.ValueString()
	}
	if !(o.LeafProtocolsMplsLdpNeighborTTLSecURIty.IsNull() || o.LeafProtocolsMplsLdpNeighborTTLSecURIty.IsUnknown()) {
		vyosData["ttl-security"] = o.LeafProtocolsMplsLdpNeighborTTLSecURIty.ValueString()
	}
	if !(o.LeafProtocolsMplsLdpNeighborSessionHoldtime.IsNull() || o.LeafProtocolsMplsLdpNeighborSessionHoldtime.IsUnknown()) {
		vyosData["session-holdtime"] = o.LeafProtocolsMplsLdpNeighborSessionHoldtime.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsMplsLdpNeighbor) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "mpls", "ldp", "neighbor"}})

	// Leafs
	if value, ok := vyosData["password"]; ok {
		o.LeafProtocolsMplsLdpNeighborPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsMplsLdpNeighborPassword = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ttl-security"]; ok {
		o.LeafProtocolsMplsLdpNeighborTTLSecURIty = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsMplsLdpNeighborTTLSecURIty = basetypes.NewStringNull()
	}
	if value, ok := vyosData["session-holdtime"]; ok {
		o.LeafProtocolsMplsLdpNeighborSessionHoldtime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsMplsLdpNeighborSessionHoldtime = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "mpls", "ldp", "neighbor"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsMplsLdpNeighbor) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"password":         types.StringType,
		"ttl_security":     types.StringType,
		"session_holdtime": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpNeighbor) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `LDP neighbor parameters

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Neighbor IPv4 address  |

`,
		},

		// LeafNodes

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor password

`,
		},

		"ttl_security": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor TTL security

|  Format  |  Description  |
|----------|---------------|
|  u32:1-254  |  TTL  |
|  disable  |  Disable neighbor TTL security  |

`,
		},

		"session_holdtime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Session IPv4 hold time

|  Format  |  Description  |
|----------|---------------|
|  u32:15-65535  |  Time in seconds  |

`,
		},

		// TagNodes

		// Nodes

	}
}
