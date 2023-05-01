// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal types.String `tfsdk:"external" json:"external,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal types.String `tfsdk:"internal" json:"internal,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal    types.String `tfsdk:"local" json:"local,omitempty"`

	// TagNodes
	TagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix *map[string]VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix `tfsdk:"prefix" json:"prefix,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `eBGP routes administrative distance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  eBGP routes administrative distance  |

`,
		},

		"internal": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `iBGP routes administrative distance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  iBGP routes administrative distance  |

`,
		},

		"local": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Locally originated BGP routes administrative distance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Locally originated BGP routes administrative distance  |

`,
		},

		// TagNodes

		"prefix": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Administrative distance for a specific BGP prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Administrative distance for a specific BGP prefix  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal.IsNull() && !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal.IsUnknown() {
		jsonData["external"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal.IsNull() && !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal.IsUnknown() {
		jsonData["internal"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal.IsNull() && !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal.IsUnknown() {
		jsonData["local"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["prefix"] = subData
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
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["external"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal = basetypes.NewStringNull()
	}

	if value, ok := jsonData["internal"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["prefix"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix = &map[string]VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
