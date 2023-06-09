// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallName describes the resource data model.
type FirewallName struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafFirewallNameDefaultAction     types.String `tfsdk:"default_action" json:"default-action,omitempty"`
	LeafFirewallNameEnableDefaultLog  types.String `tfsdk:"enable_default_log" json:"enable-default-log,omitempty"`
	LeafFirewallNameDescrIPtion       types.String `tfsdk:"description" json:"description,omitempty"`
	LeafFirewallNameDefaultJumpTarget types.String `tfsdk:"default_jump_target" json:"default-jump-target,omitempty"`

	// TagNodes
	TagFirewallNameRule *map[string]FirewallNameRule `tfsdk:"rule" json:"rule,omitempty"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallName) GetVyosPath() []string {
	return []string{
		"firewall",
		"name",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv4 firewall rule-set name

`,
		},

		// LeafNodes

		"default_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default-action for rule-set

|  Format  |  Description  |
|----------|---------------|
|  drop  |  Drop if no prior rules are hit  |
|  jump  |  Jump to another chain if no prior rules are hit  |
|  reject  |  Drop and notify source if no prior rules are hit  |
|  return  |  Return from the current chain and continue at the next rule of the last                   chain  |
|  accept  |  Accept if no prior rules are hit  |

`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"enable_default_log": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting default-action

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"default_jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined in default-action to use this
                setting

`,
		},

		// TagNodes

		"rule": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: FirewallNameRule{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Firewall rule number (IPv4)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number for this Firewall rule  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallName) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallNameDefaultAction.IsNull() && !o.LeafFirewallNameDefaultAction.IsUnknown() {
		jsonData["default-action"] = o.LeafFirewallNameDefaultAction.ValueString()
	}

	if !o.LeafFirewallNameEnableDefaultLog.IsNull() && !o.LeafFirewallNameEnableDefaultLog.IsUnknown() {
		jsonData["enable-default-log"] = o.LeafFirewallNameEnableDefaultLog.ValueString()
	}

	if !o.LeafFirewallNameDescrIPtion.IsNull() && !o.LeafFirewallNameDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafFirewallNameDescrIPtion.ValueString()
	}

	if !o.LeafFirewallNameDefaultJumpTarget.IsNull() && !o.LeafFirewallNameDefaultJumpTarget.IsUnknown() {
		jsonData["default-jump-target"] = o.LeafFirewallNameDefaultJumpTarget.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagFirewallNameRule).IsZero() {
		subJSONStr, err := json.Marshal(o.TagFirewallNameRule)
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
func (o *FirewallName) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["default-action"]; ok {
		o.LeafFirewallNameDefaultAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameDefaultAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-default-log"]; ok {
		o.LeafFirewallNameEnableDefaultLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameEnableDefaultLog = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafFirewallNameDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-jump-target"]; ok {
		o.LeafFirewallNameDefaultJumpTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameDefaultJumpTarget = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["rule"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagFirewallNameRule = &map[string]FirewallNameRule{}

		err = json.Unmarshal(subJSONStr, o.TagFirewallNameRule)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
