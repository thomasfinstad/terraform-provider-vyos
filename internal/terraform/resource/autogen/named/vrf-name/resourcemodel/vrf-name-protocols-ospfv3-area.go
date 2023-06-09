// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsOspfvthreeArea describes the resource data model.
type VrfNameProtocolsOspfvthreeArea struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeAreaExportList types.String `tfsdk:"export_list" json:"export-list,omitempty"`
	LeafVrfNameProtocolsOspfvthreeAreaImportList types.String `tfsdk:"import_list" json:"import-list,omitempty"`

	// TagNodes
	TagVrfNameProtocolsOspfvthreeAreaRange *map[string]VrfNameProtocolsOspfvthreeAreaRange `tfsdk:"range" json:"range,omitempty"`

	// Nodes
	NodeVrfNameProtocolsOspfvthreeAreaAreaType *VrfNameProtocolsOspfvthreeAreaAreaType `tfsdk:"area_type" json:"area-type,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeArea) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of export-list

`,
		},

		"import_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of import-list

`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsOspfvthreeAreaRange{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Specify IPv6 prefix (border routers only)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Specify IPv6 prefix (border routers only)  |

`,
		},

		// Nodes

		"area_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAreaAreaType{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `OSPFv3 Area type

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfvthreeArea) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsOspfvthreeAreaExportList.IsNull() && !o.LeafVrfNameProtocolsOspfvthreeAreaExportList.IsUnknown() {
		jsonData["export-list"] = o.LeafVrfNameProtocolsOspfvthreeAreaExportList.ValueString()
	}

	if !o.LeafVrfNameProtocolsOspfvthreeAreaImportList.IsNull() && !o.LeafVrfNameProtocolsOspfvthreeAreaImportList.IsUnknown() {
		jsonData["import-list"] = o.LeafVrfNameProtocolsOspfvthreeAreaImportList.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagVrfNameProtocolsOspfvthreeAreaRange).IsZero() {
		subJSONStr, err := json.Marshal(o.TagVrfNameProtocolsOspfvthreeAreaRange)
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

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeAreaAreaType).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeAreaAreaType)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["area-type"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfvthreeArea) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export-list"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeAreaExportList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeAreaExportList = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import-list"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeAreaImportList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeAreaImportList = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["range"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagVrfNameProtocolsOspfvthreeAreaRange = &map[string]VrfNameProtocolsOspfvthreeAreaRange{}

		err = json.Unmarshal(subJSONStr, o.TagVrfNameProtocolsOspfvthreeAreaRange)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["area-type"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeAreaAreaType = &VrfNameProtocolsOspfvthreeAreaAreaType{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeAreaAreaType)
		if err != nil {
			return err
		}
	}

	return nil
}
