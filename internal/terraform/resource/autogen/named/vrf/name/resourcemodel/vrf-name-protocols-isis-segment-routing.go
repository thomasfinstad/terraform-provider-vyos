// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsIsisSegmentRouting describes the resource data model.
type VrfNameProtocolsIsisSegmentRouting struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth types.String `tfsdk:"maximum_label_depth" vyos:"maximum-label-depth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsIsisSegmentRoutingPrefix bool `tfsdk:"prefix" vyos:"prefix,child"`

	// Nodes
	NodeVrfNameProtocolsIsisSegmentRoutingGlobalBlock *VrfNameProtocolsIsisSegmentRoutingGlobalBlock `tfsdk:"global_block" vyos:"global-block,omitempty"`
	NodeVrfNameProtocolsIsisSegmentRoutingLocalBlock  *VrfNameProtocolsIsisSegmentRoutingLocalBlock  `tfsdk:"local_block" vyos:"local-block,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRouting) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"maximum_label_depth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum MPLS labels allowed for this router

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-16  |  MPLS label depth  |

`,
		},

		// Nodes

		"global_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingGlobalBlock{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Segment Routing Global Block label range

`,
		},

		"local_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingLocalBlock{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Segment Routing Local Block label range

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisSegmentRouting) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth.IsNull() && !o.LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth.IsUnknown() {
		jsonData["maximum-label-depth"] = o.LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth.ValueString()
	}

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisSegmentRoutingGlobalBlock).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisSegmentRoutingGlobalBlock)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["global-block"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisSegmentRoutingLocalBlock).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisSegmentRoutingLocalBlock)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["local-block"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisSegmentRouting) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["maximum-label-depth"]; ok {
		o.LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth = basetypes.NewStringNull()
	}

	// Nodes
	if value, ok := jsonData["global-block"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsIsisSegmentRoutingGlobalBlock = &VrfNameProtocolsIsisSegmentRoutingGlobalBlock{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisSegmentRoutingGlobalBlock)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["local-block"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsIsisSegmentRoutingLocalBlock = &VrfNameProtocolsIsisSegmentRoutingLocalBlock{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisSegmentRoutingLocalBlock)
		if err != nil {
			return err
		}
	}

	return nil
}
