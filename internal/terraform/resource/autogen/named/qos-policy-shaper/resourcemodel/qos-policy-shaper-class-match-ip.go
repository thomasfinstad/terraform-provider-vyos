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

// QosPolicyShaperClassMatchIP describes the resource data model.
type QosPolicyShaperClassMatchIP struct {
	// LeafNodes
	LeafQosPolicyShaperClassMatchIPDscp      types.String `tfsdk:"dscp"`
	LeafQosPolicyShaperClassMatchIPMaxLength types.String `tfsdk:"max_length"`
	LeafQosPolicyShaperClassMatchIPProtocol  types.String `tfsdk:"protocol"`

	// TagNodes

	// Nodes
	NodeQosPolicyShaperClassMatchIPDestination types.Object `tfsdk:"destination"`
	NodeQosPolicyShaperClassMatchIPSource      types.Object `tfsdk:"source"`
	NodeQosPolicyShaperClassMatchIPTCP         types.Object `tfsdk:"tcp"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperClassMatchIP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match", "ip"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperClassMatchIPDscp.IsNull() || o.LeafQosPolicyShaperClassMatchIPDscp.IsUnknown()) {
		vyosData["dscp"] = o.LeafQosPolicyShaperClassMatchIPDscp.ValueString()
	}
	if !(o.LeafQosPolicyShaperClassMatchIPMaxLength.IsNull() || o.LeafQosPolicyShaperClassMatchIPMaxLength.IsUnknown()) {
		vyosData["max-length"] = o.LeafQosPolicyShaperClassMatchIPMaxLength.ValueString()
	}
	if !(o.LeafQosPolicyShaperClassMatchIPProtocol.IsNull() || o.LeafQosPolicyShaperClassMatchIPProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafQosPolicyShaperClassMatchIPProtocol.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeQosPolicyShaperClassMatchIPDestination.IsNull() || o.NodeQosPolicyShaperClassMatchIPDestination.IsUnknown()) {
		var subModel QosPolicyShaperClassMatchIPDestination
		diags.Append(o.NodeQosPolicyShaperClassMatchIPDestination.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["destination"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperClassMatchIPSource.IsNull() || o.NodeQosPolicyShaperClassMatchIPSource.IsUnknown()) {
		var subModel QosPolicyShaperClassMatchIPSource
		diags.Append(o.NodeQosPolicyShaperClassMatchIPSource.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["source"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperClassMatchIPTCP.IsNull() || o.NodeQosPolicyShaperClassMatchIPTCP.IsUnknown()) {
		var subModel QosPolicyShaperClassMatchIPTCP
		diags.Append(o.NodeQosPolicyShaperClassMatchIPTCP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["tcp"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperClassMatchIP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match", "ip"}})

	// Leafs
	if value, ok := vyosData["dscp"]; ok {
		o.LeafQosPolicyShaperClassMatchIPDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["max-length"]; ok {
		o.LeafQosPolicyShaperClassMatchIPMaxLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPMaxLength = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafQosPolicyShaperClassMatchIPProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPProtocol = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["destination"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperClassMatchIPDestination{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperClassMatchIPDestination = data

	} else {
		o.NodeQosPolicyShaperClassMatchIPDestination = basetypes.NewObjectNull(QosPolicyShaperClassMatchIPDestination{}.AttributeTypes())
	}
	if value, ok := vyosData["source"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperClassMatchIPSource{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperClassMatchIPSource = data

	} else {
		o.NodeQosPolicyShaperClassMatchIPSource = basetypes.NewObjectNull(QosPolicyShaperClassMatchIPSource{}.AttributeTypes())
	}
	if value, ok := vyosData["tcp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperClassMatchIPTCP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperClassMatchIPTCP = data

	} else {
		o.NodeQosPolicyShaperClassMatchIPTCP = basetypes.NewObjectNull(QosPolicyShaperClassMatchIPTCP{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match", "ip"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperClassMatchIP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"dscp":       types.StringType,
		"max_length": types.StringType,
		"protocol":   types.StringType,

		// Tags

		// Nodes
		"destination": types.ObjectType{AttrTypes: QosPolicyShaperClassMatchIPDestination{}.AttributeTypes()},
		"source":      types.ObjectType{AttrTypes: QosPolicyShaperClassMatchIPSource{}.AttributeTypes()},
		"tcp":         types.ObjectType{AttrTypes: QosPolicyShaperClassMatchIPTCP{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatchIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: QosPolicyShaperClassMatchIPDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
		},
	}
}