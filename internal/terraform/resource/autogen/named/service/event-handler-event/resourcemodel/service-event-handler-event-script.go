// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceEventHandlerEventScrIPt describes the resource data model.
type ServiceEventHandlerEventScrIPt struct {
	// LeafNodes
	LeafServiceEventHandlerEventScrIPtArguments types.String `tfsdk:"arguments" vyos:"arguments,omitempty"`
	LeafServiceEventHandlerEventScrIPtPath      types.String `tfsdk:"path" vyos:"path,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceEventHandlerEventScrIPtEnvironment bool `tfsdk:"environment" vyos:"environment,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceEventHandlerEventScrIPt) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"arguments": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script arguments

`,
		},

		"path": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Path to the script

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceEventHandlerEventScrIPt) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceEventHandlerEventScrIPtArguments.IsNull() && !o.LeafServiceEventHandlerEventScrIPtArguments.IsUnknown() {
		jsonData["arguments"] = o.LeafServiceEventHandlerEventScrIPtArguments.ValueString()
	}

	if !o.LeafServiceEventHandlerEventScrIPtPath.IsNull() && !o.LeafServiceEventHandlerEventScrIPtPath.IsUnknown() {
		jsonData["path"] = o.LeafServiceEventHandlerEventScrIPtPath.ValueString()
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
func (o *ServiceEventHandlerEventScrIPt) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["arguments"]; ok {
		o.LeafServiceEventHandlerEventScrIPtArguments = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceEventHandlerEventScrIPtArguments = basetypes.NewStringNull()
	}

	if value, ok := jsonData["path"]; ok {
		o.LeafServiceEventHandlerEventScrIPtPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceEventHandlerEventScrIPtPath = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
