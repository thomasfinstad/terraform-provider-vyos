// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesTunnelIPvsixAddress describes the resource data model.
type InterfacesTunnelIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesTunnelIPvsixAddressAutoconf           types.String `tfsdk:"autoconf" json:"autoconf,omitempty"`
	LeafInterfacesTunnelIPvsixAddressEuisixfour         types.String `tfsdk:"eui64" json:"eui64,omitempty"`
	LeafInterfacesTunnelIPvsixAddressNoDefaultLinkLocal types.String `tfsdk:"no_default_link_local" json:"no-default-link-local,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesTunnelIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"autoconf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
		},

		"eui64": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

|  Format  |  Description  |
|----------|---------------|
|  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |

`,
		},

		"no_default_link_local": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remove the default link-local address from the interface

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesTunnelIPvsixAddress) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesTunnelIPvsixAddressAutoconf.IsNull() && !o.LeafInterfacesTunnelIPvsixAddressAutoconf.IsUnknown() {
		jsonData["autoconf"] = o.LeafInterfacesTunnelIPvsixAddressAutoconf.ValueString()
	}

	if !o.LeafInterfacesTunnelIPvsixAddressEuisixfour.IsNull() && !o.LeafInterfacesTunnelIPvsixAddressEuisixfour.IsUnknown() {
		jsonData["eui64"] = o.LeafInterfacesTunnelIPvsixAddressEuisixfour.ValueString()
	}

	if !o.LeafInterfacesTunnelIPvsixAddressNoDefaultLinkLocal.IsNull() && !o.LeafInterfacesTunnelIPvsixAddressNoDefaultLinkLocal.IsUnknown() {
		jsonData["no-default-link-local"] = o.LeafInterfacesTunnelIPvsixAddressNoDefaultLinkLocal.ValueString()
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
func (o *InterfacesTunnelIPvsixAddress) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["autoconf"]; ok {
		o.LeafInterfacesTunnelIPvsixAddressAutoconf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesTunnelIPvsixAddressAutoconf = basetypes.NewStringNull()
	}

	if value, ok := jsonData["eui64"]; ok {
		o.LeafInterfacesTunnelIPvsixAddressEuisixfour = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesTunnelIPvsixAddressEuisixfour = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-link-local"]; ok {
		o.LeafInterfacesTunnelIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesTunnelIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
