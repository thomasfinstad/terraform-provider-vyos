// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceIPoeServerAuthenticationRadiusServer describes the resource data model.
type ServiceIPoeServerAuthenticationRadiusServer struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServiceIPoeServerAuthenticationRadiusServerDisable           types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerKey               types.String `tfsdk:"key" json:"key,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerPort              types.String `tfsdk:"port" json:"port,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerAcctPort          types.String `tfsdk:"acct_port" json:"acct-port,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerDisableAccounting types.String `tfsdk:"disable_accounting" json:"disable-accounting,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerFailTime          types.String `tfsdk:"fail_time" json:"fail-time,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerAuthenticationRadiusServer) GetVyosPath() []string {
	return []string{
		"service",
		"ipoe-server",
		"authentication",
		"radius",
		"server",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerAuthenticationRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `RADIUS server configuration

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  RADIUS server IPv4 address  |

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared secret key

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"acct_port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Accounting port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1813`),
			Computed: true,
		},

		"disable_accounting": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable accounting

`,
		},

		"fail_time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mark server unavailable for <n> seconds on failure

|  Format  |  Description  |
|----------|---------------|
|  u32:0-600  |  Fail time penalty  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerAuthenticationRadiusServer) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceIPoeServerAuthenticationRadiusServerDisable.IsNull() && !o.LeafServiceIPoeServerAuthenticationRadiusServerDisable.IsUnknown() {
		jsonData["disable"] = o.LeafServiceIPoeServerAuthenticationRadiusServerDisable.ValueString()
	}

	if !o.LeafServiceIPoeServerAuthenticationRadiusServerKey.IsNull() && !o.LeafServiceIPoeServerAuthenticationRadiusServerKey.IsUnknown() {
		jsonData["key"] = o.LeafServiceIPoeServerAuthenticationRadiusServerKey.ValueString()
	}

	if !o.LeafServiceIPoeServerAuthenticationRadiusServerPort.IsNull() && !o.LeafServiceIPoeServerAuthenticationRadiusServerPort.IsUnknown() {
		jsonData["port"] = o.LeafServiceIPoeServerAuthenticationRadiusServerPort.ValueString()
	}

	if !o.LeafServiceIPoeServerAuthenticationRadiusServerAcctPort.IsNull() && !o.LeafServiceIPoeServerAuthenticationRadiusServerAcctPort.IsUnknown() {
		jsonData["acct-port"] = o.LeafServiceIPoeServerAuthenticationRadiusServerAcctPort.ValueString()
	}

	if !o.LeafServiceIPoeServerAuthenticationRadiusServerDisableAccounting.IsNull() && !o.LeafServiceIPoeServerAuthenticationRadiusServerDisableAccounting.IsUnknown() {
		jsonData["disable-accounting"] = o.LeafServiceIPoeServerAuthenticationRadiusServerDisableAccounting.ValueString()
	}

	if !o.LeafServiceIPoeServerAuthenticationRadiusServerFailTime.IsNull() && !o.LeafServiceIPoeServerAuthenticationRadiusServerFailTime.IsUnknown() {
		jsonData["fail-time"] = o.LeafServiceIPoeServerAuthenticationRadiusServerFailTime.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceIPoeServerAuthenticationRadiusServer) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["disable"]; ok {
		o.LeafServiceIPoeServerAuthenticationRadiusServerDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerAuthenticationRadiusServerDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["key"]; ok {
		o.LeafServiceIPoeServerAuthenticationRadiusServerKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerAuthenticationRadiusServerKey = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafServiceIPoeServerAuthenticationRadiusServerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerAuthenticationRadiusServerPort = basetypes.NewStringNull()
	}

	if value, ok := jsonData["acct-port"]; ok {
		o.LeafServiceIPoeServerAuthenticationRadiusServerAcctPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerAuthenticationRadiusServerAcctPort = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-accounting"]; ok {
		o.LeafServiceIPoeServerAuthenticationRadiusServerDisableAccounting = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerAuthenticationRadiusServerDisableAccounting = basetypes.NewStringNull()
	}

	if value, ok := jsonData["fail-time"]; ok {
		o.LeafServiceIPoeServerAuthenticationRadiusServerFailTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerAuthenticationRadiusServerFailTime = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
