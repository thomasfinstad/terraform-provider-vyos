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

// InterfacesTunnelParameters describes the resource data model.
type InterfacesTunnelParameters struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeInterfacesTunnelParametersErspan types.Object `tfsdk:"erspan"`
	NodeInterfacesTunnelParametersIP     types.Object `tfsdk:"ip"`
	NodeInterfacesTunnelParametersIPvsix types.Object `tfsdk:"ipv6"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesTunnelParameters) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "tunnel", "parameters"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeInterfacesTunnelParametersErspan.IsNull() || o.NodeInterfacesTunnelParametersErspan.IsUnknown()) {
		var subModel InterfacesTunnelParametersErspan
		diags.Append(o.NodeInterfacesTunnelParametersErspan.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["erspan"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesTunnelParametersIP.IsNull() || o.NodeInterfacesTunnelParametersIP.IsUnknown()) {
		var subModel InterfacesTunnelParametersIP
		diags.Append(o.NodeInterfacesTunnelParametersIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesTunnelParametersIPvsix.IsNull() || o.NodeInterfacesTunnelParametersIPvsix.IsUnknown()) {
		var subModel InterfacesTunnelParametersIPvsix
		diags.Append(o.NodeInterfacesTunnelParametersIPvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesTunnelParameters) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "tunnel", "parameters"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["erspan"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesTunnelParametersErspan{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesTunnelParametersErspan = data

	} else {
		o.NodeInterfacesTunnelParametersErspan = basetypes.NewObjectNull(InterfacesTunnelParametersErspan{}.AttributeTypes())
	}
	if value, ok := vyosData["ip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesTunnelParametersIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesTunnelParametersIP = data

	} else {
		o.NodeInterfacesTunnelParametersIP = basetypes.NewObjectNull(InterfacesTunnelParametersIP{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesTunnelParametersIPvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesTunnelParametersIPvsix = data

	} else {
		o.NodeInterfacesTunnelParametersIPvsix = basetypes.NewObjectNull(InterfacesTunnelParametersIPvsix{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "tunnel", "parameters"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesTunnelParameters) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"erspan": types.ObjectType{AttrTypes: InterfacesTunnelParametersErspan{}.AttributeTypes()},
		"ip":     types.ObjectType{AttrTypes: InterfacesTunnelParametersIP{}.AttributeTypes()},
		"ipv6":   types.ObjectType{AttrTypes: InterfacesTunnelParametersIPvsix{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesTunnelParameters) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"erspan": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParametersErspan{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ERSPAN tunnel parameters

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParametersIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-specific tunnel parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParametersIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6-specific tunnel parameters

`,
		},
	}
}
