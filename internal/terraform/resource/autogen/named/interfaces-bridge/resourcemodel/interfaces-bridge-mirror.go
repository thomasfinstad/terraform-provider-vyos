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

// InterfacesBrIDgeMirror describes the resource data model.
type InterfacesBrIDgeMirror struct {
	// LeafNodes
	LeafInterfacesBrIDgeMirrorIngress types.String `tfsdk:"ingress"`
	LeafInterfacesBrIDgeMirrorEgress  types.String `tfsdk:"egress"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBrIDgeMirror) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bridge", "mirror"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBrIDgeMirrorIngress.IsNull() || o.LeafInterfacesBrIDgeMirrorIngress.IsUnknown()) {
		vyosData["ingress"] = o.LeafInterfacesBrIDgeMirrorIngress.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeMirrorEgress.IsNull() || o.LeafInterfacesBrIDgeMirrorEgress.IsUnknown()) {
		vyosData["egress"] = o.LeafInterfacesBrIDgeMirrorEgress.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBrIDgeMirror) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bridge", "mirror"}})

	// Leafs
	if value, ok := vyosData["ingress"]; ok {
		o.LeafInterfacesBrIDgeMirrorIngress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeMirrorIngress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["egress"]; ok {
		o.LeafInterfacesBrIDgeMirrorEgress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeMirrorEgress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bridge", "mirror"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBrIDgeMirror) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"ingress": types.StringType,
		"egress":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
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
