// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfName describes the resource data model.
type VrfName struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafVrfNameDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`
	LeafVrfNameDisable     types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafVrfNameTable       types.String `tfsdk:"table" json:"table,omitempty"`
	LeafVrfNameVni         types.String `tfsdk:"vni" json:"vni,omitempty"`

	// TagNodes

	// Nodes
	NodeVrfNameIP        *VrfNameIP        `tfsdk:"ip" json:"ip,omitempty"`
	NodeVrfNameIPvsix    *VrfNameIPvsix    `tfsdk:"ipv6" json:"ipv6,omitempty"`
	NodeVrfNameProtocols *VrfNameProtocols `tfsdk:"protocols" json:"protocols,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfName) GetVyosPath() []string {
	return []string{
		"vrf",
		"name",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"table": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Routing table associated with this instance

|  Format  |  Description  |
|----------|---------------|
|  u32:100-65535  |  Routing table ID  |

`,
		},

		"vni": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Network Identifier

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777214  |  VXLAN virtual network identifier  |

`,
		},

		// TagNodes

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: VrfNameIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: VrfNameIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"protocols": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocols{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Routing protocol parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfName) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameDescrIPtion.IsNull() && !o.LeafVrfNameDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafVrfNameDescrIPtion.ValueString()
	}

	if !o.LeafVrfNameDisable.IsNull() && !o.LeafVrfNameDisable.IsUnknown() {
		jsonData["disable"] = o.LeafVrfNameDisable.ValueString()
	}

	if !o.LeafVrfNameTable.IsNull() && !o.LeafVrfNameTable.IsUnknown() {
		jsonData["table"] = o.LeafVrfNameTable.ValueString()
	}

	if !o.LeafVrfNameVni.IsNull() && !o.LeafVrfNameVni.IsUnknown() {
		jsonData["vni"] = o.LeafVrfNameVni.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameIP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameIP)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ip"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameIPvsix).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameIPvsix)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipv6"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocols).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocols)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["protocols"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfName) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafVrfNameDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafVrfNameDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["table"]; ok {
		o.LeafVrfNameTable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameTable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vni"]; ok {
		o.LeafVrfNameVni = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameVni = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["ip"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameIP = &VrfNameIP{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameIP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipv6"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameIPvsix = &VrfNameIPvsix{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameIPvsix)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["protocols"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocols = &VrfNameProtocols{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocols)
		if err != nil {
			return err
		}
	}

	return nil
}
