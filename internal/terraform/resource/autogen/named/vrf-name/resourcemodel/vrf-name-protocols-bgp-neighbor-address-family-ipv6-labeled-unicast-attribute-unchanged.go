// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
