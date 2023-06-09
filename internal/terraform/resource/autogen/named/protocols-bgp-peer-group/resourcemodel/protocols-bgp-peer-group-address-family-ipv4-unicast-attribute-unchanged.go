// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedAsPath  types.String `tfsdk:"as_path" json:"as-path,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedMed     types.String `tfsdk:"med" json:"med,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedNextHop types.String `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedAsPath.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedAsPath.IsUnknown() {
		jsonData["as-path"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedAsPath.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedMed.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedMed.IsUnknown() {
		jsonData["med"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedMed.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedNextHop.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedNextHop.IsUnknown() {
		jsonData["next-hop"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedNextHop.ValueString()
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-path"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}

	if value, ok := jsonData["med"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedMed = basetypes.NewStringNull()
	}

	if value, ok := jsonData["next-hop"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
