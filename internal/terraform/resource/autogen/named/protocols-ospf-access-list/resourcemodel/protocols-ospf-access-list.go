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

// ProtocolsOspfAccessList describes the resource data model.
type ProtocolsOspfAccessList struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsOspfAccessListExport types.String `tfsdk:"export"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAccessList) GetVyosPath() []string {
	return []string{
		"protocols",
		"ospf",
		"access-list",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsOspfAccessList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "ospf", "access-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsOspfAccessListExport.IsNull() || o.LeafProtocolsOspfAccessListExport.IsUnknown()) {
		vyosData["export"] = o.LeafProtocolsOspfAccessListExport.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsOspfAccessList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "ospf", "access-list"}})

	// Leafs
	if value, ok := vyosData["export"]; ok {
		o.LeafProtocolsOspfAccessListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfAccessListExport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "ospf", "access-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsOspfAccessList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Access list to filter networks in routing updates

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access-list number  |

`,
		},

		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Filter for outgoing routing update

|  Format  |  Description  |
|----------|---------------|
|  bgp  |  Filter BGP routes  |
|  connected  |  Filter connected routes  |
|  isis  |  Filter IS-IS routes  |
|  kernel  |  Filter Kernel routes  |
|  rip  |  Filter RIP routes  |
|  static  |  Filter static routes  |

`,
		},

		// TagNodes

		// Nodes

	}
}