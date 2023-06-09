// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyRoundRobin describes the resource data model.
type QosPolicyRoundRobin struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafQosPolicyRoundRobinDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes
	TagQosPolicyRoundRobinClass *map[string]QosPolicyRoundRobinClass `tfsdk:"class" json:"class,omitempty"`

	// Nodes
	NodeQosPolicyRoundRobinDefault *QosPolicyRoundRobinDefault `tfsdk:"default" json:"default,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyRoundRobin) GetVyosPath() []string {
	return []string{
		"qos",
		"policy",
		"round-robin",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobin) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Deficit Round Robin Scheduler

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		// TagNodes

		"class": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: QosPolicyRoundRobinClass{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Class ID

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4095  |  Class Identifier  |

`,
		},

		// Nodes

		"default": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinDefault{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Default policy

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyRoundRobin) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyRoundRobinDescrIPtion.IsNull() && !o.LeafQosPolicyRoundRobinDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafQosPolicyRoundRobinDescrIPtion.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagQosPolicyRoundRobinClass).IsZero() {
		subJSONStr, err := json.Marshal(o.TagQosPolicyRoundRobinClass)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["class"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeQosPolicyRoundRobinDefault).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeQosPolicyRoundRobinDefault)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["default"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyRoundRobin) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafQosPolicyRoundRobinDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyRoundRobinDescrIPtion = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["class"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagQosPolicyRoundRobinClass = &map[string]QosPolicyRoundRobinClass{}

		err = json.Unmarshal(subJSONStr, o.TagQosPolicyRoundRobinClass)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["default"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeQosPolicyRoundRobinDefault = &QosPolicyRoundRobinDefault{}

		err = json.Unmarshal(subJSONStr, o.NodeQosPolicyRoundRobinDefault)
		if err != nil {
			return err
		}
	}

	return nil
}
