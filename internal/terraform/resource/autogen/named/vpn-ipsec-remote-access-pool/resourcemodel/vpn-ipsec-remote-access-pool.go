// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsecRemoteAccessPool describes the resource data model.
type VpnIPsecRemoteAccessPool struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafVpnIPsecRemoteAccessPoolExclude    types.String `tfsdk:"exclude" json:"exclude,omitempty"`
	LeafVpnIPsecRemoteAccessPoolPrefix     types.String `tfsdk:"prefix" json:"prefix,omitempty"`
	LeafVpnIPsecRemoteAccessPoolNameServer types.String `tfsdk:"name_server" json:"name-server,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecRemoteAccessPool) GetVyosPath() []string {
	return []string{
		"vpn",
		"ipsec",
		"remote-access",
		"pool",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessPool) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IP address pool for remote access users

`,
		},

		// LeafNodes

		"exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IPv4 or IPv6 pool prefix exclusions

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Local IPv4 pool prefix exclusion  |
|  ipv6net  |  Local IPv6 pool prefix exclusion  |

`,
		},

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IPv4 or IPv6 pool prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Local IPv4 pool prefix  |
|  ipv6net  |  Local IPv6 pool prefix  |

`,
		},

		"name_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
|  ipv6  |  Domain Name Server (DNS) IPv6 address  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecRemoteAccessPool) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVpnIPsecRemoteAccessPoolExclude.IsNull() && !o.LeafVpnIPsecRemoteAccessPoolExclude.IsUnknown() {
		jsonData["exclude"] = o.LeafVpnIPsecRemoteAccessPoolExclude.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessPoolPrefix.IsNull() && !o.LeafVpnIPsecRemoteAccessPoolPrefix.IsUnknown() {
		jsonData["prefix"] = o.LeafVpnIPsecRemoteAccessPoolPrefix.ValueString()
	}

	if !o.LeafVpnIPsecRemoteAccessPoolNameServer.IsNull() && !o.LeafVpnIPsecRemoteAccessPoolNameServer.IsUnknown() {
		jsonData["name-server"] = o.LeafVpnIPsecRemoteAccessPoolNameServer.ValueString()
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
func (o *VpnIPsecRemoteAccessPool) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["exclude"]; ok {
		o.LeafVpnIPsecRemoteAccessPoolExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessPoolExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["prefix"]; ok {
		o.LeafVpnIPsecRemoteAccessPoolPrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessPoolPrefix = basetypes.NewStringNull()
	}

	if value, ok := jsonData["name-server"]; ok {
		o.LeafVpnIPsecRemoteAccessPoolNameServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessPoolNameServer = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
