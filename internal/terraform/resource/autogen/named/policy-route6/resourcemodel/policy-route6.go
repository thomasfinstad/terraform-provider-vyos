// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRoutesix describes the resource data model.
type PolicyRoutesix struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafPolicyRoutesixDescrIPtion      types.String `tfsdk:"description" json:"description,omitempty"`
	LeafPolicyRoutesixInterface        types.String `tfsdk:"interface" json:"interface,omitempty"`
	LeafPolicyRoutesixEnableDefaultLog types.String `tfsdk:"enable_default_log" json:"enable-default-log,omitempty"`

	// TagNodes
	TagPolicyRoutesixRule *map[string]PolicyRoutesixRule `tfsdk:"rule" json:"rule,omitempty"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyRoutesix) GetVyosPath() []string {
	return []string{
		"policy",
		"route6",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Policy route rule set name for IPv6

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

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		"enable_default_log": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting default-action

`,
		},

		// TagNodes

		"rule": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: PolicyRoutesixRule{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Policy rule number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number of policy rule  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRoutesix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRoutesixDescrIPtion.IsNull() && !o.LeafPolicyRoutesixDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyRoutesixDescrIPtion.ValueString()
	}

	if !o.LeafPolicyRoutesixInterface.IsNull() && !o.LeafPolicyRoutesixInterface.IsUnknown() {
		jsonData["interface"] = o.LeafPolicyRoutesixInterface.ValueString()
	}

	if !o.LeafPolicyRoutesixEnableDefaultLog.IsNull() && !o.LeafPolicyRoutesixEnableDefaultLog.IsUnknown() {
		jsonData["enable-default-log"] = o.LeafPolicyRoutesixEnableDefaultLog.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagPolicyRoutesixRule).IsZero() {
		subJSONStr, err := json.Marshal(o.TagPolicyRoutesixRule)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["rule"] = subData
	}

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRoutesix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyRoutesixDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["interface"]; ok {
		o.LeafPolicyRoutesixInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-default-log"]; ok {
		o.LeafPolicyRoutesixEnableDefaultLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixEnableDefaultLog = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["rule"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagPolicyRoutesixRule = &map[string]PolicyRoutesixRule{}

		err = json.Unmarshal(subJSONStr, o.TagPolicyRoutesixRule)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
