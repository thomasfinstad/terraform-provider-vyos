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

// QosPolicyShaperHfscClassMatchIP describes the resource data model.
type QosPolicyShaperHfscClassMatchIP struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchIPDscp      types.String `tfsdk:"dscp"`
	LeafQosPolicyShaperHfscClassMatchIPMaxLength types.String `tfsdk:"max_length"`
	LeafQosPolicyShaperHfscClassMatchIPProtocol  types.String `tfsdk:"protocol"`

	// TagNodes

	// Nodes
	NodeQosPolicyShaperHfscClassMatchIPDestination types.Object `tfsdk:"destination"`
	NodeQosPolicyShaperHfscClassMatchIPSource      types.Object `tfsdk:"source"`
	NodeQosPolicyShaperHfscClassMatchIPTCP         types.Object `tfsdk:"tcp"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperHfscClassMatchIP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ip"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperHfscClassMatchIPDscp.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPDscp.IsUnknown()) {
		vyosData["dscp"] = o.LeafQosPolicyShaperHfscClassMatchIPDscp.ValueString()
	}
	if !(o.LeafQosPolicyShaperHfscClassMatchIPMaxLength.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPMaxLength.IsUnknown()) {
		vyosData["max-length"] = o.LeafQosPolicyShaperHfscClassMatchIPMaxLength.ValueString()
	}
	if !(o.LeafQosPolicyShaperHfscClassMatchIPProtocol.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafQosPolicyShaperHfscClassMatchIPProtocol.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeQosPolicyShaperHfscClassMatchIPDestination.IsNull() || o.NodeQosPolicyShaperHfscClassMatchIPDestination.IsUnknown()) {
		var subModel QosPolicyShaperHfscClassMatchIPDestination
		diags.Append(o.NodeQosPolicyShaperHfscClassMatchIPDestination.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["destination"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperHfscClassMatchIPSource.IsNull() || o.NodeQosPolicyShaperHfscClassMatchIPSource.IsUnknown()) {
		var subModel QosPolicyShaperHfscClassMatchIPSource
		diags.Append(o.NodeQosPolicyShaperHfscClassMatchIPSource.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["source"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperHfscClassMatchIPTCP.IsNull() || o.NodeQosPolicyShaperHfscClassMatchIPTCP.IsUnknown()) {
		var subModel QosPolicyShaperHfscClassMatchIPTCP
		diags.Append(o.NodeQosPolicyShaperHfscClassMatchIPTCP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["tcp"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperHfscClassMatchIP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ip"}})

	// Leafs
	if value, ok := vyosData["dscp"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["max-length"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPMaxLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPMaxLength = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPProtocol = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["destination"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscClassMatchIPDestination{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscClassMatchIPDestination = data

	} else {
		o.NodeQosPolicyShaperHfscClassMatchIPDestination = basetypes.NewObjectNull(QosPolicyShaperHfscClassMatchIPDestination{}.AttributeTypes())
	}
	if value, ok := vyosData["source"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscClassMatchIPSource{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscClassMatchIPSource = data

	} else {
		o.NodeQosPolicyShaperHfscClassMatchIPSource = basetypes.NewObjectNull(QosPolicyShaperHfscClassMatchIPSource{}.AttributeTypes())
	}
	if value, ok := vyosData["tcp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscClassMatchIPTCP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscClassMatchIPTCP = data

	} else {
		o.NodeQosPolicyShaperHfscClassMatchIPTCP = basetypes.NewObjectNull(QosPolicyShaperHfscClassMatchIPTCP{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ip"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperHfscClassMatchIP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"dscp":       types.StringType,
		"max_length": types.StringType,
		"protocol":   types.StringType,

		// Tags

		// Nodes
		"destination": types.ObjectType{AttrTypes: QosPolicyShaperHfscClassMatchIPDestination{}.AttributeTypes()},
		"source":      types.ObjectType{AttrTypes: QosPolicyShaperHfscClassMatchIPSource{}.AttributeTypes()},
		"tcp":         types.ObjectType{AttrTypes: QosPolicyShaperHfscClassMatchIPTCP{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on Differentiated Services Codepoint (DSCP)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  Differentiated Services Codepoint (DSCP) value   |
|  default  |  match DSCP (000000)  |
|  reliability  |  match DSCP (000001)  |
|  throughput  |  match DSCP (000010)  |
|  lowdelay  |  match DSCP (000100)  |
|  priority  |  match DSCP (001000)  |
|  immediate  |  match DSCP (010000)  |
|  flash  |  match DSCP (011000)  |
|  flash-override  |  match DSCP (100000)  |
|  critical  |  match DSCP (101000)  |
|  internet  |  match DSCP (110000)  |
|  network  |  match DSCP (111000)  |
|  AF11  |  High-throughput data  |
|  AF12  |  High-throughput data  |
|  AF13  |  High-throughput data  |
|  AF21  |  Low-latency data  |
|  AF22  |  Low-latency data  |
|  AF23  |  Low-latency data  |
|  AF31  |  Multimedia streaming  |
|  AF32  |  Multimedia streaming  |
|  AF33  |  Multimedia streaming  |
|  AF41  |  Multimedia conferencing  |
|  AF42  |  Multimedia conferencing  |
|  AF43  |  Multimedia conferencing  |
|  CS1  |  Low-priority data  |
|  CS2  |  OAM  |
|  CS3  |  Broadcast video  |
|  CS4  |  Real-time interactive  |
|  CS5  |  Signaling  |
|  CS6  |  Network control  |
|  CS7  |    |
|  EF  |  Expedited Forwarding  |

`,
		},

		"max_length": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum packet length

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Maximum packet/payload length  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Protocol name  |

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIPDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIPSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIPTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
		},
	}
}