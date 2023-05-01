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

// VrfNameProtocolsOspfSegmentRoutingLocalBlock describes the resource data model.
type VrfNameProtocolsOspfSegmentRoutingLocalBlock struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockLowLabelValue  types.String `tfsdk:"low_label_value"`
	LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockHighLabelValue types.String `tfsdk:"high_label_value"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfSegmentRoutingLocalBlock) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "segment-routing", "local-block"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockLowLabelValue.IsNull() || o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockLowLabelValue.IsUnknown()) {
		vyosData["low-label-value"] = o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockLowLabelValue.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockHighLabelValue.IsNull() || o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockHighLabelValue.IsUnknown()) {
		vyosData["high-label-value"] = o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockHighLabelValue.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfSegmentRoutingLocalBlock) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "segment-routing", "local-block"}})

	// Leafs
	if value, ok := vyosData["low-label-value"]; ok {
		o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockLowLabelValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockLowLabelValue = basetypes.NewStringNull()
	}
	if value, ok := vyosData["high-label-value"]; ok {
		o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockHighLabelValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockHighLabelValue = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "segment-routing", "local-block"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfSegmentRoutingLocalBlock) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"low_label_value":  types.StringType,
		"high_label_value": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfSegmentRoutingLocalBlock) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"low_label_value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label lower bound

|  Format  |  Description  |
|----------|---------------|
|  u32:16-1048575  |  Label value (recommended minimum value: 300)  |

`,
		},

		"high_label_value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label upper bound

|  Format  |  Description  |
|----------|---------------|
|  u32:16-1048575  |  Label value  |

`,
		},

		// TagNodes

		// Nodes

	}
}
