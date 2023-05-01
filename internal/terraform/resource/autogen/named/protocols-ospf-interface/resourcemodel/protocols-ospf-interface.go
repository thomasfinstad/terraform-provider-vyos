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

// ProtocolsOspfInterface describes the resource data model.
type ProtocolsOspfInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsOspfInterfaceArea               types.String `tfsdk:"area"`
	LeafProtocolsOspfInterfaceDeadInterval       types.String `tfsdk:"dead_interval"`
	LeafProtocolsOspfInterfaceHelloInterval      types.String `tfsdk:"hello_interval"`
	LeafProtocolsOspfInterfaceRetransmitInterval types.String `tfsdk:"retransmit_interval"`
	LeafProtocolsOspfInterfaceTransmitDelay      types.String `tfsdk:"transmit_delay"`
	LeafProtocolsOspfInterfaceCost               types.String `tfsdk:"cost"`
	LeafProtocolsOspfInterfaceMtuIgnore          types.String `tfsdk:"mtu_ignore"`
	LeafProtocolsOspfInterfacePriority           types.String `tfsdk:"priority"`
	LeafProtocolsOspfInterfaceBandwIDth          types.String `tfsdk:"bandwidth"`
	LeafProtocolsOspfInterfaceHelloMultIPlier    types.String `tfsdk:"hello_multiplier"`
	LeafProtocolsOspfInterfaceNetwork            types.String `tfsdk:"network"`

	// TagNodes

	// Nodes
	NodeProtocolsOspfInterfaceAuthentication types.Object `tfsdk:"authentication"`
	NodeProtocolsOspfInterfaceBfd            types.Object `tfsdk:"bfd"`
	NodeProtocolsOspfInterfacePassive        types.Object `tfsdk:"passive"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfInterface) GetVyosPath() []string {
	return []string{
		"protocols",
		"ospf",
		"interface",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsOspfInterface) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "ospf", "interface"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsOspfInterfaceArea.IsNull() || o.LeafProtocolsOspfInterfaceArea.IsUnknown()) {
		vyosData["area"] = o.LeafProtocolsOspfInterfaceArea.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceDeadInterval.IsNull() || o.LeafProtocolsOspfInterfaceDeadInterval.IsUnknown()) {
		vyosData["dead-interval"] = o.LeafProtocolsOspfInterfaceDeadInterval.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceHelloInterval.IsNull() || o.LeafProtocolsOspfInterfaceHelloInterval.IsUnknown()) {
		vyosData["hello-interval"] = o.LeafProtocolsOspfInterfaceHelloInterval.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceRetransmitInterval.IsNull() || o.LeafProtocolsOspfInterfaceRetransmitInterval.IsUnknown()) {
		vyosData["retransmit-interval"] = o.LeafProtocolsOspfInterfaceRetransmitInterval.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceTransmitDelay.IsNull() || o.LeafProtocolsOspfInterfaceTransmitDelay.IsUnknown()) {
		vyosData["transmit-delay"] = o.LeafProtocolsOspfInterfaceTransmitDelay.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceCost.IsNull() || o.LeafProtocolsOspfInterfaceCost.IsUnknown()) {
		vyosData["cost"] = o.LeafProtocolsOspfInterfaceCost.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceMtuIgnore.IsNull() || o.LeafProtocolsOspfInterfaceMtuIgnore.IsUnknown()) {
		vyosData["mtu-ignore"] = o.LeafProtocolsOspfInterfaceMtuIgnore.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfacePriority.IsNull() || o.LeafProtocolsOspfInterfacePriority.IsUnknown()) {
		vyosData["priority"] = o.LeafProtocolsOspfInterfacePriority.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceBandwIDth.IsNull() || o.LeafProtocolsOspfInterfaceBandwIDth.IsUnknown()) {
		vyosData["bandwidth"] = o.LeafProtocolsOspfInterfaceBandwIDth.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceHelloMultIPlier.IsNull() || o.LeafProtocolsOspfInterfaceHelloMultIPlier.IsUnknown()) {
		vyosData["hello-multiplier"] = o.LeafProtocolsOspfInterfaceHelloMultIPlier.ValueString()
	}
	if !(o.LeafProtocolsOspfInterfaceNetwork.IsNull() || o.LeafProtocolsOspfInterfaceNetwork.IsUnknown()) {
		vyosData["network"] = o.LeafProtocolsOspfInterfaceNetwork.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeProtocolsOspfInterfaceAuthentication.IsNull() || o.NodeProtocolsOspfInterfaceAuthentication.IsUnknown()) {
		var subModel ProtocolsOspfInterfaceAuthentication
		diags.Append(o.NodeProtocolsOspfInterfaceAuthentication.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["authentication"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsOspfInterfaceBfd.IsNull() || o.NodeProtocolsOspfInterfaceBfd.IsUnknown()) {
		var subModel ProtocolsOspfInterfaceBfd
		diags.Append(o.NodeProtocolsOspfInterfaceBfd.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["bfd"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsOspfInterfacePassive.IsNull() || o.NodeProtocolsOspfInterfacePassive.IsUnknown()) {
		var subModel ProtocolsOspfInterfacePassive
		diags.Append(o.NodeProtocolsOspfInterfacePassive.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["passive"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsOspfInterface) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "ospf", "interface"}})

	// Leafs
	if value, ok := vyosData["area"]; ok {
		o.LeafProtocolsOspfInterfaceArea = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceArea = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dead-interval"]; ok {
		o.LeafProtocolsOspfInterfaceDeadInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceDeadInterval = basetypes.NewStringNull()
	}
	if value, ok := vyosData["hello-interval"]; ok {
		o.LeafProtocolsOspfInterfaceHelloInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceHelloInterval = basetypes.NewStringNull()
	}
	if value, ok := vyosData["retransmit-interval"]; ok {
		o.LeafProtocolsOspfInterfaceRetransmitInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceRetransmitInterval = basetypes.NewStringNull()
	}
	if value, ok := vyosData["transmit-delay"]; ok {
		o.LeafProtocolsOspfInterfaceTransmitDelay = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceTransmitDelay = basetypes.NewStringNull()
	}
	if value, ok := vyosData["cost"]; ok {
		o.LeafProtocolsOspfInterfaceCost = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceCost = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mtu-ignore"]; ok {
		o.LeafProtocolsOspfInterfaceMtuIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceMtuIgnore = basetypes.NewStringNull()
	}
	if value, ok := vyosData["priority"]; ok {
		o.LeafProtocolsOspfInterfacePriority = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfacePriority = basetypes.NewStringNull()
	}
	if value, ok := vyosData["bandwidth"]; ok {
		o.LeafProtocolsOspfInterfaceBandwIDth = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceBandwIDth = basetypes.NewStringNull()
	}
	if value, ok := vyosData["hello-multiplier"]; ok {
		o.LeafProtocolsOspfInterfaceHelloMultIPlier = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceHelloMultIPlier = basetypes.NewStringNull()
	}
	if value, ok := vyosData["network"]; ok {
		o.LeafProtocolsOspfInterfaceNetwork = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceNetwork = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["authentication"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsOspfInterfaceAuthentication{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsOspfInterfaceAuthentication = data

	} else {
		o.NodeProtocolsOspfInterfaceAuthentication = basetypes.NewObjectNull(ProtocolsOspfInterfaceAuthentication{}.AttributeTypes())
	}
	if value, ok := vyosData["bfd"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsOspfInterfaceBfd{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsOspfInterfaceBfd = data

	} else {
		o.NodeProtocolsOspfInterfaceBfd = basetypes.NewObjectNull(ProtocolsOspfInterfaceBfd{}.AttributeTypes())
	}
	if value, ok := vyosData["passive"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsOspfInterfacePassive{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsOspfInterfacePassive = data

	} else {
		o.NodeProtocolsOspfInterfacePassive = basetypes.NewObjectNull(ProtocolsOspfInterfacePassive{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "ospf", "interface"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsOspfInterface) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"area":                types.StringType,
		"dead_interval":       types.StringType,
		"hello_interval":      types.StringType,
		"retransmit_interval": types.StringType,
		"transmit_delay":      types.StringType,
		"cost":                types.StringType,
		"mtu_ignore":          types.StringType,
		"priority":            types.StringType,
		"bandwidth":           types.StringType,
		"hello_multiplier":    types.StringType,
		"network":             types.StringType,

		// Tags

		// Nodes
		"authentication": types.ObjectType{AttrTypes: ProtocolsOspfInterfaceAuthentication{}.AttributeTypes()},
		"bfd":            types.ObjectType{AttrTypes: ProtocolsOspfInterfaceBfd{}.AttributeTypes()},
		"passive":        types.ObjectType{AttrTypes: ProtocolsOspfInterfacePassive{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		// LeafNodes

		"area": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable OSPF on this interface

|  Format  |  Description  |
|----------|---------------|
|  u32  |  OSPF area ID as decimal notation  |
|  ipv4  |  OSPF area ID in IP address notation  |

`,
		},

		"dead_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval after which a neighbor is declared dead

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Neighbor dead interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"hello_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval between hello packets

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Hello interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"retransmit_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval between retransmitting lost link state advertisements

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Retransmit interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"transmit_delay": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Link state transmit delay

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Link state transmit delay (seconds)  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"cost": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface cost

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  OSPF interface cost  |

`,
		},

		"mtu_ignore": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable Maximum Transmission Unit (MTU) mismatch detection

`,
		},

		"priority": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Router priority

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  OSPF router priority cost  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface bandwidth (Mbit/s)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-100000  |  Bandwidth in Megabit/sec (for calculating OSPF cost)  |

`,
		},

		"hello_multiplier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hello multiplier factor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of Hellos to send each second  |

`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network type

|  Format  |  Description  |
|----------|---------------|
|  broadcast  |  Broadcast network type  |
|  non-broadcast  |  Non-broadcast network type  |
|  point-to-multipoint  |  Point-to-multipoint network type  |
|  point-to-point  |  Point-to-point network type  |

`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfInterfaceAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfInterfaceBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},

		"passive": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfInterfacePassive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Suppress routing updates on an interface

`,
		},
	}
}
