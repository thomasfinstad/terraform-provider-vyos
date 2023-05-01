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

// InterfacesTunnelParametersIPvsix describes the resource data model.
type InterfacesTunnelParametersIPvsix struct {
	// LeafNodes
	LeafInterfacesTunnelParametersIPvsixEncaplimit types.String `tfsdk:"encaplimit"`
	LeafInterfacesTunnelParametersIPvsixFlowlabel  types.String `tfsdk:"flowlabel"`
	LeafInterfacesTunnelParametersIPvsixHoplimit   types.String `tfsdk:"hoplimit"`
	LeafInterfacesTunnelParametersIPvsixTclass     types.String `tfsdk:"tclass"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesTunnelParametersIPvsix) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "tunnel", "parameters", "ipv6"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesTunnelParametersIPvsixEncaplimit.IsNull() || o.LeafInterfacesTunnelParametersIPvsixEncaplimit.IsUnknown()) {
		vyosData["encaplimit"] = o.LeafInterfacesTunnelParametersIPvsixEncaplimit.ValueString()
	}
	if !(o.LeafInterfacesTunnelParametersIPvsixFlowlabel.IsNull() || o.LeafInterfacesTunnelParametersIPvsixFlowlabel.IsUnknown()) {
		vyosData["flowlabel"] = o.LeafInterfacesTunnelParametersIPvsixFlowlabel.ValueString()
	}
	if !(o.LeafInterfacesTunnelParametersIPvsixHoplimit.IsNull() || o.LeafInterfacesTunnelParametersIPvsixHoplimit.IsUnknown()) {
		vyosData["hoplimit"] = o.LeafInterfacesTunnelParametersIPvsixHoplimit.ValueString()
	}
	if !(o.LeafInterfacesTunnelParametersIPvsixTclass.IsNull() || o.LeafInterfacesTunnelParametersIPvsixTclass.IsUnknown()) {
		vyosData["tclass"] = o.LeafInterfacesTunnelParametersIPvsixTclass.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesTunnelParametersIPvsix) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "tunnel", "parameters", "ipv6"}})

	// Leafs
	if value, ok := vyosData["encaplimit"]; ok {
		o.LeafInterfacesTunnelParametersIPvsixEncaplimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesTunnelParametersIPvsixEncaplimit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["flowlabel"]; ok {
		o.LeafInterfacesTunnelParametersIPvsixFlowlabel = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesTunnelParametersIPvsixFlowlabel = basetypes.NewStringNull()
	}
	if value, ok := vyosData["hoplimit"]; ok {
		o.LeafInterfacesTunnelParametersIPvsixHoplimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesTunnelParametersIPvsixHoplimit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["tclass"]; ok {
		o.LeafInterfacesTunnelParametersIPvsixTclass = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesTunnelParametersIPvsixTclass = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "tunnel", "parameters", "ipv6"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesTunnelParametersIPvsix) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"encaplimit": types.StringType,
		"flowlabel":  types.StringType,
		"hoplimit":   types.StringType,
		"tclass":     types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesTunnelParametersIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encaplimit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set fixed encapsulation limit

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Encapsulation limit  |
|  none  |  Disable encapsulation limit  |

`,

			// Default:          stringdefault.StaticString(`4`),
			Computed: true,
		},

		"flowlabel": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the flow label to use in outgoing packets

|  Format  |  Description  |
|----------|---------------|
|  inherit  |  Copy field from original header  |
|  0x0-0x0fffff  |  Tunnel key, or hex value  |

`,
		},

		"hoplimit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hoplimit

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Hop limit  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		"tclass": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Traffic class (Tclass)

|  Format  |  Description  |
|----------|---------------|
|  0x0-0x0fffff  |  Traffic class, 'inherit' or hex value  |

`,

			// Default:          stringdefault.StaticString(`inherit`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}