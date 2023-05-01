// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborBfd describes the resource data model.
type ProtocolsBgpNeighborBfd struct {
	// LeafNodes
	LeafProtocolsBgpNeighborBfdProfile                  types.String `tfsdk:"profile" json:"profile,omitempty"`
	LeafProtocolsBgpNeighborBfdCheckControlPlaneFailure types.String `tfsdk:"check_control_plane_failure" json:"check-control-plane-failure,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborBfd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

|  Format  |  Description  |
|----------|---------------|
|  txt  |  BFD profile name  |

`,
		},

		"check_control_plane_failure": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Allow to write CBIT independence in BFD outgoing packets and read both C-BIT value of BFD and lookup BGP peer status

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborBfd) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborBfdProfile.IsNull() && !o.LeafProtocolsBgpNeighborBfdProfile.IsUnknown() {
		jsonData["profile"] = o.LeafProtocolsBgpNeighborBfdProfile.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborBfdCheckControlPlaneFailure.IsNull() && !o.LeafProtocolsBgpNeighborBfdCheckControlPlaneFailure.IsUnknown() {
		jsonData["check-control-plane-failure"] = o.LeafProtocolsBgpNeighborBfdCheckControlPlaneFailure.ValueString()
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
func (o *ProtocolsBgpNeighborBfd) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["profile"]; ok {
		o.LeafProtocolsBgpNeighborBfdProfile = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborBfdProfile = basetypes.NewStringNull()
	}

	if value, ok := jsonData["check-control-plane-failure"]; ok {
		o.LeafProtocolsBgpNeighborBfdCheckControlPlaneFailure = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborBfdCheckControlPlaneFailure = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
