// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedAsPath.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedMed.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedMed.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedNextHop.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedNextHop.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
