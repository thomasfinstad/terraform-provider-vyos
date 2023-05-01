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

// HighAvailabilityVirtualServer describes the resource data model.
type HighAvailabilityVirtualServer struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafHighAvailabilityVirtualServerAlgorithm          types.String `tfsdk:"algorithm"`
	LeafHighAvailabilityVirtualServerDelayLoop          types.String `tfsdk:"delay_loop"`
	LeafHighAvailabilityVirtualServerForwardMethod      types.String `tfsdk:"forward_method"`
	LeafHighAvailabilityVirtualServerFwmark             types.String `tfsdk:"fwmark"`
	LeafHighAvailabilityVirtualServerPort               types.String `tfsdk:"port"`
	LeafHighAvailabilityVirtualServerPersistenceTimeout types.String `tfsdk:"persistence_timeout"`
	LeafHighAvailabilityVirtualServerProtocol           types.String `tfsdk:"protocol"`

	// TagNodes
	TagHighAvailabilityVirtualServerRealServer types.Map `tfsdk:"real_server"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVirtualServer) GetVyosPath() []string {
	return []string{
		"high-availability",
		"virtual-server",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *HighAvailabilityVirtualServer) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"high-availability", "virtual-server"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafHighAvailabilityVirtualServerAlgorithm.IsNull() || o.LeafHighAvailabilityVirtualServerAlgorithm.IsUnknown()) {
		vyosData["algorithm"] = o.LeafHighAvailabilityVirtualServerAlgorithm.ValueString()
	}
	if !(o.LeafHighAvailabilityVirtualServerDelayLoop.IsNull() || o.LeafHighAvailabilityVirtualServerDelayLoop.IsUnknown()) {
		vyosData["delay-loop"] = o.LeafHighAvailabilityVirtualServerDelayLoop.ValueString()
	}
	if !(o.LeafHighAvailabilityVirtualServerForwardMethod.IsNull() || o.LeafHighAvailabilityVirtualServerForwardMethod.IsUnknown()) {
		vyosData["forward-method"] = o.LeafHighAvailabilityVirtualServerForwardMethod.ValueString()
	}
	if !(o.LeafHighAvailabilityVirtualServerFwmark.IsNull() || o.LeafHighAvailabilityVirtualServerFwmark.IsUnknown()) {
		vyosData["fwmark"] = o.LeafHighAvailabilityVirtualServerFwmark.ValueString()
	}
	if !(o.LeafHighAvailabilityVirtualServerPort.IsNull() || o.LeafHighAvailabilityVirtualServerPort.IsUnknown()) {
		vyosData["port"] = o.LeafHighAvailabilityVirtualServerPort.ValueString()
	}
	if !(o.LeafHighAvailabilityVirtualServerPersistenceTimeout.IsNull() || o.LeafHighAvailabilityVirtualServerPersistenceTimeout.IsUnknown()) {
		vyosData["persistence-timeout"] = o.LeafHighAvailabilityVirtualServerPersistenceTimeout.ValueString()
	}
	if !(o.LeafHighAvailabilityVirtualServerProtocol.IsNull() || o.LeafHighAvailabilityVirtualServerProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafHighAvailabilityVirtualServerProtocol.ValueString()
	}

	// Tags
	if !(o.TagHighAvailabilityVirtualServerRealServer.IsNull() || o.TagHighAvailabilityVirtualServerRealServer.IsUnknown()) {
		subModel := make(map[string]HighAvailabilityVirtualServerRealServer)
		diags.Append(o.TagHighAvailabilityVirtualServerRealServer.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["real-server"] = subData
	}

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *HighAvailabilityVirtualServer) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"high-availability", "virtual-server"}})

	// Leafs
	if value, ok := vyosData["algorithm"]; ok {
		o.LeafHighAvailabilityVirtualServerAlgorithm = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerAlgorithm = basetypes.NewStringNull()
	}
	if value, ok := vyosData["delay-loop"]; ok {
		o.LeafHighAvailabilityVirtualServerDelayLoop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerDelayLoop = basetypes.NewStringNull()
	}
	if value, ok := vyosData["forward-method"]; ok {
		o.LeafHighAvailabilityVirtualServerForwardMethod = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerForwardMethod = basetypes.NewStringNull()
	}
	if value, ok := vyosData["fwmark"]; ok {
		o.LeafHighAvailabilityVirtualServerFwmark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerFwmark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafHighAvailabilityVirtualServerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerPort = basetypes.NewStringNull()
	}
	if value, ok := vyosData["persistence-timeout"]; ok {
		o.LeafHighAvailabilityVirtualServerPersistenceTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerPersistenceTimeout = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafHighAvailabilityVirtualServerProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerProtocol = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["real-server"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: HighAvailabilityVirtualServerRealServer{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagHighAvailabilityVirtualServerRealServer = data
	} else {
		o.TagHighAvailabilityVirtualServerRealServer = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"high-availability", "virtual-server"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o HighAvailabilityVirtualServer) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"algorithm":           types.StringType,
		"delay_loop":          types.StringType,
		"forward_method":      types.StringType,
		"fwmark":              types.StringType,
		"port":                types.StringType,
		"persistence_timeout": types.StringType,
		"protocol":            types.StringType,

		// Tags
		"real_server": types.MapType{ElemType: types.ObjectType{AttrTypes: HighAvailabilityVirtualServerRealServer{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVirtualServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Load-balancing virtual server address

`,
		},

		// LeafNodes

		"algorithm": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Schedule algorithm (default - least-connection)

|  Format  |  Description  |
|----------|---------------|
|  round-robin  |  Round robin  |
|  weighted-round-robin  |  Weighted round robin  |
|  least-connection  |  Least connection  |
|  weighted-least-connection  |  Weighted least connection  |
|  source-hashing  |  Source hashing  |
|  destination-hashing  |  Destination hashing  |
|  locality-based-least-connection  |  Locality-Based least connection  |

`,

			// Default:          stringdefault.StaticString(`least-connection`),
			Computed: true,
		},

		"delay_loop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval between health-checks (in seconds)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-600  |  Interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"forward_method": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Forwarding method

|  Format  |  Description  |
|----------|---------------|
|  direct  |  Direct routing  |
|  nat  |  NAT  |
|  tunnel  |  Tunneling  |

`,

			// Default:          stringdefault.StaticString(`nat`),
			Computed: true,
		},

		"fwmark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Match firewall mark value  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Numeric IP port  |

`,
		},

		"persistence_timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for persistent connections

|  Format  |  Description  |
|----------|---------------|
|  u32:1-86400  |  Timeout for persistent connections  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol for port checks

|  Format  |  Description  |
|----------|---------------|
|  tcp  |  TCP  |
|  udp  |  UDP  |

`,

			// Default:          stringdefault.StaticString(`tcp`),
			Computed: true,
		},

		// TagNodes

		"real_server": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: HighAvailabilityVirtualServerRealServer{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Real server address

`,
		},

		// Nodes

	}
}