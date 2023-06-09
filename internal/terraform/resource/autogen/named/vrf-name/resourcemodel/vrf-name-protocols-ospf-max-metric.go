// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsOspfMaxMetric describes the resource data model.
type VrfNameProtocolsOspfMaxMetric struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsOspfMaxMetricRouterLsa *VrfNameProtocolsOspfMaxMetricRouterLsa `tfsdk:"router_lsa" json:"router-lsa,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfMaxMetric) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"router_lsa": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfMaxMetricRouterLsa{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise own Router-LSA with infinite distance (stub router)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfMaxMetric) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfMaxMetricRouterLsa).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfMaxMetricRouterLsa)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["router-lsa"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfMaxMetric) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["router-lsa"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfMaxMetricRouterLsa = &VrfNameProtocolsOspfMaxMetricRouterLsa{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfMaxMetricRouterLsa)
		if err != nil {
			return err
		}
	}

	return nil
}
