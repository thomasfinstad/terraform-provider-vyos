// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemConntrackIgnoreRule describes the resource data model.
type SystemConntrackIgnoreRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafSystemConntrackIgnoreRuleDescrIPtion      types.String `tfsdk:"description" json:"description,omitempty"`
	LeafSystemConntrackIgnoreRuleInboundInterface types.String `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`
	LeafSystemConntrackIgnoreRuleProtocol         types.String `tfsdk:"protocol" json:"protocol,omitempty"`

	// TagNodes

	// Nodes
	NodeSystemConntrackIgnoreRuleDestination *SystemConntrackIgnoreRuleDestination `tfsdk:"destination" json:"destination,omitempty"`
	NodeSystemConntrackIgnoreRuleSource      *SystemConntrackIgnoreRuleSource      `tfsdk:"source" json:"source,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemConntrackIgnoreRule) GetVyosPath() []string {
	return []string{
		"system",
		"conntrack",
		"ignore",
		"rule",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackIgnoreRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number of conntrack ignore rule  |

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

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to ignore connections tracking on

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Protocol name  |

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: SystemConntrackIgnoreRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: SystemConntrackIgnoreRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemConntrackIgnoreRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafSystemConntrackIgnoreRuleDescrIPtion.IsNull() && !o.LeafSystemConntrackIgnoreRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafSystemConntrackIgnoreRuleDescrIPtion.ValueString()
	}

	if !o.LeafSystemConntrackIgnoreRuleInboundInterface.IsNull() && !o.LeafSystemConntrackIgnoreRuleInboundInterface.IsUnknown() {
		jsonData["inbound-interface"] = o.LeafSystemConntrackIgnoreRuleInboundInterface.ValueString()
	}

	if !o.LeafSystemConntrackIgnoreRuleProtocol.IsNull() && !o.LeafSystemConntrackIgnoreRuleProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafSystemConntrackIgnoreRuleProtocol.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeSystemConntrackIgnoreRuleDestination).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeSystemConntrackIgnoreRuleDestination)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["destination"] = subData
	}

	if !reflect.ValueOf(o.NodeSystemConntrackIgnoreRuleSource).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeSystemConntrackIgnoreRuleSource)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["source"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemConntrackIgnoreRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafSystemConntrackIgnoreRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemConntrackIgnoreRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["inbound-interface"]; ok {
		o.LeafSystemConntrackIgnoreRuleInboundInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemConntrackIgnoreRuleInboundInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafSystemConntrackIgnoreRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemConntrackIgnoreRuleProtocol = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["destination"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeSystemConntrackIgnoreRuleDestination = &SystemConntrackIgnoreRuleDestination{}

		err = json.Unmarshal(subJSONStr, o.NodeSystemConntrackIgnoreRuleDestination)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["source"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeSystemConntrackIgnoreRuleSource = &SystemConntrackIgnoreRuleSource{}

		err = json.Unmarshal(subJSONStr, o.NodeSystemConntrackIgnoreRuleSource)
		if err != nil {
			return err
		}
	}

	return nil
}
