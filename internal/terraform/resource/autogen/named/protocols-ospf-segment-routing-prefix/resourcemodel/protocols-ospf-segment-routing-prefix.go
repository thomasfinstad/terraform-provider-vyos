// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsOspfSegmentRoutingPrefix describes the resource data model.
type ProtocolsOspfSegmentRoutingPrefix struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsOspfSegmentRoutingPrefixIndex *ProtocolsOspfSegmentRoutingPrefixIndex `tfsdk:"index" json:"index,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfSegmentRoutingPrefix) GetVyosPath() []string {
	return []string{
		"protocols",
		"ospf",
		"segment-routing",
		"prefix",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfSegmentRoutingPrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 prefix segment/label mapping

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 prefix segment  |

`,
		},

		// LeafNodes

		// TagNodes

		// Nodes

		"index": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfSegmentRoutingPrefixIndex{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify the index value of prefix segment/label ID

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsOspfSegmentRoutingPrefix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsOspfSegmentRoutingPrefixIndex).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsOspfSegmentRoutingPrefixIndex)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["index"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsOspfSegmentRoutingPrefix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["index"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsOspfSegmentRoutingPrefixIndex = &ProtocolsOspfSegmentRoutingPrefixIndex{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsOspfSegmentRoutingPrefixIndex)
		if err != nil {
			return err
		}
	}

	return nil
}
