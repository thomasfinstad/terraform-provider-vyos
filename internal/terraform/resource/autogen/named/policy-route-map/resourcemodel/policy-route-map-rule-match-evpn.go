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

// PolicyRouteMapRuleMatchEvpn describes the resource data model.
type PolicyRouteMapRuleMatchEvpn struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchEvpnDefaultRoute types.String `tfsdk:"default_route"`
	LeafPolicyRouteMapRuleMatchEvpnRd           types.String `tfsdk:"rd"`
	LeafPolicyRouteMapRuleMatchEvpnRouteType    types.String `tfsdk:"route_type"`
	LeafPolicyRouteMapRuleMatchEvpnVni          types.String `tfsdk:"vni"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteMapRuleMatchEvpn) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "evpn"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute.IsNull() || o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute.IsUnknown()) {
		vyosData["default-route"] = o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchEvpnRd.IsNull() || o.LeafPolicyRouteMapRuleMatchEvpnRd.IsUnknown()) {
		vyosData["rd"] = o.LeafPolicyRouteMapRuleMatchEvpnRd.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchEvpnRouteType.IsNull() || o.LeafPolicyRouteMapRuleMatchEvpnRouteType.IsUnknown()) {
		vyosData["route-type"] = o.LeafPolicyRouteMapRuleMatchEvpnRouteType.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchEvpnVni.IsNull() || o.LeafPolicyRouteMapRuleMatchEvpnVni.IsUnknown()) {
		vyosData["vni"] = o.LeafPolicyRouteMapRuleMatchEvpnVni.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteMapRuleMatchEvpn) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "evpn"}})

	// Leafs
	if value, ok := vyosData["default-route"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnDefaultRoute = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rd"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnRd = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnRd = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-type"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnRouteType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnRouteType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vni"]; ok {
		o.LeafPolicyRouteMapRuleMatchEvpnVni = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchEvpnVni = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "evpn"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteMapRuleMatchEvpn) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"default_route": types.StringType,
		"rd":            types.StringType,
		"route_type":    types.StringType,
		"vni":           types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchEvpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_route": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default EVPN type-5 route

`,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		"route_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match route-type

|  Format  |  Description  |
|----------|---------------|
|  macip  |  mac-ip route  |
|  multicast  |  IMET route  |
|  prefix  |  Prefix route  |

`,
		},

		"vni": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Network Identifier

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777214  |  VXLAN virtual network identifier  |

`,
		},

		// TagNodes

		// Nodes

	}
}
