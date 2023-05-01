// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
