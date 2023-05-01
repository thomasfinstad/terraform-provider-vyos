// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpnExport types.String `tfsdk:"export" json:"export,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `For routes leaked from current address-family to VPN

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpnExport.IsNull() && !o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpnExport.IsUnknown() {
		jsonData["export"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpnExport.ValueString()
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
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpnExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpnExport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
