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

// InterfacesBrIDgeVifDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesBrIDgeVifDhcpvsixOptionsPdInterface struct {
	// LeafNodes
	LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address"`
	LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceSLAID   types.String `tfsdk:"sla_id"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBrIDgeVifDhcpvsixOptionsPdInterface) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bridge", "vif", "dhcpv6-options", "pd", "interface"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceAddress.IsNull() || o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceAddress.IsUnknown()) {
		vyosData["address"] = o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceAddress.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceSLAID.IsNull() || o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceSLAID.IsUnknown()) {
		vyosData["sla-id"] = o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceSLAID.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBrIDgeVifDhcpvsixOptionsPdInterface) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bridge", "vif", "dhcpv6-options", "pd", "interface"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["sla-id"]; ok {
		o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceSLAID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifDhcpvsixOptionsPdInterfaceSLAID = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bridge", "vif", "dhcpv6-options", "pd", "interface"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBrIDgeVifDhcpvsixOptionsPdInterface) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"sla_id":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeVifDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

|  Format  |  Description  |
|----------|---------------|
|  >0  |  Used to form IPv6 interface address  |

`,
		},

		"sla_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface site-Level aggregator (SLA)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Decimal integer which fits in the length of SLA IDs  |

`,
		},

		// TagNodes

		// Nodes

	}
}