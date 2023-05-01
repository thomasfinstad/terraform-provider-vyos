// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedAsPath.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedMed.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedMed.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedNextHop.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedNextHop.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
