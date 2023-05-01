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

// VrfNameProtocolsIsisRedistributeIPvsixConnected describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixConnected struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne types.Object `tfsdk:"level_1"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo types.Object `tfsdk:"level_2"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsIsisRedistributeIPvsixConnected) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv6", "connected"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["level-1"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo.IsNull() || o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo.IsUnknown()) {
		var subModel VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo
		diags.Append(o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["level-2"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsIsisRedistributeIPvsixConnected) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv6", "connected"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["level-1"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne{}.AttributeTypes())
	}
	if value, ok := vyosData["level-2"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo = data

	} else {
		o.NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo = basetypes.NewObjectNull(VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "isis", "redistribute", "ipv6", "connected"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixConnected) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"level_1": types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne{}.AttributeTypes()},
		"level_2": types.ObjectType{AttrTypes: VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixConnected) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
		},
	}
}
