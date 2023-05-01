// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedAsPath.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedMed.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedMed.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedNextHop.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedNextHop.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
