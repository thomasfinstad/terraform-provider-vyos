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

// VrfNameProtocolsIsisRedistributeIPvfour describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvfour struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvfourBgp       types.Object `tfsdk:"bgp"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourConnected types.Object `tfsdk:"connected"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourKernel    types.Object `tfsdk:"kernel"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourOspf      types.Object `tfsdk:"ospf"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourRIP       types.Object `tfsdk:"rip"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourBabel     types.Object `tfsdk:"babel"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourStatic    types.Object `tfsdk:"static"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsIsisRedistributeIPvfour) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv4"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvfourBgp.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvfourBgp.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvfourBgp
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvfourBgp.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["bgp"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvfourConnected.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvfourConnected.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvfourConnected
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvfourConnected.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["connected"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvfourKernel.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvfourKernel.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvfourKernel
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvfourKernel.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["kernel"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvfourOspf.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvfourOspf.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvfourOspf
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvfourOspf.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ospf"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvfourRIP.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvfourRIP.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvfourRIP
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvfourRIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["rip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvfourBabel.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvfourBabel.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvfourBabel
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvfourBabel.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["babel"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvfourStatic.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvfourStatic.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvfourStatic
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvfourStatic.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["static"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsIsisRedistributeIPvfour) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv4"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["bgp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvfourBgp{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourBgp = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourBgp = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvfourBgp{}.AttributeTypes())
	}
	if value, ok := vyosData["connected"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvfourConnected{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourConnected = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourConnected = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvfourConnected{}.AttributeTypes())
	}
	if value, ok := vyosData["kernel"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvfourKernel{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourKernel = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourKernel = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvfourKernel{}.AttributeTypes())
	}
	if value, ok := vyosData["ospf"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvfourOspf{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourOspf = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourOspf = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvfourOspf{}.AttributeTypes())
	}
	if value, ok := vyosData["rip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvfourRIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourRIP = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourRIP = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvfourRIP{}.AttributeTypes())
	}
	if value, ok := vyosData["babel"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvfourBabel{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourBabel = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourBabel = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvfourBabel{}.AttributeTypes())
	}
	if value, ok := vyosData["static"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvfourStatic{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourStatic = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvfourStatic = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvfourStatic{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv4"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfour) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"bgp":       types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvfourBgp{}.AttributeTypes()},
		"connected": types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvfourConnected{}.AttributeTypes()},
		"kernel":    types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvfourKernel{}.AttributeTypes()},
		"ospf":      types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvfourOspf{}.AttributeTypes()},
		"rip":       types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvfourRIP{}.AttributeTypes()},
		"babel":     types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvfourBabel{}.AttributeTypes()},
		"static":    types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvfourStatic{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfour) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourBgp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Border Gateway Protocol (BGP)

`,
		},

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourConnected{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes into IS-IS

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourKernel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes into IS-IS

`,
		},

		"ospf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourOspf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute OSPF routes into IS-IS

`,
		},

		"rip": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourRIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute RIP routes into IS-IS

`,
		},

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourBabel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes into IS-IS

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourStatic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes into IS-IS

`,
		},
	}
}
