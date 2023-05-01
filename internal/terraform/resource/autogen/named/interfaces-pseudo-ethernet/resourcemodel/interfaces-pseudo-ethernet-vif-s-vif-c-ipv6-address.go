// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesPseudoEthernetVifSVifCIPvsixAddress describes the resource data model.
type InterfacesPseudoEthernetVifSVifCIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressAutoconf           types.String `tfsdk:"autoconf" json:"autoconf,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressEuisixfour         types.String `tfsdk:"eui64" json:"eui64,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressNoDefaultLinkLocal types.String `tfsdk:"no_default_link_local" json:"no-default-link-local,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVifSVifCIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesPseudoEthernetVifSVifCIPvsixAddress) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressAutoconf.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressAutoconf.IsUnknown() {
		jsonData["autoconf"] = o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressAutoconf.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressEuisixfour.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressEuisixfour.IsUnknown() {
		jsonData["eui64"] = o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressEuisixfour.ValueString()
	}

	if !o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressNoDefaultLinkLocal.IsNull() && !o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressNoDefaultLinkLocal.IsUnknown() {
		jsonData["no-default-link-local"] = o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressNoDefaultLinkLocal.ValueString()
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
func (o *InterfacesPseudoEthernetVifSVifCIPvsixAddress) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["autoconf"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressAutoconf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressAutoconf = basetypes.NewStringNull()
	}

	if value, ok := jsonData["eui64"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressEuisixfour = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressEuisixfour = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-default-link-local"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCIPvsixAddressNoDefaultLinkLocal = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
