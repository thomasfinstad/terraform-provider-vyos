// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ContainerNameVolume describes the resource data model.
type ContainerNameVolume struct {
	// LeafNodes
	LeafContainerNameVolumeSource      types.String `tfsdk:"source" json:"source,omitempty"`
	LeafContainerNameVolumeDestination types.String `tfsdk:"destination" json:"destination,omitempty"`
	LeafContainerNameVolumeMode        types.String `tfsdk:"mode" json:"mode,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNameVolume) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source host directory

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Source host directory  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination container directory

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination container directory  |

`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Volume access mode ro/rw

|  Format  |  Description  |
|----------|---------------|
|  ro  |  Volume mounted into the container as read-only  |
|  rw  |  Volume mounted into the container as read-write  |

`,

			// Default:          stringdefault.StaticString(`rw`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerNameVolume) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafContainerNameVolumeSource.IsNull() && !o.LeafContainerNameVolumeSource.IsUnknown() {
		jsonData["source"] = o.LeafContainerNameVolumeSource.ValueString()
	}

	if !o.LeafContainerNameVolumeDestination.IsNull() && !o.LeafContainerNameVolumeDestination.IsUnknown() {
		jsonData["destination"] = o.LeafContainerNameVolumeDestination.ValueString()
	}

	if !o.LeafContainerNameVolumeMode.IsNull() && !o.LeafContainerNameVolumeMode.IsUnknown() {
		jsonData["mode"] = o.LeafContainerNameVolumeMode.ValueString()
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
func (o *ContainerNameVolume) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["source"]; ok {
		o.LeafContainerNameVolumeSource = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameVolumeSource = basetypes.NewStringNull()
	}

	if value, ok := jsonData["destination"]; ok {
		o.LeafContainerNameVolumeDestination = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameVolumeDestination = basetypes.NewStringNull()
	}

	if value, ok := jsonData["mode"]; ok {
		o.LeafContainerNameVolumeMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameVolumeMode = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
