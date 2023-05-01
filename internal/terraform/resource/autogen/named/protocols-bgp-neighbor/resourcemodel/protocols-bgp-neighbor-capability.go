// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborCapability describes the resource data model.
type ProtocolsBgpNeighborCapability struct {
	// LeafNodes
	LeafProtocolsBgpNeighborCapabilityDynamic         types.String `tfsdk:"dynamic" json:"dynamic,omitempty"`
	LeafProtocolsBgpNeighborCapabilityExtendedNexthop types.String `tfsdk:"extended_nexthop" json:"extended-nexthop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborCapability) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dynamic": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise dynamic capability to this neighbor

`,
		},

		"extended_nexthop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise extended-nexthop capability to this neighbor

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborCapability) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborCapabilityDynamic.IsNull() && !o.LeafProtocolsBgpNeighborCapabilityDynamic.IsUnknown() {
		jsonData["dynamic"] = o.LeafProtocolsBgpNeighborCapabilityDynamic.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborCapabilityExtendedNexthop.IsNull() && !o.LeafProtocolsBgpNeighborCapabilityExtendedNexthop.IsUnknown() {
		jsonData["extended-nexthop"] = o.LeafProtocolsBgpNeighborCapabilityExtendedNexthop.ValueString()
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
func (o *ProtocolsBgpNeighborCapability) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["dynamic"]; ok {
		o.LeafProtocolsBgpNeighborCapabilityDynamic = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborCapabilityDynamic = basetypes.NewStringNull()
	}

	if value, ok := jsonData["extended-nexthop"]; ok {
		o.LeafProtocolsBgpNeighborCapabilityExtendedNexthop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborCapabilityExtendedNexthop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
