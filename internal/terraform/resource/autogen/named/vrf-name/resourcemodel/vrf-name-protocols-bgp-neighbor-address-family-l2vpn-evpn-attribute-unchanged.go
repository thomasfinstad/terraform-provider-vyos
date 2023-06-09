// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Send AS path unchanged

`,
		},

		"med": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Send multi-exit discriminator unchanged

`,
		},

		"next_hop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Send nexthop unchanged

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedAsPath.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedMed.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedMed.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedNextHop.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedNextHop.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
