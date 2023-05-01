// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal types.String `tfsdk:"external" json:"external,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal types.String `tfsdk:"internal" json:"internal,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal    types.String `tfsdk:"local" json:"local,omitempty"`

	// TagNodes
	TagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix *map[string]VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix `tfsdk:"prefix" json:"prefix,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix{}.ResourceSchemaAttributes(),
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
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal.IsNull() && !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal.IsUnknown() {
		jsonData["external"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal.IsNull() && !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal.IsUnknown() {
		jsonData["internal"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal.IsNull() && !o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal.IsUnknown() {
		jsonData["local"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix)
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
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["external"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal = basetypes.NewStringNull()
	}

	if value, ok := jsonData["internal"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal = basetypes.NewStringNull()
	}

	if value, ok := jsonData["local"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["prefix"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix = &map[string]VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
