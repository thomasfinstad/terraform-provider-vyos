// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsecRemoteAccessConnection describes the resource data model.
type VpnIPsecRemoteAccessConnection struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafVpnIPsecRemoteAccessConnectionDescrIPtion  types.String `tfsdk:"description" json:"description,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionDisable      types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionEspGroup     types.String `tfsdk:"esp_group" json:"esp-group,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionIkeGroup     types.String `tfsdk:"ike_group" json:"ike-group,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionLocalAddress types.String `tfsdk:"local_address" json:"local-address,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionTimeout      types.String `tfsdk:"timeout" json:"timeout,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionPool         types.String `tfsdk:"pool" json:"pool,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionUnique       types.String `tfsdk:"unique" json:"unique,omitempty"`

	// TagNodes

	// Nodes
	NodeVpnIPsecRemoteAccessConnectionAuthentication *VpnIPsecRemoteAccessConnectionAuthentication `tfsdk:"authentication" json:"authentication,omitempty"`
	NodeVpnIPsecRemoteAccessConnectionLocal          *VpnIPsecRemoteAccessConnectionLocal          `tfsdk:"local" json:"local,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecRemoteAccessConnection) GetVyosPath() []string {
	return []string{
		"vpn",
		"ipsec",
		"remote-access",
		"connection",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessConnection) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IKEv2 VPN connection name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Connection name  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"esp_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
		},

		"ike_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Internet Key Exchange (IKE) group name

`,
		},

		"local_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 or IPv6 address of a local interface to use for VPN

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of a local interface for VPN  |
|  ipv6  |  IPv6 address of a local interface for VPN  |
|  any  |  Allow any IPv4 address present on the system to be used for VPN  |

`,
		},

		"timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Timeout to close connection if no data is transmitted

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable inactivity checks  |
|  u32:1-86400  |  Timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`28800`),
			Computed: true,
		},

		"pool": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address pool

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Predefined IP pool name  |
|  dhcp  |  Forward requests for virtual IP addresses to a DHCP server  |
|  radius  |  Forward requests for virtual IP addresses to a RADIUS server  |

`,
		},

		"unique": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Connection uniqueness enforcement policy

|  Format  |  Description  |
|----------|---------------|
|  never  |  Never enforce connection uniqueness  |
|  keep  |  Reject new connection attempts if the same user already has an active connection  |
|  replace  |  Delete any existing connection if a new one for the same user gets established  |

`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication for remote access

`,
		},

		"local": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionLocal{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Local parameters for interesting traffic

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecRemoteAccessConnection) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVpnIPsecRemoteAccessConnectionDescrIPtion.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafVpnIPsecRemoteAccessConnectionDescrIPtion.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionDisable.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionDisable.IsUnknown() {
		jsonData["disable"] = o.LeafVpnIPsecRemoteAccessConnectionDisable.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionEspGroup.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionEspGroup.IsUnknown() {
		jsonData["esp-group"] = o.LeafVpnIPsecRemoteAccessConnectionEspGroup.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionIkeGroup.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionIkeGroup.IsUnknown() {
		jsonData["ike-group"] = o.LeafVpnIPsecRemoteAccessConnectionIkeGroup.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionLocalAddress.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionLocalAddress.IsUnknown() {
		jsonData["local-address"] = o.LeafVpnIPsecRemoteAccessConnectionLocalAddress.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionTimeout.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionTimeout.IsUnknown() {
		jsonData["timeout"] = o.LeafVpnIPsecRemoteAccessConnectionTimeout.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionPool.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionPool.IsUnknown() {
		jsonData["pool"] = o.LeafVpnIPsecRemoteAccessConnectionPool.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessConnectionUnique.IsNull() && !o.LeafVpnIPsecRemoteAccessConnectionUnique.IsUnknown() {
		jsonData["unique"] = o.LeafVpnIPsecRemoteAccessConnectionUnique.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVpnIPsecRemoteAccessConnectionAuthentication).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVpnIPsecRemoteAccessConnectionAuthentication)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["authentication"] = subData
	}

	if !reflect.ValueOf(o.NodeVpnIPsecRemoteAccessConnectionLocal).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVpnIPsecRemoteAccessConnectionLocal)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["local"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecRemoteAccessConnection) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["esp-group"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionEspGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionEspGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ike-group"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionIkeGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionIkeGroup = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local-address"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionLocalAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionLocalAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["timeout"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["pool"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionPool = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionPool = basetypes.NewStringNull()
	}

	if value, ok := jsonData["unique"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionUnique = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionUnique = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["authentication"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVpnIPsecRemoteAccessConnectionAuthentication = &VpnIPsecRemoteAccessConnectionAuthentication{}

		err = json.Unmarshal(subJSONStr, o.NodeVpnIPsecRemoteAccessConnectionAuthentication)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["local"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVpnIPsecRemoteAccessConnectionLocal = &VpnIPsecRemoteAccessConnectionLocal{}

		err = json.Unmarshal(subJSONStr, o.NodeVpnIPsecRemoteAccessConnectionLocal)
		if err != nil {
			return err
		}
	}

	return nil
}
