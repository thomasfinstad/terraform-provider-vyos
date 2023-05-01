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

// QosPolicyRoundRobinClassMatchEther describes the resource data model.
type QosPolicyRoundRobinClassMatchEther struct {
	// LeafNodes
	LeafQosPolicyRoundRobinClassMatchEtherDestination types.String `tfsdk:"destination"`
	LeafQosPolicyRoundRobinClassMatchEtherProtocol    types.String `tfsdk:"protocol"`
	LeafQosPolicyRoundRobinClassMatchEtherSource      types.String `tfsdk:"source"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyRoundRobinClassMatchEther) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "round-robin", "class", "match", "ether"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyRoundRobinClassMatchEtherDestination.IsNull() || o.LeafQosPolicyRoundRobinClassMatchEtherDestination.IsUnknown()) {
		vyosData["destination"] = o.LeafQosPolicyRoundRobinClassMatchEtherDestination.ValueString()
	}
	if !(o.LeafQosPolicyRoundRobinClassMatchEtherProtocol.IsNull() || o.LeafQosPolicyRoundRobinClassMatchEtherProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafQosPolicyRoundRobinClassMatchEtherProtocol.ValueString()
	}
	if !(o.LeafQosPolicyRoundRobinClassMatchEtherSource.IsNull() || o.LeafQosPolicyRoundRobinClassMatchEtherSource.IsUnknown()) {
		vyosData["source"] = o.LeafQosPolicyRoundRobinClassMatchEtherSource.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyRoundRobinClassMatchEther) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "round-robin", "class", "match", "ether"}})

	// Leafs
	if value, ok := vyosData["destination"]; ok {
		o.LeafQosPolicyRoundRobinClassMatchEtherDestination = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRoundRobinClassMatchEtherDestination = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafQosPolicyRoundRobinClassMatchEtherProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRoundRobinClassMatchEtherProtocol = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source"]; ok {
		o.LeafQosPolicyRoundRobinClassMatchEtherSource = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRoundRobinClassMatchEtherSource = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "round-robin", "class", "match", "ether"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyRoundRobinClassMatchEther) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"destination": types.StringType,
		"protocol":    types.StringType,
		"source":      types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinClassMatchEther) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet destination address for this match

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet protocol for this match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Ethernet protocol number  |
|  txt  |  Ethernet protocol name  |
|  all  |  Any protocol  |
|  ip  |  Internet IP (IPv4)  |
|  ipv6  |  Internet IP (IPv6)  |
|  arp  |  Address Resolution Protocol  |
|  atalk  |  Appletalk  |
|  ipx  |  Novell Internet Packet Exchange  |
|  802.1Q  |  802.1Q VLAN tag  |

`,
		},

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet source address for this match

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |

`,
		},

		// TagNodes

		// Nodes

	}
}