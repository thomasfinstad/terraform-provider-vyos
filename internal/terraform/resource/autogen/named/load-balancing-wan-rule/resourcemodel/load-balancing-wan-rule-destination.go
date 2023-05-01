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

// LoadBalancingWanRuleDestination describes the resource data model.
type LoadBalancingWanRuleDestination struct {
	// LeafNodes
	LeafLoadBalancingWanRuleDestinationAddress types.String `tfsdk:"address"`
	LeafLoadBalancingWanRuleDestinationPort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *LoadBalancingWanRuleDestination) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"load-balancing", "wan", "rule", "destination"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafLoadBalancingWanRuleDestinationAddress.IsNull() || o.LeafLoadBalancingWanRuleDestinationAddress.IsUnknown()) {
		vyosData["address"] = o.LeafLoadBalancingWanRuleDestinationAddress.ValueString()
	}
	if !(o.LeafLoadBalancingWanRuleDestinationPort.IsNull() || o.LeafLoadBalancingWanRuleDestinationPort.IsUnknown()) {
		vyosData["port"] = o.LeafLoadBalancingWanRuleDestinationPort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *LoadBalancingWanRuleDestination) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"load-balancing", "wan", "rule", "destination"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafLoadBalancingWanRuleDestinationAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRuleDestinationAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafLoadBalancingWanRuleDestinationPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRuleDestinationPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"load-balancing", "wan", "rule", "destination"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o LoadBalancingWanRuleDestination) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingWanRuleDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
|  ipv4range  |  IPv4 address range to match  |
|  !ipv4  |  Match everything except the specified address  |
|  !ipv4net  |  Match everything except the specified prefix  |
|  !ipv4range  |  Match everything except the specified range  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numeric IP port  |
|  start-end  |  Numbered port range (e.g. 1001-1005)  |
|   |   |

`,
		},

		// TagNodes

		// Nodes

	}
}
