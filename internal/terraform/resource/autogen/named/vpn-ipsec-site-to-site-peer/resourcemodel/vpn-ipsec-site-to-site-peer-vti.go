// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsecSiteToSitePeerVti describes the resource data model.
type VpnIPsecSiteToSitePeerVti struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerVtiBind     types.String `tfsdk:"bind" json:"bind,omitempty"`
	LeafVpnIPsecSiteToSitePeerVtiEspGroup types.String `tfsdk:"esp_group" json:"esp-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerVti) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bind": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VTI tunnel interface associated with this configuration

`,
		},

		"esp_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecSiteToSitePeerVti) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVpnIPsecSiteToSitePeerVtiBind.IsNull() && !o.LeafVpnIPsecSiteToSitePeerVtiBind.IsUnknown() {
		jsonData["bind"] = o.LeafVpnIPsecSiteToSitePeerVtiBind.ValueString()
	}

	if !o.LeafVpnIPsecSiteToSitePeerVtiEspGroup.IsNull() && !o.LeafVpnIPsecSiteToSitePeerVtiEspGroup.IsUnknown() {
		jsonData["esp-group"] = o.LeafVpnIPsecSiteToSitePeerVtiEspGroup.ValueString()
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
func (o *VpnIPsecSiteToSitePeerVti) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["bind"]; ok {
		o.LeafVpnIPsecSiteToSitePeerVtiBind = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerVtiBind = basetypes.NewStringNull()
	}

	if value, ok := jsonData["esp-group"]; ok {
		o.LeafVpnIPsecSiteToSitePeerVtiEspGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerVtiEspGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
