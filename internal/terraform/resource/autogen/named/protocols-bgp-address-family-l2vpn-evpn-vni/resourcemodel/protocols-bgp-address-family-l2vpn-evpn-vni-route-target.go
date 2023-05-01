// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget describes the resource data model.
type ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget struct {
	// LeafNodes
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetBoth   types.String `tfsdk:"both" json:"both,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetImport types.String `tfsdk:"import" json:"import,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetExport types.String `tfsdk:"export" json:"export,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"both": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target both import and export

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target import

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |

`,
		},

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target export

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetBoth.IsNull() && !o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetBoth.IsUnknown() {
		jsonData["both"] = o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetBoth.ValueString()
	}

	if !o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetImport.IsNull() && !o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetImport.IsUnknown() {
		jsonData["import"] = o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetImport.ValueString()
	}

	if !o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetExport.IsNull() && !o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetExport.IsUnknown() {
		jsonData["export"] = o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetExport.ValueString()
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
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["both"]; ok {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetBoth = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetBoth = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetImport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["export"]; ok {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetExport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
