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

// InterfacesMacsecMirror describes the resource data model.
type InterfacesMacsecMirror struct {
	// LeafNodes
	LeafInterfacesMacsecMirrorIngress types.String `tfsdk:"ingress"`
	LeafInterfacesMacsecMirrorEgress  types.String `tfsdk:"egress"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesMacsecMirror) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "macsec", "mirror"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesMacsecMirrorIngress.IsNull() || o.LeafInterfacesMacsecMirrorIngress.IsUnknown()) {
		vyosData["ingress"] = o.LeafInterfacesMacsecMirrorIngress.ValueString()
	}
	if !(o.LeafInterfacesMacsecMirrorEgress.IsNull() || o.LeafInterfacesMacsecMirrorEgress.IsUnknown()) {
		vyosData["egress"] = o.LeafInterfacesMacsecMirrorEgress.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesMacsecMirror) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "macsec", "mirror"}})

	// Leafs
	if value, ok := vyosData["ingress"]; ok {
		o.LeafInterfacesMacsecMirrorIngress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecMirrorIngress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["egress"]; ok {
		o.LeafInterfacesMacsecMirrorEgress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecMirrorEgress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "macsec", "mirror"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesMacsecMirror) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"ingress": types.StringType,
		"egress":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
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
