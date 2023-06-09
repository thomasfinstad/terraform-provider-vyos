// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceLldpInterface describes the resource data model.
type ServiceLldpInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServiceLldpInterfaceDisable types.String `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
	NodeServiceLldpInterfaceLocation *ServiceLldpInterfaceLocation `tfsdk:"location" json:"location,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceLldpInterface) GetVyosPath() []string {
	return []string{
		"service",
		"lldp",
		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceLldpInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Location data for interface

|  Format  |  Description  |
|----------|---------------|
|  all  |  Location data all interfaces  |
|  txt  |  Location data for a specific interface  |

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

		"location": schema.SingleNestedAttribute{
			Attributes: ServiceLldpInterfaceLocation{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `LLDP-MED location data

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceLldpInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceLldpInterfaceDisable.IsNull() && !o.LeafServiceLldpInterfaceDisable.IsUnknown() {
		jsonData["disable"] = o.LeafServiceLldpInterfaceDisable.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeServiceLldpInterfaceLocation).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeServiceLldpInterfaceLocation)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["location"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceLldpInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["disable"]; ok {
		o.LeafServiceLldpInterfaceDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceLldpInterfaceDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["location"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeServiceLldpInterfaceLocation = &ServiceLldpInterfaceLocation{}

		err = json.Unmarshal(subJSONStr, o.NodeServiceLldpInterfaceLocation)
		if err != nil {
			return err
		}
	}

	return nil
}
