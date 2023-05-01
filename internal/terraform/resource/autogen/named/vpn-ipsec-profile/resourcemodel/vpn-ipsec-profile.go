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

// VpnIPsecProfile describes the resource data model.
type VpnIPsecProfile struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafVpnIPsecProfileDisable  types.String `tfsdk:"disable"`
	LeafVpnIPsecProfileEspGroup types.String `tfsdk:"esp_group"`
	LeafVpnIPsecProfileIkeGroup types.String `tfsdk:"ike_group"`

	// TagNodes

	// Nodes
	NodeVpnIPsecProfileAuthentication types.Object `tfsdk:"authentication"`
	NodeVpnIPsecProfileBind           types.Object `tfsdk:"bind"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecProfile) GetVyosPath() []string {
	return []string{
		"vpn",
		"ipsec",
		"profile",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnIPsecProfile) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "ipsec", "profile"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnIPsecProfileDisable.IsNull() || o.LeafVpnIPsecProfileDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafVpnIPsecProfileDisable.ValueString()
	}
	if !(o.LeafVpnIPsecProfileEspGroup.IsNull() || o.LeafVpnIPsecProfileEspGroup.IsUnknown()) {
		vyosData["esp-group"] = o.LeafVpnIPsecProfileEspGroup.ValueString()
	}
	if !(o.LeafVpnIPsecProfileIkeGroup.IsNull() || o.LeafVpnIPsecProfileIkeGroup.IsUnknown()) {
		vyosData["ike-group"] = o.LeafVpnIPsecProfileIkeGroup.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeVpnIPsecProfileAuthentication.IsNull() || o.NodeVpnIPsecProfileAuthentication.IsUnknown()) {
		var subModel VpnIPsecProfileAuthentication
		diags.Append(o.NodeVpnIPsecProfileAuthentication.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["authentication"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVpnIPsecProfileBind.IsNull() || o.NodeVpnIPsecProfileBind.IsUnknown()) {
		var subModel VpnIPsecProfileBind
		diags.Append(o.NodeVpnIPsecProfileBind.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["bind"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnIPsecProfile) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "ipsec", "profile"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafVpnIPsecProfileDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecProfileDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["esp-group"]; ok {
		o.LeafVpnIPsecProfileEspGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecProfileEspGroup = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ike-group"]; ok {
		o.LeafVpnIPsecProfileIkeGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecProfileIkeGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["authentication"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VpnIPsecProfileAuthentication{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVpnIPsecProfileAuthentication = data

	} else {
		o.NodeVpnIPsecProfileAuthentication = basetypes.NewObjectNull(VpnIPsecProfileAuthentication{}.AttributeTypes())
	}
	if value, ok := vyosData["bind"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VpnIPsecProfileBind{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVpnIPsecProfileBind = data

	} else {
		o.NodeVpnIPsecProfileBind = basetypes.NewObjectNull(VpnIPsecProfileBind{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "ipsec", "profile"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnIPsecProfile) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":   types.StringType,
		"esp_group": types.StringType,
		"ike_group": types.StringType,

		// Tags

		// Nodes
		"authentication": types.ObjectType{AttrTypes: VpnIPsecProfileAuthentication{}.AttributeTypes()},
		"bind":           types.ObjectType{AttrTypes: VpnIPsecProfileBind{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecProfile) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VPN IPsec profile

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Profile name  |

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"esp_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
		},

		"ike_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Internet Key Exchange (IKE) group name

`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: VpnIPsecProfileAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},

		"bind": schema.SingleNestedAttribute{
			Attributes: VpnIPsecProfileBind{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DMVPN tunnel configuration

`,
		},
	}
}
