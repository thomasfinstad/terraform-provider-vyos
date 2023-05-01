// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsOspfvthreeRedistribute describes the resource data model.
type VrfNameProtocolsOspfvthreeRedistribute struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsOspfvthreeRedistributeBgp       *VrfNameProtocolsOspfvthreeRedistributeBgp       `tfsdk:"bgp" json:"bgp,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeConnected *VrfNameProtocolsOspfvthreeRedistributeConnected `tfsdk:"connected" json:"connected,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeKernel    *VrfNameProtocolsOspfvthreeRedistributeKernel    `tfsdk:"kernel" json:"kernel,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeRIPng     *VrfNameProtocolsOspfvthreeRedistributeRIPng     `tfsdk:"ripng" json:"ripng,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeBabel     *VrfNameProtocolsOspfvthreeRedistributeBabel     `tfsdk:"babel" json:"babel,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeStatic    *VrfNameProtocolsOspfvthreeRedistributeStatic    `tfsdk:"static" json:"static,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeRedistribute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeBgp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute BGP routes

`,
		},

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeConnected{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeKernel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes

`,
		},

		"ripng": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeRIPng{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute RIPNG routes

`,
		},

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeBabel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeStatic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfvthreeRedistribute) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeRedistributeBgp).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeRedistributeBgp)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["bgp"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeRedistributeConnected).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeRedistributeConnected)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["connected"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeRedistributeKernel).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeRedistributeKernel)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["kernel"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeRedistributeRIPng).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeRedistributeRIPng)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ripng"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeRedistributeBabel).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeRedistributeBabel)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["babel"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeRedistributeStatic).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeRedistributeStatic)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["static"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfvthreeRedistribute) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["bgp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeRedistributeBgp = &VrfNameProtocolsOspfvthreeRedistributeBgp{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeRedistributeBgp)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["connected"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeRedistributeConnected = &VrfNameProtocolsOspfvthreeRedistributeConnected{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeRedistributeConnected)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["kernel"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeRedistributeKernel = &VrfNameProtocolsOspfvthreeRedistributeKernel{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeRedistributeKernel)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ripng"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeRedistributeRIPng = &VrfNameProtocolsOspfvthreeRedistributeRIPng{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeRedistributeRIPng)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["babel"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeRedistributeBabel = &VrfNameProtocolsOspfvthreeRedistributeBabel{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeRedistributeBabel)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["static"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeRedistributeStatic = &VrfNameProtocolsOspfvthreeRedistributeStatic{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeRedistributeStatic)
		if err != nil {
			return err
		}
	}

	return nil
}
