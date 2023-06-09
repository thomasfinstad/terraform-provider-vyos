// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpListen describes the resource data model.
type VrfNameProtocolsBgpListen struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpListenLimit types.String `tfsdk:"limit" json:"limit,omitempty"`

	// TagNodes
	TagVrfNameProtocolsBgpListenRange *map[string]VrfNameProtocolsBgpListenRange `tfsdk:"range" json:"range,omitempty"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpListen) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"limit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of dynamic neighbors that can be created

|  Format  |  Description  |
|----------|---------------|
|  u32:1-5000  |  BGP neighbor limit  |

`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpListenRange{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `BGP dynamic neighbors listen range

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 dynamic neighbors listen range  |
|  ipv6net  |  IPv6 dynamic neighbors listen range  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpListen) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpListenLimit.IsNull() && !o.LeafVrfNameProtocolsBgpListenLimit.IsUnknown() {
		jsonData["limit"] = o.LeafVrfNameProtocolsBgpListenLimit.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagVrfNameProtocolsBgpListenRange).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsBgpListenRange)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["range"] = subData
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
func (o *VrfNameProtocolsBgpListen) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["limit"]; ok {
		o.LeafVrfNameProtocolsBgpListenLimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpListenLimit = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["range"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsBgpListenRange = &map[string]VrfNameProtocolsBgpListenRange{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsBgpListenRange)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
