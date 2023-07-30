// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyAccessList describes the resource data model.
type PolicyAccessList struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafPolicyAccessListDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagPolicyAccessListRule bool `tfsdk:"rule" vyos:"rule,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAccessList) GetVyosPath() []string {
	return []string{
		"policy",
		"access-list",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IP access-list filter

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-99  |  IP standard access list  |
    |  u32:100-199  |  IP extended access list  |
    |  u32:1300-1999  |  IP standard access list (expanded range)  |
    |  u32:2000-2699  |  IP extended access list (expanded range)  |

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

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyAccessList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyAccessListDescrIPtion.IsNull() && !o.LeafPolicyAccessListDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyAccessListDescrIPtion.ValueString()
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
func (o *PolicyAccessList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyAccessListDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListDescrIPtion = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
