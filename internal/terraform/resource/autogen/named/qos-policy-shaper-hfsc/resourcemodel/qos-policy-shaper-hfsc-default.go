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

// QosPolicyShaperHfscDefault describes the resource data model.
type QosPolicyShaperHfscDefault struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeQosPolicyShaperHfscDefaultLinkshare  types.Object `tfsdk:"linkshare"`
	NodeQosPolicyShaperHfscDefaultRealtime   types.Object `tfsdk:"realtime"`
	NodeQosPolicyShaperHfscDefaultUpperlimit types.Object `tfsdk:"upperlimit"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperHfscDefault) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "default"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeQosPolicyShaperHfscDefaultLinkshare.IsNull() || o.NodeQosPolicyShaperHfscDefaultLinkshare.IsUnknown()) {
		var subModel QosPolicyShaperHfscDefaultLinkshare
		diags.Append(o.NodeQosPolicyShaperHfscDefaultLinkshare.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["linkshare"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperHfscDefaultRealtime.IsNull() || o.NodeQosPolicyShaperHfscDefaultRealtime.IsUnknown()) {
		var subModel QosPolicyShaperHfscDefaultRealtime
		diags.Append(o.NodeQosPolicyShaperHfscDefaultRealtime.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["realtime"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperHfscDefaultUpperlimit.IsNull() || o.NodeQosPolicyShaperHfscDefaultUpperlimit.IsUnknown()) {
		var subModel QosPolicyShaperHfscDefaultUpperlimit
		diags.Append(o.NodeQosPolicyShaperHfscDefaultUpperlimit.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["upperlimit"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperHfscDefault) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "default"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["linkshare"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscDefaultLinkshare{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscDefaultLinkshare = data

	} else {
		o.NodeQosPolicyShaperHfscDefaultLinkshare = basetypes.NewObjectNull(QosPolicyShaperHfscDefaultLinkshare{}.AttributeTypes())
	}
	if value, ok := vyosData["realtime"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscDefaultRealtime{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscDefaultRealtime = data

	} else {
		o.NodeQosPolicyShaperHfscDefaultRealtime = basetypes.NewObjectNull(QosPolicyShaperHfscDefaultRealtime{}.AttributeTypes())
	}
	if value, ok := vyosData["upperlimit"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscDefaultUpperlimit{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscDefaultUpperlimit = data

	} else {
		o.NodeQosPolicyShaperHfscDefaultUpperlimit = basetypes.NewObjectNull(QosPolicyShaperHfscDefaultUpperlimit{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "default"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperHfscDefault) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"linkshare":  types.ObjectType{AttrTypes: QosPolicyShaperHfscDefaultLinkshare{}.AttributeTypes()},
		"realtime":   types.ObjectType{AttrTypes: QosPolicyShaperHfscDefaultRealtime{}.AttributeTypes()},
		"upperlimit": types.ObjectType{AttrTypes: QosPolicyShaperHfscDefaultUpperlimit{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"linkshare": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefaultLinkshare{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Linkshare class settings

`,
		},

		"realtime": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefaultRealtime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Realtime class settings

`,
		},

		"upperlimit": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefaultUpperlimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Upperlimit class settings

`,
		},
	}
}
