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

// InterfacesWirelessVifSVifCMirror describes the resource data model.
type InterfacesWirelessVifSVifCMirror struct {
	// LeafNodes
	LeafInterfacesWirelessVifSVifCMirrorIngress types.String `tfsdk:"ingress"`
	LeafInterfacesWirelessVifSVifCMirrorEgress  types.String `tfsdk:"egress"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessVifSVifCMirror) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "vif-c", "mirror"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessVifSVifCMirrorIngress.IsNull() || o.LeafInterfacesWirelessVifSVifCMirrorIngress.IsUnknown()) {
		vyosData["ingress"] = o.LeafInterfacesWirelessVifSVifCMirrorIngress.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCMirrorEgress.IsNull() || o.LeafInterfacesWirelessVifSVifCMirrorEgress.IsUnknown()) {
		vyosData["egress"] = o.LeafInterfacesWirelessVifSVifCMirrorEgress.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessVifSVifCMirror) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "vif-c", "mirror"}})

	// Leafs
	if value, ok := vyosData["ingress"]; ok {
		o.LeafInterfacesWirelessVifSVifCMirrorIngress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCMirrorIngress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["egress"]; ok {
		o.LeafInterfacesWirelessVifSVifCMirrorEgress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCMirrorEgress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "vif-c", "mirror"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessVifSVifCMirror) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"ingress": types.StringType,
		"egress":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSVifCMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ingress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror ingress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		"egress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror egress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		// TagNodes

		// Nodes

	}
}
