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

// PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation describes the resource data model.
type PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetLthreevpnNexthopEncapsulationGre types.String `tfsdk:"gre"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "set", "l3vpn-nexthop", "encapsulation"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteMapRuleSetLthreevpnNexthopEncapsulationGre.IsNull() || o.LeafPolicyRouteMapRuleSetLthreevpnNexthopEncapsulationGre.IsUnknown()) {
		vyosData["gre"] = o.LeafPolicyRouteMapRuleSetLthreevpnNexthopEncapsulationGre.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "set", "l3vpn-nexthop", "encapsulation"}})

	// Leafs
	if value, ok := vyosData["gre"]; ok {
		o.LeafPolicyRouteMapRuleSetLthreevpnNexthopEncapsulationGre = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleSetLthreevpnNexthopEncapsulationGre = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "set", "l3vpn-nexthop", "encapsulation"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"gre": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"gre": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Accept L3VPN traffic over GRE encapsulation

`,
		},

		// TagNodes

		// Nodes

	}
}
