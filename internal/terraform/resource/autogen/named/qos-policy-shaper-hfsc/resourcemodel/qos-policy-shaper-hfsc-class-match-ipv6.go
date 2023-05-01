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

// QosPolicyShaperHfscClassMatchIPvsix describes the resource data model.
type QosPolicyShaperHfscClassMatchIPvsix struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchIPvsixDscp      types.String `tfsdk:"dscp"`
	LeafQosPolicyShaperHfscClassMatchIPvsixMaxLength types.String `tfsdk:"max_length"`
	LeafQosPolicyShaperHfscClassMatchIPvsixProtocol  types.String `tfsdk:"protocol"`

	// TagNodes

	// Nodes
	NodeQosPolicyShaperHfscClassMatchIPvsixDestination types.Object `tfsdk:"destination"`
	NodeQosPolicyShaperHfscClassMatchIPvsixSource      types.Object `tfsdk:"source"`
	NodeQosPolicyShaperHfscClassMatchIPvsixTCP         types.Object `tfsdk:"tcp"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperHfscClassMatchIPvsix) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ipv6"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperHfscClassMatchIPvsixDscp.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPvsixDscp.IsUnknown()) {
		vyosData["dscp"] = o.LeafQosPolicyShaperHfscClassMatchIPvsixDscp.ValueString()
	}
	if !(o.LeafQosPolicyShaperHfscClassMatchIPvsixMaxLength.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPvsixMaxLength.IsUnknown()) {
		vyosData["max-length"] = o.LeafQosPolicyShaperHfscClassMatchIPvsixMaxLength.ValueString()
	}
	if !(o.LeafQosPolicyShaperHfscClassMatchIPvsixProtocol.IsNull() || o.LeafQosPolicyShaperHfscClassMatchIPvsixProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafQosPolicyShaperHfscClassMatchIPvsixProtocol.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeQosPolicyShaperHfscClassMatchIPvsixDestination.IsNull() || o.NodeQosPolicyShaperHfscClassMatchIPvsixDestination.IsUnknown()) {
		var subModel QosPolicyShaperHfscClassMatchIPvsixDestination
		diags.Append(o.NodeQosPolicyShaperHfscClassMatchIPvsixDestination.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["destination"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperHfscClassMatchIPvsixSource.IsNull() || o.NodeQosPolicyShaperHfscClassMatchIPvsixSource.IsUnknown()) {
		var subModel QosPolicyShaperHfscClassMatchIPvsixSource
		diags.Append(o.NodeQosPolicyShaperHfscClassMatchIPvsixSource.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["source"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperHfscClassMatchIPvsixTCP.IsNull() || o.NodeQosPolicyShaperHfscClassMatchIPvsixTCP.IsUnknown()) {
		var subModel QosPolicyShaperHfscClassMatchIPvsixTCP
		diags.Append(o.NodeQosPolicyShaperHfscClassMatchIPvsixTCP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["tcp"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperHfscClassMatchIPvsix) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ipv6"}})

	// Leafs
	if value, ok := vyosData["dscp"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["max-length"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixMaxLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixMaxLength = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscClassMatchIPvsixProtocol = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["destination"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscClassMatchIPvsixDestination{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscClassMatchIPvsixDestination = data

	} else {
		o.NodeQosPolicyShaperHfscClassMatchIPvsixDestination = basetypes.NewObjectNull(QosPolicyShaperHfscClassMatchIPvsixDestination{}.AttributeTypes())
	}
	if value, ok := vyosData["source"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscClassMatchIPvsixSource{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscClassMatchIPvsixSource = data

	} else {
		o.NodeQosPolicyShaperHfscClassMatchIPvsixSource = basetypes.NewObjectNull(QosPolicyShaperHfscClassMatchIPvsixSource{}.AttributeTypes())
	}
	if value, ok := vyosData["tcp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperHfscClassMatchIPvsixTCP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperHfscClassMatchIPvsixTCP = data

	} else {
		o.NodeQosPolicyShaperHfscClassMatchIPvsixTCP = basetypes.NewObjectNull(QosPolicyShaperHfscClassMatchIPvsixTCP{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper-hfsc", "class", "match", "ipv6"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPvsix) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"dscp":       types.StringType,
		"max_length": types.StringType,
		"protocol":   types.StringType,

		// Tags

		// Nodes
		"destination": types.ObjectType{AttrTypes: QosPolicyShaperHfscClassMatchIPvsixDestination{}.AttributeTypes()},
		"source":      types.ObjectType{AttrTypes: QosPolicyShaperHfscClassMatchIPvsixSource{}.AttributeTypes()},
		"tcp":         types.ObjectType{AttrTypes: QosPolicyShaperHfscClassMatchIPvsixTCP{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: QosPolicyShaperHfscClassMatchIPvsixDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIPvsixSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIPvsixTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
		},
	}
}