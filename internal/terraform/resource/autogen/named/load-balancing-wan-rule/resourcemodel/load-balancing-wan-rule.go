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

// LoadBalancingWanRule describes the resource data model.
type LoadBalancingWanRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafLoadBalancingWanRuleDescrIPtion        types.String `tfsdk:"description"`
	LeafLoadBalancingWanRuleExclude            types.String `tfsdk:"exclude"`
	LeafLoadBalancingWanRuleFailover           types.String `tfsdk:"failover"`
	LeafLoadBalancingWanRuleInboundInterface   types.String `tfsdk:"inbound_interface"`
	LeafLoadBalancingWanRulePerPacketBalancing types.String `tfsdk:"per_packet_balancing"`
	LeafLoadBalancingWanRuleProtocol           types.String `tfsdk:"protocol"`

	// TagNodes
	TagLoadBalancingWanRuleInterface types.Map `tfsdk:"interface"`

	// Nodes
	NodeLoadBalancingWanRuleDestination types.Object `tfsdk:"destination"`
	NodeLoadBalancingWanRuleLimit       types.Object `tfsdk:"limit"`
	NodeLoadBalancingWanRuleSource      types.Object `tfsdk:"source"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingWanRule) GetVyosPath() []string {
	return []string{
		"load-balancing",
		"wan",
		"rule",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *LoadBalancingWanRule) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"load-balancing", "wan", "rule"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafLoadBalancingWanRuleDescrIPtion.IsNull() || o.LeafLoadBalancingWanRuleDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafLoadBalancingWanRuleDescrIPtion.ValueString()
	}
	if !(o.LeafLoadBalancingWanRuleExclude.IsNull() || o.LeafLoadBalancingWanRuleExclude.IsUnknown()) {
		vyosData["exclude"] = o.LeafLoadBalancingWanRuleExclude.ValueString()
	}
	if !(o.LeafLoadBalancingWanRuleFailover.IsNull() || o.LeafLoadBalancingWanRuleFailover.IsUnknown()) {
		vyosData["failover"] = o.LeafLoadBalancingWanRuleFailover.ValueString()
	}
	if !(o.LeafLoadBalancingWanRuleInboundInterface.IsNull() || o.LeafLoadBalancingWanRuleInboundInterface.IsUnknown()) {
		vyosData["inbound-interface"] = o.LeafLoadBalancingWanRuleInboundInterface.ValueString()
	}
	if !(o.LeafLoadBalancingWanRulePerPacketBalancing.IsNull() || o.LeafLoadBalancingWanRulePerPacketBalancing.IsUnknown()) {
		vyosData["per-packet-balancing"] = o.LeafLoadBalancingWanRulePerPacketBalancing.ValueString()
	}
	if !(o.LeafLoadBalancingWanRuleProtocol.IsNull() || o.LeafLoadBalancingWanRuleProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafLoadBalancingWanRuleProtocol.ValueString()
	}

	// Tags
	if !(o.TagLoadBalancingWanRuleInterface.IsNull() || o.TagLoadBalancingWanRuleInterface.IsUnknown()) {
		subModel := make(map[string]LoadBalancingWanRuleInterface)
		diags.Append(o.TagLoadBalancingWanRuleInterface.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["interface"] = subData
	}

	// Nodes
	if !(o.NodeLoadBalancingWanRuleDestination.IsNull() || o.NodeLoadBalancingWanRuleDestination.IsUnknown()) {
		var subModel LoadBalancingWanRuleDestination
		diags.Append(o.NodeLoadBalancingWanRuleDestination.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["destination"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeLoadBalancingWanRuleLimit.IsNull() || o.NodeLoadBalancingWanRuleLimit.IsUnknown()) {
		var subModel LoadBalancingWanRuleLimit
		diags.Append(o.NodeLoadBalancingWanRuleLimit.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["limit"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeLoadBalancingWanRuleSource.IsNull() || o.NodeLoadBalancingWanRuleSource.IsUnknown()) {
		var subModel LoadBalancingWanRuleSource
		diags.Append(o.NodeLoadBalancingWanRuleSource.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["source"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *LoadBalancingWanRule) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"load-balancing", "wan", "rule"}})

	// Leafs
	if value, ok := vyosData["description"]; ok {
		o.LeafLoadBalancingWanRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRuleDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["exclude"]; ok {
		o.LeafLoadBalancingWanRuleExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRuleExclude = basetypes.NewStringNull()
	}
	if value, ok := vyosData["failover"]; ok {
		o.LeafLoadBalancingWanRuleFailover = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRuleFailover = basetypes.NewStringNull()
	}
	if value, ok := vyosData["inbound-interface"]; ok {
		o.LeafLoadBalancingWanRuleInboundInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRuleInboundInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["per-packet-balancing"]; ok {
		o.LeafLoadBalancingWanRulePerPacketBalancing = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRulePerPacketBalancing = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafLoadBalancingWanRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafLoadBalancingWanRuleProtocol = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["interface"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: LoadBalancingWanRuleInterface{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagLoadBalancingWanRuleInterface = data
	} else {
		o.TagLoadBalancingWanRuleInterface = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["destination"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, LoadBalancingWanRuleDestination{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeLoadBalancingWanRuleDestination = data

	} else {
		o.NodeLoadBalancingWanRuleDestination = basetypes.NewObjectNull(LoadBalancingWanRuleDestination{}.AttributeTypes())
	}
	if value, ok := vyosData["limit"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, LoadBalancingWanRuleLimit{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeLoadBalancingWanRuleLimit = data

	} else {
		o.NodeLoadBalancingWanRuleLimit = basetypes.NewObjectNull(LoadBalancingWanRuleLimit{}.AttributeTypes())
	}
	if value, ok := vyosData["source"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, LoadBalancingWanRuleSource{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeLoadBalancingWanRuleSource = data

	} else {
		o.NodeLoadBalancingWanRuleSource = basetypes.NewObjectNull(LoadBalancingWanRuleSource{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"load-balancing", "wan", "rule"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o LoadBalancingWanRule) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"description":          types.StringType,
		"exclude":              types.StringType,
		"failover":             types.StringType,
		"inbound_interface":    types.StringType,
		"per_packet_balancing": types.StringType,
		"protocol":             types.StringType,

		// Tags
		"interface": types.MapType{ElemType: types.ObjectType{AttrTypes: LoadBalancingWanRuleInterface{}.AttributeTypes()}},

		// Nodes
		"destination": types.ObjectType{AttrTypes: LoadBalancingWanRuleDestination{}.AttributeTypes()},
		"limit":       types.ObjectType{AttrTypes: LoadBalancingWanRuleLimit{}.AttributeTypes()},
		"source":      types.ObjectType{AttrTypes: LoadBalancingWanRuleSource{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingWanRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule number (1-9999)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-9999  |  Rule number  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description for this rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description for this rule  |

`,
		},

		"exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Exclude packets matching this rule from WAN load balance

`,
		},

		"failover": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable failover for packets matching this rule from WAN load balance

`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound interface name (e.g., "eth0") [REQUIRED]

`,
		},

		"per_packet_balancing": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Option to match traffic per-packet instead of the default, per-flow

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to match (protocol name, number, or "all")

|  Format  |  Description  |
|----------|---------------|
|  all  |  All IP protocols  |
|  tcp_udp  |  Both TCP and UDP  |
|  u32:0-255  |  IP protocol number  |
|  <protocol>  |  IP protocol name  |
|  !<protocol>  |  IP protocol name  |

`,
		},

		// TagNodes

		"interface": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: LoadBalancingWanRuleInterface{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Interface name [REQUIRED]

`,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: LoadBalancingWanRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: LoadBalancingWanRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable packet limit for this rule

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: LoadBalancingWanRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source information

`,
		},
	}
}
