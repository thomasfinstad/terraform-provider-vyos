// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
