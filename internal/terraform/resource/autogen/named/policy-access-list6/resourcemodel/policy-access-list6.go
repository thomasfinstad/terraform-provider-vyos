// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyAccessListsix describes the resource data model.
type PolicyAccessListsix struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafPolicyAccessListsixDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes
	TagPolicyAccessListsixRule *map[string]PolicyAccessListsixRule `tfsdk:"rule" json:"rule,omitempty"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAccessListsix) GetVyosPath() []string {
	return []string{
		"policy",
		"access-list6",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 access-list filter

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 access-list  |

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

		"rule": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: PolicyAccessListsixRule{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Rule for this access-list6

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list6 rule number  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyAccessListsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyAccessListsixDescrIPtion.IsNull() && !o.LeafPolicyAccessListsixDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyAccessListsixDescrIPtion.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagPolicyAccessListsixRule).IsZero() {
		subJSONStr, err := json.Marshal(o.TagPolicyAccessListsixRule)
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
func (o *PolicyAccessListsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyAccessListsixDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListsixDescrIPtion = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["rule"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagPolicyAccessListsixRule = &map[string]PolicyAccessListsixRule{}

		err = json.Unmarshal(subJSONStr, o.TagPolicyAccessListsixRule)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
