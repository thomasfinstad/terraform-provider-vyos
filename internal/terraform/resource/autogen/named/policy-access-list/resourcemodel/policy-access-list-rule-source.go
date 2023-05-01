// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyAccessListRuleSource describes the resource data model.
type PolicyAccessListRuleSource struct {
	// LeafNodes
	LeafPolicyAccessListRuleSourceAny         types.String `tfsdk:"any" json:"any,omitempty"`
	LeafPolicyAccessListRuleSourceHost        types.String `tfsdk:"host" json:"host,omitempty"`
	LeafPolicyAccessListRuleSourceInverseMask types.String `tfsdk:"inverse_mask" json:"inverse-mask,omitempty"`
	LeafPolicyAccessListRuleSourceNetwork     types.String `tfsdk:"network" json:"network,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"any": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Any IP address to match

`,
		},

		"host": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Single host IP address to match

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Host address to match  |

`,
		},

		"inverse_mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match (requires network be defined)

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Inverse-mask to match  |

`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match (requires inverse-mask be defined)

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Inverse-mask to match  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyAccessListRuleSource) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyAccessListRuleSourceAny.IsNull() && !o.LeafPolicyAccessListRuleSourceAny.IsUnknown() {
		jsonData["any"] = o.LeafPolicyAccessListRuleSourceAny.ValueString()
	}

	if !o.LeafPolicyAccessListRuleSourceHost.IsNull() && !o.LeafPolicyAccessListRuleSourceHost.IsUnknown() {
		jsonData["host"] = o.LeafPolicyAccessListRuleSourceHost.ValueString()
	}

	if !o.LeafPolicyAccessListRuleSourceInverseMask.IsNull() && !o.LeafPolicyAccessListRuleSourceInverseMask.IsUnknown() {
		jsonData["inverse-mask"] = o.LeafPolicyAccessListRuleSourceInverseMask.ValueString()
	}

	if !o.LeafPolicyAccessListRuleSourceNetwork.IsNull() && !o.LeafPolicyAccessListRuleSourceNetwork.IsUnknown() {
		jsonData["network"] = o.LeafPolicyAccessListRuleSourceNetwork.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyAccessListRuleSource) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["any"]; ok {
		o.LeafPolicyAccessListRuleSourceAny = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListRuleSourceAny = basetypes.NewStringNull()
	}

	if value, ok := jsonData["host"]; ok {
		o.LeafPolicyAccessListRuleSourceHost = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListRuleSourceHost = basetypes.NewStringNull()
	}

	if value, ok := jsonData["inverse-mask"]; ok {
		o.LeafPolicyAccessListRuleSourceInverseMask = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListRuleSourceInverseMask = basetypes.NewStringNull()
	}

	if value, ok := jsonData["network"]; ok {
		o.LeafPolicyAccessListRuleSourceNetwork = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListRuleSourceNetwork = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
