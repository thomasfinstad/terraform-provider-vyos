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

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeTable types.String `tfsdk:"table"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected  types.Object `tfsdk:"connected"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel     types.Object `tfsdk:"kernel"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree types.Object `tfsdk:"ospfv3"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng      types.Object `tfsdk:"ripng"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel      types.Object `tfsdk:"babel"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic     types.Object `tfsdk:"static"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "redistribute"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeTable.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeTable.IsUnknown()) {
		vyosData["table"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeTable.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["connected"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["kernel"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ospfv3"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ripng"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["babel"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["static"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "redistribute"}})

	// Leafs
	if value, ok := vyosData["table"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeTable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeTable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["connected"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}.AttributeTypes())
	}
	if value, ok := vyosData["kernel"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}.AttributeTypes())
	}
	if value, ok := vyosData["ospfv3"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree{}.AttributeTypes())
	}
	if value, ok := vyosData["ripng"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng{}.AttributeTypes())
	}
	if value, ok := vyosData["babel"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}.AttributeTypes())
	}
	if value, ok := vyosData["static"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "redistribute"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"table": types.StringType,

		// Tags

		// Nodes
		"connected": types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}.AttributeTypes()},
		"kernel":    types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}.AttributeTypes()},
		"ospfv3":    types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree{}.AttributeTypes()},
		"ripng":     types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng{}.AttributeTypes()},
		"babel":     types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}.AttributeTypes()},
		"static":    types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"table": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute non-main Kernel Routing Table

`,
		},

		// TagNodes

		// Nodes

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes into BGP

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes into BGP

`,
		},

		"ospfv3": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute OSPFv3 routes into BGP

`,
		},

		"ripng": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute RIPng routes into BGP

`,
		},

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes into BGP

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes into BGP

`,
		},
	}
}