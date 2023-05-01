// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesEthernetVifIPvsixAddress describes the resource data model.
type InterfacesEthernetVifIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesEthernetVifIPvsixAddressAutoconf           types.String `tfsdk:"autoconf" json:"autoconf,omitempty"`
	LeafInterfacesEthernetVifIPvsixAddressEuisixfour         types.String `tfsdk:"eui64" json:"eui64,omitempty"`
	LeafInterfacesEthernetVifIPvsixAddressNoDefaultLinkLocal types.String `tfsdk:"no_default_link_local" json:"no-default-link-local,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesEthernetVifIPvsixAddress) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesEthernetVifIPvsixAddressAutoconf.IsNull() && !o.LeafInterfacesEthernetVifIPvsixAddressAutoconf.IsUnknown() {
		jsonData["autoconf"] = o.LeafInterfacesEthernetVifIPvsixAddressAutoconf.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPvsixAddressEuisixfour.IsNull() && !o.LeafInterfacesEthernetVifIPvsixAddressEuisixfour.IsUnknown() {
		jsonData["eui64"] = o.LeafInterfacesEthernetVifIPvsixAddressEuisixfour.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPvsixAddressNoDefaultLinkLocal.IsNull() && !o.LeafInterfacesEthernetVifIPvsixAddressNoDefaultLinkLocal.IsUnknown() {
		jsonData["no-default-link-local"] = o.LeafInterfacesEthernetVifIPvsixAddressNoDefaultLinkLocal.ValueString()
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
func (o *InterfacesEthernetVifIPvsixAddress) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["autoconf"]; ok {
		o.LeafInterfacesEthernetVifIPvsixAddressAutoconf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPvsixAddressAutoconf = basetypes.NewStringNull()
	}

	if value, ok := jsonData["eui64"]; ok {
		o.LeafInterfacesEthernetVifIPvsixAddressEuisixfour = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPvsixAddressEuisixfour = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-link-local"]; ok {
		o.LeafInterfacesEthernetVifIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
